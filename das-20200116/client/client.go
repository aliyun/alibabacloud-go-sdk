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

type AddHDMInstanceRequest struct {
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	FlushAccount  *string `json:"FlushAccount,omitempty" xml:"FlushAccount,omitempty"`
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceArea  *string `json:"InstanceArea,omitempty" xml:"InstanceArea,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NetworkType   *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Context       *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s AddHDMInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHDMInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceRequest) SetEngine(v string) *AddHDMInstanceRequest {
	s.Engine = &v
	return s
}

func (s *AddHDMInstanceRequest) SetFlushAccount(v string) *AddHDMInstanceRequest {
	s.FlushAccount = &v
	return s
}

func (s *AddHDMInstanceRequest) SetInstanceAlias(v string) *AddHDMInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *AddHDMInstanceRequest) SetInstanceArea(v string) *AddHDMInstanceRequest {
	s.InstanceArea = &v
	return s
}

func (s *AddHDMInstanceRequest) SetInstanceId(v string) *AddHDMInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHDMInstanceRequest) SetIp(v string) *AddHDMInstanceRequest {
	s.Ip = &v
	return s
}

func (s *AddHDMInstanceRequest) SetNetworkType(v string) *AddHDMInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *AddHDMInstanceRequest) SetPassword(v string) *AddHDMInstanceRequest {
	s.Password = &v
	return s
}

func (s *AddHDMInstanceRequest) SetPort(v string) *AddHDMInstanceRequest {
	s.Port = &v
	return s
}

func (s *AddHDMInstanceRequest) SetRegion(v string) *AddHDMInstanceRequest {
	s.Region = &v
	return s
}

func (s *AddHDMInstanceRequest) SetUsername(v string) *AddHDMInstanceRequest {
	s.Username = &v
	return s
}

func (s *AddHDMInstanceRequest) SetVpcId(v string) *AddHDMInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *AddHDMInstanceRequest) SetContext(v string) *AddHDMInstanceRequest {
	s.Context = &v
	return s
}

type AddHDMInstanceResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AddHDMInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                         `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s AddHDMInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHDMInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceResponseBody) SetCode(v string) *AddHDMInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetData(v *AddHDMInstanceResponseBodyData) *AddHDMInstanceResponseBody {
	s.Data = v
	return s
}

func (s *AddHDMInstanceResponseBody) SetMessage(v string) *AddHDMInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetRequestId(v string) *AddHDMInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetSuccess(v string) *AddHDMInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetSynchro(v string) *AddHDMInstanceResponseBody {
	s.Synchro = &v
	return s
}

type AddHDMInstanceResponseBodyData struct {
	CallerUid  *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	Code       *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddHDMInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddHDMInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceResponseBodyData) SetCallerUid(v string) *AddHDMInstanceResponseBodyData {
	s.CallerUid = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetCode(v int32) *AddHDMInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetError(v string) *AddHDMInstanceResponseBodyData {
	s.Error = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetInstanceId(v string) *AddHDMInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetIp(v string) *AddHDMInstanceResponseBodyData {
	s.Ip = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetOwnerId(v string) *AddHDMInstanceResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetPort(v int32) *AddHDMInstanceResponseBodyData {
	s.Port = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetRole(v string) *AddHDMInstanceResponseBodyData {
	s.Role = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetTenantId(v string) *AddHDMInstanceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetToken(v string) *AddHDMInstanceResponseBodyData {
	s.Token = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetUuid(v string) *AddHDMInstanceResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetVpcId(v string) *AddHDMInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

type AddHDMInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHDMInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHDMInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHDMInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceResponse) SetHeaders(v map[string]*string) *AddHDMInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddHDMInstanceResponse) SetBody(v *AddHDMInstanceResponseBody) *AddHDMInstanceResponse {
	s.Body = v
	return s
}

type CreateAdamBenchTaskRequest struct {
	Description      *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	DstInstanceId    *string  `json:"DstInstanceId,omitempty" xml:"DstInstanceId,omitempty"`
	DstSuperAccount  *string  `json:"DstSuperAccount,omitempty" xml:"DstSuperAccount,omitempty"`
	DstSuperPassword *string  `json:"DstSuperPassword,omitempty" xml:"DstSuperPassword,omitempty"`
	Rate             *int32   `json:"Rate,omitempty" xml:"Rate,omitempty"`
	RequestDuration  *int64   `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	RequestStartTime *int64   `json:"RequestStartTime,omitempty" xml:"RequestStartTime,omitempty"`
	SrcEngine        *string  `json:"SrcEngine,omitempty" xml:"SrcEngine,omitempty"`
	SrcEngineVersion *string  `json:"SrcEngineVersion,omitempty" xml:"SrcEngineVersion,omitempty"`
	SrcMaxQps        *float64 `json:"SrcMaxQps,omitempty" xml:"SrcMaxQps,omitempty"`
	SrcMeanQps       *float64 `json:"SrcMeanQps,omitempty" xml:"SrcMeanQps,omitempty"`
	SrcSqlOssAddr    *string  `json:"SrcSqlOssAddr,omitempty" xml:"SrcSqlOssAddr,omitempty"`
}

func (s CreateAdamBenchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAdamBenchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAdamBenchTaskRequest) SetDescription(v string) *CreateAdamBenchTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetDstInstanceId(v string) *CreateAdamBenchTaskRequest {
	s.DstInstanceId = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetDstSuperAccount(v string) *CreateAdamBenchTaskRequest {
	s.DstSuperAccount = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetDstSuperPassword(v string) *CreateAdamBenchTaskRequest {
	s.DstSuperPassword = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetRate(v int32) *CreateAdamBenchTaskRequest {
	s.Rate = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetRequestDuration(v int64) *CreateAdamBenchTaskRequest {
	s.RequestDuration = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetRequestStartTime(v int64) *CreateAdamBenchTaskRequest {
	s.RequestStartTime = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetSrcEngine(v string) *CreateAdamBenchTaskRequest {
	s.SrcEngine = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetSrcEngineVersion(v string) *CreateAdamBenchTaskRequest {
	s.SrcEngineVersion = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetSrcMaxQps(v float64) *CreateAdamBenchTaskRequest {
	s.SrcMaxQps = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetSrcMeanQps(v float64) *CreateAdamBenchTaskRequest {
	s.SrcMeanQps = &v
	return s
}

func (s *CreateAdamBenchTaskRequest) SetSrcSqlOssAddr(v string) *CreateAdamBenchTaskRequest {
	s.SrcSqlOssAddr = &v
	return s
}

type CreateAdamBenchTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAdamBenchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAdamBenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdamBenchTaskResponseBody) SetCode(v string) *CreateAdamBenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAdamBenchTaskResponseBody) SetData(v string) *CreateAdamBenchTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAdamBenchTaskResponseBody) SetMessage(v string) *CreateAdamBenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAdamBenchTaskResponseBody) SetRequestId(v string) *CreateAdamBenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdamBenchTaskResponseBody) SetSuccess(v string) *CreateAdamBenchTaskResponseBody {
	s.Success = &v
	return s
}

type CreateAdamBenchTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAdamBenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAdamBenchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAdamBenchTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAdamBenchTaskResponse) SetHeaders(v map[string]*string) *CreateAdamBenchTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAdamBenchTaskResponse) SetBody(v *CreateAdamBenchTaskResponseBody) *CreateAdamBenchTaskResponse {
	s.Body = v
	return s
}

type CreateCacheAnalysisJobRequest struct {
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId      *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateCacheAnalysisJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobRequest) SetBackupSetId(v string) *CreateCacheAnalysisJobRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateCacheAnalysisJobRequest) SetInstanceId(v string) *CreateCacheAnalysisJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCacheAnalysisJobRequest) SetNodeId(v string) *CreateCacheAnalysisJobRequest {
	s.NodeId = &v
	return s
}

type CreateCacheAnalysisJobResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateCacheAnalysisJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCacheAnalysisJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBody) SetCode(v string) *CreateCacheAnalysisJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetData(v *CreateCacheAnalysisJobResponseBodyData) *CreateCacheAnalysisJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetMessage(v string) *CreateCacheAnalysisJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetRequestId(v string) *CreateCacheAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetSuccess(v string) *CreateCacheAnalysisJobResponseBody {
	s.Success = &v
	return s
}

type CreateCacheAnalysisJobResponseBodyData struct {
	BigKeys    *CreateCacheAnalysisJobResponseBodyDataBigKeys `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	InstanceId *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string                                        `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeId     *string                                        `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TaskState  *string                                        `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
}

func (s CreateCacheAnalysisJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetBigKeys(v *CreateCacheAnalysisJobResponseBodyDataBigKeys) *CreateCacheAnalysisJobResponseBodyData {
	s.BigKeys = v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetInstanceId(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetJobId(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetMessage(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetNodeId(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetTaskState(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.TaskState = &v
	return s
}

type CreateCacheAnalysisJobResponseBodyDataBigKeys struct {
	KeyInfo []*CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeys) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeys) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeys) SetKeyInfo(v []*CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) *CreateCacheAnalysisJobResponseBodyDataBigKeys {
	s.KeyInfo = v
	return s
}

type CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo struct {
	Bytes                *int64  `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	Count                *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Db                   *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Encoding             *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	ExpirationTimeMillis *int64  `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetBytes(v int64) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Bytes = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetCount(v int64) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Count = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetDb(v int32) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Db = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetEncoding(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Encoding = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetExpirationTimeMillis(v int64) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetKey(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Key = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetNodeId(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.NodeId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetType(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Type = &v
	return s
}

type CreateCacheAnalysisJobResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCacheAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCacheAnalysisJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponse) SetHeaders(v map[string]*string) *CreateCacheAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheAnalysisJobResponse) SetBody(v *CreateCacheAnalysisJobResponseBody) *CreateCacheAnalysisJobResponse {
	s.Body = v
	return s
}

type CreateCloudBenchTasksRequest struct {
	Amount              *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupTime          *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	ClientType          *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DstConnectionString *string `json:"DstConnectionString,omitempty" xml:"DstConnectionString,omitempty"`
	DstInstanceId       *string `json:"DstInstanceId,omitempty" xml:"DstInstanceId,omitempty"`
	DstPort             *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	DstSuperAccount     *string `json:"DstSuperAccount,omitempty" xml:"DstSuperAccount,omitempty"`
	DstSuperPassword    *string `json:"DstSuperPassword,omitempty" xml:"DstSuperPassword,omitempty"`
	DstType             *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	DtsJobClass         *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	DtsJobId            *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	EndState            *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	GatewayVpcId        *string `json:"GatewayVpcId,omitempty" xml:"GatewayVpcId,omitempty"`
	GatewayVpcIp        *string `json:"GatewayVpcIp,omitempty" xml:"GatewayVpcIp,omitempty"`
	Rate                *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	RequestDuration     *string `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	RequestEndTime      *string `json:"RequestEndTime,omitempty" xml:"RequestEndTime,omitempty"`
	RequestStartTime    *string `json:"RequestStartTime,omitempty" xml:"RequestStartTime,omitempty"`
	SmartPressureTime   *string `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	SrcInstanceId       *string `json:"SrcInstanceId,omitempty" xml:"SrcInstanceId,omitempty"`
	SrcPublicIp         *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	SrcSuperAccount     *string `json:"SrcSuperAccount,omitempty" xml:"SrcSuperAccount,omitempty"`
	SrcSuperPassword    *string `json:"SrcSuperPassword,omitempty" xml:"SrcSuperPassword,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	WorkDir             *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s CreateCloudBenchTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudBenchTasksRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksRequest) SetAmount(v string) *CreateCloudBenchTasksRequest {
	s.Amount = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetBackupId(v string) *CreateCloudBenchTasksRequest {
	s.BackupId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetBackupTime(v string) *CreateCloudBenchTasksRequest {
	s.BackupTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetClientType(v string) *CreateCloudBenchTasksRequest {
	s.ClientType = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDescription(v string) *CreateCloudBenchTasksRequest {
	s.Description = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstConnectionString(v string) *CreateCloudBenchTasksRequest {
	s.DstConnectionString = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstInstanceId(v string) *CreateCloudBenchTasksRequest {
	s.DstInstanceId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstPort(v string) *CreateCloudBenchTasksRequest {
	s.DstPort = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstSuperAccount(v string) *CreateCloudBenchTasksRequest {
	s.DstSuperAccount = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstSuperPassword(v string) *CreateCloudBenchTasksRequest {
	s.DstSuperPassword = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstType(v string) *CreateCloudBenchTasksRequest {
	s.DstType = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDtsJobClass(v string) *CreateCloudBenchTasksRequest {
	s.DtsJobClass = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDtsJobId(v string) *CreateCloudBenchTasksRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetEndState(v string) *CreateCloudBenchTasksRequest {
	s.EndState = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetGatewayVpcId(v string) *CreateCloudBenchTasksRequest {
	s.GatewayVpcId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetGatewayVpcIp(v string) *CreateCloudBenchTasksRequest {
	s.GatewayVpcIp = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRate(v string) *CreateCloudBenchTasksRequest {
	s.Rate = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRequestDuration(v string) *CreateCloudBenchTasksRequest {
	s.RequestDuration = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRequestEndTime(v string) *CreateCloudBenchTasksRequest {
	s.RequestEndTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRequestStartTime(v string) *CreateCloudBenchTasksRequest {
	s.RequestStartTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSmartPressureTime(v string) *CreateCloudBenchTasksRequest {
	s.SmartPressureTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcInstanceId(v string) *CreateCloudBenchTasksRequest {
	s.SrcInstanceId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcPublicIp(v string) *CreateCloudBenchTasksRequest {
	s.SrcPublicIp = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcSuperAccount(v string) *CreateCloudBenchTasksRequest {
	s.SrcSuperAccount = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcSuperPassword(v string) *CreateCloudBenchTasksRequest {
	s.SrcSuperPassword = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetTaskType(v string) *CreateCloudBenchTasksRequest {
	s.TaskType = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetWorkDir(v string) *CreateCloudBenchTasksRequest {
	s.WorkDir = &v
	return s
}

type CreateCloudBenchTasksResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateCloudBenchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCloudBenchTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudBenchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksResponseBody) SetCode(v string) *CreateCloudBenchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetData(v *CreateCloudBenchTasksResponseBodyData) *CreateCloudBenchTasksResponseBody {
	s.Data = v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetMessage(v string) *CreateCloudBenchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetRequestId(v string) *CreateCloudBenchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetSuccess(v string) *CreateCloudBenchTasksResponseBody {
	s.Success = &v
	return s
}

type CreateCloudBenchTasksResponseBodyData struct {
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s CreateCloudBenchTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudBenchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksResponseBodyData) SetTaskIds(v []*string) *CreateCloudBenchTasksResponseBodyData {
	s.TaskIds = v
	return s
}

type CreateCloudBenchTasksResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCloudBenchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCloudBenchTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudBenchTasksResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksResponse) SetHeaders(v map[string]*string) *CreateCloudBenchTasksResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudBenchTasksResponse) SetBody(v *CreateCloudBenchTasksResponseBody) *CreateCloudBenchTasksResponse {
	s.Body = v
	return s
}

type CreateDiagnosticReportRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticReportRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportRequest) SetDBInstanceId(v string) *CreateDiagnosticReportRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetEndTime(v string) *CreateDiagnosticReportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetStartTime(v string) *CreateDiagnosticReportRequest {
	s.StartTime = &v
	return s
}

type CreateDiagnosticReportResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDiagnosticReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponseBody) SetCode(v string) *CreateDiagnosticReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetData(v string) *CreateDiagnosticReportResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetMessage(v string) *CreateDiagnosticReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetRequestId(v string) *CreateDiagnosticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetSuccess(v string) *CreateDiagnosticReportResponseBody {
	s.Success = &v
	return s
}

type CreateDiagnosticReportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDiagnosticReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiagnosticReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticReportResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponse) SetHeaders(v map[string]*string) *CreateDiagnosticReportResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticReportResponse) SetBody(v *CreateDiagnosticReportResponseBody) *CreateDiagnosticReportResponse {
	s.Body = v
	return s
}

type CreateRequestDiagnosisRequest struct {
	Database   *string `json:"Database,omitempty" xml:"Database,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Sql        *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s CreateRequestDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRequestDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateRequestDiagnosisRequest) SetDatabase(v string) *CreateRequestDiagnosisRequest {
	s.Database = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) SetInstanceId(v string) *CreateRequestDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) SetNodeId(v string) *CreateRequestDiagnosisRequest {
	s.NodeId = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) SetSql(v string) *CreateRequestDiagnosisRequest {
	s.Sql = &v
	return s
}

type CreateRequestDiagnosisResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRequestDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRequestDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRequestDiagnosisResponseBody) SetCode(v string) *CreateRequestDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetData(v string) *CreateRequestDiagnosisResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetMessage(v string) *CreateRequestDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetRequestId(v string) *CreateRequestDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetSuccess(v string) *CreateRequestDiagnosisResponseBody {
	s.Success = &v
	return s
}

type CreateRequestDiagnosisResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRequestDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRequestDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRequestDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateRequestDiagnosisResponse) SetHeaders(v map[string]*string) *CreateRequestDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateRequestDiagnosisResponse) SetBody(v *CreateRequestDiagnosisResponseBody) *CreateRequestDiagnosisResponse {
	s.Body = v
	return s
}

type DeleteCloudBenchTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteCloudBenchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudBenchTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudBenchTaskRequest) SetTaskId(v string) *DeleteCloudBenchTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteCloudBenchTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCloudBenchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudBenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudBenchTaskResponseBody) SetCode(v string) *DeleteCloudBenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetData(v string) *DeleteCloudBenchTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetMessage(v string) *DeleteCloudBenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetRequestId(v string) *DeleteCloudBenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudBenchTaskResponseBody) SetSuccess(v string) *DeleteCloudBenchTaskResponseBody {
	s.Success = &v
	return s
}

type DeleteCloudBenchTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCloudBenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCloudBenchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudBenchTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudBenchTaskResponse) SetHeaders(v map[string]*string) *DeleteCloudBenchTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudBenchTaskResponse) SetBody(v *DeleteCloudBenchTaskResponseBody) *DeleteCloudBenchTaskResponse {
	s.Body = v
	return s
}

type DescribeCacheAnalysisJobRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeCacheAnalysisJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobRequest) SetInstanceId(v string) *DescribeCacheAnalysisJobRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobRequest) SetJobId(v string) *DescribeCacheAnalysisJobRequest {
	s.JobId = &v
	return s
}

type DescribeCacheAnalysisJobResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCacheAnalysisJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBody) SetCode(v string) *DescribeCacheAnalysisJobResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetData(v *DescribeCacheAnalysisJobResponseBodyData) *DescribeCacheAnalysisJobResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetMessage(v string) *DescribeCacheAnalysisJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetRequestId(v string) *DescribeCacheAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetSuccess(v string) *DescribeCacheAnalysisJobResponseBody {
	s.Success = &v
	return s
}

type DescribeCacheAnalysisJobResponseBodyData struct {
	BigKeys     *DescribeCacheAnalysisJobResponseBodyDataBigKeys     `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	InstanceId  *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId       *string                                              `json:"JobId,omitempty" xml:"JobId,omitempty"`
	KeyPrefixes *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes `json:"KeyPrefixes,omitempty" xml:"KeyPrefixes,omitempty" type:"Struct"`
	Message     *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeId      *string                                              `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TaskState   *string                                              `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetBigKeys(v *DescribeCacheAnalysisJobResponseBodyDataBigKeys) *DescribeCacheAnalysisJobResponseBodyData {
	s.BigKeys = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetInstanceId(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetJobId(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetKeyPrefixes(v *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) *DescribeCacheAnalysisJobResponseBodyData {
	s.KeyPrefixes = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetMessage(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetTaskState(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.TaskState = &v
	return s
}

type DescribeCacheAnalysisJobResponseBodyDataBigKeys struct {
	KeyInfo []*DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeys) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeys) SetKeyInfo(v []*DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) *DescribeCacheAnalysisJobResponseBodyDataBigKeys {
	s.KeyInfo = v
	return s
}

type DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo struct {
	Bytes                *int64  `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	Count                *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Db                   *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Encoding             *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	ExpirationTimeMillis *int64  `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Type = &v
	return s
}

type DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes struct {
	Prefix []*DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix `json:"Prefix,omitempty" xml:"Prefix,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) SetPrefix(v []*DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes {
	s.Prefix = v
	return s
}

type DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix struct {
	Bytes  *int64  `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	Count  *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	KeyNum *int64  `json:"KeyNum,omitempty" xml:"KeyNum,omitempty"`
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetKeyNum(v int64) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.KeyNum = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetPrefix(v string) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Prefix = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Type = &v
	return s
}

type DescribeCacheAnalysisJobResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCacheAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCacheAnalysisJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisJobResponse) SetBody(v *DescribeCacheAnalysisJobResponseBody) *DescribeCacheAnalysisJobResponse {
	s.Body = v
	return s
}

type DescribeCacheAnalysisJobsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo     *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCacheAnalysisJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsRequest) SetEndTime(v string) *DescribeCacheAnalysisJobsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetInstanceId(v string) *DescribeCacheAnalysisJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetPageNo(v string) *DescribeCacheAnalysisJobsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetPageSize(v string) *DescribeCacheAnalysisJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisJobsRequest) SetStartTime(v string) *DescribeCacheAnalysisJobsRequest {
	s.StartTime = &v
	return s
}

type DescribeCacheAnalysisJobsResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCacheAnalysisJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetCode(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetData(v *DescribeCacheAnalysisJobsResponseBodyData) *DescribeCacheAnalysisJobsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetMessage(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetRequestId(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetSuccess(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.Success = &v
	return s
}

type DescribeCacheAnalysisJobsResponseBodyData struct {
	Extra    *string                                        `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     *DescribeCacheAnalysisJobsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	PageNo   *int64                                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetExtra(v string) *DescribeCacheAnalysisJobsResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetList(v *DescribeCacheAnalysisJobsResponseBodyDataList) *DescribeCacheAnalysisJobsResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetPageNo(v int64) *DescribeCacheAnalysisJobsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetPageSize(v int64) *DescribeCacheAnalysisJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetTotal(v int64) *DescribeCacheAnalysisJobsResponseBodyData {
	s.Total = &v
	return s
}

type DescribeCacheAnalysisJobsResponseBodyDataList struct {
	CacheAnalysisJob []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob `json:"CacheAnalysisJob,omitempty" xml:"CacheAnalysisJob,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataList) SetCacheAnalysisJob(v []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) *DescribeCacheAnalysisJobsResponseBodyDataList {
	s.CacheAnalysisJob = v
	return s
}

type DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob struct {
	BigKeys    *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	InstanceId *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string                                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message    *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeId     *string                                                               `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TaskState  *string                                                               `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetBigKeys(v *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.BigKeys = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetInstanceId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetJobId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.JobId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetMessage(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetNodeId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetTaskState(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.TaskState = &v
	return s
}

type DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys struct {
	KeyInfo []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) SetKeyInfo(v []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys {
	s.KeyInfo = v
	return s
}

type DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo struct {
	Bytes                *int64  `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	Count                *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Db                   *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Encoding             *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	ExpirationTimeMillis *int64  `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetType(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Type = &v
	return s
}

type DescribeCacheAnalysisJobsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCacheAnalysisJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCacheAnalysisJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponse) SetBody(v *DescribeCacheAnalysisJobsResponseBody) *DescribeCacheAnalysisJobsResponse {
	s.Body = v
	return s
}

type DescribeCloudBenchTasksRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo    *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType  *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeCloudBenchTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudBenchTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksRequest) SetEndTime(v string) *DescribeCloudBenchTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetPageNo(v string) *DescribeCloudBenchTasksRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetPageSize(v string) *DescribeCloudBenchTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetStartTime(v string) *DescribeCloudBenchTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetStatus(v string) *DescribeCloudBenchTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetTaskType(v string) *DescribeCloudBenchTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeCloudBenchTasksResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCloudBenchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBody) SetCode(v string) *DescribeCloudBenchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetData(v *DescribeCloudBenchTasksResponseBodyData) *DescribeCloudBenchTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetMessage(v string) *DescribeCloudBenchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetRequestId(v string) *DescribeCloudBenchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetSuccess(v string) *DescribeCloudBenchTasksResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudBenchTasksResponseBodyData struct {
	Extra    *string                                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     *DescribeCloudBenchTasksResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	PageNo   *int32                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetExtra(v string) *DescribeCloudBenchTasksResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetList(v *DescribeCloudBenchTasksResponseBodyDataList) *DescribeCloudBenchTasksResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetPageNo(v int32) *DescribeCloudBenchTasksResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetPageSize(v int32) *DescribeCloudBenchTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetTotal(v int64) *DescribeCloudBenchTasksResponseBodyData {
	s.Total = &v
	return s
}

type DescribeCloudBenchTasksResponseBodyDataList struct {
	CloudbenchTasks []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks `json:"cloudbenchTasks,omitempty" xml:"cloudbenchTasks,omitempty" type:"Repeated"`
}

func (s DescribeCloudBenchTasksResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) SetCloudbenchTasks(v []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) *DescribeCloudBenchTasksResponseBodyDataList {
	s.CloudbenchTasks = v
	return s
}

type DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks struct {
	ArchiveJobId        *string `json:"ArchiveJobId,omitempty" xml:"ArchiveJobId,omitempty"`
	ArchiveOssTableName *string `json:"ArchiveOssTableName,omitempty" xml:"ArchiveOssTableName,omitempty"`
	ArchiveState        *int32  `json:"ArchiveState,omitempty" xml:"ArchiveState,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupType          *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BenchStep           *string `json:"BenchStep,omitempty" xml:"BenchStep,omitempty"`
	BenchStepStatus     *string `json:"BenchStepStatus,omitempty" xml:"BenchStepStatus,omitempty"`
	ClientGatewayId     *string `json:"ClientGatewayId,omitempty" xml:"ClientGatewayId,omitempty"`
	ClientType          *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DstInstanceUuid     *string `json:"DstInstanceUuid,omitempty" xml:"DstInstanceUuid,omitempty"`
	DstIp               *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	DstPort             *int32  `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	DstType             *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	DtsJobClass         *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	DtsJobId            *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsJobName          *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	DtsJobState         *int32  `json:"DtsJobState,omitempty" xml:"DtsJobState,omitempty"`
	DtsJobStatus        *string `json:"DtsJobStatus,omitempty" xml:"DtsJobStatus,omitempty"`
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EndState            *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	External            *string `json:"External,omitempty" xml:"External,omitempty"`
	Rate                *int32  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	RequestDuration     *int64  `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	SmartPressureTime   *int32  `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SqlCompleteReuse    *string `json:"SqlCompleteReuse,omitempty" xml:"SqlCompleteReuse,omitempty"`
	SrcInstanceArea     *string `json:"SrcInstanceArea,omitempty" xml:"SrcInstanceArea,omitempty"`
	SrcInstanceUuid     *string `json:"SrcInstanceUuid,omitempty" xml:"SrcInstanceUuid,omitempty"`
	SrcPublicIp         *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableSchema         *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Topic               *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkDir             *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveJobId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveJobId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveOssTableName(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveOssTableName = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveState(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBackupId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BackupId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBackupType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BackupType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBenchStep(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BenchStep = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBenchStepStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BenchStepStatus = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetClientGatewayId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ClientGatewayId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetClientType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ClientType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDescription(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Description = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstInstanceUuid(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstInstanceUuid = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstIp(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstIp = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstPort(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstPort = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobClass(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobName(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobName = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobState(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobStatus = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetEcsInstanceId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetEndState(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.EndState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetErrorCode(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetErrorMessage(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetExternal(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.External = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetRate(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Rate = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetRequestDuration(v int64) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.RequestDuration = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSmartPressureTime(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SmartPressureTime = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSource(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Source = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSqlCompleteReuse(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SqlCompleteReuse = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcInstanceArea(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcInstanceArea = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcInstanceUuid(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcInstanceUuid = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcPublicIp(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcPublicIp = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetState(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.State = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Status = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTableSchema(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TableSchema = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTaskId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTaskType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTopic(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Topic = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetUserId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.UserId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetVersion(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Version = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetWorkDir(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.WorkDir = &v
	return s
}

type DescribeCloudBenchTasksResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCloudBenchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudBenchTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudBenchTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponse) SetHeaders(v map[string]*string) *DescribeCloudBenchTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudBenchTasksResponse) SetBody(v *DescribeCloudBenchTasksResponseBody) *DescribeCloudBenchTasksResponse {
	s.Body = v
	return s
}

type DescribeCloudbenchTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCloudbenchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskRequest) SetTaskId(v string) *DescribeCloudbenchTaskRequest {
	s.TaskId = &v
	return s
}

type DescribeCloudbenchTaskResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCloudbenchTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudbenchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskResponseBody) SetCode(v string) *DescribeCloudbenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetData(v *DescribeCloudbenchTaskResponseBodyData) *DescribeCloudbenchTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetMessage(v string) *DescribeCloudbenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetRequestId(v string) *DescribeCloudbenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetSuccess(v string) *DescribeCloudbenchTaskResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudbenchTaskResponseBodyData struct {
	ArchiveJobId        *string `json:"ArchiveJobId,omitempty" xml:"ArchiveJobId,omitempty"`
	ArchiveOssTableName *string `json:"ArchiveOssTableName,omitempty" xml:"ArchiveOssTableName,omitempty"`
	ArchiveState        *int32  `json:"ArchiveState,omitempty" xml:"ArchiveState,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupType          *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BenchStep           *string `json:"BenchStep,omitempty" xml:"BenchStep,omitempty"`
	BenchStepStatus     *string `json:"BenchStepStatus,omitempty" xml:"BenchStepStatus,omitempty"`
	ClientGatewayId     *string `json:"ClientGatewayId,omitempty" xml:"ClientGatewayId,omitempty"`
	ClientType          *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DstInstanceUuid     *string `json:"DstInstanceUuid,omitempty" xml:"DstInstanceUuid,omitempty"`
	DstIp               *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	DstPort             *int32  `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	DstType             *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	DtsJobClass         *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	DtsJobId            *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsJobName          *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	DtsJobState         *int32  `json:"DtsJobState,omitempty" xml:"DtsJobState,omitempty"`
	DtsJobStatus        *string `json:"DtsJobStatus,omitempty" xml:"DtsJobStatus,omitempty"`
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EndState            *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	External            *string `json:"External,omitempty" xml:"External,omitempty"`
	Rate                *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	RequestDuration     *int64  `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	SmartPressureTime   *int32  `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SqlCompleteReuse    *string `json:"SqlCompleteReuse,omitempty" xml:"SqlCompleteReuse,omitempty"`
	SrcInstanceArea     *string `json:"SrcInstanceArea,omitempty" xml:"SrcInstanceArea,omitempty"`
	SrcInstanceUuid     *string `json:"SrcInstanceUuid,omitempty" xml:"SrcInstanceUuid,omitempty"`
	SrcPublicIp         *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableSchema         *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Topic               *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkDir             *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudbenchTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetArchiveJobId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ArchiveJobId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetArchiveOssTableName(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ArchiveOssTableName = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetArchiveState(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.ArchiveState = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBackupId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BackupId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBackupType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBenchStep(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BenchStep = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBenchStepStatus(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BenchStepStatus = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetClientGatewayId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ClientGatewayId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetClientType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDescription(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstInstanceUuid(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DstInstanceUuid = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstIp(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DstIp = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstPort(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.DstPort = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DstType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobClass(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobName(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobName = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobState(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobState = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobStatus(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobStatus = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetEcsInstanceId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetEndState(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.EndState = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetErrorCode(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetErrorMessage(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetExternal(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.External = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetRate(v int64) *DescribeCloudbenchTaskResponseBodyData {
	s.Rate = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetRequestDuration(v int64) *DescribeCloudbenchTaskResponseBodyData {
	s.RequestDuration = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSmartPressureTime(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.SmartPressureTime = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSource(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSqlCompleteReuse(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SqlCompleteReuse = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSrcInstanceArea(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SrcInstanceArea = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSrcInstanceUuid(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SrcInstanceUuid = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSrcPublicIp(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SrcPublicIp = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetState(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetStatus(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTableSchema(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.TableSchema = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTaskId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTaskType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTopic(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Topic = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetUserId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetVersion(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetWorkDir(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.WorkDir = &v
	return s
}

type DescribeCloudbenchTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCloudbenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudbenchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskResponse) SetHeaders(v map[string]*string) *DescribeCloudbenchTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudbenchTaskResponse) SetBody(v *DescribeCloudbenchTaskResponseBody) *DescribeCloudbenchTaskResponse {
	s.Body = v
	return s
}

type DescribeCloudbenchTaskConfigRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCloudbenchTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigRequest) SetTaskId(v string) *DescribeCloudbenchTaskConfigRequest {
	s.TaskId = &v
	return s
}

type DescribeCloudbenchTaskConfigResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCloudbenchTaskConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudbenchTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetCode(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetData(v *DescribeCloudbenchTaskConfigResponseBodyData) *DescribeCloudbenchTaskConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetMessage(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetRequestId(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBody) SetSuccess(v string) *DescribeCloudbenchTaskConfigResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudbenchTaskConfigResponseBodyData struct {
	ArchiveFolder *string `json:"ArchiveFolder,omitempty" xml:"ArchiveFolder,omitempty"`
	BenchCmd      *string `json:"BenchCmd,omitempty" xml:"BenchCmd,omitempty"`
	ClientJarPath *string `json:"ClientJarPath,omitempty" xml:"ClientJarPath,omitempty"`
	JarOnOss      *string `json:"JarOnOss,omitempty" xml:"JarOnOss,omitempty"`
	LoadCmd       *string `json:"LoadCmd,omitempty" xml:"LoadCmd,omitempty"`
	MetaFileName  *string `json:"MetaFileName,omitempty" xml:"MetaFileName,omitempty"`
	MetaFileOnOss *string `json:"MetaFileOnOss,omitempty" xml:"MetaFileOnOss,omitempty"`
	MetaFilePath  *string `json:"MetaFilePath,omitempty" xml:"MetaFilePath,omitempty"`
	ParseCmd      *string `json:"ParseCmd,omitempty" xml:"ParseCmd,omitempty"`
	ParseFilePath *string `json:"ParseFilePath,omitempty" xml:"ParseFilePath,omitempty"`
	RocksDbPath   *string `json:"RocksDbPath,omitempty" xml:"RocksDbPath,omitempty"`
	SqlFileName   *string `json:"SqlFileName,omitempty" xml:"SqlFileName,omitempty"`
	SqlFileOnOss  *string `json:"SqlFileOnOss,omitempty" xml:"SqlFileOnOss,omitempty"`
	SqlFilePath   *string `json:"SqlFilePath,omitempty" xml:"SqlFilePath,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkDir       *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudbenchTaskConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetArchiveFolder(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ArchiveFolder = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetBenchCmd(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.BenchCmd = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetClientJarPath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ClientJarPath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetJarOnOss(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.JarOnOss = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetLoadCmd(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.LoadCmd = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetMetaFileName(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.MetaFileName = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetMetaFileOnOss(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.MetaFileOnOss = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetMetaFilePath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.MetaFilePath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetParseCmd(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ParseCmd = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetParseFilePath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.ParseFilePath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetRocksDbPath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.RocksDbPath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetSqlFileName(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.SqlFileName = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetSqlFileOnOss(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.SqlFileOnOss = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetSqlFilePath(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.SqlFilePath = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetTaskId(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetUserId(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponseBodyData) SetWorkDir(v string) *DescribeCloudbenchTaskConfigResponseBodyData {
	s.WorkDir = &v
	return s
}

type DescribeCloudbenchTaskConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCloudbenchTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudbenchTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigResponse) SetHeaders(v map[string]*string) *DescribeCloudbenchTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponse) SetBody(v *DescribeCloudbenchTaskConfigResponseBody) *DescribeCloudbenchTaskConfigResponse {
	s.Body = v
	return s
}

type DescribeDiagnosticReportListRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo       *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosticReportListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListRequest) SetDBInstanceId(v string) *DescribeDiagnosticReportListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetEndTime(v string) *DescribeDiagnosticReportListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetPageNo(v string) *DescribeDiagnosticReportListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetPageSize(v string) *DescribeDiagnosticReportListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetStartTime(v string) *DescribeDiagnosticReportListRequest {
	s.StartTime = &v
	return s
}

type DescribeDiagnosticReportListResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DescribeDiagnosticReportListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponseBody) SetCode(v string) *DescribeDiagnosticReportListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetData(v string) *DescribeDiagnosticReportListResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetMessage(v string) *DescribeDiagnosticReportListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetRequestId(v string) *DescribeDiagnosticReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetSuccess(v string) *DescribeDiagnosticReportListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetSynchro(v string) *DescribeDiagnosticReportListResponseBody {
	s.Synchro = &v
	return s
}

type DescribeDiagnosticReportListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiagnosticReportListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosticReportListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticReportListResponse) SetBody(v *DescribeDiagnosticReportListResponseBody) *DescribeDiagnosticReportListResponse {
	s.Body = v
	return s
}

type DescribeHotBigKeysRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeHotBigKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysRequest) SetConsoleContext(v string) *DescribeHotBigKeysRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeHotBigKeysRequest) SetInstanceId(v string) *DescribeHotBigKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHotBigKeysRequest) SetNodeId(v string) *DescribeHotBigKeysRequest {
	s.NodeId = &v
	return s
}

type DescribeHotBigKeysResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeHotBigKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHotBigKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBody) SetCode(v string) *DescribeHotBigKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetData(v *DescribeHotBigKeysResponseBodyData) *DescribeHotBigKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetMessage(v string) *DescribeHotBigKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetRequestId(v string) *DescribeHotBigKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetSuccess(v string) *DescribeHotBigKeysResponseBody {
	s.Success = &v
	return s
}

type DescribeHotBigKeysResponseBodyData struct {
	BigKeyMsg *string                                    `json:"BigKeyMsg,omitempty" xml:"BigKeyMsg,omitempty"`
	BigKeys   *DescribeHotBigKeysResponseBodyDataBigKeys `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	HotKeyMsg *string                                    `json:"HotKeyMsg,omitempty" xml:"HotKeyMsg,omitempty"`
	HotKeys   *DescribeHotBigKeysResponseBodyDataHotKeys `json:"HotKeys,omitempty" xml:"HotKeys,omitempty" type:"Struct"`
}

func (s DescribeHotBigKeysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyData) SetBigKeyMsg(v string) *DescribeHotBigKeysResponseBodyData {
	s.BigKeyMsg = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetBigKeys(v *DescribeHotBigKeysResponseBodyDataBigKeys) *DescribeHotBigKeysResponseBodyData {
	s.BigKeys = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetHotKeyMsg(v string) *DescribeHotBigKeysResponseBodyData {
	s.HotKeyMsg = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetHotKeys(v *DescribeHotBigKeysResponseBodyDataHotKeys) *DescribeHotBigKeysResponseBodyData {
	s.HotKeys = v
	return s
}

type DescribeHotBigKeysResponseBodyDataBigKeys struct {
	BigKey []*DescribeHotBigKeysResponseBodyDataBigKeysBigKey `json:"BigKey,omitempty" xml:"BigKey,omitempty" type:"Repeated"`
}

func (s DescribeHotBigKeysResponseBodyDataBigKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataBigKeys) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeys) SetBigKey(v []*DescribeHotBigKeysResponseBodyDataBigKeysBigKey) *DescribeHotBigKeysResponseBodyDataBigKeys {
	s.BigKey = v
	return s
}

type DescribeHotBigKeysResponseBodyDataBigKeysBigKey struct {
	Db      *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	NodeId  *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Size    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeHotBigKeysResponseBodyDataBigKeysBigKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetDb(v int32) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.Db = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetKey(v string) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.Key = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetKeyType(v string) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetNodeId(v string) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.NodeId = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetSize(v int64) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.Size = &v
	return s
}

type DescribeHotBigKeysResponseBodyDataHotKeys struct {
	HotKey []*DescribeHotBigKeysResponseBodyDataHotKeysHotKey `json:"HotKey,omitempty" xml:"HotKey,omitempty" type:"Repeated"`
}

func (s DescribeHotBigKeysResponseBodyDataHotKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataHotKeys) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeys) SetHotKey(v []*DescribeHotBigKeysResponseBodyDataHotKeysHotKey) *DescribeHotBigKeysResponseBodyDataHotKeys {
	s.HotKey = v
	return s
}

type DescribeHotBigKeysResponseBodyDataHotKeysHotKey struct {
	Db      *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Hot     *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	Lfu     *int32  `json:"Lfu,omitempty" xml:"Lfu,omitempty"`
	NodeId  *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeHotBigKeysResponseBodyDataHotKeysHotKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetDb(v int32) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Db = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetHot(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Hot = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetKey(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Key = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetKeyType(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetLfu(v int32) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Lfu = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetNodeId(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.NodeId = &v
	return s
}

type DescribeHotBigKeysResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHotBigKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHotBigKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotBigKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponse) SetHeaders(v map[string]*string) *DescribeHotBigKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotBigKeysResponse) SetBody(v *DescribeHotBigKeysResponseBody) *DescribeHotBigKeysResponse {
	s.Body = v
	return s
}

type DescribeHotKeysRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeHotKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysRequest) SetInstanceId(v string) *DescribeHotKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHotKeysRequest) SetNodeId(v string) *DescribeHotKeysRequest {
	s.NodeId = &v
	return s
}

type DescribeHotKeysResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeHotKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHotKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponseBody) SetCode(v string) *DescribeHotKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHotKeysResponseBody) SetData(v *DescribeHotKeysResponseBodyData) *DescribeHotKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHotKeysResponseBody) SetMessage(v string) *DescribeHotKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHotKeysResponseBody) SetRequestId(v string) *DescribeHotKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotKeysResponseBody) SetSuccess(v string) *DescribeHotKeysResponseBody {
	s.Success = &v
	return s
}

type DescribeHotKeysResponseBodyData struct {
	HotKey []*DescribeHotKeysResponseBodyDataHotKey `json:"HotKey,omitempty" xml:"HotKey,omitempty" type:"Repeated"`
}

func (s DescribeHotKeysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponseBodyData) SetHotKey(v []*DescribeHotKeysResponseBodyDataHotKey) *DescribeHotKeysResponseBodyData {
	s.HotKey = v
	return s
}

type DescribeHotKeysResponseBodyDataHotKey struct {
	Db      *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Hot     *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	Size    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeHotKeysResponseBodyDataHotKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotKeysResponseBodyDataHotKey) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetDb(v int32) *DescribeHotKeysResponseBodyDataHotKey {
	s.Db = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetHot(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.Hot = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetKey(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.Key = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetKeyType(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetSize(v int64) *DescribeHotKeysResponseBodyDataHotKey {
	s.Size = &v
	return s
}

type DescribeHotKeysResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHotKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHotKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHotKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponse) SetHeaders(v map[string]*string) *DescribeHotKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotKeysResponse) SetBody(v *DescribeHotKeysResponseBody) *DescribeHotKeysResponse {
	s.Body = v
	return s
}

type DescribeInstanceDasProRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceDasProRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDasProRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDasProRequest) SetInstanceId(v string) *DescribeInstanceDasProRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceDasProResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceDasProResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDasProResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDasProResponseBody) SetCode(v string) *DescribeInstanceDasProResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetData(v bool) *DescribeInstanceDasProResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetMessage(v string) *DescribeInstanceDasProResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetRequestId(v string) *DescribeInstanceDasProResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetSuccess(v string) *DescribeInstanceDasProResponseBody {
	s.Success = &v
	return s
}

type DescribeInstanceDasProResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceDasProResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceDasProResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDasProResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDasProResponse) SetHeaders(v map[string]*string) *DescribeInstanceDasProResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDasProResponse) SetBody(v *DescribeInstanceDasProResponseBody) *DescribeInstanceDasProResponse {
	s.Body = v
	return s
}

type DescribeTopBigKeysRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTopBigKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopBigKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysRequest) SetConsoleContext(v string) *DescribeTopBigKeysRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetEndTime(v string) *DescribeTopBigKeysRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetInstanceId(v string) *DescribeTopBigKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetNodeId(v string) *DescribeTopBigKeysRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetStartTime(v string) *DescribeTopBigKeysRequest {
	s.StartTime = &v
	return s
}

type DescribeTopBigKeysResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeTopBigKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTopBigKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopBigKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponseBody) SetCode(v string) *DescribeTopBigKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetData(v *DescribeTopBigKeysResponseBodyData) *DescribeTopBigKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetMessage(v string) *DescribeTopBigKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetRequestId(v string) *DescribeTopBigKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetSuccess(v string) *DescribeTopBigKeysResponseBody {
	s.Success = &v
	return s
}

type DescribeTopBigKeysResponseBodyData struct {
	BigKey []*DescribeTopBigKeysResponseBodyDataBigKey `json:"BigKey,omitempty" xml:"BigKey,omitempty" type:"Repeated"`
}

func (s DescribeTopBigKeysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopBigKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponseBodyData) SetBigKey(v []*DescribeTopBigKeysResponseBodyDataBigKey) *DescribeTopBigKeysResponseBodyData {
	s.BigKey = v
	return s
}

type DescribeTopBigKeysResponseBodyDataBigKey struct {
	Db      *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	NodeId  *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Size    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTopBigKeysResponseBodyDataBigKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopBigKeysResponseBodyDataBigKey) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetDb(v int32) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.Db = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetKey(v string) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.Key = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetKeyType(v string) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.KeyType = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetNodeId(v string) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.NodeId = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetSize(v int64) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.Size = &v
	return s
}

type DescribeTopBigKeysResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTopBigKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopBigKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopBigKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponse) SetHeaders(v map[string]*string) *DescribeTopBigKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopBigKeysResponse) SetBody(v *DescribeTopBigKeysResponseBody) *DescribeTopBigKeysResponse {
	s.Body = v
	return s
}

type DescribeTopHotKeysRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTopHotKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopHotKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysRequest) SetConsoleContext(v string) *DescribeTopHotKeysRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetEndTime(v string) *DescribeTopHotKeysRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetInstanceId(v string) *DescribeTopHotKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetNodeId(v string) *DescribeTopHotKeysRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetStartTime(v string) *DescribeTopHotKeysRequest {
	s.StartTime = &v
	return s
}

type DescribeTopHotKeysResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeTopHotKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTopHotKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopHotKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponseBody) SetCode(v string) *DescribeTopHotKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetData(v *DescribeTopHotKeysResponseBodyData) *DescribeTopHotKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetMessage(v string) *DescribeTopHotKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetRequestId(v string) *DescribeTopHotKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetSuccess(v string) *DescribeTopHotKeysResponseBody {
	s.Success = &v
	return s
}

type DescribeTopHotKeysResponseBodyData struct {
	HotKey []*DescribeTopHotKeysResponseBodyDataHotKey `json:"HotKey,omitempty" xml:"HotKey,omitempty" type:"Repeated"`
}

func (s DescribeTopHotKeysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopHotKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponseBodyData) SetHotKey(v []*DescribeTopHotKeysResponseBodyDataHotKey) *DescribeTopHotKeysResponseBodyData {
	s.HotKey = v
	return s
}

type DescribeTopHotKeysResponseBodyDataHotKey struct {
	Db      *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Hot     *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	Lfu     *int32  `json:"Lfu,omitempty" xml:"Lfu,omitempty"`
	NodeId  *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeTopHotKeysResponseBodyDataHotKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopHotKeysResponseBodyDataHotKey) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetDb(v int32) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Db = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetHot(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Hot = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetKey(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Key = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetKeyType(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.KeyType = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetLfu(v int32) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Lfu = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetNodeId(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.NodeId = &v
	return s
}

type DescribeTopHotKeysResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTopHotKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopHotKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopHotKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponse) SetHeaders(v map[string]*string) *DescribeTopHotKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopHotKeysResponse) SetBody(v *DescribeTopHotKeysResponseBody) *DescribeTopHotKeysResponse {
	s.Body = v
	return s
}

type DisableAllSqlConcurrencyControlRulesRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableAllSqlConcurrencyControlRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAllSqlConcurrencyControlRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) SetConsoleContext(v string) *DisableAllSqlConcurrencyControlRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesRequest) SetInstanceId(v string) *DisableAllSqlConcurrencyControlRulesRequest {
	s.InstanceId = &v
	return s
}

type DisableAllSqlConcurrencyControlRulesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableAllSqlConcurrencyControlRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAllSqlConcurrencyControlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetCode(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetData(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Data = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetMessage(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetRequestId(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetSuccess(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Success = &v
	return s
}

type DisableAllSqlConcurrencyControlRulesResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableAllSqlConcurrencyControlRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAllSqlConcurrencyControlRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAllSqlConcurrencyControlRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) SetHeaders(v map[string]*string) *DisableAllSqlConcurrencyControlRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) SetBody(v *DisableAllSqlConcurrencyControlRulesResponseBody) *DisableAllSqlConcurrencyControlRulesResponse {
	s.Body = v
	return s
}

type DisableDasProRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DisableDasProRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDasProRequest) GoString() string {
	return s.String()
}

func (s *DisableDasProRequest) SetInstanceId(v string) *DisableDasProRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableDasProRequest) SetUserId(v string) *DisableDasProRequest {
	s.UserId = &v
	return s
}

type DisableDasProResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DisableDasProResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDasProResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDasProResponseBody) SetCode(v string) *DisableDasProResponseBody {
	s.Code = &v
	return s
}

func (s *DisableDasProResponseBody) SetData(v string) *DisableDasProResponseBody {
	s.Data = &v
	return s
}

func (s *DisableDasProResponseBody) SetMessage(v string) *DisableDasProResponseBody {
	s.Message = &v
	return s
}

func (s *DisableDasProResponseBody) SetRequestId(v string) *DisableDasProResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDasProResponseBody) SetSuccess(v string) *DisableDasProResponseBody {
	s.Success = &v
	return s
}

func (s *DisableDasProResponseBody) SetSynchro(v string) *DisableDasProResponseBody {
	s.Synchro = &v
	return s
}

type DisableDasProResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableDasProResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDasProResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDasProResponse) GoString() string {
	return s.String()
}

func (s *DisableDasProResponse) SetHeaders(v map[string]*string) *DisableDasProResponse {
	s.Headers = v
	return s
}

func (s *DisableDasProResponse) SetBody(v *DisableDasProResponseBody) *DisableDasProResponse {
	s.Body = v
	return s
}

type DisableSqlConcurrencyControlRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId         *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s DisableSqlConcurrencyControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlConcurrencyControlRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlConcurrencyControlRequest) SetConsoleContext(v string) *DisableSqlConcurrencyControlRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DisableSqlConcurrencyControlRequest) SetInstanceId(v string) *DisableSqlConcurrencyControlRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableSqlConcurrencyControlRequest) SetItemId(v int64) *DisableSqlConcurrencyControlRequest {
	s.ItemId = &v
	return s
}

type DisableSqlConcurrencyControlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSqlConcurrencyControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlConcurrencyControlResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSqlConcurrencyControlResponseBody) SetCode(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetData(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetMessage(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetRequestId(v string) *DisableSqlConcurrencyControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponseBody) SetSuccess(v string) *DisableSqlConcurrencyControlResponseBody {
	s.Success = &v
	return s
}

type DisableSqlConcurrencyControlResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableSqlConcurrencyControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSqlConcurrencyControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSqlConcurrencyControlResponse) GoString() string {
	return s.String()
}

func (s *DisableSqlConcurrencyControlResponse) SetHeaders(v map[string]*string) *DisableSqlConcurrencyControlResponse {
	s.Headers = v
	return s
}

func (s *DisableSqlConcurrencyControlResponse) SetBody(v *DisableSqlConcurrencyControlResponseBody) *DisableSqlConcurrencyControlResponse {
	s.Body = v
	return s
}

type EnableDasProRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SqlRetention *int32  `json:"SqlRetention,omitempty" xml:"SqlRetention,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s EnableDasProRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableDasProRequest) GoString() string {
	return s.String()
}

func (s *EnableDasProRequest) SetInstanceId(v string) *EnableDasProRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableDasProRequest) SetSqlRetention(v int32) *EnableDasProRequest {
	s.SqlRetention = &v
	return s
}

func (s *EnableDasProRequest) SetUserId(v string) *EnableDasProRequest {
	s.UserId = &v
	return s
}

type EnableDasProResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s EnableDasProResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableDasProResponseBody) GoString() string {
	return s.String()
}

func (s *EnableDasProResponseBody) SetCode(v string) *EnableDasProResponseBody {
	s.Code = &v
	return s
}

func (s *EnableDasProResponseBody) SetData(v string) *EnableDasProResponseBody {
	s.Data = &v
	return s
}

func (s *EnableDasProResponseBody) SetMessage(v string) *EnableDasProResponseBody {
	s.Message = &v
	return s
}

func (s *EnableDasProResponseBody) SetRequestId(v string) *EnableDasProResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableDasProResponseBody) SetSuccess(v string) *EnableDasProResponseBody {
	s.Success = &v
	return s
}

func (s *EnableDasProResponseBody) SetSynchro(v string) *EnableDasProResponseBody {
	s.Synchro = &v
	return s
}

type EnableDasProResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableDasProResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableDasProResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableDasProResponse) GoString() string {
	return s.String()
}

func (s *EnableDasProResponse) SetHeaders(v map[string]*string) *EnableDasProResponse {
	s.Headers = v
	return s
}

func (s *EnableDasProResponse) SetBody(v *EnableDasProResponseBody) *EnableDasProResponse {
	s.Body = v
	return s
}

type EnableSqlConcurrencyControlRequest struct {
	ConcurrencyControlTime *int64  `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
	ConsoleContext         *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxConcurrency         *int64  `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	SqlKeywords            *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
	SqlType                *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s EnableSqlConcurrencyControlRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlConcurrencyControlRequest) GoString() string {
	return s.String()
}

func (s *EnableSqlConcurrencyControlRequest) SetConcurrencyControlTime(v int64) *EnableSqlConcurrencyControlRequest {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *EnableSqlConcurrencyControlRequest) SetConsoleContext(v string) *EnableSqlConcurrencyControlRequest {
	s.ConsoleContext = &v
	return s
}

func (s *EnableSqlConcurrencyControlRequest) SetInstanceId(v string) *EnableSqlConcurrencyControlRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableSqlConcurrencyControlRequest) SetMaxConcurrency(v int64) *EnableSqlConcurrencyControlRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *EnableSqlConcurrencyControlRequest) SetSqlKeywords(v string) *EnableSqlConcurrencyControlRequest {
	s.SqlKeywords = &v
	return s
}

func (s *EnableSqlConcurrencyControlRequest) SetSqlType(v string) *EnableSqlConcurrencyControlRequest {
	s.SqlType = &v
	return s
}

type EnableSqlConcurrencyControlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlConcurrencyControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlConcurrencyControlResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSqlConcurrencyControlResponseBody) SetCode(v string) *EnableSqlConcurrencyControlResponseBody {
	s.Code = &v
	return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetData(v string) *EnableSqlConcurrencyControlResponseBody {
	s.Data = &v
	return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetMessage(v string) *EnableSqlConcurrencyControlResponseBody {
	s.Message = &v
	return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetRequestId(v string) *EnableSqlConcurrencyControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetSuccess(v string) *EnableSqlConcurrencyControlResponseBody {
	s.Success = &v
	return s
}

type EnableSqlConcurrencyControlResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSqlConcurrencyControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSqlConcurrencyControlResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSqlConcurrencyControlResponse) GoString() string {
	return s.String()
}

func (s *EnableSqlConcurrencyControlResponse) SetHeaders(v map[string]*string) *EnableSqlConcurrencyControlResponse {
	s.Headers = v
	return s
}

func (s *EnableSqlConcurrencyControlResponse) SetBody(v *EnableSqlConcurrencyControlResponseBody) *EnableSqlConcurrencyControlResponse {
	s.Body = v
	return s
}

type GetAsyncErrorRequestListByCodeRequest struct {
	End        *int64  `json:"End,omitempty" xml:"End,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetAsyncErrorRequestListByCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetEnd(v int64) *GetAsyncErrorRequestListByCodeRequest {
	s.End = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetErrorCode(v string) *GetAsyncErrorRequestListByCodeRequest {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetInstanceId(v string) *GetAsyncErrorRequestListByCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetNodeId(v string) *GetAsyncErrorRequestListByCodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetStart(v int64) *GetAsyncErrorRequestListByCodeRequest {
	s.Start = &v
	return s
}

type GetAsyncErrorRequestListByCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsyncErrorRequestListByCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetCode(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetData(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Data = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetMessage(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetRequestId(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponseBody) SetSuccess(v string) *GetAsyncErrorRequestListByCodeResponseBody {
	s.Success = &v
	return s
}

type GetAsyncErrorRequestListByCodeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncErrorRequestListByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncErrorRequestListByCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeResponse) SetHeaders(v map[string]*string) *GetAsyncErrorRequestListByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponse) SetBody(v *GetAsyncErrorRequestListByCodeResponseBody) *GetAsyncErrorRequestListByCodeResponse {
	s.Body = v
	return s
}

type GetAsyncErrorRequestStatByCodeRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	End            *string `json:"End,omitempty" xml:"End,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Start          *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetConsoleContext(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetDbName(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.DbName = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetEnd(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.End = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetInstanceId(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetNodeId(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetRole(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.Role = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetStart(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.Start = &v
	return s
}

type GetAsyncErrorRequestStatByCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetCode(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetData(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Data = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetMessage(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetRequestId(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponseBody) SetSuccess(v string) *GetAsyncErrorRequestStatByCodeResponseBody {
	s.Success = &v
	return s
}

type GetAsyncErrorRequestStatByCodeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncErrorRequestStatByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncErrorRequestStatByCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeResponse) SetHeaders(v map[string]*string) *GetAsyncErrorRequestStatByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponse) SetBody(v *GetAsyncErrorRequestStatByCodeResponseBody) *GetAsyncErrorRequestStatByCodeResponse {
	s.Body = v
	return s
}

type GetAsyncErrorRequestStatResultRequest struct {
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	End        *int64  `json:"End,omitempty" xml:"End,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SqlIdList  *string `json:"SqlIdList,omitempty" xml:"SqlIdList,omitempty"`
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetAsyncErrorRequestStatResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultRequest) SetDbName(v string) *GetAsyncErrorRequestStatResultRequest {
	s.DbName = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetEnd(v int64) *GetAsyncErrorRequestStatResultRequest {
	s.End = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetInstanceId(v string) *GetAsyncErrorRequestStatResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetNodeId(v string) *GetAsyncErrorRequestStatResultRequest {
	s.NodeId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetSqlIdList(v string) *GetAsyncErrorRequestStatResultRequest {
	s.SqlIdList = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetStart(v int64) *GetAsyncErrorRequestStatResultRequest {
	s.Start = &v
	return s
}

type GetAsyncErrorRequestStatResultResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsyncErrorRequestStatResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetCode(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetData(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetMessage(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetRequestId(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponseBody) SetSuccess(v string) *GetAsyncErrorRequestStatResultResponseBody {
	s.Success = &v
	return s
}

type GetAsyncErrorRequestStatResultResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncErrorRequestStatResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncErrorRequestStatResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultResponse) SetHeaders(v map[string]*string) *GetAsyncErrorRequestStatResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponse) SetBody(v *GetAsyncErrorRequestStatResultResponseBody) *GetAsyncErrorRequestStatResultResponse {
	s.Body = v
	return s
}

type GetAutoResourceOptimizeConfigRequest struct {
	AccessKey  *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Signature  *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Uid        *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context    *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s GetAutoResourceOptimizeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoResourceOptimizeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeConfigRequest) SetAccessKey(v string) *GetAutoResourceOptimizeConfigRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigRequest) SetInstanceId(v string) *GetAutoResourceOptimizeConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigRequest) SetSignature(v string) *GetAutoResourceOptimizeConfigRequest {
	s.Signature = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigRequest) SetUid(v string) *GetAutoResourceOptimizeConfigRequest {
	s.Uid = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigRequest) SetUserId(v string) *GetAutoResourceOptimizeConfigRequest {
	s.UserId = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigRequest) SetContext(v string) *GetAutoResourceOptimizeConfigRequest {
	s.Context = &v
	return s
}

type GetAutoResourceOptimizeConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAutoResourceOptimizeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoResourceOptimizeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeConfigResponseBody) SetCode(v string) *GetAutoResourceOptimizeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigResponseBody) SetData(v string) *GetAutoResourceOptimizeConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigResponseBody) SetMessage(v string) *GetAutoResourceOptimizeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigResponseBody) SetRequestId(v string) *GetAutoResourceOptimizeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigResponseBody) SetSuccess(v string) *GetAutoResourceOptimizeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoResourceOptimizeConfigResponseBody) SetSynchro(v string) *GetAutoResourceOptimizeConfigResponseBody {
	s.Synchro = &v
	return s
}

type GetAutoResourceOptimizeConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAutoResourceOptimizeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutoResourceOptimizeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoResourceOptimizeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeConfigResponse) SetHeaders(v map[string]*string) *GetAutoResourceOptimizeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoResourceOptimizeConfigResponse) SetBody(v *GetAutoResourceOptimizeConfigResponseBody) *GetAutoResourceOptimizeConfigResponse {
	s.Body = v
	return s
}

type GetAutonomousNotifyEventContentRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SpanId     *string `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	Context    *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s GetAutonomousNotifyEventContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventContentRequest) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventContentRequest) SetInstanceId(v string) *GetAutonomousNotifyEventContentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutonomousNotifyEventContentRequest) SetSpanId(v string) *GetAutonomousNotifyEventContentRequest {
	s.SpanId = &v
	return s
}

func (s *GetAutonomousNotifyEventContentRequest) SetContext(v string) *GetAutonomousNotifyEventContentRequest {
	s.Context = &v
	return s
}

type GetAutonomousNotifyEventContentResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutonomousNotifyEventContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetCode(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetData(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Data = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetMessage(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetRequestId(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetSuccess(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Success = &v
	return s
}

type GetAutonomousNotifyEventContentResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAutonomousNotifyEventContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutonomousNotifyEventContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventContentResponse) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventContentResponse) SetHeaders(v map[string]*string) *GetAutonomousNotifyEventContentResponse {
	s.Headers = v
	return s
}

func (s *GetAutonomousNotifyEventContentResponse) SetBody(v *GetAutonomousNotifyEventContentResponseBody) *GetAutonomousNotifyEventContentResponse {
	s.Body = v
	return s
}

type GetAutonomousNotifyEventsInRangeRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventContext *string `json:"EventContext,omitempty" xml:"EventContext,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	MinLevel     *string `json:"MinLevel,omitempty" xml:"MinLevel,omitempty"`
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PageOffset   *string `json:"PageOffset,omitempty" xml:"PageOffset,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Context      *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeRequest) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetEndTime(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.EndTime = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetEventContext(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.EventContext = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetInstanceId(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetLevel(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.Level = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetMinLevel(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.MinLevel = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetNodeId(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.NodeId = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetPageOffset(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.PageOffset = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetPageSize(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.PageSize = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetStartTime(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.StartTime = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeRequest) SetContext(v string) *GetAutonomousNotifyEventsInRangeRequest {
	s.Context = &v
	return s
}

type GetAutonomousNotifyEventsInRangeResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAutonomousNotifyEventsInRangeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetCode(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetData(v *GetAutonomousNotifyEventsInRangeResponseBodyData) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Data = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetMessage(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetRequestId(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetSuccess(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Success = &v
	return s
}

type GetAutonomousNotifyEventsInRangeResponseBodyData struct {
	Extra    *string                                               `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     *GetAutonomousNotifyEventsInRangeResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	PageNo   *int64                                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetExtra(v string) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetList(v *GetAutonomousNotifyEventsInRangeResponseBodyDataList) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.List = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetPageNo(v int64) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetPageSize(v int64) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetTotal(v int64) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.Total = &v
	return s
}

type GetAutonomousNotifyEventsInRangeResponseBodyDataList struct {
	T []*string `json:"T,omitempty" xml:"T,omitempty" type:"Repeated"`
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyDataList) SetT(v []*string) *GetAutonomousNotifyEventsInRangeResponseBodyDataList {
	s.T = v
	return s
}

type GetAutonomousNotifyEventsInRangeResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAutonomousNotifyEventsInRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutonomousNotifyEventsInRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponse) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponse) SetHeaders(v map[string]*string) *GetAutonomousNotifyEventsInRangeResponse {
	s.Headers = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponse) SetBody(v *GetAutonomousNotifyEventsInRangeResponseBody) *GetAutonomousNotifyEventsInRangeResponse {
	s.Body = v
	return s
}

type GetDasProServiceUsageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDasProServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDasProServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageRequest) SetInstanceId(v string) *GetDasProServiceUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDasProServiceUsageRequest) SetUserId(v string) *GetDasProServiceUsageRequest {
	s.UserId = &v
	return s
}

type GetDasProServiceUsageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDasProServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDasProServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageResponseBody) SetCode(v string) *GetDasProServiceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetData(v string) *GetDasProServiceUsageResponseBody {
	s.Data = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetMessage(v string) *GetDasProServiceUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetRequestId(v string) *GetDasProServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetSuccess(v string) *GetDasProServiceUsageResponseBody {
	s.Success = &v
	return s
}

type GetDasProServiceUsageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDasProServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDasProServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDasProServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageResponse) SetHeaders(v map[string]*string) *GetDasProServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetDasProServiceUsageResponse) SetBody(v *GetDasProServiceUsageResponseBody) *GetDasProServiceUsageResponse {
	s.Body = v
	return s
}

type GetEndpointSwitchTaskRequest struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context   *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth  *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetEndpointSwitchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointSwitchTaskRequest) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskRequest) SetTaskId(v string) *GetEndpointSwitchTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetUid(v string) *GetEndpointSwitchTaskRequest {
	s.Uid = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetUserId(v string) *GetEndpointSwitchTaskRequest {
	s.UserId = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetContext(v string) *GetEndpointSwitchTaskRequest {
	s.Context = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetAccessKey(v string) *GetEndpointSwitchTaskRequest {
	s.AccessKey = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetSignature(v string) *GetEndpointSwitchTaskRequest {
	s.Signature = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetSkipAuth(v string) *GetEndpointSwitchTaskRequest {
	s.SkipAuth = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetTimestamp(v string) *GetEndpointSwitchTaskRequest {
	s.Timestamp = &v
	return s
}

type GetEndpointSwitchTaskResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEndpointSwitchTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                                `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetEndpointSwitchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointSwitchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskResponseBody) SetCode(v string) *GetEndpointSwitchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetData(v *GetEndpointSwitchTaskResponseBodyData) *GetEndpointSwitchTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetMessage(v string) *GetEndpointSwitchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetRequestId(v string) *GetEndpointSwitchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetSuccess(v string) *GetEndpointSwitchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetSynchro(v string) *GetEndpointSwitchTaskResponseBody {
	s.Synchro = &v
	return s
}

type GetEndpointSwitchTaskResponseBodyData struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DbLinkId  *int64  `json:"DbLinkId,omitempty" xml:"DbLinkId,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	OriUuid   *string `json:"OriUuid,omitempty" xml:"OriUuid,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uuid      *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetEndpointSwitchTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointSwitchTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetAccountId(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetDbLinkId(v int64) *GetEndpointSwitchTaskResponseBodyData {
	s.DbLinkId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetErrMsg(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetOriUuid(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.OriUuid = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetStatus(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetTaskId(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetUuid(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.Uuid = &v
	return s
}

type GetEndpointSwitchTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEndpointSwitchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEndpointSwitchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEndpointSwitchTaskResponse) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskResponse) SetHeaders(v map[string]*string) *GetEndpointSwitchTaskResponse {
	s.Headers = v
	return s
}

func (s *GetEndpointSwitchTaskResponse) SetBody(v *GetEndpointSwitchTaskResponseBody) *GetEndpointSwitchTaskResponse {
	s.Body = v
	return s
}

type GetErrorRequestSampleRequest struct {
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	End        *int64  `json:"End,omitempty" xml:"End,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SqlId      *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetErrorRequestSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErrorRequestSampleRequest) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleRequest) SetDbName(v string) *GetErrorRequestSampleRequest {
	s.DbName = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetEnd(v int64) *GetErrorRequestSampleRequest {
	s.End = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetInstanceId(v string) *GetErrorRequestSampleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetNodeId(v string) *GetErrorRequestSampleRequest {
	s.NodeId = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetSqlId(v string) *GetErrorRequestSampleRequest {
	s.SqlId = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetStart(v int64) *GetErrorRequestSampleRequest {
	s.Start = &v
	return s
}

type GetErrorRequestSampleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetErrorRequestSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErrorRequestSampleResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleResponseBody) SetCode(v string) *GetErrorRequestSampleResponseBody {
	s.Code = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetData(v string) *GetErrorRequestSampleResponseBody {
	s.Data = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetMessage(v string) *GetErrorRequestSampleResponseBody {
	s.Message = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetRequestId(v string) *GetErrorRequestSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetSuccess(v string) *GetErrorRequestSampleResponseBody {
	s.Success = &v
	return s
}

type GetErrorRequestSampleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetErrorRequestSampleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetErrorRequestSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErrorRequestSampleResponse) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleResponse) SetHeaders(v map[string]*string) *GetErrorRequestSampleResponse {
	s.Headers = v
	return s
}

func (s *GetErrorRequestSampleResponse) SetBody(v *GetErrorRequestSampleResponseBody) *GetErrorRequestSampleResponse {
	s.Body = v
	return s
}

type GetFullRequestOriginStatByInstanceIdRequest struct {
	Asc        *bool   `json:"Asc,omitempty" xml:"Asc,omitempty"`
	End        *int64  `json:"End,omitempty" xml:"End,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SqlType    *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetAsc(v bool) *GetFullRequestOriginStatByInstanceIdRequest {
	s.Asc = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetEnd(v int64) *GetFullRequestOriginStatByInstanceIdRequest {
	s.End = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetInstanceId(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetNodeId(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.NodeId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetOrderBy(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.OrderBy = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetPageNo(v int32) *GetFullRequestOriginStatByInstanceIdRequest {
	s.PageNo = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetPageSize(v int32) *GetFullRequestOriginStatByInstanceIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetRole(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.Role = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetSqlType(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.SqlType = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetStart(v int64) *GetFullRequestOriginStatByInstanceIdRequest {
	s.Start = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdRequest) SetUserId(v string) *GetFullRequestOriginStatByInstanceIdRequest {
	s.UserId = &v
	return s
}

type GetFullRequestOriginStatByInstanceIdResponseBody struct {
	Code      *int64                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFullRequestOriginStatByInstanceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetCode(v int64) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetData(v *GetFullRequestOriginStatByInstanceIdResponseBodyData) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetMessage(v string) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetRequestId(v string) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetSuccess(v bool) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Success = &v
	return s
}

type GetFullRequestOriginStatByInstanceIdResponseBodyData struct {
	List  []*GetFullRequestOriginStatByInstanceIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64                                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) SetList(v []*GetFullRequestOriginStatByInstanceIdResponseBodyDataList) *GetFullRequestOriginStatByInstanceIdResponseBodyData {
	s.List = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) SetTotal(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyData {
	s.Total = &v
	return s
}

type GetFullRequestOriginStatByInstanceIdResponseBodyDataList struct {
	AvgExaminedRows             *float64 `json:"AvgExaminedRows,omitempty" xml:"AvgExaminedRows,omitempty"`
	AvgFetchRows                *int64   `json:"AvgFetchRows,omitempty" xml:"AvgFetchRows,omitempty"`
	AvgLockWaitTime             *float64 `json:"AvgLockWaitTime,omitempty" xml:"AvgLockWaitTime,omitempty"`
	AvgLogicalRead              *float64 `json:"AvgLogicalRead,omitempty" xml:"AvgLogicalRead,omitempty"`
	AvgPhysicalAsyncRead        *int64   `json:"AvgPhysicalAsyncRead,omitempty" xml:"AvgPhysicalAsyncRead,omitempty"`
	AvgPhysicalSyncRead         *float64 `json:"AvgPhysicalSyncRead,omitempty" xml:"AvgPhysicalSyncRead,omitempty"`
	AvgReturnedRows             *float64 `json:"AvgReturnedRows,omitempty" xml:"AvgReturnedRows,omitempty"`
	AvgRows                     *int64   `json:"AvgRows,omitempty" xml:"AvgRows,omitempty"`
	AvgRt                       *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgSqlCount                 *int64   `json:"AvgSqlCount,omitempty" xml:"AvgSqlCount,omitempty"`
	AvgUpdatedRows              *float64 `json:"AvgUpdatedRows,omitempty" xml:"AvgUpdatedRows,omitempty"`
	Count                       *int64   `json:"Count,omitempty" xml:"Count,omitempty"`
	CountRate                   *float64 `json:"CountRate,omitempty" xml:"CountRate,omitempty"`
	Database                    *string  `json:"Database,omitempty" xml:"Database,omitempty"`
	ErrorCount                  *int64   `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	ExaminedRows                *int64   `json:"ExaminedRows,omitempty" xml:"ExaminedRows,omitempty"`
	FetchRows                   *int64   `json:"FetchRows,omitempty" xml:"FetchRows,omitempty"`
	Ip                          *string  `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Key                         *string  `json:"Key,omitempty" xml:"Key,omitempty"`
	LockWaitTime                *float64 `json:"LockWaitTime,omitempty" xml:"LockWaitTime,omitempty"`
	LogicalRead                 *int64   `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	OriginHost                  *string  `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	PhysicalAsyncRead           *int64   `json:"PhysicalAsyncRead,omitempty" xml:"PhysicalAsyncRead,omitempty"`
	PhysicalSyncRead            *int64   `json:"PhysicalSyncRead,omitempty" xml:"PhysicalSyncRead,omitempty"`
	Port                        *int64   `json:"Port,omitempty" xml:"Port,omitempty"`
	Rows                        *int64   `json:"Rows,omitempty" xml:"Rows,omitempty"`
	RtGreaterThanOneSecondCount *int64   `json:"RtGreaterThanOneSecondCount,omitempty" xml:"RtGreaterThanOneSecondCount,omitempty"`
	RtRate                      *float64 `json:"RtRate,omitempty" xml:"RtRate,omitempty"`
	SqlCount                    *int64   `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	SumUpdatedRows              *int64   `json:"SumUpdatedRows,omitempty" xml:"SumUpdatedRows,omitempty"`
	Version                     *int64   `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId                       *string  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgExaminedRows(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgExaminedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgFetchRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgFetchRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgLockWaitTime(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgLockWaitTime = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgLogicalRead(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgLogicalRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgPhysicalAsyncRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgPhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgPhysicalSyncRead(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgPhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgReturnedRows(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgReturnedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgRt(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgRt = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgSqlCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgSqlCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgUpdatedRows(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgUpdatedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetCountRate(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.CountRate = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetDatabase(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Database = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetErrorCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.ErrorCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetExaminedRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.ExaminedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetFetchRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.FetchRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetIp(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetKey(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Key = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetLockWaitTime(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.LockWaitTime = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetLogicalRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.LogicalRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetOriginHost(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.OriginHost = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetPhysicalAsyncRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.PhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetPhysicalSyncRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.PhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetPort(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Port = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Rows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetRtGreaterThanOneSecondCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.RtGreaterThanOneSecondCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetRtRate(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.RtRate = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetSqlCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.SqlCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetSumUpdatedRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.SumUpdatedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetVersion(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Version = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetVpcId(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.VpcId = &v
	return s
}

type GetFullRequestOriginStatByInstanceIdResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFullRequestOriginStatByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFullRequestOriginStatByInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) SetHeaders(v map[string]*string) *GetFullRequestOriginStatByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponse) SetBody(v *GetFullRequestOriginStatByInstanceIdResponseBody) *GetFullRequestOriginStatByInstanceIdResponse {
	s.Body = v
	return s
}

type GetFullRequestStatResultByInstanceIdRequest struct {
	Asc        *bool   `json:"Asc,omitempty" xml:"Asc,omitempty"`
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	End        *int64  `json:"End,omitempty" xml:"End,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SqlId      *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlType    *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetAsc(v bool) *GetFullRequestStatResultByInstanceIdRequest {
	s.Asc = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetDbName(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.DbName = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetEnd(v int64) *GetFullRequestStatResultByInstanceIdRequest {
	s.End = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetInstanceId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetKeyword(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.Keyword = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetNodeId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.NodeId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetOrderBy(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.OrderBy = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetOriginHost(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.OriginHost = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetPageNo(v int32) *GetFullRequestStatResultByInstanceIdRequest {
	s.PageNo = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetPageSize(v int32) *GetFullRequestStatResultByInstanceIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetRole(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.Role = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetSqlId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.SqlId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetSqlType(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.SqlType = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetStart(v int64) *GetFullRequestStatResultByInstanceIdRequest {
	s.Start = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdRequest) SetUserId(v string) *GetFullRequestStatResultByInstanceIdRequest {
	s.UserId = &v
	return s
}

type GetFullRequestStatResultByInstanceIdResponseBody struct {
	Code      *int64                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFullRequestStatResultByInstanceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetCode(v int64) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetData(v *GetFullRequestStatResultByInstanceIdResponseBodyData) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetMessage(v string) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetRequestId(v string) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetSuccess(v bool) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Success = &v
	return s
}

type GetFullRequestStatResultByInstanceIdResponseBodyData struct {
	Fail      *bool                                                       `json:"Fail,omitempty" xml:"Fail,omitempty"`
	IsFinish  *bool                                                       `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	Result    *GetFullRequestStatResultByInstanceIdResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	ResultId  *string                                                     `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	State     *string                                                     `json:"State,omitempty" xml:"State,omitempty"`
	Timestamp *int64                                                      `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetFail(v bool) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetIsFinish(v bool) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetResult(v *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.Result = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetResultId(v string) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetState(v string) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.State = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetTimestamp(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.Timestamp = &v
	return s
}

type GetFullRequestStatResultByInstanceIdResponseBodyDataResult struct {
	List  []*GetFullRequestStatResultByInstanceIdResponseBodyDataResultList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64                                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) SetList(v []*GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) *GetFullRequestStatResultByInstanceIdResponseBodyDataResult {
	s.List = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) SetTotal(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResult {
	s.Total = &v
	return s
}

type GetFullRequestStatResultByInstanceIdResponseBodyDataResultList struct {
	AvgExaminedRows             *float64  `json:"AvgExaminedRows,omitempty" xml:"AvgExaminedRows,omitempty"`
	AvgFetchRows                *int64    `json:"AvgFetchRows,omitempty" xml:"AvgFetchRows,omitempty"`
	AvgLockWaitTime             *float64  `json:"AvgLockWaitTime,omitempty" xml:"AvgLockWaitTime,omitempty"`
	AvgLogicalRead              *float64  `json:"AvgLogicalRead,omitempty" xml:"AvgLogicalRead,omitempty"`
	AvgPhysicalAsyncRead        *int64    `json:"AvgPhysicalAsyncRead,omitempty" xml:"AvgPhysicalAsyncRead,omitempty"`
	AvgPhysicalSyncRead         *int64    `json:"AvgPhysicalSyncRead,omitempty" xml:"AvgPhysicalSyncRead,omitempty"`
	AvgReturnedRows             *float64  `json:"AvgReturnedRows,omitempty" xml:"AvgReturnedRows,omitempty"`
	AvgRt                       *float64  `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgSqlCount                 *int64    `json:"AvgSqlCount,omitempty" xml:"AvgSqlCount,omitempty"`
	AvgUpdatedRows              *int64    `json:"AvgUpdatedRows,omitempty" xml:"AvgUpdatedRows,omitempty"`
	Count                       *int64    `json:"Count,omitempty" xml:"Count,omitempty"`
	CountRate                   *float64  `json:"CountRate,omitempty" xml:"CountRate,omitempty"`
	Database                    *string   `json:"Database,omitempty" xml:"Database,omitempty"`
	ErrorCount                  *int64    `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	ExaminedRows                *int64    `json:"ExaminedRows,omitempty" xml:"ExaminedRows,omitempty"`
	FetchRows                   *int64    `json:"FetchRows,omitempty" xml:"FetchRows,omitempty"`
	Ip                          *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	LockWaitTime                *float64  `json:"LockWaitTime,omitempty" xml:"LockWaitTime,omitempty"`
	LogicalRead                 *int64    `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	PhysicalAsyncRead           *int64    `json:"PhysicalAsyncRead,omitempty" xml:"PhysicalAsyncRead,omitempty"`
	PhysicalSyncRead            *int64    `json:"PhysicalSyncRead,omitempty" xml:"PhysicalSyncRead,omitempty"`
	Port                        *int64    `json:"Port,omitempty" xml:"Port,omitempty"`
	Psql                        *string   `json:"Psql,omitempty" xml:"Psql,omitempty"`
	Rows                        *int64    `json:"Rows,omitempty" xml:"Rows,omitempty"`
	RtGreaterThanOneSecondCount *int64    `json:"RtGreaterThanOneSecondCount,omitempty" xml:"RtGreaterThanOneSecondCount,omitempty"`
	RtRate                      *float64  `json:"RtRate,omitempty" xml:"RtRate,omitempty"`
	SqlCount                    *int64    `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	SqlId                       *string   `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SumUpdatedRows              *int64    `json:"SumUpdatedRows,omitempty" xml:"SumUpdatedRows,omitempty"`
	Tables                      []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	Version                     *int64    `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId                       *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgExaminedRows(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgExaminedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgFetchRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgFetchRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgLockWaitTime(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgLockWaitTime = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgLogicalRead(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgLogicalRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgPhysicalAsyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgPhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgPhysicalSyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgPhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgReturnedRows(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgReturnedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgRt(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgRt = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgSqlCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgSqlCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgUpdatedRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgUpdatedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Count = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetCountRate(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.CountRate = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetDatabase(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Database = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetErrorCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.ErrorCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetExaminedRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.ExaminedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetFetchRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.FetchRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetIp(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Ip = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetLockWaitTime(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.LockWaitTime = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetLogicalRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.LogicalRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPhysicalAsyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.PhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPhysicalSyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.PhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPort(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Port = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPsql(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Psql = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Rows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetRtGreaterThanOneSecondCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.RtGreaterThanOneSecondCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetRtRate(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.RtRate = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetSqlCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.SqlCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetSqlId(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.SqlId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetSumUpdatedRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.SumUpdatedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetTables(v []*string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Tables = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetVersion(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Version = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetVpcId(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.VpcId = &v
	return s
}

type GetFullRequestStatResultByInstanceIdResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFullRequestStatResultByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFullRequestStatResultByInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponse) SetHeaders(v map[string]*string) *GetFullRequestStatResultByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponse) SetBody(v *GetFullRequestStatResultByInstanceIdResponseBody) *GetFullRequestStatResultByInstanceIdResponse {
	s.Body = v
	return s
}

type GetHDMAliyunResourceSyncResultRequest struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context   *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth  *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetTaskId(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetUid(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Uid = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetUserId(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.UserId = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetContext(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Context = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetAccessKey(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.AccessKey = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetSignature(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Signature = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetSkipAuth(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.SkipAuth = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultRequest) SetTimestamp(v string) *GetHDMAliyunResourceSyncResultRequest {
	s.Timestamp = &v
	return s
}

type GetHDMAliyunResourceSyncResultResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHDMAliyunResourceSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                                         `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetCode(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetData(v *GetHDMAliyunResourceSyncResultResponseBodyData) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetMessage(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetRequestId(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetSuccess(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetSynchro(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Synchro = &v
	return s
}

type GetHDMAliyunResourceSyncResultResponseBodyData struct {
	ErrorMsg   *string                                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Results    *string                                                   `json:"Results,omitempty" xml:"Results,omitempty"`
	SubResults *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Struct"`
	SyncStatus *string                                                   `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetErrorMsg(v string) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetResults(v string) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.Results = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetSubResults(v *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.SubResults = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetSyncStatus(v string) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.SyncStatus = &v
	return s
}

type GetHDMAliyunResourceSyncResultResponseBodyDataSubResults struct {
	ResourceSyncSubResult []*GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult `json:"ResourceSyncSubResult,omitempty" xml:"ResourceSyncSubResult,omitempty" type:"Repeated"`
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) String() string {
	return tea.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) SetResourceSyncSubResult(v []*GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults {
	s.ResourceSyncSubResult = v
	return s
}

type GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult struct {
	ErrMsg       *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncCount    *int32  `json:"SyncCount,omitempty" xml:"SyncCount,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) String() string {
	return tea.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetErrMsg(v string) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ErrMsg = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetResourceType(v string) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ResourceType = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSuccess(v bool) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.Success = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSyncCount(v int32) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.SyncCount = &v
	return s
}

type GetHDMAliyunResourceSyncResultResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHDMAliyunResourceSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHDMAliyunResourceSyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponse) SetHeaders(v map[string]*string) *GetHDMAliyunResourceSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponse) SetBody(v *GetHDMAliyunResourceSyncResultResponseBody) *GetHDMAliyunResourceSyncResultResponse {
	s.Body = v
	return s
}

type GetHDMLastAliyunResourceSyncResultRequest struct {
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context   *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth  *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetUid(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Uid = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetUserId(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.UserId = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetContext(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Context = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetAccessKey(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.AccessKey = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetSignature(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Signature = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetSkipAuth(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.SkipAuth = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultRequest) SetTimestamp(v string) *GetHDMLastAliyunResourceSyncResultRequest {
	s.Timestamp = &v
	return s
}

type GetHDMLastAliyunResourceSyncResultResponseBody struct {
	Code      *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHDMLastAliyunResourceSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                                             `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetCode(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetData(v *GetHDMLastAliyunResourceSyncResultResponseBodyData) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetMessage(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetRequestId(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetSuccess(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetSynchro(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Synchro = &v
	return s
}

type GetHDMLastAliyunResourceSyncResultResponseBodyData struct {
	ErrorMsg   *string                                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Results    *string                                                       `json:"Results,omitempty" xml:"Results,omitempty"`
	SubResults *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Struct"`
	SyncStatus *string                                                       `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetErrorMsg(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetResults(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.Results = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetSubResults(v *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.SubResults = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetSyncStatus(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.SyncStatus = &v
	return s
}

type GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults struct {
	ResourceSyncSubResult []*GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult `json:"ResourceSyncSubResult,omitempty" xml:"ResourceSyncSubResult,omitempty" type:"Repeated"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) String() string {
	return tea.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) SetResourceSyncSubResult(v []*GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults {
	s.ResourceSyncSubResult = v
	return s
}

type GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult struct {
	ErrMsg       *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncCount    *int32  `json:"SyncCount,omitempty" xml:"SyncCount,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) String() string {
	return tea.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetErrMsg(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ErrMsg = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetResourceType(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ResourceType = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSuccess(v bool) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.Success = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSyncCount(v int32) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.SyncCount = &v
	return s
}

type GetHDMLastAliyunResourceSyncResultResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHDMLastAliyunResourceSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHDMLastAliyunResourceSyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) SetHeaders(v map[string]*string) *GetHDMLastAliyunResourceSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponse) SetBody(v *GetHDMLastAliyunResourceSyncResultResponseBody) *GetHDMLastAliyunResourceSyncResultResponse {
	s.Body = v
	return s
}

type GetInstanceInspectionsRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceArea *string `json:"InstanceArea,omitempty" xml:"InstanceArea,omitempty"`
	PageNo       *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchMap    *string `json:"SearchMap,omitempty" xml:"SearchMap,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInstanceInspectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInspectionsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsRequest) SetEndTime(v string) *GetInstanceInspectionsRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetEngine(v string) *GetInstanceInspectionsRequest {
	s.Engine = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetInstanceArea(v string) *GetInstanceInspectionsRequest {
	s.InstanceArea = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetPageNo(v string) *GetInstanceInspectionsRequest {
	s.PageNo = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetPageSize(v string) *GetInstanceInspectionsRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetSearchMap(v string) *GetInstanceInspectionsRequest {
	s.SearchMap = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetStartTime(v string) *GetInstanceInspectionsRequest {
	s.StartTime = &v
	return s
}

type GetInstanceInspectionsResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetInstanceInspectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceInspectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInspectionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBody) SetCode(v string) *GetInstanceInspectionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetData(v *GetInstanceInspectionsResponseBodyData) *GetInstanceInspectionsResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetMessage(v string) *GetInstanceInspectionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetRequestId(v string) *GetInstanceInspectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetSuccess(v string) *GetInstanceInspectionsResponseBody {
	s.Success = &v
	return s
}

type GetInstanceInspectionsResponseBodyData struct {
	List     []*GetInstanceInspectionsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int64                                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyData) SetList(v []*GetInstanceInspectionsResponseBodyDataList) *GetInstanceInspectionsResponseBodyData {
	s.List = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) SetPageNo(v int64) *GetInstanceInspectionsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) SetPageSize(v int64) *GetInstanceInspectionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) SetTotal(v int64) *GetInstanceInspectionsResponseBodyData {
	s.Total = &v
	return s
}

type GetInstanceInspectionsResponseBodyDataList struct {
	Data      map[string]interface{}                              `json:"Data,omitempty" xml:"Data,omitempty"`
	EndTime   *int64                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GmtCreate *int64                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Instance  *GetInstanceInspectionsResponseBodyDataListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	Score     *int32                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	ScoreMap  map[string]interface{}                              `json:"ScoreMap,omitempty" xml:"ScoreMap,omitempty"`
	StartTime *int64                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetData(v map[string]interface{}) *GetInstanceInspectionsResponseBodyDataList {
	s.Data = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetEndTime(v int64) *GetInstanceInspectionsResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetGmtCreate(v int64) *GetInstanceInspectionsResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetInstance(v *GetInstanceInspectionsResponseBodyDataListInstance) *GetInstanceInspectionsResponseBodyDataList {
	s.Instance = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetScore(v int32) *GetInstanceInspectionsResponseBodyDataList {
	s.Score = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetScoreMap(v map[string]interface{}) *GetInstanceInspectionsResponseBodyDataList {
	s.ScoreMap = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetStartTime(v int64) *GetInstanceInspectionsResponseBodyDataList {
	s.StartTime = &v
	return s
}

type GetInstanceInspectionsResponseBodyDataListInstance struct {
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceArea  *string `json:"InstanceArea,omitempty" xml:"InstanceArea,omitempty"`
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkType   *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyDataListInstance) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyDataListInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetAccountId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.AccountId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetEngine(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Engine = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetEngineVersion(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.EngineVersion = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceArea(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceArea = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceClass(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceClass = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetNetworkType(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.NetworkType = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetNodeId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.NodeId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetRegion(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Region = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetUuid(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Uuid = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetVpcId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.VpcId = &v
	return s
}

type GetInstanceInspectionsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceInspectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceInspectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInspectionsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponse) SetHeaders(v map[string]*string) *GetInstanceInspectionsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceInspectionsResponse) SetBody(v *GetInstanceInspectionsResponseBody) *GetInstanceInspectionsResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeDataStatsRequest struct {
	Asc              *string `json:"Asc,omitempty" xml:"Asc,omitempty"`
	DbNames          *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Keywords         *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	LogicalOperator  *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	OnlyOptimizedSql *string `json:"OnlyOptimizedSql,omitempty" xml:"OnlyOptimizedSql,omitempty"`
	OrderBy          *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNo           *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rules            *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	SqlIds           *string `json:"SqlIds,omitempty" xml:"SqlIds,omitempty"`
	TagNames         *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	Time             *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetQueryOptimizeDataStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataStatsRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsRequest) SetAsc(v string) *GetQueryOptimizeDataStatsRequest {
	s.Asc = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetDbNames(v string) *GetQueryOptimizeDataStatsRequest {
	s.DbNames = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetEngine(v string) *GetQueryOptimizeDataStatsRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetInstanceIds(v string) *GetQueryOptimizeDataStatsRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetKeywords(v string) *GetQueryOptimizeDataStatsRequest {
	s.Keywords = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetLogicalOperator(v string) *GetQueryOptimizeDataStatsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetOnlyOptimizedSql(v string) *GetQueryOptimizeDataStatsRequest {
	s.OnlyOptimizedSql = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetOrderBy(v string) *GetQueryOptimizeDataStatsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetPageNo(v string) *GetQueryOptimizeDataStatsRequest {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetPageSize(v string) *GetQueryOptimizeDataStatsRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetRules(v string) *GetQueryOptimizeDataStatsRequest {
	s.Rules = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetSqlIds(v string) *GetQueryOptimizeDataStatsRequest {
	s.SqlIds = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetTagNames(v string) *GetQueryOptimizeDataStatsRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeDataStatsRequest) SetTime(v string) *GetQueryOptimizeDataStatsRequest {
	s.Time = &v
	return s
}

type GetQueryOptimizeDataStatsResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeDataStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetCode(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetData(v *GetQueryOptimizeDataStatsResponseBodyData) *GetQueryOptimizeDataStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetMessage(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetRequestId(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetSuccess(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeDataStatsResponseBodyData struct {
	Extra    *string                                          `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeDataStatsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                           `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetExtra(v string) *GetQueryOptimizeDataStatsResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetList(v []*GetQueryOptimizeDataStatsResponseBodyDataList) *GetQueryOptimizeDataStatsResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeDataStatsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeDataStatsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetTotal(v int64) *GetQueryOptimizeDataStatsResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeDataStatsResponseBodyDataList struct {
	AvgLockTime     *float64                                                 `json:"AvgLockTime,omitempty" xml:"AvgLockTime,omitempty"`
	AvgQueryTime    *float64                                                 `json:"AvgQueryTime,omitempty" xml:"AvgQueryTime,omitempty"`
	AvgRowsAffected *float64                                                 `json:"AvgRowsAffected,omitempty" xml:"AvgRowsAffected,omitempty"`
	AvgRowsExamined *float64                                                 `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty"`
	AvgRowsSent     *float64                                                 `json:"AvgRowsSent,omitempty" xml:"AvgRowsSent,omitempty"`
	Count           *int32                                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	Dbname          *string                                                  `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	InstanceId      *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxLockTime     *float64                                                 `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty"`
	MaxQueryTime    *float64                                                 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	MaxRowsAffected *int64                                                   `json:"MaxRowsAffected,omitempty" xml:"MaxRowsAffected,omitempty"`
	MaxRowsExamined *int64                                                   `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty"`
	MaxRowsSent     *int64                                                   `json:"MaxRowsSent,omitempty" xml:"MaxRowsSent,omitempty"`
	Psql            *string                                                  `json:"Psql,omitempty" xml:"Psql,omitempty"`
	RuleList        []*GetQueryOptimizeDataStatsResponseBodyDataListRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	SqlId           *string                                                  `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlSample       *string                                                  `json:"SqlSample,omitempty" xml:"SqlSample,omitempty"`
	SqlType         *string                                                  `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgLockTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgLockTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgQueryTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgQueryTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgRowsAffected(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgRowsAffected = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgRowsExamined(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgRowsExamined = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgRowsSent(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgRowsSent = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetCount(v int32) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetDbname(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.Dbname = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetInstanceId(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxLockTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxLockTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxQueryTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxQueryTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxRowsAffected(v int64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxRowsAffected = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxRowsExamined(v int64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxRowsExamined = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxRowsSent(v int64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxRowsSent = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetPsql(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.Psql = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetRuleList(v []*GetQueryOptimizeDataStatsResponseBodyDataListRuleList) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.RuleList = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetSqlId(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetSqlSample(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.SqlSample = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetSqlType(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.SqlType = &v
	return s
}

type GetQueryOptimizeDataStatsResponseBodyDataListRuleList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBodyDataListRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBodyDataListRuleList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) SetName(v string) *GetQueryOptimizeDataStatsResponseBodyDataListRuleList {
	s.Name = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) SetType(v string) *GetQueryOptimizeDataStatsResponseBodyDataListRuleList {
	s.Type = &v
	return s
}

type GetQueryOptimizeDataStatsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeDataStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeDataStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeDataStatsResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponse) SetBody(v *GetQueryOptimizeDataStatsResponseBody) *GetQueryOptimizeDataStatsResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeDataTopRequest struct {
	Engine      *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	TagNames    *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
	Time        *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQueryOptimizeDataTopRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTopRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopRequest) SetEngine(v string) *GetQueryOptimizeDataTopRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetInstanceIds(v string) *GetQueryOptimizeDataTopRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetTagNames(v string) *GetQueryOptimizeDataTopRequest {
	s.TagNames = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetTime(v string) *GetQueryOptimizeDataTopRequest {
	s.Time = &v
	return s
}

func (s *GetQueryOptimizeDataTopRequest) SetType(v string) *GetQueryOptimizeDataTopRequest {
	s.Type = &v
	return s
}

type GetQueryOptimizeDataTopResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeDataTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeDataTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponseBody) SetCode(v string) *GetQueryOptimizeDataTopResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetData(v *GetQueryOptimizeDataTopResponseBodyData) *GetQueryOptimizeDataTopResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetMessage(v string) *GetQueryOptimizeDataTopResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetRequestId(v string) *GetQueryOptimizeDataTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBody) SetSuccess(v string) *GetQueryOptimizeDataTopResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeDataTopResponseBodyData struct {
	Extra    *string                                        `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeDataTopResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeDataTopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetExtra(v string) *GetQueryOptimizeDataTopResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetList(v []*GetQueryOptimizeDataTopResponseBodyDataList) *GetQueryOptimizeDataTopResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeDataTopResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeDataTopResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyData) SetTotal(v int64) *GetQueryOptimizeDataTopResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeDataTopResponseBodyDataList struct {
	InstanceId *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Type       *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value      *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQueryOptimizeDataTopResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) SetInstanceId(v string) *GetQueryOptimizeDataTopResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) SetType(v string) *GetQueryOptimizeDataTopResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponseBodyDataList) SetValue(v float64) *GetQueryOptimizeDataTopResponseBodyDataList {
	s.Value = &v
	return s
}

type GetQueryOptimizeDataTopResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeDataTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeDataTopResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeDataTopResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeDataTopResponse) SetBody(v *GetQueryOptimizeDataTopResponseBody) *GetQueryOptimizeDataTopResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeDataTrendRequest struct {
	End         *string `json:"End,omitempty" xml:"End,omitempty"`
	Engine      *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Start       *string `json:"Start,omitempty" xml:"Start,omitempty"`
	TagNames    *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
}

func (s GetQueryOptimizeDataTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTrendRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendRequest) SetEnd(v string) *GetQueryOptimizeDataTrendRequest {
	s.End = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetEngine(v string) *GetQueryOptimizeDataTrendRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetInstanceIds(v string) *GetQueryOptimizeDataTrendRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetStart(v string) *GetQueryOptimizeDataTrendRequest {
	s.Start = &v
	return s
}

func (s *GetQueryOptimizeDataTrendRequest) SetTagNames(v string) *GetQueryOptimizeDataTrendRequest {
	s.TagNames = &v
	return s
}

type GetQueryOptimizeDataTrendResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeDataTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetCode(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetData(v *GetQueryOptimizeDataTrendResponseBodyData) *GetQueryOptimizeDataTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetMessage(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetRequestId(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBody) SetSuccess(v string) *GetQueryOptimizeDataTrendResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeDataTrendResponseBodyData struct {
	Extra    *string                                          `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeDataTrendResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                           `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetExtra(v string) *GetQueryOptimizeDataTrendResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetList(v []*GetQueryOptimizeDataTrendResponseBodyDataList) *GetQueryOptimizeDataTrendResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeDataTrendResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeDataTrendResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyData) SetTotal(v int64) *GetQueryOptimizeDataTrendResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeDataTrendResponseBodyDataList struct {
	Kpi       *string  `json:"Kpi,omitempty" xml:"Kpi,omitempty"`
	Timestamp *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) SetKpi(v string) *GetQueryOptimizeDataTrendResponseBodyDataList {
	s.Kpi = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) SetTimestamp(v int64) *GetQueryOptimizeDataTrendResponseBodyDataList {
	s.Timestamp = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponseBodyDataList) SetValue(v float64) *GetQueryOptimizeDataTrendResponseBodyDataList {
	s.Value = &v
	return s
}

type GetQueryOptimizeDataTrendResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeDataTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeDataTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeDataTrendResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponse) SetBody(v *GetQueryOptimizeDataTrendResponseBody) *GetQueryOptimizeDataTrendResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeExecErrorSampleRequest struct {
	Engine     *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SqlId      *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetEngine(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetInstanceId(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetSqlId(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetTime(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.Time = &v
	return s
}

type GetQueryOptimizeExecErrorSampleResponseBody struct {
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeExecErrorSampleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetCode(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetData(v *GetQueryOptimizeExecErrorSampleResponseBodyData) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetMessage(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetRequestId(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetSuccess(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeExecErrorSampleResponseBodyData struct {
	Extra    *string                                                `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeExecErrorSampleResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetExtra(v string) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetList(v []*GetQueryOptimizeExecErrorSampleResponseBodyDataList) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetTotal(v int64) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeExecErrorSampleResponseBodyDataList struct {
	Dbname    *string `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OrigHost  *string `json:"OrigHost,omitempty" xml:"OrigHost,omitempty"`
	SqlId     *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlText   *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetDbname(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.Dbname = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetErrorCode(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.ErrorCode = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetOrigHost(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.OrigHost = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetSqlId(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetSqlText(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.SqlText = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetTimestamp(v int64) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.Timestamp = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetUser(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.User = &v
	return s
}

type GetQueryOptimizeExecErrorSampleResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeExecErrorSampleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeExecErrorSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeExecErrorSampleResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponse) SetBody(v *GetQueryOptimizeExecErrorSampleResponseBody) *GetQueryOptimizeExecErrorSampleResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeExecErrorStatsRequest struct {
	Asc             *string `json:"Asc,omitempty" xml:"Asc,omitempty"`
	DbNames         *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	Engine          *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceIds     *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Keywords        *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	OrderBy         *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNo          *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Time            *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetAsc(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Asc = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetDbNames(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.DbNames = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetEngine(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetInstanceIds(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetKeywords(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Keywords = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetLogicalOperator(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetOrderBy(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetPageNo(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetPageSize(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsRequest) SetTime(v string) *GetQueryOptimizeExecErrorStatsRequest {
	s.Time = &v
	return s
}

type GetQueryOptimizeExecErrorStatsResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeExecErrorStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetCode(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetData(v *GetQueryOptimizeExecErrorStatsResponseBodyData) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetMessage(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetRequestId(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetSuccess(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeExecErrorStatsResponseBodyData struct {
	Extra    *string                                               `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeExecErrorStatsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetExtra(v string) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetList(v []*GetQueryOptimizeExecErrorStatsResponseBodyDataList) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetTotal(v int64) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeExecErrorStatsResponseBodyDataList struct {
	Dbname       *string `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorCount   *int64  `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	SqlId        *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlText      *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetDbname(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.Dbname = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetErrorCode(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.ErrorCode = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetErrorCount(v int64) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.ErrorCount = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetInstanceId(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetInstanceName(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetSqlId(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetSqlText(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.SqlText = &v
	return s
}

type GetQueryOptimizeExecErrorStatsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeExecErrorStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeExecErrorStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeExecErrorStatsResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponse) SetBody(v *GetQueryOptimizeExecErrorStatsResponseBody) *GetQueryOptimizeExecErrorStatsResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeRuleListRequest struct {
	Engine      *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	TagNames    *string `json:"TagNames,omitempty" xml:"TagNames,omitempty"`
}

func (s GetQueryOptimizeRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeRuleListRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListRequest) SetEngine(v string) *GetQueryOptimizeRuleListRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeRuleListRequest) SetInstanceIds(v string) *GetQueryOptimizeRuleListRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetQueryOptimizeRuleListRequest) SetTagNames(v string) *GetQueryOptimizeRuleListRequest {
	s.TagNames = &v
	return s
}

type GetQueryOptimizeRuleListResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponseBody) SetCode(v string) *GetQueryOptimizeRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetData(v *GetQueryOptimizeRuleListResponseBodyData) *GetQueryOptimizeRuleListResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetMessage(v string) *GetQueryOptimizeRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetRequestId(v string) *GetQueryOptimizeRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetSuccess(v string) *GetQueryOptimizeRuleListResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeRuleListResponseBodyData struct {
	Extra    *string                                         `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeRuleListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeRuleListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetExtra(v string) *GetQueryOptimizeRuleListResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetList(v []*GetQueryOptimizeRuleListResponseBodyDataList) *GetQueryOptimizeRuleListResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeRuleListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeRuleListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetTotal(v int64) *GetQueryOptimizeRuleListResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeRuleListResponseBodyDataList struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQueryOptimizeRuleListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) SetName(v string) *GetQueryOptimizeRuleListResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) SetRuleId(v string) *GetQueryOptimizeRuleListResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) SetType(v string) *GetQueryOptimizeRuleListResponseBodyDataList {
	s.Type = &v
	return s
}

type GetQueryOptimizeRuleListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeRuleListResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeRuleListResponse) SetBody(v *GetQueryOptimizeRuleListResponseBody) *GetQueryOptimizeRuleListResponse {
	s.Body = v
	return s
}

type GetQueryOptimizeSolutionRequest struct {
	Engine  *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	SqlId   *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s GetQueryOptimizeSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeSolutionRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionRequest) SetEngine(v string) *GetQueryOptimizeSolutionRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeSolutionRequest) SetRuleIds(v string) *GetQueryOptimizeSolutionRequest {
	s.RuleIds = &v
	return s
}

func (s *GetQueryOptimizeSolutionRequest) SetSqlId(v string) *GetQueryOptimizeSolutionRequest {
	s.SqlId = &v
	return s
}

type GetQueryOptimizeSolutionResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueryOptimizeSolutionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponseBody) SetCode(v string) *GetQueryOptimizeSolutionResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetData(v *GetQueryOptimizeSolutionResponseBodyData) *GetQueryOptimizeSolutionResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetMessage(v string) *GetQueryOptimizeSolutionResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetRequestId(v string) *GetQueryOptimizeSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBody) SetSuccess(v string) *GetQueryOptimizeSolutionResponseBody {
	s.Success = &v
	return s
}

type GetQueryOptimizeSolutionResponseBodyData struct {
	Extra    *string                                         `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List     []*GetQueryOptimizeSolutionResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeSolutionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetExtra(v string) *GetQueryOptimizeSolutionResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetList(v []*GetQueryOptimizeSolutionResponseBodyDataList) *GetQueryOptimizeSolutionResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeSolutionResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeSolutionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyData) SetTotal(v int64) *GetQueryOptimizeSolutionResponseBodyData {
	s.Total = &v
	return s
}

type GetQueryOptimizeSolutionResponseBodyDataList struct {
	Level       *string `json:"Level,omitempty" xml:"Level,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Solution    *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	SolutionExt *string `json:"SolutionExt,omitempty" xml:"SolutionExt,omitempty"`
}

func (s GetQueryOptimizeSolutionResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetLevel(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.Level = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetRuleId(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetSolution(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.Solution = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponseBodyDataList) SetSolutionExt(v string) *GetQueryOptimizeSolutionResponseBodyDataList {
	s.SolutionExt = &v
	return s
}

type GetQueryOptimizeSolutionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQueryOptimizeSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryOptimizeSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeSolutionResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeSolutionResponse) SetBody(v *GetQueryOptimizeSolutionResponseBody) *GetQueryOptimizeSolutionResponse {
	s.Body = v
	return s
}

type GetRequestDiagnosisPageRequest struct {
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetRequestDiagnosisPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisPageRequest) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageRequest) SetEndTime(v int64) *GetRequestDiagnosisPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetInstanceId(v string) *GetRequestDiagnosisPageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetNodeId(v string) *GetRequestDiagnosisPageRequest {
	s.NodeId = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetPageNo(v int32) *GetRequestDiagnosisPageRequest {
	s.PageNo = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetPageSize(v int32) *GetRequestDiagnosisPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetStartTime(v int64) *GetRequestDiagnosisPageRequest {
	s.StartTime = &v
	return s
}

type GetRequestDiagnosisPageResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRequestDiagnosisPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRequestDiagnosisPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponseBody) SetCode(v string) *GetRequestDiagnosisPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetData(v *GetRequestDiagnosisPageResponseBodyData) *GetRequestDiagnosisPageResponseBody {
	s.Data = v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetMessage(v string) *GetRequestDiagnosisPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetRequestId(v string) *GetRequestDiagnosisPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetSuccess(v string) *GetRequestDiagnosisPageResponseBody {
	s.Success = &v
	return s
}

type GetRequestDiagnosisPageResponseBodyData struct {
	Extra    *string                                        `json:"extra,omitempty" xml:"extra,omitempty"`
	List     []*GetRequestDiagnosisPageResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNo   *int64                                         `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int64                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int64                                         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetRequestDiagnosisPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetExtra(v string) *GetRequestDiagnosisPageResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetList(v []*GetRequestDiagnosisPageResponseBodyDataList) *GetRequestDiagnosisPageResponseBodyData {
	s.List = v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetPageNo(v int64) *GetRequestDiagnosisPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetPageSize(v int64) *GetRequestDiagnosisPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetTotal(v int64) *GetRequestDiagnosisPageResponseBodyData {
	s.Total = &v
	return s
}

type GetRequestDiagnosisPageResponseBodyDataList struct {
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DbSchema    *string `json:"dbSchema,omitempty" xml:"dbSchema,omitempty"`
	Engine      *string `json:"engine,omitempty" xml:"engine,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	MessageId   *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	Param       *string `json:"param,omitempty" xml:"param,omitempty"`
	Result      *string `json:"result,omitempty" xml:"result,omitempty"`
	SqlId       *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	State       *int32  `json:"state,omitempty" xml:"state,omitempty"`
	Uuid        *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s GetRequestDiagnosisPageResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetAccountId(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.AccountId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetDbSchema(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.DbSchema = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetEngine(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Engine = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetGmtCreate(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetGmtModified(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetMessageId(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.MessageId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetParam(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Param = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetResult(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Result = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetSqlId(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetState(v int32) *GetRequestDiagnosisPageResponseBodyDataList {
	s.State = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetUuid(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Uuid = &v
	return s
}

type GetRequestDiagnosisPageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRequestDiagnosisPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRequestDiagnosisPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisPageResponse) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponse) SetHeaders(v map[string]*string) *GetRequestDiagnosisPageResponse {
	s.Headers = v
	return s
}

func (s *GetRequestDiagnosisPageResponse) SetBody(v *GetRequestDiagnosisPageResponseBody) *GetRequestDiagnosisPageResponse {
	s.Body = v
	return s
}

type GetRequestDiagnosisResultRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageId  *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s GetRequestDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultRequest) SetInstanceId(v string) *GetRequestDiagnosisResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) SetMessageId(v string) *GetRequestDiagnosisResultRequest {
	s.MessageId = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) SetNodeId(v string) *GetRequestDiagnosisResultRequest {
	s.NodeId = &v
	return s
}

type GetRequestDiagnosisResultResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRequestDiagnosisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRequestDiagnosisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultResponseBody) SetCode(v string) *GetRequestDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetData(v *GetRequestDiagnosisResultResponseBodyData) *GetRequestDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetMessage(v string) *GetRequestDiagnosisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetRequestId(v string) *GetRequestDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetSuccess(v string) *GetRequestDiagnosisResultResponseBody {
	s.Success = &v
	return s
}

type GetRequestDiagnosisResultResponseBodyData struct {
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DbSchema    *string `json:"dbSchema,omitempty" xml:"dbSchema,omitempty"`
	Engine      *string `json:"engine,omitempty" xml:"engine,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	MessageId   *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	Param       *string `json:"param,omitempty" xml:"param,omitempty"`
	Result      *string `json:"result,omitempty" xml:"result,omitempty"`
	SqlId       *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	State       *int32  `json:"state,omitempty" xml:"state,omitempty"`
	Uuid        *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s GetRequestDiagnosisResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetAccountId(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetDbSchema(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.DbSchema = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetEngine(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Engine = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetGmtCreate(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetGmtModified(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetMessageId(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetParam(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Param = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetResult(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetSqlId(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetState(v int32) *GetRequestDiagnosisResultResponseBodyData {
	s.State = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetUuid(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Uuid = &v
	return s
}

type GetRequestDiagnosisResultResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRequestDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRequestDiagnosisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRequestDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetRequestDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetRequestDiagnosisResultResponse) SetBody(v *GetRequestDiagnosisResultResponseBody) *GetRequestDiagnosisResultResponse {
	s.Body = v
	return s
}

type GetResourceOptimizeHistoryListRequest struct {
	AccessKey  *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Page       *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Signature  *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Uid        *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context    *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s GetResourceOptimizeHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOptimizeHistoryListRequest) GoString() string {
	return s.String()
}

func (s *GetResourceOptimizeHistoryListRequest) SetAccessKey(v string) *GetResourceOptimizeHistoryListRequest {
	s.AccessKey = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetEndTime(v string) *GetResourceOptimizeHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetInstanceId(v string) *GetResourceOptimizeHistoryListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetPage(v int32) *GetResourceOptimizeHistoryListRequest {
	s.Page = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetPageSize(v int32) *GetResourceOptimizeHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetSignature(v string) *GetResourceOptimizeHistoryListRequest {
	s.Signature = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetStartTime(v string) *GetResourceOptimizeHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetUid(v string) *GetResourceOptimizeHistoryListRequest {
	s.Uid = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetUserId(v string) *GetResourceOptimizeHistoryListRequest {
	s.UserId = &v
	return s
}

func (s *GetResourceOptimizeHistoryListRequest) SetContext(v string) *GetResourceOptimizeHistoryListRequest {
	s.Context = &v
	return s
}

type GetResourceOptimizeHistoryListResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetResourceOptimizeHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOptimizeHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceOptimizeHistoryListResponseBody) SetCode(v string) *GetResourceOptimizeHistoryListResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceOptimizeHistoryListResponseBody) SetData(v string) *GetResourceOptimizeHistoryListResponseBody {
	s.Data = &v
	return s
}

func (s *GetResourceOptimizeHistoryListResponseBody) SetMessage(v string) *GetResourceOptimizeHistoryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceOptimizeHistoryListResponseBody) SetRequestId(v string) *GetResourceOptimizeHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceOptimizeHistoryListResponseBody) SetSuccess(v string) *GetResourceOptimizeHistoryListResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourceOptimizeHistoryListResponseBody) SetSynchro(v string) *GetResourceOptimizeHistoryListResponseBody {
	s.Synchro = &v
	return s
}

type GetResourceOptimizeHistoryListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResourceOptimizeHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceOptimizeHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOptimizeHistoryListResponse) GoString() string {
	return s.String()
}

func (s *GetResourceOptimizeHistoryListResponse) SetHeaders(v map[string]*string) *GetResourceOptimizeHistoryListResponse {
	s.Headers = v
	return s
}

func (s *GetResourceOptimizeHistoryListResponse) SetBody(v *GetResourceOptimizeHistoryListResponseBody) *GetResourceOptimizeHistoryListResponse {
	s.Body = v
	return s
}

type GetRunningSqlConcurrencyControlRulesRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo         *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesRequest) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetConsoleContext(v string) *GetRunningSqlConcurrencyControlRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetInstanceId(v string) *GetRunningSqlConcurrencyControlRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetPageNo(v int64) *GetRunningSqlConcurrencyControlRulesRequest {
	s.PageNo = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesRequest) SetPageSize(v int64) *GetRunningSqlConcurrencyControlRulesRequest {
	s.PageSize = &v
	return s
}

type GetRunningSqlConcurrencyControlRulesResponseBody struct {
	Code      *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRunningSqlConcurrencyControlRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetCode(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetData(v *GetRunningSqlConcurrencyControlRulesResponseBodyData) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Data = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetMessage(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetRequestId(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetSuccess(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Success = &v
	return s
}

type GetRunningSqlConcurrencyControlRulesResponseBodyData struct {
	List  *GetRunningSqlConcurrencyControlRulesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	Total *int64                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) SetList(v *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	s.List = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) SetTotal(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	s.Total = &v
	return s
}

type GetRunningSqlConcurrencyControlRulesResponseBodyDataList struct {
	RunningRules []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules `json:"runningRules,omitempty" xml:"runningRules,omitempty" type:"Repeated"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) SetRunningRules(v []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) *GetRunningSqlConcurrencyControlRulesResponseBodyDataList {
	s.RunningRules = v
	return s
}

type GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules struct {
	ConcurrencyControlTime *int64  `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId                 *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	KeywordsHash           *string `json:"KeywordsHash,omitempty" xml:"KeywordsHash,omitempty"`
	MaxConcurrency         *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	SqlKeywords            *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
	SqlType                *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) String() string {
	return tea.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetConcurrencyControlTime(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetInstanceId(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.InstanceId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetItemId(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.ItemId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetKeywordsHash(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.KeywordsHash = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetMaxConcurrency(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.MaxConcurrency = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetSqlKeywords(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.SqlKeywords = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetSqlType(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.SqlType = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetStartTime(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.StartTime = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetStatus(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.Status = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetUserId(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.UserId = &v
	return s
}

type GetRunningSqlConcurrencyControlRulesResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRunningSqlConcurrencyControlRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRunningSqlConcurrencyControlRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponse) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) SetHeaders(v map[string]*string) *GetRunningSqlConcurrencyControlRulesResponse {
	s.Headers = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) SetBody(v *GetRunningSqlConcurrencyControlRulesResponseBody) *GetRunningSqlConcurrencyControlRulesResponse {
	s.Body = v
	return s
}

type GetSqlAnalysisReportRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	ModuleNameList *string `json:"ModuleNameList,omitempty" xml:"ModuleNameList,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSqlAnalysisReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSqlAnalysisReportRequest) GoString() string {
	return s.String()
}

func (s *GetSqlAnalysisReportRequest) SetConsoleContext(v string) *GetSqlAnalysisReportRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlAnalysisReportRequest) SetModuleNameList(v string) *GetSqlAnalysisReportRequest {
	s.ModuleNameList = &v
	return s
}

func (s *GetSqlAnalysisReportRequest) SetTaskId(v string) *GetSqlAnalysisReportRequest {
	s.TaskId = &v
	return s
}

type GetSqlAnalysisReportResponseBody struct {
	Code      *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetSqlAnalysisReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSqlAnalysisReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlAnalysisReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlAnalysisReportResponseBody) SetCode(v int64) *GetSqlAnalysisReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBody) SetData(v []*GetSqlAnalysisReportResponseBodyData) *GetSqlAnalysisReportResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlAnalysisReportResponseBody) SetMessage(v string) *GetSqlAnalysisReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBody) SetRequestId(v string) *GetSqlAnalysisReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBody) SetSuccess(v bool) *GetSqlAnalysisReportResponseBody {
	s.Success = &v
	return s
}

type GetSqlAnalysisReportResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSqlAnalysisReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSqlAnalysisReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlAnalysisReportResponseBodyData) SetInstanceId(v string) *GetSqlAnalysisReportResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBodyData) SetModuleName(v string) *GetSqlAnalysisReportResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBodyData) SetResult(v string) *GetSqlAnalysisReportResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBodyData) SetTaskId(v string) *GetSqlAnalysisReportResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetSqlAnalysisReportResponseBodyData) SetUserId(v string) *GetSqlAnalysisReportResponseBodyData {
	s.UserId = &v
	return s
}

type GetSqlAnalysisReportResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSqlAnalysisReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSqlAnalysisReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlAnalysisReportResponse) GoString() string {
	return s.String()
}

func (s *GetSqlAnalysisReportResponse) SetHeaders(v map[string]*string) *GetSqlAnalysisReportResponse {
	s.Headers = v
	return s
}

func (s *GetSqlAnalysisReportResponse) SetBody(v *GetSqlAnalysisReportResponseBody) *GetSqlAnalysisReportResponse {
	s.Body = v
	return s
}

type GetSqlConcurrencyControlKeywordsFromSqlTextRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SqlText        *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) SetConsoleContext(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) SetInstanceId(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) SetSqlText(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest {
	s.SqlText = &v
	return s
}

type GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetCode(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetData(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Data = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetMessage(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetRequestId(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetSuccess(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Success = &v
	return s
}

type GetSqlConcurrencyControlKeywordsFromSqlTextResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) SetHeaders(v map[string]*string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) SetBody(v *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse {
	s.Body = v
	return s
}

type GetSqlConcurrencyControlRulesHistoryRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo         *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetConsoleContext(v string) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetInstanceId(v string) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetPageNo(v int64) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.PageNo = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryRequest) SetPageSize(v int64) *GetSqlConcurrencyControlRulesHistoryRequest {
	s.PageSize = &v
	return s
}

type GetSqlConcurrencyControlRulesHistoryResponseBody struct {
	Code      *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSqlConcurrencyControlRulesHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetCode(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetData(v *GetSqlConcurrencyControlRulesHistoryResponseBodyData) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetMessage(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetRequestId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetSuccess(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Success = &v
	return s
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyData struct {
	List  *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	Total *int64                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) SetList(v *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	s.List = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) SetTotal(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	s.Total = &v
	return s
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyDataList struct {
	Rules []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) SetRules(v []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList {
	s.Rules = v
	return s
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules struct {
	ConcurrencyControlTime *int64  `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId                 *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	KeywordsHash           *string `json:"KeywordsHash,omitempty" xml:"KeywordsHash,omitempty"`
	MaxConcurrency         *int64  `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	SqlKeywords            *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
	SqlType                *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetConcurrencyControlTime(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetInstanceId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetItemId(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.ItemId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetKeywordsHash(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.KeywordsHash = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetMaxConcurrency(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.MaxConcurrency = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetSqlKeywords(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.SqlKeywords = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetSqlType(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.SqlType = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetStartTime(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.StartTime = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetStatus(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.Status = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetUserId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.UserId = &v
	return s
}

type GetSqlConcurrencyControlRulesHistoryResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSqlConcurrencyControlRulesHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) SetHeaders(v map[string]*string) *GetSqlConcurrencyControlRulesHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) SetBody(v *GetSqlConcurrencyControlRulesHistoryResponseBody) *GetSqlConcurrencyControlRulesHistoryResponse {
	s.Body = v
	return s
}

type GetSqlOptimizeAdviceRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	EndDt          *string `json:"EndDt,omitempty" xml:"EndDt,omitempty"`
	Engine         *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceIds    *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	StartDt        *string `json:"StartDt,omitempty" xml:"StartDt,omitempty"`
}

func (s GetSqlOptimizeAdviceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSqlOptimizeAdviceRequest) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceRequest) SetConsoleContext(v string) *GetSqlOptimizeAdviceRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetEndDt(v string) *GetSqlOptimizeAdviceRequest {
	s.EndDt = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetEngine(v string) *GetSqlOptimizeAdviceRequest {
	s.Engine = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetInstanceIds(v string) *GetSqlOptimizeAdviceRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetSqlOptimizeAdviceRequest) SetStartDt(v string) *GetSqlOptimizeAdviceRequest {
	s.StartDt = &v
	return s
}

type GetSqlOptimizeAdviceResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSqlOptimizeAdviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSqlOptimizeAdviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlOptimizeAdviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceResponseBody) SetCode(v string) *GetSqlOptimizeAdviceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetData(v *GetSqlOptimizeAdviceResponseBodyData) *GetSqlOptimizeAdviceResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetMessage(v string) *GetSqlOptimizeAdviceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetRequestId(v string) *GetSqlOptimizeAdviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetSuccess(v string) *GetSqlOptimizeAdviceResponseBody {
	s.Success = &v
	return s
}

type GetSqlOptimizeAdviceResponseBodyData struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusCode  *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSqlOptimizeAdviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSqlOptimizeAdviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetCreateTime(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetDownloadUrl(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetExpireTime(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetStatus(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetStatusCode(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetTaskId(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.TaskId = &v
	return s
}

type GetSqlOptimizeAdviceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSqlOptimizeAdviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSqlOptimizeAdviceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlOptimizeAdviceResponse) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceResponse) SetHeaders(v map[string]*string) *GetSqlOptimizeAdviceResponse {
	s.Headers = v
	return s
}

func (s *GetSqlOptimizeAdviceResponse) SetBody(v *GetSqlOptimizeAdviceResponseBody) *GetSqlOptimizeAdviceResponse {
	s.Body = v
	return s
}

type QuerySqlAnalysisDataRequest struct {
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	PageNo         *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SqlIdList      *string `json:"SqlIdList,omitempty" xml:"SqlIdList,omitempty"`
	SqlTextFeature *string `json:"SqlTextFeature,omitempty" xml:"SqlTextFeature,omitempty"`
	SqlType        *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuerySqlAnalysisDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySqlAnalysisDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySqlAnalysisDataRequest) SetConsoleContext(v string) *QuerySqlAnalysisDataRequest {
	s.ConsoleContext = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetPageNo(v int32) *QuerySqlAnalysisDataRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetPageSize(v int32) *QuerySqlAnalysisDataRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetSqlIdList(v string) *QuerySqlAnalysisDataRequest {
	s.SqlIdList = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetSqlTextFeature(v string) *QuerySqlAnalysisDataRequest {
	s.SqlTextFeature = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetSqlType(v string) *QuerySqlAnalysisDataRequest {
	s.SqlType = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetTaskId(v string) *QuerySqlAnalysisDataRequest {
	s.TaskId = &v
	return s
}

func (s *QuerySqlAnalysisDataRequest) SetType(v string) *QuerySqlAnalysisDataRequest {
	s.Type = &v
	return s
}

type QuerySqlAnalysisDataResponseBody struct {
	Code      *int64                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QuerySqlAnalysisDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySqlAnalysisDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySqlAnalysisDataResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySqlAnalysisDataResponseBody) SetCode(v int64) *QuerySqlAnalysisDataResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBody) SetData(v *QuerySqlAnalysisDataResponseBodyData) *QuerySqlAnalysisDataResponseBody {
	s.Data = v
	return s
}

func (s *QuerySqlAnalysisDataResponseBody) SetMessage(v string) *QuerySqlAnalysisDataResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBody) SetRequestId(v string) *QuerySqlAnalysisDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBody) SetSuccess(v bool) *QuerySqlAnalysisDataResponseBody {
	s.Success = &v
	return s
}

type QuerySqlAnalysisDataResponseBodyData struct {
	List  []*QuerySqlAnalysisDataResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuerySqlAnalysisDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySqlAnalysisDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySqlAnalysisDataResponseBodyData) SetList(v []*QuerySqlAnalysisDataResponseBodyDataList) *QuerySqlAnalysisDataResponseBodyData {
	s.List = v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyData) SetTotal(v int64) *QuerySqlAnalysisDataResponseBodyData {
	s.Total = &v
	return s
}

type QuerySqlAnalysisDataResponseBodyDataList struct {
	AvgRowsExamined *int64   `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty"`
	AvgRt           *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	Cnt             *int64   `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	CurrentEndTime  *int64   `json:"CurrentEndTime,omitempty" xml:"CurrentEndTime,omitempty"`
	DbName          *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ErrorCnt        *int64   `json:"ErrorCnt,omitempty" xml:"ErrorCnt,omitempty"`
	FirstTime       *int64   `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	InstanceId      *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxRowsExamined *int64   `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty"`
	Sql             *string  `json:"Sql,omitempty" xml:"Sql,omitempty"`
	SqlId           *string  `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlTextFeature  *string  `json:"SqlTextFeature,omitempty" xml:"SqlTextFeature,omitempty"`
	SqlType         *string  `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	SumRt           *float64 `json:"SumRt,omitempty" xml:"SumRt,omitempty"`
	TableList       *string  `json:"TableList,omitempty" xml:"TableList,omitempty"`
	TaskId          *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type            *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySqlAnalysisDataResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QuerySqlAnalysisDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetAvgRowsExamined(v int64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.AvgRowsExamined = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetAvgRt(v float64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.AvgRt = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetCnt(v int64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.Cnt = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetCurrentEndTime(v int64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.CurrentEndTime = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetDbName(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.DbName = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetErrorCnt(v int64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.ErrorCnt = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetFirstTime(v int64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.FirstTime = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetInstanceId(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetMaxRowsExamined(v int64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.MaxRowsExamined = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetSql(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.Sql = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetSqlId(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetSqlTextFeature(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.SqlTextFeature = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetSqlType(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.SqlType = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetSumRt(v float64) *QuerySqlAnalysisDataResponseBodyDataList {
	s.SumRt = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetTableList(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.TableList = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetTaskId(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetType(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *QuerySqlAnalysisDataResponseBodyDataList) SetUserId(v string) *QuerySqlAnalysisDataResponseBodyDataList {
	s.UserId = &v
	return s
}

type QuerySqlAnalysisDataResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySqlAnalysisDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySqlAnalysisDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySqlAnalysisDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySqlAnalysisDataResponse) SetHeaders(v map[string]*string) *QuerySqlAnalysisDataResponse {
	s.Headers = v
	return s
}

func (s *QuerySqlAnalysisDataResponse) SetBody(v *QuerySqlAnalysisDataResponseBody) *QuerySqlAnalysisDataResponse {
	s.Body = v
	return s
}

type RunCloudBenchTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RunCloudBenchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCloudBenchTaskRequest) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskRequest) SetTaskId(v string) *RunCloudBenchTaskRequest {
	s.TaskId = &v
	return s
}

type RunCloudBenchTaskResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RunCloudBenchTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunCloudBenchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCloudBenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponseBody) SetCode(v string) *RunCloudBenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetData(v *RunCloudBenchTaskResponseBodyData) *RunCloudBenchTaskResponseBody {
	s.Data = v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetMessage(v string) *RunCloudBenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetRequestId(v string) *RunCloudBenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetSuccess(v string) *RunCloudBenchTaskResponseBody {
	s.Success = &v
	return s
}

type RunCloudBenchTaskResponseBodyData struct {
	PreCheckItem []*RunCloudBenchTaskResponseBodyDataPreCheckItem `json:"PreCheckItem,omitempty" xml:"PreCheckItem,omitempty" type:"Repeated"`
}

func (s RunCloudBenchTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunCloudBenchTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponseBodyData) SetPreCheckItem(v []*RunCloudBenchTaskResponseBodyDataPreCheckItem) *RunCloudBenchTaskResponseBodyData {
	s.PreCheckItem = v
	return s
}

type RunCloudBenchTaskResponseBodyDataPreCheckItem struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order   *int32  `json:"Order,omitempty" xml:"Order,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RunCloudBenchTaskResponseBodyDataPreCheckItem) String() string {
	return tea.Prettify(s)
}

func (s RunCloudBenchTaskResponseBodyDataPreCheckItem) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetCode(v int32) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Code = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetDetails(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Details = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetMessage(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Message = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetName(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Name = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetOrder(v int32) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Order = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetStatus(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Status = &v
	return s
}

type RunCloudBenchTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCloudBenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCloudBenchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCloudBenchTaskResponse) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponse) SetHeaders(v map[string]*string) *RunCloudBenchTaskResponse {
	s.Headers = v
	return s
}

func (s *RunCloudBenchTaskResponse) SetBody(v *RunCloudBenchTaskResponseBody) *RunCloudBenchTaskResponse {
	s.Body = v
	return s
}

type StopCloudBenchTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopCloudBenchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCloudBenchTaskRequest) GoString() string {
	return s.String()
}

func (s *StopCloudBenchTaskRequest) SetTaskId(v string) *StopCloudBenchTaskRequest {
	s.TaskId = &v
	return s
}

type StopCloudBenchTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopCloudBenchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCloudBenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudBenchTaskResponseBody) SetCode(v string) *StopCloudBenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopCloudBenchTaskResponseBody) SetData(v string) *StopCloudBenchTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopCloudBenchTaskResponseBody) SetMessage(v string) *StopCloudBenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopCloudBenchTaskResponseBody) SetRequestId(v string) *StopCloudBenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCloudBenchTaskResponseBody) SetSuccess(v string) *StopCloudBenchTaskResponseBody {
	s.Success = &v
	return s
}

type StopCloudBenchTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCloudBenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCloudBenchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCloudBenchTaskResponse) GoString() string {
	return s.String()
}

func (s *StopCloudBenchTaskResponse) SetHeaders(v map[string]*string) *StopCloudBenchTaskResponse {
	s.Headers = v
	return s
}

func (s *StopCloudBenchTaskResponse) SetBody(v *StopCloudBenchTaskResponseBody) *StopCloudBenchTaskResponse {
	s.Body = v
	return s
}

type StopOrRollbackOptimizeTaskRequest struct {
	AccessKey      *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	StopOrRollback *string `json:"StopOrRollback,omitempty" xml:"StopOrRollback,omitempty"`
	TaskType       *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskUuid       *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	Uid            *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context        *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s StopOrRollbackOptimizeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopOrRollbackOptimizeTaskRequest) GoString() string {
	return s.String()
}

func (s *StopOrRollbackOptimizeTaskRequest) SetAccessKey(v string) *StopOrRollbackOptimizeTaskRequest {
	s.AccessKey = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetInstanceId(v string) *StopOrRollbackOptimizeTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetSignature(v string) *StopOrRollbackOptimizeTaskRequest {
	s.Signature = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetStopOrRollback(v string) *StopOrRollbackOptimizeTaskRequest {
	s.StopOrRollback = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetTaskType(v string) *StopOrRollbackOptimizeTaskRequest {
	s.TaskType = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetTaskUuid(v string) *StopOrRollbackOptimizeTaskRequest {
	s.TaskUuid = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetUid(v string) *StopOrRollbackOptimizeTaskRequest {
	s.Uid = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetUserId(v string) *StopOrRollbackOptimizeTaskRequest {
	s.UserId = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskRequest) SetContext(v string) *StopOrRollbackOptimizeTaskRequest {
	s.Context = &v
	return s
}

type StopOrRollbackOptimizeTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s StopOrRollbackOptimizeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopOrRollbackOptimizeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopOrRollbackOptimizeTaskResponseBody) SetCode(v string) *StopOrRollbackOptimizeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskResponseBody) SetData(v string) *StopOrRollbackOptimizeTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskResponseBody) SetMessage(v string) *StopOrRollbackOptimizeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskResponseBody) SetRequestId(v string) *StopOrRollbackOptimizeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskResponseBody) SetSuccess(v string) *StopOrRollbackOptimizeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopOrRollbackOptimizeTaskResponseBody) SetSynchro(v string) *StopOrRollbackOptimizeTaskResponseBody {
	s.Synchro = &v
	return s
}

type StopOrRollbackOptimizeTaskResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopOrRollbackOptimizeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopOrRollbackOptimizeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopOrRollbackOptimizeTaskResponse) GoString() string {
	return s.String()
}

func (s *StopOrRollbackOptimizeTaskResponse) SetHeaders(v map[string]*string) *StopOrRollbackOptimizeTaskResponse {
	s.Headers = v
	return s
}

func (s *StopOrRollbackOptimizeTaskResponse) SetBody(v *StopOrRollbackOptimizeTaskResponseBody) *StopOrRollbackOptimizeTaskResponse {
	s.Body = v
	return s
}

type SyncHDMAliyunResourceRequest struct {
	Async                    *string `json:"Async,omitempty" xml:"Async,omitempty"`
	ResourceTypes            *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	Uid                      *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId                   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WaitForModifySecurityIps *string `json:"WaitForModifySecurityIps,omitempty" xml:"WaitForModifySecurityIps,omitempty"`
	Context                  *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey                *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature                *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth                 *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp                *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s SyncHDMAliyunResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncHDMAliyunResourceRequest) GoString() string {
	return s.String()
}

func (s *SyncHDMAliyunResourceRequest) SetAsync(v string) *SyncHDMAliyunResourceRequest {
	s.Async = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetResourceTypes(v string) *SyncHDMAliyunResourceRequest {
	s.ResourceTypes = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetUid(v string) *SyncHDMAliyunResourceRequest {
	s.Uid = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetUserId(v string) *SyncHDMAliyunResourceRequest {
	s.UserId = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetWaitForModifySecurityIps(v string) *SyncHDMAliyunResourceRequest {
	s.WaitForModifySecurityIps = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetContext(v string) *SyncHDMAliyunResourceRequest {
	s.Context = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetAccessKey(v string) *SyncHDMAliyunResourceRequest {
	s.AccessKey = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetSignature(v string) *SyncHDMAliyunResourceRequest {
	s.Signature = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetSkipAuth(v string) *SyncHDMAliyunResourceRequest {
	s.SkipAuth = &v
	return s
}

func (s *SyncHDMAliyunResourceRequest) SetTimestamp(v string) *SyncHDMAliyunResourceRequest {
	s.Timestamp = &v
	return s
}

type SyncHDMAliyunResourceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s SyncHDMAliyunResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncHDMAliyunResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SyncHDMAliyunResourceResponseBody) SetCode(v string) *SyncHDMAliyunResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SyncHDMAliyunResourceResponseBody) SetData(v string) *SyncHDMAliyunResourceResponseBody {
	s.Data = &v
	return s
}

func (s *SyncHDMAliyunResourceResponseBody) SetMessage(v string) *SyncHDMAliyunResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SyncHDMAliyunResourceResponseBody) SetRequestId(v string) *SyncHDMAliyunResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncHDMAliyunResourceResponseBody) SetSuccess(v string) *SyncHDMAliyunResourceResponseBody {
	s.Success = &v
	return s
}

func (s *SyncHDMAliyunResourceResponseBody) SetSynchro(v string) *SyncHDMAliyunResourceResponseBody {
	s.Synchro = &v
	return s
}

type SyncHDMAliyunResourceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncHDMAliyunResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncHDMAliyunResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncHDMAliyunResourceResponse) GoString() string {
	return s.String()
}

func (s *SyncHDMAliyunResourceResponse) SetHeaders(v map[string]*string) *SyncHDMAliyunResourceResponse {
	s.Headers = v
	return s
}

func (s *SyncHDMAliyunResourceResponse) SetBody(v *SyncHDMAliyunResourceResponseBody) *SyncHDMAliyunResourceResponse {
	s.Body = v
	return s
}

type TurnOffAutoResourceOptimizeRequest struct {
	AccessKey  *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Signature  *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Uid        *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context    *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s TurnOffAutoResourceOptimizeRequest) String() string {
	return tea.Prettify(s)
}

func (s TurnOffAutoResourceOptimizeRequest) GoString() string {
	return s.String()
}

func (s *TurnOffAutoResourceOptimizeRequest) SetAccessKey(v string) *TurnOffAutoResourceOptimizeRequest {
	s.AccessKey = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeRequest) SetInstanceId(v string) *TurnOffAutoResourceOptimizeRequest {
	s.InstanceId = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeRequest) SetSignature(v string) *TurnOffAutoResourceOptimizeRequest {
	s.Signature = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeRequest) SetUid(v string) *TurnOffAutoResourceOptimizeRequest {
	s.Uid = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeRequest) SetUserId(v string) *TurnOffAutoResourceOptimizeRequest {
	s.UserId = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeRequest) SetContext(v string) *TurnOffAutoResourceOptimizeRequest {
	s.Context = &v
	return s
}

type TurnOffAutoResourceOptimizeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s TurnOffAutoResourceOptimizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TurnOffAutoResourceOptimizeResponseBody) GoString() string {
	return s.String()
}

func (s *TurnOffAutoResourceOptimizeResponseBody) SetCode(v string) *TurnOffAutoResourceOptimizeResponseBody {
	s.Code = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeResponseBody) SetData(v string) *TurnOffAutoResourceOptimizeResponseBody {
	s.Data = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeResponseBody) SetMessage(v string) *TurnOffAutoResourceOptimizeResponseBody {
	s.Message = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeResponseBody) SetRequestId(v string) *TurnOffAutoResourceOptimizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeResponseBody) SetSuccess(v string) *TurnOffAutoResourceOptimizeResponseBody {
	s.Success = &v
	return s
}

func (s *TurnOffAutoResourceOptimizeResponseBody) SetSynchro(v string) *TurnOffAutoResourceOptimizeResponseBody {
	s.Synchro = &v
	return s
}

type TurnOffAutoResourceOptimizeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TurnOffAutoResourceOptimizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TurnOffAutoResourceOptimizeResponse) String() string {
	return tea.Prettify(s)
}

func (s TurnOffAutoResourceOptimizeResponse) GoString() string {
	return s.String()
}

func (s *TurnOffAutoResourceOptimizeResponse) SetHeaders(v map[string]*string) *TurnOffAutoResourceOptimizeResponse {
	s.Headers = v
	return s
}

func (s *TurnOffAutoResourceOptimizeResponse) SetBody(v *TurnOffAutoResourceOptimizeResponseBody) *TurnOffAutoResourceOptimizeResponse {
	s.Body = v
	return s
}

type UpdateAutoResourceOptimizeConfigRequest struct {
	AccessKey                *string  `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AutoDefragment           *int32   `json:"AutoDefragment,omitempty" xml:"AutoDefragment,omitempty"`
	AutoDuplicateIndexDelete *int32   `json:"AutoDuplicateIndexDelete,omitempty" xml:"AutoDuplicateIndexDelete,omitempty"`
	InstanceId               *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Signature                *string  `json:"Signature,omitempty" xml:"Signature,omitempty"`
	TableFragmentationRatio  *float32 `json:"TableFragmentationRatio,omitempty" xml:"TableFragmentationRatio,omitempty"`
	TableSpaceSize           *float32 `json:"TableSpaceSize,omitempty" xml:"TableSpaceSize,omitempty"`
	Uid                      *string  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId                   *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context                  *string  `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s UpdateAutoResourceOptimizeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoResourceOptimizeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetAccessKey(v string) *UpdateAutoResourceOptimizeConfigRequest {
	s.AccessKey = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetAutoDefragment(v int32) *UpdateAutoResourceOptimizeConfigRequest {
	s.AutoDefragment = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetAutoDuplicateIndexDelete(v int32) *UpdateAutoResourceOptimizeConfigRequest {
	s.AutoDuplicateIndexDelete = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetInstanceId(v string) *UpdateAutoResourceOptimizeConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetSignature(v string) *UpdateAutoResourceOptimizeConfigRequest {
	s.Signature = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetTableFragmentationRatio(v float32) *UpdateAutoResourceOptimizeConfigRequest {
	s.TableFragmentationRatio = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetTableSpaceSize(v float32) *UpdateAutoResourceOptimizeConfigRequest {
	s.TableSpaceSize = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetUid(v string) *UpdateAutoResourceOptimizeConfigRequest {
	s.Uid = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetUserId(v string) *UpdateAutoResourceOptimizeConfigRequest {
	s.UserId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigRequest) SetContext(v string) *UpdateAutoResourceOptimizeConfigRequest {
	s.Context = &v
	return s
}

type UpdateAutoResourceOptimizeConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateAutoResourceOptimizeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoResourceOptimizeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeConfigResponseBody) SetCode(v string) *UpdateAutoResourceOptimizeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigResponseBody) SetData(v string) *UpdateAutoResourceOptimizeConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigResponseBody) SetMessage(v string) *UpdateAutoResourceOptimizeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigResponseBody) SetRequestId(v string) *UpdateAutoResourceOptimizeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigResponseBody) SetSuccess(v string) *UpdateAutoResourceOptimizeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigResponseBody) SetSynchro(v string) *UpdateAutoResourceOptimizeConfigResponseBody {
	s.Synchro = &v
	return s
}

type UpdateAutoResourceOptimizeConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAutoResourceOptimizeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAutoResourceOptimizeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoResourceOptimizeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeConfigResponse) SetHeaders(v map[string]*string) *UpdateAutoResourceOptimizeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoResourceOptimizeConfigResponse) SetBody(v *UpdateAutoResourceOptimizeConfigResponseBody) *UpdateAutoResourceOptimizeConfigResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-shanghai": tea.String("das.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("das"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddHDMInstanceWithOptions(request *AddHDMInstanceRequest, runtime *util.RuntimeOptions) (_result *AddHDMInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.FlushAccount)) {
		query["FlushAccount"] = request.FlushAccount
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAlias)) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceArea)) {
		query["InstanceArea"] = request.InstanceArea
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddHDMInstance"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddHDMInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHDMInstance(request *AddHDMInstanceRequest) (_result *AddHDMInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddHDMInstanceResponse{}
	_body, _err := client.AddHDMInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAdamBenchTaskWithOptions(request *CreateAdamBenchTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAdamBenchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DstInstanceId)) {
		query["DstInstanceId"] = request.DstInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DstSuperAccount)) {
		query["DstSuperAccount"] = request.DstSuperAccount
	}

	if !tea.BoolValue(util.IsUnset(request.DstSuperPassword)) {
		query["DstSuperPassword"] = request.DstSuperPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Rate)) {
		query["Rate"] = request.Rate
	}

	if !tea.BoolValue(util.IsUnset(request.RequestDuration)) {
		query["RequestDuration"] = request.RequestDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RequestStartTime)) {
		query["RequestStartTime"] = request.RequestStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SrcEngine)) {
		query["SrcEngine"] = request.SrcEngine
	}

	if !tea.BoolValue(util.IsUnset(request.SrcEngineVersion)) {
		query["SrcEngineVersion"] = request.SrcEngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SrcMaxQps)) {
		query["SrcMaxQps"] = request.SrcMaxQps
	}

	if !tea.BoolValue(util.IsUnset(request.SrcMeanQps)) {
		query["SrcMeanQps"] = request.SrcMeanQps
	}

	if !tea.BoolValue(util.IsUnset(request.SrcSqlOssAddr)) {
		query["SrcSqlOssAddr"] = request.SrcSqlOssAddr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAdamBenchTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAdamBenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAdamBenchTask(request *CreateAdamBenchTaskRequest) (_result *CreateAdamBenchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAdamBenchTaskResponse{}
	_body, _err := client.CreateAdamBenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCacheAnalysisJobWithOptions(request *CreateCacheAnalysisJobRequest, runtime *util.RuntimeOptions) (_result *CreateCacheAnalysisJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCacheAnalysisJob"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCacheAnalysisJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCacheAnalysisJob(request *CreateCacheAnalysisJobRequest) (_result *CreateCacheAnalysisJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCacheAnalysisJobResponse{}
	_body, _err := client.CreateCacheAnalysisJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCloudBenchTasksWithOptions(request *CreateCloudBenchTasksRequest, runtime *util.RuntimeOptions) (_result *CreateCloudBenchTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupTime)) {
		query["BackupTime"] = request.BackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DstConnectionString)) {
		query["DstConnectionString"] = request.DstConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DstInstanceId)) {
		query["DstInstanceId"] = request.DstInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DstPort)) {
		query["DstPort"] = request.DstPort
	}

	if !tea.BoolValue(util.IsUnset(request.DstSuperAccount)) {
		query["DstSuperAccount"] = request.DstSuperAccount
	}

	if !tea.BoolValue(util.IsUnset(request.DstSuperPassword)) {
		query["DstSuperPassword"] = request.DstSuperPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DstType)) {
		query["DstType"] = request.DstType
	}

	if !tea.BoolValue(util.IsUnset(request.DtsJobClass)) {
		query["DtsJobClass"] = request.DtsJobClass
	}

	if !tea.BoolValue(util.IsUnset(request.DtsJobId)) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !tea.BoolValue(util.IsUnset(request.EndState)) {
		query["EndState"] = request.EndState
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayVpcId)) {
		query["GatewayVpcId"] = request.GatewayVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayVpcIp)) {
		query["GatewayVpcIp"] = request.GatewayVpcIp
	}

	if !tea.BoolValue(util.IsUnset(request.Rate)) {
		query["Rate"] = request.Rate
	}

	if !tea.BoolValue(util.IsUnset(request.RequestDuration)) {
		query["RequestDuration"] = request.RequestDuration
	}

	if !tea.BoolValue(util.IsUnset(request.RequestEndTime)) {
		query["RequestEndTime"] = request.RequestEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RequestStartTime)) {
		query["RequestStartTime"] = request.RequestStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SmartPressureTime)) {
		query["SmartPressureTime"] = request.SmartPressureTime
	}

	if !tea.BoolValue(util.IsUnset(request.SrcInstanceId)) {
		query["SrcInstanceId"] = request.SrcInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcPublicIp)) {
		query["SrcPublicIp"] = request.SrcPublicIp
	}

	if !tea.BoolValue(util.IsUnset(request.SrcSuperAccount)) {
		query["SrcSuperAccount"] = request.SrcSuperAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SrcSuperPassword)) {
		query["SrcSuperPassword"] = request.SrcSuperPassword
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkDir)) {
		query["WorkDir"] = request.WorkDir
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCloudBenchTasks"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCloudBenchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCloudBenchTasks(request *CreateCloudBenchTasksRequest) (_result *CreateCloudBenchTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCloudBenchTasksResponse{}
	_body, _err := client.CreateCloudBenchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDiagnosticReportWithOptions(request *CreateDiagnosticReportRequest, runtime *util.RuntimeOptions) (_result *CreateDiagnosticReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiagnosticReport"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiagnosticReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiagnosticReport(request *CreateDiagnosticReportRequest) (_result *CreateDiagnosticReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiagnosticReportResponse{}
	_body, _err := client.CreateDiagnosticReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRequestDiagnosisWithOptions(request *CreateRequestDiagnosisRequest, runtime *util.RuntimeOptions) (_result *CreateRequestDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		query["Sql"] = request.Sql
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRequestDiagnosis"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRequestDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRequestDiagnosis(request *CreateRequestDiagnosisRequest) (_result *CreateRequestDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRequestDiagnosisResponse{}
	_body, _err := client.CreateRequestDiagnosisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCloudBenchTaskWithOptions(request *DeleteCloudBenchTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteCloudBenchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCloudBenchTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCloudBenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCloudBenchTask(request *DeleteCloudBenchTaskRequest) (_result *DeleteCloudBenchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCloudBenchTaskResponse{}
	_body, _err := client.DeleteCloudBenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisJobWithOptions(request *DescribeCacheAnalysisJobRequest, runtime *util.RuntimeOptions) (_result *DescribeCacheAnalysisJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCacheAnalysisJob"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCacheAnalysisJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisJob(request *DescribeCacheAnalysisJobRequest) (_result *DescribeCacheAnalysisJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCacheAnalysisJobResponse{}
	_body, _err := client.DescribeCacheAnalysisJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisJobsWithOptions(request *DescribeCacheAnalysisJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeCacheAnalysisJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCacheAnalysisJobs"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCacheAnalysisJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCacheAnalysisJobs(request *DescribeCacheAnalysisJobsRequest) (_result *DescribeCacheAnalysisJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCacheAnalysisJobsResponse{}
	_body, _err := client.DescribeCacheAnalysisJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudBenchTasksWithOptions(request *DescribeCloudBenchTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudBenchTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
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
		Action:      tea.String("DescribeCloudBenchTasks"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudBenchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudBenchTasks(request *DescribeCloudBenchTasksRequest) (_result *DescribeCloudBenchTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudBenchTasksResponse{}
	_body, _err := client.DescribeCloudBenchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudbenchTaskWithOptions(request *DescribeCloudbenchTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudbenchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudbenchTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudbenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudbenchTask(request *DescribeCloudbenchTaskRequest) (_result *DescribeCloudbenchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudbenchTaskResponse{}
	_body, _err := client.DescribeCloudbenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudbenchTaskConfigWithOptions(request *DescribeCloudbenchTaskConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudbenchTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudbenchTaskConfig"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudbenchTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudbenchTaskConfig(request *DescribeCloudbenchTaskConfigRequest) (_result *DescribeCloudbenchTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudbenchTaskConfigResponse{}
	_body, _err := client.DescribeCloudbenchTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosticReportListWithOptions(request *DescribeDiagnosticReportListRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosticReportListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosticReportList"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosticReportListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosticReportList(request *DescribeDiagnosticReportListRequest) (_result *DescribeDiagnosticReportListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosticReportListResponse{}
	_body, _err := client.DescribeDiagnosticReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHotBigKeysWithOptions(request *DescribeHotBigKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeHotBigKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHotBigKeys"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHotBigKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHotBigKeys(request *DescribeHotBigKeysRequest) (_result *DescribeHotBigKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHotBigKeysResponse{}
	_body, _err := client.DescribeHotBigKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHotKeysWithOptions(request *DescribeHotKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeHotKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHotKeys"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHotKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHotKeys(request *DescribeHotKeysRequest) (_result *DescribeHotKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHotKeysResponse{}
	_body, _err := client.DescribeHotKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceDasProWithOptions(request *DescribeInstanceDasProRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDasProResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceDasPro"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceDasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceDasPro(request *DescribeInstanceDasProRequest) (_result *DescribeInstanceDasProResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceDasProResponse{}
	_body, _err := client.DescribeInstanceDasProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopBigKeysWithOptions(request *DescribeTopBigKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeTopBigKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopBigKeys"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopBigKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopBigKeys(request *DescribeTopBigKeysRequest) (_result *DescribeTopBigKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopBigKeysResponse{}
	_body, _err := client.DescribeTopBigKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopHotKeysWithOptions(request *DescribeTopHotKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeTopHotKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopHotKeys"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopHotKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopHotKeys(request *DescribeTopHotKeysRequest) (_result *DescribeTopHotKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopHotKeysResponse{}
	_body, _err := client.DescribeTopHotKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAllSqlConcurrencyControlRulesWithOptions(request *DisableAllSqlConcurrencyControlRulesRequest, runtime *util.RuntimeOptions) (_result *DisableAllSqlConcurrencyControlRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAllSqlConcurrencyControlRules"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAllSqlConcurrencyControlRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAllSqlConcurrencyControlRules(request *DisableAllSqlConcurrencyControlRulesRequest) (_result *DisableAllSqlConcurrencyControlRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAllSqlConcurrencyControlRulesResponse{}
	_body, _err := client.DisableAllSqlConcurrencyControlRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDasProWithOptions(request *DisableDasProRequest, runtime *util.RuntimeOptions) (_result *DisableDasProResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableDasPro"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableDasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDasPro(request *DisableDasProRequest) (_result *DisableDasProResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDasProResponse{}
	_body, _err := client.DisableDasProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSqlConcurrencyControlWithOptions(request *DisableSqlConcurrencyControlRequest, runtime *util.RuntimeOptions) (_result *DisableSqlConcurrencyControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSqlConcurrencyControl"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSqlConcurrencyControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSqlConcurrencyControl(request *DisableSqlConcurrencyControlRequest) (_result *DisableSqlConcurrencyControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSqlConcurrencyControlResponse{}
	_body, _err := client.DisableSqlConcurrencyControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableDasProWithOptions(request *EnableDasProRequest, runtime *util.RuntimeOptions) (_result *EnableDasProResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlRetention)) {
		query["SqlRetention"] = request.SqlRetention
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableDasPro"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableDasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableDasPro(request *EnableDasProRequest) (_result *EnableDasProResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableDasProResponse{}
	_body, _err := client.EnableDasProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSqlConcurrencyControlWithOptions(request *EnableSqlConcurrencyControlRequest, runtime *util.RuntimeOptions) (_result *EnableSqlConcurrencyControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConcurrencyControlTime)) {
		query["ConcurrencyControlTime"] = request.ConcurrencyControlTime
	}

	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.SqlKeywords)) {
		query["SqlKeywords"] = request.SqlKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSqlConcurrencyControl"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSqlConcurrencyControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSqlConcurrencyControl(request *EnableSqlConcurrencyControlRequest) (_result *EnableSqlConcurrencyControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSqlConcurrencyControlResponse{}
	_body, _err := client.EnableSqlConcurrencyControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncErrorRequestListByCodeWithOptions(request *GetAsyncErrorRequestListByCodeRequest, runtime *util.RuntimeOptions) (_result *GetAsyncErrorRequestListByCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		query["ErrorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncErrorRequestListByCode"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncErrorRequestListByCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncErrorRequestListByCode(request *GetAsyncErrorRequestListByCodeRequest) (_result *GetAsyncErrorRequestListByCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncErrorRequestListByCodeResponse{}
	_body, _err := client.GetAsyncErrorRequestListByCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncErrorRequestStatByCodeWithOptions(request *GetAsyncErrorRequestStatByCodeRequest, runtime *util.RuntimeOptions) (_result *GetAsyncErrorRequestStatByCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncErrorRequestStatByCode"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncErrorRequestStatByCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncErrorRequestStatByCode(request *GetAsyncErrorRequestStatByCodeRequest) (_result *GetAsyncErrorRequestStatByCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncErrorRequestStatByCodeResponse{}
	_body, _err := client.GetAsyncErrorRequestStatByCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncErrorRequestStatResultWithOptions(request *GetAsyncErrorRequestStatResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncErrorRequestStatResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlIdList)) {
		query["SqlIdList"] = request.SqlIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncErrorRequestStatResult"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncErrorRequestStatResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncErrorRequestStatResult(request *GetAsyncErrorRequestStatResultRequest) (_result *GetAsyncErrorRequestStatResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncErrorRequestStatResultResponse{}
	_body, _err := client.GetAsyncErrorRequestStatResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutoResourceOptimizeConfigWithOptions(request *GetAutoResourceOptimizeConfigRequest, runtime *util.RuntimeOptions) (_result *GetAutoResourceOptimizeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoResourceOptimizeConfig"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutoResourceOptimizeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutoResourceOptimizeConfig(request *GetAutoResourceOptimizeConfigRequest) (_result *GetAutoResourceOptimizeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoResourceOptimizeConfigResponse{}
	_body, _err := client.GetAutoResourceOptimizeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutonomousNotifyEventContentWithOptions(request *GetAutonomousNotifyEventContentRequest, runtime *util.RuntimeOptions) (_result *GetAutonomousNotifyEventContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SpanId)) {
		query["SpanId"] = request.SpanId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutonomousNotifyEventContent"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutonomousNotifyEventContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutonomousNotifyEventContent(request *GetAutonomousNotifyEventContentRequest) (_result *GetAutonomousNotifyEventContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutonomousNotifyEventContentResponse{}
	_body, _err := client.GetAutonomousNotifyEventContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutonomousNotifyEventsInRangeWithOptions(request *GetAutonomousNotifyEventsInRangeRequest, runtime *util.RuntimeOptions) (_result *GetAutonomousNotifyEventsInRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventContext)) {
		query["EventContext"] = request.EventContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.MinLevel)) {
		query["MinLevel"] = request.MinLevel
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.PageOffset)) {
		query["PageOffset"] = request.PageOffset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutonomousNotifyEventsInRange"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutonomousNotifyEventsInRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutonomousNotifyEventsInRange(request *GetAutonomousNotifyEventsInRangeRequest) (_result *GetAutonomousNotifyEventsInRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutonomousNotifyEventsInRangeResponse{}
	_body, _err := client.GetAutonomousNotifyEventsInRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDasProServiceUsageWithOptions(request *GetDasProServiceUsageRequest, runtime *util.RuntimeOptions) (_result *GetDasProServiceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDasProServiceUsage"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDasProServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDasProServiceUsage(request *GetDasProServiceUsageRequest) (_result *GetDasProServiceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDasProServiceUsageResponse{}
	_body, _err := client.GetDasProServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpointSwitchTaskWithOptions(request *GetEndpointSwitchTaskRequest, runtime *util.RuntimeOptions) (_result *GetEndpointSwitchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SkipAuth)) {
		query["skipAuth"] = request.SkipAuth
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEndpointSwitchTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEndpointSwitchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEndpointSwitchTask(request *GetEndpointSwitchTaskRequest) (_result *GetEndpointSwitchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEndpointSwitchTaskResponse{}
	_body, _err := client.GetEndpointSwitchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetErrorRequestSampleWithOptions(request *GetErrorRequestSampleRequest, runtime *util.RuntimeOptions) (_result *GetErrorRequestSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlId)) {
		query["SqlId"] = request.SqlId
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetErrorRequestSample"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetErrorRequestSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetErrorRequestSample(request *GetErrorRequestSampleRequest) (_result *GetErrorRequestSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetErrorRequestSampleResponse{}
	_body, _err := client.GetErrorRequestSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFullRequestOriginStatByInstanceIdWithOptions(request *GetFullRequestOriginStatByInstanceIdRequest, runtime *util.RuntimeOptions) (_result *GetFullRequestOriginStatByInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Asc)) {
		query["Asc"] = request.Asc
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFullRequestOriginStatByInstanceId"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFullRequestOriginStatByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFullRequestOriginStatByInstanceId(request *GetFullRequestOriginStatByInstanceIdRequest) (_result *GetFullRequestOriginStatByInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFullRequestOriginStatByInstanceIdResponse{}
	_body, _err := client.GetFullRequestOriginStatByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFullRequestStatResultByInstanceIdWithOptions(request *GetFullRequestStatResultByInstanceIdRequest, runtime *util.RuntimeOptions) (_result *GetFullRequestStatResultByInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Asc)) {
		query["Asc"] = request.Asc
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OriginHost)) {
		query["OriginHost"] = request.OriginHost
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.SqlId)) {
		query["SqlId"] = request.SqlId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFullRequestStatResultByInstanceId"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFullRequestStatResultByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFullRequestStatResultByInstanceId(request *GetFullRequestStatResultByInstanceIdRequest) (_result *GetFullRequestStatResultByInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFullRequestStatResultByInstanceIdResponse{}
	_body, _err := client.GetFullRequestStatResultByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHDMAliyunResourceSyncResultWithOptions(request *GetHDMAliyunResourceSyncResultRequest, runtime *util.RuntimeOptions) (_result *GetHDMAliyunResourceSyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SkipAuth)) {
		query["skipAuth"] = request.SkipAuth
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHDMAliyunResourceSyncResult"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHDMAliyunResourceSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHDMAliyunResourceSyncResult(request *GetHDMAliyunResourceSyncResultRequest) (_result *GetHDMAliyunResourceSyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHDMAliyunResourceSyncResultResponse{}
	_body, _err := client.GetHDMAliyunResourceSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHDMLastAliyunResourceSyncResultWithOptions(request *GetHDMLastAliyunResourceSyncResultRequest, runtime *util.RuntimeOptions) (_result *GetHDMLastAliyunResourceSyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SkipAuth)) {
		query["skipAuth"] = request.SkipAuth
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHDMLastAliyunResourceSyncResult"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHDMLastAliyunResourceSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHDMLastAliyunResourceSyncResult(request *GetHDMLastAliyunResourceSyncResultRequest) (_result *GetHDMLastAliyunResourceSyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHDMLastAliyunResourceSyncResultResponse{}
	_body, _err := client.GetHDMLastAliyunResourceSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceInspectionsWithOptions(request *GetInstanceInspectionsRequest, runtime *util.RuntimeOptions) (_result *GetInstanceInspectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceArea)) {
		query["InstanceArea"] = request.InstanceArea
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchMap)) {
		query["SearchMap"] = request.SearchMap
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceInspections"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceInspectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceInspections(request *GetInstanceInspectionsRequest) (_result *GetInstanceInspectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceInspectionsResponse{}
	_body, _err := client.GetInstanceInspectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeDataStatsWithOptions(request *GetQueryOptimizeDataStatsRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeDataStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeDataStats"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeDataStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeDataStats(request *GetQueryOptimizeDataStatsRequest) (_result *GetQueryOptimizeDataStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeDataStatsResponse{}
	_body, _err := client.GetQueryOptimizeDataStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeDataTopWithOptions(request *GetQueryOptimizeDataTopRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeDataTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeDataTop"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeDataTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeDataTop(request *GetQueryOptimizeDataTopRequest) (_result *GetQueryOptimizeDataTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeDataTopResponse{}
	_body, _err := client.GetQueryOptimizeDataTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeDataTrendWithOptions(request *GetQueryOptimizeDataTrendRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeDataTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeDataTrend"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeDataTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeDataTrend(request *GetQueryOptimizeDataTrendRequest) (_result *GetQueryOptimizeDataTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeDataTrendResponse{}
	_body, _err := client.GetQueryOptimizeDataTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeExecErrorSampleWithOptions(request *GetQueryOptimizeExecErrorSampleRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeExecErrorSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeExecErrorSample"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeExecErrorSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeExecErrorSample(request *GetQueryOptimizeExecErrorSampleRequest) (_result *GetQueryOptimizeExecErrorSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeExecErrorSampleResponse{}
	_body, _err := client.GetQueryOptimizeExecErrorSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeExecErrorStatsWithOptions(request *GetQueryOptimizeExecErrorStatsRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeExecErrorStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeExecErrorStats"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeExecErrorStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeExecErrorStats(request *GetQueryOptimizeExecErrorStatsRequest) (_result *GetQueryOptimizeExecErrorStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeExecErrorStatsResponse{}
	_body, _err := client.GetQueryOptimizeExecErrorStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeRuleListWithOptions(request *GetQueryOptimizeRuleListRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeRuleList"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeRuleList(request *GetQueryOptimizeRuleListRequest) (_result *GetQueryOptimizeRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeRuleListResponse{}
	_body, _err := client.GetQueryOptimizeRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryOptimizeSolutionWithOptions(request *GetQueryOptimizeSolutionRequest, runtime *util.RuntimeOptions) (_result *GetQueryOptimizeSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryOptimizeSolution"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryOptimizeSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryOptimizeSolution(request *GetQueryOptimizeSolutionRequest) (_result *GetQueryOptimizeSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryOptimizeSolutionResponse{}
	_body, _err := client.GetQueryOptimizeSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRequestDiagnosisPageWithOptions(request *GetRequestDiagnosisPageRequest, runtime *util.RuntimeOptions) (_result *GetRequestDiagnosisPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRequestDiagnosisPage"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRequestDiagnosisPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRequestDiagnosisPage(request *GetRequestDiagnosisPageRequest) (_result *GetRequestDiagnosisPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRequestDiagnosisPageResponse{}
	_body, _err := client.GetRequestDiagnosisPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRequestDiagnosisResultWithOptions(request *GetRequestDiagnosisResultRequest, runtime *util.RuntimeOptions) (_result *GetRequestDiagnosisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRequestDiagnosisResult"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRequestDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRequestDiagnosisResult(request *GetRequestDiagnosisResultRequest) (_result *GetRequestDiagnosisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRequestDiagnosisResultResponse{}
	_body, _err := client.GetRequestDiagnosisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceOptimizeHistoryListWithOptions(request *GetResourceOptimizeHistoryListRequest, runtime *util.RuntimeOptions) (_result *GetResourceOptimizeHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceOptimizeHistoryList"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceOptimizeHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceOptimizeHistoryList(request *GetResourceOptimizeHistoryListRequest) (_result *GetResourceOptimizeHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceOptimizeHistoryListResponse{}
	_body, _err := client.GetResourceOptimizeHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRunningSqlConcurrencyControlRulesWithOptions(request *GetRunningSqlConcurrencyControlRulesRequest, runtime *util.RuntimeOptions) (_result *GetRunningSqlConcurrencyControlRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRunningSqlConcurrencyControlRules"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRunningSqlConcurrencyControlRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRunningSqlConcurrencyControlRules(request *GetRunningSqlConcurrencyControlRulesRequest) (_result *GetRunningSqlConcurrencyControlRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRunningSqlConcurrencyControlRulesResponse{}
	_body, _err := client.GetRunningSqlConcurrencyControlRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSqlAnalysisReportWithOptions(request *GetSqlAnalysisReportRequest, runtime *util.RuntimeOptions) (_result *GetSqlAnalysisReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleNameList)) {
		query["ModuleNameList"] = request.ModuleNameList
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSqlAnalysisReport"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSqlAnalysisReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSqlAnalysisReport(request *GetSqlAnalysisReportRequest) (_result *GetSqlAnalysisReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSqlAnalysisReportResponse{}
	_body, _err := client.GetSqlAnalysisReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSqlConcurrencyControlKeywordsFromSqlTextWithOptions(request *GetSqlConcurrencyControlKeywordsFromSqlTextRequest, runtime *util.RuntimeOptions) (_result *GetSqlConcurrencyControlKeywordsFromSqlTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlText)) {
		query["SqlText"] = request.SqlText
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSqlConcurrencyControlKeywordsFromSqlText"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSqlConcurrencyControlKeywordsFromSqlTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSqlConcurrencyControlKeywordsFromSqlText(request *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) (_result *GetSqlConcurrencyControlKeywordsFromSqlTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSqlConcurrencyControlKeywordsFromSqlTextResponse{}
	_body, _err := client.GetSqlConcurrencyControlKeywordsFromSqlTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSqlConcurrencyControlRulesHistoryWithOptions(request *GetSqlConcurrencyControlRulesHistoryRequest, runtime *util.RuntimeOptions) (_result *GetSqlConcurrencyControlRulesHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSqlConcurrencyControlRulesHistory"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSqlConcurrencyControlRulesHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSqlConcurrencyControlRulesHistory(request *GetSqlConcurrencyControlRulesHistoryRequest) (_result *GetSqlConcurrencyControlRulesHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSqlConcurrencyControlRulesHistoryResponse{}
	_body, _err := client.GetSqlConcurrencyControlRulesHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSqlOptimizeAdviceWithOptions(request *GetSqlOptimizeAdviceRequest, runtime *util.RuntimeOptions) (_result *GetSqlOptimizeAdviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !tea.BoolValue(util.IsUnset(request.EndDt)) {
		query["EndDt"] = request.EndDt
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDt)) {
		query["StartDt"] = request.StartDt
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSqlOptimizeAdvice"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSqlOptimizeAdviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSqlOptimizeAdvice(request *GetSqlOptimizeAdviceRequest) (_result *GetSqlOptimizeAdviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSqlOptimizeAdviceResponse{}
	_body, _err := client.GetSqlOptimizeAdviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySqlAnalysisDataWithOptions(request *QuerySqlAnalysisDataRequest, runtime *util.RuntimeOptions) (_result *QuerySqlAnalysisDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsoleContext)) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SqlIdList)) {
		body["SqlIdList"] = request.SqlIdList
	}

	if !tea.BoolValue(util.IsUnset(request.SqlTextFeature)) {
		body["SqlTextFeature"] = request.SqlTextFeature
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		body["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySqlAnalysisData"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySqlAnalysisDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySqlAnalysisData(request *QuerySqlAnalysisDataRequest) (_result *QuerySqlAnalysisDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySqlAnalysisDataResponse{}
	_body, _err := client.QuerySqlAnalysisDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCloudBenchTaskWithOptions(request *RunCloudBenchTaskRequest, runtime *util.RuntimeOptions) (_result *RunCloudBenchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCloudBenchTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCloudBenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCloudBenchTask(request *RunCloudBenchTaskRequest) (_result *RunCloudBenchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCloudBenchTaskResponse{}
	_body, _err := client.RunCloudBenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCloudBenchTaskWithOptions(request *StopCloudBenchTaskRequest, runtime *util.RuntimeOptions) (_result *StopCloudBenchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCloudBenchTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCloudBenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCloudBenchTask(request *StopCloudBenchTaskRequest) (_result *StopCloudBenchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCloudBenchTaskResponse{}
	_body, _err := client.StopCloudBenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopOrRollbackOptimizeTaskWithOptions(request *StopOrRollbackOptimizeTaskRequest, runtime *util.RuntimeOptions) (_result *StopOrRollbackOptimizeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.StopOrRollback)) {
		query["StopOrRollback"] = request.StopOrRollback
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		query["TaskUuid"] = request.TaskUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopOrRollbackOptimizeTask"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopOrRollbackOptimizeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopOrRollbackOptimizeTask(request *StopOrRollbackOptimizeTaskRequest) (_result *StopOrRollbackOptimizeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopOrRollbackOptimizeTaskResponse{}
	_body, _err := client.StopOrRollbackOptimizeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncHDMAliyunResourceWithOptions(request *SyncHDMAliyunResourceRequest, runtime *util.RuntimeOptions) (_result *SyncHDMAliyunResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WaitForModifySecurityIps)) {
		query["WaitForModifySecurityIps"] = request.WaitForModifySecurityIps
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SkipAuth)) {
		query["skipAuth"] = request.SkipAuth
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncHDMAliyunResource"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncHDMAliyunResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncHDMAliyunResource(request *SyncHDMAliyunResourceRequest) (_result *SyncHDMAliyunResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncHDMAliyunResourceResponse{}
	_body, _err := client.SyncHDMAliyunResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TurnOffAutoResourceOptimizeWithOptions(request *TurnOffAutoResourceOptimizeRequest, runtime *util.RuntimeOptions) (_result *TurnOffAutoResourceOptimizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TurnOffAutoResourceOptimize"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TurnOffAutoResourceOptimizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TurnOffAutoResourceOptimize(request *TurnOffAutoResourceOptimizeRequest) (_result *TurnOffAutoResourceOptimizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TurnOffAutoResourceOptimizeResponse{}
	_body, _err := client.TurnOffAutoResourceOptimizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAutoResourceOptimizeConfigWithOptions(request *UpdateAutoResourceOptimizeConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateAutoResourceOptimizeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDefragment)) {
		query["AutoDefragment"] = request.AutoDefragment
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDuplicateIndexDelete)) {
		query["AutoDuplicateIndexDelete"] = request.AutoDuplicateIndexDelete
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.TableFragmentationRatio)) {
		query["TableFragmentationRatio"] = request.TableFragmentationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.TableSpaceSize)) {
		query["TableSpaceSize"] = request.TableSpaceSize
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["__context"] = request.Context
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAutoResourceOptimizeConfig"),
		Version:     tea.String("2020-01-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAutoResourceOptimizeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAutoResourceOptimizeConfig(request *UpdateAutoResourceOptimizeConfigRequest) (_result *UpdateAutoResourceOptimizeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAutoResourceOptimizeConfigResponse{}
	_body, _err := client.UpdateAutoResourceOptimizeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

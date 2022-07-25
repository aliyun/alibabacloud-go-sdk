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

type CreateSandboxInstanceRequest struct {
	BackupPlanId         *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SandboxInstanceName  *string `json:"SandboxInstanceName,omitempty" xml:"SandboxInstanceName,omitempty"`
	SandboxPassword      *string `json:"SandboxPassword,omitempty" xml:"SandboxPassword,omitempty"`
	SandboxSpecification *string `json:"SandboxSpecification,omitempty" xml:"SandboxSpecification,omitempty"`
	SandboxType          *string `json:"SandboxType,omitempty" xml:"SandboxType,omitempty"`
	SandboxUser          *string `json:"SandboxUser,omitempty" xml:"SandboxUser,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcSwitchId          *string `json:"VpcSwitchId,omitempty" xml:"VpcSwitchId,omitempty"`
}

func (s CreateSandboxInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceRequest) SetBackupPlanId(v string) *CreateSandboxInstanceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetBackupSetId(v string) *CreateSandboxInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetRestoreTime(v string) *CreateSandboxInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxInstanceName(v string) *CreateSandboxInstanceRequest {
	s.SandboxInstanceName = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxPassword(v string) *CreateSandboxInstanceRequest {
	s.SandboxPassword = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxSpecification(v string) *CreateSandboxInstanceRequest {
	s.SandboxSpecification = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxType(v string) *CreateSandboxInstanceRequest {
	s.SandboxType = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxUser(v string) *CreateSandboxInstanceRequest {
	s.SandboxUser = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetVpcId(v string) *CreateSandboxInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetVpcSwitchId(v string) *CreateSandboxInstanceRequest {
	s.VpcSwitchId = &v
	return s
}

type CreateSandboxInstanceResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *CreateSandboxInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                                `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSandboxInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceResponseBody) SetCode(v string) *CreateSandboxInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetData(v *CreateSandboxInstanceResponseBodyData) *CreateSandboxInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetErrCode(v string) *CreateSandboxInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetErrMessage(v string) *CreateSandboxInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetMessage(v string) *CreateSandboxInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetRequestId(v string) *CreateSandboxInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetSuccess(v string) *CreateSandboxInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateSandboxInstanceResponseBodyData struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateSandboxInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceResponseBodyData) SetBackupPlanId(v string) *CreateSandboxInstanceResponseBodyData {
	s.BackupPlanId = &v
	return s
}

func (s *CreateSandboxInstanceResponseBodyData) SetInstanceId(v string) *CreateSandboxInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

type CreateSandboxInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSandboxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSandboxInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceResponse) SetHeaders(v map[string]*string) *CreateSandboxInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSandboxInstanceResponse) SetStatusCode(v int32) *CreateSandboxInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSandboxInstanceResponse) SetBody(v *CreateSandboxInstanceResponseBody) *CreateSandboxInstanceResponse {
	s.Body = v
	return s
}

type DeleteSandboxInstanceRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteSandboxInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSandboxInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceRequest) SetBackupPlanId(v string) *DeleteSandboxInstanceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DeleteSandboxInstanceRequest) SetInstanceId(v string) *DeleteSandboxInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteSandboxInstanceResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSandboxInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSandboxInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceResponseBody) SetCode(v string) *DeleteSandboxInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetData(v string) *DeleteSandboxInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetErrCode(v string) *DeleteSandboxInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetErrMessage(v string) *DeleteSandboxInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetMessage(v string) *DeleteSandboxInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetRequestId(v string) *DeleteSandboxInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetSuccess(v string) *DeleteSandboxInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteSandboxInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSandboxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSandboxInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSandboxInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceResponse) SetHeaders(v map[string]*string) *DeleteSandboxInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSandboxInstanceResponse) SetStatusCode(v int32) *DeleteSandboxInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSandboxInstanceResponse) SetBody(v *DeleteSandboxInstanceResponseBody) *DeleteSandboxInstanceResponse {
	s.Body = v
	return s
}

type DescribeSandboxBackupSetsRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupSetId  *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSandboxBackupSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxBackupSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsRequest) SetBackupPlanId(v string) *DescribeSandboxBackupSetsRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetBackupSetId(v string) *DescribeSandboxBackupSetsRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetPageNumber(v string) *DescribeSandboxBackupSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetPageSize(v string) *DescribeSandboxBackupSetsRequest {
	s.PageSize = &v
	return s
}

type DescribeSandboxBackupSetsResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxBackupSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsResponseBody) SetCode(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetData(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetErrCode(v string) *DescribeSandboxBackupSetsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetErrMessage(v string) *DescribeSandboxBackupSetsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetMessage(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetRequestId(v string) *DescribeSandboxBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetSuccess(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Success = &v
	return s
}

type DescribeSandboxBackupSetsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSandboxBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSandboxBackupSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxBackupSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsResponse) SetHeaders(v map[string]*string) *DescribeSandboxBackupSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) SetStatusCode(v int32) *DescribeSandboxBackupSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) SetBody(v *DescribeSandboxBackupSetsResponseBody) *DescribeSandboxBackupSetsResponse {
	s.Body = v
	return s
}

type DescribeSandboxInstancesRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSandboxInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesRequest) SetBackupPlanId(v string) *DescribeSandboxInstancesRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetInstanceId(v string) *DescribeSandboxInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetPageNumber(v string) *DescribeSandboxInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetPageSize(v string) *DescribeSandboxInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeSandboxInstancesResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesResponseBody) SetCode(v string) *DescribeSandboxInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetData(v string) *DescribeSandboxInstancesResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetErrCode(v string) *DescribeSandboxInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetErrMessage(v string) *DescribeSandboxInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetMessage(v string) *DescribeSandboxInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetRequestId(v string) *DescribeSandboxInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetSuccess(v string) *DescribeSandboxInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeSandboxInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSandboxInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSandboxInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesResponse) SetHeaders(v map[string]*string) *DescribeSandboxInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxInstancesResponse) SetStatusCode(v int32) *DescribeSandboxInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxInstancesResponse) SetBody(v *DescribeSandboxInstancesResponseBody) *DescribeSandboxInstancesResponse {
	s.Body = v
	return s
}

type DescribeSandboxRecoveryTimeRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
}

func (s DescribeSandboxRecoveryTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeRequest) SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeRequest {
	s.BackupPlanId = &v
	return s
}

type DescribeSandboxRecoveryTimeResponseBody struct {
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *DescribeSandboxRecoveryTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                                      `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                                      `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetCode(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetData(v *DescribeSandboxRecoveryTimeResponseBodyData) *DescribeSandboxRecoveryTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetErrCode(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetErrMessage(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetMessage(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetRequestId(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetSuccess(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Success = &v
	return s
}

type DescribeSandboxRecoveryTimeResponseBodyData struct {
	BackupPlanId      *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	RecoveryEndTime   *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetRecoveryBeginTime(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.RecoveryBeginTime = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetRecoveryEndTime(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.RecoveryEndTime = &v
	return s
}

type DescribeSandboxRecoveryTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSandboxRecoveryTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSandboxRecoveryTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponse) SetHeaders(v map[string]*string) *DescribeSandboxRecoveryTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) SetStatusCode(v int32) *DescribeSandboxRecoveryTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) SetBody(v *DescribeSandboxRecoveryTimeResponseBody) *DescribeSandboxRecoveryTimeResponse {
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
		"cn-qingdao":            tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            tea.String("dbs-api.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          tea.String("dbs-api.cn-huhehaote.aliyuncs.com"),
		"cn-hangzhou":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"ap-southeast-1":        tea.String("dbs-api.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("dbs-api.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("dbs-api.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("dbs-api.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1":        tea.String("dbs-api.ap-northeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"us-east-1":             tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"eu-central-1":          tea.String("dbs-api.eu-central-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"ap-south-1":            tea.String("dbs-api.ap-south-1.aliyuncs.com"),
		"eu-west-1":             tea.String("dbs-api.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("dbs-api.me-east-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dbs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateSandboxInstanceWithOptions(request *CreateSandboxInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateSandboxInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreTime)) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxInstanceName)) {
		query["SandboxInstanceName"] = request.SandboxInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxPassword)) {
		query["SandboxPassword"] = request.SandboxPassword
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxSpecification)) {
		query["SandboxSpecification"] = request.SandboxSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxType)) {
		query["SandboxType"] = request.SandboxType
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxUser)) {
		query["SandboxUser"] = request.SandboxUser
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcSwitchId)) {
		query["VpcSwitchId"] = request.VpcSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSandboxInstance"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSandboxInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSandboxInstance(request *CreateSandboxInstanceRequest) (_result *CreateSandboxInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSandboxInstanceResponse{}
	_body, _err := client.CreateSandboxInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSandboxInstanceWithOptions(request *DeleteSandboxInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteSandboxInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSandboxInstance"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSandboxInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSandboxInstance(request *DeleteSandboxInstanceRequest) (_result *DeleteSandboxInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSandboxInstanceResponse{}
	_body, _err := client.DeleteSandboxInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSandboxBackupSetsWithOptions(request *DescribeSandboxBackupSetsRequest, runtime *util.RuntimeOptions) (_result *DescribeSandboxBackupSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
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
		Action:      tea.String("DescribeSandboxBackupSets"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSandboxBackupSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSandboxBackupSets(request *DescribeSandboxBackupSetsRequest) (_result *DescribeSandboxBackupSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSandboxBackupSetsResponse{}
	_body, _err := client.DescribeSandboxBackupSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSandboxInstancesWithOptions(request *DescribeSandboxInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeSandboxInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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
		Action:      tea.String("DescribeSandboxInstances"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSandboxInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSandboxInstances(request *DescribeSandboxInstancesRequest) (_result *DescribeSandboxInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSandboxInstancesResponse{}
	_body, _err := client.DescribeSandboxInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSandboxRecoveryTimeWithOptions(request *DescribeSandboxRecoveryTimeRequest, runtime *util.RuntimeOptions) (_result *DescribeSandboxRecoveryTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSandboxRecoveryTime"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSandboxRecoveryTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSandboxRecoveryTime(request *DescribeSandboxRecoveryTimeRequest) (_result *DescribeSandboxRecoveryTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSandboxRecoveryTimeResponse{}
	_body, _err := client.DescribeSandboxRecoveryTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

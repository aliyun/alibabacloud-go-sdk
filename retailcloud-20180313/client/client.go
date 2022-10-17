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

type AddClusterNodeRequest struct {
	ClusterInstanceId *string   `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	EcsInstanceIdList []*string `json:"EcsInstanceIdList,omitempty" xml:"EcsInstanceIdList,omitempty" type:"Repeated"`
}

func (s AddClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *AddClusterNodeRequest) SetClusterInstanceId(v string) *AddClusterNodeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *AddClusterNodeRequest) SetEcsInstanceIdList(v []*string) *AddClusterNodeRequest {
	s.EcsInstanceIdList = v
	return s
}

type AddClusterNodeResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                           `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddClusterNodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddClusterNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddClusterNodeResponseBody) SetCode(v int32) *AddClusterNodeResponseBody {
	s.Code = &v
	return s
}

func (s *AddClusterNodeResponseBody) SetErrMsg(v string) *AddClusterNodeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *AddClusterNodeResponseBody) SetRequestId(v string) *AddClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddClusterNodeResponseBody) SetResult(v *AddClusterNodeResponseBodyResult) *AddClusterNodeResponseBody {
	s.Result = v
	return s
}

func (s *AddClusterNodeResponseBody) SetSuccess(v bool) *AddClusterNodeResponseBody {
	s.Success = &v
	return s
}

type AddClusterNodeResponseBodyResult struct {
	Nonsense *int32 `json:"Nonsense,omitempty" xml:"Nonsense,omitempty"`
}

func (s AddClusterNodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddClusterNodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddClusterNodeResponseBodyResult) SetNonsense(v int32) *AddClusterNodeResponseBodyResult {
	s.Nonsense = &v
	return s
}

type AddClusterNodeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *AddClusterNodeResponse) SetHeaders(v map[string]*string) *AddClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *AddClusterNodeResponse) SetStatusCode(v int32) *AddClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClusterNodeResponse) SetBody(v *AddClusterNodeResponseBody) *AddClusterNodeResponse {
	s.Body = v
	return s
}

type AllocatePodConfigRequest struct {
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId     *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePodConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePodConfigRequest) GoString() string {
	return s.String()
}

func (s *AllocatePodConfigRequest) SetAppId(v int64) *AllocatePodConfigRequest {
	s.AppId = &v
	return s
}

func (s *AllocatePodConfigRequest) SetEnvId(v int64) *AllocatePodConfigRequest {
	s.EnvId = &v
	return s
}

func (s *AllocatePodConfigRequest) SetRequestId(v string) *AllocatePodConfigRequest {
	s.RequestId = &v
	return s
}

type AllocatePodConfigResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AllocatePodConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AllocatePodConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePodConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePodConfigResponseBody) SetCode(v int32) *AllocatePodConfigResponseBody {
	s.Code = &v
	return s
}

func (s *AllocatePodConfigResponseBody) SetErrMsg(v string) *AllocatePodConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *AllocatePodConfigResponseBody) SetRequestId(v string) *AllocatePodConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocatePodConfigResponseBody) SetResult(v *AllocatePodConfigResponseBodyResult) *AllocatePodConfigResponseBody {
	s.Result = v
	return s
}

func (s *AllocatePodConfigResponseBody) SetSuccess(v bool) *AllocatePodConfigResponseBody {
	s.Success = &v
	return s
}

type AllocatePodConfigResponseBodyResult struct {
	ConfigData *string `json:"ConfigData,omitempty" xml:"ConfigData,omitempty"`
}

func (s AllocatePodConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AllocatePodConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AllocatePodConfigResponseBodyResult) SetConfigData(v string) *AllocatePodConfigResponseBodyResult {
	s.ConfigData = &v
	return s
}

type AllocatePodConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocatePodConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocatePodConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePodConfigResponse) GoString() string {
	return s.String()
}

func (s *AllocatePodConfigResponse) SetHeaders(v map[string]*string) *AllocatePodConfigResponse {
	s.Headers = v
	return s
}

func (s *AllocatePodConfigResponse) SetStatusCode(v int32) *AllocatePodConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePodConfigResponse) SetBody(v *AllocatePodConfigResponseBody) *AllocatePodConfigResponse {
	s.Body = v
	return s
}

type BatchAddServersRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sign       *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s BatchAddServersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddServersRequest) GoString() string {
	return s.String()
}

func (s *BatchAddServersRequest) SetInstanceId(v string) *BatchAddServersRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchAddServersRequest) SetRegionId(v string) *BatchAddServersRequest {
	s.RegionId = &v
	return s
}

func (s *BatchAddServersRequest) SetSign(v string) *BatchAddServersRequest {
	s.Sign = &v
	return s
}

func (s *BatchAddServersRequest) SetVpcId(v string) *BatchAddServersRequest {
	s.VpcId = &v
	return s
}

type BatchAddServersResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                            `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *BatchAddServersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s BatchAddServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddServersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddServersResponseBody) SetCode(v int32) *BatchAddServersResponseBody {
	s.Code = &v
	return s
}

func (s *BatchAddServersResponseBody) SetErrMsg(v string) *BatchAddServersResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *BatchAddServersResponseBody) SetRequestId(v string) *BatchAddServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddServersResponseBody) SetResult(v *BatchAddServersResponseBodyResult) *BatchAddServersResponseBody {
	s.Result = v
	return s
}

type BatchAddServersResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchAddServersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchAddServersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchAddServersResponseBodyResult) SetSuccess(v bool) *BatchAddServersResponseBodyResult {
	s.Success = &v
	return s
}

type BatchAddServersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchAddServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddServersResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddServersResponse) GoString() string {
	return s.String()
}

func (s *BatchAddServersResponse) SetHeaders(v map[string]*string) *BatchAddServersResponse {
	s.Headers = v
	return s
}

func (s *BatchAddServersResponse) SetStatusCode(v int32) *BatchAddServersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddServersResponse) SetBody(v *BatchAddServersResponseBody) *BatchAddServersResponse {
	s.Body = v
	return s
}

type BindGroupRequest struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BindGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s BindGroupRequest) GoString() string {
	return s.String()
}

func (s *BindGroupRequest) SetAppId(v int64) *BindGroupRequest {
	s.AppId = &v
	return s
}

func (s *BindGroupRequest) SetBizCode(v string) *BindGroupRequest {
	s.BizCode = &v
	return s
}

func (s *BindGroupRequest) SetName(v string) *BindGroupRequest {
	s.Name = &v
	return s
}

type BindGroupResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                      `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *BindGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s BindGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BindGroupResponseBody) SetCode(v int32) *BindGroupResponseBody {
	s.Code = &v
	return s
}

func (s *BindGroupResponseBody) SetErrMsg(v string) *BindGroupResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *BindGroupResponseBody) SetRequestId(v string) *BindGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindGroupResponseBody) SetResult(v *BindGroupResponseBodyResult) *BindGroupResponseBody {
	s.Result = v
	return s
}

type BindGroupResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BindGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BindGroupResponseBodyResult) SetSuccess(v bool) *BindGroupResponseBodyResult {
	s.Success = &v
	return s
}

type BindGroupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BindGroupResponse) GoString() string {
	return s.String()
}

func (s *BindGroupResponse) SetHeaders(v map[string]*string) *BindGroupResponse {
	s.Headers = v
	return s
}

func (s *BindGroupResponse) SetStatusCode(v int32) *BindGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *BindGroupResponse) SetBody(v *BindGroupResponseBody) *BindGroupResponse {
	s.Body = v
	return s
}

type BindNodeLabelRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s BindNodeLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s BindNodeLabelRequest) GoString() string {
	return s.String()
}

func (s *BindNodeLabelRequest) SetClusterId(v string) *BindNodeLabelRequest {
	s.ClusterId = &v
	return s
}

func (s *BindNodeLabelRequest) SetInstanceId(v string) *BindNodeLabelRequest {
	s.InstanceId = &v
	return s
}

func (s *BindNodeLabelRequest) SetLabelKey(v string) *BindNodeLabelRequest {
	s.LabelKey = &v
	return s
}

func (s *BindNodeLabelRequest) SetLabelValue(v string) *BindNodeLabelRequest {
	s.LabelValue = &v
	return s
}

type BindNodeLabelResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindNodeLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindNodeLabelResponseBody) GoString() string {
	return s.String()
}

func (s *BindNodeLabelResponseBody) SetCode(v int32) *BindNodeLabelResponseBody {
	s.Code = &v
	return s
}

func (s *BindNodeLabelResponseBody) SetErrMsg(v string) *BindNodeLabelResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *BindNodeLabelResponseBody) SetRequestId(v string) *BindNodeLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindNodeLabelResponseBody) SetSuccess(v bool) *BindNodeLabelResponseBody {
	s.Success = &v
	return s
}

type BindNodeLabelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindNodeLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindNodeLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s BindNodeLabelResponse) GoString() string {
	return s.String()
}

func (s *BindNodeLabelResponse) SetHeaders(v map[string]*string) *BindNodeLabelResponse {
	s.Headers = v
	return s
}

func (s *BindNodeLabelResponse) SetStatusCode(v int32) *BindNodeLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *BindNodeLabelResponse) SetBody(v *BindNodeLabelResponseBody) *BindNodeLabelResponse {
	s.Body = v
	return s
}

type CloseDeployOrderRequest struct {
	DeployOrderId *int64 `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s CloseDeployOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDeployOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseDeployOrderRequest) SetDeployOrderId(v int64) *CloseDeployOrderRequest {
	s.DeployOrderId = &v
	return s
}

type CloseDeployOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseDeployOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDeployOrderResponseBody) SetCode(v int32) *CloseDeployOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CloseDeployOrderResponseBody) SetErrMsg(v string) *CloseDeployOrderResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CloseDeployOrderResponseBody) SetRequestId(v string) *CloseDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDeployOrderResponseBody) SetSuccess(v bool) *CloseDeployOrderResponseBody {
	s.Success = &v
	return s
}

type CloseDeployOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseDeployOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseDeployOrderResponse) SetHeaders(v map[string]*string) *CloseDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *CloseDeployOrderResponse) SetStatusCode(v int32) *CloseDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDeployOrderResponse) SetBody(v *CloseDeployOrderResponseBody) *CloseDeployOrderResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DbInstanceId    *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
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

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDbInstanceId(v string) *CreateAccountRequest {
	s.DbInstanceId = &v
	return s
}

type CreateAccountResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetCode(v int32) *CreateAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccountResponseBody) SetErrMsg(v string) *CreateAccountResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
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

type CreateAppRequest struct {
	BizCode          *string                      `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	BizTitle         *string                      `json:"BizTitle,omitempty" xml:"BizTitle,omitempty"`
	Description      *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName        *string                      `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Language         *string                      `json:"Language,omitempty" xml:"Language,omitempty"`
	MiddleWareIdList []*int32                     `json:"MiddleWareIdList,omitempty" xml:"MiddleWareIdList,omitempty" type:"Repeated"`
	Namespace        *string                      `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OperatingSystem  *string                      `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
	ServiceType      *string                      `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	StateType        *int32                       `json:"StateType,omitempty" xml:"StateType,omitempty"`
	Title            *string                      `json:"Title,omitempty" xml:"Title,omitempty"`
	UserRoles        []*CreateAppRequestUserRoles `json:"UserRoles,omitempty" xml:"UserRoles,omitempty" type:"Repeated"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetBizCode(v string) *CreateAppRequest {
	s.BizCode = &v
	return s
}

func (s *CreateAppRequest) SetBizTitle(v string) *CreateAppRequest {
	s.BizTitle = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetGroupName(v string) *CreateAppRequest {
	s.GroupName = &v
	return s
}

func (s *CreateAppRequest) SetLanguage(v string) *CreateAppRequest {
	s.Language = &v
	return s
}

func (s *CreateAppRequest) SetMiddleWareIdList(v []*int32) *CreateAppRequest {
	s.MiddleWareIdList = v
	return s
}

func (s *CreateAppRequest) SetNamespace(v string) *CreateAppRequest {
	s.Namespace = &v
	return s
}

func (s *CreateAppRequest) SetOperatingSystem(v string) *CreateAppRequest {
	s.OperatingSystem = &v
	return s
}

func (s *CreateAppRequest) SetServiceType(v string) *CreateAppRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateAppRequest) SetStateType(v int32) *CreateAppRequest {
	s.StateType = &v
	return s
}

func (s *CreateAppRequest) SetTitle(v string) *CreateAppRequest {
	s.Title = &v
	return s
}

func (s *CreateAppRequest) SetUserRoles(v []*CreateAppRequestUserRoles) *CreateAppRequest {
	s.UserRoles = v
	return s
}

type CreateAppRequestUserRoles struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreateAppRequestUserRoles) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestUserRoles) GoString() string {
	return s.String()
}

func (s *CreateAppRequestUserRoles) SetRoleName(v string) *CreateAppRequestUserRoles {
	s.RoleName = &v
	return s
}

func (s *CreateAppRequestUserRoles) SetUserId(v string) *CreateAppRequestUserRoles {
	s.UserId = &v
	return s
}

func (s *CreateAppRequestUserRoles) SetUserType(v string) *CreateAppRequestUserRoles {
	s.UserType = &v
	return s
}

type CreateAppResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                      `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetCode(v int32) *CreateAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResponseBody) SetErrMsg(v string) *CreateAppResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponseBodyResult struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) SetAppId(v int64) *CreateAppResponseBodyResult {
	s.AppId = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateAppGroupRequest struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) SetBizCode(v string) *CreateAppGroupRequest {
	s.BizCode = &v
	return s
}

func (s *CreateAppGroupRequest) SetName(v string) *CreateAppGroupRequest {
	s.Name = &v
	return s
}

type CreateAppGroupResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                           `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAppGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBody) SetCode(v int32) *CreateAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetErrMsg(v string) *CreateAppGroupResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetRequestId(v string) *CreateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetResult(v *CreateAppGroupResponseBodyResult) *CreateAppGroupResponseBody {
	s.Result = v
	return s
}

type CreateAppGroupResponseBodyResult struct {
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
}

func (s CreateAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResult) SetAppGroupId(v int64) *CreateAppGroupResponseBodyResult {
	s.AppGroupId = &v
	return s
}

type CreateAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponse) SetHeaders(v map[string]*string) *CreateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupResponse) SetStatusCode(v int32) *CreateAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppGroupResponse) SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse {
	s.Body = v
	return s
}

type CreateAppMonitorsRequest struct {
	AlarmTemplateId *int64   `json:"AlarmTemplateId,omitempty" xml:"AlarmTemplateId,omitempty"`
	AppIds          []*int64 `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	EnvType         *int32   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	MainUserId      *int64   `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	SilenceTime     *string  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
}

func (s CreateAppMonitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppMonitorsRequest) GoString() string {
	return s.String()
}

func (s *CreateAppMonitorsRequest) SetAlarmTemplateId(v int64) *CreateAppMonitorsRequest {
	s.AlarmTemplateId = &v
	return s
}

func (s *CreateAppMonitorsRequest) SetAppIds(v []*int64) *CreateAppMonitorsRequest {
	s.AppIds = v
	return s
}

func (s *CreateAppMonitorsRequest) SetEnvType(v int32) *CreateAppMonitorsRequest {
	s.EnvType = &v
	return s
}

func (s *CreateAppMonitorsRequest) SetMainUserId(v int64) *CreateAppMonitorsRequest {
	s.MainUserId = &v
	return s
}

func (s *CreateAppMonitorsRequest) SetSilenceTime(v string) *CreateAppMonitorsRequest {
	s.SilenceTime = &v
	return s
}

type CreateAppMonitorsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppMonitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppMonitorsResponseBody) SetCode(v int32) *CreateAppMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppMonitorsResponseBody) SetRequestId(v string) *CreateAppMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppMonitorsResponseBody) SetSuccess(v bool) *CreateAppMonitorsResponseBody {
	s.Success = &v
	return s
}

type CreateAppMonitorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppMonitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppMonitorsResponse) GoString() string {
	return s.String()
}

func (s *CreateAppMonitorsResponse) SetHeaders(v map[string]*string) *CreateAppMonitorsResponse {
	s.Headers = v
	return s
}

func (s *CreateAppMonitorsResponse) SetStatusCode(v int32) *CreateAppMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppMonitorsResponse) SetBody(v *CreateAppMonitorsResponseBody) *CreateAppMonitorsResponse {
	s.Body = v
	return s
}

type CreateAppResourceAllocRequest struct {
	AppEnvId  *int64  `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateAppResourceAllocRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResourceAllocRequest) GoString() string {
	return s.String()
}

func (s *CreateAppResourceAllocRequest) SetAppEnvId(v int64) *CreateAppResourceAllocRequest {
	s.AppEnvId = &v
	return s
}

func (s *CreateAppResourceAllocRequest) SetAppId(v int64) *CreateAppResourceAllocRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppResourceAllocRequest) SetClusterId(v string) *CreateAppResourceAllocRequest {
	s.ClusterId = &v
	return s
}

type CreateAppResourceAllocResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                   `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAppResourceAllocResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppResourceAllocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResourceAllocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResourceAllocResponseBody) SetCode(v int32) *CreateAppResourceAllocResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResourceAllocResponseBody) SetErrMsg(v string) *CreateAppResourceAllocResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateAppResourceAllocResponseBody) SetRequestId(v string) *CreateAppResourceAllocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResourceAllocResponseBody) SetResult(v *CreateAppResourceAllocResponseBodyResult) *CreateAppResourceAllocResponseBody {
	s.Result = v
	return s
}

func (s *CreateAppResourceAllocResponseBody) SetSuccess(v bool) *CreateAppResourceAllocResponseBody {
	s.Success = &v
	return s
}

type CreateAppResourceAllocResponseBodyResult struct {
	AppEnvId    *int64  `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ResourceDef *string `json:"ResourceDef,omitempty" xml:"ResourceDef,omitempty"`
}

func (s CreateAppResourceAllocResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResourceAllocResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResourceAllocResponseBodyResult) SetAppEnvId(v int64) *CreateAppResourceAllocResponseBodyResult {
	s.AppEnvId = &v
	return s
}

func (s *CreateAppResourceAllocResponseBodyResult) SetAppId(v int64) *CreateAppResourceAllocResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *CreateAppResourceAllocResponseBodyResult) SetClusterId(v string) *CreateAppResourceAllocResponseBodyResult {
	s.ClusterId = &v
	return s
}

func (s *CreateAppResourceAllocResponseBodyResult) SetId(v int64) *CreateAppResourceAllocResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAppResourceAllocResponseBodyResult) SetResourceDef(v string) *CreateAppResourceAllocResponseBodyResult {
	s.ResourceDef = &v
	return s
}

type CreateAppResourceAllocResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppResourceAllocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResourceAllocResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResourceAllocResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResourceAllocResponse) SetHeaders(v map[string]*string) *CreateAppResourceAllocResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResourceAllocResponse) SetStatusCode(v int32) *CreateAppResourceAllocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResourceAllocResponse) SetBody(v *CreateAppResourceAllocResponseBody) *CreateAppResourceAllocResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	BusinessCode              *string   `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	CloudMonitorFlags         *int32    `json:"CloudMonitorFlags,omitempty" xml:"CloudMonitorFlags,omitempty"`
	ClusterEnvType            *string   `json:"ClusterEnvType,omitempty" xml:"ClusterEnvType,omitempty"`
	ClusterId                 *int64    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterTitle              *string   `json:"ClusterTitle,omitempty" xml:"ClusterTitle,omitempty"`
	ClusterType               *string   `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateWithArmsIntegration *bool     `json:"CreateWithArmsIntegration,omitempty" xml:"CreateWithArmsIntegration,omitempty"`
	CreateWithLogIntegration  *bool     `json:"CreateWithLogIntegration,omitempty" xml:"CreateWithLogIntegration,omitempty"`
	KeyPair                   *string   `json:"KeyPair,omitempty" xml:"KeyPair,omitempty"`
	NetPlug                   *string   `json:"NetPlug,omitempty" xml:"NetPlug,omitempty"`
	Password                  *string   `json:"Password,omitempty" xml:"Password,omitempty"`
	PodCIDR                   *string   `json:"PodCIDR,omitempty" xml:"PodCIDR,omitempty"`
	PrivateZone               *bool     `json:"PrivateZone,omitempty" xml:"PrivateZone,omitempty"`
	PublicSlb                 *int32    `json:"PublicSlb,omitempty" xml:"PublicSlb,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName                *string   `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ServiceCIDR               *string   `json:"ServiceCIDR,omitempty" xml:"ServiceCIDR,omitempty"`
	SnatEntry                 *int32    `json:"SnatEntry,omitempty" xml:"SnatEntry,omitempty"`
	VpcId                     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Vswitchids                []*string `json:"Vswitchids,omitempty" xml:"Vswitchids,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetBusinessCode(v string) *CreateClusterRequest {
	s.BusinessCode = &v
	return s
}

func (s *CreateClusterRequest) SetCloudMonitorFlags(v int32) *CreateClusterRequest {
	s.CloudMonitorFlags = &v
	return s
}

func (s *CreateClusterRequest) SetClusterEnvType(v string) *CreateClusterRequest {
	s.ClusterEnvType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterId(v int64) *CreateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterRequest) SetClusterTitle(v string) *CreateClusterRequest {
	s.ClusterTitle = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetCreateWithArmsIntegration(v bool) *CreateClusterRequest {
	s.CreateWithArmsIntegration = &v
	return s
}

func (s *CreateClusterRequest) SetCreateWithLogIntegration(v bool) *CreateClusterRequest {
	s.CreateWithLogIntegration = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPair(v string) *CreateClusterRequest {
	s.KeyPair = &v
	return s
}

func (s *CreateClusterRequest) SetNetPlug(v string) *CreateClusterRequest {
	s.NetPlug = &v
	return s
}

func (s *CreateClusterRequest) SetPassword(v string) *CreateClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateClusterRequest) SetPodCIDR(v string) *CreateClusterRequest {
	s.PodCIDR = &v
	return s
}

func (s *CreateClusterRequest) SetPrivateZone(v bool) *CreateClusterRequest {
	s.PrivateZone = &v
	return s
}

func (s *CreateClusterRequest) SetPublicSlb(v int32) *CreateClusterRequest {
	s.PublicSlb = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetRegionName(v string) *CreateClusterRequest {
	s.RegionName = &v
	return s
}

func (s *CreateClusterRequest) SetServiceCIDR(v string) *CreateClusterRequest {
	s.ServiceCIDR = &v
	return s
}

func (s *CreateClusterRequest) SetSnatEntry(v int32) *CreateClusterRequest {
	s.SnatEntry = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetVswitchids(v []*string) *CreateClusterRequest {
	s.Vswitchids = v
	return s
}

type CreateClusterResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateClusterResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetCode(v int32) *CreateClusterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClusterResponseBody) SetErrMsg(v string) *CreateClusterResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetResult(v *CreateClusterResponseBodyResult) *CreateClusterResponseBody {
	s.Result = v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

type CreateClusterResponseBodyResult struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
}

func (s CreateClusterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBodyResult) SetClusterInstanceId(v string) *CreateClusterResponseBodyResult {
	s.ClusterInstanceId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateDbRequest struct {
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	DbDescription    *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	DbInstanceId     *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbName           *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s CreateDbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDbRequest) GoString() string {
	return s.String()
}

func (s *CreateDbRequest) SetCharacterSetName(v string) *CreateDbRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDbRequest) SetDbDescription(v string) *CreateDbRequest {
	s.DbDescription = &v
	return s
}

func (s *CreateDbRequest) SetDbInstanceId(v string) *CreateDbRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDbRequest) SetDbName(v string) *CreateDbRequest {
	s.DbName = &v
	return s
}

type CreateDbResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDbResponseBody) SetCode(v int32) *CreateDbResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDbResponseBody) SetErrMsg(v string) *CreateDbResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateDbResponseBody) SetRequestId(v string) *CreateDbResponseBody {
	s.RequestId = &v
	return s
}

type CreateDbResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDbResponse) GoString() string {
	return s.String()
}

func (s *CreateDbResponse) SetHeaders(v map[string]*string) *CreateDbResponse {
	s.Headers = v
	return s
}

func (s *CreateDbResponse) SetStatusCode(v int32) *CreateDbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDbResponse) SetBody(v *CreateDbResponseBody) *CreateDbResponse {
	s.Body = v
	return s
}

type CreateDeployConfigRequest struct {
	AppId         *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CodePath      *string   `json:"CodePath,omitempty" xml:"CodePath,omitempty"`
	ConfigMap     *string   `json:"ConfigMap,omitempty" xml:"ConfigMap,omitempty"`
	ConfigMapList []*string `json:"ConfigMapList,omitempty" xml:"ConfigMapList,omitempty" type:"Repeated"`
	CronJob       *string   `json:"CronJob,omitempty" xml:"CronJob,omitempty"`
	Deployment    *string   `json:"Deployment,omitempty" xml:"Deployment,omitempty"`
	EnvType       *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	SecretList    []*string `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Repeated"`
	StatefulSet   *string   `json:"StatefulSet,omitempty" xml:"StatefulSet,omitempty"`
}

func (s CreateDeployConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeployConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateDeployConfigRequest) SetAppId(v int64) *CreateDeployConfigRequest {
	s.AppId = &v
	return s
}

func (s *CreateDeployConfigRequest) SetCodePath(v string) *CreateDeployConfigRequest {
	s.CodePath = &v
	return s
}

func (s *CreateDeployConfigRequest) SetConfigMap(v string) *CreateDeployConfigRequest {
	s.ConfigMap = &v
	return s
}

func (s *CreateDeployConfigRequest) SetConfigMapList(v []*string) *CreateDeployConfigRequest {
	s.ConfigMapList = v
	return s
}

func (s *CreateDeployConfigRequest) SetCronJob(v string) *CreateDeployConfigRequest {
	s.CronJob = &v
	return s
}

func (s *CreateDeployConfigRequest) SetDeployment(v string) *CreateDeployConfigRequest {
	s.Deployment = &v
	return s
}

func (s *CreateDeployConfigRequest) SetEnvType(v string) *CreateDeployConfigRequest {
	s.EnvType = &v
	return s
}

func (s *CreateDeployConfigRequest) SetName(v string) *CreateDeployConfigRequest {
	s.Name = &v
	return s
}

func (s *CreateDeployConfigRequest) SetSecretList(v []*string) *CreateDeployConfigRequest {
	s.SecretList = v
	return s
}

func (s *CreateDeployConfigRequest) SetStatefulSet(v string) *CreateDeployConfigRequest {
	s.StatefulSet = &v
	return s
}

type CreateDeployConfigResponseBody struct {
	Code       *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMessage *string                               `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *CreateDeployConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateDeployConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeployConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeployConfigResponseBody) SetCode(v int32) *CreateDeployConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDeployConfigResponseBody) SetErrMessage(v string) *CreateDeployConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDeployConfigResponseBody) SetRequestId(v string) *CreateDeployConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeployConfigResponseBody) SetResult(v *CreateDeployConfigResponseBodyResult) *CreateDeployConfigResponseBody {
	s.Result = v
	return s
}

type CreateDeployConfigResponseBodyResult struct {
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SchemaId *int64  `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateDeployConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDeployConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDeployConfigResponseBodyResult) SetAppId(v int64) *CreateDeployConfigResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *CreateDeployConfigResponseBodyResult) SetName(v string) *CreateDeployConfigResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateDeployConfigResponseBodyResult) SetSchemaId(v int64) *CreateDeployConfigResponseBodyResult {
	s.SchemaId = &v
	return s
}

type CreateDeployConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeployConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeployConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeployConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateDeployConfigResponse) SetHeaders(v map[string]*string) *CreateDeployConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateDeployConfigResponse) SetStatusCode(v int32) *CreateDeployConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeployConfigResponse) SetBody(v *CreateDeployConfigResponseBody) *CreateDeployConfigResponse {
	s.Body = v
	return s
}

type CreateEciConfigRequest struct {
	AppEnvId                *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	EipBandwidth            *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EnableEciSchedulePolicy *bool  `json:"EnableEciSchedulePolicy,omitempty" xml:"EnableEciSchedulePolicy,omitempty"`
	MirrorCache             *bool  `json:"MirrorCache,omitempty" xml:"MirrorCache,omitempty"`
	NormalInstanceLimit     *int32 `json:"NormalInstanceLimit,omitempty" xml:"NormalInstanceLimit,omitempty"`
	ScheduleVirtualNode     *bool  `json:"ScheduleVirtualNode,omitempty" xml:"ScheduleVirtualNode,omitempty"`
}

func (s CreateEciConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEciConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateEciConfigRequest) SetAppEnvId(v int64) *CreateEciConfigRequest {
	s.AppEnvId = &v
	return s
}

func (s *CreateEciConfigRequest) SetEipBandwidth(v int32) *CreateEciConfigRequest {
	s.EipBandwidth = &v
	return s
}

func (s *CreateEciConfigRequest) SetEnableEciSchedulePolicy(v bool) *CreateEciConfigRequest {
	s.EnableEciSchedulePolicy = &v
	return s
}

func (s *CreateEciConfigRequest) SetMirrorCache(v bool) *CreateEciConfigRequest {
	s.MirrorCache = &v
	return s
}

func (s *CreateEciConfigRequest) SetNormalInstanceLimit(v int32) *CreateEciConfigRequest {
	s.NormalInstanceLimit = &v
	return s
}

func (s *CreateEciConfigRequest) SetScheduleVirtualNode(v bool) *CreateEciConfigRequest {
	s.ScheduleVirtualNode = &v
	return s
}

type CreateEciConfigResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                            `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateEciConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateEciConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEciConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEciConfigResponseBody) SetCode(v int32) *CreateEciConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEciConfigResponseBody) SetErrMsg(v string) *CreateEciConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEciConfigResponseBody) SetRequestId(v string) *CreateEciConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEciConfigResponseBody) SetResult(v *CreateEciConfigResponseBodyResult) *CreateEciConfigResponseBody {
	s.Result = v
	return s
}

type CreateEciConfigResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEciConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateEciConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateEciConfigResponseBodyResult) SetSuccess(v bool) *CreateEciConfigResponseBodyResult {
	s.Success = &v
	return s
}

type CreateEciConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEciConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEciConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEciConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateEciConfigResponse) SetHeaders(v map[string]*string) *CreateEciConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateEciConfigResponse) SetStatusCode(v int32) *CreateEciConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEciConfigResponse) SetBody(v *CreateEciConfigResponseBody) *CreateEciConfigResponse {
	s.Body = v
	return s
}

type CreateEnvironmentRequest struct {
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSchemaId *int64  `json:"AppSchemaId,omitempty" xml:"AppSchemaId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EnvName     *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvType     *int32  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Replicas    *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) SetAppId(v int64) *CreateEnvironmentRequest {
	s.AppId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetAppSchemaId(v int64) *CreateEnvironmentRequest {
	s.AppSchemaId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetClusterId(v string) *CreateEnvironmentRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetEnvName(v string) *CreateEnvironmentRequest {
	s.EnvName = &v
	return s
}

func (s *CreateEnvironmentRequest) SetEnvType(v int32) *CreateEnvironmentRequest {
	s.EnvType = &v
	return s
}

func (s *CreateEnvironmentRequest) SetRegion(v string) *CreateEnvironmentRequest {
	s.Region = &v
	return s
}

func (s *CreateEnvironmentRequest) SetReplicas(v int32) *CreateEnvironmentRequest {
	s.Replicas = &v
	return s
}

type CreateEnvironmentResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateEnvironmentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBody) SetCode(v int32) *CreateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetErrMsg(v string) *CreateEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetRequestId(v string) *CreateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetResult(v *CreateEnvironmentResponseBodyResult) *CreateEnvironmentResponseBody {
	s.Result = v
	return s
}

type CreateEnvironmentResponseBodyResult struct {
	AppEnvId *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
}

func (s CreateEnvironmentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBodyResult) SetAppEnvId(v int64) *CreateEnvironmentResponseBodyResult {
	s.AppEnvId = &v
	return s
}

type CreateEnvironmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponse) SetHeaders(v map[string]*string) *CreateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentResponse) SetStatusCode(v int32) *CreateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvironmentResponse) SetBody(v *CreateEnvironmentResponseBody) *CreateEnvironmentResponse {
	s.Body = v
	return s
}

type CreateNodeLabelRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s CreateNodeLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeLabelRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeLabelRequest) SetClusterId(v string) *CreateNodeLabelRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeLabelRequest) SetLabelKey(v string) *CreateNodeLabelRequest {
	s.LabelKey = &v
	return s
}

func (s *CreateNodeLabelRequest) SetLabelValue(v string) *CreateNodeLabelRequest {
	s.LabelValue = &v
	return s
}

type CreateNodeLabelResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                            `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateNodeLabelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNodeLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeLabelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeLabelResponseBody) SetCode(v int32) *CreateNodeLabelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNodeLabelResponseBody) SetErrMsg(v string) *CreateNodeLabelResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateNodeLabelResponseBody) SetRequestId(v string) *CreateNodeLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeLabelResponseBody) SetResult(v *CreateNodeLabelResponseBodyResult) *CreateNodeLabelResponseBody {
	s.Result = v
	return s
}

func (s *CreateNodeLabelResponseBody) SetSuccess(v bool) *CreateNodeLabelResponseBody {
	s.Success = &v
	return s
}

type CreateNodeLabelResponseBodyResult struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s CreateNodeLabelResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeLabelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateNodeLabelResponseBodyResult) SetClusterId(v string) *CreateNodeLabelResponseBodyResult {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeLabelResponseBodyResult) SetId(v int64) *CreateNodeLabelResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateNodeLabelResponseBodyResult) SetLabelKey(v string) *CreateNodeLabelResponseBodyResult {
	s.LabelKey = &v
	return s
}

func (s *CreateNodeLabelResponseBodyResult) SetLabelValue(v string) *CreateNodeLabelResponseBodyResult {
	s.LabelValue = &v
	return s
}

type CreateNodeLabelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNodeLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNodeLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeLabelResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeLabelResponse) SetHeaders(v map[string]*string) *CreateNodeLabelResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeLabelResponse) SetStatusCode(v int32) *CreateNodeLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeLabelResponse) SetBody(v *CreateNodeLabelResponseBody) *CreateNodeLabelResponse {
	s.Body = v
	return s
}

type CreatePersistentVolumeRequest struct {
	AccessModes       *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	Capacity          *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	MountDir          *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NFSVersion        *string `json:"NFSVersion,omitempty" xml:"NFSVersion,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NasType           *string `json:"NasType,omitempty" xml:"NasType,omitempty"`
	ReclaimPolicy     *string `json:"ReclaimPolicy,omitempty" xml:"ReclaimPolicy,omitempty"`
	StorageClass      *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s CreatePersistentVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeRequest) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeRequest) SetAccessModes(v string) *CreatePersistentVolumeRequest {
	s.AccessModes = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetCapacity(v string) *CreatePersistentVolumeRequest {
	s.Capacity = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetClusterInstanceId(v string) *CreatePersistentVolumeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetMountDir(v string) *CreatePersistentVolumeRequest {
	s.MountDir = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetMountTargetDomain(v string) *CreatePersistentVolumeRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetNFSVersion(v string) *CreatePersistentVolumeRequest {
	s.NFSVersion = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetName(v string) *CreatePersistentVolumeRequest {
	s.Name = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetNasType(v string) *CreatePersistentVolumeRequest {
	s.NasType = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetReclaimPolicy(v string) *CreatePersistentVolumeRequest {
	s.ReclaimPolicy = &v
	return s
}

func (s *CreatePersistentVolumeRequest) SetStorageClass(v string) *CreatePersistentVolumeRequest {
	s.StorageClass = &v
	return s
}

type CreatePersistentVolumeResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                   `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreatePersistentVolumeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreatePersistentVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeResponseBody) SetCode(v int32) *CreatePersistentVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersistentVolumeResponseBody) SetErrMsg(v string) *CreatePersistentVolumeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreatePersistentVolumeResponseBody) SetRequestId(v string) *CreatePersistentVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersistentVolumeResponseBody) SetResult(v *CreatePersistentVolumeResponseBodyResult) *CreatePersistentVolumeResponseBody {
	s.Result = v
	return s
}

type CreatePersistentVolumeResponseBodyResult struct {
	PersistentVolumeId *int64 `json:"PersistentVolumeId,omitempty" xml:"PersistentVolumeId,omitempty"`
}

func (s CreatePersistentVolumeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeResponseBodyResult) SetPersistentVolumeId(v int64) *CreatePersistentVolumeResponseBodyResult {
	s.PersistentVolumeId = &v
	return s
}

type CreatePersistentVolumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePersistentVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePersistentVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeResponse) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeResponse) SetHeaders(v map[string]*string) *CreatePersistentVolumeResponse {
	s.Headers = v
	return s
}

func (s *CreatePersistentVolumeResponse) SetStatusCode(v int32) *CreatePersistentVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersistentVolumeResponse) SetBody(v *CreatePersistentVolumeResponseBody) *CreatePersistentVolumeResponse {
	s.Body = v
	return s
}

type CreatePersistentVolumeClaimRequest struct {
	AccessModes  *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Capacity     *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	EnvId        *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s CreatePersistentVolumeClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeClaimRequest) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeClaimRequest) SetAccessModes(v string) *CreatePersistentVolumeClaimRequest {
	s.AccessModes = &v
	return s
}

func (s *CreatePersistentVolumeClaimRequest) SetAppId(v int64) *CreatePersistentVolumeClaimRequest {
	s.AppId = &v
	return s
}

func (s *CreatePersistentVolumeClaimRequest) SetCapacity(v string) *CreatePersistentVolumeClaimRequest {
	s.Capacity = &v
	return s
}

func (s *CreatePersistentVolumeClaimRequest) SetEnvId(v int64) *CreatePersistentVolumeClaimRequest {
	s.EnvId = &v
	return s
}

func (s *CreatePersistentVolumeClaimRequest) SetName(v string) *CreatePersistentVolumeClaimRequest {
	s.Name = &v
	return s
}

func (s *CreatePersistentVolumeClaimRequest) SetStorageClass(v string) *CreatePersistentVolumeClaimRequest {
	s.StorageClass = &v
	return s
}

type CreatePersistentVolumeClaimResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreatePersistentVolumeClaimResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreatePersistentVolumeClaimResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeClaimResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeClaimResponseBody) SetCode(v int32) *CreatePersistentVolumeClaimResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersistentVolumeClaimResponseBody) SetErrMsg(v string) *CreatePersistentVolumeClaimResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreatePersistentVolumeClaimResponseBody) SetRequestId(v string) *CreatePersistentVolumeClaimResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersistentVolumeClaimResponseBody) SetResult(v *CreatePersistentVolumeClaimResponseBodyResult) *CreatePersistentVolumeClaimResponseBody {
	s.Result = v
	return s
}

type CreatePersistentVolumeClaimResponseBodyResult struct {
	PersistentVolumeClaimId *int64 `json:"PersistentVolumeClaimId,omitempty" xml:"PersistentVolumeClaimId,omitempty"`
}

func (s CreatePersistentVolumeClaimResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeClaimResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeClaimResponseBodyResult) SetPersistentVolumeClaimId(v int64) *CreatePersistentVolumeClaimResponseBodyResult {
	s.PersistentVolumeClaimId = &v
	return s
}

type CreatePersistentVolumeClaimResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePersistentVolumeClaimResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePersistentVolumeClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePersistentVolumeClaimResponse) GoString() string {
	return s.String()
}

func (s *CreatePersistentVolumeClaimResponse) SetHeaders(v map[string]*string) *CreatePersistentVolumeClaimResponse {
	s.Headers = v
	return s
}

func (s *CreatePersistentVolumeClaimResponse) SetStatusCode(v int32) *CreatePersistentVolumeClaimResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersistentVolumeClaimResponse) SetBody(v *CreatePersistentVolumeClaimResponseBody) *CreatePersistentVolumeClaimResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	EnvId        *int64                              `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Headless     *bool                               `json:"Headless,omitempty" xml:"Headless,omitempty"`
	K8sServiceId *string                             `json:"K8sServiceId,omitempty" xml:"K8sServiceId,omitempty"`
	Name         *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PortMappings []*CreateServiceRequestPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	ServiceType  *string                             `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetEnvId(v int64) *CreateServiceRequest {
	s.EnvId = &v
	return s
}

func (s *CreateServiceRequest) SetHeadless(v bool) *CreateServiceRequest {
	s.Headless = &v
	return s
}

func (s *CreateServiceRequest) SetK8sServiceId(v string) *CreateServiceRequest {
	s.K8sServiceId = &v
	return s
}

func (s *CreateServiceRequest) SetName(v string) *CreateServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceRequest) SetPortMappings(v []*CreateServiceRequestPortMappings) *CreateServiceRequest {
	s.PortMappings = v
	return s
}

func (s *CreateServiceRequest) SetServiceType(v string) *CreateServiceRequest {
	s.ServiceType = &v
	return s
}

type CreateServiceRequestPortMappings struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePort   *int32  `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s CreateServiceRequestPortMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestPortMappings) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestPortMappings) SetName(v string) *CreateServiceRequestPortMappings {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestPortMappings) SetNodePort(v int32) *CreateServiceRequestPortMappings {
	s.NodePort = &v
	return s
}

func (s *CreateServiceRequestPortMappings) SetPort(v int32) *CreateServiceRequestPortMappings {
	s.Port = &v
	return s
}

func (s *CreateServiceRequestPortMappings) SetProtocol(v string) *CreateServiceRequestPortMappings {
	s.Protocol = &v
	return s
}

func (s *CreateServiceRequestPortMappings) SetTargetPort(v string) *CreateServiceRequestPortMappings {
	s.TargetPort = &v
	return s
}

type CreateServiceResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateServiceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetCode(v int32) *CreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceResponseBody) SetErrMsg(v string) *CreateServiceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetResult(v *CreateServiceResponseBodyResult) *CreateServiceResponseBody {
	s.Result = v
	return s
}

func (s *CreateServiceResponseBody) SetSuccess(v bool) *CreateServiceResponseBody {
	s.Success = &v
	return s
}

type CreateServiceResponseBodyResult struct {
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyResult) SetServiceId(v int64) *CreateServiceResponseBodyResult {
	s.ServiceId = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateSlbAPRequest struct {
	CookieTimeout      *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	EnvId              *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EstablishedTimeout *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol           *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServerPort     *int32  `json:"RealServerPort,omitempty" xml:"RealServerPort,omitempty"`
	SlbId              *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SslCertId          *string `json:"SslCertId,omitempty" xml:"SslCertId,omitempty"`
	StickySession      *int32  `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
}

func (s CreateSlbAPRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSlbAPRequest) GoString() string {
	return s.String()
}

func (s *CreateSlbAPRequest) SetCookieTimeout(v int32) *CreateSlbAPRequest {
	s.CookieTimeout = &v
	return s
}

func (s *CreateSlbAPRequest) SetEnvId(v int64) *CreateSlbAPRequest {
	s.EnvId = &v
	return s
}

func (s *CreateSlbAPRequest) SetEstablishedTimeout(v int32) *CreateSlbAPRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *CreateSlbAPRequest) SetListenerPort(v int32) *CreateSlbAPRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateSlbAPRequest) SetName(v string) *CreateSlbAPRequest {
	s.Name = &v
	return s
}

func (s *CreateSlbAPRequest) SetProtocol(v string) *CreateSlbAPRequest {
	s.Protocol = &v
	return s
}

func (s *CreateSlbAPRequest) SetRealServerPort(v int32) *CreateSlbAPRequest {
	s.RealServerPort = &v
	return s
}

func (s *CreateSlbAPRequest) SetSlbId(v string) *CreateSlbAPRequest {
	s.SlbId = &v
	return s
}

func (s *CreateSlbAPRequest) SetSslCertId(v string) *CreateSlbAPRequest {
	s.SslCertId = &v
	return s
}

func (s *CreateSlbAPRequest) SetStickySession(v int32) *CreateSlbAPRequest {
	s.StickySession = &v
	return s
}

type CreateSlbAPResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateSlbAPResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSlbAPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSlbAPResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlbAPResponseBody) SetCode(v int32) *CreateSlbAPResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSlbAPResponseBody) SetErrMsg(v string) *CreateSlbAPResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateSlbAPResponseBody) SetRequestId(v string) *CreateSlbAPResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlbAPResponseBody) SetResult(v *CreateSlbAPResponseBodyResult) *CreateSlbAPResponseBody {
	s.Result = v
	return s
}

func (s *CreateSlbAPResponseBody) SetSuccess(v bool) *CreateSlbAPResponseBody {
	s.Success = &v
	return s
}

type CreateSlbAPResponseBodyResult struct {
	SlbAPId *int64 `json:"SlbAPId,omitempty" xml:"SlbAPId,omitempty"`
}

func (s CreateSlbAPResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSlbAPResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSlbAPResponseBodyResult) SetSlbAPId(v int64) *CreateSlbAPResponseBodyResult {
	s.SlbAPId = &v
	return s
}

type CreateSlbAPResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSlbAPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSlbAPResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSlbAPResponse) GoString() string {
	return s.String()
}

func (s *CreateSlbAPResponse) SetHeaders(v map[string]*string) *CreateSlbAPResponse {
	s.Headers = v
	return s
}

func (s *CreateSlbAPResponse) SetStatusCode(v int32) *CreateSlbAPResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlbAPResponse) SetBody(v *CreateSlbAPResponseBody) *CreateSlbAPResponse {
	s.Body = v
	return s
}

type DeleteAppDetailRequest struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Force *bool  `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteAppDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppDetailRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppDetailRequest) SetAppId(v int64) *DeleteAppDetailRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppDetailRequest) SetForce(v bool) *DeleteAppDetailRequest {
	s.Force = &v
	return s
}

type DeleteAppDetailResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                            `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteAppDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteAppDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppDetailResponseBody) SetCode(v int32) *DeleteAppDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppDetailResponseBody) SetErrMsg(v string) *DeleteAppDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteAppDetailResponseBody) SetRequestId(v string) *DeleteAppDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppDetailResponseBody) SetResult(v *DeleteAppDetailResponseBodyResult) *DeleteAppDetailResponseBody {
	s.Result = v
	return s
}

type DeleteAppDetailResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAppDetailResponseBodyResult) SetSuccess(v bool) *DeleteAppDetailResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteAppDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppDetailResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppDetailResponse) SetHeaders(v map[string]*string) *DeleteAppDetailResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppDetailResponse) SetStatusCode(v int32) *DeleteAppDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppDetailResponse) SetBody(v *DeleteAppDetailResponseBody) *DeleteAppDetailResponse {
	s.Body = v
	return s
}

type DeleteAppEnvironmentRequest struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Force *bool  `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteAppEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvironmentRequest) SetAppId(v int64) *DeleteAppEnvironmentRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppEnvironmentRequest) SetEnvId(v int64) *DeleteAppEnvironmentRequest {
	s.EnvId = &v
	return s
}

func (s *DeleteAppEnvironmentRequest) SetForce(v bool) *DeleteAppEnvironmentRequest {
	s.Force = &v
	return s
}

type DeleteAppEnvironmentResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                 `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteAppEnvironmentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteAppEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvironmentResponseBody) SetCode(v int32) *DeleteAppEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppEnvironmentResponseBody) SetErrMsg(v string) *DeleteAppEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteAppEnvironmentResponseBody) SetRequestId(v string) *DeleteAppEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppEnvironmentResponseBody) SetResult(v *DeleteAppEnvironmentResponseBodyResult) *DeleteAppEnvironmentResponseBody {
	s.Result = v
	return s
}

type DeleteAppEnvironmentResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppEnvironmentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvironmentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvironmentResponseBodyResult) SetSuccess(v bool) *DeleteAppEnvironmentResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteAppEnvironmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvironmentResponse) SetHeaders(v map[string]*string) *DeleteAppEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppEnvironmentResponse) SetStatusCode(v int32) *DeleteAppEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppEnvironmentResponse) SetBody(v *DeleteAppEnvironmentResponseBody) *DeleteAppEnvironmentResponse {
	s.Body = v
	return s
}

type DeleteAppGroupRequest struct {
	Force   *bool  `json:"Force,omitempty" xml:"Force,omitempty"`
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupRequest) SetForce(v bool) *DeleteAppGroupRequest {
	s.Force = &v
	return s
}

func (s *DeleteAppGroupRequest) SetGroupId(v int64) *DeleteAppGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteAppGroupResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                           `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteAppGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponseBody) SetCode(v int32) *DeleteAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetErrMsg(v string) *DeleteAppGroupResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetRequestId(v string) *DeleteAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetResult(v *DeleteAppGroupResponseBodyResult) *DeleteAppGroupResponseBody {
	s.Result = v
	return s
}

type DeleteAppGroupResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponseBodyResult) SetSuccess(v bool) *DeleteAppGroupResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponse) SetHeaders(v map[string]*string) *DeleteAppGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppGroupResponse) SetStatusCode(v int32) *DeleteAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppGroupResponse) SetBody(v *DeleteAppGroupResponseBody) *DeleteAppGroupResponse {
	s.Body = v
	return s
}

type DeleteAppResourceAllocRequest struct {
	AppEnvId *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
}

func (s DeleteAppResourceAllocRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResourceAllocRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppResourceAllocRequest) SetAppEnvId(v int64) *DeleteAppResourceAllocRequest {
	s.AppEnvId = &v
	return s
}

type DeleteAppResourceAllocResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppResourceAllocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResourceAllocResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResourceAllocResponseBody) SetCode(v int32) *DeleteAppResourceAllocResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResourceAllocResponseBody) SetErrMsg(v string) *DeleteAppResourceAllocResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteAppResourceAllocResponseBody) SetRequestId(v string) *DeleteAppResourceAllocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResourceAllocResponseBody) SetSuccess(v bool) *DeleteAppResourceAllocResponseBody {
	s.Success = &v
	return s
}

type DeleteAppResourceAllocResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppResourceAllocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppResourceAllocResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResourceAllocResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResourceAllocResponse) SetHeaders(v map[string]*string) *DeleteAppResourceAllocResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResourceAllocResponse) SetStatusCode(v int32) *DeleteAppResourceAllocResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResourceAllocResponse) SetBody(v *DeleteAppResourceAllocResponseBody) *DeleteAppResourceAllocResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterInstanceId(v string) *DeleteClusterRequest {
	s.ClusterInstanceId = &v
	return s
}

type DeleteClusterResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteClusterResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetCode(v int32) *DeleteClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterResponseBody) SetErrMsg(v string) *DeleteClusterResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetResult(v *DeleteClusterResponseBodyResult) *DeleteClusterResponseBody {
	s.Result = v
	return s
}

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

type DeleteClusterResponseBodyResult struct {
	Nonsense *int32 `json:"Nonsense,omitempty" xml:"Nonsense,omitempty"`
}

func (s DeleteClusterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBodyResult) SetNonsense(v int32) *DeleteClusterResponseBodyResult {
	s.Nonsense = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetDBInstanceId(v string) *DeleteDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetCode(v int32) *DeleteDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetErrMsg(v string) *DeleteDatabaseResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetStatusCode(v int32) *DeleteDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

type DeleteDeployConfigRequest struct {
	SchemaId *int64 `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s DeleteDeployConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeployConfigRequest) SetSchemaId(v int64) *DeleteDeployConfigRequest {
	s.SchemaId = &v
	return s
}

type DeleteDeployConfigResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                               `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteDeployConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteDeployConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeployConfigResponseBody) SetCode(v int32) *DeleteDeployConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDeployConfigResponseBody) SetErrMsg(v string) *DeleteDeployConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteDeployConfigResponseBody) SetRequestId(v string) *DeleteDeployConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeployConfigResponseBody) SetResult(v *DeleteDeployConfigResponseBodyResult) *DeleteDeployConfigResponseBody {
	s.Result = v
	return s
}

type DeleteDeployConfigResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDeployConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteDeployConfigResponseBodyResult) SetSuccess(v bool) *DeleteDeployConfigResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteDeployConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeployConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeployConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeployConfigResponse) SetHeaders(v map[string]*string) *DeleteDeployConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeployConfigResponse) SetStatusCode(v int32) *DeleteDeployConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeployConfigResponse) SetBody(v *DeleteDeployConfigResponseBody) *DeleteDeployConfigResponse {
	s.Body = v
	return s
}

type DeleteNodeLabelRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Force      *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s DeleteNodeLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeLabelRequest) SetClusterId(v string) *DeleteNodeLabelRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeLabelRequest) SetForce(v bool) *DeleteNodeLabelRequest {
	s.Force = &v
	return s
}

func (s *DeleteNodeLabelRequest) SetLabelKey(v string) *DeleteNodeLabelRequest {
	s.LabelKey = &v
	return s
}

func (s *DeleteNodeLabelRequest) SetLabelValue(v string) *DeleteNodeLabelRequest {
	s.LabelValue = &v
	return s
}

type DeleteNodeLabelResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeLabelResponseBody) SetCode(v int32) *DeleteNodeLabelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNodeLabelResponseBody) SetErrMsg(v string) *DeleteNodeLabelResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteNodeLabelResponseBody) SetRequestId(v string) *DeleteNodeLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeLabelResponseBody) SetSuccess(v bool) *DeleteNodeLabelResponseBody {
	s.Success = &v
	return s
}

type DeleteNodeLabelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNodeLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodeLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeLabelResponse) SetHeaders(v map[string]*string) *DeleteNodeLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeLabelResponse) SetStatusCode(v int32) *DeleteNodeLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeLabelResponse) SetBody(v *DeleteNodeLabelResponseBody) *DeleteNodeLabelResponse {
	s.Body = v
	return s
}

type DeletePersistentVolumeRequest struct {
	ClusterInstanceId    *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	PersistentVolumeName *string `json:"PersistentVolumeName,omitempty" xml:"PersistentVolumeName,omitempty"`
}

func (s DeletePersistentVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeRequest) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeRequest) SetClusterInstanceId(v string) *DeletePersistentVolumeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *DeletePersistentVolumeRequest) SetPersistentVolumeName(v string) *DeletePersistentVolumeRequest {
	s.PersistentVolumeName = &v
	return s
}

type DeletePersistentVolumeResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                   `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeletePersistentVolumeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeletePersistentVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeResponseBody) SetCode(v int32) *DeletePersistentVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePersistentVolumeResponseBody) SetErrMsg(v string) *DeletePersistentVolumeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeletePersistentVolumeResponseBody) SetRequestId(v string) *DeletePersistentVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePersistentVolumeResponseBody) SetResult(v *DeletePersistentVolumeResponseBodyResult) *DeletePersistentVolumeResponseBody {
	s.Result = v
	return s
}

type DeletePersistentVolumeResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePersistentVolumeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeResponseBodyResult) SetSuccess(v bool) *DeletePersistentVolumeResponseBodyResult {
	s.Success = &v
	return s
}

type DeletePersistentVolumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePersistentVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePersistentVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeResponse) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeResponse) SetHeaders(v map[string]*string) *DeletePersistentVolumeResponse {
	s.Headers = v
	return s
}

func (s *DeletePersistentVolumeResponse) SetStatusCode(v int32) *DeletePersistentVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePersistentVolumeResponse) SetBody(v *DeletePersistentVolumeResponseBody) *DeletePersistentVolumeResponse {
	s.Body = v
	return s
}

type DeletePersistentVolumeClaimRequest struct {
	AppId                     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId                     *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PersistentVolumeClaimName *string `json:"PersistentVolumeClaimName,omitempty" xml:"PersistentVolumeClaimName,omitempty"`
}

func (s DeletePersistentVolumeClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeClaimRequest) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeClaimRequest) SetAppId(v int64) *DeletePersistentVolumeClaimRequest {
	s.AppId = &v
	return s
}

func (s *DeletePersistentVolumeClaimRequest) SetEnvId(v int64) *DeletePersistentVolumeClaimRequest {
	s.EnvId = &v
	return s
}

func (s *DeletePersistentVolumeClaimRequest) SetPersistentVolumeClaimName(v string) *DeletePersistentVolumeClaimRequest {
	s.PersistentVolumeClaimName = &v
	return s
}

type DeletePersistentVolumeClaimResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeletePersistentVolumeClaimResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeletePersistentVolumeClaimResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeClaimResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeClaimResponseBody) SetCode(v int32) *DeletePersistentVolumeClaimResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePersistentVolumeClaimResponseBody) SetErrMsg(v string) *DeletePersistentVolumeClaimResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeletePersistentVolumeClaimResponseBody) SetRequestId(v string) *DeletePersistentVolumeClaimResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePersistentVolumeClaimResponseBody) SetResult(v *DeletePersistentVolumeClaimResponseBodyResult) *DeletePersistentVolumeClaimResponseBody {
	s.Result = v
	return s
}

type DeletePersistentVolumeClaimResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePersistentVolumeClaimResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeClaimResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeClaimResponseBodyResult) SetSuccess(v bool) *DeletePersistentVolumeClaimResponseBodyResult {
	s.Success = &v
	return s
}

type DeletePersistentVolumeClaimResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePersistentVolumeClaimResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePersistentVolumeClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePersistentVolumeClaimResponse) GoString() string {
	return s.String()
}

func (s *DeletePersistentVolumeClaimResponse) SetHeaders(v map[string]*string) *DeletePersistentVolumeClaimResponse {
	s.Headers = v
	return s
}

func (s *DeletePersistentVolumeClaimResponse) SetStatusCode(v int32) *DeletePersistentVolumeClaimResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePersistentVolumeClaimResponse) SetBody(v *DeletePersistentVolumeClaimResponseBody) *DeletePersistentVolumeClaimResponse {
	s.Body = v
	return s
}

type DeleteRdsAccountRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s DeleteRdsAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdsAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteRdsAccountRequest) SetAccountName(v string) *DeleteRdsAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteRdsAccountRequest) SetDbInstanceId(v string) *DeleteRdsAccountRequest {
	s.DbInstanceId = &v
	return s
}

type DeleteRdsAccountResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                             `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteRdsAccountResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteRdsAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdsAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRdsAccountResponseBody) SetCode(v int32) *DeleteRdsAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRdsAccountResponseBody) SetErrMsg(v string) *DeleteRdsAccountResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteRdsAccountResponseBody) SetRequestId(v string) *DeleteRdsAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRdsAccountResponseBody) SetResult(v *DeleteRdsAccountResponseBodyResult) *DeleteRdsAccountResponseBody {
	s.Result = v
	return s
}

type DeleteRdsAccountResponseBodyResult struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRdsAccountResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdsAccountResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRdsAccountResponseBodyResult) SetRequestId(v string) *DeleteRdsAccountResponseBodyResult {
	s.RequestId = &v
	return s
}

type DeleteRdsAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRdsAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRdsAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdsAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteRdsAccountResponse) SetHeaders(v map[string]*string) *DeleteRdsAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteRdsAccountResponse) SetStatusCode(v int32) *DeleteRdsAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRdsAccountResponse) SetBody(v *DeleteRdsAccountResponseBody) *DeleteRdsAccountResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetServiceId(v int64) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

type DeleteServiceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetCode(v int32) *DeleteServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceResponseBody) SetErrMsg(v string) *DeleteServiceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceResponseBody) SetSuccess(v bool) *DeleteServiceResponseBody {
	s.Success = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteSlbAPRequest struct {
	SlbAPId *int64 `json:"SlbAPId,omitempty" xml:"SlbAPId,omitempty"`
}

func (s DeleteSlbAPRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSlbAPRequest) GoString() string {
	return s.String()
}

func (s *DeleteSlbAPRequest) SetSlbAPId(v int64) *DeleteSlbAPRequest {
	s.SlbAPId = &v
	return s
}

type DeleteSlbAPResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSlbAPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSlbAPResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSlbAPResponseBody) SetCode(v int32) *DeleteSlbAPResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSlbAPResponseBody) SetErrMsg(v string) *DeleteSlbAPResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteSlbAPResponseBody) SetRequestId(v string) *DeleteSlbAPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSlbAPResponseBody) SetSuccess(v bool) *DeleteSlbAPResponseBody {
	s.Success = &v
	return s
}

type DeleteSlbAPResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSlbAPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSlbAPResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSlbAPResponse) GoString() string {
	return s.String()
}

func (s *DeleteSlbAPResponse) SetHeaders(v map[string]*string) *DeleteSlbAPResponse {
	s.Headers = v
	return s
}

func (s *DeleteSlbAPResponse) SetStatusCode(v int32) *DeleteSlbAPResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSlbAPResponse) SetBody(v *DeleteSlbAPResponseBody) *DeleteSlbAPResponse {
	s.Body = v
	return s
}

type DeployAppRequest struct {
	ArmsFlag                *bool     `json:"ArmsFlag,omitempty" xml:"ArmsFlag,omitempty"`
	ContainerImageList      []*string `json:"ContainerImageList,omitempty" xml:"ContainerImageList,omitempty" type:"Repeated"`
	DefaultPacketOfAppGroup *string   `json:"DefaultPacketOfAppGroup,omitempty" xml:"DefaultPacketOfAppGroup,omitempty"`
	DeployPacketId          *int64    `json:"DeployPacketId,omitempty" xml:"DeployPacketId,omitempty"`
	DeployPacketUrl         *string   `json:"DeployPacketUrl,omitempty" xml:"DeployPacketUrl,omitempty"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvId                   *int64    `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	InitContainerImageList  []*string `json:"InitContainerImageList,omitempty" xml:"InitContainerImageList,omitempty" type:"Repeated"`
	Name                    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PauseType               *string   `json:"PauseType,omitempty" xml:"PauseType,omitempty"`
	TotalPartitions         *int32    `json:"TotalPartitions,omitempty" xml:"TotalPartitions,omitempty"`
	UpdateStrategyType      *string   `json:"UpdateStrategyType,omitempty" xml:"UpdateStrategyType,omitempty"`
}

func (s DeployAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployAppRequest) GoString() string {
	return s.String()
}

func (s *DeployAppRequest) SetArmsFlag(v bool) *DeployAppRequest {
	s.ArmsFlag = &v
	return s
}

func (s *DeployAppRequest) SetContainerImageList(v []*string) *DeployAppRequest {
	s.ContainerImageList = v
	return s
}

func (s *DeployAppRequest) SetDefaultPacketOfAppGroup(v string) *DeployAppRequest {
	s.DefaultPacketOfAppGroup = &v
	return s
}

func (s *DeployAppRequest) SetDeployPacketId(v int64) *DeployAppRequest {
	s.DeployPacketId = &v
	return s
}

func (s *DeployAppRequest) SetDeployPacketUrl(v string) *DeployAppRequest {
	s.DeployPacketUrl = &v
	return s
}

func (s *DeployAppRequest) SetDescription(v string) *DeployAppRequest {
	s.Description = &v
	return s
}

func (s *DeployAppRequest) SetEnvId(v int64) *DeployAppRequest {
	s.EnvId = &v
	return s
}

func (s *DeployAppRequest) SetInitContainerImageList(v []*string) *DeployAppRequest {
	s.InitContainerImageList = v
	return s
}

func (s *DeployAppRequest) SetName(v string) *DeployAppRequest {
	s.Name = &v
	return s
}

func (s *DeployAppRequest) SetPauseType(v string) *DeployAppRequest {
	s.PauseType = &v
	return s
}

func (s *DeployAppRequest) SetTotalPartitions(v int32) *DeployAppRequest {
	s.TotalPartitions = &v
	return s
}

func (s *DeployAppRequest) SetUpdateStrategyType(v string) *DeployAppRequest {
	s.UpdateStrategyType = &v
	return s
}

type DeployAppResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                      `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeployAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeployAppResponseBody) SetCode(v int32) *DeployAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeployAppResponseBody) SetErrMsg(v string) *DeployAppResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeployAppResponseBody) SetRequestId(v string) *DeployAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployAppResponseBody) SetResult(v *DeployAppResponseBodyResult) *DeployAppResponseBody {
	s.Result = v
	return s
}

func (s *DeployAppResponseBody) SetSuccess(v bool) *DeployAppResponseBody {
	s.Success = &v
	return s
}

type DeployAppResponseBodyResult struct {
	Admitted      *bool   `json:"Admitted,omitempty" xml:"Admitted,omitempty"`
	BusinessCode  *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	DeployOrderId *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s DeployAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeployAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeployAppResponseBodyResult) SetAdmitted(v bool) *DeployAppResponseBodyResult {
	s.Admitted = &v
	return s
}

func (s *DeployAppResponseBodyResult) SetBusinessCode(v string) *DeployAppResponseBodyResult {
	s.BusinessCode = &v
	return s
}

func (s *DeployAppResponseBodyResult) SetDeployOrderId(v int64) *DeployAppResponseBodyResult {
	s.DeployOrderId = &v
	return s
}

type DeployAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployAppResponse) GoString() string {
	return s.String()
}

func (s *DeployAppResponse) SetHeaders(v map[string]*string) *DeployAppResponse {
	s.Headers = v
	return s
}

func (s *DeployAppResponse) SetStatusCode(v int32) *DeployAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployAppResponse) SetBody(v *DeployAppResponseBody) *DeployAppResponse {
	s.Body = v
	return s
}

type DescribeAppDetailRequest struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAppDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppDetailRequest) SetAppId(v int64) *DescribeAppDetailRequest {
	s.AppId = &v
	return s
}

type DescribeAppDetailResponseBody struct {
	Code       *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMessage *string                              `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *DescribeAppDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAppDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppDetailResponseBody) SetCode(v int64) *DescribeAppDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppDetailResponseBody) SetErrMessage(v string) *DescribeAppDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeAppDetailResponseBody) SetRequestId(v string) *DescribeAppDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppDetailResponseBody) SetResult(v *DescribeAppDetailResponseBodyResult) *DescribeAppDetailResponseBody {
	s.Result = v
	return s
}

type DescribeAppDetailResponseBodyResult struct {
	AppId              *int64                                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppStateType       *string                                                  `json:"AppStateType,omitempty" xml:"AppStateType,omitempty"`
	BizName            *string                                                  `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizTitle           *string                                                  `json:"BizTitle,omitempty" xml:"BizTitle,omitempty"`
	DeployType         *string                                                  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Description        *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Language           *string                                                  `json:"Language,omitempty" xml:"Language,omitempty"`
	MiddleWareInfoList []*DescribeAppDetailResponseBodyResultMiddleWareInfoList `json:"MiddleWareInfoList,omitempty" xml:"MiddleWareInfoList,omitempty" type:"Repeated"`
	OperatingSystem    *string                                                  `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
	ServiceType        *string                                                  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Title              *string                                                  `json:"Title,omitempty" xml:"Title,omitempty"`
	UserRoles          []*DescribeAppDetailResponseBodyResultUserRoles          `json:"UserRoles,omitempty" xml:"UserRoles,omitempty" type:"Repeated"`
}

func (s DescribeAppDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppDetailResponseBodyResult) SetAppId(v int64) *DescribeAppDetailResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetAppStateType(v string) *DescribeAppDetailResponseBodyResult {
	s.AppStateType = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetBizName(v string) *DescribeAppDetailResponseBodyResult {
	s.BizName = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetBizTitle(v string) *DescribeAppDetailResponseBodyResult {
	s.BizTitle = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetDeployType(v string) *DescribeAppDetailResponseBodyResult {
	s.DeployType = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetDescription(v string) *DescribeAppDetailResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetLanguage(v string) *DescribeAppDetailResponseBodyResult {
	s.Language = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetMiddleWareInfoList(v []*DescribeAppDetailResponseBodyResultMiddleWareInfoList) *DescribeAppDetailResponseBodyResult {
	s.MiddleWareInfoList = v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetOperatingSystem(v string) *DescribeAppDetailResponseBodyResult {
	s.OperatingSystem = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetServiceType(v string) *DescribeAppDetailResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetTitle(v string) *DescribeAppDetailResponseBodyResult {
	s.Title = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResult) SetUserRoles(v []*DescribeAppDetailResponseBodyResultUserRoles) *DescribeAppDetailResponseBodyResult {
	s.UserRoles = v
	return s
}

type DescribeAppDetailResponseBodyResultMiddleWareInfoList struct {
	AppId *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code  *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAppDetailResponseBodyResultMiddleWareInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDetailResponseBodyResultMiddleWareInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAppDetailResponseBodyResultMiddleWareInfoList) SetAppId(v int64) *DescribeAppDetailResponseBodyResultMiddleWareInfoList {
	s.AppId = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResultMiddleWareInfoList) SetCode(v int32) *DescribeAppDetailResponseBodyResultMiddleWareInfoList {
	s.Code = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResultMiddleWareInfoList) SetName(v string) *DescribeAppDetailResponseBodyResultMiddleWareInfoList {
	s.Name = &v
	return s
}

type DescribeAppDetailResponseBodyResultUserRoles struct {
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeAppDetailResponseBodyResultUserRoles) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDetailResponseBodyResultUserRoles) GoString() string {
	return s.String()
}

func (s *DescribeAppDetailResponseBodyResultUserRoles) SetRealName(v string) *DescribeAppDetailResponseBodyResultUserRoles {
	s.RealName = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResultUserRoles) SetRoleName(v string) *DescribeAppDetailResponseBodyResultUserRoles {
	s.RoleName = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResultUserRoles) SetUserId(v string) *DescribeAppDetailResponseBodyResultUserRoles {
	s.UserId = &v
	return s
}

func (s *DescribeAppDetailResponseBodyResultUserRoles) SetUserType(v string) *DescribeAppDetailResponseBodyResultUserRoles {
	s.UserType = &v
	return s
}

type DescribeAppDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppDetailResponse) SetHeaders(v map[string]*string) *DescribeAppDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppDetailResponse) SetStatusCode(v int32) *DescribeAppDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppDetailResponse) SetBody(v *DescribeAppDetailResponseBody) *DescribeAppDetailResponse {
	s.Body = v
	return s
}

type DescribeAppEnvDeployBaselineRequest struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DescribeAppEnvDeployBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvDeployBaselineRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvDeployBaselineRequest) SetAppId(v int64) *DescribeAppEnvDeployBaselineRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineRequest) SetEnvId(v int64) *DescribeAppEnvDeployBaselineRequest {
	s.EnvId = &v
	return s
}

type DescribeAppEnvDeployBaselineResponseBody struct {
	Code      *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                         `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAppEnvDeployBaselineResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAppEnvDeployBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvDeployBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvDeployBaselineResponseBody) SetCode(v int32) *DescribeAppEnvDeployBaselineResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBody) SetErrMsg(v string) *DescribeAppEnvDeployBaselineResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBody) SetRequestId(v string) *DescribeAppEnvDeployBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBody) SetResult(v *DescribeAppEnvDeployBaselineResponseBodyResult) *DescribeAppEnvDeployBaselineResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBody) SetSuccess(v bool) *DescribeAppEnvDeployBaselineResponseBody {
	s.Success = &v
	return s
}

type DescribeAppEnvDeployBaselineResponseBodyResult struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnvId         *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PacketComment *string `json:"PacketComment,omitempty" xml:"PacketComment,omitempty"`
	PacketId      *int64  `json:"PacketId,omitempty" xml:"PacketId,omitempty"`
	PacketUrl     *string `json:"PacketUrl,omitempty" xml:"PacketUrl,omitempty"`
	SchemaId      *int64  `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s DescribeAppEnvDeployBaselineResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvDeployBaselineResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetAppId(v int64) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetCreateTime(v string) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetEnvId(v int64) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.EnvId = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetPacketComment(v string) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.PacketComment = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetPacketId(v int64) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.PacketId = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetPacketUrl(v string) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.PacketUrl = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponseBodyResult) SetSchemaId(v int64) *DescribeAppEnvDeployBaselineResponseBodyResult {
	s.SchemaId = &v
	return s
}

type DescribeAppEnvDeployBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppEnvDeployBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppEnvDeployBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvDeployBaselineResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvDeployBaselineResponse) SetHeaders(v map[string]*string) *DescribeAppEnvDeployBaselineResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponse) SetStatusCode(v int32) *DescribeAppEnvDeployBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppEnvDeployBaselineResponse) SetBody(v *DescribeAppEnvDeployBaselineResponseBody) *DescribeAppEnvDeployBaselineResponse {
	s.Body = v
	return s
}

type DescribeAppEnvironmentDetailRequest struct {
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DescribeAppEnvironmentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvironmentDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvironmentDetailRequest) SetAppId(v int64) *DescribeAppEnvironmentDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppEnvironmentDetailRequest) SetEnvId(v int64) *DescribeAppEnvironmentDetailRequest {
	s.EnvId = &v
	return s
}

type DescribeAppEnvironmentDetailResponseBody struct {
	Code      *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                         `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAppEnvironmentDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAppEnvironmentDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvironmentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvironmentDetailResponseBody) SetCode(v int32) *DescribeAppEnvironmentDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBody) SetErrMsg(v string) *DescribeAppEnvironmentDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBody) SetRequestId(v string) *DescribeAppEnvironmentDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBody) SetResult(v *DescribeAppEnvironmentDetailResponseBodyResult) *DescribeAppEnvironmentDetailResponseBody {
	s.Result = v
	return s
}

type DescribeAppEnvironmentDetailResponseBodyResult struct {
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSchemaId *int64  `json:"AppSchemaId,omitempty" xml:"AppSchemaId,omitempty"`
	EnvId       *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvName     *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvType     *int32  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	EnvTypeName *string `json:"EnvTypeName,omitempty" xml:"EnvTypeName,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Replicas    *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s DescribeAppEnvironmentDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvironmentDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetAppId(v int64) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetAppSchemaId(v int64) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.AppSchemaId = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetEnvId(v int64) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.EnvId = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetEnvName(v string) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.EnvName = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetEnvType(v int32) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.EnvType = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetEnvTypeName(v string) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.EnvTypeName = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetRegion(v string) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponseBodyResult) SetReplicas(v int32) *DescribeAppEnvironmentDetailResponseBodyResult {
	s.Replicas = &v
	return s
}

type DescribeAppEnvironmentDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppEnvironmentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppEnvironmentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvironmentDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvironmentDetailResponse) SetHeaders(v map[string]*string) *DescribeAppEnvironmentDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppEnvironmentDetailResponse) SetStatusCode(v int32) *DescribeAppEnvironmentDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppEnvironmentDetailResponse) SetBody(v *DescribeAppEnvironmentDetailResponseBody) *DescribeAppEnvironmentDetailResponse {
	s.Body = v
	return s
}

type DescribeAppGroupDeploySettingRequest struct {
	AppGroupId *int64  `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	EnvType    *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s DescribeAppGroupDeploySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDeploySettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDeploySettingRequest) SetAppGroupId(v int64) *DescribeAppGroupDeploySettingRequest {
	s.AppGroupId = &v
	return s
}

func (s *DescribeAppGroupDeploySettingRequest) SetEnvType(v string) *DescribeAppGroupDeploySettingRequest {
	s.EnvType = &v
	return s
}

type DescribeAppGroupDeploySettingResponseBody struct {
	Code      *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAppGroupDeploySettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAppGroupDeploySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDeploySettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDeploySettingResponseBody) SetCode(v int32) *DescribeAppGroupDeploySettingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppGroupDeploySettingResponseBody) SetErrMsg(v string) *DescribeAppGroupDeploySettingResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeAppGroupDeploySettingResponseBody) SetRequestId(v string) *DescribeAppGroupDeploySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppGroupDeploySettingResponseBody) SetResult(v *DescribeAppGroupDeploySettingResponseBodyResult) *DescribeAppGroupDeploySettingResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppGroupDeploySettingResponseBody) SetSuccess(v bool) *DescribeAppGroupDeploySettingResponseBody {
	s.Success = &v
	return s
}

type DescribeAppGroupDeploySettingResponseBodyResult struct {
	DefaultPacketComment *string `json:"DefaultPacketComment,omitempty" xml:"DefaultPacketComment,omitempty"`
	DefaultPacketId      *int64  `json:"DefaultPacketId,omitempty" xml:"DefaultPacketId,omitempty"`
}

func (s DescribeAppGroupDeploySettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDeploySettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDeploySettingResponseBodyResult) SetDefaultPacketComment(v string) *DescribeAppGroupDeploySettingResponseBodyResult {
	s.DefaultPacketComment = &v
	return s
}

func (s *DescribeAppGroupDeploySettingResponseBodyResult) SetDefaultPacketId(v int64) *DescribeAppGroupDeploySettingResponseBodyResult {
	s.DefaultPacketId = &v
	return s
}

type DescribeAppGroupDeploySettingResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppGroupDeploySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppGroupDeploySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupDeploySettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupDeploySettingResponse) SetHeaders(v map[string]*string) *DescribeAppGroupDeploySettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppGroupDeploySettingResponse) SetStatusCode(v int32) *DescribeAppGroupDeploySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppGroupDeploySettingResponse) SetBody(v *DescribeAppGroupDeploySettingResponseBody) *DescribeAppGroupDeploySettingResponse {
	s.Body = v
	return s
}

type DescribeAppMonitorMetricRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeployOrderId *string `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvId         *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Metric        *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	PodName       *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppMonitorMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppMonitorMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppMonitorMetricRequest) SetAppId(v int64) *DescribeAppMonitorMetricRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetDeployOrderId(v string) *DescribeAppMonitorMetricRequest {
	s.DeployOrderId = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetEndTime(v int64) *DescribeAppMonitorMetricRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetEnvId(v int64) *DescribeAppMonitorMetricRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetMetric(v string) *DescribeAppMonitorMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetPodName(v string) *DescribeAppMonitorMetricRequest {
	s.PodName = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetStartTime(v int64) *DescribeAppMonitorMetricRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAppMonitorMetricRequest) SetType(v string) *DescribeAppMonitorMetricRequest {
	s.Type = &v
	return s
}

type DescribeAppMonitorMetricResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                       `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeAppMonitorMetricResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAppMonitorMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppMonitorMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppMonitorMetricResponseBody) SetCode(v int32) *DescribeAppMonitorMetricResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppMonitorMetricResponseBody) SetErrMsg(v string) *DescribeAppMonitorMetricResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeAppMonitorMetricResponseBody) SetRequestId(v string) *DescribeAppMonitorMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppMonitorMetricResponseBody) SetResult(v []*DescribeAppMonitorMetricResponseBodyResult) *DescribeAppMonitorMetricResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppMonitorMetricResponseBody) SetSuccess(v bool) *DescribeAppMonitorMetricResponseBody {
	s.Success = &v
	return s
}

type DescribeAppMonitorMetricResponseBodyResult struct {
	Categories []*string  `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Data       []*float32 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Name       *string    `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAppMonitorMetricResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppMonitorMetricResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppMonitorMetricResponseBodyResult) SetCategories(v []*string) *DescribeAppMonitorMetricResponseBodyResult {
	s.Categories = v
	return s
}

func (s *DescribeAppMonitorMetricResponseBodyResult) SetData(v []*float32) *DescribeAppMonitorMetricResponseBodyResult {
	s.Data = v
	return s
}

func (s *DescribeAppMonitorMetricResponseBodyResult) SetName(v string) *DescribeAppMonitorMetricResponseBodyResult {
	s.Name = &v
	return s
}

type DescribeAppMonitorMetricResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppMonitorMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppMonitorMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppMonitorMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppMonitorMetricResponse) SetHeaders(v map[string]*string) *DescribeAppMonitorMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppMonitorMetricResponse) SetStatusCode(v int32) *DescribeAppMonitorMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppMonitorMetricResponse) SetBody(v *DescribeAppMonitorMetricResponseBody) *DescribeAppMonitorMetricResponse {
	s.Body = v
	return s
}

type DescribeAppResourceAllocRequest struct {
	AppEnvId *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
}

func (s DescribeAppResourceAllocRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResourceAllocRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppResourceAllocRequest) SetAppEnvId(v int64) *DescribeAppResourceAllocRequest {
	s.AppEnvId = &v
	return s
}

type DescribeAppResourceAllocResponseBody struct {
	Code      *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                     `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAppResourceAllocResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAppResourceAllocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResourceAllocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppResourceAllocResponseBody) SetCode(v int32) *DescribeAppResourceAllocResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBody) SetErrMsg(v string) *DescribeAppResourceAllocResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBody) SetRequestId(v string) *DescribeAppResourceAllocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBody) SetResult(v *DescribeAppResourceAllocResponseBodyResult) *DescribeAppResourceAllocResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppResourceAllocResponseBody) SetSuccess(v bool) *DescribeAppResourceAllocResponseBody {
	s.Success = &v
	return s
}

type DescribeAppResourceAllocResponseBodyResult struct {
	AppEnvId    *int64  `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ResourceDef *string `json:"ResourceDef,omitempty" xml:"ResourceDef,omitempty"`
}

func (s DescribeAppResourceAllocResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResourceAllocResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppResourceAllocResponseBodyResult) SetAppEnvId(v int64) *DescribeAppResourceAllocResponseBodyResult {
	s.AppEnvId = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBodyResult) SetAppId(v int64) *DescribeAppResourceAllocResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBodyResult) SetClusterId(v string) *DescribeAppResourceAllocResponseBodyResult {
	s.ClusterId = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBodyResult) SetId(v int64) *DescribeAppResourceAllocResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppResourceAllocResponseBodyResult) SetResourceDef(v string) *DescribeAppResourceAllocResponseBodyResult {
	s.ResourceDef = &v
	return s
}

type DescribeAppResourceAllocResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppResourceAllocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppResourceAllocResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResourceAllocResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppResourceAllocResponse) SetHeaders(v map[string]*string) *DescribeAppResourceAllocResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppResourceAllocResponse) SetStatusCode(v int32) *DescribeAppResourceAllocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppResourceAllocResponse) SetBody(v *DescribeAppResourceAllocResponseBody) *DescribeAppResourceAllocResponse {
	s.Body = v
	return s
}

type DescribeClusterDetailRequest struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
}

func (s DescribeClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailRequest) SetClusterInstanceId(v string) *DescribeClusterDetailRequest {
	s.ClusterInstanceId = &v
	return s
}

type DescribeClusterDetailResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                  `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeClusterDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeClusterDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBody) SetCode(v int32) *DescribeClusterDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetErrMsg(v string) *DescribeClusterDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetRequestId(v string) *DescribeClusterDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetResult(v *DescribeClusterDetailResponseBodyResult) *DescribeClusterDetailResponseBody {
	s.Result = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSuccess(v bool) *DescribeClusterDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeClusterDetailResponseBodyResult struct {
	BasicInfo        *DescribeClusterDetailResponseBodyResultBasicInfo          `json:"BasicInfo,omitempty" xml:"BasicInfo,omitempty" type:"Struct"`
	InstanceInfo     *DescribeClusterDetailResponseBodyResultInstanceInfo       `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	NetInfo          *DescribeClusterDetailResponseBodyResultNetInfo            `json:"NetInfo,omitempty" xml:"NetInfo,omitempty" type:"Struct"`
	NodeWorkLoadList []*DescribeClusterDetailResponseBodyResultNodeWorkLoadList `json:"NodeWorkLoadList,omitempty" xml:"NodeWorkLoadList,omitempty" type:"Repeated"`
	WorkLoad         *DescribeClusterDetailResponseBodyResultWorkLoad           `json:"WorkLoad,omitempty" xml:"WorkLoad,omitempty" type:"Struct"`
}

func (s DescribeClusterDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResult) SetBasicInfo(v *DescribeClusterDetailResponseBodyResultBasicInfo) *DescribeClusterDetailResponseBodyResult {
	s.BasicInfo = v
	return s
}

func (s *DescribeClusterDetailResponseBodyResult) SetInstanceInfo(v *DescribeClusterDetailResponseBodyResultInstanceInfo) *DescribeClusterDetailResponseBodyResult {
	s.InstanceInfo = v
	return s
}

func (s *DescribeClusterDetailResponseBodyResult) SetNetInfo(v *DescribeClusterDetailResponseBodyResultNetInfo) *DescribeClusterDetailResponseBodyResult {
	s.NetInfo = v
	return s
}

func (s *DescribeClusterDetailResponseBodyResult) SetNodeWorkLoadList(v []*DescribeClusterDetailResponseBodyResultNodeWorkLoadList) *DescribeClusterDetailResponseBodyResult {
	s.NodeWorkLoadList = v
	return s
}

func (s *DescribeClusterDetailResponseBodyResult) SetWorkLoad(v *DescribeClusterDetailResponseBodyResultWorkLoad) *DescribeClusterDetailResponseBodyResult {
	s.WorkLoad = v
	return s
}

type DescribeClusterDetailResponseBodyResultBasicInfo struct {
	BusinessCode            *string   `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	ClusterId               *int64    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterInstanceId       *string   `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	ClusterName             *string   `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	EcsIds                  []*string `json:"EcsIds,omitempty" xml:"EcsIds,omitempty" type:"Repeated"`
	EnvType                 *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	HasInstallArmsPilot     *bool     `json:"HasInstallArmsPilot,omitempty" xml:"HasInstallArmsPilot,omitempty"`
	HasInstallLogController *bool     `json:"HasInstallLogController,omitempty" xml:"HasInstallLogController,omitempty"`
	InstallArmsInProcess    *bool     `json:"InstallArmsInProcess,omitempty" xml:"InstallArmsInProcess,omitempty"`
	InstallLogInProcess     *bool     `json:"InstallLogInProcess,omitempty" xml:"InstallLogInProcess,omitempty"`
	MainUserId              *string   `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName              *string   `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	UserNick                *string   `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	VpcId                   *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Vswitchs                []*string `json:"Vswitchs,omitempty" xml:"Vswitchs,omitempty" type:"Repeated"`
}

func (s DescribeClusterDetailResponseBodyResultBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultBasicInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetBusinessCode(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.BusinessCode = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetClusterId(v int64) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetClusterInstanceId(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.ClusterInstanceId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetClusterName(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetEcsIds(v []*string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.EcsIds = v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetEnvType(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.EnvType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetHasInstallArmsPilot(v bool) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.HasInstallArmsPilot = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetHasInstallLogController(v bool) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.HasInstallLogController = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetInstallArmsInProcess(v bool) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.InstallArmsInProcess = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetInstallLogInProcess(v bool) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.InstallLogInProcess = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetMainUserId(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.MainUserId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetRegionId(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetRegionName(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.RegionName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetUserNick(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.UserNick = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetVpcId(v string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultBasicInfo) SetVswitchs(v []*string) *DescribeClusterDetailResponseBodyResultBasicInfo {
	s.Vswitchs = v
	return s
}

type DescribeClusterDetailResponseBodyResultInstanceInfo struct {
	AllocatePodCount     *int32                                                                     `json:"AllocatePodCount,omitempty" xml:"AllocatePodCount,omitempty"`
	AllocatedPodInfoList []*DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList `json:"AllocatedPodInfoList,omitempty" xml:"AllocatedPodInfoList,omitempty" type:"Repeated"`
	AppCount             *int32                                                                     `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	AvailablePodInfoList []*DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList `json:"AvailablePodInfoList,omitempty" xml:"AvailablePodInfoList,omitempty" type:"Repeated"`
}

func (s DescribeClusterDetailResponseBodyResultInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfo) SetAllocatePodCount(v int32) *DescribeClusterDetailResponseBodyResultInstanceInfo {
	s.AllocatePodCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfo) SetAllocatedPodInfoList(v []*DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) *DescribeClusterDetailResponseBodyResultInstanceInfo {
	s.AllocatedPodInfoList = v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfo) SetAppCount(v int32) *DescribeClusterDetailResponseBodyResultInstanceInfo {
	s.AppCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfo) SetAvailablePodInfoList(v []*DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList) *DescribeClusterDetailResponseBodyResultInstanceInfo {
	s.AvailablePodInfoList = v
	return s
}

type DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList struct {
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CupRequest *string `json:"CupRequest,omitempty" xml:"CupRequest,omitempty"`
	EnvId      *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvName    *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	MemRequest *string `json:"MemRequest,omitempty" xml:"MemRequest,omitempty"`
	PodCount   *int32  `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
}

func (s DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetAppId(v int64) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.AppId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetAppName(v string) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.AppName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetCupRequest(v string) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.CupRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetEnvId(v int64) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.EnvId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetEnvName(v string) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.EnvName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetMemRequest(v string) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.MemRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList) SetPodCount(v int32) *DescribeClusterDetailResponseBodyResultInstanceInfoAllocatedPodInfoList {
	s.PodCount = &v
	return s
}

type DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList struct {
	AvailablePodCount *int32  `json:"AvailablePodCount,omitempty" xml:"AvailablePodCount,omitempty"`
	CupRequest        *string `json:"CupRequest,omitempty" xml:"CupRequest,omitempty"`
	MemRequest        *string `json:"MemRequest,omitempty" xml:"MemRequest,omitempty"`
}

func (s DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList) SetAvailablePodCount(v int32) *DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList {
	s.AvailablePodCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList) SetCupRequest(v string) *DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList {
	s.CupRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList) SetMemRequest(v string) *DescribeClusterDetailResponseBodyResultInstanceInfoAvailablePodInfoList {
	s.MemRequest = &v
	return s
}

type DescribeClusterDetailResponseBodyResultNetInfo struct {
	NetPlugType *string `json:"NetPlugType,omitempty" xml:"NetPlugType,omitempty"`
	ProdCIDR    *string `json:"ProdCIDR,omitempty" xml:"ProdCIDR,omitempty"`
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" xml:"ServiceCIDR,omitempty"`
}

func (s DescribeClusterDetailResponseBodyResultNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultNetInfo) SetNetPlugType(v string) *DescribeClusterDetailResponseBodyResultNetInfo {
	s.NetPlugType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNetInfo) SetProdCIDR(v string) *DescribeClusterDetailResponseBodyResultNetInfo {
	s.ProdCIDR = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNetInfo) SetServiceCIDR(v string) *DescribeClusterDetailResponseBodyResultNetInfo {
	s.ServiceCIDR = &v
	return s
}

type DescribeClusterDetailResponseBodyResultNodeWorkLoadList struct {
	AppPodCount   *int32  `json:"AppPodCount,omitempty" xml:"AppPodCount,omitempty"`
	CpuRequest    *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	CpuTotal      *string `json:"CpuTotal,omitempty" xml:"CpuTotal,omitempty"`
	CpuUse        *string `json:"CpuUse,omitempty" xml:"CpuUse,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemRequest    *string `json:"MemRequest,omitempty" xml:"MemRequest,omitempty"`
	MemTotal      *string `json:"MemTotal,omitempty" xml:"MemTotal,omitempty"`
	MemUse        *string `json:"MemUse,omitempty" xml:"MemUse,omitempty"`
	NodeName      *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	PodCount      *int32  `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	Ready         *bool   `json:"Ready,omitempty" xml:"Ready,omitempty"`
	Unschedulable *bool   `json:"Unschedulable,omitempty" xml:"Unschedulable,omitempty"`
}

func (s DescribeClusterDetailResponseBodyResultNodeWorkLoadList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultNodeWorkLoadList) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetAppPodCount(v int32) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.AppPodCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetCpuRequest(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.CpuRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetCpuTotal(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.CpuTotal = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetCpuUse(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.CpuUse = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetInstanceId(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetMemRequest(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.MemRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetMemTotal(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.MemTotal = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetMemUse(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.MemUse = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetNodeName(v string) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.NodeName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetPodCount(v int32) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.PodCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetReady(v bool) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.Ready = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultNodeWorkLoadList) SetUnschedulable(v bool) *DescribeClusterDetailResponseBodyResultNodeWorkLoadList {
	s.Unschedulable = &v
	return s
}

type DescribeClusterDetailResponseBodyResultWorkLoad struct {
	AllNodeCount        *int32  `json:"AllNodeCount,omitempty" xml:"AllNodeCount,omitempty"`
	AllocateAllPodCount *int32  `json:"AllocateAllPodCount,omitempty" xml:"AllocateAllPodCount,omitempty"`
	AllocateAppPodCount *int32  `json:"AllocateAppPodCount,omitempty" xml:"AllocateAppPodCount,omitempty"`
	CpuCapacityTotal    *string `json:"CpuCapacityTotal,omitempty" xml:"CpuCapacityTotal,omitempty"`
	CpuLevel            *string `json:"CpuLevel,omitempty" xml:"CpuLevel,omitempty"`
	CpuRequest          *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	CpuRequestPercent   *string `json:"CpuRequestPercent,omitempty" xml:"CpuRequestPercent,omitempty"`
	CpuTotal            *string `json:"CpuTotal,omitempty" xml:"CpuTotal,omitempty"`
	CpuUse              *string `json:"CpuUse,omitempty" xml:"CpuUse,omitempty"`
	CpuUsePercent       *string `json:"CpuUsePercent,omitempty" xml:"CpuUsePercent,omitempty"`
	MemCapacityTotal    *string `json:"MemCapacityTotal,omitempty" xml:"MemCapacityTotal,omitempty"`
	MemLevel            *string `json:"MemLevel,omitempty" xml:"MemLevel,omitempty"`
	MemRequest          *string `json:"MemRequest,omitempty" xml:"MemRequest,omitempty"`
	MemRequestPercent   *string `json:"MemRequestPercent,omitempty" xml:"MemRequestPercent,omitempty"`
	MemTotal            *string `json:"MemTotal,omitempty" xml:"MemTotal,omitempty"`
	MemUse              *string `json:"MemUse,omitempty" xml:"MemUse,omitempty"`
	MemUsePercent       *string `json:"MemUsePercent,omitempty" xml:"MemUsePercent,omitempty"`
}

func (s DescribeClusterDetailResponseBodyResultWorkLoad) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyResultWorkLoad) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetAllNodeCount(v int32) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.AllNodeCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetAllocateAllPodCount(v int32) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.AllocateAllPodCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetAllocateAppPodCount(v int32) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.AllocateAppPodCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuCapacityTotal(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuCapacityTotal = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuLevel(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuLevel = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuRequest(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuRequestPercent(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuRequestPercent = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuTotal(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuTotal = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuUse(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuUse = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetCpuUsePercent(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.CpuUsePercent = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemCapacityTotal(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemCapacityTotal = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemLevel(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemLevel = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemRequest(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemRequest = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemRequestPercent(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemRequestPercent = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemTotal(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemTotal = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemUse(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemUse = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyResultWorkLoad) SetMemUsePercent(v string) *DescribeClusterDetailResponseBodyResultWorkLoad {
	s.MemUsePercent = &v
	return s
}

type DescribeClusterDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterDetailResponse) SetStatusCode(v int32) *DescribeClusterDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterDetailResponse) SetBody(v *DescribeClusterDetailResponseBody) *DescribeClusterDetailResponse {
	s.Body = v
	return s
}

type DescribeDatabasesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) SetInstanceId(v string) *DescribeDatabasesRequest {
	s.InstanceId = &v
	return s
}

type DescribeDatabasesResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDatabasesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBody) SetCode(v int32) *DescribeDatabasesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetErrMsg(v string) *DescribeDatabasesResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetResult(v *DescribeDatabasesResponseBodyResult) *DescribeDatabasesResponseBody {
	s.Result = v
	return s
}

type DescribeDatabasesResponseBodyResult struct {
	Databases []*DescribeDatabasesResponseBodyResultDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyResult) SetDatabases(v []*DescribeDatabasesResponseBodyResultDatabases) *DescribeDatabasesResponseBodyResult {
	s.Databases = v
	return s
}

type DescribeDatabasesResponseBodyResultDatabases struct {
	Accounts         []*DescribeDatabasesResponseBodyResultDatabasesAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	CharacterSetName *string                                                 `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	DBDescription    *string                                                 `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	DBInstanceId     *string                                                 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBName           *string                                                 `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBStatus         *string                                                 `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	Engine           *string                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDatabasesResponseBodyResultDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyResultDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetAccounts(v []*DescribeDatabasesResponseBodyResultDatabasesAccounts) *DescribeDatabasesResponseBodyResultDatabases {
	s.Accounts = v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetCharacterSetName(v string) *DescribeDatabasesResponseBodyResultDatabases {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetDBDescription(v string) *DescribeDatabasesResponseBodyResultDatabases {
	s.DBDescription = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetDBInstanceId(v string) *DescribeDatabasesResponseBodyResultDatabases {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetDBName(v string) *DescribeDatabasesResponseBodyResultDatabases {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetDBStatus(v string) *DescribeDatabasesResponseBodyResultDatabases {
	s.DBStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabases) SetEngine(v string) *DescribeDatabasesResponseBodyResultDatabases {
	s.Engine = &v
	return s
}

type DescribeDatabasesResponseBodyResultDatabasesAccounts struct {
	Account                *string `json:"Account,omitempty" xml:"Account,omitempty"`
	AccountPrivilege       *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountPrivilegeDetail *string `json:"AccountPrivilegeDetail,omitempty" xml:"AccountPrivilegeDetail,omitempty"`
}

func (s DescribeDatabasesResponseBodyResultDatabasesAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyResultDatabasesAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyResultDatabasesAccounts) SetAccount(v string) *DescribeDatabasesResponseBodyResultDatabasesAccounts {
	s.Account = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabasesAccounts) SetAccountPrivilege(v string) *DescribeDatabasesResponseBodyResultDatabasesAccounts {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDatabasesResponseBodyResultDatabasesAccounts) SetAccountPrivilegeDetail(v string) *DescribeDatabasesResponseBodyResultDatabasesAccounts {
	s.AccountPrivilegeDetail = &v
	return s
}

type DescribeDatabasesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponse) SetHeaders(v map[string]*string) *DescribeDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabasesResponse) SetStatusCode(v int32) *DescribeDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabasesResponse) SetBody(v *DescribeDatabasesResponseBody) *DescribeDatabasesResponse {
	s.Body = v
	return s
}

type DescribeDeployOrderDetailRequest struct {
	DeployOrderId *int64 `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s DescribeDeployOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeployOrderDetailRequest) SetDeployOrderId(v int64) *DescribeDeployOrderDetailRequest {
	s.DeployOrderId = &v
	return s
}

type DescribeDeployOrderDetailResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                      `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDeployOrderDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDeployOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployOrderDetailResponseBody) SetCode(v int32) *DescribeDeployOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBody) SetErrMsg(v string) *DescribeDeployOrderDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBody) SetRequestId(v string) *DescribeDeployOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBody) SetResult(v *DescribeDeployOrderDetailResponseBodyResult) *DescribeDeployOrderDetailResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDeployOrderDetailResponseBody) SetSuccess(v bool) *DescribeDeployOrderDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeDeployOrderDetailResponseBodyResult struct {
	AppInstanceType     *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	CurrentPartitionNum *int32  `json:"CurrentPartitionNum,omitempty" xml:"CurrentPartitionNum,omitempty"`
	DeployOrderId       *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	DeployPauseType     *string `json:"DeployPauseType,omitempty" xml:"DeployPauseType,omitempty"`
	DeployPauseTypeName *string `json:"DeployPauseTypeName,omitempty" xml:"DeployPauseTypeName,omitempty"`
	DeployType          *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DeployTypeName      *string `json:"DeployTypeName,omitempty" xml:"DeployTypeName,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ElapsedTime         *int32  `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvId               *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType             *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	FailureRate         *string `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	FinishAppInstanceCt *int32  `json:"FinishAppInstanceCt,omitempty" xml:"FinishAppInstanceCt,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionType       *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	PartitionTypeName   *string `json:"PartitionTypeName,omitempty" xml:"PartitionTypeName,omitempty"`
	Result              *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultName          *string `json:"ResultName,omitempty" xml:"ResultName,omitempty"`
	SchemaId            *int64  `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName          *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TotalAppInstanceCt  *int32  `json:"TotalAppInstanceCt,omitempty" xml:"TotalAppInstanceCt,omitempty"`
	TotalPartitions     *int32  `json:"TotalPartitions,omitempty" xml:"TotalPartitions,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick            *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s DescribeDeployOrderDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployOrderDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetAppInstanceType(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.AppInstanceType = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetCurrentPartitionNum(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.CurrentPartitionNum = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetDeployOrderId(v int64) *DescribeDeployOrderDetailResponseBodyResult {
	s.DeployOrderId = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetDeployPauseType(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.DeployPauseType = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetDeployPauseTypeName(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.DeployPauseTypeName = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetDeployType(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.DeployType = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetDeployTypeName(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.DeployTypeName = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetDescription(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetElapsedTime(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetEndTime(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetEnvId(v int64) *DescribeDeployOrderDetailResponseBodyResult {
	s.EnvId = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetEnvType(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.EnvType = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetFailureRate(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.FailureRate = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetFinishAppInstanceCt(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.FinishAppInstanceCt = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetName(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetPartitionType(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.PartitionType = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetPartitionTypeName(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.PartitionTypeName = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetResult(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetResultName(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.ResultName = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetSchemaId(v int64) *DescribeDeployOrderDetailResponseBodyResult {
	s.SchemaId = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetStartTime(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetStatus(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetStatusName(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.StatusName = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetTotalAppInstanceCt(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.TotalAppInstanceCt = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetTotalPartitions(v int32) *DescribeDeployOrderDetailResponseBodyResult {
	s.TotalPartitions = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetUserId(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *DescribeDeployOrderDetailResponseBodyResult) SetUserNick(v string) *DescribeDeployOrderDetailResponseBodyResult {
	s.UserNick = &v
	return s
}

type DescribeDeployOrderDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeployOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeployOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployOrderDetailResponse) SetHeaders(v map[string]*string) *DescribeDeployOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployOrderDetailResponse) SetStatusCode(v int32) *DescribeDeployOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeployOrderDetailResponse) SetBody(v *DescribeDeployOrderDetailResponseBody) *DescribeDeployOrderDetailResponse {
	s.Body = v
	return s
}

type DescribeEciConfigRequest struct {
	AppEnvId *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
}

func (s DescribeEciConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeEciConfigRequest) SetAppEnvId(v int64) *DescribeEciConfigRequest {
	s.AppEnvId = &v
	return s
}

type DescribeEciConfigResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeEciConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeEciConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEciConfigResponseBody) SetCode(v int32) *DescribeEciConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEciConfigResponseBody) SetErrMsg(v string) *DescribeEciConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeEciConfigResponseBody) SetRequestId(v string) *DescribeEciConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEciConfigResponseBody) SetResult(v *DescribeEciConfigResponseBodyResult) *DescribeEciConfigResponseBody {
	s.Result = v
	return s
}

type DescribeEciConfigResponseBodyResult struct {
	AppEnvId                *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	EipBandwidth            *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EnableEciSchedulePolicy *bool  `json:"EnableEciSchedulePolicy,omitempty" xml:"EnableEciSchedulePolicy,omitempty"`
	MirrorCache             *bool  `json:"MirrorCache,omitempty" xml:"MirrorCache,omitempty"`
	NormalInstanceLimit     *int32 `json:"NormalInstanceLimit,omitempty" xml:"NormalInstanceLimit,omitempty"`
	ScheduleVirtualNode     *bool  `json:"ScheduleVirtualNode,omitempty" xml:"ScheduleVirtualNode,omitempty"`
}

func (s DescribeEciConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeEciConfigResponseBodyResult) SetAppEnvId(v int64) *DescribeEciConfigResponseBodyResult {
	s.AppEnvId = &v
	return s
}

func (s *DescribeEciConfigResponseBodyResult) SetEipBandwidth(v int32) *DescribeEciConfigResponseBodyResult {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEciConfigResponseBodyResult) SetEnableEciSchedulePolicy(v bool) *DescribeEciConfigResponseBodyResult {
	s.EnableEciSchedulePolicy = &v
	return s
}

func (s *DescribeEciConfigResponseBodyResult) SetMirrorCache(v bool) *DescribeEciConfigResponseBodyResult {
	s.MirrorCache = &v
	return s
}

func (s *DescribeEciConfigResponseBodyResult) SetNormalInstanceLimit(v int32) *DescribeEciConfigResponseBodyResult {
	s.NormalInstanceLimit = &v
	return s
}

func (s *DescribeEciConfigResponseBodyResult) SetScheduleVirtualNode(v bool) *DescribeEciConfigResponseBodyResult {
	s.ScheduleVirtualNode = &v
	return s
}

type DescribeEciConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEciConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEciConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeEciConfigResponse) SetHeaders(v map[string]*string) *DescribeEciConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeEciConfigResponse) SetStatusCode(v int32) *DescribeEciConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEciConfigResponse) SetBody(v *DescribeEciConfigResponseBody) *DescribeEciConfigResponse {
	s.Body = v
	return s
}

type DescribeEventMonitorListRequest struct {
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvId      *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventType  *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	PageNum    *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PodName    *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventMonitorListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventMonitorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventMonitorListRequest) SetAppId(v int64) *DescribeEventMonitorListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetEndTime(v int64) *DescribeEventMonitorListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetEnvId(v int64) *DescribeEventMonitorListRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetEventLevel(v string) *DescribeEventMonitorListRequest {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetEventType(v string) *DescribeEventMonitorListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetPageNum(v int32) *DescribeEventMonitorListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetPageSize(v int32) *DescribeEventMonitorListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetPodName(v string) *DescribeEventMonitorListRequest {
	s.PodName = &v
	return s
}

func (s *DescribeEventMonitorListRequest) SetStartTime(v int64) *DescribeEventMonitorListRequest {
	s.StartTime = &v
	return s
}

type DescribeEventMonitorListResponseBody struct {
	Code       *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeEventMonitorListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventMonitorListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventMonitorListResponseBody) SetCode(v int32) *DescribeEventMonitorListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventMonitorListResponseBody) SetData(v []*DescribeEventMonitorListResponseBodyData) *DescribeEventMonitorListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventMonitorListResponseBody) SetErrorMsg(v string) *DescribeEventMonitorListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeEventMonitorListResponseBody) SetPageNumber(v int32) *DescribeEventMonitorListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventMonitorListResponseBody) SetPageSize(v int32) *DescribeEventMonitorListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventMonitorListResponseBody) SetRequestId(v string) *DescribeEventMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventMonitorListResponseBody) SetTotalCount(v int64) *DescribeEventMonitorListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEventMonitorListResponseBodyData struct {
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	HostName  *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Kind      *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	PodName   *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeEventMonitorListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventMonitorListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventMonitorListResponseBodyData) SetCount(v int32) *DescribeEventMonitorListResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetEventTime(v string) *DescribeEventMonitorListResponseBodyData {
	s.EventTime = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetHostName(v string) *DescribeEventMonitorListResponseBodyData {
	s.HostName = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetKind(v string) *DescribeEventMonitorListResponseBodyData {
	s.Kind = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetLevel(v string) *DescribeEventMonitorListResponseBodyData {
	s.Level = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetMessage(v string) *DescribeEventMonitorListResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetNameSpace(v string) *DescribeEventMonitorListResponseBodyData {
	s.NameSpace = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetPodName(v string) *DescribeEventMonitorListResponseBodyData {
	s.PodName = &v
	return s
}

func (s *DescribeEventMonitorListResponseBodyData) SetReason(v string) *DescribeEventMonitorListResponseBodyData {
	s.Reason = &v
	return s
}

type DescribeEventMonitorListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEventMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventMonitorListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventMonitorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventMonitorListResponse) SetHeaders(v map[string]*string) *DescribeEventMonitorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventMonitorListResponse) SetStatusCode(v int32) *DescribeEventMonitorListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventMonitorListResponse) SetBody(v *DescribeEventMonitorListResponseBody) *DescribeEventMonitorListResponse {
	s.Body = v
	return s
}

type DescribeJobLogRequest struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId   *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s DescribeJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobLogRequest) SetAppId(v int64) *DescribeJobLogRequest {
	s.AppId = &v
	return s
}

func (s *DescribeJobLogRequest) SetEnvId(v int64) *DescribeJobLogRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeJobLogRequest) SetPodName(v string) *DescribeJobLogRequest {
	s.PodName = &v
	return s
}

type DescribeJobLogResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                           `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeJobLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobLogResponseBody) SetCode(v int32) *DescribeJobLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobLogResponseBody) SetErrMsg(v string) *DescribeJobLogResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeJobLogResponseBody) SetRequestId(v string) *DescribeJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobLogResponseBody) SetResult(v *DescribeJobLogResponseBodyResult) *DescribeJobLogResponseBody {
	s.Result = v
	return s
}

type DescribeJobLogResponseBodyResult struct {
	AppId   *int64                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId   *int64                                    `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Events  []*DescribeJobLogResponseBodyResultEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	JobLog  *string                                   `json:"JobLog,omitempty" xml:"JobLog,omitempty"`
	PodName *string                                   `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s DescribeJobLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeJobLogResponseBodyResult) SetAppId(v int64) *DescribeJobLogResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeJobLogResponseBodyResult) SetEnvId(v int64) *DescribeJobLogResponseBodyResult {
	s.EnvId = &v
	return s
}

func (s *DescribeJobLogResponseBodyResult) SetEvents(v []*DescribeJobLogResponseBodyResultEvents) *DescribeJobLogResponseBodyResult {
	s.Events = v
	return s
}

func (s *DescribeJobLogResponseBodyResult) SetJobLog(v string) *DescribeJobLogResponseBodyResult {
	s.JobLog = &v
	return s
}

func (s *DescribeJobLogResponseBodyResult) SetPodName(v string) *DescribeJobLogResponseBodyResult {
	s.PodName = &v
	return s
}

type DescribeJobLogResponseBodyResultEvents struct {
	Action         *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	EventTime      *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Mesage         *string `json:"Mesage,omitempty" xml:"Mesage,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeJobLogResponseBodyResultEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobLogResponseBodyResultEvents) GoString() string {
	return s.String()
}

func (s *DescribeJobLogResponseBodyResultEvents) SetAction(v string) *DescribeJobLogResponseBodyResultEvents {
	s.Action = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetCount(v int32) *DescribeJobLogResponseBodyResultEvents {
	s.Count = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetEventTime(v string) *DescribeJobLogResponseBodyResultEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetFirstTimestamp(v string) *DescribeJobLogResponseBodyResultEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetLastTimestamp(v string) *DescribeJobLogResponseBodyResultEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetMesage(v string) *DescribeJobLogResponseBodyResultEvents {
	s.Mesage = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetReason(v string) *DescribeJobLogResponseBodyResultEvents {
	s.Reason = &v
	return s
}

func (s *DescribeJobLogResponseBodyResultEvents) SetType(v string) *DescribeJobLogResponseBodyResultEvents {
	s.Type = &v
	return s
}

type DescribeJobLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobLogResponse) SetHeaders(v map[string]*string) *DescribeJobLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobLogResponse) SetStatusCode(v int32) *DescribeJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobLogResponse) SetBody(v *DescribeJobLogResponseBody) *DescribeJobLogResponse {
	s.Body = v
	return s
}

type DescribePodContainerLogListRequest struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId   *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Line    *int32  `json:"Line,omitempty" xml:"Line,omitempty"`
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s DescribePodContainerLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePodContainerLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribePodContainerLogListRequest) SetAppId(v int64) *DescribePodContainerLogListRequest {
	s.AppId = &v
	return s
}

func (s *DescribePodContainerLogListRequest) SetEnvId(v int64) *DescribePodContainerLogListRequest {
	s.EnvId = &v
	return s
}

func (s *DescribePodContainerLogListRequest) SetLine(v int32) *DescribePodContainerLogListRequest {
	s.Line = &v
	return s
}

func (s *DescribePodContainerLogListRequest) SetPodName(v string) *DescribePodContainerLogListRequest {
	s.PodName = &v
	return s
}

type DescribePodContainerLogListResponseBody struct {
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePodContainerLogListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePodContainerLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePodContainerLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePodContainerLogListResponseBody) SetCode(v int32) *DescribePodContainerLogListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePodContainerLogListResponseBody) SetErrMsg(v string) *DescribePodContainerLogListResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribePodContainerLogListResponseBody) SetRequestId(v string) *DescribePodContainerLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePodContainerLogListResponseBody) SetResult(v *DescribePodContainerLogListResponseBodyResult) *DescribePodContainerLogListResponseBody {
	s.Result = v
	return s
}

func (s *DescribePodContainerLogListResponseBody) SetSuccess(v bool) *DescribePodContainerLogListResponseBody {
	s.Success = &v
	return s
}

type DescribePodContainerLogListResponseBodyResult struct {
	ContainerLogList []*DescribePodContainerLogListResponseBodyResultContainerLogList `json:"ContainerLogList,omitempty" xml:"ContainerLogList,omitempty" type:"Repeated"`
}

func (s DescribePodContainerLogListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePodContainerLogListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePodContainerLogListResponseBodyResult) SetContainerLogList(v []*DescribePodContainerLogListResponseBodyResultContainerLogList) *DescribePodContainerLogListResponseBodyResult {
	s.ContainerLogList = v
	return s
}

type DescribePodContainerLogListResponseBodyResultContainerLogList struct {
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	PodName       *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s DescribePodContainerLogListResponseBodyResultContainerLogList) String() string {
	return tea.Prettify(s)
}

func (s DescribePodContainerLogListResponseBodyResultContainerLogList) GoString() string {
	return s.String()
}

func (s *DescribePodContainerLogListResponseBodyResultContainerLogList) SetContainerName(v string) *DescribePodContainerLogListResponseBodyResultContainerLogList {
	s.ContainerName = &v
	return s
}

func (s *DescribePodContainerLogListResponseBodyResultContainerLogList) SetContent(v string) *DescribePodContainerLogListResponseBodyResultContainerLogList {
	s.Content = &v
	return s
}

func (s *DescribePodContainerLogListResponseBodyResultContainerLogList) SetPodName(v string) *DescribePodContainerLogListResponseBodyResultContainerLogList {
	s.PodName = &v
	return s
}

type DescribePodContainerLogListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePodContainerLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePodContainerLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePodContainerLogListResponse) GoString() string {
	return s.String()
}

func (s *DescribePodContainerLogListResponse) SetHeaders(v map[string]*string) *DescribePodContainerLogListResponse {
	s.Headers = v
	return s
}

func (s *DescribePodContainerLogListResponse) SetStatusCode(v int32) *DescribePodContainerLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePodContainerLogListResponse) SetBody(v *DescribePodContainerLogListResponseBody) *DescribePodContainerLogListResponse {
	s.Body = v
	return s
}

type DescribePodEventsRequest struct {
	AppInstId     *string `json:"AppInstId,omitempty" xml:"AppInstId,omitempty"`
	DeployOrderId *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s DescribePodEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePodEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribePodEventsRequest) SetAppInstId(v string) *DescribePodEventsRequest {
	s.AppInstId = &v
	return s
}

func (s *DescribePodEventsRequest) SetDeployOrderId(v int64) *DescribePodEventsRequest {
	s.DeployOrderId = &v
	return s
}

type DescribePodEventsResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePodEventsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePodEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePodEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePodEventsResponseBody) SetCode(v int32) *DescribePodEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePodEventsResponseBody) SetErrMsg(v string) *DescribePodEventsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribePodEventsResponseBody) SetRequestId(v string) *DescribePodEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePodEventsResponseBody) SetResult(v *DescribePodEventsResponseBodyResult) *DescribePodEventsResponseBody {
	s.Result = v
	return s
}

func (s *DescribePodEventsResponseBody) SetSuccess(v bool) *DescribePodEventsResponseBody {
	s.Success = &v
	return s
}

type DescribePodEventsResponseBodyResult struct {
	DeployOrderName *string                                         `json:"DeployOrderName,omitempty" xml:"DeployOrderName,omitempty"`
	PodEvents       []*DescribePodEventsResponseBodyResultPodEvents `json:"PodEvents,omitempty" xml:"PodEvents,omitempty" type:"Repeated"`
}

func (s DescribePodEventsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePodEventsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePodEventsResponseBodyResult) SetDeployOrderName(v string) *DescribePodEventsResponseBodyResult {
	s.DeployOrderName = &v
	return s
}

func (s *DescribePodEventsResponseBodyResult) SetPodEvents(v []*DescribePodEventsResponseBodyResultPodEvents) *DescribePodEventsResponseBodyResult {
	s.PodEvents = v
	return s
}

type DescribePodEventsResponseBodyResultPodEvents struct {
	Action         *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	EventTime      *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePodEventsResponseBodyResultPodEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribePodEventsResponseBodyResultPodEvents) GoString() string {
	return s.String()
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetAction(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.Action = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetCount(v int32) *DescribePodEventsResponseBodyResultPodEvents {
	s.Count = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetEventTime(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.EventTime = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetFirstTimestamp(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetLastTimestamp(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetMessage(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.Message = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetReason(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.Reason = &v
	return s
}

func (s *DescribePodEventsResponseBodyResultPodEvents) SetType(v string) *DescribePodEventsResponseBodyResultPodEvents {
	s.Type = &v
	return s
}

type DescribePodEventsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePodEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePodEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePodEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribePodEventsResponse) SetHeaders(v map[string]*string) *DescribePodEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribePodEventsResponse) SetStatusCode(v int32) *DescribePodEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePodEventsResponse) SetBody(v *DescribePodEventsResponseBody) *DescribePodEventsResponse {
	s.Body = v
	return s
}

type DescribePodLogRequest struct {
	AppInstId     *string `json:"AppInstId,omitempty" xml:"AppInstId,omitempty"`
	DeployOrderId *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s DescribePodLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePodLogRequest) GoString() string {
	return s.String()
}

func (s *DescribePodLogRequest) SetAppInstId(v string) *DescribePodLogRequest {
	s.AppInstId = &v
	return s
}

func (s *DescribePodLogRequest) SetDeployOrderId(v int64) *DescribePodLogRequest {
	s.DeployOrderId = &v
	return s
}

type DescribePodLogResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                           `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePodLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePodLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePodLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePodLogResponseBody) SetCode(v int32) *DescribePodLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePodLogResponseBody) SetErrMsg(v string) *DescribePodLogResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribePodLogResponseBody) SetRequestId(v string) *DescribePodLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePodLogResponseBody) SetResult(v *DescribePodLogResponseBodyResult) *DescribePodLogResponseBody {
	s.Result = v
	return s
}

func (s *DescribePodLogResponseBody) SetSuccess(v bool) *DescribePodLogResponseBody {
	s.Success = &v
	return s
}

type DescribePodLogResponseBodyResult struct {
	DeployOrderName *string                                           `json:"DeployOrderName,omitempty" xml:"DeployOrderName,omitempty"`
	DeployStepList  []*DescribePodLogResponseBodyResultDeployStepList `json:"DeployStepList,omitempty" xml:"DeployStepList,omitempty" type:"Repeated"`
	EnvTypeName     *string                                           `json:"EnvTypeName,omitempty" xml:"EnvTypeName,omitempty"`
}

func (s DescribePodLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePodLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePodLogResponseBodyResult) SetDeployOrderName(v string) *DescribePodLogResponseBodyResult {
	s.DeployOrderName = &v
	return s
}

func (s *DescribePodLogResponseBodyResult) SetDeployStepList(v []*DescribePodLogResponseBodyResultDeployStepList) *DescribePodLogResponseBodyResult {
	s.DeployStepList = v
	return s
}

func (s *DescribePodLogResponseBodyResult) SetEnvTypeName(v string) *DescribePodLogResponseBodyResult {
	s.EnvTypeName = &v
	return s
}

type DescribePodLogResponseBodyResultDeployStepList struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StepCode *string `json:"StepCode,omitempty" xml:"StepCode,omitempty"`
	StepLog  *string `json:"StepLog,omitempty" xml:"StepLog,omitempty"`
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s DescribePodLogResponseBodyResultDeployStepList) String() string {
	return tea.Prettify(s)
}

func (s DescribePodLogResponseBodyResultDeployStepList) GoString() string {
	return s.String()
}

func (s *DescribePodLogResponseBodyResultDeployStepList) SetStatus(v string) *DescribePodLogResponseBodyResultDeployStepList {
	s.Status = &v
	return s
}

func (s *DescribePodLogResponseBodyResultDeployStepList) SetStepCode(v string) *DescribePodLogResponseBodyResultDeployStepList {
	s.StepCode = &v
	return s
}

func (s *DescribePodLogResponseBodyResultDeployStepList) SetStepLog(v string) *DescribePodLogResponseBodyResultDeployStepList {
	s.StepLog = &v
	return s
}

func (s *DescribePodLogResponseBodyResultDeployStepList) SetStepName(v string) *DescribePodLogResponseBodyResultDeployStepList {
	s.StepName = &v
	return s
}

type DescribePodLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePodLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePodLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePodLogResponse) GoString() string {
	return s.String()
}

func (s *DescribePodLogResponse) SetHeaders(v map[string]*string) *DescribePodLogResponse {
	s.Headers = v
	return s
}

func (s *DescribePodLogResponse) SetStatusCode(v int32) *DescribePodLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePodLogResponse) SetBody(v *DescribePodLogResponseBody) *DescribePodLogResponse {
	s.Body = v
	return s
}

type DescribeRdsAccountsRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s DescribeRdsAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsAccountsRequest) SetAccountName(v string) *DescribeRdsAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeRdsAccountsRequest) SetDbInstanceId(v string) *DescribeRdsAccountsRequest {
	s.DbInstanceId = &v
	return s
}

type DescribeRdsAccountsResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeRdsAccountsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeRdsAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsAccountsResponseBody) SetCode(v int32) *DescribeRdsAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRdsAccountsResponseBody) SetErrMsg(v string) *DescribeRdsAccountsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeRdsAccountsResponseBody) SetRequestId(v string) *DescribeRdsAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsAccountsResponseBody) SetResult(v *DescribeRdsAccountsResponseBodyResult) *DescribeRdsAccountsResponseBody {
	s.Result = v
	return s
}

type DescribeRdsAccountsResponseBodyResult struct {
	Accounts []*DescribeRdsAccountsResponseBodyResultAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s DescribeRdsAccountsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsAccountsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRdsAccountsResponseBodyResult) SetAccounts(v []*DescribeRdsAccountsResponseBodyResultAccounts) *DescribeRdsAccountsResponseBodyResult {
	s.Accounts = v
	return s
}

type DescribeRdsAccountsResponseBodyResultAccounts struct {
	AccountDescription *string                                                            `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string                                                            `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string                                                            `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountType        *string                                                            `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	DBInstanceId       *string                                                            `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DatabasePrivileges []*DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Repeated"`
	PrivExceeded       *string                                                            `json:"PrivExceeded,omitempty" xml:"PrivExceeded,omitempty"`
}

func (s DescribeRdsAccountsResponseBodyResultAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsAccountsResponseBodyResultAccounts) GoString() string {
	return s.String()
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetAccountDescription(v string) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.AccountDescription = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetAccountName(v string) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetAccountStatus(v string) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetAccountType(v string) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetDBInstanceId(v string) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetDatabasePrivileges(v []*DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccounts) SetPrivExceeded(v string) *DescribeRdsAccountsResponseBodyResultAccounts {
	s.PrivExceeded = &v
	return s
}

type DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges struct {
	AccountPrivilege       *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountPrivilegeDetail *string `json:"AccountPrivilegeDetail,omitempty" xml:"AccountPrivilegeDetail,omitempty"`
	DBName                 *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges) SetAccountPrivilege(v string) *DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges) SetAccountPrivilegeDetail(v string) *DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges {
	s.AccountPrivilegeDetail = &v
	return s
}

func (s *DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges) SetDBName(v string) *DescribeRdsAccountsResponseBodyResultAccountsDatabasePrivileges {
	s.DBName = &v
	return s
}

type DescribeRdsAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRdsAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRdsAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsAccountsResponse) SetHeaders(v map[string]*string) *DescribeRdsAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsAccountsResponse) SetStatusCode(v int32) *DescribeRdsAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsAccountsResponse) SetBody(v *DescribeRdsAccountsResponseBody) *DescribeRdsAccountsResponse {
	s.Body = v
	return s
}

type DescribeServiceDetailRequest struct {
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DescribeServiceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceDetailRequest) SetServiceId(v int64) *DescribeServiceDetailRequest {
	s.ServiceId = &v
	return s
}

type DescribeServiceDetailResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                  `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeServiceDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeServiceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceDetailResponseBody) SetCode(v int32) *DescribeServiceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceDetailResponseBody) SetErrMsg(v string) *DescribeServiceDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeServiceDetailResponseBody) SetRequestId(v string) *DescribeServiceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceDetailResponseBody) SetResult(v *DescribeServiceDetailResponseBodyResult) *DescribeServiceDetailResponseBody {
	s.Result = v
	return s
}

func (s *DescribeServiceDetailResponseBody) SetSuccess(v bool) *DescribeServiceDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeServiceDetailResponseBodyResult struct {
	AppId        *int64                                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterIP    *string                                                `json:"ClusterIP,omitempty" xml:"ClusterIP,omitempty"`
	EnvId        *int64                                                 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Headless     *bool                                                  `json:"Headless,omitempty" xml:"Headless,omitempty"`
	K8sServiceId *string                                                `json:"K8sServiceId,omitempty" xml:"K8sServiceId,omitempty"`
	Name         *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PortMappings []*DescribeServiceDetailResponseBodyResultPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	ServiceId    *int64                                                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceType  *string                                                `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s DescribeServiceDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeServiceDetailResponseBodyResult) SetAppId(v int64) *DescribeServiceDetailResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetClusterIP(v string) *DescribeServiceDetailResponseBodyResult {
	s.ClusterIP = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetEnvId(v int64) *DescribeServiceDetailResponseBodyResult {
	s.EnvId = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetHeadless(v bool) *DescribeServiceDetailResponseBodyResult {
	s.Headless = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetK8sServiceId(v string) *DescribeServiceDetailResponseBodyResult {
	s.K8sServiceId = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetName(v string) *DescribeServiceDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetPortMappings(v []*DescribeServiceDetailResponseBodyResultPortMappings) *DescribeServiceDetailResponseBodyResult {
	s.PortMappings = v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetServiceId(v int64) *DescribeServiceDetailResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResult) SetServiceType(v string) *DescribeServiceDetailResponseBodyResult {
	s.ServiceType = &v
	return s
}

type DescribeServiceDetailResponseBodyResultPortMappings struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePort   *int32  `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s DescribeServiceDetailResponseBodyResultPortMappings) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDetailResponseBodyResultPortMappings) GoString() string {
	return s.String()
}

func (s *DescribeServiceDetailResponseBodyResultPortMappings) SetName(v string) *DescribeServiceDetailResponseBodyResultPortMappings {
	s.Name = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResultPortMappings) SetNodePort(v int32) *DescribeServiceDetailResponseBodyResultPortMappings {
	s.NodePort = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResultPortMappings) SetPort(v int32) *DescribeServiceDetailResponseBodyResultPortMappings {
	s.Port = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResultPortMappings) SetProtocol(v string) *DescribeServiceDetailResponseBodyResultPortMappings {
	s.Protocol = &v
	return s
}

func (s *DescribeServiceDetailResponseBodyResultPortMappings) SetTargetPort(v string) *DescribeServiceDetailResponseBodyResultPortMappings {
	s.TargetPort = &v
	return s
}

type DescribeServiceDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceDetailResponse) SetHeaders(v map[string]*string) *DescribeServiceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceDetailResponse) SetStatusCode(v int32) *DescribeServiceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceDetailResponse) SetBody(v *DescribeServiceDetailResponseBody) *DescribeServiceDetailResponse {
	s.Body = v
	return s
}

type DescribeSlbAPDetailRequest struct {
	SlbAPId *int64 `json:"SlbAPId,omitempty" xml:"SlbAPId,omitempty"`
}

func (s DescribeSlbAPDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlbAPDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlbAPDetailRequest) SetSlbAPId(v int64) *DescribeSlbAPDetailRequest {
	s.SlbAPId = &v
	return s
}

type DescribeSlbAPDetailResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeSlbAPDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSlbAPDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlbAPDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlbAPDetailResponseBody) SetCode(v int32) *DescribeSlbAPDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBody) SetErrMsg(v string) *DescribeSlbAPDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBody) SetRequestId(v string) *DescribeSlbAPDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBody) SetResult(v *DescribeSlbAPDetailResponseBodyResult) *DescribeSlbAPDetailResponseBody {
	s.Result = v
	return s
}

func (s *DescribeSlbAPDetailResponseBody) SetSuccess(v bool) *DescribeSlbAPDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeSlbAPDetailResponseBodyResult struct {
	AppId              *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CookieTimeout      *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	EnvId              *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EstablishedTimeout *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkMode        *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	Protocol           *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServerPort     *int32  `json:"RealServerPort,omitempty" xml:"RealServerPort,omitempty"`
	SlbAPId            *int64  `json:"SlbAPId,omitempty" xml:"SlbAPId,omitempty"`
	SlbId              *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbIp              *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SslCertId          *string `json:"SslCertId,omitempty" xml:"SslCertId,omitempty"`
	StickySession      *int32  `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
}

func (s DescribeSlbAPDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlbAPDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetAppId(v int64) *DescribeSlbAPDetailResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetCookieTimeout(v int32) *DescribeSlbAPDetailResponseBodyResult {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetEnvId(v int64) *DescribeSlbAPDetailResponseBodyResult {
	s.EnvId = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetEstablishedTimeout(v int32) *DescribeSlbAPDetailResponseBodyResult {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetListenerPort(v int32) *DescribeSlbAPDetailResponseBodyResult {
	s.ListenerPort = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetName(v string) *DescribeSlbAPDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetNetworkMode(v string) *DescribeSlbAPDetailResponseBodyResult {
	s.NetworkMode = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetProtocol(v string) *DescribeSlbAPDetailResponseBodyResult {
	s.Protocol = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetRealServerPort(v int32) *DescribeSlbAPDetailResponseBodyResult {
	s.RealServerPort = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetSlbAPId(v int64) *DescribeSlbAPDetailResponseBodyResult {
	s.SlbAPId = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetSlbId(v string) *DescribeSlbAPDetailResponseBodyResult {
	s.SlbId = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetSlbIp(v string) *DescribeSlbAPDetailResponseBodyResult {
	s.SlbIp = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetSslCertId(v string) *DescribeSlbAPDetailResponseBodyResult {
	s.SslCertId = &v
	return s
}

func (s *DescribeSlbAPDetailResponseBodyResult) SetStickySession(v int32) *DescribeSlbAPDetailResponseBodyResult {
	s.StickySession = &v
	return s
}

type DescribeSlbAPDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlbAPDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlbAPDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlbAPDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlbAPDetailResponse) SetHeaders(v map[string]*string) *DescribeSlbAPDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlbAPDetailResponse) SetStatusCode(v int32) *DescribeSlbAPDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlbAPDetailResponse) SetBody(v *DescribeSlbAPDetailResponseBody) *DescribeSlbAPDetailResponse {
	s.Body = v
	return s
}

type GetInstTransInfoRequest struct {
	AliyunCommodityCode *string `json:"aliyunCommodityCode,omitempty" xml:"aliyunCommodityCode,omitempty"`
	AliyunEquipId       *string `json:"aliyunEquipId,omitempty" xml:"aliyunEquipId,omitempty"`
	AliyunUid           *string `json:"aliyunUid,omitempty" xml:"aliyunUid,omitempty"`
}

func (s GetInstTransInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstTransInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInstTransInfoRequest) SetAliyunCommodityCode(v string) *GetInstTransInfoRequest {
	s.AliyunCommodityCode = &v
	return s
}

func (s *GetInstTransInfoRequest) SetAliyunEquipId(v string) *GetInstTransInfoRequest {
	s.AliyunEquipId = &v
	return s
}

func (s *GetInstTransInfoRequest) SetAliyunUid(v string) *GetInstTransInfoRequest {
	s.AliyunUid = &v
	return s
}

type GetInstTransInfoResponseBody struct {
	ChargeType  *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsAutoRenew *bool   `json:"isAutoRenew,omitempty" xml:"isAutoRenew,omitempty"`
	RenewCycle  *int32  `json:"renewCycle,omitempty" xml:"renewCycle,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstTransInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstTransInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstTransInfoResponseBody) SetChargeType(v string) *GetInstTransInfoResponseBody {
	s.ChargeType = &v
	return s
}

func (s *GetInstTransInfoResponseBody) SetEndTime(v int64) *GetInstTransInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetInstTransInfoResponseBody) SetInstanceId(v string) *GetInstTransInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstTransInfoResponseBody) SetIsAutoRenew(v bool) *GetInstTransInfoResponseBody {
	s.IsAutoRenew = &v
	return s
}

func (s *GetInstTransInfoResponseBody) SetRenewCycle(v int32) *GetInstTransInfoResponseBody {
	s.RenewCycle = &v
	return s
}

func (s *GetInstTransInfoResponseBody) SetStartTime(v int64) *GetInstTransInfoResponseBody {
	s.StartTime = &v
	return s
}

type GetInstTransInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstTransInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstTransInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstTransInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInstTransInfoResponse) SetHeaders(v map[string]*string) *GetInstTransInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInstTransInfoResponse) SetStatusCode(v int32) *GetInstTransInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstTransInfoResponse) SetBody(v *GetInstTransInfoResponseBody) *GetInstTransInfoResponse {
	s.Body = v
	return s
}

type GetRdsBackUpRequest struct {
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupType   *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetRdsBackUpRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRdsBackUpRequest) GoString() string {
	return s.String()
}

func (s *GetRdsBackUpRequest) SetBackupId(v string) *GetRdsBackUpRequest {
	s.BackupId = &v
	return s
}

func (s *GetRdsBackUpRequest) SetBackupType(v string) *GetRdsBackUpRequest {
	s.BackupType = &v
	return s
}

func (s *GetRdsBackUpRequest) SetDbInstanceId(v string) *GetRdsBackUpRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GetRdsBackUpRequest) SetPageNumber(v int32) *GetRdsBackUpRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRdsBackUpRequest) SetPageSize(v int32) *GetRdsBackUpRequest {
	s.PageSize = &v
	return s
}

type GetRdsBackUpResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                         `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetRdsBackUpResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRdsBackUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRdsBackUpResponseBody) GoString() string {
	return s.String()
}

func (s *GetRdsBackUpResponseBody) SetCode(v int32) *GetRdsBackUpResponseBody {
	s.Code = &v
	return s
}

func (s *GetRdsBackUpResponseBody) SetErrMsg(v string) *GetRdsBackUpResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetRdsBackUpResponseBody) SetRequestId(v string) *GetRdsBackUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRdsBackUpResponseBody) SetResult(v *GetRdsBackUpResponseBodyResult) *GetRdsBackUpResponseBody {
	s.Result = v
	return s
}

type GetRdsBackUpResponseBodyResult struct {
	Items            []*GetRdsBackUpResponseBodyResultItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber       *string                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *string                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalBackupSize  *int64                                 `json:"TotalBackupSize,omitempty" xml:"TotalBackupSize,omitempty"`
	TotalRecordCount *string                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s GetRdsBackUpResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRdsBackUpResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRdsBackUpResponseBodyResult) SetItems(v []*GetRdsBackUpResponseBodyResultItems) *GetRdsBackUpResponseBodyResult {
	s.Items = v
	return s
}

func (s *GetRdsBackUpResponseBodyResult) SetPageNumber(v string) *GetRdsBackUpResponseBodyResult {
	s.PageNumber = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResult) SetPageRecordCount(v string) *GetRdsBackUpResponseBodyResult {
	s.PageRecordCount = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResult) SetTotalBackupSize(v int64) *GetRdsBackUpResponseBodyResult {
	s.TotalBackupSize = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResult) SetTotalRecordCount(v string) *GetRdsBackUpResponseBodyResult {
	s.TotalRecordCount = &v
	return s
}

type GetRdsBackUpResponseBodyResultItems struct {
	BackupDBNames          *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	BackupEndTime          *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupExtractionStatus *string `json:"BackupExtractionStatus,omitempty" xml:"BackupExtractionStatus,omitempty"`
	BackupId               *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupLocation         *string `json:"BackupLocation,omitempty" xml:"BackupLocation,omitempty"`
	BackupMethod           *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupMode             *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupScale            *string `json:"BackupScale,omitempty" xml:"BackupScale,omitempty"`
	BackupSize             *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime        *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStatus           *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType             *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	HostInstanceID         *string `json:"HostInstanceID,omitempty" xml:"HostInstanceID,omitempty"`
	MetaStatus             *string `json:"MetaStatus,omitempty" xml:"MetaStatus,omitempty"`
	StoreStatus            *string `json:"StoreStatus,omitempty" xml:"StoreStatus,omitempty"`
	TotalBackupSize        *int64  `json:"TotalBackupSize,omitempty" xml:"TotalBackupSize,omitempty"`
}

func (s GetRdsBackUpResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetRdsBackUpResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupDBNames(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupDBNames = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupEndTime(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupEndTime = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupExtractionStatus(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupExtractionStatus = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupId(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupId = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupLocation(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupLocation = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupMethod(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupMethod = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupMode(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupMode = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupScale(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupScale = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupSize(v int64) *GetRdsBackUpResponseBodyResultItems {
	s.BackupSize = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupStartTime(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupStartTime = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupStatus(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupStatus = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetBackupType(v string) *GetRdsBackUpResponseBodyResultItems {
	s.BackupType = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetDBInstanceId(v string) *GetRdsBackUpResponseBodyResultItems {
	s.DBInstanceId = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetHostInstanceID(v string) *GetRdsBackUpResponseBodyResultItems {
	s.HostInstanceID = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetMetaStatus(v string) *GetRdsBackUpResponseBodyResultItems {
	s.MetaStatus = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetStoreStatus(v string) *GetRdsBackUpResponseBodyResultItems {
	s.StoreStatus = &v
	return s
}

func (s *GetRdsBackUpResponseBodyResultItems) SetTotalBackupSize(v int64) *GetRdsBackUpResponseBodyResultItems {
	s.TotalBackupSize = &v
	return s
}

type GetRdsBackUpResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRdsBackUpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRdsBackUpResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRdsBackUpResponse) GoString() string {
	return s.String()
}

func (s *GetRdsBackUpResponse) SetHeaders(v map[string]*string) *GetRdsBackUpResponse {
	s.Headers = v
	return s
}

func (s *GetRdsBackUpResponse) SetStatusCode(v int32) *GetRdsBackUpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRdsBackUpResponse) SetBody(v *GetRdsBackUpResponseBody) *GetRdsBackUpResponse {
	s.Body = v
	return s
}

type GrantDbToAccountRequest struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	DbInstanceId     *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbName           *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s GrantDbToAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantDbToAccountRequest) GoString() string {
	return s.String()
}

func (s *GrantDbToAccountRequest) SetAccountName(v string) *GrantDbToAccountRequest {
	s.AccountName = &v
	return s
}

func (s *GrantDbToAccountRequest) SetAccountPrivilege(v string) *GrantDbToAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *GrantDbToAccountRequest) SetDbInstanceId(v string) *GrantDbToAccountRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GrantDbToAccountRequest) SetDbName(v string) *GrantDbToAccountRequest {
	s.DbName = &v
	return s
}

type GrantDbToAccountResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantDbToAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantDbToAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GrantDbToAccountResponseBody) SetCode(v int32) *GrantDbToAccountResponseBody {
	s.Code = &v
	return s
}

func (s *GrantDbToAccountResponseBody) SetErrMsg(v string) *GrantDbToAccountResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GrantDbToAccountResponseBody) SetRequestId(v string) *GrantDbToAccountResponseBody {
	s.RequestId = &v
	return s
}

type GrantDbToAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantDbToAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantDbToAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantDbToAccountResponse) GoString() string {
	return s.String()
}

func (s *GrantDbToAccountResponse) SetHeaders(v map[string]*string) *GrantDbToAccountResponse {
	s.Headers = v
	return s
}

func (s *GrantDbToAccountResponse) SetStatusCode(v int32) *GrantDbToAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantDbToAccountResponse) SetBody(v *GrantDbToAccountResponseBody) *GrantDbToAccountResponse {
	s.Body = v
	return s
}

type ListAppRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppRequest) GoString() string {
	return s.String()
}

func (s *ListAppRequest) SetPageNumber(v int32) *ListAppRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppRequest) SetPageSize(v int32) *ListAppRequest {
	s.PageSize = &v
	return s
}

type ListAppResponseBody struct {
	Code       *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppResponseBody) SetCode(v int32) *ListAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppResponseBody) SetData(v []*ListAppResponseBodyData) *ListAppResponseBody {
	s.Data = v
	return s
}

func (s *ListAppResponseBody) SetErrorMsg(v string) *ListAppResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppResponseBody) SetRequestId(v string) *ListAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppResponseBody) SetTotalCount(v int32) *ListAppResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppResponseBodyData struct {
	AppId           *int64                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppStateType    *string                                  `json:"AppStateType,omitempty" xml:"AppStateType,omitempty"`
	BizName         *string                                  `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizTitle        *string                                  `json:"BizTitle,omitempty" xml:"BizTitle,omitempty"`
	DeployType      *string                                  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Description     *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Language        *string                                  `json:"Language,omitempty" xml:"Language,omitempty"`
	MiddleWareList  []*ListAppResponseBodyDataMiddleWareList `json:"MiddleWareList,omitempty" xml:"MiddleWareList,omitempty" type:"Repeated"`
	OperatingSystem *string                                  `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
	ServiceType     *string                                  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Title           *string                                  `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppResponseBodyData) SetAppId(v int64) *ListAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppResponseBodyData) SetAppStateType(v string) *ListAppResponseBodyData {
	s.AppStateType = &v
	return s
}

func (s *ListAppResponseBodyData) SetBizName(v string) *ListAppResponseBodyData {
	s.BizName = &v
	return s
}

func (s *ListAppResponseBodyData) SetBizTitle(v string) *ListAppResponseBodyData {
	s.BizTitle = &v
	return s
}

func (s *ListAppResponseBodyData) SetDeployType(v string) *ListAppResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *ListAppResponseBodyData) SetDescription(v string) *ListAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAppResponseBodyData) SetLanguage(v string) *ListAppResponseBodyData {
	s.Language = &v
	return s
}

func (s *ListAppResponseBodyData) SetMiddleWareList(v []*ListAppResponseBodyDataMiddleWareList) *ListAppResponseBodyData {
	s.MiddleWareList = v
	return s
}

func (s *ListAppResponseBodyData) SetOperatingSystem(v string) *ListAppResponseBodyData {
	s.OperatingSystem = &v
	return s
}

func (s *ListAppResponseBodyData) SetServiceType(v string) *ListAppResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ListAppResponseBodyData) SetTitle(v string) *ListAppResponseBodyData {
	s.Title = &v
	return s
}

type ListAppResponseBodyDataMiddleWareList struct {
	AppId *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code  *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAppResponseBodyDataMiddleWareList) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponseBodyDataMiddleWareList) GoString() string {
	return s.String()
}

func (s *ListAppResponseBodyDataMiddleWareList) SetAppId(v int64) *ListAppResponseBodyDataMiddleWareList {
	s.AppId = &v
	return s
}

func (s *ListAppResponseBodyDataMiddleWareList) SetCode(v int32) *ListAppResponseBodyDataMiddleWareList {
	s.Code = &v
	return s
}

func (s *ListAppResponseBodyDataMiddleWareList) SetName(v string) *ListAppResponseBodyDataMiddleWareList {
	s.Name = &v
	return s
}

type ListAppResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponse) GoString() string {
	return s.String()
}

func (s *ListAppResponse) SetHeaders(v map[string]*string) *ListAppResponse {
	s.Headers = v
	return s
}

func (s *ListAppResponse) SetStatusCode(v int32) *ListAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppResponse) SetBody(v *ListAppResponseBody) *ListAppResponse {
	s.Body = v
	return s
}

type ListAppCmsGroupsRequest struct {
	AppId      *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId      *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppCmsGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppCmsGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAppCmsGroupsRequest) SetAppId(v int64) *ListAppCmsGroupsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppCmsGroupsRequest) SetEnvId(v int64) *ListAppCmsGroupsRequest {
	s.EnvId = &v
	return s
}

func (s *ListAppCmsGroupsRequest) SetPageNumber(v int32) *ListAppCmsGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppCmsGroupsRequest) SetPageSize(v int32) *ListAppCmsGroupsRequest {
	s.PageSize = &v
	return s
}

type ListAppCmsGroupsResponseBody struct {
	Code       *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppCmsGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppCmsGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppCmsGroupsResponseBody) SetCode(v int32) *ListAppCmsGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppCmsGroupsResponseBody) SetData(v []*string) *ListAppCmsGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppCmsGroupsResponseBody) SetErrorMsg(v string) *ListAppCmsGroupsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppCmsGroupsResponseBody) SetPageNumber(v int32) *ListAppCmsGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppCmsGroupsResponseBody) SetPageSize(v int32) *ListAppCmsGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppCmsGroupsResponseBody) SetRequestId(v string) *ListAppCmsGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppCmsGroupsResponseBody) SetTotalCount(v int64) *ListAppCmsGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppCmsGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppCmsGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppCmsGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppCmsGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAppCmsGroupsResponse) SetHeaders(v map[string]*string) *ListAppCmsGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAppCmsGroupsResponse) SetStatusCode(v int32) *ListAppCmsGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppCmsGroupsResponse) SetBody(v *ListAppCmsGroupsResponseBody) *ListAppCmsGroupsResponse {
	s.Body = v
	return s
}

type ListAppEnvironmentRequest struct {
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvName    *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvType    *int32  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *ListAppEnvironmentRequest) SetAppId(v int64) *ListAppEnvironmentRequest {
	s.AppId = &v
	return s
}

func (s *ListAppEnvironmentRequest) SetEnvName(v string) *ListAppEnvironmentRequest {
	s.EnvName = &v
	return s
}

func (s *ListAppEnvironmentRequest) SetEnvType(v int32) *ListAppEnvironmentRequest {
	s.EnvType = &v
	return s
}

func (s *ListAppEnvironmentRequest) SetPageNumber(v int32) *ListAppEnvironmentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppEnvironmentRequest) SetPageSize(v int32) *ListAppEnvironmentRequest {
	s.PageSize = &v
	return s
}

type ListAppEnvironmentResponseBody struct {
	Code       *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAppEnvironmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppEnvironmentResponseBody) SetCode(v int32) *ListAppEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppEnvironmentResponseBody) SetData(v []*ListAppEnvironmentResponseBodyData) *ListAppEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *ListAppEnvironmentResponseBody) SetErrorMsg(v string) *ListAppEnvironmentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppEnvironmentResponseBody) SetPageNumber(v int32) *ListAppEnvironmentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppEnvironmentResponseBody) SetPageSize(v int32) *ListAppEnvironmentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppEnvironmentResponseBody) SetRequestId(v string) *ListAppEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppEnvironmentResponseBody) SetTotalCount(v int64) *ListAppEnvironmentResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppEnvironmentResponseBodyData struct {
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSchemaId *int64  `json:"AppSchemaId,omitempty" xml:"AppSchemaId,omitempty"`
	EnvId       *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvName     *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvType     *int32  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	EnvTypeName *string `json:"EnvTypeName,omitempty" xml:"EnvTypeName,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListAppEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppEnvironmentResponseBodyData) SetAppId(v int64) *ListAppEnvironmentResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppEnvironmentResponseBodyData) SetAppSchemaId(v int64) *ListAppEnvironmentResponseBodyData {
	s.AppSchemaId = &v
	return s
}

func (s *ListAppEnvironmentResponseBodyData) SetEnvId(v int64) *ListAppEnvironmentResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *ListAppEnvironmentResponseBodyData) SetEnvName(v string) *ListAppEnvironmentResponseBodyData {
	s.EnvName = &v
	return s
}

func (s *ListAppEnvironmentResponseBodyData) SetEnvType(v int32) *ListAppEnvironmentResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *ListAppEnvironmentResponseBodyData) SetEnvTypeName(v string) *ListAppEnvironmentResponseBodyData {
	s.EnvTypeName = &v
	return s
}

func (s *ListAppEnvironmentResponseBodyData) SetRegion(v string) *ListAppEnvironmentResponseBodyData {
	s.Region = &v
	return s
}

type ListAppEnvironmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *ListAppEnvironmentResponse) SetHeaders(v map[string]*string) *ListAppEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *ListAppEnvironmentResponse) SetStatusCode(v int32) *ListAppEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppEnvironmentResponse) SetBody(v *ListAppEnvironmentResponseBody) *ListAppEnvironmentResponse {
	s.Body = v
	return s
}

type ListAppGroupRequest struct {
	BizCode    *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupRequest) SetBizCode(v string) *ListAppGroupRequest {
	s.BizCode = &v
	return s
}

func (s *ListAppGroupRequest) SetPageNumber(v int32) *ListAppGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupRequest) SetPageSize(v int32) *ListAppGroupRequest {
	s.PageSize = &v
	return s
}

type ListAppGroupResponseBody struct {
	Code       *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAppGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupResponseBody) SetCode(v int32) *ListAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppGroupResponseBody) SetData(v []*ListAppGroupResponseBodyData) *ListAppGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListAppGroupResponseBody) SetErrorMsg(v string) *ListAppGroupResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppGroupResponseBody) SetPageNumber(v int32) *ListAppGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupResponseBody) SetPageSize(v int32) *ListAppGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupResponseBody) SetRequestId(v string) *ListAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupResponseBody) SetTotalCount(v int64) *ListAppGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppGroupResponseBodyData struct {
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAppGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppGroupResponseBodyData) SetBizName(v string) *ListAppGroupResponseBodyData {
	s.BizName = &v
	return s
}

func (s *ListAppGroupResponseBodyData) SetId(v int64) *ListAppGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppGroupResponseBodyData) SetName(v string) *ListAppGroupResponseBodyData {
	s.Name = &v
	return s
}

type ListAppGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupResponse) SetHeaders(v map[string]*string) *ListAppGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupResponse) SetStatusCode(v int32) *ListAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppGroupResponse) SetBody(v *ListAppGroupResponseBody) *ListAppGroupResponse {
	s.Body = v
	return s
}

type ListAppGroupMappingRequest struct {
	BizCode    *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppGroupMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMappingRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupMappingRequest) SetBizCode(v string) *ListAppGroupMappingRequest {
	s.BizCode = &v
	return s
}

func (s *ListAppGroupMappingRequest) SetName(v string) *ListAppGroupMappingRequest {
	s.Name = &v
	return s
}

func (s *ListAppGroupMappingRequest) SetPageNumber(v int32) *ListAppGroupMappingRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupMappingRequest) SetPageSize(v int32) *ListAppGroupMappingRequest {
	s.PageSize = &v
	return s
}

type ListAppGroupMappingResponseBody struct {
	Code       *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAppGroupMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppGroupMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupMappingResponseBody) SetCode(v int32) *ListAppGroupMappingResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppGroupMappingResponseBody) SetData(v []*ListAppGroupMappingResponseBodyData) *ListAppGroupMappingResponseBody {
	s.Data = v
	return s
}

func (s *ListAppGroupMappingResponseBody) SetErrorMsg(v string) *ListAppGroupMappingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppGroupMappingResponseBody) SetPageNumber(v int32) *ListAppGroupMappingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupMappingResponseBody) SetPageSize(v int32) *ListAppGroupMappingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupMappingResponseBody) SetRequestId(v string) *ListAppGroupMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupMappingResponseBody) SetTotalCount(v int64) *ListAppGroupMappingResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppGroupMappingResponseBodyData struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GroupId *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAppGroupMappingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppGroupMappingResponseBodyData) SetAppId(v int64) *ListAppGroupMappingResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppGroupMappingResponseBodyData) SetGroupId(v int64) *ListAppGroupMappingResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListAppGroupMappingResponseBodyData) SetId(v int64) *ListAppGroupMappingResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppGroupMappingResponseBodyData) SetName(v string) *ListAppGroupMappingResponseBodyData {
	s.Name = &v
	return s
}

type ListAppGroupMappingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppGroupMappingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppGroupMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupMappingResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupMappingResponse) SetHeaders(v map[string]*string) *ListAppGroupMappingResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupMappingResponse) SetStatusCode(v int32) *ListAppGroupMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppGroupMappingResponse) SetBody(v *ListAppGroupMappingResponseBody) *ListAppGroupMappingResponse {
	s.Body = v
	return s
}

type ListAppInstanceRequest struct {
	AppId      *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId      *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceRequest) SetAppId(v int64) *ListAppInstanceRequest {
	s.AppId = &v
	return s
}

func (s *ListAppInstanceRequest) SetEnvId(v int64) *ListAppInstanceRequest {
	s.EnvId = &v
	return s
}

func (s *ListAppInstanceRequest) SetPageNumber(v int32) *ListAppInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceRequest) SetPageSize(v int32) *ListAppInstanceRequest {
	s.PageSize = &v
	return s
}

type ListAppInstanceResponseBody struct {
	Code       *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAppInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrMsg     *string                            `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceResponseBody) SetCode(v int32) *ListAppInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppInstanceResponseBody) SetData(v []*ListAppInstanceResponseBodyData) *ListAppInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListAppInstanceResponseBody) SetErrMsg(v string) *ListAppInstanceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListAppInstanceResponseBody) SetPageNumber(v int32) *ListAppInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceResponseBody) SetPageSize(v int32) *ListAppInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceResponseBody) SetRequestId(v string) *ListAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceResponseBody) SetTotalCount(v int64) *ListAppInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppInstanceResponseBodyData struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Health        *string `json:"Health,omitempty" xml:"Health,omitempty"`
	HostIp        *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	Limits        *string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	PodIp         *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	Requests      *string `json:"Requests,omitempty" xml:"Requests,omitempty"`
	RestartCount  *int32  `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	Spec          *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppInstanceResponseBodyData) SetAppInstanceId(v string) *ListAppInstanceResponseBodyData {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetCreateTime(v string) *ListAppInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetHealth(v string) *ListAppInstanceResponseBodyData {
	s.Health = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetHostIp(v string) *ListAppInstanceResponseBodyData {
	s.HostIp = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetLimits(v string) *ListAppInstanceResponseBodyData {
	s.Limits = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetPodIp(v string) *ListAppInstanceResponseBodyData {
	s.PodIp = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetRequests(v string) *ListAppInstanceResponseBodyData {
	s.Requests = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetRestartCount(v int32) *ListAppInstanceResponseBodyData {
	s.RestartCount = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetSpec(v string) *ListAppInstanceResponseBodyData {
	s.Spec = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetStatus(v string) *ListAppInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAppInstanceResponseBodyData) SetVersion(v string) *ListAppInstanceResponseBodyData {
	s.Version = &v
	return s
}

type ListAppInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceResponse) SetHeaders(v map[string]*string) *ListAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceResponse) SetStatusCode(v int32) *ListAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceResponse) SetBody(v *ListAppInstanceResponseBody) *ListAppInstanceResponse {
	s.Body = v
	return s
}

type ListAppResourceAllocsRequest struct {
	AppEnvId   *int64  `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppResourceAllocsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppResourceAllocsRequest) GoString() string {
	return s.String()
}

func (s *ListAppResourceAllocsRequest) SetAppEnvId(v int64) *ListAppResourceAllocsRequest {
	s.AppEnvId = &v
	return s
}

func (s *ListAppResourceAllocsRequest) SetAppId(v int64) *ListAppResourceAllocsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppResourceAllocsRequest) SetClusterId(v string) *ListAppResourceAllocsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAppResourceAllocsRequest) SetPageNumber(v int32) *ListAppResourceAllocsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppResourceAllocsRequest) SetPageSize(v int32) *ListAppResourceAllocsRequest {
	s.PageSize = &v
	return s
}

type ListAppResourceAllocsResponseBody struct {
	Code       *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAppResourceAllocsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppResourceAllocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppResourceAllocsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppResourceAllocsResponseBody) SetCode(v int32) *ListAppResourceAllocsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppResourceAllocsResponseBody) SetData(v []*ListAppResourceAllocsResponseBodyData) *ListAppResourceAllocsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppResourceAllocsResponseBody) SetErrorMsg(v string) *ListAppResourceAllocsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppResourceAllocsResponseBody) SetPageNumber(v int32) *ListAppResourceAllocsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppResourceAllocsResponseBody) SetPageSize(v int32) *ListAppResourceAllocsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppResourceAllocsResponseBody) SetRequestId(v string) *ListAppResourceAllocsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppResourceAllocsResponseBody) SetTotalCount(v int64) *ListAppResourceAllocsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppResourceAllocsResponseBodyData struct {
	AppEnvId    *int64  `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ResourceDef *string `json:"ResourceDef,omitempty" xml:"ResourceDef,omitempty"`
}

func (s ListAppResourceAllocsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppResourceAllocsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppResourceAllocsResponseBodyData) SetAppEnvId(v int64) *ListAppResourceAllocsResponseBodyData {
	s.AppEnvId = &v
	return s
}

func (s *ListAppResourceAllocsResponseBodyData) SetAppId(v int64) *ListAppResourceAllocsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppResourceAllocsResponseBodyData) SetClusterId(v string) *ListAppResourceAllocsResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListAppResourceAllocsResponseBodyData) SetId(v int64) *ListAppResourceAllocsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppResourceAllocsResponseBodyData) SetResourceDef(v string) *ListAppResourceAllocsResponseBodyData {
	s.ResourceDef = &v
	return s
}

type ListAppResourceAllocsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppResourceAllocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppResourceAllocsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppResourceAllocsResponse) GoString() string {
	return s.String()
}

func (s *ListAppResourceAllocsResponse) SetHeaders(v map[string]*string) *ListAppResourceAllocsResponse {
	s.Headers = v
	return s
}

func (s *ListAppResourceAllocsResponse) SetStatusCode(v int32) *ListAppResourceAllocsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppResourceAllocsResponse) SetBody(v *ListAppResourceAllocsResponseBody) *ListAppResourceAllocsResponse {
	s.Body = v
	return s
}

type ListAvailableClusterNodeRequest struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAvailableClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableClusterNodeRequest) SetClusterInstanceId(v string) *ListAvailableClusterNodeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *ListAvailableClusterNodeRequest) SetPageNum(v int32) *ListAvailableClusterNodeRequest {
	s.PageNum = &v
	return s
}

func (s *ListAvailableClusterNodeRequest) SetPageSize(v int32) *ListAvailableClusterNodeRequest {
	s.PageSize = &v
	return s
}

type ListAvailableClusterNodeResponseBody struct {
	Code       *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListAvailableClusterNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvailableClusterNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableClusterNodeResponseBody) SetCode(v int32) *ListAvailableClusterNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBody) SetData(v []*ListAvailableClusterNodeResponseBodyData) *ListAvailableClusterNodeResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableClusterNodeResponseBody) SetErrorMsg(v string) *ListAvailableClusterNodeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBody) SetPageNumber(v int32) *ListAvailableClusterNodeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBody) SetPageSize(v int32) *ListAvailableClusterNodeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBody) SetRequestId(v string) *ListAvailableClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBody) SetTotalCount(v int64) *ListAvailableClusterNodeResponseBody {
	s.TotalCount = &v
	return s
}

type ListAvailableClusterNodeResponseBodyData struct {
	BusinessCode            *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	EcsConfiguration        *string `json:"EcsConfiguration,omitempty" xml:"EcsConfiguration,omitempty"`
	EcsCpu                  *string `json:"EcsCpu,omitempty" xml:"EcsCpu,omitempty"`
	EcsEip                  *string `json:"EcsEip,omitempty" xml:"EcsEip,omitempty"`
	EcsExpiredTime          *string `json:"EcsExpiredTime,omitempty" xml:"EcsExpiredTime,omitempty"`
	EcsLocalStorageCapacity *string `json:"EcsLocalStorageCapacity,omitempty" xml:"EcsLocalStorageCapacity,omitempty"`
	EcsMemory               *string `json:"EcsMemory,omitempty" xml:"EcsMemory,omitempty"`
	EcsOsType               *string `json:"EcsOsType,omitempty" xml:"EcsOsType,omitempty"`
	EcsPrivateIp            *string `json:"EcsPrivateIp,omitempty" xml:"EcsPrivateIp,omitempty"`
	EcsPublicIp             *string `json:"EcsPublicIp,omitempty" xml:"EcsPublicIp,omitempty"`
	EcsZone                 *string `json:"EcsZone,omitempty" xml:"EcsZone,omitempty"`
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName            *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceNetworkType     *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetMaxBandwidthIn  *string `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *string `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	OSName                  *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAvailableClusterNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableClusterNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableClusterNodeResponseBodyData) SetBusinessCode(v string) *ListAvailableClusterNodeResponseBodyData {
	s.BusinessCode = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsConfiguration(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsConfiguration = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsCpu(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsCpu = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsEip(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsEip = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsExpiredTime(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsExpiredTime = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsLocalStorageCapacity(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsLocalStorageCapacity = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsMemory(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsMemory = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsOsType(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsOsType = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsPrivateIp(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsPrivateIp = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsPublicIp(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsPublicIp = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetEcsZone(v string) *ListAvailableClusterNodeResponseBodyData {
	s.EcsZone = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetInstanceId(v string) *ListAvailableClusterNodeResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetInstanceName(v string) *ListAvailableClusterNodeResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetInstanceNetworkType(v string) *ListAvailableClusterNodeResponseBodyData {
	s.InstanceNetworkType = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetInstanceType(v string) *ListAvailableClusterNodeResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetInternetMaxBandwidthIn(v string) *ListAvailableClusterNodeResponseBodyData {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetInternetMaxBandwidthOut(v string) *ListAvailableClusterNodeResponseBodyData {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetOSName(v string) *ListAvailableClusterNodeResponseBodyData {
	s.OSName = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetRegionId(v string) *ListAvailableClusterNodeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListAvailableClusterNodeResponseBodyData) SetVpcId(v string) *ListAvailableClusterNodeResponseBodyData {
	s.VpcId = &v
	return s
}

type ListAvailableClusterNodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAvailableClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableClusterNodeResponse) SetHeaders(v map[string]*string) *ListAvailableClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableClusterNodeResponse) SetStatusCode(v int32) *ListAvailableClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableClusterNodeResponse) SetBody(v *ListAvailableClusterNodeResponseBody) *ListAvailableClusterNodeResponse {
	s.Body = v
	return s
}

type ListClusterRequest struct {
	BusinessCode *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	EnvType      *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterRequest) GoString() string {
	return s.String()
}

func (s *ListClusterRequest) SetBusinessCode(v string) *ListClusterRequest {
	s.BusinessCode = &v
	return s
}

func (s *ListClusterRequest) SetEnvType(v string) *ListClusterRequest {
	s.EnvType = &v
	return s
}

func (s *ListClusterRequest) SetPageNum(v int32) *ListClusterRequest {
	s.PageNum = &v
	return s
}

func (s *ListClusterRequest) SetPageSize(v int32) *ListClusterRequest {
	s.PageSize = &v
	return s
}

type ListClusterResponseBody struct {
	Code       *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                        `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBody) SetCode(v int32) *ListClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterResponseBody) SetData(v []*ListClusterResponseBodyData) *ListClusterResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterResponseBody) SetErrorMsg(v string) *ListClusterResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListClusterResponseBody) SetPageNumber(v int32) *ListClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterResponseBody) SetPageSize(v int32) *ListClusterResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterResponseBody) SetRequestId(v string) *ListClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterResponseBody) SetTotalCount(v int64) *ListClusterResponseBody {
	s.TotalCount = &v
	return s
}

type ListClusterResponseBodyData struct {
	BusinessCode *string   `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	ClusterTitle *string   `json:"ClusterTitle,omitempty" xml:"ClusterTitle,omitempty"`
	CreateStatus *string   `json:"CreateStatus,omitempty" xml:"CreateStatus,omitempty"`
	EcsIds       []*string `json:"EcsIds,omitempty" xml:"EcsIds,omitempty" type:"Repeated"`
	EnvType      *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Id           *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeyPair      *string   `json:"KeyPair,omitempty" xml:"KeyPair,omitempty"`
	NetPlug      *string   `json:"NetPlug,omitempty" xml:"NetPlug,omitempty"`
	Password     *string   `json:"Password,omitempty" xml:"Password,omitempty"`
	PodCIDR      *string   `json:"PodCIDR,omitempty" xml:"PodCIDR,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName   *string   `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ServiceCIDR  *string   `json:"ServiceCIDR,omitempty" xml:"ServiceCIDR,omitempty"`
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId        *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchIds   []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
	WorkLoadCpu  *string   `json:"WorkLoadCpu,omitempty" xml:"WorkLoadCpu,omitempty"`
	WorkLoadMem  *string   `json:"WorkLoadMem,omitempty" xml:"WorkLoadMem,omitempty"`
}

func (s ListClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBodyData) SetBusinessCode(v string) *ListClusterResponseBodyData {
	s.BusinessCode = &v
	return s
}

func (s *ListClusterResponseBodyData) SetClusterTitle(v string) *ListClusterResponseBodyData {
	s.ClusterTitle = &v
	return s
}

func (s *ListClusterResponseBodyData) SetCreateStatus(v string) *ListClusterResponseBodyData {
	s.CreateStatus = &v
	return s
}

func (s *ListClusterResponseBodyData) SetEcsIds(v []*string) *ListClusterResponseBodyData {
	s.EcsIds = v
	return s
}

func (s *ListClusterResponseBodyData) SetEnvType(v string) *ListClusterResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *ListClusterResponseBodyData) SetId(v int64) *ListClusterResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListClusterResponseBodyData) SetInstanceId(v string) *ListClusterResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListClusterResponseBodyData) SetKeyPair(v string) *ListClusterResponseBodyData {
	s.KeyPair = &v
	return s
}

func (s *ListClusterResponseBodyData) SetNetPlug(v string) *ListClusterResponseBodyData {
	s.NetPlug = &v
	return s
}

func (s *ListClusterResponseBodyData) SetPassword(v string) *ListClusterResponseBodyData {
	s.Password = &v
	return s
}

func (s *ListClusterResponseBodyData) SetPodCIDR(v string) *ListClusterResponseBodyData {
	s.PodCIDR = &v
	return s
}

func (s *ListClusterResponseBodyData) SetRegionId(v string) *ListClusterResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListClusterResponseBodyData) SetRegionName(v string) *ListClusterResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListClusterResponseBodyData) SetServiceCIDR(v string) *ListClusterResponseBodyData {
	s.ServiceCIDR = &v
	return s
}

func (s *ListClusterResponseBodyData) SetStatus(v string) *ListClusterResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListClusterResponseBodyData) SetVpcId(v string) *ListClusterResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *ListClusterResponseBodyData) SetVswitchIds(v []*string) *ListClusterResponseBodyData {
	s.VswitchIds = v
	return s
}

func (s *ListClusterResponseBodyData) SetWorkLoadCpu(v string) *ListClusterResponseBodyData {
	s.WorkLoadCpu = &v
	return s
}

func (s *ListClusterResponseBodyData) SetWorkLoadMem(v string) *ListClusterResponseBodyData {
	s.WorkLoadMem = &v
	return s
}

type ListClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponse) GoString() string {
	return s.String()
}

func (s *ListClusterResponse) SetHeaders(v map[string]*string) *ListClusterResponse {
	s.Headers = v
	return s
}

func (s *ListClusterResponse) SetStatusCode(v int32) *ListClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterResponse) SetBody(v *ListClusterResponseBody) *ListClusterResponse {
	s.Body = v
	return s
}

type ListClusterNodeRequest struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *ListClusterNodeRequest) SetClusterInstanceId(v string) *ListClusterNodeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *ListClusterNodeRequest) SetPageNum(v int32) *ListClusterNodeRequest {
	s.PageNum = &v
	return s
}

func (s *ListClusterNodeRequest) SetPageSize(v int32) *ListClusterNodeRequest {
	s.PageSize = &v
	return s
}

type ListClusterNodeResponseBody struct {
	Code       *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListClusterNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                            `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNodeResponseBody) SetCode(v int32) *ListClusterNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterNodeResponseBody) SetData(v []*ListClusterNodeResponseBodyData) *ListClusterNodeResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterNodeResponseBody) SetErrorMsg(v string) *ListClusterNodeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListClusterNodeResponseBody) SetPageNumber(v int32) *ListClusterNodeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterNodeResponseBody) SetPageSize(v int32) *ListClusterNodeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterNodeResponseBody) SetRequestId(v string) *ListClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterNodeResponseBody) SetTotalCount(v int64) *ListClusterNodeResponseBody {
	s.TotalCount = &v
	return s
}

type ListClusterNodeResponseBodyData struct {
	BusinessCode            *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	EcsConfiguration        *string `json:"EcsConfiguration,omitempty" xml:"EcsConfiguration,omitempty"`
	EcsCpu                  *string `json:"EcsCpu,omitempty" xml:"EcsCpu,omitempty"`
	EcsEip                  *string `json:"EcsEip,omitempty" xml:"EcsEip,omitempty"`
	EcsExpiredTime          *string `json:"EcsExpiredTime,omitempty" xml:"EcsExpiredTime,omitempty"`
	EcsLocalStorageCapacity *string `json:"EcsLocalStorageCapacity,omitempty" xml:"EcsLocalStorageCapacity,omitempty"`
	EcsMemory               *string `json:"EcsMemory,omitempty" xml:"EcsMemory,omitempty"`
	EcsOsType               *string `json:"EcsOsType,omitempty" xml:"EcsOsType,omitempty"`
	EcsPrivateIp            *string `json:"EcsPrivateIp,omitempty" xml:"EcsPrivateIp,omitempty"`
	EcsPublicIp             *string `json:"EcsPublicIp,omitempty" xml:"EcsPublicIp,omitempty"`
	EcsZone                 *string `json:"EcsZone,omitempty" xml:"EcsZone,omitempty"`
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName            *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceNetworkType     *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetMaxBandwidthIn  *string `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *string `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	OSName                  *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClusterNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterNodeResponseBodyData) SetBusinessCode(v string) *ListClusterNodeResponseBodyData {
	s.BusinessCode = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsConfiguration(v string) *ListClusterNodeResponseBodyData {
	s.EcsConfiguration = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsCpu(v string) *ListClusterNodeResponseBodyData {
	s.EcsCpu = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsEip(v string) *ListClusterNodeResponseBodyData {
	s.EcsEip = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsExpiredTime(v string) *ListClusterNodeResponseBodyData {
	s.EcsExpiredTime = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsLocalStorageCapacity(v string) *ListClusterNodeResponseBodyData {
	s.EcsLocalStorageCapacity = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsMemory(v string) *ListClusterNodeResponseBodyData {
	s.EcsMemory = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsOsType(v string) *ListClusterNodeResponseBodyData {
	s.EcsOsType = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsPrivateIp(v string) *ListClusterNodeResponseBodyData {
	s.EcsPrivateIp = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsPublicIp(v string) *ListClusterNodeResponseBodyData {
	s.EcsPublicIp = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetEcsZone(v string) *ListClusterNodeResponseBodyData {
	s.EcsZone = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetInstanceId(v string) *ListClusterNodeResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetInstanceName(v string) *ListClusterNodeResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetInstanceNetworkType(v string) *ListClusterNodeResponseBodyData {
	s.InstanceNetworkType = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetInstanceType(v string) *ListClusterNodeResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetInternetMaxBandwidthIn(v string) *ListClusterNodeResponseBodyData {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetInternetMaxBandwidthOut(v string) *ListClusterNodeResponseBodyData {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetOSName(v string) *ListClusterNodeResponseBodyData {
	s.OSName = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetRegionId(v string) *ListClusterNodeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListClusterNodeResponseBodyData) SetVpcId(v string) *ListClusterNodeResponseBodyData {
	s.VpcId = &v
	return s
}

type ListClusterNodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNodeResponse) SetHeaders(v map[string]*string) *ListClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNodeResponse) SetStatusCode(v int32) *ListClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNodeResponse) SetBody(v *ListClusterNodeResponseBody) *ListClusterNodeResponse {
	s.Body = v
	return s
}

type ListDeployConfigRequest struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDeployConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployConfigRequest) GoString() string {
	return s.String()
}

func (s *ListDeployConfigRequest) SetAppId(v int64) *ListDeployConfigRequest {
	s.AppId = &v
	return s
}

func (s *ListDeployConfigRequest) SetEnvType(v string) *ListDeployConfigRequest {
	s.EnvType = &v
	return s
}

func (s *ListDeployConfigRequest) SetId(v int64) *ListDeployConfigRequest {
	s.Id = &v
	return s
}

func (s *ListDeployConfigRequest) SetName(v string) *ListDeployConfigRequest {
	s.Name = &v
	return s
}

type ListDeployConfigResponseBody struct {
	Code       *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListDeployConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeployConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployConfigResponseBody) SetCode(v int32) *ListDeployConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployConfigResponseBody) SetData(v []*ListDeployConfigResponseBodyData) *ListDeployConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListDeployConfigResponseBody) SetErrorMsg(v string) *ListDeployConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDeployConfigResponseBody) SetPageNumber(v int32) *ListDeployConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeployConfigResponseBody) SetPageSize(v int32) *ListDeployConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeployConfigResponseBody) SetRequestId(v string) *ListDeployConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployConfigResponseBody) SetTotalCount(v int64) *ListDeployConfigResponseBody {
	s.TotalCount = &v
	return s
}

type ListDeployConfigResponseBodyData struct {
	AppId             *int64                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerCodePath *ListDeployConfigResponseBodyDataContainerCodePath `json:"ContainerCodePath,omitempty" xml:"ContainerCodePath,omitempty" type:"Struct"`
	ContainerYamlConf *ListDeployConfigResponseBodyDataContainerYamlConf `json:"ContainerYamlConf,omitempty" xml:"ContainerYamlConf,omitempty" type:"Struct"`
	EnvType           *string                                            `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Id                *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Name              *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDeployConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeployConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeployConfigResponseBodyData) SetAppId(v int64) *ListDeployConfigResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListDeployConfigResponseBodyData) SetContainerCodePath(v *ListDeployConfigResponseBodyDataContainerCodePath) *ListDeployConfigResponseBodyData {
	s.ContainerCodePath = v
	return s
}

func (s *ListDeployConfigResponseBodyData) SetContainerYamlConf(v *ListDeployConfigResponseBodyDataContainerYamlConf) *ListDeployConfigResponseBodyData {
	s.ContainerYamlConf = v
	return s
}

func (s *ListDeployConfigResponseBodyData) SetEnvType(v string) *ListDeployConfigResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *ListDeployConfigResponseBodyData) SetId(v int64) *ListDeployConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDeployConfigResponseBodyData) SetName(v string) *ListDeployConfigResponseBodyData {
	s.Name = &v
	return s
}

type ListDeployConfigResponseBodyDataContainerCodePath struct {
	CodePath *string `json:"CodePath,omitempty" xml:"CodePath,omitempty"`
}

func (s ListDeployConfigResponseBodyDataContainerCodePath) String() string {
	return tea.Prettify(s)
}

func (s ListDeployConfigResponseBodyDataContainerCodePath) GoString() string {
	return s.String()
}

func (s *ListDeployConfigResponseBodyDataContainerCodePath) SetCodePath(v string) *ListDeployConfigResponseBodyDataContainerCodePath {
	s.CodePath = &v
	return s
}

type ListDeployConfigResponseBodyDataContainerYamlConf struct {
	ConfigMap     *string   `json:"ConfigMap,omitempty" xml:"ConfigMap,omitempty"`
	ConfigMapList []*string `json:"ConfigMapList,omitempty" xml:"ConfigMapList,omitempty" type:"Repeated"`
	CronJob       *string   `json:"CronJob,omitempty" xml:"CronJob,omitempty"`
	Deployment    *string   `json:"Deployment,omitempty" xml:"Deployment,omitempty"`
	SecretList    []*string `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Repeated"`
	StatefulSet   *string   `json:"StatefulSet,omitempty" xml:"StatefulSet,omitempty"`
}

func (s ListDeployConfigResponseBodyDataContainerYamlConf) String() string {
	return tea.Prettify(s)
}

func (s ListDeployConfigResponseBodyDataContainerYamlConf) GoString() string {
	return s.String()
}

func (s *ListDeployConfigResponseBodyDataContainerYamlConf) SetConfigMap(v string) *ListDeployConfigResponseBodyDataContainerYamlConf {
	s.ConfigMap = &v
	return s
}

func (s *ListDeployConfigResponseBodyDataContainerYamlConf) SetConfigMapList(v []*string) *ListDeployConfigResponseBodyDataContainerYamlConf {
	s.ConfigMapList = v
	return s
}

func (s *ListDeployConfigResponseBodyDataContainerYamlConf) SetCronJob(v string) *ListDeployConfigResponseBodyDataContainerYamlConf {
	s.CronJob = &v
	return s
}

func (s *ListDeployConfigResponseBodyDataContainerYamlConf) SetDeployment(v string) *ListDeployConfigResponseBodyDataContainerYamlConf {
	s.Deployment = &v
	return s
}

func (s *ListDeployConfigResponseBodyDataContainerYamlConf) SetSecretList(v []*string) *ListDeployConfigResponseBodyDataContainerYamlConf {
	s.SecretList = v
	return s
}

func (s *ListDeployConfigResponseBodyDataContainerYamlConf) SetStatefulSet(v string) *ListDeployConfigResponseBodyDataContainerYamlConf {
	s.StatefulSet = &v
	return s
}

type ListDeployConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeployConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployConfigResponse) GoString() string {
	return s.String()
}

func (s *ListDeployConfigResponse) SetHeaders(v map[string]*string) *ListDeployConfigResponse {
	s.Headers = v
	return s
}

func (s *ListDeployConfigResponse) SetStatusCode(v int32) *ListDeployConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeployConfigResponse) SetBody(v *ListDeployConfigResponseBody) *ListDeployConfigResponse {
	s.Body = v
	return s
}

type ListDeployOrdersRequest struct {
	AppId                         *int64   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeployCategory                *string  `json:"DeployCategory,omitempty" xml:"DeployCategory,omitempty"`
	DeployType                    *string  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	EndTimeGreaterThan            *string  `json:"EndTimeGreaterThan,omitempty" xml:"EndTimeGreaterThan,omitempty"`
	EndTimeGreaterThanOrEqualTo   *string  `json:"EndTimeGreaterThanOrEqualTo,omitempty" xml:"EndTimeGreaterThanOrEqualTo,omitempty"`
	EndTimeLessThan               *string  `json:"EndTimeLessThan,omitempty" xml:"EndTimeLessThan,omitempty"`
	EndTimeLessThanOrEqualTo      *string  `json:"EndTimeLessThanOrEqualTo,omitempty" xml:"EndTimeLessThanOrEqualTo,omitempty"`
	EnvId                         *int64   `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType                       *string  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	PageNumber                    *int32   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                      *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PartitionType                 *string  `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	PauseType                     *string  `json:"PauseType,omitempty" xml:"PauseType,omitempty"`
	ResultList                    []*int32 `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	StartTimeGreaterThan          *string  `json:"StartTimeGreaterThan,omitempty" xml:"StartTimeGreaterThan,omitempty"`
	StartTimeGreaterThanOrEqualTo *string  `json:"StartTimeGreaterThanOrEqualTo,omitempty" xml:"StartTimeGreaterThanOrEqualTo,omitempty"`
	StartTimeLessThan             *string  `json:"StartTimeLessThan,omitempty" xml:"StartTimeLessThan,omitempty"`
	StartTimeLessThanOrEqualTo    *string  `json:"StartTimeLessThanOrEqualTo,omitempty" xml:"StartTimeLessThanOrEqualTo,omitempty"`
	Status                        *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusList                    []*int32 `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListDeployOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListDeployOrdersRequest) SetAppId(v int64) *ListDeployOrdersRequest {
	s.AppId = &v
	return s
}

func (s *ListDeployOrdersRequest) SetDeployCategory(v string) *ListDeployOrdersRequest {
	s.DeployCategory = &v
	return s
}

func (s *ListDeployOrdersRequest) SetDeployType(v string) *ListDeployOrdersRequest {
	s.DeployType = &v
	return s
}

func (s *ListDeployOrdersRequest) SetEndTimeGreaterThan(v string) *ListDeployOrdersRequest {
	s.EndTimeGreaterThan = &v
	return s
}

func (s *ListDeployOrdersRequest) SetEndTimeGreaterThanOrEqualTo(v string) *ListDeployOrdersRequest {
	s.EndTimeGreaterThanOrEqualTo = &v
	return s
}

func (s *ListDeployOrdersRequest) SetEndTimeLessThan(v string) *ListDeployOrdersRequest {
	s.EndTimeLessThan = &v
	return s
}

func (s *ListDeployOrdersRequest) SetEndTimeLessThanOrEqualTo(v string) *ListDeployOrdersRequest {
	s.EndTimeLessThanOrEqualTo = &v
	return s
}

func (s *ListDeployOrdersRequest) SetEnvId(v int64) *ListDeployOrdersRequest {
	s.EnvId = &v
	return s
}

func (s *ListDeployOrdersRequest) SetEnvType(v string) *ListDeployOrdersRequest {
	s.EnvType = &v
	return s
}

func (s *ListDeployOrdersRequest) SetPageNumber(v int32) *ListDeployOrdersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeployOrdersRequest) SetPageSize(v int32) *ListDeployOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeployOrdersRequest) SetPartitionType(v string) *ListDeployOrdersRequest {
	s.PartitionType = &v
	return s
}

func (s *ListDeployOrdersRequest) SetPauseType(v string) *ListDeployOrdersRequest {
	s.PauseType = &v
	return s
}

func (s *ListDeployOrdersRequest) SetResultList(v []*int32) *ListDeployOrdersRequest {
	s.ResultList = v
	return s
}

func (s *ListDeployOrdersRequest) SetStartTimeGreaterThan(v string) *ListDeployOrdersRequest {
	s.StartTimeGreaterThan = &v
	return s
}

func (s *ListDeployOrdersRequest) SetStartTimeGreaterThanOrEqualTo(v string) *ListDeployOrdersRequest {
	s.StartTimeGreaterThanOrEqualTo = &v
	return s
}

func (s *ListDeployOrdersRequest) SetStartTimeLessThan(v string) *ListDeployOrdersRequest {
	s.StartTimeLessThan = &v
	return s
}

func (s *ListDeployOrdersRequest) SetStartTimeLessThanOrEqualTo(v string) *ListDeployOrdersRequest {
	s.StartTimeLessThanOrEqualTo = &v
	return s
}

func (s *ListDeployOrdersRequest) SetStatus(v int32) *ListDeployOrdersRequest {
	s.Status = &v
	return s
}

func (s *ListDeployOrdersRequest) SetStatusList(v []*int32) *ListDeployOrdersRequest {
	s.StatusList = v
	return s
}

type ListDeployOrdersResponseBody struct {
	Code       *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListDeployOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeployOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployOrdersResponseBody) SetCode(v int32) *ListDeployOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployOrdersResponseBody) SetData(v []*ListDeployOrdersResponseBodyData) *ListDeployOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListDeployOrdersResponseBody) SetErrorMsg(v string) *ListDeployOrdersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDeployOrdersResponseBody) SetPageNumber(v int32) *ListDeployOrdersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeployOrdersResponseBody) SetPageSize(v int32) *ListDeployOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeployOrdersResponseBody) SetRequestId(v string) *ListDeployOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployOrdersResponseBody) SetTotalCount(v int64) *ListDeployOrdersResponseBody {
	s.TotalCount = &v
	return s
}

type ListDeployOrdersResponseBodyData struct {
	AppInstanceType     *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	CurrentPartitionNum *int32  `json:"CurrentPartitionNum,omitempty" xml:"CurrentPartitionNum,omitempty"`
	DeployOrderId       *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	DeployPauseType     *string `json:"DeployPauseType,omitempty" xml:"DeployPauseType,omitempty"`
	DeployPauseTypeName *string `json:"DeployPauseTypeName,omitempty" xml:"DeployPauseTypeName,omitempty"`
	DeployType          *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DeployTypeName      *string `json:"DeployTypeName,omitempty" xml:"DeployTypeName,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ElapsedTime         *int32  `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvId               *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType             *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	FailureRate         *string `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	FinishAppInstanceCt *int32  `json:"FinishAppInstanceCt,omitempty" xml:"FinishAppInstanceCt,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionType       *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	PartitionTypeName   *string `json:"PartitionTypeName,omitempty" xml:"PartitionTypeName,omitempty"`
	Result              *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultName          *string `json:"ResultName,omitempty" xml:"ResultName,omitempty"`
	SchemaId            *int64  `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName          *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TotalAppInstanceCt  *int32  `json:"TotalAppInstanceCt,omitempty" xml:"TotalAppInstanceCt,omitempty"`
	TotalPartitions     *int32  `json:"TotalPartitions,omitempty" xml:"TotalPartitions,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick            *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s ListDeployOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeployOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeployOrdersResponseBodyData) SetAppInstanceType(v string) *ListDeployOrdersResponseBodyData {
	s.AppInstanceType = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetCurrentPartitionNum(v int32) *ListDeployOrdersResponseBodyData {
	s.CurrentPartitionNum = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetDeployOrderId(v int64) *ListDeployOrdersResponseBodyData {
	s.DeployOrderId = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetDeployPauseType(v string) *ListDeployOrdersResponseBodyData {
	s.DeployPauseType = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetDeployPauseTypeName(v string) *ListDeployOrdersResponseBodyData {
	s.DeployPauseTypeName = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetDeployType(v string) *ListDeployOrdersResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetDeployTypeName(v string) *ListDeployOrdersResponseBodyData {
	s.DeployTypeName = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetDescription(v string) *ListDeployOrdersResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetElapsedTime(v int32) *ListDeployOrdersResponseBodyData {
	s.ElapsedTime = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetEndTime(v string) *ListDeployOrdersResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetEnvId(v int64) *ListDeployOrdersResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetEnvType(v string) *ListDeployOrdersResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetFailureRate(v string) *ListDeployOrdersResponseBodyData {
	s.FailureRate = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetFinishAppInstanceCt(v int32) *ListDeployOrdersResponseBodyData {
	s.FinishAppInstanceCt = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetName(v string) *ListDeployOrdersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetPartitionType(v string) *ListDeployOrdersResponseBodyData {
	s.PartitionType = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetPartitionTypeName(v string) *ListDeployOrdersResponseBodyData {
	s.PartitionTypeName = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetResult(v int32) *ListDeployOrdersResponseBodyData {
	s.Result = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetResultName(v string) *ListDeployOrdersResponseBodyData {
	s.ResultName = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetSchemaId(v int64) *ListDeployOrdersResponseBodyData {
	s.SchemaId = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetStartTime(v string) *ListDeployOrdersResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetStatus(v int32) *ListDeployOrdersResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetStatusName(v string) *ListDeployOrdersResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetTotalAppInstanceCt(v int32) *ListDeployOrdersResponseBodyData {
	s.TotalAppInstanceCt = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetTotalPartitions(v int32) *ListDeployOrdersResponseBodyData {
	s.TotalPartitions = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetUserId(v string) *ListDeployOrdersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDeployOrdersResponseBodyData) SetUserNick(v string) *ListDeployOrdersResponseBodyData {
	s.UserNick = &v
	return s
}

type ListDeployOrdersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeployOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListDeployOrdersResponse) SetHeaders(v map[string]*string) *ListDeployOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListDeployOrdersResponse) SetStatusCode(v int32) *ListDeployOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeployOrdersResponse) SetBody(v *ListDeployOrdersResponseBody) *ListDeployOrdersResponse {
	s.Body = v
	return s
}

type ListJobHistoriesRequest struct {
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId      *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListJobHistoriesRequest) SetAppId(v int64) *ListJobHistoriesRequest {
	s.AppId = &v
	return s
}

func (s *ListJobHistoriesRequest) SetEnvId(v int64) *ListJobHistoriesRequest {
	s.EnvId = &v
	return s
}

func (s *ListJobHistoriesRequest) SetPageNumber(v int32) *ListJobHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobHistoriesRequest) SetPageSize(v int32) *ListJobHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobHistoriesRequest) SetStatus(v string) *ListJobHistoriesRequest {
	s.Status = &v
	return s
}

type ListJobHistoriesResponseBody struct {
	Code       *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListJobHistoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobHistoriesResponseBody) SetCode(v int32) *ListJobHistoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobHistoriesResponseBody) SetData(v []*ListJobHistoriesResponseBodyData) *ListJobHistoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListJobHistoriesResponseBody) SetErrorMsg(v string) *ListJobHistoriesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListJobHistoriesResponseBody) SetPageNumber(v int32) *ListJobHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobHistoriesResponseBody) SetPageSize(v int32) *ListJobHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobHistoriesResponseBody) SetRequestId(v string) *ListJobHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobHistoriesResponseBody) SetTotalCount(v int64) *ListJobHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobHistoriesResponseBodyData struct {
	ActiveDeadlineSeconds *int32    `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	BackoffLimit          *int32    `json:"BackoffLimit,omitempty" xml:"BackoffLimit,omitempty"`
	CompletionTime        *string   `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	Completions           *int32    `json:"Completions,omitempty" xml:"Completions,omitempty"`
	Failed                *int32    `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Message               *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Name                  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Parallelism           *int32    `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	PodList               []*string `json:"PodList,omitempty" xml:"PodList,omitempty" type:"Repeated"`
	StartTime             *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Succeeded             *int32    `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s ListJobHistoriesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobHistoriesResponseBodyData) SetActiveDeadlineSeconds(v int32) *ListJobHistoriesResponseBodyData {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetBackoffLimit(v int32) *ListJobHistoriesResponseBodyData {
	s.BackoffLimit = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetCompletionTime(v string) *ListJobHistoriesResponseBodyData {
	s.CompletionTime = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetCompletions(v int32) *ListJobHistoriesResponseBodyData {
	s.Completions = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetFailed(v int32) *ListJobHistoriesResponseBodyData {
	s.Failed = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetMessage(v string) *ListJobHistoriesResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetName(v string) *ListJobHistoriesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetParallelism(v int32) *ListJobHistoriesResponseBodyData {
	s.Parallelism = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetPodList(v []*string) *ListJobHistoriesResponseBodyData {
	s.PodList = v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetStartTime(v string) *ListJobHistoriesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListJobHistoriesResponseBodyData) SetSucceeded(v int32) *ListJobHistoriesResponseBodyData {
	s.Succeeded = &v
	return s
}

type ListJobHistoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListJobHistoriesResponse) SetHeaders(v map[string]*string) *ListJobHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListJobHistoriesResponse) SetStatusCode(v int32) *ListJobHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobHistoriesResponse) SetBody(v *ListJobHistoriesResponseBody) *ListJobHistoriesResponse {
	s.Body = v
	return s
}

type ListNodeLabelBindingsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNodeLabelBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeLabelBindingsRequest) SetClusterId(v string) *ListNodeLabelBindingsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeLabelBindingsRequest) SetInstanceId(v string) *ListNodeLabelBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNodeLabelBindingsRequest) SetLabelKey(v string) *ListNodeLabelBindingsRequest {
	s.LabelKey = &v
	return s
}

func (s *ListNodeLabelBindingsRequest) SetLabelValue(v string) *ListNodeLabelBindingsRequest {
	s.LabelValue = &v
	return s
}

func (s *ListNodeLabelBindingsRequest) SetPageNumber(v int32) *ListNodeLabelBindingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeLabelBindingsRequest) SetPageSize(v int32) *ListNodeLabelBindingsRequest {
	s.PageSize = &v
	return s
}

type ListNodeLabelBindingsResponseBody struct {
	Code       *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListNodeLabelBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeLabelBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeLabelBindingsResponseBody) SetCode(v int32) *ListNodeLabelBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBody) SetData(v []*ListNodeLabelBindingsResponseBodyData) *ListNodeLabelBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeLabelBindingsResponseBody) SetErrorMsg(v string) *ListNodeLabelBindingsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBody) SetPageNumber(v int32) *ListNodeLabelBindingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBody) SetPageSize(v int32) *ListNodeLabelBindingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBody) SetRequestId(v string) *ListNodeLabelBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBody) SetTotalCount(v int64) *ListNodeLabelBindingsResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodeLabelBindingsResponseBodyData struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LabelKey     *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue   *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s ListNodeLabelBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeLabelBindingsResponseBodyData) SetId(v int64) *ListNodeLabelBindingsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBodyData) SetInstanceId(v string) *ListNodeLabelBindingsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBodyData) SetInstanceType(v string) *ListNodeLabelBindingsResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBodyData) SetLabelKey(v string) *ListNodeLabelBindingsResponseBodyData {
	s.LabelKey = &v
	return s
}

func (s *ListNodeLabelBindingsResponseBodyData) SetLabelValue(v string) *ListNodeLabelBindingsResponseBodyData {
	s.LabelValue = &v
	return s
}

type ListNodeLabelBindingsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeLabelBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeLabelBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeLabelBindingsResponse) SetHeaders(v map[string]*string) *ListNodeLabelBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeLabelBindingsResponse) SetStatusCode(v int32) *ListNodeLabelBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeLabelBindingsResponse) SetBody(v *ListNodeLabelBindingsResponseBody) *ListNodeLabelBindingsResponse {
	s.Body = v
	return s
}

type ListNodeLabelsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNodeLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeLabelsRequest) SetClusterId(v string) *ListNodeLabelsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeLabelsRequest) SetLabelKey(v string) *ListNodeLabelsRequest {
	s.LabelKey = &v
	return s
}

func (s *ListNodeLabelsRequest) SetPageNumber(v int32) *ListNodeLabelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeLabelsRequest) SetPageSize(v int32) *ListNodeLabelsRequest {
	s.PageSize = &v
	return s
}

type ListNodeLabelsResponseBody struct {
	Code       *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListNodeLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                           `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeLabelsResponseBody) SetCode(v int32) *ListNodeLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodeLabelsResponseBody) SetData(v []*ListNodeLabelsResponseBodyData) *ListNodeLabelsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeLabelsResponseBody) SetErrorMsg(v string) *ListNodeLabelsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListNodeLabelsResponseBody) SetPageNumber(v int32) *ListNodeLabelsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodeLabelsResponseBody) SetPageSize(v int32) *ListNodeLabelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodeLabelsResponseBody) SetRequestId(v string) *ListNodeLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeLabelsResponseBody) SetTotalCount(v int64) *ListNodeLabelsResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodeLabelsResponseBodyData struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s ListNodeLabelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeLabelsResponseBodyData) SetClusterId(v string) *ListNodeLabelsResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListNodeLabelsResponseBodyData) SetId(v int64) *ListNodeLabelsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListNodeLabelsResponseBodyData) SetLabelKey(v string) *ListNodeLabelsResponseBodyData {
	s.LabelKey = &v
	return s
}

func (s *ListNodeLabelsResponseBodyData) SetLabelValue(v string) *ListNodeLabelsResponseBodyData {
	s.LabelValue = &v
	return s
}

type ListNodeLabelsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeLabelsResponse) SetHeaders(v map[string]*string) *ListNodeLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeLabelsResponse) SetStatusCode(v int32) *ListNodeLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeLabelsResponse) SetBody(v *ListNodeLabelsResponseBody) *ListNodeLabelsResponse {
	s.Body = v
	return s
}

type ListPersistentVolumeRequest struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPersistentVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeRequest) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeRequest) SetClusterInstanceId(v string) *ListPersistentVolumeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *ListPersistentVolumeRequest) SetPageNumber(v int32) *ListPersistentVolumeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersistentVolumeRequest) SetPageSize(v int32) *ListPersistentVolumeRequest {
	s.PageSize = &v
	return s
}

type ListPersistentVolumeResponseBody struct {
	Code       *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListPersistentVolumeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrMsg     *string                                 `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPersistentVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeResponseBody) SetCode(v int32) *ListPersistentVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersistentVolumeResponseBody) SetData(v []*ListPersistentVolumeResponseBodyData) *ListPersistentVolumeResponseBody {
	s.Data = v
	return s
}

func (s *ListPersistentVolumeResponseBody) SetErrMsg(v string) *ListPersistentVolumeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListPersistentVolumeResponseBody) SetPageNumber(v int32) *ListPersistentVolumeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersistentVolumeResponseBody) SetPageSize(v int32) *ListPersistentVolumeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersistentVolumeResponseBody) SetRequestId(v string) *ListPersistentVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersistentVolumeResponseBody) SetTotalCount(v int64) *ListPersistentVolumeResponseBody {
	s.TotalCount = &v
	return s
}

type ListPersistentVolumeResponseBodyData struct {
	AccessModes   *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	Capacity      *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MountDir      *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PvcName       *string `json:"PvcName,omitempty" xml:"PvcName,omitempty"`
	Reason        *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReclaimPolicy *string `json:"ReclaimPolicy,omitempty" xml:"ReclaimPolicy,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageClass  *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s ListPersistentVolumeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeResponseBodyData) SetAccessModes(v string) *ListPersistentVolumeResponseBodyData {
	s.AccessModes = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetCapacity(v string) *ListPersistentVolumeResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetCreateTime(v string) *ListPersistentVolumeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetMountDir(v string) *ListPersistentVolumeResponseBodyData {
	s.MountDir = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetName(v string) *ListPersistentVolumeResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetPvcName(v string) *ListPersistentVolumeResponseBodyData {
	s.PvcName = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetReason(v string) *ListPersistentVolumeResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetReclaimPolicy(v string) *ListPersistentVolumeResponseBodyData {
	s.ReclaimPolicy = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetStatus(v string) *ListPersistentVolumeResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListPersistentVolumeResponseBodyData) SetStorageClass(v string) *ListPersistentVolumeResponseBodyData {
	s.StorageClass = &v
	return s
}

type ListPersistentVolumeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPersistentVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersistentVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeResponse) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeResponse) SetHeaders(v map[string]*string) *ListPersistentVolumeResponse {
	s.Headers = v
	return s
}

func (s *ListPersistentVolumeResponse) SetStatusCode(v int32) *ListPersistentVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPersistentVolumeResponse) SetBody(v *ListPersistentVolumeResponseBody) *ListPersistentVolumeResponse {
	s.Body = v
	return s
}

type ListPersistentVolumeClaimRequest struct {
	AppId      *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId      *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPersistentVolumeClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeClaimRequest) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeClaimRequest) SetAppId(v int64) *ListPersistentVolumeClaimRequest {
	s.AppId = &v
	return s
}

func (s *ListPersistentVolumeClaimRequest) SetEnvId(v int64) *ListPersistentVolumeClaimRequest {
	s.EnvId = &v
	return s
}

func (s *ListPersistentVolumeClaimRequest) SetPageNumber(v int32) *ListPersistentVolumeClaimRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersistentVolumeClaimRequest) SetPageSize(v int32) *ListPersistentVolumeClaimRequest {
	s.PageSize = &v
	return s
}

type ListPersistentVolumeClaimResponseBody struct {
	Code       *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListPersistentVolumeClaimResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPersistentVolumeClaimResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeClaimResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeClaimResponseBody) SetCode(v int32) *ListPersistentVolumeClaimResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBody) SetData(v []*ListPersistentVolumeClaimResponseBodyData) *ListPersistentVolumeClaimResponseBody {
	s.Data = v
	return s
}

func (s *ListPersistentVolumeClaimResponseBody) SetErrorMsg(v string) *ListPersistentVolumeClaimResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBody) SetPageNumber(v int32) *ListPersistentVolumeClaimResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBody) SetPageSize(v int32) *ListPersistentVolumeClaimResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBody) SetRequestId(v string) *ListPersistentVolumeClaimResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBody) SetTotalCount(v int64) *ListPersistentVolumeClaimResponseBody {
	s.TotalCount = &v
	return s
}

type ListPersistentVolumeClaimResponseBodyData struct {
	AccessModes  *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	Capacity     *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	VolumeName   *string `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
}

func (s ListPersistentVolumeClaimResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeClaimResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetAccessModes(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.AccessModes = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetCapacity(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetCreateTime(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetName(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetStatus(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetStorageClass(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.StorageClass = &v
	return s
}

func (s *ListPersistentVolumeClaimResponseBodyData) SetVolumeName(v string) *ListPersistentVolumeClaimResponseBodyData {
	s.VolumeName = &v
	return s
}

type ListPersistentVolumeClaimResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPersistentVolumeClaimResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersistentVolumeClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersistentVolumeClaimResponse) GoString() string {
	return s.String()
}

func (s *ListPersistentVolumeClaimResponse) SetHeaders(v map[string]*string) *ListPersistentVolumeClaimResponse {
	s.Headers = v
	return s
}

func (s *ListPersistentVolumeClaimResponse) SetStatusCode(v int32) *ListPersistentVolumeClaimResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPersistentVolumeClaimResponse) SetBody(v *ListPersistentVolumeClaimResponseBody) *ListPersistentVolumeClaimResponse {
	s.Body = v
	return s
}

type ListPodsRequest struct {
	DeployOrderId *int64   `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	PageNumber    *int32   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResultList    []*int32 `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	StatusList    []*int32 `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListPodsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPodsRequest) GoString() string {
	return s.String()
}

func (s *ListPodsRequest) SetDeployOrderId(v int64) *ListPodsRequest {
	s.DeployOrderId = &v
	return s
}

func (s *ListPodsRequest) SetPageNumber(v int32) *ListPodsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPodsRequest) SetPageSize(v int32) *ListPodsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPodsRequest) SetResultList(v []*int32) *ListPodsRequest {
	s.ResultList = v
	return s
}

func (s *ListPodsRequest) SetStatusList(v []*int32) *ListPodsRequest {
	s.StatusList = v
	return s
}

type ListPodsResponseBody struct {
	Code       *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListPodsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPodsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPodsResponseBody) SetCode(v int32) *ListPodsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPodsResponseBody) SetData(v []*ListPodsResponseBodyData) *ListPodsResponseBody {
	s.Data = v
	return s
}

func (s *ListPodsResponseBody) SetErrorMsg(v string) *ListPodsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListPodsResponseBody) SetPageNumber(v int32) *ListPodsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPodsResponseBody) SetPageSize(v int32) *ListPodsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPodsResponseBody) SetRequestId(v string) *ListPodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPodsResponseBody) SetTotalCount(v int64) *ListPodsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPodsResponseBodyData struct {
	AppInstanceId      *string                                `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	DeployOrderId      *int64                                 `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	DeployPartitionNum *int32                                 `json:"DeployPartitionNum,omitempty" xml:"DeployPartitionNum,omitempty"`
	DeploySteps        []*ListPodsResponseBodyDataDeploySteps `json:"DeploySteps,omitempty" xml:"DeploySteps,omitempty" type:"Repeated"`
	GroupName          *string                                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HostIp             *string                                `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	HostName           *string                                `json:"HostName,omitempty" xml:"HostName,omitempty"`
	PodIp              *string                                `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	Region             *string                                `json:"Region,omitempty" xml:"Region,omitempty"`
	Result             *int32                                 `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultName         *string                                `json:"ResultName,omitempty" xml:"ResultName,omitempty"`
	StartTime          *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *int32                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName         *string                                `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	UpdateTime         *string                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPodsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPodsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPodsResponseBodyData) SetAppInstanceId(v string) *ListPodsResponseBodyData {
	s.AppInstanceId = &v
	return s
}

func (s *ListPodsResponseBodyData) SetDeployOrderId(v int64) *ListPodsResponseBodyData {
	s.DeployOrderId = &v
	return s
}

func (s *ListPodsResponseBodyData) SetDeployPartitionNum(v int32) *ListPodsResponseBodyData {
	s.DeployPartitionNum = &v
	return s
}

func (s *ListPodsResponseBodyData) SetDeploySteps(v []*ListPodsResponseBodyDataDeploySteps) *ListPodsResponseBodyData {
	s.DeploySteps = v
	return s
}

func (s *ListPodsResponseBodyData) SetGroupName(v string) *ListPodsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListPodsResponseBodyData) SetHostIp(v string) *ListPodsResponseBodyData {
	s.HostIp = &v
	return s
}

func (s *ListPodsResponseBodyData) SetHostName(v string) *ListPodsResponseBodyData {
	s.HostName = &v
	return s
}

func (s *ListPodsResponseBodyData) SetPodIp(v string) *ListPodsResponseBodyData {
	s.PodIp = &v
	return s
}

func (s *ListPodsResponseBodyData) SetRegion(v string) *ListPodsResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListPodsResponseBodyData) SetResult(v int32) *ListPodsResponseBodyData {
	s.Result = &v
	return s
}

func (s *ListPodsResponseBodyData) SetResultName(v string) *ListPodsResponseBodyData {
	s.ResultName = &v
	return s
}

func (s *ListPodsResponseBodyData) SetStartTime(v string) *ListPodsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListPodsResponseBodyData) SetStatus(v int32) *ListPodsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListPodsResponseBodyData) SetStatusName(v string) *ListPodsResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *ListPodsResponseBodyData) SetUpdateTime(v string) *ListPodsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListPodsResponseBodyDataDeploySteps struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StepCode *string `json:"StepCode,omitempty" xml:"StepCode,omitempty"`
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s ListPodsResponseBodyDataDeploySteps) String() string {
	return tea.Prettify(s)
}

func (s ListPodsResponseBodyDataDeploySteps) GoString() string {
	return s.String()
}

func (s *ListPodsResponseBodyDataDeploySteps) SetStatus(v string) *ListPodsResponseBodyDataDeploySteps {
	s.Status = &v
	return s
}

func (s *ListPodsResponseBodyDataDeploySteps) SetStepCode(v string) *ListPodsResponseBodyDataDeploySteps {
	s.StepCode = &v
	return s
}

func (s *ListPodsResponseBodyDataDeploySteps) SetStepName(v string) *ListPodsResponseBodyDataDeploySteps {
	s.StepName = &v
	return s
}

type ListPodsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPodsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPodsResponse) GoString() string {
	return s.String()
}

func (s *ListPodsResponse) SetHeaders(v map[string]*string) *ListPodsResponse {
	s.Headers = v
	return s
}

func (s *ListPodsResponse) SetStatusCode(v int32) *ListPodsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPodsResponse) SetBody(v *ListPodsResponseBody) *ListPodsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId       *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetAppId(v int64) *ListServicesRequest {
	s.AppId = &v
	return s
}

func (s *ListServicesRequest) SetEnvId(v int64) *ListServicesRequest {
	s.EnvId = &v
	return s
}

func (s *ListServicesRequest) SetName(v string) *ListServicesRequest {
	s.Name = &v
	return s
}

func (s *ListServicesRequest) SetPageNumber(v int32) *ListServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int32) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetServiceType(v string) *ListServicesRequest {
	s.ServiceType = &v
	return s
}

type ListServicesResponseBody struct {
	ClusterIP  *string                         `json:"ClusterIP,omitempty" xml:"ClusterIP,omitempty"`
	Code       *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetClusterIP(v string) *ListServicesResponseBody {
	s.ClusterIP = &v
	return s
}

func (s *ListServicesResponseBody) SetCode(v int32) *ListServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListServicesResponseBody) SetData(v []*ListServicesResponseBodyData) *ListServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListServicesResponseBody) SetErrorMsg(v string) *ListServicesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListServicesResponseBody) SetPageNumber(v int32) *ListServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServicesResponseBody) SetPageSize(v int32) *ListServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int64) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponseBodyData struct {
	AppId        *int64                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClusterIP    *string                                     `json:"ClusterIP,omitempty" xml:"ClusterIP,omitempty"`
	EnvId        *int64                                      `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Headless     *bool                                       `json:"Headless,omitempty" xml:"Headless,omitempty"`
	K8sServiceId *string                                     `json:"K8sServiceId,omitempty" xml:"K8sServiceId,omitempty"`
	Name         *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	PortMappings []*ListServicesResponseBodyDataPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	ServiceId    *int64                                      `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceType  *string                                     `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyData) SetAppId(v int64) *ListServicesResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListServicesResponseBodyData) SetClusterIP(v string) *ListServicesResponseBodyData {
	s.ClusterIP = &v
	return s
}

func (s *ListServicesResponseBodyData) SetEnvId(v int64) *ListServicesResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *ListServicesResponseBodyData) SetHeadless(v bool) *ListServicesResponseBodyData {
	s.Headless = &v
	return s
}

func (s *ListServicesResponseBodyData) SetK8sServiceId(v string) *ListServicesResponseBodyData {
	s.K8sServiceId = &v
	return s
}

func (s *ListServicesResponseBodyData) SetName(v string) *ListServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyData) SetPortMappings(v []*ListServicesResponseBodyDataPortMappings) *ListServicesResponseBodyData {
	s.PortMappings = v
	return s
}

func (s *ListServicesResponseBodyData) SetServiceId(v int64) *ListServicesResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyData) SetServiceType(v string) *ListServicesResponseBodyData {
	s.ServiceType = &v
	return s
}

type ListServicesResponseBodyDataPortMappings struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePort   *int32  `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s ListServicesResponseBodyDataPortMappings) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyDataPortMappings) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyDataPortMappings) SetName(v string) *ListServicesResponseBodyDataPortMappings {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyDataPortMappings) SetNodePort(v int32) *ListServicesResponseBodyDataPortMappings {
	s.NodePort = &v
	return s
}

func (s *ListServicesResponseBodyDataPortMappings) SetPort(v int32) *ListServicesResponseBodyDataPortMappings {
	s.Port = &v
	return s
}

func (s *ListServicesResponseBodyDataPortMappings) SetProtocol(v string) *ListServicesResponseBodyDataPortMappings {
	s.Protocol = &v
	return s
}

func (s *ListServicesResponseBodyDataPortMappings) SetTargetPort(v string) *ListServicesResponseBodyDataPortMappings {
	s.TargetPort = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListSlbAPsRequest struct {
	AppId        *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId        *int64    `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkMode  *string   `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProtocolList []*string `json:"ProtocolList,omitempty" xml:"ProtocolList,omitempty" type:"Repeated"`
	SlbId        *string   `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
}

func (s ListSlbAPsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSlbAPsRequest) GoString() string {
	return s.String()
}

func (s *ListSlbAPsRequest) SetAppId(v int64) *ListSlbAPsRequest {
	s.AppId = &v
	return s
}

func (s *ListSlbAPsRequest) SetEnvId(v int64) *ListSlbAPsRequest {
	s.EnvId = &v
	return s
}

func (s *ListSlbAPsRequest) SetName(v string) *ListSlbAPsRequest {
	s.Name = &v
	return s
}

func (s *ListSlbAPsRequest) SetNetworkMode(v string) *ListSlbAPsRequest {
	s.NetworkMode = &v
	return s
}

func (s *ListSlbAPsRequest) SetPageNumber(v int32) *ListSlbAPsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSlbAPsRequest) SetPageSize(v int32) *ListSlbAPsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSlbAPsRequest) SetProtocolList(v []*string) *ListSlbAPsRequest {
	s.ProtocolList = v
	return s
}

func (s *ListSlbAPsRequest) SetSlbId(v string) *ListSlbAPsRequest {
	s.SlbId = &v
	return s
}

type ListSlbAPsResponseBody struct {
	Code       *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListSlbAPsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSlbAPsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlbAPsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlbAPsResponseBody) SetCode(v int32) *ListSlbAPsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSlbAPsResponseBody) SetData(v []*ListSlbAPsResponseBodyData) *ListSlbAPsResponseBody {
	s.Data = v
	return s
}

func (s *ListSlbAPsResponseBody) SetErrorMsg(v string) *ListSlbAPsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListSlbAPsResponseBody) SetPageNumber(v int32) *ListSlbAPsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSlbAPsResponseBody) SetPageSize(v int32) *ListSlbAPsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSlbAPsResponseBody) SetRequestId(v string) *ListSlbAPsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlbAPsResponseBody) SetTotalCount(v int64) *ListSlbAPsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSlbAPsResponseBodyData struct {
	AppId              *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CookieTimeout      *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	EnvId              *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EstablishedTimeout *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkMode        *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	Protocol           *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServerPort     *int32  `json:"RealServerPort,omitempty" xml:"RealServerPort,omitempty"`
	SlbAPId            *int64  `json:"SlbAPId,omitempty" xml:"SlbAPId,omitempty"`
	SlbId              *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbIp              *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SslCertId          *string `json:"SslCertId,omitempty" xml:"SslCertId,omitempty"`
	StickySession      *int32  `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
}

func (s ListSlbAPsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSlbAPsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSlbAPsResponseBodyData) SetAppId(v int64) *ListSlbAPsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetCookieTimeout(v int32) *ListSlbAPsResponseBodyData {
	s.CookieTimeout = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetEnvId(v int64) *ListSlbAPsResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetEstablishedTimeout(v int32) *ListSlbAPsResponseBodyData {
	s.EstablishedTimeout = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetListenerPort(v int32) *ListSlbAPsResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetName(v string) *ListSlbAPsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetNetworkMode(v string) *ListSlbAPsResponseBodyData {
	s.NetworkMode = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetProtocol(v string) *ListSlbAPsResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetRealServerPort(v int32) *ListSlbAPsResponseBodyData {
	s.RealServerPort = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetSlbAPId(v int64) *ListSlbAPsResponseBodyData {
	s.SlbAPId = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetSlbId(v string) *ListSlbAPsResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetSlbIp(v string) *ListSlbAPsResponseBodyData {
	s.SlbIp = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetSslCertId(v string) *ListSlbAPsResponseBodyData {
	s.SslCertId = &v
	return s
}

func (s *ListSlbAPsResponseBodyData) SetStickySession(v int32) *ListSlbAPsResponseBodyData {
	s.StickySession = &v
	return s
}

type ListSlbAPsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSlbAPsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlbAPsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlbAPsResponse) GoString() string {
	return s.String()
}

func (s *ListSlbAPsResponse) SetHeaders(v map[string]*string) *ListSlbAPsResponse {
	s.Headers = v
	return s
}

func (s *ListSlbAPsResponse) SetStatusCode(v int32) *ListSlbAPsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlbAPsResponse) SetBody(v *ListSlbAPsResponseBody) *ListSlbAPsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
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
	Code       *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg   *string                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	PageNumber *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetCode(v int32) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v []*ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetErrorMsg(v string) *ListUsersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

type ListUsersResponseBodyData struct {
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetRealName(v string) *ListUsersResponseBodyData {
	s.RealName = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserId(v string) *ListUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserType(v string) *ListUsersResponseBodyData {
	s.UserType = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ModifyServiceRequest struct {
	Name         *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PortMappings []*ModifyServiceRequestPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	ServiceId    *int64                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ModifyServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceRequest) SetName(v string) *ModifyServiceRequest {
	s.Name = &v
	return s
}

func (s *ModifyServiceRequest) SetPortMappings(v []*ModifyServiceRequestPortMappings) *ModifyServiceRequest {
	s.PortMappings = v
	return s
}

func (s *ModifyServiceRequest) SetServiceId(v int64) *ModifyServiceRequest {
	s.ServiceId = &v
	return s
}

type ModifyServiceRequestPortMappings struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePort   *int32  `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s ModifyServiceRequestPortMappings) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceRequestPortMappings) GoString() string {
	return s.String()
}

func (s *ModifyServiceRequestPortMappings) SetName(v string) *ModifyServiceRequestPortMappings {
	s.Name = &v
	return s
}

func (s *ModifyServiceRequestPortMappings) SetNodePort(v int32) *ModifyServiceRequestPortMappings {
	s.NodePort = &v
	return s
}

func (s *ModifyServiceRequestPortMappings) SetPort(v int32) *ModifyServiceRequestPortMappings {
	s.Port = &v
	return s
}

func (s *ModifyServiceRequestPortMappings) SetProtocol(v string) *ModifyServiceRequestPortMappings {
	s.Protocol = &v
	return s
}

func (s *ModifyServiceRequestPortMappings) SetTargetPort(v string) *ModifyServiceRequestPortMappings {
	s.TargetPort = &v
	return s
}

type ModifyServiceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceResponseBody) SetCode(v int32) *ModifyServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyServiceResponseBody) SetErrMsg(v string) *ModifyServiceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ModifyServiceResponseBody) SetRequestId(v string) *ModifyServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyServiceResponseBody) SetSuccess(v bool) *ModifyServiceResponseBody {
	s.Success = &v
	return s
}

type ModifyServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceResponse) SetHeaders(v map[string]*string) *ModifyServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceResponse) SetStatusCode(v int32) *ModifyServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceResponse) SetBody(v *ModifyServiceResponseBody) *ModifyServiceResponse {
	s.Body = v
	return s
}

type ModifySlbAPRequest struct {
	CookieTimeout      *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	EstablishedTimeout *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RealServerPort     *int32  `json:"RealServerPort,omitempty" xml:"RealServerPort,omitempty"`
	SlbAPId            *int64  `json:"SlbAPId,omitempty" xml:"SlbAPId,omitempty"`
	SslCertId          *string `json:"SslCertId,omitempty" xml:"SslCertId,omitempty"`
	StickySession      *int32  `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
}

func (s ModifySlbAPRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySlbAPRequest) GoString() string {
	return s.String()
}

func (s *ModifySlbAPRequest) SetCookieTimeout(v int32) *ModifySlbAPRequest {
	s.CookieTimeout = &v
	return s
}

func (s *ModifySlbAPRequest) SetEstablishedTimeout(v int32) *ModifySlbAPRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *ModifySlbAPRequest) SetName(v string) *ModifySlbAPRequest {
	s.Name = &v
	return s
}

func (s *ModifySlbAPRequest) SetRealServerPort(v int32) *ModifySlbAPRequest {
	s.RealServerPort = &v
	return s
}

func (s *ModifySlbAPRequest) SetSlbAPId(v int64) *ModifySlbAPRequest {
	s.SlbAPId = &v
	return s
}

func (s *ModifySlbAPRequest) SetSslCertId(v string) *ModifySlbAPRequest {
	s.SslCertId = &v
	return s
}

func (s *ModifySlbAPRequest) SetStickySession(v int32) *ModifySlbAPRequest {
	s.StickySession = &v
	return s
}

type ModifySlbAPResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySlbAPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySlbAPResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySlbAPResponseBody) SetCode(v int32) *ModifySlbAPResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySlbAPResponseBody) SetErrMsg(v string) *ModifySlbAPResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ModifySlbAPResponseBody) SetRequestId(v string) *ModifySlbAPResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySlbAPResponseBody) SetSuccess(v bool) *ModifySlbAPResponseBody {
	s.Success = &v
	return s
}

type ModifySlbAPResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySlbAPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySlbAPResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySlbAPResponse) GoString() string {
	return s.String()
}

func (s *ModifySlbAPResponse) SetHeaders(v map[string]*string) *ModifySlbAPResponse {
	s.Headers = v
	return s
}

func (s *ModifySlbAPResponse) SetStatusCode(v int32) *ModifySlbAPResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySlbAPResponse) SetBody(v *ModifySlbAPResponseBody) *ModifySlbAPResponse {
	s.Body = v
	return s
}

type OfflineAppEnvironmentRequest struct {
	AppId     *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeletePvc *bool  `json:"DeletePvc,omitempty" xml:"DeletePvc,omitempty"`
	EnvId     *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s OfflineAppEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineAppEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *OfflineAppEnvironmentRequest) SetAppId(v int64) *OfflineAppEnvironmentRequest {
	s.AppId = &v
	return s
}

func (s *OfflineAppEnvironmentRequest) SetDeletePvc(v bool) *OfflineAppEnvironmentRequest {
	s.DeletePvc = &v
	return s
}

func (s *OfflineAppEnvironmentRequest) SetEnvId(v int64) *OfflineAppEnvironmentRequest {
	s.EnvId = &v
	return s
}

type OfflineAppEnvironmentResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                  `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *OfflineAppEnvironmentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineAppEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineAppEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineAppEnvironmentResponseBody) SetCode(v int32) *OfflineAppEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineAppEnvironmentResponseBody) SetErrMsg(v string) *OfflineAppEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *OfflineAppEnvironmentResponseBody) SetRequestId(v string) *OfflineAppEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineAppEnvironmentResponseBody) SetResult(v *OfflineAppEnvironmentResponseBodyResult) *OfflineAppEnvironmentResponseBody {
	s.Result = v
	return s
}

func (s *OfflineAppEnvironmentResponseBody) SetSuccess(v bool) *OfflineAppEnvironmentResponseBody {
	s.Success = &v
	return s
}

type OfflineAppEnvironmentResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineAppEnvironmentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OfflineAppEnvironmentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OfflineAppEnvironmentResponseBodyResult) SetSuccess(v bool) *OfflineAppEnvironmentResponseBodyResult {
	s.Success = &v
	return s
}

type OfflineAppEnvironmentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OfflineAppEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OfflineAppEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineAppEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *OfflineAppEnvironmentResponse) SetHeaders(v map[string]*string) *OfflineAppEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *OfflineAppEnvironmentResponse) SetStatusCode(v int32) *OfflineAppEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineAppEnvironmentResponse) SetBody(v *OfflineAppEnvironmentResponseBody) *OfflineAppEnvironmentResponse {
	s.Body = v
	return s
}

type QueryClusterDetailRequest struct {
	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
}

func (s QueryClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailRequest) SetClusterInstanceId(v string) *QueryClusterDetailRequest {
	s.ClusterInstanceId = &v
	return s
}

type QueryClusterDetailResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                               `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryClusterDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBody) SetCode(v int32) *QueryClusterDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetErrMsg(v string) *QueryClusterDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetRequestId(v string) *QueryClusterDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetResult(v *QueryClusterDetailResponseBodyResult) *QueryClusterDetailResponseBody {
	s.Result = v
	return s
}

func (s *QueryClusterDetailResponseBody) SetSuccess(v bool) *QueryClusterDetailResponseBody {
	s.Success = &v
	return s
}

type QueryClusterDetailResponseBodyResult struct {
	BasicInfo        *QueryClusterDetailResponseBodyResultBasicInfo    `json:"BasicInfo,omitempty" xml:"BasicInfo,omitempty" type:"Struct"`
	InstanceInfo     *QueryClusterDetailResponseBodyResultInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	NetInfo          *QueryClusterDetailResponseBodyResultNetInfo      `json:"NetInfo,omitempty" xml:"NetInfo,omitempty" type:"Struct"`
	NodeWorkLoadList []*string                                         `json:"NodeWorkLoadList,omitempty" xml:"NodeWorkLoadList,omitempty" type:"Repeated"`
	WorkLoad         *QueryClusterDetailResponseBodyResultWorkLoad     `json:"WorkLoad,omitempty" xml:"WorkLoad,omitempty" type:"Struct"`
}

func (s QueryClusterDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyResult) SetBasicInfo(v *QueryClusterDetailResponseBodyResultBasicInfo) *QueryClusterDetailResponseBodyResult {
	s.BasicInfo = v
	return s
}

func (s *QueryClusterDetailResponseBodyResult) SetInstanceInfo(v *QueryClusterDetailResponseBodyResultInstanceInfo) *QueryClusterDetailResponseBodyResult {
	s.InstanceInfo = v
	return s
}

func (s *QueryClusterDetailResponseBodyResult) SetNetInfo(v *QueryClusterDetailResponseBodyResultNetInfo) *QueryClusterDetailResponseBodyResult {
	s.NetInfo = v
	return s
}

func (s *QueryClusterDetailResponseBodyResult) SetNodeWorkLoadList(v []*string) *QueryClusterDetailResponseBodyResult {
	s.NodeWorkLoadList = v
	return s
}

func (s *QueryClusterDetailResponseBodyResult) SetWorkLoad(v *QueryClusterDetailResponseBodyResultWorkLoad) *QueryClusterDetailResponseBodyResult {
	s.WorkLoad = v
	return s
}

type QueryClusterDetailResponseBodyResultBasicInfo struct {
	BusinessCode            *string   `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	ClusterId               *int64    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterInstanceId       *string   `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	ClusterName             *string   `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	EcsIds                  []*string `json:"EcsIds,omitempty" xml:"EcsIds,omitempty" type:"Repeated"`
	EnvType                 *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	HasInstallArmsPilot     *bool     `json:"HasInstallArmsPilot,omitempty" xml:"HasInstallArmsPilot,omitempty"`
	HasInstallLogController *bool     `json:"HasInstallLogController,omitempty" xml:"HasInstallLogController,omitempty"`
	InstallArmsInProcess    *bool     `json:"InstallArmsInProcess,omitempty" xml:"InstallArmsInProcess,omitempty"`
	InstallLogInProcess     *bool     `json:"InstallLogInProcess,omitempty" xml:"InstallLogInProcess,omitempty"`
	MainUserId              *string   `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName              *string   `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	UserNick                *string   `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	VpcId                   *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Vswitchs                []*string `json:"Vswitchs,omitempty" xml:"Vswitchs,omitempty" type:"Repeated"`
}

func (s QueryClusterDetailResponseBodyResultBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyResultBasicInfo) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetBusinessCode(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.BusinessCode = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetClusterId(v int64) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.ClusterId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetClusterInstanceId(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.ClusterInstanceId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetClusterName(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.ClusterName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetEcsIds(v []*string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.EcsIds = v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetEnvType(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.EnvType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetHasInstallArmsPilot(v bool) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.HasInstallArmsPilot = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetHasInstallLogController(v bool) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.HasInstallLogController = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetInstallArmsInProcess(v bool) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.InstallArmsInProcess = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetInstallLogInProcess(v bool) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.InstallLogInProcess = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetMainUserId(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.MainUserId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetRegionId(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.RegionId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetRegionName(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.RegionName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetUserNick(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.UserNick = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetVpcId(v string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.VpcId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultBasicInfo) SetVswitchs(v []*string) *QueryClusterDetailResponseBodyResultBasicInfo {
	s.Vswitchs = v
	return s
}

type QueryClusterDetailResponseBodyResultInstanceInfo struct {
	AllocatePodCount     *int32    `json:"AllocatePodCount,omitempty" xml:"AllocatePodCount,omitempty"`
	AllocatedPodInfoList []*string `json:"AllocatedPodInfoList,omitempty" xml:"AllocatedPodInfoList,omitempty" type:"Repeated"`
	AppCount             *int32    `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	AvailablePodInfoList []*string `json:"AvailablePodInfoList,omitempty" xml:"AvailablePodInfoList,omitempty" type:"Repeated"`
}

func (s QueryClusterDetailResponseBodyResultInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyResultInstanceInfo) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyResultInstanceInfo) SetAllocatePodCount(v int32) *QueryClusterDetailResponseBodyResultInstanceInfo {
	s.AllocatePodCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultInstanceInfo) SetAllocatedPodInfoList(v []*string) *QueryClusterDetailResponseBodyResultInstanceInfo {
	s.AllocatedPodInfoList = v
	return s
}

func (s *QueryClusterDetailResponseBodyResultInstanceInfo) SetAppCount(v int32) *QueryClusterDetailResponseBodyResultInstanceInfo {
	s.AppCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultInstanceInfo) SetAvailablePodInfoList(v []*string) *QueryClusterDetailResponseBodyResultInstanceInfo {
	s.AvailablePodInfoList = v
	return s
}

type QueryClusterDetailResponseBodyResultNetInfo struct {
	NetPlugType *string `json:"NetPlugType,omitempty" xml:"NetPlugType,omitempty"`
	ProdCIDR    *string `json:"ProdCIDR,omitempty" xml:"ProdCIDR,omitempty"`
	ServiceCIDR *string `json:"ServiceCIDR,omitempty" xml:"ServiceCIDR,omitempty"`
}

func (s QueryClusterDetailResponseBodyResultNetInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyResultNetInfo) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyResultNetInfo) SetNetPlugType(v string) *QueryClusterDetailResponseBodyResultNetInfo {
	s.NetPlugType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultNetInfo) SetProdCIDR(v string) *QueryClusterDetailResponseBodyResultNetInfo {
	s.ProdCIDR = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultNetInfo) SetServiceCIDR(v string) *QueryClusterDetailResponseBodyResultNetInfo {
	s.ServiceCIDR = &v
	return s
}

type QueryClusterDetailResponseBodyResultWorkLoad struct {
	AllNodeCount        *int32  `json:"AllNodeCount,omitempty" xml:"AllNodeCount,omitempty"`
	AllocateAllPodCount *int32  `json:"AllocateAllPodCount,omitempty" xml:"AllocateAllPodCount,omitempty"`
	AllocateAppPodCount *int32  `json:"AllocateAppPodCount,omitempty" xml:"AllocateAppPodCount,omitempty"`
	CpuCapacityTotal    *string `json:"CpuCapacityTotal,omitempty" xml:"CpuCapacityTotal,omitempty"`
	CpuLevel            *string `json:"CpuLevel,omitempty" xml:"CpuLevel,omitempty"`
	CpuRequest          *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	CpuRequestPercent   *string `json:"CpuRequestPercent,omitempty" xml:"CpuRequestPercent,omitempty"`
	CpuTotal            *string `json:"CpuTotal,omitempty" xml:"CpuTotal,omitempty"`
	CpuUse              *string `json:"CpuUse,omitempty" xml:"CpuUse,omitempty"`
	CpuUsePercent       *string `json:"CpuUsePercent,omitempty" xml:"CpuUsePercent,omitempty"`
	MemCapacityTotal    *string `json:"MemCapacityTotal,omitempty" xml:"MemCapacityTotal,omitempty"`
	MemLevel            *string `json:"MemLevel,omitempty" xml:"MemLevel,omitempty"`
	MemRequest          *string `json:"MemRequest,omitempty" xml:"MemRequest,omitempty"`
	MemRequestPercent   *string `json:"MemRequestPercent,omitempty" xml:"MemRequestPercent,omitempty"`
	MemTotal            *string `json:"MemTotal,omitempty" xml:"MemTotal,omitempty"`
	MemUse              *string `json:"MemUse,omitempty" xml:"MemUse,omitempty"`
	MemUsePercent       *string `json:"MemUsePercent,omitempty" xml:"MemUsePercent,omitempty"`
}

func (s QueryClusterDetailResponseBodyResultWorkLoad) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyResultWorkLoad) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetAllNodeCount(v int32) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.AllNodeCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetAllocateAllPodCount(v int32) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.AllocateAllPodCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetAllocateAppPodCount(v int32) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.AllocateAppPodCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuCapacityTotal(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuCapacityTotal = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuLevel(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuLevel = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuRequest(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuRequest = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuRequestPercent(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuRequestPercent = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuTotal(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuTotal = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuUse(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuUse = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetCpuUsePercent(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.CpuUsePercent = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemCapacityTotal(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemCapacityTotal = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemLevel(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemLevel = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemRequest(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemRequest = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemRequestPercent(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemRequestPercent = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemTotal(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemTotal = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemUse(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemUse = &v
	return s
}

func (s *QueryClusterDetailResponseBodyResultWorkLoad) SetMemUsePercent(v string) *QueryClusterDetailResponseBodyResultWorkLoad {
	s.MemUsePercent = &v
	return s
}

type QueryClusterDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponse) SetHeaders(v map[string]*string) *QueryClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterDetailResponse) SetStatusCode(v int32) *QueryClusterDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClusterDetailResponse) SetBody(v *QueryClusterDetailResponseBody) *QueryClusterDetailResponse {
	s.Body = v
	return s
}

type RebuildAppInstanceRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	EnvId         *int64  `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s RebuildAppInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebuildAppInstanceRequest) SetAppId(v int64) *RebuildAppInstanceRequest {
	s.AppId = &v
	return s
}

func (s *RebuildAppInstanceRequest) SetAppInstanceId(v string) *RebuildAppInstanceRequest {
	s.AppInstanceId = &v
	return s
}

func (s *RebuildAppInstanceRequest) SetEnvId(v int64) *RebuildAppInstanceRequest {
	s.EnvId = &v
	return s
}

type RebuildAppInstanceResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                               `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *RebuildAppInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s RebuildAppInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildAppInstanceResponseBody) SetCode(v int32) *RebuildAppInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RebuildAppInstanceResponseBody) SetErrMsg(v string) *RebuildAppInstanceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *RebuildAppInstanceResponseBody) SetRequestId(v string) *RebuildAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildAppInstanceResponseBody) SetResult(v *RebuildAppInstanceResponseBodyResult) *RebuildAppInstanceResponseBody {
	s.Result = v
	return s
}

type RebuildAppInstanceResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RebuildAppInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RebuildAppInstanceResponseBodyResult) SetSuccess(v bool) *RebuildAppInstanceResponseBodyResult {
	s.Success = &v
	return s
}

type RebuildAppInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebuildAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebuildAppInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebuildAppInstanceResponse) SetHeaders(v map[string]*string) *RebuildAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebuildAppInstanceResponse) SetStatusCode(v int32) *RebuildAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildAppInstanceResponse) SetBody(v *RebuildAppInstanceResponseBody) *RebuildAppInstanceResponse {
	s.Body = v
	return s
}

type RemoveClusterNodeRequest struct {
	ClusterInstanceId *string   `json:"ClusterInstanceId,omitempty" xml:"ClusterInstanceId,omitempty"`
	EcsInstanceIdList []*string `json:"EcsInstanceIdList,omitempty" xml:"EcsInstanceIdList,omitempty" type:"Repeated"`
}

func (s RemoveClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodeRequest) SetClusterInstanceId(v string) *RemoveClusterNodeRequest {
	s.ClusterInstanceId = &v
	return s
}

func (s *RemoveClusterNodeRequest) SetEcsInstanceIdList(v []*string) *RemoveClusterNodeRequest {
	s.EcsInstanceIdList = v
	return s
}

type RemoveClusterNodeResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *RemoveClusterNodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveClusterNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodeResponseBody) SetCode(v int32) *RemoveClusterNodeResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveClusterNodeResponseBody) SetErrMsg(v string) *RemoveClusterNodeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *RemoveClusterNodeResponseBody) SetRequestId(v string) *RemoveClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClusterNodeResponseBody) SetResult(v *RemoveClusterNodeResponseBodyResult) *RemoveClusterNodeResponseBody {
	s.Result = v
	return s
}

func (s *RemoveClusterNodeResponseBody) SetSuccess(v bool) *RemoveClusterNodeResponseBody {
	s.Success = &v
	return s
}

type RemoveClusterNodeResponseBodyResult struct {
	Nonsense *int32 `json:"Nonsense,omitempty" xml:"Nonsense,omitempty"`
}

func (s RemoveClusterNodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodeResponseBodyResult) SetNonsense(v int32) *RemoveClusterNodeResponseBodyResult {
	s.Nonsense = &v
	return s
}

type RemoveClusterNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodeResponse) SetHeaders(v map[string]*string) *RemoveClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterNodeResponse) SetStatusCode(v int32) *RemoveClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterNodeResponse) SetBody(v *RemoveClusterNodeResponseBody) *RemoveClusterNodeResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DbInstanceId    *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
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

func (s *ResetAccountPasswordRequest) SetDbInstanceId(v string) *ResetAccountPasswordRequest {
	s.DbInstanceId = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetCode(v int32) *ResetAccountPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetErrMsg(v string) *ResetAccountPasswordResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
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

type ResourceStatusNotifyRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ResourceStatusNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ResourceStatusNotifyRequest) GoString() string {
	return s.String()
}

func (s *ResourceStatusNotifyRequest) SetData(v string) *ResourceStatusNotifyRequest {
	s.Data = &v
	return s
}

type ResourceStatusNotifyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ResourceStatusNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ResourceStatusNotifyResponse) GoString() string {
	return s.String()
}

func (s *ResourceStatusNotifyResponse) SetHeaders(v map[string]*string) *ResourceStatusNotifyResponse {
	s.Headers = v
	return s
}

func (s *ResourceStatusNotifyResponse) SetStatusCode(v int32) *ResourceStatusNotifyResponse {
	s.StatusCode = &v
	return s
}

type RestartAppInstanceRequest struct {
	AppId             *int64   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceIdList []*int64 `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	EnvId             *int64   `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s RestartAppInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartAppInstanceRequest) SetAppId(v int64) *RestartAppInstanceRequest {
	s.AppId = &v
	return s
}

func (s *RestartAppInstanceRequest) SetAppInstanceIdList(v []*int64) *RestartAppInstanceRequest {
	s.AppInstanceIdList = v
	return s
}

func (s *RestartAppInstanceRequest) SetEnvId(v int64) *RestartAppInstanceRequest {
	s.EnvId = &v
	return s
}

type RestartAppInstanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RestartAppInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartAppInstanceResponseBody) SetCode(v int32) *RestartAppInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RestartAppInstanceResponseBody) SetErrMsg(v string) *RestartAppInstanceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *RestartAppInstanceResponseBody) SetRequestId(v string) *RestartAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartAppInstanceResponseBody) SetResult(v string) *RestartAppInstanceResponseBody {
	s.Result = &v
	return s
}

type RestartAppInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartAppInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartAppInstanceResponse) SetHeaders(v map[string]*string) *RestartAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartAppInstanceResponse) SetStatusCode(v int32) *RestartAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartAppInstanceResponse) SetBody(v *RestartAppInstanceResponseBody) *RestartAppInstanceResponse {
	s.Body = v
	return s
}

type ResumeDeployRequest struct {
	DeployOrderId *int64 `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s ResumeDeployRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeDeployRequest) GoString() string {
	return s.String()
}

func (s *ResumeDeployRequest) SetDeployOrderId(v int64) *ResumeDeployRequest {
	s.DeployOrderId = &v
	return s
}

type ResumeDeployResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeDeployResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeDeployResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeDeployResponseBody) SetCode(v int32) *ResumeDeployResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeDeployResponseBody) SetErrMsg(v string) *ResumeDeployResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ResumeDeployResponseBody) SetRequestId(v string) *ResumeDeployResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeDeployResponseBody) SetSuccess(v bool) *ResumeDeployResponseBody {
	s.Success = &v
	return s
}

type ResumeDeployResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeDeployResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeDeployResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeDeployResponse) GoString() string {
	return s.String()
}

func (s *ResumeDeployResponse) SetHeaders(v map[string]*string) *ResumeDeployResponse {
	s.Headers = v
	return s
}

func (s *ResumeDeployResponse) SetStatusCode(v int32) *ResumeDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeDeployResponse) SetBody(v *ResumeDeployResponseBody) *ResumeDeployResponse {
	s.Body = v
	return s
}

type ScaleAppRequest struct {
	EnvId           *int64 `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Replicas        *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	TotalPartitions *int32 `json:"TotalPartitions,omitempty" xml:"TotalPartitions,omitempty"`
}

func (s ScaleAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleAppRequest) GoString() string {
	return s.String()
}

func (s *ScaleAppRequest) SetEnvId(v int64) *ScaleAppRequest {
	s.EnvId = &v
	return s
}

func (s *ScaleAppRequest) SetReplicas(v int32) *ScaleAppRequest {
	s.Replicas = &v
	return s
}

func (s *ScaleAppRequest) SetTotalPartitions(v int32) *ScaleAppRequest {
	s.TotalPartitions = &v
	return s
}

type ScaleAppResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                     `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ScaleAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScaleAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleAppResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleAppResponseBody) SetCode(v int32) *ScaleAppResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleAppResponseBody) SetErrMsg(v string) *ScaleAppResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ScaleAppResponseBody) SetRequestId(v string) *ScaleAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleAppResponseBody) SetResult(v *ScaleAppResponseBodyResult) *ScaleAppResponseBody {
	s.Result = v
	return s
}

func (s *ScaleAppResponseBody) SetSuccess(v bool) *ScaleAppResponseBody {
	s.Success = &v
	return s
}

type ScaleAppResponseBodyResult struct {
	Admitted      *bool   `json:"Admitted,omitempty" xml:"Admitted,omitempty"`
	BusinessCode  *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	DeployOrderId *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
}

func (s ScaleAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ScaleAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ScaleAppResponseBodyResult) SetAdmitted(v bool) *ScaleAppResponseBodyResult {
	s.Admitted = &v
	return s
}

func (s *ScaleAppResponseBodyResult) SetBusinessCode(v string) *ScaleAppResponseBodyResult {
	s.BusinessCode = &v
	return s
}

func (s *ScaleAppResponseBodyResult) SetDeployOrderId(v int64) *ScaleAppResponseBodyResult {
	s.DeployOrderId = &v
	return s
}

type ScaleAppResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScaleAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleAppResponse) GoString() string {
	return s.String()
}

func (s *ScaleAppResponse) SetHeaders(v map[string]*string) *ScaleAppResponse {
	s.Headers = v
	return s
}

func (s *ScaleAppResponse) SetStatusCode(v int32) *ScaleAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleAppResponse) SetBody(v *ScaleAppResponseBody) *ScaleAppResponse {
	s.Body = v
	return s
}

type SetDeployPauseTypeRequest struct {
	DeployOrderId   *int64  `json:"DeployOrderId,omitempty" xml:"DeployOrderId,omitempty"`
	DeployPauseType *string `json:"DeployPauseType,omitempty" xml:"DeployPauseType,omitempty"`
}

func (s SetDeployPauseTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeployPauseTypeRequest) GoString() string {
	return s.String()
}

func (s *SetDeployPauseTypeRequest) SetDeployOrderId(v int64) *SetDeployPauseTypeRequest {
	s.DeployOrderId = &v
	return s
}

func (s *SetDeployPauseTypeRequest) SetDeployPauseType(v string) *SetDeployPauseTypeRequest {
	s.DeployPauseType = &v
	return s
}

type SetDeployPauseTypeResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDeployPauseTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeployPauseTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeployPauseTypeResponseBody) SetCode(v int32) *SetDeployPauseTypeResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeployPauseTypeResponseBody) SetErrMsg(v string) *SetDeployPauseTypeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SetDeployPauseTypeResponseBody) SetRequestId(v string) *SetDeployPauseTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeployPauseTypeResponseBody) SetSuccess(v bool) *SetDeployPauseTypeResponseBody {
	s.Success = &v
	return s
}

type SetDeployPauseTypeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDeployPauseTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDeployPauseTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeployPauseTypeResponse) GoString() string {
	return s.String()
}

func (s *SetDeployPauseTypeResponse) SetHeaders(v map[string]*string) *SetDeployPauseTypeResponse {
	s.Headers = v
	return s
}

func (s *SetDeployPauseTypeResponse) SetStatusCode(v int32) *SetDeployPauseTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeployPauseTypeResponse) SetBody(v *SetDeployPauseTypeResponseBody) *SetDeployPauseTypeResponse {
	s.Body = v
	return s
}

type SubmitInfoRequest struct {
	CallerUid   *int64                          `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	EcsDescList []*SubmitInfoRequestEcsDescList `json:"EcsDescList,omitempty" xml:"EcsDescList,omitempty" type:"Repeated"`
	MainUserId  *int64                          `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitInfoRequest) SetCallerUid(v int64) *SubmitInfoRequest {
	s.CallerUid = &v
	return s
}

func (s *SubmitInfoRequest) SetEcsDescList(v []*SubmitInfoRequestEcsDescList) *SubmitInfoRequest {
	s.EcsDescList = v
	return s
}

func (s *SubmitInfoRequest) SetMainUserId(v int64) *SubmitInfoRequest {
	s.MainUserId = &v
	return s
}

func (s *SubmitInfoRequest) SetRequestId(v string) *SubmitInfoRequest {
	s.RequestId = &v
	return s
}

type SubmitInfoRequestEcsDescList struct {
	AppType             *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	BussinessDesc       *string `json:"BussinessDesc,omitempty" xml:"BussinessDesc,omitempty"`
	BussinessType       *string `json:"BussinessType,omitempty" xml:"BussinessType,omitempty"`
	EnvType             *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	MiddleWareDesc      *string `json:"MiddleWareDesc,omitempty" xml:"MiddleWareDesc,omitempty"`
	OtherMiddleWareDesc *string `json:"OtherMiddleWareDesc,omitempty" xml:"OtherMiddleWareDesc,omitempty"`
	ResourceId          *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SubmitInfoRequestEcsDescList) String() string {
	return tea.Prettify(s)
}

func (s SubmitInfoRequestEcsDescList) GoString() string {
	return s.String()
}

func (s *SubmitInfoRequestEcsDescList) SetAppType(v string) *SubmitInfoRequestEcsDescList {
	s.AppType = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetBussinessDesc(v string) *SubmitInfoRequestEcsDescList {
	s.BussinessDesc = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetBussinessType(v string) *SubmitInfoRequestEcsDescList {
	s.BussinessType = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetEnvType(v string) *SubmitInfoRequestEcsDescList {
	s.EnvType = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetMiddleWareDesc(v string) *SubmitInfoRequestEcsDescList {
	s.MiddleWareDesc = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetOtherMiddleWareDesc(v string) *SubmitInfoRequestEcsDescList {
	s.OtherMiddleWareDesc = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetResourceId(v string) *SubmitInfoRequestEcsDescList {
	s.ResourceId = &v
	return s
}

func (s *SubmitInfoRequestEcsDescList) SetUserId(v string) *SubmitInfoRequestEcsDescList {
	s.UserId = &v
	return s
}

type SubmitInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitInfoResponseBody) SetCode(v int32) *SubmitInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitInfoResponseBody) SetErrMsg(v string) *SubmitInfoResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SubmitInfoResponseBody) SetRequestId(v string) *SubmitInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitInfoResponseBody) SetSuccess(v bool) *SubmitInfoResponseBody {
	s.Success = &v
	return s
}

type SubmitInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitInfoResponse) SetHeaders(v map[string]*string) *SubmitInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitInfoResponse) SetStatusCode(v int32) *SubmitInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitInfoResponse) SetBody(v *SubmitInfoResponseBody) *SubmitInfoResponse {
	s.Body = v
	return s
}

type SyncPodInfoRequest struct {
	PodName     *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SideCarType *string `json:"SideCarType,omitempty" xml:"SideCarType,omitempty"`
	Status      *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SyncPodInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncPodInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncPodInfoRequest) SetPodName(v string) *SyncPodInfoRequest {
	s.PodName = &v
	return s
}

func (s *SyncPodInfoRequest) SetReason(v string) *SyncPodInfoRequest {
	s.Reason = &v
	return s
}

func (s *SyncPodInfoRequest) SetRequestId(v string) *SyncPodInfoRequest {
	s.RequestId = &v
	return s
}

func (s *SyncPodInfoRequest) SetSideCarType(v string) *SyncPodInfoRequest {
	s.SideCarType = &v
	return s
}

func (s *SyncPodInfoRequest) SetStatus(v bool) *SyncPodInfoRequest {
	s.Status = &v
	return s
}

func (s *SyncPodInfoRequest) SetTaskId(v int64) *SyncPodInfoRequest {
	s.TaskId = &v
	return s
}

type SyncPodInfoResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *SyncPodInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SyncPodInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncPodInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncPodInfoResponseBody) SetCode(v int32) *SyncPodInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncPodInfoResponseBody) SetErrMsg(v string) *SyncPodInfoResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SyncPodInfoResponseBody) SetRequestId(v string) *SyncPodInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncPodInfoResponseBody) SetResult(v *SyncPodInfoResponseBodyResult) *SyncPodInfoResponseBody {
	s.Result = v
	return s
}

type SyncPodInfoResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncPodInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SyncPodInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SyncPodInfoResponseBodyResult) SetSuccess(v bool) *SyncPodInfoResponseBodyResult {
	s.Success = &v
	return s
}

type SyncPodInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncPodInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncPodInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncPodInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncPodInfoResponse) SetHeaders(v map[string]*string) *SyncPodInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncPodInfoResponse) SetStatusCode(v int32) *SyncPodInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncPodInfoResponse) SetBody(v *SyncPodInfoResponseBody) *SyncPodInfoResponse {
	s.Body = v
	return s
}

type UnbindGroupRequest struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UnbindGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindGroupRequest) GoString() string {
	return s.String()
}

func (s *UnbindGroupRequest) SetAppId(v int64) *UnbindGroupRequest {
	s.AppId = &v
	return s
}

func (s *UnbindGroupRequest) SetBizCode(v string) *UnbindGroupRequest {
	s.BizCode = &v
	return s
}

func (s *UnbindGroupRequest) SetName(v string) *UnbindGroupRequest {
	s.Name = &v
	return s
}

type UnbindGroupResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                        `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UnbindGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UnbindGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindGroupResponseBody) SetCode(v int32) *UnbindGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindGroupResponseBody) SetErrMsg(v string) *UnbindGroupResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UnbindGroupResponseBody) SetRequestId(v string) *UnbindGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindGroupResponseBody) SetResult(v *UnbindGroupResponseBodyResult) *UnbindGroupResponseBody {
	s.Result = v
	return s
}

type UnbindGroupResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UnbindGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UnbindGroupResponseBodyResult) SetSuccess(v bool) *UnbindGroupResponseBodyResult {
	s.Success = &v
	return s
}

type UnbindGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindGroupResponse) GoString() string {
	return s.String()
}

func (s *UnbindGroupResponse) SetHeaders(v map[string]*string) *UnbindGroupResponse {
	s.Headers = v
	return s
}

func (s *UnbindGroupResponse) SetStatusCode(v int32) *UnbindGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindGroupResponse) SetBody(v *UnbindGroupResponseBody) *UnbindGroupResponse {
	s.Body = v
	return s
}

type UnbindNodeLabelRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s UnbindNodeLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindNodeLabelRequest) GoString() string {
	return s.String()
}

func (s *UnbindNodeLabelRequest) SetClusterId(v string) *UnbindNodeLabelRequest {
	s.ClusterId = &v
	return s
}

func (s *UnbindNodeLabelRequest) SetInstanceId(v string) *UnbindNodeLabelRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindNodeLabelRequest) SetLabelKey(v string) *UnbindNodeLabelRequest {
	s.LabelKey = &v
	return s
}

func (s *UnbindNodeLabelRequest) SetLabelValue(v string) *UnbindNodeLabelRequest {
	s.LabelValue = &v
	return s
}

type UnbindNodeLabelResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindNodeLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindNodeLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindNodeLabelResponseBody) SetCode(v int32) *UnbindNodeLabelResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindNodeLabelResponseBody) SetErrMsg(v string) *UnbindNodeLabelResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UnbindNodeLabelResponseBody) SetRequestId(v string) *UnbindNodeLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindNodeLabelResponseBody) SetSuccess(v bool) *UnbindNodeLabelResponseBody {
	s.Success = &v
	return s
}

type UnbindNodeLabelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindNodeLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindNodeLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindNodeLabelResponse) GoString() string {
	return s.String()
}

func (s *UnbindNodeLabelResponse) SetHeaders(v map[string]*string) *UnbindNodeLabelResponse {
	s.Headers = v
	return s
}

func (s *UnbindNodeLabelResponse) SetStatusCode(v int32) *UnbindNodeLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindNodeLabelResponse) SetBody(v *UnbindNodeLabelResponseBody) *UnbindNodeLabelResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	AppId            *int64                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizTitle         *string                      `json:"BizTitle,omitempty" xml:"BizTitle,omitempty"`
	Description      *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Language         *string                      `json:"Language,omitempty" xml:"Language,omitempty"`
	MiddleWareIdList []*int32                     `json:"MiddleWareIdList,omitempty" xml:"MiddleWareIdList,omitempty" type:"Repeated"`
	OperatingSystem  *string                      `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
	ServiceType      *string                      `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	UserRoles        []*UpdateAppRequestUserRoles `json:"UserRoles,omitempty" xml:"UserRoles,omitempty" type:"Repeated"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetAppId(v int64) *UpdateAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppRequest) SetBizTitle(v string) *UpdateAppRequest {
	s.BizTitle = &v
	return s
}

func (s *UpdateAppRequest) SetDescription(v string) *UpdateAppRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppRequest) SetLanguage(v string) *UpdateAppRequest {
	s.Language = &v
	return s
}

func (s *UpdateAppRequest) SetMiddleWareIdList(v []*int32) *UpdateAppRequest {
	s.MiddleWareIdList = v
	return s
}

func (s *UpdateAppRequest) SetOperatingSystem(v string) *UpdateAppRequest {
	s.OperatingSystem = &v
	return s
}

func (s *UpdateAppRequest) SetServiceType(v string) *UpdateAppRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateAppRequest) SetUserRoles(v []*UpdateAppRequestUserRoles) *UpdateAppRequest {
	s.UserRoles = v
	return s
}

type UpdateAppRequestUserRoles struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s UpdateAppRequestUserRoles) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestUserRoles) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestUserRoles) SetRoleName(v string) *UpdateAppRequestUserRoles {
	s.RoleName = &v
	return s
}

func (s *UpdateAppRequestUserRoles) SetUserId(v string) *UpdateAppRequestUserRoles {
	s.UserId = &v
	return s
}

func (s *UpdateAppRequestUserRoles) SetUserType(v string) *UpdateAppRequestUserRoles {
	s.UserType = &v
	return s
}

type UpdateAppResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                      `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetCode(v int32) *UpdateAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppResponseBody) SetErrMsg(v string) *UpdateAppResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetResult(v *UpdateAppResponseBodyResult) *UpdateAppResponseBody {
	s.Result = v
	return s
}

type UpdateAppResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBodyResult) SetSuccess(v bool) *UpdateAppResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateAppMonitorsRequest struct {
	MainUserId  *int64   `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	MonitorIds  []*int64 `json:"MonitorIds,omitempty" xml:"MonitorIds,omitempty" type:"Repeated"`
	SilenceTime *string  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	TemplateId  *int64   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateAppMonitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppMonitorsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppMonitorsRequest) SetMainUserId(v int64) *UpdateAppMonitorsRequest {
	s.MainUserId = &v
	return s
}

func (s *UpdateAppMonitorsRequest) SetMonitorIds(v []*int64) *UpdateAppMonitorsRequest {
	s.MonitorIds = v
	return s
}

func (s *UpdateAppMonitorsRequest) SetSilenceTime(v string) *UpdateAppMonitorsRequest {
	s.SilenceTime = &v
	return s
}

func (s *UpdateAppMonitorsRequest) SetTemplateId(v int64) *UpdateAppMonitorsRequest {
	s.TemplateId = &v
	return s
}

type UpdateAppMonitorsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppMonitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppMonitorsResponseBody) SetCode(v string) *UpdateAppMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppMonitorsResponseBody) SetMsg(v string) *UpdateAppMonitorsResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateAppMonitorsResponseBody) SetRequestId(v string) *UpdateAppMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppMonitorsResponseBody) SetSuccess(v bool) *UpdateAppMonitorsResponseBody {
	s.Success = &v
	return s
}

type UpdateAppMonitorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppMonitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppMonitorsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppMonitorsResponse) SetHeaders(v map[string]*string) *UpdateAppMonitorsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppMonitorsResponse) SetStatusCode(v int32) *UpdateAppMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppMonitorsResponse) SetBody(v *UpdateAppMonitorsResponseBody) *UpdateAppMonitorsResponse {
	s.Body = v
	return s
}

type UpdateDeployConfigRequest struct {
	AppId         *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CodePath      *string   `json:"CodePath,omitempty" xml:"CodePath,omitempty"`
	ConfigMap     *string   `json:"ConfigMap,omitempty" xml:"ConfigMap,omitempty"`
	ConfigMapList []*string `json:"ConfigMapList,omitempty" xml:"ConfigMapList,omitempty" type:"Repeated"`
	CronJob       *string   `json:"CronJob,omitempty" xml:"CronJob,omitempty"`
	Deployment    *string   `json:"Deployment,omitempty" xml:"Deployment,omitempty"`
	Id            *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	SecretList    []*string `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Repeated"`
	StatefulSet   *string   `json:"StatefulSet,omitempty" xml:"StatefulSet,omitempty"`
}

func (s UpdateDeployConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeployConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeployConfigRequest) SetAppId(v int64) *UpdateDeployConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateDeployConfigRequest) SetCodePath(v string) *UpdateDeployConfigRequest {
	s.CodePath = &v
	return s
}

func (s *UpdateDeployConfigRequest) SetConfigMap(v string) *UpdateDeployConfigRequest {
	s.ConfigMap = &v
	return s
}

func (s *UpdateDeployConfigRequest) SetConfigMapList(v []*string) *UpdateDeployConfigRequest {
	s.ConfigMapList = v
	return s
}

func (s *UpdateDeployConfigRequest) SetCronJob(v string) *UpdateDeployConfigRequest {
	s.CronJob = &v
	return s
}

func (s *UpdateDeployConfigRequest) SetDeployment(v string) *UpdateDeployConfigRequest {
	s.Deployment = &v
	return s
}

func (s *UpdateDeployConfigRequest) SetId(v int64) *UpdateDeployConfigRequest {
	s.Id = &v
	return s
}

func (s *UpdateDeployConfigRequest) SetSecretList(v []*string) *UpdateDeployConfigRequest {
	s.SecretList = v
	return s
}

func (s *UpdateDeployConfigRequest) SetStatefulSet(v string) *UpdateDeployConfigRequest {
	s.StatefulSet = &v
	return s
}

type UpdateDeployConfigResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                               `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateDeployConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateDeployConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeployConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeployConfigResponseBody) SetCode(v int32) *UpdateDeployConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDeployConfigResponseBody) SetErrMsg(v string) *UpdateDeployConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateDeployConfigResponseBody) SetRequestId(v string) *UpdateDeployConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeployConfigResponseBody) SetResult(v *UpdateDeployConfigResponseBodyResult) *UpdateDeployConfigResponseBody {
	s.Result = v
	return s
}

type UpdateDeployConfigResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDeployConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeployConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateDeployConfigResponseBodyResult) SetSuccess(v bool) *UpdateDeployConfigResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateDeployConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDeployConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeployConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeployConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeployConfigResponse) SetHeaders(v map[string]*string) *UpdateDeployConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeployConfigResponse) SetStatusCode(v int32) *UpdateDeployConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeployConfigResponse) SetBody(v *UpdateDeployConfigResponseBody) *UpdateDeployConfigResponse {
	s.Body = v
	return s
}

type UpdateEciConfigRequest struct {
	AppEnvId                *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	EipBandwidth            *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EnableEciSchedulePolicy *bool  `json:"EnableEciSchedulePolicy,omitempty" xml:"EnableEciSchedulePolicy,omitempty"`
	MirrorCache             *bool  `json:"MirrorCache,omitempty" xml:"MirrorCache,omitempty"`
	NormalInstanceLimit     *int32 `json:"NormalInstanceLimit,omitempty" xml:"NormalInstanceLimit,omitempty"`
	ScheduleVirtualNode     *bool  `json:"ScheduleVirtualNode,omitempty" xml:"ScheduleVirtualNode,omitempty"`
}

func (s UpdateEciConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEciConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateEciConfigRequest) SetAppEnvId(v int64) *UpdateEciConfigRequest {
	s.AppEnvId = &v
	return s
}

func (s *UpdateEciConfigRequest) SetEipBandwidth(v int32) *UpdateEciConfigRequest {
	s.EipBandwidth = &v
	return s
}

func (s *UpdateEciConfigRequest) SetEnableEciSchedulePolicy(v bool) *UpdateEciConfigRequest {
	s.EnableEciSchedulePolicy = &v
	return s
}

func (s *UpdateEciConfigRequest) SetMirrorCache(v bool) *UpdateEciConfigRequest {
	s.MirrorCache = &v
	return s
}

func (s *UpdateEciConfigRequest) SetNormalInstanceLimit(v int32) *UpdateEciConfigRequest {
	s.NormalInstanceLimit = &v
	return s
}

func (s *UpdateEciConfigRequest) SetScheduleVirtualNode(v bool) *UpdateEciConfigRequest {
	s.ScheduleVirtualNode = &v
	return s
}

type UpdateEciConfigResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                            `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateEciConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateEciConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEciConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEciConfigResponseBody) SetCode(v int32) *UpdateEciConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEciConfigResponseBody) SetErrMsg(v string) *UpdateEciConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateEciConfigResponseBody) SetRequestId(v string) *UpdateEciConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEciConfigResponseBody) SetResult(v *UpdateEciConfigResponseBodyResult) *UpdateEciConfigResponseBody {
	s.Result = v
	return s
}

type UpdateEciConfigResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEciConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateEciConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateEciConfigResponseBodyResult) SetSuccess(v bool) *UpdateEciConfigResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateEciConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEciConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEciConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEciConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateEciConfigResponse) SetHeaders(v map[string]*string) *UpdateEciConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateEciConfigResponse) SetStatusCode(v int32) *UpdateEciConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEciConfigResponse) SetBody(v *UpdateEciConfigResponseBody) *UpdateEciConfigResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentRequest struct {
	AppEnvId    *int64 `json:"AppEnvId,omitempty" xml:"AppEnvId,omitempty"`
	AppId       *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSchemaId *int64 `json:"AppSchemaId,omitempty" xml:"AppSchemaId,omitempty"`
	Replicas    *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) SetAppEnvId(v int64) *UpdateEnvironmentRequest {
	s.AppEnvId = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetAppId(v int64) *UpdateEnvironmentRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetAppSchemaId(v int64) *UpdateEnvironmentRequest {
	s.AppSchemaId = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetReplicas(v int32) *UpdateEnvironmentRequest {
	s.Replicas = &v
	return s
}

type UpdateEnvironmentResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateEnvironmentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) SetCode(v int32) *UpdateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetErrMsg(v string) *UpdateEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetRequestId(v string) *UpdateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetResult(v *UpdateEnvironmentResponseBodyResult) *UpdateEnvironmentResponseBody {
	s.Result = v
	return s
}

type UpdateEnvironmentResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEnvironmentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBodyResult) SetSuccess(v bool) *UpdateEnvironmentResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateEnvironmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentResponse) SetStatusCode(v int32) *UpdateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvironmentResponse) SetBody(v *UpdateEnvironmentResponseBody) *UpdateEnvironmentResponse {
	s.Body = v
	return s
}

type UpdateNormalDeployConfigRequest struct {
	AppId                    *int64                                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerResourceLimit   *UpdateNormalDeployConfigRequestContainerResourceLimit   `json:"ContainerResourceLimit,omitempty" xml:"ContainerResourceLimit,omitempty" type:"Struct"`
	ContainerResourceRequest *UpdateNormalDeployConfigRequestContainerResourceRequest `json:"ContainerResourceRequest,omitempty" xml:"ContainerResourceRequest,omitempty" type:"Struct"`
	Id                       *int64                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateNormalDeployConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigRequest) SetAppId(v int64) *UpdateNormalDeployConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateNormalDeployConfigRequest) SetContainerResourceLimit(v *UpdateNormalDeployConfigRequestContainerResourceLimit) *UpdateNormalDeployConfigRequest {
	s.ContainerResourceLimit = v
	return s
}

func (s *UpdateNormalDeployConfigRequest) SetContainerResourceRequest(v *UpdateNormalDeployConfigRequestContainerResourceRequest) *UpdateNormalDeployConfigRequest {
	s.ContainerResourceRequest = v
	return s
}

func (s *UpdateNormalDeployConfigRequest) SetId(v int64) *UpdateNormalDeployConfigRequest {
	s.Id = &v
	return s
}

type UpdateNormalDeployConfigRequestContainerResourceLimit struct {
	Cpu    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s UpdateNormalDeployConfigRequestContainerResourceLimit) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigRequestContainerResourceLimit) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigRequestContainerResourceLimit) SetCpu(v string) *UpdateNormalDeployConfigRequestContainerResourceLimit {
	s.Cpu = &v
	return s
}

func (s *UpdateNormalDeployConfigRequestContainerResourceLimit) SetMemory(v string) *UpdateNormalDeployConfigRequestContainerResourceLimit {
	s.Memory = &v
	return s
}

type UpdateNormalDeployConfigRequestContainerResourceRequest struct {
	Cpu    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s UpdateNormalDeployConfigRequestContainerResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigRequestContainerResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigRequestContainerResourceRequest) SetCpu(v string) *UpdateNormalDeployConfigRequestContainerResourceRequest {
	s.Cpu = &v
	return s
}

func (s *UpdateNormalDeployConfigRequestContainerResourceRequest) SetMemory(v string) *UpdateNormalDeployConfigRequestContainerResourceRequest {
	s.Memory = &v
	return s
}

type UpdateNormalDeployConfigShrinkRequest struct {
	AppId                          *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerResourceLimitShrink   *string `json:"ContainerResourceLimit,omitempty" xml:"ContainerResourceLimit,omitempty"`
	ContainerResourceRequestShrink *string `json:"ContainerResourceRequest,omitempty" xml:"ContainerResourceRequest,omitempty"`
	Id                             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateNormalDeployConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigShrinkRequest) SetAppId(v int64) *UpdateNormalDeployConfigShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateNormalDeployConfigShrinkRequest) SetContainerResourceLimitShrink(v string) *UpdateNormalDeployConfigShrinkRequest {
	s.ContainerResourceLimitShrink = &v
	return s
}

func (s *UpdateNormalDeployConfigShrinkRequest) SetContainerResourceRequestShrink(v string) *UpdateNormalDeployConfigShrinkRequest {
	s.ContainerResourceRequestShrink = &v
	return s
}

func (s *UpdateNormalDeployConfigShrinkRequest) SetId(v int64) *UpdateNormalDeployConfigShrinkRequest {
	s.Id = &v
	return s
}

type UpdateNormalDeployConfigResponseBody struct {
	Code      *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrMsg    *string                                     `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateNormalDeployConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateNormalDeployConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigResponseBody) SetCode(v int32) *UpdateNormalDeployConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNormalDeployConfigResponseBody) SetErrMsg(v string) *UpdateNormalDeployConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateNormalDeployConfigResponseBody) SetRequestId(v string) *UpdateNormalDeployConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNormalDeployConfigResponseBody) SetResult(v *UpdateNormalDeployConfigResponseBodyResult) *UpdateNormalDeployConfigResponseBody {
	s.Result = v
	return s
}

type UpdateNormalDeployConfigResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNormalDeployConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigResponseBodyResult) SetSuccess(v bool) *UpdateNormalDeployConfigResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateNormalDeployConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNormalDeployConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNormalDeployConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNormalDeployConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateNormalDeployConfigResponse) SetHeaders(v map[string]*string) *UpdateNormalDeployConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateNormalDeployConfigResponse) SetStatusCode(v int32) *UpdateNormalDeployConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNormalDeployConfigResponse) SetBody(v *UpdateNormalDeployConfigResponseBody) *UpdateNormalDeployConfigResponse {
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
		"ap-northeast-1":              tea.String("retailcloud.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("retailcloud.aliyuncs.com"),
		"ap-south-1":                  tea.String("retailcloud.aliyuncs.com"),
		"ap-southeast-1":              tea.String("retailcloud.aliyuncs.com"),
		"ap-southeast-2":              tea.String("retailcloud.aliyuncs.com"),
		"ap-southeast-3":              tea.String("retailcloud.aliyuncs.com"),
		"ap-southeast-5":              tea.String("retailcloud.aliyuncs.com"),
		"cn-beijing":                  tea.String("retailcloud.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("retailcloud.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("retailcloud.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("retailcloud.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("retailcloud.aliyuncs.com"),
		"cn-chengdu":                  tea.String("retailcloud.aliyuncs.com"),
		"cn-edge-1":                   tea.String("retailcloud.aliyuncs.com"),
		"cn-fujian":                   tea.String("retailcloud.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("retailcloud.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("retailcloud.aliyuncs.com"),
		"cn-hongkong":                 tea.String("retailcloud.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("retailcloud.aliyuncs.com"),
		"cn-huhehaote":                tea.String("retailcloud.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("retailcloud.aliyuncs.com"),
		"cn-qingdao":                  tea.String("retailcloud.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("retailcloud.aliyuncs.com"),
		"cn-shanghai":                 tea.String("retailcloud.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("retailcloud.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("retailcloud.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("retailcloud.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("retailcloud.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("retailcloud.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("retailcloud.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("retailcloud.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("retailcloud.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("retailcloud.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("retailcloud.aliyuncs.com"),
		"cn-wuhan":                    tea.String("retailcloud.aliyuncs.com"),
		"cn-yushanfang":               tea.String("retailcloud.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("retailcloud.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("retailcloud.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("retailcloud.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("retailcloud.aliyuncs.com"),
		"eu-central-1":                tea.String("retailcloud.aliyuncs.com"),
		"eu-west-1":                   tea.String("retailcloud.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("retailcloud.aliyuncs.com"),
		"me-east-1":                   tea.String("retailcloud.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("retailcloud.aliyuncs.com"),
		"us-east-1":                   tea.String("retailcloud.aliyuncs.com"),
		"us-west-1":                   tea.String("retailcloud.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("retailcloud"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddClusterNodeWithOptions(request *AddClusterNodeRequest, runtime *util.RuntimeOptions) (_result *AddClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		query["ClusterInstanceId"] = request.ClusterInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceIdList)) {
		query["EcsInstanceIdList"] = request.EcsInstanceIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddClusterNode"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddClusterNode(request *AddClusterNodeRequest) (_result *AddClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddClusterNodeResponse{}
	_body, _err := client.AddClusterNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocatePodConfigWithOptions(request *AllocatePodConfigRequest, runtime *util.RuntimeOptions) (_result *AllocatePodConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocatePodConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocatePodConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocatePodConfig(request *AllocatePodConfigRequest) (_result *AllocatePodConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePodConfigResponse{}
	_body, _err := client.AllocatePodConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddServersWithOptions(request *BatchAddServersRequest, runtime *util.RuntimeOptions) (_result *BatchAddServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddServers"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddServers(request *BatchAddServersRequest) (_result *BatchAddServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddServersResponse{}
	_body, _err := client.BatchAddServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindGroupWithOptions(request *BindGroupRequest, runtime *util.RuntimeOptions) (_result *BindGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindGroup"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindGroup(request *BindGroupRequest) (_result *BindGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindGroupResponse{}
	_body, _err := client.BindGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindNodeLabelWithOptions(request *BindNodeLabelRequest, runtime *util.RuntimeOptions) (_result *BindNodeLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKey)) {
		query["LabelKey"] = request.LabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelValue)) {
		query["LabelValue"] = request.LabelValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindNodeLabel"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindNodeLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindNodeLabel(request *BindNodeLabelRequest) (_result *BindNodeLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindNodeLabelResponse{}
	_body, _err := client.BindNodeLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDeployOrderWithOptions(request *CloseDeployOrderRequest, runtime *util.RuntimeOptions) (_result *CloseDeployOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseDeployOrder"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseDeployOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDeployOrder(request *CloseDeployOrderRequest) (_result *CloseDeployOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDeployOrderResponse{}
	_body, _err := client.CloseDeployOrderWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		body["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		body["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2018-03-13"),
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

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizTitle)) {
		body["BizTitle"] = request.BizTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MiddleWareIdList)) {
		body["MiddleWareIdList"] = request.MiddleWareIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingSystem)) {
		body["OperatingSystem"] = request.OperatingSystem
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		body["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.StateType)) {
		body["StateType"] = request.StateType
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserRoles)) {
		body["UserRoles"] = request.UserRoles
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppGroupWithOptions(request *CreateAppGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppGroup"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppGroup(request *CreateAppGroupRequest) (_result *CreateAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CreateAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppMonitorsWithOptions(request *CreateAppMonitorsRequest, runtime *util.RuntimeOptions) (_result *CreateAppMonitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmTemplateId)) {
		query["AlarmTemplateId"] = request.AlarmTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.MainUserId)) {
		query["MainUserId"] = request.MainUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["AppIds"] = request.AppIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppMonitors"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppMonitors(request *CreateAppMonitorsRequest) (_result *CreateAppMonitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppMonitorsResponse{}
	_body, _err := client.CreateAppMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppResourceAllocWithOptions(request *CreateAppResourceAllocRequest, runtime *util.RuntimeOptions) (_result *CreateAppResourceAllocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppEnvId)) {
		query["AppEnvId"] = request.AppEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppResourceAlloc"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResourceAllocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppResourceAlloc(request *CreateAppResourceAllocRequest) (_result *CreateAppResourceAllocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResourceAllocResponse{}
	_body, _err := client.CreateAppResourceAllocWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessCode)) {
		query["BusinessCode"] = request.BusinessCode
	}

	if !tea.BoolValue(util.IsUnset(request.CloudMonitorFlags)) {
		query["CloudMonitorFlags"] = request.CloudMonitorFlags
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterEnvType)) {
		query["ClusterEnvType"] = request.ClusterEnvType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterTitle)) {
		query["ClusterTitle"] = request.ClusterTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateWithArmsIntegration)) {
		query["CreateWithArmsIntegration"] = request.CreateWithArmsIntegration
	}

	if !tea.BoolValue(util.IsUnset(request.CreateWithLogIntegration)) {
		query["CreateWithLogIntegration"] = request.CreateWithLogIntegration
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPair)) {
		query["KeyPair"] = request.KeyPair
	}

	if !tea.BoolValue(util.IsUnset(request.NetPlug)) {
		query["NetPlug"] = request.NetPlug
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PodCIDR)) {
		query["PodCIDR"] = request.PodCIDR
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateZone)) {
		query["PrivateZone"] = request.PrivateZone
	}

	if !tea.BoolValue(util.IsUnset(request.PublicSlb)) {
		query["PublicSlb"] = request.PublicSlb
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionName)) {
		query["RegionName"] = request.RegionName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCIDR)) {
		query["ServiceCIDR"] = request.ServiceCIDR
	}

	if !tea.BoolValue(util.IsUnset(request.SnatEntry)) {
		query["SnatEntry"] = request.SnatEntry
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Vswitchids)) {
		query["Vswitchids"] = request.Vswitchids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2018-03-13"),
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

func (client *Client) CreateDbWithOptions(request *CreateDbRequest, runtime *util.RuntimeOptions) (_result *CreateDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CharacterSetName)) {
		body["CharacterSetName"] = request.CharacterSetName
	}

	if !tea.BoolValue(util.IsUnset(request.DbDescription)) {
		body["DbDescription"] = request.DbDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDb"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDb(request *CreateDbRequest) (_result *CreateDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDbResponse{}
	_body, _err := client.CreateDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeployConfigWithOptions(request *CreateDeployConfigRequest, runtime *util.RuntimeOptions) (_result *CreateDeployConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CodePath)) {
		query["CodePath"] = request.CodePath
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigMap)) {
		query["ConfigMap"] = request.ConfigMap
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigMapList)) {
		query["ConfigMapList"] = request.ConfigMapList
	}

	if !tea.BoolValue(util.IsUnset(request.CronJob)) {
		query["CronJob"] = request.CronJob
	}

	if !tea.BoolValue(util.IsUnset(request.Deployment)) {
		query["Deployment"] = request.Deployment
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecretList)) {
		query["SecretList"] = request.SecretList
	}

	if !tea.BoolValue(util.IsUnset(request.StatefulSet)) {
		query["StatefulSet"] = request.StatefulSet
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeployConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeployConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeployConfig(request *CreateDeployConfigRequest) (_result *CreateDeployConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeployConfigResponse{}
	_body, _err := client.CreateDeployConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEciConfigWithOptions(request *CreateEciConfigRequest, runtime *util.RuntimeOptions) (_result *CreateEciConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppEnvId)) {
		query["AppEnvId"] = request.AppEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EnableEciSchedulePolicy)) {
		query["EnableEciSchedulePolicy"] = request.EnableEciSchedulePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MirrorCache)) {
		query["MirrorCache"] = request.MirrorCache
	}

	if !tea.BoolValue(util.IsUnset(request.NormalInstanceLimit)) {
		query["NormalInstanceLimit"] = request.NormalInstanceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleVirtualNode)) {
		query["ScheduleVirtualNode"] = request.ScheduleVirtualNode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEciConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEciConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEciConfig(request *CreateEciConfigRequest) (_result *CreateEciConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEciConfigResponse{}
	_body, _err := client.CreateEciConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentWithOptions(request *CreateEnvironmentRequest, runtime *util.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSchemaId)) {
		query["AppSchemaId"] = request.AppSchemaId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvName)) {
		query["EnvName"] = request.EnvName
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		query["Replicas"] = request.Replicas
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEnvironment"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironment(request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNodeLabelWithOptions(request *CreateNodeLabelRequest, runtime *util.RuntimeOptions) (_result *CreateNodeLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKey)) {
		query["LabelKey"] = request.LabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelValue)) {
		query["LabelValue"] = request.LabelValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNodeLabel"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNodeLabel(request *CreateNodeLabelRequest) (_result *CreateNodeLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeLabelResponse{}
	_body, _err := client.CreateNodeLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePersistentVolumeWithOptions(request *CreatePersistentVolumeRequest, runtime *util.RuntimeOptions) (_result *CreatePersistentVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessModes)) {
		body["AccessModes"] = request.AccessModes
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		body["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		body["ClusterInstanceId"] = request.ClusterInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MountDir)) {
		body["MountDir"] = request.MountDir
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		body["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.NFSVersion)) {
		body["NFSVersion"] = request.NFSVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NasType)) {
		body["NasType"] = request.NasType
	}

	if !tea.BoolValue(util.IsUnset(request.ReclaimPolicy)) {
		body["ReclaimPolicy"] = request.ReclaimPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.StorageClass)) {
		body["StorageClass"] = request.StorageClass
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePersistentVolume"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePersistentVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePersistentVolume(request *CreatePersistentVolumeRequest) (_result *CreatePersistentVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePersistentVolumeResponse{}
	_body, _err := client.CreatePersistentVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePersistentVolumeClaimWithOptions(request *CreatePersistentVolumeClaimRequest, runtime *util.RuntimeOptions) (_result *CreatePersistentVolumeClaimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessModes)) {
		query["AccessModes"] = request.AccessModes
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StorageClass)) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePersistentVolumeClaim"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePersistentVolumeClaimResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePersistentVolumeClaim(request *CreatePersistentVolumeClaimRequest) (_result *CreatePersistentVolumeClaimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePersistentVolumeClaimResponse{}
	_body, _err := client.CreatePersistentVolumeClaimWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Headless)) {
		query["Headless"] = request.Headless
	}

	if !tea.BoolValue(util.IsUnset(request.K8sServiceId)) {
		query["K8sServiceId"] = request.K8sServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortMappings)) {
		body["PortMappings"] = request.PortMappings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSlbAPWithOptions(request *CreateSlbAPRequest, runtime *util.RuntimeOptions) (_result *CreateSlbAPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EstablishedTimeout)) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RealServerPort)) {
		query["RealServerPort"] = request.RealServerPort
	}

	if !tea.BoolValue(util.IsUnset(request.SlbId)) {
		query["SlbId"] = request.SlbId
	}

	if !tea.BoolValue(util.IsUnset(request.SslCertId)) {
		query["SslCertId"] = request.SslCertId
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSlbAP"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSlbAPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSlbAP(request *CreateSlbAPRequest) (_result *CreateSlbAPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSlbAPResponse{}
	_body, _err := client.CreateSlbAPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppDetailWithOptions(request *DeleteAppDetailRequest, runtime *util.RuntimeOptions) (_result *DeleteAppDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppDetail(request *DeleteAppDetailRequest) (_result *DeleteAppDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppDetailResponse{}
	_body, _err := client.DeleteAppDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppEnvironmentWithOptions(request *DeleteAppEnvironmentRequest, runtime *util.RuntimeOptions) (_result *DeleteAppEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppEnvironment"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppEnvironment(request *DeleteAppEnvironmentRequest) (_result *DeleteAppEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppEnvironmentResponse{}
	_body, _err := client.DeleteAppEnvironmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppGroupWithOptions(request *DeleteAppGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppGroup"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppGroup(request *DeleteAppGroupRequest) (_result *DeleteAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppGroupResponse{}
	_body, _err := client.DeleteAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppResourceAllocWithOptions(request *DeleteAppResourceAllocRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResourceAllocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppEnvId)) {
		query["AppEnvId"] = request.AppEnvId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppResourceAlloc"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResourceAllocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppResourceAlloc(request *DeleteAppResourceAllocRequest) (_result *DeleteAppResourceAllocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResourceAllocResponse{}
	_body, _err := client.DeleteAppResourceAllocWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		query["ClusterInstanceId"] = request.ClusterInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		body["DBName"] = request.DBName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabase"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeployConfigWithOptions(request *DeleteDeployConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDeployConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaId)) {
		query["SchemaId"] = request.SchemaId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeployConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeployConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeployConfig(request *DeleteDeployConfigRequest) (_result *DeleteDeployConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeployConfigResponse{}
	_body, _err := client.DeleteDeployConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodeLabelWithOptions(request *DeleteNodeLabelRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKey)) {
		query["LabelKey"] = request.LabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelValue)) {
		query["LabelValue"] = request.LabelValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodeLabel"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNodeLabel(request *DeleteNodeLabelRequest) (_result *DeleteNodeLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeLabelResponse{}
	_body, _err := client.DeleteNodeLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePersistentVolumeWithOptions(request *DeletePersistentVolumeRequest, runtime *util.RuntimeOptions) (_result *DeletePersistentVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		body["ClusterInstanceId"] = request.ClusterInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PersistentVolumeName)) {
		body["PersistentVolumeName"] = request.PersistentVolumeName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePersistentVolume"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePersistentVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePersistentVolume(request *DeletePersistentVolumeRequest) (_result *DeletePersistentVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePersistentVolumeResponse{}
	_body, _err := client.DeletePersistentVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePersistentVolumeClaimWithOptions(request *DeletePersistentVolumeClaimRequest, runtime *util.RuntimeOptions) (_result *DeletePersistentVolumeClaimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.PersistentVolumeClaimName)) {
		query["PersistentVolumeClaimName"] = request.PersistentVolumeClaimName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePersistentVolumeClaim"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePersistentVolumeClaimResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePersistentVolumeClaim(request *DeletePersistentVolumeClaimRequest) (_result *DeletePersistentVolumeClaimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePersistentVolumeClaimResponse{}
	_body, _err := client.DeletePersistentVolumeClaimWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRdsAccountWithOptions(request *DeleteRdsAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteRdsAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRdsAccount"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRdsAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRdsAccount(request *DeleteRdsAccountRequest) (_result *DeleteRdsAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRdsAccountResponse{}
	_body, _err := client.DeleteRdsAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSlbAPWithOptions(request *DeleteSlbAPRequest, runtime *util.RuntimeOptions) (_result *DeleteSlbAPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SlbAPId)) {
		query["SlbAPId"] = request.SlbAPId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSlbAP"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSlbAPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSlbAP(request *DeleteSlbAPRequest) (_result *DeleteSlbAPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSlbAPResponse{}
	_body, _err := client.DeleteSlbAPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployAppWithOptions(request *DeployAppRequest, runtime *util.RuntimeOptions) (_result *DeployAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArmsFlag)) {
		query["ArmsFlag"] = request.ArmsFlag
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerImageList)) {
		query["ContainerImageList"] = request.ContainerImageList
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultPacketOfAppGroup)) {
		query["DefaultPacketOfAppGroup"] = request.DefaultPacketOfAppGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DeployPacketId)) {
		query["DeployPacketId"] = request.DeployPacketId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployPacketUrl)) {
		query["DeployPacketUrl"] = request.DeployPacketUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainerImageList)) {
		query["InitContainerImageList"] = request.InitContainerImageList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PauseType)) {
		query["PauseType"] = request.PauseType
	}

	if !tea.BoolValue(util.IsUnset(request.TotalPartitions)) {
		query["TotalPartitions"] = request.TotalPartitions
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStrategyType)) {
		query["UpdateStrategyType"] = request.UpdateStrategyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApp"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployApp(request *DeployAppRequest) (_result *DeployAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployAppResponse{}
	_body, _err := client.DeployAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppDetailWithOptions(request *DescribeAppDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeAppDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppDetail(request *DescribeAppDetailRequest) (_result *DescribeAppDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppDetailResponse{}
	_body, _err := client.DescribeAppDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppEnvDeployBaselineWithOptions(request *DescribeAppEnvDeployBaselineRequest, runtime *util.RuntimeOptions) (_result *DescribeAppEnvDeployBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppEnvDeployBaseline"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppEnvDeployBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppEnvDeployBaseline(request *DescribeAppEnvDeployBaselineRequest) (_result *DescribeAppEnvDeployBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppEnvDeployBaselineResponse{}
	_body, _err := client.DescribeAppEnvDeployBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppEnvironmentDetailWithOptions(request *DescribeAppEnvironmentDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeAppEnvironmentDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppEnvironmentDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppEnvironmentDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppEnvironmentDetail(request *DescribeAppEnvironmentDetailRequest) (_result *DescribeAppEnvironmentDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppEnvironmentDetailResponse{}
	_body, _err := client.DescribeAppEnvironmentDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppGroupDeploySettingWithOptions(request *DescribeAppGroupDeploySettingRequest, runtime *util.RuntimeOptions) (_result *DescribeAppGroupDeploySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppGroupId)) {
		query["AppGroupId"] = request.AppGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppGroupDeploySetting"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppGroupDeploySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppGroupDeploySetting(request *DescribeAppGroupDeploySettingRequest) (_result *DescribeAppGroupDeploySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppGroupDeploySettingResponse{}
	_body, _err := client.DescribeAppGroupDeploySettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppMonitorMetricWithOptions(request *DescribeAppMonitorMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeAppMonitorMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.PodName)) {
		query["PodName"] = request.PodName
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
		Action:      tea.String("DescribeAppMonitorMetric"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppMonitorMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppMonitorMetric(request *DescribeAppMonitorMetricRequest) (_result *DescribeAppMonitorMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppMonitorMetricResponse{}
	_body, _err := client.DescribeAppMonitorMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppResourceAllocWithOptions(request *DescribeAppResourceAllocRequest, runtime *util.RuntimeOptions) (_result *DescribeAppResourceAllocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppResourceAlloc"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppResourceAllocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppResourceAlloc(request *DescribeAppResourceAllocRequest) (_result *DescribeAppResourceAllocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppResourceAllocResponse{}
	_body, _err := client.DescribeAppResourceAllocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterDetailWithOptions(request *DescribeClusterDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		query["ClusterInstanceId"] = request.ClusterInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (_result *DescribeClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DescribeClusterDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabasesWithOptions(request *DescribeDatabasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabases"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (_result *DescribeDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DescribeDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeployOrderDetailWithOptions(request *DescribeDeployOrderDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDeployOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeployOrderDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeployOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeployOrderDetail(request *DescribeDeployOrderDetailRequest) (_result *DescribeDeployOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeployOrderDetailResponse{}
	_body, _err := client.DescribeDeployOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEciConfigWithOptions(request *DescribeEciConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeEciConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEciConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEciConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEciConfig(request *DescribeEciConfigRequest) (_result *DescribeEciConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEciConfigResponse{}
	_body, _err := client.DescribeEciConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventMonitorListWithOptions(request *DescribeEventMonitorListRequest, runtime *util.RuntimeOptions) (_result *DescribeEventMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EventLevel)) {
		query["EventLevel"] = request.EventLevel
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PodName)) {
		query["PodName"] = request.PodName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventMonitorList"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventMonitorListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventMonitorList(request *DescribeEventMonitorListRequest) (_result *DescribeEventMonitorListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventMonitorListResponse{}
	_body, _err := client.DescribeEventMonitorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobLogWithOptions(request *DescribeJobLogRequest, runtime *util.RuntimeOptions) (_result *DescribeJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJobLog"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobLog(request *DescribeJobLogRequest) (_result *DescribeJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobLogResponse{}
	_body, _err := client.DescribeJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePodContainerLogListWithOptions(request *DescribePodContainerLogListRequest, runtime *util.RuntimeOptions) (_result *DescribePodContainerLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["Line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.PodName)) {
		query["PodName"] = request.PodName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePodContainerLogList"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePodContainerLogListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePodContainerLogList(request *DescribePodContainerLogListRequest) (_result *DescribePodContainerLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePodContainerLogListResponse{}
	_body, _err := client.DescribePodContainerLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePodEventsWithOptions(request *DescribePodEventsRequest, runtime *util.RuntimeOptions) (_result *DescribePodEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstId)) {
		query["AppInstId"] = request.AppInstId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePodEvents"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePodEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePodEvents(request *DescribePodEventsRequest) (_result *DescribePodEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePodEventsResponse{}
	_body, _err := client.DescribePodEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePodLogWithOptions(request *DescribePodLogRequest, runtime *util.RuntimeOptions) (_result *DescribePodLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstId)) {
		body["AppInstId"] = request.AppInstId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		body["DeployOrderId"] = request.DeployOrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePodLog"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePodLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePodLog(request *DescribePodLogRequest) (_result *DescribePodLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePodLogResponse{}
	_body, _err := client.DescribePodLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsAccountsWithOptions(request *DescribeRdsAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdsAccounts"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdsAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsAccounts(request *DescribeRdsAccountsRequest) (_result *DescribeRdsAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsAccountsResponse{}
	_body, _err := client.DescribeRdsAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceDetailWithOptions(request *DescribeServiceDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceDetail(request *DescribeServiceDetailRequest) (_result *DescribeServiceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceDetailResponse{}
	_body, _err := client.DescribeServiceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlbAPDetailWithOptions(request *DescribeSlbAPDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeSlbAPDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SlbAPId)) {
		query["SlbAPId"] = request.SlbAPId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlbAPDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlbAPDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlbAPDetail(request *DescribeSlbAPDetailRequest) (_result *DescribeSlbAPDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlbAPDetailResponse{}
	_body, _err := client.DescribeSlbAPDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstTransInfoWithOptions(request *GetInstTransInfoRequest, runtime *util.RuntimeOptions) (_result *GetInstTransInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunCommodityCode)) {
		body["aliyunCommodityCode"] = request.AliyunCommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunEquipId)) {
		body["aliyunEquipId"] = request.AliyunEquipId
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunUid)) {
		body["aliyunUid"] = request.AliyunUid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstTransInfo"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstTransInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstTransInfo(request *GetInstTransInfoRequest) (_result *GetInstTransInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstTransInfoResponse{}
	_body, _err := client.GetInstTransInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRdsBackUpWithOptions(request *GetRdsBackUpRequest, runtime *util.RuntimeOptions) (_result *GetRdsBackUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		body["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		body["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRdsBackUp"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRdsBackUpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRdsBackUp(request *GetRdsBackUpRequest) (_result *GetRdsBackUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRdsBackUpResponse{}
	_body, _err := client.GetRdsBackUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantDbToAccountWithOptions(request *GrantDbToAccountRequest, runtime *util.RuntimeOptions) (_result *GrantDbToAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPrivilege)) {
		body["AccountPrivilege"] = request.AccountPrivilege
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantDbToAccount"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantDbToAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantDbToAccount(request *GrantDbToAccountRequest) (_result *GrantDbToAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantDbToAccountResponse{}
	_body, _err := client.GrantDbToAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppWithOptions(request *ListAppRequest, runtime *util.RuntimeOptions) (_result *ListAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApp"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApp(request *ListAppRequest) (_result *ListAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppResponse{}
	_body, _err := client.ListAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppCmsGroupsWithOptions(request *ListAppCmsGroupsRequest, runtime *util.RuntimeOptions) (_result *ListAppCmsGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppCmsGroups"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppCmsGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppCmsGroups(request *ListAppCmsGroupsRequest) (_result *ListAppCmsGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppCmsGroupsResponse{}
	_body, _err := client.ListAppCmsGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppEnvironmentWithOptions(request *ListAppEnvironmentRequest, runtime *util.RuntimeOptions) (_result *ListAppEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppEnvironment"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppEnvironment(request *ListAppEnvironmentRequest) (_result *ListAppEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppEnvironmentResponse{}
	_body, _err := client.ListAppEnvironmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppGroupWithOptions(request *ListAppGroupRequest, runtime *util.RuntimeOptions) (_result *ListAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["BizCode"] = request.BizCode
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
		Action:      tea.String("ListAppGroup"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppGroup(request *ListAppGroupRequest) (_result *ListAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppGroupResponse{}
	_body, _err := client.ListAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppGroupMappingWithOptions(request *ListAppGroupMappingRequest, runtime *util.RuntimeOptions) (_result *ListAppGroupMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
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
		Action:      tea.String("ListAppGroupMapping"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppGroupMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppGroupMapping(request *ListAppGroupMappingRequest) (_result *ListAppGroupMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppGroupMappingResponse{}
	_body, _err := client.ListAppGroupMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppInstanceWithOptions(request *ListAppInstanceRequest, runtime *util.RuntimeOptions) (_result *ListAppInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppInstance"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppInstance(request *ListAppInstanceRequest) (_result *ListAppInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInstanceResponse{}
	_body, _err := client.ListAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppResourceAllocsWithOptions(request *ListAppResourceAllocsRequest, runtime *util.RuntimeOptions) (_result *ListAppResourceAllocsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppResourceAllocs"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppResourceAllocsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppResourceAllocs(request *ListAppResourceAllocsRequest) (_result *ListAppResourceAllocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppResourceAllocsResponse{}
	_body, _err := client.ListAppResourceAllocsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableClusterNodeWithOptions(request *ListAvailableClusterNodeRequest, runtime *util.RuntimeOptions) (_result *ListAvailableClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableClusterNode"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableClusterNode(request *ListAvailableClusterNodeRequest) (_result *ListAvailableClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableClusterNodeResponse{}
	_body, _err := client.ListAvailableClusterNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterWithOptions(request *ListClusterRequest, runtime *util.RuntimeOptions) (_result *ListClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCluster"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCluster(request *ListClusterRequest) (_result *ListClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterResponse{}
	_body, _err := client.ListClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterNodeWithOptions(request *ListClusterNodeRequest, runtime *util.RuntimeOptions) (_result *ListClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterNode"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterNode(request *ListClusterNodeRequest) (_result *ListClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterNodeResponse{}
	_body, _err := client.ListClusterNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployConfigWithOptions(request *ListDeployConfigRequest, runtime *util.RuntimeOptions) (_result *ListDeployConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeployConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployConfig(request *ListDeployConfigRequest) (_result *ListDeployConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployConfigResponse{}
	_body, _err := client.ListDeployConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployOrdersWithOptions(request *ListDeployOrdersRequest, runtime *util.RuntimeOptions) (_result *ListDeployOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployCategory)) {
		query["DeployCategory"] = request.DeployCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGreaterThan)) {
		query["EndTimeGreaterThan"] = request.EndTimeGreaterThan
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGreaterThanOrEqualTo)) {
		query["EndTimeGreaterThanOrEqualTo"] = request.EndTimeGreaterThanOrEqualTo
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeLessThan)) {
		query["EndTimeLessThan"] = request.EndTimeLessThan
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeLessThanOrEqualTo)) {
		query["EndTimeLessThanOrEqualTo"] = request.EndTimeLessThanOrEqualTo
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionType)) {
		query["PartitionType"] = request.PartitionType
	}

	if !tea.BoolValue(util.IsUnset(request.PauseType)) {
		query["PauseType"] = request.PauseType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeGreaterThan)) {
		query["StartTimeGreaterThan"] = request.StartTimeGreaterThan
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeGreaterThanOrEqualTo)) {
		query["StartTimeGreaterThanOrEqualTo"] = request.StartTimeGreaterThanOrEqualTo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeLessThan)) {
		query["StartTimeLessThan"] = request.StartTimeLessThan
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeLessThanOrEqualTo)) {
		query["StartTimeLessThanOrEqualTo"] = request.StartTimeLessThanOrEqualTo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResultList)) {
		body["ResultList"] = request.ResultList
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["StatusList"] = request.StatusList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployOrders"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeployOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployOrders(request *ListDeployOrdersRequest) (_result *ListDeployOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployOrdersResponse{}
	_body, _err := client.ListDeployOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobHistoriesWithOptions(request *ListJobHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListJobHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobHistories"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobHistories(request *ListJobHistoriesRequest) (_result *ListJobHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobHistoriesResponse{}
	_body, _err := client.ListJobHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeLabelBindingsWithOptions(request *ListNodeLabelBindingsRequest, runtime *util.RuntimeOptions) (_result *ListNodeLabelBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeLabelBindings"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeLabelBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeLabelBindings(request *ListNodeLabelBindingsRequest) (_result *ListNodeLabelBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeLabelBindingsResponse{}
	_body, _err := client.ListNodeLabelBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeLabelsWithOptions(request *ListNodeLabelsRequest, runtime *util.RuntimeOptions) (_result *ListNodeLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKey)) {
		query["LabelKey"] = request.LabelKey
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
		Action:      tea.String("ListNodeLabels"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeLabels(request *ListNodeLabelsRequest) (_result *ListNodeLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeLabelsResponse{}
	_body, _err := client.ListNodeLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersistentVolumeWithOptions(request *ListPersistentVolumeRequest, runtime *util.RuntimeOptions) (_result *ListPersistentVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		body["ClusterInstanceId"] = request.ClusterInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPersistentVolume"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPersistentVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersistentVolume(request *ListPersistentVolumeRequest) (_result *ListPersistentVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersistentVolumeResponse{}
	_body, _err := client.ListPersistentVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersistentVolumeClaimWithOptions(request *ListPersistentVolumeClaimRequest, runtime *util.RuntimeOptions) (_result *ListPersistentVolumeClaimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
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
		Action:      tea.String("ListPersistentVolumeClaim"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPersistentVolumeClaimResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersistentVolumeClaim(request *ListPersistentVolumeClaimRequest) (_result *ListPersistentVolumeClaimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersistentVolumeClaimResponse{}
	_body, _err := client.ListPersistentVolumeClaimWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPodsWithOptions(request *ListPodsRequest, runtime *util.RuntimeOptions) (_result *ListPodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResultList)) {
		body["ResultList"] = request.ResultList
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["StatusList"] = request.StatusList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPods"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPodsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPods(request *ListPodsRequest) (_result *ListPodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPodsResponse{}
	_body, _err := client.ListPodsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSlbAPsWithOptions(request *ListSlbAPsRequest, runtime *util.RuntimeOptions) (_result *ListSlbAPsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkMode)) {
		query["NetworkMode"] = request.NetworkMode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SlbId)) {
		query["SlbId"] = request.SlbId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProtocolList)) {
		body["ProtocolList"] = request.ProtocolList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSlbAPs"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSlbAPsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlbAPs(request *ListSlbAPsRequest) (_result *ListSlbAPsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSlbAPsResponse{}
	_body, _err := client.ListSlbAPsWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyServiceWithOptions(request *ModifyServiceRequest, runtime *util.RuntimeOptions) (_result *ModifyServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortMappings)) {
		body["PortMappings"] = request.PortMappings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyService"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyService(request *ModifyServiceRequest) (_result *ModifyServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyServiceResponse{}
	_body, _err := client.ModifyServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySlbAPWithOptions(request *ModifySlbAPRequest, runtime *util.RuntimeOptions) (_result *ModifySlbAPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.EstablishedTimeout)) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RealServerPort)) {
		query["RealServerPort"] = request.RealServerPort
	}

	if !tea.BoolValue(util.IsUnset(request.SlbAPId)) {
		query["SlbAPId"] = request.SlbAPId
	}

	if !tea.BoolValue(util.IsUnset(request.SslCertId)) {
		query["SslCertId"] = request.SslCertId
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySlbAP"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySlbAPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySlbAP(request *ModifySlbAPRequest) (_result *ModifySlbAPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySlbAPResponse{}
	_body, _err := client.ModifySlbAPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OfflineAppEnvironmentWithOptions(request *OfflineAppEnvironmentRequest, runtime *util.RuntimeOptions) (_result *OfflineAppEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeletePvc)) {
		query["DeletePvc"] = request.DeletePvc
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineAppEnvironment"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OfflineAppEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OfflineAppEnvironment(request *OfflineAppEnvironmentRequest) (_result *OfflineAppEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OfflineAppEnvironmentResponse{}
	_body, _err := client.OfflineAppEnvironmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClusterDetailWithOptions(request *QueryClusterDetailRequest, runtime *util.RuntimeOptions) (_result *QueryClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryClusterDetail"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClusterDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClusterDetail(request *QueryClusterDetailRequest) (_result *QueryClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryClusterDetailResponse{}
	_body, _err := client.QueryClusterDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildAppInstanceWithOptions(request *RebuildAppInstanceRequest, runtime *util.RuntimeOptions) (_result *RebuildAppInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebuildAppInstance"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebuildAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildAppInstance(request *RebuildAppInstanceRequest) (_result *RebuildAppInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebuildAppInstanceResponse{}
	_body, _err := client.RebuildAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveClusterNodeWithOptions(request *RemoveClusterNodeRequest, runtime *util.RuntimeOptions) (_result *RemoveClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterInstanceId)) {
		query["ClusterInstanceId"] = request.ClusterInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceIdList)) {
		query["EcsInstanceIdList"] = request.EcsInstanceIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveClusterNode"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClusterNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveClusterNode(request *RemoveClusterNodeRequest) (_result *RemoveClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveClusterNodeResponse{}
	_body, _err := client.RemoveClusterNodeWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		body["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DbInstanceId)) {
		body["DbInstanceId"] = request.DbInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2018-03-13"),
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

func (client *Client) ResourceStatusNotifyWithOptions(request *ResourceStatusNotifyRequest, runtime *util.RuntimeOptions) (_result *ResourceStatusNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResourceStatusNotify"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &ResourceStatusNotifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResourceStatusNotify(request *ResourceStatusNotifyRequest) (_result *ResourceStatusNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResourceStatusNotifyResponse{}
	_body, _err := client.ResourceStatusNotifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartAppInstanceWithOptions(request *RestartAppInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartAppInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceIdList)) {
		query["AppInstanceIdList"] = request.AppInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartAppInstance"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartAppInstance(request *RestartAppInstanceRequest) (_result *RestartAppInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartAppInstanceResponse{}
	_body, _err := client.RestartAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeDeployWithOptions(request *ResumeDeployRequest, runtime *util.RuntimeOptions) (_result *ResumeDeployResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeDeploy"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeDeployResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeDeploy(request *ResumeDeployRequest) (_result *ResumeDeployResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeDeployResponse{}
	_body, _err := client.ResumeDeployWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleAppWithOptions(request *ScaleAppRequest, runtime *util.RuntimeOptions) (_result *ScaleAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		query["Replicas"] = request.Replicas
	}

	if !tea.BoolValue(util.IsUnset(request.TotalPartitions)) {
		query["TotalPartitions"] = request.TotalPartitions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScaleApp"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScaleAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleApp(request *ScaleAppRequest) (_result *ScaleAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleAppResponse{}
	_body, _err := client.ScaleAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeployPauseTypeWithOptions(request *SetDeployPauseTypeRequest, runtime *util.RuntimeOptions) (_result *SetDeployPauseTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployOrderId)) {
		query["DeployOrderId"] = request.DeployOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployPauseType)) {
		query["DeployPauseType"] = request.DeployPauseType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeployPauseType"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeployPauseTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeployPauseType(request *SetDeployPauseTypeRequest) (_result *SetDeployPauseTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeployPauseTypeResponse{}
	_body, _err := client.SetDeployPauseTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitInfoWithOptions(request *SubmitInfoRequest, runtime *util.RuntimeOptions) (_result *SubmitInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["CallerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.MainUserId)) {
		query["MainUserId"] = request.MainUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcsDescList)) {
		body["EcsDescList"] = request.EcsDescList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitInfo"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitInfo(request *SubmitInfoRequest) (_result *SubmitInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitInfoResponse{}
	_body, _err := client.SubmitInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncPodInfoWithOptions(request *SyncPodInfoRequest, runtime *util.RuntimeOptions) (_result *SyncPodInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PodName)) {
		query["PodName"] = request.PodName
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SideCarType)) {
		query["SideCarType"] = request.SideCarType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncPodInfo"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncPodInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncPodInfo(request *SyncPodInfoRequest) (_result *SyncPodInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncPodInfoResponse{}
	_body, _err := client.SyncPodInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindGroupWithOptions(request *UnbindGroupRequest, runtime *util.RuntimeOptions) (_result *UnbindGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindGroup"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindGroup(request *UnbindGroupRequest) (_result *UnbindGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindGroupResponse{}
	_body, _err := client.UnbindGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindNodeLabelWithOptions(request *UnbindNodeLabelRequest, runtime *util.RuntimeOptions) (_result *UnbindNodeLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKey)) {
		query["LabelKey"] = request.LabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelValue)) {
		query["LabelValue"] = request.LabelValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindNodeLabel"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindNodeLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindNodeLabel(request *UnbindNodeLabelRequest) (_result *UnbindNodeLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindNodeLabelResponse{}
	_body, _err := client.UnbindNodeLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizTitle)) {
		body["BizTitle"] = request.BizTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MiddleWareIdList)) {
		body["MiddleWareIdList"] = request.MiddleWareIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingSystem)) {
		body["OperatingSystem"] = request.OperatingSystem
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		body["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.UserRoles)) {
		body["UserRoles"] = request.UserRoles
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppMonitorsWithOptions(request *UpdateAppMonitorsRequest, runtime *util.RuntimeOptions) (_result *UpdateAppMonitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MainUserId)) {
		body["MainUserId"] = request.MainUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorIds)) {
		body["MonitorIds"] = request.MonitorIds
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		body["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppMonitors"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppMonitors(request *UpdateAppMonitorsRequest) (_result *UpdateAppMonitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppMonitorsResponse{}
	_body, _err := client.UpdateAppMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeployConfigWithOptions(request *UpdateDeployConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateDeployConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CodePath)) {
		query["CodePath"] = request.CodePath
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigMap)) {
		query["ConfigMap"] = request.ConfigMap
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigMapList)) {
		query["ConfigMapList"] = request.ConfigMapList
	}

	if !tea.BoolValue(util.IsUnset(request.CronJob)) {
		query["CronJob"] = request.CronJob
	}

	if !tea.BoolValue(util.IsUnset(request.Deployment)) {
		query["Deployment"] = request.Deployment
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.SecretList)) {
		query["SecretList"] = request.SecretList
	}

	if !tea.BoolValue(util.IsUnset(request.StatefulSet)) {
		query["StatefulSet"] = request.StatefulSet
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDeployConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeployConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeployConfig(request *UpdateDeployConfigRequest) (_result *UpdateDeployConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeployConfigResponse{}
	_body, _err := client.UpdateDeployConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEciConfigWithOptions(request *UpdateEciConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateEciConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppEnvId)) {
		query["AppEnvId"] = request.AppEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EnableEciSchedulePolicy)) {
		query["EnableEciSchedulePolicy"] = request.EnableEciSchedulePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MirrorCache)) {
		query["MirrorCache"] = request.MirrorCache
	}

	if !tea.BoolValue(util.IsUnset(request.NormalInstanceLimit)) {
		query["NormalInstanceLimit"] = request.NormalInstanceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleVirtualNode)) {
		query["ScheduleVirtualNode"] = request.ScheduleVirtualNode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEciConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEciConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEciConfig(request *UpdateEciConfigRequest) (_result *UpdateEciConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEciConfigResponse{}
	_body, _err := client.UpdateEciConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentWithOptions(request *UpdateEnvironmentRequest, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppEnvId)) {
		query["AppEnvId"] = request.AppEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSchemaId)) {
		query["AppSchemaId"] = request.AppSchemaId
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		query["Replicas"] = request.Replicas
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnvironment"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironment(request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNormalDeployConfigWithOptions(tmpReq *UpdateNormalDeployConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateNormalDeployConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateNormalDeployConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ContainerResourceLimit))) {
		request.ContainerResourceLimitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ContainerResourceLimit), tea.String("ContainerResourceLimit"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ContainerResourceRequest))) {
		request.ContainerResourceRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ContainerResourceRequest), tea.String("ContainerResourceRequest"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerResourceLimitShrink)) {
		query["ContainerResourceLimit"] = request.ContainerResourceLimitShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerResourceRequestShrink)) {
		query["ContainerResourceRequest"] = request.ContainerResourceRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNormalDeployConfig"),
		Version:     tea.String("2018-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNormalDeployConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNormalDeployConfig(request *UpdateNormalDeployConfigRequest) (_result *UpdateNormalDeployConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNormalDeployConfigResponse{}
	_body, _err := client.UpdateNormalDeployConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

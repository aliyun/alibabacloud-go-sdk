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

type AddGrafanaRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddGrafanaRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaRequest) GoString() string {
	return s.String()
}

func (s *AddGrafanaRequest) SetClusterId(v string) *AddGrafanaRequest {
	s.ClusterId = &v
	return s
}

func (s *AddGrafanaRequest) SetIntegration(v string) *AddGrafanaRequest {
	s.Integration = &v
	return s
}

func (s *AddGrafanaRequest) SetRegionId(v string) *AddGrafanaRequest {
	s.RegionId = &v
	return s
}

type AddGrafanaResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponseBody) SetData(v string) *AddGrafanaResponseBody {
	s.Data = &v
	return s
}

func (s *AddGrafanaResponseBody) SetRequestId(v string) *AddGrafanaResponseBody {
	s.RequestId = &v
	return s
}

type AddGrafanaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGrafanaResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaResponse) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponse) SetHeaders(v map[string]*string) *AddGrafanaResponse {
	s.Headers = v
	return s
}

func (s *AddGrafanaResponse) SetStatusCode(v int32) *AddGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGrafanaResponse) SetBody(v *AddGrafanaResponseBody) *AddGrafanaResponse {
	s.Body = v
	return s
}

type AddIntegrationRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationRequest) GoString() string {
	return s.String()
}

func (s *AddIntegrationRequest) SetClusterId(v string) *AddIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *AddIntegrationRequest) SetIntegration(v string) *AddIntegrationRequest {
	s.Integration = &v
	return s
}

func (s *AddIntegrationRequest) SetRegionId(v string) *AddIntegrationRequest {
	s.RegionId = &v
	return s
}

type AddIntegrationResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponseBody) SetData(v string) *AddIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *AddIntegrationResponseBody) SetRequestId(v string) *AddIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type AddIntegrationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationResponse) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponse) SetHeaders(v map[string]*string) *AddIntegrationResponse {
	s.Headers = v
	return s
}

func (s *AddIntegrationResponse) SetStatusCode(v int32) *AddIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIntegrationResponse) SetBody(v *AddIntegrationResponseBody) *AddIntegrationResponse {
	s.Body = v
	return s
}

type ApplyScenarioRequest struct {
	AppId        *string                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Config       map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario     *string                `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign         *string                `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SnDump       *bool                  `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce      *bool                  `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	SnStat       *bool                  `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnTransfer   *bool                  `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	UpdateOption *bool                  `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioRequest) SetAppId(v string) *ApplyScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioRequest) SetConfig(v map[string]interface{}) *ApplyScenarioRequest {
	s.Config = v
	return s
}

func (s *ApplyScenarioRequest) SetName(v string) *ApplyScenarioRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioRequest) SetRegionId(v string) *ApplyScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioRequest) SetScenario(v string) *ApplyScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioRequest) SetSign(v string) *ApplyScenarioRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnDump(v bool) *ApplyScenarioRequest {
	s.SnDump = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnForce(v bool) *ApplyScenarioRequest {
	s.SnForce = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnStat(v bool) *ApplyScenarioRequest {
	s.SnStat = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnTransfer(v bool) *ApplyScenarioRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioRequest) SetUpdateOption(v bool) *ApplyScenarioRequest {
	s.UpdateOption = &v
	return s
}

type ApplyScenarioShrinkRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario     *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign         *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SnDump       *bool   `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce      *bool   `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	SnStat       *bool   `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnTransfer   *bool   `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	UpdateOption *bool   `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioShrinkRequest) SetAppId(v string) *ApplyScenarioShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetConfigShrink(v string) *ApplyScenarioShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetName(v string) *ApplyScenarioShrinkRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetRegionId(v string) *ApplyScenarioShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetScenario(v string) *ApplyScenarioShrinkRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSign(v string) *ApplyScenarioShrinkRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnDump(v bool) *ApplyScenarioShrinkRequest {
	s.SnDump = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnForce(v bool) *ApplyScenarioShrinkRequest {
	s.SnForce = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnStat(v bool) *ApplyScenarioShrinkRequest {
	s.SnStat = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnTransfer(v bool) *ApplyScenarioShrinkRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetUpdateOption(v bool) *ApplyScenarioShrinkRequest {
	s.UpdateOption = &v
	return s
}

type ApplyScenarioResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ApplyScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyScenarioResponseBody) SetRequestId(v string) *ApplyScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyScenarioResponseBody) SetResult(v string) *ApplyScenarioResponseBody {
	s.Result = &v
	return s
}

type ApplyScenarioResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyScenarioResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioResponse) GoString() string {
	return s.String()
}

func (s *ApplyScenarioResponse) SetHeaders(v map[string]*string) *ApplyScenarioResponse {
	s.Headers = v
	return s
}

func (s *ApplyScenarioResponse) SetStatusCode(v int32) *ApplyScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyScenarioResponse) SetBody(v *ApplyScenarioResponseBody) *ApplyScenarioResponse {
	s.Body = v
	return s
}

type CheckDataConsistencyRequest struct {
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentTimestamp *int64  `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	Pid              *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckDataConsistencyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDataConsistencyRequest) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyRequest) SetAppType(v string) *CheckDataConsistencyRequest {
	s.AppType = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetCurrentTimestamp(v int64) *CheckDataConsistencyRequest {
	s.CurrentTimestamp = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetPid(v string) *CheckDataConsistencyRequest {
	s.Pid = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetProxyUserId(v string) *CheckDataConsistencyRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetRegionId(v string) *CheckDataConsistencyRequest {
	s.RegionId = &v
	return s
}

type CheckDataConsistencyResponseBody struct {
	IsDataConsistency *bool   `json:"IsDataConsistency,omitempty" xml:"IsDataConsistency,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDataConsistencyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDataConsistencyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyResponseBody) SetIsDataConsistency(v bool) *CheckDataConsistencyResponseBody {
	s.IsDataConsistency = &v
	return s
}

func (s *CheckDataConsistencyResponseBody) SetRequestId(v string) *CheckDataConsistencyResponseBody {
	s.RequestId = &v
	return s
}

type CheckDataConsistencyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDataConsistencyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDataConsistencyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDataConsistencyResponse) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyResponse) SetHeaders(v map[string]*string) *CheckDataConsistencyResponse {
	s.Headers = v
	return s
}

func (s *CheckDataConsistencyResponse) SetStatusCode(v int32) *CheckDataConsistencyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataConsistencyResponse) SetBody(v *CheckDataConsistencyResponseBody) *CheckDataConsistencyResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleForDeletingRequest struct {
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SPIRegionId    *string `json:"SPIRegionId,omitempty" xml:"SPIRegionId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CheckServiceLinkedRoleForDeletingRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetDeletionTaskId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRoleArn(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetSPIRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.SPIRegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetServiceName(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.ServiceName = &v
	return s
}

type CheckServiceLinkedRoleForDeletingResponseBody struct {
	Deletable  *bool                                                      `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleUsages []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetDeletable(v bool) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.Deletable = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRoleUsages(v []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RoleUsages = v
	return s
}

type CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages struct {
	Region    *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) SetRegion(v string) *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	s.Region = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) SetResources(v []*string) *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	s.Resources = v
	return s
}

type CheckServiceLinkedRoleForDeletingResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckServiceLinkedRoleForDeletingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleForDeletingResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleForDeletingResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleForDeletingResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetBody(v *CheckServiceLinkedRoleForDeletingResponseBody) *CheckServiceLinkedRoleForDeletingResponse {
	s.Body = v
	return s
}

type CheckServiceStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SvcCode  *string `json:"SvcCode,omitempty" xml:"SvcCode,omitempty"`
}

func (s CheckServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusRequest) SetRegionId(v string) *CheckServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceStatusRequest) SetSvcCode(v string) *CheckServiceStatusRequest {
	s.SvcCode = &v
	return s
}

type CheckServiceStatusResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusResponseBody) SetData(v string) *CheckServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckServiceStatusResponseBody) SetRequestId(v string) *CheckServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type CheckServiceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusResponse) SetHeaders(v map[string]*string) *CheckServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceStatusResponse) SetStatusCode(v int32) *CheckServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceStatusResponse) SetBody(v *CheckServiceStatusResponseBody) *CheckServiceStatusResponse {
	s.Body = v
	return s
}

type ConfigAppRequest struct {
	AppIds   *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	Enable   *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppRequest) GoString() string {
	return s.String()
}

func (s *ConfigAppRequest) SetAppIds(v string) *ConfigAppRequest {
	s.AppIds = &v
	return s
}

func (s *ConfigAppRequest) SetEnable(v string) *ConfigAppRequest {
	s.Enable = &v
	return s
}

func (s *ConfigAppRequest) SetRegionId(v string) *ConfigAppRequest {
	s.RegionId = &v
	return s
}

type ConfigAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAppResponseBody) SetData(v string) *ConfigAppResponseBody {
	s.Data = &v
	return s
}

func (s *ConfigAppResponseBody) SetRequestId(v string) *ConfigAppResponseBody {
	s.RequestId = &v
	return s
}

type ConfigAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppResponse) GoString() string {
	return s.String()
}

func (s *ConfigAppResponse) SetHeaders(v map[string]*string) *ConfigAppResponse {
	s.Headers = v
	return s
}

func (s *ConfigAppResponse) SetStatusCode(v int32) *ConfigAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigAppResponse) SetBody(v *ConfigAppResponseBody) *ConfigAppResponse {
	s.Body = v
	return s
}

type CreateAlertContactRequest struct {
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
}

func (s CreateAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactRequest) SetContactName(v string) *CreateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateAlertContactRequest) SetDingRobotWebhookUrl(v string) *CreateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *CreateAlertContactRequest) SetEmail(v string) *CreateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *CreateAlertContactRequest) SetPhoneNum(v string) *CreateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateAlertContactRequest) SetRegionId(v string) *CreateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertContactRequest) SetSystemNoc(v bool) *CreateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

type CreateAlertContactResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponseBody) SetContactId(v string) *CreateAlertContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateAlertContactResponseBody) SetRequestId(v string) *CreateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponse) SetHeaders(v map[string]*string) *CreateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactResponse) SetStatusCode(v int32) *CreateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactResponse) SetBody(v *CreateAlertContactResponseBody) *CreateAlertContactResponse {
	s.Body = v
	return s
}

type CreateAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupRequest) SetContactGroupName(v string) *CreateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetContactIds(v string) *CreateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetRegionId(v string) *CreateAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type CreateAlertContactGroupResponseBody struct {
	ContactGroupId *string `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponseBody) SetContactGroupId(v string) *CreateAlertContactGroupResponseBody {
	s.ContactGroupId = &v
	return s
}

func (s *CreateAlertContactGroupResponseBody) SetRequestId(v string) *CreateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponse) SetHeaders(v map[string]*string) *CreateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactGroupResponse) SetStatusCode(v int32) *CreateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactGroupResponse) SetBody(v *CreateAlertContactGroupResponseBody) *CreateAlertContactGroupResponse {
	s.Body = v
	return s
}

type CreateAlertTemplateRequest struct {
	Annotations      *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Labels           *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId         *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Rule             *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	TemplateProvider *string `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAlertTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateRequest) SetAnnotations(v string) *CreateAlertTemplateRequest {
	s.Annotations = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetLabels(v string) *CreateAlertTemplateRequest {
	s.Labels = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetMatchExpressions(v string) *CreateAlertTemplateRequest {
	s.MatchExpressions = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetMessage(v string) *CreateAlertTemplateRequest {
	s.Message = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetName(v string) *CreateAlertTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetParentId(v string) *CreateAlertTemplateRequest {
	s.ParentId = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetRegionId(v string) *CreateAlertTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetRule(v string) *CreateAlertTemplateRequest {
	s.Rule = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetTemplateProvider(v string) *CreateAlertTemplateRequest {
	s.TemplateProvider = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetType(v string) *CreateAlertTemplateRequest {
	s.Type = &v
	return s
}

type CreateAlertTemplateResponseBody struct {
	AlertTemplate *CreateAlertTemplateResponseBodyAlertTemplate `json:"AlertTemplate,omitempty" xml:"AlertTemplate,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBody) SetAlertTemplate(v *CreateAlertTemplateResponseBodyAlertTemplate) *CreateAlertTemplateResponseBody {
	s.AlertTemplate = v
	return s
}

func (s *CreateAlertTemplateResponseBody) SetRequestId(v string) *CreateAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertTemplateResponseBodyAlertTemplate struct {
	AlertProvider            *string                                                               `json:"AlertProvider,omitempty" xml:"AlertProvider,omitempty"`
	Annotations              map[string]interface{}                                                `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Id                       *int32                                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelMatchExpressionGrid *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Labels                   map[string]interface{}                                                `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message                  *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId                 *int32                                                                `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Public                   *bool                                                                 `json:"Public,omitempty" xml:"Public,omitempty"`
	Rule                     *string                                                               `json:"Rule,omitempty" xml:"Rule,omitempty"`
	Status                   *bool                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateProvider         *string                                                               `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	Type                     *string                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId                   *string                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplate) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetAlertProvider(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.AlertProvider = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetAnnotations(v map[string]interface{}) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Annotations = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetId(v int32) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Id = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetLabelMatchExpressionGrid(v *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetLabels(v map[string]interface{}) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Labels = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetMessage(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Message = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetName(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Name = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetParentId(v int32) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.ParentId = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetPublic(v bool) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Public = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetRule(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Rule = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetStatus(v bool) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Status = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetTemplateProvider(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.TemplateProvider = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetType(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Type = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetUserId(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.UserId = &v
	return s
}

type CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid struct {
	LabelMatchExpressionGroups []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

type CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	LabelMatchExpressions []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

type CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

type CreateAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponse) SetHeaders(v map[string]*string) *CreateAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertTemplateResponse) SetStatusCode(v int32) *CreateAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertTemplateResponse) SetBody(v *CreateAlertTemplateResponseBody) *CreateAlertTemplateResponse {
	s.Body = v
	return s
}

type CreateDispatchRuleRequest struct {
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleRequest) SetDispatchRule(v string) *CreateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *CreateDispatchRuleRequest) SetRegionId(v string) *CreateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type CreateDispatchRuleResponseBody struct {
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleResponseBody) SetDispatchRuleId(v int64) *CreateDispatchRuleResponseBody {
	s.DispatchRuleId = &v
	return s
}

func (s *CreateDispatchRuleResponseBody) SetRequestId(v string) *CreateDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleResponse) SetHeaders(v map[string]*string) *CreateDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDispatchRuleResponse) SetStatusCode(v int32) *CreateDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDispatchRuleResponse) SetBody(v *CreateDispatchRuleResponseBody) *CreateDispatchRuleResponse {
	s.Body = v
	return s
}

type CreatePrometheusAlertRuleRequest struct {
	AlertName      *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleRequest) SetAlertName(v string) *CreatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetAnnotations(v string) *CreatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetClusterId(v string) *CreatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDuration(v string) *CreatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetExpression(v string) *CreatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetLabels(v string) *CreatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetMessage(v string) *CreatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetNotifyType(v string) *CreatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetRegionId(v string) *CreatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetType(v string) *CreatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

type CreatePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *CreatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetRequestId(v string) *CreatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

type CreatePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *CreatePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) SetStatusCode(v int32) *CreatePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) SetBody(v *CreatePrometheusAlertRuleResponseBody) *CreatePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type CreateRetcodeAppRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
}

func (s CreateRetcodeAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppRequest) SetRegionId(v string) *CreateRetcodeAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppName(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppType(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppType = &v
	return s
}

type CreateRetcodeAppResponseBody struct {
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetcodeAppDataBean *CreateRetcodeAppResponseBodyRetcodeAppDataBean `json:"RetcodeAppDataBean,omitempty" xml:"RetcodeAppDataBean,omitempty" type:"Struct"`
}

func (s CreateRetcodeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBody) SetRequestId(v string) *CreateRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetRetcodeAppDataBean(v *CreateRetcodeAppResponseBodyRetcodeAppDataBean) *CreateRetcodeAppResponseBody {
	s.RetcodeAppDataBean = v
	return s
}

type CreateRetcodeAppResponseBodyRetcodeAppDataBean struct {
	AppId *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetAppId(v int64) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.AppId = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetPid(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.Pid = &v
	return s
}

type CreateRetcodeAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRetcodeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponse) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponse) SetHeaders(v map[string]*string) *CreateRetcodeAppResponse {
	s.Headers = v
	return s
}

func (s *CreateRetcodeAppResponse) SetStatusCode(v int32) *CreateRetcodeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRetcodeAppResponse) SetBody(v *CreateRetcodeAppResponseBody) *CreateRetcodeAppResponse {
	s.Body = v
	return s
}

type CreateWehookRequest struct {
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateWehookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWehookRequest) GoString() string {
	return s.String()
}

func (s *CreateWehookRequest) SetBody(v string) *CreateWehookRequest {
	s.Body = &v
	return s
}

func (s *CreateWehookRequest) SetContactName(v string) *CreateWehookRequest {
	s.ContactName = &v
	return s
}

func (s *CreateWehookRequest) SetHttpHeaders(v string) *CreateWehookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *CreateWehookRequest) SetHttpParams(v string) *CreateWehookRequest {
	s.HttpParams = &v
	return s
}

func (s *CreateWehookRequest) SetMethod(v string) *CreateWehookRequest {
	s.Method = &v
	return s
}

func (s *CreateWehookRequest) SetRegionId(v string) *CreateWehookRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWehookRequest) SetUrl(v string) *CreateWehookRequest {
	s.Url = &v
	return s
}

type CreateWehookResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWehookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWehookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWehookResponseBody) SetContactId(v string) *CreateWehookResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateWehookResponseBody) SetRequestId(v string) *CreateWehookResponseBody {
	s.RequestId = &v
	return s
}

type CreateWehookResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWehookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWehookResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWehookResponse) GoString() string {
	return s.String()
}

func (s *CreateWehookResponse) SetHeaders(v map[string]*string) *CreateWehookResponse {
	s.Headers = v
	return s
}

func (s *CreateWehookResponse) SetStatusCode(v int32) *CreateWehookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWehookResponse) SetBody(v *CreateWehookResponseBody) *CreateWehookResponse {
	s.Body = v
	return s
}

type DeleteAlertContactRequest struct {
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactRequest) SetContactId(v int64) *DeleteAlertContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteAlertContactRequest) SetRegionId(v string) *DeleteAlertContactRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponseBody) SetIsSuccess(v bool) *DeleteAlertContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertContactResponseBody) SetRequestId(v string) *DeleteAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponse) SetHeaders(v map[string]*string) *DeleteAlertContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactResponse) SetStatusCode(v int32) *DeleteAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactResponse) SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse {
	s.Body = v
	return s
}

type DeleteAlertContactGroupRequest struct {
	ContactGroupId *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupRequest) SetContactGroupId(v int64) *DeleteAlertContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *DeleteAlertContactGroupRequest) SetRegionId(v string) *DeleteAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertContactGroupResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponseBody) SetIsSuccess(v bool) *DeleteAlertContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) SetRequestId(v string) *DeleteAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponse) SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetStatusCode(v int32) *DeleteAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetBody(v *DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse {
	s.Body = v
	return s
}

type DeleteAlertRulesRequest struct {
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesRequest) SetAlertIds(v string) *DeleteAlertRulesRequest {
	s.AlertIds = &v
	return s
}

func (s *DeleteAlertRulesRequest) SetRegionId(v string) *DeleteAlertRulesRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertRulesResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesResponseBody) SetIsSuccess(v bool) *DeleteAlertRulesResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertRulesResponseBody) SetRequestId(v string) *DeleteAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesResponse) SetHeaders(v map[string]*string) *DeleteAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertRulesResponse) SetStatusCode(v int32) *DeleteAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertRulesResponse) SetBody(v *DeleteAlertRulesResponseBody) *DeleteAlertRulesResponse {
	s.Body = v
	return s
}

type DeleteAlertTemplateRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertTemplateRequest) SetId(v int64) *DeleteAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertTemplateRequest) SetRegionId(v string) *DeleteAlertTemplateRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAlertTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertTemplateResponseBody) SetRequestId(v string) *DeleteAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertTemplateResponseBody) SetSuccess(v bool) *DeleteAlertTemplateResponseBody {
	s.Success = &v
	return s
}

type DeleteAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertTemplateResponse) SetHeaders(v map[string]*string) *DeleteAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertTemplateResponse) SetStatusCode(v int32) *DeleteAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertTemplateResponse) SetBody(v *DeleteAlertTemplateResponseBody) *DeleteAlertTemplateResponse {
	s.Body = v
	return s
}

type DeleteDispatchRuleRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleRequest) SetId(v string) *DeleteDispatchRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDispatchRuleRequest) SetRegionId(v string) *DeleteDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteDispatchRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleResponseBody) SetRequestId(v string) *DeleteDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDispatchRuleResponseBody) SetSuccess(v bool) *DeleteDispatchRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleResponse) SetHeaders(v map[string]*string) *DeleteDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDispatchRuleResponse) SetStatusCode(v int32) *DeleteDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDispatchRuleResponse) SetBody(v *DeleteDispatchRuleResponseBody) *DeleteDispatchRuleResponse {
	s.Body = v
	return s
}

type DeleteGrafanaResourceRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteGrafanaResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGrafanaResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceRequest) SetClusterId(v string) *DeleteGrafanaResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetClusterName(v string) *DeleteGrafanaResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetUserId(v string) *DeleteGrafanaResourceRequest {
	s.UserId = &v
	return s
}

type DeleteGrafanaResourceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGrafanaResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGrafanaResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponseBody) SetData(v string) *DeleteGrafanaResourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) SetRequestId(v string) *DeleteGrafanaResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGrafanaResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGrafanaResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGrafanaResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGrafanaResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponse) SetHeaders(v map[string]*string) *DeleteGrafanaResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGrafanaResourceResponse) SetStatusCode(v int32) *DeleteGrafanaResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGrafanaResourceResponse) SetBody(v *DeleteGrafanaResourceResponseBody) *DeleteGrafanaResourceResponse {
	s.Body = v
	return s
}

type DeletePrometheusAlertRuleRequest struct {
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DeletePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleRequest) SetAlertId(v int64) *DeletePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

type DeletePrometheusAlertRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponseBody) SetRequestId(v string) *DeletePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) SetSuccess(v bool) *DeletePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

type DeletePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *DeletePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) SetStatusCode(v int32) *DeletePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) SetBody(v *DeletePrometheusAlertRuleResponseBody) *DeletePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type DeleteRetcodeAppRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRetcodeAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppRequest) SetAppId(v string) *DeleteRetcodeAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRetcodeAppRequest) SetRegionId(v string) *DeleteRetcodeAppRequest {
	s.RegionId = &v
	return s
}

type DeleteRetcodeAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRetcodeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponseBody) SetData(v string) *DeleteRetcodeAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetRequestId(v string) *DeleteRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRetcodeAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRetcodeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponse) SetHeaders(v map[string]*string) *DeleteRetcodeAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteRetcodeAppResponse) SetStatusCode(v int32) *DeleteRetcodeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRetcodeAppResponse) SetBody(v *DeleteRetcodeAppResponseBody) *DeleteRetcodeAppResponse {
	s.Body = v
	return s
}

type DeleteScenarioRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScenarioId *int64  `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
}

func (s DeleteScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioRequest) GoString() string {
	return s.String()
}

func (s *DeleteScenarioRequest) SetRegionId(v string) *DeleteScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScenarioRequest) SetScenarioId(v int64) *DeleteScenarioRequest {
	s.ScenarioId = &v
	return s
}

type DeleteScenarioResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponseBody) SetRequestId(v string) *DeleteScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScenarioResponseBody) SetResult(v bool) *DeleteScenarioResponseBody {
	s.Result = &v
	return s
}

type DeleteScenarioResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScenarioResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioResponse) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponse) SetHeaders(v map[string]*string) *DeleteScenarioResponse {
	s.Headers = v
	return s
}

func (s *DeleteScenarioResponse) SetStatusCode(v int32) *DeleteScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScenarioResponse) SetBody(v *DeleteScenarioResponseBody) *DeleteScenarioResponse {
	s.Body = v
	return s
}

type DeleteTraceAppRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteTraceAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppRequest) SetAppId(v string) *DeleteTraceAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetPid(v string) *DeleteTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *DeleteTraceAppRequest) SetRegionId(v string) *DeleteTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetType(v string) *DeleteTraceAppRequest {
	s.Type = &v
	return s
}

type DeleteTraceAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTraceAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponseBody) SetData(v string) *DeleteTraceAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetRequestId(v string) *DeleteTraceAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTraceAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTraceAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponse) SetHeaders(v map[string]*string) *DeleteTraceAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteTraceAppResponse) SetStatusCode(v int32) *DeleteTraceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTraceAppResponse) SetBody(v *DeleteTraceAppResponseBody) *DeleteTraceAppResponse {
	s.Body = v
	return s
}

type DescribeDispatchRuleRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleRequest) SetId(v string) *DescribeDispatchRuleRequest {
	s.Id = &v
	return s
}

func (s *DescribeDispatchRuleRequest) SetRegionId(v string) *DescribeDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type DescribeDispatchRuleResponseBody struct {
	DispatchRule *DescribeDispatchRuleResponseBodyDispatchRule `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBody) SetDispatchRule(v *DescribeDispatchRuleResponseBodyDispatchRule) *DescribeDispatchRuleResponseBody {
	s.DispatchRule = v
	return s
}

func (s *DescribeDispatchRuleResponseBody) SetRequestId(v string) *DescribeDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRule struct {
	GroupRules               []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules             `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
	LabelMatchExpressionGrid *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyRules              []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules            `json:"NotifyRules,omitempty" xml:"NotifyRules,omitempty" type:"Repeated"`
	RuleId                   *int64                                                                `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	State                    *string                                                               `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetGroupRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.GroupRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetLabelMatchExpressionGrid(v *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.Name = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetNotifyRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.NotifyRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetRuleId(v int64) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetState(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.State = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleGroupRules struct {
	GroupId        *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupInterval  *int64    `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	GroupWaitTime  *int64    `json:"GroupWaitTime,omitempty" xml:"GroupWaitTime,omitempty"`
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupId(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupInterval(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupInterval = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupWaitTime(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupWaitTime = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupingFields(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupingFields = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid struct {
	LabelMatchExpressionGroups []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	LabelMatchExpressions []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules struct {
	NotifyChannels []*string                                                               `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	NotifyObjects  []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyChannels(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyChannels = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyObjects(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyObjects = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyObjectId *string `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.Name = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyObjectId(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyType(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyType = &v
	return s
}

type DescribeDispatchRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponse) SetHeaders(v map[string]*string) *DescribeDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDispatchRuleResponse) SetStatusCode(v int32) *DescribeDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDispatchRuleResponse) SetBody(v *DescribeDispatchRuleResponseBody) *DescribeDispatchRuleResponse {
	s.Body = v
	return s
}

type DescribePrometheusAlertRuleRequest struct {
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DescribePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleRequest) SetAlertId(v int64) *DescribePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

type DescribePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) *DescribePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetRequestId(v string) *DescribePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                   `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                  `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                   `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                  `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                  `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

type DescribePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *DescribePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) SetStatusCode(v int32) *DescribePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) SetBody(v *DescribePrometheusAlertRuleResponseBody) *DescribePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type DescribeTraceLicenseKeyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTraceLicenseKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyRequest) SetRegionId(v string) *DescribeTraceLicenseKeyRequest {
	s.RegionId = &v
	return s
}

type DescribeTraceLicenseKeyResponseBody struct {
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTraceLicenseKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponseBody) SetLicenseKey(v string) *DescribeTraceLicenseKeyResponseBody {
	s.LicenseKey = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponseBody) SetRequestId(v string) *DescribeTraceLicenseKeyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTraceLicenseKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTraceLicenseKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTraceLicenseKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponse) SetHeaders(v map[string]*string) *DescribeTraceLicenseKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) SetStatusCode(v int32) *DescribeTraceLicenseKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) SetBody(v *DescribeTraceLicenseKeyResponseBody) *DescribeTraceLicenseKeyResponse {
	s.Body = v
	return s
}

type DescribeTraceLocationRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTraceLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLocationRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationRequest) SetRegionId(v string) *DescribeTraceLocationRequest {
	s.RegionId = &v
	return s
}

type DescribeTraceLocationResponseBody struct {
	RegionConfigs []*DescribeTraceLocationResponseBodyRegionConfigs `json:"RegionConfigs,omitempty" xml:"RegionConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTraceLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponseBody) SetRegionConfigs(v []*DescribeTraceLocationResponseBodyRegionConfigs) *DescribeTraceLocationResponseBody {
	s.RegionConfigs = v
	return s
}

func (s *DescribeTraceLocationResponseBody) SetRequestId(v string) *DescribeTraceLocationResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTraceLocationResponseBodyRegionConfigs struct {
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeTraceLocationResponseBodyRegionConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLocationResponseBodyRegionConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) SetRegionNo(v string) *DescribeTraceLocationResponseBodyRegionConfigs {
	s.RegionNo = &v
	return s
}

func (s *DescribeTraceLocationResponseBodyRegionConfigs) SetUrl(v string) *DescribeTraceLocationResponseBodyRegionConfigs {
	s.Url = &v
	return s
}

type DescribeTraceLocationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTraceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTraceLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLocationResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponse) SetHeaders(v map[string]*string) *DescribeTraceLocationResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceLocationResponse) SetStatusCode(v int32) *DescribeTraceLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceLocationResponse) SetBody(v *DescribeTraceLocationResponseBody) *DescribeTraceLocationResponse {
	s.Body = v
	return s
}

type DisableAlertTemplateRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableAlertTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *DisableAlertTemplateRequest) SetId(v int64) *DisableAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *DisableAlertTemplateRequest) SetRegionId(v string) *DisableAlertTemplateRequest {
	s.RegionId = &v
	return s
}

type DisableAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableAlertTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAlertTemplateResponseBody) SetRequestId(v string) *DisableAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAlertTemplateResponseBody) SetSuccess(v bool) *DisableAlertTemplateResponseBody {
	s.Success = &v
	return s
}

type DisableAlertTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAlertTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *DisableAlertTemplateResponse) SetHeaders(v map[string]*string) *DisableAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *DisableAlertTemplateResponse) SetStatusCode(v int32) *DisableAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAlertTemplateResponse) SetBody(v *DisableAlertTemplateResponseBody) *DisableAlertTemplateResponse {
	s.Body = v
	return s
}

type EnableAlertTemplateRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAlertTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *EnableAlertTemplateRequest) SetId(v int64) *EnableAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *EnableAlertTemplateRequest) SetRegionId(v string) *EnableAlertTemplateRequest {
	s.RegionId = &v
	return s
}

type EnableAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAlertTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAlertTemplateResponseBody) SetRequestId(v string) *EnableAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableAlertTemplateResponseBody) SetSuccess(v bool) *EnableAlertTemplateResponseBody {
	s.Success = &v
	return s
}

type EnableAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAlertTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *EnableAlertTemplateResponse) SetHeaders(v map[string]*string) *EnableAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *EnableAlertTemplateResponse) SetStatusCode(v int32) *EnableAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAlertTemplateResponse) SetBody(v *EnableAlertTemplateResponseBody) *EnableAlertTemplateResponse {
	s.Body = v
	return s
}

type ExportPrometheusRulesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportPrometheusRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportPrometheusRulesRequest) GoString() string {
	return s.String()
}

func (s *ExportPrometheusRulesRequest) SetClusterId(v string) *ExportPrometheusRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ExportPrometheusRulesRequest) SetName(v string) *ExportPrometheusRulesRequest {
	s.Name = &v
	return s
}

func (s *ExportPrometheusRulesRequest) SetNameSpace(v string) *ExportPrometheusRulesRequest {
	s.NameSpace = &v
	return s
}

func (s *ExportPrometheusRulesRequest) SetRegionId(v string) *ExportPrometheusRulesRequest {
	s.RegionId = &v
	return s
}

type ExportPrometheusRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportPrometheusRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportPrometheusRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ExportPrometheusRulesResponseBody) SetData(v string) *ExportPrometheusRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ExportPrometheusRulesResponseBody) SetRequestId(v string) *ExportPrometheusRulesResponseBody {
	s.RequestId = &v
	return s
}

type ExportPrometheusRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExportPrometheusRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportPrometheusRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportPrometheusRulesResponse) GoString() string {
	return s.String()
}

func (s *ExportPrometheusRulesResponse) SetHeaders(v map[string]*string) *ExportPrometheusRulesResponse {
	s.Headers = v
	return s
}

func (s *ExportPrometheusRulesResponse) SetStatusCode(v int32) *ExportPrometheusRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportPrometheusRulesResponse) SetBody(v *ExportPrometheusRulesResponseBody) *ExportPrometheusRulesResponse {
	s.Body = v
	return s
}

type GetAgentDownloadUrlRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAgentDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlRequest) SetRegionId(v string) *GetAgentDownloadUrlRequest {
	s.RegionId = &v
	return s
}

type GetAgentDownloadUrlResponseBody struct {
	ArmsAgentDownloadUrl *string `json:"ArmsAgentDownloadUrl,omitempty" xml:"ArmsAgentDownloadUrl,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlResponseBody) SetArmsAgentDownloadUrl(v string) *GetAgentDownloadUrlResponseBody {
	s.ArmsAgentDownloadUrl = &v
	return s
}

func (s *GetAgentDownloadUrlResponseBody) SetRequestId(v string) *GetAgentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetAgentDownloadUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetAgentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentDownloadUrlResponse) SetStatusCode(v int32) *GetAgentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentDownloadUrlResponse) SetBody(v *GetAgentDownloadUrlResponseBody) *GetAgentDownloadUrlResponse {
	s.Body = v
	return s
}

type GetAppApiByPageRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IntervalMills *int32  `json:"IntervalMills,omitempty" xml:"IntervalMills,omitempty"`
	PId           *string `json:"PId,omitempty" xml:"PId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppApiByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageRequest) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageRequest) SetCurrentPage(v int32) *GetAppApiByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAppApiByPageRequest) SetEndTime(v int64) *GetAppApiByPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppApiByPageRequest) SetIntervalMills(v int32) *GetAppApiByPageRequest {
	s.IntervalMills = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPId(v string) *GetAppApiByPageRequest {
	s.PId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPageSize(v int32) *GetAppApiByPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageRequest) SetRegionId(v string) *GetAppApiByPageRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetStartTime(v int64) *GetAppApiByPageRequest {
	s.StartTime = &v
	return s
}

type GetAppApiByPageResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAppApiByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppApiByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBody) SetCode(v int32) *GetAppApiByPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody {
	s.Data = v
	return s
}

func (s *GetAppApiByPageResponseBody) SetMessage(v string) *GetAppApiByPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetRequestId(v string) *GetAppApiByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetSuccess(v bool) *GetAppApiByPageResponseBody {
	s.Success = &v
	return s
}

type GetAppApiByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *string                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAppApiByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBodyData) SetItems(v []map[string]interface{}) *GetAppApiByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPage(v int32) *GetAppApiByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPageSize(v int32) *GetAppApiByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetTotal(v string) *GetAppApiByPageResponseBodyData {
	s.Total = &v
	return s
}

type GetAppApiByPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppApiByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppApiByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponse) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponse) SetHeaders(v map[string]*string) *GetAppApiByPageResponse {
	s.Headers = v
	return s
}

func (s *GetAppApiByPageResponse) SetStatusCode(v int32) *GetAppApiByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppApiByPageResponse) SetBody(v *GetAppApiByPageResponseBody) *GetAppApiByPageResponse {
	s.Body = v
	return s
}

type GetConsistencySnapshotRequest struct {
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentTimestamp *int64  `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	Pid              *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsistencySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsistencySnapshotRequest) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotRequest) SetAppType(v string) *GetConsistencySnapshotRequest {
	s.AppType = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetCurrentTimestamp(v int64) *GetConsistencySnapshotRequest {
	s.CurrentTimestamp = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetPid(v string) *GetConsistencySnapshotRequest {
	s.Pid = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetProxyUserId(v string) *GetConsistencySnapshotRequest {
	s.ProxyUserId = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetRegionId(v string) *GetConsistencySnapshotRequest {
	s.RegionId = &v
	return s
}

type GetConsistencySnapshotResponseBody struct {
	ConsistencyResultJsonStr *string `json:"ConsistencyResultJsonStr,omitempty" xml:"ConsistencyResultJsonStr,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConsistencySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsistencySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotResponseBody) SetConsistencyResultJsonStr(v string) *GetConsistencySnapshotResponseBody {
	s.ConsistencyResultJsonStr = &v
	return s
}

func (s *GetConsistencySnapshotResponseBody) SetRequestId(v string) *GetConsistencySnapshotResponseBody {
	s.RequestId = &v
	return s
}

type GetConsistencySnapshotResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConsistencySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConsistencySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsistencySnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotResponse) SetHeaders(v map[string]*string) *GetConsistencySnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetConsistencySnapshotResponse) SetStatusCode(v int32) *GetConsistencySnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsistencySnapshotResponse) SetBody(v *GetConsistencySnapshotResponseBody) *GetConsistencySnapshotResponse {
	s.Body = v
	return s
}

type GetIntegrationTokenRequest struct {
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIntegrationTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationTokenRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationTokenRequest) SetProductType(v string) *GetIntegrationTokenRequest {
	s.ProductType = &v
	return s
}

func (s *GetIntegrationTokenRequest) SetRegionId(v string) *GetIntegrationTokenRequest {
	s.RegionId = &v
	return s
}

type GetIntegrationTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetIntegrationTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationTokenResponseBody) SetRequestId(v string) *GetIntegrationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegrationTokenResponseBody) SetToken(v string) *GetIntegrationTokenResponseBody {
	s.Token = &v
	return s
}

type GetIntegrationTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIntegrationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIntegrationTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationTokenResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationTokenResponse) SetHeaders(v map[string]*string) *GetIntegrationTokenResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationTokenResponse) SetStatusCode(v int32) *GetIntegrationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationTokenResponse) SetBody(v *GetIntegrationTokenResponseBody) *GetIntegrationTokenResponse {
	s.Body = v
	return s
}

type GetMultipleTraceRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceIDs []*string `json:"TraceIDs,omitempty" xml:"TraceIDs,omitempty" type:"Repeated"`
}

func (s GetMultipleTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceRequest) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceRequest) SetRegionId(v string) *GetMultipleTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultipleTraceRequest) SetTraceIDs(v []*string) *GetMultipleTraceRequest {
	s.TraceIDs = v
	return s
}

type GetMultipleTraceResponseBody struct {
	MultiCallChainInfos []*GetMultipleTraceResponseBodyMultiCallChainInfos `json:"MultiCallChainInfos,omitempty" xml:"MultiCallChainInfos,omitempty" type:"Repeated"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultipleTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBody) SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody {
	s.MultiCallChainInfos = v
	return s
}

func (s *GetMultipleTraceResponseBody) SetRequestId(v string) *GetMultipleTraceResponseBody {
	s.RequestId = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfos struct {
	Spans   []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
	TraceID *string                                                 `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetSpans(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.Spans = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.TraceID = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpans struct {
	Duration      *int64                                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                                               `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	OperationName *string                                                             `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                                             `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                                             `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType       *int32                                                              `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServiceIp     *string                                                             `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                                             `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                                             `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                                                              `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                                             `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetDuration(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Duration = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetHaveStack(v bool) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.HaveStack = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetLogEventList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.LogEventList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetOperationName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.OperationName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetParentSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetResultCode(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ResultCode = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcType(v int32) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcType = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceIp(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceIp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.SpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TraceID = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList struct {
	TagEntryList []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp    *int64                                                                          `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.Timestamp = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Value = &v
	return s
}

type GetMultipleTraceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMultipleTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMultipleTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponse) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponse) SetHeaders(v map[string]*string) *GetMultipleTraceResponse {
	s.Headers = v
	return s
}

func (s *GetMultipleTraceResponse) SetStatusCode(v int32) *GetMultipleTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultipleTraceResponse) SetBody(v *GetMultipleTraceResponseBody) *GetMultipleTraceResponse {
	s.Body = v
	return s
}

type GetPrometheusApiTokenRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPrometheusApiTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusApiTokenRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenRequest) SetRegionId(v string) *GetPrometheusApiTokenRequest {
	s.RegionId = &v
	return s
}

type GetPrometheusApiTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetPrometheusApiTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusApiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenResponseBody) SetRequestId(v string) *GetPrometheusApiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusApiTokenResponseBody) SetToken(v string) *GetPrometheusApiTokenResponseBody {
	s.Token = &v
	return s
}

type GetPrometheusApiTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPrometheusApiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrometheusApiTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusApiTokenResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenResponse) SetHeaders(v map[string]*string) *GetPrometheusApiTokenResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusApiTokenResponse) SetStatusCode(v int32) *GetPrometheusApiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusApiTokenResponse) SetBody(v *GetPrometheusApiTokenResponseBody) *GetPrometheusApiTokenResponse {
	s.Body = v
	return s
}

type GetRetcodeShareUrlRequest struct {
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetRetcodeShareUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRetcodeShareUrlRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlRequest) SetPid(v string) *GetRetcodeShareUrlRequest {
	s.Pid = &v
	return s
}

type GetRetcodeShareUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRetcodeShareUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRetcodeShareUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponseBody) SetRequestId(v string) *GetRetcodeShareUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeShareUrlResponseBody) SetUrl(v string) *GetRetcodeShareUrlResponseBody {
	s.Url = &v
	return s
}

type GetRetcodeShareUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRetcodeShareUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRetcodeShareUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRetcodeShareUrlResponse) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponse) SetHeaders(v map[string]*string) *GetRetcodeShareUrlResponse {
	s.Headers = v
	return s
}

func (s *GetRetcodeShareUrlResponse) SetStatusCode(v int32) *GetRetcodeShareUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRetcodeShareUrlResponse) SetBody(v *GetRetcodeShareUrlResponseBody) *GetRetcodeShareUrlResponse {
	s.Body = v
	return s
}

type GetStackRequest struct {
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RpcID    *string `json:"RpcID,omitempty" xml:"RpcID,omitempty"`
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) SetPid(v string) *GetStackRequest {
	s.Pid = &v
	return s
}

func (s *GetStackRequest) SetRegionId(v string) *GetStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackRequest) SetRpcID(v string) *GetStackRequest {
	s.RpcID = &v
	return s
}

func (s *GetStackRequest) SetTraceID(v string) *GetStackRequest {
	s.TraceID = &v
	return s
}

type GetStackResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackInfo []*GetStackResponseBodyStackInfo `json:"StackInfo,omitempty" xml:"StackInfo,omitempty" type:"Repeated"`
}

func (s GetStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetStackInfo(v []*GetStackResponseBodyStackInfo) *GetStackResponseBody {
	s.StackInfo = v
	return s
}

type GetStackResponseBodyStackInfo struct {
	Api         *string                               `json:"Api,omitempty" xml:"Api,omitempty"`
	Duration    *int64                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Exception   *string                               `json:"Exception,omitempty" xml:"Exception,omitempty"`
	ExtInfo     *GetStackResponseBodyStackInfoExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	Line        *string                               `json:"Line,omitempty" xml:"Line,omitempty"`
	RpcId       *string                               `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName *string                               `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStackResponseBodyStackInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyStackInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfo) SetApi(v string) *GetStackResponseBodyStackInfo {
	s.Api = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetDuration(v int64) *GetStackResponseBodyStackInfo {
	s.Duration = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetException(v string) *GetStackResponseBodyStackInfo {
	s.Exception = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetExtInfo(v *GetStackResponseBodyStackInfoExtInfo) *GetStackResponseBodyStackInfo {
	s.ExtInfo = v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetLine(v string) *GetStackResponseBodyStackInfo {
	s.Line = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetRpcId(v string) *GetStackResponseBodyStackInfo {
	s.RpcId = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetServiceName(v string) *GetStackResponseBodyStackInfo {
	s.ServiceName = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetStartTime(v int64) *GetStackResponseBodyStackInfo {
	s.StartTime = &v
	return s
}

type GetStackResponseBodyStackInfoExtInfo struct {
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStackResponseBodyStackInfoExtInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyStackInfoExtInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetInfo(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Info = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetType(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Type = &v
	return s
}

type GetStackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponse) GoString() string {
	return s.String()
}

func (s *GetStackResponse) SetHeaders(v map[string]*string) *GetStackResponse {
	s.Headers = v
	return s
}

func (s *GetStackResponse) SetStatusCode(v int32) *GetStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackResponse) SetBody(v *GetStackResponseBody) *GetStackResponse {
	s.Body = v
	return s
}

type GetTraceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

type GetTraceResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spans     []*GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSpans(v []*GetTraceResponseBodySpans) *GetTraceResponseBody {
	s.Spans = v
	return s
}

type GetTraceResponseBodySpans struct {
	Duration      *int64                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                    `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  []*GetTraceResponseBodySpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	OperationName *string                                  `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                  `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                  `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType       *int32                                   `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServiceIp     *string                                  `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                  `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                  `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  []*GetTraceResponseBodySpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                  `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpans) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) SetDuration(v int64) *GetTraceResponseBodySpans {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetHaveStack(v bool) *GetTraceResponseBodySpans {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetLogEventList(v []*GetTraceResponseBodySpansLogEventList) *GetTraceResponseBodySpans {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetOperationName(v string) *GetTraceResponseBodySpans {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetParentSpanId(v string) *GetTraceResponseBodySpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetResultCode(v string) *GetTraceResponseBodySpans {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcId(v string) *GetTraceResponseBodySpans {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcType(v int32) *GetTraceResponseBodySpans {
	s.RpcType = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceIp(v string) *GetTraceResponseBodySpans {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceName(v string) *GetTraceResponseBodySpans {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetSpanId(v string) *GetTraceResponseBodySpans {
	s.SpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTagEntryList(v []*GetTraceResponseBodySpansTagEntryList) *GetTraceResponseBodySpans {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetTimestamp(v int64) *GetTraceResponseBodySpans {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTraceID(v string) *GetTraceResponseBodySpans {
	s.TraceID = &v
	return s
}

type GetTraceResponseBodySpansLogEventList struct {
	TagEntryList []*GetTraceResponseBodySpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp    *int64                                               `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventList) SetTagEntryList(v []*GetTraceResponseBodySpansLogEventListTagEntryList) *GetTraceResponseBodySpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansLogEventList) SetTimestamp(v int64) *GetTraceResponseBodySpansLogEventList {
	s.Timestamp = &v
	return s
}

type GetTraceResponseBodySpansLogEventListTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetKey(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetValue(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

type GetTraceResponseBodySpansTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansTagEntryList) SetKey(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansTagEntryList) SetValue(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Value = &v
	return s
}

type GetTraceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetStatusCode(v int32) *GetTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

type GetTraceAppRequest struct {
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAppRequest) SetPid(v string) *GetTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *GetTraceAppRequest) SetRegionId(v string) *GetTraceAppRequest {
	s.RegionId = &v
	return s
}

type GetTraceAppResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceApp  *GetTraceAppResponseBodyTraceApp `json:"TraceApp,omitempty" xml:"TraceApp,omitempty" type:"Struct"`
}

func (s GetTraceAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBody) SetRequestId(v string) *GetTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceAppResponseBody) SetTraceApp(v *GetTraceAppResponseBodyTraceApp) *GetTraceAppResponseBody {
	s.TraceApp = v
	return s
}

type GetTraceAppResponseBodyTraceApp struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTraceAppResponseBodyTraceApp) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponseBodyTraceApp) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppId(v int64) *GetTraceAppResponseBodyTraceApp {
	s.AppId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppName(v string) *GetTraceAppResponseBodyTraceApp {
	s.AppName = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetCreateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.CreateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetLabels(v []*string) *GetTraceAppResponseBodyTraceApp {
	s.Labels = v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetPid(v string) *GetTraceAppResponseBodyTraceApp {
	s.Pid = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetRegionId(v string) *GetTraceAppResponseBodyTraceApp {
	s.RegionId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetShow(v bool) *GetTraceAppResponseBodyTraceApp {
	s.Show = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetType(v string) *GetTraceAppResponseBodyTraceApp {
	s.Type = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUpdateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.UpdateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUserId(v string) *GetTraceAppResponseBodyTraceApp {
	s.UserId = &v
	return s
}

type GetTraceAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponse) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponse) SetHeaders(v map[string]*string) *GetTraceAppResponse {
	s.Headers = v
	return s
}

func (s *GetTraceAppResponse) SetStatusCode(v int32) *GetTraceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceAppResponse) SetBody(v *GetTraceAppResponseBody) *GetTraceAppResponse {
	s.Body = v
	return s
}

type ImportAppAlertRulesRequest struct {
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	Pids                *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	TemplateAlertId     *string `json:"TemplateAlertId,omitempty" xml:"TemplateAlertId,omitempty"`
}

func (s ImportAppAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequest) SetContactGroupIds(v string) *ImportAppAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetIsAutoStart(v bool) *ImportAppAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetPids(v string) *ImportAppAlertRulesRequest {
	s.Pids = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetRegionId(v string) *ImportAppAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplateAlertId(v string) *ImportAppAlertRulesRequest {
	s.TemplateAlertId = &v
	return s
}

type ImportAppAlertRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportAppAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponseBody) SetData(v string) *ImportAppAlertRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportAppAlertRulesResponseBody) SetRequestId(v string) *ImportAppAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type ImportAppAlertRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportAppAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportAppAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponse) SetHeaders(v map[string]*string) *ImportAppAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportAppAlertRulesResponse) SetStatusCode(v int32) *ImportAppAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAppAlertRulesResponse) SetBody(v *ImportAppAlertRulesResponseBody) *ImportAppAlertRulesResponse {
	s.Body = v
	return s
}

type ImportCustomAlertRulesRequest struct {
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	TemplateAlertConfig *string `json:"TemplateAlertConfig,omitempty" xml:"TemplateAlertConfig,omitempty"`
}

func (s ImportCustomAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportCustomAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesRequest) SetContactGroupIds(v string) *ImportCustomAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetIsAutoStart(v bool) *ImportCustomAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetRegionId(v string) *ImportCustomAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportCustomAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetTemplateAlertConfig(v string) *ImportCustomAlertRulesRequest {
	s.TemplateAlertConfig = &v
	return s
}

type ImportCustomAlertRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportCustomAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportCustomAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesResponseBody) SetData(v string) *ImportCustomAlertRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportCustomAlertRulesResponseBody) SetRequestId(v string) *ImportCustomAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type ImportCustomAlertRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportCustomAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportCustomAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportCustomAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesResponse) SetHeaders(v map[string]*string) *ImportCustomAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportCustomAlertRulesResponse) SetStatusCode(v int32) *ImportCustomAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCustomAlertRulesResponse) SetBody(v *ImportCustomAlertRulesResponseBody) *ImportCustomAlertRulesResponse {
	s.Body = v
	return s
}

type ImportPrometheusRulesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ImportPrometheusRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportPrometheusRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesRequest) SetClusterId(v string) *ImportPrometheusRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetContent(v string) *ImportPrometheusRulesRequest {
	s.Content = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetName(v string) *ImportPrometheusRulesRequest {
	s.Name = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetNameSpace(v string) *ImportPrometheusRulesRequest {
	s.NameSpace = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetRegionId(v string) *ImportPrometheusRulesRequest {
	s.RegionId = &v
	return s
}

type ImportPrometheusRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportPrometheusRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportPrometheusRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesResponseBody) SetData(v string) *ImportPrometheusRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportPrometheusRulesResponseBody) SetRequestId(v string) *ImportPrometheusRulesResponseBody {
	s.RequestId = &v
	return s
}

type ImportPrometheusRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportPrometheusRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportPrometheusRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportPrometheusRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesResponse) SetHeaders(v map[string]*string) *ImportPrometheusRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportPrometheusRulesResponse) SetStatusCode(v int32) *ImportPrometheusRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportPrometheusRulesResponse) SetBody(v *ImportPrometheusRulesResponseBody) *ImportPrometheusRulesResponse {
	s.Body = v
	return s
}

type ListActivatedAlertsRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActivatedAlertsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsRequest) SetCurrentPage(v int32) *ListActivatedAlertsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetFilter(v string) *ListActivatedAlertsRequest {
	s.Filter = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetPageSize(v int32) *ListActivatedAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetRegionId(v string) *ListActivatedAlertsRequest {
	s.RegionId = &v
	return s
}

type ListActivatedAlertsResponseBody struct {
	Page      *ListActivatedAlertsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListActivatedAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBody) SetPage(v *ListActivatedAlertsResponseBodyPage) *ListActivatedAlertsResponseBody {
	s.Page = v
	return s
}

func (s *ListActivatedAlertsResponseBody) SetRequestId(v string) *ListActivatedAlertsResponseBody {
	s.RequestId = &v
	return s
}

type ListActivatedAlertsResponseBodyPage struct {
	Alerts   []*ListActivatedAlertsResponseBodyPageAlerts `json:"Alerts,omitempty" xml:"Alerts,omitempty" type:"Repeated"`
	Page     *int32                                       `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPage) SetAlerts(v []*ListActivatedAlertsResponseBodyPageAlerts) *ListActivatedAlertsResponseBodyPage {
	s.Alerts = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetPage(v int32) *ListActivatedAlertsResponseBodyPage {
	s.Page = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetPageSize(v int32) *ListActivatedAlertsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetTotal(v int32) *ListActivatedAlertsResponseBodyPage {
	s.Total = &v
	return s
}

type ListActivatedAlertsResponseBodyPageAlerts struct {
	AlertId            *string                                                   `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName          *string                                                   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertType          *string                                                   `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Count              *int32                                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime         *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DispatchRules      []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	EndsAt             *int64                                                    `json:"EndsAt,omitempty" xml:"EndsAt,omitempty"`
	ExpandFields       map[string]interface{}                                    `json:"ExpandFields,omitempty" xml:"ExpandFields,omitempty"`
	IntegrationName    *string                                                   `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationType    *string                                                   `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	InvolvedObjectKind *string                                                   `json:"InvolvedObjectKind,omitempty" xml:"InvolvedObjectKind,omitempty"`
	InvolvedObjectName *string                                                   `json:"InvolvedObjectName,omitempty" xml:"InvolvedObjectName,omitempty"`
	Message            *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Severity           *string                                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartsAt           *int64                                                    `json:"StartsAt,omitempty" xml:"StartsAt,omitempty"`
	Status             *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPageAlerts) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPageAlerts) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertId(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertId = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertType(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertType = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetCount(v int32) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Count = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetCreateTime(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.CreateTime = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetDispatchRules(v []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules) *ListActivatedAlertsResponseBodyPageAlerts {
	s.DispatchRules = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetEndsAt(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.EndsAt = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetExpandFields(v map[string]interface{}) *ListActivatedAlertsResponseBodyPageAlerts {
	s.ExpandFields = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetIntegrationName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.IntegrationName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetIntegrationType(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.IntegrationType = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetInvolvedObjectKind(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.InvolvedObjectKind = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetInvolvedObjectName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.InvolvedObjectName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetMessage(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Message = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetSeverity(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Severity = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetStartsAt(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.StartsAt = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetStatus(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Status = &v
	return s
}

type ListActivatedAlertsResponseBodyPageAlertsDispatchRules struct {
	RuleId   *int32  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPageAlertsDispatchRules) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPageAlertsDispatchRules) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) SetRuleId(v int32) *ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	s.RuleId = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) SetRuleName(v string) *ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	s.RuleName = &v
	return s
}

type ListActivatedAlertsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListActivatedAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActivatedAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponse) SetHeaders(v map[string]*string) *ListActivatedAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListActivatedAlertsResponse) SetStatusCode(v int32) *ListActivatedAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActivatedAlertsResponse) SetBody(v *ListActivatedAlertsResponseBody) *ListActivatedAlertsResponse {
	s.Body = v
	return s
}

type ListAlertTemplatesRequest struct {
	AlertProvider    *string `json:"AlertProvider,omitempty" xml:"AlertProvider,omitempty"`
	Labels           *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateProvider *string `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesRequest) SetAlertProvider(v string) *ListAlertTemplatesRequest {
	s.AlertProvider = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetLabels(v string) *ListAlertTemplatesRequest {
	s.Labels = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetName(v string) *ListAlertTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetRegionId(v string) *ListAlertTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetStatus(v bool) *ListAlertTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetTemplateProvider(v string) *ListAlertTemplatesRequest {
	s.TemplateProvider = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetType(v string) *ListAlertTemplatesRequest {
	s.Type = &v
	return s
}

type ListAlertTemplatesResponseBody struct {
	AlertTemplates []*ListAlertTemplatesResponseBodyAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlertTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBody) SetAlertTemplates(v []*ListAlertTemplatesResponseBodyAlertTemplates) *ListAlertTemplatesResponseBody {
	s.AlertTemplates = v
	return s
}

func (s *ListAlertTemplatesResponseBody) SetRequestId(v string) *ListAlertTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListAlertTemplatesResponseBodyAlertTemplates struct {
	AlertProvider            *string                                                               `json:"AlertProvider,omitempty" xml:"AlertProvider,omitempty"`
	Annotations              map[string]interface{}                                                `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Id                       *int32                                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelMatchExpressionGrid *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Labels                   map[string]interface{}                                                `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message                  *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId                 *int32                                                                `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Public                   *bool                                                                 `json:"Public,omitempty" xml:"Public,omitempty"`
	Rule                     *string                                                               `json:"Rule,omitempty" xml:"Rule,omitempty"`
	Status                   *bool                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateProvider         *string                                                               `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	Type                     *string                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId                   *string                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplates) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetAlertProvider(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.AlertProvider = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetAnnotations(v map[string]interface{}) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Annotations = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetId(v int32) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Id = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetLabelMatchExpressionGrid(v *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetLabels(v map[string]interface{}) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Labels = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetMessage(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Message = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetName(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Name = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetParentId(v int32) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.ParentId = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetPublic(v bool) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Public = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetRule(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Rule = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetStatus(v bool) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Status = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetTemplateProvider(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.TemplateProvider = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetType(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Type = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetUserId(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.UserId = &v
	return s
}

type ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid struct {
	LabelMatchExpressionGroups []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

type ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	LabelMatchExpressions []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

type ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

type ListAlertTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlertTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlertTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlertTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponse) SetHeaders(v map[string]*string) *ListAlertTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertTemplatesResponse) SetStatusCode(v int32) *ListAlertTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertTemplatesResponse) SetBody(v *ListAlertTemplatesResponseBody) *ListAlertTemplatesResponse {
	s.Body = v
	return s
}

type ListClusterFromGrafanaRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClusterFromGrafanaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaRequest) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaRequest) SetRegionId(v string) *ListClusterFromGrafanaRequest {
	s.RegionId = &v
	return s
}

type ListClusterFromGrafanaResponseBody struct {
	PromClusterList []*ListClusterFromGrafanaResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterFromGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBody) SetPromClusterList(v []*ListClusterFromGrafanaResponseBodyPromClusterList) *ListClusterFromGrafanaResponseBody {
	s.PromClusterList = v
	return s
}

func (s *ListClusterFromGrafanaResponseBody) SetRequestId(v string) *ListClusterFromGrafanaResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterFromGrafanaResponseBodyPromClusterList struct {
	AgentStatus           *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType           *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ControllerId          *string `json:"ControllerId,omitempty" xml:"ControllerId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallTime           *int64  `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	IsControllerInstalled *bool   `json:"IsControllerInstalled,omitempty" xml:"IsControllerInstalled,omitempty"`
	LastHeartBeatTime     *int64  `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	NodeNum               *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	PluginsJsonArray      *string `json:"PluginsJsonArray,omitempty" xml:"PluginsJsonArray,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StateJson             *string `json:"StateJson,omitempty" xml:"StateJson,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetAgentStatus(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterName(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterType(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetControllerId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetCreateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetExtra(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetId(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetInstallTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetLastHeartBeatTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.LastHeartBeatTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetNodeNum(v int32) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.NodeNum = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetOptions(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetRegionId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetStateJson(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.StateJson = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUpdateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUserId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

type ListClusterFromGrafanaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterFromGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterFromGrafanaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponse) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponse) SetHeaders(v map[string]*string) *ListClusterFromGrafanaResponse {
	s.Headers = v
	return s
}

func (s *ListClusterFromGrafanaResponse) SetStatusCode(v int32) *ListClusterFromGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterFromGrafanaResponse) SetBody(v *ListClusterFromGrafanaResponseBody) *ListClusterFromGrafanaResponse {
	s.Body = v
	return s
}

type ListDashboardsRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType    *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RecreateSwitch *bool   `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDashboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsRequest) SetClusterId(v string) *ListDashboardsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsRequest) SetClusterType(v string) *ListDashboardsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDashboardsRequest) SetProduct(v string) *ListDashboardsRequest {
	s.Product = &v
	return s
}

func (s *ListDashboardsRequest) SetRecreateSwitch(v bool) *ListDashboardsRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *ListDashboardsRequest) SetRegionId(v string) *ListDashboardsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDashboardsRequest) SetTitle(v string) *ListDashboardsRequest {
	s.Title = &v
	return s
}

type ListDashboardsResponseBody struct {
	DashboardVos []*ListDashboardsResponseBodyDashboardVos `json:"DashboardVos,omitempty" xml:"DashboardVos,omitempty" type:"Repeated"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDashboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBody) SetDashboardVos(v []*ListDashboardsResponseBodyDashboardVos) *ListDashboardsResponseBody {
	s.DashboardVos = v
	return s
}

func (s *ListDashboardsResponseBody) SetRequestId(v string) *ListDashboardsResponseBody {
	s.RequestId = &v
	return s
}

type ListDashboardsResponseBodyDashboardVos struct {
	DashboardType  *string   `json:"DashboardType,omitempty" xml:"DashboardType,omitempty"`
	Exporter       *string   `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArmsExporter *bool     `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	Kind           *string   `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedUpdate     *bool     `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	Tags           []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Time           *string   `json:"Time,omitempty" xml:"Time,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid            *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Version        *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVos) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVos) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVos) SetDashboardType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.DashboardType = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetExporter(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetId(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetKind(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Kind = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetName(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Name = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetNeedUpdate(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.NeedUpdate = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVos {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTime(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTitle(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUid(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Uid = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetVersion(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Version = &v
	return s
}

type ListDashboardsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponse) SetHeaders(v map[string]*string) *ListDashboardsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardsResponse) SetStatusCode(v int32) *ListDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDashboardsResponse) SetBody(v *ListDashboardsResponseBody) *ListDashboardsResponse {
	s.Body = v
	return s
}

type ListDispatchRuleRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	System   *bool   `json:"System,omitempty" xml:"System,omitempty"`
}

func (s ListDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleRequest) SetName(v string) *ListDispatchRuleRequest {
	s.Name = &v
	return s
}

func (s *ListDispatchRuleRequest) SetRegionId(v string) *ListDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListDispatchRuleRequest) SetSystem(v bool) *ListDispatchRuleRequest {
	s.System = &v
	return s
}

type ListDispatchRuleResponseBody struct {
	DispatchRules []*ListDispatchRuleResponseBodyDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponseBody) SetDispatchRules(v []*ListDispatchRuleResponseBodyDispatchRules) *ListDispatchRuleResponseBody {
	s.DispatchRules = v
	return s
}

func (s *ListDispatchRuleResponseBody) SetRequestId(v string) *ListDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

type ListDispatchRuleResponseBodyDispatchRules struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	State  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDispatchRuleResponseBodyDispatchRules) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleResponseBodyDispatchRules) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetName(v string) *ListDispatchRuleResponseBodyDispatchRules {
	s.Name = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetRuleId(v int64) *ListDispatchRuleResponseBodyDispatchRules {
	s.RuleId = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetState(v string) *ListDispatchRuleResponseBodyDispatchRules {
	s.State = &v
	return s
}

type ListDispatchRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponse) SetHeaders(v map[string]*string) *ListDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *ListDispatchRuleResponse) SetStatusCode(v int32) *ListDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDispatchRuleResponse) SetBody(v *ListDispatchRuleResponseBody) *ListDispatchRuleResponse {
	s.Body = v
	return s
}

type ListPromClustersRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPromClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPromClustersRequest) GoString() string {
	return s.String()
}

func (s *ListPromClustersRequest) SetRegionId(v string) *ListPromClustersRequest {
	s.RegionId = &v
	return s
}

type ListPromClustersResponseBody struct {
	PromClusterList []*ListPromClustersResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPromClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPromClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponseBody) SetPromClusterList(v []*ListPromClustersResponseBodyPromClusterList) *ListPromClustersResponseBody {
	s.PromClusterList = v
	return s
}

func (s *ListPromClustersResponseBody) SetRequestId(v string) *ListPromClustersResponseBody {
	s.RequestId = &v
	return s
}

type ListPromClustersResponseBodyPromClusterList struct {
	AgentStatus           *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType           *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ControllerId          *string `json:"ControllerId,omitempty" xml:"ControllerId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallTime           *int64  `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	IsControllerInstalled *bool   `json:"IsControllerInstalled,omitempty" xml:"IsControllerInstalled,omitempty"`
	LastHeartBeatTime     *int64  `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	NodeNum               *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	PluginsJsonArray      *string `json:"PluginsJsonArray,omitempty" xml:"PluginsJsonArray,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StateJson             *string `json:"StateJson,omitempty" xml:"StateJson,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPromClustersResponseBodyPromClusterList) String() string {
	return tea.Prettify(s)
}

func (s ListPromClustersResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponseBodyPromClusterList) SetAgentStatus(v string) *ListPromClustersResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterName(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterType(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetControllerId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetCreateTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetExtra(v string) *ListPromClustersResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetId(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetInstallTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListPromClustersResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetLastHeartBeatTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.LastHeartBeatTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetNodeNum(v int32) *ListPromClustersResponseBodyPromClusterList {
	s.NodeNum = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetOptions(v string) *ListPromClustersResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListPromClustersResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetRegionId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetStateJson(v string) *ListPromClustersResponseBodyPromClusterList {
	s.StateJson = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetUpdateTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetUserId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

type ListPromClustersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPromClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPromClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPromClustersResponse) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponse) SetHeaders(v map[string]*string) *ListPromClustersResponse {
	s.Headers = v
	return s
}

func (s *ListPromClustersResponse) SetStatusCode(v int32) *ListPromClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromClustersResponse) SetBody(v *ListPromClustersResponseBody) *ListPromClustersResponse {
	s.Body = v
	return s
}

type ListPrometheusAlertRulesRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesRequest) SetClusterId(v string) *ListPrometheusAlertRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetMatchExpressions(v string) *ListPrometheusAlertRulesRequest {
	s.MatchExpressions = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetName(v string) *ListPrometheusAlertRulesRequest {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetRegionId(v string) *ListPrometheusAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetStatus(v int32) *ListPrometheusAlertRulesRequest {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetType(v string) *ListPrometheusAlertRulesRequest {
	s.Type = &v
	return s
}

type ListPrometheusAlertRulesResponseBody struct {
	PrometheusAlertRules []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules `json:"PrometheusAlertRules,omitempty" xml:"PrometheusAlertRules,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBody) SetPrometheusAlertRules(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) *ListPrometheusAlertRulesResponseBody {
	s.PrometheusAlertRules = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetRequestId(v string) *ListPrometheusAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRules struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAnnotations(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetClusterId(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDispatchRuleId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.DispatchRuleId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDuration(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetExpression(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetLabels(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetMessage(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Message = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetNotifyType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.NotifyType = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetStatus(v int32) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Type = &v
	return s
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Value = &v
	return s
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Value = &v
	return s
}

type ListPrometheusAlertRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrometheusAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrometheusAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponse) SetHeaders(v map[string]*string) *ListPrometheusAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusAlertRulesResponse) SetStatusCode(v int32) *ListPrometheusAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusAlertRulesResponse) SetBody(v *ListPrometheusAlertRulesResponseBody) *ListPrometheusAlertRulesResponse {
	s.Body = v
	return s
}

type ListPrometheusAlertTemplatesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPrometheusAlertTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesRequest) SetClusterId(v string) *ListPrometheusAlertTemplatesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertTemplatesRequest) SetRegionId(v string) *ListPrometheusAlertTemplatesRequest {
	s.RegionId = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBody struct {
	PrometheusAlertTemplates []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates `json:"PrometheusAlertTemplates,omitempty" xml:"PrometheusAlertTemplates,omitempty" type:"Repeated"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBody) SetPrometheusAlertTemplates(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) *ListPrometheusAlertTemplatesResponseBody {
	s.PrometheusAlertTemplates = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBody) SetRequestId(v string) *ListPrometheusAlertTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates struct {
	AlertName   *string                                                                        `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	Description *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *string                                                                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression  *string                                                                        `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels      []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Type        *string                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Version     *string                                                                        `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetAlertName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetAnnotations(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetDescription(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Description = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetDuration(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetExpression(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetLabels(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetType(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetVersion(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Version = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) SetName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) SetValue(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	s.Value = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) SetName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) SetValue(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	s.Value = &v
	return s
}

type ListPrometheusAlertTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrometheusAlertTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrometheusAlertTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponse) SetHeaders(v map[string]*string) *ListPrometheusAlertTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) SetStatusCode(v int32) *ListPrometheusAlertTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) SetBody(v *ListPrometheusAlertTemplatesResponseBody) *ListPrometheusAlertTemplatesResponse {
	s.Body = v
	return s
}

type ListRetcodeAppsRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListRetcodeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsRequest) SetRegionId(v string) *ListRetcodeAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRetcodeAppsRequest) SetSecurityToken(v string) *ListRetcodeAppsRequest {
	s.SecurityToken = &v
	return s
}

type ListRetcodeAppsResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetcodeApps []*ListRetcodeAppsResponseBodyRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
}

func (s ListRetcodeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBody) SetRequestId(v string) *ListRetcodeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRetcodeAppsResponseBody) SetRetcodeApps(v []*ListRetcodeAppsResponseBodyRetcodeApps) *ListRetcodeAppsResponseBody {
	s.RetcodeApps = v
	return s
}

type ListRetcodeAppsResponseBodyRetcodeApps struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Pid     *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppId(v int64) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppId = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppName(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppName = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetPid(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.Pid = &v
	return s
}

type ListRetcodeAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRetcodeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRetcodeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponse) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponse) SetHeaders(v map[string]*string) *ListRetcodeAppsResponse {
	s.Headers = v
	return s
}

func (s *ListRetcodeAppsResponse) SetStatusCode(v int32) *ListRetcodeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRetcodeAppsResponse) SetBody(v *ListRetcodeAppsResponseBody) *ListRetcodeAppsResponse {
	s.Body = v
	return s
}

type ListScenarioRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign     *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s ListScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioRequest) GoString() string {
	return s.String()
}

func (s *ListScenarioRequest) SetAppId(v string) *ListScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ListScenarioRequest) SetName(v string) *ListScenarioRequest {
	s.Name = &v
	return s
}

func (s *ListScenarioRequest) SetRegionId(v string) *ListScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ListScenarioRequest) SetScenario(v string) *ListScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ListScenarioRequest) SetSign(v string) *ListScenarioRequest {
	s.Sign = &v
	return s
}

type ListScenarioResponseBody struct {
	ArmsScenarios []*ListScenarioResponseBodyArmsScenarios `json:"ArmsScenarios,omitempty" xml:"ArmsScenarios,omitempty" type:"Repeated"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBody) SetArmsScenarios(v []*ListScenarioResponseBodyArmsScenarios) *ListScenarioResponseBody {
	s.ArmsScenarios = v
	return s
}

func (s *ListScenarioResponseBody) SetRequestId(v string) *ListScenarioResponseBody {
	s.RequestId = &v
	return s
}

type ListScenarioResponseBodyArmsScenarios struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sign       *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListScenarioResponseBodyArmsScenarios) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponseBodyArmsScenarios) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBodyArmsScenarios) SetAppId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.AppId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetCreateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.CreateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetExtensions(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Extensions = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetId(v int64) *ListScenarioResponseBodyArmsScenarios {
	s.Id = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetName(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Name = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetRegionId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.RegionId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetSign(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Sign = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUpdateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UpdateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUserId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UserId = &v
	return s
}

type ListScenarioResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScenarioResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponse) GoString() string {
	return s.String()
}

func (s *ListScenarioResponse) SetHeaders(v map[string]*string) *ListScenarioResponse {
	s.Headers = v
	return s
}

func (s *ListScenarioResponse) SetStatusCode(v int32) *ListScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScenarioResponse) SetBody(v *ListScenarioResponseBody) *ListScenarioResponse {
	s.Body = v
	return s
}

type ListServerlessTopNAppsRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderBy   *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListServerlessTopNAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerlessTopNAppsRequest) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsRequest) SetEndTime(v int64) *ListServerlessTopNAppsRequest {
	s.EndTime = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetLimit(v int32) *ListServerlessTopNAppsRequest {
	s.Limit = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetOrderBy(v string) *ListServerlessTopNAppsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetRegionId(v string) *ListServerlessTopNAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetStartTime(v int64) *ListServerlessTopNAppsRequest {
	s.StartTime = &v
	return s
}

type ListServerlessTopNAppsResponseBody struct {
	Data      []*ListServerlessTopNAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServerlessTopNAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerlessTopNAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsResponseBody) SetData(v []*ListServerlessTopNAppsResponseBodyData) *ListServerlessTopNAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListServerlessTopNAppsResponseBody) SetRequestId(v string) *ListServerlessTopNAppsResponseBody {
	s.RequestId = &v
	return s
}

type ListServerlessTopNAppsResponseBodyData struct {
	Count  *int32   `json:"Count,omitempty" xml:"Count,omitempty"`
	Error  *int32   `json:"Error,omitempty" xml:"Error,omitempty"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid    *string  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Region *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Rt     *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
}

func (s ListServerlessTopNAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServerlessTopNAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsResponseBodyData) SetCount(v int32) *ListServerlessTopNAppsResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetError(v int32) *ListServerlessTopNAppsResponseBodyData {
	s.Error = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetName(v string) *ListServerlessTopNAppsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetPid(v string) *ListServerlessTopNAppsResponseBodyData {
	s.Pid = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetRegion(v string) *ListServerlessTopNAppsResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetRt(v float32) *ListServerlessTopNAppsResponseBodyData {
	s.Rt = &v
	return s
}

type ListServerlessTopNAppsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServerlessTopNAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerlessTopNAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerlessTopNAppsResponse) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsResponse) SetHeaders(v map[string]*string) *ListServerlessTopNAppsResponse {
	s.Headers = v
	return s
}

func (s *ListServerlessTopNAppsResponse) SetStatusCode(v int32) *ListServerlessTopNAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerlessTopNAppsResponse) SetBody(v *ListServerlessTopNAppsResponseBody) *ListServerlessTopNAppsResponse {
	s.Body = v
	return s
}

type ListTraceAppsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTraceAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsRequest) GoString() string {
	return s.String()
}

func (s *ListTraceAppsRequest) SetRegionId(v string) *ListTraceAppsRequest {
	s.RegionId = &v
	return s
}

type ListTraceAppsResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceApps []*ListTraceAppsResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s ListTraceAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBody) SetCode(v int32) *ListTraceAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetMessage(v string) *ListTraceAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetRequestId(v string) *ListTraceAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetSuccess(v bool) *ListTraceAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody {
	s.TraceApps = v
	return s
}

type ListTraceAppsResponseBodyTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListTraceAppsResponseBodyTraceApps) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppId(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppName(v string) *ListTraceAppsResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetCreateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetLabels(v []*string) *ListTraceAppsResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetPid(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetRegionId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetShow(v bool) *ListTraceAppsResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetType(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUpdateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUserId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.UserId = &v
	return s
}

type ListTraceAppsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTraceAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTraceAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponse) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponse) SetHeaders(v map[string]*string) *ListTraceAppsResponse {
	s.Headers = v
	return s
}

func (s *ListTraceAppsResponse) SetStatusCode(v int32) *ListTraceAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTraceAppsResponse) SetBody(v *ListTraceAppsResponseBody) *ListTraceAppsResponse {
	s.Body = v
	return s
}

type OpenArmsDefaultSLRRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenArmsDefaultSLRRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsDefaultSLRRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRRequest) SetRegionId(v string) *OpenArmsDefaultSLRRequest {
	s.RegionId = &v
	return s
}

type OpenArmsDefaultSLRResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsDefaultSLRResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsDefaultSLRResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRResponseBody) SetData(v string) *OpenArmsDefaultSLRResponseBody {
	s.Data = &v
	return s
}

func (s *OpenArmsDefaultSLRResponseBody) SetRequestId(v string) *OpenArmsDefaultSLRResponseBody {
	s.RequestId = &v
	return s
}

type OpenArmsDefaultSLRResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenArmsDefaultSLRResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenArmsDefaultSLRResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsDefaultSLRResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRResponse) SetHeaders(v map[string]*string) *OpenArmsDefaultSLRResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsDefaultSLRResponse) SetStatusCode(v int32) *OpenArmsDefaultSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsDefaultSLRResponse) SetBody(v *OpenArmsDefaultSLRResponseBody) *OpenArmsDefaultSLRResponse {
	s.Body = v
	return s
}

type OpenArmsServiceRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OpenArmsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceRequest) SetOwnerId(v int64) *OpenArmsServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenArmsServiceRequest) SetType(v string) *OpenArmsServiceRequest {
	s.Type = &v
	return s
}

type OpenArmsServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceResponseBody) SetOrderId(v string) *OpenArmsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenArmsServiceResponseBody) SetRequestId(v string) *OpenArmsServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenArmsServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenArmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenArmsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceResponse) SetHeaders(v map[string]*string) *OpenArmsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsServiceResponse) SetStatusCode(v int32) *OpenArmsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsServiceResponse) SetBody(v *OpenArmsServiceResponseBody) *OpenArmsServiceResponse {
	s.Body = v
	return s
}

type OpenVClusterRequest struct {
	ClusterType    *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RecreateSwitch *bool   `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenVClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenVClusterRequest) GoString() string {
	return s.String()
}

func (s *OpenVClusterRequest) SetClusterType(v string) *OpenVClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *OpenVClusterRequest) SetLength(v int32) *OpenVClusterRequest {
	s.Length = &v
	return s
}

func (s *OpenVClusterRequest) SetProduct(v string) *OpenVClusterRequest {
	s.Product = &v
	return s
}

func (s *OpenVClusterRequest) SetRecreateSwitch(v bool) *OpenVClusterRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *OpenVClusterRequest) SetRegionId(v string) *OpenVClusterRequest {
	s.RegionId = &v
	return s
}

type OpenVClusterResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVClusterResponseBody) SetData(v string) *OpenVClusterResponseBody {
	s.Data = &v
	return s
}

func (s *OpenVClusterResponseBody) SetRequestId(v string) *OpenVClusterResponseBody {
	s.RequestId = &v
	return s
}

type OpenVClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenVClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenVClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenVClusterResponse) GoString() string {
	return s.String()
}

func (s *OpenVClusterResponse) SetHeaders(v map[string]*string) *OpenVClusterResponse {
	s.Headers = v
	return s
}

func (s *OpenVClusterResponse) SetStatusCode(v int32) *OpenVClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVClusterResponse) SetBody(v *OpenVClusterResponseBody) *OpenVClusterResponse {
	s.Body = v
	return s
}

type OpenXtraceDefaultSLRRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenXtraceDefaultSLRRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceDefaultSLRRequest) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRRequest) SetRegionId(v string) *OpenXtraceDefaultSLRRequest {
	s.RegionId = &v
	return s
}

type OpenXtraceDefaultSLRResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenXtraceDefaultSLRResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceDefaultSLRResponseBody) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRResponseBody) SetData(v string) *OpenXtraceDefaultSLRResponseBody {
	s.Data = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponseBody) SetRequestId(v string) *OpenXtraceDefaultSLRResponseBody {
	s.RequestId = &v
	return s
}

type OpenXtraceDefaultSLRResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenXtraceDefaultSLRResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenXtraceDefaultSLRResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceDefaultSLRResponse) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRResponse) SetHeaders(v map[string]*string) *OpenXtraceDefaultSLRResponse {
	s.Headers = v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) SetStatusCode(v int32) *OpenXtraceDefaultSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) SetBody(v *OpenXtraceDefaultSLRResponseBody) *OpenXtraceDefaultSLRResponse {
	s.Body = v
	return s
}

type QueryDatasetRequest struct {
	DatasetId     *int64                             `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DateStr       *string                            `json:"DateStr,omitempty" xml:"DateStr,omitempty"`
	Dimensions    []*QueryDatasetRequestDimensions   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	HungryMode    *bool                              `json:"HungryMode,omitempty" xml:"HungryMode,omitempty"`
	IntervalInSec *int32                             `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	IsDrillDown   *bool                              `json:"IsDrillDown,omitempty" xml:"IsDrillDown,omitempty"`
	Limit         *int32                             `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MaxTime       *int64                             `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	Measures      []*string                          `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	MinTime       *int64                             `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	OptionalDims  []*QueryDatasetRequestOptionalDims `json:"OptionalDims,omitempty" xml:"OptionalDims,omitempty" type:"Repeated"`
	OrderByKey    *string                            `json:"OrderByKey,omitempty" xml:"OrderByKey,omitempty"`
	ProxyUserId   *string                            `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ReduceTail    *bool                              `json:"ReduceTail,omitempty" xml:"ReduceTail,omitempty"`
	RequiredDims  []*QueryDatasetRequestRequiredDims `json:"RequiredDims,omitempty" xml:"RequiredDims,omitempty" type:"Repeated"`
}

func (s QueryDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequest) SetDatasetId(v int64) *QueryDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetRequest) SetDateStr(v string) *QueryDatasetRequest {
	s.DateStr = &v
	return s
}

func (s *QueryDatasetRequest) SetDimensions(v []*QueryDatasetRequestDimensions) *QueryDatasetRequest {
	s.Dimensions = v
	return s
}

func (s *QueryDatasetRequest) SetHungryMode(v bool) *QueryDatasetRequest {
	s.HungryMode = &v
	return s
}

func (s *QueryDatasetRequest) SetIntervalInSec(v int32) *QueryDatasetRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryDatasetRequest) SetIsDrillDown(v bool) *QueryDatasetRequest {
	s.IsDrillDown = &v
	return s
}

func (s *QueryDatasetRequest) SetLimit(v int32) *QueryDatasetRequest {
	s.Limit = &v
	return s
}

func (s *QueryDatasetRequest) SetMaxTime(v int64) *QueryDatasetRequest {
	s.MaxTime = &v
	return s
}

func (s *QueryDatasetRequest) SetMeasures(v []*string) *QueryDatasetRequest {
	s.Measures = v
	return s
}

func (s *QueryDatasetRequest) SetMinTime(v int64) *QueryDatasetRequest {
	s.MinTime = &v
	return s
}

func (s *QueryDatasetRequest) SetOptionalDims(v []*QueryDatasetRequestOptionalDims) *QueryDatasetRequest {
	s.OptionalDims = v
	return s
}

func (s *QueryDatasetRequest) SetOrderByKey(v string) *QueryDatasetRequest {
	s.OrderByKey = &v
	return s
}

func (s *QueryDatasetRequest) SetProxyUserId(v string) *QueryDatasetRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryDatasetRequest) SetReduceTail(v bool) *QueryDatasetRequest {
	s.ReduceTail = &v
	return s
}

func (s *QueryDatasetRequest) SetRequiredDims(v []*QueryDatasetRequestRequiredDims) *QueryDatasetRequest {
	s.RequiredDims = v
	return s
}

type QueryDatasetRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDatasetRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetRequestDimensions) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequestDimensions) SetKey(v string) *QueryDatasetRequestDimensions {
	s.Key = &v
	return s
}

func (s *QueryDatasetRequestDimensions) SetType(v string) *QueryDatasetRequestDimensions {
	s.Type = &v
	return s
}

func (s *QueryDatasetRequestDimensions) SetValue(v string) *QueryDatasetRequestDimensions {
	s.Value = &v
	return s
}

type QueryDatasetRequestOptionalDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDatasetRequestOptionalDims) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetRequestOptionalDims) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequestOptionalDims) SetKey(v string) *QueryDatasetRequestOptionalDims {
	s.Key = &v
	return s
}

func (s *QueryDatasetRequestOptionalDims) SetType(v string) *QueryDatasetRequestOptionalDims {
	s.Type = &v
	return s
}

func (s *QueryDatasetRequestOptionalDims) SetValue(v string) *QueryDatasetRequestOptionalDims {
	s.Value = &v
	return s
}

type QueryDatasetRequestRequiredDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDatasetRequestRequiredDims) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetRequestRequiredDims) GoString() string {
	return s.String()
}

func (s *QueryDatasetRequestRequiredDims) SetKey(v string) *QueryDatasetRequestRequiredDims {
	s.Key = &v
	return s
}

func (s *QueryDatasetRequestRequiredDims) SetType(v string) *QueryDatasetRequestRequiredDims {
	s.Type = &v
	return s
}

func (s *QueryDatasetRequestRequiredDims) SetValue(v string) *QueryDatasetRequestRequiredDims {
	s.Value = &v
	return s
}

type QueryDatasetResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetResponseBody) SetData(v string) *QueryDatasetResponseBody {
	s.Data = &v
	return s
}

func (s *QueryDatasetResponseBody) SetRequestId(v string) *QueryDatasetResponseBody {
	s.RequestId = &v
	return s
}

type QueryDatasetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetResponse) SetHeaders(v map[string]*string) *QueryDatasetResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetResponse) SetStatusCode(v int32) *QueryDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetResponse) SetBody(v *QueryDatasetResponseBody) *QueryDatasetResponse {
	s.Body = v
	return s
}

type QueryMetricRequest struct {
	ConsistencyDataKey       *string                      `json:"ConsistencyDataKey,omitempty" xml:"ConsistencyDataKey,omitempty"`
	ConsistencyQueryStrategy *string                      `json:"ConsistencyQueryStrategy,omitempty" xml:"ConsistencyQueryStrategy,omitempty"`
	Dimensions               []*string                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	EndTime                  *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters                  []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IntervalInSec            *int32                       `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	Limit                    *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Measures                 []*string                    `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Metric                   *string                      `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order                    *string                      `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy                  *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	ProxyUserId              *string                      `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	StartTime                *int64                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricRequest) SetConsistencyDataKey(v string) *QueryMetricRequest {
	s.ConsistencyDataKey = &v
	return s
}

func (s *QueryMetricRequest) SetConsistencyQueryStrategy(v string) *QueryMetricRequest {
	s.ConsistencyQueryStrategy = &v
	return s
}

func (s *QueryMetricRequest) SetDimensions(v []*string) *QueryMetricRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricRequest) SetEndTime(v int64) *QueryMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricRequest) SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricRequest) SetIntervalInSec(v int32) *QueryMetricRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricRequest) SetLimit(v int32) *QueryMetricRequest {
	s.Limit = &v
	return s
}

func (s *QueryMetricRequest) SetMeasures(v []*string) *QueryMetricRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricRequest) SetMetric(v string) *QueryMetricRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricRequest) SetOrder(v string) *QueryMetricRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricRequest) SetOrderBy(v string) *QueryMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricRequest) SetProxyUserId(v string) *QueryMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryMetricRequest) SetStartTime(v int64) *QueryMetricRequest {
	s.StartTime = &v
	return s
}

type QueryMetricRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricRequestFilters) SetKey(v string) *QueryMetricRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricRequestFilters) SetValue(v string) *QueryMetricRequestFilters {
	s.Value = &v
	return s
}

type QueryMetricResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricResponseBody) SetData(v string) *QueryMetricResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMetricResponseBody) SetRequestId(v string) *QueryMetricResponseBody {
	s.RequestId = &v
	return s
}

type QueryMetricResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricResponse) SetHeaders(v map[string]*string) *QueryMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricResponse) SetStatusCode(v int32) *QueryMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricResponse) SetBody(v *QueryMetricResponseBody) *QueryMetricResponse {
	s.Body = v
	return s
}

type QueryMetricByPageRequest struct {
	CurrentPage   *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CustomFilters []*string                          `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty" type:"Repeated"`
	Dimensions    []*string                          `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	EndTime       *int64                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters       []*QueryMetricByPageRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IntervalInSec *int32                             `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	Measures      []*string                          `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Metric        *string                            `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order         *string                            `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy       *string                            `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageSize      *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime     *int64                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequest) SetCurrentPage(v int32) *QueryMetricByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMetricByPageRequest) SetCustomFilters(v []*string) *QueryMetricByPageRequest {
	s.CustomFilters = v
	return s
}

func (s *QueryMetricByPageRequest) SetDimensions(v []*string) *QueryMetricByPageRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricByPageRequest) SetEndTime(v int64) *QueryMetricByPageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricByPageRequest) SetFilters(v []*QueryMetricByPageRequestFilters) *QueryMetricByPageRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricByPageRequest) SetIntervalInSec(v int32) *QueryMetricByPageRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricByPageRequest) SetMeasures(v []*string) *QueryMetricByPageRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricByPageRequest) SetMetric(v string) *QueryMetricByPageRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrder(v string) *QueryMetricByPageRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrderBy(v string) *QueryMetricByPageRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricByPageRequest) SetPageSize(v int32) *QueryMetricByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageRequest) SetStartTime(v int64) *QueryMetricByPageRequest {
	s.StartTime = &v
	return s
}

type QueryMetricByPageRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricByPageRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequestFilters) SetKey(v string) *QueryMetricByPageRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricByPageRequestFilters) SetValue(v string) *QueryMetricByPageRequestFilters {
	s.Value = &v
	return s
}

type QueryMetricByPageResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryMetricByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBody) SetCode(v string) *QueryMetricByPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryMetricByPageResponseBody) SetMessage(v string) *QueryMetricByPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetRequestId(v string) *QueryMetricByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetSuccess(v bool) *QueryMetricByPageResponseBody {
	s.Success = &v
	return s
}

type QueryMetricByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMetricByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBodyData) SetItems(v []map[string]interface{}) *QueryMetricByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPage(v int32) *QueryMetricByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPageSize(v int32) *QueryMetricByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetTotal(v int32) *QueryMetricByPageResponseBodyData {
	s.Total = &v
	return s
}

type QueryMetricByPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMetricByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMetricByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponse) SetHeaders(v map[string]*string) *QueryMetricByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricByPageResponse) SetStatusCode(v int32) *QueryMetricByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricByPageResponse) SetBody(v *QueryMetricByPageResponseBody) *QueryMetricByPageResponse {
	s.Body = v
	return s
}

type SaveTraceAppConfigRequest struct {
	Pid      *string                              `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Settings []*SaveTraceAppConfigRequestSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s SaveTraceAppConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigRequest) SetPid(v string) *SaveTraceAppConfigRequest {
	s.Pid = &v
	return s
}

func (s *SaveTraceAppConfigRequest) SetSettings(v []*SaveTraceAppConfigRequestSettings) *SaveTraceAppConfigRequest {
	s.Settings = v
	return s
}

type SaveTraceAppConfigRequestSettings struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveTraceAppConfigRequestSettings) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigRequestSettings) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigRequestSettings) SetKey(v string) *SaveTraceAppConfigRequestSettings {
	s.Key = &v
	return s
}

func (s *SaveTraceAppConfigRequestSettings) SetValue(v string) *SaveTraceAppConfigRequestSettings {
	s.Value = &v
	return s
}

type SaveTraceAppConfigResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveTraceAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponseBody) SetData(v string) *SaveTraceAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetRequestId(v string) *SaveTraceAppConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveTraceAppConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTraceAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTraceAppConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponse) SetHeaders(v map[string]*string) *SaveTraceAppConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveTraceAppConfigResponse) SetStatusCode(v int32) *SaveTraceAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTraceAppConfigResponse) SetBody(v *SaveTraceAppConfigResponseBody) *SaveTraceAppConfigResponse {
	s.Body = v
	return s
}

type SearchAlertContactRequest struct {
	ContactIds  *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactRequest) SetContactIds(v string) *SearchAlertContactRequest {
	s.ContactIds = &v
	return s
}

func (s *SearchAlertContactRequest) SetContactName(v string) *SearchAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactRequest) SetCurrentPage(v string) *SearchAlertContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertContactRequest) SetEmail(v string) *SearchAlertContactRequest {
	s.Email = &v
	return s
}

func (s *SearchAlertContactRequest) SetPageSize(v string) *SearchAlertContactRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactRequest) SetPhone(v string) *SearchAlertContactRequest {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactRequest) SetRegionId(v string) *SearchAlertContactRequest {
	s.RegionId = &v
	return s
}

type SearchAlertContactResponseBody struct {
	PageBean  *SearchAlertContactResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBody) SetPageBean(v *SearchAlertContactResponseBodyPageBean) *SearchAlertContactResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertContactResponseBody) SetRequestId(v string) *SearchAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertContactResponseBodyPageBean struct {
	Contacts   []*SearchAlertContactResponseBodyPageBeanContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBean) SetContacts(v []*SearchAlertContactResponseBodyPageBeanContacts) *SearchAlertContactResponseBodyPageBean {
	s.Contacts = v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertContactResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetPageSize(v int32) *SearchAlertContactResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertContactResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchAlertContactResponseBodyPageBeanContacts struct {
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Webhook     *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBeanContacts) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBeanContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactId(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactName(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetCreateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetDingRobot(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetEmail(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetPhone(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetSystemNoc(v bool) *SearchAlertContactResponseBodyPageBeanContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUpdateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUserId(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetWebhook(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Webhook = &v
	return s
}

type SearchAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponse) SetHeaders(v map[string]*string) *SearchAlertContactResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactResponse) SetStatusCode(v int32) *SearchAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactResponse) SetBody(v *SearchAlertContactResponseBody) *SearchAlertContactResponse {
	s.Body = v
	return s
}

type SearchAlertContactGroupRequest struct {
	ContactGroupIds  *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactId        *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName      *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	IsDetail         *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupRequest) SetContactGroupIds(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactGroupName(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactId(v int64) *SearchAlertContactGroupRequest {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactName(v string) *SearchAlertContactGroupRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetIsDetail(v bool) *SearchAlertContactGroupRequest {
	s.IsDetail = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetRegionId(v string) *SearchAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type SearchAlertContactGroupResponseBody struct {
	ContactGroups []*SearchAlertContactGroupResponseBodyContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBody) SetContactGroups(v []*SearchAlertContactGroupResponseBodyContactGroups) *SearchAlertContactGroupResponseBody {
	s.ContactGroups = v
	return s
}

func (s *SearchAlertContactGroupResponseBody) SetRequestId(v string) *SearchAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertContactGroupResponseBodyContactGroups struct {
	ContactGroupId   *int64                                                      `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string                                                     `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Contacts         []*SearchAlertContactGroupResponseBodyContactGroupsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	CreateTime       *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime       *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId           *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroups) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupId(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupName(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContacts(v []*SearchAlertContactGroupResponseBodyContactGroupsContacts) *SearchAlertContactGroupResponseBodyContactGroups {
	s.Contacts = v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UserId = &v
	return s
}

type SearchAlertContactGroupResponseBodyContactGroupsContacts struct {
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactId(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactName(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetDingRobot(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetEmail(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetPhone(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetSystemNoc(v bool) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UserId = &v
	return s
}

type SearchAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponse) SetHeaders(v map[string]*string) *SearchAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactGroupResponse) SetStatusCode(v int32) *SearchAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactGroupResponse) SetBody(v *SearchAlertContactGroupResponseBody) *SearchAlertContactGroupResponse {
	s.Body = v
	return s
}

type SearchAlertHistoriesRequest struct {
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesRequest) SetAlertId(v int64) *SearchAlertHistoriesRequest {
	s.AlertId = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetAlertType(v int32) *SearchAlertHistoriesRequest {
	s.AlertType = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetCurrentPage(v int32) *SearchAlertHistoriesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetEndTime(v int64) *SearchAlertHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetPageSize(v int32) *SearchAlertHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetRegionId(v string) *SearchAlertHistoriesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetStartTime(v int64) *SearchAlertHistoriesRequest {
	s.StartTime = &v
	return s
}

type SearchAlertHistoriesResponseBody struct {
	PageBean  *SearchAlertHistoriesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBody) SetPageBean(v *SearchAlertHistoriesResponseBodyPageBean) *SearchAlertHistoriesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertHistoriesResponseBody) SetRequestId(v string) *SearchAlertHistoriesResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertHistoriesResponseBodyPageBean struct {
	AlarmHistories []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	PageNumber     *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetAlarmHistories(v []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) *SearchAlertHistoriesResponseBodyPageBean {
	s.AlarmHistories = v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchAlertHistoriesResponseBodyPageBeanAlarmHistories struct {
	AlarmContent      *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	AlarmResponseCode *int32  `json:"AlarmResponseCode,omitempty" xml:"AlarmResponseCode,omitempty"`
	AlarmSources      *string `json:"AlarmSources,omitempty" xml:"AlarmSources,omitempty"`
	AlarmTime         *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AlarmType         *int32  `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	Emails            *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Phones            *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	StrategyId        *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Target            *string `json:"Target,omitempty" xml:"Target,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmContent(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmContent = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmResponseCode(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmResponseCode = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmSources(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmSources = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmTime(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmTime = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmType(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmType = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetEmails(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Emails = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetId(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Id = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetPhones(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Phones = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetStrategyId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.StrategyId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetTarget(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Target = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetUserId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.UserId = &v
	return s
}

type SearchAlertHistoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponse) SetHeaders(v map[string]*string) *SearchAlertHistoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertHistoriesResponse) SetStatusCode(v int32) *SearchAlertHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertHistoriesResponse) SetBody(v *SearchAlertHistoriesResponseBody) *SearchAlertHistoriesResponse {
	s.Body = v
	return s
}

type SearchAlertRulesRequest struct {
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesRequest) SetAppType(v string) *SearchAlertRulesRequest {
	s.AppType = &v
	return s
}

func (s *SearchAlertRulesRequest) SetCurrentPage(v int32) *SearchAlertRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPageSize(v int32) *SearchAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPid(v string) *SearchAlertRulesRequest {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesRequest) SetRegionId(v string) *SearchAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetTitle(v string) *SearchAlertRulesRequest {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesRequest) SetType(v string) *SearchAlertRulesRequest {
	s.Type = &v
	return s
}

type SearchAlertRulesResponseBody struct {
	PageBean  *SearchAlertRulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBody) SetPageBean(v *SearchAlertRulesResponseBodyPageBean) *SearchAlertRulesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertRulesResponseBody) SetRequestId(v string) *SearchAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertRulesResponseBodyPageBean struct {
	AlertRules []*SearchAlertRulesResponseBodyPageBeanAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBean) SetAlertRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRules) *SearchAlertRulesResponseBodyPageBean {
	s.AlertRules = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRules struct {
	AlarmContext       *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext `json:"AlarmContext,omitempty" xml:"AlarmContext,omitempty" type:"Struct"`
	AlertLevel         *string                                                     `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertRule          *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule    `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	AlertTitle         *string                                                     `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	AlertType          *int32                                                      `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertVersion       *int32                                                      `json:"AlertVersion,omitempty" xml:"AlertVersion,omitempty"`
	AlertWay           []*string                                                   `json:"AlertWay,omitempty" xml:"AlertWay,omitempty" type:"Repeated"`
	AlertWays          []*string                                                   `json:"AlertWays,omitempty" xml:"AlertWays,omitempty" type:"Repeated"`
	Config             *string                                                     `json:"Config,omitempty" xml:"Config,omitempty"`
	ContactGroupIdList *string                                                     `json:"ContactGroupIdList,omitempty" xml:"ContactGroupIdList,omitempty"`
	ContactGroupIds    *string                                                     `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	CreateTime         *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id                 *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	MetricParam        *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam  `json:"MetricParam,omitempty" xml:"MetricParam,omitempty" type:"Struct"`
	Notice             *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice       `json:"Notice,omitempty" xml:"Notice,omitempty" type:"Struct"`
	RegionId           *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId             *int64                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus         *string                                                     `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Title              *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime         *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId             *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlarmContext(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlarmContext = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertLevel(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertLevel = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertRule(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRule = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertVersion(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertVersion = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWay(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWay = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWays(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWays = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetConfig(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Config = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIdList(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIdList = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIds(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetCreateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Id = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetMetricParam(v *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricParam = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetNotice(v *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Notice = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Status = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskStatus = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUpdateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext struct {
	AlarmContentSubTitle *string `json:"AlarmContentSubTitle,omitempty" xml:"AlarmContentSubTitle,omitempty"`
	AlarmContentTemplate *string `json:"AlarmContentTemplate,omitempty" xml:"AlarmContentTemplate,omitempty"`
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SubTitle             *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentSubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentTemplate(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentTemplate = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetContent(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.Content = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.SubTitle = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule struct {
	Operator *string                                                         `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Rules    []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Rules = v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules struct {
	Aggregates *string  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	Alias      *string  `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Measure    *string  `json:"Measure,omitempty" xml:"Measure,omitempty"`
	NValue     *int32   `json:"NValue,omitempty" xml:"NValue,omitempty"`
	Operator   *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value      *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAggregates(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Aggregates = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAlias(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Alias = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetMeasure(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Measure = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetNValue(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.NValue = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetValue(v float32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Value = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam struct {
	AppGroupId *string                                                                `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	AppId      *string                                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Dimensions []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Pid        *string                                                                `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type       *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppGroupId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppGroupId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetDimensions(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Dimensions = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetPid(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Type = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetKey(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Key = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetValue(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Value = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesNotice struct {
	EndTime         *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NoticeEndTime   *int64 `json:"NoticeEndTime,omitempty" xml:"NoticeEndTime,omitempty"`
	NoticeStartTime *int64 `json:"NoticeStartTime,omitempty" xml:"NoticeStartTime,omitempty"`
	StartTime       *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.EndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeEndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeStartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.StartTime = &v
	return s
}

type SearchAlertRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponse) SetHeaders(v map[string]*string) *SearchAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertRulesResponse) SetStatusCode(v int32) *SearchAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertRulesResponse) SetBody(v *SearchAlertRulesResponseBody) *SearchAlertRulesResponse {
	s.Body = v
	return s
}

type SearchEventsRequest struct {
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsTrigger   *int32  `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsRequest) GoString() string {
	return s.String()
}

func (s *SearchEventsRequest) SetAlertId(v int64) *SearchEventsRequest {
	s.AlertId = &v
	return s
}

func (s *SearchEventsRequest) SetAlertType(v int32) *SearchEventsRequest {
	s.AlertType = &v
	return s
}

func (s *SearchEventsRequest) SetAppType(v string) *SearchEventsRequest {
	s.AppType = &v
	return s
}

func (s *SearchEventsRequest) SetCurrentPage(v int32) *SearchEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchEventsRequest) SetEndTime(v int64) *SearchEventsRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEventsRequest) SetIsTrigger(v int32) *SearchEventsRequest {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsRequest) SetPageSize(v int32) *SearchEventsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEventsRequest) SetPid(v string) *SearchEventsRequest {
	s.Pid = &v
	return s
}

func (s *SearchEventsRequest) SetRegionId(v string) *SearchEventsRequest {
	s.RegionId = &v
	return s
}

func (s *SearchEventsRequest) SetStartTime(v int64) *SearchEventsRequest {
	s.StartTime = &v
	return s
}

type SearchEventsResponseBody struct {
	IsTrigger *int32                            `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	PageBean  *SearchEventsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBody) SetIsTrigger(v int32) *SearchEventsResponseBody {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsResponseBody) SetPageBean(v *SearchEventsResponseBodyPageBean) *SearchEventsResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchEventsResponseBody) SetRequestId(v string) *SearchEventsResponseBody {
	s.RequestId = &v
	return s
}

type SearchEventsResponseBodyPageBean struct {
	Event      []*SearchEventsResponseBodyPageBeanEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchEventsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBean) SetEvent(v []*SearchEventsResponseBodyPageBeanEvent) *SearchEventsResponseBodyPageBean {
	s.Event = v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetPageNumber(v int32) *SearchEventsResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetPageSize(v int32) *SearchEventsResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetTotalCount(v int32) *SearchEventsResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchEventsResponseBodyPageBeanEvent struct {
	AlertId    *int64    `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName  *string   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertRule  *string   `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	AlertType  *int32    `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	EventLevel *string   `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventTime  *int64    `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Id         *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Links      []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	Message    *string   `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SearchEventsResponseBodyPageBeanEvent) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBodyPageBeanEvent) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertId = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertName(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertName = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertRule(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertRule = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertType(v int32) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertType = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventLevel(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.EventLevel = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventTime(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.EventTime = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.Id = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetLinks(v []*string) *SearchEventsResponseBodyPageBeanEvent {
	s.Links = v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetMessage(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.Message = &v
	return s
}

type SearchEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponse) GoString() string {
	return s.String()
}

func (s *SearchEventsResponse) SetHeaders(v map[string]*string) *SearchEventsResponse {
	s.Headers = v
	return s
}

func (s *SearchEventsResponse) SetStatusCode(v int32) *SearchEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchEventsResponse) SetBody(v *SearchEventsResponseBody) *SearchEventsResponse {
	s.Body = v
	return s
}

type SearchRetcodeAppByPageRequest struct {
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
}

func (s SearchRetcodeAppByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageRequest) SetPageNumber(v int32) *SearchRetcodeAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetPageSize(v int32) *SearchRetcodeAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRegionId(v string) *SearchRetcodeAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest {
	s.RetcodeAppName = &v
	return s
}

type SearchRetcodeAppByPageResponseBody struct {
	PageBean  *SearchRetcodeAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBody) SetPageBean(v *SearchRetcodeAppByPageResponseBodyPageBean) *SearchRetcodeAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBody) SetRequestId(v string) *SearchRetcodeAppByPageResponseBody {
	s.RequestId = &v
	return s
}

type SearchRetcodeAppByPageResponseBodyPageBean struct {
	PageNumber  *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RetcodeApps []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
	TotalCount  *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetRetcodeApps(v []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.RetcodeApps = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps struct {
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid        *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppId(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppName(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppName = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetCreateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.CreateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetPid(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Pid = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRegionId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Type = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUpdateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUserId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UserId = &v
	return s
}

type SearchRetcodeAppByPageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchRetcodeAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchRetcodeAppByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponse) SetHeaders(v map[string]*string) *SearchRetcodeAppByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchRetcodeAppByPageResponse) SetStatusCode(v int32) *SearchRetcodeAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchRetcodeAppByPageResponse) SetBody(v *SearchRetcodeAppByPageResponseBody) *SearchRetcodeAppByPageResponse {
	s.Body = v
	return s
}

type SearchTraceAppByNameRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameRequest) SetRegionId(v string) *SearchTraceAppByNameRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameRequest) SetTraceAppName(v string) *SearchTraceAppByNameRequest {
	s.TraceAppName = &v
	return s
}

type SearchTraceAppByNameResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceApps []*SearchTraceAppByNameResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBody) SetRequestId(v string) *SearchTraceAppByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBody) SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody {
	s.TraceApps = v
	return s
}

type SearchTraceAppByNameResponseBodyTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByNameResponseBodyTraceApps) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppId(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppName(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetCreateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetLabels(v []*string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetPid(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetRegionId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetShow(v bool) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetType(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUpdateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUserId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UserId = &v
	return s
}

type SearchTraceAppByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTraceAppByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTraceAppByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponse) SetHeaders(v map[string]*string) *SearchTraceAppByNameResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceAppByNameResponse) SetStatusCode(v int32) *SearchTraceAppByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceAppByNameResponse) SetBody(v *SearchTraceAppByNameResponseBody) *SearchTraceAppByNameResponse {
	s.Body = v
	return s
}

type SearchTraceAppByPageRequest struct {
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageRequest) SetPageNumber(v int32) *SearchTraceAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetPageSize(v int32) *SearchTraceAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetRegionId(v string) *SearchTraceAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetTraceAppName(v string) *SearchTraceAppByPageRequest {
	s.TraceAppName = &v
	return s
}

type SearchTraceAppByPageResponseBody struct {
	PageBean  *SearchTraceAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTraceAppByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBody) SetPageBean(v *SearchTraceAppByPageResponseBodyPageBean) *SearchTraceAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTraceAppByPageResponseBody) SetRequestId(v string) *SearchTraceAppByPageResponseBody {
	s.RequestId = &v
	return s
}

type SearchTraceAppByPageResponseBodyPageBean struct {
	PageNumber *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TraceApps  []*SearchTraceAppByPageResponseBodyPageBeanTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTraceApps(v []*SearchTraceAppByPageResponseBodyPageBeanTraceApps) *SearchTraceAppByPageResponseBodyPageBean {
	s.TraceApps = v
	return s
}

type SearchTraceAppByPageResponseBodyPageBeanTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppId(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppName(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetCreateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetLabels(v []*string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetPid(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetRegionId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetShow(v bool) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetType(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUpdateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUserId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UserId = &v
	return s
}

type SearchTraceAppByPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTraceAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTraceAppByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponse) SetHeaders(v map[string]*string) *SearchTraceAppByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceAppByPageResponse) SetStatusCode(v int32) *SearchTraceAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceAppByPageResponse) SetBody(v *SearchTraceAppByPageResponseBody) *SearchTraceAppByPageResponse {
	s.Body = v
	return s
}

type SearchTracesRequest struct {
	EndTime          *int64                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExclusionFilters []*SearchTracesRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	MinDuration      *int64                                 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName    *string                                `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	RegionId         *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse          *bool                                  `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp        *string                                `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName      *string                                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime        *int64                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag              []*SearchTracesRequestTag              `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

type SearchTracesRequestExclusionFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestExclusionFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestExclusionFilters) SetKey(v string) *SearchTracesRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestExclusionFilters) SetValue(v string) *SearchTracesRequestExclusionFilters {
	s.Value = &v
	return s
}

type SearchTracesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

type SearchTracesResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceInfos []*SearchTracesResponseBodyTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBody) SetRequestId(v string) *SearchTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTracesResponseBody) SetTraceInfos(v []*SearchTracesResponseBodyTraceInfos) *SearchTracesResponseBody {
	s.TraceInfos = v
	return s
}

type SearchTracesResponseBodyTraceInfos struct {
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyTraceInfos) SetDuration(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetOperationName(v string) *SearchTracesResponseBodyTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceIp(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceName(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTimestamp(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTraceID(v string) *SearchTracesResponseBodyTraceInfos {
	s.TraceID = &v
	return s
}

type SearchTracesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) SetHeaders(v map[string]*string) *SearchTracesResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesResponse) SetStatusCode(v int32) *SearchTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
	s.Body = v
	return s
}

type SearchTracesByPageRequest struct {
	EndTime          *int64                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExclusionFilters []*SearchTracesByPageRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	MinDuration      *int64                                       `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName    *string                                      `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId         *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse          *bool                                        `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp        *string                                      `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName      *string                                      `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime        *int64                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchTracesByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequest) SetEndTime(v int64) *SearchTracesByPageRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesByPageRequest) SetMinDuration(v int64) *SearchTracesByPageRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOperationName(v string) *SearchTracesByPageRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageNumber(v int32) *SearchTracesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageSize(v int32) *SearchTracesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageRequest) SetRegionId(v string) *SearchTracesByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesByPageRequest) SetReverse(v bool) *SearchTracesByPageRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceIp(v string) *SearchTracesByPageRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceName(v string) *SearchTracesByPageRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetStartTime(v int64) *SearchTracesByPageRequest {
	s.StartTime = &v
	return s
}

type SearchTracesByPageRequestExclusionFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesByPageRequestExclusionFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequestExclusionFilters) SetKey(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesByPageRequestExclusionFilters) SetValue(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Value = &v
	return s
}

type SearchTracesByPageResponseBody struct {
	PageBean  *SearchTracesByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBody) SetPageBean(v *SearchTracesByPageResponseBodyPageBean) *SearchTracesByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTracesByPageResponseBody) SetRequestId(v string) *SearchTracesByPageResponseBody {
	s.RequestId = &v
	return s
}

type SearchTracesByPageResponseBodyPageBean struct {
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	TraceInfos []*SearchTracesByPageResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTotal(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTraceInfos(v []*SearchTracesByPageResponseBodyPageBeanTraceInfos) *SearchTracesByPageResponseBodyPageBean {
	s.TraceInfos = v
	return s
}

type SearchTracesByPageResponseBodyPageBeanTraceInfos struct {
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetDuration(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetOperationName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceIp(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTimestamp(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTraceID(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.TraceID = &v
	return s
}

type SearchTracesByPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTracesByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTracesByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponse) SetHeaders(v map[string]*string) *SearchTracesByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesByPageResponse) SetStatusCode(v int32) *SearchTracesByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesByPageResponse) SetBody(v *SearchTracesByPageResponseBody) *SearchTracesByPageResponse {
	s.Body = v
	return s
}

type SendCustomIncidentsRequest struct {
	Incidents   *string `json:"Incidents,omitempty" xml:"Incidents,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SendCustomIncidentsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomIncidentsRequest) GoString() string {
	return s.String()
}

func (s *SendCustomIncidentsRequest) SetIncidents(v string) *SendCustomIncidentsRequest {
	s.Incidents = &v
	return s
}

func (s *SendCustomIncidentsRequest) SetProductType(v string) *SendCustomIncidentsRequest {
	s.ProductType = &v
	return s
}

func (s *SendCustomIncidentsRequest) SetRegionId(v string) *SendCustomIncidentsRequest {
	s.RegionId = &v
	return s
}

type SendCustomIncidentsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCustomIncidentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomIncidentsResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomIncidentsResponseBody) SetRequestId(v string) *SendCustomIncidentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomIncidentsResponseBody) SetSuccess(v bool) *SendCustomIncidentsResponseBody {
	s.Success = &v
	return s
}

type SendCustomIncidentsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendCustomIncidentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomIncidentsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomIncidentsResponse) GoString() string {
	return s.String()
}

func (s *SendCustomIncidentsResponse) SetHeaders(v map[string]*string) *SendCustomIncidentsResponse {
	s.Headers = v
	return s
}

func (s *SendCustomIncidentsResponse) SetStatusCode(v int32) *SendCustomIncidentsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomIncidentsResponse) SetBody(v *SendCustomIncidentsResponseBody) *SendCustomIncidentsResponse {
	s.Body = v
	return s
}

type SendMseIncidentRequest struct {
	Incidents *string `json:"Incidents,omitempty" xml:"Incidents,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SendMseIncidentRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMseIncidentRequest) GoString() string {
	return s.String()
}

func (s *SendMseIncidentRequest) SetIncidents(v string) *SendMseIncidentRequest {
	s.Incidents = &v
	return s
}

func (s *SendMseIncidentRequest) SetRegionId(v string) *SendMseIncidentRequest {
	s.RegionId = &v
	return s
}

type SendMseIncidentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMseIncidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMseIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *SendMseIncidentResponseBody) SetRequestId(v string) *SendMseIncidentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMseIncidentResponseBody) SetSuccess(v bool) *SendMseIncidentResponseBody {
	s.Success = &v
	return s
}

type SendMseIncidentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMseIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMseIncidentResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMseIncidentResponse) GoString() string {
	return s.String()
}

func (s *SendMseIncidentResponse) SetHeaders(v map[string]*string) *SendMseIncidentResponse {
	s.Headers = v
	return s
}

func (s *SendMseIncidentResponse) SetStatusCode(v int32) *SendMseIncidentResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMseIncidentResponse) SetBody(v *SendMseIncidentResponseBody) *SendMseIncidentResponse {
	s.Body = v
	return s
}

type SetRetcodeShareStatusRequest struct {
	Pid    *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Status *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetRetcodeShareStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRetcodeShareStatusRequest) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusRequest) SetPid(v string) *SetRetcodeShareStatusRequest {
	s.Pid = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) SetStatus(v bool) *SetRetcodeShareStatusRequest {
	s.Status = &v
	return s
}

type SetRetcodeShareStatusResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRetcodeShareStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRetcodeShareStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusResponseBody) SetIsSuccess(v bool) *SetRetcodeShareStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SetRetcodeShareStatusResponseBody) SetRequestId(v string) *SetRetcodeShareStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetRetcodeShareStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetRetcodeShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRetcodeShareStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRetcodeShareStatusResponse) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusResponse) SetHeaders(v map[string]*string) *SetRetcodeShareStatusResponse {
	s.Headers = v
	return s
}

func (s *SetRetcodeShareStatusResponse) SetStatusCode(v int32) *SetRetcodeShareStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRetcodeShareStatusResponse) SetBody(v *SetRetcodeShareStatusResponseBody) *SetRetcodeShareStatusResponse {
	s.Body = v
	return s
}

type StartAlertRequest struct {
	AlertId  *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAlertRequest) GoString() string {
	return s.String()
}

func (s *StartAlertRequest) SetAlertId(v string) *StartAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StartAlertRequest) SetRegionId(v string) *StartAlertRequest {
	s.RegionId = &v
	return s
}

type StartAlertResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StartAlertResponseBody) SetIsSuccess(v bool) *StartAlertResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *StartAlertResponseBody) SetRequestId(v string) *StartAlertResponseBody {
	s.RequestId = &v
	return s
}

type StartAlertResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAlertResponse) GoString() string {
	return s.String()
}

func (s *StartAlertResponse) SetHeaders(v map[string]*string) *StartAlertResponse {
	s.Headers = v
	return s
}

func (s *StartAlertResponse) SetStatusCode(v int32) *StartAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAlertResponse) SetBody(v *StartAlertResponseBody) *StartAlertResponse {
	s.Body = v
	return s
}

type StopAlertRequest struct {
	AlertId  *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAlertRequest) GoString() string {
	return s.String()
}

func (s *StopAlertRequest) SetAlertId(v string) *StopAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StopAlertRequest) SetRegionId(v string) *StopAlertRequest {
	s.RegionId = &v
	return s
}

type StopAlertResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StopAlertResponseBody) SetIsSuccess(v bool) *StopAlertResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *StopAlertResponseBody) SetRequestId(v string) *StopAlertResponseBody {
	s.RequestId = &v
	return s
}

type StopAlertResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAlertResponse) GoString() string {
	return s.String()
}

func (s *StopAlertResponse) SetHeaders(v map[string]*string) *StopAlertResponse {
	s.Headers = v
	return s
}

func (s *StopAlertResponse) SetStatusCode(v int32) *StopAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAlertResponse) SetBody(v *StopAlertResponseBody) *StopAlertResponse {
	s.Body = v
	return s
}

type UpdateAlertContactRequest struct {
	ContactId           *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
}

func (s UpdateAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactRequest) SetContactId(v int64) *UpdateAlertContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetContactName(v string) *UpdateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateAlertContactRequest) SetDingRobotWebhookUrl(v string) *UpdateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *UpdateAlertContactRequest) SetEmail(v string) *UpdateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateAlertContactRequest) SetPhoneNum(v string) *UpdateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *UpdateAlertContactRequest) SetRegionId(v string) *UpdateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetSystemNoc(v bool) *UpdateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

type UpdateAlertContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactResponseBody) SetIsSuccess(v bool) *UpdateAlertContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateAlertContactResponseBody) SetRequestId(v string) *UpdateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactResponse) SetHeaders(v map[string]*string) *UpdateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertContactResponse) SetStatusCode(v int32) *UpdateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertContactResponse) SetBody(v *UpdateAlertContactResponseBody) *UpdateAlertContactResponse {
	s.Body = v
	return s
}

type UpdateAlertContactGroupRequest struct {
	ContactGroupId   *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupId(v int64) *UpdateAlertContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupName(v string) *UpdateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactIds(v string) *UpdateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetRegionId(v string) *UpdateAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type UpdateAlertContactGroupResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupResponseBody) SetIsSuccess(v bool) *UpdateAlertContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateAlertContactGroupResponseBody) SetRequestId(v string) *UpdateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupResponse) SetHeaders(v map[string]*string) *UpdateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertContactGroupResponse) SetStatusCode(v int32) *UpdateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertContactGroupResponse) SetBody(v *UpdateAlertContactGroupResponseBody) *UpdateAlertContactGroupResponse {
	s.Body = v
	return s
}

type UpdateAlertRuleRequest struct {
	AlertId             *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
}

func (s UpdateAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequest) SetAlertId(v int64) *UpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetContactGroupIds(v string) *UpdateAlertRuleRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetIsAutoStart(v bool) *UpdateAlertRuleRequest {
	s.IsAutoStart = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetRegionId(v string) *UpdateAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest {
	s.TemplageAlertConfig = &v
	return s
}

type UpdateAlertRuleResponseBody struct {
	AlertId   *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponseBody) SetAlertId(v int64) *UpdateAlertRuleResponseBody {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetData(v string) *UpdateAlertRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetRequestId(v string) *UpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlertRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponse) SetHeaders(v map[string]*string) *UpdateAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertRuleResponse) SetStatusCode(v int32) *UpdateAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertRuleResponse) SetBody(v *UpdateAlertRuleResponseBody) *UpdateAlertRuleResponse {
	s.Body = v
	return s
}

type UpdateAlertTemplateRequest struct {
	Annotations      *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Labels           *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Rule             *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	Status           *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAlertTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertTemplateRequest) SetAnnotations(v string) *UpdateAlertTemplateRequest {
	s.Annotations = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetId(v int64) *UpdateAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetLabels(v string) *UpdateAlertTemplateRequest {
	s.Labels = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetMatchExpressions(v string) *UpdateAlertTemplateRequest {
	s.MatchExpressions = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetMessage(v string) *UpdateAlertTemplateRequest {
	s.Message = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetName(v string) *UpdateAlertTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetRegionId(v string) *UpdateAlertTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetRule(v string) *UpdateAlertTemplateRequest {
	s.Rule = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetStatus(v bool) *UpdateAlertTemplateRequest {
	s.Status = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetType(v string) *UpdateAlertTemplateRequest {
	s.Type = &v
	return s
}

type UpdateAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAlertTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertTemplateResponseBody) SetRequestId(v string) *UpdateAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertTemplateResponseBody) SetSuccess(v bool) *UpdateAlertTemplateResponseBody {
	s.Success = &v
	return s
}

type UpdateAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertTemplateResponse) SetHeaders(v map[string]*string) *UpdateAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertTemplateResponse) SetStatusCode(v int32) *UpdateAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertTemplateResponse) SetBody(v *UpdateAlertTemplateResponseBody) *UpdateAlertTemplateResponse {
	s.Body = v
	return s
}

type UpdateDispatchRuleRequest struct {
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleRequest) SetDispatchRule(v string) *UpdateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *UpdateDispatchRuleRequest) SetRegionId(v string) *UpdateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type UpdateDispatchRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleResponseBody) SetRequestId(v string) *UpdateDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDispatchRuleResponseBody) SetSuccess(v bool) *UpdateDispatchRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleResponse) SetHeaders(v map[string]*string) *UpdateDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDispatchRuleResponse) SetStatusCode(v int32) *UpdateDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDispatchRuleResponse) SetBody(v *UpdateDispatchRuleResponseBody) *UpdateDispatchRuleResponse {
	s.Body = v
	return s
}

type UpdatePrometheusAlertRuleRequest struct {
	AlertId        *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleRequest) SetAlertId(v int64) *UpdatePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetAlertName(v string) *UpdatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetAnnotations(v string) *UpdatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetClusterId(v string) *UpdatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetDuration(v string) *UpdatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetExpression(v string) *UpdatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetLabels(v string) *UpdatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetMessage(v string) *UpdatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetNotifyType(v string) *UpdatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetRegionId(v string) *UpdatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetType(v string) *UpdatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *UpdatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetRequestId(v string) *UpdatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

type UpdatePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *UpdatePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) SetStatusCode(v int32) *UpdatePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) SetBody(v *UpdatePrometheusAlertRuleResponseBody) *UpdatePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type UpdateWebhookRequest struct {
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebhookRequest) SetBody(v string) *UpdateWebhookRequest {
	s.Body = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactId(v int64) *UpdateWebhookRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactName(v string) *UpdateWebhookRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpHeaders(v string) *UpdateWebhookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpParams(v string) *UpdateWebhookRequest {
	s.HttpParams = &v
	return s
}

func (s *UpdateWebhookRequest) SetMethod(v string) *UpdateWebhookRequest {
	s.Method = &v
	return s
}

func (s *UpdateWebhookRequest) SetRegionId(v string) *UpdateWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWebhookRequest) SetUrl(v string) *UpdateWebhookRequest {
	s.Url = &v
	return s
}

type UpdateWebhookResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponseBody) SetIsSuccess(v bool) *UpdateWebhookResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetRequestId(v string) *UpdateWebhookResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponse) SetHeaders(v map[string]*string) *UpdateWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebhookResponse) SetStatusCode(v int32) *UpdateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebhookResponse) SetBody(v *UpdateWebhookResponseBody) *UpdateWebhookResponse {
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
		"ap-northeast-2-pop":          tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("arms.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("arms.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("arms.aliyuncs.com"),
		"cn-edge-1":                   tea.String("arms.aliyuncs.com"),
		"cn-fujian":                   tea.String("arms.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("arms.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("arms.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("arms.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("arms.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("arms.aliyuncs.com"),
		"cn-wuhan":                    tea.String("arms.aliyuncs.com"),
		"cn-yushanfang":               tea.String("arms.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("arms.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("arms.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("arms.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("arms.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("arms.aliyuncs.com"),
		"me-east-1":                   tea.String("arms.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("arms.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("arms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddGrafanaWithOptions(request *AddGrafanaRequest, runtime *util.RuntimeOptions) (_result *AddGrafanaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Integration)) {
		query["Integration"] = request.Integration
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGrafana"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGrafanaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGrafana(request *AddGrafanaRequest) (_result *AddGrafanaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGrafanaResponse{}
	_body, _err := client.AddGrafanaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddIntegrationWithOptions(request *AddIntegrationRequest, runtime *util.RuntimeOptions) (_result *AddIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Integration)) {
		query["Integration"] = request.Integration
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIntegration"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIntegration(request *AddIntegrationRequest) (_result *AddIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIntegrationResponse{}
	_body, _err := client.AddIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyScenarioWithOptions(tmpReq *ApplyScenarioRequest, runtime *util.RuntimeOptions) (_result *ApplyScenarioResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyScenarioShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Config)) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, tea.String("Config"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigShrink)) {
		query["Config"] = request.ConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		query["Scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.SnDump)) {
		query["SnDump"] = request.SnDump
	}

	if !tea.BoolValue(util.IsUnset(request.SnForce)) {
		query["SnForce"] = request.SnForce
	}

	if !tea.BoolValue(util.IsUnset(request.SnStat)) {
		query["SnStat"] = request.SnStat
	}

	if !tea.BoolValue(util.IsUnset(request.SnTransfer)) {
		query["SnTransfer"] = request.SnTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateOption)) {
		query["UpdateOption"] = request.UpdateOption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyScenario"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyScenario(request *ApplyScenarioRequest) (_result *ApplyScenarioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyScenarioResponse{}
	_body, _err := client.ApplyScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDataConsistencyWithOptions(request *CheckDataConsistencyRequest, runtime *util.RuntimeOptions) (_result *CheckDataConsistencyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentTimestamp)) {
		query["CurrentTimestamp"] = request.CurrentTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUserId)) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDataConsistency"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDataConsistencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDataConsistency(request *CheckDataConsistencyRequest) (_result *CheckDataConsistencyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDataConsistencyResponse{}
	_body, _err := client.CheckDataConsistencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleForDeletingWithOptions(request *CheckServiceLinkedRoleForDeletingRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionTaskId)) {
		query["DeletionTaskId"] = request.DeletionTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.SPIRegionId)) {
		query["SPIRegionId"] = request.SPIRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRoleForDeleting"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleForDeleting(request *CheckServiceLinkedRoleForDeletingRequest) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.CheckServiceLinkedRoleForDeletingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceStatusWithOptions(request *CheckServiceStatusRequest, runtime *util.RuntimeOptions) (_result *CheckServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SvcCode)) {
		query["SvcCode"] = request.SvcCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceStatus"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceStatus(request *CheckServiceStatusRequest) (_result *CheckServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceStatusResponse{}
	_body, _err := client.CheckServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigAppWithOptions(request *ConfigAppRequest, runtime *util.RuntimeOptions) (_result *ConfigAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigApp"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigApp(request *ConfigAppRequest) (_result *ConfigAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigAppResponse{}
	_body, _err := client.ConfigAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertContactWithOptions(request *CreateAlertContactRequest, runtime *util.RuntimeOptions) (_result *CreateAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.DingRobotWebhookUrl)) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemNoc)) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertContact"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertContact(request *CreateAlertContactRequest) (_result *CreateAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CreateAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertContactGroupWithOptions(request *CreateAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertContactGroup"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertContactGroup(request *CreateAlertContactGroupRequest) (_result *CreateAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CreateAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertTemplateWithOptions(request *CreateAlertTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateAlertTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.MatchExpressions)) {
		query["MatchExpressions"] = request.MatchExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		query["Rule"] = request.Rule
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateProvider)) {
		query["TemplateProvider"] = request.TemplateProvider
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertTemplate"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertTemplate(request *CreateAlertTemplateRequest) (_result *CreateAlertTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertTemplateResponse{}
	_body, _err := client.CreateAlertTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDispatchRuleWithOptions(request *CreateDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DispatchRule)) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDispatchRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDispatchRule(request *CreateDispatchRuleRequest) (_result *CreateDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDispatchRuleResponse{}
	_body, _err := client.CreateDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePrometheusAlertRuleWithOptions(request *CreatePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *CreatePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchRuleId)) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrometheusAlertRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePrometheusAlertRule(request *CreatePrometheusAlertRuleRequest) (_result *CreatePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrometheusAlertRuleResponse{}
	_body, _err := client.CreatePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRetcodeAppWithOptions(request *CreateRetcodeAppRequest, runtime *util.RuntimeOptions) (_result *CreateRetcodeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetcodeAppName)) {
		query["RetcodeAppName"] = request.RetcodeAppName
	}

	if !tea.BoolValue(util.IsUnset(request.RetcodeAppType)) {
		query["RetcodeAppType"] = request.RetcodeAppType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRetcodeApp"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRetcodeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRetcodeApp(request *CreateRetcodeAppRequest) (_result *CreateRetcodeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRetcodeAppResponse{}
	_body, _err := client.CreateRetcodeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWehookWithOptions(request *CreateWehookRequest, runtime *util.RuntimeOptions) (_result *CreateWehookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.HttpHeaders)) {
		query["HttpHeaders"] = request.HttpHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.HttpParams)) {
		query["HttpParams"] = request.HttpParams
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWehook"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWehookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWehook(request *CreateWehookRequest) (_result *CreateWehookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWehookResponse{}
	_body, _err := client.CreateWehookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertContactWithOptions(request *DeleteAlertContactRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertContact"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertContact(request *DeleteAlertContactRequest) (_result *DeleteAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.DeleteAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertContactGroupWithOptions(request *DeleteAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupId)) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertContactGroup"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertContactGroup(request *DeleteAlertContactGroupRequest) (_result *DeleteAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.DeleteAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertRulesWithOptions(request *DeleteAlertRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertIds)) {
		query["AlertIds"] = request.AlertIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertRules(request *DeleteAlertRulesRequest) (_result *DeleteAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertRulesResponse{}
	_body, _err := client.DeleteAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertTemplateWithOptions(request *DeleteAlertTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertTemplate"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertTemplate(request *DeleteAlertTemplateRequest) (_result *DeleteAlertTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertTemplateResponse{}
	_body, _err := client.DeleteAlertTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDispatchRuleWithOptions(request *DeleteDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDispatchRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDispatchRule(request *DeleteDispatchRuleRequest) (_result *DeleteDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDispatchRuleResponse{}
	_body, _err := client.DeleteDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGrafanaResourceWithOptions(request *DeleteGrafanaResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteGrafanaResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGrafanaResource"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGrafanaResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGrafanaResource(request *DeleteGrafanaResourceRequest) (_result *DeleteGrafanaResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGrafanaResourceResponse{}
	_body, _err := client.DeleteGrafanaResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePrometheusAlertRuleWithOptions(request *DeletePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *DeletePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrometheusAlertRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePrometheusAlertRule(request *DeletePrometheusAlertRuleRequest) (_result *DeletePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrometheusAlertRuleResponse{}
	_body, _err := client.DeletePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRetcodeAppWithOptions(request *DeleteRetcodeAppRequest, runtime *util.RuntimeOptions) (_result *DeleteRetcodeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRetcodeApp"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRetcodeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRetcodeApp(request *DeleteRetcodeAppRequest) (_result *DeleteRetcodeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRetcodeAppResponse{}
	_body, _err := client.DeleteRetcodeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScenarioWithOptions(request *DeleteScenarioRequest, runtime *util.RuntimeOptions) (_result *DeleteScenarioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioId)) {
		query["ScenarioId"] = request.ScenarioId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScenario"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScenario(request *DeleteScenarioRequest) (_result *DeleteScenarioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.DeleteScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTraceAppWithOptions(request *DeleteTraceAppRequest, runtime *util.RuntimeOptions) (_result *DeleteTraceAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTraceApp"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTraceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTraceApp(request *DeleteTraceAppRequest) (_result *DeleteTraceAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTraceAppResponse{}
	_body, _err := client.DeleteTraceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDispatchRuleWithOptions(request *DescribeDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDispatchRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDispatchRule(request *DescribeDispatchRuleRequest) (_result *DescribeDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDispatchRuleResponse{}
	_body, _err := client.DescribeDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePrometheusAlertRuleWithOptions(request *DescribePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *DescribePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrometheusAlertRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrometheusAlertRule(request *DescribePrometheusAlertRuleRequest) (_result *DescribePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePrometheusAlertRuleResponse{}
	_body, _err := client.DescribePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTraceLicenseKeyWithOptions(request *DescribeTraceLicenseKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeTraceLicenseKeyResponse, _err error) {
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
		Action:      tea.String("DescribeTraceLicenseKey"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTraceLicenseKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraceLicenseKey(request *DescribeTraceLicenseKeyRequest) (_result *DescribeTraceLicenseKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTraceLicenseKeyResponse{}
	_body, _err := client.DescribeTraceLicenseKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTraceLocationWithOptions(request *DescribeTraceLocationRequest, runtime *util.RuntimeOptions) (_result *DescribeTraceLocationResponse, _err error) {
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
		Action:      tea.String("DescribeTraceLocation"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTraceLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraceLocation(request *DescribeTraceLocationRequest) (_result *DescribeTraceLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTraceLocationResponse{}
	_body, _err := client.DescribeTraceLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAlertTemplateWithOptions(request *DisableAlertTemplateRequest, runtime *util.RuntimeOptions) (_result *DisableAlertTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAlertTemplate"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAlertTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAlertTemplate(request *DisableAlertTemplateRequest) (_result *DisableAlertTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAlertTemplateResponse{}
	_body, _err := client.DisableAlertTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAlertTemplateWithOptions(request *EnableAlertTemplateRequest, runtime *util.RuntimeOptions) (_result *EnableAlertTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAlertTemplate"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAlertTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAlertTemplate(request *EnableAlertTemplateRequest) (_result *EnableAlertTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAlertTemplateResponse{}
	_body, _err := client.EnableAlertTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportPrometheusRulesWithOptions(request *ExportPrometheusRulesRequest, runtime *util.RuntimeOptions) (_result *ExportPrometheusRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportPrometheusRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportPrometheusRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportPrometheusRules(request *ExportPrometheusRulesRequest) (_result *ExportPrometheusRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportPrometheusRulesResponse{}
	_body, _err := client.ExportPrometheusRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentDownloadUrlWithOptions(request *GetAgentDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetAgentDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentDownloadUrl"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentDownloadUrl(request *GetAgentDownloadUrlRequest) (_result *GetAgentDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentDownloadUrlResponse{}
	_body, _err := client.GetAgentDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppApiByPageWithOptions(request *GetAppApiByPageRequest, runtime *util.RuntimeOptions) (_result *GetAppApiByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalMills)) {
		query["IntervalMills"] = request.IntervalMills
	}

	if !tea.BoolValue(util.IsUnset(request.PId)) {
		query["PId"] = request.PId
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
		Action:      tea.String("GetAppApiByPage"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppApiByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppApiByPage(request *GetAppApiByPageRequest) (_result *GetAppApiByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppApiByPageResponse{}
	_body, _err := client.GetAppApiByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsistencySnapshotWithOptions(request *GetConsistencySnapshotRequest, runtime *util.RuntimeOptions) (_result *GetConsistencySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentTimestamp)) {
		query["CurrentTimestamp"] = request.CurrentTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUserId)) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsistencySnapshot"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsistencySnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsistencySnapshot(request *GetConsistencySnapshotRequest) (_result *GetConsistencySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsistencySnapshotResponse{}
	_body, _err := client.GetConsistencySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIntegrationTokenWithOptions(request *GetIntegrationTokenRequest, runtime *util.RuntimeOptions) (_result *GetIntegrationTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIntegrationToken"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIntegrationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIntegrationToken(request *GetIntegrationTokenRequest) (_result *GetIntegrationTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIntegrationTokenResponse{}
	_body, _err := client.GetIntegrationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMultipleTraceWithOptions(request *GetMultipleTraceRequest, runtime *util.RuntimeOptions) (_result *GetMultipleTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceIDs)) {
		query["TraceIDs"] = request.TraceIDs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultipleTrace"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultipleTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMultipleTrace(request *GetMultipleTraceRequest) (_result *GetMultipleTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultipleTraceResponse{}
	_body, _err := client.GetMultipleTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrometheusApiTokenWithOptions(request *GetPrometheusApiTokenRequest, runtime *util.RuntimeOptions) (_result *GetPrometheusApiTokenResponse, _err error) {
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
		Action:      tea.String("GetPrometheusApiToken"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrometheusApiTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrometheusApiToken(request *GetPrometheusApiTokenRequest) (_result *GetPrometheusApiTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrometheusApiTokenResponse{}
	_body, _err := client.GetPrometheusApiTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRetcodeShareUrlWithOptions(request *GetRetcodeShareUrlRequest, runtime *util.RuntimeOptions) (_result *GetRetcodeShareUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRetcodeShareUrl"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRetcodeShareUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRetcodeShareUrl(request *GetRetcodeShareUrlRequest) (_result *GetRetcodeShareUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRetcodeShareUrlResponse{}
	_body, _err := client.GetRetcodeShareUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackWithOptions(request *GetStackRequest, runtime *util.RuntimeOptions) (_result *GetStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RpcID)) {
		query["RpcID"] = request.RpcID
	}

	if !tea.BoolValue(util.IsUnset(request.TraceID)) {
		query["TraceID"] = request.TraceID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStack"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStack(request *GetStackRequest) (_result *GetStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackResponse{}
	_body, _err := client.GetStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceWithOptions(request *GetTraceRequest, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceID)) {
		query["TraceID"] = request.TraceID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrace"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrace(request *GetTraceRequest) (_result *GetTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceAppWithOptions(request *GetTraceAppRequest, runtime *util.RuntimeOptions) (_result *GetTraceAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTraceApp"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTraceApp(request *GetTraceAppRequest) (_result *GetTraceAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceAppResponse{}
	_body, _err := client.GetTraceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportAppAlertRulesWithOptions(request *ImportAppAlertRulesRequest, runtime *util.RuntimeOptions) (_result *ImportAppAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoStart)) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.Pids)) {
		query["Pids"] = request.Pids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplageAlertConfig)) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateAlertId)) {
		query["TemplateAlertId"] = request.TemplateAlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportAppAlertRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportAppAlertRules(request *ImportAppAlertRulesRequest) (_result *ImportAppAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.ImportAppAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportCustomAlertRulesWithOptions(request *ImportCustomAlertRulesRequest, runtime *util.RuntimeOptions) (_result *ImportCustomAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoStart)) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplageAlertConfig)) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateAlertConfig)) {
		query["TemplateAlertConfig"] = request.TemplateAlertConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportCustomAlertRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportCustomAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportCustomAlertRules(request *ImportCustomAlertRulesRequest) (_result *ImportCustomAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportCustomAlertRulesResponse{}
	_body, _err := client.ImportCustomAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportPrometheusRulesWithOptions(request *ImportPrometheusRulesRequest, runtime *util.RuntimeOptions) (_result *ImportPrometheusRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportPrometheusRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportPrometheusRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportPrometheusRules(request *ImportPrometheusRulesRequest) (_result *ImportPrometheusRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportPrometheusRulesResponse{}
	_body, _err := client.ImportPrometheusRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActivatedAlertsWithOptions(request *ListActivatedAlertsRequest, runtime *util.RuntimeOptions) (_result *ListActivatedAlertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
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
		Action:      tea.String("ListActivatedAlerts"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActivatedAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActivatedAlerts(request *ListActivatedAlertsRequest) (_result *ListActivatedAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActivatedAlertsResponse{}
	_body, _err := client.ListActivatedAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlertTemplatesWithOptions(request *ListAlertTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListAlertTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertProvider)) {
		query["AlertProvider"] = request.AlertProvider
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateProvider)) {
		query["TemplateProvider"] = request.TemplateProvider
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlertTemplates"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlertTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlertTemplates(request *ListAlertTemplatesRequest) (_result *ListAlertTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlertTemplatesResponse{}
	_body, _err := client.ListAlertTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterFromGrafanaWithOptions(request *ListClusterFromGrafanaRequest, runtime *util.RuntimeOptions) (_result *ListClusterFromGrafanaResponse, _err error) {
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
		Action:      tea.String("ListClusterFromGrafana"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterFromGrafanaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterFromGrafana(request *ListClusterFromGrafanaRequest) (_result *ListClusterFromGrafanaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterFromGrafanaResponse{}
	_body, _err := client.ListClusterFromGrafanaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardsWithOptions(request *ListDashboardsRequest, runtime *util.RuntimeOptions) (_result *ListDashboardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RecreateSwitch)) {
		query["RecreateSwitch"] = request.RecreateSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDashboards"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboards(request *ListDashboardsRequest) (_result *ListDashboardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDashboardsResponse{}
	_body, _err := client.ListDashboardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDispatchRuleWithOptions(request *ListDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *ListDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.System)) {
		query["System"] = request.System
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDispatchRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDispatchRule(request *ListDispatchRuleRequest) (_result *ListDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDispatchRuleResponse{}
	_body, _err := client.ListDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPromClustersWithOptions(request *ListPromClustersRequest, runtime *util.RuntimeOptions) (_result *ListPromClustersResponse, _err error) {
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
		Action:      tea.String("ListPromClusters"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPromClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPromClusters(request *ListPromClustersRequest) (_result *ListPromClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPromClustersResponse{}
	_body, _err := client.ListPromClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrometheusAlertRulesWithOptions(request *ListPrometheusAlertRulesRequest, runtime *util.RuntimeOptions) (_result *ListPrometheusAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MatchExpressions)) {
		query["MatchExpressions"] = request.MatchExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrometheusAlertRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrometheusAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrometheusAlertRules(request *ListPrometheusAlertRulesRequest) (_result *ListPrometheusAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrometheusAlertRulesResponse{}
	_body, _err := client.ListPrometheusAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrometheusAlertTemplatesWithOptions(request *ListPrometheusAlertTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListPrometheusAlertTemplatesResponse, _err error) {
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
		Action:      tea.String("ListPrometheusAlertTemplates"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrometheusAlertTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrometheusAlertTemplates(request *ListPrometheusAlertTemplatesRequest) (_result *ListPrometheusAlertTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrometheusAlertTemplatesResponse{}
	_body, _err := client.ListPrometheusAlertTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRetcodeAppsWithOptions(request *ListRetcodeAppsRequest, runtime *util.RuntimeOptions) (_result *ListRetcodeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRetcodeApps"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRetcodeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRetcodeApps(request *ListRetcodeAppsRequest) (_result *ListRetcodeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRetcodeAppsResponse{}
	_body, _err := client.ListRetcodeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScenarioWithOptions(request *ListScenarioRequest, runtime *util.RuntimeOptions) (_result *ListScenarioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		query["Scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScenario"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScenario(request *ListScenarioRequest) (_result *ListScenarioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScenarioResponse{}
	_body, _err := client.ListScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerlessTopNAppsWithOptions(request *ListServerlessTopNAppsRequest, runtime *util.RuntimeOptions) (_result *ListServerlessTopNAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
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
		Action:      tea.String("ListServerlessTopNApps"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerlessTopNAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerlessTopNApps(request *ListServerlessTopNAppsRequest) (_result *ListServerlessTopNAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerlessTopNAppsResponse{}
	_body, _err := client.ListServerlessTopNAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTraceAppsWithOptions(request *ListTraceAppsRequest, runtime *util.RuntimeOptions) (_result *ListTraceAppsResponse, _err error) {
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
		Action:      tea.String("ListTraceApps"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTraceAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTraceApps(request *ListTraceAppsRequest) (_result *ListTraceAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTraceAppsResponse{}
	_body, _err := client.ListTraceAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenArmsDefaultSLRWithOptions(request *OpenArmsDefaultSLRRequest, runtime *util.RuntimeOptions) (_result *OpenArmsDefaultSLRResponse, _err error) {
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
		Action:      tea.String("OpenArmsDefaultSLR"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenArmsDefaultSLRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenArmsDefaultSLR(request *OpenArmsDefaultSLRRequest) (_result *OpenArmsDefaultSLRResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenArmsDefaultSLRResponse{}
	_body, _err := client.OpenArmsDefaultSLRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenArmsServiceWithOptions(request *OpenArmsServiceRequest, runtime *util.RuntimeOptions) (_result *OpenArmsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenArmsService"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenArmsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenArmsService(request *OpenArmsServiceRequest) (_result *OpenArmsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenArmsServiceResponse{}
	_body, _err := client.OpenArmsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenVClusterWithOptions(request *OpenVClusterRequest, runtime *util.RuntimeOptions) (_result *OpenVClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RecreateSwitch)) {
		query["RecreateSwitch"] = request.RecreateSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenVCluster"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenVClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenVCluster(request *OpenVClusterRequest) (_result *OpenVClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenVClusterResponse{}
	_body, _err := client.OpenVClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenXtraceDefaultSLRWithOptions(request *OpenXtraceDefaultSLRRequest, runtime *util.RuntimeOptions) (_result *OpenXtraceDefaultSLRResponse, _err error) {
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
		Action:      tea.String("OpenXtraceDefaultSLR"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenXtraceDefaultSLRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenXtraceDefaultSLR(request *OpenXtraceDefaultSLRRequest) (_result *OpenXtraceDefaultSLRResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenXtraceDefaultSLRResponse{}
	_body, _err := client.OpenXtraceDefaultSLRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDatasetWithOptions(request *QueryDatasetRequest, runtime *util.RuntimeOptions) (_result *QueryDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.DateStr)) {
		query["DateStr"] = request.DateStr
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.HungryMode)) {
		query["HungryMode"] = request.HungryMode
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.IsDrillDown)) {
		query["IsDrillDown"] = request.IsDrillDown
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		query["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		query["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.OptionalDims)) {
		query["OptionalDims"] = request.OptionalDims
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByKey)) {
		query["OrderByKey"] = request.OrderByKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUserId)) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ReduceTail)) {
		query["ReduceTail"] = request.ReduceTail
	}

	if !tea.BoolValue(util.IsUnset(request.RequiredDims)) {
		query["RequiredDims"] = request.RequiredDims
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDataset"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDataset(request *QueryDatasetRequest) (_result *QueryDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDatasetResponse{}
	_body, _err := client.QueryDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMetricWithOptions(request *QueryMetricRequest, runtime *util.RuntimeOptions) (_result *QueryMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsistencyDataKey)) {
		query["ConsistencyDataKey"] = request.ConsistencyDataKey
	}

	if !tea.BoolValue(util.IsUnset(request.ConsistencyQueryStrategy)) {
		query["ConsistencyQueryStrategy"] = request.ConsistencyQueryStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUserId)) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetric"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMetric(request *QueryMetricRequest) (_result *QueryMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricResponse{}
	_body, _err := client.QueryMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMetricByPageWithOptions(request *QueryMetricByPageRequest, runtime *util.RuntimeOptions) (_result *QueryMetricByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFilters)) {
		query["CustomFilters"] = request.CustomFilters
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
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
		Action:      tea.String("QueryMetricByPage"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMetricByPage(request *QueryMetricByPageRequest) (_result *QueryMetricByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricByPageResponse{}
	_body, _err := client.QueryMetricByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTraceAppConfigWithOptions(request *SaveTraceAppConfigRequest, runtime *util.RuntimeOptions) (_result *SaveTraceAppConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Settings)) {
		query["Settings"] = request.Settings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTraceAppConfig"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTraceAppConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTraceAppConfig(request *SaveTraceAppConfigRequest) (_result *SaveTraceAppConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTraceAppConfigResponse{}
	_body, _err := client.SaveTraceAppConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertContactWithOptions(request *SearchAlertContactRequest, runtime *util.RuntimeOptions) (_result *SearchAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertContact"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertContact(request *SearchAlertContactRequest) (_result *SearchAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.SearchAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertContactGroupWithOptions(request *SearchAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *SearchAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDetail)) {
		query["IsDetail"] = request.IsDetail
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertContactGroup"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertContactGroup(request *SearchAlertContactGroupRequest) (_result *SearchAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.SearchAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertHistoriesWithOptions(request *SearchAlertHistoriesRequest, runtime *util.RuntimeOptions) (_result *SearchAlertHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		query["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
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
		Action:      tea.String("SearchAlertHistories"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertHistories(request *SearchAlertHistoriesRequest) (_result *SearchAlertHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertHistoriesResponse{}
	_body, _err := client.SearchAlertHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertRulesWithOptions(request *SearchAlertRulesRequest, runtime *util.RuntimeOptions) (_result *SearchAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertRules"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertRules(request *SearchAlertRulesRequest) (_result *SearchAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertRulesResponse{}
	_body, _err := client.SearchAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEventsWithOptions(request *SearchEventsRequest, runtime *util.RuntimeOptions) (_result *SearchEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		query["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsTrigger)) {
		query["IsTrigger"] = request.IsTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
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
		Action:      tea.String("SearchEvents"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEvents(request *SearchEventsRequest) (_result *SearchEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchEventsResponse{}
	_body, _err := client.SearchEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchRetcodeAppByPageWithOptions(request *SearchRetcodeAppByPageRequest, runtime *util.RuntimeOptions) (_result *SearchRetcodeAppByPageResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RetcodeAppName)) {
		query["RetcodeAppName"] = request.RetcodeAppName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchRetcodeAppByPage"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchRetcodeAppByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchRetcodeAppByPage(request *SearchRetcodeAppByPageRequest) (_result *SearchRetcodeAppByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchRetcodeAppByPageResponse{}
	_body, _err := client.SearchRetcodeAppByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTraceAppByNameWithOptions(request *SearchTraceAppByNameRequest, runtime *util.RuntimeOptions) (_result *SearchTraceAppByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceAppName)) {
		query["TraceAppName"] = request.TraceAppName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraceAppByName"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTraceAppByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraceAppByName(request *SearchTraceAppByNameRequest) (_result *SearchTraceAppByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTraceAppByNameResponse{}
	_body, _err := client.SearchTraceAppByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTraceAppByPageWithOptions(request *SearchTraceAppByPageRequest, runtime *util.RuntimeOptions) (_result *SearchTraceAppByPageResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TraceAppName)) {
		query["TraceAppName"] = request.TraceAppName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraceAppByPage"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTraceAppByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraceAppByPage(request *SearchTraceAppByPageRequest) (_result *SearchTraceAppByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTraceAppByPageResponse{}
	_body, _err := client.SearchTraceAppByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTracesWithOptions(request *SearchTracesRequest, runtime *util.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusionFilters)) {
		query["ExclusionFilters"] = request.ExclusionFilters
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraces"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraces(request *SearchTracesRequest) (_result *SearchTracesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesResponse{}
	_body, _err := client.SearchTracesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTracesByPageWithOptions(request *SearchTracesByPageRequest, runtime *util.RuntimeOptions) (_result *SearchTracesByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusionFilters)) {
		query["ExclusionFilters"] = request.ExclusionFilters
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
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

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTracesByPage"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTracesByPage(request *SearchTracesByPageRequest) (_result *SearchTracesByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.SearchTracesByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCustomIncidentsWithOptions(request *SendCustomIncidentsRequest, runtime *util.RuntimeOptions) (_result *SendCustomIncidentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Incidents)) {
		query["Incidents"] = request.Incidents
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomIncidents"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCustomIncidentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomIncidents(request *SendCustomIncidentsRequest) (_result *SendCustomIncidentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomIncidentsResponse{}
	_body, _err := client.SendCustomIncidentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMseIncidentWithOptions(request *SendMseIncidentRequest, runtime *util.RuntimeOptions) (_result *SendMseIncidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Incidents)) {
		query["Incidents"] = request.Incidents
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMseIncident"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMseIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMseIncident(request *SendMseIncidentRequest) (_result *SendMseIncidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMseIncidentResponse{}
	_body, _err := client.SendMseIncidentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRetcodeShareStatusWithOptions(request *SetRetcodeShareStatusRequest, runtime *util.RuntimeOptions) (_result *SetRetcodeShareStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetRetcodeShareStatus"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRetcodeShareStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRetcodeShareStatus(request *SetRetcodeShareStatusRequest) (_result *SetRetcodeShareStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRetcodeShareStatusResponse{}
	_body, _err := client.SetRetcodeShareStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAlertWithOptions(request *StartAlertRequest, runtime *util.RuntimeOptions) (_result *StartAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAlert"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAlert(request *StartAlertRequest) (_result *StartAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAlertResponse{}
	_body, _err := client.StartAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAlertWithOptions(request *StopAlertRequest, runtime *util.RuntimeOptions) (_result *StopAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAlert"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAlert(request *StopAlertRequest) (_result *StopAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAlertResponse{}
	_body, _err := client.StopAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertContactWithOptions(request *UpdateAlertContactRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.DingRobotWebhookUrl)) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemNoc)) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertContact"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertContact(request *UpdateAlertContactRequest) (_result *UpdateAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertContactResponse{}
	_body, _err := client.UpdateAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertContactGroupWithOptions(request *UpdateAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupId)) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertContactGroup"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertContactGroup(request *UpdateAlertContactGroupRequest) (_result *UpdateAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertContactGroupResponse{}
	_body, _err := client.UpdateAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertRuleWithOptions(request *UpdateAlertRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoStart)) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplageAlertConfig)) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertRule(request *UpdateAlertRuleRequest) (_result *UpdateAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.UpdateAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertTemplateWithOptions(request *UpdateAlertTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.MatchExpressions)) {
		query["MatchExpressions"] = request.MatchExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		query["Rule"] = request.Rule
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertTemplate"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertTemplate(request *UpdateAlertTemplateRequest) (_result *UpdateAlertTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertTemplateResponse{}
	_body, _err := client.UpdateAlertTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDispatchRuleWithOptions(request *UpdateDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DispatchRule)) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDispatchRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDispatchRule(request *UpdateDispatchRuleRequest) (_result *UpdateDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDispatchRuleResponse{}
	_body, _err := client.UpdateDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePrometheusAlertRuleWithOptions(request *UpdatePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *UpdatePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchRuleId)) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrometheusAlertRule"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePrometheusAlertRule(request *UpdatePrometheusAlertRuleRequest) (_result *UpdatePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePrometheusAlertRuleResponse{}
	_body, _err := client.UpdatePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWebhookWithOptions(request *UpdateWebhookRequest, runtime *util.RuntimeOptions) (_result *UpdateWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.HttpHeaders)) {
		query["HttpHeaders"] = request.HttpHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.HttpParams)) {
		query["HttpParams"] = request.HttpParams
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWebhook"),
		Version:     tea.String("2021-04-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWebhook(request *UpdateWebhookRequest) (_result *UpdateWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.UpdateWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

type AddGrafanaRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
}

func (s AddGrafanaRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaRequest) GoString() string {
	return s.String()
}

func (s *AddGrafanaRequest) SetRegionId(v string) *AddGrafanaRequest {
	s.RegionId = &v
	return s
}

func (s *AddGrafanaRequest) SetClusterId(v string) *AddGrafanaRequest {
	s.ClusterId = &v
	return s
}

func (s *AddGrafanaRequest) SetIntegration(v string) *AddGrafanaRequest {
	s.Integration = &v
	return s
}

type AddGrafanaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s AddGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponseBody) SetRequestId(v string) *AddGrafanaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGrafanaResponseBody) SetData(v string) *AddGrafanaResponseBody {
	s.Data = &v
	return s
}

type AddGrafanaResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddGrafanaResponse) SetBody(v *AddGrafanaResponseBody) *AddGrafanaResponse {
	s.Body = v
	return s
}

type AddIntegrationRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
}

func (s AddIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationRequest) GoString() string {
	return s.String()
}

func (s *AddIntegrationRequest) SetRegionId(v string) *AddIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *AddIntegrationRequest) SetClusterId(v string) *AddIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *AddIntegrationRequest) SetIntegration(v string) *AddIntegrationRequest {
	s.Integration = &v
	return s
}

type AddIntegrationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s AddIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponseBody) SetRequestId(v string) *AddIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIntegrationResponseBody) SetData(v string) *AddIntegrationResponseBody {
	s.Data = &v
	return s
}

type AddIntegrationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddIntegrationResponse) SetBody(v *AddIntegrationResponseBody) *AddIntegrationResponse {
	s.Body = v
	return s
}

type ApplyScenarioRequest struct {
	RegionId     *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario     *string                `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	AppId        *string                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Sign         *string                `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Config       map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	SnTransfer   *bool                  `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	SnStat       *bool                  `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnDump       *bool                  `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce      *bool                  `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	UpdateOption *bool                  `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioRequest) SetRegionId(v string) *ApplyScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioRequest) SetScenario(v string) *ApplyScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioRequest) SetName(v string) *ApplyScenarioRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioRequest) SetAppId(v string) *ApplyScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioRequest) SetSign(v string) *ApplyScenarioRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioRequest) SetConfig(v map[string]interface{}) *ApplyScenarioRequest {
	s.Config = v
	return s
}

func (s *ApplyScenarioRequest) SetSnTransfer(v bool) *ApplyScenarioRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnStat(v bool) *ApplyScenarioRequest {
	s.SnStat = &v
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

func (s *ApplyScenarioRequest) SetUpdateOption(v bool) *ApplyScenarioRequest {
	s.UpdateOption = &v
	return s
}

type ApplyScenarioShrinkRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario     *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Sign         *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	SnTransfer   *bool   `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	SnStat       *bool   `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnDump       *bool   `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce      *bool   `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	UpdateOption *bool   `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioShrinkRequest) SetRegionId(v string) *ApplyScenarioShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetScenario(v string) *ApplyScenarioShrinkRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetName(v string) *ApplyScenarioShrinkRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetAppId(v string) *ApplyScenarioShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSign(v string) *ApplyScenarioShrinkRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetConfigShrink(v string) *ApplyScenarioShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnTransfer(v bool) *ApplyScenarioShrinkRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnStat(v bool) *ApplyScenarioShrinkRequest {
	s.SnStat = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ApplyScenarioResponse) SetBody(v *ApplyScenarioResponseBody) *ApplyScenarioResponse {
	s.Body = v
	return s
}

type CheckDataConsistencyRequest struct {
	CurrentTimestamp *int64  `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Pid              *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s CheckDataConsistencyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDataConsistencyRequest) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyRequest) SetCurrentTimestamp(v int64) *CheckDataConsistencyRequest {
	s.CurrentTimestamp = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetRegionId(v string) *CheckDataConsistencyRequest {
	s.RegionId = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetProxyUserId(v string) *CheckDataConsistencyRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetPid(v string) *CheckDataConsistencyRequest {
	s.Pid = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetAppType(v string) *CheckDataConsistencyRequest {
	s.AppType = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckDataConsistencyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckDataConsistencyResponse) SetBody(v *CheckDataConsistencyResponseBody) *CheckDataConsistencyResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleForDeletingRequest struct {
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SPIRegionId    *string `json:"SPIRegionId,omitempty" xml:"SPIRegionId,omitempty"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckServiceLinkedRoleForDeletingRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRoleArn(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetServiceName(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.ServiceName = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetSPIRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.SPIRegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetDeletionTaskId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RegionId = &v
	return s
}

type CheckServiceLinkedRoleForDeletingResponseBody struct {
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Deletable  *bool                                                      `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	RoleUsages []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetDeletable(v bool) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.Deletable = &v
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckServiceLinkedRoleForDeletingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckServiceLinkedRoleForDeletingResponse) SetBody(v *CheckServiceLinkedRoleForDeletingResponseBody) *CheckServiceLinkedRoleForDeletingResponse {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ConfigAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAppResponseBody) SetRequestId(v string) *ConfigAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigAppResponseBody) SetData(v string) *ConfigAppResponseBody {
	s.Data = &v
	return s
}

type ConfigAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigAppResponse) SetBody(v *ConfigAppResponseBody) *ConfigAppResponse {
	s.Body = v
	return s
}

type CreateAlertContactRequest struct {
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId         *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
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

func (s *CreateAlertContactRequest) SetPhoneNum(v string) *CreateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateAlertContactRequest) SetEmail(v string) *CreateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *CreateAlertContactRequest) SetDingRobotWebhookUrl(v string) *CreateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *CreateAlertContactRequest) SetSystemNoc(v bool) *CreateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

func (s *CreateAlertContactRequest) SetRegionId(v string) *CreateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertContactRequest) SetProxyUserId(v string) *CreateAlertContactRequest {
	s.ProxyUserId = &v
	return s
}

type CreateAlertContactResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s CreateAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponseBody) SetRequestId(v string) *CreateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertContactResponseBody) SetContactId(v string) *CreateAlertContactResponseBody {
	s.ContactId = &v
	return s
}

type CreateAlertContactResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAlertContactResponse) SetBody(v *CreateAlertContactResponseBody) *CreateAlertContactResponse {
	s.Body = v
	return s
}

type CreateAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
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

func (s *CreateAlertContactGroupRequest) SetProxyUserId(v string) *CreateAlertContactGroupRequest {
	s.ProxyUserId = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAlertContactGroupResponse) SetBody(v *CreateAlertContactGroupResponseBody) *CreateAlertContactGroupResponse {
	s.Body = v
	return s
}

type CreateRetcodeAppRequest struct {
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateRetcodeAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppName(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppType(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppType = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRegionId(v string) *CreateRetcodeAppRequest {
	s.RegionId = &v
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
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppId *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetPid(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.Pid = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetAppId(v int64) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.AppId = &v
	return s
}

type CreateRetcodeAppResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRetcodeAppResponse) SetBody(v *CreateRetcodeAppResponseBody) *CreateRetcodeAppResponse {
	s.Body = v
	return s
}

type CreateWehookRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s CreateWehookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWehookRequest) GoString() string {
	return s.String()
}

func (s *CreateWehookRequest) SetContactName(v string) *CreateWehookRequest {
	s.ContactName = &v
	return s
}

func (s *CreateWehookRequest) SetMethod(v string) *CreateWehookRequest {
	s.Method = &v
	return s
}

func (s *CreateWehookRequest) SetUrl(v string) *CreateWehookRequest {
	s.Url = &v
	return s
}

func (s *CreateWehookRequest) SetHttpParams(v string) *CreateWehookRequest {
	s.HttpParams = &v
	return s
}

func (s *CreateWehookRequest) SetHttpHeaders(v string) *CreateWehookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *CreateWehookRequest) SetRegionId(v string) *CreateWehookRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWehookRequest) SetProxyUserId(v string) *CreateWehookRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CreateWehookRequest) SetBody(v string) *CreateWehookRequest {
	s.Body = &v
	return s
}

type CreateWehookResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s CreateWehookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWehookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWehookResponseBody) SetRequestId(v string) *CreateWehookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWehookResponseBody) SetContactId(v string) *CreateWehookResponseBody {
	s.ContactId = &v
	return s
}

type CreateWehookResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWehookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateWehookResponse) SetBody(v *CreateWehookResponseBody) *CreateWehookResponse {
	s.Body = v
	return s
}

type DeleteAlertContactRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s DeleteAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactRequest) SetRegionId(v string) *DeleteAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlertContactRequest) SetProxyUserId(v string) *DeleteAlertContactRequest {
	s.ProxyUserId = &v
	return s
}

func (s *DeleteAlertContactRequest) SetContactId(v int64) *DeleteAlertContactRequest {
	s.ContactId = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAlertContactResponse) SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse {
	s.Body = v
	return s
}

type DeleteAlertContactGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId    *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactGroupId *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
}

func (s DeleteAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupRequest) SetRegionId(v string) *DeleteAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlertContactGroupRequest) SetProxyUserId(v string) *DeleteAlertContactGroupRequest {
	s.ProxyUserId = &v
	return s
}

func (s *DeleteAlertContactGroupRequest) SetContactGroupId(v int64) *DeleteAlertContactGroupRequest {
	s.ContactGroupId = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAlertContactGroupResponse) SetBody(v *DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse {
	s.Body = v
	return s
}

type DeleteAlertRulesRequest struct {
	AlertIds    *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DeleteAlertRulesRequest) SetProxyUserId(v string) *DeleteAlertRulesRequest {
	s.ProxyUserId = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAlertRulesResponse) SetBody(v *DeleteAlertRulesResponseBody) *DeleteAlertRulesResponse {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteRetcodeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponseBody) SetRequestId(v string) *DeleteRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetData(v string) *DeleteRetcodeAppResponseBody {
	s.Data = &v
	return s
}

type DeleteRetcodeAppResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRetcodeAppResponse) SetBody(v *DeleteRetcodeAppResponseBody) *DeleteRetcodeAppResponse {
	s.Body = v
	return s
}

type DeleteScenarioRequest struct {
	ScenarioId *int64  `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioRequest) GoString() string {
	return s.String()
}

func (s *DeleteScenarioRequest) SetScenarioId(v int64) *DeleteScenarioRequest {
	s.ScenarioId = &v
	return s
}

func (s *DeleteScenarioRequest) SetRegionId(v string) *DeleteScenarioRequest {
	s.RegionId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteScenarioResponse) SetBody(v *DeleteScenarioResponseBody) *DeleteScenarioResponse {
	s.Body = v
	return s
}

type DeleteTraceAppRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
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

func (s *DeleteTraceAppRequest) SetRegionId(v string) *DeleteTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetType(v string) *DeleteTraceAppRequest {
	s.Type = &v
	return s
}

func (s *DeleteTraceAppRequest) SetPid(v string) *DeleteTraceAppRequest {
	s.Pid = &v
	return s
}

type DeleteTraceAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteTraceAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponseBody) SetRequestId(v string) *DeleteTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetData(v string) *DeleteTraceAppResponseBody {
	s.Data = &v
	return s
}

type DeleteTraceAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTraceAppResponse) SetBody(v *DeleteTraceAppResponseBody) *DeleteTraceAppResponse {
	s.Body = v
	return s
}

type DescribeDispatchRuleRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
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

func (s *DescribeDispatchRuleRequest) SetProxyUserId(v string) *DescribeDispatchRuleRequest {
	s.ProxyUserId = &v
	return s
}

type DescribeDispatchRuleResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DispatchRule *DescribeDispatchRuleResponseBodyDispatchRule `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty" type:"Struct"`
}

func (s DescribeDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBody) SetRequestId(v string) *DescribeDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBody) SetDispatchRule(v *DescribeDispatchRuleResponseBodyDispatchRule) *DescribeDispatchRuleResponseBody {
	s.DispatchRule = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRule struct {
	State                    *string                                                               `json:"State,omitempty" xml:"State,omitempty"`
	LabelMatchExpressionGrid *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	GroupRules               []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules             `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
	RuleId                   *int64                                                                `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	NotifyRules              []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules            `json:"NotifyRules,omitempty" xml:"NotifyRules,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetState(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.State = &v
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

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetGroupRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.GroupRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetRuleId(v int64) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetNotifyRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.NotifyRules = v
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
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
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

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleGroupRules struct {
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
	GroupId        *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupInterval  *int64    `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	GroupWaitTime  *int64    `json:"GroupWaitTime,omitempty" xml:"GroupWaitTime,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupingFields(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupingFields = v
	return s
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

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules struct {
	NotifyObjects  []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
	NotifyChannels []*string                                                               `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyObjects(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyObjects = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyChannels(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyChannels = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects struct {
	NotifyObjectId *string `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyObjectId(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyType(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyType = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.Name = &v
	return s
}

type DescribeDispatchRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDispatchRuleResponse) SetBody(v *DescribeDispatchRuleResponseBody) *DescribeDispatchRuleResponse {
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
}

func (s DescribeTraceLicenseKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponseBody) SetRequestId(v string) *DescribeTraceLicenseKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponseBody) SetLicenseKey(v string) *DescribeTraceLicenseKeyResponseBody {
	s.LicenseKey = &v
	return s
}

type DescribeTraceLicenseKeyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTraceLicenseKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionConfigs []*DescribeTraceLocationResponseBodyRegionConfigs `json:"RegionConfigs,omitempty" xml:"RegionConfigs,omitempty" type:"Repeated"`
}

func (s DescribeTraceLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLocationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationResponseBody) SetRequestId(v string) *DescribeTraceLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceLocationResponseBody) SetRegionConfigs(v []*DescribeTraceLocationResponseBodyRegionConfigs) *DescribeTraceLocationResponseBody {
	s.RegionConfigs = v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTraceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTraceLocationResponse) SetBody(v *DescribeTraceLocationResponseBody) *DescribeTraceLocationResponse {
	s.Body = v
	return s
}

type ExportPrometheusRulesRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ExportPrometheusRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportPrometheusRulesRequest) GoString() string {
	return s.String()
}

func (s *ExportPrometheusRulesRequest) SetRegionId(v string) *ExportPrometheusRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ExportPrometheusRulesRequest) SetClusterId(v string) *ExportPrometheusRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ExportPrometheusRulesRequest) SetNameSpace(v string) *ExportPrometheusRulesRequest {
	s.NameSpace = &v
	return s
}

func (s *ExportPrometheusRulesRequest) SetName(v string) *ExportPrometheusRulesRequest {
	s.Name = &v
	return s
}

type ExportPrometheusRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ExportPrometheusRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportPrometheusRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ExportPrometheusRulesResponseBody) SetRequestId(v string) *ExportPrometheusRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportPrometheusRulesResponseBody) SetData(v string) *ExportPrometheusRulesResponseBody {
	s.Data = &v
	return s
}

type ExportPrometheusRulesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportPrometheusRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAgentDownloadUrlResponse) SetBody(v *GetAgentDownloadUrlResponseBody) *GetAgentDownloadUrlResponse {
	s.Body = v
	return s
}

type GetAppApiByPageRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IntervalMills *int32  `json:"IntervalMills,omitempty" xml:"IntervalMills,omitempty"`
	PId           *string `json:"PId,omitempty" xml:"PId,omitempty"`
}

func (s GetAppApiByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageRequest) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageRequest) SetRegionId(v string) *GetAppApiByPageRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetStartTime(v int64) *GetAppApiByPageRequest {
	s.StartTime = &v
	return s
}

func (s *GetAppApiByPageRequest) SetEndTime(v int64) *GetAppApiByPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppApiByPageRequest) SetCurrentPage(v int32) *GetAppApiByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPageSize(v int32) *GetAppApiByPageRequest {
	s.PageSize = &v
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

type GetAppApiByPageResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAppApiByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppApiByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBody) SetMessage(v string) *GetAppApiByPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetRequestId(v string) *GetAppApiByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody {
	s.Data = v
	return s
}

func (s *GetAppApiByPageResponseBody) SetCode(v int32) *GetAppApiByPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetSuccess(v bool) *GetAppApiByPageResponseBody {
	s.Success = &v
	return s
}

type GetAppApiByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *string                  `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *GetAppApiByPageResponseBodyData) SetPageSize(v int32) *GetAppApiByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetTotal(v string) *GetAppApiByPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPage(v int32) *GetAppApiByPageResponseBodyData {
	s.Page = &v
	return s
}

type GetAppApiByPageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppApiByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAppApiByPageResponse) SetBody(v *GetAppApiByPageResponseBody) *GetAppApiByPageResponse {
	s.Body = v
	return s
}

type GetConsistencySnapshotRequest struct {
	CurrentTimestamp *int64  `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Pid              *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s GetConsistencySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsistencySnapshotRequest) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotRequest) SetCurrentTimestamp(v int64) *GetConsistencySnapshotRequest {
	s.CurrentTimestamp = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetRegionId(v string) *GetConsistencySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetProxyUserId(v string) *GetConsistencySnapshotRequest {
	s.ProxyUserId = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetPid(v string) *GetConsistencySnapshotRequest {
	s.Pid = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetAppType(v string) *GetConsistencySnapshotRequest {
	s.AppType = &v
	return s
}

type GetConsistencySnapshotResponseBody struct {
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConsistencyResultJsonStr *string `json:"ConsistencyResultJsonStr,omitempty" xml:"ConsistencyResultJsonStr,omitempty"`
}

func (s GetConsistencySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsistencySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotResponseBody) SetRequestId(v string) *GetConsistencySnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsistencySnapshotResponseBody) SetConsistencyResultJsonStr(v string) *GetConsistencySnapshotResponseBody {
	s.ConsistencyResultJsonStr = &v
	return s
}

type GetConsistencySnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConsistencySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetConsistencySnapshotResponse) SetBody(v *GetConsistencySnapshotResponseBody) *GetConsistencySnapshotResponse {
	s.Body = v
	return s
}

type GetIntegrationTokenRequest struct {
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
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

func (s *GetIntegrationTokenRequest) SetProxyUserId(v string) *GetIntegrationTokenRequest {
	s.ProxyUserId = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIntegrationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MultiCallChainInfos []*GetMultipleTraceResponseBodyMultiCallChainInfos `json:"MultiCallChainInfos,omitempty" xml:"MultiCallChainInfos,omitempty" type:"Repeated"`
}

func (s GetMultipleTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBody) SetRequestId(v string) *GetMultipleTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultipleTraceResponseBody) SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody {
	s.MultiCallChainInfos = v
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
	OperationName *string                                                             `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ResultCode    *string                                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Timestamp     *int64                                                              `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	RpcType       *int32                                                              `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	TagEntryList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	LogEventList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	HaveStack     *bool                                                               `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	ServiceIp     *string                                                             `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Duration      *int64                                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RpcId         *string                                                             `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName   *string                                                             `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TraceID       *string                                                             `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetOperationName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.OperationName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetResultCode(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ResultCode = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcType(v int32) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcType = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetLogEventList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.LogEventList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetHaveStack(v bool) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.HaveStack = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceIp(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceIp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetDuration(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Duration = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TraceID = &v
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

type GetMultipleTraceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMultipleTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPrometheusApiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRetcodeShareUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRetcodeShareUrlResponse) SetBody(v *GetRetcodeShareUrlResponseBody) *GetRetcodeShareUrlResponse {
	s.Body = v
	return s
}

type GetStackRequest struct {
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RpcID    *string `json:"RpcID,omitempty" xml:"RpcID,omitempty"`
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) SetTraceID(v string) *GetStackRequest {
	s.TraceID = &v
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

func (s *GetStackRequest) SetPid(v string) *GetStackRequest {
	s.Pid = &v
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
	StartTime   *int64                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Exception   *string                               `json:"Exception,omitempty" xml:"Exception,omitempty"`
	Api         *string                               `json:"Api,omitempty" xml:"Api,omitempty"`
	Line        *string                               `json:"Line,omitempty" xml:"Line,omitempty"`
	Duration    *int64                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RpcId       *string                               `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName *string                               `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ExtInfo     *GetStackResponseBodyStackInfoExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
}

func (s GetStackResponseBodyStackInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyStackInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfo) SetStartTime(v int64) *GetStackResponseBodyStackInfo {
	s.StartTime = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetException(v string) *GetStackResponseBodyStackInfo {
	s.Exception = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetApi(v string) *GetStackResponseBodyStackInfo {
	s.Api = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetLine(v string) *GetStackResponseBodyStackInfo {
	s.Line = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetDuration(v int64) *GetStackResponseBodyStackInfo {
	s.Duration = &v
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

func (s *GetStackResponseBodyStackInfo) SetExtInfo(v *GetStackResponseBodyStackInfoExtInfo) *GetStackResponseBodyStackInfo {
	s.ExtInfo = v
	return s
}

type GetStackResponseBodyStackInfoExtInfo struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
}

func (s GetStackResponseBodyStackInfoExtInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyStackInfoExtInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetType(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Type = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetInfo(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Info = &v
	return s
}

type GetStackResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetStackResponse) SetBody(v *GetStackResponseBody) *GetStackResponse {
	s.Body = v
	return s
}

type GetTraceRequest struct {
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
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
	OperationName *string                                  `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Timestamp     *int64                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	RpcType       *int32                                   `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	TagEntryList  []*GetTraceResponseBodySpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	LogEventList  []*GetTraceResponseBodySpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	HaveStack     *bool                                    `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	ServiceIp     *string                                  `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Duration      *int64                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RpcId         *string                                  `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName   *string                                  `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TraceID       *string                                  `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpans) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) SetOperationName(v string) *GetTraceResponseBodySpans {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetResultCode(v string) *GetTraceResponseBodySpans {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTimestamp(v int64) *GetTraceResponseBodySpans {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcType(v int32) *GetTraceResponseBodySpans {
	s.RpcType = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTagEntryList(v []*GetTraceResponseBodySpansTagEntryList) *GetTraceResponseBodySpans {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetLogEventList(v []*GetTraceResponseBodySpansLogEventList) *GetTraceResponseBodySpans {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetHaveStack(v bool) *GetTraceResponseBodySpans {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceIp(v string) *GetTraceResponseBodySpans {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetDuration(v int64) *GetTraceResponseBodySpans {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcId(v string) *GetTraceResponseBodySpans {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceName(v string) *GetTraceResponseBodySpans {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTraceID(v string) *GetTraceResponseBodySpans {
	s.TraceID = &v
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

type GetTraceResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceAppResponseBodyTraceApp) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponseBodyTraceApp) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBodyTraceApp) SetType(v string) *GetTraceAppResponseBodyTraceApp {
	s.Type = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppName(v string) *GetTraceAppResponseBodyTraceApp {
	s.AppName = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUpdateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.UpdateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetLabels(v []*string) *GetTraceAppResponseBodyTraceApp {
	s.Labels = v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetShow(v bool) *GetTraceAppResponseBodyTraceApp {
	s.Show = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetCreateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.CreateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetPid(v string) *GetTraceAppResponseBodyTraceApp {
	s.Pid = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppId(v int64) *GetTraceAppResponseBodyTraceApp {
	s.AppId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUserId(v string) *GetTraceAppResponseBodyTraceApp {
	s.UserId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetRegionId(v string) *GetTraceAppResponseBodyTraceApp {
	s.RegionId = &v
	return s
}

type GetTraceAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTraceAppResponse) SetBody(v *GetTraceAppResponseBody) *GetTraceAppResponse {
	s.Body = v
	return s
}

type ImportAppAlertRulesRequest struct {
	TemplateAlertId     *string `json:"TemplateAlertId,omitempty" xml:"TemplateAlertId,omitempty"`
	Pids                *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	ProxyUserId         *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s ImportAppAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequest) SetTemplateAlertId(v string) *ImportAppAlertRulesRequest {
	s.TemplateAlertId = &v
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

func (s *ImportAppAlertRulesRequest) SetContactGroupIds(v string) *ImportAppAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetIsAutoStart(v bool) *ImportAppAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetProxyUserId(v string) *ImportAppAlertRulesRequest {
	s.ProxyUserId = &v
	return s
}

type ImportAppAlertRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ImportAppAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponseBody) SetRequestId(v string) *ImportAppAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportAppAlertRulesResponseBody) SetData(v string) *ImportAppAlertRulesResponseBody {
	s.Data = &v
	return s
}

type ImportAppAlertRulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportAppAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImportAppAlertRulesResponse) SetBody(v *ImportAppAlertRulesResponseBody) *ImportAppAlertRulesResponse {
	s.Body = v
	return s
}

type ImportCustomAlertRulesRequest struct {
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	TemplateAlertConfig *string `json:"TemplateAlertConfig,omitempty" xml:"TemplateAlertConfig,omitempty"`
	ProxyUserId         *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
}

func (s ImportCustomAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportCustomAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesRequest) SetRegionId(v string) *ImportCustomAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetContactGroupIds(v string) *ImportCustomAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetIsAutoStart(v bool) *ImportCustomAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetTemplateAlertConfig(v string) *ImportCustomAlertRulesRequest {
	s.TemplateAlertConfig = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetProxyUserId(v string) *ImportCustomAlertRulesRequest {
	s.ProxyUserId = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportCustomAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

type ImportCustomAlertRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ImportCustomAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportCustomAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesResponseBody) SetRequestId(v string) *ImportCustomAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCustomAlertRulesResponseBody) SetData(v string) *ImportCustomAlertRulesResponseBody {
	s.Data = &v
	return s
}

type ImportCustomAlertRulesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportCustomAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImportCustomAlertRulesResponse) SetBody(v *ImportCustomAlertRulesResponseBody) *ImportCustomAlertRulesResponse {
	s.Body = v
	return s
}

type ImportPrometheusRulesRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ImportPrometheusRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportPrometheusRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesRequest) SetRegionId(v string) *ImportPrometheusRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetClusterId(v string) *ImportPrometheusRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetNameSpace(v string) *ImportPrometheusRulesRequest {
	s.NameSpace = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetName(v string) *ImportPrometheusRulesRequest {
	s.Name = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetContent(v string) *ImportPrometheusRulesRequest {
	s.Content = &v
	return s
}

type ImportPrometheusRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ImportPrometheusRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportPrometheusRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesResponseBody) SetRequestId(v string) *ImportPrometheusRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportPrometheusRulesResponseBody) SetData(v string) *ImportPrometheusRulesResponseBody {
	s.Data = &v
	return s
}

type ImportPrometheusRulesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportPrometheusRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImportPrometheusRulesResponse) SetBody(v *ImportPrometheusRulesResponseBody) *ImportPrometheusRulesResponse {
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
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromClusterList []*ListClusterFromGrafanaResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
}

func (s ListClusterFromGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBody) SetRequestId(v string) *ListClusterFromGrafanaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBody) SetPromClusterList(v []*ListClusterFromGrafanaResponseBodyPromClusterList) *ListClusterFromGrafanaResponseBody {
	s.PromClusterList = v
	return s
}

type ListClusterFromGrafanaResponseBodyPromClusterList struct {
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	IsControllerInstalled *bool   `json:"IsControllerInstalled,omitempty" xml:"IsControllerInstalled,omitempty"`
	AgentStatus           *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	ControllerId          *string `json:"ControllerId,omitempty" xml:"ControllerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstallTime           *int64  `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	PluginsJsonArray      *string `json:"PluginsJsonArray,omitempty" xml:"PluginsJsonArray,omitempty"`
	ClusterType           *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	StateJson             *string `json:"StateJson,omitempty" xml:"StateJson,omitempty"`
	LastHeartBeatTime     *int64  `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	NodeNum               *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUpdateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetCreateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUserId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetOptions(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetAgentStatus(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetExtra(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetControllerId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetRegionId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetInstallTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterType(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterName(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetStateJson(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.StateJson = &v
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

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetId(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

type ListClusterFromGrafanaResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterFromGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListClusterFromGrafanaResponse) SetBody(v *ListClusterFromGrafanaResponseBody) *ListClusterFromGrafanaResponse {
	s.Body = v
	return s
}

type ListDashboardsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDashboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsRequest) SetRegionId(v string) *ListDashboardsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDashboardsRequest) SetClusterId(v string) *ListDashboardsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsRequest) SetClusterType(v string) *ListDashboardsRequest {
	s.ClusterType = &v
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
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Time           *string   `json:"Time,omitempty" xml:"Time,omitempty"`
	Exporter       *string   `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	IsArmsExporter *bool     `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Tags           []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Uid            *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVos) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVos) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVos) SetType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTime(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetExporter(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVos {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTitle(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetId(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUid(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Uid = &v
	return s
}

type ListDashboardsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDashboardsResponse) SetBody(v *ListDashboardsResponseBody) *ListDashboardsResponse {
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
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromClusterList []*ListPromClustersResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
}

func (s ListPromClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPromClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponseBody) SetRequestId(v string) *ListPromClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromClustersResponseBody) SetPromClusterList(v []*ListPromClustersResponseBodyPromClusterList) *ListPromClustersResponseBody {
	s.PromClusterList = v
	return s
}

type ListPromClustersResponseBodyPromClusterList struct {
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	IsControllerInstalled *bool   `json:"IsControllerInstalled,omitempty" xml:"IsControllerInstalled,omitempty"`
	AgentStatus           *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	ControllerId          *string `json:"ControllerId,omitempty" xml:"ControllerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstallTime           *int64  `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	PluginsJsonArray      *string `json:"PluginsJsonArray,omitempty" xml:"PluginsJsonArray,omitempty"`
	ClusterType           *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	StateJson             *string `json:"StateJson,omitempty" xml:"StateJson,omitempty"`
	LastHeartBeatTime     *int64  `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	NodeNum               *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListPromClustersResponseBodyPromClusterList) String() string {
	return tea.Prettify(s)
}

func (s ListPromClustersResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListPromClustersResponseBodyPromClusterList) SetUpdateTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetCreateTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetUserId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetOptions(v string) *ListPromClustersResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListPromClustersResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetAgentStatus(v string) *ListPromClustersResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetExtra(v string) *ListPromClustersResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetControllerId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetRegionId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetInstallTime(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListPromClustersResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterType(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterName(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetStateJson(v string) *ListPromClustersResponseBodyPromClusterList {
	s.StateJson = &v
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

func (s *ListPromClustersResponseBodyPromClusterList) SetId(v int64) *ListPromClustersResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListPromClustersResponseBodyPromClusterList) SetClusterId(v string) *ListPromClustersResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

type ListPromClustersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPromClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListPromClustersResponse) SetBody(v *ListPromClustersResponseBody) *ListPromClustersResponse {
	s.Body = v
	return s
}

type ListRetcodeAppsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRetcodeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsRequest) SetSecurityToken(v string) *ListRetcodeAppsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListRetcodeAppsRequest) SetRegionId(v string) *ListRetcodeAppsRequest {
	s.RegionId = &v
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
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid     *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppName(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppName = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppId(v int64) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppId = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetPid(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.Pid = &v
	return s
}

type ListRetcodeAppsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRetcodeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRetcodeAppsResponse) SetBody(v *ListRetcodeAppsResponseBody) *ListRetcodeAppsResponse {
	s.Body = v
	return s
}

type ListScenarioRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Sign     *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s ListScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioRequest) GoString() string {
	return s.String()
}

func (s *ListScenarioRequest) SetRegionId(v string) *ListScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ListScenarioRequest) SetScenario(v string) *ListScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ListScenarioRequest) SetName(v string) *ListScenarioRequest {
	s.Name = &v
	return s
}

func (s *ListScenarioRequest) SetAppId(v string) *ListScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ListScenarioRequest) SetSign(v string) *ListScenarioRequest {
	s.Sign = &v
	return s
}

type ListScenarioResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ArmsScenarios []*ListScenarioResponseBodyArmsScenarios `json:"ArmsScenarios,omitempty" xml:"ArmsScenarios,omitempty" type:"Repeated"`
}

func (s ListScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBody) SetRequestId(v string) *ListScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenarioResponseBody) SetArmsScenarios(v []*ListScenarioResponseBodyArmsScenarios) *ListScenarioResponseBody {
	s.ArmsScenarios = v
	return s
}

type ListScenarioResponseBodyArmsScenarios struct {
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Sign       *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListScenarioResponseBodyArmsScenarios) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponseBodyArmsScenarios) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUpdateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UpdateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetAppId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.AppId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetSign(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Sign = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetCreateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.CreateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUserId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UserId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetExtensions(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Extensions = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetName(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Name = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetId(v int64) *ListScenarioResponseBodyArmsScenarios {
	s.Id = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetRegionId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.RegionId = &v
	return s
}

type ListScenarioResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListScenarioResponse) SetBody(v *ListScenarioResponseBody) *ListScenarioResponse {
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
	TraceApps []*ListTraceAppsResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTraceAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBody) SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody {
	s.TraceApps = v
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

func (s *ListTraceAppsResponseBody) SetCode(v int32) *ListTraceAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetSuccess(v bool) *ListTraceAppsResponseBody {
	s.Success = &v
	return s
}

type ListTraceAppsResponseBodyTraceApps struct {
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTraceAppsResponseBodyTraceApps) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBodyTraceApps) SetType(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppName(v string) *ListTraceAppsResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUpdateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetLabels(v []*string) *ListTraceAppsResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetShow(v bool) *ListTraceAppsResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetCreateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetPid(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppId(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUserId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.UserId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetRegionId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

type ListTraceAppsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTraceAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTraceAppsResponse) SetBody(v *ListTraceAppsResponseBody) *ListTraceAppsResponse {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s OpenArmsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceResponseBody) SetRequestId(v string) *OpenArmsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenArmsServiceResponseBody) SetOrderId(v string) *OpenArmsServiceResponseBody {
	s.OrderId = &v
	return s
}

type OpenArmsServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenArmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenArmsServiceResponse) SetBody(v *OpenArmsServiceResponseBody) *OpenArmsServiceResponse {
	s.Body = v
	return s
}

type QueryDatasetRequest struct {
	DatasetId     *int64                             `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	IntervalInSec *int32                             `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	DateStr       *string                            `json:"DateStr,omitempty" xml:"DateStr,omitempty"`
	MinTime       *int64                             `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	MaxTime       *int64                             `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	IsDrillDown   *bool                              `json:"IsDrillDown,omitempty" xml:"IsDrillDown,omitempty"`
	OrderByKey    *string                            `json:"OrderByKey,omitempty" xml:"OrderByKey,omitempty"`
	Limit         *int32                             `json:"Limit,omitempty" xml:"Limit,omitempty"`
	ReduceTail    *bool                              `json:"ReduceTail,omitempty" xml:"ReduceTail,omitempty"`
	HungryMode    *bool                              `json:"HungryMode,omitempty" xml:"HungryMode,omitempty"`
	ProxyUserId   *string                            `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Measures      []*string                          `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Dimensions    []*QueryDatasetRequestDimensions   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	RequiredDims  []*QueryDatasetRequestRequiredDims `json:"RequiredDims,omitempty" xml:"RequiredDims,omitempty" type:"Repeated"`
	OptionalDims  []*QueryDatasetRequestOptionalDims `json:"OptionalDims,omitempty" xml:"OptionalDims,omitempty" type:"Repeated"`
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

func (s *QueryDatasetRequest) SetIntervalInSec(v int32) *QueryDatasetRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryDatasetRequest) SetDateStr(v string) *QueryDatasetRequest {
	s.DateStr = &v
	return s
}

func (s *QueryDatasetRequest) SetMinTime(v int64) *QueryDatasetRequest {
	s.MinTime = &v
	return s
}

func (s *QueryDatasetRequest) SetMaxTime(v int64) *QueryDatasetRequest {
	s.MaxTime = &v
	return s
}

func (s *QueryDatasetRequest) SetIsDrillDown(v bool) *QueryDatasetRequest {
	s.IsDrillDown = &v
	return s
}

func (s *QueryDatasetRequest) SetOrderByKey(v string) *QueryDatasetRequest {
	s.OrderByKey = &v
	return s
}

func (s *QueryDatasetRequest) SetLimit(v int32) *QueryDatasetRequest {
	s.Limit = &v
	return s
}

func (s *QueryDatasetRequest) SetReduceTail(v bool) *QueryDatasetRequest {
	s.ReduceTail = &v
	return s
}

func (s *QueryDatasetRequest) SetHungryMode(v bool) *QueryDatasetRequest {
	s.HungryMode = &v
	return s
}

func (s *QueryDatasetRequest) SetProxyUserId(v string) *QueryDatasetRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryDatasetRequest) SetMeasures(v []*string) *QueryDatasetRequest {
	s.Measures = v
	return s
}

func (s *QueryDatasetRequest) SetDimensions(v []*QueryDatasetRequestDimensions) *QueryDatasetRequest {
	s.Dimensions = v
	return s
}

func (s *QueryDatasetRequest) SetRequiredDims(v []*QueryDatasetRequestRequiredDims) *QueryDatasetRequest {
	s.RequiredDims = v
	return s
}

func (s *QueryDatasetRequest) SetOptionalDims(v []*QueryDatasetRequestOptionalDims) *QueryDatasetRequest {
	s.OptionalDims = v
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

type QueryDatasetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetResponseBody) SetRequestId(v string) *QueryDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetResponseBody) SetData(v string) *QueryDatasetResponseBody {
	s.Data = &v
	return s
}

type QueryDatasetResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDatasetResponse) SetBody(v *QueryDatasetResponseBody) *QueryDatasetResponse {
	s.Body = v
	return s
}

type QueryMetricRequest struct {
	IntervalInSec            *int32                       `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	StartTime                *int64                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime                  *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderBy                  *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Limit                    *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Metric                   *string                      `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order                    *string                      `json:"Order,omitempty" xml:"Order,omitempty"`
	ProxyUserId              *string                      `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ConsistencyDataKey       *string                      `json:"ConsistencyDataKey,omitempty" xml:"ConsistencyDataKey,omitempty"`
	ConsistencyQueryStrategy *string                      `json:"ConsistencyQueryStrategy,omitempty" xml:"ConsistencyQueryStrategy,omitempty"`
	Filters                  []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	Dimensions               []*string                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Measures                 []*string                    `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
}

func (s QueryMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricRequest) SetIntervalInSec(v int32) *QueryMetricRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricRequest) SetStartTime(v int64) *QueryMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMetricRequest) SetEndTime(v int64) *QueryMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricRequest) SetOrderBy(v string) *QueryMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricRequest) SetLimit(v int32) *QueryMetricRequest {
	s.Limit = &v
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

func (s *QueryMetricRequest) SetProxyUserId(v string) *QueryMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryMetricRequest) SetConsistencyDataKey(v string) *QueryMetricRequest {
	s.ConsistencyDataKey = &v
	return s
}

func (s *QueryMetricRequest) SetConsistencyQueryStrategy(v string) *QueryMetricRequest {
	s.ConsistencyQueryStrategy = &v
	return s
}

func (s *QueryMetricRequest) SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricRequest) SetDimensions(v []*string) *QueryMetricRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricRequest) SetMeasures(v []*string) *QueryMetricRequest {
	s.Measures = v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricResponseBody) SetRequestId(v string) *QueryMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricResponseBody) SetData(v string) *QueryMetricResponseBody {
	s.Data = &v
	return s
}

type QueryMetricResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMetricResponse) SetBody(v *QueryMetricResponseBody) *QueryMetricResponse {
	s.Body = v
	return s
}

type QueryMetricByPageRequest struct {
	IntervalInSec            *int32                             `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	StartTime                *int64                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime                  *int64                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderBy                  *string                            `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Metric                   *string                            `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order                    *string                            `json:"Order,omitempty" xml:"Order,omitempty"`
	ProxyUserId              *string                            `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ConsistencyDataKey       *string                            `json:"ConsistencyDataKey,omitempty" xml:"ConsistencyDataKey,omitempty"`
	ConsistencyQueryStrategy *string                            `json:"ConsistencyQueryStrategy,omitempty" xml:"ConsistencyQueryStrategy,omitempty"`
	CurrentPage              *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize                 *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Filters                  []*QueryMetricByPageRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	Dimensions               []*string                          `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Measures                 []*string                          `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
}

func (s QueryMetricByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequest) SetIntervalInSec(v int32) *QueryMetricByPageRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricByPageRequest) SetStartTime(v int64) *QueryMetricByPageRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMetricByPageRequest) SetEndTime(v int64) *QueryMetricByPageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrderBy(v string) *QueryMetricByPageRequest {
	s.OrderBy = &v
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

func (s *QueryMetricByPageRequest) SetProxyUserId(v string) *QueryMetricByPageRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryMetricByPageRequest) SetConsistencyDataKey(v string) *QueryMetricByPageRequest {
	s.ConsistencyDataKey = &v
	return s
}

func (s *QueryMetricByPageRequest) SetConsistencyQueryStrategy(v string) *QueryMetricByPageRequest {
	s.ConsistencyQueryStrategy = &v
	return s
}

func (s *QueryMetricByPageRequest) SetCurrentPage(v int32) *QueryMetricByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMetricByPageRequest) SetPageSize(v int32) *QueryMetricByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageRequest) SetFilters(v []*QueryMetricByPageRequestFilters) *QueryMetricByPageRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricByPageRequest) SetDimensions(v []*string) *QueryMetricByPageRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricByPageRequest) SetMeasures(v []*string) *QueryMetricByPageRequest {
	s.Measures = v
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
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryMetricByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBody) SetMessage(v string) *QueryMetricByPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetRequestId(v string) *QueryMetricByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryMetricByPageResponseBody) SetCode(v string) *QueryMetricByPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetSuccess(v bool) *QueryMetricByPageResponseBody {
	s.Success = &v
	return s
}

type QueryMetricByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
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

func (s *QueryMetricByPageResponseBodyData) SetPageSize(v int32) *QueryMetricByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetTotal(v int32) *QueryMetricByPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPage(v int32) *QueryMetricByPageResponseBodyData {
	s.Page = &v
	return s
}

type QueryMetricByPageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMetricByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s SaveTraceAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponseBody) SetRequestId(v string) *SaveTraceAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetData(v string) *SaveTraceAppConfigResponseBody {
	s.Data = &v
	return s
}

type SaveTraceAppConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveTraceAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SaveTraceAppConfigResponse) SetBody(v *SaveTraceAppConfigResponseBody) *SaveTraceAppConfigResponse {
	s.Body = v
	return s
}

type SearchAlertContactRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactIds  *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
}

func (s SearchAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactRequest) SetContactName(v string) *SearchAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactRequest) SetPhone(v string) *SearchAlertContactRequest {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactRequest) SetEmail(v string) *SearchAlertContactRequest {
	s.Email = &v
	return s
}

func (s *SearchAlertContactRequest) SetCurrentPage(v string) *SearchAlertContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertContactRequest) SetPageSize(v string) *SearchAlertContactRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactRequest) SetRegionId(v string) *SearchAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertContactRequest) SetProxyUserId(v string) *SearchAlertContactRequest {
	s.ProxyUserId = &v
	return s
}

func (s *SearchAlertContactRequest) SetContactIds(v string) *SearchAlertContactRequest {
	s.ContactIds = &v
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
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Webhook     *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBeanContacts) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBeanContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUpdateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetDingRobot(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetWebhook(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Webhook = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetEmail(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactId(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetCreateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUserId(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactName(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetSystemNoc(v bool) *SearchAlertContactResponseBodyPageBeanContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetPhone(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Phone = &v
	return s
}

type SearchAlertContactResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchAlertContactResponse) SetBody(v *SearchAlertContactResponseBody) *SearchAlertContactResponse {
	s.Body = v
	return s
}

type SearchAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactName      *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactId        *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactGroupIds  *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsDetail         *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
}

func (s SearchAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupRequest) SetContactGroupName(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetRegionId(v string) *SearchAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetProxyUserId(v string) *SearchAlertContactGroupRequest {
	s.ProxyUserId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactName(v string) *SearchAlertContactGroupRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactId(v int64) *SearchAlertContactGroupRequest {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactGroupIds(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetIsDetail(v bool) *SearchAlertContactGroupRequest {
	s.IsDetail = &v
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
	UpdateTime       *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ContactGroupName *string                                                     `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Contacts         []*SearchAlertContactGroupResponseBodyContactGroupsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	ContactGroupId   *int64                                                      `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	CreateTime       *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId           *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroups) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UpdateTime = &v
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

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupId(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UserId = &v
	return s
}

type SearchAlertContactGroupResponseBodyContactGroupsContacts struct {
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UpdateTime = &v
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

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactId(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactName(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetSystemNoc(v bool) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetPhone(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Phone = &v
	return s
}

type SearchAlertContactGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchAlertContactGroupResponse) SetBody(v *SearchAlertContactGroupResponseBody) *SearchAlertContactGroupResponse {
	s.Body = v
	return s
}

type SearchAlertHistoriesRequest struct {
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s SearchAlertHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesRequest) SetProxyUserId(v string) *SearchAlertHistoriesRequest {
	s.ProxyUserId = &v
	return s
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

func (s *SearchAlertHistoriesRequest) SetEndTime(v int64) *SearchAlertHistoriesRequest {
	s.EndTime = &v
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
	AlarmTime         *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	StrategyId        *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	AlarmResponseCode *int32  `json:"AlarmResponseCode,omitempty" xml:"AlarmResponseCode,omitempty"`
	Emails            *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AlarmSources      *string `json:"AlarmSources,omitempty" xml:"AlarmSources,omitempty"`
	AlarmContent      *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	Phones            *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	AlarmType         *int32  `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	Target            *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmTime(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmTime = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetStrategyId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.StrategyId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmResponseCode(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmResponseCode = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetEmails(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Emails = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetUserId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.UserId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmSources(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmSources = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmContent(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmContent = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetPhones(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Phones = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmType(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmType = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetTarget(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Target = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetId(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Id = &v
	return s
}

type SearchAlertHistoriesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchAlertHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchAlertHistoriesResponse) SetBody(v *SearchAlertHistoriesResponseBody) *SearchAlertHistoriesResponse {
	s.Body = v
	return s
}

type SearchAlertRulesRequest struct {
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s SearchAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesRequest) SetProxyUserId(v string) *SearchAlertRulesRequest {
	s.ProxyUserId = &v
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

func (s *SearchAlertRulesRequest) SetCurrentPage(v int32) *SearchAlertRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPageSize(v int32) *SearchAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesRequest) SetRegionId(v string) *SearchAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPid(v string) *SearchAlertRulesRequest {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesRequest) SetAppType(v string) *SearchAlertRulesRequest {
	s.AppType = &v
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
	Status             *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	AlarmContext       *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext `json:"AlarmContext,omitempty" xml:"AlarmContext,omitempty" type:"Struct"`
	UpdateTime         *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ContactGroupIdList *string                                                     `json:"ContactGroupIdList,omitempty" xml:"ContactGroupIdList,omitempty"`
	Notice             *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice       `json:"Notice,omitempty" xml:"Notice,omitempty" type:"Struct"`
	CreateTime         *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AlertTitle         *string                                                     `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	UserId             *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AlertVersion       *int32                                                      `json:"AlertVersion,omitempty" xml:"AlertVersion,omitempty"`
	AlertRule          *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule    `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	MetricParam        *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam  `json:"MetricParam,omitempty" xml:"MetricParam,omitempty" type:"Struct"`
	AlertType          *int32                                                      `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	ContactGroupIds    *string                                                     `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	Config             *string                                                     `json:"Config,omitempty" xml:"Config,omitempty"`
	RegionId           *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AlertLevel         *string                                                     `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertWay           []*string                                                   `json:"AlertWay,omitempty" xml:"AlertWay,omitempty" type:"Repeated"`
	TaskStatus         *string                                                     `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Title              *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	TaskId             *int64                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Id                 *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	AlertWays          []*string                                                   `json:"AlertWays,omitempty" xml:"AlertWays,omitempty" type:"Repeated"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Status = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlarmContext(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlarmContext = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUpdateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIdList(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIdList = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetNotice(v *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Notice = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetCreateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertVersion(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertVersion = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertRule(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRule = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetMetricParam(v *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricParam = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIds(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetConfig(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Config = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertLevel(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertLevel = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWay(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWay = v
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

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Id = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWays(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWays = v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext struct {
	AlarmContentSubTitle *string `json:"AlarmContentSubTitle,omitempty" xml:"AlarmContentSubTitle,omitempty"`
	AlarmContentTemplate *string `json:"AlarmContentTemplate,omitempty" xml:"AlarmContentTemplate,omitempty"`
	SubTitle             *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
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

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.SubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetContent(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.Content = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesNotice struct {
	EndTime         *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NoticeEndTime   *int64 `json:"NoticeEndTime,omitempty" xml:"NoticeEndTime,omitempty"`
	StartTime       *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	NoticeStartTime *int64 `json:"NoticeStartTime,omitempty" xml:"NoticeStartTime,omitempty"`
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

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.StartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeStartTime = &v
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
	Measure    *string  `json:"Measure,omitempty" xml:"Measure,omitempty"`
	Value      *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
	Aggregates *string  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	NValue     *int32   `json:"NValue,omitempty" xml:"NValue,omitempty"`
	Operator   *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Alias      *string  `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetMeasure(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Measure = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetValue(v float32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Value = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAggregates(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Aggregates = &v
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

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAlias(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Alias = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam struct {
	Type       *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
	AppGroupId *string                                                                `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	AppId      *string                                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid        *string                                                                `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Dimensions []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppGroupId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppGroupId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetPid(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetDimensions(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Dimensions = v
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

type SearchAlertRulesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchAlertRulesResponse) SetBody(v *SearchAlertRulesResponseBody) *SearchAlertRulesResponse {
	s.Body = v
	return s
}

type SearchEventsRequest struct {
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	IsTrigger   *int32  `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s SearchEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsRequest) GoString() string {
	return s.String()
}

func (s *SearchEventsRequest) SetProxyUserId(v string) *SearchEventsRequest {
	s.ProxyUserId = &v
	return s
}

func (s *SearchEventsRequest) SetAlertId(v int64) *SearchEventsRequest {
	s.AlertId = &v
	return s
}

func (s *SearchEventsRequest) SetPid(v string) *SearchEventsRequest {
	s.Pid = &v
	return s
}

func (s *SearchEventsRequest) SetCurrentPage(v int32) *SearchEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchEventsRequest) SetPageSize(v int32) *SearchEventsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEventsRequest) SetRegionId(v string) *SearchEventsRequest {
	s.RegionId = &v
	return s
}

func (s *SearchEventsRequest) SetAppType(v string) *SearchEventsRequest {
	s.AppType = &v
	return s
}

func (s *SearchEventsRequest) SetAlertType(v int32) *SearchEventsRequest {
	s.AlertType = &v
	return s
}

func (s *SearchEventsRequest) SetIsTrigger(v int32) *SearchEventsRequest {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsRequest) SetStartTime(v int64) *SearchEventsRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEventsRequest) SetEndTime(v int64) *SearchEventsRequest {
	s.EndTime = &v
	return s
}

type SearchEventsResponseBody struct {
	PageBean  *SearchEventsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsTrigger *int32                            `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
}

func (s SearchEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBody) SetPageBean(v *SearchEventsResponseBodyPageBean) *SearchEventsResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchEventsResponseBody) SetRequestId(v string) *SearchEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEventsResponseBody) SetIsTrigger(v int32) *SearchEventsResponseBody {
	s.IsTrigger = &v
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
	EventTime  *int64    `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Links      []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	EventLevel *string   `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	AlertRule  *string   `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	Message    *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	AlertType  *int32    `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertName  *string   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Id         *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	AlertId    *int64    `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s SearchEventsResponseBodyPageBeanEvent) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBodyPageBeanEvent) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventTime(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.EventTime = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetLinks(v []*string) *SearchEventsResponseBodyPageBeanEvent {
	s.Links = v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventLevel(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.EventLevel = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertRule(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertRule = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetMessage(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.Message = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertType(v int32) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertType = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertName(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertName = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.Id = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertId = &v
	return s
}

type SearchEventsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchEventsResponse) SetBody(v *SearchEventsResponseBody) *SearchEventsResponse {
	s.Body = v
	return s
}

type SearchRetcodeAppByPageRequest struct {
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchRetcodeAppByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageRequest) SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest {
	s.RetcodeAppName = &v
	return s
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
	TotalCount  *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RetcodeApps []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
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

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetRetcodeApps(v []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.RetcodeApps = v
	return s
}

type SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps struct {
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid        *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Type = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppName(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppName = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUpdateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetCreateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.CreateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppId(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetPid(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Pid = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUserId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UserId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRegionId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RegionId = &v
	return s
}

type SearchRetcodeAppByPageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchRetcodeAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchRetcodeAppByPageResponse) SetBody(v *SearchRetcodeAppByPageResponseBody) *SearchRetcodeAppByPageResponse {
	s.Body = v
	return s
}

type SearchTraceAppByNameRequest struct {
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchTraceAppByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameRequest) SetTraceAppName(v string) *SearchTraceAppByNameRequest {
	s.TraceAppName = &v
	return s
}

func (s *SearchTraceAppByNameRequest) SetRegionId(v string) *SearchTraceAppByNameRequest {
	s.RegionId = &v
	return s
}

type SearchTraceAppByNameResponseBody struct {
	TraceApps []*SearchTraceAppByNameResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTraceAppByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBody) SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody {
	s.TraceApps = v
	return s
}

func (s *SearchTraceAppByNameResponseBody) SetRequestId(v string) *SearchTraceAppByNameResponseBody {
	s.RequestId = &v
	return s
}

type SearchTraceAppByNameResponseBodyTraceApps struct {
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchTraceAppByNameResponseBodyTraceApps) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetType(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppName(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUpdateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetLabels(v []*string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetShow(v bool) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetCreateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetPid(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppId(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUserId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UserId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetRegionId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

type SearchTraceAppByNameResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTraceAppByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchTraceAppByNameResponse) SetBody(v *SearchTraceAppByNameResponseBody) *SearchTraceAppByNameResponse {
	s.Body = v
	return s
}

type SearchTraceAppByPageRequest struct {
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchTraceAppByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageRequest) SetTraceAppName(v string) *SearchTraceAppByPageRequest {
	s.TraceAppName = &v
	return s
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
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetType(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppName(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUpdateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetLabels(v []*string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetShow(v bool) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetCreateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetPid(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppId(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUserId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UserId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetRegionId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.RegionId = &v
	return s
}

type SearchTraceAppByPageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTraceAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchTraceAppByPageResponse) SetBody(v *SearchTraceAppByPageResponseBody) *SearchTraceAppByPageResponse {
	s.Body = v
	return s
}

type SearchTracesRequest struct {
	StartTime        *int64                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *int64                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId         *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName      *string                                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	OperationName    *string                                `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	MinDuration      *int64                                 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	Reverse          *bool                                  `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp        *string                                `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Tag              []*SearchTracesRequestTag              `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ExclusionFilters []*SearchTracesRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
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

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

func (s *SearchTracesRequest) SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest {
	s.ExclusionFilters = v
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
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyTraceInfos) SetOperationName(v string) *SearchTracesResponseBodyTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceIp(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetDuration(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTimestamp(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceName(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTraceID(v string) *SearchTracesResponseBodyTraceInfos {
	s.TraceID = &v
	return s
}

type SearchTracesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
	s.Body = v
	return s
}

type SearchTracesByPageRequest struct {
	StartTime        *int64                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *int64                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId         *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName      *string                                      `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	OperationName    *string                                      `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	MinDuration      *int64                                       `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	Reverse          *bool                                        `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp        *string                                      `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ExclusionFilters []*SearchTracesByPageRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequest) SetStartTime(v int64) *SearchTracesByPageRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetEndTime(v int64) *SearchTracesByPageRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetRegionId(v string) *SearchTracesByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceName(v string) *SearchTracesByPageRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOperationName(v string) *SearchTracesByPageRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetMinDuration(v int64) *SearchTracesByPageRequest {
	s.MinDuration = &v
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

func (s *SearchTracesByPageRequest) SetPageNumber(v int32) *SearchTracesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageSize(v int32) *SearchTracesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageRequest) SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest {
	s.ExclusionFilters = v
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
	TraceInfos []*SearchTracesByPageResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchTracesByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTraceInfos(v []*SearchTracesByPageResponseBodyPageBeanTraceInfos) *SearchTracesByPageResponseBodyPageBean {
	s.TraceInfos = v
	return s
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

type SearchTracesByPageResponseBodyPageBeanTraceInfos struct {
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetOperationName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceIp(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetDuration(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTimestamp(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTraceID(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.TraceID = &v
	return s
}

type SearchTracesByPageResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTracesByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SearchTracesByPageResponse) SetBody(v *SearchTracesByPageResponseBody) *SearchTracesByPageResponse {
	s.Body = v
	return s
}

type SendCustomIncidentsRequest struct {
	Incidents   *string `json:"Incidents,omitempty" xml:"Incidents,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
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

func (s *SendCustomIncidentsRequest) SetRegionId(v string) *SendCustomIncidentsRequest {
	s.RegionId = &v
	return s
}

func (s *SendCustomIncidentsRequest) SetProxyUserId(v string) *SendCustomIncidentsRequest {
	s.ProxyUserId = &v
	return s
}

func (s *SendCustomIncidentsRequest) SetProductType(v string) *SendCustomIncidentsRequest {
	s.ProductType = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCustomIncidentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendCustomIncidentsResponse) SetBody(v *SendCustomIncidentsResponseBody) *SendCustomIncidentsResponse {
	s.Body = v
	return s
}

type SendMseIncidentRequest struct {
	Incidents   *string `json:"Incidents,omitempty" xml:"Incidents,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
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

func (s *SendMseIncidentRequest) SetProxyUserId(v string) *SendMseIncidentRequest {
	s.ProxyUserId = &v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMseIncidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRetcodeShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetRetcodeShareStatusResponse) SetBody(v *SetRetcodeShareStatusResponseBody) *SetRetcodeShareStatusResponse {
	s.Body = v
	return s
}

type StartAlertRequest struct {
	AlertId     *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *StartAlertRequest) SetProxyUserId(v string) *StartAlertRequest {
	s.ProxyUserId = &v
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartAlertResponse) SetBody(v *StartAlertResponseBody) *StartAlertResponse {
	s.Body = v
	return s
}

type StopAlertRequest struct {
	AlertId     *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *StopAlertRequest) SetProxyUserId(v string) *StopAlertRequest {
	s.ProxyUserId = &v
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopAlertResponse) SetBody(v *StopAlertResponseBody) *StopAlertResponse {
	s.Body = v
	return s
}

type UpdateAlertContactRequest struct {
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId         *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactId           *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s UpdateAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactRequest) SetContactName(v string) *UpdateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateAlertContactRequest) SetPhoneNum(v string) *UpdateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *UpdateAlertContactRequest) SetEmail(v string) *UpdateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateAlertContactRequest) SetDingRobotWebhookUrl(v string) *UpdateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *UpdateAlertContactRequest) SetSystemNoc(v bool) *UpdateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

func (s *UpdateAlertContactRequest) SetRegionId(v string) *UpdateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetProxyUserId(v string) *UpdateAlertContactRequest {
	s.ProxyUserId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetContactId(v int64) *UpdateAlertContactRequest {
	s.ContactId = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAlertContactResponse) SetBody(v *UpdateAlertContactResponseBody) *UpdateAlertContactResponse {
	s.Body = v
	return s
}

type UpdateAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactGroupId   *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
}

func (s UpdateAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupRequest) GoString() string {
	return s.String()
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

func (s *UpdateAlertContactGroupRequest) SetProxyUserId(v string) *UpdateAlertContactGroupRequest {
	s.ProxyUserId = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupId(v int64) *UpdateAlertContactGroupRequest {
	s.ContactGroupId = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAlertContactGroupResponse) SetBody(v *UpdateAlertContactGroupResponseBody) *UpdateAlertContactGroupResponse {
	s.Body = v
	return s
}

type UpdateAlertRuleRequest struct {
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	ProxyUserId         *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	AlertId             *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s UpdateAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequest) SetRegionId(v string) *UpdateAlertRuleRequest {
	s.RegionId = &v
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

func (s *UpdateAlertRuleRequest) SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetProxyUserId(v string) *UpdateAlertRuleRequest {
	s.ProxyUserId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetAlertId(v int64) *UpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

type UpdateAlertRuleResponseBody struct {
	AlertId   *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s *UpdateAlertRuleResponseBody) SetRequestId(v string) *UpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetData(v string) *UpdateAlertRuleResponseBody {
	s.Data = &v
	return s
}

type UpdateAlertRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAlertRuleResponse) SetBody(v *UpdateAlertRuleResponseBody) *UpdateAlertRuleResponse {
	s.Body = v
	return s
}

type UpdateWebhookRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s UpdateWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebhookRequest) SetContactName(v string) *UpdateWebhookRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateWebhookRequest) SetMethod(v string) *UpdateWebhookRequest {
	s.Method = &v
	return s
}

func (s *UpdateWebhookRequest) SetUrl(v string) *UpdateWebhookRequest {
	s.Url = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpParams(v string) *UpdateWebhookRequest {
	s.HttpParams = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpHeaders(v string) *UpdateWebhookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *UpdateWebhookRequest) SetRegionId(v string) *UpdateWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWebhookRequest) SetProxyUserId(v string) *UpdateWebhookRequest {
	s.ProxyUserId = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactId(v int64) *UpdateWebhookRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateWebhookRequest) SetBody(v string) *UpdateWebhookRequest {
	s.Body = &v
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
		"ap-northeast-1":        tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("arms.aliyuncs.com"),
		"cn-huhehaote":          tea.String("arms.aliyuncs.com"),
		"eu-central-1":          tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("arms.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("arms.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("arms.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("arms.aliyuncs.com"),
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGrafanaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGrafana"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddIntegrationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddIntegration"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyScenarioResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyScenario"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckDataConsistencyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckDataConsistency"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckServiceLinkedRoleForDeleting"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ConfigAppWithOptions(request *ConfigAppRequest, runtime *util.RuntimeOptions) (_result *ConfigAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigApp"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlertContact"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlertContactGroup"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateRetcodeAppWithOptions(request *CreateRetcodeAppRequest, runtime *util.RuntimeOptions) (_result *CreateRetcodeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRetcodeAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRetcodeApp"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWehookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWehook"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlertContact"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlertContactGroup"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlertRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlertRules"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteRetcodeAppWithOptions(request *DeleteRetcodeAppRequest, runtime *util.RuntimeOptions) (_result *DeleteRetcodeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRetcodeAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRetcodeApp"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScenario"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTraceAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTraceApp"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDispatchRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDispatchRule"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTraceLicenseKeyWithOptions(request *DescribeTraceLicenseKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeTraceLicenseKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTraceLicenseKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTraceLicenseKey"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTraceLocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTraceLocation"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ExportPrometheusRulesWithOptions(request *ExportPrometheusRulesRequest, runtime *util.RuntimeOptions) (_result *ExportPrometheusRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportPrometheusRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportPrometheusRules"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &GetAgentDownloadUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentDownloadUrl"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppApiByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAppApiByPage"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConsistencySnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConsistencySnapshot"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIntegrationTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIntegrationToken"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMultipleTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMultipleTrace"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPrometheusApiTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPrometheusApiToken"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRetcodeShareUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRetcodeShareUrl"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStack"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrace"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTraceAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTraceApp"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportAppAlertRules"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportCustomAlertRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportCustomAlertRules"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportPrometheusRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportPrometheusRules"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListClusterFromGrafanaWithOptions(request *ListClusterFromGrafanaRequest, runtime *util.RuntimeOptions) (_result *ListClusterFromGrafanaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterFromGrafanaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterFromGrafana"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDashboardsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDashboards"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListPromClustersWithOptions(request *ListPromClustersRequest, runtime *util.RuntimeOptions) (_result *ListPromClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPromClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPromClusters"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListRetcodeAppsWithOptions(request *ListRetcodeAppsRequest, runtime *util.RuntimeOptions) (_result *ListRetcodeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRetcodeAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRetcodeApps"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListScenarioResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScenario"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTraceAppsWithOptions(request *ListTraceAppsRequest, runtime *util.RuntimeOptions) (_result *ListTraceAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTraceAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTraceApps"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) OpenArmsServiceWithOptions(request *OpenArmsServiceRequest, runtime *util.RuntimeOptions) (_result *OpenArmsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenArmsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenArmsService"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryDatasetWithOptions(request *QueryDatasetRequest, runtime *util.RuntimeOptions) (_result *QueryDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDataset"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMetric"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMetricByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMetricByPage"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveTraceAppConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveTraceAppConfig"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchAlertContact"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchAlertContactGroup"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchAlertHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchAlertHistories"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchAlertRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchAlertRules"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchEvents"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchRetcodeAppByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchRetcodeAppByPage"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTraceAppByNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTraceAppByName"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTraceAppByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTraceAppByPage"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTraces"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTracesByPage"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCustomIncidentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCustomIncidents"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendMseIncidentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendMseIncident"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetRetcodeShareStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetRetcodeShareStatus"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartAlert"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopAlert"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAlertContactResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAlertContact"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAlertContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAlertContactGroup"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAlertRule"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateWebhookWithOptions(request *UpdateWebhookRequest, runtime *util.RuntimeOptions) (_result *UpdateWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateWebhook"), tea.String("2019-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

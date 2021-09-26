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

type AddLogMaskConfigRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskName        *string `json:"MaskName,omitempty" xml:"MaskName,omitempty"`
	MaskRegex       *string `json:"MaskRegex,omitempty" xml:"MaskRegex,omitempty"`
	MaskTxt         *string `json:"MaskTxt,omitempty" xml:"MaskTxt,omitempty"`
	MaskDescription *string `json:"MaskDescription,omitempty" xml:"MaskDescription,omitempty"`
}

func (s AddLogMaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLogMaskConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLogMaskConfigRequest) SetRegionId(v string) *AddLogMaskConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLogMaskConfigRequest) SetInstanceId(v string) *AddLogMaskConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *AddLogMaskConfigRequest) SetMaskName(v string) *AddLogMaskConfigRequest {
	s.MaskName = &v
	return s
}

func (s *AddLogMaskConfigRequest) SetMaskRegex(v string) *AddLogMaskConfigRequest {
	s.MaskRegex = &v
	return s
}

func (s *AddLogMaskConfigRequest) SetMaskTxt(v string) *AddLogMaskConfigRequest {
	s.MaskTxt = &v
	return s
}

func (s *AddLogMaskConfigRequest) SetMaskDescription(v string) *AddLogMaskConfigRequest {
	s.MaskDescription = &v
	return s
}

type AddLogMaskConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLogMaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLogMaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLogMaskConfigResponseBody) SetRequestId(v string) *AddLogMaskConfigResponseBody {
	s.RequestId = &v
	return s
}

type AddLogMaskConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLogMaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLogMaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLogMaskConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLogMaskConfigResponse) SetHeaders(v map[string]*string) *AddLogMaskConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLogMaskConfigResponse) SetBody(v *AddLogMaskConfigResponseBody) *AddLogMaskConfigResponse {
	s.Body = v
	return s
}

type AssociateDbToRuleRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleIds         *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	RuleDbRelations *string `json:"RuleDbRelations,omitempty" xml:"RuleDbRelations,omitempty"`
	OperType        *string `json:"OperType,omitempty" xml:"OperType,omitempty"`
}

func (s AssociateDbToRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateDbToRuleRequest) GoString() string {
	return s.String()
}

func (s *AssociateDbToRuleRequest) SetRegionId(v string) *AssociateDbToRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateDbToRuleRequest) SetInstanceId(v string) *AssociateDbToRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateDbToRuleRequest) SetRuleIds(v string) *AssociateDbToRuleRequest {
	s.RuleIds = &v
	return s
}

func (s *AssociateDbToRuleRequest) SetRuleDbRelations(v string) *AssociateDbToRuleRequest {
	s.RuleDbRelations = &v
	return s
}

func (s *AssociateDbToRuleRequest) SetOperType(v string) *AssociateDbToRuleRequest {
	s.OperType = &v
	return s
}

type AssociateDbToRuleResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateDbToRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateDbToRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateDbToRuleResponseBody) SetServerData(v string) *AssociateDbToRuleResponseBody {
	s.ServerData = &v
	return s
}

func (s *AssociateDbToRuleResponseBody) SetRequestId(v string) *AssociateDbToRuleResponseBody {
	s.RequestId = &v
	return s
}

type AssociateDbToRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateDbToRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateDbToRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateDbToRuleResponse) GoString() string {
	return s.String()
}

func (s *AssociateDbToRuleResponse) SetHeaders(v map[string]*string) *AssociateDbToRuleResponse {
	s.Headers = v
	return s
}

func (s *AssociateDbToRuleResponse) SetBody(v *AssociateDbToRuleResponseBody) *AssociateDbToRuleResponse {
	s.Body = v
	return s
}

type AssociateRuleToDbRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JsonParam  *string `json:"JsonParam,omitempty" xml:"JsonParam,omitempty"`
}

func (s AssociateRuleToDbRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateRuleToDbRequest) GoString() string {
	return s.String()
}

func (s *AssociateRuleToDbRequest) SetRegionId(v string) *AssociateRuleToDbRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateRuleToDbRequest) SetInstanceId(v string) *AssociateRuleToDbRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateRuleToDbRequest) SetJsonParam(v string) *AssociateRuleToDbRequest {
	s.JsonParam = &v
	return s
}

type AssociateRuleToDbResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *AssociateRuleToDbResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s AssociateRuleToDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateRuleToDbResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateRuleToDbResponseBody) SetRequestId(v string) *AssociateRuleToDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateRuleToDbResponseBody) SetServerData(v *AssociateRuleToDbResponseBodyServerData) *AssociateRuleToDbResponseBody {
	s.ServerData = v
	return s
}

type AssociateRuleToDbResponseBodyServerData struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
}

func (s AssociateRuleToDbResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s AssociateRuleToDbResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *AssociateRuleToDbResponseBodyServerData) SetJsonResult(v string) *AssociateRuleToDbResponseBodyServerData {
	s.JsonResult = &v
	return s
}

type AssociateRuleToDbResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateRuleToDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateRuleToDbResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateRuleToDbResponse) GoString() string {
	return s.String()
}

func (s *AssociateRuleToDbResponse) SetHeaders(v map[string]*string) *AssociateRuleToDbResponse {
	s.Headers = v
	return s
}

func (s *AssociateRuleToDbResponse) SetBody(v *AssociateRuleToDbResponseBody) *AssociateRuleToDbResponse {
	s.Body = v
	return s
}

type ChangeAgentStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AgentId     *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatus *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
}

func (s ChangeAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeAgentStatusRequest) SetRegionId(v string) *ChangeAgentStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeAgentStatusRequest) SetInstanceId(v string) *ChangeAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeAgentStatusRequest) SetAgentId(v string) *ChangeAgentStatusRequest {
	s.AgentId = &v
	return s
}

func (s *ChangeAgentStatusRequest) SetAgentStatus(v int32) *ChangeAgentStatusRequest {
	s.AgentStatus = &v
	return s
}

type ChangeAgentStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAgentStatusResponseBody) SetRequestId(v string) *ChangeAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

type ChangeAgentStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeAgentStatusResponse) SetHeaders(v map[string]*string) *ChangeAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeAgentStatusResponse) SetBody(v *ChangeAgentStatusResponseBody) *ChangeAgentStatusResponse {
	s.Body = v
	return s
}

type ChangeLogMaskConfigStateRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskId     *int32  `json:"MaskId,omitempty" xml:"MaskId,omitempty"`
	MaskState  *int32  `json:"MaskState,omitempty" xml:"MaskState,omitempty"`
}

func (s ChangeLogMaskConfigStateRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeLogMaskConfigStateRequest) GoString() string {
	return s.String()
}

func (s *ChangeLogMaskConfigStateRequest) SetRegionId(v string) *ChangeLogMaskConfigStateRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeLogMaskConfigStateRequest) SetInstanceId(v string) *ChangeLogMaskConfigStateRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeLogMaskConfigStateRequest) SetMaskId(v int32) *ChangeLogMaskConfigStateRequest {
	s.MaskId = &v
	return s
}

func (s *ChangeLogMaskConfigStateRequest) SetMaskState(v int32) *ChangeLogMaskConfigStateRequest {
	s.MaskState = &v
	return s
}

type ChangeLogMaskConfigStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeLogMaskConfigStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeLogMaskConfigStateResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeLogMaskConfigStateResponseBody) SetRequestId(v string) *ChangeLogMaskConfigStateResponseBody {
	s.RequestId = &v
	return s
}

type ChangeLogMaskConfigStateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeLogMaskConfigStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeLogMaskConfigStateResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeLogMaskConfigStateResponse) GoString() string {
	return s.String()
}

func (s *ChangeLogMaskConfigStateResponse) SetHeaders(v map[string]*string) *ChangeLogMaskConfigStateResponse {
	s.Headers = v
	return s
}

func (s *ChangeLogMaskConfigStateResponse) SetBody(v *ChangeLogMaskConfigStateResponseBody) *ChangeLogMaskConfigStateResponse {
	s.Body = v
	return s
}

type ChangeRulePriorityRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	RuleInfos  *string `json:"RuleInfos,omitempty" xml:"RuleInfos,omitempty"`
}

func (s ChangeRulePriorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeRulePriorityRequest) GoString() string {
	return s.String()
}

func (s *ChangeRulePriorityRequest) SetRegionId(v string) *ChangeRulePriorityRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeRulePriorityRequest) SetInstanceId(v string) *ChangeRulePriorityRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeRulePriorityRequest) SetDbId(v string) *ChangeRulePriorityRequest {
	s.DbId = &v
	return s
}

func (s *ChangeRulePriorityRequest) SetRuleInfos(v string) *ChangeRulePriorityRequest {
	s.RuleInfos = &v
	return s
}

type ChangeRulePriorityResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeRulePriorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeRulePriorityResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeRulePriorityResponseBody) SetServerData(v string) *ChangeRulePriorityResponseBody {
	s.ServerData = &v
	return s
}

func (s *ChangeRulePriorityResponseBody) SetRequestId(v string) *ChangeRulePriorityResponseBody {
	s.RequestId = &v
	return s
}

type ChangeRulePriorityResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeRulePriorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeRulePriorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeRulePriorityResponse) GoString() string {
	return s.String()
}

func (s *ChangeRulePriorityResponse) SetHeaders(v map[string]*string) *ChangeRulePriorityResponse {
	s.Headers = v
	return s
}

func (s *ChangeRulePriorityResponse) SetBody(v *ChangeRulePriorityResponseBody) *ChangeRulePriorityResponse {
	s.Body = v
	return s
}

type ChangeRuleStatusRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JsonParam  *string `json:"JsonParam,omitempty" xml:"JsonParam,omitempty"`
}

func (s ChangeRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeRuleStatusRequest) SetRegionId(v string) *ChangeRuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeRuleStatusRequest) SetInstanceId(v string) *ChangeRuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeRuleStatusRequest) SetJsonParam(v string) *ChangeRuleStatusRequest {
	s.JsonParam = &v
	return s
}

type ChangeRuleStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeRuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeRuleStatusResponseBody) SetRequestId(v string) *ChangeRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type ChangeRuleStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeRuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeRuleStatusResponse) SetHeaders(v map[string]*string) *ChangeRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeRuleStatusResponse) SetBody(v *ChangeRuleStatusResponseBody) *ChangeRuleStatusResponse {
	s.Body = v
	return s
}

type CheckMailRegisteredRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mail       *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
}

func (s CheckMailRegisteredRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMailRegisteredRequest) GoString() string {
	return s.String()
}

func (s *CheckMailRegisteredRequest) SetRegionId(v string) *CheckMailRegisteredRequest {
	s.RegionId = &v
	return s
}

func (s *CheckMailRegisteredRequest) SetInstanceId(v string) *CheckMailRegisteredRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckMailRegisteredRequest) SetMail(v string) *CheckMailRegisteredRequest {
	s.Mail = &v
	return s
}

type CheckMailRegisteredResponseBody struct {
	Registered *bool   `json:"Registered,omitempty" xml:"Registered,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMailRegisteredResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMailRegisteredResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMailRegisteredResponseBody) SetRegistered(v bool) *CheckMailRegisteredResponseBody {
	s.Registered = &v
	return s
}

func (s *CheckMailRegisteredResponseBody) SetRequestId(v string) *CheckMailRegisteredResponseBody {
	s.RequestId = &v
	return s
}

type CheckMailRegisteredResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckMailRegisteredResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMailRegisteredResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMailRegisteredResponse) GoString() string {
	return s.String()
}

func (s *CheckMailRegisteredResponse) SetHeaders(v map[string]*string) *CheckMailRegisteredResponse {
	s.Headers = v
	return s
}

func (s *CheckMailRegisteredResponse) SetBody(v *CheckMailRegisteredResponseBody) *CheckMailRegisteredResponse {
	s.Body = v
	return s
}

type ClearAgentRecordsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AgentIds   []*string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
}

func (s ClearAgentRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearAgentRecordsRequest) GoString() string {
	return s.String()
}

func (s *ClearAgentRecordsRequest) SetRegionId(v string) *ClearAgentRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *ClearAgentRecordsRequest) SetInstanceId(v string) *ClearAgentRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearAgentRecordsRequest) SetAgentIds(v []*string) *ClearAgentRecordsRequest {
	s.AgentIds = v
	return s
}

type ClearAgentRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearAgentRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearAgentRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAgentRecordsResponseBody) SetRequestId(v string) *ClearAgentRecordsResponseBody {
	s.RequestId = &v
	return s
}

type ClearAgentRecordsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearAgentRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearAgentRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearAgentRecordsResponse) GoString() string {
	return s.String()
}

func (s *ClearAgentRecordsResponse) SetHeaders(v map[string]*string) *ClearAgentRecordsResponse {
	s.Headers = v
	return s
}

func (s *ClearAgentRecordsResponse) SetBody(v *ClearAgentRecordsResponseBody) *ClearAgentRecordsResponse {
	s.Body = v
	return s
}

type ConfigInstanceNetworkRequest struct {
	InstanceId          *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PublicAccessControl *int32    `json:"PublicAccessControl,omitempty" xml:"PublicAccessControl,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PrivateWhiteList    []*string `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	PublicWhiteList     []*string `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	SecurityGroupIds    []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s ConfigInstanceNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceNetworkRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceNetworkRequest) SetInstanceId(v string) *ConfigInstanceNetworkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetPublicAccessControl(v int32) *ConfigInstanceNetworkRequest {
	s.PublicAccessControl = &v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetRegionId(v string) *ConfigInstanceNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetPrivateWhiteList(v []*string) *ConfigInstanceNetworkRequest {
	s.PrivateWhiteList = v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetPublicWhiteList(v []*string) *ConfigInstanceNetworkRequest {
	s.PublicWhiteList = v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetSecurityGroupIds(v []*string) *ConfigInstanceNetworkRequest {
	s.SecurityGroupIds = v
	return s
}

type ConfigInstanceNetworkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceNetworkResponseBody) SetRequestId(v string) *ConfigInstanceNetworkResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceNetworkResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigInstanceNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigInstanceNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceNetworkResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceNetworkResponse) SetHeaders(v map[string]*string) *ConfigInstanceNetworkResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceNetworkResponse) SetBody(v *ConfigInstanceNetworkResponseBody) *ConfigInstanceNetworkResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbName        *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbInstanceId  *string   `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	AssetType     *int32    `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	DbType        *int32    `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbVersion     *int32    `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	DbCertificate *string   `json:"DbCertificate,omitempty" xml:"DbCertificate,omitempty"`
	DbUsername    *string   `json:"DbUsername,omitempty" xml:"DbUsername,omitempty"`
	DbPassword    *string   `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbAddresses   []*string `json:"DbAddresses,omitempty" xml:"DbAddresses,omitempty" type:"Repeated"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetRegionId(v string) *CreateDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataSourceRequest) SetInstanceId(v string) *CreateDataSourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbName(v string) *CreateDataSourceRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbInstanceId(v string) *CreateDataSourceRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDataSourceRequest) SetAssetType(v int32) *CreateDataSourceRequest {
	s.AssetType = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbType(v int32) *CreateDataSourceRequest {
	s.DbType = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbVersion(v int32) *CreateDataSourceRequest {
	s.DbVersion = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbCertificate(v string) *CreateDataSourceRequest {
	s.DbCertificate = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbUsername(v string) *CreateDataSourceRequest {
	s.DbUsername = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbPassword(v string) *CreateDataSourceRequest {
	s.DbPassword = &v
	return s
}

func (s *CreateDataSourceRequest) SetDbAddresses(v []*string) *CreateDataSourceRequest {
	s.DbAddresses = v
	return s
}

type CreateDataSourceResponseBody struct {
	DbId      *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetDbId(v int32) *CreateDataSourceResponseBody {
	s.DbId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type CreateGradeProtectionReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s CreateGradeProtectionReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGradeProtectionReportRequest) GoString() string {
	return s.String()
}

func (s *CreateGradeProtectionReportRequest) SetRegionId(v string) *CreateGradeProtectionReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGradeProtectionReportRequest) SetInstanceId(v string) *CreateGradeProtectionReportRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGradeProtectionReportRequest) SetDbId(v int32) *CreateGradeProtectionReportRequest {
	s.DbId = &v
	return s
}

func (s *CreateGradeProtectionReportRequest) SetBeginDate(v string) *CreateGradeProtectionReportRequest {
	s.BeginDate = &v
	return s
}

func (s *CreateGradeProtectionReportRequest) SetEndDate(v string) *CreateGradeProtectionReportRequest {
	s.EndDate = &v
	return s
}

type CreateGradeProtectionReportResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *CreateGradeProtectionReportResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s CreateGradeProtectionReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGradeProtectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGradeProtectionReportResponseBody) SetRequestId(v string) *CreateGradeProtectionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGradeProtectionReportResponseBody) SetServerData(v *CreateGradeProtectionReportResponseBodyServerData) *CreateGradeProtectionReportResponseBody {
	s.ServerData = v
	return s
}

type CreateGradeProtectionReportResponseBodyServerData struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreateGradeProtectionReportResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s CreateGradeProtectionReportResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *CreateGradeProtectionReportResponseBodyServerData) SetFileName(v string) *CreateGradeProtectionReportResponseBodyServerData {
	s.FileName = &v
	return s
}

type CreateGradeProtectionReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGradeProtectionReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGradeProtectionReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGradeProtectionReportResponse) GoString() string {
	return s.String()
}

func (s *CreateGradeProtectionReportResponse) SetHeaders(v map[string]*string) *CreateGradeProtectionReportResponse {
	s.Headers = v
	return s
}

func (s *CreateGradeProtectionReportResponse) SetBody(v *CreateGradeProtectionReportResponseBody) *CreateGradeProtectionReportResponse {
	s.Body = v
	return s
}

type CreateIntegratedReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s CreateIntegratedReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedReportRequest) GoString() string {
	return s.String()
}

func (s *CreateIntegratedReportRequest) SetRegionId(v string) *CreateIntegratedReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIntegratedReportRequest) SetInstanceId(v string) *CreateIntegratedReportRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIntegratedReportRequest) SetDbId(v int32) *CreateIntegratedReportRequest {
	s.DbId = &v
	return s
}

func (s *CreateIntegratedReportRequest) SetBeginDate(v string) *CreateIntegratedReportRequest {
	s.BeginDate = &v
	return s
}

func (s *CreateIntegratedReportRequest) SetEndDate(v string) *CreateIntegratedReportRequest {
	s.EndDate = &v
	return s
}

type CreateIntegratedReportResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *CreateIntegratedReportResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s CreateIntegratedReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntegratedReportResponseBody) SetRequestId(v string) *CreateIntegratedReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntegratedReportResponseBody) SetServerData(v *CreateIntegratedReportResponseBodyServerData) *CreateIntegratedReportResponseBody {
	s.ServerData = v
	return s
}

type CreateIntegratedReportResponseBodyServerData struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreateIntegratedReportResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedReportResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *CreateIntegratedReportResponseBodyServerData) SetFileName(v string) *CreateIntegratedReportResponseBodyServerData {
	s.FileName = &v
	return s
}

type CreateIntegratedReportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIntegratedReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntegratedReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedReportResponse) GoString() string {
	return s.String()
}

func (s *CreateIntegratedReportResponse) SetHeaders(v map[string]*string) *CreateIntegratedReportResponse {
	s.Headers = v
	return s
}

func (s *CreateIntegratedReportResponse) SetBody(v *CreateIntegratedReportResponseBody) *CreateIntegratedReportResponse {
	s.Body = v
	return s
}

type CreateLogAlarmTaskRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskName      *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ToMailList    []*string `json:"ToMailList,omitempty" xml:"ToMailList,omitempty" type:"Repeated"`
	DbIdList      []*string `json:"DbIdList,omitempty" xml:"DbIdList,omitempty" type:"Repeated"`
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
}

func (s CreateLogAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLogAlarmTaskRequest) SetRegionId(v string) *CreateLogAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLogAlarmTaskRequest) SetInstanceId(v string) *CreateLogAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLogAlarmTaskRequest) SetTaskName(v string) *CreateLogAlarmTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateLogAlarmTaskRequest) SetToMailList(v []*string) *CreateLogAlarmTaskRequest {
	s.ToMailList = v
	return s
}

func (s *CreateLogAlarmTaskRequest) SetDbIdList(v []*string) *CreateLogAlarmTaskRequest {
	s.DbIdList = v
	return s
}

func (s *CreateLogAlarmTaskRequest) SetRiskLevelList(v []*string) *CreateLogAlarmTaskRequest {
	s.RiskLevelList = v
	return s
}

type CreateLogAlarmTaskResponseBody struct {
	TaskId    *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLogAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogAlarmTaskResponseBody) SetTaskId(v int32) *CreateLogAlarmTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateLogAlarmTaskResponseBody) SetRequestId(v string) *CreateLogAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateLogAlarmTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLogAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLogAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLogAlarmTaskResponse) SetHeaders(v map[string]*string) *CreateLogAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLogAlarmTaskResponse) SetBody(v *CreateLogAlarmTaskResponseBody) *CreateLogAlarmTaskResponse {
	s.Body = v
	return s
}

type CreatePCIReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s CreatePCIReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePCIReportRequest) GoString() string {
	return s.String()
}

func (s *CreatePCIReportRequest) SetRegionId(v string) *CreatePCIReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePCIReportRequest) SetInstanceId(v string) *CreatePCIReportRequest {
	s.InstanceId = &v
	return s
}

func (s *CreatePCIReportRequest) SetDbId(v int32) *CreatePCIReportRequest {
	s.DbId = &v
	return s
}

func (s *CreatePCIReportRequest) SetBeginDate(v string) *CreatePCIReportRequest {
	s.BeginDate = &v
	return s
}

func (s *CreatePCIReportRequest) SetEndDate(v string) *CreatePCIReportRequest {
	s.EndDate = &v
	return s
}

type CreatePCIReportResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *CreatePCIReportResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s CreatePCIReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePCIReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePCIReportResponseBody) SetRequestId(v string) *CreatePCIReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePCIReportResponseBody) SetServerData(v *CreatePCIReportResponseBodyServerData) *CreatePCIReportResponseBody {
	s.ServerData = v
	return s
}

type CreatePCIReportResponseBodyServerData struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreatePCIReportResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s CreatePCIReportResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *CreatePCIReportResponseBodyServerData) SetFileName(v string) *CreatePCIReportResponseBodyServerData {
	s.FileName = &v
	return s
}

type CreatePCIReportResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePCIReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePCIReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePCIReportResponse) GoString() string {
	return s.String()
}

func (s *CreatePCIReportResponse) SetHeaders(v map[string]*string) *CreatePCIReportResponse {
	s.Headers = v
	return s
}

func (s *CreatePCIReportResponse) SetBody(v *CreatePCIReportResponseBody) *CreatePCIReportResponse {
	s.Body = v
	return s
}

type CreateReportPushTaskRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ScheduleType *string   `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	ScheduleTime *string   `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	JobStatus    *int32    `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	ReportType   []*string `json:"ReportType,omitempty" xml:"ReportType,omitempty" type:"Repeated"`
	DbList       []*string `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	EmailList    []*string `json:"EmailList,omitempty" xml:"EmailList,omitempty" type:"Repeated"`
}

func (s CreateReportPushTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReportPushTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateReportPushTaskRequest) SetRegionId(v string) *CreateReportPushTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateReportPushTaskRequest) SetInstanceId(v string) *CreateReportPushTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateReportPushTaskRequest) SetScheduleType(v string) *CreateReportPushTaskRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateReportPushTaskRequest) SetScheduleTime(v string) *CreateReportPushTaskRequest {
	s.ScheduleTime = &v
	return s
}

func (s *CreateReportPushTaskRequest) SetJobStatus(v int32) *CreateReportPushTaskRequest {
	s.JobStatus = &v
	return s
}

func (s *CreateReportPushTaskRequest) SetReportType(v []*string) *CreateReportPushTaskRequest {
	s.ReportType = v
	return s
}

func (s *CreateReportPushTaskRequest) SetDbList(v []*string) *CreateReportPushTaskRequest {
	s.DbList = v
	return s
}

func (s *CreateReportPushTaskRequest) SetEmailList(v []*string) *CreateReportPushTaskRequest {
	s.EmailList = v
	return s
}

type CreateReportPushTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReportPushTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReportPushTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportPushTaskResponseBody) SetRequestId(v string) *CreateReportPushTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateReportPushTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateReportPushTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReportPushTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReportPushTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateReportPushTaskResponse) SetHeaders(v map[string]*string) *CreateReportPushTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateReportPushTaskResponse) SetBody(v *CreateReportPushTaskResponseBody) *CreateReportPushTaskResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleInfo              *string `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty"`
	RuleDbRelation        *string `json:"RuleDbRelation,omitempty" xml:"RuleDbRelation,omitempty"`
	RuleGroup             *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	EffectCurrentDBStatus *string `json:"EffectCurrentDBStatus,omitempty" xml:"EffectCurrentDBStatus,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetRegionId(v string) *CreateRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRuleRequest) SetInstanceId(v string) *CreateRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRuleRequest) SetRuleInfo(v string) *CreateRuleRequest {
	s.RuleInfo = &v
	return s
}

func (s *CreateRuleRequest) SetRuleDbRelation(v string) *CreateRuleRequest {
	s.RuleDbRelation = &v
	return s
}

func (s *CreateRuleRequest) SetRuleGroup(v string) *CreateRuleRequest {
	s.RuleGroup = &v
	return s
}

func (s *CreateRuleRequest) SetEffectCurrentDBStatus(v string) *CreateRuleRequest {
	s.EffectCurrentDBStatus = &v
	return s
}

type CreateRuleResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetServerData(v string) *CreateRuleResponseBody {
	s.ServerData = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateRuleGroupRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
}

func (s CreateRuleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleGroupRequest) SetRegionId(v string) *CreateRuleGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRuleGroupRequest) SetInstanceId(v string) *CreateRuleGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRuleGroupRequest) SetGroupName(v string) *CreateRuleGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateRuleGroupRequest) SetGroupType(v string) *CreateRuleGroupRequest {
	s.GroupType = &v
	return s
}

type CreateRuleGroupResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRuleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleGroupResponseBody) SetServerData(v string) *CreateRuleGroupResponseBody {
	s.ServerData = &v
	return s
}

func (s *CreateRuleGroupResponseBody) SetRequestId(v string) *CreateRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateRuleGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleGroupResponse) SetHeaders(v map[string]*string) *CreateRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleGroupResponse) SetBody(v *CreateRuleGroupResponseBody) *CreateRuleGroupResponse {
	s.Body = v
	return s
}

type CreateSOXReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s CreateSOXReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSOXReportRequest) GoString() string {
	return s.String()
}

func (s *CreateSOXReportRequest) SetRegionId(v string) *CreateSOXReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSOXReportRequest) SetInstanceId(v string) *CreateSOXReportRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSOXReportRequest) SetDbId(v int32) *CreateSOXReportRequest {
	s.DbId = &v
	return s
}

func (s *CreateSOXReportRequest) SetBeginDate(v string) *CreateSOXReportRequest {
	s.BeginDate = &v
	return s
}

func (s *CreateSOXReportRequest) SetEndDate(v string) *CreateSOXReportRequest {
	s.EndDate = &v
	return s
}

type CreateSOXReportResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *CreateSOXReportResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s CreateSOXReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSOXReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSOXReportResponseBody) SetRequestId(v string) *CreateSOXReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSOXReportResponseBody) SetServerData(v *CreateSOXReportResponseBodyServerData) *CreateSOXReportResponseBody {
	s.ServerData = v
	return s
}

type CreateSOXReportResponseBodyServerData struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreateSOXReportResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s CreateSOXReportResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *CreateSOXReportResponseBodyServerData) SetFileName(v string) *CreateSOXReportResponseBodyServerData {
	s.FileName = &v
	return s
}

type CreateSOXReportResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSOXReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSOXReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSOXReportResponse) GoString() string {
	return s.String()
}

func (s *CreateSOXReportResponse) SetHeaders(v map[string]*string) *CreateSOXReportResponse {
	s.Headers = v
	return s
}

func (s *CreateSOXReportResponse) SetBody(v *CreateSOXReportResponseBody) *CreateSOXReportResponse {
	s.Body = v
	return s
}

type CreateSqlRuleRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JsonParam  *string `json:"JsonParam,omitempty" xml:"JsonParam,omitempty"`
}

func (s CreateSqlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlRuleRequest) SetRegionId(v string) *CreateSqlRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSqlRuleRequest) SetInstanceId(v string) *CreateSqlRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSqlRuleRequest) SetJsonParam(v string) *CreateSqlRuleRequest {
	s.JsonParam = &v
	return s
}

type CreateSqlRuleResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *CreateSqlRuleResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s CreateSqlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlRuleResponseBody) SetRequestId(v string) *CreateSqlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlRuleResponseBody) SetServerData(v *CreateSqlRuleResponseBodyServerData) *CreateSqlRuleResponseBody {
	s.ServerData = v
	return s
}

type CreateSqlRuleResponseBodyServerData struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
}

func (s CreateSqlRuleResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlRuleResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *CreateSqlRuleResponseBodyServerData) SetJsonResult(v string) *CreateSqlRuleResponseBodyServerData {
	s.JsonResult = &v
	return s
}

type CreateSqlRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSqlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSqlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlRuleResponse) SetHeaders(v map[string]*string) *CreateSqlRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlRuleResponse) SetBody(v *CreateSqlRuleResponseBody) *CreateSqlRuleResponse {
	s.Body = v
	return s
}

type CreateSystemAlarmTaskRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskName   *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ToMailList []*string `json:"ToMailList,omitempty" xml:"ToMailList,omitempty" type:"Repeated"`
}

func (s CreateSystemAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSystemAlarmTaskRequest) SetRegionId(v string) *CreateSystemAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSystemAlarmTaskRequest) SetInstanceId(v string) *CreateSystemAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSystemAlarmTaskRequest) SetTaskName(v string) *CreateSystemAlarmTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSystemAlarmTaskRequest) SetToMailList(v []*string) *CreateSystemAlarmTaskRequest {
	s.ToMailList = v
	return s
}

type CreateSystemAlarmTaskResponseBody struct {
	TaskId    *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSystemAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSystemAlarmTaskResponseBody) SetTaskId(v int32) *CreateSystemAlarmTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateSystemAlarmTaskResponseBody) SetRequestId(v string) *CreateSystemAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateSystemAlarmTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSystemAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSystemAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSystemAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSystemAlarmTaskResponse) SetHeaders(v map[string]*string) *CreateSystemAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSystemAlarmTaskResponse) SetBody(v *CreateSystemAlarmTaskResponseBody) *CreateSystemAlarmTaskResponse {
	s.Body = v
	return s
}

type DeleteAlarmTaskRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskIds    []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DeleteAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmTaskRequest) SetRegionId(v string) *DeleteAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlarmTaskRequest) SetInstanceId(v string) *DeleteAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAlarmTaskRequest) SetTaskIds(v []*string) *DeleteAlarmTaskRequest {
	s.TaskIds = v
	return s
}

type DeleteAlarmTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmTaskResponseBody) SetRequestId(v string) *DeleteAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlarmTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmTaskResponse) SetHeaders(v map[string]*string) *DeleteAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmTaskResponse) SetBody(v *DeleteAlarmTaskResponseBody) *DeleteAlarmTaskResponse {
	s.Body = v
	return s
}

type DeleteDataSourceRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbIds      []*string `json:"DbIds,omitempty" xml:"DbIds,omitempty" type:"Repeated"`
}

func (s DeleteDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) SetRegionId(v string) *DeleteDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetInstanceId(v string) *DeleteDataSourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetDbIds(v []*string) *DeleteDataSourceRequest {
	s.DbIds = v
	return s
}

type DeleteDataSourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type DeleteReportPushTaskRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *int32  `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteReportPushTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportPushTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteReportPushTaskRequest) SetRegionId(v string) *DeleteReportPushTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteReportPushTaskRequest) SetInstanceId(v string) *DeleteReportPushTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteReportPushTaskRequest) SetJobId(v int32) *DeleteReportPushTaskRequest {
	s.JobId = &v
	return s
}

type DeleteReportPushTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReportPushTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportPushTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReportPushTaskResponseBody) SetRequestId(v string) *DeleteReportPushTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteReportPushTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteReportPushTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteReportPushTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportPushTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteReportPushTaskResponse) SetHeaders(v map[string]*string) *DeleteReportPushTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteReportPushTaskResponse) SetBody(v *DeleteReportPushTaskResponseBody) *DeleteReportPushTaskResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType   *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetRegionId(v string) *DeleteRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRuleRequest) SetInstanceId(v string) *DeleteRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v string) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleType(v string) *DeleteRuleRequest {
	s.RuleType = &v
	return s
}

type DeleteRuleResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetServerData(v string) *DeleteRuleResponseBody {
	s.ServerData = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DeleteRuleGroupRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupKeyId *string `json:"GroupKeyId,omitempty" xml:"GroupKeyId,omitempty"`
}

func (s DeleteRuleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleGroupRequest) SetRegionId(v string) *DeleteRuleGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRuleGroupRequest) SetInstanceId(v string) *DeleteRuleGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRuleGroupRequest) SetGroupKeyId(v string) *DeleteRuleGroupRequest {
	s.GroupKeyId = &v
	return s
}

type DeleteRuleGroupResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleGroupResponseBody) SetServerData(v string) *DeleteRuleGroupResponseBody {
	s.ServerData = &v
	return s
}

func (s *DeleteRuleGroupResponseBody) SetRequestId(v string) *DeleteRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleGroupResponse) SetHeaders(v map[string]*string) *DeleteRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleGroupResponse) SetBody(v *DeleteRuleGroupResponseBody) *DeleteRuleGroupResponse {
	s.Body = v
	return s
}

type DeregisterTemplatesFromRuleRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SentenceParam *string `json:"SentenceParam,omitempty" xml:"SentenceParam,omitempty"`
}

func (s DeregisterTemplatesFromRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterTemplatesFromRuleRequest) GoString() string {
	return s.String()
}

func (s *DeregisterTemplatesFromRuleRequest) SetRegionId(v string) *DeregisterTemplatesFromRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeregisterTemplatesFromRuleRequest) SetInstanceId(v string) *DeregisterTemplatesFromRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeregisterTemplatesFromRuleRequest) SetSentenceParam(v string) *DeregisterTemplatesFromRuleRequest {
	s.SentenceParam = &v
	return s
}

type DeregisterTemplatesFromRuleResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterTemplatesFromRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeregisterTemplatesFromRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterTemplatesFromRuleResponseBody) SetServerData(v string) *DeregisterTemplatesFromRuleResponseBody {
	s.ServerData = &v
	return s
}

func (s *DeregisterTemplatesFromRuleResponseBody) SetRequestId(v string) *DeregisterTemplatesFromRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeregisterTemplatesFromRuleResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeregisterTemplatesFromRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeregisterTemplatesFromRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterTemplatesFromRuleResponse) GoString() string {
	return s.String()
}

func (s *DeregisterTemplatesFromRuleResponse) SetHeaders(v map[string]*string) *DeregisterTemplatesFromRuleResponse {
	s.Headers = v
	return s
}

func (s *DeregisterTemplatesFromRuleResponse) SetBody(v *DeregisterTemplatesFromRuleResponseBody) *DeregisterTemplatesFromRuleResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	VpcId               *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	InternetIp          *string   `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	NetworkType         *string   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ImageVersionName    *string   `json:"ImageVersionName,omitempty" xml:"ImageVersionName,omitempty"`
	SeriesCode          *string   `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description         *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	AccessType          *int32    `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	EcsStatus           *string   `json:"EcsStatus,omitempty" xml:"EcsStatus,omitempty"`
	Operatable          *bool     `json:"Operatable,omitempty" xml:"Operatable,omitempty"`
	PlanUpgradeStatus   *int32    `json:"PlanUpgradeStatus,omitempty" xml:"PlanUpgradeStatus,omitempty"`
	ExpireTime          *int64    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Upgradeable         *bool     `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	InstanceId          *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetEndpoint    *string   `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetIp          *string   `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Renewable           *bool     `json:"Renewable,omitempty" xml:"Renewable,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint    *string   `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	StartTime           *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UpgradeStatus       *int32    `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	PlanUpgradeable     *bool     `json:"PlanUpgradeable,omitempty" xml:"PlanUpgradeable,omitempty"`
	InstanceStatus      *string   `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode         *string   `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PublicAccessControl *int32    `json:"PublicAccessControl,omitempty" xml:"PublicAccessControl,omitempty"`
	PublicWhiteList     []*string `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	SecurityGroupIds    []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	PrivateWhiteList    []*string `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetIp(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetIp = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetNetworkType(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetImageVersionName(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ImageVersionName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSeriesCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAccessType(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AccessType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetEcsStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.EcsStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetOperatable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Operatable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPlanUpgradeStatus(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PlanUpgradeStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetUpgradeable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Upgradeable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetIp(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRenewable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Renewable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetUpgradeStatus(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPlanUpgradeable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PlanUpgradeable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicAccessControl(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicAccessControl = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSecurityGroupIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPrivateWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PrivateWhiteList = v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	InstanceStatus  *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNo          *int32                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	CurrentPage     *int32                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId      []*string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNo(v int32) *DescribeInstancesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId         *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	InternetIp        *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ImageVersionName  *string `json:"ImageVersionName,omitempty" xml:"ImageVersionName,omitempty"`
	SeriesCode        *string `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsStatus         *string `json:"EcsStatus,omitempty" xml:"EcsStatus,omitempty"`
	Operatable        *bool   `json:"Operatable,omitempty" xml:"Operatable,omitempty"`
	PlanUpgradeStatus *int32  `json:"PlanUpgradeStatus,omitempty" xml:"PlanUpgradeStatus,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Upgradeable       *bool   `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	Legacy            *bool   `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetEndpoint  *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetIp        *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Renewable         *bool   `json:"Renewable,omitempty" xml:"Renewable,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint  *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	StartTime         *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UpgradeStatus     *int32  `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	PlanUpgradeable   *bool   `json:"PlanUpgradeable,omitempty" xml:"PlanUpgradeable,omitempty"`
	InstanceStatus    *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode       *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetIp(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetNetworkType(v string) *DescribeInstancesResponseBodyInstances {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersionName(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersionName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeriesCode(v string) *DescribeInstancesResponseBodyInstances {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEcsStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.EcsStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetOperatable(v bool) *DescribeInstancesResponseBodyInstances {
	s.Operatable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPlanUpgradeStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.PlanUpgradeStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUpgradeable(v bool) *DescribeInstancesResponseBodyInstances {
	s.Upgradeable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLegacy(v bool) *DescribeInstancesResponseBodyInstances {
	s.Legacy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetIp(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRenewable(v bool) *DescribeInstancesResponseBodyInstances {
	s.Renewable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUpgradeStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPlanUpgradeable(v bool) *DescribeInstancesResponseBodyInstances {
	s.PlanUpgradeable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLicenseCode(v string) *DescribeInstancesResponseBodyInstances {
	s.LicenseCode = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeLoginTicketRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLoginTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoginTicketRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketRequest) SetInstanceId(v string) *DescribeLoginTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLoginTicketRequest) SetRegionId(v string) *DescribeLoginTicketRequest {
	s.RegionId = &v
	return s
}

type DescribeLoginTicketResponseBody struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LoginTicket *DescribeLoginTicketResponseBodyLoginTicket `json:"LoginTicket,omitempty" xml:"LoginTicket,omitempty" type:"Struct"`
}

func (s DescribeLoginTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoginTicketResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponseBody) SetRequestId(v string) *DescribeLoginTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoginTicketResponseBody) SetLoginTicket(v *DescribeLoginTicketResponseBodyLoginTicket) *DescribeLoginTicketResponseBody {
	s.LoginTicket = v
	return s
}

type DescribeLoginTicketResponseBodyLoginTicket struct {
	Ticket      *string                                            `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	Certificate *string                                            `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Zones       []*DescribeLoginTicketResponseBodyLoginTicketZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeLoginTicketResponseBodyLoginTicket) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoginTicketResponseBodyLoginTicket) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) SetTicket(v string) *DescribeLoginTicketResponseBodyLoginTicket {
	s.Ticket = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) SetCertificate(v string) *DescribeLoginTicketResponseBodyLoginTicket {
	s.Certificate = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) SetZones(v []*DescribeLoginTicketResponseBodyLoginTicketZones) *DescribeLoginTicketResponseBodyLoginTicket {
	s.Zones = v
	return s
}

type DescribeLoginTicketResponseBodyLoginTicketZones struct {
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
}

func (s DescribeLoginTicketResponseBodyLoginTicketZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoginTicketResponseBodyLoginTicketZones) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) SetZoneId(v string) *DescribeLoginTicketResponseBodyLoginTicketZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) SetLocalName(v string) *DescribeLoginTicketResponseBodyLoginTicketZones {
	s.LocalName = &v
	return s
}

type DescribeLoginTicketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLoginTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoginTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoginTicketResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponse) SetHeaders(v map[string]*string) *DescribeLoginTicketResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoginTicketResponse) SetBody(v *DescribeLoginTicketResponseBody) *DescribeLoginTicketResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
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

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
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

type DescribeSyncInfoRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSyncInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoRequest) SetRegionId(v string) *DescribeSyncInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncInfoRequest) SetInstanceId(v string) *DescribeSyncInfoRequest {
	s.InstanceId = &v
	return s
}

type DescribeSyncInfoResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceInfo *DescribeSyncInfoResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
}

func (s DescribeSyncInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoResponseBody) SetRequestId(v string) *DescribeSyncInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncInfoResponseBody) SetInstanceInfo(v *DescribeSyncInfoResponseBodyInstanceInfo) *DescribeSyncInfoResponseBody {
	s.InstanceInfo = v
	return s
}

type DescribeSyncInfoResponseBodyInstanceInfo struct {
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	VswitchId           *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	RegionNo            *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	ImageVersionName    *string `json:"ImageVersionName,omitempty" xml:"ImageVersionName,omitempty"`
	PlanCode            *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	EcsUuid             *string `json:"EcsUuid,omitempty" xml:"EcsUuid,omitempty"`
	AccessType          *int32  `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	EcsStatus           *string `json:"EcsStatus,omitempty" xml:"EcsStatus,omitempty"`
	PlanUpgradeStatus   *int32  `json:"PlanUpgradeStatus,omitempty" xml:"PlanUpgradeStatus,omitempty"`
	ZoneNo              *string `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
	Aliuid              *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	ProductName         *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	EcsEip              *string `json:"EcsEip,omitempty" xml:"EcsEip,omitempty"`
	ExpireTime          *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	EcsInternetIp       *string `json:"EcsInternetIp,omitempty" xml:"EcsInternetIp,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Renewable           *bool   `json:"Renewable,omitempty" xml:"Renewable,omitempty"`
	EcsIntranetIp       *string `json:"EcsIntranetIp,omitempty" xml:"EcsIntranetIp,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RegionName          *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	UpgradeStatus       *int32  `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	PlanUpgradeable     *string `json:"PlanUpgradeable,omitempty" xml:"PlanUpgradeable,omitempty"`
	CustomName          *string `json:"CustomName,omitempty" xml:"CustomName,omitempty"`
	EcsNetworkType      *string `json:"EcsNetworkType,omitempty" xml:"EcsNetworkType,omitempty"`
	PublicAccessControl *int32  `json:"PublicAccessControl,omitempty" xml:"PublicAccessControl,omitempty"`
	VendorCode          *string `json:"VendorCode,omitempty" xml:"VendorCode,omitempty"`
	PlanName            *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeSyncInfoResponseBodyInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncInfoResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetStatus(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetVswitchId(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetRegionNo(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.RegionNo = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsInstanceId(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetImageVersionName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ImageVersionName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanCode(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanCode = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsUuid(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsUuid = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetAccessType(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.AccessType = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsStatus(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsStatus = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanUpgradeStatus(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanUpgradeStatus = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetZoneNo(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ZoneNo = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetAliuid(v int64) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.Aliuid = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetProductName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ProductName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsEip(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsEip = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetExpireTime(v int64) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsInternetIp(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsInternetIp = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetInstanceId(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetRenewable(v bool) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.Renewable = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsIntranetIp(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsIntranetIp = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetStartTime(v int64) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetRegionName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.RegionName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetUpgradeStatus(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanUpgradeable(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanUpgradeable = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetCustomName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.CustomName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsNetworkType(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsNetworkType = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPublicAccessControl(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PublicAccessControl = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetVendorCode(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.VendorCode = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetProductCode(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ProductCode = &v
	return s
}

type DescribeSyncInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSyncInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSyncInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoResponse) SetHeaders(v map[string]*string) *DescribeSyncInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncInfoResponse) SetBody(v *DescribeSyncInfoResponseBody) *DescribeSyncInfoResponse {
	s.Body = v
	return s
}

type DisableLogMaskConfigsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskIdList []*string `json:"MaskIdList,omitempty" xml:"MaskIdList,omitempty" type:"Repeated"`
}

func (s DisableLogMaskConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLogMaskConfigsRequest) GoString() string {
	return s.String()
}

func (s *DisableLogMaskConfigsRequest) SetRegionId(v string) *DisableLogMaskConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DisableLogMaskConfigsRequest) SetInstanceId(v string) *DisableLogMaskConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableLogMaskConfigsRequest) SetMaskIdList(v []*string) *DisableLogMaskConfigsRequest {
	s.MaskIdList = v
	return s
}

type DisableLogMaskConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLogMaskConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLogMaskConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLogMaskConfigsResponseBody) SetRequestId(v string) *DisableLogMaskConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DisableLogMaskConfigsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableLogMaskConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLogMaskConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLogMaskConfigsResponse) GoString() string {
	return s.String()
}

func (s *DisableLogMaskConfigsResponse) SetHeaders(v map[string]*string) *DisableLogMaskConfigsResponse {
	s.Headers = v
	return s
}

func (s *DisableLogMaskConfigsResponse) SetBody(v *DisableLogMaskConfigsResponseBody) *DisableLogMaskConfigsResponse {
	s.Body = v
	return s
}

type DissociateRulesFromDbRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JsonParam  *string `json:"JsonParam,omitempty" xml:"JsonParam,omitempty"`
}

func (s DissociateRulesFromDbRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateRulesFromDbRequest) GoString() string {
	return s.String()
}

func (s *DissociateRulesFromDbRequest) SetRegionId(v string) *DissociateRulesFromDbRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateRulesFromDbRequest) SetInstanceId(v string) *DissociateRulesFromDbRequest {
	s.InstanceId = &v
	return s
}

func (s *DissociateRulesFromDbRequest) SetJsonParam(v string) *DissociateRulesFromDbRequest {
	s.JsonParam = &v
	return s
}

type DissociateRulesFromDbResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *DissociateRulesFromDbResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s DissociateRulesFromDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateRulesFromDbResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateRulesFromDbResponseBody) SetRequestId(v string) *DissociateRulesFromDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateRulesFromDbResponseBody) SetServerData(v *DissociateRulesFromDbResponseBodyServerData) *DissociateRulesFromDbResponseBody {
	s.ServerData = v
	return s
}

type DissociateRulesFromDbResponseBodyServerData struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
}

func (s DissociateRulesFromDbResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s DissociateRulesFromDbResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *DissociateRulesFromDbResponseBodyServerData) SetJsonResult(v string) *DissociateRulesFromDbResponseBodyServerData {
	s.JsonResult = &v
	return s
}

type DissociateRulesFromDbResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateRulesFromDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateRulesFromDbResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateRulesFromDbResponse) GoString() string {
	return s.String()
}

func (s *DissociateRulesFromDbResponse) SetHeaders(v map[string]*string) *DissociateRulesFromDbResponse {
	s.Headers = v
	return s
}

func (s *DissociateRulesFromDbResponse) SetBody(v *DissociateRulesFromDbResponseBody) *DissociateRulesFromDbResponse {
	s.Body = v
	return s
}

type DissociateTemplatesFromRuleRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JsonParam  *string `json:"JsonParam,omitempty" xml:"JsonParam,omitempty"`
}

func (s DissociateTemplatesFromRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateTemplatesFromRuleRequest) GoString() string {
	return s.String()
}

func (s *DissociateTemplatesFromRuleRequest) SetRegionId(v string) *DissociateTemplatesFromRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateTemplatesFromRuleRequest) SetInstanceId(v string) *DissociateTemplatesFromRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DissociateTemplatesFromRuleRequest) SetJsonParam(v string) *DissociateTemplatesFromRuleRequest {
	s.JsonParam = &v
	return s
}

type DissociateTemplatesFromRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateTemplatesFromRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateTemplatesFromRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateTemplatesFromRuleResponseBody) SetRequestId(v string) *DissociateTemplatesFromRuleResponseBody {
	s.RequestId = &v
	return s
}

type DissociateTemplatesFromRuleResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateTemplatesFromRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateTemplatesFromRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateTemplatesFromRuleResponse) GoString() string {
	return s.String()
}

func (s *DissociateTemplatesFromRuleResponse) SetHeaders(v map[string]*string) *DissociateTemplatesFromRuleResponse {
	s.Headers = v
	return s
}

func (s *DissociateTemplatesFromRuleResponse) SetBody(v *DissociateTemplatesFromRuleResponseBody) *DissociateTemplatesFromRuleResponse {
	s.Body = v
	return s
}

type EditLogMaskConfigRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskId          *int32  `json:"MaskId,omitempty" xml:"MaskId,omitempty"`
	MaskName        *string `json:"MaskName,omitempty" xml:"MaskName,omitempty"`
	MaskRegex       *string `json:"MaskRegex,omitempty" xml:"MaskRegex,omitempty"`
	MaskTxt         *string `json:"MaskTxt,omitempty" xml:"MaskTxt,omitempty"`
	MaskDescription *string `json:"MaskDescription,omitempty" xml:"MaskDescription,omitempty"`
}

func (s EditLogMaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s EditLogMaskConfigRequest) GoString() string {
	return s.String()
}

func (s *EditLogMaskConfigRequest) SetRegionId(v string) *EditLogMaskConfigRequest {
	s.RegionId = &v
	return s
}

func (s *EditLogMaskConfigRequest) SetInstanceId(v string) *EditLogMaskConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *EditLogMaskConfigRequest) SetMaskId(v int32) *EditLogMaskConfigRequest {
	s.MaskId = &v
	return s
}

func (s *EditLogMaskConfigRequest) SetMaskName(v string) *EditLogMaskConfigRequest {
	s.MaskName = &v
	return s
}

func (s *EditLogMaskConfigRequest) SetMaskRegex(v string) *EditLogMaskConfigRequest {
	s.MaskRegex = &v
	return s
}

func (s *EditLogMaskConfigRequest) SetMaskTxt(v string) *EditLogMaskConfigRequest {
	s.MaskTxt = &v
	return s
}

func (s *EditLogMaskConfigRequest) SetMaskDescription(v string) *EditLogMaskConfigRequest {
	s.MaskDescription = &v
	return s
}

type EditLogMaskConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditLogMaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditLogMaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *EditLogMaskConfigResponseBody) SetRequestId(v string) *EditLogMaskConfigResponseBody {
	s.RequestId = &v
	return s
}

type EditLogMaskConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditLogMaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditLogMaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s EditLogMaskConfigResponse) GoString() string {
	return s.String()
}

func (s *EditLogMaskConfigResponse) SetHeaders(v map[string]*string) *EditLogMaskConfigResponse {
	s.Headers = v
	return s
}

func (s *EditLogMaskConfigResponse) SetBody(v *EditLogMaskConfigResponseBody) *EditLogMaskConfigResponse {
	s.Body = v
	return s
}

type EnableLogMaskConfigsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskIdList []*string `json:"MaskIdList,omitempty" xml:"MaskIdList,omitempty" type:"Repeated"`
}

func (s EnableLogMaskConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLogMaskConfigsRequest) GoString() string {
	return s.String()
}

func (s *EnableLogMaskConfigsRequest) SetRegionId(v string) *EnableLogMaskConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *EnableLogMaskConfigsRequest) SetInstanceId(v string) *EnableLogMaskConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableLogMaskConfigsRequest) SetMaskIdList(v []*string) *EnableLogMaskConfigsRequest {
	s.MaskIdList = v
	return s
}

type EnableLogMaskConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLogMaskConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLogMaskConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLogMaskConfigsResponseBody) SetRequestId(v string) *EnableLogMaskConfigsResponseBody {
	s.RequestId = &v
	return s
}

type EnableLogMaskConfigsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableLogMaskConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLogMaskConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLogMaskConfigsResponse) GoString() string {
	return s.String()
}

func (s *EnableLogMaskConfigsResponse) SetHeaders(v map[string]*string) *EnableLogMaskConfigsResponse {
	s.Headers = v
	return s
}

func (s *EnableLogMaskConfigsResponse) SetBody(v *EnableLogMaskConfigsResponseBody) *EnableLogMaskConfigsResponse {
	s.Body = v
	return s
}

type GetAgentFileUrlRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentFileUrlRequest) SetRegionId(v string) *GetAgentFileUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetAgentFileUrlRequest) SetInstanceId(v string) *GetAgentFileUrlRequest {
	s.InstanceId = &v
	return s
}

type GetAgentFileUrlResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LinuxFileUrl *string `json:"LinuxFileUrl,omitempty" xml:"LinuxFileUrl,omitempty"`
	AccessToken  *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	WinFileUrl   *string `json:"WinFileUrl,omitempty" xml:"WinFileUrl,omitempty"`
}

func (s GetAgentFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentFileUrlResponseBody) SetRequestId(v string) *GetAgentFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentFileUrlResponseBody) SetLinuxFileUrl(v string) *GetAgentFileUrlResponseBody {
	s.LinuxFileUrl = &v
	return s
}

func (s *GetAgentFileUrlResponseBody) SetAccessToken(v string) *GetAgentFileUrlResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetAgentFileUrlResponseBody) SetWinFileUrl(v string) *GetAgentFileUrlResponseBody {
	s.WinFileUrl = &v
	return s
}

type GetAgentFileUrlResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentFileUrlResponse) SetHeaders(v map[string]*string) *GetAgentFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentFileUrlResponse) SetBody(v *GetAgentFileUrlResponseBody) *GetAgentFileUrlResponse {
	s.Body = v
	return s
}

type GetAgentListRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AgentIp     *string `json:"AgentIp,omitempty" xml:"AgentIp,omitempty"`
	AgentStatus *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	AgentOs     *string `json:"AgentOs,omitempty" xml:"AgentOs,omitempty"`
}

func (s GetAgentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentListRequest) GoString() string {
	return s.String()
}

func (s *GetAgentListRequest) SetRegionId(v string) *GetAgentListRequest {
	s.RegionId = &v
	return s
}

func (s *GetAgentListRequest) SetInstanceId(v string) *GetAgentListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentListRequest) SetAgentIp(v string) *GetAgentListRequest {
	s.AgentIp = &v
	return s
}

func (s *GetAgentListRequest) SetAgentStatus(v int32) *GetAgentListRequest {
	s.AgentStatus = &v
	return s
}

func (s *GetAgentListRequest) SetAgentOs(v string) *GetAgentListRequest {
	s.AgentOs = &v
	return s
}

type GetAgentListResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AgentList []*GetAgentListResponseBodyAgentList `json:"AgentList,omitempty" xml:"AgentList,omitempty" type:"Repeated"`
}

func (s GetAgentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentListResponseBody) SetRequestId(v string) *GetAgentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentListResponseBody) SetAgentList(v []*GetAgentListResponseBodyAgentList) *GetAgentListResponseBody {
	s.AgentList = v
	return s
}

type GetAgentListResponseBodyAgentList struct {
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	PrivateIp       *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	RmagentMem      *int32  `json:"RmagentMem,omitempty" xml:"RmagentMem,omitempty"`
	AgentId         *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	RmagentCpu      *int32  `json:"RmagentCpu,omitempty" xml:"RmagentCpu,omitempty"`
	FirstLoginTime  *string `json:"FirstLoginTime,omitempty" xml:"FirstLoginTime,omitempty"`
	AgentOs         *string `json:"AgentOs,omitempty" xml:"AgentOs,omitempty"`
	AgentStatus     *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	AgentPort       *string `json:"AgentPort,omitempty" xml:"AgentPort,omitempty"`
	AgentVersion    *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	SendPackets     *int64  `json:"SendPackets,omitempty" xml:"SendPackets,omitempty"`
	SendBytes       *int64  `json:"SendBytes,omitempty" xml:"SendBytes,omitempty"`
	LastActiveTime  *string `json:"LastActiveTime,omitempty" xml:"LastActiveTime,omitempty"`
	SendPacketCount *int64  `json:"SendPacketCount,omitempty" xml:"SendPacketCount,omitempty"`
	EcsId           *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	PublicIp        *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	SendByteCount   *int64  `json:"SendByteCount,omitempty" xml:"SendByteCount,omitempty"`
}

func (s GetAgentListResponseBodyAgentList) String() string {
	return tea.Prettify(s)
}

func (s GetAgentListResponseBodyAgentList) GoString() string {
	return s.String()
}

func (s *GetAgentListResponseBodyAgentList) SetVpcId(v string) *GetAgentListResponseBodyAgentList {
	s.VpcId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetPrivateIp(v string) *GetAgentListResponseBodyAgentList {
	s.PrivateIp = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetRmagentMem(v int32) *GetAgentListResponseBodyAgentList {
	s.RmagentMem = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentId(v string) *GetAgentListResponseBodyAgentList {
	s.AgentId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetRmagentCpu(v int32) *GetAgentListResponseBodyAgentList {
	s.RmagentCpu = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetFirstLoginTime(v string) *GetAgentListResponseBodyAgentList {
	s.FirstLoginTime = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentOs(v string) *GetAgentListResponseBodyAgentList {
	s.AgentOs = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentStatus(v int32) *GetAgentListResponseBodyAgentList {
	s.AgentStatus = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentPort(v string) *GetAgentListResponseBodyAgentList {
	s.AgentPort = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentVersion(v string) *GetAgentListResponseBodyAgentList {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendPackets(v int64) *GetAgentListResponseBodyAgentList {
	s.SendPackets = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendBytes(v int64) *GetAgentListResponseBodyAgentList {
	s.SendBytes = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetLastActiveTime(v string) *GetAgentListResponseBodyAgentList {
	s.LastActiveTime = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendPacketCount(v int64) *GetAgentListResponseBodyAgentList {
	s.SendPacketCount = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetEcsId(v string) *GetAgentListResponseBodyAgentList {
	s.EcsId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetPublicIp(v string) *GetAgentListResponseBodyAgentList {
	s.PublicIp = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendByteCount(v int64) *GetAgentListResponseBodyAgentList {
	s.SendByteCount = &v
	return s
}

type GetAgentListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentListResponse) GoString() string {
	return s.String()
}

func (s *GetAgentListResponse) SetHeaders(v map[string]*string) *GetAgentListResponse {
	s.Headers = v
	return s
}

func (s *GetAgentListResponse) SetBody(v *GetAgentListResponseBody) *GetAgentListResponse {
	s.Body = v
	return s
}

type GetAppointOperationRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAppointOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppointOperationRequest) GoString() string {
	return s.String()
}

func (s *GetAppointOperationRequest) SetRegionId(v string) *GetAppointOperationRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppointOperationRequest) SetInstanceId(v string) *GetAppointOperationRequest {
	s.InstanceId = &v
	return s
}

type GetAppointOperationResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppointOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppointOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppointOperationResponseBody) SetServerData(v string) *GetAppointOperationResponseBody {
	s.ServerData = &v
	return s
}

func (s *GetAppointOperationResponseBody) SetRequestId(v string) *GetAppointOperationResponseBody {
	s.RequestId = &v
	return s
}

type GetAppointOperationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppointOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppointOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppointOperationResponse) GoString() string {
	return s.String()
}

func (s *GetAppointOperationResponse) SetHeaders(v map[string]*string) *GetAppointOperationResponse {
	s.Headers = v
	return s
}

func (s *GetAppointOperationResponse) SetBody(v *GetAppointOperationResponseBody) *GetAppointOperationResponse {
	s.Body = v
	return s
}

type GetAuditCountRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetAuditCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountRequest) GoString() string {
	return s.String()
}

func (s *GetAuditCountRequest) SetRegionId(v string) *GetAuditCountRequest {
	s.RegionId = &v
	return s
}

func (s *GetAuditCountRequest) SetInstanceId(v string) *GetAuditCountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuditCountRequest) SetDbId(v int32) *GetAuditCountRequest {
	s.DbId = &v
	return s
}

func (s *GetAuditCountRequest) SetBeginDate(v string) *GetAuditCountRequest {
	s.BeginDate = &v
	return s
}

func (s *GetAuditCountRequest) SetEndDate(v string) *GetAuditCountRequest {
	s.EndDate = &v
	return s
}

type GetAuditCountResponseBody struct {
	SessionCount *int64  `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	SqlCount     *int64  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	RiskCount    *int64  `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuditCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditCountResponseBody) SetSessionCount(v int64) *GetAuditCountResponseBody {
	s.SessionCount = &v
	return s
}

func (s *GetAuditCountResponseBody) SetSqlCount(v int64) *GetAuditCountResponseBody {
	s.SqlCount = &v
	return s
}

func (s *GetAuditCountResponseBody) SetRiskCount(v int64) *GetAuditCountResponseBody {
	s.RiskCount = &v
	return s
}

func (s *GetAuditCountResponseBody) SetRequestId(v string) *GetAuditCountResponseBody {
	s.RequestId = &v
	return s
}

type GetAuditCountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuditCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuditCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountResponse) GoString() string {
	return s.String()
}

func (s *GetAuditCountResponse) SetHeaders(v map[string]*string) *GetAuditCountResponse {
	s.Headers = v
	return s
}

func (s *GetAuditCountResponse) SetBody(v *GetAuditCountResponseBody) *GetAuditCountResponse {
	s.Body = v
	return s
}

type GetAuditCountDistributionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetAuditCountDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionRequest) SetRegionId(v string) *GetAuditCountDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetInstanceId(v string) *GetAuditCountDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetDbId(v int32) *GetAuditCountDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetBeginDate(v string) *GetAuditCountDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetEndDate(v string) *GetAuditCountDistributionRequest {
	s.EndDate = &v
	return s
}

type GetAuditCountDistributionResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetAuditCountDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetAuditCountDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionResponseBody) SetRequestId(v string) *GetAuditCountDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditCountDistributionResponseBody) SetTimeList(v []*GetAuditCountDistributionResponseBodyTimeList) *GetAuditCountDistributionResponseBody {
	s.TimeList = v
	return s
}

type GetAuditCountDistributionResponseBodyTimeList struct {
	SessionCount *int64  `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	BeginDate    *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	SqlCount     *int64  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	RiskCount    *int64  `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
}

func (s GetAuditCountDistributionResponseBodyTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetSessionCount(v int64) *GetAuditCountDistributionResponseBodyTimeList {
	s.SessionCount = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetBeginDate(v string) *GetAuditCountDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetSqlCount(v int64) *GetAuditCountDistributionResponseBodyTimeList {
	s.SqlCount = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetEndDate(v string) *GetAuditCountDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetAuditCountDistributionResponseBodyTimeList) SetRiskCount(v int64) *GetAuditCountDistributionResponseBodyTimeList {
	s.RiskCount = &v
	return s
}

type GetAuditCountDistributionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuditCountDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuditCountDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditCountDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionResponse) SetHeaders(v map[string]*string) *GetAuditCountDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetAuditCountDistributionResponse) SetBody(v *GetAuditCountDistributionResponseBody) *GetAuditCountDistributionResponse {
	s.Body = v
	return s
}

type GetBaseTemplateListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetBaseTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBaseTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetBaseTemplateListRequest) SetRegionId(v string) *GetBaseTemplateListRequest {
	s.RegionId = &v
	return s
}

func (s *GetBaseTemplateListRequest) SetInstanceId(v string) *GetBaseTemplateListRequest {
	s.InstanceId = &v
	return s
}

type GetBaseTemplateListResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateList []*GetBaseTemplateListResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
}

func (s GetBaseTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBaseTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaseTemplateListResponseBody) SetRequestId(v string) *GetBaseTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaseTemplateListResponseBody) SetTemplateList(v []*GetBaseTemplateListResponseBodyTemplateList) *GetBaseTemplateListResponseBody {
	s.TemplateList = v
	return s
}

type GetBaseTemplateListResponseBodyTemplateList struct {
	DbTypeName      *string `json:"DbTypeName,omitempty" xml:"DbTypeName,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	SqlTypeName     *string `json:"SqlTypeName,omitempty" xml:"SqlTypeName,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateState   *string `json:"TemplateState,omitempty" xml:"TemplateState,omitempty"`
}

func (s GetBaseTemplateListResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s GetBaseTemplateListResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *GetBaseTemplateListResponseBodyTemplateList) SetDbTypeName(v string) *GetBaseTemplateListResponseBodyTemplateList {
	s.DbTypeName = &v
	return s
}

func (s *GetBaseTemplateListResponseBodyTemplateList) SetTemplateContent(v string) *GetBaseTemplateListResponseBodyTemplateList {
	s.TemplateContent = &v
	return s
}

func (s *GetBaseTemplateListResponseBodyTemplateList) SetSqlTypeName(v string) *GetBaseTemplateListResponseBodyTemplateList {
	s.SqlTypeName = &v
	return s
}

func (s *GetBaseTemplateListResponseBodyTemplateList) SetTemplateId(v string) *GetBaseTemplateListResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *GetBaseTemplateListResponseBodyTemplateList) SetTemplateState(v string) *GetBaseTemplateListResponseBodyTemplateList {
	s.TemplateState = &v
	return s
}

type GetBaseTemplateListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBaseTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBaseTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBaseTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetBaseTemplateListResponse) SetHeaders(v map[string]*string) *GetBaseTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetBaseTemplateListResponse) SetBody(v *GetBaseTemplateListResponseBody) *GetBaseTemplateListResponse {
	s.Body = v
	return s
}

type GetDasUsageRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDasUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageRequest) GoString() string {
	return s.String()
}

func (s *GetDasUsageRequest) SetRegionId(v string) *GetDasUsageRequest {
	s.RegionId = &v
	return s
}

func (s *GetDasUsageRequest) SetInstanceId(v string) *GetDasUsageRequest {
	s.InstanceId = &v
	return s
}

type GetDasUsageResponseBody struct {
	Over1sSqlCount    *int64                                `json:"Over1sSqlCount,omitempty" xml:"Over1sSqlCount,omitempty"`
	RequestId         *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionTotalCount *int64                                `json:"SessionTotalCount,omitempty" xml:"SessionTotalCount,omitempty"`
	CreateTime        *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SqlTotalCount     *int64                                `json:"SqlTotalCount,omitempty" xml:"SqlTotalCount,omitempty"`
	AuditAsset        *GetDasUsageResponseBodyAuditAsset    `json:"AuditAsset,omitempty" xml:"AuditAsset,omitempty" type:"Struct"`
	ConsoleAccess     *GetDasUsageResponseBodyConsoleAccess `json:"ConsoleAccess,omitempty" xml:"ConsoleAccess,omitempty" type:"Struct"`
	Agent             *GetDasUsageResponseBodyAgent         `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	NoticeState       *GetDasUsageResponseBodyNoticeState   `json:"NoticeState,omitempty" xml:"NoticeState,omitempty" type:"Struct"`
	RiskAsset         *GetDasUsageResponseBodyRiskAsset     `json:"RiskAsset,omitempty" xml:"RiskAsset,omitempty" type:"Struct"`
	RiskRule          *GetDasUsageResponseBodyRiskRule      `json:"RiskRule,omitempty" xml:"RiskRule,omitempty" type:"Struct"`
	WpAsset           *GetDasUsageResponseBodyWpAsset       `json:"WpAsset,omitempty" xml:"WpAsset,omitempty" type:"Struct"`
}

func (s GetDasUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBody) SetOver1sSqlCount(v int64) *GetDasUsageResponseBody {
	s.Over1sSqlCount = &v
	return s
}

func (s *GetDasUsageResponseBody) SetRequestId(v string) *GetDasUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDasUsageResponseBody) SetSessionTotalCount(v int64) *GetDasUsageResponseBody {
	s.SessionTotalCount = &v
	return s
}

func (s *GetDasUsageResponseBody) SetCreateTime(v string) *GetDasUsageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDasUsageResponseBody) SetSqlTotalCount(v int64) *GetDasUsageResponseBody {
	s.SqlTotalCount = &v
	return s
}

func (s *GetDasUsageResponseBody) SetAuditAsset(v *GetDasUsageResponseBodyAuditAsset) *GetDasUsageResponseBody {
	s.AuditAsset = v
	return s
}

func (s *GetDasUsageResponseBody) SetConsoleAccess(v *GetDasUsageResponseBodyConsoleAccess) *GetDasUsageResponseBody {
	s.ConsoleAccess = v
	return s
}

func (s *GetDasUsageResponseBody) SetAgent(v *GetDasUsageResponseBodyAgent) *GetDasUsageResponseBody {
	s.Agent = v
	return s
}

func (s *GetDasUsageResponseBody) SetNoticeState(v *GetDasUsageResponseBodyNoticeState) *GetDasUsageResponseBody {
	s.NoticeState = v
	return s
}

func (s *GetDasUsageResponseBody) SetRiskAsset(v *GetDasUsageResponseBodyRiskAsset) *GetDasUsageResponseBody {
	s.RiskAsset = v
	return s
}

func (s *GetDasUsageResponseBody) SetRiskRule(v *GetDasUsageResponseBodyRiskRule) *GetDasUsageResponseBody {
	s.RiskRule = v
	return s
}

func (s *GetDasUsageResponseBody) SetWpAsset(v *GetDasUsageResponseBodyWpAsset) *GetDasUsageResponseBody {
	s.WpAsset = v
	return s
}

type GetDasUsageResponseBodyAuditAsset struct {
	DbCount *int32                                      `json:"DbCount,omitempty" xml:"DbCount,omitempty"`
	DbTypes []*GetDasUsageResponseBodyAuditAssetDbTypes `json:"DbTypes,omitempty" xml:"DbTypes,omitempty" type:"Repeated"`
}

func (s GetDasUsageResponseBodyAuditAsset) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyAuditAsset) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyAuditAsset) SetDbCount(v int32) *GetDasUsageResponseBodyAuditAsset {
	s.DbCount = &v
	return s
}

func (s *GetDasUsageResponseBodyAuditAsset) SetDbTypes(v []*GetDasUsageResponseBodyAuditAssetDbTypes) *GetDasUsageResponseBodyAuditAsset {
	s.DbTypes = v
	return s
}

type GetDasUsageResponseBodyAuditAssetDbTypes struct {
	DbType  *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbCount *int32  `json:"DbCount,omitempty" xml:"DbCount,omitempty"`
}

func (s GetDasUsageResponseBodyAuditAssetDbTypes) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyAuditAssetDbTypes) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyAuditAssetDbTypes) SetDbType(v string) *GetDasUsageResponseBodyAuditAssetDbTypes {
	s.DbType = &v
	return s
}

func (s *GetDasUsageResponseBodyAuditAssetDbTypes) SetDbCount(v int32) *GetDasUsageResponseBodyAuditAssetDbTypes {
	s.DbCount = &v
	return s
}

type GetDasUsageResponseBodyConsoleAccess struct {
	LastAccessTime    *string `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	Day90AccessCount  *int32  `json:"Day90AccessCount,omitempty" xml:"Day90AccessCount,omitempty"`
	Day15AccessCount  *int32  `json:"Day15AccessCount,omitempty" xml:"Day15AccessCount,omitempty"`
	Day30AccessCount  *int32  `json:"Day30AccessCount,omitempty" xml:"Day30AccessCount,omitempty"`
	Day180AccessCount *int32  `json:"Day180AccessCount,omitempty" xml:"Day180AccessCount,omitempty"`
}

func (s GetDasUsageResponseBodyConsoleAccess) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyConsoleAccess) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyConsoleAccess) SetLastAccessTime(v string) *GetDasUsageResponseBodyConsoleAccess {
	s.LastAccessTime = &v
	return s
}

func (s *GetDasUsageResponseBodyConsoleAccess) SetDay90AccessCount(v int32) *GetDasUsageResponseBodyConsoleAccess {
	s.Day90AccessCount = &v
	return s
}

func (s *GetDasUsageResponseBodyConsoleAccess) SetDay15AccessCount(v int32) *GetDasUsageResponseBodyConsoleAccess {
	s.Day15AccessCount = &v
	return s
}

func (s *GetDasUsageResponseBodyConsoleAccess) SetDay30AccessCount(v int32) *GetDasUsageResponseBodyConsoleAccess {
	s.Day30AccessCount = &v
	return s
}

func (s *GetDasUsageResponseBodyConsoleAccess) SetDay180AccessCount(v int32) *GetDasUsageResponseBodyConsoleAccess {
	s.Day180AccessCount = &v
	return s
}

type GetDasUsageResponseBodyAgent struct {
	HasFlow     *bool   `json:"HasFlow,omitempty" xml:"HasFlow,omitempty"`
	InstState   *string `json:"InstState,omitempty" xml:"InstState,omitempty"`
	ActiveCount *int32  `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
}

func (s GetDasUsageResponseBodyAgent) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyAgent) SetHasFlow(v bool) *GetDasUsageResponseBodyAgent {
	s.HasFlow = &v
	return s
}

func (s *GetDasUsageResponseBodyAgent) SetInstState(v string) *GetDasUsageResponseBodyAgent {
	s.InstState = &v
	return s
}

func (s *GetDasUsageResponseBodyAgent) SetActiveCount(v int32) *GetDasUsageResponseBodyAgent {
	s.ActiveCount = &v
	return s
}

type GetDasUsageResponseBodyNoticeState struct {
	RiskNotice *bool `json:"RiskNotice,omitempty" xml:"RiskNotice,omitempty"`
	SysNotice  *bool `json:"SysNotice,omitempty" xml:"SysNotice,omitempty"`
}

func (s GetDasUsageResponseBodyNoticeState) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyNoticeState) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyNoticeState) SetRiskNotice(v bool) *GetDasUsageResponseBodyNoticeState {
	s.RiskNotice = &v
	return s
}

func (s *GetDasUsageResponseBodyNoticeState) SetSysNotice(v bool) *GetDasUsageResponseBodyNoticeState {
	s.SysNotice = &v
	return s
}

type GetDasUsageResponseBodyRiskAsset struct {
	RiskDBCount      *int32 `json:"RiskDBCount,omitempty" xml:"RiskDBCount,omitempty"`
	Day30RiskDBCount *int32 `json:"Day30RiskDBCount,omitempty" xml:"Day30RiskDBCount,omitempty"`
}

func (s GetDasUsageResponseBodyRiskAsset) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyRiskAsset) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyRiskAsset) SetRiskDBCount(v int32) *GetDasUsageResponseBodyRiskAsset {
	s.RiskDBCount = &v
	return s
}

func (s *GetDasUsageResponseBodyRiskAsset) SetDay30RiskDBCount(v int32) *GetDasUsageResponseBodyRiskAsset {
	s.Day30RiskDBCount = &v
	return s
}

type GetDasUsageResponseBodyRiskRule struct {
	Day30RiskCount  *int64                                          `json:"Day30RiskCount,omitempty" xml:"Day30RiskCount,omitempty"`
	RiskCount       *int64                                          `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	MaxHitRule      *GetDasUsageResponseBodyRiskRuleMaxHitRule      `json:"MaxHitRule,omitempty" xml:"MaxHitRule,omitempty" type:"Struct"`
	Day30MaxHitRule *GetDasUsageResponseBodyRiskRuleDay30MaxHitRule `json:"Day30MaxHitRule,omitempty" xml:"Day30MaxHitRule,omitempty" type:"Struct"`
}

func (s GetDasUsageResponseBodyRiskRule) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyRiskRule) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyRiskRule) SetDay30RiskCount(v int64) *GetDasUsageResponseBodyRiskRule {
	s.Day30RiskCount = &v
	return s
}

func (s *GetDasUsageResponseBodyRiskRule) SetRiskCount(v int64) *GetDasUsageResponseBodyRiskRule {
	s.RiskCount = &v
	return s
}

func (s *GetDasUsageResponseBodyRiskRule) SetMaxHitRule(v *GetDasUsageResponseBodyRiskRuleMaxHitRule) *GetDasUsageResponseBodyRiskRule {
	s.MaxHitRule = v
	return s
}

func (s *GetDasUsageResponseBodyRiskRule) SetDay30MaxHitRule(v *GetDasUsageResponseBodyRiskRuleDay30MaxHitRule) *GetDasUsageResponseBodyRiskRule {
	s.Day30MaxHitRule = v
	return s
}

type GetDasUsageResponseBodyRiskRuleMaxHitRule struct {
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetDasUsageResponseBodyRiskRuleMaxHitRule) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyRiskRuleMaxHitRule) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyRiskRuleMaxHitRule) SetCount(v int64) *GetDasUsageResponseBodyRiskRuleMaxHitRule {
	s.Count = &v
	return s
}

func (s *GetDasUsageResponseBodyRiskRuleMaxHitRule) SetRuleName(v string) *GetDasUsageResponseBodyRiskRuleMaxHitRule {
	s.RuleName = &v
	return s
}

type GetDasUsageResponseBodyRiskRuleDay30MaxHitRule struct {
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetDasUsageResponseBodyRiskRuleDay30MaxHitRule) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyRiskRuleDay30MaxHitRule) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyRiskRuleDay30MaxHitRule) SetCount(v int64) *GetDasUsageResponseBodyRiskRuleDay30MaxHitRule {
	s.Count = &v
	return s
}

func (s *GetDasUsageResponseBodyRiskRuleDay30MaxHitRule) SetRuleName(v string) *GetDasUsageResponseBodyRiskRuleDay30MaxHitRule {
	s.RuleName = &v
	return s
}

type GetDasUsageResponseBodyWpAsset struct {
	AvgTime *int64  `json:"AvgTime,omitempty" xml:"AvgTime,omitempty"`
	DbName  *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
}

func (s GetDasUsageResponseBodyWpAsset) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponseBodyWpAsset) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponseBodyWpAsset) SetAvgTime(v int64) *GetDasUsageResponseBodyWpAsset {
	s.AvgTime = &v
	return s
}

func (s *GetDasUsageResponseBodyWpAsset) SetDbName(v string) *GetDasUsageResponseBodyWpAsset {
	s.DbName = &v
	return s
}

type GetDasUsageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDasUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDasUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDasUsageResponse) GoString() string {
	return s.String()
}

func (s *GetDasUsageResponse) SetHeaders(v map[string]*string) *GetDasUsageResponse {
	s.Headers = v
	return s
}

func (s *GetDasUsageResponse) SetBody(v *GetDasUsageResponseBody) *GetDasUsageResponse {
	s.Body = v
	return s
}

type GetDBAuditCountListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetDBAuditCountListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDBAuditCountListRequest) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListRequest) SetRegionId(v string) *GetDBAuditCountListRequest {
	s.RegionId = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetInstanceId(v string) *GetDBAuditCountListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetDbId(v int32) *GetDBAuditCountListRequest {
	s.DbId = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetBeginDate(v string) *GetDBAuditCountListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetEndDate(v string) *GetDBAuditCountListRequest {
	s.EndDate = &v
	return s
}

type GetDBAuditCountListResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbList    []*GetDBAuditCountListResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s GetDBAuditCountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDBAuditCountListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListResponseBody) SetRequestId(v string) *GetDBAuditCountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBAuditCountListResponseBody) SetDbList(v []*GetDBAuditCountListResponseBodyDbList) *GetDBAuditCountListResponseBody {
	s.DbList = v
	return s
}

type GetDBAuditCountListResponseBodyDbList struct {
	SessionCount  *int64    `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	DbId          *int32    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName        *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbType        *int32    `json:"DbType,omitempty" xml:"DbType,omitempty"`
	SqlCount      *int64    `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	DbTypeName    *string   `json:"DbTypeName,omitempty" xml:"DbTypeName,omitempty"`
	RiskCount     *int64    `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	DbVersionName *string   `json:"DbVersionName,omitempty" xml:"DbVersionName,omitempty"`
	AssetType     *int32    `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	DbVersion     *int32    `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	DbAddresses   []*string `json:"DbAddresses,omitempty" xml:"DbAddresses,omitempty" type:"Repeated"`
}

func (s GetDBAuditCountListResponseBodyDbList) String() string {
	return tea.Prettify(s)
}

func (s GetDBAuditCountListResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListResponseBodyDbList) SetSessionCount(v int64) *GetDBAuditCountListResponseBodyDbList {
	s.SessionCount = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbId(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.DbId = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbName(v string) *GetDBAuditCountListResponseBodyDbList {
	s.DbName = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbType(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.DbType = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetSqlCount(v int64) *GetDBAuditCountListResponseBodyDbList {
	s.SqlCount = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbTypeName(v string) *GetDBAuditCountListResponseBodyDbList {
	s.DbTypeName = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetRiskCount(v int64) *GetDBAuditCountListResponseBodyDbList {
	s.RiskCount = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbVersionName(v string) *GetDBAuditCountListResponseBodyDbList {
	s.DbVersionName = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetAssetType(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.AssetType = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbVersion(v int32) *GetDBAuditCountListResponseBodyDbList {
	s.DbVersion = &v
	return s
}

func (s *GetDBAuditCountListResponseBodyDbList) SetDbAddresses(v []*string) *GetDBAuditCountListResponseBodyDbList {
	s.DbAddresses = v
	return s
}

type GetDBAuditCountListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDBAuditCountListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDBAuditCountListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDBAuditCountListResponse) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListResponse) SetHeaders(v map[string]*string) *GetDBAuditCountListResponse {
	s.Headers = v
	return s
}

func (s *GetDBAuditCountListResponse) SetBody(v *GetDBAuditCountListResponseBody) *GetDBAuditCountListResponse {
	s.Body = v
	return s
}

type GetGroupDetailRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupKeyId *string `json:"GroupKeyId,omitempty" xml:"GroupKeyId,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
}

func (s GetGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGroupDetailRequest) SetRegionId(v string) *GetGroupDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetGroupDetailRequest) SetInstanceId(v string) *GetGroupDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetGroupDetailRequest) SetGroupKeyId(v string) *GetGroupDetailRequest {
	s.GroupKeyId = &v
	return s
}

func (s *GetGroupDetailRequest) SetGroupType(v string) *GetGroupDetailRequest {
	s.GroupType = &v
	return s
}

type GetGroupDetailResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBody) SetServerData(v string) *GetGroupDetailResponseBody {
	s.ServerData = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetRequestId(v string) *GetGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupDetailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponse) SetHeaders(v map[string]*string) *GetGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGroupDetailResponse) SetBody(v *GetGroupDetailResponseBody) *GetGroupDetailResponse {
	s.Body = v
	return s
}

type GetLicenseRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetLicenseRequest) SetRegionId(v string) *GetLicenseRequest {
	s.RegionId = &v
	return s
}

func (s *GetLicenseRequest) SetInstanceId(v string) *GetLicenseRequest {
	s.InstanceId = &v
	return s
}

type GetLicenseResponseBody struct {
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AssetLimit     *int32  `json:"AssetLimit,omitempty" xml:"AssetLimit,omitempty"`
	AssetLimitUsed *int32  `json:"AssetLimitUsed,omitempty" xml:"AssetLimitUsed,omitempty"`
}

func (s GetLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBody) SetStartTime(v string) *GetLicenseResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetLicenseResponseBody) SetRequestId(v string) *GetLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicenseResponseBody) SetAssetLimit(v int32) *GetLicenseResponseBody {
	s.AssetLimit = &v
	return s
}

func (s *GetLicenseResponseBody) SetAssetLimitUsed(v int32) *GetLicenseResponseBody {
	s.AssetLimitUsed = &v
	return s
}

type GetLicenseResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseResponse) SetHeaders(v map[string]*string) *GetLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseResponse) SetBody(v *GetLicenseResponseBody) *GetLicenseResponse {
	s.Body = v
	return s
}

type GetLogDetailRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	SqlId      *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s GetLogDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogDetailRequest) GoString() string {
	return s.String()
}

func (s *GetLogDetailRequest) SetRegionId(v string) *GetLogDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogDetailRequest) SetInstanceId(v string) *GetLogDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogDetailRequest) SetBeginDate(v string) *GetLogDetailRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogDetailRequest) SetEndDate(v string) *GetLogDetailRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogDetailRequest) SetSqlId(v string) *GetLogDetailRequest {
	s.SqlId = &v
	return s
}

type GetLogDetailResponseBody struct {
	ClientPort        *int32    `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	AppClientIp       *string   `json:"AppClientIp,omitempty" xml:"AppClientIp,omitempty"`
	ExecCostUS        *int32    `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	SessionLogoutTime *string   `json:"SessionLogoutTime,omitempty" xml:"SessionLogoutTime,omitempty"`
	ClientOsUser      *string   `json:"ClientOsUser,omitempty" xml:"ClientOsUser,omitempty"`
	RuleId            *int32    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SqlId             *string   `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SessionId         *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SqlType           *string   `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	AppUsername       *string   `json:"AppUsername,omitempty" xml:"AppUsername,omitempty"`
	RiskLevel         *int32    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	DbId              *int32    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	RuleType          *int32    `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	RuleKeyId         *int32    `json:"RuleKeyId,omitempty" xml:"RuleKeyId,omitempty"`
	SqlResult         *string   `json:"SqlResult,omitempty" xml:"SqlResult,omitempty"`
	AffectRows        *int32    `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	SqlTypeName       *string   `json:"SqlTypeName,omitempty" xml:"SqlTypeName,omitempty"`
	Schema            *string   `json:"Schema,omitempty" xml:"Schema,omitempty"`
	SessionLoginTime  *string   `json:"SessionLoginTime,omitempty" xml:"SessionLoginTime,omitempty"`
	DbUser            *string   `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	ServerMac         *string   `json:"ServerMac,omitempty" xml:"ServerMac,omitempty"`
	DbServer          *string   `json:"DbServer,omitempty" xml:"DbServer,omitempty"`
	RuleName          *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	ResponseCode      *string   `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	SqlContent        *string   `json:"SqlContent,omitempty" xml:"SqlContent,omitempty"`
	InstName          *string   `json:"InstName,omitempty" xml:"InstName,omitempty"`
	TemplateContent   *string   `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	ClientProgram     *string   `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	CaptureTime       *string   `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	ClientIp          *string   `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientMac         *string   `json:"ClientMac,omitempty" xml:"ClientMac,omitempty"`
	TemplateId        *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ResponseText      *string   `json:"ResponseText,omitempty" xml:"ResponseText,omitempty"`
	FetchCostUS       *int32    `json:"FetchCostUS,omitempty" xml:"FetchCostUS,omitempty"`
	TemplateRules     []*string `json:"TemplateRules,omitempty" xml:"TemplateRules,omitempty" type:"Repeated"`
}

func (s GetLogDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogDetailResponseBody) SetClientPort(v int32) *GetLogDetailResponseBody {
	s.ClientPort = &v
	return s
}

func (s *GetLogDetailResponseBody) SetAppClientIp(v string) *GetLogDetailResponseBody {
	s.AppClientIp = &v
	return s
}

func (s *GetLogDetailResponseBody) SetExecCostUS(v int32) *GetLogDetailResponseBody {
	s.ExecCostUS = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSessionLogoutTime(v string) *GetLogDetailResponseBody {
	s.SessionLogoutTime = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientOsUser(v string) *GetLogDetailResponseBody {
	s.ClientOsUser = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleId(v int32) *GetLogDetailResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRequestId(v string) *GetLogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlId(v string) *GetLogDetailResponseBody {
	s.SqlId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSessionId(v string) *GetLogDetailResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlType(v string) *GetLogDetailResponseBody {
	s.SqlType = &v
	return s
}

func (s *GetLogDetailResponseBody) SetAppUsername(v string) *GetLogDetailResponseBody {
	s.AppUsername = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRiskLevel(v int32) *GetLogDetailResponseBody {
	s.RiskLevel = &v
	return s
}

func (s *GetLogDetailResponseBody) SetDbId(v int32) *GetLogDetailResponseBody {
	s.DbId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleType(v int32) *GetLogDetailResponseBody {
	s.RuleType = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleKeyId(v int32) *GetLogDetailResponseBody {
	s.RuleKeyId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlResult(v string) *GetLogDetailResponseBody {
	s.SqlResult = &v
	return s
}

func (s *GetLogDetailResponseBody) SetAffectRows(v int32) *GetLogDetailResponseBody {
	s.AffectRows = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlTypeName(v string) *GetLogDetailResponseBody {
	s.SqlTypeName = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSchema(v string) *GetLogDetailResponseBody {
	s.Schema = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSessionLoginTime(v string) *GetLogDetailResponseBody {
	s.SessionLoginTime = &v
	return s
}

func (s *GetLogDetailResponseBody) SetDbUser(v string) *GetLogDetailResponseBody {
	s.DbUser = &v
	return s
}

func (s *GetLogDetailResponseBody) SetServerMac(v string) *GetLogDetailResponseBody {
	s.ServerMac = &v
	return s
}

func (s *GetLogDetailResponseBody) SetDbServer(v string) *GetLogDetailResponseBody {
	s.DbServer = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleName(v string) *GetLogDetailResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetLogDetailResponseBody) SetResponseCode(v string) *GetLogDetailResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlContent(v string) *GetLogDetailResponseBody {
	s.SqlContent = &v
	return s
}

func (s *GetLogDetailResponseBody) SetInstName(v string) *GetLogDetailResponseBody {
	s.InstName = &v
	return s
}

func (s *GetLogDetailResponseBody) SetTemplateContent(v string) *GetLogDetailResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientProgram(v string) *GetLogDetailResponseBody {
	s.ClientProgram = &v
	return s
}

func (s *GetLogDetailResponseBody) SetCaptureTime(v string) *GetLogDetailResponseBody {
	s.CaptureTime = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientIp(v string) *GetLogDetailResponseBody {
	s.ClientIp = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientMac(v string) *GetLogDetailResponseBody {
	s.ClientMac = &v
	return s
}

func (s *GetLogDetailResponseBody) SetTemplateId(v string) *GetLogDetailResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetResponseText(v string) *GetLogDetailResponseBody {
	s.ResponseText = &v
	return s
}

func (s *GetLogDetailResponseBody) SetFetchCostUS(v int32) *GetLogDetailResponseBody {
	s.FetchCostUS = &v
	return s
}

func (s *GetLogDetailResponseBody) SetTemplateRules(v []*string) *GetLogDetailResponseBody {
	s.TemplateRules = v
	return s
}

type GetLogDetailResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogDetailResponse) GoString() string {
	return s.String()
}

func (s *GetLogDetailResponse) SetHeaders(v map[string]*string) *GetLogDetailResponse {
	s.Headers = v
	return s
}

func (s *GetLogDetailResponse) SetBody(v *GetLogDetailResponseBody) *GetLogDetailResponse {
	s.Body = v
	return s
}

type GetLogListRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId        *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId              *int32    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate         *string   `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate           *string   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageNumber        *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SqlId             *string   `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlKey            *string   `json:"SqlKey,omitempty" xml:"SqlKey,omitempty"`
	SessionId         *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TemplateId        *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	IsSuccess         *string   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	MinAffectRows     *int32    `json:"MinAffectRows,omitempty" xml:"MinAffectRows,omitempty"`
	MaxAffectRows     *int32    `json:"MaxAffectRows,omitempty" xml:"MaxAffectRows,omitempty"`
	MinExecCostUS     *int32    `json:"MinExecCostUS,omitempty" xml:"MinExecCostUS,omitempty"`
	MaxExecCostUS     *int32    `json:"MaxExecCostUS,omitempty" xml:"MaxExecCostUS,omitempty"`
	RuleName          *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	ClientIpList      []*string `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	DbUserList        []*string `json:"DbUserList,omitempty" xml:"DbUserList,omitempty" type:"Repeated"`
	DbHostList        []*string `json:"DbHostList,omitempty" xml:"DbHostList,omitempty" type:"Repeated"`
	RiskLevelList     []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	RuleTypeList      []*string `json:"RuleTypeList,omitempty" xml:"RuleTypeList,omitempty" type:"Repeated"`
	SqlTypeList       []*string `json:"SqlTypeList,omitempty" xml:"SqlTypeList,omitempty" type:"Repeated"`
	ClientProgramList []*string `json:"ClientProgramList,omitempty" xml:"ClientProgramList,omitempty" type:"Repeated"`
}

func (s GetLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogListRequest) GoString() string {
	return s.String()
}

func (s *GetLogListRequest) SetRegionId(v string) *GetLogListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogListRequest) SetInstanceId(v string) *GetLogListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogListRequest) SetDbId(v int32) *GetLogListRequest {
	s.DbId = &v
	return s
}

func (s *GetLogListRequest) SetBeginDate(v string) *GetLogListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogListRequest) SetEndDate(v string) *GetLogListRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogListRequest) SetPageNumber(v int32) *GetLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLogListRequest) SetPageSize(v int32) *GetLogListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLogListRequest) SetSqlId(v string) *GetLogListRequest {
	s.SqlId = &v
	return s
}

func (s *GetLogListRequest) SetSqlKey(v string) *GetLogListRequest {
	s.SqlKey = &v
	return s
}

func (s *GetLogListRequest) SetSessionId(v string) *GetLogListRequest {
	s.SessionId = &v
	return s
}

func (s *GetLogListRequest) SetTemplateId(v string) *GetLogListRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLogListRequest) SetIsSuccess(v string) *GetLogListRequest {
	s.IsSuccess = &v
	return s
}

func (s *GetLogListRequest) SetMinAffectRows(v int32) *GetLogListRequest {
	s.MinAffectRows = &v
	return s
}

func (s *GetLogListRequest) SetMaxAffectRows(v int32) *GetLogListRequest {
	s.MaxAffectRows = &v
	return s
}

func (s *GetLogListRequest) SetMinExecCostUS(v int32) *GetLogListRequest {
	s.MinExecCostUS = &v
	return s
}

func (s *GetLogListRequest) SetMaxExecCostUS(v int32) *GetLogListRequest {
	s.MaxExecCostUS = &v
	return s
}

func (s *GetLogListRequest) SetRuleName(v string) *GetLogListRequest {
	s.RuleName = &v
	return s
}

func (s *GetLogListRequest) SetClientIpList(v []*string) *GetLogListRequest {
	s.ClientIpList = v
	return s
}

func (s *GetLogListRequest) SetDbUserList(v []*string) *GetLogListRequest {
	s.DbUserList = v
	return s
}

func (s *GetLogListRequest) SetDbHostList(v []*string) *GetLogListRequest {
	s.DbHostList = v
	return s
}

func (s *GetLogListRequest) SetRiskLevelList(v []*string) *GetLogListRequest {
	s.RiskLevelList = v
	return s
}

func (s *GetLogListRequest) SetRuleTypeList(v []*string) *GetLogListRequest {
	s.RuleTypeList = v
	return s
}

func (s *GetLogListRequest) SetSqlTypeList(v []*string) *GetLogListRequest {
	s.SqlTypeList = v
	return s
}

func (s *GetLogListRequest) SetClientProgramList(v []*string) *GetLogListRequest {
	s.ClientProgramList = v
	return s
}

type GetLogListResponseBody struct {
	EndDate    *string                          `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BeginDate  *string                          `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	Incomplete *string                          `json:"Incomplete,omitempty" xml:"Incomplete,omitempty"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Results    []*GetLogListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s GetLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogListResponseBody) SetEndDate(v string) *GetLogListResponseBody {
	s.EndDate = &v
	return s
}

func (s *GetLogListResponseBody) SetRequestId(v string) *GetLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogListResponseBody) SetBeginDate(v string) *GetLogListResponseBody {
	s.BeginDate = &v
	return s
}

func (s *GetLogListResponseBody) SetIncomplete(v string) *GetLogListResponseBody {
	s.Incomplete = &v
	return s
}

func (s *GetLogListResponseBody) SetPageNumber(v int32) *GetLogListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetLogListResponseBody) SetPageSize(v int32) *GetLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLogListResponseBody) SetTotalCount(v int32) *GetLogListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetLogListResponseBody) SetResults(v []*GetLogListResponseBodyResults) *GetLogListResponseBody {
	s.Results = v
	return s
}

type GetLogListResponseBodyResults struct {
	ClientPort        *int32  `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	ExecCostUS        *int32  `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	AppClientIp       *string `json:"AppClientIp,omitempty" xml:"AppClientIp,omitempty"`
	SessionLogoutTime *string `json:"SessionLogoutTime,omitempty" xml:"SessionLogoutTime,omitempty"`
	ClientOsUser      *string `json:"ClientOsUser,omitempty" xml:"ClientOsUser,omitempty"`
	RuleId            *int32  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SqlId             *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SessionId         *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SqlType           *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	RiskLevel         *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	AppUsername       *string `json:"AppUsername,omitempty" xml:"AppUsername,omitempty"`
	DbId              *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	RuleType          *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	RuleKeyId         *int32  `json:"RuleKeyId,omitempty" xml:"RuleKeyId,omitempty"`
	AffectRows        *int32  `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	Schema            *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	SessionLoginTime  *string `json:"SessionLoginTime,omitempty" xml:"SessionLoginTime,omitempty"`
	DbUser            *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	ServerMac         *string `json:"ServerMac,omitempty" xml:"ServerMac,omitempty"`
	DbServer          *string `json:"DbServer,omitempty" xml:"DbServer,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SqlContent        *string `json:"SqlContent,omitempty" xml:"SqlContent,omitempty"`
	ResponseCode      *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	InstName          *string `json:"InstName,omitempty" xml:"InstName,omitempty"`
	ClientProgram     *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	CaptureTime       *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	ClientIp          *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientMac         *string `json:"ClientMac,omitempty" xml:"ClientMac,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FetchCostUS       *int32  `json:"FetchCostUS,omitempty" xml:"FetchCostUS,omitempty"`
	ResponseText      *string `json:"ResponseText,omitempty" xml:"ResponseText,omitempty"`
}

func (s GetLogListResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetLogListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetLogListResponseBodyResults) SetClientPort(v int32) *GetLogListResponseBodyResults {
	s.ClientPort = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetExecCostUS(v int32) *GetLogListResponseBodyResults {
	s.ExecCostUS = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetAppClientIp(v string) *GetLogListResponseBodyResults {
	s.AppClientIp = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSessionLogoutTime(v string) *GetLogListResponseBodyResults {
	s.SessionLogoutTime = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientOsUser(v string) *GetLogListResponseBodyResults {
	s.ClientOsUser = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleId(v int32) *GetLogListResponseBodyResults {
	s.RuleId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSqlId(v string) *GetLogListResponseBodyResults {
	s.SqlId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSessionId(v string) *GetLogListResponseBodyResults {
	s.SessionId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSqlType(v string) *GetLogListResponseBodyResults {
	s.SqlType = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRiskLevel(v int32) *GetLogListResponseBodyResults {
	s.RiskLevel = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetAppUsername(v string) *GetLogListResponseBodyResults {
	s.AppUsername = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetDbId(v int32) *GetLogListResponseBodyResults {
	s.DbId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleType(v int32) *GetLogListResponseBodyResults {
	s.RuleType = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleKeyId(v int32) *GetLogListResponseBodyResults {
	s.RuleKeyId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetAffectRows(v int32) *GetLogListResponseBodyResults {
	s.AffectRows = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSchema(v string) *GetLogListResponseBodyResults {
	s.Schema = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSessionLoginTime(v string) *GetLogListResponseBodyResults {
	s.SessionLoginTime = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetDbUser(v string) *GetLogListResponseBodyResults {
	s.DbUser = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetServerMac(v string) *GetLogListResponseBodyResults {
	s.ServerMac = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetDbServer(v string) *GetLogListResponseBodyResults {
	s.DbServer = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleName(v string) *GetLogListResponseBodyResults {
	s.RuleName = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSqlContent(v string) *GetLogListResponseBodyResults {
	s.SqlContent = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetResponseCode(v string) *GetLogListResponseBodyResults {
	s.ResponseCode = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetInstName(v string) *GetLogListResponseBodyResults {
	s.InstName = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientProgram(v string) *GetLogListResponseBodyResults {
	s.ClientProgram = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetCaptureTime(v string) *GetLogListResponseBodyResults {
	s.CaptureTime = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientIp(v string) *GetLogListResponseBodyResults {
	s.ClientIp = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientMac(v string) *GetLogListResponseBodyResults {
	s.ClientMac = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetTemplateId(v string) *GetLogListResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetFetchCostUS(v int32) *GetLogListResponseBodyResults {
	s.FetchCostUS = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetResponseText(v string) *GetLogListResponseBodyResults {
	s.ResponseText = &v
	return s
}

type GetLogListResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogListResponse) GoString() string {
	return s.String()
}

func (s *GetLogListResponse) SetHeaders(v map[string]*string) *GetLogListResponse {
	s.Headers = v
	return s
}

func (s *GetLogListResponse) SetBody(v *GetLogListResponseBody) *GetLogListResponse {
	s.Body = v
	return s
}

type GetLogMaskConfigsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskName   *string `json:"MaskName,omitempty" xml:"MaskName,omitempty"`
	MaskType   *int32  `json:"MaskType,omitempty" xml:"MaskType,omitempty"`
	MaskState  *int32  `json:"MaskState,omitempty" xml:"MaskState,omitempty"`
}

func (s GetLogMaskConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogMaskConfigsRequest) GoString() string {
	return s.String()
}

func (s *GetLogMaskConfigsRequest) SetRegionId(v string) *GetLogMaskConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogMaskConfigsRequest) SetInstanceId(v string) *GetLogMaskConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogMaskConfigsRequest) SetMaskName(v string) *GetLogMaskConfigsRequest {
	s.MaskName = &v
	return s
}

func (s *GetLogMaskConfigsRequest) SetMaskType(v int32) *GetLogMaskConfigsRequest {
	s.MaskType = &v
	return s
}

func (s *GetLogMaskConfigsRequest) SetMaskState(v int32) *GetLogMaskConfigsRequest {
	s.MaskState = &v
	return s
}

type GetLogMaskConfigsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Configs   []*GetLogMaskConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s GetLogMaskConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogMaskConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogMaskConfigsResponseBody) SetRequestId(v string) *GetLogMaskConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogMaskConfigsResponseBody) SetConfigs(v []*GetLogMaskConfigsResponseBodyConfigs) *GetLogMaskConfigsResponseBody {
	s.Configs = v
	return s
}

type GetLogMaskConfigsResponseBodyConfigs struct {
	MaskDescription *string `json:"MaskDescription,omitempty" xml:"MaskDescription,omitempty"`
	MaskState       *int32  `json:"MaskState,omitempty" xml:"MaskState,omitempty"`
	MaskName        *string `json:"MaskName,omitempty" xml:"MaskName,omitempty"`
	MaskRegex       *string `json:"MaskRegex,omitempty" xml:"MaskRegex,omitempty"`
	MaskTxt         *string `json:"MaskTxt,omitempty" xml:"MaskTxt,omitempty"`
	MaskId          *int32  `json:"MaskId,omitempty" xml:"MaskId,omitempty"`
	MaskType        *int32  `json:"MaskType,omitempty" xml:"MaskType,omitempty"`
}

func (s GetLogMaskConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetLogMaskConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskDescription(v string) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskDescription = &v
	return s
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskState(v int32) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskState = &v
	return s
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskName(v string) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskName = &v
	return s
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskRegex(v string) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskRegex = &v
	return s
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskTxt(v string) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskTxt = &v
	return s
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskId(v int32) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskId = &v
	return s
}

func (s *GetLogMaskConfigsResponseBodyConfigs) SetMaskType(v int32) *GetLogMaskConfigsResponseBodyConfigs {
	s.MaskType = &v
	return s
}

type GetLogMaskConfigsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogMaskConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogMaskConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogMaskConfigsResponse) GoString() string {
	return s.String()
}

func (s *GetLogMaskConfigsResponse) SetHeaders(v map[string]*string) *GetLogMaskConfigsResponse {
	s.Headers = v
	return s
}

func (s *GetLogMaskConfigsResponse) SetBody(v *GetLogMaskConfigsResponseBody) *GetLogMaskConfigsResponse {
	s.Body = v
	return s
}

type GetLogQueryConditionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	IsRisk     *string `json:"IsRisk,omitempty" xml:"IsRisk,omitempty"`
	IsSuccess  *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetLogQueryConditionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogQueryConditionRequest) GoString() string {
	return s.String()
}

func (s *GetLogQueryConditionRequest) SetRegionId(v string) *GetLogQueryConditionRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogQueryConditionRequest) SetInstanceId(v string) *GetLogQueryConditionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogQueryConditionRequest) SetDbId(v int32) *GetLogQueryConditionRequest {
	s.DbId = &v
	return s
}

func (s *GetLogQueryConditionRequest) SetBeginDate(v string) *GetLogQueryConditionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogQueryConditionRequest) SetEndDate(v string) *GetLogQueryConditionRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogQueryConditionRequest) SetIsRisk(v string) *GetLogQueryConditionRequest {
	s.IsRisk = &v
	return s
}

func (s *GetLogQueryConditionRequest) SetIsSuccess(v string) *GetLogQueryConditionRequest {
	s.IsSuccess = &v
	return s
}

type GetLogQueryConditionResponseBody struct {
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbUserList        []*string                                      `json:"DbUserList,omitempty" xml:"DbUserList,omitempty" type:"Repeated"`
	ClientIpList      []*string                                      `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	ClientProgramList []*string                                      `json:"ClientProgramList,omitempty" xml:"ClientProgramList,omitempty" type:"Repeated"`
	DbServerList      []*string                                      `json:"DbServerList,omitempty" xml:"DbServerList,omitempty" type:"Repeated"`
	SqlTypeList       []*GetLogQueryConditionResponseBodySqlTypeList `json:"SqlTypeList,omitempty" xml:"SqlTypeList,omitempty" type:"Repeated"`
}

func (s GetLogQueryConditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogQueryConditionResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogQueryConditionResponseBody) SetRequestId(v string) *GetLogQueryConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogQueryConditionResponseBody) SetDbUserList(v []*string) *GetLogQueryConditionResponseBody {
	s.DbUserList = v
	return s
}

func (s *GetLogQueryConditionResponseBody) SetClientIpList(v []*string) *GetLogQueryConditionResponseBody {
	s.ClientIpList = v
	return s
}

func (s *GetLogQueryConditionResponseBody) SetClientProgramList(v []*string) *GetLogQueryConditionResponseBody {
	s.ClientProgramList = v
	return s
}

func (s *GetLogQueryConditionResponseBody) SetDbServerList(v []*string) *GetLogQueryConditionResponseBody {
	s.DbServerList = v
	return s
}

func (s *GetLogQueryConditionResponseBody) SetSqlTypeList(v []*GetLogQueryConditionResponseBodySqlTypeList) *GetLogQueryConditionResponseBody {
	s.SqlTypeList = v
	return s
}

type GetLogQueryConditionResponseBodySqlTypeList struct {
	SqlKeyword *string `json:"SqlKeyword,omitempty" xml:"SqlKeyword,omitempty"`
	SqlType    *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s GetLogQueryConditionResponseBodySqlTypeList) String() string {
	return tea.Prettify(s)
}

func (s GetLogQueryConditionResponseBodySqlTypeList) GoString() string {
	return s.String()
}

func (s *GetLogQueryConditionResponseBodySqlTypeList) SetSqlKeyword(v string) *GetLogQueryConditionResponseBodySqlTypeList {
	s.SqlKeyword = &v
	return s
}

func (s *GetLogQueryConditionResponseBodySqlTypeList) SetSqlType(v string) *GetLogQueryConditionResponseBodySqlTypeList {
	s.SqlType = &v
	return s
}

type GetLogQueryConditionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogQueryConditionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogQueryConditionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogQueryConditionResponse) GoString() string {
	return s.String()
}

func (s *GetLogQueryConditionResponse) SetHeaders(v map[string]*string) *GetLogQueryConditionResponse {
	s.Headers = v
	return s
}

func (s *GetLogQueryConditionResponse) SetBody(v *GetLogQueryConditionResponseBody) *GetLogQueryConditionResponse {
	s.Body = v
	return s
}

type GetLogStatisticsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId            *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate       *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate         *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ByDbId          *int32  `json:"ByDbId,omitempty" xml:"ByDbId,omitempty"`
	ByClientIp      *int32  `json:"ByClientIp,omitempty" xml:"ByClientIp,omitempty"`
	ByDbUser        *int32  `json:"ByDbUser,omitempty" xml:"ByDbUser,omitempty"`
	ByDbServer      *int32  `json:"ByDbServer,omitempty" xml:"ByDbServer,omitempty"`
	ByClientProgram *int32  `json:"ByClientProgram,omitempty" xml:"ByClientProgram,omitempty"`
	BySqlType       *int32  `json:"BySqlType,omitempty" xml:"BySqlType,omitempty"`
}

func (s GetLogStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetLogStatisticsRequest) SetRegionId(v string) *GetLogStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogStatisticsRequest) SetInstanceId(v string) *GetLogStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogStatisticsRequest) SetDbId(v int32) *GetLogStatisticsRequest {
	s.DbId = &v
	return s
}

func (s *GetLogStatisticsRequest) SetBeginDate(v string) *GetLogStatisticsRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogStatisticsRequest) SetEndDate(v string) *GetLogStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogStatisticsRequest) SetByDbId(v int32) *GetLogStatisticsRequest {
	s.ByDbId = &v
	return s
}

func (s *GetLogStatisticsRequest) SetByClientIp(v int32) *GetLogStatisticsRequest {
	s.ByClientIp = &v
	return s
}

func (s *GetLogStatisticsRequest) SetByDbUser(v int32) *GetLogStatisticsRequest {
	s.ByDbUser = &v
	return s
}

func (s *GetLogStatisticsRequest) SetByDbServer(v int32) *GetLogStatisticsRequest {
	s.ByDbServer = &v
	return s
}

func (s *GetLogStatisticsRequest) SetByClientProgram(v int32) *GetLogStatisticsRequest {
	s.ByClientProgram = &v
	return s
}

func (s *GetLogStatisticsRequest) SetBySqlType(v int32) *GetLogStatisticsRequest {
	s.BySqlType = &v
	return s
}

type GetLogStatisticsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CountList []*GetLogStatisticsResponseBodyCountList `json:"CountList,omitempty" xml:"CountList,omitempty" type:"Repeated"`
}

func (s GetLogStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogStatisticsResponseBody) SetRequestId(v string) *GetLogStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogStatisticsResponseBody) SetCountList(v []*GetLogStatisticsResponseBodyCountList) *GetLogStatisticsResponseBody {
	s.CountList = v
	return s
}

type GetLogStatisticsResponseBodyCountList struct {
	SqlKeyword    *string `json:"SqlKeyword,omitempty" xml:"SqlKeyword,omitempty"`
	DbId          *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName        *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ClientProgram *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	DbUser        *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	SqlCount      *int64  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	RiskCount     *int64  `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	DbServer      *string `json:"DbServer,omitempty" xml:"DbServer,omitempty"`
	SqlType       *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s GetLogStatisticsResponseBodyCountList) String() string {
	return tea.Prettify(s)
}

func (s GetLogStatisticsResponseBodyCountList) GoString() string {
	return s.String()
}

func (s *GetLogStatisticsResponseBodyCountList) SetSqlKeyword(v string) *GetLogStatisticsResponseBodyCountList {
	s.SqlKeyword = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetDbId(v int32) *GetLogStatisticsResponseBodyCountList {
	s.DbId = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetDbName(v string) *GetLogStatisticsResponseBodyCountList {
	s.DbName = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetClientProgram(v string) *GetLogStatisticsResponseBodyCountList {
	s.ClientProgram = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetDbUser(v string) *GetLogStatisticsResponseBodyCountList {
	s.DbUser = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetClientIp(v string) *GetLogStatisticsResponseBodyCountList {
	s.ClientIp = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetSqlCount(v int64) *GetLogStatisticsResponseBodyCountList {
	s.SqlCount = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetRiskCount(v int64) *GetLogStatisticsResponseBodyCountList {
	s.RiskCount = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetDbServer(v string) *GetLogStatisticsResponseBodyCountList {
	s.DbServer = &v
	return s
}

func (s *GetLogStatisticsResponseBodyCountList) SetSqlType(v string) *GetLogStatisticsResponseBodyCountList {
	s.SqlType = &v
	return s
}

type GetLogStatisticsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetLogStatisticsResponse) SetHeaders(v map[string]*string) *GetLogStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetLogStatisticsResponse) SetBody(v *GetLogStatisticsResponseBody) *GetLogStatisticsResponse {
	s.Body = v
	return s
}

type GetLogTopDistributionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetLogTopDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetLogTopDistributionRequest) SetRegionId(v string) *GetLogTopDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogTopDistributionRequest) SetInstanceId(v string) *GetLogTopDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogTopDistributionRequest) SetDbId(v int32) *GetLogTopDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetLogTopDistributionRequest) SetBeginDate(v string) *GetLogTopDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogTopDistributionRequest) SetEndDate(v string) *GetLogTopDistributionRequest {
	s.EndDate = &v
	return s
}

type GetLogTopDistributionResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetLogTopDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetLogTopDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogTopDistributionResponseBody) SetRequestId(v string) *GetLogTopDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogTopDistributionResponseBody) SetTimeList(v []*GetLogTopDistributionResponseBodyTimeList) *GetLogTopDistributionResponseBody {
	s.TimeList = v
	return s
}

type GetLogTopDistributionResponseBodyTimeList struct {
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	SqlCount  *int64  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	TopTime   *string `json:"TopTime,omitempty" xml:"TopTime,omitempty"`
}

func (s GetLogTopDistributionResponseBodyTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetLogTopDistributionResponseBodyTimeList) SetBeginDate(v string) *GetLogTopDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetLogTopDistributionResponseBodyTimeList) SetSqlCount(v int64) *GetLogTopDistributionResponseBodyTimeList {
	s.SqlCount = &v
	return s
}

func (s *GetLogTopDistributionResponseBodyTimeList) SetEndDate(v string) *GetLogTopDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetLogTopDistributionResponseBodyTimeList) SetTopTime(v string) *GetLogTopDistributionResponseBodyTimeList {
	s.TopTime = &v
	return s
}

type GetLogTopDistributionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogTopDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogTopDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetLogTopDistributionResponse) SetHeaders(v map[string]*string) *GetLogTopDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetLogTopDistributionResponse) SetBody(v *GetLogTopDistributionResponseBody) *GetLogTopDistributionResponse {
	s.Body = v
	return s
}

type GetLogTopStatisticsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	TopTime    *string `json:"TopTime,omitempty" xml:"TopTime,omitempty"`
}

func (s GetLogTopStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetLogTopStatisticsRequest) SetRegionId(v string) *GetLogTopStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogTopStatisticsRequest) SetInstanceId(v string) *GetLogTopStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogTopStatisticsRequest) SetDbId(v int32) *GetLogTopStatisticsRequest {
	s.DbId = &v
	return s
}

func (s *GetLogTopStatisticsRequest) SetTopTime(v string) *GetLogTopStatisticsRequest {
	s.TopTime = &v
	return s
}

type GetLogTopStatisticsResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CountList []*GetLogTopStatisticsResponseBodyCountList `json:"CountList,omitempty" xml:"CountList,omitempty" type:"Repeated"`
}

func (s GetLogTopStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogTopStatisticsResponseBody) SetRequestId(v string) *GetLogTopStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogTopStatisticsResponseBody) SetCountList(v []*GetLogTopStatisticsResponseBodyCountList) *GetLogTopStatisticsResponseBody {
	s.CountList = v
	return s
}

type GetLogTopStatisticsResponseBodyCountList struct {
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	DbUser        *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	SqlCount      *int64  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	ClientProgram *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
}

func (s GetLogTopStatisticsResponseBodyCountList) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopStatisticsResponseBodyCountList) GoString() string {
	return s.String()
}

func (s *GetLogTopStatisticsResponseBodyCountList) SetClientIp(v string) *GetLogTopStatisticsResponseBodyCountList {
	s.ClientIp = &v
	return s
}

func (s *GetLogTopStatisticsResponseBodyCountList) SetDbUser(v string) *GetLogTopStatisticsResponseBodyCountList {
	s.DbUser = &v
	return s
}

func (s *GetLogTopStatisticsResponseBodyCountList) SetSqlCount(v int64) *GetLogTopStatisticsResponseBodyCountList {
	s.SqlCount = &v
	return s
}

func (s *GetLogTopStatisticsResponseBodyCountList) SetClientProgram(v string) *GetLogTopStatisticsResponseBodyCountList {
	s.ClientProgram = &v
	return s
}

type GetLogTopStatisticsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogTopStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogTopStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogTopStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetLogTopStatisticsResponse) SetHeaders(v map[string]*string) *GetLogTopStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetLogTopStatisticsResponse) SetBody(v *GetLogTopStatisticsResponseBody) *GetLogTopStatisticsResponse {
	s.Body = v
	return s
}

type GetLogTypeDistributionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetLogTypeDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogTypeDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionRequest) SetRegionId(v string) *GetLogTypeDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetInstanceId(v string) *GetLogTypeDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetDbId(v int32) *GetLogTypeDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetBeginDate(v string) *GetLogTypeDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetEndDate(v string) *GetLogTypeDistributionRequest {
	s.EndDate = &v
	return s
}

type GetLogTypeDistributionResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetLogTypeDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetLogTypeDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogTypeDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionResponseBody) SetRequestId(v string) *GetLogTypeDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogTypeDistributionResponseBody) SetTimeList(v []*GetLogTypeDistributionResponseBodyTimeList) *GetLogTypeDistributionResponseBody {
	s.TimeList = v
	return s
}

type GetLogTypeDistributionResponseBodyTimeList struct {
	EndDate        *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ExecCostUS     *string `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	InsertSqlCount *int32  `json:"InsertSqlCount,omitempty" xml:"InsertSqlCount,omitempty"`
	SelectSqlCount *int32  `json:"SelectSqlCount,omitempty" xml:"SelectSqlCount,omitempty"`
	DeleteSqlCount *int32  `json:"DeleteSqlCount,omitempty" xml:"DeleteSqlCount,omitempty"`
	BeginDate      *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	OtherSqlCount  *int32  `json:"OtherSqlCount,omitempty" xml:"OtherSqlCount,omitempty"`
	SqlCount       *int32  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	UpdateSqlCount *int32  `json:"UpdateSqlCount,omitempty" xml:"UpdateSqlCount,omitempty"`
}

func (s GetLogTypeDistributionResponseBodyTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetLogTypeDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetEndDate(v string) *GetLogTypeDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetExecCostUS(v string) *GetLogTypeDistributionResponseBodyTimeList {
	s.ExecCostUS = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetInsertSqlCount(v int32) *GetLogTypeDistributionResponseBodyTimeList {
	s.InsertSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetSelectSqlCount(v int32) *GetLogTypeDistributionResponseBodyTimeList {
	s.SelectSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetDeleteSqlCount(v int32) *GetLogTypeDistributionResponseBodyTimeList {
	s.DeleteSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetBeginDate(v string) *GetLogTypeDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetOtherSqlCount(v int32) *GetLogTypeDistributionResponseBodyTimeList {
	s.OtherSqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetSqlCount(v int32) *GetLogTypeDistributionResponseBodyTimeList {
	s.SqlCount = &v
	return s
}

func (s *GetLogTypeDistributionResponseBodyTimeList) SetUpdateSqlCount(v int32) *GetLogTypeDistributionResponseBodyTimeList {
	s.UpdateSqlCount = &v
	return s
}

type GetLogTypeDistributionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogTypeDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogTypeDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogTypeDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionResponse) SetHeaders(v map[string]*string) *GetLogTypeDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetLogTypeDistributionResponse) SetBody(v *GetLogTypeDistributionResponseBody) *GetLogTypeDistributionResponse {
	s.Body = v
	return s
}

type GetNewSqlTemplateListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetNewSqlTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNewSqlTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetNewSqlTemplateListRequest) SetRegionId(v string) *GetNewSqlTemplateListRequest {
	s.RegionId = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetInstanceId(v string) *GetNewSqlTemplateListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetDbId(v int32) *GetNewSqlTemplateListRequest {
	s.DbId = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetBeginDate(v string) *GetNewSqlTemplateListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetEndDate(v string) *GetNewSqlTemplateListRequest {
	s.EndDate = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetPageNumber(v int32) *GetNewSqlTemplateListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetPageSize(v int32) *GetNewSqlTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *GetNewSqlTemplateListRequest) SetTemplateId(v string) *GetNewSqlTemplateListRequest {
	s.TemplateId = &v
	return s
}

type GetNewSqlTemplateListResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Results    []*GetNewSqlTemplateListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s GetNewSqlTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNewSqlTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetNewSqlTemplateListResponseBody) SetRequestId(v string) *GetNewSqlTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNewSqlTemplateListResponseBody) SetPageNumber(v int32) *GetNewSqlTemplateListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetNewSqlTemplateListResponseBody) SetPageSize(v int32) *GetNewSqlTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetNewSqlTemplateListResponseBody) SetTotalCount(v int64) *GetNewSqlTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetNewSqlTemplateListResponseBody) SetResults(v []*GetNewSqlTemplateListResponseBodyResults) *GetNewSqlTemplateListResponseBody {
	s.Results = v
	return s
}

type GetNewSqlTemplateListResponseBodyResults struct {
	TemplateContent  *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	TemplateId       *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FirstCaptureTime *string `json:"FirstCaptureTime,omitempty" xml:"FirstCaptureTime,omitempty"`
}

func (s GetNewSqlTemplateListResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetNewSqlTemplateListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetNewSqlTemplateListResponseBodyResults) SetTemplateContent(v string) *GetNewSqlTemplateListResponseBodyResults {
	s.TemplateContent = &v
	return s
}

func (s *GetNewSqlTemplateListResponseBodyResults) SetTemplateId(v string) *GetNewSqlTemplateListResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *GetNewSqlTemplateListResponseBodyResults) SetFirstCaptureTime(v string) *GetNewSqlTemplateListResponseBodyResults {
	s.FirstCaptureTime = &v
	return s
}

type GetNewSqlTemplateListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNewSqlTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNewSqlTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNewSqlTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetNewSqlTemplateListResponse) SetHeaders(v map[string]*string) *GetNewSqlTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetNewSqlTemplateListResponse) SetBody(v *GetNewSqlTemplateListResponseBody) *GetNewSqlTemplateListResponse {
	s.Body = v
	return s
}

type GetReportFileUrlRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetReportFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetReportFileUrlRequest) SetRegionId(v string) *GetReportFileUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetReportFileUrlRequest) SetInstanceId(v string) *GetReportFileUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetReportFileUrlRequest) SetFileName(v string) *GetReportFileUrlRequest {
	s.FileName = &v
	return s
}

type GetReportFileUrlResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *GetReportFileUrlResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s GetReportFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportFileUrlResponseBody) SetRequestId(v string) *GetReportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportFileUrlResponseBody) SetServerData(v *GetReportFileUrlResponseBodyServerData) *GetReportFileUrlResponseBody {
	s.ServerData = v
	return s
}

type GetReportFileUrlResponseBodyServerData struct {
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetReportFileUrlResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s GetReportFileUrlResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *GetReportFileUrlResponseBodyServerData) SetFileUrl(v string) *GetReportFileUrlResponseBodyServerData {
	s.FileUrl = &v
	return s
}

type GetReportFileUrlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetReportFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetReportFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReportFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetReportFileUrlResponse) SetHeaders(v map[string]*string) *GetReportFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetReportFileUrlResponse) SetBody(v *GetReportFileUrlResponseBody) *GetReportFileUrlResponse {
	s.Body = v
	return s
}

type GetRiskLevelDistributionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetRiskLevelDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRiskLevelDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionRequest) SetRegionId(v string) *GetRiskLevelDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetInstanceId(v string) *GetRiskLevelDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetDbId(v int32) *GetRiskLevelDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetBeginDate(v string) *GetRiskLevelDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetEndDate(v string) *GetRiskLevelDistributionRequest {
	s.EndDate = &v
	return s
}

type GetRiskLevelDistributionResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetRiskLevelDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetRiskLevelDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRiskLevelDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionResponseBody) SetRequestId(v string) *GetRiskLevelDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBody) SetTimeList(v []*GetRiskLevelDistributionResponseBodyTimeList) *GetRiskLevelDistributionResponseBody {
	s.TimeList = v
	return s
}

type GetRiskLevelDistributionResponseBodyTimeList struct {
	MiddleRiskCount *int64  `json:"MiddleRiskCount,omitempty" xml:"MiddleRiskCount,omitempty"`
	HighRiskCount   *int64  `json:"HighRiskCount,omitempty" xml:"HighRiskCount,omitempty"`
	EndDate         *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	BeginDate       *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	RiskCount       *int64  `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	LowRiskCount    *int64  `json:"LowRiskCount,omitempty" xml:"LowRiskCount,omitempty"`
}

func (s GetRiskLevelDistributionResponseBodyTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetRiskLevelDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetMiddleRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.MiddleRiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetHighRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.HighRiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetEndDate(v string) *GetRiskLevelDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetBeginDate(v string) *GetRiskLevelDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.RiskCount = &v
	return s
}

func (s *GetRiskLevelDistributionResponseBodyTimeList) SetLowRiskCount(v int64) *GetRiskLevelDistributionResponseBodyTimeList {
	s.LowRiskCount = &v
	return s
}

type GetRiskLevelDistributionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRiskLevelDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRiskLevelDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRiskLevelDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionResponse) SetHeaders(v map[string]*string) *GetRiskLevelDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetRiskLevelDistributionResponse) SetBody(v *GetRiskLevelDistributionResponseBody) *GetRiskLevelDistributionResponse {
	s.Body = v
	return s
}

type GetRiskStatisticsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId        *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate   *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ByDbId      *int32  `json:"ByDbId,omitempty" xml:"ByDbId,omitempty"`
	ByRiskLevel *int32  `json:"ByRiskLevel,omitempty" xml:"ByRiskLevel,omitempty"`
	ByRuleType  *int32  `json:"ByRuleType,omitempty" xml:"ByRuleType,omitempty"`
	ByRuleId    *int32  `json:"ByRuleId,omitempty" xml:"ByRuleId,omitempty"`
}

func (s GetRiskStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRiskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetRiskStatisticsRequest) SetRegionId(v string) *GetRiskStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetInstanceId(v string) *GetRiskStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetDbId(v int32) *GetRiskStatisticsRequest {
	s.DbId = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetBeginDate(v string) *GetRiskStatisticsRequest {
	s.BeginDate = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetEndDate(v string) *GetRiskStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetByDbId(v int32) *GetRiskStatisticsRequest {
	s.ByDbId = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetByRiskLevel(v int32) *GetRiskStatisticsRequest {
	s.ByRiskLevel = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetByRuleType(v int32) *GetRiskStatisticsRequest {
	s.ByRuleType = &v
	return s
}

func (s *GetRiskStatisticsRequest) SetByRuleId(v int32) *GetRiskStatisticsRequest {
	s.ByRuleId = &v
	return s
}

type GetRiskStatisticsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetRiskStatisticsResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetRiskStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRiskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskStatisticsResponseBody) SetRequestId(v string) *GetRiskStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskStatisticsResponseBody) SetTimeList(v []*GetRiskStatisticsResponseBodyTimeList) *GetRiskStatisticsResponseBody {
	s.TimeList = v
	return s
}

type GetRiskStatisticsResponseBodyTimeList struct {
	RiskLevel       *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	DbId            *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RuleType        *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	LastCaptureTime *string `json:"LastCaptureTime,omitempty" xml:"LastCaptureTime,omitempty"`
	RiskCount       *int64  `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId          *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetRiskStatisticsResponseBodyTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetRiskStatisticsResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetRiskLevel(v int32) *GetRiskStatisticsResponseBodyTimeList {
	s.RiskLevel = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetDbId(v int32) *GetRiskStatisticsResponseBodyTimeList {
	s.DbId = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetDbName(v string) *GetRiskStatisticsResponseBodyTimeList {
	s.DbName = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetRuleType(v int32) *GetRiskStatisticsResponseBodyTimeList {
	s.RuleType = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetLastCaptureTime(v string) *GetRiskStatisticsResponseBodyTimeList {
	s.LastCaptureTime = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetRiskCount(v int64) *GetRiskStatisticsResponseBodyTimeList {
	s.RiskCount = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetRuleName(v string) *GetRiskStatisticsResponseBodyTimeList {
	s.RuleName = &v
	return s
}

func (s *GetRiskStatisticsResponseBodyTimeList) SetRuleId(v string) *GetRiskStatisticsResponseBodyTimeList {
	s.RuleId = &v
	return s
}

type GetRiskStatisticsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRiskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRiskStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRiskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetRiskStatisticsResponse) SetHeaders(v map[string]*string) *GetRiskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetRiskStatisticsResponse) SetBody(v *GetRiskStatisticsResponseBody) *GetRiskStatisticsResponse {
	s.Body = v
	return s
}

type GetRuleDetailRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *int32  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleKeyId  *int32  `json:"RuleKeyId,omitempty" xml:"RuleKeyId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
}

func (s GetRuleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetRuleDetailRequest) SetRegionId(v string) *GetRuleDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetRuleDetailRequest) SetInstanceId(v string) *GetRuleDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRuleDetailRequest) SetRuleId(v int32) *GetRuleDetailRequest {
	s.RuleId = &v
	return s
}

func (s *GetRuleDetailRequest) SetRuleKeyId(v int32) *GetRuleDetailRequest {
	s.RuleKeyId = &v
	return s
}

func (s *GetRuleDetailRequest) SetDbId(v int32) *GetRuleDetailRequest {
	s.DbId = &v
	return s
}

type GetRuleDetailResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *GetRuleDetailResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBody) SetRequestId(v string) *GetRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetServerData(v *GetRuleDetailResponseBodyServerData) *GetRuleDetailResponseBody {
	s.ServerData = v
	return s
}

type GetRuleDetailResponseBodyServerData struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
}

func (s GetRuleDetailResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyServerData) SetJsonResult(v string) *GetRuleDetailResponseBodyServerData {
	s.JsonResult = &v
	return s
}

type GetRuleDetailResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRuleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponse) SetHeaders(v map[string]*string) *GetRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetRuleDetailResponse) SetBody(v *GetRuleDetailResponseBody) *GetRuleDetailResponse {
	s.Body = v
	return s
}

type GetRuleGroupNameRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
}

func (s GetRuleGroupNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleGroupNameRequest) GoString() string {
	return s.String()
}

func (s *GetRuleGroupNameRequest) SetRegionId(v string) *GetRuleGroupNameRequest {
	s.RegionId = &v
	return s
}

func (s *GetRuleGroupNameRequest) SetInstanceId(v string) *GetRuleGroupNameRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRuleGroupNameRequest) SetGroupName(v string) *GetRuleGroupNameRequest {
	s.GroupName = &v
	return s
}

func (s *GetRuleGroupNameRequest) SetGroupType(v string) *GetRuleGroupNameRequest {
	s.GroupType = &v
	return s
}

type GetRuleGroupNameResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRuleGroupNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleGroupNameResponseBody) SetServerData(v string) *GetRuleGroupNameResponseBody {
	s.ServerData = &v
	return s
}

func (s *GetRuleGroupNameResponseBody) SetRequestId(v string) *GetRuleGroupNameResponseBody {
	s.RequestId = &v
	return s
}

type GetRuleGroupNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRuleGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRuleGroupNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleGroupNameResponse) GoString() string {
	return s.String()
}

func (s *GetRuleGroupNameResponse) SetHeaders(v map[string]*string) *GetRuleGroupNameResponse {
	s.Headers = v
	return s
}

func (s *GetRuleGroupNameResponse) SetBody(v *GetRuleGroupNameResponseBody) *GetRuleGroupNameResponse {
	s.Body = v
	return s
}

type GetSessionDetailRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetSessionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSessionDetailRequest) SetRegionId(v string) *GetSessionDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionDetailRequest) SetInstanceId(v string) *GetSessionDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSessionDetailRequest) SetSessionId(v string) *GetSessionDetailRequest {
	s.SessionId = &v
	return s
}

type GetSessionDetailResponseBody struct {
	LoginTime     *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	DbId          *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	LoginCode     *int32  `json:"LoginCode,omitempty" xml:"LoginCode,omitempty"`
	ClientPort    *int32  `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	LoginMessage  *string `json:"LoginMessage,omitempty" xml:"LoginMessage,omitempty"`
	DbUser        *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	LogoutTime    *string `json:"LogoutTime,omitempty" xml:"LogoutTime,omitempty"`
	ServerPort    *int32  `json:"ServerPort,omitempty" xml:"ServerPort,omitempty"`
	ClientOsUser  *string `json:"ClientOsUser,omitempty" xml:"ClientOsUser,omitempty"`
	ServerMac     *string `json:"ServerMac,omitempty" xml:"ServerMac,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClientProgram *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ClientMac     *string `json:"ClientMac,omitempty" xml:"ClientMac,omitempty"`
}

func (s GetSessionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionDetailResponseBody) SetLoginTime(v string) *GetSessionDetailResponseBody {
	s.LoginTime = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetDbId(v int32) *GetSessionDetailResponseBody {
	s.DbId = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetLoginCode(v int32) *GetSessionDetailResponseBody {
	s.LoginCode = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetClientPort(v int32) *GetSessionDetailResponseBody {
	s.ClientPort = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetLoginMessage(v string) *GetSessionDetailResponseBody {
	s.LoginMessage = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetDbUser(v string) *GetSessionDetailResponseBody {
	s.DbUser = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetLogoutTime(v string) *GetSessionDetailResponseBody {
	s.LogoutTime = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetServerPort(v int32) *GetSessionDetailResponseBody {
	s.ServerPort = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetClientOsUser(v string) *GetSessionDetailResponseBody {
	s.ClientOsUser = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetServerMac(v string) *GetSessionDetailResponseBody {
	s.ServerMac = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetRequestId(v string) *GetSessionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetClientProgram(v string) *GetSessionDetailResponseBody {
	s.ClientProgram = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetClientIp(v string) *GetSessionDetailResponseBody {
	s.ClientIp = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetServerIp(v string) *GetSessionDetailResponseBody {
	s.ServerIp = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetSessionId(v string) *GetSessionDetailResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetSessionDetailResponseBody) SetClientMac(v string) *GetSessionDetailResponseBody {
	s.ClientMac = &v
	return s
}

type GetSessionDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSessionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSessionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSessionDetailResponse) SetHeaders(v map[string]*string) *GetSessionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSessionDetailResponse) SetBody(v *GetSessionDetailResponseBody) *GetSessionDetailResponse {
	s.Body = v
	return s
}

type GetSessionDistributionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetSessionDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionRequest) SetRegionId(v string) *GetSessionDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionDistributionRequest) SetInstanceId(v string) *GetSessionDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSessionDistributionRequest) SetDbId(v int32) *GetSessionDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetSessionDistributionRequest) SetBeginDate(v string) *GetSessionDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetSessionDistributionRequest) SetEndDate(v string) *GetSessionDistributionRequest {
	s.EndDate = &v
	return s
}

type GetSessionDistributionResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeList  []*GetSessionDistributionResponseBodyTimeList `json:"TimeList,omitempty" xml:"TimeList,omitempty" type:"Repeated"`
}

func (s GetSessionDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionResponseBody) SetRequestId(v string) *GetSessionDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionDistributionResponseBody) SetTimeList(v []*GetSessionDistributionResponseBodyTimeList) *GetSessionDistributionResponseBody {
	s.TimeList = v
	return s
}

type GetSessionDistributionResponseBodyTimeList struct {
	BeginDate          *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	LoginSessionCount  *int32  `json:"LoginSessionCount,omitempty" xml:"LoginSessionCount,omitempty"`
	ActiveSessionCount *int32  `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	EndDate            *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ErrorSessionCount  *int32  `json:"ErrorSessionCount,omitempty" xml:"ErrorSessionCount,omitempty"`
}

func (s GetSessionDistributionResponseBodyTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDistributionResponseBodyTimeList) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionResponseBodyTimeList) SetBeginDate(v string) *GetSessionDistributionResponseBodyTimeList {
	s.BeginDate = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetLoginSessionCount(v int32) *GetSessionDistributionResponseBodyTimeList {
	s.LoginSessionCount = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetActiveSessionCount(v int32) *GetSessionDistributionResponseBodyTimeList {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetEndDate(v string) *GetSessionDistributionResponseBodyTimeList {
	s.EndDate = &v
	return s
}

func (s *GetSessionDistributionResponseBodyTimeList) SetErrorSessionCount(v int32) *GetSessionDistributionResponseBodyTimeList {
	s.ErrorSessionCount = &v
	return s
}

type GetSessionDistributionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSessionDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSessionDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionResponse) SetHeaders(v map[string]*string) *GetSessionDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetSessionDistributionResponse) SetBody(v *GetSessionDistributionResponseBody) *GetSessionDistributionResponse {
	s.Body = v
	return s
}

type GetSessionListRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId        *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId              *int32    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate         *string   `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate           *string   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageNumber        *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SessionId         *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ActionList        []*string `json:"ActionList,omitempty" xml:"ActionList,omitempty" type:"Repeated"`
	ClientIpList      []*string `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	DbUserList        []*string `json:"DbUserList,omitempty" xml:"DbUserList,omitempty" type:"Repeated"`
	DbHostList        []*string `json:"DbHostList,omitempty" xml:"DbHostList,omitempty" type:"Repeated"`
	ClientProgramList []*string `json:"ClientProgramList,omitempty" xml:"ClientProgramList,omitempty" type:"Repeated"`
}

func (s GetSessionListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSessionListRequest) GoString() string {
	return s.String()
}

func (s *GetSessionListRequest) SetRegionId(v string) *GetSessionListRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionListRequest) SetInstanceId(v string) *GetSessionListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSessionListRequest) SetDbId(v int32) *GetSessionListRequest {
	s.DbId = &v
	return s
}

func (s *GetSessionListRequest) SetBeginDate(v string) *GetSessionListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetSessionListRequest) SetEndDate(v string) *GetSessionListRequest {
	s.EndDate = &v
	return s
}

func (s *GetSessionListRequest) SetPageNumber(v int32) *GetSessionListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSessionListRequest) SetPageSize(v int32) *GetSessionListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSessionListRequest) SetSessionId(v string) *GetSessionListRequest {
	s.SessionId = &v
	return s
}

func (s *GetSessionListRequest) SetActionList(v []*string) *GetSessionListRequest {
	s.ActionList = v
	return s
}

func (s *GetSessionListRequest) SetClientIpList(v []*string) *GetSessionListRequest {
	s.ClientIpList = v
	return s
}

func (s *GetSessionListRequest) SetDbUserList(v []*string) *GetSessionListRequest {
	s.DbUserList = v
	return s
}

func (s *GetSessionListRequest) SetDbHostList(v []*string) *GetSessionListRequest {
	s.DbHostList = v
	return s
}

func (s *GetSessionListRequest) SetClientProgramList(v []*string) *GetSessionListRequest {
	s.ClientProgramList = v
	return s
}

type GetSessionListResponseBody struct {
	EndDate    *string                              `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BeginDate  *string                              `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	Incomplete *string                              `json:"Incomplete,omitempty" xml:"Incomplete,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Results    []*GetSessionListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s GetSessionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionListResponseBody) SetEndDate(v string) *GetSessionListResponseBody {
	s.EndDate = &v
	return s
}

func (s *GetSessionListResponseBody) SetRequestId(v string) *GetSessionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionListResponseBody) SetBeginDate(v string) *GetSessionListResponseBody {
	s.BeginDate = &v
	return s
}

func (s *GetSessionListResponseBody) SetIncomplete(v string) *GetSessionListResponseBody {
	s.Incomplete = &v
	return s
}

func (s *GetSessionListResponseBody) SetPageNumber(v int32) *GetSessionListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetSessionListResponseBody) SetPageSize(v int32) *GetSessionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSessionListResponseBody) SetTotalCount(v int64) *GetSessionListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSessionListResponseBody) SetResults(v []*GetSessionListResponseBodyResults) *GetSessionListResponseBody {
	s.Results = v
	return s
}

type GetSessionListResponseBodyResults struct {
	DbId          *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	LoginCode     *int32  `json:"LoginCode,omitempty" xml:"LoginCode,omitempty"`
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ClientPort    *int32  `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	LoginMessage  *string `json:"LoginMessage,omitempty" xml:"LoginMessage,omitempty"`
	DbUser        *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	ServerPort    *int32  `json:"ServerPort,omitempty" xml:"ServerPort,omitempty"`
	ClientOsUser  *string `json:"ClientOsUser,omitempty" xml:"ClientOsUser,omitempty"`
	ServerMac     *string `json:"ServerMac,omitempty" xml:"ServerMac,omitempty"`
	ClientProgram *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	CaptureTime   *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ClientMac     *string `json:"ClientMac,omitempty" xml:"ClientMac,omitempty"`
}

func (s GetSessionListResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetSessionListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetSessionListResponseBodyResults) SetDbId(v int32) *GetSessionListResponseBodyResults {
	s.DbId = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetLoginCode(v int32) *GetSessionListResponseBodyResults {
	s.LoginCode = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetAction(v string) *GetSessionListResponseBodyResults {
	s.Action = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientPort(v int32) *GetSessionListResponseBodyResults {
	s.ClientPort = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetLoginMessage(v string) *GetSessionListResponseBodyResults {
	s.LoginMessage = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetDbUser(v string) *GetSessionListResponseBodyResults {
	s.DbUser = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetServerPort(v int32) *GetSessionListResponseBodyResults {
	s.ServerPort = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientOsUser(v string) *GetSessionListResponseBodyResults {
	s.ClientOsUser = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetServerMac(v string) *GetSessionListResponseBodyResults {
	s.ServerMac = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientProgram(v string) *GetSessionListResponseBodyResults {
	s.ClientProgram = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetCaptureTime(v string) *GetSessionListResponseBodyResults {
	s.CaptureTime = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientIp(v string) *GetSessionListResponseBodyResults {
	s.ClientIp = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetServerIp(v string) *GetSessionListResponseBodyResults {
	s.ServerIp = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetSessionId(v string) *GetSessionListResponseBodyResults {
	s.SessionId = &v
	return s
}

func (s *GetSessionListResponseBodyResults) SetClientMac(v string) *GetSessionListResponseBodyResults {
	s.ClientMac = &v
	return s
}

type GetSessionListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSessionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSessionListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionListResponse) GoString() string {
	return s.String()
}

func (s *GetSessionListResponse) SetHeaders(v map[string]*string) *GetSessionListResponse {
	s.Headers = v
	return s
}

func (s *GetSessionListResponse) SetBody(v *GetSessionListResponseBody) *GetSessionListResponse {
	s.Body = v
	return s
}

type GetSessionQueryConditionRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetSessionQueryConditionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSessionQueryConditionRequest) GoString() string {
	return s.String()
}

func (s *GetSessionQueryConditionRequest) SetRegionId(v string) *GetSessionQueryConditionRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionQueryConditionRequest) SetInstanceId(v string) *GetSessionQueryConditionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSessionQueryConditionRequest) SetDbId(v int32) *GetSessionQueryConditionRequest {
	s.DbId = &v
	return s
}

func (s *GetSessionQueryConditionRequest) SetBeginDate(v string) *GetSessionQueryConditionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetSessionQueryConditionRequest) SetEndDate(v string) *GetSessionQueryConditionRequest {
	s.EndDate = &v
	return s
}

type GetSessionQueryConditionResponseBody struct {
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbUserList        []*string `json:"DbUserList,omitempty" xml:"DbUserList,omitempty" type:"Repeated"`
	ClientIpList      []*string `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	ClientProgramList []*string `json:"ClientProgramList,omitempty" xml:"ClientProgramList,omitempty" type:"Repeated"`
	DbServerList      []*string `json:"DbServerList,omitempty" xml:"DbServerList,omitempty" type:"Repeated"`
}

func (s GetSessionQueryConditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionQueryConditionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionQueryConditionResponseBody) SetRequestId(v string) *GetSessionQueryConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionQueryConditionResponseBody) SetDbUserList(v []*string) *GetSessionQueryConditionResponseBody {
	s.DbUserList = v
	return s
}

func (s *GetSessionQueryConditionResponseBody) SetClientIpList(v []*string) *GetSessionQueryConditionResponseBody {
	s.ClientIpList = v
	return s
}

func (s *GetSessionQueryConditionResponseBody) SetClientProgramList(v []*string) *GetSessionQueryConditionResponseBody {
	s.ClientProgramList = v
	return s
}

func (s *GetSessionQueryConditionResponseBody) SetDbServerList(v []*string) *GetSessionQueryConditionResponseBody {
	s.DbServerList = v
	return s
}

type GetSessionQueryConditionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSessionQueryConditionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSessionQueryConditionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionQueryConditionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionQueryConditionResponse) SetHeaders(v map[string]*string) *GetSessionQueryConditionResponse {
	s.Headers = v
	return s
}

func (s *GetSessionQueryConditionResponse) SetBody(v *GetSessionQueryConditionResponseBody) *GetSessionQueryConditionResponse {
	s.Body = v
	return s
}

type GetSqlTemplateListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SqlId      *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetSqlTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSqlTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetSqlTemplateListRequest) SetRegionId(v string) *GetSqlTemplateListRequest {
	s.RegionId = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetInstanceId(v string) *GetSqlTemplateListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetDbId(v int32) *GetSqlTemplateListRequest {
	s.DbId = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetBeginDate(v string) *GetSqlTemplateListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetEndDate(v string) *GetSqlTemplateListRequest {
	s.EndDate = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetPageNumber(v int32) *GetSqlTemplateListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetPageSize(v int32) *GetSqlTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetSqlId(v string) *GetSqlTemplateListRequest {
	s.SqlId = &v
	return s
}

func (s *GetSqlTemplateListRequest) SetTemplateId(v string) *GetSqlTemplateListRequest {
	s.TemplateId = &v
	return s
}

type GetSqlTemplateListResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Results    []*GetSqlTemplateListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s GetSqlTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlTemplateListResponseBody) SetRequestId(v string) *GetSqlTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlTemplateListResponseBody) SetPageNumber(v int32) *GetSqlTemplateListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetSqlTemplateListResponseBody) SetPageSize(v int32) *GetSqlTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSqlTemplateListResponseBody) SetTotalCount(v int64) *GetSqlTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSqlTemplateListResponseBody) SetResults(v []*GetSqlTemplateListResponseBodyResults) *GetSqlTemplateListResponseBody {
	s.Results = v
	return s
}

type GetSqlTemplateListResponseBodyResults struct {
	TemplateContent *string                                          `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	LastCaptureTime *string                                          `json:"LastCaptureTime,omitempty" xml:"LastCaptureTime,omitempty"`
	CaptureCount    *int64                                           `json:"CaptureCount,omitempty" xml:"CaptureCount,omitempty"`
	TemplateId      *string                                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	SqlType         *int32                                           `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	RuleList        []*GetSqlTemplateListResponseBodyResultsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s GetSqlTemplateListResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetSqlTemplateListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetSqlTemplateListResponseBodyResults) SetTemplateContent(v string) *GetSqlTemplateListResponseBodyResults {
	s.TemplateContent = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResults) SetLastCaptureTime(v string) *GetSqlTemplateListResponseBodyResults {
	s.LastCaptureTime = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResults) SetCaptureCount(v int64) *GetSqlTemplateListResponseBodyResults {
	s.CaptureCount = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResults) SetTemplateId(v string) *GetSqlTemplateListResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResults) SetSqlType(v int32) *GetSqlTemplateListResponseBodyResults {
	s.SqlType = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResults) SetRuleList(v []*GetSqlTemplateListResponseBodyResultsRuleList) *GetSqlTemplateListResponseBodyResults {
	s.RuleList = v
	return s
}

type GetSqlTemplateListResponseBodyResultsRuleList struct {
	RiskLevel *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RuleState *int32  `json:"RuleState,omitempty" xml:"RuleState,omitempty"`
	DbId      *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	RuleName  *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetSqlTemplateListResponseBodyResultsRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetSqlTemplateListResponseBodyResultsRuleList) GoString() string {
	return s.String()
}

func (s *GetSqlTemplateListResponseBodyResultsRuleList) SetRiskLevel(v int32) *GetSqlTemplateListResponseBodyResultsRuleList {
	s.RiskLevel = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResultsRuleList) SetRuleState(v int32) *GetSqlTemplateListResponseBodyResultsRuleList {
	s.RuleState = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResultsRuleList) SetDbId(v int32) *GetSqlTemplateListResponseBodyResultsRuleList {
	s.DbId = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResultsRuleList) SetRuleName(v string) *GetSqlTemplateListResponseBodyResultsRuleList {
	s.RuleName = &v
	return s
}

func (s *GetSqlTemplateListResponseBodyResultsRuleList) SetRuleId(v string) *GetSqlTemplateListResponseBodyResultsRuleList {
	s.RuleId = &v
	return s
}

type GetSqlTemplateListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSqlTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSqlTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetSqlTemplateListResponse) SetHeaders(v map[string]*string) *GetSqlTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetSqlTemplateListResponse) SetBody(v *GetSqlTemplateListResponseBody) *GetSqlTemplateListResponse {
	s.Body = v
	return s
}

type GetTopSqlTemplateListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderType  *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s GetTopSqlTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopSqlTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetTopSqlTemplateListRequest) SetRegionId(v string) *GetTopSqlTemplateListRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopSqlTemplateListRequest) SetInstanceId(v string) *GetTopSqlTemplateListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopSqlTemplateListRequest) SetDbId(v int32) *GetTopSqlTemplateListRequest {
	s.DbId = &v
	return s
}

func (s *GetTopSqlTemplateListRequest) SetBeginDate(v string) *GetTopSqlTemplateListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetTopSqlTemplateListRequest) SetEndDate(v string) *GetTopSqlTemplateListRequest {
	s.EndDate = &v
	return s
}

func (s *GetTopSqlTemplateListRequest) SetOrderType(v int32) *GetTopSqlTemplateListRequest {
	s.OrderType = &v
	return s
}

type GetTopSqlTemplateListResponseBody struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateList []*GetTopSqlTemplateListResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
}

func (s GetTopSqlTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopSqlTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopSqlTemplateListResponseBody) SetRequestId(v string) *GetTopSqlTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBody) SetTemplateList(v []*GetTopSqlTemplateListResponseBodyTemplateList) *GetTopSqlTemplateListResponseBody {
	s.TemplateList = v
	return s
}

type GetTopSqlTemplateListResponseBodyTemplateList struct {
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	ExecCostUS      *string `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	AffectRows      *string `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	ExecCostUSMean  *string `json:"ExecCostUSMean,omitempty" xml:"ExecCostUSMean,omitempty"`
	LastCaptureTime *string `json:"LastCaptureTime,omitempty" xml:"LastCaptureTime,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CaptureCount    *string `json:"CaptureCount,omitempty" xml:"CaptureCount,omitempty"`
}

func (s GetTopSqlTemplateListResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s GetTopSqlTemplateListResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetTemplateContent(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.TemplateContent = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetExecCostUS(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.ExecCostUS = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetAffectRows(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.AffectRows = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetExecCostUSMean(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.ExecCostUSMean = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetLastCaptureTime(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.LastCaptureTime = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetTemplateId(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *GetTopSqlTemplateListResponseBodyTemplateList) SetCaptureCount(v string) *GetTopSqlTemplateListResponseBodyTemplateList {
	s.CaptureCount = &v
	return s
}

type GetTopSqlTemplateListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTopSqlTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTopSqlTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopSqlTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetTopSqlTemplateListResponse) SetHeaders(v map[string]*string) *GetTopSqlTemplateListResponse {
	s.Headers = v
	return s
}

func (s *GetTopSqlTemplateListResponse) SetBody(v *GetTopSqlTemplateListResponseBody) *GetTopSqlTemplateListResponse {
	s.Body = v
	return s
}

type GradeProtectionReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GradeProtectionReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GradeProtectionReportRequest) GoString() string {
	return s.String()
}

func (s *GradeProtectionReportRequest) SetRegionId(v string) *GradeProtectionReportRequest {
	s.RegionId = &v
	return s
}

func (s *GradeProtectionReportRequest) SetInstanceId(v string) *GradeProtectionReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GradeProtectionReportRequest) SetDbId(v int32) *GradeProtectionReportRequest {
	s.DbId = &v
	return s
}

func (s *GradeProtectionReportRequest) SetBeginDate(v string) *GradeProtectionReportRequest {
	s.BeginDate = &v
	return s
}

func (s *GradeProtectionReportRequest) SetEndDate(v string) *GradeProtectionReportRequest {
	s.EndDate = &v
	return s
}

type GradeProtectionReportResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GradeProtectionReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GradeProtectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *GradeProtectionReportResponseBody) SetServerData(v string) *GradeProtectionReportResponseBody {
	s.ServerData = &v
	return s
}

func (s *GradeProtectionReportResponseBody) SetRequestId(v string) *GradeProtectionReportResponseBody {
	s.RequestId = &v
	return s
}

type GradeProtectionReportResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GradeProtectionReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GradeProtectionReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GradeProtectionReportResponse) GoString() string {
	return s.String()
}

func (s *GradeProtectionReportResponse) SetHeaders(v map[string]*string) *GradeProtectionReportResponse {
	s.Headers = v
	return s
}

func (s *GradeProtectionReportResponse) SetBody(v *GradeProtectionReportResponseBody) *GradeProtectionReportResponse {
	s.Body = v
	return s
}

type ImportDataSourceRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DataJson   *string `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
}

func (s ImportDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ImportDataSourceRequest) SetRegionId(v string) *ImportDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *ImportDataSourceRequest) SetInstanceId(v string) *ImportDataSourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportDataSourceRequest) SetDataJson(v string) *ImportDataSourceRequest {
	s.DataJson = &v
	return s
}

type ImportDataSourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDataSourceResponseBody) SetRequestId(v string) *ImportDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type ImportDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ImportDataSourceResponse) SetHeaders(v map[string]*string) *ImportDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ImportDataSourceResponse) SetBody(v *ImportDataSourceResponseBody) *ImportDataSourceResponse {
	s.Body = v
	return s
}

type IntegratedReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s IntegratedReportRequest) String() string {
	return tea.Prettify(s)
}

func (s IntegratedReportRequest) GoString() string {
	return s.String()
}

func (s *IntegratedReportRequest) SetRegionId(v string) *IntegratedReportRequest {
	s.RegionId = &v
	return s
}

func (s *IntegratedReportRequest) SetInstanceId(v string) *IntegratedReportRequest {
	s.InstanceId = &v
	return s
}

func (s *IntegratedReportRequest) SetDbId(v int32) *IntegratedReportRequest {
	s.DbId = &v
	return s
}

func (s *IntegratedReportRequest) SetBeginDate(v string) *IntegratedReportRequest {
	s.BeginDate = &v
	return s
}

func (s *IntegratedReportRequest) SetEndDate(v string) *IntegratedReportRequest {
	s.EndDate = &v
	return s
}

type IntegratedReportResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IntegratedReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IntegratedReportResponseBody) GoString() string {
	return s.String()
}

func (s *IntegratedReportResponseBody) SetServerData(v string) *IntegratedReportResponseBody {
	s.ServerData = &v
	return s
}

func (s *IntegratedReportResponseBody) SetRequestId(v string) *IntegratedReportResponseBody {
	s.RequestId = &v
	return s
}

type IntegratedReportResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IntegratedReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IntegratedReportResponse) String() string {
	return tea.Prettify(s)
}

func (s IntegratedReportResponse) GoString() string {
	return s.String()
}

func (s *IntegratedReportResponse) SetHeaders(v map[string]*string) *IntegratedReportResponse {
	s.Headers = v
	return s
}

func (s *IntegratedReportResponse) SetBody(v *IntegratedReportResponseBody) *IntegratedReportResponse {
	s.Body = v
	return s
}

type ListAssociatedRulesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType   *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	RuleDefn   *string `json:"RuleDefn,omitempty" xml:"RuleDefn,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
}

func (s ListAssociatedRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAssociatedRulesRequest) SetRegionId(v string) *ListAssociatedRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAssociatedRulesRequest) SetInstanceId(v string) *ListAssociatedRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAssociatedRulesRequest) SetRuleName(v string) *ListAssociatedRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListAssociatedRulesRequest) SetRuleType(v string) *ListAssociatedRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListAssociatedRulesRequest) SetRuleDefn(v string) *ListAssociatedRulesRequest {
	s.RuleDefn = &v
	return s
}

func (s *ListAssociatedRulesRequest) SetDbId(v int32) *ListAssociatedRulesRequest {
	s.DbId = &v
	return s
}

type ListAssociatedRulesResponseBody struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *ListAssociatedRulesResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s ListAssociatedRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssociatedRulesResponseBody) SetRequestId(v string) *ListAssociatedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssociatedRulesResponseBody) SetServerData(v *ListAssociatedRulesResponseBodyServerData) *ListAssociatedRulesResponseBody {
	s.ServerData = v
	return s
}

type ListAssociatedRulesResponseBodyServerData struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
}

func (s ListAssociatedRulesResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedRulesResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *ListAssociatedRulesResponseBodyServerData) SetJsonResult(v string) *ListAssociatedRulesResponseBodyServerData {
	s.JsonResult = &v
	return s
}

type ListAssociatedRulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAssociatedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAssociatedRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssociatedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAssociatedRulesResponse) SetHeaders(v map[string]*string) *ListAssociatedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAssociatedRulesResponse) SetBody(v *ListAssociatedRulesResponseBody) *ListAssociatedRulesResponse {
	s.Body = v
	return s
}

type ListDataSourceAttributeRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbIds      []*string `json:"DbIds,omitempty" xml:"DbIds,omitempty" type:"Repeated"`
}

func (s ListDataSourceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeRequest) SetRegionId(v string) *ListDataSourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourceAttributeRequest) SetInstanceId(v string) *ListDataSourceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataSourceAttributeRequest) SetDbIds(v []*string) *ListDataSourceAttributeRequest {
	s.DbIds = v
	return s
}

type ListDataSourceAttributeResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbList    []*ListDataSourceAttributeResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListDataSourceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeResponseBody) SetRequestId(v string) *ListDataSourceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceAttributeResponseBody) SetDbList(v []*ListDataSourceAttributeResponseBodyDbList) *ListDataSourceAttributeResponseBody {
	s.DbList = v
	return s
}

type ListDataSourceAttributeResponseBodyDbList struct {
	ResultAuditMode    *string `json:"ResultAuditMode,omitempty" xml:"ResultAuditMode,omitempty"`
	DbId               *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	ResultAuditMaxSize *int32  `json:"ResultAuditMaxSize,omitempty" xml:"ResultAuditMaxSize,omitempty"`
	AuditMode          *string `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	ResultAuditMaxLine *int32  `json:"ResultAuditMaxLine,omitempty" xml:"ResultAuditMaxLine,omitempty"`
}

func (s ListDataSourceAttributeResponseBodyDbList) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceAttributeResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetResultAuditMode(v string) *ListDataSourceAttributeResponseBodyDbList {
	s.ResultAuditMode = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetDbId(v int32) *ListDataSourceAttributeResponseBodyDbList {
	s.DbId = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetResultAuditMaxSize(v int32) *ListDataSourceAttributeResponseBodyDbList {
	s.ResultAuditMaxSize = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetAuditMode(v string) *ListDataSourceAttributeResponseBodyDbList {
	s.AuditMode = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetResultAuditMaxLine(v int32) *ListDataSourceAttributeResponseBodyDbList {
	s.ResultAuditMaxLine = &v
	return s
}

type ListDataSourceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeResponse) SetHeaders(v map[string]*string) *ListDataSourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceAttributeResponse) SetBody(v *ListDataSourceAttributeResponseBody) *ListDataSourceAttributeResponse {
	s.Body = v
	return s
}

type ListDataSourcesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) SetRegionId(v string) *ListDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourcesRequest) SetInstanceId(v string) *ListDataSourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataSourcesRequest) SetDbId(v int32) *ListDataSourcesRequest {
	s.DbId = &v
	return s
}

type ListDataSourcesResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DbList    []*ListDataSourcesResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetDbList(v []*ListDataSourcesResponseBodyDbList) *ListDataSourcesResponseBody {
	s.DbList = v
	return s
}

type ListDataSourcesResponseBodyDbList struct {
	DbId          *int32    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	CreateTime    *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbUsername    *string   `json:"DbUsername,omitempty" xml:"DbUsername,omitempty"`
	DbCertificate *string   `json:"DbCertificate,omitempty" xml:"DbCertificate,omitempty"`
	DbInstanceId  *string   `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	AssetType     *int32    `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	DbVersion     *int32    `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	DbName        *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbType        *int32    `json:"DbType,omitempty" xml:"DbType,omitempty"`
	AuditSwitch   *int32    `json:"AuditSwitch,omitempty" xml:"AuditSwitch,omitempty"`
	DbTypeName    *string   `json:"DbTypeName,omitempty" xml:"DbTypeName,omitempty"`
	DbVersionName *string   `json:"DbVersionName,omitempty" xml:"DbVersionName,omitempty"`
	DbAddresses   []*string `json:"DbAddresses,omitempty" xml:"DbAddresses,omitempty" type:"Repeated"`
}

func (s ListDataSourcesResponseBodyDbList) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDbList) SetDbId(v int32) *ListDataSourcesResponseBodyDbList {
	s.DbId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetCreateTime(v string) *ListDataSourcesResponseBodyDbList {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbUsername(v string) *ListDataSourcesResponseBodyDbList {
	s.DbUsername = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbCertificate(v string) *ListDataSourcesResponseBodyDbList {
	s.DbCertificate = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbInstanceId(v string) *ListDataSourcesResponseBodyDbList {
	s.DbInstanceId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetAssetType(v int32) *ListDataSourcesResponseBodyDbList {
	s.AssetType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbVersion(v int32) *ListDataSourcesResponseBodyDbList {
	s.DbVersion = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbName(v string) *ListDataSourcesResponseBodyDbList {
	s.DbName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbType(v int32) *ListDataSourcesResponseBodyDbList {
	s.DbType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetAuditSwitch(v int32) *ListDataSourcesResponseBodyDbList {
	s.AuditSwitch = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbTypeName(v string) *ListDataSourcesResponseBodyDbList {
	s.DbTypeName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbVersionName(v string) *ListDataSourcesResponseBodyDbList {
	s.DbVersionName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDbList) SetDbAddresses(v []*string) *ListDataSourcesResponseBodyDbList {
	s.DbAddresses = v
	return s
}

type ListDataSourcesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListLogAlarmTasksRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListLogAlarmTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogAlarmTasksRequest) GoString() string {
	return s.String()
}

func (s *ListLogAlarmTasksRequest) SetRegionId(v string) *ListLogAlarmTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListLogAlarmTasksRequest) SetInstanceId(v string) *ListLogAlarmTasksRequest {
	s.InstanceId = &v
	return s
}

type ListLogAlarmTasksResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList  []*ListLogAlarmTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s ListLogAlarmTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogAlarmTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogAlarmTasksResponseBody) SetRequestId(v string) *ListLogAlarmTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogAlarmTasksResponseBody) SetTaskList(v []*ListLogAlarmTasksResponseBodyTaskList) *ListLogAlarmTasksResponseBody {
	s.TaskList = v
	return s
}

type ListLogAlarmTasksResponseBodyTaskList struct {
	TaskStatus    *int32    `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskId        *int32    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CreateTime    *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskName      *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ToMailList    []*string `json:"ToMailList,omitempty" xml:"ToMailList,omitempty" type:"Repeated"`
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	DbIdList      []*string `json:"DbIdList,omitempty" xml:"DbIdList,omitempty" type:"Repeated"`
}

func (s ListLogAlarmTasksResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListLogAlarmTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetTaskStatus(v int32) *ListLogAlarmTasksResponseBodyTaskList {
	s.TaskStatus = &v
	return s
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetTaskId(v int32) *ListLogAlarmTasksResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetCreateTime(v string) *ListLogAlarmTasksResponseBodyTaskList {
	s.CreateTime = &v
	return s
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetTaskName(v string) *ListLogAlarmTasksResponseBodyTaskList {
	s.TaskName = &v
	return s
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetToMailList(v []*string) *ListLogAlarmTasksResponseBodyTaskList {
	s.ToMailList = v
	return s
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetRiskLevelList(v []*string) *ListLogAlarmTasksResponseBodyTaskList {
	s.RiskLevelList = v
	return s
}

func (s *ListLogAlarmTasksResponseBodyTaskList) SetDbIdList(v []*string) *ListLogAlarmTasksResponseBodyTaskList {
	s.DbIdList = v
	return s
}

type ListLogAlarmTasksResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogAlarmTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogAlarmTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogAlarmTasksResponse) GoString() string {
	return s.String()
}

func (s *ListLogAlarmTasksResponse) SetHeaders(v map[string]*string) *ListLogAlarmTasksResponse {
	s.Headers = v
	return s
}

func (s *ListLogAlarmTasksResponse) SetBody(v *ListLogAlarmTasksResponseBody) *ListLogAlarmTasksResponse {
	s.Body = v
	return s
}

type ListReportPushTasksRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListReportPushTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReportPushTasksRequest) GoString() string {
	return s.String()
}

func (s *ListReportPushTasksRequest) SetRegionId(v string) *ListReportPushTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListReportPushTasksRequest) SetInstanceId(v string) *ListReportPushTasksRequest {
	s.InstanceId = &v
	return s
}

type ListReportPushTasksResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList  []*ListReportPushTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s ListReportPushTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReportPushTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportPushTasksResponseBody) SetRequestId(v string) *ListReportPushTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReportPushTasksResponseBody) SetTaskList(v []*ListReportPushTasksResponseBodyTaskList) *ListReportPushTasksResponseBody {
	s.TaskList = v
	return s
}

type ListReportPushTasksResponseBodyTaskList struct {
	ScheduleTime *string   `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	JobStatus    *int32    `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	JobId        *int32    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ScheduleType *string   `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	DbList       []*string `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	ReportType   []*string `json:"ReportType,omitempty" xml:"ReportType,omitempty" type:"Repeated"`
	EmailList    []*string `json:"EmailList,omitempty" xml:"EmailList,omitempty" type:"Repeated"`
}

func (s ListReportPushTasksResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListReportPushTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListReportPushTasksResponseBodyTaskList) SetScheduleTime(v string) *ListReportPushTasksResponseBodyTaskList {
	s.ScheduleTime = &v
	return s
}

func (s *ListReportPushTasksResponseBodyTaskList) SetJobStatus(v int32) *ListReportPushTasksResponseBodyTaskList {
	s.JobStatus = &v
	return s
}

func (s *ListReportPushTasksResponseBodyTaskList) SetJobId(v int32) *ListReportPushTasksResponseBodyTaskList {
	s.JobId = &v
	return s
}

func (s *ListReportPushTasksResponseBodyTaskList) SetScheduleType(v string) *ListReportPushTasksResponseBodyTaskList {
	s.ScheduleType = &v
	return s
}

func (s *ListReportPushTasksResponseBodyTaskList) SetDbList(v []*string) *ListReportPushTasksResponseBodyTaskList {
	s.DbList = v
	return s
}

func (s *ListReportPushTasksResponseBodyTaskList) SetReportType(v []*string) *ListReportPushTasksResponseBodyTaskList {
	s.ReportType = v
	return s
}

func (s *ListReportPushTasksResponseBodyTaskList) SetEmailList(v []*string) *ListReportPushTasksResponseBodyTaskList {
	s.EmailList = v
	return s
}

type ListReportPushTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListReportPushTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListReportPushTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReportPushTasksResponse) GoString() string {
	return s.String()
}

func (s *ListReportPushTasksResponse) SetHeaders(v map[string]*string) *ListReportPushTasksResponse {
	s.Headers = v
	return s
}

func (s *ListReportPushTasksResponse) SetBody(v *ListReportPushTasksResponseBody) *ListReportPushTasksResponse {
	s.Body = v
	return s
}

type ListRuleGroupsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListRuleGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRuleGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListRuleGroupsRequest) SetRegionId(v string) *ListRuleGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRuleGroupsRequest) SetInstanceId(v string) *ListRuleGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRuleGroupsRequest) SetRuleName(v string) *ListRuleGroupsRequest {
	s.RuleName = &v
	return s
}

type ListRuleGroupsResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRuleGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRuleGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRuleGroupsResponseBody) SetServerData(v string) *ListRuleGroupsResponseBody {
	s.ServerData = &v
	return s
}

func (s *ListRuleGroupsResponseBody) SetRequestId(v string) *ListRuleGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListRuleGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRuleGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRuleGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRuleGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListRuleGroupsResponse) SetHeaders(v map[string]*string) *ListRuleGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListRuleGroupsResponse) SetBody(v *ListRuleGroupsResponseBody) *ListRuleGroupsResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId           *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	ForegroundType *string `json:"ForegroundType,omitempty" xml:"ForegroundType,omitempty"`
	RuleName       *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType       *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	RuleGroupId    *string `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	RiskLevel      *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RuleQuote      *int32  `json:"RuleQuote,omitempty" xml:"RuleQuote,omitempty"`
	RuleState      *int32  `json:"RuleState,omitempty" xml:"RuleState,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetRegionId(v string) *ListRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListRulesRequest) SetInstanceId(v string) *ListRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRulesRequest) SetDbId(v string) *ListRulesRequest {
	s.DbId = &v
	return s
}

func (s *ListRulesRequest) SetForegroundType(v string) *ListRulesRequest {
	s.ForegroundType = &v
	return s
}

func (s *ListRulesRequest) SetRuleName(v string) *ListRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListRulesRequest) SetRuleType(v string) *ListRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListRulesRequest) SetRuleGroupId(v string) *ListRulesRequest {
	s.RuleGroupId = &v
	return s
}

func (s *ListRulesRequest) SetRiskLevel(v string) *ListRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListRulesRequest) SetRuleQuote(v int32) *ListRulesRequest {
	s.RuleQuote = &v
	return s
}

func (s *ListRulesRequest) SetRuleState(v int32) *ListRulesRequest {
	s.RuleState = &v
	return s
}

type ListRulesResponseBody struct {
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData *ListRulesResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Struct"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetServerData(v *ListRulesResponseBodyServerData) *ListRulesResponseBody {
	s.ServerData = v
	return s
}

type ListRulesResponseBodyServerData struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
}

func (s ListRulesResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyServerData) SetJsonResult(v string) *ListRulesResponseBodyServerData {
	s.JsonResult = &v
	return s
}

type ListRulesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListSqlTypeKeysForRuleRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSqlTypeKeysForRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSqlTypeKeysForRuleRequest) GoString() string {
	return s.String()
}

func (s *ListSqlTypeKeysForRuleRequest) SetRegionId(v string) *ListSqlTypeKeysForRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListSqlTypeKeysForRuleRequest) SetInstanceId(v string) *ListSqlTypeKeysForRuleRequest {
	s.InstanceId = &v
	return s
}

type ListSqlTypeKeysForRuleResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSqlTypeKeysForRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSqlTypeKeysForRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSqlTypeKeysForRuleResponseBody) SetServerData(v string) *ListSqlTypeKeysForRuleResponseBody {
	s.ServerData = &v
	return s
}

func (s *ListSqlTypeKeysForRuleResponseBody) SetRequestId(v string) *ListSqlTypeKeysForRuleResponseBody {
	s.RequestId = &v
	return s
}

type ListSqlTypeKeysForRuleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSqlTypeKeysForRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSqlTypeKeysForRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSqlTypeKeysForRuleResponse) GoString() string {
	return s.String()
}

func (s *ListSqlTypeKeysForRuleResponse) SetHeaders(v map[string]*string) *ListSqlTypeKeysForRuleResponse {
	s.Headers = v
	return s
}

func (s *ListSqlTypeKeysForRuleResponse) SetBody(v *ListSqlTypeKeysForRuleResponseBody) *ListSqlTypeKeysForRuleResponse {
	s.Body = v
	return s
}

type ListSqlTypesForRuleRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TypeId     *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
	Sqltype1   *string `json:"Sqltype1,omitempty" xml:"Sqltype1,omitempty"`
	KeyWorld   *string `json:"KeyWorld,omitempty" xml:"KeyWorld,omitempty"`
}

func (s ListSqlTypesForRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSqlTypesForRuleRequest) GoString() string {
	return s.String()
}

func (s *ListSqlTypesForRuleRequest) SetRegionId(v string) *ListSqlTypesForRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListSqlTypesForRuleRequest) SetInstanceId(v string) *ListSqlTypesForRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSqlTypesForRuleRequest) SetTypeId(v string) *ListSqlTypesForRuleRequest {
	s.TypeId = &v
	return s
}

func (s *ListSqlTypesForRuleRequest) SetSqltype1(v string) *ListSqlTypesForRuleRequest {
	s.Sqltype1 = &v
	return s
}

func (s *ListSqlTypesForRuleRequest) SetKeyWorld(v string) *ListSqlTypesForRuleRequest {
	s.KeyWorld = &v
	return s
}

type ListSqlTypesForRuleResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSqlTypesForRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSqlTypesForRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSqlTypesForRuleResponseBody) SetServerData(v string) *ListSqlTypesForRuleResponseBody {
	s.ServerData = &v
	return s
}

func (s *ListSqlTypesForRuleResponseBody) SetRequestId(v string) *ListSqlTypesForRuleResponseBody {
	s.RequestId = &v
	return s
}

type ListSqlTypesForRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSqlTypesForRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSqlTypesForRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSqlTypesForRuleResponse) GoString() string {
	return s.String()
}

func (s *ListSqlTypesForRuleResponse) SetHeaders(v map[string]*string) *ListSqlTypesForRuleResponse {
	s.Headers = v
	return s
}

func (s *ListSqlTypesForRuleResponse) SetBody(v *ListSqlTypesForRuleResponseBody) *ListSqlTypesForRuleResponse {
	s.Body = v
	return s
}

type ListSupportDbTypesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSupportDbTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSupportDbTypesRequest) GoString() string {
	return s.String()
}

func (s *ListSupportDbTypesRequest) SetRegionId(v string) *ListSupportDbTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupportDbTypesRequest) SetInstanceId(v string) *ListSupportDbTypesRequest {
	s.InstanceId = &v
	return s
}

type ListSupportDbTypesResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TypeList  []*ListSupportDbTypesResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s ListSupportDbTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSupportDbTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportDbTypesResponseBody) SetRequestId(v string) *ListSupportDbTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportDbTypesResponseBody) SetTypeList(v []*ListSupportDbTypesResponseBodyTypeList) *ListSupportDbTypesResponseBody {
	s.TypeList = v
	return s
}

type ListSupportDbTypesResponseBodyTypeList struct {
	DbType     *int32                                              `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DbTypeName *string                                             `json:"DbTypeName,omitempty" xml:"DbTypeName,omitempty"`
	DbVersions []*ListSupportDbTypesResponseBodyTypeListDbVersions `json:"DbVersions,omitempty" xml:"DbVersions,omitempty" type:"Repeated"`
}

func (s ListSupportDbTypesResponseBodyTypeList) String() string {
	return tea.Prettify(s)
}

func (s ListSupportDbTypesResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *ListSupportDbTypesResponseBodyTypeList) SetDbType(v int32) *ListSupportDbTypesResponseBodyTypeList {
	s.DbType = &v
	return s
}

func (s *ListSupportDbTypesResponseBodyTypeList) SetDbTypeName(v string) *ListSupportDbTypesResponseBodyTypeList {
	s.DbTypeName = &v
	return s
}

func (s *ListSupportDbTypesResponseBodyTypeList) SetDbVersions(v []*ListSupportDbTypesResponseBodyTypeListDbVersions) *ListSupportDbTypesResponseBodyTypeList {
	s.DbVersions = v
	return s
}

type ListSupportDbTypesResponseBodyTypeListDbVersions struct {
	DbVersionName *string `json:"DbVersionName,omitempty" xml:"DbVersionName,omitempty"`
	DbVersion     *int32  `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
}

func (s ListSupportDbTypesResponseBodyTypeListDbVersions) String() string {
	return tea.Prettify(s)
}

func (s ListSupportDbTypesResponseBodyTypeListDbVersions) GoString() string {
	return s.String()
}

func (s *ListSupportDbTypesResponseBodyTypeListDbVersions) SetDbVersionName(v string) *ListSupportDbTypesResponseBodyTypeListDbVersions {
	s.DbVersionName = &v
	return s
}

func (s *ListSupportDbTypesResponseBodyTypeListDbVersions) SetDbVersion(v int32) *ListSupportDbTypesResponseBodyTypeListDbVersions {
	s.DbVersion = &v
	return s
}

type ListSupportDbTypesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSupportDbTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSupportDbTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSupportDbTypesResponse) GoString() string {
	return s.String()
}

func (s *ListSupportDbTypesResponse) SetHeaders(v map[string]*string) *ListSupportDbTypesResponse {
	s.Headers = v
	return s
}

func (s *ListSupportDbTypesResponse) SetBody(v *ListSupportDbTypesResponseBody) *ListSupportDbTypesResponse {
	s.Body = v
	return s
}

type ListSystemAlarmsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AlarmType  *int32  `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s ListSystemAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsRequest) SetRegionId(v string) *ListSystemAlarmsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetInstanceId(v string) *ListSystemAlarmsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetAlarmType(v int32) *ListSystemAlarmsRequest {
	s.AlarmType = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetBeginDate(v string) *ListSystemAlarmsRequest {
	s.BeginDate = &v
	return s
}

func (s *ListSystemAlarmsRequest) SetEndDate(v string) *ListSystemAlarmsRequest {
	s.EndDate = &v
	return s
}

type ListSystemAlarmsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Alarms    []*ListSystemAlarmsResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
}

func (s ListSystemAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsResponseBody) SetRequestId(v string) *ListSystemAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemAlarmsResponseBody) SetAlarms(v []*ListSystemAlarmsResponseBodyAlarms) *ListSystemAlarmsResponseBody {
	s.Alarms = v
	return s
}

type ListSystemAlarmsResponseBodyAlarms struct {
	ReadMark    *int32  `json:"ReadMark,omitempty" xml:"ReadMark,omitempty"`
	AlarmDetail *string `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty"`
	AlarmType   *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	AlarmId     *int32  `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListSystemAlarmsResponseBodyAlarms) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmsResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetReadMark(v int32) *ListSystemAlarmsResponseBodyAlarms {
	s.ReadMark = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetAlarmDetail(v string) *ListSystemAlarmsResponseBodyAlarms {
	s.AlarmDetail = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetAlarmType(v string) *ListSystemAlarmsResponseBodyAlarms {
	s.AlarmType = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetAlarmId(v int32) *ListSystemAlarmsResponseBodyAlarms {
	s.AlarmId = &v
	return s
}

func (s *ListSystemAlarmsResponseBodyAlarms) SetCreateTime(v string) *ListSystemAlarmsResponseBodyAlarms {
	s.CreateTime = &v
	return s
}

type ListSystemAlarmsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSystemAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsResponse) SetHeaders(v map[string]*string) *ListSystemAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ListSystemAlarmsResponse) SetBody(v *ListSystemAlarmsResponseBody) *ListSystemAlarmsResponse {
	s.Body = v
	return s
}

type ListSystemAlarmTasksRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSystemAlarmTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmTasksRequest) SetRegionId(v string) *ListSystemAlarmTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListSystemAlarmTasksRequest) SetInstanceId(v string) *ListSystemAlarmTasksRequest {
	s.InstanceId = &v
	return s
}

type ListSystemAlarmTasksResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList  []*ListSystemAlarmTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s ListSystemAlarmTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmTasksResponseBody) SetRequestId(v string) *ListSystemAlarmTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemAlarmTasksResponseBody) SetTaskList(v []*ListSystemAlarmTasksResponseBodyTaskList) *ListSystemAlarmTasksResponseBody {
	s.TaskList = v
	return s
}

type ListSystemAlarmTasksResponseBodyTaskList struct {
	TaskStatus *int32    `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskId     *int32    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CreateTime *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskName   *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ToMailList []*string `json:"ToMailList,omitempty" xml:"ToMailList,omitempty" type:"Repeated"`
}

func (s ListSystemAlarmTasksResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmTasksResponseBodyTaskList) SetTaskStatus(v int32) *ListSystemAlarmTasksResponseBodyTaskList {
	s.TaskStatus = &v
	return s
}

func (s *ListSystemAlarmTasksResponseBodyTaskList) SetTaskId(v int32) *ListSystemAlarmTasksResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListSystemAlarmTasksResponseBodyTaskList) SetCreateTime(v string) *ListSystemAlarmTasksResponseBodyTaskList {
	s.CreateTime = &v
	return s
}

func (s *ListSystemAlarmTasksResponseBodyTaskList) SetTaskName(v string) *ListSystemAlarmTasksResponseBodyTaskList {
	s.TaskName = &v
	return s
}

func (s *ListSystemAlarmTasksResponseBodyTaskList) SetToMailList(v []*string) *ListSystemAlarmTasksResponseBodyTaskList {
	s.ToMailList = v
	return s
}

type ListSystemAlarmTasksResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSystemAlarmTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemAlarmTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemAlarmTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmTasksResponse) SetHeaders(v map[string]*string) *ListSystemAlarmTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSystemAlarmTasksResponse) SetBody(v *ListSystemAlarmTasksResponseBody) *ListSystemAlarmTasksResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

type ListTagKeysResponseBody struct {
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TagKeys     []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
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

type ListTemplatesForSqlRuleRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SqlType1       *string `json:"SqlType1,omitempty" xml:"SqlType1,omitempty"`
	ChoseCondition *string `json:"ChoseCondition,omitempty" xml:"ChoseCondition,omitempty"`
	DbId           *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	RuleId         *int32  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTemplatesForSqlRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesForSqlRuleRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesForSqlRuleRequest) SetRegionId(v string) *ListTemplatesForSqlRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetInstanceId(v string) *ListTemplatesForSqlRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetSqlType1(v string) *ListTemplatesForSqlRuleRequest {
	s.SqlType1 = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetChoseCondition(v string) *ListTemplatesForSqlRuleRequest {
	s.ChoseCondition = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetDbId(v int32) *ListTemplatesForSqlRuleRequest {
	s.DbId = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetRuleId(v int32) *ListTemplatesForSqlRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetPageNumber(v int32) *ListTemplatesForSqlRuleRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesForSqlRuleRequest) SetPageSize(v int32) *ListTemplatesForSqlRuleRequest {
	s.PageSize = &v
	return s
}

type ListTemplatesForSqlRuleResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerData []*ListTemplatesForSqlRuleResponseBodyServerData `json:"ServerData,omitempty" xml:"ServerData,omitempty" type:"Repeated"`
}

func (s ListTemplatesForSqlRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesForSqlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesForSqlRuleResponseBody) SetRequestId(v string) *ListTemplatesForSqlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesForSqlRuleResponseBody) SetServerData(v []*ListTemplatesForSqlRuleResponseBodyServerData) *ListTemplatesForSqlRuleResponseBody {
	s.ServerData = v
	return s
}

type ListTemplatesForSqlRuleResponseBodyServerData struct {
	CountTimely  *string `json:"CountTimely,omitempty" xml:"CountTimely,omitempty"`
	BlackOrWhite *int32  `json:"BlackOrWhite,omitempty" xml:"BlackOrWhite,omitempty"`
	SqlText      *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	Sqltype1     *string `json:"Sqltype1,omitempty" xml:"Sqltype1,omitempty"`
	SqlId        *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s ListTemplatesForSqlRuleResponseBodyServerData) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesForSqlRuleResponseBodyServerData) GoString() string {
	return s.String()
}

func (s *ListTemplatesForSqlRuleResponseBodyServerData) SetCountTimely(v string) *ListTemplatesForSqlRuleResponseBodyServerData {
	s.CountTimely = &v
	return s
}

func (s *ListTemplatesForSqlRuleResponseBodyServerData) SetBlackOrWhite(v int32) *ListTemplatesForSqlRuleResponseBodyServerData {
	s.BlackOrWhite = &v
	return s
}

func (s *ListTemplatesForSqlRuleResponseBodyServerData) SetSqlText(v string) *ListTemplatesForSqlRuleResponseBodyServerData {
	s.SqlText = &v
	return s
}

func (s *ListTemplatesForSqlRuleResponseBodyServerData) SetSqltype1(v string) *ListTemplatesForSqlRuleResponseBodyServerData {
	s.Sqltype1 = &v
	return s
}

func (s *ListTemplatesForSqlRuleResponseBodyServerData) SetSqlId(v string) *ListTemplatesForSqlRuleResponseBodyServerData {
	s.SqlId = &v
	return s
}

type ListTemplatesForSqlRuleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesForSqlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesForSqlRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesForSqlRuleResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesForSqlRuleResponse) SetHeaders(v map[string]*string) *ListTemplatesForSqlRuleResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesForSqlRuleResponse) SetBody(v *ListTemplatesForSqlRuleResponseBody) *ListTemplatesForSqlRuleResponse {
	s.Body = v
	return s
}

type ListUsedSqlTypesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType   *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListUsedSqlTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsedSqlTypesRequest) GoString() string {
	return s.String()
}

func (s *ListUsedSqlTypesRequest) SetRegionId(v string) *ListUsedSqlTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsedSqlTypesRequest) SetInstanceId(v string) *ListUsedSqlTypesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsedSqlTypesRequest) SetRuleId(v string) *ListUsedSqlTypesRequest {
	s.RuleId = &v
	return s
}

func (s *ListUsedSqlTypesRequest) SetRuleType(v string) *ListUsedSqlTypesRequest {
	s.RuleType = &v
	return s
}

type ListUsedSqlTypesResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUsedSqlTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsedSqlTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsedSqlTypesResponseBody) SetServerData(v string) *ListUsedSqlTypesResponseBody {
	s.ServerData = &v
	return s
}

func (s *ListUsedSqlTypesResponseBody) SetRequestId(v string) *ListUsedSqlTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListUsedSqlTypesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsedSqlTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsedSqlTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsedSqlTypesResponse) GoString() string {
	return s.String()
}

func (s *ListUsedSqlTypesResponse) SetHeaders(v map[string]*string) *ListUsedSqlTypesResponse {
	s.Headers = v
	return s
}

func (s *ListUsedSqlTypesResponse) SetBody(v *ListUsedSqlTypesResponseBody) *ListUsedSqlTypesResponse {
	s.Body = v
	return s
}

type ModifyBaseTemplateStateRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateState *int32    `json:"TemplateState,omitempty" xml:"TemplateState,omitempty"`
	TemplateIds   []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s ModifyBaseTemplateStateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBaseTemplateStateRequest) GoString() string {
	return s.String()
}

func (s *ModifyBaseTemplateStateRequest) SetRegionId(v string) *ModifyBaseTemplateStateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBaseTemplateStateRequest) SetInstanceId(v string) *ModifyBaseTemplateStateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBaseTemplateStateRequest) SetTemplateState(v int32) *ModifyBaseTemplateStateRequest {
	s.TemplateState = &v
	return s
}

func (s *ModifyBaseTemplateStateRequest) SetTemplateIds(v []*string) *ModifyBaseTemplateStateRequest {
	s.TemplateIds = v
	return s
}

type ModifyBaseTemplateStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBaseTemplateStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBaseTemplateStateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBaseTemplateStateResponseBody) SetRequestId(v string) *ModifyBaseTemplateStateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBaseTemplateStateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBaseTemplateStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBaseTemplateStateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBaseTemplateStateResponse) GoString() string {
	return s.String()
}

func (s *ModifyBaseTemplateStateResponse) SetHeaders(v map[string]*string) *ModifyBaseTemplateStateResponse {
	s.Headers = v
	return s
}

func (s *ModifyBaseTemplateStateResponse) SetBody(v *ModifyBaseTemplateStateResponseBody) *ModifyBaseTemplateStateResponse {
	s.Body = v
	return s
}

type ModifyCustomNameRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CustomName    *string `json:"CustomName,omitempty" xml:"CustomName,omitempty"`
}

func (s ModifyCustomNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCustomNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomNameRequest) SetSourceIp(v string) *ModifyCustomNameRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyCustomNameRequest) SetCommodityCode(v string) *ModifyCustomNameRequest {
	s.CommodityCode = &v
	return s
}

func (s *ModifyCustomNameRequest) SetInstanceId(v string) *ModifyCustomNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCustomNameRequest) SetCustomName(v string) *ModifyCustomNameRequest {
	s.CustomName = &v
	return s
}

type ModifyCustomNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCustomNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomNameResponseBody) SetRequestId(v string) *ModifyCustomNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCustomNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCustomNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCustomNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCustomNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomNameResponse) SetHeaders(v map[string]*string) *ModifyCustomNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomNameResponse) SetBody(v *ModifyCustomNameResponseBody) *ModifyCustomNameResponse {
	s.Body = v
	return s
}

type ModifyDataSourceRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId          *int32    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName        *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbVersion     *int32    `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	DbCertificate *string   `json:"DbCertificate,omitempty" xml:"DbCertificate,omitempty"`
	DbUsername    *string   `json:"DbUsername,omitempty" xml:"DbUsername,omitempty"`
	DbPassword    *string   `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	DbInstanceId  *string   `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbAddresses   []*string `json:"DbAddresses,omitempty" xml:"DbAddresses,omitempty" type:"Repeated"`
}

func (s ModifyDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceRequest) SetRegionId(v string) *ModifyDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDataSourceRequest) SetInstanceId(v string) *ModifyDataSourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbId(v int32) *ModifyDataSourceRequest {
	s.DbId = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbName(v string) *ModifyDataSourceRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbVersion(v int32) *ModifyDataSourceRequest {
	s.DbVersion = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbCertificate(v string) *ModifyDataSourceRequest {
	s.DbCertificate = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbUsername(v string) *ModifyDataSourceRequest {
	s.DbUsername = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbPassword(v string) *ModifyDataSourceRequest {
	s.DbPassword = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbInstanceId(v string) *ModifyDataSourceRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDbAddresses(v []*string) *ModifyDataSourceRequest {
	s.DbAddresses = v
	return s
}

type ModifyDataSourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBody) SetRequestId(v string) *ModifyDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponse) SetHeaders(v map[string]*string) *ModifyDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceResponse) SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse {
	s.Body = v
	return s
}

type ModifyDataSourceAttributeRequest struct {
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId         *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AuditMode          *string   `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	ResultAuditMode    *string   `json:"ResultAuditMode,omitempty" xml:"ResultAuditMode,omitempty"`
	ResultAuditMaxLine *int32    `json:"ResultAuditMaxLine,omitempty" xml:"ResultAuditMaxLine,omitempty"`
	ResultAuditMaxSize *int32    `json:"ResultAuditMaxSize,omitempty" xml:"ResultAuditMaxSize,omitempty"`
	DbIds              []*string `json:"DbIds,omitempty" xml:"DbIds,omitempty" type:"Repeated"`
}

func (s ModifyDataSourceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceAttributeRequest) SetRegionId(v string) *ModifyDataSourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDataSourceAttributeRequest) SetInstanceId(v string) *ModifyDataSourceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDataSourceAttributeRequest) SetAuditMode(v string) *ModifyDataSourceAttributeRequest {
	s.AuditMode = &v
	return s
}

func (s *ModifyDataSourceAttributeRequest) SetResultAuditMode(v string) *ModifyDataSourceAttributeRequest {
	s.ResultAuditMode = &v
	return s
}

func (s *ModifyDataSourceAttributeRequest) SetResultAuditMaxLine(v int32) *ModifyDataSourceAttributeRequest {
	s.ResultAuditMaxLine = &v
	return s
}

func (s *ModifyDataSourceAttributeRequest) SetResultAuditMaxSize(v int32) *ModifyDataSourceAttributeRequest {
	s.ResultAuditMaxSize = &v
	return s
}

func (s *ModifyDataSourceAttributeRequest) SetDbIds(v []*string) *ModifyDataSourceAttributeRequest {
	s.DbIds = v
	return s
}

type ModifyDataSourceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataSourceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceAttributeResponseBody) SetRequestId(v string) *ModifyDataSourceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataSourceAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDataSourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataSourceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceAttributeResponse) SetHeaders(v map[string]*string) *ModifyDataSourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceAttributeResponse) SetBody(v *ModifyDataSourceAttributeResponseBody) *ModifyDataSourceAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRegionId(v string) *ModifyInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyLogAlarmTaskRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskName      *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskId        *int32    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ToMailList    []*string `json:"ToMailList,omitempty" xml:"ToMailList,omitempty" type:"Repeated"`
	DbIdList      []*string `json:"DbIdList,omitempty" xml:"DbIdList,omitempty" type:"Repeated"`
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
}

func (s ModifyLogAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogAlarmTaskRequest) SetRegionId(v string) *ModifyLogAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLogAlarmTaskRequest) SetInstanceId(v string) *ModifyLogAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLogAlarmTaskRequest) SetTaskName(v string) *ModifyLogAlarmTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyLogAlarmTaskRequest) SetTaskId(v int32) *ModifyLogAlarmTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyLogAlarmTaskRequest) SetToMailList(v []*string) *ModifyLogAlarmTaskRequest {
	s.ToMailList = v
	return s
}

func (s *ModifyLogAlarmTaskRequest) SetDbIdList(v []*string) *ModifyLogAlarmTaskRequest {
	s.DbIdList = v
	return s
}

func (s *ModifyLogAlarmTaskRequest) SetRiskLevelList(v []*string) *ModifyLogAlarmTaskRequest {
	s.RiskLevelList = v
	return s
}

type ModifyLogAlarmTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogAlarmTaskResponseBody) SetRequestId(v string) *ModifyLogAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogAlarmTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLogAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogAlarmTaskResponse) SetHeaders(v map[string]*string) *ModifyLogAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogAlarmTaskResponse) SetBody(v *ModifyLogAlarmTaskResponseBody) *ModifyLogAlarmTaskResponse {
	s.Body = v
	return s
}

type ModifyPlanRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlanRequest) SetCommodityCode(v string) *ModifyPlanRequest {
	s.CommodityCode = &v
	return s
}

func (s *ModifyPlanRequest) SetInstanceId(v string) *ModifyPlanRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPlanRequest) SetRegionId(v string) *ModifyPlanRequest {
	s.RegionId = &v
	return s
}

type ModifyPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlanResponseBody) SetRequestId(v string) *ModifyPlanResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlanResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlanResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlanResponse) SetHeaders(v map[string]*string) *ModifyPlanResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlanResponse) SetBody(v *ModifyPlanResponseBody) *ModifyPlanResponse {
	s.Body = v
	return s
}

type ModifyReportPushTaskRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *int32    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ScheduleType *string   `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	ScheduleTime *string   `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	JobStatus    *int32    `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	ReportType   []*string `json:"ReportType,omitempty" xml:"ReportType,omitempty" type:"Repeated"`
	DbList       []*string `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	EmailList    []*string `json:"EmailList,omitempty" xml:"EmailList,omitempty" type:"Repeated"`
}

func (s ModifyReportPushTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyReportPushTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyReportPushTaskRequest) SetRegionId(v string) *ModifyReportPushTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyReportPushTaskRequest) SetInstanceId(v string) *ModifyReportPushTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyReportPushTaskRequest) SetJobId(v int32) *ModifyReportPushTaskRequest {
	s.JobId = &v
	return s
}

func (s *ModifyReportPushTaskRequest) SetScheduleType(v string) *ModifyReportPushTaskRequest {
	s.ScheduleType = &v
	return s
}

func (s *ModifyReportPushTaskRequest) SetScheduleTime(v string) *ModifyReportPushTaskRequest {
	s.ScheduleTime = &v
	return s
}

func (s *ModifyReportPushTaskRequest) SetJobStatus(v int32) *ModifyReportPushTaskRequest {
	s.JobStatus = &v
	return s
}

func (s *ModifyReportPushTaskRequest) SetReportType(v []*string) *ModifyReportPushTaskRequest {
	s.ReportType = v
	return s
}

func (s *ModifyReportPushTaskRequest) SetDbList(v []*string) *ModifyReportPushTaskRequest {
	s.DbList = v
	return s
}

func (s *ModifyReportPushTaskRequest) SetEmailList(v []*string) *ModifyReportPushTaskRequest {
	s.EmailList = v
	return s
}

type ModifyReportPushTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReportPushTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyReportPushTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReportPushTaskResponseBody) SetRequestId(v string) *ModifyReportPushTaskResponseBody {
	s.RequestId = &v
	return s
}

type ModifyReportPushTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyReportPushTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyReportPushTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyReportPushTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyReportPushTaskResponse) SetHeaders(v map[string]*string) *ModifyReportPushTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyReportPushTaskResponse) SetBody(v *ModifyReportPushTaskResponseBody) *ModifyReportPushTaskResponse {
	s.Body = v
	return s
}

type ModifyRuleGroupRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	GroupKeyId *string `json:"GroupKeyId,omitempty" xml:"GroupKeyId,omitempty"`
}

func (s ModifyRuleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleGroupRequest) SetRegionId(v string) *ModifyRuleGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRuleGroupRequest) SetInstanceId(v string) *ModifyRuleGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRuleGroupRequest) SetGroupId(v string) *ModifyRuleGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyRuleGroupRequest) SetGroupName(v string) *ModifyRuleGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyRuleGroupRequest) SetGroupType(v string) *ModifyRuleGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyRuleGroupRequest) SetGroupKeyId(v string) *ModifyRuleGroupRequest {
	s.GroupKeyId = &v
	return s
}

type ModifyRuleGroupResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRuleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleGroupResponseBody) SetServerData(v string) *ModifyRuleGroupResponseBody {
	s.ServerData = &v
	return s
}

func (s *ModifyRuleGroupResponseBody) SetRequestId(v string) *ModifyRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRuleGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRuleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyRuleGroupResponse) SetHeaders(v map[string]*string) *ModifyRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyRuleGroupResponse) SetBody(v *ModifyRuleGroupResponseBody) *ModifyRuleGroupResponse {
	s.Body = v
	return s
}

type ModifySystemAlarmTaskRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskName   *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskId     *int32    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ToMailList []*string `json:"ToMailList,omitempty" xml:"ToMailList,omitempty" type:"Repeated"`
}

func (s ModifySystemAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifySystemAlarmTaskRequest) SetRegionId(v string) *ModifySystemAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySystemAlarmTaskRequest) SetInstanceId(v string) *ModifySystemAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySystemAlarmTaskRequest) SetTaskName(v string) *ModifySystemAlarmTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ModifySystemAlarmTaskRequest) SetTaskId(v int32) *ModifySystemAlarmTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifySystemAlarmTaskRequest) SetToMailList(v []*string) *ModifySystemAlarmTaskRequest {
	s.ToMailList = v
	return s
}

type ModifySystemAlarmTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySystemAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySystemAlarmTaskResponseBody) SetRequestId(v string) *ModifySystemAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type ModifySystemAlarmTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySystemAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySystemAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySystemAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifySystemAlarmTaskResponse) SetHeaders(v map[string]*string) *ModifySystemAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifySystemAlarmTaskResponse) SetBody(v *ModifySystemAlarmTaskResponseBody) *ModifySystemAlarmTaskResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
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

type PciReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s PciReportRequest) String() string {
	return tea.Prettify(s)
}

func (s PciReportRequest) GoString() string {
	return s.String()
}

func (s *PciReportRequest) SetRegionId(v string) *PciReportRequest {
	s.RegionId = &v
	return s
}

func (s *PciReportRequest) SetInstanceId(v string) *PciReportRequest {
	s.InstanceId = &v
	return s
}

func (s *PciReportRequest) SetDbId(v int32) *PciReportRequest {
	s.DbId = &v
	return s
}

func (s *PciReportRequest) SetBeginDate(v string) *PciReportRequest {
	s.BeginDate = &v
	return s
}

func (s *PciReportRequest) SetEndDate(v string) *PciReportRequest {
	s.EndDate = &v
	return s
}

type PciReportResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PciReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PciReportResponseBody) GoString() string {
	return s.String()
}

func (s *PciReportResponseBody) SetServerData(v string) *PciReportResponseBody {
	s.ServerData = &v
	return s
}

func (s *PciReportResponseBody) SetRequestId(v string) *PciReportResponseBody {
	s.RequestId = &v
	return s
}

type PciReportResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PciReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PciReportResponse) String() string {
	return tea.Prettify(s)
}

func (s PciReportResponse) GoString() string {
	return s.String()
}

func (s *PciReportResponse) SetHeaders(v map[string]*string) *PciReportResponse {
	s.Headers = v
	return s
}

func (s *PciReportResponse) SetBody(v *PciReportResponseBody) *PciReportResponse {
	s.Body = v
	return s
}

type PutLoginCountRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PutLoginCountRequest) String() string {
	return tea.Prettify(s)
}

func (s PutLoginCountRequest) GoString() string {
	return s.String()
}

func (s *PutLoginCountRequest) SetRegionId(v string) *PutLoginCountRequest {
	s.RegionId = &v
	return s
}

func (s *PutLoginCountRequest) SetInstanceId(v string) *PutLoginCountRequest {
	s.InstanceId = &v
	return s
}

type PutLoginCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutLoginCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutLoginCountResponseBody) GoString() string {
	return s.String()
}

func (s *PutLoginCountResponseBody) SetRequestId(v string) *PutLoginCountResponseBody {
	s.RequestId = &v
	return s
}

type PutLoginCountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutLoginCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutLoginCountResponse) String() string {
	return tea.Prettify(s)
}

func (s PutLoginCountResponse) GoString() string {
	return s.String()
}

func (s *PutLoginCountResponse) SetHeaders(v map[string]*string) *PutLoginCountResponse {
	s.Headers = v
	return s
}

func (s *PutLoginCountResponse) SetBody(v *PutLoginCountResponseBody) *PutLoginCountResponse {
	s.Body = v
	return s
}

type ReadMarkSystemAlarmsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AlarmIds   []*string `json:"AlarmIds,omitempty" xml:"AlarmIds,omitempty" type:"Repeated"`
}

func (s ReadMarkSystemAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMarkSystemAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ReadMarkSystemAlarmsRequest) SetRegionId(v string) *ReadMarkSystemAlarmsRequest {
	s.RegionId = &v
	return s
}

func (s *ReadMarkSystemAlarmsRequest) SetInstanceId(v string) *ReadMarkSystemAlarmsRequest {
	s.InstanceId = &v
	return s
}

func (s *ReadMarkSystemAlarmsRequest) SetAlarmIds(v []*string) *ReadMarkSystemAlarmsRequest {
	s.AlarmIds = v
	return s
}

type ReadMarkSystemAlarmsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReadMarkSystemAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadMarkSystemAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMarkSystemAlarmsResponseBody) SetRequestId(v string) *ReadMarkSystemAlarmsResponseBody {
	s.RequestId = &v
	return s
}

type ReadMarkSystemAlarmsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReadMarkSystemAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReadMarkSystemAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMarkSystemAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ReadMarkSystemAlarmsResponse) SetHeaders(v map[string]*string) *ReadMarkSystemAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ReadMarkSystemAlarmsResponse) SetBody(v *ReadMarkSystemAlarmsResponseBody) *ReadMarkSystemAlarmsResponse {
	s.Body = v
	return s
}

type RefundInstanceRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RefundInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundInstanceRequest) SetInstanceId(v string) *RefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundInstanceRequest) SetServiceCode(v string) *RefundInstanceRequest {
	s.ServiceCode = &v
	return s
}

func (s *RefundInstanceRequest) SetRegionId(v string) *RefundInstanceRequest {
	s.RegionId = &v
	return s
}

type RefundInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponseBody) SetRequestId(v string) *RefundInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RefundInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceResponse) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponse) SetHeaders(v map[string]*string) *RefundInstanceResponse {
	s.Headers = v
	return s
}

func (s *RefundInstanceResponse) SetBody(v *RefundInstanceResponseBody) *RefundInstanceResponse {
	s.Body = v
	return s
}

type RegisterNoticeMailRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mail       *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	Vcode      *string `json:"Vcode,omitempty" xml:"Vcode,omitempty"`
}

func (s RegisterNoticeMailRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterNoticeMailRequest) GoString() string {
	return s.String()
}

func (s *RegisterNoticeMailRequest) SetRegionId(v string) *RegisterNoticeMailRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterNoticeMailRequest) SetInstanceId(v string) *RegisterNoticeMailRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterNoticeMailRequest) SetMail(v string) *RegisterNoticeMailRequest {
	s.Mail = &v
	return s
}

func (s *RegisterNoticeMailRequest) SetVcode(v string) *RegisterNoticeMailRequest {
	s.Vcode = &v
	return s
}

type RegisterNoticeMailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterNoticeMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterNoticeMailResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterNoticeMailResponseBody) SetRequestId(v string) *RegisterNoticeMailResponseBody {
	s.RequestId = &v
	return s
}

type RegisterNoticeMailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterNoticeMailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterNoticeMailResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterNoticeMailResponse) GoString() string {
	return s.String()
}

func (s *RegisterNoticeMailResponse) SetHeaders(v map[string]*string) *RegisterNoticeMailResponse {
	s.Headers = v
	return s
}

func (s *RegisterNoticeMailResponse) SetBody(v *RegisterNoticeMailResponseBody) *RegisterNoticeMailResponse {
	s.Body = v
	return s
}

type RemoveLogMaskConfigRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaskId     *int32  `json:"MaskId,omitempty" xml:"MaskId,omitempty"`
}

func (s RemoveLogMaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveLogMaskConfigRequest) GoString() string {
	return s.String()
}

func (s *RemoveLogMaskConfigRequest) SetRegionId(v string) *RemoveLogMaskConfigRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveLogMaskConfigRequest) SetInstanceId(v string) *RemoveLogMaskConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveLogMaskConfigRequest) SetMaskId(v int32) *RemoveLogMaskConfigRequest {
	s.MaskId = &v
	return s
}

type RemoveLogMaskConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveLogMaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveLogMaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveLogMaskConfigResponseBody) SetRequestId(v string) *RemoveLogMaskConfigResponseBody {
	s.RequestId = &v
	return s
}

type RemoveLogMaskConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveLogMaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveLogMaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveLogMaskConfigResponse) GoString() string {
	return s.String()
}

func (s *RemoveLogMaskConfigResponse) SetHeaders(v map[string]*string) *RemoveLogMaskConfigResponse {
	s.Headers = v
	return s
}

func (s *RemoveLogMaskConfigResponse) SetBody(v *RemoveLogMaskConfigResponseBody) *RemoveLogMaskConfigResponse {
	s.Body = v
	return s
}

type SendVerifyCodeMailRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mail       *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
}

func (s SendVerifyCodeMailRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeMailRequest) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeMailRequest) SetRegionId(v string) *SendVerifyCodeMailRequest {
	s.RegionId = &v
	return s
}

func (s *SendVerifyCodeMailRequest) SetInstanceId(v string) *SendVerifyCodeMailRequest {
	s.InstanceId = &v
	return s
}

func (s *SendVerifyCodeMailRequest) SetMail(v string) *SendVerifyCodeMailRequest {
	s.Mail = &v
	return s
}

type SendVerifyCodeMailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerifyCodeMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeMailResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeMailResponseBody) SetRequestId(v string) *SendVerifyCodeMailResponseBody {
	s.RequestId = &v
	return s
}

type SendVerifyCodeMailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendVerifyCodeMailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendVerifyCodeMailResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeMailResponse) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeMailResponse) SetHeaders(v map[string]*string) *SendVerifyCodeMailResponse {
	s.Headers = v
	return s
}

func (s *SendVerifyCodeMailResponse) SetBody(v *SendVerifyCodeMailResponseBody) *SendVerifyCodeMailResponse {
	s.Body = v
	return s
}

type SoxReportRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s SoxReportRequest) String() string {
	return tea.Prettify(s)
}

func (s SoxReportRequest) GoString() string {
	return s.String()
}

func (s *SoxReportRequest) SetRegionId(v string) *SoxReportRequest {
	s.RegionId = &v
	return s
}

func (s *SoxReportRequest) SetInstanceId(v string) *SoxReportRequest {
	s.InstanceId = &v
	return s
}

func (s *SoxReportRequest) SetDbId(v int32) *SoxReportRequest {
	s.DbId = &v
	return s
}

func (s *SoxReportRequest) SetBeginDate(v string) *SoxReportRequest {
	s.BeginDate = &v
	return s
}

func (s *SoxReportRequest) SetEndDate(v string) *SoxReportRequest {
	s.EndDate = &v
	return s
}

type SoxReportResponseBody struct {
	ServerData *string `json:"ServerData,omitempty" xml:"ServerData,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SoxReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SoxReportResponseBody) GoString() string {
	return s.String()
}

func (s *SoxReportResponseBody) SetServerData(v string) *SoxReportResponseBody {
	s.ServerData = &v
	return s
}

func (s *SoxReportResponseBody) SetRequestId(v string) *SoxReportResponseBody {
	s.RequestId = &v
	return s
}

type SoxReportResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SoxReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SoxReportResponse) String() string {
	return tea.Prettify(s)
}

func (s SoxReportResponse) GoString() string {
	return s.String()
}

func (s *SoxReportResponse) SetHeaders(v map[string]*string) *SoxReportResponse {
	s.Headers = v
	return s
}

func (s *SoxReportResponse) SetBody(v *SoxReportResponseBody) *SoxReportResponse {
	s.Body = v
	return s
}

type StartAlarmTaskRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskIds    []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s StartAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *StartAlarmTaskRequest) SetRegionId(v string) *StartAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StartAlarmTaskRequest) SetInstanceId(v string) *StartAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StartAlarmTaskRequest) SetTaskIds(v []*string) *StartAlarmTaskRequest {
	s.TaskIds = v
	return s
}

type StartAlarmTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartAlarmTaskResponseBody) SetRequestId(v string) *StartAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartAlarmTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *StartAlarmTaskResponse) SetHeaders(v map[string]*string) *StartAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *StartAlarmTaskResponse) SetBody(v *StartAlarmTaskResponseBody) *StartAlarmTaskResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VswitchId  *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetVswitchId(v string) *StartInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

type StartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type StopAlarmTaskRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskIds    []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s StopAlarmTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAlarmTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAlarmTaskRequest) SetRegionId(v string) *StopAlarmTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopAlarmTaskRequest) SetInstanceId(v string) *StopAlarmTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StopAlarmTaskRequest) SetTaskIds(v []*string) *StopAlarmTaskRequest {
	s.TaskIds = v
	return s
}

type StopAlarmTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAlarmTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAlarmTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopAlarmTaskResponseBody) SetRequestId(v string) *StopAlarmTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopAlarmTaskResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopAlarmTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAlarmTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAlarmTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAlarmTaskResponse) SetHeaders(v map[string]*string) *StopAlarmTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAlarmTaskResponse) SetBody(v *StopAlarmTaskResponseBody) *StopAlarmTaskResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeInstanceVersionRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionRequest) SetInstanceId(v string) *UpgradeInstanceVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRegionId(v string) *UpgradeInstanceVersionRequest {
	s.RegionId = &v
	return s
}

type UpgradeInstanceVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponseBody) SetRequestId(v string) *UpgradeInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeInstanceVersionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponse) SetHeaders(v map[string]*string) *UpgradeInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("yundun-dbaudit"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddLogMaskConfigWithOptions(request *AddLogMaskConfigRequest, runtime *util.RuntimeOptions) (_result *AddLogMaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddLogMaskConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddLogMaskConfig"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLogMaskConfig(request *AddLogMaskConfigRequest) (_result *AddLogMaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLogMaskConfigResponse{}
	_body, _err := client.AddLogMaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateDbToRuleWithOptions(request *AssociateDbToRuleRequest, runtime *util.RuntimeOptions) (_result *AssociateDbToRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateDbToRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateDbToRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateDbToRule(request *AssociateDbToRuleRequest) (_result *AssociateDbToRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateDbToRuleResponse{}
	_body, _err := client.AssociateDbToRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateRuleToDbWithOptions(request *AssociateRuleToDbRequest, runtime *util.RuntimeOptions) (_result *AssociateRuleToDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateRuleToDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateRuleToDb"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateRuleToDb(request *AssociateRuleToDbRequest) (_result *AssociateRuleToDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateRuleToDbResponse{}
	_body, _err := client.AssociateRuleToDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeAgentStatusWithOptions(request *ChangeAgentStatusRequest, runtime *util.RuntimeOptions) (_result *ChangeAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeAgentStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeAgentStatus"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeAgentStatus(request *ChangeAgentStatusRequest) (_result *ChangeAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeAgentStatusResponse{}
	_body, _err := client.ChangeAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeLogMaskConfigStateWithOptions(request *ChangeLogMaskConfigStateRequest, runtime *util.RuntimeOptions) (_result *ChangeLogMaskConfigStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeLogMaskConfigStateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeLogMaskConfigState"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeLogMaskConfigState(request *ChangeLogMaskConfigStateRequest) (_result *ChangeLogMaskConfigStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeLogMaskConfigStateResponse{}
	_body, _err := client.ChangeLogMaskConfigStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeRulePriorityWithOptions(request *ChangeRulePriorityRequest, runtime *util.RuntimeOptions) (_result *ChangeRulePriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeRulePriorityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeRulePriority"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeRulePriority(request *ChangeRulePriorityRequest) (_result *ChangeRulePriorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeRulePriorityResponse{}
	_body, _err := client.ChangeRulePriorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeRuleStatusWithOptions(request *ChangeRuleStatusRequest, runtime *util.RuntimeOptions) (_result *ChangeRuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeRuleStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeRuleStatus"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeRuleStatus(request *ChangeRuleStatusRequest) (_result *ChangeRuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeRuleStatusResponse{}
	_body, _err := client.ChangeRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMailRegisteredWithOptions(request *CheckMailRegisteredRequest, runtime *util.RuntimeOptions) (_result *CheckMailRegisteredResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckMailRegisteredResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckMailRegistered"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMailRegistered(request *CheckMailRegisteredRequest) (_result *CheckMailRegisteredResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMailRegisteredResponse{}
	_body, _err := client.CheckMailRegisteredWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearAgentRecordsWithOptions(request *ClearAgentRecordsRequest, runtime *util.RuntimeOptions) (_result *ClearAgentRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClearAgentRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClearAgentRecords"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearAgentRecords(request *ClearAgentRecordsRequest) (_result *ClearAgentRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearAgentRecordsResponse{}
	_body, _err := client.ClearAgentRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigInstanceNetworkWithOptions(request *ConfigInstanceNetworkRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigInstanceNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigInstanceNetwork"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigInstanceNetwork(request *ConfigInstanceNetworkRequest) (_result *ConfigInstanceNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceNetworkResponse{}
	_body, _err := client.ConfigInstanceNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataSource"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGradeProtectionReportWithOptions(request *CreateGradeProtectionReportRequest, runtime *util.RuntimeOptions) (_result *CreateGradeProtectionReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGradeProtectionReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGradeProtectionReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGradeProtectionReport(request *CreateGradeProtectionReportRequest) (_result *CreateGradeProtectionReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGradeProtectionReportResponse{}
	_body, _err := client.CreateGradeProtectionReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntegratedReportWithOptions(request *CreateIntegratedReportRequest, runtime *util.RuntimeOptions) (_result *CreateIntegratedReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIntegratedReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIntegratedReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntegratedReport(request *CreateIntegratedReportRequest) (_result *CreateIntegratedReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntegratedReportResponse{}
	_body, _err := client.CreateIntegratedReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogAlarmTaskWithOptions(request *CreateLogAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *CreateLogAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLogAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLogAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogAlarmTask(request *CreateLogAlarmTaskRequest) (_result *CreateLogAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLogAlarmTaskResponse{}
	_body, _err := client.CreateLogAlarmTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePCIReportWithOptions(request *CreatePCIReportRequest, runtime *util.RuntimeOptions) (_result *CreatePCIReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePCIReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePCIReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePCIReport(request *CreatePCIReportRequest) (_result *CreatePCIReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePCIReportResponse{}
	_body, _err := client.CreatePCIReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateReportPushTaskWithOptions(request *CreateReportPushTaskRequest, runtime *util.RuntimeOptions) (_result *CreateReportPushTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateReportPushTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateReportPushTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReportPushTask(request *CreateReportPushTaskRequest) (_result *CreateReportPushTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReportPushTaskResponse{}
	_body, _err := client.CreateReportPushTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleGroupWithOptions(request *CreateRuleGroupRequest, runtime *util.RuntimeOptions) (_result *CreateRuleGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRuleGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRuleGroup"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRuleGroup(request *CreateRuleGroupRequest) (_result *CreateRuleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleGroupResponse{}
	_body, _err := client.CreateRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSOXReportWithOptions(request *CreateSOXReportRequest, runtime *util.RuntimeOptions) (_result *CreateSOXReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSOXReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSOXReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSOXReport(request *CreateSOXReportRequest) (_result *CreateSOXReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSOXReportResponse{}
	_body, _err := client.CreateSOXReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSqlRuleWithOptions(request *CreateSqlRuleRequest, runtime *util.RuntimeOptions) (_result *CreateSqlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSqlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSqlRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSqlRule(request *CreateSqlRuleRequest) (_result *CreateSqlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSqlRuleResponse{}
	_body, _err := client.CreateSqlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSystemAlarmTaskWithOptions(request *CreateSystemAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *CreateSystemAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSystemAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSystemAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSystemAlarmTask(request *CreateSystemAlarmTaskRequest) (_result *CreateSystemAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSystemAlarmTaskResponse{}
	_body, _err := client.CreateSystemAlarmTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlarmTaskWithOptions(request *DeleteAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlarmTask(request *DeleteAlarmTaskRequest) (_result *DeleteAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlarmTaskResponse{}
	_body, _err := client.DeleteAlarmTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDataSource"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteReportPushTaskWithOptions(request *DeleteReportPushTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteReportPushTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteReportPushTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteReportPushTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteReportPushTask(request *DeleteReportPushTaskRequest) (_result *DeleteReportPushTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteReportPushTaskResponse{}
	_body, _err := client.DeleteReportPushTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleGroupWithOptions(request *DeleteRuleGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRuleGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRuleGroup"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRuleGroup(request *DeleteRuleGroupRequest) (_result *DeleteRuleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleGroupResponse{}
	_body, _err := client.DeleteRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeregisterTemplatesFromRuleWithOptions(request *DeregisterTemplatesFromRuleRequest, runtime *util.RuntimeOptions) (_result *DeregisterTemplatesFromRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeregisterTemplatesFromRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeregisterTemplatesFromRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeregisterTemplatesFromRule(request *DeregisterTemplatesFromRuleRequest) (_result *DeregisterTemplatesFromRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeregisterTemplatesFromRuleResponse{}
	_body, _err := client.DeregisterTemplatesFromRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAttribute"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeLoginTicketWithOptions(request *DescribeLoginTicketRequest, runtime *util.RuntimeOptions) (_result *DescribeLoginTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLoginTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLoginTicket"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoginTicket(request *DescribeLoginTicketRequest) (_result *DescribeLoginTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoginTicketResponse{}
	_body, _err := client.DescribeLoginTicketWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSyncInfoWithOptions(request *DescribeSyncInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSyncInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSyncInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSyncInfo"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSyncInfo(request *DescribeSyncInfoRequest) (_result *DescribeSyncInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSyncInfoResponse{}
	_body, _err := client.DescribeSyncInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLogMaskConfigsWithOptions(request *DisableLogMaskConfigsRequest, runtime *util.RuntimeOptions) (_result *DisableLogMaskConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableLogMaskConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableLogMaskConfigs"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLogMaskConfigs(request *DisableLogMaskConfigsRequest) (_result *DisableLogMaskConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLogMaskConfigsResponse{}
	_body, _err := client.DisableLogMaskConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateRulesFromDbWithOptions(request *DissociateRulesFromDbRequest, runtime *util.RuntimeOptions) (_result *DissociateRulesFromDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateRulesFromDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateRulesFromDb"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateRulesFromDb(request *DissociateRulesFromDbRequest) (_result *DissociateRulesFromDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateRulesFromDbResponse{}
	_body, _err := client.DissociateRulesFromDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateTemplatesFromRuleWithOptions(request *DissociateTemplatesFromRuleRequest, runtime *util.RuntimeOptions) (_result *DissociateTemplatesFromRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateTemplatesFromRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateTemplatesFromRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateTemplatesFromRule(request *DissociateTemplatesFromRuleRequest) (_result *DissociateTemplatesFromRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateTemplatesFromRuleResponse{}
	_body, _err := client.DissociateTemplatesFromRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditLogMaskConfigWithOptions(request *EditLogMaskConfigRequest, runtime *util.RuntimeOptions) (_result *EditLogMaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditLogMaskConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditLogMaskConfig"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditLogMaskConfig(request *EditLogMaskConfigRequest) (_result *EditLogMaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditLogMaskConfigResponse{}
	_body, _err := client.EditLogMaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLogMaskConfigsWithOptions(request *EnableLogMaskConfigsRequest, runtime *util.RuntimeOptions) (_result *EnableLogMaskConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableLogMaskConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableLogMaskConfigs"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLogMaskConfigs(request *EnableLogMaskConfigsRequest) (_result *EnableLogMaskConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLogMaskConfigsResponse{}
	_body, _err := client.EnableLogMaskConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentFileUrlWithOptions(request *GetAgentFileUrlRequest, runtime *util.RuntimeOptions) (_result *GetAgentFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAgentFileUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentFileUrl"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentFileUrl(request *GetAgentFileUrlRequest) (_result *GetAgentFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentFileUrlResponse{}
	_body, _err := client.GetAgentFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentListWithOptions(request *GetAgentListRequest, runtime *util.RuntimeOptions) (_result *GetAgentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAgentListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgentList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentList(request *GetAgentListRequest) (_result *GetAgentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentListResponse{}
	_body, _err := client.GetAgentListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppointOperationWithOptions(request *GetAppointOperationRequest, runtime *util.RuntimeOptions) (_result *GetAppointOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppointOperationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAppointOperation"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppointOperation(request *GetAppointOperationRequest) (_result *GetAppointOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppointOperationResponse{}
	_body, _err := client.GetAppointOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuditCountWithOptions(request *GetAuditCountRequest, runtime *util.RuntimeOptions) (_result *GetAuditCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuditCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuditCount"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuditCount(request *GetAuditCountRequest) (_result *GetAuditCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditCountResponse{}
	_body, _err := client.GetAuditCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuditCountDistributionWithOptions(request *GetAuditCountDistributionRequest, runtime *util.RuntimeOptions) (_result *GetAuditCountDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuditCountDistributionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuditCountDistribution"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuditCountDistribution(request *GetAuditCountDistributionRequest) (_result *GetAuditCountDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditCountDistributionResponse{}
	_body, _err := client.GetAuditCountDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBaseTemplateListWithOptions(request *GetBaseTemplateListRequest, runtime *util.RuntimeOptions) (_result *GetBaseTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBaseTemplateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBaseTemplateList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBaseTemplateList(request *GetBaseTemplateListRequest) (_result *GetBaseTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBaseTemplateListResponse{}
	_body, _err := client.GetBaseTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDasUsageWithOptions(request *GetDasUsageRequest, runtime *util.RuntimeOptions) (_result *GetDasUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDasUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDasUsage"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDasUsage(request *GetDasUsageRequest) (_result *GetDasUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDasUsageResponse{}
	_body, _err := client.GetDasUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDBAuditCountListWithOptions(request *GetDBAuditCountListRequest, runtime *util.RuntimeOptions) (_result *GetDBAuditCountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDBAuditCountListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDBAuditCountList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDBAuditCountList(request *GetDBAuditCountListRequest) (_result *GetDBAuditCountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDBAuditCountListResponse{}
	_body, _err := client.GetDBAuditCountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupDetailWithOptions(request *GetGroupDetailRequest, runtime *util.RuntimeOptions) (_result *GetGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetGroupDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGroupDetail"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroupDetail(request *GetGroupDetailRequest) (_result *GetGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupDetailResponse{}
	_body, _err := client.GetGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLicenseWithOptions(request *GetLicenseRequest, runtime *util.RuntimeOptions) (_result *GetLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLicense"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLicense(request *GetLicenseRequest) (_result *GetLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLicenseResponse{}
	_body, _err := client.GetLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogDetailWithOptions(request *GetLogDetailRequest, runtime *util.RuntimeOptions) (_result *GetLogDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogDetail"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogDetail(request *GetLogDetailRequest) (_result *GetLogDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogDetailResponse{}
	_body, _err := client.GetLogDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogListWithOptions(request *GetLogListRequest, runtime *util.RuntimeOptions) (_result *GetLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogList(request *GetLogListRequest) (_result *GetLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogListResponse{}
	_body, _err := client.GetLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogMaskConfigsWithOptions(request *GetLogMaskConfigsRequest, runtime *util.RuntimeOptions) (_result *GetLogMaskConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogMaskConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogMaskConfigs"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogMaskConfigs(request *GetLogMaskConfigsRequest) (_result *GetLogMaskConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogMaskConfigsResponse{}
	_body, _err := client.GetLogMaskConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogQueryConditionWithOptions(request *GetLogQueryConditionRequest, runtime *util.RuntimeOptions) (_result *GetLogQueryConditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogQueryConditionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogQueryCondition"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogQueryCondition(request *GetLogQueryConditionRequest) (_result *GetLogQueryConditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogQueryConditionResponse{}
	_body, _err := client.GetLogQueryConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogStatisticsWithOptions(request *GetLogStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetLogStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogStatistics"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogStatistics(request *GetLogStatisticsRequest) (_result *GetLogStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogStatisticsResponse{}
	_body, _err := client.GetLogStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogTopDistributionWithOptions(request *GetLogTopDistributionRequest, runtime *util.RuntimeOptions) (_result *GetLogTopDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogTopDistributionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogTopDistribution"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogTopDistribution(request *GetLogTopDistributionRequest) (_result *GetLogTopDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogTopDistributionResponse{}
	_body, _err := client.GetLogTopDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogTopStatisticsWithOptions(request *GetLogTopStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetLogTopStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogTopStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogTopStatistics"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogTopStatistics(request *GetLogTopStatisticsRequest) (_result *GetLogTopStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogTopStatisticsResponse{}
	_body, _err := client.GetLogTopStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogTypeDistributionWithOptions(request *GetLogTypeDistributionRequest, runtime *util.RuntimeOptions) (_result *GetLogTypeDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogTypeDistributionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogTypeDistribution"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogTypeDistribution(request *GetLogTypeDistributionRequest) (_result *GetLogTypeDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogTypeDistributionResponse{}
	_body, _err := client.GetLogTypeDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNewSqlTemplateListWithOptions(request *GetNewSqlTemplateListRequest, runtime *util.RuntimeOptions) (_result *GetNewSqlTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetNewSqlTemplateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetNewSqlTemplateList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNewSqlTemplateList(request *GetNewSqlTemplateListRequest) (_result *GetNewSqlTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNewSqlTemplateListResponse{}
	_body, _err := client.GetNewSqlTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetReportFileUrlWithOptions(request *GetReportFileUrlRequest, runtime *util.RuntimeOptions) (_result *GetReportFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetReportFileUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetReportFileUrl"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetReportFileUrl(request *GetReportFileUrlRequest) (_result *GetReportFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReportFileUrlResponse{}
	_body, _err := client.GetReportFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRiskLevelDistributionWithOptions(request *GetRiskLevelDistributionRequest, runtime *util.RuntimeOptions) (_result *GetRiskLevelDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRiskLevelDistributionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRiskLevelDistribution"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRiskLevelDistribution(request *GetRiskLevelDistributionRequest) (_result *GetRiskLevelDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRiskLevelDistributionResponse{}
	_body, _err := client.GetRiskLevelDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRiskStatisticsWithOptions(request *GetRiskStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetRiskStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRiskStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRiskStatistics"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRiskStatistics(request *GetRiskStatisticsRequest) (_result *GetRiskStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRiskStatisticsResponse{}
	_body, _err := client.GetRiskStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleDetailWithOptions(request *GetRuleDetailRequest, runtime *util.RuntimeOptions) (_result *GetRuleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRuleDetail"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRuleDetail(request *GetRuleDetailRequest) (_result *GetRuleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.GetRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleGroupNameWithOptions(request *GetRuleGroupNameRequest, runtime *util.RuntimeOptions) (_result *GetRuleGroupNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRuleGroupNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRuleGroupName"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRuleGroupName(request *GetRuleGroupNameRequest) (_result *GetRuleGroupNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleGroupNameResponse{}
	_body, _err := client.GetRuleGroupNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSessionDetailWithOptions(request *GetSessionDetailRequest, runtime *util.RuntimeOptions) (_result *GetSessionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSessionDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSessionDetail"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSessionDetail(request *GetSessionDetailRequest) (_result *GetSessionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSessionDetailResponse{}
	_body, _err := client.GetSessionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSessionDistributionWithOptions(request *GetSessionDistributionRequest, runtime *util.RuntimeOptions) (_result *GetSessionDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSessionDistributionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSessionDistribution"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSessionDistribution(request *GetSessionDistributionRequest) (_result *GetSessionDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSessionDistributionResponse{}
	_body, _err := client.GetSessionDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSessionListWithOptions(request *GetSessionListRequest, runtime *util.RuntimeOptions) (_result *GetSessionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSessionListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSessionList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSessionList(request *GetSessionListRequest) (_result *GetSessionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSessionListResponse{}
	_body, _err := client.GetSessionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSessionQueryConditionWithOptions(request *GetSessionQueryConditionRequest, runtime *util.RuntimeOptions) (_result *GetSessionQueryConditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSessionQueryConditionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSessionQueryCondition"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSessionQueryCondition(request *GetSessionQueryConditionRequest) (_result *GetSessionQueryConditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSessionQueryConditionResponse{}
	_body, _err := client.GetSessionQueryConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSqlTemplateListWithOptions(request *GetSqlTemplateListRequest, runtime *util.RuntimeOptions) (_result *GetSqlTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSqlTemplateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSqlTemplateList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSqlTemplateList(request *GetSqlTemplateListRequest) (_result *GetSqlTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSqlTemplateListResponse{}
	_body, _err := client.GetSqlTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopSqlTemplateListWithOptions(request *GetTopSqlTemplateListRequest, runtime *util.RuntimeOptions) (_result *GetTopSqlTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTopSqlTemplateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTopSqlTemplateList"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopSqlTemplateList(request *GetTopSqlTemplateListRequest) (_result *GetTopSqlTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopSqlTemplateListResponse{}
	_body, _err := client.GetTopSqlTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GradeProtectionReportWithOptions(request *GradeProtectionReportRequest, runtime *util.RuntimeOptions) (_result *GradeProtectionReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GradeProtectionReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GradeProtectionReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GradeProtectionReport(request *GradeProtectionReportRequest) (_result *GradeProtectionReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GradeProtectionReportResponse{}
	_body, _err := client.GradeProtectionReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportDataSourceWithOptions(request *ImportDataSourceRequest, runtime *util.RuntimeOptions) (_result *ImportDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportDataSource"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportDataSource(request *ImportDataSourceRequest) (_result *ImportDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportDataSourceResponse{}
	_body, _err := client.ImportDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IntegratedReportWithOptions(request *IntegratedReportRequest, runtime *util.RuntimeOptions) (_result *IntegratedReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IntegratedReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IntegratedReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IntegratedReport(request *IntegratedReportRequest) (_result *IntegratedReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IntegratedReportResponse{}
	_body, _err := client.IntegratedReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAssociatedRulesWithOptions(request *ListAssociatedRulesRequest, runtime *util.RuntimeOptions) (_result *ListAssociatedRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAssociatedRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAssociatedRules"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAssociatedRules(request *ListAssociatedRulesRequest) (_result *ListAssociatedRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAssociatedRulesResponse{}
	_body, _err := client.ListAssociatedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceAttributeWithOptions(request *ListDataSourceAttributeRequest, runtime *util.RuntimeOptions) (_result *ListDataSourceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDataSourceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataSourceAttribute"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSourceAttribute(request *ListDataSourceAttributeRequest) (_result *ListDataSourceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourceAttributeResponse{}
	_body, _err := client.ListDataSourceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourcesWithOptions(request *ListDataSourcesRequest, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataSources"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSources(request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogAlarmTasksWithOptions(request *ListLogAlarmTasksRequest, runtime *util.RuntimeOptions) (_result *ListLogAlarmTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLogAlarmTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLogAlarmTasks"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogAlarmTasks(request *ListLogAlarmTasksRequest) (_result *ListLogAlarmTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLogAlarmTasksResponse{}
	_body, _err := client.ListLogAlarmTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListReportPushTasksWithOptions(request *ListReportPushTasksRequest, runtime *util.RuntimeOptions) (_result *ListReportPushTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListReportPushTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListReportPushTasks"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListReportPushTasks(request *ListReportPushTasksRequest) (_result *ListReportPushTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListReportPushTasksResponse{}
	_body, _err := client.ListReportPushTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRuleGroupsWithOptions(request *ListRuleGroupsRequest, runtime *util.RuntimeOptions) (_result *ListRuleGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRuleGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRuleGroups"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRuleGroups(request *ListRuleGroupsRequest) (_result *ListRuleGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRuleGroupsResponse{}
	_body, _err := client.ListRuleGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRules"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSqlTypeKeysForRuleWithOptions(request *ListSqlTypeKeysForRuleRequest, runtime *util.RuntimeOptions) (_result *ListSqlTypeKeysForRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSqlTypeKeysForRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSqlTypeKeysForRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSqlTypeKeysForRule(request *ListSqlTypeKeysForRuleRequest) (_result *ListSqlTypeKeysForRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSqlTypeKeysForRuleResponse{}
	_body, _err := client.ListSqlTypeKeysForRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSqlTypesForRuleWithOptions(request *ListSqlTypesForRuleRequest, runtime *util.RuntimeOptions) (_result *ListSqlTypesForRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSqlTypesForRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSqlTypesForRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSqlTypesForRule(request *ListSqlTypesForRuleRequest) (_result *ListSqlTypesForRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSqlTypesForRuleResponse{}
	_body, _err := client.ListSqlTypesForRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSupportDbTypesWithOptions(request *ListSupportDbTypesRequest, runtime *util.RuntimeOptions) (_result *ListSupportDbTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSupportDbTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSupportDbTypes"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSupportDbTypes(request *ListSupportDbTypesRequest) (_result *ListSupportDbTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSupportDbTypesResponse{}
	_body, _err := client.ListSupportDbTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemAlarmsWithOptions(request *ListSystemAlarmsRequest, runtime *util.RuntimeOptions) (_result *ListSystemAlarmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSystemAlarmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSystemAlarms"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemAlarms(request *ListSystemAlarmsRequest) (_result *ListSystemAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemAlarmsResponse{}
	_body, _err := client.ListSystemAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemAlarmTasksWithOptions(request *ListSystemAlarmTasksRequest, runtime *util.RuntimeOptions) (_result *ListSystemAlarmTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSystemAlarmTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSystemAlarmTasks"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemAlarmTasks(request *ListSystemAlarmTasksRequest) (_result *ListSystemAlarmTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemAlarmTasksResponse{}
	_body, _err := client.ListSystemAlarmTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTemplatesForSqlRuleWithOptions(request *ListTemplatesForSqlRuleRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesForSqlRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesForSqlRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplatesForSqlRule"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplatesForSqlRule(request *ListTemplatesForSqlRuleRequest) (_result *ListTemplatesForSqlRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesForSqlRuleResponse{}
	_body, _err := client.ListTemplatesForSqlRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsedSqlTypesWithOptions(request *ListUsedSqlTypesRequest, runtime *util.RuntimeOptions) (_result *ListUsedSqlTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUsedSqlTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsedSqlTypes"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsedSqlTypes(request *ListUsedSqlTypesRequest) (_result *ListUsedSqlTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsedSqlTypesResponse{}
	_body, _err := client.ListUsedSqlTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBaseTemplateStateWithOptions(request *ModifyBaseTemplateStateRequest, runtime *util.RuntimeOptions) (_result *ModifyBaseTemplateStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBaseTemplateStateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBaseTemplateState"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBaseTemplateState(request *ModifyBaseTemplateStateRequest) (_result *ModifyBaseTemplateStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBaseTemplateStateResponse{}
	_body, _err := client.ModifyBaseTemplateStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCustomNameWithOptions(request *ModifyCustomNameRequest, runtime *util.RuntimeOptions) (_result *ModifyCustomNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCustomNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCustomName"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCustomName(request *ModifyCustomNameRequest) (_result *ModifyCustomNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCustomNameResponse{}
	_body, _err := client.ModifyCustomNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataSourceWithOptions(request *ModifyDataSourceRequest, runtime *util.RuntimeOptions) (_result *ModifyDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDataSource"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataSource(request *ModifyDataSourceRequest) (_result *ModifyDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.ModifyDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataSourceAttributeWithOptions(request *ModifyDataSourceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDataSourceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDataSourceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDataSourceAttribute"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataSourceAttribute(request *ModifyDataSourceAttributeRequest) (_result *ModifyDataSourceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataSourceAttributeResponse{}
	_body, _err := client.ModifyDataSourceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAttribute"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogAlarmTaskWithOptions(request *ModifyLogAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyLogAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLogAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLogAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogAlarmTask(request *ModifyLogAlarmTaskRequest) (_result *ModifyLogAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogAlarmTaskResponse{}
	_body, _err := client.ModifyLogAlarmTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPlanWithOptions(request *ModifyPlanRequest, runtime *util.RuntimeOptions) (_result *ModifyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPlan"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPlan(request *ModifyPlanRequest) (_result *ModifyPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlanResponse{}
	_body, _err := client.ModifyPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyReportPushTaskWithOptions(request *ModifyReportPushTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyReportPushTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyReportPushTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyReportPushTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyReportPushTask(request *ModifyReportPushTaskRequest) (_result *ModifyReportPushTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyReportPushTaskResponse{}
	_body, _err := client.ModifyReportPushTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRuleGroupWithOptions(request *ModifyRuleGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyRuleGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRuleGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRuleGroup"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRuleGroup(request *ModifyRuleGroupRequest) (_result *ModifyRuleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRuleGroupResponse{}
	_body, _err := client.ModifyRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySystemAlarmTaskWithOptions(request *ModifySystemAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *ModifySystemAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySystemAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySystemAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySystemAlarmTask(request *ModifySystemAlarmTaskRequest) (_result *ModifySystemAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySystemAlarmTaskResponse{}
	_body, _err := client.ModifySystemAlarmTaskWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PciReportWithOptions(request *PciReportRequest, runtime *util.RuntimeOptions) (_result *PciReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PciReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PciReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PciReport(request *PciReportRequest) (_result *PciReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PciReportResponse{}
	_body, _err := client.PciReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutLoginCountWithOptions(request *PutLoginCountRequest, runtime *util.RuntimeOptions) (_result *PutLoginCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutLoginCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutLoginCount"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutLoginCount(request *PutLoginCountRequest) (_result *PutLoginCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutLoginCountResponse{}
	_body, _err := client.PutLoginCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReadMarkSystemAlarmsWithOptions(request *ReadMarkSystemAlarmsRequest, runtime *util.RuntimeOptions) (_result *ReadMarkSystemAlarmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReadMarkSystemAlarmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReadMarkSystemAlarms"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReadMarkSystemAlarms(request *ReadMarkSystemAlarmsRequest) (_result *ReadMarkSystemAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadMarkSystemAlarmsResponse{}
	_body, _err := client.ReadMarkSystemAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundInstanceWithOptions(request *RefundInstanceRequest, runtime *util.RuntimeOptions) (_result *RefundInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefundInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefundInstance"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundInstance(request *RefundInstanceRequest) (_result *RefundInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundInstanceResponse{}
	_body, _err := client.RefundInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterNoticeMailWithOptions(request *RegisterNoticeMailRequest, runtime *util.RuntimeOptions) (_result *RegisterNoticeMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterNoticeMailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterNoticeMail"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterNoticeMail(request *RegisterNoticeMailRequest) (_result *RegisterNoticeMailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterNoticeMailResponse{}
	_body, _err := client.RegisterNoticeMailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveLogMaskConfigWithOptions(request *RemoveLogMaskConfigRequest, runtime *util.RuntimeOptions) (_result *RemoveLogMaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveLogMaskConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveLogMaskConfig"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveLogMaskConfig(request *RemoveLogMaskConfigRequest) (_result *RemoveLogMaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveLogMaskConfigResponse{}
	_body, _err := client.RemoveLogMaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerifyCodeMailWithOptions(request *SendVerifyCodeMailRequest, runtime *util.RuntimeOptions) (_result *SendVerifyCodeMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendVerifyCodeMailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendVerifyCodeMail"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerifyCodeMail(request *SendVerifyCodeMailRequest) (_result *SendVerifyCodeMailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerifyCodeMailResponse{}
	_body, _err := client.SendVerifyCodeMailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SoxReportWithOptions(request *SoxReportRequest, runtime *util.RuntimeOptions) (_result *SoxReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SoxReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SoxReport"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SoxReport(request *SoxReportRequest) (_result *SoxReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SoxReportResponse{}
	_body, _err := client.SoxReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAlarmTaskWithOptions(request *StartAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *StartAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAlarmTask(request *StartAlarmTaskRequest) (_result *StartAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAlarmTaskResponse{}
	_body, _err := client.StartAlarmTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartInstance"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAlarmTaskWithOptions(request *StopAlarmTaskRequest, runtime *util.RuntimeOptions) (_result *StopAlarmTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopAlarmTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopAlarmTask"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAlarmTask(request *StopAlarmTaskRequest) (_result *StopAlarmTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAlarmTaskResponse{}
	_body, _err := client.StopAlarmTaskWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpgradeInstanceVersionWithOptions(request *UpgradeInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeInstanceVersion"), tea.String("2018-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstanceVersion(request *UpgradeInstanceVersionRequest) (_result *UpgradeInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.UpgradeInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

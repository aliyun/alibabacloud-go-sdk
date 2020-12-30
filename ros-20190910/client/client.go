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

type CancelUpdateStackRequest struct {
	StackId    *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CancelType *string `json:"CancelType,omitempty" xml:"CancelType,omitempty"`
}

func (s CancelUpdateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelUpdateStackRequest) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackRequest) SetStackId(v string) *CancelUpdateStackRequest {
	s.StackId = &v
	return s
}

func (s *CancelUpdateStackRequest) SetRegionId(v string) *CancelUpdateStackRequest {
	s.RegionId = &v
	return s
}

func (s *CancelUpdateStackRequest) SetCancelType(v string) *CancelUpdateStackRequest {
	s.CancelType = &v
	return s
}

type CancelUpdateStackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUpdateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelUpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackResponseBody) SetRequestId(v string) *CancelUpdateStackResponseBody {
	s.RequestId = &v
	return s
}

type CancelUpdateStackResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelUpdateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelUpdateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelUpdateStackResponse) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackResponse) SetHeaders(v map[string]*string) *CancelUpdateStackResponse {
	s.Headers = v
	return s
}

func (s *CancelUpdateStackResponse) SetBody(v *CancelUpdateStackResponseBody) *CancelUpdateStackResponse {
	s.Body = v
	return s
}

type ContinueCreateStackRequest struct {
	StackId             *string                                 `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId            *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RamRoleName         *string                                 `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	Mode                *string                                 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	TemplateBody        *string                                 `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateURL         *string                                 `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	DryRun              *bool                                   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	TemplateId          *string                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion     *string                                 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	RecreatingResources []*string                               `json:"RecreatingResources,omitempty" xml:"RecreatingResources,omitempty" type:"Repeated"`
	Parameters          []*ContinueCreateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s ContinueCreateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackRequest) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackRequest) SetStackId(v string) *ContinueCreateStackRequest {
	s.StackId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetRegionId(v string) *ContinueCreateStackRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetRamRoleName(v string) *ContinueCreateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *ContinueCreateStackRequest) SetMode(v string) *ContinueCreateStackRequest {
	s.Mode = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateBody(v string) *ContinueCreateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateURL(v string) *ContinueCreateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *ContinueCreateStackRequest) SetDryRun(v bool) *ContinueCreateStackRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateId(v string) *ContinueCreateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateVersion(v string) *ContinueCreateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ContinueCreateStackRequest) SetRecreatingResources(v []*string) *ContinueCreateStackRequest {
	s.RecreatingResources = v
	return s
}

func (s *ContinueCreateStackRequest) SetParameters(v []*ContinueCreateStackRequestParameters) *ContinueCreateStackRequest {
	s.Parameters = v
	return s
}

type ContinueCreateStackRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ContinueCreateStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackRequestParameters) SetParameterKey(v string) *ContinueCreateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *ContinueCreateStackRequestParameters) SetParameterValue(v string) *ContinueCreateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type ContinueCreateStackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackId   *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ContinueCreateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponseBody) SetRequestId(v string) *ContinueCreateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueCreateStackResponseBody) SetStackId(v string) *ContinueCreateStackResponseBody {
	s.StackId = &v
	return s
}

type ContinueCreateStackResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContinueCreateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinueCreateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackResponse) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponse) SetHeaders(v map[string]*string) *ContinueCreateStackResponse {
	s.Headers = v
	return s
}

func (s *ContinueCreateStackResponse) SetBody(v *ContinueCreateStackResponseBody) *ContinueCreateStackResponse {
	s.Body = v
	return s
}

type CreateChangeSetRequest struct {
	StackId                     *string                                    `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ChannelId                   *string                                    `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	StackPolicyURL              *string                                    `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	StackPolicyBody             *string                                    `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	StackName                   *string                                    `json:"StackName,omitempty" xml:"StackName,omitempty"`
	UsePreviousParameters       *bool                                      `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
	ChangeSetType               *string                                    `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	Description                 *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId                    *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken                 *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TemplateURL                 *string                                    `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	StackPolicyDuringUpdateURL  *string                                    `json:"StackPolicyDuringUpdateURL,omitempty" xml:"StackPolicyDuringUpdateURL,omitempty"`
	TemplateBody                *string                                    `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	UpdateAllowPolicy           *string                                    `json:"UpdateAllowPolicy,omitempty" xml:"UpdateAllowPolicy,omitempty"`
	TimeoutInMinutes            *int64                                     `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	ActivityId                  *string                                    `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	OrderSource                 *string                                    `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	DisableRollback             *bool                                      `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	ChangeSetName               *string                                    `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	StackPolicyDuringUpdateBody *string                                    `json:"StackPolicyDuringUpdateBody,omitempty" xml:"StackPolicyDuringUpdateBody,omitempty"`
	RamRoleName                 *string                                    `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ReplacementOption           *string                                    `json:"ReplacementOption,omitempty" xml:"ReplacementOption,omitempty"`
	TemplateId                  *string                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion             *string                                    `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters                  []*CreateChangeSetRequestParameters        `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	NotificationURLs            []*string                                  `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	ResourcesToImport           []*CreateChangeSetRequestResourcesToImport `json:"ResourcesToImport,omitempty" xml:"ResourcesToImport,omitempty" type:"Repeated"`
}

func (s CreateChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetRequest) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequest) SetStackId(v string) *CreateChangeSetRequest {
	s.StackId = &v
	return s
}

func (s *CreateChangeSetRequest) SetChannelId(v string) *CreateChangeSetRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyURL(v string) *CreateChangeSetRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyBody(v string) *CreateChangeSetRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackName(v string) *CreateChangeSetRequest {
	s.StackName = &v
	return s
}

func (s *CreateChangeSetRequest) SetUsePreviousParameters(v bool) *CreateChangeSetRequest {
	s.UsePreviousParameters = &v
	return s
}

func (s *CreateChangeSetRequest) SetChangeSetType(v string) *CreateChangeSetRequest {
	s.ChangeSetType = &v
	return s
}

func (s *CreateChangeSetRequest) SetDescription(v string) *CreateChangeSetRequest {
	s.Description = &v
	return s
}

func (s *CreateChangeSetRequest) SetRegionId(v string) *CreateChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateChangeSetRequest) SetClientToken(v string) *CreateChangeSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateURL(v string) *CreateChangeSetRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyDuringUpdateURL(v string) *CreateChangeSetRequest {
	s.StackPolicyDuringUpdateURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateBody(v string) *CreateChangeSetRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetUpdateAllowPolicy(v string) *CreateChangeSetRequest {
	s.UpdateAllowPolicy = &v
	return s
}

func (s *CreateChangeSetRequest) SetTimeoutInMinutes(v int64) *CreateChangeSetRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateChangeSetRequest) SetActivityId(v string) *CreateChangeSetRequest {
	s.ActivityId = &v
	return s
}

func (s *CreateChangeSetRequest) SetOrderSource(v string) *CreateChangeSetRequest {
	s.OrderSource = &v
	return s
}

func (s *CreateChangeSetRequest) SetDisableRollback(v bool) *CreateChangeSetRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateChangeSetRequest) SetChangeSetName(v string) *CreateChangeSetRequest {
	s.ChangeSetName = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyDuringUpdateBody(v string) *CreateChangeSetRequest {
	s.StackPolicyDuringUpdateBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetRamRoleName(v string) *CreateChangeSetRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateChangeSetRequest) SetReplacementOption(v string) *CreateChangeSetRequest {
	s.ReplacementOption = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateId(v string) *CreateChangeSetRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateVersion(v string) *CreateChangeSetRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateChangeSetRequest) SetParameters(v []*CreateChangeSetRequestParameters) *CreateChangeSetRequest {
	s.Parameters = v
	return s
}

func (s *CreateChangeSetRequest) SetNotificationURLs(v []*string) *CreateChangeSetRequest {
	s.NotificationURLs = v
	return s
}

func (s *CreateChangeSetRequest) SetResourcesToImport(v []*CreateChangeSetRequestResourcesToImport) *CreateChangeSetRequest {
	s.ResourcesToImport = v
	return s
}

type CreateChangeSetRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateChangeSetRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestParameters) SetParameterKey(v string) *CreateChangeSetRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateChangeSetRequestParameters) SetParameterValue(v string) *CreateChangeSetRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateChangeSetRequestResourcesToImport struct {
	ResourceIdentifier *string `json:"ResourceIdentifier,omitempty" xml:"ResourceIdentifier,omitempty"`
	LogicalResourceId  *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateChangeSetRequestResourcesToImport) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetRequestResourcesToImport) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestResourcesToImport) SetResourceIdentifier(v string) *CreateChangeSetRequestResourcesToImport {
	s.ResourceIdentifier = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) SetLogicalResourceId(v string) *CreateChangeSetRequestResourcesToImport {
	s.LogicalResourceId = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) SetResourceType(v string) *CreateChangeSetRequestResourcesToImport {
	s.ResourceType = &v
	return s
}

type CreateChangeSetResponseBody struct {
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackId     *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CreateChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChangeSetResponseBody) SetChangeSetId(v string) *CreateChangeSetResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *CreateChangeSetResponseBody) SetRequestId(v string) *CreateChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChangeSetResponseBody) SetStackId(v string) *CreateChangeSetResponseBody {
	s.StackId = &v
	return s
}

type CreateChangeSetResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetResponse) GoString() string {
	return s.String()
}

func (s *CreateChangeSetResponse) SetHeaders(v map[string]*string) *CreateChangeSetResponse {
	s.Headers = v
	return s
}

func (s *CreateChangeSetResponse) SetBody(v *CreateChangeSetResponseBody) *CreateChangeSetResponse {
	s.Body = v
	return s
}

type CreateStackRequest struct {
	DisableRollback    *bool                           `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	ChannelId          *string                         `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	TemplateBody       *string                         `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	StackPolicyURL     *string                         `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	TimeoutInMinutes   *int64                          `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	StackPolicyBody    *string                         `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	StackName          *string                         `json:"StackName,omitempty" xml:"StackName,omitempty"`
	RegionId           *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ActivityId         *string                         `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	OrderSource        *string                         `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	ClientToken        *string                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TemplateURL        *string                         `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	RamRoleName        *string                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	DeletionProtection *string                         `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	CreateOption       *string                         `json:"CreateOption,omitempty" xml:"CreateOption,omitempty"`
	TemplateId         *string                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion    *string                         `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters         []*CreateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	NotificationURLs   []*string                       `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
}

func (s CreateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackRequest) GoString() string {
	return s.String()
}

func (s *CreateStackRequest) SetDisableRollback(v bool) *CreateStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackRequest) SetChannelId(v string) *CreateStackRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateBody(v string) *CreateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateStackRequest) SetStackPolicyURL(v string) *CreateStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *CreateStackRequest) SetTimeoutInMinutes(v int64) *CreateStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateStackRequest) SetStackPolicyBody(v string) *CreateStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *CreateStackRequest) SetStackName(v string) *CreateStackRequest {
	s.StackName = &v
	return s
}

func (s *CreateStackRequest) SetRegionId(v string) *CreateStackRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackRequest) SetActivityId(v string) *CreateStackRequest {
	s.ActivityId = &v
	return s
}

func (s *CreateStackRequest) SetOrderSource(v string) *CreateStackRequest {
	s.OrderSource = &v
	return s
}

func (s *CreateStackRequest) SetClientToken(v string) *CreateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackRequest) SetTemplateURL(v string) *CreateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateStackRequest) SetRamRoleName(v string) *CreateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateStackRequest) SetDeletionProtection(v string) *CreateStackRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateStackRequest) SetCreateOption(v string) *CreateStackRequest {
	s.CreateOption = &v
	return s
}

func (s *CreateStackRequest) SetTemplateId(v string) *CreateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateVersion(v string) *CreateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStackRequest) SetParameters(v []*CreateStackRequestParameters) *CreateStackRequest {
	s.Parameters = v
	return s
}

func (s *CreateStackRequest) SetNotificationURLs(v []*string) *CreateStackRequest {
	s.NotificationURLs = v
	return s
}

type CreateStackRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackRequestParameters) SetParameterKey(v string) *CreateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackRequestParameters) SetParameterValue(v string) *CreateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateStackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackId   *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CreateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackResponseBody) SetRequestId(v string) *CreateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackResponseBody) SetStackId(v string) *CreateStackResponseBody {
	s.StackId = &v
	return s
}

type CreateStackResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStackResponse) GoString() string {
	return s.String()
}

func (s *CreateStackResponse) SetHeaders(v map[string]*string) *CreateStackResponse {
	s.Headers = v
	return s
}

func (s *CreateStackResponse) SetBody(v *CreateStackResponseBody) *CreateStackResponse {
	s.Body = v
	return s
}

type CreateStackGroupRequest struct {
	RegionId               *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName         *string                              `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	Description            *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateBody           *string                              `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateURL            *string                              `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	ClientToken            *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AdministrationRoleName *string                              `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	ExecutionRoleName      *string                              `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	TemplateId             *string                              `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion        *string                              `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters             []*CreateStackGroupRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s CreateStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequest) SetRegionId(v string) *CreateStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackGroupRequest) SetStackGroupName(v string) *CreateStackGroupRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackGroupRequest) SetDescription(v string) *CreateStackGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateBody(v string) *CreateStackGroupRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateURL(v string) *CreateStackGroupRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateStackGroupRequest) SetClientToken(v string) *CreateStackGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackGroupRequest) SetAdministrationRoleName(v string) *CreateStackGroupRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *CreateStackGroupRequest) SetExecutionRoleName(v string) *CreateStackGroupRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateId(v string) *CreateStackGroupRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateVersion(v string) *CreateStackGroupRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStackGroupRequest) SetParameters(v []*CreateStackGroupRequestParameters) *CreateStackGroupRequest {
	s.Parameters = v
	return s
}

type CreateStackGroupRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackGroupRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestParameters) SetParameterKey(v string) *CreateStackGroupRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackGroupRequestParameters) SetParameterValue(v string) *CreateStackGroupRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateStackGroupResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
}

func (s CreateStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackGroupResponseBody) SetRequestId(v string) *CreateStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackGroupResponseBody) SetStackGroupId(v string) *CreateStackGroupResponseBody {
	s.StackGroupId = &v
	return s
}

type CreateStackGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateStackGroupResponse) SetHeaders(v map[string]*string) *CreateStackGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateStackGroupResponse) SetBody(v *CreateStackGroupResponseBody) *CreateStackGroupResponse {
	s.Body = v
	return s
}

type CreateStackInstancesRequest struct {
	RegionId             *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName       *string                                          `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountIds           map[string]interface{}                           `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIds            map[string]interface{}                           `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ClientToken          *string                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription *string                                          `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferences map[string]interface{}                           `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	TimeoutInMinutes     *int64                                           `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	DisableRollback      *bool                                            `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	ParameterOverrides   []*CreateStackInstancesRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
}

func (s CreateStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequest) SetRegionId(v string) *CreateStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackInstancesRequest) SetStackGroupName(v string) *CreateStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackInstancesRequest) SetAccountIds(v map[string]interface{}) *CreateStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *CreateStackInstancesRequest) SetRegionIds(v map[string]interface{}) *CreateStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *CreateStackInstancesRequest) SetClientToken(v string) *CreateStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackInstancesRequest) SetOperationDescription(v string) *CreateStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *CreateStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *CreateStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *CreateStackInstancesRequest) SetTimeoutInMinutes(v int64) *CreateStackInstancesRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateStackInstancesRequest) SetDisableRollback(v bool) *CreateStackInstancesRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackInstancesRequest) SetParameterOverrides(v []*CreateStackInstancesRequestParameterOverrides) *CreateStackInstancesRequest {
	s.ParameterOverrides = v
	return s
}

type CreateStackInstancesRequestParameterOverrides struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackInstancesRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequestParameterOverrides) SetParameterKey(v string) *CreateStackInstancesRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackInstancesRequestParameterOverrides) SetParameterValue(v string) *CreateStackInstancesRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type CreateStackInstancesShrinkRequest struct {
	RegionId                   *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName             *string                                                `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountIdsShrink           *string                                                `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIdsShrink            *string                                                `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ClientToken                *string                                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription       *string                                                `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferencesShrink *string                                                `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	TimeoutInMinutes           *int64                                                 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	DisableRollback            *bool                                                  `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	ParameterOverrides         []*CreateStackInstancesShrinkRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
}

func (s CreateStackInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesShrinkRequest) SetRegionId(v string) *CreateStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetStackGroupName(v string) *CreateStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetClientToken(v string) *CreateStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetOperationDescription(v string) *CreateStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *CreateStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetTimeoutInMinutes(v int64) *CreateStackInstancesShrinkRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetDisableRollback(v bool) *CreateStackInstancesShrinkRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetParameterOverrides(v []*CreateStackInstancesShrinkRequestParameterOverrides) *CreateStackInstancesShrinkRequest {
	s.ParameterOverrides = v
	return s
}

type CreateStackInstancesShrinkRequestParameterOverrides struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackInstancesShrinkRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesShrinkRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) SetParameterKey(v string) *CreateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) SetParameterValue(v string) *CreateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type CreateStackInstancesResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s CreateStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesResponseBody) SetRequestId(v string) *CreateStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackInstancesResponseBody) SetOperationId(v string) *CreateStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

type CreateStackInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesResponse) SetHeaders(v map[string]*string) *CreateStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateStackInstancesResponse) SetBody(v *CreateStackInstancesResponseBody) *CreateStackInstancesResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	TemplateURL  *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetTemplateURL(v string) *CreateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateBody(v string) *CreateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

type CreateTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteChangeSetRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
}

func (s DeleteChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetRequest) SetRegionId(v string) *DeleteChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteChangeSetRequest) SetChangeSetId(v string) *DeleteChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

type DeleteChangeSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetResponseBody) SetRequestId(v string) *DeleteChangeSetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChangeSetResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetResponse) SetHeaders(v map[string]*string) *DeleteChangeSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteChangeSetResponse) SetBody(v *DeleteChangeSetResponseBody) *DeleteChangeSetResponse {
	s.Body = v
	return s
}

type DeleteStackRequest struct {
	StackId            *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RetainAllResources *bool     `json:"RetainAllResources,omitempty" xml:"RetainAllResources,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RamRoleName        *string   `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RetainResources    []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
}

func (s DeleteStackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackRequest) SetStackId(v string) *DeleteStackRequest {
	s.StackId = &v
	return s
}

func (s *DeleteStackRequest) SetRetainAllResources(v bool) *DeleteStackRequest {
	s.RetainAllResources = &v
	return s
}

func (s *DeleteStackRequest) SetRegionId(v string) *DeleteStackRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackRequest) SetRamRoleName(v string) *DeleteStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *DeleteStackRequest) SetRetainResources(v []*string) *DeleteStackRequest {
	s.RetainResources = v
	return s
}

type DeleteStackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackResponseBody) SetRequestId(v string) *DeleteStackResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStackResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackResponse) SetHeaders(v map[string]*string) *DeleteStackResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackResponse) SetBody(v *DeleteStackResponseBody) *DeleteStackResponse {
	s.Body = v
	return s
}

type DeleteStackGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DeleteStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupRequest) SetRegionId(v string) *DeleteStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackGroupRequest) SetStackGroupName(v string) *DeleteStackGroupRequest {
	s.StackGroupName = &v
	return s
}

type DeleteStackGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupResponseBody) SetRequestId(v string) *DeleteStackGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStackGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupResponse) SetHeaders(v map[string]*string) *DeleteStackGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackGroupResponse) SetBody(v *DeleteStackGroupResponseBody) *DeleteStackGroupResponse {
	s.Body = v
	return s
}

type DeleteStackInstancesRequest struct {
	RegionId             *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName       *string                `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountIds           map[string]interface{} `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIds            map[string]interface{} `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	RetainStacks         *bool                  `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	ClientToken          *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription *string                `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
}

func (s DeleteStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesRequest) SetRegionId(v string) *DeleteStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetStackGroupName(v string) *DeleteStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetAccountIds(v map[string]interface{}) *DeleteStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *DeleteStackInstancesRequest) SetRegionIds(v map[string]interface{}) *DeleteStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *DeleteStackInstancesRequest) SetRetainStacks(v bool) *DeleteStackInstancesRequest {
	s.RetainStacks = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetClientToken(v string) *DeleteStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetOperationDescription(v string) *DeleteStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *DeleteStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

type DeleteStackInstancesShrinkRequest struct {
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName             *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountIdsShrink           *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIdsShrink            *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	RetainStacks               *bool   `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription       *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
}

func (s DeleteStackInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesShrinkRequest) SetRegionId(v string) *DeleteStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetStackGroupName(v string) *DeleteStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetRetainStacks(v bool) *DeleteStackInstancesShrinkRequest {
	s.RetainStacks = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetClientToken(v string) *DeleteStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetOperationDescription(v string) *DeleteStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

type DeleteStackInstancesResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s DeleteStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesResponseBody) SetRequestId(v string) *DeleteStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStackInstancesResponseBody) SetOperationId(v string) *DeleteStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

type DeleteStackInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesResponse) SetHeaders(v map[string]*string) *DeleteStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackInstancesResponse) SetBody(v *DeleteStackInstancesResponseBody) *DeleteStackInstancesResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetTemplateId(v string) *DeleteTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
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

type DetectStackDriftRequest struct {
	StackId           *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
}

func (s DetectStackDriftRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackDriftRequest) SetStackId(v string) *DetectStackDriftRequest {
	s.StackId = &v
	return s
}

func (s *DetectStackDriftRequest) SetRegionId(v string) *DetectStackDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackDriftRequest) SetClientToken(v string) *DetectStackDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackDriftRequest) SetLogicalResourceId(v []*string) *DetectStackDriftRequest {
	s.LogicalResourceId = v
	return s
}

type DetectStackDriftResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
}

func (s DetectStackDriftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectStackDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackDriftResponseBody) SetRequestId(v string) *DetectStackDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackDriftResponseBody) SetDriftDetectionId(v string) *DetectStackDriftResponseBody {
	s.DriftDetectionId = &v
	return s
}

type DetectStackDriftResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectStackDriftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectStackDriftResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectStackDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackDriftResponse) SetHeaders(v map[string]*string) *DetectStackDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackDriftResponse) SetBody(v *DetectStackDriftResponseBody) *DetectStackDriftResponse {
	s.Body = v
	return s
}

type DetectStackGroupDriftRequest struct {
	ClientToken          *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId             *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName       *string                `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
}

func (s DetectStackGroupDriftRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftRequest) SetClientToken(v string) *DetectStackGroupDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetRegionId(v string) *DetectStackGroupDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetStackGroupName(v string) *DetectStackGroupDriftRequest {
	s.StackGroupName = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetOperationPreferences(v map[string]interface{}) *DetectStackGroupDriftRequest {
	s.OperationPreferences = v
	return s
}

type DetectStackGroupDriftShrinkRequest struct {
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName             *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
}

func (s DetectStackGroupDriftShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftShrinkRequest) SetClientToken(v string) *DetectStackGroupDriftShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackGroupDriftShrinkRequest) SetRegionId(v string) *DetectStackGroupDriftShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackGroupDriftShrinkRequest) SetStackGroupName(v string) *DetectStackGroupDriftShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *DetectStackGroupDriftShrinkRequest) SetOperationPreferencesShrink(v string) *DetectStackGroupDriftShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

type DetectStackGroupDriftResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s DetectStackGroupDriftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftResponseBody) SetRequestId(v string) *DetectStackGroupDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackGroupDriftResponseBody) SetOperationId(v string) *DetectStackGroupDriftResponseBody {
	s.OperationId = &v
	return s
}

type DetectStackGroupDriftResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectStackGroupDriftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectStackGroupDriftResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftResponse) SetHeaders(v map[string]*string) *DetectStackGroupDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackGroupDriftResponse) SetBody(v *DetectStackGroupDriftResponseBody) *DetectStackGroupDriftResponse {
	s.Body = v
	return s
}

type DetectStackResourceDriftRequest struct {
	StackId           *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
}

func (s DetectStackResourceDriftRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftRequest) SetStackId(v string) *DetectStackResourceDriftRequest {
	s.StackId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetClientToken(v string) *DetectStackResourceDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetRegionId(v string) *DetectStackResourceDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetLogicalResourceId(v string) *DetectStackResourceDriftRequest {
	s.LogicalResourceId = &v
	return s
}

type DetectStackResourceDriftResponseBody struct {
	LogicalResourceId   *string                                                    `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	ResourceDriftStatus *string                                                    `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	PropertyDifferences []*DetectStackResourceDriftResponseBodyPropertyDifferences `json:"PropertyDifferences,omitempty" xml:"PropertyDifferences,omitempty" type:"Repeated"`
	RequestId           *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PhysicalResourceId  *string                                                    `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	ExpectedProperties  *string                                                    `json:"ExpectedProperties,omitempty" xml:"ExpectedProperties,omitempty"`
	DriftDetectionTime  *string                                                    `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	ResourceType        *string                                                    `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ActualProperties    *string                                                    `json:"ActualProperties,omitempty" xml:"ActualProperties,omitempty"`
	StackId             *string                                                    `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DetectStackResourceDriftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponseBody) SetLogicalResourceId(v string) *DetectStackResourceDriftResponseBody {
	s.LogicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetResourceDriftStatus(v string) *DetectStackResourceDriftResponseBody {
	s.ResourceDriftStatus = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetPropertyDifferences(v []*DetectStackResourceDriftResponseBodyPropertyDifferences) *DetectStackResourceDriftResponseBody {
	s.PropertyDifferences = v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetRequestId(v string) *DetectStackResourceDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetPhysicalResourceId(v string) *DetectStackResourceDriftResponseBody {
	s.PhysicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetExpectedProperties(v string) *DetectStackResourceDriftResponseBody {
	s.ExpectedProperties = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetDriftDetectionTime(v string) *DetectStackResourceDriftResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetResourceType(v string) *DetectStackResourceDriftResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetActualProperties(v string) *DetectStackResourceDriftResponseBody {
	s.ActualProperties = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetStackId(v string) *DetectStackResourceDriftResponseBody {
	s.StackId = &v
	return s
}

type DetectStackResourceDriftResponseBodyPropertyDifferences struct {
	ActualValue    *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	DifferenceType *string `json:"DifferenceType,omitempty" xml:"DifferenceType,omitempty"`
	PropertyPath   *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
	ExpectedValue  *string `json:"ExpectedValue,omitempty" xml:"ExpectedValue,omitempty"`
}

func (s DetectStackResourceDriftResponseBodyPropertyDifferences) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftResponseBodyPropertyDifferences) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetActualValue(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.ActualValue = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetDifferenceType(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.DifferenceType = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetPropertyPath(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.PropertyPath = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetExpectedValue(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.ExpectedValue = &v
	return s
}

type DetectStackResourceDriftResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectStackResourceDriftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectStackResourceDriftResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponse) SetHeaders(v map[string]*string) *DetectStackResourceDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackResourceDriftResponse) SetBody(v *DetectStackResourceDriftResponseBody) *DetectStackResourceDriftResponse {
	s.Body = v
	return s
}

type ExecuteChangeSetRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ExecuteChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteChangeSetRequest) GoString() string {
	return s.String()
}

func (s *ExecuteChangeSetRequest) SetRegionId(v string) *ExecuteChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *ExecuteChangeSetRequest) SetChangeSetId(v string) *ExecuteChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

func (s *ExecuteChangeSetRequest) SetClientToken(v string) *ExecuteChangeSetRequest {
	s.ClientToken = &v
	return s
}

type ExecuteChangeSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteChangeSetResponseBody) SetRequestId(v string) *ExecuteChangeSetResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteChangeSetResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteChangeSetResponse) GoString() string {
	return s.String()
}

func (s *ExecuteChangeSetResponse) SetHeaders(v map[string]*string) *ExecuteChangeSetResponse {
	s.Headers = v
	return s
}

func (s *ExecuteChangeSetResponse) SetBody(v *ExecuteChangeSetResponseBody) *ExecuteChangeSetResponse {
	s.Body = v
	return s
}

type GenerateTemplatePolicyRequest struct {
	TemplateURL     *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	TemplateBody    *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GenerateTemplatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyRequest) SetTemplateURL(v string) *GenerateTemplatePolicyRequest {
	s.TemplateURL = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateBody(v string) *GenerateTemplatePolicyRequest {
	s.TemplateBody = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateId(v string) *GenerateTemplatePolicyRequest {
	s.TemplateId = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateVersion(v string) *GenerateTemplatePolicyRequest {
	s.TemplateVersion = &v
	return s
}

type GenerateTemplatePolicyResponseBody struct {
	Policy    *GenerateTemplatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTemplatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBody) SetPolicy(v *GenerateTemplatePolicyResponseBodyPolicy) *GenerateTemplatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) SetRequestId(v string) *GenerateTemplatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateTemplatePolicyResponseBodyPolicy struct {
	Version   *string                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	Statement []*GenerateTemplatePolicyResponseBodyPolicyStatement `json:"Statement,omitempty" xml:"Statement,omitempty" type:"Repeated"`
}

func (s GenerateTemplatePolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) SetVersion(v string) *GenerateTemplatePolicyResponseBodyPolicy {
	s.Version = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) SetStatement(v []*GenerateTemplatePolicyResponseBodyPolicyStatement) *GenerateTemplatePolicyResponseBodyPolicy {
	s.Statement = v
	return s
}

type GenerateTemplatePolicyResponseBodyPolicyStatement struct {
	Effect   *string   `json:"Effect,omitempty" xml:"Effect,omitempty"`
	Resource *string   `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Action   []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
}

func (s GenerateTemplatePolicyResponseBodyPolicyStatement) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicyStatement) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetEffect(v string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Effect = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetResource(v string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Resource = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetAction(v []*string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Action = v
	return s
}

type GenerateTemplatePolicyResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateTemplatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateTemplatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponse) SetHeaders(v map[string]*string) *GenerateTemplatePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateTemplatePolicyResponse) SetBody(v *GenerateTemplatePolicyResponseBody) *GenerateTemplatePolicyResponse {
	s.Body = v
	return s
}

type GetChangeSetRequest struct {
	ShowTemplate *bool   `json:"ShowTemplate,omitempty" xml:"ShowTemplate,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChangeSetId  *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
}

func (s GetChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetRequest) GoString() string {
	return s.String()
}

func (s *GetChangeSetRequest) SetShowTemplate(v bool) *GetChangeSetRequest {
	s.ShowTemplate = &v
	return s
}

func (s *GetChangeSetRequest) SetRegionId(v string) *GetChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *GetChangeSetRequest) SetChangeSetId(v string) *GetChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

type GetChangeSetResponseBody struct {
	Status           *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Changes          []map[string]interface{}              `json:"Changes,omitempty" xml:"Changes,omitempty" type:"Repeated"`
	Description      *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Parameters       []*GetChangeSetResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusReason     *string                               `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	CreateTime       *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateBody     *string                               `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	ChangeSetName    *string                               `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	ChangeSetId      *string                               `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	ExecutionStatus  *string                               `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	ChangeSetType    *string                               `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	RegionId         *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DisableRollback  *bool                                 `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	StackName        *string                               `json:"StackName,omitempty" xml:"StackName,omitempty"`
	TimeoutInMinutes *int32                                `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	StackId          *string                               `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBody) SetStatus(v string) *GetChangeSetResponseBody {
	s.Status = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChanges(v []map[string]interface{}) *GetChangeSetResponseBody {
	s.Changes = v
	return s
}

func (s *GetChangeSetResponseBody) SetDescription(v string) *GetChangeSetResponseBody {
	s.Description = &v
	return s
}

func (s *GetChangeSetResponseBody) SetParameters(v []*GetChangeSetResponseBodyParameters) *GetChangeSetResponseBody {
	s.Parameters = v
	return s
}

func (s *GetChangeSetResponseBody) SetRequestId(v string) *GetChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStatusReason(v string) *GetChangeSetResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetChangeSetResponseBody) SetCreateTime(v string) *GetChangeSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChangeSetResponseBody) SetTemplateBody(v string) *GetChangeSetResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetName(v string) *GetChangeSetResponseBody {
	s.ChangeSetName = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetId(v string) *GetChangeSetResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetExecutionStatus(v string) *GetChangeSetResponseBody {
	s.ExecutionStatus = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetType(v string) *GetChangeSetResponseBody {
	s.ChangeSetType = &v
	return s
}

func (s *GetChangeSetResponseBody) SetRegionId(v string) *GetChangeSetResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetDisableRollback(v bool) *GetChangeSetResponseBody {
	s.DisableRollback = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStackName(v string) *GetChangeSetResponseBody {
	s.StackName = &v
	return s
}

func (s *GetChangeSetResponseBody) SetTimeoutInMinutes(v int32) *GetChangeSetResponseBody {
	s.TimeoutInMinutes = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStackId(v string) *GetChangeSetResponseBody {
	s.StackId = &v
	return s
}

type GetChangeSetResponseBodyParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetChangeSetResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyParameters) SetParameterKey(v string) *GetChangeSetResponseBodyParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetChangeSetResponseBodyParameters) SetParameterValue(v string) *GetChangeSetResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

type GetChangeSetResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponse) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponse) SetHeaders(v map[string]*string) *GetChangeSetResponse {
	s.Headers = v
	return s
}

func (s *GetChangeSetResponse) SetBody(v *GetChangeSetResponseBody) *GetChangeSetResponse {
	s.Body = v
	return s
}

type GetResourceTypeRequest struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeRequest) SetResourceType(v string) *GetResourceTypeRequest {
	s.ResourceType = &v
	return s
}

type GetResourceTypeResponseBody struct {
	RequestId             *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Attributes            map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	ResourceType          *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Properties            map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	SupportDriftDetection *bool                  `json:"SupportDriftDetection,omitempty" xml:"SupportDriftDetection,omitempty"`
}

func (s GetResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBody) SetRequestId(v string) *GetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetAttributes(v map[string]interface{}) *GetResourceTypeResponseBody {
	s.Attributes = v
	return s
}

func (s *GetResourceTypeResponseBody) SetResourceType(v string) *GetResourceTypeResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetProperties(v map[string]interface{}) *GetResourceTypeResponseBody {
	s.Properties = v
	return s
}

func (s *GetResourceTypeResponseBody) SetSupportDriftDetection(v bool) *GetResourceTypeResponseBody {
	s.SupportDriftDetection = &v
	return s
}

type GetResourceTypeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponse) SetHeaders(v map[string]*string) *GetResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeResponse) SetBody(v *GetResourceTypeResponseBody) *GetResourceTypeResponse {
	s.Body = v
	return s
}

type GetResourceTypeTemplateRequest struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceTypeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateRequest) SetResourceType(v string) *GetResourceTypeTemplateRequest {
	s.ResourceType = &v
	return s
}

type GetResourceTypeTemplateResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateBody map[string]interface{} `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
}

func (s GetResourceTypeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateResponseBody) SetRequestId(v string) *GetResourceTypeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeTemplateResponseBody) SetTemplateBody(v map[string]interface{}) *GetResourceTypeTemplateResponseBody {
	s.TemplateBody = v
	return s
}

type GetResourceTypeTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResourceTypeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceTypeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateResponse) SetHeaders(v map[string]*string) *GetResourceTypeTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeTemplateResponse) SetBody(v *GetResourceTypeTemplateResponseBody) *GetResourceTypeTemplateResponse {
	s.Body = v
	return s
}

type GetStackRequest struct {
	StackId     *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s GetStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) SetStackId(v string) *GetStackRequest {
	s.StackId = &v
	return s
}

func (s *GetStackRequest) SetRegionId(v string) *GetStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackRequest) SetClientToken(v string) *GetStackRequest {
	s.ClientToken = &v
	return s
}

type GetStackResponseBody struct {
	Status              *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Description         *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Parameters          []*GetStackResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId           *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusReason        *string                           `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	ParentStackId       *string                           `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	CreateTime          *string                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeletionProtection  *string                           `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	RootStackId         *string                           `json:"RootStackId,omitempty" xml:"RootStackId,omitempty"`
	TemplateDescription *string                           `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	StackType           *string                           `json:"StackType,omitempty" xml:"StackType,omitempty"`
	RamRoleName         *string                           `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	UpdateTime          *string                           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Outputs             []map[string]interface{}          `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	DriftDetectionTime  *string                           `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	RegionId            *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackDriftStatus    *string                           `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	NotificationURLs    []*string                         `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	DisableRollback     *bool                             `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	StackName           *string                           `json:"StackName,omitempty" xml:"StackName,omitempty"`
	TimeoutInMinutes    *int32                            `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	StackId             *string                           `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) SetStatus(v string) *GetStackResponseBody {
	s.Status = &v
	return s
}

func (s *GetStackResponseBody) SetDescription(v string) *GetStackResponseBody {
	s.Description = &v
	return s
}

func (s *GetStackResponseBody) SetParameters(v []*GetStackResponseBodyParameters) *GetStackResponseBody {
	s.Parameters = v
	return s
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetStatusReason(v string) *GetStackResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetStackResponseBody) SetParentStackId(v string) *GetStackResponseBody {
	s.ParentStackId = &v
	return s
}

func (s *GetStackResponseBody) SetCreateTime(v string) *GetStackResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStackResponseBody) SetDeletionProtection(v string) *GetStackResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetStackResponseBody) SetRootStackId(v string) *GetStackResponseBody {
	s.RootStackId = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateDescription(v string) *GetStackResponseBody {
	s.TemplateDescription = &v
	return s
}

func (s *GetStackResponseBody) SetStackType(v string) *GetStackResponseBody {
	s.StackType = &v
	return s
}

func (s *GetStackResponseBody) SetRamRoleName(v string) *GetStackResponseBody {
	s.RamRoleName = &v
	return s
}

func (s *GetStackResponseBody) SetUpdateTime(v string) *GetStackResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetStackResponseBody) SetOutputs(v []map[string]interface{}) *GetStackResponseBody {
	s.Outputs = v
	return s
}

func (s *GetStackResponseBody) SetDriftDetectionTime(v string) *GetStackResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackResponseBody) SetRegionId(v string) *GetStackResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetStackResponseBody) SetStackDriftStatus(v string) *GetStackResponseBody {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackResponseBody) SetNotificationURLs(v []*string) *GetStackResponseBody {
	s.NotificationURLs = v
	return s
}

func (s *GetStackResponseBody) SetDisableRollback(v bool) *GetStackResponseBody {
	s.DisableRollback = &v
	return s
}

func (s *GetStackResponseBody) SetStackName(v string) *GetStackResponseBody {
	s.StackName = &v
	return s
}

func (s *GetStackResponseBody) SetTimeoutInMinutes(v int32) *GetStackResponseBody {
	s.TimeoutInMinutes = &v
	return s
}

func (s *GetStackResponseBody) SetStackId(v string) *GetStackResponseBody {
	s.StackId = &v
	return s
}

type GetStackResponseBodyParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyParameters) SetParameterKey(v string) *GetStackResponseBodyParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetStackResponseBodyParameters) SetParameterValue(v string) *GetStackResponseBodyParameters {
	s.ParameterValue = &v
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

type GetStackDriftDetectionStatusRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
}

func (s GetStackDriftDetectionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackDriftDetectionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusRequest) SetRegionId(v string) *GetStackDriftDetectionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackDriftDetectionStatusRequest) SetDriftDetectionId(v string) *GetStackDriftDetectionStatusRequest {
	s.DriftDetectionId = &v
	return s
}

type GetStackDriftDetectionStatusResponseBody struct {
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DriftDetectionTime         *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	DriftDetectionStatusReason *string `json:"DriftDetectionStatusReason,omitempty" xml:"DriftDetectionStatusReason,omitempty"`
	DriftedStackResourceCount  *int32  `json:"DriftedStackResourceCount,omitempty" xml:"DriftedStackResourceCount,omitempty"`
	StackDriftStatus           *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	DriftDetectionStatus       *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	StackId                    *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	DriftDetectionId           *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
}

func (s GetStackDriftDetectionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackDriftDetectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusResponseBody) SetRequestId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionTime(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionStatusReason(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionStatusReason = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftedStackResourceCount(v int32) *GetStackDriftDetectionStatusResponseBody {
	s.DriftedStackResourceCount = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetStackDriftStatus(v string) *GetStackDriftDetectionStatusResponseBody {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionStatus(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetStackId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.StackId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionId = &v
	return s
}

type GetStackDriftDetectionStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackDriftDetectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackDriftDetectionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackDriftDetectionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusResponse) SetHeaders(v map[string]*string) *GetStackDriftDetectionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetStackDriftDetectionStatusResponse) SetBody(v *GetStackDriftDetectionStatusResponseBody) *GetStackDriftDetectionStatusResponse {
	s.Body = v
	return s
}

type GetStackGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s GetStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupRequest) GoString() string {
	return s.String()
}

func (s *GetStackGroupRequest) SetRegionId(v string) *GetStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackGroupRequest) SetStackGroupName(v string) *GetStackGroupRequest {
	s.StackGroupName = &v
	return s
}

type GetStackGroupResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackGroup *GetStackGroupResponseBodyStackGroup `json:"StackGroup,omitempty" xml:"StackGroup,omitempty" type:"Struct"`
}

func (s GetStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBody) SetRequestId(v string) *GetStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackGroupResponseBody) SetStackGroup(v *GetStackGroupResponseBodyStackGroup) *GetStackGroupResponseBody {
	s.StackGroup = v
	return s
}

type GetStackGroupResponseBodyStackGroup struct {
	StackGroupId                   *string                                                            `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	Status                         *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	AdministrationRoleName         *string                                                            `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	Parameters                     []*GetStackGroupResponseBodyStackGroupParameters                   `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	Description                    *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	StackGroupName                 *string                                                            `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	ExecutionRoleName              *string                                                            `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	TemplateBody                   *string                                                            `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	StackGroupDriftDetectionDetail *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail `json:"StackGroupDriftDetectionDetail,omitempty" xml:"StackGroupDriftDetectionDetail,omitempty" type:"Struct"`
}

func (s GetStackGroupResponseBodyStackGroup) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroup) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupId(v string) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStatus(v string) *GetStackGroupResponseBodyStackGroup {
	s.Status = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetAdministrationRoleName(v string) *GetStackGroupResponseBodyStackGroup {
	s.AdministrationRoleName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetParameters(v []*GetStackGroupResponseBodyStackGroupParameters) *GetStackGroupResponseBodyStackGroup {
	s.Parameters = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetDescription(v string) *GetStackGroupResponseBodyStackGroup {
	s.Description = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupName(v string) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetExecutionRoleName(v string) *GetStackGroupResponseBodyStackGroup {
	s.ExecutionRoleName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetTemplateBody(v string) *GetStackGroupResponseBodyStackGroup {
	s.TemplateBody = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupDriftDetectionDetail(v *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupDriftDetectionDetail = v
	return s
}

type GetStackGroupResponseBodyStackGroupParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupParameters) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupParameters) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupParameters) SetParameterKey(v string) *GetStackGroupResponseBodyStackGroupParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupParameters) SetParameterValue(v string) *GetStackGroupResponseBodyStackGroupParameters {
	s.ParameterValue = &v
	return s
}

type GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail struct {
	DriftDetectionTime            *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	TotalStackInstancesCount      *int32  `json:"TotalStackInstancesCount,omitempty" xml:"TotalStackInstancesCount,omitempty"`
	FailedStackInstancesCount     *int32  `json:"FailedStackInstancesCount,omitempty" xml:"FailedStackInstancesCount,omitempty"`
	DriftDetectionStatus          *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	StackGroupDriftStatus         *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	InProgressStackInstancesCount *int32  `json:"InProgressStackInstancesCount,omitempty" xml:"InProgressStackInstancesCount,omitempty"`
	InSyncStackInstancesCount     *int32  `json:"InSyncStackInstancesCount,omitempty" xml:"InSyncStackInstancesCount,omitempty"`
	CancelledStackInstancesCount  *int32  `json:"CancelledStackInstancesCount,omitempty" xml:"CancelledStackInstancesCount,omitempty"`
	DriftedStackInstancesCount    *int32  `json:"DriftedStackInstancesCount,omitempty" xml:"DriftedStackInstancesCount,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftDetectionTime(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetTotalStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.TotalStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetFailedStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.FailedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftDetectionStatus(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetStackGroupDriftStatus(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetInProgressStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.InProgressStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetInSyncStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.InSyncStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetCancelledStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.CancelledStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftedStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftedStackInstancesCount = &v
	return s
}

type GetStackGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponse) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponse) SetHeaders(v map[string]*string) *GetStackGroupResponse {
	s.Headers = v
	return s
}

func (s *GetStackGroupResponse) SetBody(v *GetStackGroupResponseBody) *GetStackGroupResponse {
	s.Body = v
	return s
}

type GetStackGroupOperationRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s GetStackGroupOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationRequest) SetRegionId(v string) *GetStackGroupOperationRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackGroupOperationRequest) SetOperationId(v string) *GetStackGroupOperationRequest {
	s.OperationId = &v
	return s
}

type GetStackGroupOperationResponseBody struct {
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackGroupOperation *GetStackGroupOperationResponseBodyStackGroupOperation `json:"StackGroupOperation,omitempty" xml:"StackGroupOperation,omitempty" type:"Struct"`
}

func (s GetStackGroupOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBody) SetRequestId(v string) *GetStackGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackGroupOperationResponseBody) SetStackGroupOperation(v *GetStackGroupOperationResponseBodyStackGroupOperation) *GetStackGroupOperationResponseBody {
	s.StackGroupOperation = v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperation struct {
	Status                         *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	StackGroupId                   *string                                                                              `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	Action                         *string                                                                              `json:"Action,omitempty" xml:"Action,omitempty"`
	CreateTime                     *string                                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RetainStacks                   *bool                                                                                `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	StackGroupName                 *string                                                                              `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	OperationId                    *string                                                                              `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	OperationDescription           *string                                                                              `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	StackGroupDriftDetectionDetail *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail `json:"StackGroupDriftDetectionDetail,omitempty" xml:"StackGroupDriftDetectionDetail,omitempty" type:"Struct"`
	OperationPreferences           *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences           `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty" type:"Struct"`
	EndTime                        *string                                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutionRoleName              *string                                                                              `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	AdministratorRoleName          *string                                                                              `json:"AdministratorRoleName,omitempty" xml:"AdministratorRoleName,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperation) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperation) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.Status = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupId(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetAction(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.Action = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetCreateTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.CreateTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetRetainStacks(v bool) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.RetainStacks = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationId(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationId = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationDescription(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationDescription = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupDriftDetectionDetail(v *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupDriftDetectionDetail = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationPreferences(v *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationPreferences = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetEndTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.EndTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetExecutionRoleName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.ExecutionRoleName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetAdministratorRoleName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.AdministratorRoleName = &v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail struct {
	DriftDetectionTime            *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	TotalStackInstancesCount      *int32  `json:"TotalStackInstancesCount,omitempty" xml:"TotalStackInstancesCount,omitempty"`
	FailedStackInstancesCount     *int32  `json:"FailedStackInstancesCount,omitempty" xml:"FailedStackInstancesCount,omitempty"`
	DriftDetectionStatus          *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	StackGroupDriftStatus         *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	InProgressStackInstancesCount *int32  `json:"InProgressStackInstancesCount,omitempty" xml:"InProgressStackInstancesCount,omitempty"`
	InSyncStackInstancesCount     *int32  `json:"InSyncStackInstancesCount,omitempty" xml:"InSyncStackInstancesCount,omitempty"`
	CancelledStackInstancesCount  *int32  `json:"CancelledStackInstancesCount,omitempty" xml:"CancelledStackInstancesCount,omitempty"`
	DriftedStackInstancesCount    *int32  `json:"DriftedStackInstancesCount,omitempty" xml:"DriftedStackInstancesCount,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftDetectionTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetTotalStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.TotalStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetFailedStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.FailedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftDetectionStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetStackGroupDriftStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetInProgressStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.InProgressStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetInSyncStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.InSyncStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetCancelledStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.CancelledStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftedStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftedStackInstancesCount = &v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences struct {
	MaxConcurrentCount         *int32    `json:"MaxConcurrentCount,omitempty" xml:"MaxConcurrentCount,omitempty"`
	FailureToleranceCount      *int32    `json:"FailureToleranceCount,omitempty" xml:"FailureToleranceCount,omitempty"`
	MaxConcurrentPercentage    *int32    `json:"MaxConcurrentPercentage,omitempty" xml:"MaxConcurrentPercentage,omitempty"`
	RegionIdsOrder             []*string `json:"RegionIdsOrder,omitempty" xml:"RegionIdsOrder,omitempty" type:"Repeated"`
	FailureTolerancePercentage *int32    `json:"FailureTolerancePercentage,omitempty" xml:"FailureTolerancePercentage,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetMaxConcurrentCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.MaxConcurrentCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetFailureToleranceCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.FailureToleranceCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetMaxConcurrentPercentage(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.MaxConcurrentPercentage = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetRegionIdsOrder(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.RegionIdsOrder = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetFailureTolerancePercentage(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.FailureTolerancePercentage = &v
	return s
}

type GetStackGroupOperationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackGroupOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponse) SetHeaders(v map[string]*string) *GetStackGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *GetStackGroupOperationResponse) SetBody(v *GetStackGroupOperationResponseBody) *GetStackGroupOperationResponse {
	s.Body = v
	return s
}

type GetStackInstanceRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName         *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	StackInstanceAccountId *string `json:"StackInstanceAccountId,omitempty" xml:"StackInstanceAccountId,omitempty"`
	StackInstanceRegionId  *string `json:"StackInstanceRegionId,omitempty" xml:"StackInstanceRegionId,omitempty"`
}

func (s GetStackInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetStackInstanceRequest) SetRegionId(v string) *GetStackInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackGroupName(v string) *GetStackInstanceRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackInstanceAccountId(v string) *GetStackInstanceRequest {
	s.StackInstanceAccountId = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackInstanceRegionId(v string) *GetStackInstanceRequest {
	s.StackInstanceRegionId = &v
	return s
}

type GetStackInstanceResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackInstance *GetStackInstanceResponseBodyStackInstance `json:"StackInstance,omitempty" xml:"StackInstance,omitempty" type:"Struct"`
}

func (s GetStackInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBody) SetRequestId(v string) *GetStackInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackInstanceResponseBody) SetStackInstance(v *GetStackInstanceResponseBodyStackInstance) *GetStackInstanceResponseBody {
	s.StackInstance = v
	return s
}

type GetStackInstanceResponseBodyStackInstance struct {
	Status             *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StackGroupId       *string                                                        `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	StackId            *string                                                        `json:"StackId,omitempty" xml:"StackId,omitempty"`
	DriftDetectionTime *string                                                        `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	StackDriftStatus   *string                                                        `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	StatusReason       *string                                                        `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	ParameterOverrides []*GetStackInstanceResponseBodyStackInstanceParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	StackGroupName     *string                                                        `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountId          *string                                                        `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RegionId           *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStackInstanceResponseBodyStackInstance) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponseBodyStackInstance) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStatus(v string) *GetStackInstanceResponseBodyStackInstance {
	s.Status = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackGroupId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackGroupId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetDriftDetectionTime(v string) *GetStackInstanceResponseBodyStackInstance {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackDriftStatus(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStatusReason(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StatusReason = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetParameterOverrides(v []*GetStackInstanceResponseBodyStackInstanceParameterOverrides) *GetStackInstanceResponseBodyStackInstance {
	s.ParameterOverrides = v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackGroupName(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackGroupName = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetAccountId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.AccountId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetRegionId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.RegionId = &v
	return s
}

type GetStackInstanceResponseBodyStackInstanceParameterOverrides struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackInstanceResponseBodyStackInstanceParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponseBodyStackInstanceParameterOverrides) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) SetParameterKey(v string) *GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) SetParameterValue(v string) *GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	s.ParameterValue = &v
	return s
}

type GetStackInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponse) SetHeaders(v map[string]*string) *GetStackInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetStackInstanceResponse) SetBody(v *GetStackInstanceResponseBody) *GetStackInstanceResponse {
	s.Body = v
	return s
}

type GetStackPolicyRequest struct {
	StackId  *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStackPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetStackPolicyRequest) SetStackId(v string) *GetStackPolicyRequest {
	s.StackId = &v
	return s
}

func (s *GetStackPolicyRequest) SetRegionId(v string) *GetStackPolicyRequest {
	s.RegionId = &v
	return s
}

type GetStackPolicyResponseBody struct {
	RequestId       *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackPolicyBody map[string]interface{} `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
}

func (s GetStackPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackPolicyResponseBody) SetRequestId(v string) *GetStackPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackPolicyResponseBody) SetStackPolicyBody(v map[string]interface{}) *GetStackPolicyResponseBody {
	s.StackPolicyBody = v
	return s
}

type GetStackPolicyResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetStackPolicyResponse) SetHeaders(v map[string]*string) *GetStackPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetStackPolicyResponse) SetBody(v *GetStackPolicyResponseBody) *GetStackPolicyResponse {
	s.Body = v
	return s
}

type GetStackResourceRequest struct {
	StackId                *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShowResourceAttributes *bool   `json:"ShowResourceAttributes,omitempty" xml:"ShowResourceAttributes,omitempty"`
	LogicalResourceId      *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
}

func (s GetStackResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceRequest) GoString() string {
	return s.String()
}

func (s *GetStackResourceRequest) SetStackId(v string) *GetStackResourceRequest {
	s.StackId = &v
	return s
}

func (s *GetStackResourceRequest) SetClientToken(v string) *GetStackResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStackResourceRequest) SetRegionId(v string) *GetStackResourceRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackResourceRequest) SetShowResourceAttributes(v bool) *GetStackResourceRequest {
	s.ShowResourceAttributes = &v
	return s
}

func (s *GetStackResourceRequest) SetLogicalResourceId(v string) *GetStackResourceRequest {
	s.LogicalResourceId = &v
	return s
}

type GetStackResourceResponseBody struct {
	Status              *string                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Description         *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId           *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusReason        *string                  `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	PhysicalResourceId  *string                  `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	CreateTime          *string                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Metadata            map[string]interface{}   `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	ResourceType        *string                  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceAttributes  []map[string]interface{} `json:"ResourceAttributes,omitempty" xml:"ResourceAttributes,omitempty" type:"Repeated"`
	LogicalResourceId   *string                  `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	ResourceDriftStatus *string                  `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	UpdateTime          *string                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DriftDetectionTime  *string                  `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	StackName           *string                  `json:"StackName,omitempty" xml:"StackName,omitempty"`
	StackId             *string                  `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponseBody) SetStatus(v string) *GetStackResourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetStackResourceResponseBody) SetDescription(v string) *GetStackResourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetStackResourceResponseBody) SetRequestId(v string) *GetStackResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStatusReason(v string) *GetStackResourceResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetStackResourceResponseBody) SetPhysicalResourceId(v string) *GetStackResourceResponseBody {
	s.PhysicalResourceId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetCreateTime(v string) *GetStackResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetMetadata(v map[string]interface{}) *GetStackResourceResponseBody {
	s.Metadata = v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceType(v string) *GetStackResourceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceAttributes(v []map[string]interface{}) *GetStackResourceResponseBody {
	s.ResourceAttributes = v
	return s
}

func (s *GetStackResourceResponseBody) SetLogicalResourceId(v string) *GetStackResourceResponseBody {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceDriftStatus(v string) *GetStackResourceResponseBody {
	s.ResourceDriftStatus = &v
	return s
}

func (s *GetStackResourceResponseBody) SetUpdateTime(v string) *GetStackResourceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetDriftDetectionTime(v string) *GetStackResourceResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStackName(v string) *GetStackResourceResponseBody {
	s.StackName = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStackId(v string) *GetStackResourceResponseBody {
	s.StackId = &v
	return s
}

type GetStackResourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStackResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceResponse) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponse) SetHeaders(v map[string]*string) *GetStackResourceResponse {
	s.Headers = v
	return s
}

func (s *GetStackResourceResponse) SetBody(v *GetStackResourceResponseBody) *GetStackResourceResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	StackId           *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChangeSetId       *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion   *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateStage     *string `json:"TemplateStage,omitempty" xml:"TemplateStage,omitempty"`
	IncludePermission *string `json:"IncludePermission,omitempty" xml:"IncludePermission,omitempty"`
	StackGroupName    *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetStackId(v string) *GetTemplateRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetChangeSetId(v string) *GetTemplateRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateVersion(v string) *GetTemplateRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateStage(v string) *GetTemplateRequest {
	s.TemplateStage = &v
	return s
}

func (s *GetTemplateRequest) SetIncludePermission(v string) *GetTemplateRequest {
	s.IncludePermission = &v
	return s
}

func (s *GetTemplateRequest) SetStackGroupName(v string) *GetTemplateRequest {
	s.StackGroupName = &v
	return s
}

type GetTemplateResponseBody struct {
	TemplateARN     *string                               `json:"TemplateARN,omitempty" xml:"TemplateARN,omitempty"`
	Description     *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId       *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime      *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StackGroupName  *string                               `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	TemplateVersion *string                               `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateBody    *string                               `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	ChangeSetId     *string                               `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	OwnerId         *string                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	UpdateTime      *string                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Permissions     []*GetTemplateResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	TemplateName    *string                               `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	RegionId        *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateId      *string                               `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	StackId         *string                               `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ShareType       *string                               `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetTemplateARN(v string) *GetTemplateResponseBody {
	s.TemplateARN = &v
	return s
}

func (s *GetTemplateResponseBody) SetDescription(v string) *GetTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetCreateTime(v string) *GetTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetStackGroupName(v string) *GetTemplateResponseBody {
	s.StackGroupName = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateVersion(v string) *GetTemplateResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateBody(v string) *GetTemplateResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateResponseBody) SetChangeSetId(v string) *GetTemplateResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateResponseBody) SetOwnerId(v string) *GetTemplateResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetTemplateResponseBody) SetUpdateTime(v string) *GetTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetPermissions(v []*GetTemplateResponseBodyPermissions) *GetTemplateResponseBody {
	s.Permissions = v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateName(v string) *GetTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBody) SetRegionId(v string) *GetTemplateResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateId(v string) *GetTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBody) SetStackId(v string) *GetTemplateResponseBody {
	s.StackId = &v
	return s
}

func (s *GetTemplateResponseBody) SetShareType(v string) *GetTemplateResponseBody {
	s.ShareType = &v
	return s
}

type GetTemplateResponseBodyPermissions struct {
	VersionOption   *string `json:"VersionOption,omitempty" xml:"VersionOption,omitempty"`
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ShareOption     *string `json:"ShareOption,omitempty" xml:"ShareOption,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyPermissions) SetVersionOption(v string) *GetTemplateResponseBodyPermissions {
	s.VersionOption = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetAccountId(v string) *GetTemplateResponseBodyPermissions {
	s.AccountId = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetShareOption(v string) *GetTemplateResponseBodyPermissions {
	s.ShareOption = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetTemplateVersion(v string) *GetTemplateResponseBodyPermissions {
	s.TemplateVersion = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetTemplateEstimateCostRequest struct {
	TemplateURL     *string                                     `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateBody    *string                                     `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	ClientToken     *string                                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TemplateId      *string                                     `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion *string                                     `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters      []*GetTemplateEstimateCostRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetTemplateEstimateCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostRequest) SetTemplateURL(v string) *GetTemplateEstimateCostRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetRegionId(v string) *GetTemplateEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateBody(v string) *GetTemplateEstimateCostRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetClientToken(v string) *GetTemplateEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateVersion(v string) *GetTemplateEstimateCostRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetParameters(v []*GetTemplateEstimateCostRequestParameters) *GetTemplateEstimateCostRequest {
	s.Parameters = v
	return s
}

type GetTemplateEstimateCostRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateEstimateCostRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostRequestParameters) SetParameterKey(v string) *GetTemplateEstimateCostRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateEstimateCostRequestParameters) SetParameterValue(v string) *GetTemplateEstimateCostRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateEstimateCostResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetTemplateEstimateCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostResponseBody) SetRequestId(v string) *GetTemplateEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetTemplateEstimateCostResponseBody {
	s.Resources = v
	return s
}

type GetTemplateEstimateCostResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateEstimateCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostResponse) SetHeaders(v map[string]*string) *GetTemplateEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateEstimateCostResponse) SetBody(v *GetTemplateEstimateCostResponseBody) *GetTemplateEstimateCostResponse {
	s.Body = v
	return s
}

type GetTemplateSummaryRequest struct {
	StackId         *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateBody    *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateURL     *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	ChangeSetId     *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	StackGroupName  *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s GetTemplateSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryRequest) SetStackId(v string) *GetTemplateSummaryRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateBody(v string) *GetTemplateSummaryRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetRegionId(v string) *GetTemplateSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateId(v string) *GetTemplateSummaryRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateURL(v string) *GetTemplateSummaryRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetChangeSetId(v string) *GetTemplateSummaryRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateVersion(v string) *GetTemplateSummaryRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetStackGroupName(v string) *GetTemplateSummaryRequest {
	s.StackGroupName = &v
	return s
}

type GetTemplateSummaryResponseBody struct {
	ResourceTypes               []*string                                                    `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	Description                 *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Parameters                  []map[string]interface{}                                     `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId                   *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version                     *string                                                      `json:"Version,omitempty" xml:"Version,omitempty"`
	Metadata                    map[string]interface{}                                       `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	ResourceIdentifierSummaries []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries `json:"ResourceIdentifierSummaries,omitempty" xml:"ResourceIdentifierSummaries,omitempty" type:"Repeated"`
}

func (s GetTemplateSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponseBody) SetResourceTypes(v []*string) *GetTemplateSummaryResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetDescription(v string) *GetTemplateSummaryResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetParameters(v []map[string]interface{}) *GetTemplateSummaryResponseBody {
	s.Parameters = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetRequestId(v string) *GetTemplateSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetVersion(v string) *GetTemplateSummaryResponseBody {
	s.Version = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetMetadata(v map[string]interface{}) *GetTemplateSummaryResponseBody {
	s.Metadata = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetResourceIdentifierSummaries(v []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries) *GetTemplateSummaryResponseBody {
	s.ResourceIdentifierSummaries = v
	return s
}

type GetTemplateSummaryResponseBodyResourceIdentifierSummaries struct {
	ResourceType        *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	LogicalResourceIds  []*string `json:"LogicalResourceIds,omitempty" xml:"LogicalResourceIds,omitempty" type:"Repeated"`
	ResourceIdentifiers []*string `json:"ResourceIdentifiers,omitempty" xml:"ResourceIdentifiers,omitempty" type:"Repeated"`
}

func (s GetTemplateSummaryResponseBodyResourceIdentifierSummaries) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryResponseBodyResourceIdentifierSummaries) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetResourceType(v string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetLogicalResourceIds(v []*string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.LogicalResourceIds = v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetResourceIdentifiers(v []*string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.ResourceIdentifiers = v
	return s
}

type GetTemplateSummaryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponse) SetHeaders(v map[string]*string) *GetTemplateSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateSummaryResponse) SetBody(v *GetTemplateSummaryResponseBody) *GetTemplateSummaryResponse {
	s.Body = v
	return s
}

type ListChangeSetsRequest struct {
	StackId         *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PageSize        *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber      *int64    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ChangeSetId     *string   `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	Status          []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	ChangeSetName   []*string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty" type:"Repeated"`
	ExecutionStatus []*string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty" type:"Repeated"`
}

func (s ListChangeSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsRequest) GoString() string {
	return s.String()
}

func (s *ListChangeSetsRequest) SetStackId(v string) *ListChangeSetsRequest {
	s.StackId = &v
	return s
}

func (s *ListChangeSetsRequest) SetPageSize(v int64) *ListChangeSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChangeSetsRequest) SetRegionId(v string) *ListChangeSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListChangeSetsRequest) SetPageNumber(v int64) *ListChangeSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChangeSetsRequest) SetChangeSetId(v string) *ListChangeSetsRequest {
	s.ChangeSetId = &v
	return s
}

func (s *ListChangeSetsRequest) SetStatus(v []*string) *ListChangeSetsRequest {
	s.Status = v
	return s
}

func (s *ListChangeSetsRequest) SetChangeSetName(v []*string) *ListChangeSetsRequest {
	s.ChangeSetName = v
	return s
}

func (s *ListChangeSetsRequest) SetExecutionStatus(v []*string) *ListChangeSetsRequest {
	s.ExecutionStatus = v
	return s
}

type ListChangeSetsResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ChangeSets []*ListChangeSetsResponseBodyChangeSets `json:"ChangeSets,omitempty" xml:"ChangeSets,omitempty" type:"Repeated"`
}

func (s ListChangeSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBody) SetTotalCount(v int32) *ListChangeSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetPageSize(v int32) *ListChangeSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetRequestId(v string) *ListChangeSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetPageNumber(v int32) *ListChangeSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetChangeSets(v []*ListChangeSetsResponseBodyChangeSets) *ListChangeSetsResponseBody {
	s.ChangeSets = v
	return s
}

type ListChangeSetsResponseBodyChangeSets struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StackId         *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ChangeSetName   *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ChangeSetType   *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	StatusReason    *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChangeSetId     *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	StackName       *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListChangeSetsResponseBodyChangeSets) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsResponseBodyChangeSets) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStatus(v string) *ListChangeSetsResponseBodyChangeSets {
	s.Status = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStackId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StackId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetName(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetName = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetDescription(v string) *ListChangeSetsResponseBodyChangeSets {
	s.Description = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetType(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetType = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStatusReason(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StatusReason = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetCreateTime(v string) *ListChangeSetsResponseBodyChangeSets {
	s.CreateTime = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStackName(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StackName = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetExecutionStatus(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ExecutionStatus = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetRegionId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.RegionId = &v
	return s
}

type ListChangeSetsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListChangeSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChangeSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsResponse) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponse) SetHeaders(v map[string]*string) *ListChangeSetsResponse {
	s.Headers = v
	return s
}

func (s *ListChangeSetsResponse) SetBody(v *ListChangeSetsResponseBody) *ListChangeSetsResponse {
	s.Body = v
	return s
}

type ListResourceTypesResponseBody struct {
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*string) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type ListStackEventsRequest struct {
	StackId           *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PageSize          *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber        *int64    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Status            []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	ResourceType      []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
}

func (s ListStackEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsRequest) GoString() string {
	return s.String()
}

func (s *ListStackEventsRequest) SetStackId(v string) *ListStackEventsRequest {
	s.StackId = &v
	return s
}

func (s *ListStackEventsRequest) SetPageSize(v int64) *ListStackEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackEventsRequest) SetRegionId(v string) *ListStackEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackEventsRequest) SetPageNumber(v int64) *ListStackEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackEventsRequest) SetStatus(v []*string) *ListStackEventsRequest {
	s.Status = v
	return s
}

func (s *ListStackEventsRequest) SetResourceType(v []*string) *ListStackEventsRequest {
	s.ResourceType = v
	return s
}

func (s *ListStackEventsRequest) SetLogicalResourceId(v []*string) *ListStackEventsRequest {
	s.LogicalResourceId = v
	return s
}

type ListStackEventsResponseBody struct {
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Events     []*ListStackEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s ListStackEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponseBody) SetTotalCount(v int32) *ListStackEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackEventsResponseBody) SetPageSize(v int32) *ListStackEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackEventsResponseBody) SetRequestId(v string) *ListStackEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackEventsResponseBody) SetPageNumber(v int32) *ListStackEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackEventsResponseBody) SetEvents(v []*ListStackEventsResponseBodyEvents) *ListStackEventsResponseBody {
	s.Events = v
	return s
}

type ListStackEventsResponseBodyEvents struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	EventId            *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	LogicalResourceId  *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	StackId            *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StatusReason       *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StackName          *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
}

func (s ListStackEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponseBodyEvents) SetStatus(v string) *ListStackEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetEventId(v string) *ListStackEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetLogicalResourceId(v string) *ListStackEventsResponseBodyEvents {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStackId(v string) *ListStackEventsResponseBodyEvents {
	s.StackId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetPhysicalResourceId(v string) *ListStackEventsResponseBodyEvents {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetResourceType(v string) *ListStackEventsResponseBodyEvents {
	s.ResourceType = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStatusReason(v string) *ListStackEventsResponseBodyEvents {
	s.StatusReason = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetCreateTime(v string) *ListStackEventsResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStackName(v string) *ListStackEventsResponseBodyEvents {
	s.StackName = &v
	return s
}

type ListStackEventsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsResponse) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponse) SetHeaders(v map[string]*string) *ListStackEventsResponse {
	s.Headers = v
	return s
}

func (s *ListStackEventsResponse) SetBody(v *ListStackEventsResponseBody) *ListStackEventsResponse {
	s.Body = v
	return s
}

type ListStackGroupOperationResultsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListStackGroupOperationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsRequest) SetRegionId(v string) *ListStackGroupOperationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetOperationId(v string) *ListStackGroupOperationResultsRequest {
	s.OperationId = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetPageSize(v int64) *ListStackGroupOperationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetPageNumber(v int64) *ListStackGroupOperationResultsRequest {
	s.PageNumber = &v
	return s
}

type ListStackGroupOperationResultsResponseBody struct {
	TotalCount                 *int32                                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId                  *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize                   *int32                                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber                 *int32                                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StackGroupOperationResults []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults `json:"StackGroupOperationResults,omitempty" xml:"StackGroupOperationResults,omitempty" type:"Repeated"`
}

func (s ListStackGroupOperationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponseBody) SetTotalCount(v int32) *ListStackGroupOperationResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetRequestId(v string) *ListStackGroupOperationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetPageSize(v int32) *ListStackGroupOperationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetPageNumber(v int32) *ListStackGroupOperationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetStackGroupOperationResults(v []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) *ListStackGroupOperationResultsResponseBody {
	s.StackGroupOperationResults = v
	return s
}

type ListStackGroupOperationResultsResponseBodyStackGroupOperationResults struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetStatus(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.Status = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetStatusReason(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.StatusReason = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetAccountId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.AccountId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetRegionId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.RegionId = &v
	return s
}

type ListStackGroupOperationResultsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackGroupOperationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackGroupOperationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponse) SetHeaders(v map[string]*string) *ListStackGroupOperationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupOperationResultsResponse) SetBody(v *ListStackGroupOperationResultsResponseBody) *ListStackGroupOperationResultsResponse {
	s.Body = v
	return s
}

type ListStackGroupOperationsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListStackGroupOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsRequest) SetRegionId(v string) *ListStackGroupOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetStackGroupName(v string) *ListStackGroupOperationsRequest {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetPageSize(v int64) *ListStackGroupOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetPageNumber(v int64) *ListStackGroupOperationsRequest {
	s.PageNumber = &v
	return s
}

type ListStackGroupOperationsResponseBody struct {
	TotalCount           *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize             *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StackGroupOperations []*ListStackGroupOperationsResponseBodyStackGroupOperations `json:"StackGroupOperations,omitempty" xml:"StackGroupOperations,omitempty" type:"Repeated"`
}

func (s ListStackGroupOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponseBody) SetTotalCount(v int32) *ListStackGroupOperationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetRequestId(v string) *ListStackGroupOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetPageSize(v int32) *ListStackGroupOperationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetPageNumber(v int32) *ListStackGroupOperationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetStackGroupOperations(v []*ListStackGroupOperationsResponseBodyStackGroupOperations) *ListStackGroupOperationsResponseBody {
	s.StackGroupOperations = v
	return s
}

type ListStackGroupOperationsResponseBodyStackGroupOperations struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StackGroupId         *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Action               *string `json:"Action,omitempty" xml:"Action,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StackGroupName       *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	OperationId          *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
}

func (s ListStackGroupOperationsResponseBodyStackGroupOperations) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsResponseBodyStackGroupOperations) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStatus(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.Status = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStackGroupId(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.StackGroupId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetEndTime(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.EndTime = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetAction(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.Action = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetCreateTime(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.CreateTime = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStackGroupName(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetOperationId(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.OperationId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetOperationDescription(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.OperationDescription = &v
	return s
}

type ListStackGroupOperationsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackGroupOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackGroupOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponse) SetHeaders(v map[string]*string) *ListStackGroupOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupOperationsResponse) SetBody(v *ListStackGroupOperationsResponseBody) *ListStackGroupOperationsResponse {
	s.Body = v
	return s
}

type ListStackGroupsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListStackGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupsRequest) SetRegionId(v string) *ListStackGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupsRequest) SetStatus(v string) *ListStackGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListStackGroupsRequest) SetPageSize(v int64) *ListStackGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupsRequest) SetPageNumber(v int64) *ListStackGroupsRequest {
	s.PageNumber = &v
	return s
}

type ListStackGroupsResponseBody struct {
	StackGroups []*ListStackGroupsResponseBodyStackGroups `json:"StackGroups,omitempty" xml:"StackGroups,omitempty" type:"Repeated"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListStackGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBody) SetStackGroups(v []*ListStackGroupsResponseBodyStackGroups) *ListStackGroupsResponseBody {
	s.StackGroups = v
	return s
}

func (s *ListStackGroupsResponseBody) SetTotalCount(v int32) *ListStackGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetRequestId(v string) *ListStackGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetPageSize(v int32) *ListStackGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetPageNumber(v int32) *ListStackGroupsResponseBody {
	s.PageNumber = &v
	return s
}

type ListStackGroupsResponseBodyStackGroups struct {
	StackGroupId          *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DriftDetectionTime    *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	StackGroupName        *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s ListStackGroupsResponseBodyStackGroups) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroups) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupId(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupId = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStatus(v string) *ListStackGroupsResponseBodyStackGroups {
	s.Status = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetDriftDetectionTime(v string) *ListStackGroupsResponseBodyStackGroups {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetDescription(v string) *ListStackGroupsResponseBodyStackGroups {
	s.Description = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupDriftStatus(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupName(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupName = &v
	return s
}

type ListStackGroupsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponse) SetHeaders(v map[string]*string) *ListStackGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupsResponse) SetBody(v *ListStackGroupsResponseBody) *ListStackGroupsResponse {
	s.Body = v
	return s
}

type ListStackInstancesRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName         *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	StackInstanceAccountId *string `json:"StackInstanceAccountId,omitempty" xml:"StackInstanceAccountId,omitempty"`
	StackInstanceRegionId  *string `json:"StackInstanceRegionId,omitempty" xml:"StackInstanceRegionId,omitempty"`
	PageSize               *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber             *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListStackInstancesRequest) SetRegionId(v string) *ListStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackGroupName(v string) *ListStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackInstanceAccountId(v string) *ListStackInstancesRequest {
	s.StackInstanceAccountId = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackInstanceRegionId(v string) *ListStackInstancesRequest {
	s.StackInstanceRegionId = &v
	return s
}

func (s *ListStackInstancesRequest) SetPageSize(v int64) *ListStackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackInstancesRequest) SetPageNumber(v int64) *ListStackInstancesRequest {
	s.PageNumber = &v
	return s
}

type ListStackInstancesResponseBody struct {
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StackInstances []*ListStackInstancesResponseBodyStackInstances `json:"StackInstances,omitempty" xml:"StackInstances,omitempty" type:"Repeated"`
}

func (s ListStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponseBody) SetTotalCount(v int32) *ListStackInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetRequestId(v string) *ListStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetPageSize(v int32) *ListStackInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetPageNumber(v int32) *ListStackInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetStackInstances(v []*ListStackInstancesResponseBodyStackInstances) *ListStackInstancesResponseBody {
	s.StackInstances = v
	return s
}

type ListStackInstancesResponseBodyStackInstances struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StackGroupId       *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	StackId            *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	StackDriftStatus   *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	StatusReason       *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	StackGroupName     *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountId          *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStackInstancesResponseBodyStackInstances) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesResponseBodyStackInstances) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStatus(v string) *ListStackInstancesResponseBodyStackInstances {
	s.Status = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackGroupId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackGroupId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetDriftDetectionTime(v string) *ListStackInstancesResponseBodyStackInstances {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackDriftStatus(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackDriftStatus = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStatusReason(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StatusReason = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackGroupName(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackGroupName = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetAccountId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.AccountId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetRegionId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.RegionId = &v
	return s
}

type ListStackInstancesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponse) SetHeaders(v map[string]*string) *ListStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListStackInstancesResponse) SetBody(v *ListStackInstancesResponseBody) *ListStackInstancesResponse {
	s.Body = v
	return s
}

type ListStackOperationRisksRequest struct {
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackId            *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	OperationType      *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ClientToken        *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RamRoleName        *string   `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RetainAllResources *bool     `json:"RetainAllResources,omitempty" xml:"RetainAllResources,omitempty"`
	RetainResources    []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
}

func (s ListStackOperationRisksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksRequest) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksRequest) SetRegionId(v string) *ListStackOperationRisksRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetStackId(v string) *ListStackOperationRisksRequest {
	s.StackId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetOperationType(v string) *ListStackOperationRisksRequest {
	s.OperationType = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetClientToken(v string) *ListStackOperationRisksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRamRoleName(v string) *ListStackOperationRisksRequest {
	s.RamRoleName = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRetainAllResources(v bool) *ListStackOperationRisksRequest {
	s.RetainAllResources = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRetainResources(v []*string) *ListStackOperationRisksRequest {
	s.RetainResources = v
	return s
}

type ListStackOperationRisksResponseBody struct {
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskResources []*ListStackOperationRisksResponseBodyRiskResources `json:"RiskResources,omitempty" xml:"RiskResources,omitempty" type:"Repeated"`
}

func (s ListStackOperationRisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponseBody) SetRequestId(v string) *ListStackOperationRisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackOperationRisksResponseBody) SetRiskResources(v []*ListStackOperationRisksResponseBodyRiskResources) *ListStackOperationRisksResponseBody {
	s.RiskResources = v
	return s
}

type ListStackOperationRisksResponseBodyRiskResources struct {
	LogicalResourceId  *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RiskType           *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	Reason             *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListStackOperationRisksResponseBodyRiskResources) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksResponseBodyRiskResources) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetLogicalResourceId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetPhysicalResourceId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetRequestId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.RequestId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetResourceType(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.ResourceType = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetCode(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Code = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetMessage(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Message = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetRiskType(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.RiskType = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetReason(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Reason = &v
	return s
}

type ListStackOperationRisksResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackOperationRisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackOperationRisksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksResponse) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponse) SetHeaders(v map[string]*string) *ListStackOperationRisksResponse {
	s.Headers = v
	return s
}

func (s *ListStackOperationRisksResponse) SetBody(v *ListStackOperationRisksResponseBody) *ListStackOperationRisksResponse {
	s.Body = v
	return s
}

type ListStackResourceDriftsRequest struct {
	StackId             *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults          *int64    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceDriftStatus []*string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty" type:"Repeated"`
}

func (s ListStackResourceDriftsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsRequest) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsRequest) SetStackId(v string) *ListStackResourceDriftsRequest {
	s.StackId = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetRegionId(v string) *ListStackResourceDriftsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetMaxResults(v int64) *ListStackResourceDriftsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetNextToken(v string) *ListStackResourceDriftsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetResourceDriftStatus(v []*string) *ListStackResourceDriftsRequest {
	s.ResourceDriftStatus = v
	return s
}

type ListStackResourceDriftsResponseBody struct {
	ResourceDrifts []*ListStackResourceDriftsResponseBodyResourceDrifts `json:"ResourceDrifts,omitempty" xml:"ResourceDrifts,omitempty" type:"Repeated"`
	NextToken      *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListStackResourceDriftsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBody) SetResourceDrifts(v []*ListStackResourceDriftsResponseBodyResourceDrifts) *ListStackResourceDriftsResponseBody {
	s.ResourceDrifts = v
	return s
}

func (s *ListStackResourceDriftsResponseBody) SetNextToken(v string) *ListStackResourceDriftsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStackResourceDriftsResponseBody) SetRequestId(v string) *ListStackResourceDriftsResponseBody {
	s.RequestId = &v
	return s
}

type ListStackResourceDriftsResponseBodyResourceDrifts struct {
	LogicalResourceId   *string                                                                 `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	StackId             *string                                                                 `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PhysicalResourceId  *string                                                                 `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	DriftDetectionTime  *string                                                                 `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	ResourceType        *string                                                                 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ExpectedProperties  *string                                                                 `json:"ExpectedProperties,omitempty" xml:"ExpectedProperties,omitempty"`
	ResourceDriftStatus *string                                                                 `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	PropertyDifferences []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences `json:"PropertyDifferences,omitempty" xml:"PropertyDifferences,omitempty" type:"Repeated"`
	ActualProperties    *string                                                                 `json:"ActualProperties,omitempty" xml:"ActualProperties,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDrifts) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDrifts) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetLogicalResourceId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetStackId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.StackId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetPhysicalResourceId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetDriftDetectionTime(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetResourceType(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ResourceType = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetExpectedProperties(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ExpectedProperties = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetResourceDriftStatus(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ResourceDriftStatus = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetPropertyDifferences(v []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.PropertyDifferences = v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetActualProperties(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ActualProperties = &v
	return s
}

type ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences struct {
	ActualValue    *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	DifferenceType *string `json:"DifferenceType,omitempty" xml:"DifferenceType,omitempty"`
	PropertyPath   *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
	ExpectedValue  *string `json:"ExpectedValue,omitempty" xml:"ExpectedValue,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetActualValue(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.ActualValue = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetDifferenceType(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.DifferenceType = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetPropertyPath(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.PropertyPath = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetExpectedValue(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.ExpectedValue = &v
	return s
}

type ListStackResourceDriftsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackResourceDriftsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackResourceDriftsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponse) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponse) SetHeaders(v map[string]*string) *ListStackResourceDriftsResponse {
	s.Headers = v
	return s
}

func (s *ListStackResourceDriftsResponse) SetBody(v *ListStackResourceDriftsResponseBody) *ListStackResourceDriftsResponse {
	s.Body = v
	return s
}

type ListStackResourcesRequest struct {
	StackId  *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStackResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListStackResourcesRequest) SetStackId(v string) *ListStackResourcesRequest {
	s.StackId = &v
	return s
}

func (s *ListStackResourcesRequest) SetRegionId(v string) *ListStackResourcesRequest {
	s.RegionId = &v
	return s
}

type ListStackResourcesResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*ListStackResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListStackResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBody) SetRequestId(v string) *ListStackResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackResourcesResponseBody) SetResources(v []*ListStackResourcesResponseBodyResources) *ListStackResourcesResponseBody {
	s.Resources = v
	return s
}

type ListStackResourcesResponseBodyResources struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	LogicalResourceId   *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	StackId             *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PhysicalResourceId  *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	DriftDetectionTime  *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	ResourceType        *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	StatusReason        *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StackName           *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
}

func (s ListStackResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBodyResources) SetStatus(v string) *ListStackResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetLogicalResourceId(v string) *ListStackResourcesResponseBodyResources {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetUpdateTime(v string) *ListStackResourcesResponseBodyResources {
	s.UpdateTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStackId(v string) *ListStackResourcesResponseBodyResources {
	s.StackId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetPhysicalResourceId(v string) *ListStackResourcesResponseBodyResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetDriftDetectionTime(v string) *ListStackResourcesResponseBodyResources {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetResourceType(v string) *ListStackResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetResourceDriftStatus(v string) *ListStackResourcesResponseBodyResources {
	s.ResourceDriftStatus = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStatusReason(v string) *ListStackResourcesResponseBodyResources {
	s.StatusReason = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetCreateTime(v string) *ListStackResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStackName(v string) *ListStackResourcesResponseBodyResources {
	s.StackName = &v
	return s
}

type ListStackResourcesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStackResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponse) SetHeaders(v map[string]*string) *ListStackResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListStackResourcesResponse) SetBody(v *ListStackResourcesResponseBody) *ListStackResourcesResponse {
	s.Body = v
	return s
}

type ListStacksRequest struct {
	PageSize        *int64                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentStackId   *string                 `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	RegionId        *string                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber      *int64                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ShowNestedStack *bool                   `json:"ShowNestedStack,omitempty" xml:"ShowNestedStack,omitempty"`
	StackId         *string                 `json:"StackId,omitempty" xml:"StackId,omitempty"`
	Status          []*string               `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	StackName       []*string               `json:"StackName,omitempty" xml:"StackName,omitempty" type:"Repeated"`
	Tag             []*ListStacksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListStacksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStacksRequest) GoString() string {
	return s.String()
}

func (s *ListStacksRequest) SetPageSize(v int64) *ListStacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListStacksRequest) SetParentStackId(v string) *ListStacksRequest {
	s.ParentStackId = &v
	return s
}

func (s *ListStacksRequest) SetRegionId(v string) *ListStacksRequest {
	s.RegionId = &v
	return s
}

func (s *ListStacksRequest) SetPageNumber(v int64) *ListStacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStacksRequest) SetShowNestedStack(v bool) *ListStacksRequest {
	s.ShowNestedStack = &v
	return s
}

func (s *ListStacksRequest) SetStackId(v string) *ListStacksRequest {
	s.StackId = &v
	return s
}

func (s *ListStacksRequest) SetStatus(v []*string) *ListStacksRequest {
	s.Status = v
	return s
}

func (s *ListStacksRequest) SetStackName(v []*string) *ListStacksRequest {
	s.StackName = v
	return s
}

func (s *ListStacksRequest) SetTag(v []*ListStacksRequestTag) *ListStacksRequest {
	s.Tag = v
	return s
}

type ListStacksRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStacksRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListStacksRequestTag) GoString() string {
	return s.String()
}

func (s *ListStacksRequestTag) SetKey(v string) *ListStacksRequestTag {
	s.Key = &v
	return s
}

func (s *ListStacksRequestTag) SetValue(v string) *ListStacksRequestTag {
	s.Value = &v
	return s
}

type ListStacksResponseBody struct {
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Stacks     []*ListStacksResponseBodyStacks `json:"Stacks,omitempty" xml:"Stacks,omitempty" type:"Repeated"`
}

func (s ListStacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBody) SetTotalCount(v int32) *ListStacksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStacksResponseBody) SetPageSize(v int32) *ListStacksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStacksResponseBody) SetRequestId(v string) *ListStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStacksResponseBody) SetPageNumber(v int32) *ListStacksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStacksResponseBody) SetStacks(v []*ListStacksResponseBodyStacks) *ListStacksResponseBody {
	s.Stacks = v
	return s
}

type ListStacksResponseBodyStacks struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	StatusReason       *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DisableRollback    *bool   `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	StackName          *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	TimeoutInMinutes   *int32  `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ParentStackId      *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	StackId            *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	StackDriftStatus   *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	StackType          *string `json:"StackType,omitempty" xml:"StackType,omitempty"`
}

func (s ListStacksResponseBodyStacks) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponseBodyStacks) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacks) SetStatus(v string) *ListStacksResponseBodyStacks {
	s.Status = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetUpdateTime(v string) *ListStacksResponseBodyStacks {
	s.UpdateTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDriftDetectionTime(v string) *ListStacksResponseBodyStacks {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStatusReason(v string) *ListStacksResponseBodyStacks {
	s.StatusReason = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetCreateTime(v string) *ListStacksResponseBodyStacks {
	s.CreateTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDisableRollback(v bool) *ListStacksResponseBodyStacks {
	s.DisableRollback = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackName(v string) *ListStacksResponseBodyStacks {
	s.StackName = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetTimeoutInMinutes(v int32) *ListStacksResponseBodyStacks {
	s.TimeoutInMinutes = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetRegionId(v string) *ListStacksResponseBodyStacks {
	s.RegionId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetParentStackId(v string) *ListStacksResponseBodyStacks {
	s.ParentStackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackId(v string) *ListStacksResponseBodyStacks {
	s.StackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackDriftStatus(v string) *ListStacksResponseBodyStacks {
	s.StackDriftStatus = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackType(v string) *ListStacksResponseBodyStacks {
	s.StackType = &v
	return s
}

type ListStacksResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStacksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponse) GoString() string {
	return s.String()
}

func (s *ListStacksResponse) SetHeaders(v map[string]*string) *ListStacksResponse {
	s.Headers = v
	return s
}

func (s *ListStacksResponse) SetBody(v *ListStacksResponseBody) *ListStacksResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

type ListTagKeysResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Keys      []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
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

type ListTagValuesRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Values    []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	PageNumber   *int64                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateName *string                    `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	ShareType    *string                    `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	Tag          []*ListTemplatesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetPageNumber(v int64) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int64) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) SetShareType(v string) *ListTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesRequest) SetTag(v []*ListTemplatesRequestTag) *ListTemplatesRequest {
	s.Tag = v
	return s
}

type ListTemplatesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplatesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequestTag) SetKey(v string) *ListTemplatesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTemplatesRequestTag) SetValue(v string) *ListTemplatesRequestTag {
	s.Value = &v
	return s
}

type ListTemplatesResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Templates  []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageNumber(v int32) *ListTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	TemplateARN     *string `json:"TemplateARN,omitempty" xml:"TemplateARN,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateARN(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateARN = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetOwnerId(v string) *ListTemplatesResponseBodyTemplates {
	s.OwnerId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetShareType(v string) *ListTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListTemplateVersionsRequest struct {
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplateVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsRequest) SetNextToken(v string) *ListTemplateVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetMaxResults(v int64) *ListTemplateVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetTemplateId(v string) *ListTemplateVersionsRequest {
	s.TemplateId = &v
	return s
}

type ListTemplateVersionsResponseBody struct {
	Versions  []*ListTemplateVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
	NextToken *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTemplateVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBody) SetVersions(v []*ListTemplateVersionsResponseBodyVersions) *ListTemplateVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetNextToken(v string) *ListTemplateVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetRequestId(v string) *ListTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListTemplateVersionsResponseBodyVersions struct {
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListTemplateVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBodyVersions) SetUpdateTime(v string) *ListTemplateVersionsResponseBodyVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetDescription(v string) *ListTemplateVersionsResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetCreateTime(v string) *ListTemplateVersionsResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateName(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateId(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateId = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateVersion(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateVersion = &v
	return s
}

type ListTemplateVersionsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplateVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponse) SetHeaders(v map[string]*string) *ListTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateVersionsResponse) SetBody(v *ListTemplateVersionsResponseBody) *ListTemplateVersionsResponse {
	s.Body = v
	return s
}

type PreviewStackRequest struct {
	DisableRollback  *bool                            `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	TimeoutInMinutes *int64                           `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	TemplateBody     *string                          `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	StackPolicyURL   *string                          `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	RegionId         *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackPolicyBody  *string                          `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	StackName        *string                          `json:"StackName,omitempty" xml:"StackName,omitempty"`
	ClientToken      *string                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TemplateURL      *string                          `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	TemplateId       *string                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion  *string                          `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters       []*PreviewStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s PreviewStackRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackRequest) GoString() string {
	return s.String()
}

func (s *PreviewStackRequest) SetDisableRollback(v bool) *PreviewStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *PreviewStackRequest) SetTimeoutInMinutes(v int64) *PreviewStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateBody(v string) *PreviewStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *PreviewStackRequest) SetStackPolicyURL(v string) *PreviewStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *PreviewStackRequest) SetRegionId(v string) *PreviewStackRequest {
	s.RegionId = &v
	return s
}

func (s *PreviewStackRequest) SetStackPolicyBody(v string) *PreviewStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *PreviewStackRequest) SetStackName(v string) *PreviewStackRequest {
	s.StackName = &v
	return s
}

func (s *PreviewStackRequest) SetClientToken(v string) *PreviewStackRequest {
	s.ClientToken = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateURL(v string) *PreviewStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateId(v string) *PreviewStackRequest {
	s.TemplateId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateVersion(v string) *PreviewStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *PreviewStackRequest) SetParameters(v []*PreviewStackRequestParameters) *PreviewStackRequest {
	s.Parameters = v
	return s
}

type PreviewStackRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s PreviewStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackRequestParameters) GoString() string {
	return s.String()
}

func (s *PreviewStackRequestParameters) SetParameterKey(v string) *PreviewStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *PreviewStackRequestParameters) SetParameterValue(v string) *PreviewStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type PreviewStackResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stack     *PreviewStackResponseBodyStack `json:"Stack,omitempty" xml:"Stack,omitempty" type:"Struct"`
}

func (s PreviewStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBody) SetRequestId(v string) *PreviewStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewStackResponseBody) SetStack(v *PreviewStackResponseBodyStack) *PreviewStackResponseBody {
	s.Stack = v
	return s
}

type PreviewStackResponseBodyStack struct {
	TemplateDescription *string                                    `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	Parameters          []*PreviewStackResponseBodyStackParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	Description         *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableRollback     *bool                                      `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	StackName           *string                                    `json:"StackName,omitempty" xml:"StackName,omitempty"`
	TimeoutInMinutes    *int32                                     `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	StackPolicyBody     map[string]interface{}                     `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	Resources           []*PreviewStackResponseBodyStackResources  `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	RegionId            *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PreviewStackResponseBodyStack) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStack) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStack) SetTemplateDescription(v string) *PreviewStackResponseBodyStack {
	s.TemplateDescription = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetParameters(v []*PreviewStackResponseBodyStackParameters) *PreviewStackResponseBodyStack {
	s.Parameters = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetDescription(v string) *PreviewStackResponseBodyStack {
	s.Description = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetDisableRollback(v bool) *PreviewStackResponseBodyStack {
	s.DisableRollback = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetStackName(v string) *PreviewStackResponseBodyStack {
	s.StackName = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetTimeoutInMinutes(v int32) *PreviewStackResponseBodyStack {
	s.TimeoutInMinutes = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetStackPolicyBody(v map[string]interface{}) *PreviewStackResponseBodyStack {
	s.StackPolicyBody = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetResources(v []*PreviewStackResponseBodyStackResources) *PreviewStackResponseBodyStack {
	s.Resources = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetRegionId(v string) *PreviewStackResponseBodyStack {
	s.RegionId = &v
	return s
}

type PreviewStackResponseBodyStackParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s PreviewStackResponseBodyStackParameters) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStackParameters) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackParameters) SetParameterKey(v string) *PreviewStackResponseBodyStackParameters {
	s.ParameterKey = &v
	return s
}

func (s *PreviewStackResponseBodyStackParameters) SetParameterValue(v string) *PreviewStackResponseBodyStackParameters {
	s.ParameterValue = &v
	return s
}

type PreviewStackResponseBodyStackResources struct {
	LogicalResourceId *string                `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	ResourceType      *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Description       *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Stack             map[string]interface{} `json:"Stack,omitempty" xml:"Stack,omitempty"`
	RequiredBy        []*string              `json:"RequiredBy,omitempty" xml:"RequiredBy,omitempty" type:"Repeated"`
	Properties        map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s PreviewStackResponseBodyStackResources) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStackResources) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackResources) SetLogicalResourceId(v string) *PreviewStackResponseBodyStackResources {
	s.LogicalResourceId = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetResourceType(v string) *PreviewStackResponseBodyStackResources {
	s.ResourceType = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetDescription(v string) *PreviewStackResponseBodyStackResources {
	s.Description = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetStack(v map[string]interface{}) *PreviewStackResponseBodyStackResources {
	s.Stack = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetRequiredBy(v []*string) *PreviewStackResponseBodyStackResources {
	s.RequiredBy = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetProperties(v map[string]interface{}) *PreviewStackResponseBodyStackResources {
	s.Properties = v
	return s
}

type PreviewStackResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PreviewStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviewStackResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponse) GoString() string {
	return s.String()
}

func (s *PreviewStackResponse) SetHeaders(v map[string]*string) *PreviewStackResponse {
	s.Headers = v
	return s
}

func (s *PreviewStackResponse) SetBody(v *PreviewStackResponseBody) *PreviewStackResponse {
	s.Body = v
	return s
}

type SetDeletionProtectionRequest struct {
	StackId            *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionRequest) SetStackId(v string) *SetDeletionProtectionRequest {
	s.StackId = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetDeletionProtection(v string) *SetDeletionProtectionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetRegionId(v string) *SetDeletionProtectionRequest {
	s.RegionId = &v
	return s
}

type SetDeletionProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponseBody) SetRequestId(v string) *SetDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetDeletionProtectionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponse) SetHeaders(v map[string]*string) *SetDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetDeletionProtectionResponse) SetBody(v *SetDeletionProtectionResponseBody) *SetDeletionProtectionResponse {
	s.Body = v
	return s
}

type SetStackPolicyRequest struct {
	StackId         *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	StackPolicyURL  *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
}

func (s SetStackPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetStackPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetStackPolicyRequest) SetStackId(v string) *SetStackPolicyRequest {
	s.StackId = &v
	return s
}

func (s *SetStackPolicyRequest) SetRegionId(v string) *SetStackPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackPolicyBody(v string) *SetStackPolicyRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackPolicyURL(v string) *SetStackPolicyRequest {
	s.StackPolicyURL = &v
	return s
}

type SetStackPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetStackPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetStackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetStackPolicyResponseBody) SetRequestId(v string) *SetStackPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetStackPolicyResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetStackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetStackPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetStackPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetStackPolicyResponse) SetHeaders(v map[string]*string) *SetStackPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetStackPolicyResponse) SetBody(v *SetStackPolicyResponseBody) *SetStackPolicyResponse {
	s.Body = v
	return s
}

type SetTemplatePermissionRequest struct {
	ShareOption     *string   `json:"ShareOption,omitempty" xml:"ShareOption,omitempty"`
	VersionOption   *string   `json:"VersionOption,omitempty" xml:"VersionOption,omitempty"`
	TemplateVersion *string   `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TemplateId      *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	AccountIds      []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
}

func (s SetTemplatePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTemplatePermissionRequest) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionRequest) SetShareOption(v string) *SetTemplatePermissionRequest {
	s.ShareOption = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetVersionOption(v string) *SetTemplatePermissionRequest {
	s.VersionOption = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetTemplateVersion(v string) *SetTemplatePermissionRequest {
	s.TemplateVersion = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetTemplateId(v string) *SetTemplatePermissionRequest {
	s.TemplateId = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetAccountIds(v []*string) *SetTemplatePermissionRequest {
	s.AccountIds = v
	return s
}

type SetTemplatePermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTemplatePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTemplatePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionResponseBody) SetRequestId(v string) *SetTemplatePermissionResponseBody {
	s.RequestId = &v
	return s
}

type SetTemplatePermissionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetTemplatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetTemplatePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTemplatePermissionResponse) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionResponse) SetHeaders(v map[string]*string) *SetTemplatePermissionResponse {
	s.Headers = v
	return s
}

func (s *SetTemplatePermissionResponse) SetBody(v *SetTemplatePermissionResponseBody) *SetTemplatePermissionResponse {
	s.Body = v
	return s
}

type SignalResourceRequest struct {
	StackId           *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UniqueId          *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Data              *string `json:"Data,omitempty" xml:"Data,omitempty"`
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
}

func (s SignalResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s SignalResourceRequest) GoString() string {
	return s.String()
}

func (s *SignalResourceRequest) SetStackId(v string) *SignalResourceRequest {
	s.StackId = &v
	return s
}

func (s *SignalResourceRequest) SetStatus(v string) *SignalResourceRequest {
	s.Status = &v
	return s
}

func (s *SignalResourceRequest) SetRegionId(v string) *SignalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *SignalResourceRequest) SetUniqueId(v string) *SignalResourceRequest {
	s.UniqueId = &v
	return s
}

func (s *SignalResourceRequest) SetClientToken(v string) *SignalResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *SignalResourceRequest) SetData(v string) *SignalResourceRequest {
	s.Data = &v
	return s
}

func (s *SignalResourceRequest) SetLogicalResourceId(v string) *SignalResourceRequest {
	s.LogicalResourceId = &v
	return s
}

type SignalResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SignalResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SignalResourceResponseBody) SetRequestId(v string) *SignalResourceResponseBody {
	s.RequestId = &v
	return s
}

type SignalResourceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SignalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignalResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s SignalResourceResponse) GoString() string {
	return s.String()
}

func (s *SignalResourceResponse) SetHeaders(v map[string]*string) *SignalResourceResponse {
	s.Headers = v
	return s
}

func (s *SignalResourceResponse) SetBody(v *SignalResourceResponseBody) *SignalResourceResponse {
	s.Body = v
	return s
}

type StopStackGroupOperationRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s StopStackGroupOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStackGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationRequest) SetRegionId(v string) *StopStackGroupOperationRequest {
	s.RegionId = &v
	return s
}

func (s *StopStackGroupOperationRequest) SetOperationId(v string) *StopStackGroupOperationRequest {
	s.OperationId = &v
	return s
}

type StopStackGroupOperationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopStackGroupOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStackGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationResponseBody) SetRequestId(v string) *StopStackGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

type StopStackGroupOperationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopStackGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStackGroupOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStackGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationResponse) SetHeaders(v map[string]*string) *StopStackGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *StopStackGroupOperationResponse) SetBody(v *StopStackGroupOperationResponseBody) *StopStackGroupOperationResponse {
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

type UpdateStackRequest struct {
	StackId                     *string                         `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ClientToken                 *string                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StackPolicyDuringUpdateBody *string                         `json:"StackPolicyDuringUpdateBody,omitempty" xml:"StackPolicyDuringUpdateBody,omitempty"`
	TimeoutInMinutes            *int64                          `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	TemplateBody                *string                         `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	StackPolicyURL              *string                         `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	UpdateAllowPolicy           *string                         `json:"UpdateAllowPolicy,omitempty" xml:"UpdateAllowPolicy,omitempty"`
	StackPolicyDuringUpdateURL  *string                         `json:"StackPolicyDuringUpdateURL,omitempty" xml:"StackPolicyDuringUpdateURL,omitempty"`
	StackPolicyBody             *string                         `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	UsePreviousParameters       *bool                           `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
	RegionId                    *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DisableRollback             *bool                           `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	EnableRecover               *bool                           `json:"EnableRecover,omitempty" xml:"EnableRecover,omitempty"`
	TemplateURL                 *string                         `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	RamRoleName                 *string                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ReplacementOption           *string                         `json:"ReplacementOption,omitempty" xml:"ReplacementOption,omitempty"`
	TemplateId                  *string                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion             *string                         `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters                  []*UpdateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackRequest) SetStackId(v string) *UpdateStackRequest {
	s.StackId = &v
	return s
}

func (s *UpdateStackRequest) SetClientToken(v string) *UpdateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyDuringUpdateBody(v string) *UpdateStackRequest {
	s.StackPolicyDuringUpdateBody = &v
	return s
}

func (s *UpdateStackRequest) SetTimeoutInMinutes(v int64) *UpdateStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateBody(v string) *UpdateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyURL(v string) *UpdateStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *UpdateStackRequest) SetUpdateAllowPolicy(v string) *UpdateStackRequest {
	s.UpdateAllowPolicy = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyDuringUpdateURL(v string) *UpdateStackRequest {
	s.StackPolicyDuringUpdateURL = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyBody(v string) *UpdateStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *UpdateStackRequest) SetUsePreviousParameters(v bool) *UpdateStackRequest {
	s.UsePreviousParameters = &v
	return s
}

func (s *UpdateStackRequest) SetRegionId(v string) *UpdateStackRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackRequest) SetDisableRollback(v bool) *UpdateStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *UpdateStackRequest) SetEnableRecover(v bool) *UpdateStackRequest {
	s.EnableRecover = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateURL(v string) *UpdateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackRequest) SetRamRoleName(v string) *UpdateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *UpdateStackRequest) SetReplacementOption(v string) *UpdateStackRequest {
	s.ReplacementOption = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateId(v string) *UpdateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateVersion(v string) *UpdateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackRequest) SetParameters(v []*UpdateStackRequestParameters) *UpdateStackRequest {
	s.Parameters = v
	return s
}

type UpdateStackRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackRequestParameters) SetParameterKey(v string) *UpdateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackRequestParameters) SetParameterValue(v string) *UpdateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateStackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackId   *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s UpdateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackResponseBody) SetRequestId(v string) *UpdateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackResponseBody) SetStackId(v string) *UpdateStackResponseBody {
	s.StackId = &v
	return s
}

type UpdateStackResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackResponse) SetHeaders(v map[string]*string) *UpdateStackResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackResponse) SetBody(v *UpdateStackResponseBody) *UpdateStackResponse {
	s.Body = v
	return s
}

type UpdateStackGroupRequest struct {
	RegionId               *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName         *string                              `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	Description            *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	AccountIds             map[string]interface{}               `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIds              map[string]interface{}               `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	TemplateBody           *string                              `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateURL            *string                              `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	ClientToken            *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription   *string                              `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferences   map[string]interface{}               `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	AdministrationRoleName *string                              `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	ExecutionRoleName      *string                              `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	TemplateId             *string                              `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion        *string                              `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters             []*UpdateStackGroupRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequest) SetRegionId(v string) *UpdateStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackGroupRequest) SetStackGroupName(v string) *UpdateStackGroupRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackGroupRequest) SetDescription(v string) *UpdateStackGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateStackGroupRequest) SetAccountIds(v map[string]interface{}) *UpdateStackGroupRequest {
	s.AccountIds = v
	return s
}

func (s *UpdateStackGroupRequest) SetRegionIds(v map[string]interface{}) *UpdateStackGroupRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateBody(v string) *UpdateStackGroupRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateURL(v string) *UpdateStackGroupRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackGroupRequest) SetClientToken(v string) *UpdateStackGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackGroupRequest) SetOperationDescription(v string) *UpdateStackGroupRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackGroupRequest) SetOperationPreferences(v map[string]interface{}) *UpdateStackGroupRequest {
	s.OperationPreferences = v
	return s
}

func (s *UpdateStackGroupRequest) SetAdministrationRoleName(v string) *UpdateStackGroupRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *UpdateStackGroupRequest) SetExecutionRoleName(v string) *UpdateStackGroupRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateId(v string) *UpdateStackGroupRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateVersion(v string) *UpdateStackGroupRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackGroupRequest) SetParameters(v []*UpdateStackGroupRequestParameters) *UpdateStackGroupRequest {
	s.Parameters = v
	return s
}

type UpdateStackGroupRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackGroupRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestParameters) SetParameterKey(v string) *UpdateStackGroupRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackGroupRequestParameters) SetParameterValue(v string) *UpdateStackGroupRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateStackGroupShrinkRequest struct {
	RegionId                   *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName             *string                                    `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	Description                *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	AccountIdsShrink           *string                                    `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIdsShrink            *string                                    `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	TemplateBody               *string                                    `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateURL                *string                                    `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	ClientToken                *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription       *string                                    `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferencesShrink *string                                    `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	AdministrationRoleName     *string                                    `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	ExecutionRoleName          *string                                    `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	TemplateId                 *string                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion            *string                                    `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Parameters                 []*UpdateStackGroupShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateStackGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupShrinkRequest) SetRegionId(v string) *UpdateStackGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetStackGroupName(v string) *UpdateStackGroupShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetDescription(v string) *UpdateStackGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetAccountIdsShrink(v string) *UpdateStackGroupShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetRegionIdsShrink(v string) *UpdateStackGroupShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateBody(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateURL(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetClientToken(v string) *UpdateStackGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetOperationDescription(v string) *UpdateStackGroupShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetOperationPreferencesShrink(v string) *UpdateStackGroupShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetAdministrationRoleName(v string) *UpdateStackGroupShrinkRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetExecutionRoleName(v string) *UpdateStackGroupShrinkRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateId(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateVersion(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetParameters(v []*UpdateStackGroupShrinkRequestParameters) *UpdateStackGroupShrinkRequest {
	s.Parameters = v
	return s
}

type UpdateStackGroupShrinkRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackGroupShrinkRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupShrinkRequestParameters) SetParameterKey(v string) *UpdateStackGroupShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackGroupShrinkRequestParameters) SetParameterValue(v string) *UpdateStackGroupShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateStackGroupResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s UpdateStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupResponseBody) SetRequestId(v string) *UpdateStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackGroupResponseBody) SetOperationId(v string) *UpdateStackGroupResponseBody {
	s.OperationId = &v
	return s
}

type UpdateStackGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupResponse) SetHeaders(v map[string]*string) *UpdateStackGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackGroupResponse) SetBody(v *UpdateStackGroupResponseBody) *UpdateStackGroupResponse {
	s.Body = v
	return s
}

type UpdateStackInstancesRequest struct {
	RegionId             *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName       *string                                          `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountIds           map[string]interface{}                           `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIds            map[string]interface{}                           `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ClientToken          *string                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription *string                                          `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferences map[string]interface{}                           `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	TimeoutInMinutes     *int64                                           `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	ParameterOverrides   []*UpdateStackInstancesRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
}

func (s UpdateStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequest) SetRegionId(v string) *UpdateStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetStackGroupName(v string) *UpdateStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetAccountIds(v map[string]interface{}) *UpdateStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *UpdateStackInstancesRequest) SetRegionIds(v map[string]interface{}) *UpdateStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateStackInstancesRequest) SetClientToken(v string) *UpdateStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetOperationDescription(v string) *UpdateStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *UpdateStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *UpdateStackInstancesRequest) SetTimeoutInMinutes(v int64) *UpdateStackInstancesRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetParameterOverrides(v []*UpdateStackInstancesRequestParameterOverrides) *UpdateStackInstancesRequest {
	s.ParameterOverrides = v
	return s
}

type UpdateStackInstancesRequestParameterOverrides struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackInstancesRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequestParameterOverrides) SetParameterKey(v string) *UpdateStackInstancesRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackInstancesRequestParameterOverrides) SetParameterValue(v string) *UpdateStackInstancesRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type UpdateStackInstancesShrinkRequest struct {
	RegionId                   *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StackGroupName             *string                                                `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	AccountIdsShrink           *string                                                `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	RegionIdsShrink            *string                                                `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ClientToken                *string                                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OperationDescription       *string                                                `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationPreferencesShrink *string                                                `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	TimeoutInMinutes           *int64                                                 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	ParameterOverrides         []*UpdateStackInstancesShrinkRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
}

func (s UpdateStackInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesShrinkRequest) SetRegionId(v string) *UpdateStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetStackGroupName(v string) *UpdateStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetClientToken(v string) *UpdateStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetOperationDescription(v string) *UpdateStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetTimeoutInMinutes(v int64) *UpdateStackInstancesShrinkRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetParameterOverrides(v []*UpdateStackInstancesShrinkRequestParameterOverrides) *UpdateStackInstancesShrinkRequest {
	s.ParameterOverrides = v
	return s
}

type UpdateStackInstancesShrinkRequestParameterOverrides struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackInstancesShrinkRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesShrinkRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) SetParameterKey(v string) *UpdateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) SetParameterValue(v string) *UpdateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type UpdateStackInstancesResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s UpdateStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesResponseBody) SetRequestId(v string) *UpdateStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackInstancesResponseBody) SetOperationId(v string) *UpdateStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

type UpdateStackInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesResponse) SetHeaders(v map[string]*string) *UpdateStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackInstancesResponse) SetBody(v *UpdateStackInstancesResponseBody) *UpdateStackInstancesResponse {
	s.Body = v
	return s
}

type UpdateStackTemplateByResourcesRequest struct {
	StackId           *string   `json:"StackId,omitempty" xml:"StackId,omitempty"`
	DryRun            *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TemplateFormat    *string   `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
}

func (s UpdateStackTemplateByResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackTemplateByResourcesRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesRequest) SetStackId(v string) *UpdateStackTemplateByResourcesRequest {
	s.StackId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetDryRun(v bool) *UpdateStackTemplateByResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetRegionId(v string) *UpdateStackTemplateByResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetClientToken(v string) *UpdateStackTemplateByResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetTemplateFormat(v string) *UpdateStackTemplateByResourcesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetLogicalResourceId(v []*string) *UpdateStackTemplateByResourcesRequest {
	s.LogicalResourceId = v
	return s
}

type UpdateStackTemplateByResourcesResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NewTemplateBody *string `json:"NewTemplateBody,omitempty" xml:"NewTemplateBody,omitempty"`
	OldTemplateBody *string `json:"OldTemplateBody,omitempty" xml:"OldTemplateBody,omitempty"`
}

func (s UpdateStackTemplateByResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackTemplateByResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetRequestId(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetNewTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.NewTemplateBody = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetOldTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.OldTemplateBody = &v
	return s
}

type UpdateStackTemplateByResourcesResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateStackTemplateByResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackTemplateByResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackTemplateByResourcesResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesResponse) SetHeaders(v map[string]*string) *UpdateStackTemplateByResourcesResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackTemplateByResourcesResponse) SetBody(v *UpdateStackTemplateByResourcesResponseBody) *UpdateStackTemplateByResourcesResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	TemplateURL  *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetTemplateURL(v string) *UpdateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateBody(v string) *UpdateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

type UpdateTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateId(v string) *UpdateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type ValidateTemplateRequest struct {
	TemplateURL  *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
}

func (s ValidateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateRequest) GoString() string {
	return s.String()
}

func (s *ValidateTemplateRequest) SetTemplateURL(v string) *ValidateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *ValidateTemplateRequest) SetRegionId(v string) *ValidateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateTemplateRequest) SetTemplateBody(v string) *ValidateTemplateRequest {
	s.TemplateBody = &v
	return s
}

type ValidateTemplateResponseBody struct {
	Description *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Parameters  []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId   *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBody) SetDescription(v string) *ValidateTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *ValidateTemplateResponseBody) SetParameters(v []map[string]interface{}) *ValidateTemplateResponseBody {
	s.Parameters = v
	return s
}

func (s *ValidateTemplateResponseBody) SetRequestId(v string) *ValidateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ValidateTemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponse) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponse) SetHeaders(v map[string]*string) *ValidateTemplateResponse {
	s.Headers = v
	return s
}

func (s *ValidateTemplateResponse) SetBody(v *ValidateTemplateResponseBody) *ValidateTemplateResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ros"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelUpdateStackWithOptions(request *CancelUpdateStackRequest, runtime *util.RuntimeOptions) (_result *CancelUpdateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelUpdateStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelUpdateStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelUpdateStack(request *CancelUpdateStackRequest) (_result *CancelUpdateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelUpdateStackResponse{}
	_body, _err := client.CancelUpdateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContinueCreateStackWithOptions(request *ContinueCreateStackRequest, runtime *util.RuntimeOptions) (_result *ContinueCreateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ContinueCreateStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ContinueCreateStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContinueCreateStack(request *ContinueCreateStackRequest) (_result *ContinueCreateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueCreateStackResponse{}
	_body, _err := client.ContinueCreateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateChangeSetWithOptions(request *CreateChangeSetRequest, runtime *util.RuntimeOptions) (_result *CreateChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateChangeSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateChangeSet"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateChangeSet(request *CreateChangeSetRequest) (_result *CreateChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChangeSetResponse{}
	_body, _err := client.CreateChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStackWithOptions(request *CreateStackRequest, runtime *util.RuntimeOptions) (_result *CreateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStack(request *CreateStackRequest) (_result *CreateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStackResponse{}
	_body, _err := client.CreateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStackGroupWithOptions(request *CreateStackGroupRequest, runtime *util.RuntimeOptions) (_result *CreateStackGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStackGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStackGroup"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStackGroup(request *CreateStackGroupRequest) (_result *CreateStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStackGroupResponse{}
	_body, _err := client.CreateStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStackInstancesWithOptions(tmpReq *CreateStackInstancesRequest, runtime *util.RuntimeOptions) (_result *CreateStackInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStackInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStackInstances"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStackInstances(request *CreateStackInstancesRequest) (_result *CreateStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStackInstancesResponse{}
	_body, _err := client.CreateStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTemplate"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteChangeSetWithOptions(request *DeleteChangeSetRequest, runtime *util.RuntimeOptions) (_result *DeleteChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteChangeSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteChangeSet"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteChangeSet(request *DeleteChangeSetRequest) (_result *DeleteChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChangeSetResponse{}
	_body, _err := client.DeleteChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStackWithOptions(request *DeleteStackRequest, runtime *util.RuntimeOptions) (_result *DeleteStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStack(request *DeleteStackRequest) (_result *DeleteStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStackResponse{}
	_body, _err := client.DeleteStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStackGroupWithOptions(request *DeleteStackGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteStackGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteStackGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteStackGroup"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStackGroup(request *DeleteStackGroupRequest) (_result *DeleteStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStackGroupResponse{}
	_body, _err := client.DeleteStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStackInstancesWithOptions(tmpReq *DeleteStackInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteStackInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteStackInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteStackInstances"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStackInstances(request *DeleteStackInstancesRequest) (_result *DeleteStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStackInstancesResponse{}
	_body, _err := client.DeleteStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTemplate"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DetectStackDriftWithOptions(request *DetectStackDriftRequest, runtime *util.RuntimeOptions) (_result *DetectStackDriftResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectStackDriftResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectStackDrift"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectStackDrift(request *DetectStackDriftRequest) (_result *DetectStackDriftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectStackDriftResponse{}
	_body, _err := client.DetectStackDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectStackGroupDriftWithOptions(tmpReq *DetectStackGroupDriftRequest, runtime *util.RuntimeOptions) (_result *DetectStackGroupDriftResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectStackGroupDriftShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectStackGroupDriftResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectStackGroupDrift"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectStackGroupDrift(request *DetectStackGroupDriftRequest) (_result *DetectStackGroupDriftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectStackGroupDriftResponse{}
	_body, _err := client.DetectStackGroupDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectStackResourceDriftWithOptions(request *DetectStackResourceDriftRequest, runtime *util.RuntimeOptions) (_result *DetectStackResourceDriftResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectStackResourceDriftResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectStackResourceDrift"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectStackResourceDrift(request *DetectStackResourceDriftRequest) (_result *DetectStackResourceDriftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectStackResourceDriftResponse{}
	_body, _err := client.DetectStackResourceDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteChangeSetWithOptions(request *ExecuteChangeSetRequest, runtime *util.RuntimeOptions) (_result *ExecuteChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteChangeSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteChangeSet"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteChangeSet(request *ExecuteChangeSetRequest) (_result *ExecuteChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteChangeSetResponse{}
	_body, _err := client.ExecuteChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateTemplatePolicyWithOptions(request *GenerateTemplatePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateTemplatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateTemplatePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateTemplatePolicy"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateTemplatePolicy(request *GenerateTemplatePolicyRequest) (_result *GenerateTemplatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateTemplatePolicyResponse{}
	_body, _err := client.GenerateTemplatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetChangeSetWithOptions(request *GetChangeSetRequest, runtime *util.RuntimeOptions) (_result *GetChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetChangeSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetChangeSet"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChangeSet(request *GetChangeSetRequest) (_result *GetChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChangeSetResponse{}
	_body, _err := client.GetChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceTypeWithOptions(request *GetResourceTypeRequest, runtime *util.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetResourceType"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceType(request *GetResourceTypeRequest) (_result *GetResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.GetResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceTypeTemplateWithOptions(request *GetResourceTypeTemplateRequest, runtime *util.RuntimeOptions) (_result *GetResourceTypeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetResourceTypeTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetResourceTypeTemplate"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceTypeTemplate(request *GetResourceTypeTemplateRequest) (_result *GetResourceTypeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceTypeTemplateResponse{}
	_body, _err := client.GetResourceTypeTemplateWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("GetStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetStackDriftDetectionStatusWithOptions(request *GetStackDriftDetectionStatusRequest, runtime *util.RuntimeOptions) (_result *GetStackDriftDetectionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackDriftDetectionStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStackDriftDetectionStatus"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStackDriftDetectionStatus(request *GetStackDriftDetectionStatusRequest) (_result *GetStackDriftDetectionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackDriftDetectionStatusResponse{}
	_body, _err := client.GetStackDriftDetectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackGroupWithOptions(request *GetStackGroupRequest, runtime *util.RuntimeOptions) (_result *GetStackGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStackGroup"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStackGroup(request *GetStackGroupRequest) (_result *GetStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackGroupResponse{}
	_body, _err := client.GetStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackGroupOperationWithOptions(request *GetStackGroupOperationRequest, runtime *util.RuntimeOptions) (_result *GetStackGroupOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackGroupOperationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStackGroupOperation"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStackGroupOperation(request *GetStackGroupOperationRequest) (_result *GetStackGroupOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackGroupOperationResponse{}
	_body, _err := client.GetStackGroupOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackInstanceWithOptions(request *GetStackInstanceRequest, runtime *util.RuntimeOptions) (_result *GetStackInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStackInstance"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStackInstance(request *GetStackInstanceRequest) (_result *GetStackInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackInstanceResponse{}
	_body, _err := client.GetStackInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackPolicyWithOptions(request *GetStackPolicyRequest, runtime *util.RuntimeOptions) (_result *GetStackPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStackPolicy"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStackPolicy(request *GetStackPolicyRequest) (_result *GetStackPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackPolicyResponse{}
	_body, _err := client.GetStackPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackResourceWithOptions(request *GetStackResourceRequest, runtime *util.RuntimeOptions) (_result *GetStackResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStackResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStackResource"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStackResource(request *GetStackResourceRequest) (_result *GetStackResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackResourceResponse{}
	_body, _err := client.GetStackResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplate"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateEstimateCostWithOptions(request *GetTemplateEstimateCostRequest, runtime *util.RuntimeOptions) (_result *GetTemplateEstimateCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateEstimateCostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplateEstimateCost"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplateEstimateCost(request *GetTemplateEstimateCostRequest) (_result *GetTemplateEstimateCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateEstimateCostResponse{}
	_body, _err := client.GetTemplateEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateSummaryWithOptions(request *GetTemplateSummaryRequest, runtime *util.RuntimeOptions) (_result *GetTemplateSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplateSummary"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplateSummary(request *GetTemplateSummaryRequest) (_result *GetTemplateSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateSummaryResponse{}
	_body, _err := client.GetTemplateSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChangeSetsWithOptions(request *ListChangeSetsRequest, runtime *util.RuntimeOptions) (_result *ListChangeSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListChangeSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListChangeSets"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChangeSets(request *ListChangeSetsRequest) (_result *ListChangeSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChangeSetsResponse{}
	_body, _err := client.ListChangeSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListResourceTypes"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTypes() (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackEventsWithOptions(request *ListStackEventsRequest, runtime *util.RuntimeOptions) (_result *ListStackEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackEvents"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackEvents(request *ListStackEventsRequest) (_result *ListStackEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackEventsResponse{}
	_body, _err := client.ListStackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackGroupOperationResultsWithOptions(request *ListStackGroupOperationResultsRequest, runtime *util.RuntimeOptions) (_result *ListStackGroupOperationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackGroupOperationResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackGroupOperationResults"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackGroupOperationResults(request *ListStackGroupOperationResultsRequest) (_result *ListStackGroupOperationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackGroupOperationResultsResponse{}
	_body, _err := client.ListStackGroupOperationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackGroupOperationsWithOptions(request *ListStackGroupOperationsRequest, runtime *util.RuntimeOptions) (_result *ListStackGroupOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackGroupOperationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackGroupOperations"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackGroupOperations(request *ListStackGroupOperationsRequest) (_result *ListStackGroupOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackGroupOperationsResponse{}
	_body, _err := client.ListStackGroupOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackGroupsWithOptions(request *ListStackGroupsRequest, runtime *util.RuntimeOptions) (_result *ListStackGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackGroups"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackGroups(request *ListStackGroupsRequest) (_result *ListStackGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackGroupsResponse{}
	_body, _err := client.ListStackGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackInstancesWithOptions(request *ListStackInstancesRequest, runtime *util.RuntimeOptions) (_result *ListStackInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackInstances"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackInstances(request *ListStackInstancesRequest) (_result *ListStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackInstancesResponse{}
	_body, _err := client.ListStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackOperationRisksWithOptions(request *ListStackOperationRisksRequest, runtime *util.RuntimeOptions) (_result *ListStackOperationRisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackOperationRisksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackOperationRisks"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackOperationRisks(request *ListStackOperationRisksRequest) (_result *ListStackOperationRisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackOperationRisksResponse{}
	_body, _err := client.ListStackOperationRisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackResourceDriftsWithOptions(request *ListStackResourceDriftsRequest, runtime *util.RuntimeOptions) (_result *ListStackResourceDriftsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackResourceDriftsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackResourceDrifts"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackResourceDrifts(request *ListStackResourceDriftsRequest) (_result *ListStackResourceDriftsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackResourceDriftsResponse{}
	_body, _err := client.ListStackResourceDriftsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackResourcesWithOptions(request *ListStackResourcesRequest, runtime *util.RuntimeOptions) (_result *ListStackResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStackResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStackResources"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackResources(request *ListStackResourcesRequest) (_result *ListStackResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackResourcesResponse{}
	_body, _err := client.ListStackResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStacksWithOptions(request *ListStacksRequest, runtime *util.RuntimeOptions) (_result *ListStacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStacksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStacks"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStacks(request *ListStacksRequest) (_result *ListStacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStacksResponse{}
	_body, _err := client.ListStacksWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplates"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplateVersionsWithOptions(request *ListTemplateVersionsRequest, runtime *util.RuntimeOptions) (_result *ListTemplateVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplateVersions"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (_result *ListTemplateVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.ListTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreviewStackWithOptions(request *PreviewStackRequest, runtime *util.RuntimeOptions) (_result *PreviewStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PreviewStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PreviewStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreviewStack(request *PreviewStackRequest) (_result *PreviewStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreviewStackResponse{}
	_body, _err := client.PreviewStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeletionProtectionWithOptions(request *SetDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *SetDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDeletionProtection"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeletionProtection(request *SetDeletionProtectionRequest) (_result *SetDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.SetDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetStackPolicyWithOptions(request *SetStackPolicyRequest, runtime *util.RuntimeOptions) (_result *SetStackPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetStackPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetStackPolicy"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetStackPolicy(request *SetStackPolicyRequest) (_result *SetStackPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetStackPolicyResponse{}
	_body, _err := client.SetStackPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetTemplatePermissionWithOptions(request *SetTemplatePermissionRequest, runtime *util.RuntimeOptions) (_result *SetTemplatePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetTemplatePermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetTemplatePermission"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetTemplatePermission(request *SetTemplatePermissionRequest) (_result *SetTemplatePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTemplatePermissionResponse{}
	_body, _err := client.SetTemplatePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignalResourceWithOptions(request *SignalResourceRequest, runtime *util.RuntimeOptions) (_result *SignalResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SignalResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SignalResource"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignalResource(request *SignalResourceRequest) (_result *SignalResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SignalResourceResponse{}
	_body, _err := client.SignalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStackGroupOperationWithOptions(request *StopStackGroupOperationRequest, runtime *util.RuntimeOptions) (_result *StopStackGroupOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopStackGroupOperationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopStackGroupOperation"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStackGroupOperation(request *StopStackGroupOperationRequest) (_result *StopStackGroupOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopStackGroupOperationResponse{}
	_body, _err := client.StopStackGroupOperationWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateStackWithOptions(request *UpdateStackRequest, runtime *util.RuntimeOptions) (_result *UpdateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateStackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateStack"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStack(request *UpdateStackRequest) (_result *UpdateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackResponse{}
	_body, _err := client.UpdateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStackGroupWithOptions(tmpReq *UpdateStackGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateStackGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStackGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateStackGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateStackGroup"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStackGroup(request *UpdateStackGroupRequest) (_result *UpdateStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackGroupResponse{}
	_body, _err := client.UpdateStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStackInstancesWithOptions(tmpReq *UpdateStackInstancesRequest, runtime *util.RuntimeOptions) (_result *UpdateStackInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateStackInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateStackInstances"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStackInstances(request *UpdateStackInstancesRequest) (_result *UpdateStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackInstancesResponse{}
	_body, _err := client.UpdateStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStackTemplateByResourcesWithOptions(request *UpdateStackTemplateByResourcesRequest, runtime *util.RuntimeOptions) (_result *UpdateStackTemplateByResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateStackTemplateByResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateStackTemplateByResources"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStackTemplateByResources(request *UpdateStackTemplateByResourcesRequest) (_result *UpdateStackTemplateByResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackTemplateByResourcesResponse{}
	_body, _err := client.UpdateStackTemplateByResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(request *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTemplate"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateTemplateWithOptions(request *ValidateTemplateRequest, runtime *util.RuntimeOptions) (_result *ValidateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateTemplate"), tea.String("2019-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateTemplate(request *ValidateTemplateRequest) (_result *ValidateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateTemplateResponse{}
	_body, _err := client.ValidateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

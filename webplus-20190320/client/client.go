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

type AbortChangeRequest struct {
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s AbortChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeRequest) GoString() string {
	return s.String()
}

func (s *AbortChangeRequest) SetChangeId(v string) *AbortChangeRequest {
	s.ChangeId = &v
	return s
}

type AbortChangeResponseBody struct {
	EnvChange *AbortChangeResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AbortChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeResponseBody) GoString() string {
	return s.String()
}

func (s *AbortChangeResponseBody) SetEnvChange(v *AbortChangeResponseBodyEnvChange) *AbortChangeResponseBody {
	s.EnvChange = v
	return s
}

func (s *AbortChangeResponseBody) SetMessage(v string) *AbortChangeResponseBody {
	s.Message = &v
	return s
}

func (s *AbortChangeResponseBody) SetRequestId(v string) *AbortChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortChangeResponseBody) SetCode(v string) *AbortChangeResponseBody {
	s.Code = &v
	return s
}

type AbortChangeResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s AbortChangeResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *AbortChangeResponseBodyEnvChange) SetStartTime(v string) *AbortChangeResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *AbortChangeResponseBodyEnvChange) SetChangeId(v string) *AbortChangeResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *AbortChangeResponseBodyEnvChange) SetEnvId(v string) *AbortChangeResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type AbortChangeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbortChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeResponse) GoString() string {
	return s.String()
}

func (s *AbortChangeResponse) SetHeaders(v map[string]*string) *AbortChangeResponse {
	s.Headers = v
	return s
}

func (s *AbortChangeResponse) SetBody(v *AbortChangeResponseBody) *AbortChangeResponse {
	s.Body = v
	return s
}

type CreateAppEnvRequest struct {
	EnvName         *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvDescription  *string `json:"EnvDescription,omitempty" xml:"EnvDescription,omitempty"`
	StackId         *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PkgVersionId    *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	OptionSettings  *string `json:"OptionSettings,omitempty" xml:"OptionSettings,omitempty"`
	ProfileName     *string `json:"ProfileName,omitempty" xml:"ProfileName,omitempty"`
	SourceEnvId     *string `json:"SourceEnvId,omitempty" xml:"SourceEnvId,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ExtraProperties *string `json:"ExtraProperties,omitempty" xml:"ExtraProperties,omitempty"`
}

func (s CreateAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppEnvRequest) GoString() string {
	return s.String()
}

func (s *CreateAppEnvRequest) SetEnvName(v string) *CreateAppEnvRequest {
	s.EnvName = &v
	return s
}

func (s *CreateAppEnvRequest) SetEnvDescription(v string) *CreateAppEnvRequest {
	s.EnvDescription = &v
	return s
}

func (s *CreateAppEnvRequest) SetStackId(v string) *CreateAppEnvRequest {
	s.StackId = &v
	return s
}

func (s *CreateAppEnvRequest) SetAppId(v string) *CreateAppEnvRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppEnvRequest) SetPkgVersionId(v string) *CreateAppEnvRequest {
	s.PkgVersionId = &v
	return s
}

func (s *CreateAppEnvRequest) SetOptionSettings(v string) *CreateAppEnvRequest {
	s.OptionSettings = &v
	return s
}

func (s *CreateAppEnvRequest) SetProfileName(v string) *CreateAppEnvRequest {
	s.ProfileName = &v
	return s
}

func (s *CreateAppEnvRequest) SetSourceEnvId(v string) *CreateAppEnvRequest {
	s.SourceEnvId = &v
	return s
}

func (s *CreateAppEnvRequest) SetTemplateId(v string) *CreateAppEnvRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateAppEnvRequest) SetDryRun(v bool) *CreateAppEnvRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAppEnvRequest) SetExtraProperties(v string) *CreateAppEnvRequest {
	s.ExtraProperties = &v
	return s
}

type CreateAppEnvResponseBody struct {
	Message         *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvChangeDetail *CreateAppEnvResponseBodyEnvChangeDetail `json:"EnvChangeDetail,omitempty" xml:"EnvChangeDetail,omitempty" type:"Struct"`
	Code            *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppEnvResponseBody) SetMessage(v string) *CreateAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppEnvResponseBody) SetRequestId(v string) *CreateAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppEnvResponseBody) SetEnvChangeDetail(v *CreateAppEnvResponseBodyEnvChangeDetail) *CreateAppEnvResponseBody {
	s.EnvChangeDetail = v
	return s
}

func (s *CreateAppEnvResponseBody) SetCode(v string) *CreateAppEnvResponseBody {
	s.Code = &v
	return s
}

type CreateAppEnvResponseBodyEnvChangeDetail struct {
	StartTime  *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId   *string                                            `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId      *string                                            `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Operations *CreateAppEnvResponseBodyEnvChangeDetailOperations `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Struct"`
}

func (s CreateAppEnvResponseBodyEnvChangeDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateAppEnvResponseBodyEnvChangeDetail) GoString() string {
	return s.String()
}

func (s *CreateAppEnvResponseBodyEnvChangeDetail) SetStartTime(v string) *CreateAppEnvResponseBodyEnvChangeDetail {
	s.StartTime = &v
	return s
}

func (s *CreateAppEnvResponseBodyEnvChangeDetail) SetChangeId(v string) *CreateAppEnvResponseBodyEnvChangeDetail {
	s.ChangeId = &v
	return s
}

func (s *CreateAppEnvResponseBodyEnvChangeDetail) SetEnvId(v string) *CreateAppEnvResponseBodyEnvChangeDetail {
	s.EnvId = &v
	return s
}

func (s *CreateAppEnvResponseBodyEnvChangeDetail) SetOperations(v *CreateAppEnvResponseBodyEnvChangeDetailOperations) *CreateAppEnvResponseBodyEnvChangeDetail {
	s.Operations = v
	return s
}

type CreateAppEnvResponseBodyEnvChangeDetailOperations struct {
	Operation []*CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Repeated"`
}

func (s CreateAppEnvResponseBodyEnvChangeDetailOperations) String() string {
	return tea.Prettify(s)
}

func (s CreateAppEnvResponseBodyEnvChangeDetailOperations) GoString() string {
	return s.String()
}

func (s *CreateAppEnvResponseBodyEnvChangeDetailOperations) SetOperation(v []*CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation) *CreateAppEnvResponseBodyEnvChangeDetailOperations {
	s.Operation = v
	return s
}

type CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation struct {
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation) String() string {
	return tea.Prettify(s)
}

func (s CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation) GoString() string {
	return s.String()
}

func (s *CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationDescription(v string) *CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationDescription = &v
	return s
}

func (s *CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationType(v string) *CreateAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationType = &v
	return s
}

type CreateAppEnvResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppEnvResponse) GoString() string {
	return s.String()
}

func (s *CreateAppEnvResponse) SetHeaders(v map[string]*string) *CreateAppEnvResponse {
	s.Headers = v
	return s
}

func (s *CreateAppEnvResponse) SetBody(v *CreateAppEnvResponseBody) *CreateAppEnvResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppDescription     *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	CategoryName       *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	UsingSharedStorage *bool   `json:"UsingSharedStorage,omitempty" xml:"UsingSharedStorage,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetAppDescription(v string) *CreateApplicationRequest {
	s.AppDescription = &v
	return s
}

func (s *CreateApplicationRequest) SetCategoryName(v string) *CreateApplicationRequest {
	s.CategoryName = &v
	return s
}

func (s *CreateApplicationRequest) SetUsingSharedStorage(v bool) *CreateApplicationRequest {
	s.UsingSharedStorage = &v
	return s
}

type CreateApplicationResponseBody struct {
	Message     *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Application *CreateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetCode(v string) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody {
	s.Application = v
	return s
}

type CreateApplicationResponseBodyApplication struct {
	CreateUsername     *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUsername     *string `json:"UpdateUsername,omitempty" xml:"UpdateUsername,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CategoryName       *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	UsingSharedStorage *bool   `json:"UsingSharedStorage,omitempty" xml:"UsingSharedStorage,omitempty"`
	AppDescription     *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
}

func (s CreateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplication) SetCreateUsername(v string) *CreateApplicationResponseBodyApplication {
	s.CreateUsername = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppName(v string) *CreateApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateTime(v int64) *CreateApplicationResponseBodyApplication {
	s.UpdateTime = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateUsername(v string) *CreateApplicationResponseBodyApplication {
	s.UpdateUsername = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetCreateTime(v int64) *CreateApplicationResponseBodyApplication {
	s.CreateTime = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppId(v string) *CreateApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetCategoryName(v string) *CreateApplicationResponseBodyApplication {
	s.CategoryName = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUsingSharedStorage(v bool) *CreateApplicationResponseBodyApplication {
	s.UsingSharedStorage = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppDescription(v string) *CreateApplicationResponseBodyApplication {
	s.AppDescription = &v
	return s
}

type CreateApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateConfigTemplateRequest struct {
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StackId             *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	SourceTemplateId    *string `json:"SourceTemplateId,omitempty" xml:"SourceTemplateId,omitempty"`
	SourceEnvId         *string `json:"SourceEnvId,omitempty" xml:"SourceEnvId,omitempty"`
	ProfileName         *string `json:"ProfileName,omitempty" xml:"ProfileName,omitempty"`
	PkgVersionId        *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	OptionSettings      *string `json:"OptionSettings,omitempty" xml:"OptionSettings,omitempty"`
}

func (s CreateConfigTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateRequest) SetTemplateName(v string) *CreateConfigTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetTemplateDescription(v string) *CreateConfigTemplateRequest {
	s.TemplateDescription = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetAppId(v string) *CreateConfigTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetStackId(v string) *CreateConfigTemplateRequest {
	s.StackId = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetSourceTemplateId(v string) *CreateConfigTemplateRequest {
	s.SourceTemplateId = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetSourceEnvId(v string) *CreateConfigTemplateRequest {
	s.SourceEnvId = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetProfileName(v string) *CreateConfigTemplateRequest {
	s.ProfileName = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetPkgVersionId(v string) *CreateConfigTemplateRequest {
	s.PkgVersionId = &v
	return s
}

func (s *CreateConfigTemplateRequest) SetOptionSettings(v string) *CreateConfigTemplateRequest {
	s.OptionSettings = &v
	return s
}

type CreateConfigTemplateResponseBody struct {
	ConfigTemplate *CreateConfigTemplateResponseBodyConfigTemplate `json:"ConfigTemplate,omitempty" xml:"ConfigTemplate,omitempty" type:"Struct"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateConfigTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateResponseBody) SetConfigTemplate(v *CreateConfigTemplateResponseBodyConfigTemplate) *CreateConfigTemplateResponseBody {
	s.ConfigTemplate = v
	return s
}

func (s *CreateConfigTemplateResponseBody) SetMessage(v string) *CreateConfigTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConfigTemplateResponseBody) SetRequestId(v string) *CreateConfigTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigTemplateResponseBody) SetCode(v string) *CreateConfigTemplateResponseBody {
	s.Code = &v
	return s
}

type CreateConfigTemplateResponseBodyConfigTemplate struct {
	UpdateTime          *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateType        *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	StackName           *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	PkgVersionId        *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StackId             *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PkgVersionLabel     *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateConfigTemplateResponseBodyConfigTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigTemplateResponseBodyConfigTemplate) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetUpdateTime(v int64) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.UpdateTime = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetCreateTime(v int64) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.CreateTime = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetTemplateType(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.TemplateType = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetStackName(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.StackName = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetPkgVersionId(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.PkgVersionId = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetTemplateName(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetTemplateDescription(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.TemplateDescription = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetAppName(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.AppName = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetStackId(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.StackId = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetPkgVersionLabel(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.PkgVersionLabel = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetAppId(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.AppId = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyConfigTemplate) SetTemplateId(v string) *CreateConfigTemplateResponseBodyConfigTemplate {
	s.TemplateId = &v
	return s
}

type CreateConfigTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateResponse) SetHeaders(v map[string]*string) *CreateConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigTemplateResponse) SetBody(v *CreateConfigTemplateResponseBody) *CreateConfigTemplateResponse {
	s.Body = v
	return s
}

type CreateOrderRequest struct {
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetProductName(v string) *CreateOrderRequest {
	s.ProductName = &v
	return s
}

type CreateOrderResponseBody struct {
	OrderDetail *CreateOrderResponseBodyOrderDetail `json:"OrderDetail,omitempty" xml:"OrderDetail,omitempty" type:"Struct"`
	Message     *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetOrderDetail(v *CreateOrderResponseBodyOrderDetail) *CreateOrderResponseBody {
	s.OrderDetail = v
	return s
}

func (s *CreateOrderResponseBody) SetMessage(v string) *CreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetCode(v string) *CreateOrderResponseBody {
	s.Code = &v
	return s
}

type CreateOrderResponseBodyOrderDetail struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateOrderResponseBodyOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyOrderDetail) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyOrderDetail) SetData(v string) *CreateOrderResponseBodyOrderDetail {
	s.Data = &v
	return s
}

func (s *CreateOrderResponseBodyOrderDetail) SetRequestId(v string) *CreateOrderResponseBodyOrderDetail {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBodyOrderDetail) SetSuccess(v bool) *CreateOrderResponseBodyOrderDetail {
	s.Success = &v
	return s
}

func (s *CreateOrderResponseBodyOrderDetail) SetCode(v string) *CreateOrderResponseBodyOrderDetail {
	s.Code = &v
	return s
}

func (s *CreateOrderResponseBodyOrderDetail) SetMessage(v string) *CreateOrderResponseBodyOrderDetail {
	s.Message = &v
	return s
}

type CreateOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CreatePkgVersionRequest struct {
	PkgVersionLabel       *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	PkgVersionDescription *string `json:"PkgVersionDescription,omitempty" xml:"PkgVersionDescription,omitempty"`
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PackageSource         *string `json:"PackageSource,omitempty" xml:"PackageSource,omitempty"`
}

func (s CreatePkgVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePkgVersionRequest) GoString() string {
	return s.String()
}

func (s *CreatePkgVersionRequest) SetPkgVersionLabel(v string) *CreatePkgVersionRequest {
	s.PkgVersionLabel = &v
	return s
}

func (s *CreatePkgVersionRequest) SetPkgVersionDescription(v string) *CreatePkgVersionRequest {
	s.PkgVersionDescription = &v
	return s
}

func (s *CreatePkgVersionRequest) SetAppId(v string) *CreatePkgVersionRequest {
	s.AppId = &v
	return s
}

func (s *CreatePkgVersionRequest) SetPackageSource(v string) *CreatePkgVersionRequest {
	s.PackageSource = &v
	return s
}

type CreatePkgVersionResponseBody struct {
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PkgVersion *CreatePkgVersionResponseBodyPkgVersion `json:"PkgVersion,omitempty" xml:"PkgVersion,omitempty" type:"Struct"`
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreatePkgVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePkgVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePkgVersionResponseBody) SetMessage(v string) *CreatePkgVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePkgVersionResponseBody) SetRequestId(v string) *CreatePkgVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePkgVersionResponseBody) SetPkgVersion(v *CreatePkgVersionResponseBodyPkgVersion) *CreatePkgVersionResponseBody {
	s.PkgVersion = v
	return s
}

func (s *CreatePkgVersionResponseBody) SetCode(v string) *CreatePkgVersionResponseBody {
	s.Code = &v
	return s
}

type CreatePkgVersionResponseBodyPkgVersion struct {
	AppName               *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	PkgVersionLabel       *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PackageSource         *string `json:"PackageSource,omitempty" xml:"PackageSource,omitempty"`
	PkgVersionId          *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	PkgVersionDescription *string `json:"PkgVersionDescription,omitempty" xml:"PkgVersionDescription,omitempty"`
}

func (s CreatePkgVersionResponseBodyPkgVersion) String() string {
	return tea.Prettify(s)
}

func (s CreatePkgVersionResponseBodyPkgVersion) GoString() string {
	return s.String()
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetAppName(v string) *CreatePkgVersionResponseBodyPkgVersion {
	s.AppName = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetUpdateTime(v int64) *CreatePkgVersionResponseBodyPkgVersion {
	s.UpdateTime = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetPkgVersionLabel(v string) *CreatePkgVersionResponseBodyPkgVersion {
	s.PkgVersionLabel = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetCreateTime(v int64) *CreatePkgVersionResponseBodyPkgVersion {
	s.CreateTime = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetAppId(v string) *CreatePkgVersionResponseBodyPkgVersion {
	s.AppId = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetPackageSource(v string) *CreatePkgVersionResponseBodyPkgVersion {
	s.PackageSource = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetPkgVersionId(v string) *CreatePkgVersionResponseBodyPkgVersion {
	s.PkgVersionId = &v
	return s
}

func (s *CreatePkgVersionResponseBodyPkgVersion) SetPkgVersionDescription(v string) *CreatePkgVersionResponseBodyPkgVersion {
	s.PkgVersionDescription = &v
	return s
}

type CreatePkgVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePkgVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePkgVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePkgVersionResponse) GoString() string {
	return s.String()
}

func (s *CreatePkgVersionResponse) SetHeaders(v map[string]*string) *CreatePkgVersionResponse {
	s.Headers = v
	return s
}

func (s *CreatePkgVersionResponse) SetBody(v *CreatePkgVersionResponseBody) *CreatePkgVersionResponse {
	s.Body = v
	return s
}

type CreateStorageResponseBody struct {
	Storage   *CreateStorageResponseBodyStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStorageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageResponseBody) SetStorage(v *CreateStorageResponseBodyStorage) *CreateStorageResponseBody {
	s.Storage = v
	return s
}

func (s *CreateStorageResponseBody) SetMessage(v string) *CreateStorageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStorageResponseBody) SetRequestId(v string) *CreateStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageResponseBody) SetCode(v string) *CreateStorageResponseBody {
	s.Code = &v
	return s
}

type CreateStorageResponseBodyStorage struct {
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s CreateStorageResponseBodyStorage) String() string {
	return tea.Prettify(s)
}

func (s CreateStorageResponseBodyStorage) GoString() string {
	return s.String()
}

func (s *CreateStorageResponseBodyStorage) SetUpdateTime(v int64) *CreateStorageResponseBodyStorage {
	s.UpdateTime = &v
	return s
}

func (s *CreateStorageResponseBodyStorage) SetCreateTime(v int64) *CreateStorageResponseBodyStorage {
	s.CreateTime = &v
	return s
}

func (s *CreateStorageResponseBodyStorage) SetBucketName(v string) *CreateStorageResponseBodyStorage {
	s.BucketName = &v
	return s
}

type CreateStorageResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStorageResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageResponse) SetHeaders(v map[string]*string) *CreateStorageResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageResponse) SetBody(v *CreateStorageResponseBody) *CreateStorageResponse {
	s.Body = v
	return s
}

type DeleteAppEnvRequest struct {
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DeleteAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvRequest) SetEnvId(v string) *DeleteAppEnvRequest {
	s.EnvId = &v
	return s
}

type DeleteAppEnvResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvResponseBody) SetMessage(v string) *DeleteAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppEnvResponseBody) SetRequestId(v string) *DeleteAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppEnvResponseBody) SetCode(v string) *DeleteAppEnvResponseBody {
	s.Code = &v
	return s
}

type DeleteAppEnvResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppEnvResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppEnvResponse) SetHeaders(v map[string]*string) *DeleteAppEnvResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppEnvResponse) SetBody(v *DeleteAppEnvResponseBody) *DeleteAppEnvResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

type DeleteApplicationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetCode(v string) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeleteChangeRequest struct {
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s DeleteChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeRequest) GoString() string {
	return s.String()
}

func (s *DeleteChangeRequest) SetChangeId(v string) *DeleteChangeRequest {
	s.ChangeId = &v
	return s
}

type DeleteChangeResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChangeResponseBody) SetMessage(v string) *DeleteChangeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChangeResponseBody) SetRequestId(v string) *DeleteChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChangeResponseBody) SetCode(v string) *DeleteChangeResponseBody {
	s.Code = &v
	return s
}

type DeleteChangeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeResponse) GoString() string {
	return s.String()
}

func (s *DeleteChangeResponse) SetHeaders(v map[string]*string) *DeleteChangeResponse {
	s.Headers = v
	return s
}

func (s *DeleteChangeResponse) SetBody(v *DeleteChangeResponseBody) *DeleteChangeResponse {
	s.Body = v
	return s
}

type DeleteConfigTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteConfigTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigTemplateRequest) SetTemplateId(v string) *DeleteConfigTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteConfigTemplateResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteConfigTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigTemplateResponseBody) SetMessage(v string) *DeleteConfigTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigTemplateResponseBody) SetRequestId(v string) *DeleteConfigTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigTemplateResponseBody) SetCode(v string) *DeleteConfigTemplateResponseBody {
	s.Code = &v
	return s
}

type DeleteConfigTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigTemplateResponse) SetHeaders(v map[string]*string) *DeleteConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigTemplateResponse) SetBody(v *DeleteConfigTemplateResponseBody) *DeleteConfigTemplateResponse {
	s.Body = v
	return s
}

type DeletePkgVersionRequest struct {
	PkgVersionId *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
}

func (s DeletePkgVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePkgVersionRequest) GoString() string {
	return s.String()
}

func (s *DeletePkgVersionRequest) SetPkgVersionId(v string) *DeletePkgVersionRequest {
	s.PkgVersionId = &v
	return s
}

type DeletePkgVersionResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeletePkgVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePkgVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePkgVersionResponseBody) SetMessage(v string) *DeletePkgVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePkgVersionResponseBody) SetRequestId(v string) *DeletePkgVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePkgVersionResponseBody) SetCode(v string) *DeletePkgVersionResponseBody {
	s.Code = &v
	return s
}

type DeletePkgVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePkgVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePkgVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePkgVersionResponse) GoString() string {
	return s.String()
}

func (s *DeletePkgVersionResponse) SetHeaders(v map[string]*string) *DeletePkgVersionResponse {
	s.Headers = v
	return s
}

func (s *DeletePkgVersionResponse) SetBody(v *DeletePkgVersionResponseBody) *DeletePkgVersionResponse {
	s.Body = v
	return s
}

type DeployAppEnvRequest struct {
	EnvId               *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	BatchSize           *int32  `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	BatchPercent        *int32  `json:"BatchPercent,omitempty" xml:"BatchPercent,omitempty"`
	BatchInterval       *int32  `json:"BatchInterval,omitempty" xml:"BatchInterval,omitempty"`
	PauseBetweenBatches *bool   `json:"PauseBetweenBatches,omitempty" xml:"PauseBetweenBatches,omitempty"`
	PkgVersionId        *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
}

func (s DeployAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployAppEnvRequest) GoString() string {
	return s.String()
}

func (s *DeployAppEnvRequest) SetEnvId(v string) *DeployAppEnvRequest {
	s.EnvId = &v
	return s
}

func (s *DeployAppEnvRequest) SetBatchSize(v int32) *DeployAppEnvRequest {
	s.BatchSize = &v
	return s
}

func (s *DeployAppEnvRequest) SetBatchPercent(v int32) *DeployAppEnvRequest {
	s.BatchPercent = &v
	return s
}

func (s *DeployAppEnvRequest) SetBatchInterval(v int32) *DeployAppEnvRequest {
	s.BatchInterval = &v
	return s
}

func (s *DeployAppEnvRequest) SetPauseBetweenBatches(v bool) *DeployAppEnvRequest {
	s.PauseBetweenBatches = &v
	return s
}

func (s *DeployAppEnvRequest) SetPkgVersionId(v string) *DeployAppEnvRequest {
	s.PkgVersionId = &v
	return s
}

type DeployAppEnvResponseBody struct {
	EnvChange *DeployAppEnvResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeployAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *DeployAppEnvResponseBody) SetEnvChange(v *DeployAppEnvResponseBodyEnvChange) *DeployAppEnvResponseBody {
	s.EnvChange = v
	return s
}

func (s *DeployAppEnvResponseBody) SetMessage(v string) *DeployAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *DeployAppEnvResponseBody) SetRequestId(v string) *DeployAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployAppEnvResponseBody) SetCode(v string) *DeployAppEnvResponseBody {
	s.Code = &v
	return s
}

type DeployAppEnvResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DeployAppEnvResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s DeployAppEnvResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *DeployAppEnvResponseBodyEnvChange) SetStartTime(v string) *DeployAppEnvResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *DeployAppEnvResponseBodyEnvChange) SetChangeId(v string) *DeployAppEnvResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *DeployAppEnvResponseBodyEnvChange) SetEnvId(v string) *DeployAppEnvResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type DeployAppEnvResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployAppEnvResponse) GoString() string {
	return s.String()
}

func (s *DeployAppEnvResponse) SetHeaders(v map[string]*string) *DeployAppEnvResponse {
	s.Headers = v
	return s
}

func (s *DeployAppEnvResponse) SetBody(v *DeployAppEnvResponseBody) *DeployAppEnvResponse {
	s.Body = v
	return s
}

type DescribeAppEnvInstanceHealthRequest struct {
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DescribeAppEnvInstanceHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvInstanceHealthRequest) SetEnvId(v string) *DescribeAppEnvInstanceHealthRequest {
	s.EnvId = &v
	return s
}

type DescribeAppEnvInstanceHealthResponseBody struct {
	Message           *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code              *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	EnvInstanceHealth *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth `json:"EnvInstanceHealth,omitempty" xml:"EnvInstanceHealth,omitempty" type:"Struct"`
}

func (s DescribeAppEnvInstanceHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvInstanceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvInstanceHealthResponseBody) SetMessage(v string) *DescribeAppEnvInstanceHealthResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBody) SetRequestId(v string) *DescribeAppEnvInstanceHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBody) SetCode(v string) *DescribeAppEnvInstanceHealthResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBody) SetEnvInstanceHealth(v *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) *DescribeAppEnvInstanceHealthResponseBody {
	s.EnvInstanceHealth = v
	return s
}

type DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth struct {
	EnableHealthCheck  *bool                                                                        `json:"EnableHealthCheck,omitempty" xml:"EnableHealthCheck,omitempty"`
	EnvName            *string                                                                      `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvId              *string                                                                      `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	InstanceHealthList *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList `json:"InstanceHealthList,omitempty" xml:"InstanceHealthList,omitempty" type:"Struct"`
}

func (s DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) SetEnableHealthCheck(v bool) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth {
	s.EnableHealthCheck = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) SetEnvName(v string) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth {
	s.EnvName = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) SetEnvId(v string) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth {
	s.EnvId = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth) SetInstanceHealthList(v *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealth {
	s.InstanceHealthList = v
	return s
}

type DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList struct {
	InstanceHealth []*DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth `json:"InstanceHealth,omitempty" xml:"InstanceHealth,omitempty" type:"Repeated"`
}

func (s DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList) SetInstanceHealth(v []*DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthList {
	s.InstanceHealth = v
	return s
}

type DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth struct {
	AppStatus        *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DisconnectedTime *string `json:"DisconnectedTime,omitempty" xml:"DisconnectedTime,omitempty"`
	AgentStatus      *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
}

func (s DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) SetAppStatus(v string) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth {
	s.AppStatus = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) SetInstanceId(v string) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth {
	s.InstanceId = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) SetDisconnectedTime(v string) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth {
	s.DisconnectedTime = &v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth) SetAgentStatus(v string) *DescribeAppEnvInstanceHealthResponseBodyEnvInstanceHealthInstanceHealthListInstanceHealth {
	s.AgentStatus = &v
	return s
}

type DescribeAppEnvInstanceHealthResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppEnvInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppEnvInstanceHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvInstanceHealthResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvInstanceHealthResponse) SetHeaders(v map[string]*string) *DescribeAppEnvInstanceHealthResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppEnvInstanceHealthResponse) SetBody(v *DescribeAppEnvInstanceHealthResponseBody) *DescribeAppEnvInstanceHealthResponse {
	s.Body = v
	return s
}

type DescribeAppEnvsRequest struct {
	EnvId             *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	IncludeTerminated *bool   `json:"IncludeTerminated,omitempty" xml:"IncludeTerminated,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	EnvName           *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	EnvSearch         *string `json:"EnvSearch,omitempty" xml:"EnvSearch,omitempty"`
	RecentUpdated     *bool   `json:"RecentUpdated,omitempty" xml:"RecentUpdated,omitempty"`
	StackSearch       *string `json:"StackSearch,omitempty" xml:"StackSearch,omitempty"`
}

func (s DescribeAppEnvsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvsRequest) SetEnvId(v string) *DescribeAppEnvsRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetAppId(v string) *DescribeAppEnvsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetIncludeTerminated(v bool) *DescribeAppEnvsRequest {
	s.IncludeTerminated = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetPageSize(v int32) *DescribeAppEnvsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetPageNumber(v int32) *DescribeAppEnvsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetEnvName(v string) *DescribeAppEnvsRequest {
	s.EnvName = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetEnvSearch(v string) *DescribeAppEnvsRequest {
	s.EnvSearch = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetRecentUpdated(v bool) *DescribeAppEnvsRequest {
	s.RecentUpdated = &v
	return s
}

func (s *DescribeAppEnvsRequest) SetStackSearch(v string) *DescribeAppEnvsRequest {
	s.StackSearch = &v
	return s
}

type DescribeAppEnvsResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	AppEnvs    *DescribeAppEnvsResponseBodyAppEnvs `json:"AppEnvs,omitempty" xml:"AppEnvs,omitempty" type:"Struct"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeAppEnvsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvsResponseBody) SetTotalCount(v int32) *DescribeAppEnvsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppEnvsResponseBody) SetRequestId(v string) *DescribeAppEnvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppEnvsResponseBody) SetMessage(v string) *DescribeAppEnvsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppEnvsResponseBody) SetPageSize(v int32) *DescribeAppEnvsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppEnvsResponseBody) SetPageNumber(v int32) *DescribeAppEnvsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppEnvsResponseBody) SetAppEnvs(v *DescribeAppEnvsResponseBodyAppEnvs) *DescribeAppEnvsResponseBody {
	s.AppEnvs = v
	return s
}

func (s *DescribeAppEnvsResponseBody) SetCode(v string) *DescribeAppEnvsResponseBody {
	s.Code = &v
	return s
}

type DescribeAppEnvsResponseBodyAppEnvs struct {
	AppEnv []*DescribeAppEnvsResponseBodyAppEnvsAppEnv `json:"AppEnv,omitempty" xml:"AppEnv,omitempty" type:"Repeated"`
}

func (s DescribeAppEnvsResponseBodyAppEnvs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvsResponseBodyAppEnvs) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvsResponseBodyAppEnvs) SetAppEnv(v []*DescribeAppEnvsResponseBodyAppEnvsAppEnv) *DescribeAppEnvsResponseBodyAppEnvs {
	s.AppEnv = v
	return s
}

type DescribeAppEnvsResponseBodyAppEnvsAppEnv struct {
	UpdateTime           *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	TotalInstances       *int64  `json:"TotalInstances,omitempty" xml:"TotalInstances,omitempty"`
	PkgVersionStorageKey *string `json:"PkgVersionStorageKey,omitempty" xml:"PkgVersionStorageKey,omitempty"`
	LatestChangeId       *string `json:"LatestChangeId,omitempty" xml:"LatestChangeId,omitempty"`
	EnvStatus            *string `json:"EnvStatus,omitempty" xml:"EnvStatus,omitempty"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastEnvStatus        *string `json:"LastEnvStatus,omitempty" xml:"LastEnvStatus,omitempty"`
	PkgVersionId         *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	EnvDescription       *string `json:"EnvDescription,omitempty" xml:"EnvDescription,omitempty"`
	ApplyingChange       *bool   `json:"ApplyingChange,omitempty" xml:"ApplyingChange,omitempty"`
	EnvType              *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateUsername       *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DataRoot             *string `json:"DataRoot,omitempty" xml:"DataRoot,omitempty"`
	StorageBase          *string `json:"StorageBase,omitempty" xml:"StorageBase,omitempty"`
	UpdateUsername       *string `json:"UpdateUsername,omitempty" xml:"UpdateUsername,omitempty"`
	EnvName              *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	LogBase              *string `json:"LogBase,omitempty" xml:"LogBase,omitempty"`
	StackName            *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	CategoryName         *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	UsingSharedStorage   *bool   `json:"UsingSharedStorage,omitempty" xml:"UsingSharedStorage,omitempty"`
	ChangeBanner         *string `json:"ChangeBanner,omitempty" xml:"ChangeBanner,omitempty"`
	StackId              *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PkgVersionLabel      *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	EnvId                *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	AbortingChange       *bool   `json:"AbortingChange,omitempty" xml:"AbortingChange,omitempty"`
}

func (s DescribeAppEnvsResponseBodyAppEnvsAppEnv) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvsResponseBodyAppEnvsAppEnv) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetUpdateTime(v int64) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.UpdateTime = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetTotalInstances(v int64) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.TotalInstances = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetPkgVersionStorageKey(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.PkgVersionStorageKey = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetLatestChangeId(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.LatestChangeId = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetEnvStatus(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.EnvStatus = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetCreateTime(v int64) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetLastEnvStatus(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.LastEnvStatus = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetPkgVersionId(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.PkgVersionId = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetEnvDescription(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.EnvDescription = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetApplyingChange(v bool) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.ApplyingChange = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetEnvType(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.EnvType = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetAppName(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.AppName = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetCreateUsername(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.CreateUsername = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetAppId(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.AppId = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetDataRoot(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.DataRoot = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetStorageBase(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.StorageBase = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetUpdateUsername(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.UpdateUsername = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetEnvName(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.EnvName = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetLogBase(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.LogBase = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetStackName(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.StackName = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetCategoryName(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.CategoryName = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetUsingSharedStorage(v bool) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.UsingSharedStorage = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetChangeBanner(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.ChangeBanner = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetStackId(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.StackId = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetPkgVersionLabel(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.PkgVersionLabel = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetEnvId(v string) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.EnvId = &v
	return s
}

func (s *DescribeAppEnvsResponseBodyAppEnvsAppEnv) SetAbortingChange(v bool) *DescribeAppEnvsResponseBodyAppEnvsAppEnv {
	s.AbortingChange = &v
	return s
}

type DescribeAppEnvsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppEnvsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvsResponse) SetHeaders(v map[string]*string) *DescribeAppEnvsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppEnvsResponse) SetBody(v *DescribeAppEnvsResponseBody) *DescribeAppEnvsResponse {
	s.Body = v
	return s
}

type DescribeAppEnvStatusRequest struct {
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DescribeAppEnvStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvStatusRequest) SetEnvId(v string) *DescribeAppEnvStatusRequest {
	s.EnvId = &v
	return s
}

type DescribeAppEnvStatusResponseBody struct {
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvStatus *DescribeAppEnvStatusResponseBodyEnvStatus `json:"EnvStatus,omitempty" xml:"EnvStatus,omitempty" type:"Struct"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeAppEnvStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvStatusResponseBody) SetMessage(v string) *DescribeAppEnvStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBody) SetRequestId(v string) *DescribeAppEnvStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBody) SetEnvStatus(v *DescribeAppEnvStatusResponseBodyEnvStatus) *DescribeAppEnvStatusResponseBody {
	s.EnvStatus = v
	return s
}

func (s *DescribeAppEnvStatusResponseBody) SetCode(v string) *DescribeAppEnvStatusResponseBody {
	s.Code = &v
	return s
}

type DescribeAppEnvStatusResponseBodyEnvStatus struct {
	ChangeBanner        *string                                                       `json:"ChangeBanner,omitempty" xml:"ChangeBanner,omitempty"`
	LatestChangeId      *string                                                       `json:"LatestChangeId,omitempty" xml:"LatestChangeId,omitempty"`
	EnvStatus           *string                                                       `json:"EnvStatus,omitempty" xml:"EnvStatus,omitempty"`
	EnvName             *string                                                       `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	InstanceAgentStatus *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus `json:"InstanceAgentStatus,omitempty" xml:"InstanceAgentStatus,omitempty" type:"Struct"`
	LastEnvStatus       *string                                                       `json:"LastEnvStatus,omitempty" xml:"LastEnvStatus,omitempty"`
	InstanceAppStatus   *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus   `json:"InstanceAppStatus,omitempty" xml:"InstanceAppStatus,omitempty" type:"Struct"`
	EnvId               *string                                                       `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	AbortingChange      *bool                                                         `json:"AbortingChange,omitempty" xml:"AbortingChange,omitempty"`
	ApplyingChange      *bool                                                         `json:"ApplyingChange,omitempty" xml:"ApplyingChange,omitempty"`
}

func (s DescribeAppEnvStatusResponseBodyEnvStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvStatusResponseBodyEnvStatus) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetChangeBanner(v string) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.ChangeBanner = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetLatestChangeId(v string) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.LatestChangeId = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetEnvStatus(v string) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.EnvStatus = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetEnvName(v string) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.EnvName = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetInstanceAgentStatus(v *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.InstanceAgentStatus = v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetLastEnvStatus(v string) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.LastEnvStatus = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetInstanceAppStatus(v *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.InstanceAppStatus = v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetEnvId(v string) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.EnvId = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetAbortingChange(v bool) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.AbortingChange = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatus) SetApplyingChange(v bool) *DescribeAppEnvStatusResponseBodyEnvStatus {
	s.ApplyingChange = &v
	return s
}

type DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus struct {
	ConnectedInstances    *int32 `json:"ConnectedInstances,omitempty" xml:"ConnectedInstances,omitempty"`
	DisconnectedInstances *int32 `json:"DisconnectedInstances,omitempty" xml:"DisconnectedInstances,omitempty"`
}

func (s DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus) SetConnectedInstances(v int32) *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus {
	s.ConnectedInstances = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus) SetDisconnectedInstances(v int32) *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAgentStatus {
	s.DisconnectedInstances = &v
	return s
}

type DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus struct {
	HealthyInstances   *int32 `json:"HealthyInstances,omitempty" xml:"HealthyInstances,omitempty"`
	StoppedInstances   *int32 `json:"StoppedInstances,omitempty" xml:"StoppedInstances,omitempty"`
	UnhealthyInstances *int32 `json:"UnhealthyInstances,omitempty" xml:"UnhealthyInstances,omitempty"`
	UnknownInstances   *int32 `json:"UnknownInstances,omitempty" xml:"UnknownInstances,omitempty"`
}

func (s DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) SetHealthyInstances(v int32) *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus {
	s.HealthyInstances = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) SetStoppedInstances(v int32) *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus {
	s.StoppedInstances = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) SetUnhealthyInstances(v int32) *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus {
	s.UnhealthyInstances = &v
	return s
}

func (s *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus) SetUnknownInstances(v int32) *DescribeAppEnvStatusResponseBodyEnvStatusInstanceAppStatus {
	s.UnknownInstances = &v
	return s
}

type DescribeAppEnvStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppEnvStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppEnvStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppEnvStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppEnvStatusResponse) SetHeaders(v map[string]*string) *DescribeAppEnvStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppEnvStatusResponse) SetBody(v *DescribeAppEnvStatusResponseBody) *DescribeAppEnvStatusResponse {
	s.Body = v
	return s
}

type DescribeApplicationsRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppSearch      *string `json:"AppSearch,omitempty" xml:"AppSearch,omitempty"`
	EnvSearch      *string `json:"EnvSearch,omitempty" xml:"EnvSearch,omitempty"`
	StackSearch    *string `json:"StackSearch,omitempty" xml:"StackSearch,omitempty"`
	CategorySearch *string `json:"CategorySearch,omitempty" xml:"CategorySearch,omitempty"`
}

func (s DescribeApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsRequest) SetAppId(v string) *DescribeApplicationsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationsRequest) SetPageSize(v int32) *DescribeApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationsRequest) SetPageNumber(v int32) *DescribeApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationsRequest) SetAppName(v string) *DescribeApplicationsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeApplicationsRequest) SetAppSearch(v string) *DescribeApplicationsRequest {
	s.AppSearch = &v
	return s
}

func (s *DescribeApplicationsRequest) SetEnvSearch(v string) *DescribeApplicationsRequest {
	s.EnvSearch = &v
	return s
}

func (s *DescribeApplicationsRequest) SetStackSearch(v string) *DescribeApplicationsRequest {
	s.StackSearch = &v
	return s
}

func (s *DescribeApplicationsRequest) SetCategorySearch(v string) *DescribeApplicationsRequest {
	s.CategorySearch = &v
	return s
}

type DescribeApplicationsResponseBody struct {
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Applications *DescribeApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBody) SetTotalCount(v int32) *DescribeApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetRequestId(v string) *DescribeApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetMessage(v string) *DescribeApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetPageSize(v int32) *DescribeApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetApplications(v *DescribeApplicationsResponseBodyApplications) *DescribeApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *DescribeApplicationsResponseBody) SetPageNumber(v int32) *DescribeApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetCode(v string) *DescribeApplicationsResponseBody {
	s.Code = &v
	return s
}

type DescribeApplicationsResponseBodyApplications struct {
	Application []*DescribeApplicationsResponseBodyApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s DescribeApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBodyApplications) SetApplication(v []*DescribeApplicationsResponseBodyApplicationsApplication) *DescribeApplicationsResponseBodyApplications {
	s.Application = v
	return s
}

type DescribeApplicationsResponseBodyApplicationsApplication struct {
	TotalEnvs          *int32  `json:"TotalEnvs,omitempty" xml:"TotalEnvs,omitempty"`
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUsername     *string `json:"UpdateUsername,omitempty" xml:"UpdateUsername,omitempty"`
	RunningEnvs        *int32  `json:"RunningEnvs,omitempty" xml:"RunningEnvs,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CategoryName       *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	UsingSharedStorage *string `json:"UsingSharedStorage,omitempty" xml:"UsingSharedStorage,omitempty"`
	CreateUsername     *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TerminatedEnvs     *int32  `json:"TerminatedEnvs,omitempty" xml:"TerminatedEnvs,omitempty"`
	AppDescription     *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
}

func (s DescribeApplicationsResponseBodyApplicationsApplication) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationsResponseBodyApplicationsApplication) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetTotalEnvs(v int32) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.TotalEnvs = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetUpdateTime(v int64) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.UpdateTime = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetUpdateUsername(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.UpdateUsername = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetRunningEnvs(v int32) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.RunningEnvs = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetCreateTime(v int64) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetCategoryName(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.CategoryName = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetUsingSharedStorage(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.UsingSharedStorage = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetCreateUsername(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.CreateUsername = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetAppName(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.AppName = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetAppId(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetTerminatedEnvs(v int32) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.TerminatedEnvs = &v
	return s
}

func (s *DescribeApplicationsResponseBodyApplicationsApplication) SetAppDescription(v string) *DescribeApplicationsResponseBodyApplicationsApplication {
	s.AppDescription = &v
	return s
}

type DescribeApplicationsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponse) SetHeaders(v map[string]*string) *DescribeApplicationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationsResponse) SetBody(v *DescribeApplicationsResponseBody) *DescribeApplicationsResponse {
	s.Body = v
	return s
}

type DescribeCategoriesResponseBody struct {
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Categories *DescribeCategoriesResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Struct"`
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoriesResponseBody) SetMessage(v string) *DescribeCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCategoriesResponseBody) SetRequestId(v string) *DescribeCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoriesResponseBody) SetCategories(v *DescribeCategoriesResponseBodyCategories) *DescribeCategoriesResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeCategoriesResponseBody) SetCode(v string) *DescribeCategoriesResponseBody {
	s.Code = &v
	return s
}

type DescribeCategoriesResponseBodyCategories struct {
	Category []*DescribeCategoriesResponseBodyCategoriesCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s DescribeCategoriesResponseBodyCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoriesResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *DescribeCategoriesResponseBodyCategories) SetCategory(v []*DescribeCategoriesResponseBodyCategoriesCategory) *DescribeCategoriesResponseBodyCategories {
	s.Category = v
	return s
}

type DescribeCategoriesResponseBodyCategoriesCategory struct {
	CategoryLogo        *string                                                       `json:"CategoryLogo,omitempty" xml:"CategoryLogo,omitempty"`
	UpdateTime          *string                                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DemoProjects        *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects `json:"DemoProjects,omitempty" xml:"DemoProjects,omitempty" type:"Struct"`
	CreateTime          *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CategoryId          *string                                                       `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName        *string                                                       `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CategoryDescription *string                                                       `json:"CategoryDescription,omitempty" xml:"CategoryDescription,omitempty"`
}

func (s DescribeCategoriesResponseBodyCategoriesCategory) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoriesResponseBodyCategoriesCategory) GoString() string {
	return s.String()
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetCategoryLogo(v string) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.CategoryLogo = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetUpdateTime(v string) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetDemoProjects(v *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.DemoProjects = v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetCreateTime(v string) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.CreateTime = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetCategoryId(v string) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.CategoryId = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetCategoryName(v string) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.CategoryName = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategory) SetCategoryDescription(v string) *DescribeCategoriesResponseBodyCategoriesCategory {
	s.CategoryDescription = &v
	return s
}

type DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects struct {
	DemoProject []*DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject `json:"DemoProject,omitempty" xml:"DemoProject,omitempty" type:"Repeated"`
}

func (s DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects) GoString() string {
	return s.String()
}

func (s *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects) SetDemoProject(v []*DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjects {
	s.DemoProject = v
	return s
}

type DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject struct {
	SourceUrl          *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	PackageDownloadUrl *string `json:"PackageDownloadUrl,omitempty" xml:"PackageDownloadUrl,omitempty"`
	PackageUrl         *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) GoString() string {
	return s.String()
}

func (s *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) SetSourceUrl(v string) *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject {
	s.SourceUrl = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) SetPackageDownloadUrl(v string) *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject {
	s.PackageDownloadUrl = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) SetPackageUrl(v string) *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject {
	s.PackageUrl = &v
	return s
}

func (s *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject) SetRegionId(v string) *DescribeCategoriesResponseBodyCategoriesCategoryDemoProjectsDemoProject {
	s.RegionId = &v
	return s
}

type DescribeCategoriesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoriesResponse) SetHeaders(v map[string]*string) *DescribeCategoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoriesResponse) SetBody(v *DescribeCategoriesResponseBody) *DescribeCategoriesResponse {
	s.Body = v
	return s
}

type DescribeChangeRequest struct {
	EnvId    *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s DescribeChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangeRequest) SetEnvId(v string) *DescribeChangeRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeChangeRequest) SetChangeId(v string) *DescribeChangeRequest {
	s.ChangeId = &v
	return s
}

type DescribeChangeResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Change    *DescribeChangeResponseBodyChange `json:"Change,omitempty" xml:"Change,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangeResponseBody) SetMessage(v string) *DescribeChangeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeChangeResponseBody) SetRequestId(v string) *DescribeChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangeResponseBody) SetChange(v *DescribeChangeResponseBodyChange) *DescribeChangeResponseBody {
	s.Change = v
	return s
}

func (s *DescribeChangeResponseBody) SetCode(v string) *DescribeChangeResponseBody {
	s.Code = &v
	return s
}

type DescribeChangeResponseBodyChange struct {
	ChangePaused      *bool   `json:"ChangePaused,omitempty" xml:"ChangePaused,omitempty"`
	ChangeDescription *string `json:"ChangeDescription,omitempty" xml:"ChangeDescription,omitempty"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChangeTimedout    *bool   `json:"ChangeTimedout,omitempty" xml:"ChangeTimedout,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChangeMessage     *string `json:"ChangeMessage,omitempty" xml:"ChangeMessage,omitempty"`
	ActionName        *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ChangeFinished    *bool   `json:"ChangeFinished,omitempty" xml:"ChangeFinished,omitempty"`
	CreateUsername    *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	ChangeId          *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	ChangeAborted     *bool   `json:"ChangeAborted,omitempty" xml:"ChangeAborted,omitempty"`
	ChangeSucceed     *bool   `json:"ChangeSucceed,omitempty" xml:"ChangeSucceed,omitempty"`
	EnvId             *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ChangeName        *string `json:"ChangeName,omitempty" xml:"ChangeName,omitempty"`
}

func (s DescribeChangeResponseBodyChange) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeResponseBodyChange) GoString() string {
	return s.String()
}

func (s *DescribeChangeResponseBodyChange) SetChangePaused(v bool) *DescribeChangeResponseBodyChange {
	s.ChangePaused = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeDescription(v string) *DescribeChangeResponseBodyChange {
	s.ChangeDescription = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetFinishTime(v int64) *DescribeChangeResponseBodyChange {
	s.FinishTime = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetUpdateTime(v int64) *DescribeChangeResponseBodyChange {
	s.UpdateTime = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeTimedout(v bool) *DescribeChangeResponseBodyChange {
	s.ChangeTimedout = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetCreateTime(v int64) *DescribeChangeResponseBodyChange {
	s.CreateTime = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeMessage(v string) *DescribeChangeResponseBodyChange {
	s.ChangeMessage = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetActionName(v string) *DescribeChangeResponseBodyChange {
	s.ActionName = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeFinished(v bool) *DescribeChangeResponseBodyChange {
	s.ChangeFinished = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetCreateUsername(v string) *DescribeChangeResponseBodyChange {
	s.CreateUsername = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeId(v string) *DescribeChangeResponseBodyChange {
	s.ChangeId = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeAborted(v bool) *DescribeChangeResponseBodyChange {
	s.ChangeAborted = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeSucceed(v bool) *DescribeChangeResponseBodyChange {
	s.ChangeSucceed = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetEnvId(v string) *DescribeChangeResponseBodyChange {
	s.EnvId = &v
	return s
}

func (s *DescribeChangeResponseBodyChange) SetChangeName(v string) *DescribeChangeResponseBodyChange {
	s.ChangeName = &v
	return s
}

type DescribeChangeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangeResponse) SetHeaders(v map[string]*string) *DescribeChangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangeResponse) SetBody(v *DescribeChangeResponseBody) *DescribeChangeResponse {
	s.Body = v
	return s
}

type DescribeChangesRequest struct {
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeChangesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangesRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangesRequest) SetEnvId(v string) *DescribeChangesRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeChangesRequest) SetActionName(v string) *DescribeChangesRequest {
	s.ActionName = &v
	return s
}

func (s *DescribeChangesRequest) SetPageSize(v int32) *DescribeChangesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeChangesRequest) SetPageNumber(v int32) *DescribeChangesRequest {
	s.PageNumber = &v
	return s
}

type DescribeChangesResponseBody struct {
	Changes    *DescribeChangesResponseBodyChanges `json:"Changes,omitempty" xml:"Changes,omitempty" type:"Struct"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeChangesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangesResponseBody) SetChanges(v *DescribeChangesResponseBodyChanges) *DescribeChangesResponseBody {
	s.Changes = v
	return s
}

func (s *DescribeChangesResponseBody) SetTotalCount(v int32) *DescribeChangesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeChangesResponseBody) SetRequestId(v string) *DescribeChangesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangesResponseBody) SetMessage(v string) *DescribeChangesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeChangesResponseBody) SetPageSize(v int32) *DescribeChangesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeChangesResponseBody) SetPageNumber(v int32) *DescribeChangesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangesResponseBody) SetCode(v string) *DescribeChangesResponseBody {
	s.Code = &v
	return s
}

type DescribeChangesResponseBodyChanges struct {
	Change []*DescribeChangesResponseBodyChangesChange `json:"Change,omitempty" xml:"Change,omitempty" type:"Repeated"`
}

func (s DescribeChangesResponseBodyChanges) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangesResponseBodyChanges) GoString() string {
	return s.String()
}

func (s *DescribeChangesResponseBodyChanges) SetChange(v []*DescribeChangesResponseBodyChangesChange) *DescribeChangesResponseBodyChanges {
	s.Change = v
	return s
}

type DescribeChangesResponseBodyChangesChange struct {
	ChangePaused      *bool   `json:"ChangePaused,omitempty" xml:"ChangePaused,omitempty"`
	ChangeDescription *string `json:"ChangeDescription,omitempty" xml:"ChangeDescription,omitempty"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChangeTimedout    *bool   `json:"ChangeTimedout,omitempty" xml:"ChangeTimedout,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChangeMessage     *string `json:"ChangeMessage,omitempty" xml:"ChangeMessage,omitempty"`
	ActionName        *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ChangeFinished    *bool   `json:"ChangeFinished,omitempty" xml:"ChangeFinished,omitempty"`
	CreateUsername    *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	ChangeId          *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	ChangeAborted     *bool   `json:"ChangeAborted,omitempty" xml:"ChangeAborted,omitempty"`
	ChangeSucceed     *bool   `json:"ChangeSucceed,omitempty" xml:"ChangeSucceed,omitempty"`
	EnvId             *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ChangeName        *string `json:"ChangeName,omitempty" xml:"ChangeName,omitempty"`
}

func (s DescribeChangesResponseBodyChangesChange) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangesResponseBodyChangesChange) GoString() string {
	return s.String()
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangePaused(v bool) *DescribeChangesResponseBodyChangesChange {
	s.ChangePaused = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeDescription(v string) *DescribeChangesResponseBodyChangesChange {
	s.ChangeDescription = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetFinishTime(v int64) *DescribeChangesResponseBodyChangesChange {
	s.FinishTime = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetUpdateTime(v int64) *DescribeChangesResponseBodyChangesChange {
	s.UpdateTime = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeTimedout(v bool) *DescribeChangesResponseBodyChangesChange {
	s.ChangeTimedout = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetCreateTime(v int64) *DescribeChangesResponseBodyChangesChange {
	s.CreateTime = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeMessage(v string) *DescribeChangesResponseBodyChangesChange {
	s.ChangeMessage = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetActionName(v string) *DescribeChangesResponseBodyChangesChange {
	s.ActionName = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeFinished(v bool) *DescribeChangesResponseBodyChangesChange {
	s.ChangeFinished = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetCreateUsername(v string) *DescribeChangesResponseBodyChangesChange {
	s.CreateUsername = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeId(v string) *DescribeChangesResponseBodyChangesChange {
	s.ChangeId = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeAborted(v bool) *DescribeChangesResponseBodyChangesChange {
	s.ChangeAborted = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeSucceed(v bool) *DescribeChangesResponseBodyChangesChange {
	s.ChangeSucceed = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetEnvId(v string) *DescribeChangesResponseBodyChangesChange {
	s.EnvId = &v
	return s
}

func (s *DescribeChangesResponseBodyChangesChange) SetChangeName(v string) *DescribeChangesResponseBodyChangesChange {
	s.ChangeName = &v
	return s
}

type DescribeChangesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeChangesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChangesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangesResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangesResponse) SetHeaders(v map[string]*string) *DescribeChangesResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangesResponse) SetBody(v *DescribeChangesResponseBody) *DescribeChangesResponse {
	s.Body = v
	return s
}

type DescribeConfigIndexRequest struct {
	StackId     *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	EnvId       *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ProfileName *string `json:"ProfileName,omitempty" xml:"ProfileName,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeConfigIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexRequest) SetStackId(v string) *DescribeConfigIndexRequest {
	s.StackId = &v
	return s
}

func (s *DescribeConfigIndexRequest) SetEnvId(v string) *DescribeConfigIndexRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeConfigIndexRequest) SetProfileName(v string) *DescribeConfigIndexRequest {
	s.ProfileName = &v
	return s
}

func (s *DescribeConfigIndexRequest) SetTemplateId(v string) *DescribeConfigIndexRequest {
	s.TemplateId = &v
	return s
}

type DescribeConfigIndexResponseBody struct {
	Message      *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConfigGroups *DescribeConfigIndexResponseBodyConfigGroups `json:"ConfigGroups,omitempty" xml:"ConfigGroups,omitempty" type:"Struct"`
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	StackName    *string                                      `json:"StackName,omitempty" xml:"StackName,omitempty"`
	StackId      *string                                      `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DescribeConfigIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBody) SetMessage(v string) *DescribeConfigIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigIndexResponseBody) SetRequestId(v string) *DescribeConfigIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigIndexResponseBody) SetConfigGroups(v *DescribeConfigIndexResponseBodyConfigGroups) *DescribeConfigIndexResponseBody {
	s.ConfigGroups = v
	return s
}

func (s *DescribeConfigIndexResponseBody) SetCode(v string) *DescribeConfigIndexResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeConfigIndexResponseBody) SetStackName(v string) *DescribeConfigIndexResponseBody {
	s.StackName = &v
	return s
}

func (s *DescribeConfigIndexResponseBody) SetStackId(v string) *DescribeConfigIndexResponseBody {
	s.StackId = &v
	return s
}

type DescribeConfigIndexResponseBodyConfigGroups struct {
	ConfigGroup []*DescribeConfigIndexResponseBodyConfigGroupsConfigGroup `json:"ConfigGroup,omitempty" xml:"ConfigGroup,omitempty" type:"Repeated"`
}

func (s DescribeConfigIndexResponseBodyConfigGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBodyConfigGroups) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBodyConfigGroups) SetConfigGroup(v []*DescribeConfigIndexResponseBodyConfigGroupsConfigGroup) *DescribeConfigIndexResponseBodyConfigGroups {
	s.ConfigGroup = v
	return s
}

type DescribeConfigIndexResponseBodyConfigGroupsConfigGroup struct {
	ConfigPaths *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths `json:"ConfigPaths,omitempty" xml:"ConfigPaths,omitempty" type:"Struct"`
	GroupName   *string                                                            `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupLabel  *string                                                            `json:"GroupLabel,omitempty" xml:"GroupLabel,omitempty"`
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroup) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroup) SetConfigPaths(v *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroup {
	s.ConfigPaths = v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroup) SetGroupName(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroup) SetGroupLabel(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroup {
	s.GroupLabel = &v
	return s
}

type DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths struct {
	ConfigPath []*DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath `json:"ConfigPath,omitempty" xml:"ConfigPath,omitempty" type:"Repeated"`
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths) SetConfigPath(v []*DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPaths {
	s.ConfigPath = v
	return s
}

type DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath struct {
	PathName      *string                                                                                   `json:"PathName,omitempty" xml:"PathName,omitempty"`
	HiddenPath    *bool                                                                                     `json:"HiddenPath,omitempty" xml:"HiddenPath,omitempty"`
	ConfigOptions *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions `json:"ConfigOptions,omitempty" xml:"ConfigOptions,omitempty" type:"Struct"`
	PathLabel     *string                                                                                   `json:"PathLabel,omitempty" xml:"PathLabel,omitempty"`
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) SetPathName(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath {
	s.PathName = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) SetHiddenPath(v bool) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath {
	s.HiddenPath = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) SetConfigOptions(v *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath {
	s.ConfigOptions = v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath) SetPathLabel(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPath {
	s.PathLabel = &v
	return s
}

type DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions struct {
	ConfigOption []*DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption `json:"ConfigOption,omitempty" xml:"ConfigOption,omitempty" type:"Repeated"`
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions) SetConfigOption(v []*DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptions {
	s.ConfigOption = v
	return s
}

type DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption struct {
	RegexDesc         *string `json:"RegexDesc,omitempty" xml:"RegexDesc,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	EditorType        *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	OptionLabel       *string `json:"OptionLabel,omitempty" xml:"OptionLabel,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	MaxLength         *int32  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	RegexPattern      *string `json:"RegexPattern,omitempty" xml:"RegexPattern,omitempty"`
	ChangeSeverity    *string `json:"ChangeSeverity,omitempty" xml:"ChangeSeverity,omitempty"`
	OptionDescription *string `json:"OptionDescription,omitempty" xml:"OptionDescription,omitempty"`
	OptionName        *string `json:"OptionName,omitempty" xml:"OptionName,omitempty"`
	PathName          *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	HiddenOption      *bool   `json:"HiddenOption,omitempty" xml:"HiddenOption,omitempty"`
	ValueType         *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	MinLength         *int32  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	ValueOptions      *string `json:"ValueOptions,omitempty" xml:"ValueOptions,omitempty"`
	ReadonlyOption    *bool   `json:"ReadonlyOption,omitempty" xml:"ReadonlyOption,omitempty"`
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetRegexDesc(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.RegexDesc = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetMaxValue(v int64) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.MaxValue = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetEditorType(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.EditorType = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetMinValue(v int64) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.MinValue = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetOptionLabel(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.OptionLabel = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetDefaultValue(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.DefaultValue = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetMaxLength(v int32) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.MaxLength = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetRegexPattern(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.RegexPattern = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetChangeSeverity(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.ChangeSeverity = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetOptionDescription(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.OptionDescription = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetOptionName(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.OptionName = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetPathName(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.PathName = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetHiddenOption(v bool) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.HiddenOption = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetValueType(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.ValueType = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetMinLength(v int32) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.MinLength = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetValueOptions(v string) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.ValueOptions = &v
	return s
}

func (s *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption) SetReadonlyOption(v bool) *DescribeConfigIndexResponseBodyConfigGroupsConfigGroupConfigPathsConfigPathConfigOptionsConfigOption {
	s.ReadonlyOption = &v
	return s
}

type DescribeConfigIndexResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConfigIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigIndexResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigIndexResponse) SetHeaders(v map[string]*string) *DescribeConfigIndexResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigIndexResponse) SetBody(v *DescribeConfigIndexResponseBody) *DescribeConfigIndexResponse {
	s.Body = v
	return s
}

type DescribeConfigOptionsRequest struct {
	StackId     *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	EnvId       *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ProfileName *string `json:"ProfileName,omitempty" xml:"ProfileName,omitempty"`
}

func (s DescribeConfigOptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOptionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigOptionsRequest) SetStackId(v string) *DescribeConfigOptionsRequest {
	s.StackId = &v
	return s
}

func (s *DescribeConfigOptionsRequest) SetEnvId(v string) *DescribeConfigOptionsRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeConfigOptionsRequest) SetProfileName(v string) *DescribeConfigOptionsRequest {
	s.ProfileName = &v
	return s
}

type DescribeConfigOptionsResponseBody struct {
	StackConfigOption *DescribeConfigOptionsResponseBodyStackConfigOption `json:"StackConfigOption,omitempty" xml:"StackConfigOption,omitempty" type:"Struct"`
	Message           *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code              *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeConfigOptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigOptionsResponseBody) SetStackConfigOption(v *DescribeConfigOptionsResponseBodyStackConfigOption) *DescribeConfigOptionsResponseBody {
	s.StackConfigOption = v
	return s
}

func (s *DescribeConfigOptionsResponseBody) SetMessage(v string) *DescribeConfigOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigOptionsResponseBody) SetRequestId(v string) *DescribeConfigOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigOptionsResponseBody) SetCode(v string) *DescribeConfigOptionsResponseBody {
	s.Code = &v
	return s
}

type DescribeConfigOptionsResponseBodyStackConfigOption struct {
	StackId       *string                                                          `json:"StackId,omitempty" xml:"StackId,omitempty"`
	ConfigOptions *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions `json:"ConfigOptions,omitempty" xml:"ConfigOptions,omitempty" type:"Struct"`
	StackName     *string                                                          `json:"StackName,omitempty" xml:"StackName,omitempty"`
}

func (s DescribeConfigOptionsResponseBodyStackConfigOption) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOptionsResponseBodyStackConfigOption) GoString() string {
	return s.String()
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOption) SetStackId(v string) *DescribeConfigOptionsResponseBodyStackConfigOption {
	s.StackId = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOption) SetConfigOptions(v *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions) *DescribeConfigOptionsResponseBodyStackConfigOption {
	s.ConfigOptions = v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOption) SetStackName(v string) *DescribeConfigOptionsResponseBodyStackConfigOption {
	s.StackName = &v
	return s
}

type DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions struct {
	ConfigOption []*DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption `json:"ConfigOption,omitempty" xml:"ConfigOption,omitempty" type:"Repeated"`
}

func (s DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions) SetConfigOption(v []*DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptions {
	s.ConfigOption = v
	return s
}

type DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption struct {
	RegexDesc         *string `json:"RegexDesc,omitempty" xml:"RegexDesc,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	EditorType        *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	MaxLength         *int32  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	OptionLabel       *string `json:"OptionLabel,omitempty" xml:"OptionLabel,omitempty"`
	RegexPattern      *string `json:"RegexPattern,omitempty" xml:"RegexPattern,omitempty"`
	ChangeSeverity    *string `json:"ChangeSeverity,omitempty" xml:"ChangeSeverity,omitempty"`
	OptionDescription *string `json:"OptionDescription,omitempty" xml:"OptionDescription,omitempty"`
	OptionName        *string `json:"OptionName,omitempty" xml:"OptionName,omitempty"`
	PathName          *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	HiddenOption      *bool   `json:"HiddenOption,omitempty" xml:"HiddenOption,omitempty"`
	ValueType         *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	MinLength         *int32  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	ValueOptions      *string `json:"ValueOptions,omitempty" xml:"ValueOptions,omitempty"`
	ReadonlyOption    *bool   `json:"ReadonlyOption,omitempty" xml:"ReadonlyOption,omitempty"`
}

func (s DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) GoString() string {
	return s.String()
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetRegexDesc(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.RegexDesc = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetMaxValue(v int64) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.MaxValue = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetEditorType(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.EditorType = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetMinValue(v int64) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.MinValue = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetDefaultValue(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.DefaultValue = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetMaxLength(v int32) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.MaxLength = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetOptionLabel(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.OptionLabel = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetRegexPattern(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.RegexPattern = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetChangeSeverity(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.ChangeSeverity = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetOptionDescription(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.OptionDescription = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetOptionName(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.OptionName = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetPathName(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.PathName = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetHiddenOption(v bool) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.HiddenOption = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetValueType(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.ValueType = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetMinLength(v int32) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.MinLength = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetValueOptions(v string) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.ValueOptions = &v
	return s
}

func (s *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption) SetReadonlyOption(v bool) *DescribeConfigOptionsResponseBodyStackConfigOptionConfigOptionsConfigOption {
	s.ReadonlyOption = &v
	return s
}

type DescribeConfigOptionsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConfigOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOptionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigOptionsResponse) SetHeaders(v map[string]*string) *DescribeConfigOptionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigOptionsResponse) SetBody(v *DescribeConfigOptionsResponseBody) *DescribeConfigOptionsResponse {
	s.Body = v
	return s
}

type DescribeConfigSettingsRequest struct {
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	PathName   *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	OptionName *string `json:"OptionName,omitempty" xml:"OptionName,omitempty"`
}

func (s DescribeConfigSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigSettingsRequest) SetEnvId(v string) *DescribeConfigSettingsRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeConfigSettingsRequest) SetTemplateId(v string) *DescribeConfigSettingsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeConfigSettingsRequest) SetPathName(v string) *DescribeConfigSettingsRequest {
	s.PathName = &v
	return s
}

func (s *DescribeConfigSettingsRequest) SetOptionName(v string) *DescribeConfigSettingsRequest {
	s.OptionName = &v
	return s
}

type DescribeConfigSettingsResponseBody struct {
	ConfigSettings *DescribeConfigSettingsResponseBodyConfigSettings `json:"ConfigSettings,omitempty" xml:"ConfigSettings,omitempty" type:"Struct"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeConfigSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigSettingsResponseBody) SetConfigSettings(v *DescribeConfigSettingsResponseBodyConfigSettings) *DescribeConfigSettingsResponseBody {
	s.ConfigSettings = v
	return s
}

func (s *DescribeConfigSettingsResponseBody) SetMessage(v string) *DescribeConfigSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigSettingsResponseBody) SetRequestId(v string) *DescribeConfigSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigSettingsResponseBody) SetCode(v string) *DescribeConfigSettingsResponseBody {
	s.Code = &v
	return s
}

type DescribeConfigSettingsResponseBodyConfigSettings struct {
	ConfigSetting []*DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting `json:"ConfigSetting,omitempty" xml:"ConfigSetting,omitempty" type:"Repeated"`
}

func (s DescribeConfigSettingsResponseBodyConfigSettings) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigSettingsResponseBodyConfigSettings) GoString() string {
	return s.String()
}

func (s *DescribeConfigSettingsResponseBodyConfigSettings) SetConfigSetting(v []*DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting) *DescribeConfigSettingsResponseBodyConfigSettings {
	s.ConfigSetting = v
	return s
}

type DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting struct {
	OptionName   *string `json:"OptionName,omitempty" xml:"OptionName,omitempty"`
	PathName     *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	SettingValue *string `json:"SettingValue,omitempty" xml:"SettingValue,omitempty"`
}

func (s DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting) GoString() string {
	return s.String()
}

func (s *DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting) SetOptionName(v string) *DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting {
	s.OptionName = &v
	return s
}

func (s *DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting) SetPathName(v string) *DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting {
	s.PathName = &v
	return s
}

func (s *DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting) SetSettingValue(v string) *DescribeConfigSettingsResponseBodyConfigSettingsConfigSetting {
	s.SettingValue = &v
	return s
}

type DescribeConfigSettingsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConfigSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigSettingsResponse) SetHeaders(v map[string]*string) *DescribeConfigSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigSettingsResponse) SetBody(v *DescribeConfigSettingsResponseBody) *DescribeConfigSettingsResponse {
	s.Body = v
	return s
}

type DescribeConfigTemplatesRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateSearch *string `json:"TemplateSearch,omitempty" xml:"TemplateSearch,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeConfigTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigTemplatesRequest) SetAppId(v string) *DescribeConfigTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeConfigTemplatesRequest) SetTemplateName(v string) *DescribeConfigTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeConfigTemplatesRequest) SetTemplateSearch(v string) *DescribeConfigTemplatesRequest {
	s.TemplateSearch = &v
	return s
}

func (s *DescribeConfigTemplatesRequest) SetPageSize(v int32) *DescribeConfigTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConfigTemplatesRequest) SetPageNumber(v int32) *DescribeConfigTemplatesRequest {
	s.PageNumber = &v
	return s
}

type DescribeConfigTemplatesResponseBody struct {
	ConfigTemplates *DescribeConfigTemplatesResponseBodyConfigTemplates `json:"ConfigTemplates,omitempty" xml:"ConfigTemplates,omitempty" type:"Struct"`
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message         *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Code            *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeConfigTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigTemplatesResponseBody) SetConfigTemplates(v *DescribeConfigTemplatesResponseBodyConfigTemplates) *DescribeConfigTemplatesResponseBody {
	s.ConfigTemplates = v
	return s
}

func (s *DescribeConfigTemplatesResponseBody) SetTotalCount(v int32) *DescribeConfigTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBody) SetRequestId(v string) *DescribeConfigTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBody) SetMessage(v string) *DescribeConfigTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBody) SetPageSize(v int32) *DescribeConfigTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBody) SetPageNumber(v int32) *DescribeConfigTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBody) SetCode(v string) *DescribeConfigTemplatesResponseBody {
	s.Code = &v
	return s
}

type DescribeConfigTemplatesResponseBodyConfigTemplates struct {
	ConfigTemplate []*DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate `json:"ConfigTemplate,omitempty" xml:"ConfigTemplate,omitempty" type:"Repeated"`
}

func (s DescribeConfigTemplatesResponseBodyConfigTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigTemplatesResponseBodyConfigTemplates) GoString() string {
	return s.String()
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplates) SetConfigTemplate(v []*DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) *DescribeConfigTemplatesResponseBodyConfigTemplates {
	s.ConfigTemplate = v
	return s
}

type DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate struct {
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime          *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	StackId             *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	PkgVersionLabel     *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StackName           *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	PkgVersionId        *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) GoString() string {
	return s.String()
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetTemplateDescription(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.TemplateDescription = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetAppName(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.AppName = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetUpdateTime(v int64) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.UpdateTime = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetStackId(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.StackId = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetPkgVersionLabel(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.PkgVersionLabel = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetCreateTime(v int64) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.CreateTime = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetAppId(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.AppId = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetStackName(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.StackName = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetPkgVersionId(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.PkgVersionId = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetTemplateName(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.TemplateName = &v
	return s
}

func (s *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate) SetTemplateId(v string) *DescribeConfigTemplatesResponseBodyConfigTemplatesConfigTemplate {
	s.TemplateId = &v
	return s
}

type DescribeConfigTemplatesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConfigTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigTemplatesResponse) SetHeaders(v map[string]*string) *DescribeConfigTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigTemplatesResponse) SetBody(v *DescribeConfigTemplatesResponseBody) *DescribeConfigTemplatesResponse {
	s.Body = v
	return s
}

type DescribeEnvResourceRequest struct {
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DescribeEnvResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceRequest) SetEnvId(v string) *DescribeEnvResourceRequest {
	s.EnvId = &v
	return s
}

type DescribeEnvResourceResponseBody struct {
	Message     *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	EnvResource *DescribeEnvResourceResponseBodyEnvResource `json:"EnvResource,omitempty" xml:"EnvResource,omitempty" type:"Struct"`
}

func (s DescribeEnvResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBody) SetMessage(v string) *DescribeEnvResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvResourceResponseBody) SetRequestId(v string) *DescribeEnvResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvResourceResponseBody) SetCode(v string) *DescribeEnvResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvResourceResponseBody) SetEnvResource(v *DescribeEnvResourceResponseBodyEnvResource) *DescribeEnvResourceResponseBody {
	s.EnvResource = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResource struct {
	LaunchTemplateId      *string                                                          `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	EnvName               *string                                                          `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	VSwitches             *DescribeEnvResourceResponseBodyEnvResourceVSwitches             `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
	Vpc                   *DescribeEnvResourceResponseBodyEnvResourceVpc                   `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
	MonitorGroup          *DescribeEnvResourceResponseBodyEnvResourceMonitorGroup          `json:"MonitorGroup,omitempty" xml:"MonitorGroup,omitempty" type:"Struct"`
	LaunchConfigurationId *string                                                          `json:"LaunchConfigurationId,omitempty" xml:"LaunchConfigurationId,omitempty"`
	LoadBalancers         *DescribeEnvResourceResponseBodyEnvResourceLoadBalancers         `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Struct"`
	Instances             *DescribeEnvResourceResponseBodyEnvResourceInstances             `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	DefaultSecurityGroups *DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups `json:"DefaultSecurityGroups,omitempty" xml:"DefaultSecurityGroups,omitempty" type:"Struct"`
	ScalingGroup          *DescribeEnvResourceResponseBodyEnvResourceScalingGroup          `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
	Domains               *DescribeEnvResourceResponseBodyEnvResourceDomains               `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	RdsInstances          *DescribeEnvResourceResponseBodyEnvResourceRdsInstances          `json:"RdsInstances,omitempty" xml:"RdsInstances,omitempty" type:"Struct"`
	EnvId                 *string                                                          `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResource) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetLaunchTemplateId(v string) *DescribeEnvResourceResponseBodyEnvResource {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetEnvName(v string) *DescribeEnvResourceResponseBodyEnvResource {
	s.EnvName = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetVSwitches(v *DescribeEnvResourceResponseBodyEnvResourceVSwitches) *DescribeEnvResourceResponseBodyEnvResource {
	s.VSwitches = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetVpc(v *DescribeEnvResourceResponseBodyEnvResourceVpc) *DescribeEnvResourceResponseBodyEnvResource {
	s.Vpc = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetMonitorGroup(v *DescribeEnvResourceResponseBodyEnvResourceMonitorGroup) *DescribeEnvResourceResponseBodyEnvResource {
	s.MonitorGroup = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetLaunchConfigurationId(v string) *DescribeEnvResourceResponseBodyEnvResource {
	s.LaunchConfigurationId = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetLoadBalancers(v *DescribeEnvResourceResponseBodyEnvResourceLoadBalancers) *DescribeEnvResourceResponseBodyEnvResource {
	s.LoadBalancers = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetInstances(v *DescribeEnvResourceResponseBodyEnvResourceInstances) *DescribeEnvResourceResponseBodyEnvResource {
	s.Instances = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetDefaultSecurityGroups(v *DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups) *DescribeEnvResourceResponseBodyEnvResource {
	s.DefaultSecurityGroups = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetScalingGroup(v *DescribeEnvResourceResponseBodyEnvResourceScalingGroup) *DescribeEnvResourceResponseBodyEnvResource {
	s.ScalingGroup = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetDomains(v *DescribeEnvResourceResponseBodyEnvResourceDomains) *DescribeEnvResourceResponseBodyEnvResource {
	s.Domains = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetRdsInstances(v *DescribeEnvResourceResponseBodyEnvResourceRdsInstances) *DescribeEnvResourceResponseBodyEnvResource {
	s.RdsInstances = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResource) SetEnvId(v string) *DescribeEnvResourceResponseBodyEnvResource {
	s.EnvId = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceVSwitches struct {
	VSwitch []*DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceVSwitches) SetVSwitch(v []*DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch) *DescribeEnvResourceResponseBodyEnvResourceVSwitches {
	s.VSwitch = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceVSwitchesVSwitch {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceVpc struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceVpc) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceVpc) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceVpc {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceMonitorGroup struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceMonitorGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceMonitorGroup) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceMonitorGroup) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceMonitorGroup {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceLoadBalancers struct {
	LoadBalancer []*DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancers) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancers) SetLoadBalancer(v []*DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancers {
	s.LoadBalancer = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer struct {
	Imported    *bool                                                                         `json:"Imported,omitempty" xml:"Imported,omitempty"`
	Protocol    *string                                                                       `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	AddressType *string                                                                       `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Listeners   *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
	Port        *int32                                                                        `json:"Port,omitempty" xml:"Port,omitempty"`
	Id          *string                                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) SetImported(v bool) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer {
	s.Imported = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) SetProtocol(v string) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer {
	s.Protocol = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) SetAddressType(v string) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer {
	s.AddressType = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) SetListeners(v *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer {
	s.Listeners = v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) SetPort(v int32) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer {
	s.Port = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancer {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners struct {
	Listener []*DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners) SetListener(v []*DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListeners {
	s.Listener = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener) SetProtocol(v string) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener {
	s.Protocol = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener) SetPort(v int32) *DescribeEnvResourceResponseBodyEnvResourceLoadBalancersLoadBalancerListenersListener {
	s.Port = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceInstances struct {
	Instance []*DescribeEnvResourceResponseBodyEnvResourceInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceInstances) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceInstances) SetInstance(v []*DescribeEnvResourceResponseBodyEnvResourceInstancesInstance) *DescribeEnvResourceResponseBodyEnvResourceInstances {
	s.Instance = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceInstancesInstance struct {
	Imported *bool   `json:"Imported,omitempty" xml:"Imported,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceInstancesInstance) SetImported(v bool) *DescribeEnvResourceResponseBodyEnvResourceInstancesInstance {
	s.Imported = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceInstancesInstance) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceInstancesInstance {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups struct {
	SecurityGroup []*DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups) SetSecurityGroup(v []*DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup) *DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroups {
	s.SecurityGroup = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceDefaultSecurityGroupsSecurityGroup {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceScalingGroup struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceScalingGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceScalingGroup) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceScalingGroup) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceScalingGroup {
	s.Id = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceDomains struct {
	Domain []*DescribeEnvResourceResponseBodyEnvResourceDomainsDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceDomains) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDomains) SetDomain(v []*DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) *DescribeEnvResourceResponseBodyEnvResourceDomains {
	s.Domain = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceDomainsDomain struct {
	IsDefault  *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	HostedBy   *string `json:"HostedBy,omitempty" xml:"HostedBy,omitempty"`
	SubDomain  *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ManagedBy  *string `json:"ManagedBy,omitempty" xml:"ManagedBy,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) SetIsDefault(v bool) *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain {
	s.IsDefault = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) SetHostedBy(v string) *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain {
	s.HostedBy = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) SetSubDomain(v string) *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain {
	s.SubDomain = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) SetDomainName(v string) *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain) SetManagedBy(v string) *DescribeEnvResourceResponseBodyEnvResourceDomainsDomain {
	s.ManagedBy = &v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceRdsInstances struct {
	RdsInstance []*DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceRdsInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceRdsInstances) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceRdsInstances) SetRdsInstance(v []*DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) *DescribeEnvResourceResponseBodyEnvResourceRdsInstances {
	s.RdsInstance = v
	return s
}

type DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance struct {
	Imported     *bool   `json:"Imported,omitempty" xml:"Imported,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) SetImported(v bool) *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance {
	s.Imported = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) SetDatabaseName(v string) *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance {
	s.DatabaseName = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) SetId(v string) *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance {
	s.Id = &v
	return s
}

func (s *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance) SetAccountName(v string) *DescribeEnvResourceResponseBodyEnvResourceRdsInstancesRdsInstance {
	s.AccountName = &v
	return s
}

type DescribeEnvResourceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEnvResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnvResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnvResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvResourceResponse) SetHeaders(v map[string]*string) *DescribeEnvResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvResourceResponse) SetBody(v *DescribeEnvResourceResponseBody) *DescribeEnvResourceResponse {
	s.Body = v
	return s
}

type DescribeEventsRequest struct {
	EnvId              *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	StartTime          *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime            *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ChangeId           *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	LastChangeEvents   *bool   `json:"LastChangeEvents,omitempty" xml:"LastChangeEvents,omitempty"`
	ReverseByTimestamp *bool   `json:"ReverseByTimestamp,omitempty" xml:"ReverseByTimestamp,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetEnvId(v string) *DescribeEventsRequest {
	s.EnvId = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v int64) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetEndTime(v int64) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetPageNumber(v int32) *DescribeEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsRequest) SetChangeId(v string) *DescribeEventsRequest {
	s.ChangeId = &v
	return s
}

func (s *DescribeEventsRequest) SetLastChangeEvents(v bool) *DescribeEventsRequest {
	s.LastChangeEvents = &v
	return s
}

func (s *DescribeEventsRequest) SetReverseByTimestamp(v bool) *DescribeEventsRequest {
	s.ReverseByTimestamp = &v
	return s
}

type DescribeEventsResponseBody struct {
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Events     *DescribeEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Struct"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetMessage(v string) *DescribeEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int32) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageNumber(v int32) *DescribeEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsResponseBody) SetEvents(v *DescribeEventsResponseBodyEvents) *DescribeEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsResponseBody) SetCode(v string) *DescribeEventsResponseBody {
	s.Code = &v
	return s
}

type DescribeEventsResponseBodyEvents struct {
	Event []*DescribeEventsResponseBodyEventsEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
}

func (s DescribeEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEvents) SetEvent(v []*DescribeEventsResponseBodyEventsEvent) *DescribeEventsResponseBodyEvents {
	s.Event = v
	return s
}

type DescribeEventsResponseBodyEventsEvent struct {
	PrimaryUserName *string                                           `json:"PrimaryUserName,omitempty" xml:"PrimaryUserName,omitempty"`
	ObjectType      *string                                           `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	EnventName      *string                                           `json:"EnventName,omitempty" xml:"EnventName,omitempty"`
	EventTimestamp  *int64                                            `json:"EventTimestamp,omitempty" xml:"EventTimestamp,omitempty"`
	SecondUserName  *string                                           `json:"SecondUserName,omitempty" xml:"SecondUserName,omitempty"`
	MsgCode         *string                                           `json:"MsgCode,omitempty" xml:"MsgCode,omitempty"`
	ObjectName      *string                                           `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	EventMessage    *string                                           `json:"EventMessage,omitempty" xml:"EventMessage,omitempty"`
	EventId         *string                                           `json:"EventId,omitempty" xml:"EventId,omitempty"`
	ObjectAttrs     *DescribeEventsResponseBodyEventsEventObjectAttrs `json:"ObjectAttrs,omitempty" xml:"ObjectAttrs,omitempty" type:"Struct"`
	AppId           *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EventLevel      *string                                           `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	ObjectId        *string                                           `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	MsgParams       *DescribeEventsResponseBodyEventsEventMsgParams   `json:"MsgParams,omitempty" xml:"MsgParams,omitempty" type:"Struct"`
	PrimaryUserId   *string                                           `json:"PrimaryUserId,omitempty" xml:"PrimaryUserId,omitempty"`
	EnvId           *string                                           `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	TraceId         *string                                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	SecondUserId    *string                                           `json:"SecondUserId,omitempty" xml:"SecondUserId,omitempty"`
}

func (s DescribeEventsResponseBodyEventsEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEventsEvent) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventsEvent) SetPrimaryUserName(v string) *DescribeEventsResponseBodyEventsEvent {
	s.PrimaryUserName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetObjectType(v string) *DescribeEventsResponseBodyEventsEvent {
	s.ObjectType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetEnventName(v string) *DescribeEventsResponseBodyEventsEvent {
	s.EnventName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetEventTimestamp(v int64) *DescribeEventsResponseBodyEventsEvent {
	s.EventTimestamp = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetSecondUserName(v string) *DescribeEventsResponseBodyEventsEvent {
	s.SecondUserName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetMsgCode(v string) *DescribeEventsResponseBodyEventsEvent {
	s.MsgCode = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetObjectName(v string) *DescribeEventsResponseBodyEventsEvent {
	s.ObjectName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetEventMessage(v string) *DescribeEventsResponseBodyEventsEvent {
	s.EventMessage = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetEventId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetObjectAttrs(v *DescribeEventsResponseBodyEventsEventObjectAttrs) *DescribeEventsResponseBodyEventsEvent {
	s.ObjectAttrs = v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetAppId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.AppId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetEventLevel(v string) *DescribeEventsResponseBodyEventsEvent {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetObjectId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.ObjectId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetMsgParams(v *DescribeEventsResponseBodyEventsEventMsgParams) *DescribeEventsResponseBodyEventsEvent {
	s.MsgParams = v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetPrimaryUserId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.PrimaryUserId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetEnvId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.EnvId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetTraceId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.TraceId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEvent) SetSecondUserId(v string) *DescribeEventsResponseBodyEventsEvent {
	s.SecondUserId = &v
	return s
}

type DescribeEventsResponseBodyEventsEventObjectAttrs struct {
	ObjectAttr []*DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr `json:"ObjectAttr,omitempty" xml:"ObjectAttr,omitempty" type:"Repeated"`
}

func (s DescribeEventsResponseBodyEventsEventObjectAttrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEventsEventObjectAttrs) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventsEventObjectAttrs) SetObjectAttr(v []*DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr) *DescribeEventsResponseBodyEventsEventObjectAttrs {
	s.ObjectAttr = v
	return s
}

type DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr struct {
	AttributeName  *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr) SetAttributeName(v string) *DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr {
	s.AttributeName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr) SetAttributeValue(v string) *DescribeEventsResponseBodyEventsEventObjectAttrsObjectAttr {
	s.AttributeValue = &v
	return s
}

type DescribeEventsResponseBodyEventsEventMsgParams struct {
	MsgParam []*string `json:"MsgParam,omitempty" xml:"MsgParam,omitempty" type:"Repeated"`
}

func (s DescribeEventsResponseBodyEventsEventMsgParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEventsEventMsgParams) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventsEventMsgParams) SetMsgParam(v []*string) *DescribeEventsResponseBodyEventsEventMsgParams {
	s.MsgParam = v
	return s
}

type DescribeEventsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type DescribeGatherLogResultRequest struct {
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s DescribeGatherLogResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultRequest) SetChangeId(v string) *DescribeGatherLogResultRequest {
	s.ChangeId = &v
	return s
}

type DescribeGatherLogResultResponseBody struct {
	GatherLogResult *DescribeGatherLogResultResponseBodyGatherLogResult `json:"GatherLogResult,omitempty" xml:"GatherLogResult,omitempty" type:"Struct"`
	Message         *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code            *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeGatherLogResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultResponseBody) SetGatherLogResult(v *DescribeGatherLogResultResponseBodyGatherLogResult) *DescribeGatherLogResultResponseBody {
	s.GatherLogResult = v
	return s
}

func (s *DescribeGatherLogResultResponseBody) SetMessage(v string) *DescribeGatherLogResultResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatherLogResultResponseBody) SetRequestId(v string) *DescribeGatherLogResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatherLogResultResponseBody) SetCode(v string) *DescribeGatherLogResultResponseBody {
	s.Code = &v
	return s
}

type DescribeGatherLogResultResponseBodyGatherLogResult struct {
	LogPath         *string                                                            `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	Change          *DescribeGatherLogResultResponseBodyGatherLogResultChange          `json:"Change,omitempty" xml:"Change,omitempty" type:"Struct"`
	InstanceResults *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults `json:"InstanceResults,omitempty" xml:"InstanceResults,omitempty" type:"Struct"`
}

func (s DescribeGatherLogResultResponseBodyGatherLogResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultResponseBodyGatherLogResult) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResult) SetLogPath(v string) *DescribeGatherLogResultResponseBodyGatherLogResult {
	s.LogPath = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResult) SetChange(v *DescribeGatherLogResultResponseBodyGatherLogResultChange) *DescribeGatherLogResultResponseBodyGatherLogResult {
	s.Change = v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResult) SetInstanceResults(v *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults) *DescribeGatherLogResultResponseBodyGatherLogResult {
	s.InstanceResults = v
	return s
}

type DescribeGatherLogResultResponseBodyGatherLogResultChange struct {
	ChangePaused      *bool   `json:"ChangePaused,omitempty" xml:"ChangePaused,omitempty"`
	ChangeDescription *string `json:"ChangeDescription,omitempty" xml:"ChangeDescription,omitempty"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChangeTimedout    *bool   `json:"ChangeTimedout,omitempty" xml:"ChangeTimedout,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChangeMessage     *string `json:"ChangeMessage,omitempty" xml:"ChangeMessage,omitempty"`
	ActionName        *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ChangeFinished    *bool   `json:"ChangeFinished,omitempty" xml:"ChangeFinished,omitempty"`
	CreateUsername    *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	ChangeSucceeded   *bool   `json:"ChangeSucceeded,omitempty" xml:"ChangeSucceeded,omitempty"`
	ChangeId          *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	ChangeAborted     *bool   `json:"ChangeAborted,omitempty" xml:"ChangeAborted,omitempty"`
	EnvId             *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ChangeName        *string `json:"ChangeName,omitempty" xml:"ChangeName,omitempty"`
}

func (s DescribeGatherLogResultResponseBodyGatherLogResultChange) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultResponseBodyGatherLogResultChange) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangePaused(v bool) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangePaused = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeDescription(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeDescription = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetFinishTime(v int64) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.FinishTime = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetUpdateTime(v int64) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeTimedout(v bool) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeTimedout = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetCreateTime(v int64) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeMessage(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeMessage = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetActionName(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ActionName = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeFinished(v bool) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeFinished = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetCreateUsername(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.CreateUsername = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeSucceeded(v bool) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeSucceeded = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeId(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeId = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeAborted(v bool) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeAborted = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetEnvId(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.EnvId = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultChange) SetChangeName(v string) *DescribeGatherLogResultResponseBodyGatherLogResultChange {
	s.ChangeName = &v
	return s
}

type DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults struct {
	InstanceResult []*DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult `json:"InstanceResult,omitempty" xml:"InstanceResult,omitempty" type:"Repeated"`
}

func (s DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults) SetInstanceResult(v []*DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResults {
	s.InstanceResult = v
	return s
}

type DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult struct {
	TaskMessage   *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	StorageKey    *string `json:"StorageKey,omitempty" xml:"StorageKey,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskSucceeded *bool   `json:"TaskSucceeded,omitempty" xml:"TaskSucceeded,omitempty"`
}

func (s DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) SetTaskMessage(v string) *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult {
	s.TaskMessage = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) SetStorageKey(v string) *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult {
	s.StorageKey = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) SetInstanceId(v string) *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult) SetTaskSucceeded(v bool) *DescribeGatherLogResultResponseBodyGatherLogResultInstanceResultsInstanceResult {
	s.TaskSucceeded = &v
	return s
}

type DescribeGatherLogResultResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGatherLogResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatherLogResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherLogResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatherLogResultResponse) SetHeaders(v map[string]*string) *DescribeGatherLogResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatherLogResultResponse) SetBody(v *DescribeGatherLogResultResponseBody) *DescribeGatherLogResultResponse {
	s.Body = v
	return s
}

type DescribeGatherStatsResultRequest struct {
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s DescribeGatherStatsResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultRequest) SetChangeId(v string) *DescribeGatherStatsResultRequest {
	s.ChangeId = &v
	return s
}

type DescribeGatherStatsResultResponseBody struct {
	Message           *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GatherStatsResult *DescribeGatherStatsResultResponseBodyGatherStatsResult `json:"GatherStatsResult,omitempty" xml:"GatherStatsResult,omitempty" type:"Struct"`
	Code              *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeGatherStatsResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultResponseBody) SetMessage(v string) *DescribeGatherStatsResultResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBody) SetRequestId(v string) *DescribeGatherStatsResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBody) SetGatherStatsResult(v *DescribeGatherStatsResultResponseBodyGatherStatsResult) *DescribeGatherStatsResultResponseBody {
	s.GatherStatsResult = v
	return s
}

func (s *DescribeGatherStatsResultResponseBody) SetCode(v string) *DescribeGatherStatsResultResponseBody {
	s.Code = &v
	return s
}

type DescribeGatherStatsResultResponseBodyGatherStatsResult struct {
	Change          *DescribeGatherStatsResultResponseBodyGatherStatsResultChange          `json:"Change,omitempty" xml:"Change,omitempty" type:"Struct"`
	InstanceResults *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults `json:"InstanceResults,omitempty" xml:"InstanceResults,omitempty" type:"Struct"`
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResult) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResult) SetChange(v *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) *DescribeGatherStatsResultResponseBodyGatherStatsResult {
	s.Change = v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResult) SetInstanceResults(v *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults) *DescribeGatherStatsResultResponseBodyGatherStatsResult {
	s.InstanceResults = v
	return s
}

type DescribeGatherStatsResultResponseBodyGatherStatsResultChange struct {
	ChangePaused      *bool   `json:"ChangePaused,omitempty" xml:"ChangePaused,omitempty"`
	ChangeDescription *string `json:"ChangeDescription,omitempty" xml:"ChangeDescription,omitempty"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChangeTimedout    *bool   `json:"ChangeTimedout,omitempty" xml:"ChangeTimedout,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChangeMessage     *string `json:"ChangeMessage,omitempty" xml:"ChangeMessage,omitempty"`
	ActionName        *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ChangeFinished    *bool   `json:"ChangeFinished,omitempty" xml:"ChangeFinished,omitempty"`
	CreateUsername    *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	ChangeSucceeded   *bool   `json:"ChangeSucceeded,omitempty" xml:"ChangeSucceeded,omitempty"`
	ChangeId          *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	ChangeAborted     *bool   `json:"ChangeAborted,omitempty" xml:"ChangeAborted,omitempty"`
	EnvId             *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ChangeName        *string `json:"ChangeName,omitempty" xml:"ChangeName,omitempty"`
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResultChange) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResultChange) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangePaused(v bool) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangePaused = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeDescription(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeDescription = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetFinishTime(v int64) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.FinishTime = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetUpdateTime(v int64) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeTimedout(v bool) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeTimedout = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetCreateTime(v int64) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeMessage(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeMessage = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetActionName(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ActionName = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeFinished(v bool) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeFinished = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetCreateUsername(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.CreateUsername = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeSucceeded(v bool) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeSucceeded = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeId(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeId = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeAborted(v bool) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeAborted = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetEnvId(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.EnvId = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultChange) SetChangeName(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultChange {
	s.ChangeName = &v
	return s
}

type DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults struct {
	InstanceResult []*DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult `json:"InstanceResult,omitempty" xml:"InstanceResult,omitempty" type:"Repeated"`
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults) SetInstanceResult(v []*DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResults {
	s.InstanceResult = v
	return s
}

type DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult struct {
	TaskMessage   *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	StorageKey    *string `json:"StorageKey,omitempty" xml:"StorageKey,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskSucceeded *bool   `json:"TaskSucceeded,omitempty" xml:"TaskSucceeded,omitempty"`
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) SetTaskMessage(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult {
	s.TaskMessage = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) SetStorageKey(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult {
	s.StorageKey = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) SetInstanceId(v string) *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult) SetTaskSucceeded(v bool) *DescribeGatherStatsResultResponseBodyGatherStatsResultInstanceResultsInstanceResult {
	s.TaskSucceeded = &v
	return s
}

type DescribeGatherStatsResultResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGatherStatsResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatherStatsResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatherStatsResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatherStatsResultResponse) SetHeaders(v map[string]*string) *DescribeGatherStatsResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatherStatsResultResponse) SetBody(v *DescribeGatherStatsResultResponseBody) *DescribeGatherStatsResultResponse {
	s.Body = v
	return s
}

type DescribeInstanceHealthRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHealthRequest) SetInstanceId(v string) *DescribeInstanceHealthRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceHealthResponseBody struct {
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceHealth *DescribeInstanceHealthResponseBodyInstanceHealth `json:"InstanceHealth,omitempty" xml:"InstanceHealth,omitempty" type:"Struct"`
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeInstanceHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHealthResponseBody) SetMessage(v string) *DescribeInstanceHealthResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceHealthResponseBody) SetRequestId(v string) *DescribeInstanceHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceHealthResponseBody) SetInstanceHealth(v *DescribeInstanceHealthResponseBodyInstanceHealth) *DescribeInstanceHealthResponseBody {
	s.InstanceHealth = v
	return s
}

func (s *DescribeInstanceHealthResponseBody) SetCode(v string) *DescribeInstanceHealthResponseBody {
	s.Code = &v
	return s
}

type DescribeInstanceHealthResponseBodyInstanceHealth struct {
	AppStatus        *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DisconnectedTime *int64  `json:"DisconnectedTime,omitempty" xml:"DisconnectedTime,omitempty"`
	AgentStatus      *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
}

func (s DescribeInstanceHealthResponseBodyInstanceHealth) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceHealthResponseBodyInstanceHealth) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHealthResponseBodyInstanceHealth) SetAppStatus(v string) *DescribeInstanceHealthResponseBodyInstanceHealth {
	s.AppStatus = &v
	return s
}

func (s *DescribeInstanceHealthResponseBodyInstanceHealth) SetInstanceId(v string) *DescribeInstanceHealthResponseBodyInstanceHealth {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceHealthResponseBodyInstanceHealth) SetDisconnectedTime(v int64) *DescribeInstanceHealthResponseBodyInstanceHealth {
	s.DisconnectedTime = &v
	return s
}

func (s *DescribeInstanceHealthResponseBodyInstanceHealth) SetAgentStatus(v string) *DescribeInstanceHealthResponseBodyInstanceHealth {
	s.AgentStatus = &v
	return s
}

type DescribeInstanceHealthResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceHealthResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHealthResponse) SetHeaders(v map[string]*string) *DescribeInstanceHealthResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceHealthResponse) SetBody(v *DescribeInstanceHealthResponseBody) *DescribeInstanceHealthResponse {
	s.Body = v
	return s
}

type DescribePkgVersionsRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PkgVersionLabel  *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	PkgVersionSearch *string `json:"PkgVersionSearch,omitempty" xml:"PkgVersionSearch,omitempty"`
}

func (s DescribePkgVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePkgVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePkgVersionsRequest) SetAppId(v string) *DescribePkgVersionsRequest {
	s.AppId = &v
	return s
}

func (s *DescribePkgVersionsRequest) SetPageSize(v int32) *DescribePkgVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePkgVersionsRequest) SetPageNumber(v int32) *DescribePkgVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePkgVersionsRequest) SetPkgVersionLabel(v string) *DescribePkgVersionsRequest {
	s.PkgVersionLabel = &v
	return s
}

func (s *DescribePkgVersionsRequest) SetPkgVersionSearch(v string) *DescribePkgVersionsRequest {
	s.PkgVersionSearch = &v
	return s
}

type DescribePkgVersionsResponseBody struct {
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message     *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PkgVersions *DescribePkgVersionsResponseBodyPkgVersions `json:"PkgVersions,omitempty" xml:"PkgVersions,omitempty" type:"Struct"`
	Code        *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribePkgVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePkgVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePkgVersionsResponseBody) SetTotalCount(v int32) *DescribePkgVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePkgVersionsResponseBody) SetRequestId(v string) *DescribePkgVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePkgVersionsResponseBody) SetMessage(v string) *DescribePkgVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePkgVersionsResponseBody) SetPageSize(v int32) *DescribePkgVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePkgVersionsResponseBody) SetPageNumber(v int32) *DescribePkgVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePkgVersionsResponseBody) SetPkgVersions(v *DescribePkgVersionsResponseBodyPkgVersions) *DescribePkgVersionsResponseBody {
	s.PkgVersions = v
	return s
}

func (s *DescribePkgVersionsResponseBody) SetCode(v string) *DescribePkgVersionsResponseBody {
	s.Code = &v
	return s
}

type DescribePkgVersionsResponseBodyPkgVersions struct {
	PkgVersion []*DescribePkgVersionsResponseBodyPkgVersionsPkgVersion `json:"PkgVersion,omitempty" xml:"PkgVersion,omitempty" type:"Repeated"`
}

func (s DescribePkgVersionsResponseBodyPkgVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribePkgVersionsResponseBodyPkgVersions) GoString() string {
	return s.String()
}

func (s *DescribePkgVersionsResponseBodyPkgVersions) SetPkgVersion(v []*DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) *DescribePkgVersionsResponseBodyPkgVersions {
	s.PkgVersion = v
	return s
}

type DescribePkgVersionsResponseBodyPkgVersionsPkgVersion struct {
	CreateUsername        *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	AppName               *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	PkgVersionLabel       *string `json:"PkgVersionLabel,omitempty" xml:"PkgVersionLabel,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PackageSource         *string `json:"PackageSource,omitempty" xml:"PackageSource,omitempty"`
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PackageETag           *string `json:"PackageETag,omitempty" xml:"PackageETag,omitempty"`
	PkgVersionId          *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	PkgVersionDescription *string `json:"PkgVersionDescription,omitempty" xml:"PkgVersionDescription,omitempty"`
}

func (s DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) GoString() string {
	return s.String()
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetCreateUsername(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.CreateUsername = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetAppName(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.AppName = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetUpdateTime(v int64) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.UpdateTime = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetPkgVersionLabel(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.PkgVersionLabel = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetCreateTime(v int64) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.CreateTime = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetPackageSource(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.PackageSource = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetAppId(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.AppId = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetPackageETag(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.PackageETag = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetPkgVersionId(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.PkgVersionId = &v
	return s
}

func (s *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion) SetPkgVersionDescription(v string) *DescribePkgVersionsResponseBodyPkgVersionsPkgVersion {
	s.PkgVersionDescription = &v
	return s
}

type DescribePkgVersionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePkgVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePkgVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePkgVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePkgVersionsResponse) SetHeaders(v map[string]*string) *DescribePkgVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePkgVersionsResponse) SetBody(v *DescribePkgVersionsResponseBody) *DescribePkgVersionsResponse {
	s.Body = v
	return s
}

type DescribePublicConfigTemplatesRequest struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribePublicConfigTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePublicConfigTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribePublicConfigTemplatesRequest) SetCategoryName(v string) *DescribePublicConfigTemplatesRequest {
	s.CategoryName = &v
	return s
}

func (s *DescribePublicConfigTemplatesRequest) SetPageSize(v int32) *DescribePublicConfigTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePublicConfigTemplatesRequest) SetPageNumber(v int32) *DescribePublicConfigTemplatesRequest {
	s.PageNumber = &v
	return s
}

type DescribePublicConfigTemplatesResponseBody struct {
	TotalCount            *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message               *string                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize              *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PublicConfigTemplates *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates `json:"PublicConfigTemplates,omitempty" xml:"PublicConfigTemplates,omitempty" type:"Struct"`
	Code                  *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribePublicConfigTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePublicConfigTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePublicConfigTemplatesResponseBody) SetTotalCount(v int32) *DescribePublicConfigTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBody) SetRequestId(v string) *DescribePublicConfigTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBody) SetMessage(v string) *DescribePublicConfigTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBody) SetPageSize(v int32) *DescribePublicConfigTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBody) SetPageNumber(v int32) *DescribePublicConfigTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBody) SetPublicConfigTemplates(v *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates) *DescribePublicConfigTemplatesResponseBody {
	s.PublicConfigTemplates = v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBody) SetCode(v string) *DescribePublicConfigTemplatesResponseBody {
	s.Code = &v
	return s
}

type DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates struct {
	PublicConfigTemplate []*DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate `json:"PublicConfigTemplate,omitempty" xml:"PublicConfigTemplate,omitempty" type:"Repeated"`
}

func (s DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates) GoString() string {
	return s.String()
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates) SetPublicConfigTemplate(v []*DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplates {
	s.PublicConfigTemplate = v
	return s
}

type DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate struct {
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	UpdateTime          *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	StackId             *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateLogo        *string `json:"TemplateLogo,omitempty" xml:"TemplateLogo,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PackageSource       *string `json:"PackageSource,omitempty" xml:"PackageSource,omitempty"`
	StackName           *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	CategoryName        *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) GoString() string {
	return s.String()
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetTemplateDescription(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.TemplateDescription = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetUpdateTime(v int64) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.UpdateTime = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetStackId(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.StackId = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetTemplateLogo(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.TemplateLogo = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetCreateTime(v int64) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.CreateTime = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetPackageSource(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.PackageSource = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetStackName(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.StackName = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetTemplateName(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.TemplateName = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetCategoryName(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.CategoryName = &v
	return s
}

func (s *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate) SetTemplateId(v string) *DescribePublicConfigTemplatesResponseBodyPublicConfigTemplatesPublicConfigTemplate {
	s.TemplateId = &v
	return s
}

type DescribePublicConfigTemplatesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePublicConfigTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePublicConfigTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePublicConfigTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribePublicConfigTemplatesResponse) SetHeaders(v map[string]*string) *DescribePublicConfigTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribePublicConfigTemplatesResponse) SetBody(v *DescribePublicConfigTemplatesResponseBody) *DescribePublicConfigTemplatesResponse {
	s.Body = v
	return s
}

type DescribeStacksRequest struct {
	RecommendedOnly *bool   `json:"RecommendedOnly,omitempty" xml:"RecommendedOnly,omitempty"`
	CategoryName    *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeStacksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStacksRequest) GoString() string {
	return s.String()
}

func (s *DescribeStacksRequest) SetRecommendedOnly(v bool) *DescribeStacksRequest {
	s.RecommendedOnly = &v
	return s
}

func (s *DescribeStacksRequest) SetCategoryName(v string) *DescribeStacksRequest {
	s.CategoryName = &v
	return s
}

func (s *DescribeStacksRequest) SetPageSize(v int32) *DescribeStacksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStacksRequest) SetPageNumber(v int32) *DescribeStacksRequest {
	s.PageNumber = &v
	return s
}

type DescribeStacksResponseBody struct {
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Stacks     *DescribeStacksResponseBodyStacks `json:"Stacks,omitempty" xml:"Stacks,omitempty" type:"Struct"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeStacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStacksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStacksResponseBody) SetTotalCount(v int32) *DescribeStacksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStacksResponseBody) SetRequestId(v string) *DescribeStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStacksResponseBody) SetMessage(v string) *DescribeStacksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStacksResponseBody) SetPageSize(v int32) *DescribeStacksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStacksResponseBody) SetPageNumber(v int32) *DescribeStacksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStacksResponseBody) SetStacks(v *DescribeStacksResponseBodyStacks) *DescribeStacksResponseBody {
	s.Stacks = v
	return s
}

func (s *DescribeStacksResponseBody) SetCode(v string) *DescribeStacksResponseBody {
	s.Code = &v
	return s
}

type DescribeStacksResponseBodyStacks struct {
	Stack []*DescribeStacksResponseBodyStacksStack `json:"Stack,omitempty" xml:"Stack,omitempty" type:"Repeated"`
}

func (s DescribeStacksResponseBodyStacks) String() string {
	return tea.Prettify(s)
}

func (s DescribeStacksResponseBodyStacks) GoString() string {
	return s.String()
}

func (s *DescribeStacksResponseBodyStacks) SetStack(v []*DescribeStacksResponseBodyStacksStack) *DescribeStacksResponseBodyStacks {
	s.Stack = v
	return s
}

type DescribeStacksResponseBodyStacksStack struct {
	UpdateTime       *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	StackId          *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecommendedStack *bool   `json:"RecommendedStack,omitempty" xml:"RecommendedStack,omitempty"`
	StackName        *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	CategoryName     *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	VersionCode      *int32  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	LatestStack      *bool   `json:"LatestStack,omitempty" xml:"LatestStack,omitempty"`
}

func (s DescribeStacksResponseBodyStacksStack) String() string {
	return tea.Prettify(s)
}

func (s DescribeStacksResponseBodyStacksStack) GoString() string {
	return s.String()
}

func (s *DescribeStacksResponseBodyStacksStack) SetUpdateTime(v int64) *DescribeStacksResponseBodyStacksStack {
	s.UpdateTime = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetStackId(v string) *DescribeStacksResponseBodyStacksStack {
	s.StackId = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetCreateTime(v int64) *DescribeStacksResponseBodyStacksStack {
	s.CreateTime = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetRecommendedStack(v bool) *DescribeStacksResponseBodyStacksStack {
	s.RecommendedStack = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetStackName(v string) *DescribeStacksResponseBodyStacksStack {
	s.StackName = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetCategoryName(v string) *DescribeStacksResponseBodyStacksStack {
	s.CategoryName = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetVersionCode(v int32) *DescribeStacksResponseBodyStacksStack {
	s.VersionCode = &v
	return s
}

func (s *DescribeStacksResponseBodyStacksStack) SetLatestStack(v bool) *DescribeStacksResponseBodyStacksStack {
	s.LatestStack = &v
	return s
}

type DescribeStacksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStacksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStacksResponse) GoString() string {
	return s.String()
}

func (s *DescribeStacksResponse) SetHeaders(v map[string]*string) *DescribeStacksResponse {
	s.Headers = v
	return s
}

func (s *DescribeStacksResponse) SetBody(v *DescribeStacksResponseBody) *DescribeStacksResponse {
	s.Body = v
	return s
}

type DescribeStorageRequest struct {
	UsingSharedStorage *bool `json:"UsingSharedStorage,omitempty" xml:"UsingSharedStorage,omitempty"`
}

func (s DescribeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageRequest) SetUsingSharedStorage(v bool) *DescribeStorageRequest {
	s.UsingSharedStorage = &v
	return s
}

type DescribeStorageResponseBody struct {
	Storage   *DescribeStorageResponseBodyStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponseBody) SetStorage(v *DescribeStorageResponseBodyStorage) *DescribeStorageResponseBody {
	s.Storage = v
	return s
}

func (s *DescribeStorageResponseBody) SetMessage(v string) *DescribeStorageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStorageResponseBody) SetRequestId(v string) *DescribeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageResponseBody) SetCode(v string) *DescribeStorageResponseBody {
	s.Code = &v
	return s
}

type DescribeStorageResponseBodyStorage struct {
	UpdateTime   *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	KeyPrefix    *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	PkgKeyPrefix *string `json:"PkgKeyPrefix,omitempty" xml:"PkgKeyPrefix,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	BucketName   *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s DescribeStorageResponseBodyStorage) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponseBodyStorage) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponseBodyStorage) SetUpdateTime(v int64) *DescribeStorageResponseBodyStorage {
	s.UpdateTime = &v
	return s
}

func (s *DescribeStorageResponseBodyStorage) SetKeyPrefix(v string) *DescribeStorageResponseBodyStorage {
	s.KeyPrefix = &v
	return s
}

func (s *DescribeStorageResponseBodyStorage) SetPkgKeyPrefix(v string) *DescribeStorageResponseBodyStorage {
	s.PkgKeyPrefix = &v
	return s
}

func (s *DescribeStorageResponseBodyStorage) SetCreateTime(v int64) *DescribeStorageResponseBodyStorage {
	s.CreateTime = &v
	return s
}

func (s *DescribeStorageResponseBodyStorage) SetBucketName(v string) *DescribeStorageResponseBodyStorage {
	s.BucketName = &v
	return s
}

type DescribeStorageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponse) SetHeaders(v map[string]*string) *DescribeStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageResponse) SetBody(v *DescribeStorageResponseBody) *DescribeStorageResponse {
	s.Body = v
	return s
}

type GatherAppEnvLogRequest struct {
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	TargetInstances *string `json:"TargetInstances,omitempty" xml:"TargetInstances,omitempty"`
	LogPath         *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
}

func (s GatherAppEnvLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvLogRequest) GoString() string {
	return s.String()
}

func (s *GatherAppEnvLogRequest) SetEnvId(v string) *GatherAppEnvLogRequest {
	s.EnvId = &v
	return s
}

func (s *GatherAppEnvLogRequest) SetTargetInstances(v string) *GatherAppEnvLogRequest {
	s.TargetInstances = &v
	return s
}

func (s *GatherAppEnvLogRequest) SetLogPath(v string) *GatherAppEnvLogRequest {
	s.LogPath = &v
	return s
}

type GatherAppEnvLogResponseBody struct {
	EnvChange *GatherAppEnvLogResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GatherAppEnvLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvLogResponseBody) GoString() string {
	return s.String()
}

func (s *GatherAppEnvLogResponseBody) SetEnvChange(v *GatherAppEnvLogResponseBodyEnvChange) *GatherAppEnvLogResponseBody {
	s.EnvChange = v
	return s
}

func (s *GatherAppEnvLogResponseBody) SetMessage(v string) *GatherAppEnvLogResponseBody {
	s.Message = &v
	return s
}

func (s *GatherAppEnvLogResponseBody) SetRequestId(v string) *GatherAppEnvLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GatherAppEnvLogResponseBody) SetCode(v string) *GatherAppEnvLogResponseBody {
	s.Code = &v
	return s
}

type GatherAppEnvLogResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s GatherAppEnvLogResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvLogResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *GatherAppEnvLogResponseBodyEnvChange) SetStartTime(v string) *GatherAppEnvLogResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *GatherAppEnvLogResponseBodyEnvChange) SetChangeId(v string) *GatherAppEnvLogResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *GatherAppEnvLogResponseBodyEnvChange) SetEnvId(v string) *GatherAppEnvLogResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type GatherAppEnvLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GatherAppEnvLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GatherAppEnvLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvLogResponse) GoString() string {
	return s.String()
}

func (s *GatherAppEnvLogResponse) SetHeaders(v map[string]*string) *GatherAppEnvLogResponse {
	s.Headers = v
	return s
}

func (s *GatherAppEnvLogResponse) SetBody(v *GatherAppEnvLogResponseBody) *GatherAppEnvLogResponse {
	s.Body = v
	return s
}

type GatherAppEnvStatsRequest struct {
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	TargetInstances *string `json:"TargetInstances,omitempty" xml:"TargetInstances,omitempty"`
}

func (s GatherAppEnvStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvStatsRequest) GoString() string {
	return s.String()
}

func (s *GatherAppEnvStatsRequest) SetEnvId(v string) *GatherAppEnvStatsRequest {
	s.EnvId = &v
	return s
}

func (s *GatherAppEnvStatsRequest) SetTargetInstances(v string) *GatherAppEnvStatsRequest {
	s.TargetInstances = &v
	return s
}

type GatherAppEnvStatsResponseBody struct {
	EnvChange *GatherAppEnvStatsResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GatherAppEnvStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GatherAppEnvStatsResponseBody) SetEnvChange(v *GatherAppEnvStatsResponseBodyEnvChange) *GatherAppEnvStatsResponseBody {
	s.EnvChange = v
	return s
}

func (s *GatherAppEnvStatsResponseBody) SetMessage(v string) *GatherAppEnvStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GatherAppEnvStatsResponseBody) SetRequestId(v string) *GatherAppEnvStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GatherAppEnvStatsResponseBody) SetCode(v string) *GatherAppEnvStatsResponseBody {
	s.Code = &v
	return s
}

type GatherAppEnvStatsResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s GatherAppEnvStatsResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvStatsResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *GatherAppEnvStatsResponseBodyEnvChange) SetStartTime(v string) *GatherAppEnvStatsResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *GatherAppEnvStatsResponseBodyEnvChange) SetChangeId(v string) *GatherAppEnvStatsResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *GatherAppEnvStatsResponseBodyEnvChange) SetEnvId(v string) *GatherAppEnvStatsResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type GatherAppEnvStatsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GatherAppEnvStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GatherAppEnvStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GatherAppEnvStatsResponse) GoString() string {
	return s.String()
}

func (s *GatherAppEnvStatsResponse) SetHeaders(v map[string]*string) *GatherAppEnvStatsResponse {
	s.Headers = v
	return s
}

func (s *GatherAppEnvStatsResponse) SetBody(v *GatherAppEnvStatsResponseBody) *GatherAppEnvStatsResponse {
	s.Body = v
	return s
}

type PauseChangeRequest struct {
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s PauseChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseChangeRequest) GoString() string {
	return s.String()
}

func (s *PauseChangeRequest) SetChangeId(v string) *PauseChangeRequest {
	s.ChangeId = &v
	return s
}

type PauseChangeResponseBody struct {
	EnvChange *PauseChangeResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s PauseChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseChangeResponseBody) GoString() string {
	return s.String()
}

func (s *PauseChangeResponseBody) SetEnvChange(v *PauseChangeResponseBodyEnvChange) *PauseChangeResponseBody {
	s.EnvChange = v
	return s
}

func (s *PauseChangeResponseBody) SetMessage(v string) *PauseChangeResponseBody {
	s.Message = &v
	return s
}

func (s *PauseChangeResponseBody) SetRequestId(v string) *PauseChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseChangeResponseBody) SetCode(v string) *PauseChangeResponseBody {
	s.Code = &v
	return s
}

type PauseChangeResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s PauseChangeResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s PauseChangeResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *PauseChangeResponseBodyEnvChange) SetStartTime(v string) *PauseChangeResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *PauseChangeResponseBodyEnvChange) SetChangeId(v string) *PauseChangeResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *PauseChangeResponseBodyEnvChange) SetEnvId(v string) *PauseChangeResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type PauseChangeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PauseChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PauseChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseChangeResponse) GoString() string {
	return s.String()
}

func (s *PauseChangeResponse) SetHeaders(v map[string]*string) *PauseChangeResponse {
	s.Headers = v
	return s
}

func (s *PauseChangeResponse) SetBody(v *PauseChangeResponseBody) *PauseChangeResponse {
	s.Body = v
	return s
}

type RebuildAppEnvRequest struct {
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	DryRun *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s RebuildAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppEnvRequest) GoString() string {
	return s.String()
}

func (s *RebuildAppEnvRequest) SetEnvId(v string) *RebuildAppEnvRequest {
	s.EnvId = &v
	return s
}

func (s *RebuildAppEnvRequest) SetDryRun(v bool) *RebuildAppEnvRequest {
	s.DryRun = &v
	return s
}

type RebuildAppEnvResponseBody struct {
	Message         *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvChangeDetail *RebuildAppEnvResponseBodyEnvChangeDetail `json:"EnvChangeDetail,omitempty" xml:"EnvChangeDetail,omitempty" type:"Struct"`
	Code            *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RebuildAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildAppEnvResponseBody) SetMessage(v string) *RebuildAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *RebuildAppEnvResponseBody) SetRequestId(v string) *RebuildAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildAppEnvResponseBody) SetEnvChangeDetail(v *RebuildAppEnvResponseBodyEnvChangeDetail) *RebuildAppEnvResponseBody {
	s.EnvChangeDetail = v
	return s
}

func (s *RebuildAppEnvResponseBody) SetCode(v string) *RebuildAppEnvResponseBody {
	s.Code = &v
	return s
}

type RebuildAppEnvResponseBodyEnvChangeDetail struct {
	StartTime  *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId   *string                                             `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId      *string                                             `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Operations *RebuildAppEnvResponseBodyEnvChangeDetailOperations `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Struct"`
}

func (s RebuildAppEnvResponseBodyEnvChangeDetail) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppEnvResponseBodyEnvChangeDetail) GoString() string {
	return s.String()
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetail) SetStartTime(v string) *RebuildAppEnvResponseBodyEnvChangeDetail {
	s.StartTime = &v
	return s
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetail) SetChangeId(v string) *RebuildAppEnvResponseBodyEnvChangeDetail {
	s.ChangeId = &v
	return s
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetail) SetEnvId(v string) *RebuildAppEnvResponseBodyEnvChangeDetail {
	s.EnvId = &v
	return s
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetail) SetOperations(v *RebuildAppEnvResponseBodyEnvChangeDetailOperations) *RebuildAppEnvResponseBodyEnvChangeDetail {
	s.Operations = v
	return s
}

type RebuildAppEnvResponseBodyEnvChangeDetailOperations struct {
	Operation []*RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Repeated"`
}

func (s RebuildAppEnvResponseBodyEnvChangeDetailOperations) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppEnvResponseBodyEnvChangeDetailOperations) GoString() string {
	return s.String()
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetailOperations) SetOperation(v []*RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation) *RebuildAppEnvResponseBodyEnvChangeDetailOperations {
	s.Operation = v
	return s
}

type RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation struct {
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation) GoString() string {
	return s.String()
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationDescription(v string) *RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationDescription = &v
	return s
}

func (s *RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationType(v string) *RebuildAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationType = &v
	return s
}

type RebuildAppEnvResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebuildAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebuildAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildAppEnvResponse) GoString() string {
	return s.String()
}

func (s *RebuildAppEnvResponse) SetHeaders(v map[string]*string) *RebuildAppEnvResponse {
	s.Headers = v
	return s
}

func (s *RebuildAppEnvResponse) SetBody(v *RebuildAppEnvResponseBody) *RebuildAppEnvResponse {
	s.Body = v
	return s
}

type RestartAppEnvRequest struct {
	EnvId               *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	BatchSize           *int32  `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	BatchPercent        *int32  `json:"BatchPercent,omitempty" xml:"BatchPercent,omitempty"`
	BatchInterval       *int32  `json:"BatchInterval,omitempty" xml:"BatchInterval,omitempty"`
	PauseBetweenBatches *bool   `json:"PauseBetweenBatches,omitempty" xml:"PauseBetweenBatches,omitempty"`
}

func (s RestartAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartAppEnvRequest) GoString() string {
	return s.String()
}

func (s *RestartAppEnvRequest) SetEnvId(v string) *RestartAppEnvRequest {
	s.EnvId = &v
	return s
}

func (s *RestartAppEnvRequest) SetBatchSize(v int32) *RestartAppEnvRequest {
	s.BatchSize = &v
	return s
}

func (s *RestartAppEnvRequest) SetBatchPercent(v int32) *RestartAppEnvRequest {
	s.BatchPercent = &v
	return s
}

func (s *RestartAppEnvRequest) SetBatchInterval(v int32) *RestartAppEnvRequest {
	s.BatchInterval = &v
	return s
}

func (s *RestartAppEnvRequest) SetPauseBetweenBatches(v bool) *RestartAppEnvRequest {
	s.PauseBetweenBatches = &v
	return s
}

type RestartAppEnvResponseBody struct {
	EnvChange *RestartAppEnvResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RestartAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *RestartAppEnvResponseBody) SetEnvChange(v *RestartAppEnvResponseBodyEnvChange) *RestartAppEnvResponseBody {
	s.EnvChange = v
	return s
}

func (s *RestartAppEnvResponseBody) SetMessage(v string) *RestartAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *RestartAppEnvResponseBody) SetRequestId(v string) *RestartAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartAppEnvResponseBody) SetCode(v string) *RestartAppEnvResponseBody {
	s.Code = &v
	return s
}

type RestartAppEnvResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s RestartAppEnvResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s RestartAppEnvResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *RestartAppEnvResponseBodyEnvChange) SetStartTime(v string) *RestartAppEnvResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *RestartAppEnvResponseBodyEnvChange) SetChangeId(v string) *RestartAppEnvResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *RestartAppEnvResponseBodyEnvChange) SetEnvId(v string) *RestartAppEnvResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type RestartAppEnvResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartAppEnvResponse) GoString() string {
	return s.String()
}

func (s *RestartAppEnvResponse) SetHeaders(v map[string]*string) *RestartAppEnvResponse {
	s.Headers = v
	return s
}

func (s *RestartAppEnvResponse) SetBody(v *RestartAppEnvResponseBody) *RestartAppEnvResponse {
	s.Body = v
	return s
}

type ResumeChangeRequest struct {
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
}

func (s ResumeChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeChangeRequest) GoString() string {
	return s.String()
}

func (s *ResumeChangeRequest) SetChangeId(v string) *ResumeChangeRequest {
	s.ChangeId = &v
	return s
}

type ResumeChangeResponseBody struct {
	EnvChange *ResumeChangeResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ResumeChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeChangeResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeChangeResponseBody) SetEnvChange(v *ResumeChangeResponseBodyEnvChange) *ResumeChangeResponseBody {
	s.EnvChange = v
	return s
}

func (s *ResumeChangeResponseBody) SetMessage(v string) *ResumeChangeResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeChangeResponseBody) SetRequestId(v string) *ResumeChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeChangeResponseBody) SetCode(v string) *ResumeChangeResponseBody {
	s.Code = &v
	return s
}

type ResumeChangeResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s ResumeChangeResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s ResumeChangeResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *ResumeChangeResponseBodyEnvChange) SetStartTime(v string) *ResumeChangeResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *ResumeChangeResponseBodyEnvChange) SetChangeId(v string) *ResumeChangeResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *ResumeChangeResponseBodyEnvChange) SetEnvId(v string) *ResumeChangeResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type ResumeChangeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeChangeResponse) GoString() string {
	return s.String()
}

func (s *ResumeChangeResponse) SetHeaders(v map[string]*string) *ResumeChangeResponse {
	s.Headers = v
	return s
}

func (s *ResumeChangeResponse) SetBody(v *ResumeChangeResponseBody) *ResumeChangeResponse {
	s.Body = v
	return s
}

type StartAppEnvRequest struct {
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s StartAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAppEnvRequest) GoString() string {
	return s.String()
}

func (s *StartAppEnvRequest) SetEnvId(v string) *StartAppEnvRequest {
	s.EnvId = &v
	return s
}

type StartAppEnvResponseBody struct {
	EnvChange *StartAppEnvResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StartAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *StartAppEnvResponseBody) SetEnvChange(v *StartAppEnvResponseBodyEnvChange) *StartAppEnvResponseBody {
	s.EnvChange = v
	return s
}

func (s *StartAppEnvResponseBody) SetMessage(v string) *StartAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *StartAppEnvResponseBody) SetRequestId(v string) *StartAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAppEnvResponseBody) SetCode(v string) *StartAppEnvResponseBody {
	s.Code = &v
	return s
}

type StartAppEnvResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s StartAppEnvResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s StartAppEnvResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *StartAppEnvResponseBodyEnvChange) SetStartTime(v string) *StartAppEnvResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *StartAppEnvResponseBodyEnvChange) SetChangeId(v string) *StartAppEnvResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *StartAppEnvResponseBodyEnvChange) SetEnvId(v string) *StartAppEnvResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type StartAppEnvResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAppEnvResponse) GoString() string {
	return s.String()
}

func (s *StartAppEnvResponse) SetHeaders(v map[string]*string) *StartAppEnvResponse {
	s.Headers = v
	return s
}

func (s *StartAppEnvResponse) SetBody(v *StartAppEnvResponseBody) *StartAppEnvResponse {
	s.Body = v
	return s
}

type StopAppEnvRequest struct {
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s StopAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppEnvRequest) GoString() string {
	return s.String()
}

func (s *StopAppEnvRequest) SetEnvId(v string) *StopAppEnvRequest {
	s.EnvId = &v
	return s
}

type StopAppEnvResponseBody struct {
	EnvChange *StopAppEnvResponseBodyEnvChange `json:"EnvChange,omitempty" xml:"EnvChange,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StopAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppEnvResponseBody) SetEnvChange(v *StopAppEnvResponseBodyEnvChange) *StopAppEnvResponseBody {
	s.EnvChange = v
	return s
}

func (s *StopAppEnvResponseBody) SetMessage(v string) *StopAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *StopAppEnvResponseBody) SetRequestId(v string) *StopAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAppEnvResponseBody) SetCode(v string) *StopAppEnvResponseBody {
	s.Code = &v
	return s
}

type StopAppEnvResponseBodyEnvChange struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId  *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s StopAppEnvResponseBodyEnvChange) String() string {
	return tea.Prettify(s)
}

func (s StopAppEnvResponseBodyEnvChange) GoString() string {
	return s.String()
}

func (s *StopAppEnvResponseBodyEnvChange) SetStartTime(v string) *StopAppEnvResponseBodyEnvChange {
	s.StartTime = &v
	return s
}

func (s *StopAppEnvResponseBodyEnvChange) SetChangeId(v string) *StopAppEnvResponseBodyEnvChange {
	s.ChangeId = &v
	return s
}

func (s *StopAppEnvResponseBodyEnvChange) SetEnvId(v string) *StopAppEnvResponseBodyEnvChange {
	s.EnvId = &v
	return s
}

type StopAppEnvResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppEnvResponse) GoString() string {
	return s.String()
}

func (s *StopAppEnvResponse) SetHeaders(v map[string]*string) *StopAppEnvResponse {
	s.Headers = v
	return s
}

func (s *StopAppEnvResponse) SetBody(v *StopAppEnvResponseBody) *StopAppEnvResponse {
	s.Body = v
	return s
}

type TerminateAppEnvRequest struct {
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s TerminateAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateAppEnvRequest) GoString() string {
	return s.String()
}

func (s *TerminateAppEnvRequest) SetEnvId(v string) *TerminateAppEnvRequest {
	s.EnvId = &v
	return s
}

func (s *TerminateAppEnvRequest) SetDryRun(v string) *TerminateAppEnvRequest {
	s.DryRun = &v
	return s
}

type TerminateAppEnvResponseBody struct {
	Message         *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvChangeDetail *TerminateAppEnvResponseBodyEnvChangeDetail `json:"EnvChangeDetail,omitempty" xml:"EnvChangeDetail,omitempty" type:"Struct"`
	Code            *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TerminateAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateAppEnvResponseBody) SetMessage(v string) *TerminateAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *TerminateAppEnvResponseBody) SetRequestId(v string) *TerminateAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateAppEnvResponseBody) SetEnvChangeDetail(v *TerminateAppEnvResponseBodyEnvChangeDetail) *TerminateAppEnvResponseBody {
	s.EnvChangeDetail = v
	return s
}

func (s *TerminateAppEnvResponseBody) SetCode(v string) *TerminateAppEnvResponseBody {
	s.Code = &v
	return s
}

type TerminateAppEnvResponseBodyEnvChangeDetail struct {
	StartTime  *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId   *string                                               `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId      *string                                               `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Operations *TerminateAppEnvResponseBodyEnvChangeDetailOperations `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Struct"`
}

func (s TerminateAppEnvResponseBodyEnvChangeDetail) String() string {
	return tea.Prettify(s)
}

func (s TerminateAppEnvResponseBodyEnvChangeDetail) GoString() string {
	return s.String()
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetail) SetStartTime(v string) *TerminateAppEnvResponseBodyEnvChangeDetail {
	s.StartTime = &v
	return s
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetail) SetChangeId(v string) *TerminateAppEnvResponseBodyEnvChangeDetail {
	s.ChangeId = &v
	return s
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetail) SetEnvId(v string) *TerminateAppEnvResponseBodyEnvChangeDetail {
	s.EnvId = &v
	return s
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetail) SetOperations(v *TerminateAppEnvResponseBodyEnvChangeDetailOperations) *TerminateAppEnvResponseBodyEnvChangeDetail {
	s.Operations = v
	return s
}

type TerminateAppEnvResponseBodyEnvChangeDetailOperations struct {
	Operation []*TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Repeated"`
}

func (s TerminateAppEnvResponseBodyEnvChangeDetailOperations) String() string {
	return tea.Prettify(s)
}

func (s TerminateAppEnvResponseBodyEnvChangeDetailOperations) GoString() string {
	return s.String()
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetailOperations) SetOperation(v []*TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation) *TerminateAppEnvResponseBodyEnvChangeDetailOperations {
	s.Operation = v
	return s
}

type TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation struct {
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation) String() string {
	return tea.Prettify(s)
}

func (s TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation) GoString() string {
	return s.String()
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationDescription(v string) *TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationDescription = &v
	return s
}

func (s *TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationType(v string) *TerminateAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationType = &v
	return s
}

type TerminateAppEnvResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TerminateAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TerminateAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateAppEnvResponse) GoString() string {
	return s.String()
}

func (s *TerminateAppEnvResponse) SetHeaders(v map[string]*string) *TerminateAppEnvResponse {
	s.Headers = v
	return s
}

func (s *TerminateAppEnvResponse) SetBody(v *TerminateAppEnvResponseBody) *TerminateAppEnvResponse {
	s.Body = v
	return s
}

type UpdateAppEnvRequest struct {
	EnvDescription      *string `json:"EnvDescription,omitempty" xml:"EnvDescription,omitempty"`
	EnvId               *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PkgVersionId        *string `json:"PkgVersionId,omitempty" xml:"PkgVersionId,omitempty"`
	OptionSettings      *string `json:"OptionSettings,omitempty" xml:"OptionSettings,omitempty"`
	StackId             *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ExtraProperties     *string `json:"ExtraProperties,omitempty" xml:"ExtraProperties,omitempty"`
	BatchSize           *string `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	BatchPercent        *string `json:"BatchPercent,omitempty" xml:"BatchPercent,omitempty"`
	BatchInterval       *string `json:"BatchInterval,omitempty" xml:"BatchInterval,omitempty"`
	PauseBetweenBatches *bool   `json:"PauseBetweenBatches,omitempty" xml:"PauseBetweenBatches,omitempty"`
}

func (s UpdateAppEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppEnvRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppEnvRequest) SetEnvDescription(v string) *UpdateAppEnvRequest {
	s.EnvDescription = &v
	return s
}

func (s *UpdateAppEnvRequest) SetEnvId(v string) *UpdateAppEnvRequest {
	s.EnvId = &v
	return s
}

func (s *UpdateAppEnvRequest) SetPkgVersionId(v string) *UpdateAppEnvRequest {
	s.PkgVersionId = &v
	return s
}

func (s *UpdateAppEnvRequest) SetOptionSettings(v string) *UpdateAppEnvRequest {
	s.OptionSettings = &v
	return s
}

func (s *UpdateAppEnvRequest) SetStackId(v string) *UpdateAppEnvRequest {
	s.StackId = &v
	return s
}

func (s *UpdateAppEnvRequest) SetDryRun(v bool) *UpdateAppEnvRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAppEnvRequest) SetExtraProperties(v string) *UpdateAppEnvRequest {
	s.ExtraProperties = &v
	return s
}

func (s *UpdateAppEnvRequest) SetBatchSize(v string) *UpdateAppEnvRequest {
	s.BatchSize = &v
	return s
}

func (s *UpdateAppEnvRequest) SetBatchPercent(v string) *UpdateAppEnvRequest {
	s.BatchPercent = &v
	return s
}

func (s *UpdateAppEnvRequest) SetBatchInterval(v string) *UpdateAppEnvRequest {
	s.BatchInterval = &v
	return s
}

func (s *UpdateAppEnvRequest) SetPauseBetweenBatches(v bool) *UpdateAppEnvRequest {
	s.PauseBetweenBatches = &v
	return s
}

type UpdateAppEnvResponseBody struct {
	Message         *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvChangeDetail *UpdateAppEnvResponseBodyEnvChangeDetail `json:"EnvChangeDetail,omitempty" xml:"EnvChangeDetail,omitempty" type:"Struct"`
	Code            *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateAppEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppEnvResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppEnvResponseBody) SetMessage(v string) *UpdateAppEnvResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppEnvResponseBody) SetRequestId(v string) *UpdateAppEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppEnvResponseBody) SetEnvChangeDetail(v *UpdateAppEnvResponseBodyEnvChangeDetail) *UpdateAppEnvResponseBody {
	s.EnvChangeDetail = v
	return s
}

func (s *UpdateAppEnvResponseBody) SetCode(v string) *UpdateAppEnvResponseBody {
	s.Code = &v
	return s
}

type UpdateAppEnvResponseBodyEnvChangeDetail struct {
	StartTime  *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChangeId   *string                                            `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	EnvId      *string                                            `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Operations *UpdateAppEnvResponseBodyEnvChangeDetailOperations `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Struct"`
}

func (s UpdateAppEnvResponseBodyEnvChangeDetail) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppEnvResponseBodyEnvChangeDetail) GoString() string {
	return s.String()
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetail) SetStartTime(v string) *UpdateAppEnvResponseBodyEnvChangeDetail {
	s.StartTime = &v
	return s
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetail) SetChangeId(v string) *UpdateAppEnvResponseBodyEnvChangeDetail {
	s.ChangeId = &v
	return s
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetail) SetEnvId(v string) *UpdateAppEnvResponseBodyEnvChangeDetail {
	s.EnvId = &v
	return s
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetail) SetOperations(v *UpdateAppEnvResponseBodyEnvChangeDetailOperations) *UpdateAppEnvResponseBodyEnvChangeDetail {
	s.Operations = v
	return s
}

type UpdateAppEnvResponseBodyEnvChangeDetailOperations struct {
	Operation []*UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Repeated"`
}

func (s UpdateAppEnvResponseBodyEnvChangeDetailOperations) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppEnvResponseBodyEnvChangeDetailOperations) GoString() string {
	return s.String()
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetailOperations) SetOperation(v []*UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation) *UpdateAppEnvResponseBodyEnvChangeDetailOperations {
	s.Operation = v
	return s
}

type UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation struct {
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation) GoString() string {
	return s.String()
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationDescription(v string) *UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationDescription = &v
	return s
}

func (s *UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation) SetOperationType(v string) *UpdateAppEnvResponseBodyEnvChangeDetailOperationsOperation {
	s.OperationType = &v
	return s
}

type UpdateAppEnvResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppEnvResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppEnvResponse) SetHeaders(v map[string]*string) *UpdateAppEnvResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppEnvResponse) SetBody(v *UpdateAppEnvResponseBody) *UpdateAppEnvResponse {
	s.Body = v
	return s
}

type UpdateApplicationRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) SetAppId(v string) *UpdateApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationRequest) SetAppDescription(v string) *UpdateApplicationRequest {
	s.AppDescription = &v
	return s
}

type UpdateApplicationResponseBody struct {
	Message     *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Application *UpdateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
}

func (s UpdateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) SetMessage(v string) *UpdateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetCode(v string) *UpdateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody {
	s.Application = v
	return s
}

type UpdateApplicationResponseBodyApplication struct {
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateUsername *string `json:"CreateUsername,omitempty" xml:"CreateUsername,omitempty"`
	UpdateTime     *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUsername *string `json:"UpdateUsername,omitempty" xml:"UpdateUsername,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CategoryName   *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
}

func (s UpdateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplication) SetAppName(v string) *UpdateApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetCreateUsername(v string) *UpdateApplicationResponseBodyApplication {
	s.CreateUsername = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetUpdateTime(v int64) *UpdateApplicationResponseBodyApplication {
	s.UpdateTime = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetUpdateUsername(v string) *UpdateApplicationResponseBodyApplication {
	s.UpdateUsername = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetCreateTime(v int64) *UpdateApplicationResponseBodyApplication {
	s.CreateTime = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppId(v string) *UpdateApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetCategoryName(v string) *UpdateApplicationResponseBodyApplication {
	s.CategoryName = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppDescription(v string) *UpdateApplicationResponseBodyApplication {
	s.AppDescription = &v
	return s
}

type UpdateApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponse) SetHeaders(v map[string]*string) *UpdateApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationResponse) SetBody(v *UpdateApplicationResponseBody) *UpdateApplicationResponse {
	s.Body = v
	return s
}

type UpdateConfigTemplateRequest struct {
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	OptionSettings      *string `json:"OptionSettings,omitempty" xml:"OptionSettings,omitempty"`
}

func (s UpdateConfigTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateRequest) SetTemplateDescription(v string) *UpdateConfigTemplateRequest {
	s.TemplateDescription = &v
	return s
}

func (s *UpdateConfigTemplateRequest) SetTemplateId(v string) *UpdateConfigTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateConfigTemplateRequest) SetOptionSettings(v string) *UpdateConfigTemplateRequest {
	s.OptionSettings = &v
	return s
}

type UpdateConfigTemplateResponseBody struct {
	ConfigTemplate *UpdateConfigTemplateResponseBodyConfigTemplate `json:"ConfigTemplate,omitempty" xml:"ConfigTemplate,omitempty" type:"Struct"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateConfigTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateResponseBody) SetConfigTemplate(v *UpdateConfigTemplateResponseBodyConfigTemplate) *UpdateConfigTemplateResponseBody {
	s.ConfigTemplate = v
	return s
}

func (s *UpdateConfigTemplateResponseBody) SetMessage(v string) *UpdateConfigTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigTemplateResponseBody) SetRequestId(v string) *UpdateConfigTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigTemplateResponseBody) SetCode(v string) *UpdateConfigTemplateResponseBody {
	s.Code = &v
	return s
}

type UpdateConfigTemplateResponseBodyConfigTemplate struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UpdateTime   *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	StackId      *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StackName    *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateConfigTemplateResponseBodyConfigTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigTemplateResponseBodyConfigTemplate) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetAppName(v string) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.AppName = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetUpdateTime(v int64) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.UpdateTime = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetStackId(v string) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.StackId = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetCreateTime(v int64) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.CreateTime = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetAppId(v string) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.AppId = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetStackName(v string) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.StackName = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetTemplateName(v string) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.TemplateName = &v
	return s
}

func (s *UpdateConfigTemplateResponseBodyConfigTemplate) SetTemplateId(v string) *UpdateConfigTemplateResponseBodyConfigTemplate {
	s.TemplateId = &v
	return s
}

type UpdateConfigTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateResponse) SetHeaders(v map[string]*string) *UpdateConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigTemplateResponse) SetBody(v *UpdateConfigTemplateResponseBody) *UpdateConfigTemplateResponse {
	s.Body = v
	return s
}

type ValidateConfigSettingRequest struct {
	EnvId          *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	StackId        *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	OptionSettings *string `json:"OptionSettings,omitempty" xml:"OptionSettings,omitempty"`
}

func (s ValidateConfigSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateConfigSettingRequest) GoString() string {
	return s.String()
}

func (s *ValidateConfigSettingRequest) SetEnvId(v string) *ValidateConfigSettingRequest {
	s.EnvId = &v
	return s
}

func (s *ValidateConfigSettingRequest) SetTemplateId(v string) *ValidateConfigSettingRequest {
	s.TemplateId = &v
	return s
}

func (s *ValidateConfigSettingRequest) SetStackId(v string) *ValidateConfigSettingRequest {
	s.StackId = &v
	return s
}

func (s *ValidateConfigSettingRequest) SetOptionSettings(v string) *ValidateConfigSettingRequest {
	s.OptionSettings = &v
	return s
}

type ValidateConfigSettingResponseBody struct {
	Message                 *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConfigValidationResults *ValidateConfigSettingResponseBodyConfigValidationResults `json:"ConfigValidationResults,omitempty" xml:"ConfigValidationResults,omitempty" type:"Struct"`
	Code                    *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ValidateConfigSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateConfigSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateConfigSettingResponseBody) SetMessage(v string) *ValidateConfigSettingResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateConfigSettingResponseBody) SetRequestId(v string) *ValidateConfigSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateConfigSettingResponseBody) SetConfigValidationResults(v *ValidateConfigSettingResponseBodyConfigValidationResults) *ValidateConfigSettingResponseBody {
	s.ConfigValidationResults = v
	return s
}

func (s *ValidateConfigSettingResponseBody) SetCode(v string) *ValidateConfigSettingResponseBody {
	s.Code = &v
	return s
}

type ValidateConfigSettingResponseBodyConfigValidationResults struct {
	ConfigValidationResult []*ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult `json:"ConfigValidationResult,omitempty" xml:"ConfigValidationResult,omitempty" type:"Repeated"`
}

func (s ValidateConfigSettingResponseBodyConfigValidationResults) String() string {
	return tea.Prettify(s)
}

func (s ValidateConfigSettingResponseBodyConfigValidationResults) GoString() string {
	return s.String()
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResults) SetConfigValidationResult(v []*ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) *ValidateConfigSettingResponseBodyConfigValidationResults {
	s.ConfigValidationResult = v
	return s
}

type ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult struct {
	OptionName     *string                                                                                     `json:"OptionName,omitempty" xml:"OptionName,omitempty"`
	PathName       *string                                                                                     `json:"PathName,omitempty" xml:"PathName,omitempty"`
	ResultMessage  *string                                                                                     `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultSeverity *string                                                                                     `json:"ResultSeverity,omitempty" xml:"ResultSeverity,omitempty"`
	ConfigOption   *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption `json:"ConfigOption,omitempty" xml:"ConfigOption,omitempty" type:"Struct"`
}

func (s ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) String() string {
	return tea.Prettify(s)
}

func (s ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) GoString() string {
	return s.String()
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) SetOptionName(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult {
	s.OptionName = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) SetPathName(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult {
	s.PathName = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) SetResultMessage(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult {
	s.ResultMessage = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) SetResultSeverity(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult {
	s.ResultSeverity = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult) SetConfigOption(v *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResult {
	s.ConfigOption = v
	return s
}

type ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption struct {
	RegexDesc         *string `json:"RegexDesc,omitempty" xml:"RegexDesc,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	EditorType        *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	MaxLength         *int32  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	OptionLabel       *string `json:"OptionLabel,omitempty" xml:"OptionLabel,omitempty"`
	RegexPattern      *string `json:"RegexPattern,omitempty" xml:"RegexPattern,omitempty"`
	ChangeSeverity    *string `json:"ChangeSeverity,omitempty" xml:"ChangeSeverity,omitempty"`
	OptionDescription *string `json:"OptionDescription,omitempty" xml:"OptionDescription,omitempty"`
	OptionName        *string `json:"OptionName,omitempty" xml:"OptionName,omitempty"`
	PathName          *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	HiddenOption      *bool   `json:"HiddenOption,omitempty" xml:"HiddenOption,omitempty"`
	ValueType         *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	MinLength         *int32  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	ValueOptions      *string `json:"ValueOptions,omitempty" xml:"ValueOptions,omitempty"`
	ReadonlyOption    *bool   `json:"ReadonlyOption,omitempty" xml:"ReadonlyOption,omitempty"`
}

func (s ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) String() string {
	return tea.Prettify(s)
}

func (s ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) GoString() string {
	return s.String()
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetRegexDesc(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.RegexDesc = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetMaxValue(v int64) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.MaxValue = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetEditorType(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.EditorType = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetMinValue(v int64) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.MinValue = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetDefaultValue(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.DefaultValue = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetMaxLength(v int32) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.MaxLength = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetOptionLabel(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.OptionLabel = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetRegexPattern(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.RegexPattern = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetChangeSeverity(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.ChangeSeverity = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetOptionDescription(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.OptionDescription = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetOptionName(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.OptionName = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetPathName(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.PathName = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetHiddenOption(v bool) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.HiddenOption = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetValueType(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.ValueType = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetMinLength(v int32) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.MinLength = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetValueOptions(v string) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.ValueOptions = &v
	return s
}

func (s *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption) SetReadonlyOption(v bool) *ValidateConfigSettingResponseBodyConfigValidationResultsConfigValidationResultConfigOption {
	s.ReadonlyOption = &v
	return s
}

type ValidateConfigSettingResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateConfigSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateConfigSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateConfigSettingResponse) GoString() string {
	return s.String()
}

func (s *ValidateConfigSettingResponse) SetHeaders(v map[string]*string) *ValidateConfigSettingResponse {
	s.Headers = v
	return s
}

func (s *ValidateConfigSettingResponse) SetBody(v *ValidateConfigSettingResponseBody) *ValidateConfigSettingResponse {
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
		"cn-beijing":            tea.String("webplus.cn-hangzhou.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("webplus.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("webplus.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           tea.String("webplus.cn-hangzhou.aliyuncs.com"),
		"ap-northeast-1":        tea.String("webplus.aliyuncs.com"),
		"ap-south-1":            tea.String("webplus.aliyuncs.com"),
		"ap-southeast-1":        tea.String("webplus.aliyuncs.com"),
		"ap-southeast-2":        tea.String("webplus.aliyuncs.com"),
		"ap-southeast-3":        tea.String("webplus.aliyuncs.com"),
		"ap-southeast-5":        tea.String("webplus.aliyuncs.com"),
		"cn-chengdu":            tea.String("webplus.aliyuncs.com"),
		"cn-hongkong":           tea.String("webplus-vpc.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          tea.String("webplus.aliyuncs.com"),
		"cn-qingdao":            tea.String("webplus.aliyuncs.com"),
		"eu-central-1":          tea.String("webplus.aliyuncs.com"),
		"eu-west-1":             tea.String("webplus.aliyuncs.com"),
		"me-east-1":             tea.String("webplus.aliyuncs.com"),
		"us-east-1":             tea.String("webplus.aliyuncs.com"),
		"us-west-1":             tea.String("webplus.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("webplus.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("webplus.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("webplus.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("webplus.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("webplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AbortChange(request *AbortChangeRequest) (_result *AbortChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortChangeResponse{}
	_body, _err := client.AbortChangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortChangeWithOptions(request *AbortChangeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AbortChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		body["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AbortChangeResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("AbortChange"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/change/abort"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppEnv(request *CreateAppEnvRequest) (_result *CreateAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppEnvResponse{}
	_body, _err := client.CreateAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppEnvWithOptions(request *CreateAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvName)) {
		body["EnvName"] = request.EnvName
	}

	if !tea.BoolValue(util.IsUnset(request.EnvDescription)) {
		body["EnvDescription"] = request.EnvDescription
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		body["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionId)) {
		body["PkgVersionId"] = request.PkgVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionSettings)) {
		body["OptionSettings"] = request.OptionSettings
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileName)) {
		body["ProfileName"] = request.ProfileName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEnvId)) {
		body["SourceEnvId"] = request.SourceEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraProperties)) {
		body["ExtraProperties"] = request.ExtraProperties
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreateAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppDescription)) {
		body["AppDescription"] = request.AppDescription
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["CategoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.UsingSharedStorage)) {
		body["UsingSharedStorage"] = request.UsingSharedStorage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreateApplication"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/application"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfigTemplate(request *CreateConfigTemplateRequest) (_result *CreateConfigTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigTemplateResponse{}
	_body, _err := client.CreateConfigTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigTemplateWithOptions(request *CreateConfigTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConfigTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateDescription)) {
		body["TemplateDescription"] = request.TemplateDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		body["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTemplateId)) {
		body["SourceTemplateId"] = request.SourceTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEnvId)) {
		body["SourceEnvId"] = request.SourceEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileName)) {
		body["ProfileName"] = request.ProfileName
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionId)) {
		body["PkgVersionId"] = request.PkgVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionSettings)) {
		body["OptionSettings"] = request.OptionSettings
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateConfigTemplateResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreateConfigTemplate"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/configTemplate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["ProductName"] = request.ProductName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreateOrder"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/paas/createOrder"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePkgVersion(request *CreatePkgVersionRequest) (_result *CreatePkgVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePkgVersionResponse{}
	_body, _err := client.CreatePkgVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePkgVersionWithOptions(request *CreatePkgVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePkgVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PkgVersionLabel)) {
		body["PkgVersionLabel"] = request.PkgVersionLabel
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionDescription)) {
		body["PkgVersionDescription"] = request.PkgVersionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageSource)) {
		body["PackageSource"] = request.PackageSource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreatePkgVersionResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("CreatePkgVersion"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/pkgVersion"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStorage() (_result *CreateStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateStorageResponse{}
	_body, _err := client.CreateStorageWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStorageWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateStorageResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateStorageResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateStorage"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/storage"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppEnv(request *DeleteAppEnvRequest) (_result *DeleteAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAppEnvResponse{}
	_body, _err := client.DeleteAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppEnvWithOptions(request *DeleteAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteAppEnvResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteApplication"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v1/wam/application"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteChange(request *DeleteChangeRequest) (_result *DeleteChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteChangeResponse{}
	_body, _err := client.DeleteChangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteChangeWithOptions(request *DeleteChangeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		query["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteChangeResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteChange"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v1/wam/change"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (_result *DeleteConfigTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigTemplateResponse{}
	_body, _err := client.DeleteConfigTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigTemplateWithOptions(request *DeleteConfigTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteConfigTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteConfigTemplate"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v1/wam/configTemplate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePkgVersion(request *DeletePkgVersionRequest) (_result *DeletePkgVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePkgVersionResponse{}
	_body, _err := client.DeletePkgVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePkgVersionWithOptions(request *DeletePkgVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePkgVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PkgVersionId)) {
		query["PkgVersionId"] = request.PkgVersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeletePkgVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePkgVersion"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v1/wam/pkgVersion"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployAppEnv(request *DeployAppEnvRequest) (_result *DeployAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployAppEnvResponse{}
	_body, _err := client.DeployAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployAppEnvWithOptions(request *DeployAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeployAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.BatchSize)) {
		body["BatchSize"] = request.BatchSize
	}

	if !tea.BoolValue(util.IsUnset(request.BatchPercent)) {
		body["BatchPercent"] = request.BatchPercent
	}

	if !tea.BoolValue(util.IsUnset(request.BatchInterval)) {
		body["BatchInterval"] = request.BatchInterval
	}

	if !tea.BoolValue(util.IsUnset(request.PauseBetweenBatches)) {
		body["PauseBetweenBatches"] = request.PauseBetweenBatches
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionId)) {
		body["PkgVersionId"] = request.PkgVersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &DeployAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("DeployAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/deploy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppEnvInstanceHealth(request *DescribeAppEnvInstanceHealthRequest) (_result *DescribeAppEnvInstanceHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppEnvInstanceHealthResponse{}
	_body, _err := client.DescribeAppEnvInstanceHealthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppEnvInstanceHealthWithOptions(request *DescribeAppEnvInstanceHealthRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppEnvInstanceHealthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeAppEnvInstanceHealthResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeAppEnvInstanceHealth"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/instanceHealth"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppEnvs(request *DescribeAppEnvsRequest) (_result *DescribeAppEnvsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppEnvsResponse{}
	_body, _err := client.DescribeAppEnvsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppEnvsWithOptions(request *DescribeAppEnvsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppEnvsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeTerminated)) {
		query["IncludeTerminated"] = request.IncludeTerminated
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.EnvName)) {
		query["EnvName"] = request.EnvName
	}

	if !tea.BoolValue(util.IsUnset(request.EnvSearch)) {
		query["EnvSearch"] = request.EnvSearch
	}

	if !tea.BoolValue(util.IsUnset(request.RecentUpdated)) {
		query["RecentUpdated"] = request.RecentUpdated
	}

	if !tea.BoolValue(util.IsUnset(request.StackSearch)) {
		query["StackSearch"] = request.StackSearch
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeAppEnvsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeAppEnvs"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppEnvStatus(request *DescribeAppEnvStatusRequest) (_result *DescribeAppEnvStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppEnvStatusResponse{}
	_body, _err := client.DescribeAppEnvStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppEnvStatusWithOptions(request *DescribeAppEnvStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppEnvStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeAppEnvStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeAppEnvStatus"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/status"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplications(request *DescribeApplicationsRequest) (_result *DescribeApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationsResponse{}
	_body, _err := client.DescribeApplicationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationsWithOptions(request *DescribeApplicationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppSearch)) {
		query["AppSearch"] = request.AppSearch
	}

	if !tea.BoolValue(util.IsUnset(request.EnvSearch)) {
		query["EnvSearch"] = request.EnvSearch
	}

	if !tea.BoolValue(util.IsUnset(request.StackSearch)) {
		query["StackSearch"] = request.StackSearch
	}

	if !tea.BoolValue(util.IsUnset(request.CategorySearch)) {
		query["CategorySearch"] = request.CategorySearch
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeApplicationsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeApplications"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/application"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCategories() (_result *DescribeCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCategoriesResponse{}
	_body, _err := client.DescribeCategoriesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCategoriesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCategoriesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeCategoriesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeCategories"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/category"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChange(request *DescribeChangeRequest) (_result *DescribeChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChangeResponse{}
	_body, _err := client.DescribeChangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChangeWithOptions(request *DescribeChangeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		query["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeChangeResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeChange"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/changeInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChanges(request *DescribeChangesRequest) (_result *DescribeChangesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChangesResponse{}
	_body, _err := client.DescribeChangesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChangesWithOptions(request *DescribeChangesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChangesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeChangesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeChanges"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/change"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigIndex(request *DescribeConfigIndexRequest) (_result *DescribeConfigIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigIndexResponse{}
	_body, _err := client.DescribeConfigIndexWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigIndexWithOptions(request *DescribeConfigIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileName)) {
		query["ProfileName"] = request.ProfileName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeConfigIndexResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeConfigIndex"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/config/configIndex"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigOptions(request *DescribeConfigOptionsRequest) (_result *DescribeConfigOptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigOptionsResponse{}
	_body, _err := client.DescribeConfigOptionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigOptionsWithOptions(request *DescribeConfigOptionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileName)) {
		query["ProfileName"] = request.ProfileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeConfigOptionsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeConfigOptions"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/config/configOption"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigSettings(request *DescribeConfigSettingsRequest) (_result *DescribeConfigSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigSettingsResponse{}
	_body, _err := client.DescribeConfigSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigSettingsWithOptions(request *DescribeConfigSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.PathName)) {
		query["PathName"] = request.PathName
	}

	if !tea.BoolValue(util.IsUnset(request.OptionName)) {
		query["OptionName"] = request.OptionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeConfigSettingsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeConfigSettings"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/config/configSetting"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigTemplates(request *DescribeConfigTemplatesRequest) (_result *DescribeConfigTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigTemplatesResponse{}
	_body, _err := client.DescribeConfigTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigTemplatesWithOptions(request *DescribeConfigTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateSearch)) {
		query["TemplateSearch"] = request.TemplateSearch
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeConfigTemplatesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeConfigTemplates"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/configTemplate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnvResource(request *DescribeEnvResourceRequest) (_result *DescribeEnvResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEnvResourceResponse{}
	_body, _err := client.DescribeEnvResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnvResourceWithOptions(request *DescribeEnvResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeEnvResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeEnvResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeEnvResource"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/envResource"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		query["ChangeId"] = request.ChangeId
	}

	if !tea.BoolValue(util.IsUnset(request.LastChangeEvents)) {
		query["LastChangeEvents"] = request.LastChangeEvents
	}

	if !tea.BoolValue(util.IsUnset(request.ReverseByTimestamp)) {
		query["ReverseByTimestamp"] = request.ReverseByTimestamp
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeEvents"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/event"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatherLogResult(request *DescribeGatherLogResultRequest) (_result *DescribeGatherLogResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeGatherLogResultResponse{}
	_body, _err := client.DescribeGatherLogResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatherLogResultWithOptions(request *DescribeGatherLogResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeGatherLogResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		query["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeGatherLogResultResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeGatherLogResult"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/gatherLog"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatherStatsResult(request *DescribeGatherStatsResultRequest) (_result *DescribeGatherStatsResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeGatherStatsResultResponse{}
	_body, _err := client.DescribeGatherStatsResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatherStatsResultWithOptions(request *DescribeGatherStatsResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeGatherStatsResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		query["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeGatherStatsResultResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeGatherStatsResult"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/gatherStats"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceHealth(request *DescribeInstanceHealthRequest) (_result *DescribeInstanceHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceHealthResponse{}
	_body, _err := client.DescribeInstanceHealthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceHealthWithOptions(request *DescribeInstanceHealthRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceHealthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeInstanceHealthResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeInstanceHealth"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/instance/health"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePkgVersions(request *DescribePkgVersionsRequest) (_result *DescribePkgVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePkgVersionsResponse{}
	_body, _err := client.DescribePkgVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePkgVersionsWithOptions(request *DescribePkgVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribePkgVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionLabel)) {
		query["PkgVersionLabel"] = request.PkgVersionLabel
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionSearch)) {
		query["PkgVersionSearch"] = request.PkgVersionSearch
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribePkgVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribePkgVersions"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/pkgVersion"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePublicConfigTemplates(request *DescribePublicConfigTemplatesRequest) (_result *DescribePublicConfigTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePublicConfigTemplatesResponse{}
	_body, _err := client.DescribePublicConfigTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePublicConfigTemplatesWithOptions(request *DescribePublicConfigTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribePublicConfigTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		query["CategoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribePublicConfigTemplatesResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribePublicConfigTemplates"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/publicConfigTemplate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStacks(request *DescribeStacksRequest) (_result *DescribeStacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeStacksResponse{}
	_body, _err := client.DescribeStacksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStacksWithOptions(request *DescribeStacksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeStacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecommendedOnly)) {
		query["RecommendedOnly"] = request.RecommendedOnly
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		query["CategoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeStacksResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeStacks"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/stack"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStorage(request *DescribeStorageRequest) (_result *DescribeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeStorageResponse{}
	_body, _err := client.DescribeStorageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStorageWithOptions(request *DescribeStorageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UsingSharedStorage)) {
		query["UsingSharedStorage"] = request.UsingSharedStorage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeStorageResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeStorage"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v1/wam/storage"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GatherAppEnvLog(request *GatherAppEnvLogRequest) (_result *GatherAppEnvLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GatherAppEnvLogResponse{}
	_body, _err := client.GatherAppEnvLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GatherAppEnvLogWithOptions(request *GatherAppEnvLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GatherAppEnvLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstances)) {
		body["TargetInstances"] = request.TargetInstances
	}

	if !tea.BoolValue(util.IsUnset(request.LogPath)) {
		body["LogPath"] = request.LogPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GatherAppEnvLogResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("GatherAppEnvLog"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/gatherLog"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GatherAppEnvStats(request *GatherAppEnvStatsRequest) (_result *GatherAppEnvStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GatherAppEnvStatsResponse{}
	_body, _err := client.GatherAppEnvStatsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GatherAppEnvStatsWithOptions(request *GatherAppEnvStatsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GatherAppEnvStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstances)) {
		body["TargetInstances"] = request.TargetInstances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GatherAppEnvStatsResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("GatherAppEnvStats"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/gatherStats"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseChange(request *PauseChangeRequest) (_result *PauseChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseChangeResponse{}
	_body, _err := client.PauseChangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseChangeWithOptions(request *PauseChangeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PauseChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		body["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &PauseChangeResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("PauseChange"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/change/pause"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildAppEnv(request *RebuildAppEnvRequest) (_result *RebuildAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebuildAppEnvResponse{}
	_body, _err := client.RebuildAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildAppEnvWithOptions(request *RebuildAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RebuildAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RebuildAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("RebuildAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/rebuild"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartAppEnv(request *RestartAppEnvRequest) (_result *RestartAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartAppEnvResponse{}
	_body, _err := client.RestartAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartAppEnvWithOptions(request *RestartAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.BatchSize)) {
		body["BatchSize"] = request.BatchSize
	}

	if !tea.BoolValue(util.IsUnset(request.BatchPercent)) {
		body["BatchPercent"] = request.BatchPercent
	}

	if !tea.BoolValue(util.IsUnset(request.BatchInterval)) {
		body["BatchInterval"] = request.BatchInterval
	}

	if !tea.BoolValue(util.IsUnset(request.PauseBetweenBatches)) {
		body["PauseBetweenBatches"] = request.PauseBetweenBatches
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RestartAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("RestartAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/restart"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeChange(request *ResumeChangeRequest) (_result *ResumeChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeChangeResponse{}
	_body, _err := client.ResumeChangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeChangeWithOptions(request *ResumeChangeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeId)) {
		body["ChangeId"] = request.ChangeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ResumeChangeResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ResumeChange"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/change/resume"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAppEnv(request *StartAppEnvRequest) (_result *StartAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAppEnvResponse{}
	_body, _err := client.StartAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAppEnvWithOptions(request *StartAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &StartAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("StartAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAppEnv(request *StopAppEnvRequest) (_result *StopAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopAppEnvResponse{}
	_body, _err := client.StopAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppEnvWithOptions(request *StopAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &StopAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("StopAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TerminateAppEnv(request *TerminateAppEnvRequest) (_result *TerminateAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TerminateAppEnvResponse{}
	_body, _err := client.TerminateAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TerminateAppEnvWithOptions(request *TerminateAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TerminateAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &TerminateAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("TerminateAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv/terminate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppEnv(request *UpdateAppEnvRequest) (_result *UpdateAppEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAppEnvResponse{}
	_body, _err := client.UpdateAppEnvWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppEnvWithOptions(request *UpdateAppEnvRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAppEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvDescription)) {
		body["EnvDescription"] = request.EnvDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.PkgVersionId)) {
		body["PkgVersionId"] = request.PkgVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionSettings)) {
		body["OptionSettings"] = request.OptionSettings
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		body["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraProperties)) {
		body["ExtraProperties"] = request.ExtraProperties
	}

	if !tea.BoolValue(util.IsUnset(request.BatchSize)) {
		body["BatchSize"] = request.BatchSize
	}

	if !tea.BoolValue(util.IsUnset(request.BatchPercent)) {
		body["BatchPercent"] = request.BatchPercent
	}

	if !tea.BoolValue(util.IsUnset(request.BatchInterval)) {
		body["BatchInterval"] = request.BatchInterval
	}

	if !tea.BoolValue(util.IsUnset(request.PauseBetweenBatches)) {
		body["PauseBetweenBatches"] = request.PauseBetweenBatches
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateAppEnvResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("UpdateAppEnv"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v1/wam/appEnv"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplication(request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationWithOptions(request *UpdateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppDescription)) {
		body["AppDescription"] = request.AppDescription
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("UpdateApplication"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v1/wam/application"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (_result *UpdateConfigTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigTemplateResponse{}
	_body, _err := client.UpdateConfigTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigTemplateWithOptions(request *UpdateConfigTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConfigTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateDescription)) {
		body["TemplateDescription"] = request.TemplateDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionSettings)) {
		body["OptionSettings"] = request.OptionSettings
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateConfigTemplateResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("UpdateConfigTemplate"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v1/wam/configTemplate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateConfigSetting(request *ValidateConfigSettingRequest) (_result *ValidateConfigSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateConfigSettingResponse{}
	_body, _err := client.ValidateConfigSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateConfigSettingWithOptions(request *ValidateConfigSettingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateConfigSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		body["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionSettings)) {
		body["OptionSettings"] = request.OptionSettings
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ValidateConfigSettingResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ValidateConfigSetting"), tea.String("2019-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v1/wam/config/configSetting/validate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

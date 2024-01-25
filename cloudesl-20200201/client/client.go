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

type ActivateApDeviceRequest struct {
	ApMac       *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s ActivateApDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateApDeviceRequest) GoString() string {
	return s.String()
}

func (s *ActivateApDeviceRequest) SetApMac(v string) *ActivateApDeviceRequest {
	s.ApMac = &v
	return s
}

func (s *ActivateApDeviceRequest) SetExtraParams(v string) *ActivateApDeviceRequest {
	s.ExtraParams = &v
	return s
}

func (s *ActivateApDeviceRequest) SetStoreId(v string) *ActivateApDeviceRequest {
	s.StoreId = &v
	return s
}

type ActivateApDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateApDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateApDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateApDeviceResponseBody) SetCode(v string) *ActivateApDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetDynamicCode(v string) *ActivateApDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetDynamicMessage(v string) *ActivateApDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetErrorCode(v string) *ActivateApDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetErrorMessage(v string) *ActivateApDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetMessage(v string) *ActivateApDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetRequestId(v string) *ActivateApDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetSuccess(v bool) *ActivateApDeviceResponseBody {
	s.Success = &v
	return s
}

type ActivateApDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateApDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateApDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateApDeviceResponse) GoString() string {
	return s.String()
}

func (s *ActivateApDeviceResponse) SetHeaders(v map[string]*string) *ActivateApDeviceResponse {
	s.Headers = v
	return s
}

func (s *ActivateApDeviceResponse) SetStatusCode(v int32) *ActivateApDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateApDeviceResponse) SetBody(v *ActivateApDeviceResponseBody) *ActivateApDeviceResponse {
	s.Body = v
	return s
}

type AddApDeviceRequest struct {
	ApMac        *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ExtraParams  *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	StoreId      *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s AddApDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddApDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddApDeviceRequest) SetApMac(v string) *AddApDeviceRequest {
	s.ApMac = &v
	return s
}

func (s *AddApDeviceRequest) SetClientToken(v string) *AddApDeviceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddApDeviceRequest) SetExtraParams(v string) *AddApDeviceRequest {
	s.ExtraParams = &v
	return s
}

func (s *AddApDeviceRequest) SetRemark(v string) *AddApDeviceRequest {
	s.Remark = &v
	return s
}

func (s *AddApDeviceRequest) SetSerialNumber(v string) *AddApDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *AddApDeviceRequest) SetStoreId(v string) *AddApDeviceRequest {
	s.StoreId = &v
	return s
}

type AddApDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddApDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddApDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddApDeviceResponseBody) SetCode(v string) *AddApDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddApDeviceResponseBody) SetDynamicCode(v string) *AddApDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddApDeviceResponseBody) SetDynamicMessage(v string) *AddApDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddApDeviceResponseBody) SetErrorCode(v string) *AddApDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddApDeviceResponseBody) SetErrorMessage(v string) *AddApDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddApDeviceResponseBody) SetMessage(v string) *AddApDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *AddApDeviceResponseBody) SetRequestId(v string) *AddApDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddApDeviceResponseBody) SetSuccess(v bool) *AddApDeviceResponseBody {
	s.Success = &v
	return s
}

type AddApDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddApDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddApDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddApDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddApDeviceResponse) SetHeaders(v map[string]*string) *AddApDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddApDeviceResponse) SetStatusCode(v int32) *AddApDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddApDeviceResponse) SetBody(v *AddApDeviceResponseBody) *AddApDeviceResponse {
	s.Body = v
	return s
}

type AddCompanyTemplateRequest struct {
	DeviceType       *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EslSize          *string `json:"EslSize,omitempty" xml:"EslSize,omitempty"`
	ExtraParams      *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IfDefault        *bool   `json:"IfDefault,omitempty" xml:"IfDefault,omitempty"`
	IfMember         *bool   `json:"IfMember,omitempty" xml:"IfMember,omitempty"`
	IfOutOfInventory *bool   `json:"IfOutOfInventory,omitempty" xml:"IfOutOfInventory,omitempty"`
	IfPromotion      *bool   `json:"IfPromotion,omitempty" xml:"IfPromotion,omitempty"`
	IfSourceCode     *bool   `json:"IfSourceCode,omitempty" xml:"IfSourceCode,omitempty"`
	Layout           *int32  `json:"Layout,omitempty" xml:"Layout,omitempty"`
	Scene            *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateSceneId  *string `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
	TemplateType     *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateVersion  *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Vendor           *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s AddCompanyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCompanyTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddCompanyTemplateRequest) SetDeviceType(v string) *AddCompanyTemplateRequest {
	s.DeviceType = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetEslSize(v string) *AddCompanyTemplateRequest {
	s.EslSize = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetExtraParams(v string) *AddCompanyTemplateRequest {
	s.ExtraParams = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetGroupId(v string) *AddCompanyTemplateRequest {
	s.GroupId = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetIfDefault(v bool) *AddCompanyTemplateRequest {
	s.IfDefault = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetIfMember(v bool) *AddCompanyTemplateRequest {
	s.IfMember = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetIfOutOfInventory(v bool) *AddCompanyTemplateRequest {
	s.IfOutOfInventory = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetIfPromotion(v bool) *AddCompanyTemplateRequest {
	s.IfPromotion = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetIfSourceCode(v bool) *AddCompanyTemplateRequest {
	s.IfSourceCode = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetLayout(v int32) *AddCompanyTemplateRequest {
	s.Layout = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetScene(v string) *AddCompanyTemplateRequest {
	s.Scene = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetTemplateName(v string) *AddCompanyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetTemplateSceneId(v string) *AddCompanyTemplateRequest {
	s.TemplateSceneId = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetTemplateType(v string) *AddCompanyTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetTemplateVersion(v string) *AddCompanyTemplateRequest {
	s.TemplateVersion = &v
	return s
}

func (s *AddCompanyTemplateRequest) SetVendor(v string) *AddCompanyTemplateRequest {
	s.Vendor = &v
	return s
}

type AddCompanyTemplateResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCompanyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCompanyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddCompanyTemplateResponseBody) SetCode(v string) *AddCompanyTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetDynamicCode(v string) *AddCompanyTemplateResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetDynamicMessage(v string) *AddCompanyTemplateResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetErrorCode(v string) *AddCompanyTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetErrorMessage(v string) *AddCompanyTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetMessage(v string) *AddCompanyTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetRequestId(v string) *AddCompanyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCompanyTemplateResponseBody) SetSuccess(v bool) *AddCompanyTemplateResponseBody {
	s.Success = &v
	return s
}

type AddCompanyTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCompanyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCompanyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCompanyTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddCompanyTemplateResponse) SetHeaders(v map[string]*string) *AddCompanyTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddCompanyTemplateResponse) SetStatusCode(v int32) *AddCompanyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCompanyTemplateResponse) SetBody(v *AddCompanyTemplateResponseBody) *AddCompanyTemplateResponse {
	s.Body = v
	return s
}

type AddUserRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserRequest) GoString() string {
	return s.String()
}

func (s *AddUserRequest) SetClientToken(v string) *AddUserRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserRequest) SetExtraParams(v string) *AddUserRequest {
	s.ExtraParams = &v
	return s
}

func (s *AddUserRequest) SetUserId(v string) *AddUserRequest {
	s.UserId = &v
	return s
}

type AddUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserResponseBody) SetCode(v string) *AddUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddUserResponseBody) SetDynamicCode(v string) *AddUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddUserResponseBody) SetDynamicMessage(v string) *AddUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddUserResponseBody) SetErrorCode(v string) *AddUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddUserResponseBody) SetErrorMessage(v string) *AddUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddUserResponseBody) SetMessage(v string) *AddUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddUserResponseBody) SetRequestId(v string) *AddUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserResponseBody) SetSuccess(v bool) *AddUserResponseBody {
	s.Success = &v
	return s
}

type AddUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponse) GoString() string {
	return s.String()
}

func (s *AddUserResponse) SetHeaders(v map[string]*string) *AddUserResponse {
	s.Headers = v
	return s
}

func (s *AddUserResponse) SetStatusCode(v int32) *AddUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserResponse) SetBody(v *AddUserResponseBody) *AddUserResponse {
	s.Body = v
	return s
}

type ApplyCompanyTemplateVersionToStoresRequest struct {
	Stores          *string `json:"Stores,omitempty" xml:"Stores,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ApplyCompanyTemplateVersionToStoresRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCompanyTemplateVersionToStoresRequest) GoString() string {
	return s.String()
}

func (s *ApplyCompanyTemplateVersionToStoresRequest) SetStores(v string) *ApplyCompanyTemplateVersionToStoresRequest {
	s.Stores = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresRequest) SetTemplateVersion(v string) *ApplyCompanyTemplateVersionToStoresRequest {
	s.TemplateVersion = &v
	return s
}

type ApplyCompanyTemplateVersionToStoresResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyCompanyTemplateVersionToStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyCompanyTemplateVersionToStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetCode(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetDynamicCode(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetDynamicMessage(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetErrorCode(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetErrorMessage(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetMessage(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetRequestId(v string) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponseBody) SetSuccess(v bool) *ApplyCompanyTemplateVersionToStoresResponseBody {
	s.Success = &v
	return s
}

type ApplyCompanyTemplateVersionToStoresResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCompanyTemplateVersionToStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCompanyTemplateVersionToStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyCompanyTemplateVersionToStoresResponse) GoString() string {
	return s.String()
}

func (s *ApplyCompanyTemplateVersionToStoresResponse) SetHeaders(v map[string]*string) *ApplyCompanyTemplateVersionToStoresResponse {
	s.Headers = v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponse) SetStatusCode(v int32) *ApplyCompanyTemplateVersionToStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCompanyTemplateVersionToStoresResponse) SetBody(v *ApplyCompanyTemplateVersionToStoresResponseBody) *ApplyCompanyTemplateVersionToStoresResponse {
	s.Body = v
	return s
}

type AssignUserRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	Stores      *string `json:"Stores,omitempty" xml:"Stores,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType    *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AssignUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignUserRequest) GoString() string {
	return s.String()
}

func (s *AssignUserRequest) SetExtraParams(v string) *AssignUserRequest {
	s.ExtraParams = &v
	return s
}

func (s *AssignUserRequest) SetStores(v string) *AssignUserRequest {
	s.Stores = &v
	return s
}

func (s *AssignUserRequest) SetUserId(v string) *AssignUserRequest {
	s.UserId = &v
	return s
}

func (s *AssignUserRequest) SetUserType(v string) *AssignUserRequest {
	s.UserType = &v
	return s
}

type AssignUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignUserResponseBody) GoString() string {
	return s.String()
}

func (s *AssignUserResponseBody) SetCode(v string) *AssignUserResponseBody {
	s.Code = &v
	return s
}

func (s *AssignUserResponseBody) SetDynamicCode(v string) *AssignUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AssignUserResponseBody) SetDynamicMessage(v string) *AssignUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AssignUserResponseBody) SetErrorCode(v string) *AssignUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AssignUserResponseBody) SetErrorMessage(v string) *AssignUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AssignUserResponseBody) SetMessage(v string) *AssignUserResponseBody {
	s.Message = &v
	return s
}

func (s *AssignUserResponseBody) SetRequestId(v string) *AssignUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignUserResponseBody) SetSuccess(v bool) *AssignUserResponseBody {
	s.Success = &v
	return s
}

type AssignUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignUserResponse) GoString() string {
	return s.String()
}

func (s *AssignUserResponse) SetHeaders(v map[string]*string) *AssignUserResponse {
	s.Headers = v
	return s
}

func (s *AssignUserResponse) SetStatusCode(v int32) *AssignUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignUserResponse) SetBody(v *AssignUserResponseBody) *AssignUserResponse {
	s.Body = v
	return s
}

type BatchInsertItemsRequest struct {
	ExtraParams  *string                            `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	ItemInfo     []*BatchInsertItemsRequestItemInfo `json:"ItemInfo,omitempty" xml:"ItemInfo,omitempty" type:"Repeated"`
	StoreId      *string                            `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	SyncByItemId *bool                              `json:"SyncByItemId,omitempty" xml:"SyncByItemId,omitempty"`
}

func (s BatchInsertItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsRequest) SetExtraParams(v string) *BatchInsertItemsRequest {
	s.ExtraParams = &v
	return s
}

func (s *BatchInsertItemsRequest) SetItemInfo(v []*BatchInsertItemsRequestItemInfo) *BatchInsertItemsRequest {
	s.ItemInfo = v
	return s
}

func (s *BatchInsertItemsRequest) SetStoreId(v string) *BatchInsertItemsRequest {
	s.StoreId = &v
	return s
}

func (s *BatchInsertItemsRequest) SetSyncByItemId(v bool) *BatchInsertItemsRequest {
	s.SyncByItemId = &v
	return s
}

type BatchInsertItemsRequestItemInfo struct {
	ActionPrice       *int32  `json:"ActionPrice,omitempty" xml:"ActionPrice,omitempty"`
	BeClearance       *bool   `json:"BeClearance,omitempty" xml:"BeClearance,omitempty"`
	BeMember          *bool   `json:"BeMember,omitempty" xml:"BeMember,omitempty"`
	BePromotion       *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	BeSourceCode      *bool   `json:"BeSourceCode,omitempty" xml:"BeSourceCode,omitempty"`
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	CategoryName      *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CustomizeFeatureA *string `json:"CustomizeFeatureA,omitempty" xml:"CustomizeFeatureA,omitempty"`
	CustomizeFeatureB *string `json:"CustomizeFeatureB,omitempty" xml:"CustomizeFeatureB,omitempty"`
	CustomizeFeatureC *string `json:"CustomizeFeatureC,omitempty" xml:"CustomizeFeatureC,omitempty"`
	CustomizeFeatureD *string `json:"CustomizeFeatureD,omitempty" xml:"CustomizeFeatureD,omitempty"`
	CustomizeFeatureE *string `json:"CustomizeFeatureE,omitempty" xml:"CustomizeFeatureE,omitempty"`
	CustomizeFeatureF *string `json:"CustomizeFeatureF,omitempty" xml:"CustomizeFeatureF,omitempty"`
	CustomizeFeatureG *string `json:"CustomizeFeatureG,omitempty" xml:"CustomizeFeatureG,omitempty"`
	CustomizeFeatureH *string `json:"CustomizeFeatureH,omitempty" xml:"CustomizeFeatureH,omitempty"`
	CustomizeFeatureI *string `json:"CustomizeFeatureI,omitempty" xml:"CustomizeFeatureI,omitempty"`
	CustomizeFeatureJ *string `json:"CustomizeFeatureJ,omitempty" xml:"CustomizeFeatureJ,omitempty"`
	CustomizeFeatureK *string `json:"CustomizeFeatureK,omitempty" xml:"CustomizeFeatureK,omitempty"`
	CustomizeFeatureL *string `json:"CustomizeFeatureL,omitempty" xml:"CustomizeFeatureL,omitempty"`
	CustomizeFeatureM *string `json:"CustomizeFeatureM,omitempty" xml:"CustomizeFeatureM,omitempty"`
	CustomizeFeatureN *string `json:"CustomizeFeatureN,omitempty" xml:"CustomizeFeatureN,omitempty"`
	CustomizeFeatureO *string `json:"CustomizeFeatureO,omitempty" xml:"CustomizeFeatureO,omitempty"`
	CustomizeFeatureP *string `json:"CustomizeFeatureP,omitempty" xml:"CustomizeFeatureP,omitempty"`
	CustomizeFeatureQ *string `json:"CustomizeFeatureQ,omitempty" xml:"CustomizeFeatureQ,omitempty"`
	CustomizeFeatureR *string `json:"CustomizeFeatureR,omitempty" xml:"CustomizeFeatureR,omitempty"`
	CustomizeFeatureS *string `json:"CustomizeFeatureS,omitempty" xml:"CustomizeFeatureS,omitempty"`
	CustomizeFeatureT *string `json:"CustomizeFeatureT,omitempty" xml:"CustomizeFeatureT,omitempty"`
	CustomizeFeatureU *string `json:"CustomizeFeatureU,omitempty" xml:"CustomizeFeatureU,omitempty"`
	CustomizeFeatureV *string `json:"CustomizeFeatureV,omitempty" xml:"CustomizeFeatureV,omitempty"`
	CustomizeFeatureW *string `json:"CustomizeFeatureW,omitempty" xml:"CustomizeFeatureW,omitempty"`
	CustomizeFeatureX *string `json:"CustomizeFeatureX,omitempty" xml:"CustomizeFeatureX,omitempty"`
	CustomizeFeatureY *string `json:"CustomizeFeatureY,omitempty" xml:"CustomizeFeatureY,omitempty"`
	CustomizeFeatureZ *string `json:"CustomizeFeatureZ,omitempty" xml:"CustomizeFeatureZ,omitempty"`
	EnergyEfficiency  *string `json:"EnergyEfficiency,omitempty" xml:"EnergyEfficiency,omitempty"`
	ForestFirstId     *string `json:"ForestFirstId,omitempty" xml:"ForestFirstId,omitempty"`
	ForestSecondId    *string `json:"ForestSecondId,omitempty" xml:"ForestSecondId,omitempty"`
	InventoryStatus   *string `json:"InventoryStatus,omitempty" xml:"InventoryStatus,omitempty"`
	ItemBarCode       *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemInfoIndex     *int32  `json:"ItemInfoIndex,omitempty" xml:"ItemInfoIndex,omitempty"`
	ItemPicUrl        *string `json:"ItemPicUrl,omitempty" xml:"ItemPicUrl,omitempty"`
	ItemQrCode        *string `json:"ItemQrCode,omitempty" xml:"ItemQrCode,omitempty"`
	ItemShortTitle    *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	ItemTitle         *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	Manufacturer      *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	Material          *string `json:"Material,omitempty" xml:"Material,omitempty"`
	MemberPrice       *int32  `json:"MemberPrice,omitempty" xml:"MemberPrice,omitempty"`
	ModelNumber       *string `json:"ModelNumber,omitempty" xml:"ModelNumber,omitempty"`
	OriginalPrice     *int32  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceUnit         *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	ProductionPlace   *string `json:"ProductionPlace,omitempty" xml:"ProductionPlace,omitempty"`
	PromotionEnd      *string `json:"PromotionEnd,omitempty" xml:"PromotionEnd,omitempty"`
	PromotionReason   *string `json:"PromotionReason,omitempty" xml:"PromotionReason,omitempty"`
	PromotionStart    *string `json:"PromotionStart,omitempty" xml:"PromotionStart,omitempty"`
	PromotionText     *string `json:"PromotionText,omitempty" xml:"PromotionText,omitempty"`
	Rank              *string `json:"Rank,omitempty" xml:"Rank,omitempty"`
	SaleSpec          *string `json:"SaleSpec,omitempty" xml:"SaleSpec,omitempty"`
	SalesPrice        *int32  `json:"SalesPrice,omitempty" xml:"SalesPrice,omitempty"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SourceCode        *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SuggestPrice      *int32  `json:"SuggestPrice,omitempty" xml:"SuggestPrice,omitempty"`
	SupplierName      *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	TaxFee            *string `json:"TaxFee,omitempty" xml:"TaxFee,omitempty"`
	TemplateSceneId   *string `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
}

func (s BatchInsertItemsRequestItemInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsRequestItemInfo) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsRequestItemInfo) SetActionPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.ActionPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBeClearance(v bool) *BatchInsertItemsRequestItemInfo {
	s.BeClearance = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBeMember(v bool) *BatchInsertItemsRequestItemInfo {
	s.BeMember = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBePromotion(v bool) *BatchInsertItemsRequestItemInfo {
	s.BePromotion = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBeSourceCode(v bool) *BatchInsertItemsRequestItemInfo {
	s.BeSourceCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBrandName(v string) *BatchInsertItemsRequestItemInfo {
	s.BrandName = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCategoryName(v string) *BatchInsertItemsRequestItemInfo {
	s.CategoryName = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureA(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureA = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureB(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureB = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureC(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureC = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureD(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureD = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureE(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureE = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureF(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureF = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureG(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureG = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureH(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureH = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureI(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureI = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureJ(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureJ = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureK(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureK = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureL(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureL = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureM(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureM = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureN(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureN = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureO(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureO = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureP(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureP = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureQ(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureQ = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureR(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureR = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureS(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureS = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureT(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureT = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureU(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureU = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureV(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureV = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureW(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureW = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureX(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureX = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureY(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureY = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureZ(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureZ = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetEnergyEfficiency(v string) *BatchInsertItemsRequestItemInfo {
	s.EnergyEfficiency = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetForestFirstId(v string) *BatchInsertItemsRequestItemInfo {
	s.ForestFirstId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetForestSecondId(v string) *BatchInsertItemsRequestItemInfo {
	s.ForestSecondId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetInventoryStatus(v string) *BatchInsertItemsRequestItemInfo {
	s.InventoryStatus = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemBarCode(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemBarCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemId(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemInfoIndex(v int32) *BatchInsertItemsRequestItemInfo {
	s.ItemInfoIndex = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemPicUrl(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemPicUrl = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemQrCode(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemQrCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemShortTitle(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemShortTitle = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemTitle(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemTitle = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetManufacturer(v string) *BatchInsertItemsRequestItemInfo {
	s.Manufacturer = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetMaterial(v string) *BatchInsertItemsRequestItemInfo {
	s.Material = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetMemberPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.MemberPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetModelNumber(v string) *BatchInsertItemsRequestItemInfo {
	s.ModelNumber = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetOriginalPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.OriginalPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPriceUnit(v string) *BatchInsertItemsRequestItemInfo {
	s.PriceUnit = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetProductionPlace(v string) *BatchInsertItemsRequestItemInfo {
	s.ProductionPlace = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionEnd(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionEnd = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionReason(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionReason = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionStart(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionStart = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionText(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionText = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetRank(v string) *BatchInsertItemsRequestItemInfo {
	s.Rank = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSaleSpec(v string) *BatchInsertItemsRequestItemInfo {
	s.SaleSpec = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSalesPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.SalesPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSkuId(v string) *BatchInsertItemsRequestItemInfo {
	s.SkuId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSourceCode(v string) *BatchInsertItemsRequestItemInfo {
	s.SourceCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSuggestPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.SuggestPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSupplierName(v string) *BatchInsertItemsRequestItemInfo {
	s.SupplierName = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetTaxFee(v string) *BatchInsertItemsRequestItemInfo {
	s.TaxFee = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetTemplateSceneId(v string) *BatchInsertItemsRequestItemInfo {
	s.TemplateSceneId = &v
	return s
}

type BatchInsertItemsResponseBody struct {
	BatchResults   []*BatchInsertItemsResponseBodyBatchResults `json:"BatchResults,omitempty" xml:"BatchResults,omitempty" type:"Repeated"`
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                     `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchInsertItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponseBody) SetBatchResults(v []*BatchInsertItemsResponseBodyBatchResults) *BatchInsertItemsResponseBody {
	s.BatchResults = v
	return s
}

func (s *BatchInsertItemsResponseBody) SetCode(v string) *BatchInsertItemsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetDynamicCode(v string) *BatchInsertItemsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetDynamicMessage(v string) *BatchInsertItemsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetErrorCode(v string) *BatchInsertItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetErrorMessage(v string) *BatchInsertItemsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetMessage(v string) *BatchInsertItemsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetRequestId(v string) *BatchInsertItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetSuccess(v bool) *BatchInsertItemsResponseBody {
	s.Success = &v
	return s
}

type BatchInsertItemsResponseBodyBatchResults struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchInsertItemsResponseBodyBatchResults) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponseBodyBatchResults) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponseBodyBatchResults) SetErrorCode(v string) *BatchInsertItemsResponseBodyBatchResults {
	s.ErrorCode = &v
	return s
}

func (s *BatchInsertItemsResponseBodyBatchResults) SetIndex(v int32) *BatchInsertItemsResponseBodyBatchResults {
	s.Index = &v
	return s
}

func (s *BatchInsertItemsResponseBodyBatchResults) SetMessage(v string) *BatchInsertItemsResponseBodyBatchResults {
	s.Message = &v
	return s
}

func (s *BatchInsertItemsResponseBodyBatchResults) SetSuccess(v bool) *BatchInsertItemsResponseBodyBatchResults {
	s.Success = &v
	return s
}

type BatchInsertItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchInsertItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchInsertItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponse) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponse) SetHeaders(v map[string]*string) *BatchInsertItemsResponse {
	s.Headers = v
	return s
}

func (s *BatchInsertItemsResponse) SetStatusCode(v int32) *BatchInsertItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchInsertItemsResponse) SetBody(v *BatchInsertItemsResponseBody) *BatchInsertItemsResponse {
	s.Body = v
	return s
}

type BindEslDeviceRequest struct {
	Column        *string `json:"Column,omitempty" xml:"Column,omitempty"`
	ContainerId   *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	EslBarCode    *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ExtraParams   *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	ItemBarCode   *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	Layer         *int32  `json:"Layer,omitempty" xml:"Layer,omitempty"`
	LayoutId      *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	LayoutName    *string `json:"LayoutName,omitempty" xml:"LayoutName,omitempty"`
	Shelf         *string `json:"Shelf,omitempty" xml:"Shelf,omitempty"`
	StoreId       *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s BindEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindEslDeviceRequest) SetColumn(v string) *BindEslDeviceRequest {
	s.Column = &v
	return s
}

func (s *BindEslDeviceRequest) SetContainerId(v string) *BindEslDeviceRequest {
	s.ContainerId = &v
	return s
}

func (s *BindEslDeviceRequest) SetContainerName(v string) *BindEslDeviceRequest {
	s.ContainerName = &v
	return s
}

func (s *BindEslDeviceRequest) SetEslBarCode(v string) *BindEslDeviceRequest {
	s.EslBarCode = &v
	return s
}

func (s *BindEslDeviceRequest) SetExtraParams(v string) *BindEslDeviceRequest {
	s.ExtraParams = &v
	return s
}

func (s *BindEslDeviceRequest) SetItemBarCode(v string) *BindEslDeviceRequest {
	s.ItemBarCode = &v
	return s
}

func (s *BindEslDeviceRequest) SetLayer(v int32) *BindEslDeviceRequest {
	s.Layer = &v
	return s
}

func (s *BindEslDeviceRequest) SetLayoutId(v string) *BindEslDeviceRequest {
	s.LayoutId = &v
	return s
}

func (s *BindEslDeviceRequest) SetLayoutName(v string) *BindEslDeviceRequest {
	s.LayoutName = &v
	return s
}

func (s *BindEslDeviceRequest) SetShelf(v string) *BindEslDeviceRequest {
	s.Shelf = &v
	return s
}

func (s *BindEslDeviceRequest) SetStoreId(v string) *BindEslDeviceRequest {
	s.StoreId = &v
	return s
}

type BindEslDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindEslDeviceResponseBody) SetCode(v string) *BindEslDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetDynamicCode(v string) *BindEslDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetDynamicMessage(v string) *BindEslDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetErrorCode(v string) *BindEslDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetErrorMessage(v string) *BindEslDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetMessage(v string) *BindEslDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetRequestId(v string) *BindEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetSuccess(v bool) *BindEslDeviceResponseBody {
	s.Success = &v
	return s
}

type BindEslDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindEslDeviceResponse) SetHeaders(v map[string]*string) *BindEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindEslDeviceResponse) SetStatusCode(v int32) *BindEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEslDeviceResponse) SetBody(v *BindEslDeviceResponseBody) *BindEslDeviceResponse {
	s.Body = v
	return s
}

type CreateStoreRequest struct {
	AutoUnbindDays       *int32  `json:"AutoUnbindDays,omitempty" xml:"AutoUnbindDays,omitempty"`
	AutoUnbindOfflineEsl *bool   `json:"AutoUnbindOfflineEsl,omitempty" xml:"AutoUnbindOfflineEsl,omitempty"`
	BarCodeEncode        *int32  `json:"BarCodeEncode,omitempty" xml:"BarCodeEncode,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ExtraParams          *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	ParentId             *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Phone                *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	StoreName            *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	TimeZone             *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UserStoreCode        *string `json:"UserStoreCode,omitempty" xml:"UserStoreCode,omitempty"`
}

func (s CreateStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateStoreRequest) SetAutoUnbindDays(v int32) *CreateStoreRequest {
	s.AutoUnbindDays = &v
	return s
}

func (s *CreateStoreRequest) SetAutoUnbindOfflineEsl(v bool) *CreateStoreRequest {
	s.AutoUnbindOfflineEsl = &v
	return s
}

func (s *CreateStoreRequest) SetBarCodeEncode(v int32) *CreateStoreRequest {
	s.BarCodeEncode = &v
	return s
}

func (s *CreateStoreRequest) SetClientToken(v string) *CreateStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStoreRequest) SetExtraParams(v string) *CreateStoreRequest {
	s.ExtraParams = &v
	return s
}

func (s *CreateStoreRequest) SetParentId(v string) *CreateStoreRequest {
	s.ParentId = &v
	return s
}

func (s *CreateStoreRequest) SetPhone(v string) *CreateStoreRequest {
	s.Phone = &v
	return s
}

func (s *CreateStoreRequest) SetStoreName(v string) *CreateStoreRequest {
	s.StoreName = &v
	return s
}

func (s *CreateStoreRequest) SetTimeZone(v string) *CreateStoreRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateStoreRequest) SetUserStoreCode(v string) *CreateStoreRequest {
	s.UserStoreCode = &v
	return s
}

type CreateStoreResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StoreId        *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoreResponseBody) SetCode(v string) *CreateStoreResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStoreResponseBody) SetDynamicCode(v string) *CreateStoreResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateStoreResponseBody) SetDynamicMessage(v string) *CreateStoreResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateStoreResponseBody) SetErrorCode(v string) *CreateStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateStoreResponseBody) SetErrorMessage(v string) *CreateStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStoreResponseBody) SetMessage(v string) *CreateStoreResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStoreResponseBody) SetRequestId(v string) *CreateStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStoreResponseBody) SetStoreId(v string) *CreateStoreResponseBody {
	s.StoreId = &v
	return s
}

func (s *CreateStoreResponseBody) SetSuccess(v bool) *CreateStoreResponseBody {
	s.Success = &v
	return s
}

type CreateStoreResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateStoreResponse) SetHeaders(v map[string]*string) *CreateStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateStoreResponse) SetStatusCode(v int32) *CreateStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoreResponse) SetBody(v *CreateStoreResponseBody) *CreateStoreResponse {
	s.Body = v
	return s
}

type DeleteApDeviceRequest struct {
	ApMac       *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteApDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteApDeviceRequest) SetApMac(v string) *DeleteApDeviceRequest {
	s.ApMac = &v
	return s
}

func (s *DeleteApDeviceRequest) SetExtraParams(v string) *DeleteApDeviceRequest {
	s.ExtraParams = &v
	return s
}

func (s *DeleteApDeviceRequest) SetStoreId(v string) *DeleteApDeviceRequest {
	s.StoreId = &v
	return s
}

type DeleteApDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApDeviceResponseBody) SetCode(v string) *DeleteApDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetDynamicCode(v string) *DeleteApDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetDynamicMessage(v string) *DeleteApDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetErrorCode(v string) *DeleteApDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetErrorMessage(v string) *DeleteApDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetMessage(v string) *DeleteApDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetRequestId(v string) *DeleteApDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetSuccess(v bool) *DeleteApDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteApDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteApDeviceResponse) SetHeaders(v map[string]*string) *DeleteApDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteApDeviceResponse) SetStatusCode(v int32) *DeleteApDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApDeviceResponse) SetBody(v *DeleteApDeviceResponseBody) *DeleteApDeviceResponse {
	s.Body = v
	return s
}

type DeleteCompanyTemplateRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteCompanyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompanyTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCompanyTemplateRequest) SetExtraParams(v string) *DeleteCompanyTemplateRequest {
	s.ExtraParams = &v
	return s
}

func (s *DeleteCompanyTemplateRequest) SetTemplateId(v string) *DeleteCompanyTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteCompanyTemplateResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCompanyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompanyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCompanyTemplateResponseBody) SetCode(v string) *DeleteCompanyTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetDynamicCode(v string) *DeleteCompanyTemplateResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetDynamicMessage(v string) *DeleteCompanyTemplateResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetErrorCode(v string) *DeleteCompanyTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetErrorMessage(v string) *DeleteCompanyTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetMessage(v string) *DeleteCompanyTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetRequestId(v string) *DeleteCompanyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCompanyTemplateResponseBody) SetSuccess(v bool) *DeleteCompanyTemplateResponseBody {
	s.Success = &v
	return s
}

type DeleteCompanyTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCompanyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCompanyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompanyTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCompanyTemplateResponse) SetHeaders(v map[string]*string) *DeleteCompanyTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCompanyTemplateResponse) SetStatusCode(v int32) *DeleteCompanyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCompanyTemplateResponse) SetBody(v *DeleteCompanyTemplateResponseBody) *DeleteCompanyTemplateResponse {
	s.Body = v
	return s
}

type DeleteItemRequest struct {
	ItemBarCode     *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	StoreId         *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	UnbindEslDevice *bool   `json:"UnbindEslDevice,omitempty" xml:"UnbindEslDevice,omitempty"`
}

func (s DeleteItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteItemRequest) SetItemBarCode(v string) *DeleteItemRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DeleteItemRequest) SetStoreId(v string) *DeleteItemRequest {
	s.StoreId = &v
	return s
}

func (s *DeleteItemRequest) SetUnbindEslDevice(v bool) *DeleteItemRequest {
	s.UnbindEslDevice = &v
	return s
}

type DeleteItemResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteItemResponseBody) SetCode(v string) *DeleteItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteItemResponseBody) SetDynamicCode(v string) *DeleteItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteItemResponseBody) SetDynamicMessage(v string) *DeleteItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteItemResponseBody) SetErrorCode(v string) *DeleteItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteItemResponseBody) SetErrorMessage(v string) *DeleteItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteItemResponseBody) SetMessage(v string) *DeleteItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteItemResponseBody) SetRequestId(v string) *DeleteItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteItemResponseBody) SetSuccess(v bool) *DeleteItemResponseBody {
	s.Success = &v
	return s
}

type DeleteItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteItemResponse) SetHeaders(v map[string]*string) *DeleteItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteItemResponse) SetStatusCode(v int32) *DeleteItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteItemResponse) SetBody(v *DeleteItemResponseBody) *DeleteItemResponse {
	s.Body = v
	return s
}

type DeleteStoreRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteStoreRequest) SetExtraParams(v string) *DeleteStoreRequest {
	s.ExtraParams = &v
	return s
}

func (s *DeleteStoreRequest) SetStoreId(v string) *DeleteStoreRequest {
	s.StoreId = &v
	return s
}

type DeleteStoreResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStoreResponseBody) SetCode(v string) *DeleteStoreResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStoreResponseBody) SetDynamicCode(v string) *DeleteStoreResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteStoreResponseBody) SetDynamicMessage(v string) *DeleteStoreResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteStoreResponseBody) SetErrorCode(v string) *DeleteStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteStoreResponseBody) SetErrorMessage(v string) *DeleteStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteStoreResponseBody) SetMessage(v string) *DeleteStoreResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStoreResponseBody) SetRequestId(v string) *DeleteStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStoreResponseBody) SetSuccess(v bool) *DeleteStoreResponseBody {
	s.Success = &v
	return s
}

type DeleteStoreResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteStoreResponse) SetHeaders(v map[string]*string) *DeleteStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteStoreResponse) SetStatusCode(v int32) *DeleteStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStoreResponse) SetBody(v *DeleteStoreResponseBody) *DeleteStoreResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetExtraParams(v string) *DeleteUserRequest {
	s.ExtraParams = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetDynamicCode(v string) *DeleteUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetDynamicMessage(v string) *DeleteUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorCode(v string) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorMessage(v string) *DeleteUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DescribeApDevicesRequest struct {
	ApMac       *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	BeActivate  *bool   `json:"BeActivate,omitempty" xml:"BeActivate,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status      *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeApDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesRequest) SetApMac(v string) *DescribeApDevicesRequest {
	s.ApMac = &v
	return s
}

func (s *DescribeApDevicesRequest) SetBeActivate(v bool) *DescribeApDevicesRequest {
	s.BeActivate = &v
	return s
}

func (s *DescribeApDevicesRequest) SetExtraParams(v string) *DescribeApDevicesRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeApDevicesRequest) SetModel(v string) *DescribeApDevicesRequest {
	s.Model = &v
	return s
}

func (s *DescribeApDevicesRequest) SetPageNumber(v int32) *DescribeApDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApDevicesRequest) SetPageSize(v int32) *DescribeApDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApDevicesRequest) SetStatus(v bool) *DescribeApDevicesRequest {
	s.Status = &v
	return s
}

func (s *DescribeApDevicesRequest) SetStoreId(v string) *DescribeApDevicesRequest {
	s.StoreId = &v
	return s
}

type DescribeApDevicesResponseBody struct {
	ApDevices      []*DescribeApDevicesResponseBodyApDevices `json:"ApDevices,omitempty" xml:"ApDevices,omitempty" type:"Repeated"`
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponseBody) SetApDevices(v []*DescribeApDevicesResponseBodyApDevices) *DescribeApDevicesResponseBody {
	s.ApDevices = v
	return s
}

func (s *DescribeApDevicesResponseBody) SetCode(v string) *DescribeApDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetDynamicCode(v string) *DescribeApDevicesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetDynamicMessage(v string) *DescribeApDevicesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetErrorCode(v string) *DescribeApDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetErrorMessage(v string) *DescribeApDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetMessage(v string) *DescribeApDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetPageNumber(v int32) *DescribeApDevicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetPageSize(v int32) *DescribeApDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetRequestId(v string) *DescribeApDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetSuccess(v bool) *DescribeApDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetTotalCount(v int32) *DescribeApDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApDevicesResponseBodyApDevices struct {
	BeActivate *bool   `json:"BeActivate,omitempty" xml:"BeActivate,omitempty"`
	Mac        *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status     *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	StoreId    *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeApDevicesResponseBodyApDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponseBodyApDevices) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponseBodyApDevices) SetBeActivate(v bool) *DescribeApDevicesResponseBodyApDevices {
	s.BeActivate = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevices) SetMac(v string) *DescribeApDevicesResponseBodyApDevices {
	s.Mac = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevices) SetModel(v string) *DescribeApDevicesResponseBodyApDevices {
	s.Model = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevices) SetRemark(v string) *DescribeApDevicesResponseBodyApDevices {
	s.Remark = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevices) SetStatus(v bool) *DescribeApDevicesResponseBodyApDevices {
	s.Status = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevices) SetStoreId(v string) *DescribeApDevicesResponseBodyApDevices {
	s.StoreId = &v
	return s
}

type DescribeApDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponse) SetHeaders(v map[string]*string) *DescribeApDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApDevicesResponse) SetStatusCode(v int32) *DescribeApDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApDevicesResponse) SetBody(v *DescribeApDevicesResponseBody) *DescribeApDevicesResponse {
	s.Body = v
	return s
}

type DescribeAvailableEslModelsRequest struct {
	ModelId    *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAvailableEslModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableEslModelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEslModelsRequest) SetModelId(v string) *DescribeAvailableEslModelsRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeAvailableEslModelsRequest) SetName(v string) *DescribeAvailableEslModelsRequest {
	s.Name = &v
	return s
}

func (s *DescribeAvailableEslModelsRequest) SetPageNumber(v int32) *DescribeAvailableEslModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableEslModelsRequest) SetPageSize(v int32) *DescribeAvailableEslModelsRequest {
	s.PageSize = &v
	return s
}

type DescribeAvailableEslModelsResponseBody struct {
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EslModels      []*DescribeAvailableEslModelsResponseBodyEslModels `json:"EslModels,omitempty" xml:"EslModels,omitempty" type:"Repeated"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAvailableEslModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableEslModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEslModelsResponseBody) SetCode(v string) *DescribeAvailableEslModelsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetDynamicCode(v string) *DescribeAvailableEslModelsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetDynamicMessage(v string) *DescribeAvailableEslModelsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetErrorCode(v string) *DescribeAvailableEslModelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetErrorMessage(v string) *DescribeAvailableEslModelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetEslModels(v []*DescribeAvailableEslModelsResponseBodyEslModels) *DescribeAvailableEslModelsResponseBody {
	s.EslModels = v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetMessage(v string) *DescribeAvailableEslModelsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetPageNumber(v int32) *DescribeAvailableEslModelsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetPageSize(v int32) *DescribeAvailableEslModelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetRequestId(v string) *DescribeAvailableEslModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetSuccess(v bool) *DescribeAvailableEslModelsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBody) SetTotalCount(v int32) *DescribeAvailableEslModelsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAvailableEslModelsResponseBodyEslModels struct {
	DeviceType   *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EslSize      *string `json:"EslSize,omitempty" xml:"EslSize,omitempty"`
	ModelId      *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScreenHeight *int32  `json:"ScreenHeight,omitempty" xml:"ScreenHeight,omitempty"`
	ScreenWidth  *int32  `json:"ScreenWidth,omitempty" xml:"ScreenWidth,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeAvailableEslModelsResponseBodyEslModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableEslModelsResponseBodyEslModels) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetDeviceType(v string) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.DeviceType = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetEslSize(v string) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.EslSize = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetModelId(v string) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.ModelId = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetName(v string) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.Name = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetScreenHeight(v int32) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.ScreenHeight = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetScreenWidth(v int32) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.ScreenWidth = &v
	return s
}

func (s *DescribeAvailableEslModelsResponseBodyEslModels) SetVendor(v string) *DescribeAvailableEslModelsResponseBodyEslModels {
	s.Vendor = &v
	return s
}

type DescribeAvailableEslModelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableEslModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableEslModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableEslModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEslModelsResponse) SetHeaders(v map[string]*string) *DescribeAvailableEslModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableEslModelsResponse) SetStatusCode(v int32) *DescribeAvailableEslModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableEslModelsResponse) SetBody(v *DescribeAvailableEslModelsResponseBody) *DescribeAvailableEslModelsResponse {
	s.Body = v
	return s
}

type DescribeBindersRequest struct {
	EslBarCode  *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemTitle   *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeBindersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindersRequest) GoString() string {
	return s.String()
}

func (s *DescribeBindersRequest) SetEslBarCode(v string) *DescribeBindersRequest {
	s.EslBarCode = &v
	return s
}

func (s *DescribeBindersRequest) SetExtraParams(v string) *DescribeBindersRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeBindersRequest) SetItemBarCode(v string) *DescribeBindersRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeBindersRequest) SetItemTitle(v string) *DescribeBindersRequest {
	s.ItemTitle = &v
	return s
}

func (s *DescribeBindersRequest) SetPageNumber(v int32) *DescribeBindersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBindersRequest) SetPageSize(v int32) *DescribeBindersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBindersRequest) SetStoreId(v string) *DescribeBindersRequest {
	s.StoreId = &v
	return s
}

type DescribeBindersResponseBody struct {
	Code             *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode      *string                                        `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage   *string                                        `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode        *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage     *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EslItemBindInfos []*DescribeBindersResponseBodyEslItemBindInfos `json:"EslItemBindInfos,omitempty" xml:"EslItemBindInfos,omitempty" type:"Repeated"`
	Message          *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber       *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount       *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBindersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBindersResponseBody) SetCode(v string) *DescribeBindersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBindersResponseBody) SetDynamicCode(v string) *DescribeBindersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeBindersResponseBody) SetDynamicMessage(v string) *DescribeBindersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeBindersResponseBody) SetErrorCode(v string) *DescribeBindersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeBindersResponseBody) SetErrorMessage(v string) *DescribeBindersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeBindersResponseBody) SetEslItemBindInfos(v []*DescribeBindersResponseBodyEslItemBindInfos) *DescribeBindersResponseBody {
	s.EslItemBindInfos = v
	return s
}

func (s *DescribeBindersResponseBody) SetMessage(v string) *DescribeBindersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBindersResponseBody) SetPageNumber(v int32) *DescribeBindersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBindersResponseBody) SetPageSize(v int32) *DescribeBindersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBindersResponseBody) SetRequestId(v string) *DescribeBindersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBindersResponseBody) SetSuccess(v bool) *DescribeBindersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBindersResponseBody) SetTotalCount(v int32) *DescribeBindersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBindersResponseBodyEslItemBindInfos struct {
	ActionPrice    *string `json:"ActionPrice,omitempty" xml:"ActionPrice,omitempty"`
	BePromotion    *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	BindId         *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
	ContainerName  *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	EslBarCode     *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	EslConnectAp   *string `json:"EslConnectAp,omitempty" xml:"EslConnectAp,omitempty"`
	EslModel       *string `json:"EslModel,omitempty" xml:"EslModel,omitempty"`
	EslPic         *string `json:"EslPic,omitempty" xml:"EslPic,omitempty"`
	EslStatus      *string `json:"EslStatus,omitempty" xml:"EslStatus,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ItemBarCode    *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId         *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemShortTitle *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	ItemTitle      *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	OriginalPrice  *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceUnit      *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	PromotionEnd   *string `json:"PromotionEnd,omitempty" xml:"PromotionEnd,omitempty"`
	PromotionStart *string `json:"PromotionStart,omitempty" xml:"PromotionStart,omitempty"`
	PromotionText  *string `json:"PromotionText,omitempty" xml:"PromotionText,omitempty"`
	// SkuID。
	SkuId           *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	StoreId         *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateSceneId *string `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
}

func (s DescribeBindersResponseBodyEslItemBindInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindersResponseBodyEslItemBindInfos) GoString() string {
	return s.String()
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetActionPrice(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.ActionPrice = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetBePromotion(v bool) *DescribeBindersResponseBodyEslItemBindInfos {
	s.BePromotion = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetBindId(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.BindId = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetContainerName(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.ContainerName = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetEslBarCode(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.EslBarCode = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetEslConnectAp(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.EslConnectAp = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetEslModel(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.EslModel = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetEslPic(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.EslPic = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetEslStatus(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.EslStatus = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetGmtModified(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.GmtModified = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetItemBarCode(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetItemId(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.ItemId = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetItemShortTitle(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.ItemShortTitle = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetItemTitle(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.ItemTitle = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetOriginalPrice(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetPriceUnit(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.PriceUnit = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetPromotionEnd(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.PromotionEnd = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetPromotionStart(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.PromotionStart = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetPromotionText(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.PromotionText = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetSkuId(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.SkuId = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetStoreId(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.StoreId = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetTemplateId(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.TemplateId = &v
	return s
}

func (s *DescribeBindersResponseBodyEslItemBindInfos) SetTemplateSceneId(v string) *DescribeBindersResponseBodyEslItemBindInfos {
	s.TemplateSceneId = &v
	return s
}

type DescribeBindersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBindersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBindersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindersResponse) GoString() string {
	return s.String()
}

func (s *DescribeBindersResponse) SetHeaders(v map[string]*string) *DescribeBindersResponse {
	s.Headers = v
	return s
}

func (s *DescribeBindersResponse) SetStatusCode(v int32) *DescribeBindersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBindersResponse) SetBody(v *DescribeBindersResponseBody) *DescribeBindersResponse {
	s.Body = v
	return s
}

type DescribeCompanyTemplateVersionsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCompanyTemplateVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCompanyTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCompanyTemplateVersionsRequest) SetPageNumber(v int32) *DescribeCompanyTemplateVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsRequest) SetPageSize(v int32) *DescribeCompanyTemplateVersionsRequest {
	s.PageSize = &v
	return s
}

type DescribeCompanyTemplateVersionsResponseBody struct {
	Code           *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                                `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                                `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Versions       []*DescribeCompanyTemplateVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeCompanyTemplateVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCompanyTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetCode(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetDynamicCode(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetDynamicMessage(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetErrorCode(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetErrorMessage(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetMessage(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetPageNumber(v int32) *DescribeCompanyTemplateVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetPageSize(v int32) *DescribeCompanyTemplateVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetRequestId(v string) *DescribeCompanyTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetSuccess(v bool) *DescribeCompanyTemplateVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetTotalCount(v int32) *DescribeCompanyTemplateVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponseBody) SetVersions(v []*DescribeCompanyTemplateVersionsResponseBodyVersions) *DescribeCompanyTemplateVersionsResponseBody {
	s.Versions = v
	return s
}

type DescribeCompanyTemplateVersionsResponseBodyVersions struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeCompanyTemplateVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCompanyTemplateVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *DescribeCompanyTemplateVersionsResponseBodyVersions) SetVersion(v string) *DescribeCompanyTemplateVersionsResponseBodyVersions {
	s.Version = &v
	return s
}

type DescribeCompanyTemplateVersionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCompanyTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCompanyTemplateVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCompanyTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCompanyTemplateVersionsResponse) SetHeaders(v map[string]*string) *DescribeCompanyTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponse) SetStatusCode(v int32) *DescribeCompanyTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCompanyTemplateVersionsResponse) SetBody(v *DescribeCompanyTemplateVersionsResponseBody) *DescribeCompanyTemplateVersionsResponse {
	s.Body = v
	return s
}

type DescribeEslDeviceRequest struct {
	FromDate   *string `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId    *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	ToDate     *string `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
}

func (s DescribeEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEslDeviceRequest) SetFromDate(v string) *DescribeEslDeviceRequest {
	s.FromDate = &v
	return s
}

func (s *DescribeEslDeviceRequest) SetPageNumber(v int64) *DescribeEslDeviceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslDeviceRequest) SetPageSize(v int64) *DescribeEslDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEslDeviceRequest) SetStoreId(v string) *DescribeEslDeviceRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeEslDeviceRequest) SetToDate(v string) *DescribeEslDeviceRequest {
	s.ToDate = &v
	return s
}

type DescribeEslDeviceResponseBody struct {
	EslDetails []*DescribeEslDeviceResponseBodyEslDetails `json:"EslDetails,omitempty" xml:"EslDetails,omitempty" type:"Repeated"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEslDeviceResponseBody) SetEslDetails(v []*DescribeEslDeviceResponseBodyEslDetails) *DescribeEslDeviceResponseBody {
	s.EslDetails = v
	return s
}

func (s *DescribeEslDeviceResponseBody) SetPageNumber(v int64) *DescribeEslDeviceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslDeviceResponseBody) SetPageSize(v int64) *DescribeEslDeviceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEslDeviceResponseBody) SetRequestId(v string) *DescribeEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEslDeviceResponseBody) SetSuccess(v bool) *DescribeEslDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEslDeviceResponseBody) SetTotalCount(v int64) *DescribeEslDeviceResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEslDeviceResponseBodyEslDetails struct {
	EslBarCode     *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ItemBarCode    *int64  `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId         *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemShortTitle *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StoreId        *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeEslDeviceResponseBodyEslDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDeviceResponseBodyEslDetails) GoString() string {
	return s.String()
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetEslBarCode(v string) *DescribeEslDeviceResponseBodyEslDetails {
	s.EslBarCode = &v
	return s
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetItemBarCode(v int64) *DescribeEslDeviceResponseBodyEslDetails {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetItemId(v int64) *DescribeEslDeviceResponseBodyEslDetails {
	s.ItemId = &v
	return s
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetItemShortTitle(v string) *DescribeEslDeviceResponseBodyEslDetails {
	s.ItemShortTitle = &v
	return s
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetLastUpdateTime(v string) *DescribeEslDeviceResponseBodyEslDetails {
	s.LastUpdateTime = &v
	return s
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetStatus(v string) *DescribeEslDeviceResponseBodyEslDetails {
	s.Status = &v
	return s
}

func (s *DescribeEslDeviceResponseBodyEslDetails) SetStoreId(v string) *DescribeEslDeviceResponseBodyEslDetails {
	s.StoreId = &v
	return s
}

type DescribeEslDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEslDeviceResponse) SetHeaders(v map[string]*string) *DescribeEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *DescribeEslDeviceResponse) SetStatusCode(v int32) *DescribeEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEslDeviceResponse) SetBody(v *DescribeEslDeviceResponseBody) *DescribeEslDeviceResponse {
	s.Body = v
	return s
}

type DescribeEslDevicesRequest struct {
	EslBarCode       *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	EslStatus        *string `json:"EslStatus,omitempty" xml:"EslStatus,omitempty"`
	ExtraParams      *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	FromBatteryLevel *int32  `json:"FromBatteryLevel,omitempty" xml:"FromBatteryLevel,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId          *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	ToBatteryLevel   *int32  `json:"ToBatteryLevel,omitempty" xml:"ToBatteryLevel,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeEncode       *string `json:"TypeEncode,omitempty" xml:"TypeEncode,omitempty"`
}

func (s DescribeEslDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesRequest) SetEslBarCode(v string) *DescribeEslDevicesRequest {
	s.EslBarCode = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetEslStatus(v string) *DescribeEslDevicesRequest {
	s.EslStatus = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetExtraParams(v string) *DescribeEslDevicesRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetFromBatteryLevel(v int32) *DescribeEslDevicesRequest {
	s.FromBatteryLevel = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetPageNumber(v int32) *DescribeEslDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetPageSize(v int32) *DescribeEslDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetStoreId(v string) *DescribeEslDevicesRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetToBatteryLevel(v int32) *DescribeEslDevicesRequest {
	s.ToBatteryLevel = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetType(v string) *DescribeEslDevicesRequest {
	s.Type = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetTypeEncode(v string) *DescribeEslDevicesRequest {
	s.TypeEncode = &v
	return s
}

type DescribeEslDevicesResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                     `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EslDevices     []*DescribeEslDevicesResponseBodyEslDevices `json:"EslDevices,omitempty" xml:"EslDevices,omitempty" type:"Repeated"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEslDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponseBody) SetCode(v string) *DescribeEslDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetDynamicCode(v string) *DescribeEslDevicesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetDynamicMessage(v string) *DescribeEslDevicesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetErrorCode(v string) *DescribeEslDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetErrorMessage(v string) *DescribeEslDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetEslDevices(v []*DescribeEslDevicesResponseBodyEslDevices) *DescribeEslDevicesResponseBody {
	s.EslDevices = v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetMessage(v string) *DescribeEslDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetPageNumber(v int32) *DescribeEslDevicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetPageSize(v int32) *DescribeEslDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetRequestId(v string) *DescribeEslDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetSuccess(v bool) *DescribeEslDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetTotalCount(v int32) *DescribeEslDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEslDevicesResponseBodyEslDevices struct {
	BatteryLevel        *int32  `json:"BatteryLevel,omitempty" xml:"BatteryLevel,omitempty"`
	EslBarCode          *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	EslSignal           *int32  `json:"EslSignal,omitempty" xml:"EslSignal,omitempty"`
	EslStatus           *string `json:"EslStatus,omitempty" xml:"EslStatus,omitempty"`
	LastCommunicateTime *string `json:"LastCommunicateTime,omitempty" xml:"LastCommunicateTime,omitempty"`
	LayoutId            *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	LayoutName          *string `json:"LayoutName,omitempty" xml:"LayoutName,omitempty"`
	Mac                 *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Model               *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ScreenHeight        *int32  `json:"ScreenHeight,omitempty" xml:"ScreenHeight,omitempty"`
	ScreenWidth         *int32  `json:"ScreenWidth,omitempty" xml:"ScreenWidth,omitempty"`
	StoreId             *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeEncode          *string `json:"TypeEncode,omitempty" xml:"TypeEncode,omitempty"`
}

func (s DescribeEslDevicesResponseBodyEslDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponseBodyEslDevices) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetBatteryLevel(v int32) *DescribeEslDevicesResponseBodyEslDevices {
	s.BatteryLevel = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetEslBarCode(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.EslBarCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetEslSignal(v int32) *DescribeEslDevicesResponseBodyEslDevices {
	s.EslSignal = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetEslStatus(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.EslStatus = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetLastCommunicateTime(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.LastCommunicateTime = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetLayoutId(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.LayoutId = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetLayoutName(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.LayoutName = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetMac(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.Mac = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetModel(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.Model = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetScreenHeight(v int32) *DescribeEslDevicesResponseBodyEslDevices {
	s.ScreenHeight = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetScreenWidth(v int32) *DescribeEslDevicesResponseBodyEslDevices {
	s.ScreenWidth = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetStoreId(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.StoreId = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetType(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.Type = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetTypeEncode(v string) *DescribeEslDevicesResponseBodyEslDevices {
	s.TypeEncode = &v
	return s
}

type DescribeEslDevicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEslDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEslDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponse) SetHeaders(v map[string]*string) *DescribeEslDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEslDevicesResponse) SetStatusCode(v int32) *DescribeEslDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEslDevicesResponse) SetBody(v *DescribeEslDevicesResponseBody) *DescribeEslDevicesResponse {
	s.Body = v
	return s
}

type DescribeEslModelByTemplateVersionRequest struct {
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s DescribeEslModelByTemplateVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslModelByTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeEslModelByTemplateVersionRequest) SetPageNumber(v int32) *DescribeEslModelByTemplateVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionRequest) SetPageSize(v int32) *DescribeEslModelByTemplateVersionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionRequest) SetTemplateVersion(v string) *DescribeEslModelByTemplateVersionRequest {
	s.TemplateVersion = &v
	return s
}

type DescribeEslModelByTemplateVersionResponseBody struct {
	Code           *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                                   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                                   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EslModels      []*DescribeEslModelByTemplateVersionResponseBodyEslModels `json:"EslModels,omitempty" xml:"EslModels,omitempty" type:"Repeated"`
	Message        *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEslModelByTemplateVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslModelByTemplateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetCode(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetDynamicCode(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetDynamicMessage(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetErrorCode(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetErrorMessage(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetEslModels(v []*DescribeEslModelByTemplateVersionResponseBodyEslModels) *DescribeEslModelByTemplateVersionResponseBody {
	s.EslModels = v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetMessage(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetPageNumber(v int32) *DescribeEslModelByTemplateVersionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetPageSize(v int32) *DescribeEslModelByTemplateVersionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetRequestId(v string) *DescribeEslModelByTemplateVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetSuccess(v bool) *DescribeEslModelByTemplateVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBody) SetTotalCount(v int32) *DescribeEslModelByTemplateVersionResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEslModelByTemplateVersionResponseBodyEslModels struct {
	DeviceType      *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EslPhysicalSize *string `json:"EslPhysicalSize,omitempty" xml:"EslPhysicalSize,omitempty"`
	EslSize         *string `json:"EslSize,omitempty" xml:"EslSize,omitempty"`
	Image           *string `json:"Image,omitempty" xml:"Image,omitempty"`
	ModelId         *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScreenHeight    *int32  `json:"ScreenHeight,omitempty" xml:"ScreenHeight,omitempty"`
	ScreenWidth     *int32  `json:"ScreenWidth,omitempty" xml:"ScreenWidth,omitempty"`
	Vendor          *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeEslModelByTemplateVersionResponseBodyEslModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslModelByTemplateVersionResponseBodyEslModels) GoString() string {
	return s.String()
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetDeviceType(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.DeviceType = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetEslPhysicalSize(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.EslPhysicalSize = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetEslSize(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.EslSize = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetImage(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.Image = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetModelId(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.ModelId = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetName(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.Name = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetScreenHeight(v int32) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.ScreenHeight = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetScreenWidth(v int32) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.ScreenWidth = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponseBodyEslModels) SetVendor(v string) *DescribeEslModelByTemplateVersionResponseBodyEslModels {
	s.Vendor = &v
	return s
}

type DescribeEslModelByTemplateVersionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEslModelByTemplateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEslModelByTemplateVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslModelByTemplateVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeEslModelByTemplateVersionResponse) SetHeaders(v map[string]*string) *DescribeEslModelByTemplateVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponse) SetStatusCode(v int32) *DescribeEslModelByTemplateVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEslModelByTemplateVersionResponse) SetBody(v *DescribeEslModelByTemplateVersionResponseBody) *DescribeEslModelByTemplateVersionResponse {
	s.Body = v
	return s
}

type DescribeItemsRequest struct {
	BePromotion *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId      *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle   *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SkuID。
	SkuId   *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeItemsRequest) SetBePromotion(v bool) *DescribeItemsRequest {
	s.BePromotion = &v
	return s
}

func (s *DescribeItemsRequest) SetExtraParams(v string) *DescribeItemsRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeItemsRequest) SetItemBarCode(v string) *DescribeItemsRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeItemsRequest) SetItemId(v string) *DescribeItemsRequest {
	s.ItemId = &v
	return s
}

func (s *DescribeItemsRequest) SetItemTitle(v string) *DescribeItemsRequest {
	s.ItemTitle = &v
	return s
}

func (s *DescribeItemsRequest) SetPageNumber(v int32) *DescribeItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeItemsRequest) SetPageSize(v int32) *DescribeItemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeItemsRequest) SetSkuId(v string) *DescribeItemsRequest {
	s.SkuId = &v
	return s
}

func (s *DescribeItemsRequest) SetStoreId(v string) *DescribeItemsRequest {
	s.StoreId = &v
	return s
}

type DescribeItemsResponseBody struct {
	Code            *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode     *string                           `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage  *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode       *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Items           []*DescribeItemsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Message         *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber      *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TemplateSceneId *string                           `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
	TotalCount      *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponseBody) SetCode(v string) *DescribeItemsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeItemsResponseBody) SetDynamicCode(v string) *DescribeItemsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeItemsResponseBody) SetDynamicMessage(v string) *DescribeItemsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeItemsResponseBody) SetErrorCode(v string) *DescribeItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeItemsResponseBody) SetErrorMessage(v string) *DescribeItemsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeItemsResponseBody) SetItems(v []*DescribeItemsResponseBodyItems) *DescribeItemsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeItemsResponseBody) SetMessage(v string) *DescribeItemsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeItemsResponseBody) SetPageNumber(v int32) *DescribeItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeItemsResponseBody) SetPageSize(v int32) *DescribeItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeItemsResponseBody) SetRequestId(v string) *DescribeItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeItemsResponseBody) SetSuccess(v bool) *DescribeItemsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeItemsResponseBody) SetTemplateSceneId(v string) *DescribeItemsResponseBody {
	s.TemplateSceneId = &v
	return s
}

func (s *DescribeItemsResponseBody) SetTotalCount(v int32) *DescribeItemsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeItemsResponseBodyItems struct {
	ActionPrice       *int32  `json:"ActionPrice,omitempty" xml:"ActionPrice,omitempty"`
	BeClearance       *bool   `json:"BeClearance,omitempty" xml:"BeClearance,omitempty"`
	BeMember          *bool   `json:"BeMember,omitempty" xml:"BeMember,omitempty"`
	BePromotion       *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	BeSourceCode      *bool   `json:"BeSourceCode,omitempty" xml:"BeSourceCode,omitempty"`
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	CategoryName      *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CustomizeFeatureA *string `json:"CustomizeFeatureA,omitempty" xml:"CustomizeFeatureA,omitempty"`
	CustomizeFeatureB *string `json:"CustomizeFeatureB,omitempty" xml:"CustomizeFeatureB,omitempty"`
	CustomizeFeatureC *string `json:"CustomizeFeatureC,omitempty" xml:"CustomizeFeatureC,omitempty"`
	CustomizeFeatureD *string `json:"CustomizeFeatureD,omitempty" xml:"CustomizeFeatureD,omitempty"`
	CustomizeFeatureE *string `json:"CustomizeFeatureE,omitempty" xml:"CustomizeFeatureE,omitempty"`
	CustomizeFeatureF *string `json:"CustomizeFeatureF,omitempty" xml:"CustomizeFeatureF,omitempty"`
	CustomizeFeatureG *string `json:"CustomizeFeatureG,omitempty" xml:"CustomizeFeatureG,omitempty"`
	CustomizeFeatureH *string `json:"CustomizeFeatureH,omitempty" xml:"CustomizeFeatureH,omitempty"`
	CustomizeFeatureI *string `json:"CustomizeFeatureI,omitempty" xml:"CustomizeFeatureI,omitempty"`
	CustomizeFeatureJ *string `json:"CustomizeFeatureJ,omitempty" xml:"CustomizeFeatureJ,omitempty"`
	CustomizeFeatureK *string `json:"CustomizeFeatureK,omitempty" xml:"CustomizeFeatureK,omitempty"`
	CustomizeFeatureL *string `json:"CustomizeFeatureL,omitempty" xml:"CustomizeFeatureL,omitempty"`
	CustomizeFeatureM *string `json:"CustomizeFeatureM,omitempty" xml:"CustomizeFeatureM,omitempty"`
	CustomizeFeatureN *string `json:"CustomizeFeatureN,omitempty" xml:"CustomizeFeatureN,omitempty"`
	CustomizeFeatureO *string `json:"CustomizeFeatureO,omitempty" xml:"CustomizeFeatureO,omitempty"`
	CustomizeFeatureP *string `json:"CustomizeFeatureP,omitempty" xml:"CustomizeFeatureP,omitempty"`
	CustomizeFeatureQ *string `json:"CustomizeFeatureQ,omitempty" xml:"CustomizeFeatureQ,omitempty"`
	CustomizeFeatureR *string `json:"CustomizeFeatureR,omitempty" xml:"CustomizeFeatureR,omitempty"`
	CustomizeFeatureS *string `json:"CustomizeFeatureS,omitempty" xml:"CustomizeFeatureS,omitempty"`
	CustomizeFeatureT *string `json:"CustomizeFeatureT,omitempty" xml:"CustomizeFeatureT,omitempty"`
	CustomizeFeatureU *string `json:"CustomizeFeatureU,omitempty" xml:"CustomizeFeatureU,omitempty"`
	CustomizeFeatureV *string `json:"CustomizeFeatureV,omitempty" xml:"CustomizeFeatureV,omitempty"`
	CustomizeFeatureW *string `json:"CustomizeFeatureW,omitempty" xml:"CustomizeFeatureW,omitempty"`
	CustomizeFeatureX *string `json:"CustomizeFeatureX,omitempty" xml:"CustomizeFeatureX,omitempty"`
	CustomizeFeatureY *string `json:"CustomizeFeatureY,omitempty" xml:"CustomizeFeatureY,omitempty"`
	CustomizeFeatureZ *string `json:"CustomizeFeatureZ,omitempty" xml:"CustomizeFeatureZ,omitempty"`
	EnergyEfficiency  *string `json:"EnergyEfficiency,omitempty" xml:"EnergyEfficiency,omitempty"`
	ForestFirstId     *string `json:"ForestFirstId,omitempty" xml:"ForestFirstId,omitempty"`
	ForestSecondId    *string `json:"ForestSecondId,omitempty" xml:"ForestSecondId,omitempty"`
	GmtCreate         *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified       *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InventoryStatus   *string `json:"InventoryStatus,omitempty" xml:"InventoryStatus,omitempty"`
	ItemBarCode       *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemInfoIndex     *int32  `json:"ItemInfoIndex,omitempty" xml:"ItemInfoIndex,omitempty"`
	ItemPicUrl        *string `json:"ItemPicUrl,omitempty" xml:"ItemPicUrl,omitempty"`
	ItemQrCode        *string `json:"ItemQrCode,omitempty" xml:"ItemQrCode,omitempty"`
	ItemShortTitle    *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	ItemTitle         *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	Manufacturer      *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	Material          *string `json:"Material,omitempty" xml:"Material,omitempty"`
	MemberPrice       *int32  `json:"MemberPrice,omitempty" xml:"MemberPrice,omitempty"`
	ModelNumber       *string `json:"ModelNumber,omitempty" xml:"ModelNumber,omitempty"`
	OriginalPrice     *int32  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceUnit         *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	ProductionPlace   *string `json:"ProductionPlace,omitempty" xml:"ProductionPlace,omitempty"`
	PromotionEnd      *string `json:"PromotionEnd,omitempty" xml:"PromotionEnd,omitempty"`
	PromotionReason   *string `json:"PromotionReason,omitempty" xml:"PromotionReason,omitempty"`
	PromotionStart    *string `json:"PromotionStart,omitempty" xml:"PromotionStart,omitempty"`
	PromotionText     *string `json:"PromotionText,omitempty" xml:"PromotionText,omitempty"`
	Rank              *string `json:"Rank,omitempty" xml:"Rank,omitempty"`
	SaleSpec          *string `json:"SaleSpec,omitempty" xml:"SaleSpec,omitempty"`
	SalesPrice        *int32  `json:"SalesPrice,omitempty" xml:"SalesPrice,omitempty"`
	// SKuID。
	SkuId           *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SourceCode      *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	SuggestPrice    *int32  `json:"SuggestPrice,omitempty" xml:"SuggestPrice,omitempty"`
	SupplierName    *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	TaxFee          *string `json:"TaxFee,omitempty" xml:"TaxFee,omitempty"`
	TemplateSceneId *string `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
}

func (s DescribeItemsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponseBodyItems) SetActionPrice(v int32) *DescribeItemsResponseBodyItems {
	s.ActionPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetBeClearance(v bool) *DescribeItemsResponseBodyItems {
	s.BeClearance = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetBeMember(v bool) *DescribeItemsResponseBodyItems {
	s.BeMember = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetBePromotion(v bool) *DescribeItemsResponseBodyItems {
	s.BePromotion = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetBeSourceCode(v bool) *DescribeItemsResponseBodyItems {
	s.BeSourceCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetBrandName(v string) *DescribeItemsResponseBodyItems {
	s.BrandName = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCategoryName(v string) *DescribeItemsResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureA(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureA = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureB(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureB = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureC(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureC = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureD(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureD = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureE(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureE = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureF(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureF = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureG(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureG = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureH(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureH = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureI(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureI = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureJ(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureJ = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureK(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureK = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureL(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureL = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureM(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureM = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureN(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureN = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureO(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureO = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureP(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureP = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureQ(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureQ = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureR(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureR = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureS(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureS = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureT(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureT = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureU(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureU = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureV(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureV = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureW(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureW = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureX(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureX = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureY(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureY = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetCustomizeFeatureZ(v string) *DescribeItemsResponseBodyItems {
	s.CustomizeFeatureZ = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetEnergyEfficiency(v string) *DescribeItemsResponseBodyItems {
	s.EnergyEfficiency = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetForestFirstId(v string) *DescribeItemsResponseBodyItems {
	s.ForestFirstId = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetForestSecondId(v string) *DescribeItemsResponseBodyItems {
	s.ForestSecondId = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetGmtCreate(v string) *DescribeItemsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetGmtModified(v string) *DescribeItemsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetInventoryStatus(v string) *DescribeItemsResponseBodyItems {
	s.InventoryStatus = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemBarCode(v string) *DescribeItemsResponseBodyItems {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemId(v string) *DescribeItemsResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemInfoIndex(v int32) *DescribeItemsResponseBodyItems {
	s.ItemInfoIndex = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemPicUrl(v string) *DescribeItemsResponseBodyItems {
	s.ItemPicUrl = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemQrCode(v string) *DescribeItemsResponseBodyItems {
	s.ItemQrCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemShortTitle(v string) *DescribeItemsResponseBodyItems {
	s.ItemShortTitle = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetItemTitle(v string) *DescribeItemsResponseBodyItems {
	s.ItemTitle = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetManufacturer(v string) *DescribeItemsResponseBodyItems {
	s.Manufacturer = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetMaterial(v string) *DescribeItemsResponseBodyItems {
	s.Material = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetMemberPrice(v int32) *DescribeItemsResponseBodyItems {
	s.MemberPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetModelNumber(v string) *DescribeItemsResponseBodyItems {
	s.ModelNumber = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetOriginalPrice(v int32) *DescribeItemsResponseBodyItems {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetPriceUnit(v string) *DescribeItemsResponseBodyItems {
	s.PriceUnit = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetProductionPlace(v string) *DescribeItemsResponseBodyItems {
	s.ProductionPlace = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetPromotionEnd(v string) *DescribeItemsResponseBodyItems {
	s.PromotionEnd = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetPromotionReason(v string) *DescribeItemsResponseBodyItems {
	s.PromotionReason = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetPromotionStart(v string) *DescribeItemsResponseBodyItems {
	s.PromotionStart = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetPromotionText(v string) *DescribeItemsResponseBodyItems {
	s.PromotionText = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetRank(v string) *DescribeItemsResponseBodyItems {
	s.Rank = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetSaleSpec(v string) *DescribeItemsResponseBodyItems {
	s.SaleSpec = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetSalesPrice(v int32) *DescribeItemsResponseBodyItems {
	s.SalesPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetSkuId(v string) *DescribeItemsResponseBodyItems {
	s.SkuId = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetSourceCode(v string) *DescribeItemsResponseBodyItems {
	s.SourceCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetSuggestPrice(v int32) *DescribeItemsResponseBodyItems {
	s.SuggestPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetSupplierName(v string) *DescribeItemsResponseBodyItems {
	s.SupplierName = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetTaxFee(v string) *DescribeItemsResponseBodyItems {
	s.TaxFee = &v
	return s
}

func (s *DescribeItemsResponseBodyItems) SetTemplateSceneId(v string) *DescribeItemsResponseBodyItems {
	s.TemplateSceneId = &v
	return s
}

type DescribeItemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponse) SetHeaders(v map[string]*string) *DescribeItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeItemsResponse) SetStatusCode(v int32) *DescribeItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeItemsResponse) SetBody(v *DescribeItemsResponseBody) *DescribeItemsResponse {
	s.Body = v
	return s
}

type DescribeStoreByTemplateVersionRequest struct {
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s DescribeStoreByTemplateVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreByTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoreByTemplateVersionRequest) SetTemplateVersion(v string) *DescribeStoreByTemplateVersionRequest {
	s.TemplateVersion = &v
	return s
}

type DescribeStoreByTemplateVersionResponseBody struct {
	Code           *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stores         []*DescribeStoreByTemplateVersionResponseBodyStores `json:"Stores,omitempty" xml:"Stores,omitempty" type:"Repeated"`
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStoreByTemplateVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreByTemplateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetCode(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetDynamicCode(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetDynamicMessage(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetErrorCode(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetErrorMessage(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetMessage(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetRequestId(v string) *DescribeStoreByTemplateVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetStores(v []*DescribeStoreByTemplateVersionResponseBodyStores) *DescribeStoreByTemplateVersionResponseBody {
	s.Stores = v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBody) SetSuccess(v bool) *DescribeStoreByTemplateVersionResponseBody {
	s.Success = &v
	return s
}

type DescribeStoreByTemplateVersionResponseBodyStores struct {
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Level           *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId        *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	StoreId         *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName       *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TimeZone        *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UserStoreCode   *string `json:"UserStoreCode,omitempty" xml:"UserStoreCode,omitempty"`
}

func (s DescribeStoreByTemplateVersionResponseBodyStores) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreByTemplateVersionResponseBodyStores) GoString() string {
	return s.String()
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetGmtModified(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.GmtModified = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetLevel(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.Level = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetParentId(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.ParentId = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetPhone(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.Phone = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetStoreId(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.StoreId = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetStoreName(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.StoreName = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetTemplateVersion(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.TemplateVersion = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetTimeZone(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.TimeZone = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponseBodyStores) SetUserStoreCode(v string) *DescribeStoreByTemplateVersionResponseBodyStores {
	s.UserStoreCode = &v
	return s
}

type DescribeStoreByTemplateVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoreByTemplateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoreByTemplateVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreByTemplateVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoreByTemplateVersionResponse) SetHeaders(v map[string]*string) *DescribeStoreByTemplateVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoreByTemplateVersionResponse) SetStatusCode(v int32) *DescribeStoreByTemplateVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoreByTemplateVersionResponse) SetBody(v *DescribeStoreByTemplateVersionResponseBody) *DescribeStoreByTemplateVersionResponse {
	s.Body = v
	return s
}

type DescribeStoreConfigRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeStoreConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoreConfigRequest) SetExtraParams(v string) *DescribeStoreConfigRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeStoreConfigRequest) SetStoreId(v string) *DescribeStoreConfigRequest {
	s.StoreId = &v
	return s
}

type DescribeStoreConfigResponseBody struct {
	Code            *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode     *string                                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage  *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode       *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message         *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StoreConfigInfo *DescribeStoreConfigResponseBodyStoreConfigInfo `json:"StoreConfigInfo,omitempty" xml:"StoreConfigInfo,omitempty" type:"Struct"`
	Success         *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStoreConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoreConfigResponseBody) SetCode(v string) *DescribeStoreConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetDynamicCode(v string) *DescribeStoreConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetDynamicMessage(v string) *DescribeStoreConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetErrorCode(v string) *DescribeStoreConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetErrorMessage(v string) *DescribeStoreConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetMessage(v string) *DescribeStoreConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetRequestId(v string) *DescribeStoreConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetStoreConfigInfo(v *DescribeStoreConfigResponseBodyStoreConfigInfo) *DescribeStoreConfigResponseBody {
	s.StoreConfigInfo = v
	return s
}

func (s *DescribeStoreConfigResponseBody) SetSuccess(v bool) *DescribeStoreConfigResponseBody {
	s.Success = &v
	return s
}

type DescribeStoreConfigResponseBodyStoreConfigInfo struct {
	EnableNotification      *bool                                                              `json:"EnableNotification,omitempty" xml:"EnableNotification,omitempty"`
	NotificationSilentTimes *string                                                            `json:"NotificationSilentTimes,omitempty" xml:"NotificationSilentTimes,omitempty"`
	NotificationWebHook     *string                                                            `json:"NotificationWebHook,omitempty" xml:"NotificationWebHook,omitempty"`
	StoreId                 *string                                                            `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	SubscribeContents       []*DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents `json:"SubscribeContents,omitempty" xml:"SubscribeContents,omitempty" type:"Repeated"`
}

func (s DescribeStoreConfigResponseBodyStoreConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreConfigResponseBodyStoreConfigInfo) GoString() string {
	return s.String()
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfo) SetEnableNotification(v bool) *DescribeStoreConfigResponseBodyStoreConfigInfo {
	s.EnableNotification = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfo) SetNotificationSilentTimes(v string) *DescribeStoreConfigResponseBodyStoreConfigInfo {
	s.NotificationSilentTimes = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfo) SetNotificationWebHook(v string) *DescribeStoreConfigResponseBodyStoreConfigInfo {
	s.NotificationWebHook = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfo) SetStoreId(v string) *DescribeStoreConfigResponseBodyStoreConfigInfo {
	s.StoreId = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfo) SetSubscribeContents(v []*DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) *DescribeStoreConfigResponseBodyStoreConfigInfo {
	s.SubscribeContents = v
	return s
}

type DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents struct {
	AtAll        *bool   `json:"AtAll,omitempty" xml:"AtAll,omitempty"`
	AtMobileList *string `json:"AtMobileList,omitempty" xml:"AtMobileList,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Enable       *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Threshold    *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) GoString() string {
	return s.String()
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) SetAtAll(v bool) *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents {
	s.AtAll = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) SetAtMobileList(v string) *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents {
	s.AtMobileList = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) SetCategory(v string) *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents {
	s.Category = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) SetEnable(v bool) *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents {
	s.Enable = &v
	return s
}

func (s *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents) SetThreshold(v string) *DescribeStoreConfigResponseBodyStoreConfigInfoSubscribeContents {
	s.Threshold = &v
	return s
}

type DescribeStoreConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoreConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoreConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoreConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoreConfigResponse) SetHeaders(v map[string]*string) *DescribeStoreConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoreConfigResponse) SetStatusCode(v int32) *DescribeStoreConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoreConfigResponse) SetBody(v *DescribeStoreConfigResponseBody) *DescribeStoreConfigResponse {
	s.Body = v
	return s
}

type DescribeStoresRequest struct {
	ExtraParams     *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	FromDate        *string `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId         *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName       *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	ToDate          *string `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
	UserStoreCode   *string `json:"UserStoreCode,omitempty" xml:"UserStoreCode,omitempty"`
}

func (s DescribeStoresRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoresRequest) SetExtraParams(v string) *DescribeStoresRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeStoresRequest) SetFromDate(v string) *DescribeStoresRequest {
	s.FromDate = &v
	return s
}

func (s *DescribeStoresRequest) SetPageNumber(v int32) *DescribeStoresRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoresRequest) SetPageSize(v int32) *DescribeStoresRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoresRequest) SetStoreId(v string) *DescribeStoresRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeStoresRequest) SetStoreName(v string) *DescribeStoresRequest {
	s.StoreName = &v
	return s
}

func (s *DescribeStoresRequest) SetTemplateVersion(v string) *DescribeStoresRequest {
	s.TemplateVersion = &v
	return s
}

func (s *DescribeStoresRequest) SetToDate(v string) *DescribeStoresRequest {
	s.ToDate = &v
	return s
}

func (s *DescribeStoresRequest) SetUserStoreCode(v string) *DescribeStoresRequest {
	s.UserStoreCode = &v
	return s
}

type DescribeStoresResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stores         []*DescribeStoresResponseBodyStores `json:"Stores,omitempty" xml:"Stores,omitempty" type:"Repeated"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponseBody) SetCode(v string) *DescribeStoresResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStoresResponseBody) SetDynamicCode(v string) *DescribeStoresResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeStoresResponseBody) SetDynamicMessage(v string) *DescribeStoresResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeStoresResponseBody) SetErrorCode(v string) *DescribeStoresResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeStoresResponseBody) SetErrorMessage(v string) *DescribeStoresResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeStoresResponseBody) SetMessage(v string) *DescribeStoresResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStoresResponseBody) SetPageNumber(v int32) *DescribeStoresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoresResponseBody) SetPageSize(v int32) *DescribeStoresResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoresResponseBody) SetRequestId(v string) *DescribeStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoresResponseBody) SetStores(v []*DescribeStoresResponseBodyStores) *DescribeStoresResponseBody {
	s.Stores = v
	return s
}

func (s *DescribeStoresResponseBody) SetSuccess(v bool) *DescribeStoresResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeStoresResponseBody) SetTotalCount(v int32) *DescribeStoresResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeStoresResponseBodyStores struct {
	AutoUnbindDays       *int32  `json:"AutoUnbindDays,omitempty" xml:"AutoUnbindDays,omitempty"`
	AutoUnbindOfflineEsl *bool   `json:"AutoUnbindOfflineEsl,omitempty" xml:"AutoUnbindOfflineEsl,omitempty"`
	BarCodeEncode        *int32  `json:"BarCodeEncode,omitempty" xml:"BarCodeEncode,omitempty"`
	GmtCreate            *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Level                *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId             *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Phone                *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	StoreId              *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName            *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	TemplateVersion      *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TimeZone             *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UserStoreCode        *string `json:"UserStoreCode,omitempty" xml:"UserStoreCode,omitempty"`
}

func (s DescribeStoresResponseBodyStores) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponseBodyStores) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponseBodyStores) SetAutoUnbindDays(v int32) *DescribeStoresResponseBodyStores {
	s.AutoUnbindDays = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetAutoUnbindOfflineEsl(v bool) *DescribeStoresResponseBodyStores {
	s.AutoUnbindOfflineEsl = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetBarCodeEncode(v int32) *DescribeStoresResponseBodyStores {
	s.BarCodeEncode = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetGmtCreate(v string) *DescribeStoresResponseBodyStores {
	s.GmtCreate = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetGmtModified(v string) *DescribeStoresResponseBodyStores {
	s.GmtModified = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetLevel(v string) *DescribeStoresResponseBodyStores {
	s.Level = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetParentId(v string) *DescribeStoresResponseBodyStores {
	s.ParentId = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetPhone(v string) *DescribeStoresResponseBodyStores {
	s.Phone = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetStoreId(v string) *DescribeStoresResponseBodyStores {
	s.StoreId = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetStoreName(v string) *DescribeStoresResponseBodyStores {
	s.StoreName = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetTemplateVersion(v string) *DescribeStoresResponseBodyStores {
	s.TemplateVersion = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetTimeZone(v string) *DescribeStoresResponseBodyStores {
	s.TimeZone = &v
	return s
}

func (s *DescribeStoresResponseBodyStores) SetUserStoreCode(v string) *DescribeStoresResponseBodyStores {
	s.UserStoreCode = &v
	return s
}

type DescribeStoresResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponse) SetHeaders(v map[string]*string) *DescribeStoresResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoresResponse) SetStatusCode(v int32) *DescribeStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoresResponse) SetBody(v *DescribeStoresResponseBody) *DescribeStoresResponse {
	s.Body = v
	return s
}

type DescribeTemplateByModelRequest struct {
	DeviceType      *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EslSize         *string `json:"EslSize,omitempty" xml:"EslSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s DescribeTemplateByModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateByModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateByModelRequest) SetDeviceType(v string) *DescribeTemplateByModelRequest {
	s.DeviceType = &v
	return s
}

func (s *DescribeTemplateByModelRequest) SetEslSize(v string) *DescribeTemplateByModelRequest {
	s.EslSize = &v
	return s
}

func (s *DescribeTemplateByModelRequest) SetPageNumber(v int32) *DescribeTemplateByModelRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplateByModelRequest) SetPageSize(v int32) *DescribeTemplateByModelRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplateByModelRequest) SetTemplateVersion(v string) *DescribeTemplateByModelRequest {
	s.TemplateVersion = &v
	return s
}

type DescribeTemplateByModelResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                     `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Items          []*DescribeTemplateByModelResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTemplateByModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateByModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateByModelResponseBody) SetCode(v string) *DescribeTemplateByModelResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetDynamicCode(v string) *DescribeTemplateByModelResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetDynamicMessage(v string) *DescribeTemplateByModelResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetErrorCode(v string) *DescribeTemplateByModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetErrorMessage(v string) *DescribeTemplateByModelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetItems(v []*DescribeTemplateByModelResponseBodyItems) *DescribeTemplateByModelResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetMessage(v string) *DescribeTemplateByModelResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetPageNumber(v int32) *DescribeTemplateByModelResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetPageSize(v int32) *DescribeTemplateByModelResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetRequestId(v string) *DescribeTemplateByModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetSuccess(v bool) *DescribeTemplateByModelResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTemplateByModelResponseBody) SetTotalCount(v int32) *DescribeTemplateByModelResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTemplateByModelResponseBodyItems struct {
	BasePicture     *string `json:"BasePicture,omitempty" xml:"BasePicture,omitempty"`
	Brand           *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	EslSize         *string `json:"EslSize,omitempty" xml:"EslSize,omitempty"`
	EslType         *string `json:"EslType,omitempty" xml:"EslType,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Layout          *string `json:"Layout,omitempty" xml:"Layout,omitempty"`
	Scene           *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateSceneId *string `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeTemplateByModelResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateByModelResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTemplateByModelResponseBodyItems) SetBasePicture(v string) *DescribeTemplateByModelResponseBodyItems {
	s.BasePicture = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetBrand(v string) *DescribeTemplateByModelResponseBodyItems {
	s.Brand = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetEslSize(v string) *DescribeTemplateByModelResponseBodyItems {
	s.EslSize = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetEslType(v string) *DescribeTemplateByModelResponseBodyItems {
	s.EslType = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetHeight(v int64) *DescribeTemplateByModelResponseBodyItems {
	s.Height = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetLayout(v string) *DescribeTemplateByModelResponseBodyItems {
	s.Layout = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetScene(v string) *DescribeTemplateByModelResponseBodyItems {
	s.Scene = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetTemplateId(v string) *DescribeTemplateByModelResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetTemplateName(v string) *DescribeTemplateByModelResponseBodyItems {
	s.TemplateName = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetTemplateSceneId(v string) *DescribeTemplateByModelResponseBodyItems {
	s.TemplateSceneId = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetTemplateVersion(v string) *DescribeTemplateByModelResponseBodyItems {
	s.TemplateVersion = &v
	return s
}

func (s *DescribeTemplateByModelResponseBodyItems) SetWidth(v int64) *DescribeTemplateByModelResponseBodyItems {
	s.Width = &v
	return s
}

type DescribeTemplateByModelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateByModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateByModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateByModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateByModelResponse) SetHeaders(v map[string]*string) *DescribeTemplateByModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateByModelResponse) SetStatusCode(v int32) *DescribeTemplateByModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateByModelResponse) SetBody(v *DescribeTemplateByModelResponseBody) *DescribeTemplateByModelResponse {
	s.Body = v
	return s
}

type DescribeUserLogRequest struct {
	EslBarCode      *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ExtraParams     *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	FromDate        *string `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	ItemBarCode     *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemShortTitle  *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	LogId           *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	OperationType   *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId         *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	ToDate          *string `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeUserLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLogRequest) SetEslBarCode(v string) *DescribeUserLogRequest {
	s.EslBarCode = &v
	return s
}

func (s *DescribeUserLogRequest) SetExtraParams(v string) *DescribeUserLogRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeUserLogRequest) SetFromDate(v string) *DescribeUserLogRequest {
	s.FromDate = &v
	return s
}

func (s *DescribeUserLogRequest) SetItemBarCode(v string) *DescribeUserLogRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeUserLogRequest) SetItemShortTitle(v string) *DescribeUserLogRequest {
	s.ItemShortTitle = &v
	return s
}

func (s *DescribeUserLogRequest) SetLogId(v string) *DescribeUserLogRequest {
	s.LogId = &v
	return s
}

func (s *DescribeUserLogRequest) SetOperationStatus(v string) *DescribeUserLogRequest {
	s.OperationStatus = &v
	return s
}

func (s *DescribeUserLogRequest) SetOperationType(v string) *DescribeUserLogRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeUserLogRequest) SetPageNumber(v int32) *DescribeUserLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserLogRequest) SetPageSize(v int32) *DescribeUserLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserLogRequest) SetStoreId(v string) *DescribeUserLogRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeUserLogRequest) SetToDate(v string) *DescribeUserLogRequest {
	s.ToDate = &v
	return s
}

func (s *DescribeUserLogRequest) SetUserId(v string) *DescribeUserLogRequest {
	s.UserId = &v
	return s
}

type DescribeUserLogResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserLogs       []*DescribeUserLogResponseBodyUserLogs `json:"UserLogs,omitempty" xml:"UserLogs,omitempty" type:"Repeated"`
}

func (s DescribeUserLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponseBody) SetCode(v string) *DescribeUserLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetDynamicCode(v string) *DescribeUserLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetDynamicMessage(v string) *DescribeUserLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetErrorCode(v string) *DescribeUserLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetErrorMessage(v string) *DescribeUserLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetMessage(v string) *DescribeUserLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetPageNumber(v int32) *DescribeUserLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetPageSize(v int32) *DescribeUserLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetRequestId(v string) *DescribeUserLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetSuccess(v bool) *DescribeUserLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetTotalCount(v int32) *DescribeUserLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetUserLogs(v []*DescribeUserLogResponseBodyUserLogs) *DescribeUserLogResponseBody {
	s.UserLogs = v
	return s
}

type DescribeUserLogResponseBodyUserLogs struct {
	ActionPrice           *string `json:"ActionPrice,omitempty" xml:"ActionPrice,omitempty"`
	BePromotion           *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	EslBarCode            *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	EslSignal             *int32  `json:"EslSignal,omitempty" xml:"EslSignal,omitempty"`
	GmtCreate             *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified           *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	I18nResultKey         *string `json:"I18nResultKey,omitempty" xml:"I18nResultKey,omitempty"`
	ItemBarCode           *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId                *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemShortTitle        *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	LogId                 *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	OperationResponseTime *string `json:"OperationResponseTime,omitempty" xml:"OperationResponseTime,omitempty"`
	OperationSendTime     *string `json:"OperationSendTime,omitempty" xml:"OperationSendTime,omitempty"`
	OperationStatus       *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	OperationType         *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	PriceUnit             *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	ResultCode            *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	SpendTime             *string `json:"SpendTime,omitempty" xml:"SpendTime,omitempty"`
	StoreId               *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeUserLogResponseBodyUserLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponseBodyUserLogs) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponseBodyUserLogs) SetActionPrice(v string) *DescribeUserLogResponseBodyUserLogs {
	s.ActionPrice = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetBePromotion(v bool) *DescribeUserLogResponseBodyUserLogs {
	s.BePromotion = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetEslBarCode(v string) *DescribeUserLogResponseBodyUserLogs {
	s.EslBarCode = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetEslSignal(v int32) *DescribeUserLogResponseBodyUserLogs {
	s.EslSignal = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetGmtCreate(v string) *DescribeUserLogResponseBodyUserLogs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetGmtModified(v string) *DescribeUserLogResponseBodyUserLogs {
	s.GmtModified = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetI18nResultKey(v string) *DescribeUserLogResponseBodyUserLogs {
	s.I18nResultKey = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetItemBarCode(v string) *DescribeUserLogResponseBodyUserLogs {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetItemId(v string) *DescribeUserLogResponseBodyUserLogs {
	s.ItemId = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetItemShortTitle(v string) *DescribeUserLogResponseBodyUserLogs {
	s.ItemShortTitle = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetLogId(v string) *DescribeUserLogResponseBodyUserLogs {
	s.LogId = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetOperationResponseTime(v string) *DescribeUserLogResponseBodyUserLogs {
	s.OperationResponseTime = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetOperationSendTime(v string) *DescribeUserLogResponseBodyUserLogs {
	s.OperationSendTime = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetOperationStatus(v string) *DescribeUserLogResponseBodyUserLogs {
	s.OperationStatus = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetOperationType(v string) *DescribeUserLogResponseBodyUserLogs {
	s.OperationType = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetPriceUnit(v string) *DescribeUserLogResponseBodyUserLogs {
	s.PriceUnit = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetResultCode(v string) *DescribeUserLogResponseBodyUserLogs {
	s.ResultCode = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetSpendTime(v string) *DescribeUserLogResponseBodyUserLogs {
	s.SpendTime = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetStoreId(v string) *DescribeUserLogResponseBodyUserLogs {
	s.StoreId = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogs) SetUserId(v string) *DescribeUserLogResponseBodyUserLogs {
	s.UserId = &v
	return s
}

type DescribeUserLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponse) SetHeaders(v map[string]*string) *DescribeUserLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLogResponse) SetStatusCode(v int32) *DescribeUserLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserLogResponse) SetBody(v *DescribeUserLogResponseBody) *DescribeUserLogResponse {
	s.Body = v
	return s
}

type DescribeUsersRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType    *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersRequest) SetExtraParams(v string) *DescribeUsersRequest {
	s.ExtraParams = &v
	return s
}

func (s *DescribeUsersRequest) SetPageNumber(v int32) *DescribeUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUsersRequest) SetPageSize(v int32) *DescribeUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUsersRequest) SetUserId(v string) *DescribeUsersRequest {
	s.UserId = &v
	return s
}

func (s *DescribeUsersRequest) SetUserName(v string) *DescribeUsersRequest {
	s.UserName = &v
	return s
}

func (s *DescribeUsersRequest) SetUserType(v string) *DescribeUsersRequest {
	s.UserType = &v
	return s
}

type DescribeUsersResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                           `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users          []*DescribeUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBody) SetCode(v string) *DescribeUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUsersResponseBody) SetDynamicCode(v string) *DescribeUsersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeUsersResponseBody) SetDynamicMessage(v string) *DescribeUsersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeUsersResponseBody) SetErrorCode(v string) *DescribeUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUsersResponseBody) SetErrorMessage(v string) *DescribeUsersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUsersResponseBody) SetMessage(v string) *DescribeUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUsersResponseBody) SetPageNumber(v int32) *DescribeUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUsersResponseBody) SetPageSize(v int32) *DescribeUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUsersResponseBody) SetRequestId(v string) *DescribeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersResponseBody) SetSuccess(v bool) *DescribeUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUsersResponseBody) SetTotalCount(v int32) *DescribeUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUsersResponseBody) SetUsers(v []*DescribeUsersResponseBodyUsers) *DescribeUsersResponseBody {
	s.Users = v
	return s
}

type DescribeUsersResponseBodyUsers struct {
	Bid           *string                                        `json:"Bid,omitempty" xml:"Bid,omitempty"`
	DingTalkInfos []*DescribeUsersResponseBodyUsersDingTalkInfos `json:"DingTalkInfos,omitempty" xml:"DingTalkInfos,omitempty" type:"Repeated"`
	OwnerId       *string                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Stores        *string                                        `json:"Stores,omitempty" xml:"Stores,omitempty"`
	UserId        *string                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string                                        `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType      *string                                        `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsers) SetBid(v string) *DescribeUsersResponseBodyUsers {
	s.Bid = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetDingTalkInfos(v []*DescribeUsersResponseBodyUsersDingTalkInfos) *DescribeUsersResponseBodyUsers {
	s.DingTalkInfos = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOwnerId(v string) *DescribeUsersResponseBodyUsers {
	s.OwnerId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetStores(v string) *DescribeUsersResponseBodyUsers {
	s.Stores = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetUserId(v string) *DescribeUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetUserName(v string) *DescribeUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetUserType(v string) *DescribeUsersResponseBodyUsers {
	s.UserType = &v
	return s
}

type DescribeUsersResponseBodyUsersDingTalkInfos struct {
	DingTalkCompanyId *string `json:"DingTalkCompanyId,omitempty" xml:"DingTalkCompanyId,omitempty"`
	DingTalkUserId    *string `json:"DingTalkUserId,omitempty" xml:"DingTalkUserId,omitempty"`
}

func (s DescribeUsersResponseBodyUsersDingTalkInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersDingTalkInfos) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersDingTalkInfos) SetDingTalkCompanyId(v string) *DescribeUsersResponseBodyUsersDingTalkInfos {
	s.DingTalkCompanyId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersDingTalkInfos) SetDingTalkUserId(v string) *DescribeUsersResponseBodyUsersDingTalkInfos {
	s.DingTalkUserId = &v
	return s
}

type DescribeUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponse) SetHeaders(v map[string]*string) *DescribeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersResponse) SetStatusCode(v int32) *DescribeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersResponse) SetBody(v *DescribeUsersResponseBody) *DescribeUsersResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetExtraParams(v string) *GetUserRequest {
	s.ExtraParams = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	Code           *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                  `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	User           *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetDynamicCode(v string) *GetUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetUserResponseBody) SetDynamicMessage(v string) *GetUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetUserResponseBody) SetErrorCode(v string) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetErrorMessage(v string) *GetUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponseBodyUser struct {
	Bid           *string                                 `json:"Bid,omitempty" xml:"Bid,omitempty"`
	DingTalkInfos []*GetUserResponseBodyUserDingTalkInfos `json:"DingTalkInfos,omitempty" xml:"DingTalkInfos,omitempty" type:"Repeated"`
	OwnerId       *string                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Stores        *string                                 `json:"Stores,omitempty" xml:"Stores,omitempty"`
	UserId        *string                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string                                 `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType      *string                                 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetBid(v string) *GetUserResponseBodyUser {
	s.Bid = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDingTalkInfos(v []*GetUserResponseBodyUserDingTalkInfos) *GetUserResponseBodyUser {
	s.DingTalkInfos = v
	return s
}

func (s *GetUserResponseBodyUser) SetOwnerId(v string) *GetUserResponseBodyUser {
	s.OwnerId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetStores(v string) *GetUserResponseBodyUser {
	s.Stores = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserType(v string) *GetUserResponseBodyUser {
	s.UserType = &v
	return s
}

type GetUserResponseBodyUserDingTalkInfos struct {
	DingTalkCompanyId *string `json:"DingTalkCompanyId,omitempty" xml:"DingTalkCompanyId,omitempty"`
	DingTalkUserId    *string `json:"DingTalkUserId,omitempty" xml:"DingTalkUserId,omitempty"`
}

func (s GetUserResponseBodyUserDingTalkInfos) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserDingTalkInfos) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserDingTalkInfos) SetDingTalkCompanyId(v string) *GetUserResponseBodyUserDingTalkInfos {
	s.DingTalkCompanyId = &v
	return s
}

func (s *GetUserResponseBodyUserDingTalkInfos) SetDingTalkUserId(v string) *GetUserResponseBodyUserDingTalkInfos {
	s.DingTalkUserId = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type QueryTemplateListByGroupIdRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTemplateListByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateListByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *QueryTemplateListByGroupIdRequest) SetGroupId(v string) *QueryTemplateListByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *QueryTemplateListByGroupIdRequest) SetPageNumber(v int32) *QueryTemplateListByGroupIdRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTemplateListByGroupIdRequest) SetPageSize(v int32) *QueryTemplateListByGroupIdRequest {
	s.PageSize = &v
	return s
}

type QueryTemplateListByGroupIdResponseBody struct {
	Code           *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TemplateList   []*QueryTemplateListByGroupIdResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
	TotalCount     *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryTemplateListByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateListByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateListByGroupIdResponseBody) SetCode(v string) *QueryTemplateListByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetDynamicCode(v string) *QueryTemplateListByGroupIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetDynamicMessage(v string) *QueryTemplateListByGroupIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetErrorCode(v string) *QueryTemplateListByGroupIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetErrorMessage(v string) *QueryTemplateListByGroupIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetMessage(v string) *QueryTemplateListByGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetPageNumber(v int32) *QueryTemplateListByGroupIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetPageSize(v int32) *QueryTemplateListByGroupIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetRequestId(v string) *QueryTemplateListByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetSuccess(v bool) *QueryTemplateListByGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetTemplateList(v []*QueryTemplateListByGroupIdResponseBodyTemplateList) *QueryTemplateListByGroupIdResponseBody {
	s.TemplateList = v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBody) SetTotalCount(v int32) *QueryTemplateListByGroupIdResponseBody {
	s.TotalCount = &v
	return s
}

type QueryTemplateListByGroupIdResponseBodyTemplateList struct {
	BasePicture     *string `json:"BasePicture,omitempty" xml:"BasePicture,omitempty"`
	Brand           *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	EslSize         *string `json:"EslSize,omitempty" xml:"EslSize,omitempty"`
	EslType         *string `json:"EslType,omitempty" xml:"EslType,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Layout          *string `json:"Layout,omitempty" xml:"Layout,omitempty"`
	Relation        *bool   `json:"Relation,omitempty" xml:"Relation,omitempty"`
	Scene           *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateSceneId *string `json:"TemplateSceneId,omitempty" xml:"TemplateSceneId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryTemplateListByGroupIdResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateListByGroupIdResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetBasePicture(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.BasePicture = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetBrand(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.Brand = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetEslSize(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.EslSize = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetEslType(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.EslType = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetGroupId(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.GroupId = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetHeight(v int64) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.Height = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetLayout(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.Layout = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetRelation(v bool) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.Relation = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetScene(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.Scene = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetTemplateId(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetTemplateName(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.TemplateName = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetTemplateSceneId(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.TemplateSceneId = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetTemplateVersion(v string) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.TemplateVersion = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponseBodyTemplateList) SetWidth(v int64) *QueryTemplateListByGroupIdResponseBodyTemplateList {
	s.Width = &v
	return s
}

type QueryTemplateListByGroupIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTemplateListByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTemplateListByGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateListByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *QueryTemplateListByGroupIdResponse) SetHeaders(v map[string]*string) *QueryTemplateListByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *QueryTemplateListByGroupIdResponse) SetStatusCode(v int32) *QueryTemplateListByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTemplateListByGroupIdResponse) SetBody(v *QueryTemplateListByGroupIdResponseBody) *QueryTemplateListByGroupIdResponse {
	s.Body = v
	return s
}

type SyncAddMaterialRequest struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SyncAddMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncAddMaterialRequest) GoString() string {
	return s.String()
}

func (s *SyncAddMaterialRequest) SetContent(v string) *SyncAddMaterialRequest {
	s.Content = &v
	return s
}

func (s *SyncAddMaterialRequest) SetName(v string) *SyncAddMaterialRequest {
	s.Name = &v
	return s
}

type SyncAddMaterialResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result         *SyncAddMaterialResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncAddMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncAddMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *SyncAddMaterialResponseBody) SetCode(v string) *SyncAddMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetDynamicCode(v string) *SyncAddMaterialResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetDynamicMessage(v string) *SyncAddMaterialResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetErrorCode(v string) *SyncAddMaterialResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetErrorMessage(v string) *SyncAddMaterialResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetMessage(v string) *SyncAddMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetRequestId(v string) *SyncAddMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncAddMaterialResponseBody) SetResult(v *SyncAddMaterialResponseBodyResult) *SyncAddMaterialResponseBody {
	s.Result = v
	return s
}

func (s *SyncAddMaterialResponseBody) SetSuccess(v bool) *SyncAddMaterialResponseBody {
	s.Success = &v
	return s
}

type SyncAddMaterialResponseBodyResult struct {
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncAddMaterialResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SyncAddMaterialResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SyncAddMaterialResponseBodyResult) SetDynamicCode(v string) *SyncAddMaterialResponseBodyResult {
	s.DynamicCode = &v
	return s
}

func (s *SyncAddMaterialResponseBodyResult) SetDynamicMessage(v string) *SyncAddMaterialResponseBodyResult {
	s.DynamicMessage = &v
	return s
}

func (s *SyncAddMaterialResponseBodyResult) SetErrorCode(v string) *SyncAddMaterialResponseBodyResult {
	s.ErrorCode = &v
	return s
}

func (s *SyncAddMaterialResponseBodyResult) SetMessage(v string) *SyncAddMaterialResponseBodyResult {
	s.Message = &v
	return s
}

func (s *SyncAddMaterialResponseBodyResult) SetSuccess(v bool) *SyncAddMaterialResponseBodyResult {
	s.Success = &v
	return s
}

type SyncAddMaterialResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncAddMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncAddMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncAddMaterialResponse) GoString() string {
	return s.String()
}

func (s *SyncAddMaterialResponse) SetHeaders(v map[string]*string) *SyncAddMaterialResponse {
	s.Headers = v
	return s
}

func (s *SyncAddMaterialResponse) SetStatusCode(v int32) *SyncAddMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncAddMaterialResponse) SetBody(v *SyncAddMaterialResponseBody) *SyncAddMaterialResponse {
	s.Body = v
	return s
}

type UnassignUserRequest struct {
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnassignUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassignUserRequest) GoString() string {
	return s.String()
}

func (s *UnassignUserRequest) SetExtraParams(v string) *UnassignUserRequest {
	s.ExtraParams = &v
	return s
}

func (s *UnassignUserRequest) SetUserId(v string) *UnassignUserRequest {
	s.UserId = &v
	return s
}

type UnassignUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnassignUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnassignUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignUserResponseBody) SetCode(v string) *UnassignUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnassignUserResponseBody) SetDynamicCode(v string) *UnassignUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnassignUserResponseBody) SetDynamicMessage(v string) *UnassignUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnassignUserResponseBody) SetErrorCode(v string) *UnassignUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnassignUserResponseBody) SetErrorMessage(v string) *UnassignUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnassignUserResponseBody) SetMessage(v string) *UnassignUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnassignUserResponseBody) SetRequestId(v string) *UnassignUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassignUserResponseBody) SetSuccess(v bool) *UnassignUserResponseBody {
	s.Success = &v
	return s
}

type UnassignUserResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassignUserResponse) GoString() string {
	return s.String()
}

func (s *UnassignUserResponse) SetHeaders(v map[string]*string) *UnassignUserResponse {
	s.Headers = v
	return s
}

func (s *UnassignUserResponse) SetStatusCode(v int32) *UnassignUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignUserResponse) SetBody(v *UnassignUserResponseBody) *UnassignUserResponse {
	s.Body = v
	return s
}

type UnbindEslDeviceRequest struct {
	Column        *string `json:"Column,omitempty" xml:"Column,omitempty"`
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	EslBarCode    *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ExtraParams   *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	ItemBarCode   *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	Layer         *int32  `json:"Layer,omitempty" xml:"Layer,omitempty"`
	Shelf         *string `json:"Shelf,omitempty" xml:"Shelf,omitempty"`
	StoreId       *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s UnbindEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceRequest) SetColumn(v string) *UnbindEslDeviceRequest {
	s.Column = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetContainerName(v string) *UnbindEslDeviceRequest {
	s.ContainerName = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetEslBarCode(v string) *UnbindEslDeviceRequest {
	s.EslBarCode = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetExtraParams(v string) *UnbindEslDeviceRequest {
	s.ExtraParams = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetItemBarCode(v string) *UnbindEslDeviceRequest {
	s.ItemBarCode = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetLayer(v int32) *UnbindEslDeviceRequest {
	s.Layer = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetShelf(v string) *UnbindEslDeviceRequest {
	s.Shelf = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetStoreId(v string) *UnbindEslDeviceRequest {
	s.StoreId = &v
	return s
}

type UnbindEslDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceResponseBody) SetCode(v string) *UnbindEslDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetDynamicCode(v string) *UnbindEslDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetDynamicMessage(v string) *UnbindEslDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetErrorCode(v string) *UnbindEslDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetErrorMessage(v string) *UnbindEslDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetMessage(v string) *UnbindEslDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetRequestId(v string) *UnbindEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetSuccess(v bool) *UnbindEslDeviceResponseBody {
	s.Success = &v
	return s
}

type UnbindEslDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceResponse) SetHeaders(v map[string]*string) *UnbindEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindEslDeviceResponse) SetStatusCode(v int32) *UnbindEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEslDeviceResponse) SetBody(v *UnbindEslDeviceResponseBody) *UnbindEslDeviceResponse {
	s.Body = v
	return s
}

type UpdateEslDeviceLightRequest struct {
	EslBarCode  *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	Frequency   *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	LedColor    *string `json:"LedColor,omitempty" xml:"LedColor,omitempty"`
	LightUpTime *int32  `json:"LightUpTime,omitempty" xml:"LightUpTime,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s UpdateEslDeviceLightRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightRequest) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightRequest) SetEslBarCode(v string) *UpdateEslDeviceLightRequest {
	s.EslBarCode = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetExtraParams(v string) *UpdateEslDeviceLightRequest {
	s.ExtraParams = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetFrequency(v string) *UpdateEslDeviceLightRequest {
	s.Frequency = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetItemBarCode(v string) *UpdateEslDeviceLightRequest {
	s.ItemBarCode = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetLedColor(v string) *UpdateEslDeviceLightRequest {
	s.LedColor = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetLightUpTime(v int32) *UpdateEslDeviceLightRequest {
	s.LightUpTime = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetStoreId(v string) *UpdateEslDeviceLightRequest {
	s.StoreId = &v
	return s
}

type UpdateEslDeviceLightResponseBody struct {
	Code              *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode       *string                                              `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage    *string                                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode         *string                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FailCount         *int32                                               `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	LightFailEslInfos []*UpdateEslDeviceLightResponseBodyLightFailEslInfos `json:"LightFailEslInfos,omitempty" xml:"LightFailEslInfos,omitempty" type:"Repeated"`
	Message           *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	SuccessCount      *int32                                               `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s UpdateEslDeviceLightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightResponseBody) SetCode(v string) *UpdateEslDeviceLightResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetDynamicCode(v string) *UpdateEslDeviceLightResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetDynamicMessage(v string) *UpdateEslDeviceLightResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetErrorCode(v string) *UpdateEslDeviceLightResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetErrorMessage(v string) *UpdateEslDeviceLightResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetFailCount(v int32) *UpdateEslDeviceLightResponseBody {
	s.FailCount = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetLightFailEslInfos(v []*UpdateEslDeviceLightResponseBodyLightFailEslInfos) *UpdateEslDeviceLightResponseBody {
	s.LightFailEslInfos = v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetMessage(v string) *UpdateEslDeviceLightResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetRequestId(v string) *UpdateEslDeviceLightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetSuccess(v bool) *UpdateEslDeviceLightResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetSuccessCount(v int32) *UpdateEslDeviceLightResponseBody {
	s.SuccessCount = &v
	return s
}

type UpdateEslDeviceLightResponseBodyLightFailEslInfos struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EslBarCode   *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
}

func (s UpdateEslDeviceLightResponseBodyLightFailEslInfos) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightResponseBodyLightFailEslInfos) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightResponseBodyLightFailEslInfos) SetErrorMessage(v string) *UpdateEslDeviceLightResponseBodyLightFailEslInfos {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBodyLightFailEslInfos) SetEslBarCode(v string) *UpdateEslDeviceLightResponseBodyLightFailEslInfos {
	s.EslBarCode = &v
	return s
}

type UpdateEslDeviceLightResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEslDeviceLightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEslDeviceLightResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightResponse) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightResponse) SetHeaders(v map[string]*string) *UpdateEslDeviceLightResponse {
	s.Headers = v
	return s
}

func (s *UpdateEslDeviceLightResponse) SetStatusCode(v int32) *UpdateEslDeviceLightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEslDeviceLightResponse) SetBody(v *UpdateEslDeviceLightResponseBody) *UpdateEslDeviceLightResponse {
	s.Body = v
	return s
}

type UpdateStoreRequest struct {
	AutoUnbindDays       *int32  `json:"AutoUnbindDays,omitempty" xml:"AutoUnbindDays,omitempty"`
	AutoUnbindOfflineEsl *bool   `json:"AutoUnbindOfflineEsl,omitempty" xml:"AutoUnbindOfflineEsl,omitempty"`
	BarCodeEncode        *int32  `json:"BarCodeEncode,omitempty" xml:"BarCodeEncode,omitempty"`
	ExtraParams          *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	Phone                *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	StoreId              *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName            *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	TemplateVersion      *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	Timezone             *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	UserStoreCode        *string `json:"UserStoreCode,omitempty" xml:"UserStoreCode,omitempty"`
}

func (s UpdateStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoreRequest) SetAutoUnbindDays(v int32) *UpdateStoreRequest {
	s.AutoUnbindDays = &v
	return s
}

func (s *UpdateStoreRequest) SetAutoUnbindOfflineEsl(v bool) *UpdateStoreRequest {
	s.AutoUnbindOfflineEsl = &v
	return s
}

func (s *UpdateStoreRequest) SetBarCodeEncode(v int32) *UpdateStoreRequest {
	s.BarCodeEncode = &v
	return s
}

func (s *UpdateStoreRequest) SetExtraParams(v string) *UpdateStoreRequest {
	s.ExtraParams = &v
	return s
}

func (s *UpdateStoreRequest) SetPhone(v string) *UpdateStoreRequest {
	s.Phone = &v
	return s
}

func (s *UpdateStoreRequest) SetStoreId(v string) *UpdateStoreRequest {
	s.StoreId = &v
	return s
}

func (s *UpdateStoreRequest) SetStoreName(v string) *UpdateStoreRequest {
	s.StoreName = &v
	return s
}

func (s *UpdateStoreRequest) SetTemplateVersion(v string) *UpdateStoreRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStoreRequest) SetTimezone(v string) *UpdateStoreRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateStoreRequest) SetUserStoreCode(v string) *UpdateStoreRequest {
	s.UserStoreCode = &v
	return s
}

type UpdateStoreResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoreResponseBody) SetCode(v string) *UpdateStoreResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStoreResponseBody) SetDynamicCode(v string) *UpdateStoreResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateStoreResponseBody) SetDynamicMessage(v string) *UpdateStoreResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateStoreResponseBody) SetErrorCode(v string) *UpdateStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateStoreResponseBody) SetErrorMessage(v string) *UpdateStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateStoreResponseBody) SetMessage(v string) *UpdateStoreResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStoreResponseBody) SetRequestId(v string) *UpdateStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStoreResponseBody) SetSuccess(v bool) *UpdateStoreResponseBody {
	s.Success = &v
	return s
}

type UpdateStoreResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoreResponse) SetHeaders(v map[string]*string) *UpdateStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoreResponse) SetStatusCode(v int32) *UpdateStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStoreResponse) SetBody(v *UpdateStoreResponseBody) *UpdateStoreResponse {
	s.Body = v
	return s
}

type UpdateStoreConfigRequest struct {
	EnableNotification      *bool   `json:"EnableNotification,omitempty" xml:"EnableNotification,omitempty"`
	ExtraParams             *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	NotificationSilentTimes *string `json:"NotificationSilentTimes,omitempty" xml:"NotificationSilentTimes,omitempty"`
	NotificationWebHook     *string `json:"NotificationWebHook,omitempty" xml:"NotificationWebHook,omitempty"`
	StoreId                 *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	SubscribeContents       *string `json:"SubscribeContents,omitempty" xml:"SubscribeContents,omitempty"`
}

func (s UpdateStoreConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoreConfigRequest) SetEnableNotification(v bool) *UpdateStoreConfigRequest {
	s.EnableNotification = &v
	return s
}

func (s *UpdateStoreConfigRequest) SetExtraParams(v string) *UpdateStoreConfigRequest {
	s.ExtraParams = &v
	return s
}

func (s *UpdateStoreConfigRequest) SetNotificationSilentTimes(v string) *UpdateStoreConfigRequest {
	s.NotificationSilentTimes = &v
	return s
}

func (s *UpdateStoreConfigRequest) SetNotificationWebHook(v string) *UpdateStoreConfigRequest {
	s.NotificationWebHook = &v
	return s
}

func (s *UpdateStoreConfigRequest) SetStoreId(v string) *UpdateStoreConfigRequest {
	s.StoreId = &v
	return s
}

func (s *UpdateStoreConfigRequest) SetSubscribeContents(v string) *UpdateStoreConfigRequest {
	s.SubscribeContents = &v
	return s
}

type UpdateStoreConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateStoreConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoreConfigResponseBody) SetCode(v string) *UpdateStoreConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetDynamicCode(v string) *UpdateStoreConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetDynamicMessage(v string) *UpdateStoreConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetErrorCode(v string) *UpdateStoreConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetErrorMessage(v string) *UpdateStoreConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetMessage(v string) *UpdateStoreConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetRequestId(v string) *UpdateStoreConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStoreConfigResponseBody) SetSuccess(v bool) *UpdateStoreConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateStoreConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStoreConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStoreConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoreConfigResponse) SetHeaders(v map[string]*string) *UpdateStoreConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoreConfigResponse) SetStatusCode(v int32) *UpdateStoreConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStoreConfigResponse) SetBody(v *UpdateStoreConfigResponseBody) *UpdateStoreConfigResponse {
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
		"ap-northeast-1":              tea.String("cloudesl.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cloudesl.aliyuncs.com"),
		"ap-south-1":                  tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-1":              tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-2":              tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing":                  tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cloudesl.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cloudesl.aliyuncs.com"),
		"cn-fujian":                   tea.String("cloudesl.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cloudesl.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cloudesl.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cloudesl.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cloudesl.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cloudesl.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cloudesl.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cloudesl.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cloudesl.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cloudesl.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cloudesl.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cloudesl.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cloudesl.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cloudesl.aliyuncs.com"),
		"eu-central-1":                tea.String("cloudesl.aliyuncs.com"),
		"eu-west-1":                   tea.String("cloudesl.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cloudesl.aliyuncs.com"),
		"me-east-1":                   tea.String("cloudesl.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cloudesl.aliyuncs.com"),
		"us-east-1":                   tea.String("cloudesl.aliyuncs.com"),
		"us-west-1":                   tea.String("cloudesl.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudesl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ActivateApDeviceWithOptions(request *ActivateApDeviceRequest, runtime *util.RuntimeOptions) (_result *ActivateApDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateApDevice"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateApDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateApDevice(request *ActivateApDeviceRequest) (_result *ActivateApDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateApDeviceResponse{}
	_body, _err := client.ActivateApDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddApDeviceWithOptions(request *AddApDeviceRequest, runtime *util.RuntimeOptions) (_result *AddApDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddApDevice"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddApDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddApDevice(request *AddApDeviceRequest) (_result *AddApDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddApDeviceResponse{}
	_body, _err := client.AddApDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCompanyTemplateWithOptions(request *AddCompanyTemplateRequest, runtime *util.RuntimeOptions) (_result *AddCompanyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.EslSize)) {
		body["EslSize"] = request.EslSize
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IfDefault)) {
		body["IfDefault"] = request.IfDefault
	}

	if !tea.BoolValue(util.IsUnset(request.IfMember)) {
		body["IfMember"] = request.IfMember
	}

	if !tea.BoolValue(util.IsUnset(request.IfOutOfInventory)) {
		body["IfOutOfInventory"] = request.IfOutOfInventory
	}

	if !tea.BoolValue(util.IsUnset(request.IfPromotion)) {
		body["IfPromotion"] = request.IfPromotion
	}

	if !tea.BoolValue(util.IsUnset(request.IfSourceCode)) {
		body["IfSourceCode"] = request.IfSourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Layout)) {
		body["Layout"] = request.Layout
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateSceneId)) {
		body["TemplateSceneId"] = request.TemplateSceneId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		body["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCompanyTemplate"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCompanyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCompanyTemplate(request *AddCompanyTemplateRequest) (_result *AddCompanyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCompanyTemplateResponse{}
	_body, _err := client.AddCompanyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserWithOptions(request *AddUserRequest, runtime *util.RuntimeOptions) (_result *AddUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUser"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUser(request *AddUserRequest) (_result *AddUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserResponse{}
	_body, _err := client.AddUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyCompanyTemplateVersionToStoresWithOptions(request *ApplyCompanyTemplateVersionToStoresRequest, runtime *util.RuntimeOptions) (_result *ApplyCompanyTemplateVersionToStoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Stores)) {
		body["Stores"] = request.Stores
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyCompanyTemplateVersionToStores"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyCompanyTemplateVersionToStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyCompanyTemplateVersionToStores(request *ApplyCompanyTemplateVersionToStoresRequest) (_result *ApplyCompanyTemplateVersionToStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyCompanyTemplateVersionToStoresResponse{}
	_body, _err := client.ApplyCompanyTemplateVersionToStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignUserWithOptions(request *AssignUserRequest, runtime *util.RuntimeOptions) (_result *AssignUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.Stores)) {
		body["Stores"] = request.Stores
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignUser"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignUser(request *AssignUserRequest) (_result *AssignUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignUserResponse{}
	_body, _err := client.AssignUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchInsertItemsWithOptions(request *BatchInsertItemsRequest, runtime *util.RuntimeOptions) (_result *BatchInsertItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.ItemInfo)) {
		body["ItemInfo"] = request.ItemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncByItemId)) {
		body["SyncByItemId"] = request.SyncByItemId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchInsertItems"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchInsertItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchInsertItems(request *BatchInsertItemsRequest) (_result *BatchInsertItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchInsertItemsResponse{}
	_body, _err := client.BatchInsertItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindEslDeviceWithOptions(request *BindEslDeviceRequest, runtime *util.RuntimeOptions) (_result *BindEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Column)) {
		body["Column"] = request.Column
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerId)) {
		body["ContainerId"] = request.ContainerId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		body["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.Layer)) {
		body["Layer"] = request.Layer
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutId)) {
		body["LayoutId"] = request.LayoutId
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutName)) {
		body["LayoutName"] = request.LayoutName
	}

	if !tea.BoolValue(util.IsUnset(request.Shelf)) {
		body["Shelf"] = request.Shelf
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindEslDevice"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindEslDevice(request *BindEslDeviceRequest) (_result *BindEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindEslDeviceResponse{}
	_body, _err := client.BindEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStoreWithOptions(request *CreateStoreRequest, runtime *util.RuntimeOptions) (_result *CreateStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoUnbindDays)) {
		body["AutoUnbindDays"] = request.AutoUnbindDays
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUnbindOfflineEsl)) {
		body["AutoUnbindOfflineEsl"] = request.AutoUnbindOfflineEsl
	}

	if !tea.BoolValue(util.IsUnset(request.BarCodeEncode)) {
		body["BarCodeEncode"] = request.BarCodeEncode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.StoreName)) {
		body["StoreName"] = request.StoreName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["TimeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.UserStoreCode)) {
		body["UserStoreCode"] = request.UserStoreCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStore"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStore(request *CreateStoreRequest) (_result *CreateStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStoreResponse{}
	_body, _err := client.CreateStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApDeviceWithOptions(request *DeleteApDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteApDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApDevice"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApDevice(request *DeleteApDeviceRequest) (_result *DeleteApDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApDeviceResponse{}
	_body, _err := client.DeleteApDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCompanyTemplateWithOptions(request *DeleteCompanyTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteCompanyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCompanyTemplate"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCompanyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCompanyTemplate(request *DeleteCompanyTemplateRequest) (_result *DeleteCompanyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCompanyTemplateResponse{}
	_body, _err := client.DeleteCompanyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteItemWithOptions(request *DeleteItemRequest, runtime *util.RuntimeOptions) (_result *DeleteItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.UnbindEslDevice)) {
		body["UnbindEslDevice"] = request.UnbindEslDevice
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteItem"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteItem(request *DeleteItemRequest) (_result *DeleteItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteItemResponse{}
	_body, _err := client.DeleteItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStoreWithOptions(request *DeleteStoreRequest, runtime *util.RuntimeOptions) (_result *DeleteStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStore"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStore(request *DeleteStoreRequest) (_result *DeleteStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStoreResponse{}
	_body, _err := client.DeleteStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApDevicesWithOptions(request *DescribeApDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeApDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.BeActivate)) {
		body["BeActivate"] = request.BeActivate
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApDevices"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApDevices(request *DescribeApDevicesRequest) (_result *DescribeApDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApDevicesResponse{}
	_body, _err := client.DescribeApDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableEslModelsWithOptions(request *DescribeAvailableEslModelsRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableEslModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
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
		Action:      tea.String("DescribeAvailableEslModels"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableEslModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableEslModels(request *DescribeAvailableEslModelsRequest) (_result *DescribeAvailableEslModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableEslModelsResponse{}
	_body, _err := client.DescribeAvailableEslModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBindersWithOptions(request *DescribeBindersRequest, runtime *util.RuntimeOptions) (_result *DescribeBindersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTitle)) {
		body["ItemTitle"] = request.ItemTitle
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBinders"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBindersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBinders(request *DescribeBindersRequest) (_result *DescribeBindersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBindersResponse{}
	_body, _err := client.DescribeBindersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCompanyTemplateVersionsWithOptions(request *DescribeCompanyTemplateVersionsRequest, runtime *util.RuntimeOptions) (_result *DescribeCompanyTemplateVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("DescribeCompanyTemplateVersions"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCompanyTemplateVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCompanyTemplateVersions(request *DescribeCompanyTemplateVersionsRequest) (_result *DescribeCompanyTemplateVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCompanyTemplateVersionsResponse{}
	_body, _err := client.DescribeCompanyTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEslDeviceWithOptions(request *DescribeEslDeviceRequest, runtime *util.RuntimeOptions) (_result *DescribeEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEslDevice"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEslDevice(request *DescribeEslDeviceRequest) (_result *DescribeEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEslDeviceResponse{}
	_body, _err := client.DescribeEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEslDevicesWithOptions(request *DescribeEslDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeEslDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.EslStatus)) {
		body["EslStatus"] = request.EslStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.FromBatteryLevel)) {
		body["FromBatteryLevel"] = request.FromBatteryLevel
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.ToBatteryLevel)) {
		body["ToBatteryLevel"] = request.ToBatteryLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.TypeEncode)) {
		body["TypeEncode"] = request.TypeEncode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEslDevices"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEslDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEslDevices(request *DescribeEslDevicesRequest) (_result *DescribeEslDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEslDevicesResponse{}
	_body, _err := client.DescribeEslDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEslModelByTemplateVersionWithOptions(request *DescribeEslModelByTemplateVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeEslModelByTemplateVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEslModelByTemplateVersion"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEslModelByTemplateVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEslModelByTemplateVersion(request *DescribeEslModelByTemplateVersionRequest) (_result *DescribeEslModelByTemplateVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEslModelByTemplateVersionResponse{}
	_body, _err := client.DescribeEslModelByTemplateVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeItemsWithOptions(request *DescribeItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BePromotion)) {
		body["BePromotion"] = request.BePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		body["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTitle)) {
		body["ItemTitle"] = request.ItemTitle
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SkuId)) {
		body["SkuId"] = request.SkuId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeItems"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeItems(request *DescribeItemsRequest) (_result *DescribeItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeItemsResponse{}
	_body, _err := client.DescribeItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStoreByTemplateVersionWithOptions(request *DescribeStoreByTemplateVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeStoreByTemplateVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStoreByTemplateVersion"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStoreByTemplateVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStoreByTemplateVersion(request *DescribeStoreByTemplateVersionRequest) (_result *DescribeStoreByTemplateVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStoreByTemplateVersionResponse{}
	_body, _err := client.DescribeStoreByTemplateVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStoreConfigWithOptions(request *DescribeStoreConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeStoreConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStoreConfig"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStoreConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStoreConfig(request *DescribeStoreConfigRequest) (_result *DescribeStoreConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStoreConfigResponse{}
	_body, _err := client.DescribeStoreConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStoresWithOptions(request *DescribeStoresRequest, runtime *util.RuntimeOptions) (_result *DescribeStoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreName)) {
		body["StoreName"] = request.StoreName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserStoreCode)) {
		body["UserStoreCode"] = request.UserStoreCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStores"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStores(request *DescribeStoresRequest) (_result *DescribeStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStoresResponse{}
	_body, _err := client.DescribeStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateByModelWithOptions(request *DescribeTemplateByModelRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateByModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.EslSize)) {
		body["EslSize"] = request.EslSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplateByModel"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplateByModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplateByModel(request *DescribeTemplateByModelRequest) (_result *DescribeTemplateByModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateByModelResponse{}
	_body, _err := client.DescribeTemplateByModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserLogWithOptions(request *DescribeUserLogRequest, runtime *util.RuntimeOptions) (_result *DescribeUserLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemShortTitle)) {
		body["ItemShortTitle"] = request.ItemShortTitle
	}

	if !tea.BoolValue(util.IsUnset(request.LogId)) {
		body["LogId"] = request.LogId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationStatus)) {
		body["OperationStatus"] = request.OperationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		body["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserLog"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserLog(request *DescribeUserLogRequest) (_result *DescribeUserLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserLogResponse{}
	_body, _err := client.DescribeUserLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsersWithOptions(request *DescribeUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsers"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsers(request *DescribeUsersRequest) (_result *DescribeUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersResponse{}
	_body, _err := client.DescribeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTemplateListByGroupIdWithOptions(request *QueryTemplateListByGroupIdRequest, runtime *util.RuntimeOptions) (_result *QueryTemplateListByGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
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
		Action:      tea.String("QueryTemplateListByGroupId"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTemplateListByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTemplateListByGroupId(request *QueryTemplateListByGroupIdRequest) (_result *QueryTemplateListByGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTemplateListByGroupIdResponse{}
	_body, _err := client.QueryTemplateListByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncAddMaterialWithOptions(request *SyncAddMaterialRequest, runtime *util.RuntimeOptions) (_result *SyncAddMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncAddMaterial"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncAddMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncAddMaterial(request *SyncAddMaterialRequest) (_result *SyncAddMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncAddMaterialResponse{}
	_body, _err := client.SyncAddMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnassignUserWithOptions(request *UnassignUserRequest, runtime *util.RuntimeOptions) (_result *UnassignUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnassignUser"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnassignUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnassignUser(request *UnassignUserRequest) (_result *UnassignUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassignUserResponse{}
	_body, _err := client.UnassignUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindEslDeviceWithOptions(request *UnbindEslDeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Column)) {
		body["Column"] = request.Column
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		body["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.Layer)) {
		body["Layer"] = request.Layer
	}

	if !tea.BoolValue(util.IsUnset(request.Shelf)) {
		body["Shelf"] = request.Shelf
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEslDevice"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindEslDevice(request *UnbindEslDeviceRequest) (_result *UnbindEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindEslDeviceResponse{}
	_body, _err := client.UnbindEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEslDeviceLightWithOptions(request *UpdateEslDeviceLightRequest, runtime *util.RuntimeOptions) (_result *UpdateEslDeviceLightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		body["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.LedColor)) {
		body["LedColor"] = request.LedColor
	}

	if !tea.BoolValue(util.IsUnset(request.LightUpTime)) {
		body["LightUpTime"] = request.LightUpTime
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEslDeviceLight"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEslDeviceLightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEslDeviceLight(request *UpdateEslDeviceLightRequest) (_result *UpdateEslDeviceLightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEslDeviceLightResponse{}
	_body, _err := client.UpdateEslDeviceLightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStoreWithOptions(request *UpdateStoreRequest, runtime *util.RuntimeOptions) (_result *UpdateStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoUnbindDays)) {
		body["AutoUnbindDays"] = request.AutoUnbindDays
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUnbindOfflineEsl)) {
		body["AutoUnbindOfflineEsl"] = request.AutoUnbindOfflineEsl
	}

	if !tea.BoolValue(util.IsUnset(request.BarCodeEncode)) {
		body["BarCodeEncode"] = request.BarCodeEncode
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreName)) {
		body["StoreName"] = request.StoreName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		body["Timezone"] = request.Timezone
	}

	if !tea.BoolValue(util.IsUnset(request.UserStoreCode)) {
		body["UserStoreCode"] = request.UserStoreCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStore"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStore(request *UpdateStoreRequest) (_result *UpdateStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStoreResponse{}
	_body, _err := client.UpdateStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStoreConfigWithOptions(request *UpdateStoreConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateStoreConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableNotification)) {
		body["EnableNotification"] = request.EnableNotification
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraParams)) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationSilentTimes)) {
		body["NotificationSilentTimes"] = request.NotificationSilentTimes
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationWebHook)) {
		body["NotificationWebHook"] = request.NotificationWebHook
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeContents)) {
		body["SubscribeContents"] = request.SubscribeContents
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStoreConfig"),
		Version:     tea.String("2020-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStoreConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStoreConfig(request *UpdateStoreConfigRequest) (_result *UpdateStoreConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStoreConfigResponse{}
	_body, _err := client.UpdateStoreConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

type DoIotChgBindOrUnBindRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imei                 *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	MidChannelId         *string `json:"MidChannelId,omitempty" xml:"MidChannelId,omitempty"`
	NewImei              *string `json:"NewImei,omitempty" xml:"NewImei,omitempty"`
	OpionType            *string `json:"OpionType,omitempty" xml:"OpionType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotChgBindOrUnBindRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotChgBindOrUnBindRequest) GoString() string {
	return s.String()
}

func (s *DoIotChgBindOrUnBindRequest) SetIccid(v string) *DoIotChgBindOrUnBindRequest {
	s.Iccid = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetImei(v string) *DoIotChgBindOrUnBindRequest {
	s.Imei = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetMidChannelId(v string) *DoIotChgBindOrUnBindRequest {
	s.MidChannelId = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetNewImei(v string) *DoIotChgBindOrUnBindRequest {
	s.NewImei = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetOpionType(v string) *DoIotChgBindOrUnBindRequest {
	s.OpionType = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetOwnerId(v int64) *DoIotChgBindOrUnBindRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetResourceOwnerAccount(v string) *DoIotChgBindOrUnBindRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotChgBindOrUnBindRequest) SetResourceOwnerId(v int64) *DoIotChgBindOrUnBindRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotChgBindOrUnBindResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	IotModBind *DoIotChgBindOrUnBindResponseBodyIotModBind `json:"IotModBind,omitempty" xml:"IotModBind,omitempty" type:"Struct"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoIotChgBindOrUnBindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotChgBindOrUnBindResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotChgBindOrUnBindResponseBody) SetCode(v string) *DoIotChgBindOrUnBindResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotChgBindOrUnBindResponseBody) SetIotModBind(v *DoIotChgBindOrUnBindResponseBodyIotModBind) *DoIotChgBindOrUnBindResponseBody {
	s.IotModBind = v
	return s
}

func (s *DoIotChgBindOrUnBindResponseBody) SetMessage(v string) *DoIotChgBindOrUnBindResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotChgBindOrUnBindResponseBody) SetRequestId(v string) *DoIotChgBindOrUnBindResponseBody {
	s.RequestId = &v
	return s
}

type DoIotChgBindOrUnBindResponseBodyIotModBind struct {
	IsModSuccess *bool `json:"IsModSuccess,omitempty" xml:"IsModSuccess,omitempty"`
}

func (s DoIotChgBindOrUnBindResponseBodyIotModBind) String() string {
	return tea.Prettify(s)
}

func (s DoIotChgBindOrUnBindResponseBodyIotModBind) GoString() string {
	return s.String()
}

func (s *DoIotChgBindOrUnBindResponseBodyIotModBind) SetIsModSuccess(v bool) *DoIotChgBindOrUnBindResponseBodyIotModBind {
	s.IsModSuccess = &v
	return s
}

type DoIotChgBindOrUnBindResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotChgBindOrUnBindResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotChgBindOrUnBindResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotChgBindOrUnBindResponse) GoString() string {
	return s.String()
}

func (s *DoIotChgBindOrUnBindResponse) SetHeaders(v map[string]*string) *DoIotChgBindOrUnBindResponse {
	s.Headers = v
	return s
}

func (s *DoIotChgBindOrUnBindResponse) SetStatusCode(v int32) *DoIotChgBindOrUnBindResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotChgBindOrUnBindResponse) SetBody(v *DoIotChgBindOrUnBindResponseBody) *DoIotChgBindOrUnBindResponse {
	s.Body = v
	return s
}

type DoIotIsImeiExistRequest struct {
	Imei                 *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotIsImeiExistRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotIsImeiExistRequest) GoString() string {
	return s.String()
}

func (s *DoIotIsImeiExistRequest) SetImei(v string) *DoIotIsImeiExistRequest {
	s.Imei = &v
	return s
}

func (s *DoIotIsImeiExistRequest) SetOwnerId(v int64) *DoIotIsImeiExistRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotIsImeiExistRequest) SetResourceOwnerAccount(v string) *DoIotIsImeiExistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotIsImeiExistRequest) SetResourceOwnerId(v int64) *DoIotIsImeiExistRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotIsImeiExistResponseBody struct {
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	IotImeiExist *DoIotIsImeiExistResponseBodyIotImeiExist `json:"IotImeiExist,omitempty" xml:"IotImeiExist,omitempty" type:"Struct"`
	Message      *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoIotIsImeiExistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotIsImeiExistResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotIsImeiExistResponseBody) SetCode(v string) *DoIotIsImeiExistResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotIsImeiExistResponseBody) SetIotImeiExist(v *DoIotIsImeiExistResponseBodyIotImeiExist) *DoIotIsImeiExistResponseBody {
	s.IotImeiExist = v
	return s
}

func (s *DoIotIsImeiExistResponseBody) SetMessage(v string) *DoIotIsImeiExistResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotIsImeiExistResponseBody) SetRequestId(v string) *DoIotIsImeiExistResponseBody {
	s.RequestId = &v
	return s
}

type DoIotIsImeiExistResponseBodyIotImeiExist struct {
	IsImeiExist *bool `json:"IsImeiExist,omitempty" xml:"IsImeiExist,omitempty"`
}

func (s DoIotIsImeiExistResponseBodyIotImeiExist) String() string {
	return tea.Prettify(s)
}

func (s DoIotIsImeiExistResponseBodyIotImeiExist) GoString() string {
	return s.String()
}

func (s *DoIotIsImeiExistResponseBodyIotImeiExist) SetIsImeiExist(v bool) *DoIotIsImeiExistResponseBodyIotImeiExist {
	s.IsImeiExist = &v
	return s
}

type DoIotIsImeiExistResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotIsImeiExistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotIsImeiExistResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotIsImeiExistResponse) GoString() string {
	return s.String()
}

func (s *DoIotIsImeiExistResponse) SetHeaders(v map[string]*string) *DoIotIsImeiExistResponse {
	s.Headers = v
	return s
}

func (s *DoIotIsImeiExistResponse) SetStatusCode(v int32) *DoIotIsImeiExistResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotIsImeiExistResponse) SetBody(v *DoIotIsImeiExistResponseBody) *DoIotIsImeiExistResponse {
	s.Body = v
	return s
}

type DoIotPostImeiRequest struct {
	Comments             *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	DeviceType           *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Imei                 *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotPostImeiRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotPostImeiRequest) GoString() string {
	return s.String()
}

func (s *DoIotPostImeiRequest) SetComments(v string) *DoIotPostImeiRequest {
	s.Comments = &v
	return s
}

func (s *DoIotPostImeiRequest) SetDeviceType(v string) *DoIotPostImeiRequest {
	s.DeviceType = &v
	return s
}

func (s *DoIotPostImeiRequest) SetImei(v string) *DoIotPostImeiRequest {
	s.Imei = &v
	return s
}

func (s *DoIotPostImeiRequest) SetOwnerId(v int64) *DoIotPostImeiRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotPostImeiRequest) SetResourceOwnerAccount(v string) *DoIotPostImeiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotPostImeiRequest) SetResourceOwnerId(v int64) *DoIotPostImeiRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotPostImeiResponseBody struct {
	Code        *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	IotPostImei *DoIotPostImeiResponseBodyIotPostImei `json:"IotPostImei,omitempty" xml:"IotPostImei,omitempty" type:"Struct"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoIotPostImeiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotPostImeiResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotPostImeiResponseBody) SetCode(v string) *DoIotPostImeiResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotPostImeiResponseBody) SetIotPostImei(v *DoIotPostImeiResponseBodyIotPostImei) *DoIotPostImeiResponseBody {
	s.IotPostImei = v
	return s
}

func (s *DoIotPostImeiResponseBody) SetMessage(v string) *DoIotPostImeiResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotPostImeiResponseBody) SetRequestId(v string) *DoIotPostImeiResponseBody {
	s.RequestId = &v
	return s
}

type DoIotPostImeiResponseBodyIotPostImei struct {
	IsPostSuccess *bool `json:"IsPostSuccess,omitempty" xml:"IsPostSuccess,omitempty"`
}

func (s DoIotPostImeiResponseBodyIotPostImei) String() string {
	return tea.Prettify(s)
}

func (s DoIotPostImeiResponseBodyIotPostImei) GoString() string {
	return s.String()
}

func (s *DoIotPostImeiResponseBodyIotPostImei) SetIsPostSuccess(v bool) *DoIotPostImeiResponseBodyIotPostImei {
	s.IsPostSuccess = &v
	return s
}

type DoIotPostImeiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotPostImeiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotPostImeiResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotPostImeiResponse) GoString() string {
	return s.String()
}

func (s *DoIotPostImeiResponse) SetHeaders(v map[string]*string) *DoIotPostImeiResponse {
	s.Headers = v
	return s
}

func (s *DoIotPostImeiResponse) SetStatusCode(v int32) *DoIotPostImeiResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotPostImeiResponse) SetBody(v *DoIotPostImeiResponseBody) *DoIotPostImeiResponse {
	s.Body = v
	return s
}

type DoIotSetAbsoluteRemindConfigRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ConfigInfo           *string `json:"ConfigInfo,omitempty" xml:"ConfigInfo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotSetAbsoluteRemindConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotSetAbsoluteRemindConfigRequest) GoString() string {
	return s.String()
}

func (s *DoIotSetAbsoluteRemindConfigRequest) SetBizId(v string) *DoIotSetAbsoluteRemindConfigRequest {
	s.BizId = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigRequest) SetBizType(v string) *DoIotSetAbsoluteRemindConfigRequest {
	s.BizType = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigRequest) SetConfigInfo(v string) *DoIotSetAbsoluteRemindConfigRequest {
	s.ConfigInfo = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigRequest) SetOwnerId(v int64) *DoIotSetAbsoluteRemindConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigRequest) SetResourceOwnerAccount(v string) *DoIotSetAbsoluteRemindConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigRequest) SetResourceOwnerId(v int64) *DoIotSetAbsoluteRemindConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotSetAbsoluteRemindConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoIotSetAbsoluteRemindConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotSetAbsoluteRemindConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotSetAbsoluteRemindConfigResponseBody) SetCode(v string) *DoIotSetAbsoluteRemindConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigResponseBody) SetMessage(v string) *DoIotSetAbsoluteRemindConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigResponseBody) SetRequestId(v string) *DoIotSetAbsoluteRemindConfigResponseBody {
	s.RequestId = &v
	return s
}

type DoIotSetAbsoluteRemindConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotSetAbsoluteRemindConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotSetAbsoluteRemindConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotSetAbsoluteRemindConfigResponse) GoString() string {
	return s.String()
}

func (s *DoIotSetAbsoluteRemindConfigResponse) SetHeaders(v map[string]*string) *DoIotSetAbsoluteRemindConfigResponse {
	s.Headers = v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigResponse) SetStatusCode(v int32) *DoIotSetAbsoluteRemindConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotSetAbsoluteRemindConfigResponse) SetBody(v *DoIotSetAbsoluteRemindConfigResponseBody) *DoIotSetAbsoluteRemindConfigResponse {
	s.Body = v
	return s
}

type DoIotSetRemindConfigRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ConfigInfo           *string `json:"ConfigInfo,omitempty" xml:"ConfigInfo,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotSetRemindConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotSetRemindConfigRequest) GoString() string {
	return s.String()
}

func (s *DoIotSetRemindConfigRequest) SetBizId(v string) *DoIotSetRemindConfigRequest {
	s.BizId = &v
	return s
}

func (s *DoIotSetRemindConfigRequest) SetBizType(v string) *DoIotSetRemindConfigRequest {
	s.BizType = &v
	return s
}

func (s *DoIotSetRemindConfigRequest) SetConfigInfo(v string) *DoIotSetRemindConfigRequest {
	s.ConfigInfo = &v
	return s
}

func (s *DoIotSetRemindConfigRequest) SetOperationType(v string) *DoIotSetRemindConfigRequest {
	s.OperationType = &v
	return s
}

func (s *DoIotSetRemindConfigRequest) SetOwnerId(v int64) *DoIotSetRemindConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotSetRemindConfigRequest) SetResourceOwnerAccount(v string) *DoIotSetRemindConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotSetRemindConfigRequest) SetResourceOwnerId(v int64) *DoIotSetRemindConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotSetRemindConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoIotSetRemindConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotSetRemindConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotSetRemindConfigResponseBody) SetCode(v string) *DoIotSetRemindConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotSetRemindConfigResponseBody) SetMessage(v string) *DoIotSetRemindConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotSetRemindConfigResponseBody) SetRequestId(v string) *DoIotSetRemindConfigResponseBody {
	s.RequestId = &v
	return s
}

type DoIotSetRemindConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotSetRemindConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotSetRemindConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotSetRemindConfigResponse) GoString() string {
	return s.String()
}

func (s *DoIotSetRemindConfigResponse) SetHeaders(v map[string]*string) *DoIotSetRemindConfigResponse {
	s.Headers = v
	return s
}

func (s *DoIotSetRemindConfigResponse) SetStatusCode(v int32) *DoIotSetRemindConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotSetRemindConfigResponse) SetBody(v *DoIotSetRemindConfigResponseBody) *DoIotSetRemindConfigResponse {
	s.Body = v
	return s
}

type DoIotUnbindResumeRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotUnbindResumeRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotUnbindResumeRequest) GoString() string {
	return s.String()
}

func (s *DoIotUnbindResumeRequest) SetIccid(v string) *DoIotUnbindResumeRequest {
	s.Iccid = &v
	return s
}

func (s *DoIotUnbindResumeRequest) SetOwnerId(v int64) *DoIotUnbindResumeRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotUnbindResumeRequest) SetResourceOwnerAccount(v string) *DoIotUnbindResumeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotUnbindResumeRequest) SetResourceOwnerId(v int64) *DoIotUnbindResumeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotUnbindResumeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoIotUnbindResumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotUnbindResumeResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotUnbindResumeResponseBody) SetCode(v string) *DoIotUnbindResumeResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotUnbindResumeResponseBody) SetData(v bool) *DoIotUnbindResumeResponseBody {
	s.Data = &v
	return s
}

func (s *DoIotUnbindResumeResponseBody) SetMessage(v string) *DoIotUnbindResumeResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotUnbindResumeResponseBody) SetRequestId(v string) *DoIotUnbindResumeResponseBody {
	s.RequestId = &v
	return s
}

type DoIotUnbindResumeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotUnbindResumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotUnbindResumeResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotUnbindResumeResponse) GoString() string {
	return s.String()
}

func (s *DoIotUnbindResumeResponse) SetHeaders(v map[string]*string) *DoIotUnbindResumeResponse {
	s.Headers = v
	return s
}

func (s *DoIotUnbindResumeResponse) SetStatusCode(v int32) *DoIotUnbindResumeResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotUnbindResumeResponse) SetBody(v *DoIotUnbindResumeResponseBody) *DoIotUnbindResumeResponse {
	s.Body = v
	return s
}

type DoIotUserStopResumeRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptionType           *string `json:"OptionType,omitempty" xml:"OptionType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DoIotUserStopResumeRequest) String() string {
	return tea.Prettify(s)
}

func (s DoIotUserStopResumeRequest) GoString() string {
	return s.String()
}

func (s *DoIotUserStopResumeRequest) SetIccid(v string) *DoIotUserStopResumeRequest {
	s.Iccid = &v
	return s
}

func (s *DoIotUserStopResumeRequest) SetOptionType(v string) *DoIotUserStopResumeRequest {
	s.OptionType = &v
	return s
}

func (s *DoIotUserStopResumeRequest) SetOwnerId(v int64) *DoIotUserStopResumeRequest {
	s.OwnerId = &v
	return s
}

func (s *DoIotUserStopResumeRequest) SetResourceOwnerAccount(v string) *DoIotUserStopResumeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoIotUserStopResumeRequest) SetResourceOwnerId(v int64) *DoIotUserStopResumeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DoIotUserStopResumeResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DoIotUserStopResumeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DoIotUserStopResumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoIotUserStopResumeResponseBody) GoString() string {
	return s.String()
}

func (s *DoIotUserStopResumeResponseBody) SetCode(v string) *DoIotUserStopResumeResponseBody {
	s.Code = &v
	return s
}

func (s *DoIotUserStopResumeResponseBody) SetMessage(v string) *DoIotUserStopResumeResponseBody {
	s.Message = &v
	return s
}

func (s *DoIotUserStopResumeResponseBody) SetRequestId(v string) *DoIotUserStopResumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoIotUserStopResumeResponseBody) SetResult(v *DoIotUserStopResumeResponseBodyResult) *DoIotUserStopResumeResponseBody {
	s.Result = v
	return s
}

type DoIotUserStopResumeResponseBodyResult struct {
	ControlResult *bool `json:"ControlResult,omitempty" xml:"ControlResult,omitempty"`
}

func (s DoIotUserStopResumeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DoIotUserStopResumeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DoIotUserStopResumeResponseBodyResult) SetControlResult(v bool) *DoIotUserStopResumeResponseBodyResult {
	s.ControlResult = &v
	return s
}

type DoIotUserStopResumeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoIotUserStopResumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoIotUserStopResumeResponse) String() string {
	return tea.Prettify(s)
}

func (s DoIotUserStopResumeResponse) GoString() string {
	return s.String()
}

func (s *DoIotUserStopResumeResponse) SetHeaders(v map[string]*string) *DoIotUserStopResumeResponse {
	s.Headers = v
	return s
}

func (s *DoIotUserStopResumeResponse) SetStatusCode(v int32) *DoIotUserStopResumeResponse {
	s.StatusCode = &v
	return s
}

func (s *DoIotUserStopResumeResponse) SetBody(v *DoIotUserStopResumeResponseBody) *DoIotUserStopResumeResponse {
	s.Body = v
	return s
}

type DoSendIotSmsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam        *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s DoSendIotSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DoSendIotSmsRequest) GoString() string {
	return s.String()
}

func (s *DoSendIotSmsRequest) SetOwnerId(v int64) *DoSendIotSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *DoSendIotSmsRequest) SetPhoneNumbers(v string) *DoSendIotSmsRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *DoSendIotSmsRequest) SetResourceOwnerAccount(v string) *DoSendIotSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoSendIotSmsRequest) SetResourceOwnerId(v int64) *DoSendIotSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DoSendIotSmsRequest) SetSignName(v string) *DoSendIotSmsRequest {
	s.SignName = &v
	return s
}

func (s *DoSendIotSmsRequest) SetTemplateCode(v string) *DoSendIotSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *DoSendIotSmsRequest) SetTemplateParam(v string) *DoSendIotSmsRequest {
	s.TemplateParam = &v
	return s
}

type DoSendIotSmsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DoSendIotSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoSendIotSmsResponseBody) GoString() string {
	return s.String()
}

func (s *DoSendIotSmsResponseBody) SetCode(v string) *DoSendIotSmsResponseBody {
	s.Code = &v
	return s
}

func (s *DoSendIotSmsResponseBody) SetMessage(v string) *DoSendIotSmsResponseBody {
	s.Message = &v
	return s
}

func (s *DoSendIotSmsResponseBody) SetModule(v string) *DoSendIotSmsResponseBody {
	s.Module = &v
	return s
}

func (s *DoSendIotSmsResponseBody) SetRequestId(v string) *DoSendIotSmsResponseBody {
	s.RequestId = &v
	return s
}

type DoSendIotSmsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoSendIotSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoSendIotSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DoSendIotSmsResponse) GoString() string {
	return s.String()
}

func (s *DoSendIotSmsResponse) SetHeaders(v map[string]*string) *DoSendIotSmsResponse {
	s.Headers = v
	return s
}

func (s *DoSendIotSmsResponse) SetStatusCode(v int32) *DoSendIotSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DoSendIotSmsResponse) SetBody(v *DoSendIotSmsResponseBody) *DoSendIotSmsResponse {
	s.Body = v
	return s
}

type QueryCardFlowInfoRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCardFlowInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCardFlowInfoRequest) SetIccid(v string) *QueryCardFlowInfoRequest {
	s.Iccid = &v
	return s
}

func (s *QueryCardFlowInfoRequest) SetOwnerId(v int64) *QueryCardFlowInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCardFlowInfoRequest) SetResourceOwnerAccount(v string) *QueryCardFlowInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCardFlowInfoRequest) SetResourceOwnerId(v int64) *QueryCardFlowInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCardFlowInfoResponseBody struct {
	CardFlowInfos *QueryCardFlowInfoResponseBodyCardFlowInfos `json:"CardFlowInfos,omitempty" xml:"CardFlowInfos,omitempty" type:"Struct"`
	Code          *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCardFlowInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardFlowInfoResponseBody) SetCardFlowInfos(v *QueryCardFlowInfoResponseBodyCardFlowInfos) *QueryCardFlowInfoResponseBody {
	s.CardFlowInfos = v
	return s
}

func (s *QueryCardFlowInfoResponseBody) SetCode(v string) *QueryCardFlowInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardFlowInfoResponseBody) SetMessage(v string) *QueryCardFlowInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCardFlowInfoResponseBody) SetRequestId(v string) *QueryCardFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryCardFlowInfoResponseBodyCardFlowInfos struct {
	CardFlowInfo []*QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo `json:"CardFlowInfo,omitempty" xml:"CardFlowInfo,omitempty" type:"Repeated"`
}

func (s QueryCardFlowInfoResponseBodyCardFlowInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryCardFlowInfoResponseBodyCardFlowInfos) GoString() string {
	return s.String()
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfos) SetCardFlowInfo(v []*QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) *QueryCardFlowInfoResponseBodyCardFlowInfos {
	s.CardFlowInfo = v
	return s
}

type QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo struct {
	ExpireDate   *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	FlowResource *int64  `json:"FlowResource,omitempty" xml:"FlowResource,omitempty"`
	FlowUsed     *int64  `json:"FlowUsed,omitempty" xml:"FlowUsed,omitempty"`
	ResName      *string `json:"ResName,omitempty" xml:"ResName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RestOfFlow   *int64  `json:"RestOfFlow,omitempty" xml:"RestOfFlow,omitempty"`
	SmsUsed      *int64  `json:"SmsUsed,omitempty" xml:"SmsUsed,omitempty"`
	ValidDate    *string `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
	VoiceTotal   *int64  `json:"VoiceTotal,omitempty" xml:"VoiceTotal,omitempty"`
	VoiceUsed    *int64  `json:"VoiceUsed,omitempty" xml:"VoiceUsed,omitempty"`
}

func (s QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) GoString() string {
	return s.String()
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetExpireDate(v string) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.ExpireDate = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetFlowResource(v int64) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.FlowResource = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetFlowUsed(v int64) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.FlowUsed = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetResName(v string) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.ResName = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetResourceType(v string) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.ResourceType = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetRestOfFlow(v int64) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.RestOfFlow = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetSmsUsed(v int64) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.SmsUsed = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetValidDate(v string) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.ValidDate = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetVoiceTotal(v int64) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.VoiceTotal = &v
	return s
}

func (s *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo) SetVoiceUsed(v int64) *QueryCardFlowInfoResponseBodyCardFlowInfosCardFlowInfo {
	s.VoiceUsed = &v
	return s
}

type QueryCardFlowInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardFlowInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCardFlowInfoResponse) SetHeaders(v map[string]*string) *QueryCardFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCardFlowInfoResponse) SetStatusCode(v int32) *QueryCardFlowInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardFlowInfoResponse) SetBody(v *QueryCardFlowInfoResponseBody) *QueryCardFlowInfoResponse {
	s.Body = v
	return s
}

type QueryCardHistoryFlowInfoRequest struct {
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryCardHistoryFlowInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardHistoryFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCardHistoryFlowInfoRequest) SetEndTime(v string) *QueryCardHistoryFlowInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCardHistoryFlowInfoRequest) SetIccid(v string) *QueryCardHistoryFlowInfoRequest {
	s.Iccid = &v
	return s
}

func (s *QueryCardHistoryFlowInfoRequest) SetOwnerId(v int64) *QueryCardHistoryFlowInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCardHistoryFlowInfoRequest) SetResourceOwnerAccount(v string) *QueryCardHistoryFlowInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCardHistoryFlowInfoRequest) SetResourceOwnerId(v int64) *QueryCardHistoryFlowInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCardHistoryFlowInfoRequest) SetStartTime(v string) *QueryCardHistoryFlowInfoRequest {
	s.StartTime = &v
	return s
}

type QueryCardHistoryFlowInfoResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryCardHistoryFlowInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCardHistoryFlowInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardHistoryFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardHistoryFlowInfoResponseBody) SetCode(v string) *QueryCardHistoryFlowInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBody) SetData(v []*QueryCardHistoryFlowInfoResponseBodyData) *QueryCardHistoryFlowInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBody) SetMessage(v string) *QueryCardHistoryFlowInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBody) SetRequestId(v string) *QueryCardHistoryFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBody) SetSuccess(v bool) *QueryCardHistoryFlowInfoResponseBody {
	s.Success = &v
	return s
}

type QueryCardHistoryFlowInfoResponseBodyData struct {
	CurValue     *int64                                                  `json:"CurValue,omitempty" xml:"CurValue,omitempty"`
	DayUsageList []*QueryCardHistoryFlowInfoResponseBodyDataDayUsageList `json:"DayUsageList,omitempty" xml:"DayUsageList,omitempty" type:"Repeated"`
	Month        *int64                                                  `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s QueryCardHistoryFlowInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCardHistoryFlowInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCardHistoryFlowInfoResponseBodyData) SetCurValue(v int64) *QueryCardHistoryFlowInfoResponseBodyData {
	s.CurValue = &v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBodyData) SetDayUsageList(v []*QueryCardHistoryFlowInfoResponseBodyDataDayUsageList) *QueryCardHistoryFlowInfoResponseBodyData {
	s.DayUsageList = v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBodyData) SetMonth(v int64) *QueryCardHistoryFlowInfoResponseBodyData {
	s.Month = &v
	return s
}

type QueryCardHistoryFlowInfoResponseBodyDataDayUsageList struct {
	Day   *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryCardHistoryFlowInfoResponseBodyDataDayUsageList) String() string {
	return tea.Prettify(s)
}

func (s QueryCardHistoryFlowInfoResponseBodyDataDayUsageList) GoString() string {
	return s.String()
}

func (s *QueryCardHistoryFlowInfoResponseBodyDataDayUsageList) SetDay(v int32) *QueryCardHistoryFlowInfoResponseBodyDataDayUsageList {
	s.Day = &v
	return s
}

func (s *QueryCardHistoryFlowInfoResponseBodyDataDayUsageList) SetValue(v int64) *QueryCardHistoryFlowInfoResponseBodyDataDayUsageList {
	s.Value = &v
	return s
}

type QueryCardHistoryFlowInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardHistoryFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardHistoryFlowInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardHistoryFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCardHistoryFlowInfoResponse) SetHeaders(v map[string]*string) *QueryCardHistoryFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCardHistoryFlowInfoResponse) SetStatusCode(v int32) *QueryCardHistoryFlowInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardHistoryFlowInfoResponse) SetBody(v *QueryCardHistoryFlowInfoResponseBody) *QueryCardHistoryFlowInfoResponse {
	s.Body = v
	return s
}

type QueryCardInfoRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCardInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCardInfoRequest) SetIccid(v string) *QueryCardInfoRequest {
	s.Iccid = &v
	return s
}

func (s *QueryCardInfoRequest) SetOwnerId(v int64) *QueryCardInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCardInfoRequest) SetResourceOwnerAccount(v string) *QueryCardInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCardInfoRequest) SetResourceOwnerId(v int64) *QueryCardInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCardInfoResponseBody struct {
	CardInfo  *QueryCardInfoResponseBodyCardInfo `json:"CardInfo,omitempty" xml:"CardInfo,omitempty" type:"Struct"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCardInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardInfoResponseBody) SetCardInfo(v *QueryCardInfoResponseBodyCardInfo) *QueryCardInfoResponseBody {
	s.CardInfo = v
	return s
}

func (s *QueryCardInfoResponseBody) SetCode(v string) *QueryCardInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardInfoResponseBody) SetMessage(v string) *QueryCardInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCardInfoResponseBody) SetRequestId(v string) *QueryCardInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryCardInfoResponseBodyCardInfo struct {
	FirstActiveTime *string `json:"FirstActiveTime,omitempty" xml:"FirstActiveTime,omitempty"`
	GprsStatus      *string `json:"GprsStatus,omitempty" xml:"GprsStatus,omitempty"`
	Iccid           *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi            *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	Msisdn          *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	OpenTime        *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	SmsStatus       *string `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	VoiceStatus     *string `json:"VoiceStatus,omitempty" xml:"VoiceStatus,omitempty"`
}

func (s QueryCardInfoResponseBodyCardInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCardInfoResponseBodyCardInfo) GoString() string {
	return s.String()
}

func (s *QueryCardInfoResponseBodyCardInfo) SetFirstActiveTime(v string) *QueryCardInfoResponseBodyCardInfo {
	s.FirstActiveTime = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetGprsStatus(v string) *QueryCardInfoResponseBodyCardInfo {
	s.GprsStatus = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetIccid(v string) *QueryCardInfoResponseBodyCardInfo {
	s.Iccid = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetImsi(v string) *QueryCardInfoResponseBodyCardInfo {
	s.Imsi = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetMsisdn(v string) *QueryCardInfoResponseBodyCardInfo {
	s.Msisdn = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetOpenTime(v string) *QueryCardInfoResponseBodyCardInfo {
	s.OpenTime = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetSmsStatus(v string) *QueryCardInfoResponseBodyCardInfo {
	s.SmsStatus = &v
	return s
}

func (s *QueryCardInfoResponseBodyCardInfo) SetVoiceStatus(v string) *QueryCardInfoResponseBodyCardInfo {
	s.VoiceStatus = &v
	return s
}

type QueryCardInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCardInfoResponse) SetHeaders(v map[string]*string) *QueryCardInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCardInfoResponse) SetStatusCode(v int32) *QueryCardInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardInfoResponse) SetBody(v *QueryCardInfoResponseBody) *QueryCardInfoResponse {
	s.Body = v
	return s
}

type QueryCardStatusRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCardStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryCardStatusRequest) SetIccid(v string) *QueryCardStatusRequest {
	s.Iccid = &v
	return s
}

func (s *QueryCardStatusRequest) SetOwnerId(v int64) *QueryCardStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCardStatusRequest) SetResourceOwnerAccount(v string) *QueryCardStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCardStatusRequest) SetResourceOwnerId(v int64) *QueryCardStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCardStatusResponseBody struct {
	CardStatus *QueryCardStatusResponseBodyCardStatus `json:"CardStatus,omitempty" xml:"CardStatus,omitempty" type:"Struct"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCardStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardStatusResponseBody) SetCardStatus(v *QueryCardStatusResponseBodyCardStatus) *QueryCardStatusResponseBody {
	s.CardStatus = v
	return s
}

func (s *QueryCardStatusResponseBody) SetCode(v string) *QueryCardStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardStatusResponseBody) SetMessage(v string) *QueryCardStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCardStatusResponseBody) SetRequestId(v string) *QueryCardStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryCardStatusResponseBodyCardStatus struct {
	Iccid      *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Msisdn     *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s QueryCardStatusResponseBodyCardStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryCardStatusResponseBodyCardStatus) GoString() string {
	return s.String()
}

func (s *QueryCardStatusResponseBodyCardStatus) SetIccid(v string) *QueryCardStatusResponseBodyCardStatus {
	s.Iccid = &v
	return s
}

func (s *QueryCardStatusResponseBodyCardStatus) SetMsisdn(v string) *QueryCardStatusResponseBodyCardStatus {
	s.Msisdn = &v
	return s
}

func (s *QueryCardStatusResponseBodyCardStatus) SetUserStatus(v string) *QueryCardStatusResponseBodyCardStatus {
	s.UserStatus = &v
	return s
}

type QueryCardStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryCardStatusResponse) SetHeaders(v map[string]*string) *QueryCardStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryCardStatusResponse) SetStatusCode(v int32) *QueryCardStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardStatusResponse) SetBody(v *QueryCardStatusResponseBody) *QueryCardStatusResponse {
	s.Body = v
	return s
}

type QueryCardsInfoRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCardsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardsInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCardsInfoRequest) SetIccid(v string) *QueryCardsInfoRequest {
	s.Iccid = &v
	return s
}

func (s *QueryCardsInfoRequest) SetOwnerId(v int64) *QueryCardsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCardsInfoRequest) SetResourceOwnerAccount(v string) *QueryCardsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCardsInfoRequest) SetResourceOwnerId(v int64) *QueryCardsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCardsInfoResponseBody struct {
	CardsInfo []*QueryCardsInfoResponseBodyCardsInfo `json:"CardsInfo,omitempty" xml:"CardsInfo,omitempty" type:"Repeated"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCardsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardsInfoResponseBody) SetCardsInfo(v []*QueryCardsInfoResponseBodyCardsInfo) *QueryCardsInfoResponseBody {
	s.CardsInfo = v
	return s
}

func (s *QueryCardsInfoResponseBody) SetCode(v string) *QueryCardsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardsInfoResponseBody) SetMessage(v string) *QueryCardsInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCardsInfoResponseBody) SetRequestId(v string) *QueryCardsInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryCardsInfoResponseBodyCardsInfo struct {
	FirstActiveTime *string `json:"FirstActiveTime,omitempty" xml:"FirstActiveTime,omitempty"`
	GprsStatus      *string `json:"GprsStatus,omitempty" xml:"GprsStatus,omitempty"`
	Iccid           *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi            *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	Msisdn          *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	OpenTime        *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	SmsStatus       *string `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	VoiceStatus     *string `json:"VoiceStatus,omitempty" xml:"VoiceStatus,omitempty"`
}

func (s QueryCardsInfoResponseBodyCardsInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCardsInfoResponseBodyCardsInfo) GoString() string {
	return s.String()
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetFirstActiveTime(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.FirstActiveTime = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetGprsStatus(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.GprsStatus = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetIccid(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.Iccid = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetImsi(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.Imsi = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetMsisdn(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.Msisdn = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetOpenTime(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.OpenTime = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetSmsStatus(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.SmsStatus = &v
	return s
}

func (s *QueryCardsInfoResponseBodyCardsInfo) SetVoiceStatus(v string) *QueryCardsInfoResponseBodyCardsInfo {
	s.VoiceStatus = &v
	return s
}

type QueryCardsInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardsInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCardsInfoResponse) SetHeaders(v map[string]*string) *QueryCardsInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCardsInfoResponse) SetStatusCode(v int32) *QueryCardsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardsInfoResponse) SetBody(v *QueryCardsInfoResponseBody) *QueryCardsInfoResponse {
	s.Body = v
	return s
}

type QueryIotCardOfferDtlRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryIotCardOfferDtlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIotCardOfferDtlRequest) GoString() string {
	return s.String()
}

func (s *QueryIotCardOfferDtlRequest) SetIccid(v string) *QueryIotCardOfferDtlRequest {
	s.Iccid = &v
	return s
}

func (s *QueryIotCardOfferDtlRequest) SetOwnerId(v int64) *QueryIotCardOfferDtlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryIotCardOfferDtlRequest) SetResourceOwnerAccount(v string) *QueryIotCardOfferDtlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryIotCardOfferDtlRequest) SetResourceOwnerId(v int64) *QueryIotCardOfferDtlRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryIotCardOfferDtlResponseBody struct {
	CardOfferDetail *QueryIotCardOfferDtlResponseBodyCardOfferDetail `json:"CardOfferDetail,omitempty" xml:"CardOfferDetail,omitempty" type:"Struct"`
	Code            *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryIotCardOfferDtlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIotCardOfferDtlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIotCardOfferDtlResponseBody) SetCardOfferDetail(v *QueryIotCardOfferDtlResponseBodyCardOfferDetail) *QueryIotCardOfferDtlResponseBody {
	s.CardOfferDetail = v
	return s
}

func (s *QueryIotCardOfferDtlResponseBody) SetCode(v string) *QueryIotCardOfferDtlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryIotCardOfferDtlResponseBody) SetMessage(v string) *QueryIotCardOfferDtlResponseBody {
	s.Message = &v
	return s
}

func (s *QueryIotCardOfferDtlResponseBody) SetRequestId(v string) *QueryIotCardOfferDtlResponseBody {
	s.RequestId = &v
	return s
}

type QueryIotCardOfferDtlResponseBodyCardOfferDetail struct {
	Detail []*QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Repeated"`
}

func (s QueryIotCardOfferDtlResponseBodyCardOfferDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryIotCardOfferDtlResponseBodyCardOfferDetail) GoString() string {
	return s.String()
}

func (s *QueryIotCardOfferDtlResponseBodyCardOfferDetail) SetDetail(v []*QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) *QueryIotCardOfferDtlResponseBodyCardOfferDetail {
	s.Detail = v
	return s
}

type QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail struct {
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OfferId       *string `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
	OfferName     *string `json:"OfferName,omitempty" xml:"OfferName,omitempty"`
	OrderTime     *string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
}

func (s QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) GoString() string {
	return s.String()
}

func (s *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) SetEffectiveTime(v string) *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail {
	s.EffectiveTime = &v
	return s
}

func (s *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) SetExpireTime(v string) *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail {
	s.ExpireTime = &v
	return s
}

func (s *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) SetOfferId(v string) *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail {
	s.OfferId = &v
	return s
}

func (s *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) SetOfferName(v string) *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail {
	s.OfferName = &v
	return s
}

func (s *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail) SetOrderTime(v string) *QueryIotCardOfferDtlResponseBodyCardOfferDetailDetail {
	s.OrderTime = &v
	return s
}

type QueryIotCardOfferDtlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryIotCardOfferDtlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryIotCardOfferDtlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIotCardOfferDtlResponse) GoString() string {
	return s.String()
}

func (s *QueryIotCardOfferDtlResponse) SetHeaders(v map[string]*string) *QueryIotCardOfferDtlResponse {
	s.Headers = v
	return s
}

func (s *QueryIotCardOfferDtlResponse) SetStatusCode(v int32) *QueryIotCardOfferDtlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIotCardOfferDtlResponse) SetBody(v *QueryIotCardOfferDtlResponseBody) *QueryIotCardOfferDtlResponse {
	s.Body = v
	return s
}

type QueryPersonalInfoRequest struct {
	Iccid                *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPersonalInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPersonalInfoRequest) SetIccid(v string) *QueryPersonalInfoRequest {
	s.Iccid = &v
	return s
}

func (s *QueryPersonalInfoRequest) SetOwnerId(v int64) *QueryPersonalInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPersonalInfoRequest) SetResourceOwnerAccount(v string) *QueryPersonalInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPersonalInfoRequest) SetResourceOwnerId(v int64) *QueryPersonalInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPersonalInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPersonalInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPersonalInfoResponseBody) SetCode(v string) *QueryPersonalInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPersonalInfoResponseBody) SetMessage(v string) *QueryPersonalInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPersonalInfoResponseBody) SetModule(v string) *QueryPersonalInfoResponseBody {
	s.Module = &v
	return s
}

func (s *QueryPersonalInfoResponseBody) SetRequestId(v string) *QueryPersonalInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryPersonalInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPersonalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPersonalInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPersonalInfoResponse) SetHeaders(v map[string]*string) *QueryPersonalInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPersonalInfoResponse) SetStatusCode(v int32) *QueryPersonalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPersonalInfoResponse) SetBody(v *QueryPersonalInfoResponseBody) *QueryPersonalInfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dyiotapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DoIotChgBindOrUnBindWithOptions(request *DoIotChgBindOrUnBindRequest, runtime *util.RuntimeOptions) (_result *DoIotChgBindOrUnBindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.Imei)) {
		query["Imei"] = request.Imei
	}

	if !tea.BoolValue(util.IsUnset(request.MidChannelId)) {
		query["MidChannelId"] = request.MidChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.NewImei)) {
		query["NewImei"] = request.NewImei
	}

	if !tea.BoolValue(util.IsUnset(request.OpionType)) {
		query["OpionType"] = request.OpionType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotChgBindOrUnBind"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotChgBindOrUnBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotChgBindOrUnBind(request *DoIotChgBindOrUnBindRequest) (_result *DoIotChgBindOrUnBindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotChgBindOrUnBindResponse{}
	_body, _err := client.DoIotChgBindOrUnBindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoIotIsImeiExistWithOptions(request *DoIotIsImeiExistRequest, runtime *util.RuntimeOptions) (_result *DoIotIsImeiExistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Imei)) {
		query["Imei"] = request.Imei
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotIsImeiExist"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotIsImeiExistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotIsImeiExist(request *DoIotIsImeiExistRequest) (_result *DoIotIsImeiExistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotIsImeiExistResponse{}
	_body, _err := client.DoIotIsImeiExistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoIotPostImeiWithOptions(request *DoIotPostImeiRequest, runtime *util.RuntimeOptions) (_result *DoIotPostImeiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.Imei)) {
		query["Imei"] = request.Imei
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotPostImei"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotPostImeiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotPostImei(request *DoIotPostImeiRequest) (_result *DoIotPostImeiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotPostImeiResponse{}
	_body, _err := client.DoIotPostImeiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoIotSetAbsoluteRemindConfigWithOptions(request *DoIotSetAbsoluteRemindConfigRequest, runtime *util.RuntimeOptions) (_result *DoIotSetAbsoluteRemindConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigInfo)) {
		query["ConfigInfo"] = request.ConfigInfo
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotSetAbsoluteRemindConfig"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotSetAbsoluteRemindConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotSetAbsoluteRemindConfig(request *DoIotSetAbsoluteRemindConfigRequest) (_result *DoIotSetAbsoluteRemindConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotSetAbsoluteRemindConfigResponse{}
	_body, _err := client.DoIotSetAbsoluteRemindConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoIotSetRemindConfigWithOptions(request *DoIotSetRemindConfigRequest, runtime *util.RuntimeOptions) (_result *DoIotSetRemindConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigInfo)) {
		query["ConfigInfo"] = request.ConfigInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotSetRemindConfig"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotSetRemindConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotSetRemindConfig(request *DoIotSetRemindConfigRequest) (_result *DoIotSetRemindConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotSetRemindConfigResponse{}
	_body, _err := client.DoIotSetRemindConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoIotUnbindResumeWithOptions(request *DoIotUnbindResumeRequest, runtime *util.RuntimeOptions) (_result *DoIotUnbindResumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotUnbindResume"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotUnbindResumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotUnbindResume(request *DoIotUnbindResumeRequest) (_result *DoIotUnbindResumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotUnbindResumeResponse{}
	_body, _err := client.DoIotUnbindResumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoIotUserStopResumeWithOptions(request *DoIotUserStopResumeRequest, runtime *util.RuntimeOptions) (_result *DoIotUserStopResumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptionType)) {
		query["OptionType"] = request.OptionType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoIotUserStopResume"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoIotUserStopResumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoIotUserStopResume(request *DoIotUserStopResumeRequest) (_result *DoIotUserStopResumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoIotUserStopResumeResponse{}
	_body, _err := client.DoIotUserStopResumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoSendIotSmsWithOptions(request *DoSendIotSmsRequest, runtime *util.RuntimeOptions) (_result *DoSendIotSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParam)) {
		query["TemplateParam"] = request.TemplateParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoSendIotSms"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoSendIotSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoSendIotSms(request *DoSendIotSmsRequest) (_result *DoSendIotSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoSendIotSmsResponse{}
	_body, _err := client.DoSendIotSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardFlowInfoWithOptions(request *QueryCardFlowInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCardFlowInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardFlowInfo"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardFlowInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardFlowInfo(request *QueryCardFlowInfoRequest) (_result *QueryCardFlowInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardFlowInfoResponse{}
	_body, _err := client.QueryCardFlowInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardHistoryFlowInfoWithOptions(request *QueryCardHistoryFlowInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCardHistoryFlowInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardHistoryFlowInfo"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardHistoryFlowInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardHistoryFlowInfo(request *QueryCardHistoryFlowInfoRequest) (_result *QueryCardHistoryFlowInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardHistoryFlowInfoResponse{}
	_body, _err := client.QueryCardHistoryFlowInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardInfoWithOptions(request *QueryCardInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCardInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardInfo"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardInfo(request *QueryCardInfoRequest) (_result *QueryCardInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardInfoResponse{}
	_body, _err := client.QueryCardInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardStatusWithOptions(request *QueryCardStatusRequest, runtime *util.RuntimeOptions) (_result *QueryCardStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardStatus"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardStatus(request *QueryCardStatusRequest) (_result *QueryCardStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardStatusResponse{}
	_body, _err := client.QueryCardStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardsInfoWithOptions(request *QueryCardsInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCardsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardsInfo"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardsInfo(request *QueryCardsInfoRequest) (_result *QueryCardsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardsInfoResponse{}
	_body, _err := client.QueryCardsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryIotCardOfferDtlWithOptions(request *QueryIotCardOfferDtlRequest, runtime *util.RuntimeOptions) (_result *QueryIotCardOfferDtlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryIotCardOfferDtl"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryIotCardOfferDtlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIotCardOfferDtl(request *QueryIotCardOfferDtlRequest) (_result *QueryIotCardOfferDtlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryIotCardOfferDtlResponse{}
	_body, _err := client.QueryIotCardOfferDtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPersonalInfoWithOptions(request *QueryPersonalInfoRequest, runtime *util.RuntimeOptions) (_result *QueryPersonalInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPersonalInfo"),
		Version:     tea.String("2017-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPersonalInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPersonalInfo(request *QueryPersonalInfoRequest) (_result *QueryPersonalInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPersonalInfoResponse{}
	_body, _err := client.QueryPersonalInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

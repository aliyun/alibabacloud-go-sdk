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

type DeleteSecretBlacklistRequest struct {
	PoolKey    *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	BlackNo    *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	BlackType  *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	WayControl *string `json:"WayControl,omitempty" xml:"WayControl,omitempty"`
}

func (s DeleteSecretBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistRequest) SetPoolKey(v string) *DeleteSecretBlacklistRequest {
	s.PoolKey = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetBlackNo(v string) *DeleteSecretBlacklistRequest {
	s.BlackNo = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetRemark(v string) *DeleteSecretBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetBlackType(v string) *DeleteSecretBlacklistRequest {
	s.BlackType = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetWayControl(v string) *DeleteSecretBlacklistRequest {
	s.WayControl = &v
	return s
}

type DeleteSecretBlacklistResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecretBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistResponseBody) SetCode(v string) *DeleteSecretBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) SetMessage(v string) *DeleteSecretBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) SetRequestId(v string) *DeleteSecretBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecretBlacklistResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecretBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecretBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistResponse) SetHeaders(v map[string]*string) *DeleteSecretBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretBlacklistResponse) SetBody(v *DeleteSecretBlacklistResponseBody) *DeleteSecretBlacklistResponse {
	s.Body = v
	return s
}

type BindAxnExtensionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	Extension            *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	Expiration           *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpectCity           *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	IsRecordingEnabled   *bool   `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OutOrderId           *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	CallDisplayType      *int32  `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	RingConfig           *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
	ASRStatus            *bool   `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	ASRModelId           *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	CallRestrict         *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
}

func (s BindAxnExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionRequest) SetOwnerId(v int64) *BindAxnExtensionRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionRequest) SetResourceOwnerId(v int64) *BindAxnExtensionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPoolKey(v string) *BindAxnExtensionRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPhoneNoA(v string) *BindAxnExtensionRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExtension(v string) *BindAxnExtensionRequest {
	s.Extension = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPhoneNoB(v string) *BindAxnExtensionRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxnExtensionRequest) SetPhoneNoX(v string) *BindAxnExtensionRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExpiration(v string) *BindAxnExtensionRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionRequest) SetExpectCity(v string) *BindAxnExtensionRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxnExtensionRequest) SetIsRecordingEnabled(v bool) *BindAxnExtensionRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxnExtensionRequest) SetOutId(v string) *BindAxnExtensionRequest {
	s.OutId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetOutOrderId(v string) *BindAxnExtensionRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetCallDisplayType(v int32) *BindAxnExtensionRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxnExtensionRequest) SetRingConfig(v string) *BindAxnExtensionRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxnExtensionRequest) SetASRStatus(v bool) *BindAxnExtensionRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxnExtensionRequest) SetASRModelId(v string) *BindAxnExtensionRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxnExtensionRequest) SetCallRestrict(v string) *BindAxnExtensionRequest {
	s.CallRestrict = &v
	return s
}

type BindAxnExtensionResponseBody struct {
	Code          *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindDTO *BindAxnExtensionResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxnExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponseBody) SetCode(v string) *BindAxnExtensionResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetMessage(v string) *BindAxnExtensionResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetRequestId(v string) *BindAxnExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetSecretBindDTO(v *BindAxnExtensionResponseBodySecretBindDTO) *BindAxnExtensionResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxnExtensionResponseBodySecretBindDTO struct {
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	SubsId    *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo  *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BindAxnExtensionResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetExtension(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetSubsId(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

type BindAxnExtensionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindAxnExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAxnExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxnExtensionResponse) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponse) SetHeaders(v map[string]*string) *BindAxnExtensionResponse {
	s.Headers = v
	return s
}

func (s *BindAxnExtensionResponse) SetBody(v *BindAxnExtensionResponseBody) *BindAxnExtensionResponse {
	s.Body = v
	return s
}

type UpdatePhoneSwitchRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SubsId               *int64  `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	SwitchStatus         *int32  `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s UpdatePhoneSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneSwitchRequest) SetOwnerId(v int64) *UpdatePhoneSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePhoneSwitchRequest) SetResourceOwnerAccount(v string) *UpdatePhoneSwitchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePhoneSwitchRequest) SetResourceOwnerId(v int64) *UpdatePhoneSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePhoneSwitchRequest) SetPoolKey(v string) *UpdatePhoneSwitchRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdatePhoneSwitchRequest) SetSubsId(v int64) *UpdatePhoneSwitchRequest {
	s.SubsId = &v
	return s
}

func (s *UpdatePhoneSwitchRequest) SetSecretNo(v string) *UpdatePhoneSwitchRequest {
	s.SecretNo = &v
	return s
}

func (s *UpdatePhoneSwitchRequest) SetSwitchStatus(v int32) *UpdatePhoneSwitchRequest {
	s.SwitchStatus = &v
	return s
}

type UpdatePhoneSwitchResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneSwitchResponseBody) SetCode(v string) *UpdatePhoneSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneSwitchResponseBody) SetMessage(v string) *UpdatePhoneSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneSwitchResponseBody) SetData(v string) *UpdatePhoneSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePhoneSwitchResponseBody) SetRequestId(v string) *UpdatePhoneSwitchResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePhoneSwitchResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePhoneSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePhoneSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneSwitchResponse) SetHeaders(v map[string]*string) *UpdatePhoneSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneSwitchResponse) SetBody(v *UpdatePhoneSwitchResponseBody) *UpdatePhoneSwitchResponse {
	s.Body = v
	return s
}

type GetTotalPublicUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CheckSubs            *bool   `json:"CheckSubs,omitempty" xml:"CheckSubs,omitempty"`
	PartnerKey           *string `json:"PartnerKey,omitempty" xml:"PartnerKey,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallTime             *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
}

func (s GetTotalPublicUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlRequest) SetOwnerId(v int64) *GetTotalPublicUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetResourceOwnerAccount(v string) *GetTotalPublicUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetResourceOwnerId(v int64) *GetTotalPublicUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCheckSubs(v bool) *GetTotalPublicUrlRequest {
	s.CheckSubs = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetPartnerKey(v string) *GetTotalPublicUrlRequest {
	s.PartnerKey = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCallId(v string) *GetTotalPublicUrlRequest {
	s.CallId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCallTime(v string) *GetTotalPublicUrlRequest {
	s.CallTime = &v
	return s
}

type GetTotalPublicUrlResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetTotalPublicUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetTotalPublicUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponseBody) SetCode(v string) *GetTotalPublicUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetMessage(v string) *GetTotalPublicUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetRequestId(v string) *GetTotalPublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetData(v *GetTotalPublicUrlResponseBodyData) *GetTotalPublicUrlResponseBody {
	s.Data = v
	return s
}

type GetTotalPublicUrlResponseBodyData struct {
	RingPublicUrl  *string `json:"RingPublicUrl,omitempty" xml:"RingPublicUrl,omitempty"`
	PhonePublicUrl *string `json:"PhonePublicUrl,omitempty" xml:"PhonePublicUrl,omitempty"`
}

func (s GetTotalPublicUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponseBodyData) SetRingPublicUrl(v string) *GetTotalPublicUrlResponseBodyData {
	s.RingPublicUrl = &v
	return s
}

func (s *GetTotalPublicUrlResponseBodyData) SetPhonePublicUrl(v string) *GetTotalPublicUrlResponseBodyData {
	s.PhonePublicUrl = &v
	return s
}

type GetTotalPublicUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTotalPublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTotalPublicUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTotalPublicUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponse) SetHeaders(v map[string]*string) *GetTotalPublicUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTotalPublicUrlResponse) SetBody(v *GetTotalPublicUrlResponseBody) *GetTotalPublicUrlResponse {
	s.Body = v
	return s
}

type BuySecretNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	DisplayPool          *bool   `json:"DisplayPool,omitempty" xml:"DisplayPool,omitempty"`
}

func (s BuySecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoRequest) GoString() string {
	return s.String()
}

func (s *BuySecretNoRequest) SetOwnerId(v int64) *BuySecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *BuySecretNoRequest) SetResourceOwnerAccount(v string) *BuySecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BuySecretNoRequest) SetResourceOwnerId(v int64) *BuySecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BuySecretNoRequest) SetPoolKey(v string) *BuySecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *BuySecretNoRequest) SetSpecId(v int64) *BuySecretNoRequest {
	s.SpecId = &v
	return s
}

func (s *BuySecretNoRequest) SetCity(v string) *BuySecretNoRequest {
	s.City = &v
	return s
}

func (s *BuySecretNoRequest) SetSecretNo(v string) *BuySecretNoRequest {
	s.SecretNo = &v
	return s
}

func (s *BuySecretNoRequest) SetDisplayPool(v bool) *BuySecretNoRequest {
	s.DisplayPool = &v
	return s
}

type BuySecretNoResponseBody struct {
	Code             *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBuyInfoDTO *BuySecretNoResponseBodySecretBuyInfoDTO `json:"SecretBuyInfoDTO,omitempty" xml:"SecretBuyInfoDTO,omitempty" type:"Struct"`
}

func (s BuySecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponseBody) SetCode(v string) *BuySecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *BuySecretNoResponseBody) SetMessage(v string) *BuySecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *BuySecretNoResponseBody) SetRequestId(v string) *BuySecretNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuySecretNoResponseBody) SetSecretBuyInfoDTO(v *BuySecretNoResponseBodySecretBuyInfoDTO) *BuySecretNoResponseBody {
	s.SecretBuyInfoDTO = v
	return s
}

type BuySecretNoResponseBodySecretBuyInfoDTO struct {
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BuySecretNoResponseBodySecretBuyInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoResponseBodySecretBuyInfoDTO) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponseBodySecretBuyInfoDTO) SetSecretNo(v string) *BuySecretNoResponseBodySecretBuyInfoDTO {
	s.SecretNo = &v
	return s
}

type BuySecretNoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BuySecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BuySecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s BuySecretNoResponse) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponse) SetHeaders(v map[string]*string) *BuySecretNoResponse {
	s.Headers = v
	return s
}

func (s *BuySecretNoResponse) SetBody(v *BuySecretNoResponseBody) *BuySecretNoResponse {
	s.Body = v
	return s
}

type CreateSubscriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	BindToken            *string `json:"BindToken,omitempty" xml:"BindToken,omitempty"`
}

func (s CreateSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequest) SetOwnerId(v int64) *CreateSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetResourceOwnerAccount(v string) *CreateSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSubscriptionRequest) SetResourceOwnerId(v int64) *CreateSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetPoolKey(v string) *CreateSubscriptionRequest {
	s.PoolKey = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSecretNo(v string) *CreateSubscriptionRequest {
	s.SecretNo = &v
	return s
}

func (s *CreateSubscriptionRequest) SetPhoneNo(v string) *CreateSubscriptionRequest {
	s.PhoneNo = &v
	return s
}

func (s *CreateSubscriptionRequest) SetBindToken(v string) *CreateSubscriptionRequest {
	s.BindToken = &v
	return s
}

type CreateSubscriptionResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponseBody) SetCode(v string) *CreateSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetMessage(v string) *CreateSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetRequestId(v string) *CreateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetData(v *CreateSubscriptionResponseBodyData) *CreateSubscriptionResponseBody {
	s.Data = v
	return s
}

type CreateSubscriptionResponseBodyData struct {
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s CreateSubscriptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponseBodyData) SetSubsId(v string) *CreateSubscriptionResponseBodyData {
	s.SubsId = &v
	return s
}

type CreateSubscriptionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponse) SetHeaders(v map[string]*string) *CreateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionResponse) SetBody(v *CreateSubscriptionResponseBody) *CreateSubscriptionResponse {
	s.Body = v
	return s
}

type GetFaceVerifyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VerifyToken          *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s GetFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *GetFaceVerifyRequest) SetOwnerId(v int64) *GetFaceVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFaceVerifyRequest) SetResourceOwnerAccount(v string) *GetFaceVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFaceVerifyRequest) SetResourceOwnerId(v int64) *GetFaceVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFaceVerifyRequest) SetVerifyToken(v string) *GetFaceVerifyRequest {
	s.VerifyToken = &v
	return s
}

type GetFaceVerifyResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetFaceVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceVerifyResponseBody) SetCode(v string) *GetFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetFaceVerifyResponseBody) SetMessage(v string) *GetFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetFaceVerifyResponseBody) SetRequestId(v string) *GetFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFaceVerifyResponseBody) SetData(v *GetFaceVerifyResponseBodyData) *GetFaceVerifyResponseBody {
	s.Data = v
	return s
}

type GetFaceVerifyResponseBodyData struct {
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s GetFaceVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFaceVerifyResponseBodyData) SetVerifyResult(v string) *GetFaceVerifyResponseBodyData {
	s.VerifyResult = &v
	return s
}

type GetFaceVerifyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *GetFaceVerifyResponse) SetHeaders(v map[string]*string) *GetFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *GetFaceVerifyResponse) SetBody(v *GetFaceVerifyResponseBody) *GetFaceVerifyResponse {
	s.Body = v
	return s
}

type QuerySubscriptionDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubsId               *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
}

func (s QuerySubscriptionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailRequest) SetOwnerId(v int64) *QuerySubscriptionDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetResourceOwnerAccount(v string) *QuerySubscriptionDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetResourceOwnerId(v int64) *QuerySubscriptionDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetPoolKey(v string) *QuerySubscriptionDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetProductType(v string) *QuerySubscriptionDetailRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetSubsId(v string) *QuerySubscriptionDetailRequest {
	s.SubsId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetPhoneNoX(v string) *QuerySubscriptionDetailRequest {
	s.PhoneNoX = &v
	return s
}

type QuerySubscriptionDetailResponseBody struct {
	Code                *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message             *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindDetailDTO *QuerySubscriptionDetailResponseBodySecretBindDetailDTO `json:"SecretBindDetailDTO,omitempty" xml:"SecretBindDetailDTO,omitempty" type:"Struct"`
}

func (s QuerySubscriptionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponseBody) SetCode(v string) *QuerySubscriptionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetMessage(v string) *QuerySubscriptionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetRequestId(v string) *QuerySubscriptionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetSecretBindDetailDTO(v *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) *QuerySubscriptionDetailResponseBody {
	s.SecretBindDetailDTO = v
	return s
}

type QuerySubscriptionDetailResponseBodySecretBindDetailDTO struct {
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Extension    *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	PhoneNoB     *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	GroupId      *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NeedRecord   *bool   `json:"NeedRecord,omitempty" xml:"NeedRecord,omitempty"`
	ASRStatus    *bool   `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ExpireDate   *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	SubsId       *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	PhoneNoX     *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	ASRModelId   *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	PhoneNoA     *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
}

func (s QuerySubscriptionDetailResponseBodySecretBindDetailDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetStatus(v int64) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.Status = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetExtension(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.Extension = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoB(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoB = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetGroupId(v int64) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.GroupId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetNeedRecord(v bool) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.NeedRecord = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetASRStatus(v bool) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ASRStatus = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetGmtCreate(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.GmtCreate = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetExpireDate(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ExpireDate = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetSubsId(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.SubsId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoX(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetCallRestrict(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.CallRestrict = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetASRModelId(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ASRModelId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoA(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoA = &v
	return s
}

type QuerySubscriptionDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySubscriptionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySubscriptionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscriptionDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponse) SetHeaders(v map[string]*string) *QuerySubscriptionDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySubscriptionDetailResponse) SetBody(v *QuerySubscriptionDetailResponseBody) *QuerySubscriptionDetailResponse {
	s.Body = v
	return s
}

type QueryCallStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SubsId               *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	CallNo               *string `json:"CallNo,omitempty" xml:"CallNo,omitempty"`
}

func (s QueryCallStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryCallStatusRequest) SetOwnerId(v int64) *QueryCallStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallStatusRequest) SetResourceOwnerAccount(v string) *QueryCallStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallStatusRequest) SetResourceOwnerId(v int64) *QueryCallStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallStatusRequest) SetPoolKey(v string) *QueryCallStatusRequest {
	s.PoolKey = &v
	return s
}

func (s *QueryCallStatusRequest) SetSubsId(v string) *QueryCallStatusRequest {
	s.SubsId = &v
	return s
}

func (s *QueryCallStatusRequest) SetCallNo(v string) *QueryCallStatusRequest {
	s.CallNo = &v
	return s
}

type QueryCallStatusResponseBody struct {
	Code                *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message             *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretCallStatusDTO *QueryCallStatusResponseBodySecretCallStatusDTO `json:"SecretCallStatusDTO,omitempty" xml:"SecretCallStatusDTO,omitempty" type:"Struct"`
}

func (s QueryCallStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallStatusResponseBody) SetCode(v string) *QueryCallStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallStatusResponseBody) SetMessage(v string) *QueryCallStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallStatusResponseBody) SetRequestId(v string) *QueryCallStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallStatusResponseBody) SetSecretCallStatusDTO(v *QueryCallStatusResponseBodySecretCallStatusDTO) *QueryCallStatusResponseBody {
	s.SecretCallStatusDTO = v
	return s
}

type QueryCallStatusResponseBodySecretCallStatusDTO struct {
	CalledNo *string `json:"CalledNo,omitempty" xml:"CalledNo,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryCallStatusResponseBodySecretCallStatusDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryCallStatusResponseBodySecretCallStatusDTO) GoString() string {
	return s.String()
}

func (s *QueryCallStatusResponseBodySecretCallStatusDTO) SetCalledNo(v string) *QueryCallStatusResponseBodySecretCallStatusDTO {
	s.CalledNo = &v
	return s
}

func (s *QueryCallStatusResponseBodySecretCallStatusDTO) SetStatus(v int32) *QueryCallStatusResponseBodySecretCallStatusDTO {
	s.Status = &v
	return s
}

type QueryCallStatusResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCallStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCallStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryCallStatusResponse) SetHeaders(v map[string]*string) *QueryCallStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryCallStatusResponse) SetBody(v *QueryCallStatusResponseBody) *QueryCallStatusResponse {
	s.Body = v
	return s
}

type UpdatePhoneNumberRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SubsId               *int64  `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	NewPhoneNo           *string `json:"NewPhoneNo,omitempty" xml:"NewPhoneNo,omitempty"`
	BindToken            *string `json:"BindToken,omitempty" xml:"BindToken,omitempty"`
}

func (s UpdatePhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneNumberRequest) SetOwnerId(v int64) *UpdatePhoneNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetResourceOwnerAccount(v string) *UpdatePhoneNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetResourceOwnerId(v int64) *UpdatePhoneNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetPoolKey(v string) *UpdatePhoneNumberRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetSubsId(v int64) *UpdatePhoneNumberRequest {
	s.SubsId = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetSecretNo(v string) *UpdatePhoneNumberRequest {
	s.SecretNo = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetNewPhoneNo(v string) *UpdatePhoneNumberRequest {
	s.NewPhoneNo = &v
	return s
}

func (s *UpdatePhoneNumberRequest) SetBindToken(v string) *UpdatePhoneNumberRequest {
	s.BindToken = &v
	return s
}

type UpdatePhoneNumberResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneNumberResponseBody) SetCode(v string) *UpdatePhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneNumberResponseBody) SetMessage(v string) *UpdatePhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneNumberResponseBody) SetData(v string) *UpdatePhoneNumberResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePhoneNumberResponseBody) SetRequestId(v string) *UpdatePhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePhoneNumberResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneNumberResponse) SetHeaders(v map[string]*string) *UpdatePhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneNumberResponse) SetBody(v *UpdatePhoneNumberResponseBody) *UpdatePhoneNumberResponse {
	s.Body = v
	return s
}

type OperateBlackNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	BlackNo              *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	OperateType          *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Tips                 *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s OperateBlackNoRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateBlackNoRequest) GoString() string {
	return s.String()
}

func (s *OperateBlackNoRequest) SetOwnerId(v int64) *OperateBlackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateBlackNoRequest) SetResourceOwnerAccount(v string) *OperateBlackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateBlackNoRequest) SetResourceOwnerId(v int64) *OperateBlackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateBlackNoRequest) SetPoolKey(v string) *OperateBlackNoRequest {
	s.PoolKey = &v
	return s
}

func (s *OperateBlackNoRequest) SetBlackNo(v string) *OperateBlackNoRequest {
	s.BlackNo = &v
	return s
}

func (s *OperateBlackNoRequest) SetOperateType(v string) *OperateBlackNoRequest {
	s.OperateType = &v
	return s
}

func (s *OperateBlackNoRequest) SetTips(v string) *OperateBlackNoRequest {
	s.Tips = &v
	return s
}

type OperateBlackNoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateBlackNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateBlackNoResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBlackNoResponseBody) SetCode(v string) *OperateBlackNoResponseBody {
	s.Code = &v
	return s
}

func (s *OperateBlackNoResponseBody) SetMessage(v string) *OperateBlackNoResponseBody {
	s.Message = &v
	return s
}

func (s *OperateBlackNoResponseBody) SetRequestId(v string) *OperateBlackNoResponseBody {
	s.RequestId = &v
	return s
}

type OperateBlackNoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateBlackNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateBlackNoResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateBlackNoResponse) GoString() string {
	return s.String()
}

func (s *OperateBlackNoResponse) SetHeaders(v map[string]*string) *OperateBlackNoResponse {
	s.Headers = v
	return s
}

func (s *OperateBlackNoResponse) SetBody(v *OperateBlackNoResponseBody) *OperateBlackNoResponse {
	s.Body = v
	return s
}

type GetSecretAsrDetailRequest struct {
	CallId   *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	PoolKey  *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
}

func (s GetSecretAsrDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailRequest) SetCallId(v string) *GetSecretAsrDetailRequest {
	s.CallId = &v
	return s
}

func (s *GetSecretAsrDetailRequest) SetCallTime(v string) *GetSecretAsrDetailRequest {
	s.CallTime = &v
	return s
}

func (s *GetSecretAsrDetailRequest) SetPoolKey(v string) *GetSecretAsrDetailRequest {
	s.PoolKey = &v
	return s
}

type GetSecretAsrDetailResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GetSecretAsrDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretAsrDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBody) SetCode(v string) *GetSecretAsrDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetMessage(v string) *GetSecretAsrDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetData(v *GetSecretAsrDetailResponseBodyData) *GetSecretAsrDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetRequestId(v string) *GetSecretAsrDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetSecretAsrDetailResponseBodyData struct {
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BusinessId  *string                                        `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessKey *string                                        `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	Code        *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg         *string                                        `json:"Msg,omitempty" xml:"Msg,omitempty"`
	BizDuration *int64                                         `json:"BizDuration,omitempty" xml:"BizDuration,omitempty"`
	Type        *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Sentences   []*GetSecretAsrDetailResponseBodyDataSentences `json:"Sentences,omitempty" xml:"Sentences,omitempty" type:"Repeated"`
}

func (s GetSecretAsrDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBodyData) SetRequestId(v string) *GetSecretAsrDetailResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBusinessId(v string) *GetSecretAsrDetailResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBusinessKey(v string) *GetSecretAsrDetailResponseBodyData {
	s.BusinessKey = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetCode(v string) *GetSecretAsrDetailResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetMsg(v string) *GetSecretAsrDetailResponseBodyData {
	s.Msg = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBizDuration(v int64) *GetSecretAsrDetailResponseBodyData {
	s.BizDuration = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetType(v string) *GetSecretAsrDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetSentences(v []*GetSecretAsrDetailResponseBodyDataSentences) *GetSecretAsrDetailResponseBodyData {
	s.Sentences = v
	return s
}

type GetSecretAsrDetailResponseBodyDataSentences struct {
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SilenceDuration *int64  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	BeginTime       *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Text            *string `json:"Text,omitempty" xml:"Text,omitempty"`
	ChannelId       *int32  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	SpeechRate      *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	EmotionValue    *string `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
}

func (s GetSecretAsrDetailResponseBodyDataSentences) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponseBodyDataSentences) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetEndTime(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.EndTime = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetSilenceDuration(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.SilenceDuration = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetBeginTime(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.BeginTime = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetText(v string) *GetSecretAsrDetailResponseBodyDataSentences {
	s.Text = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetChannelId(v int32) *GetSecretAsrDetailResponseBodyDataSentences {
	s.ChannelId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetSpeechRate(v int32) *GetSecretAsrDetailResponseBodyDataSentences {
	s.SpeechRate = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetEmotionValue(v string) *GetSecretAsrDetailResponseBodyDataSentences {
	s.EmotionValue = &v
	return s
}

type GetSecretAsrDetailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecretAsrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecretAsrDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponse) SetHeaders(v map[string]*string) *GetSecretAsrDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSecretAsrDetailResponse) SetBody(v *GetSecretAsrDetailResponseBody) *GetSecretAsrDetailResponse {
	s.Body = v
	return s
}

type UpdateSubscriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubsId               *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	CallRestrict         *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	Expiration           *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	CallDisplayType      *int32  `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	IsRecordingEnabled   *bool   `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	OperateType          *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	RingConfig           *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
	ASRStatus            *bool   `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	ASRModelId           *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
}

func (s UpdateSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) SetOwnerId(v int64) *UpdateSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetResourceOwnerAccount(v string) *UpdateSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetResourceOwnerId(v int64) *UpdateSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPoolKey(v string) *UpdateSubscriptionRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetProductType(v string) *UpdateSubscriptionRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetSubsId(v string) *UpdateSubscriptionRequest {
	s.SubsId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPhoneNoX(v string) *UpdateSubscriptionRequest {
	s.PhoneNoX = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPhoneNoA(v string) *UpdateSubscriptionRequest {
	s.PhoneNoA = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPhoneNoB(v string) *UpdateSubscriptionRequest {
	s.PhoneNoB = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetGroupId(v string) *UpdateSubscriptionRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetCallRestrict(v string) *UpdateSubscriptionRequest {
	s.CallRestrict = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetExpiration(v string) *UpdateSubscriptionRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetCallDisplayType(v int32) *UpdateSubscriptionRequest {
	s.CallDisplayType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetOutId(v string) *UpdateSubscriptionRequest {
	s.OutId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetIsRecordingEnabled(v bool) *UpdateSubscriptionRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetOperateType(v string) *UpdateSubscriptionRequest {
	s.OperateType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetRingConfig(v string) *UpdateSubscriptionRequest {
	s.RingConfig = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetASRStatus(v bool) *UpdateSubscriptionRequest {
	s.ASRStatus = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetASRModelId(v string) *UpdateSubscriptionRequest {
	s.ASRModelId = &v
	return s
}

type UpdateSubscriptionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) SetCode(v string) *UpdateSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetMessage(v string) *UpdateSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSubscriptionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponse) SetHeaders(v map[string]*string) *UpdateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscriptionResponse) SetBody(v *UpdateSubscriptionResponseBody) *UpdateSubscriptionResponse {
	s.Body = v
	return s
}

type LockSecretNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s LockSecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s LockSecretNoRequest) GoString() string {
	return s.String()
}

func (s *LockSecretNoRequest) SetOwnerId(v int64) *LockSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *LockSecretNoRequest) SetResourceOwnerAccount(v string) *LockSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockSecretNoRequest) SetResourceOwnerId(v int64) *LockSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockSecretNoRequest) SetPoolKey(v string) *LockSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *LockSecretNoRequest) SetSecretNo(v string) *LockSecretNoRequest {
	s.SecretNo = &v
	return s
}

type LockSecretNoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockSecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *LockSecretNoResponseBody) SetCode(v string) *LockSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *LockSecretNoResponseBody) SetMessage(v string) *LockSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *LockSecretNoResponseBody) SetRequestId(v string) *LockSecretNoResponseBody {
	s.RequestId = &v
	return s
}

type LockSecretNoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LockSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockSecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s LockSecretNoResponse) GoString() string {
	return s.String()
}

func (s *LockSecretNoResponse) SetHeaders(v map[string]*string) *LockSecretNoResponse {
	s.Headers = v
	return s
}

func (s *LockSecretNoResponse) SetBody(v *LockSecretNoResponseBody) *LockSecretNoResponse {
	s.Body = v
	return s
}

type UnlockSecretNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UnlockSecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockSecretNoRequest) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoRequest) SetOwnerId(v int64) *UnlockSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockSecretNoRequest) SetResourceOwnerAccount(v string) *UnlockSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockSecretNoRequest) SetResourceOwnerId(v int64) *UnlockSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockSecretNoRequest) SetPoolKey(v string) *UnlockSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *UnlockSecretNoRequest) SetSecretNo(v string) *UnlockSecretNoRequest {
	s.SecretNo = &v
	return s
}

type UnlockSecretNoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockSecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoResponseBody) SetCode(v string) *UnlockSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *UnlockSecretNoResponseBody) SetMessage(v string) *UnlockSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *UnlockSecretNoResponseBody) SetRequestId(v string) *UnlockSecretNoResponseBody {
	s.RequestId = &v
	return s
}

type UnlockSecretNoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnlockSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockSecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockSecretNoResponse) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoResponse) SetHeaders(v map[string]*string) *UnlockSecretNoResponse {
	s.Headers = v
	return s
}

func (s *UnlockSecretNoResponse) SetBody(v *UnlockSecretNoResponseBody) *UnlockSecretNoResponse {
	s.Body = v
	return s
}

type ReleaseSecretNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s ReleaseSecretNoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSecretNoRequest) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoRequest) SetOwnerId(v int64) *ReleaseSecretNoRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetResourceOwnerAccount(v string) *ReleaseSecretNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetResourceOwnerId(v int64) *ReleaseSecretNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetPoolKey(v string) *ReleaseSecretNoRequest {
	s.PoolKey = &v
	return s
}

func (s *ReleaseSecretNoRequest) SetSecretNo(v string) *ReleaseSecretNoRequest {
	s.SecretNo = &v
	return s
}

type ReleaseSecretNoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseSecretNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoResponseBody) SetCode(v string) *ReleaseSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) SetMessage(v string) *ReleaseSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) SetRequestId(v string) *ReleaseSecretNoResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseSecretNoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseSecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseSecretNoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSecretNoResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoResponse) SetHeaders(v map[string]*string) *ReleaseSecretNoResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSecretNoResponse) SetBody(v *ReleaseSecretNoResponseBody) *ReleaseSecretNoResponse {
	s.Body = v
	return s
}

type OperateAxgGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OperateType          *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Numbers              *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
}

func (s OperateAxgGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupRequest) SetOwnerId(v int64) *OperateAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetResourceOwnerAccount(v string) *OperateAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateAxgGroupRequest) SetResourceOwnerId(v int64) *OperateAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetPoolKey(v string) *OperateAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *OperateAxgGroupRequest) SetGroupId(v int64) *OperateAxgGroupRequest {
	s.GroupId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetOperateType(v string) *OperateAxgGroupRequest {
	s.OperateType = &v
	return s
}

func (s *OperateAxgGroupRequest) SetNumbers(v string) *OperateAxgGroupRequest {
	s.Numbers = &v
	return s
}

type OperateAxgGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateAxgGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupResponseBody) SetCode(v string) *OperateAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *OperateAxgGroupResponseBody) SetMessage(v string) *OperateAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *OperateAxgGroupResponseBody) SetRequestId(v string) *OperateAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

type OperateAxgGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateAxgGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupResponse) SetHeaders(v map[string]*string) *OperateAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *OperateAxgGroupResponse) SetBody(v *OperateAxgGroupResponseBody) *OperateAxgGroupResponse {
	s.Body = v
	return s
}

type UpdateDefaultBRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SubsId               *int64  `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
}

func (s UpdateDefaultBRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDefaultBRequest) GoString() string {
	return s.String()
}

func (s *UpdateDefaultBRequest) SetOwnerId(v int64) *UpdateDefaultBRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDefaultBRequest) SetResourceOwnerAccount(v string) *UpdateDefaultBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateDefaultBRequest) SetResourceOwnerId(v int64) *UpdateDefaultBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateDefaultBRequest) SetPoolKey(v string) *UpdateDefaultBRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdateDefaultBRequest) SetSubsId(v int64) *UpdateDefaultBRequest {
	s.SubsId = &v
	return s
}

func (s *UpdateDefaultBRequest) SetSecretNo(v string) *UpdateDefaultBRequest {
	s.SecretNo = &v
	return s
}

func (s *UpdateDefaultBRequest) SetPhoneNo(v string) *UpdateDefaultBRequest {
	s.PhoneNo = &v
	return s
}

type UpdateDefaultBResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDefaultBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDefaultBResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDefaultBResponseBody) SetCode(v string) *UpdateDefaultBResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDefaultBResponseBody) SetMessage(v string) *UpdateDefaultBResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDefaultBResponseBody) SetData(v string) *UpdateDefaultBResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDefaultBResponseBody) SetRequestId(v string) *UpdateDefaultBResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDefaultBResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDefaultBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDefaultBResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDefaultBResponse) GoString() string {
	return s.String()
}

func (s *UpdateDefaultBResponse) SetHeaders(v map[string]*string) *UpdateDefaultBResponse {
	s.Headers = v
	return s
}

func (s *UpdateDefaultBResponse) SetBody(v *UpdateDefaultBResponseBody) *UpdateDefaultBResponse {
	s.Body = v
	return s
}

type QueryPhoneNoAByTrackNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TrackNo              *string `json:"trackNo,omitempty" xml:"trackNo,omitempty"`
	CabinetNo            *string `json:"CabinetNo,omitempty" xml:"CabinetNo,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
}

func (s QueryPhoneNoAByTrackNoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoRequest) SetOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetResourceOwnerAccount(v string) *QueryPhoneNoAByTrackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetResourceOwnerId(v int64) *QueryPhoneNoAByTrackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetTrackNo(v string) *QueryPhoneNoAByTrackNoRequest {
	s.TrackNo = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetCabinetNo(v string) *QueryPhoneNoAByTrackNoRequest {
	s.CabinetNo = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoRequest) SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoRequest {
	s.PhoneNoX = &v
	return s
}

type QueryPhoneNoAByTrackNoResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    []*QueryPhoneNoAByTrackNoResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s QueryPhoneNoAByTrackNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetCode(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetMessage(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetRequestId(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetModule(v []*QueryPhoneNoAByTrackNoResponseBodyModule) *QueryPhoneNoAByTrackNoResponseBody {
	s.Module = v
	return s
}

type QueryPhoneNoAByTrackNoResponseBodyModule struct {
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	PhoneNoX  *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	PhoneNoA  *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetExtension(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.Extension = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.PhoneNoX = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetPhoneNoA(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.PhoneNoA = &v
	return s
}

type QueryPhoneNoAByTrackNoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPhoneNoAByTrackNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPhoneNoAByTrackNoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponse) SetHeaders(v map[string]*string) *QueryPhoneNoAByTrackNoResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponse) SetBody(v *QueryPhoneNoAByTrackNoResponseBody) *QueryPhoneNoAByTrackNoResponse {
	s.Body = v
	return s
}

type AddAxnTrackNoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	TrackNo              *string `json:"trackNo,omitempty" xml:"trackNo,omitempty"`
	SubsId               *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s AddAxnTrackNoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAxnTrackNoRequest) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoRequest) SetOwnerId(v int64) *AddAxnTrackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetResourceOwnerAccount(v string) *AddAxnTrackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetResourceOwnerId(v int64) *AddAxnTrackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetPoolKey(v string) *AddAxnTrackNoRequest {
	s.PoolKey = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetPhoneNoX(v string) *AddAxnTrackNoRequest {
	s.PhoneNoX = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetTrackNo(v string) *AddAxnTrackNoRequest {
	s.TrackNo = &v
	return s
}

func (s *AddAxnTrackNoRequest) SetSubsId(v string) *AddAxnTrackNoRequest {
	s.SubsId = &v
	return s
}

type AddAxnTrackNoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAxnTrackNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAxnTrackNoResponseBody) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoResponseBody) SetCode(v string) *AddAxnTrackNoResponseBody {
	s.Code = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) SetMessage(v string) *AddAxnTrackNoResponseBody {
	s.Message = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) SetRequestId(v string) *AddAxnTrackNoResponseBody {
	s.RequestId = &v
	return s
}

type AddAxnTrackNoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAxnTrackNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAxnTrackNoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAxnTrackNoResponse) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoResponse) SetHeaders(v map[string]*string) *AddAxnTrackNoResponse {
	s.Headers = v
	return s
}

func (s *AddAxnTrackNoResponse) SetBody(v *AddAxnTrackNoResponseBody) *AddAxnTrackNoResponse {
	s.Body = v
	return s
}

type QuerySecretNoRemainRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s QuerySecretNoRemainRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainRequest) SetOwnerId(v int64) *QuerySecretNoRemainRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetResourceOwnerAccount(v string) *QuerySecretNoRemainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetResourceOwnerId(v int64) *QuerySecretNoRemainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetSpecId(v int64) *QuerySecretNoRemainRequest {
	s.SpecId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetCity(v string) *QuerySecretNoRemainRequest {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetSecretNo(v string) *QuerySecretNoRemainRequest {
	s.SecretNo = &v
	return s
}

type QuerySecretNoRemainResponseBody struct {
	Code            *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretRemainDTO *QuerySecretNoRemainResponseBodySecretRemainDTO `json:"SecretRemainDTO,omitempty" xml:"SecretRemainDTO,omitempty" type:"Struct"`
}

func (s QuerySecretNoRemainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBody) SetCode(v string) *QuerySecretNoRemainResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetMessage(v string) *QuerySecretNoRemainResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetRequestId(v string) *QuerySecretNoRemainResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetSecretRemainDTO(v *QuerySecretNoRemainResponseBodySecretRemainDTO) *QuerySecretNoRemainResponseBody {
	s.SecretRemainDTO = v
	return s
}

type QuerySecretNoRemainResponseBodySecretRemainDTO struct {
	Amount        *int64                                                       `json:"Amount,omitempty" xml:"Amount,omitempty"`
	City          *string                                                      `json:"City,omitempty" xml:"City,omitempty"`
	RemainDTOList *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList `json:"RemainDTOList,omitempty" xml:"RemainDTOList,omitempty" type:"Struct"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetAmount(v int64) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.Amount = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetCity(v string) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetRemainDTOList(v *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.RemainDTOList = v
	return s
}

type QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList struct {
	RemainDTO []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO `json:"remainDTO,omitempty" xml:"remainDTO,omitempty" type:"Repeated"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) SetRemainDTO(v []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList {
	s.RemainDTO = v
	return s
}

type QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO struct {
	Amount *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	City   *string `json:"City,omitempty" xml:"City,omitempty"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) SetAmount(v int64) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	s.Amount = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) SetCity(v string) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	s.City = &v
	return s
}

type QuerySecretNoRemainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySecretNoRemainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySecretNoRemainResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySecretNoRemainResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponse) SetHeaders(v map[string]*string) *QuerySecretNoRemainResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretNoRemainResponse) SetBody(v *QuerySecretNoRemainResponseBody) *QuerySecretNoRemainResponse {
	s.Body = v
	return s
}

type CreateAxgGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Numbers              *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
}

func (s CreateAxgGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupRequest) SetOwnerId(v int64) *CreateAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAxgGroupRequest) SetResourceOwnerAccount(v string) *CreateAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAxgGroupRequest) SetResourceOwnerId(v int64) *CreateAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAxgGroupRequest) SetPoolKey(v string) *CreateAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *CreateAxgGroupRequest) SetName(v string) *CreateAxgGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateAxgGroupRequest) SetRemark(v string) *CreateAxgGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateAxgGroupRequest) SetNumbers(v string) *CreateAxgGroupRequest {
	s.Numbers = &v
	return s
}

type CreateAxgGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateAxgGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupResponseBody) SetCode(v string) *CreateAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetMessage(v string) *CreateAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetRequestId(v string) *CreateAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetGroupId(v int64) *CreateAxgGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateAxgGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAxgGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupResponse) SetHeaders(v map[string]*string) *CreateAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAxgGroupResponse) SetBody(v *CreateAxgGroupResponseBody) *CreateAxgGroupResponse {
	s.Body = v
	return s
}

type BindAxbRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	Expiration           *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpectCity           *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	IsRecordingEnabled   *bool   `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OutOrderId           *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	CallRestrict         *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	CallDisplayType      *int32  `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	RingConfig           *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
	ASRStatus            *bool   `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	ASRModelId           *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	CallTimeout          *int32  `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
}

func (s BindAxbRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxbRequest) GoString() string {
	return s.String()
}

func (s *BindAxbRequest) SetOwnerId(v int64) *BindAxbRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbRequest) SetResourceOwnerAccount(v string) *BindAxbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbRequest) SetResourceOwnerId(v int64) *BindAxbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbRequest) SetPoolKey(v string) *BindAxbRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxbRequest) SetPhoneNoA(v string) *BindAxbRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxbRequest) SetPhoneNoB(v string) *BindAxbRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxbRequest) SetPhoneNoX(v string) *BindAxbRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxbRequest) SetExpiration(v string) *BindAxbRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbRequest) SetExpectCity(v string) *BindAxbRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxbRequest) SetIsRecordingEnabled(v bool) *BindAxbRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxbRequest) SetOutId(v string) *BindAxbRequest {
	s.OutId = &v
	return s
}

func (s *BindAxbRequest) SetOutOrderId(v string) *BindAxbRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxbRequest) SetCallRestrict(v string) *BindAxbRequest {
	s.CallRestrict = &v
	return s
}

func (s *BindAxbRequest) SetCallDisplayType(v int32) *BindAxbRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxbRequest) SetRingConfig(v string) *BindAxbRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxbRequest) SetASRStatus(v bool) *BindAxbRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxbRequest) SetASRModelId(v string) *BindAxbRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxbRequest) SetCallTimeout(v int32) *BindAxbRequest {
	s.CallTimeout = &v
	return s
}

type BindAxbResponseBody struct {
	Code          *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindDTO *BindAxbResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxbResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxbResponseBody) SetCode(v string) *BindAxbResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxbResponseBody) SetMessage(v string) *BindAxbResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxbResponseBody) SetRequestId(v string) *BindAxbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxbResponseBody) SetSecretBindDTO(v *BindAxbResponseBodySecretBindDTO) *BindAxbResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxbResponseBodySecretBindDTO struct {
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	SubsId    *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo  *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BindAxbResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxbResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxbResponseBodySecretBindDTO) SetExtension(v string) *BindAxbResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) SetSubsId(v string) *BindAxbResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxbResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

type BindAxbResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindAxbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAxbResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxbResponse) GoString() string {
	return s.String()
}

func (s *BindAxbResponse) SetHeaders(v map[string]*string) *BindAxbResponse {
	s.Headers = v
	return s
}

func (s *BindAxbResponse) SetBody(v *BindAxbResponseBody) *BindAxbResponse {
	s.Body = v
	return s
}

type UnbindSubscriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubsId               *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UnbindSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionRequest) SetOwnerId(v int64) *UnbindSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetResourceOwnerAccount(v string) *UnbindSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetResourceOwnerId(v int64) *UnbindSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetPoolKey(v string) *UnbindSubscriptionRequest {
	s.PoolKey = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetProductType(v string) *UnbindSubscriptionRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetSubsId(v string) *UnbindSubscriptionRequest {
	s.SubsId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetSecretNo(v string) *UnbindSubscriptionRequest {
	s.SecretNo = &v
	return s
}

type UnbindSubscriptionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChargeId  *string `json:"ChargeId,omitempty" xml:"ChargeId,omitempty"`
}

func (s UnbindSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionResponseBody) SetCode(v string) *UnbindSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetMessage(v string) *UnbindSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetRequestId(v string) *UnbindSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSubscriptionResponseBody) SetChargeId(v string) *UnbindSubscriptionResponseBody {
	s.ChargeId = &v
	return s
}

type UnbindSubscriptionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionResponse) SetHeaders(v map[string]*string) *UnbindSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UnbindSubscriptionResponse) SetBody(v *UnbindSubscriptionResponseBody) *UnbindSubscriptionResponse {
	s.Body = v
	return s
}

type QuerySubsIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
}

func (s QuerySubsIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubsIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySubsIdRequest) SetOwnerId(v int64) *QuerySubsIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySubsIdRequest) SetResourceOwnerAccount(v string) *QuerySubsIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySubsIdRequest) SetResourceOwnerId(v int64) *QuerySubsIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySubsIdRequest) SetPoolKey(v string) *QuerySubsIdRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySubsIdRequest) SetPhoneNoX(v string) *QuerySubsIdRequest {
	s.PhoneNoX = &v
	return s
}

type QuerySubsIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubsId    *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubsIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubsIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubsIdResponseBody) SetCode(v string) *QuerySubsIdResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetMessage(v string) *QuerySubsIdResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetRequestId(v string) *QuerySubsIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetSubsId(v string) *QuerySubsIdResponseBody {
	s.SubsId = &v
	return s
}

type QuerySubsIdResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySubsIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySubsIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubsIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySubsIdResponse) SetHeaders(v map[string]*string) *QuerySubsIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySubsIdResponse) SetBody(v *QuerySubsIdResponseBody) *QuerySubsIdResponse {
	s.Body = v
	return s
}

type ConfirmSendSmsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s ConfirmSendSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmSendSmsRequest) GoString() string {
	return s.String()
}

func (s *ConfirmSendSmsRequest) SetOwnerId(v int64) *ConfirmSendSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfirmSendSmsRequest) SetResourceOwnerAccount(v string) *ConfirmSendSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfirmSendSmsRequest) SetResourceOwnerId(v int64) *ConfirmSendSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfirmSendSmsRequest) SetPoolKey(v string) *ConfirmSendSmsRequest {
	s.PoolKey = &v
	return s
}

func (s *ConfirmSendSmsRequest) SetSecretNo(v string) *ConfirmSendSmsRequest {
	s.SecretNo = &v
	return s
}

func (s *ConfirmSendSmsRequest) SetCallId(v string) *ConfirmSendSmsRequest {
	s.CallId = &v
	return s
}

type ConfirmSendSmsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmSendSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmSendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmSendSmsResponseBody) SetCode(v string) *ConfirmSendSmsResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmSendSmsResponseBody) SetMessage(v string) *ConfirmSendSmsResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmSendSmsResponseBody) SetData(v string) *ConfirmSendSmsResponseBody {
	s.Data = &v
	return s
}

func (s *ConfirmSendSmsResponseBody) SetRequestId(v string) *ConfirmSendSmsResponseBody {
	s.RequestId = &v
	return s
}

type ConfirmSendSmsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmSendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmSendSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmSendSmsResponse) GoString() string {
	return s.String()
}

func (s *ConfirmSendSmsResponse) SetHeaders(v map[string]*string) *ConfirmSendSmsResponse {
	s.Headers = v
	return s
}

func (s *ConfirmSendSmsResponse) SetBody(v *ConfirmSendSmsResponseBody) *ConfirmSendSmsResponse {
	s.Body = v
	return s
}

type GetSubscriptionDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	SubsId               *int64  `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s GetSubscriptionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionDetailRequest) SetOwnerId(v int64) *GetSubscriptionDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSubscriptionDetailRequest) SetResourceOwnerAccount(v string) *GetSubscriptionDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSubscriptionDetailRequest) SetResourceOwnerId(v int64) *GetSubscriptionDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSubscriptionDetailRequest) SetPoolKey(v string) *GetSubscriptionDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *GetSubscriptionDetailRequest) SetSubsId(v int64) *GetSubscriptionDetailRequest {
	s.SubsId = &v
	return s
}

func (s *GetSubscriptionDetailRequest) SetSecretNo(v string) *GetSubscriptionDetailRequest {
	s.SecretNo = &v
	return s
}

type GetSubscriptionDetailResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetSubscriptionDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSubscriptionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionDetailResponseBody) SetCode(v string) *GetSubscriptionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionDetailResponseBody) SetMessage(v string) *GetSubscriptionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionDetailResponseBody) SetRequestId(v string) *GetSubscriptionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionDetailResponseBody) SetData(v *GetSubscriptionDetailResponseBodyData) *GetSubscriptionDetailResponseBody {
	s.Data = v
	return s
}

type GetSubscriptionDetailResponseBodyData struct {
	PhoneNo      *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	City         *string `json:"City,omitempty" xml:"City,omitempty"`
	SwitchStatus *int32  `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	SubsId       *int64  `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo     *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	Vendor       *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	Province     *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetSubscriptionDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionDetailResponseBodyData) SetPhoneNo(v string) *GetSubscriptionDetailResponseBodyData {
	s.PhoneNo = &v
	return s
}

func (s *GetSubscriptionDetailResponseBodyData) SetCity(v string) *GetSubscriptionDetailResponseBodyData {
	s.City = &v
	return s
}

func (s *GetSubscriptionDetailResponseBodyData) SetSwitchStatus(v int32) *GetSubscriptionDetailResponseBodyData {
	s.SwitchStatus = &v
	return s
}

func (s *GetSubscriptionDetailResponseBodyData) SetSubsId(v int64) *GetSubscriptionDetailResponseBodyData {
	s.SubsId = &v
	return s
}

func (s *GetSubscriptionDetailResponseBodyData) SetSecretNo(v string) *GetSubscriptionDetailResponseBodyData {
	s.SecretNo = &v
	return s
}

func (s *GetSubscriptionDetailResponseBodyData) SetVendor(v string) *GetSubscriptionDetailResponseBodyData {
	s.Vendor = &v
	return s
}

func (s *GetSubscriptionDetailResponseBodyData) SetProvince(v string) *GetSubscriptionDetailResponseBodyData {
	s.Province = &v
	return s
}

type GetSubscriptionDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSubscriptionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubscriptionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionDetailResponse) SetHeaders(v map[string]*string) *GetSubscriptionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionDetailResponse) SetBody(v *GetSubscriptionDetailResponseBody) *GetSubscriptionDetailResponse {
	s.Body = v
	return s
}

type BindAxgRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	Expiration           *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpectCity           *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	IsRecordingEnabled   *bool   `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OutOrderId           *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	CallDisplayType      *int32  `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	RingConfig           *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
	ASRStatus            *bool   `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	ASRModelId           *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	CallRestrict         *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
}

func (s BindAxgRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxgRequest) GoString() string {
	return s.String()
}

func (s *BindAxgRequest) SetOwnerId(v int64) *BindAxgRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxgRequest) SetResourceOwnerAccount(v string) *BindAxgRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxgRequest) SetResourceOwnerId(v int64) *BindAxgRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxgRequest) SetPoolKey(v string) *BindAxgRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxgRequest) SetPhoneNoA(v string) *BindAxgRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxgRequest) SetGroupId(v string) *BindAxgRequest {
	s.GroupId = &v
	return s
}

func (s *BindAxgRequest) SetPhoneNoB(v string) *BindAxgRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxgRequest) SetPhoneNoX(v string) *BindAxgRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxgRequest) SetExpiration(v string) *BindAxgRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxgRequest) SetExpectCity(v string) *BindAxgRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxgRequest) SetIsRecordingEnabled(v bool) *BindAxgRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxgRequest) SetOutId(v string) *BindAxgRequest {
	s.OutId = &v
	return s
}

func (s *BindAxgRequest) SetOutOrderId(v string) *BindAxgRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxgRequest) SetCallDisplayType(v int32) *BindAxgRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxgRequest) SetRingConfig(v string) *BindAxgRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxgRequest) SetASRStatus(v bool) *BindAxgRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxgRequest) SetASRModelId(v string) *BindAxgRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxgRequest) SetCallRestrict(v string) *BindAxgRequest {
	s.CallRestrict = &v
	return s
}

type BindAxgResponseBody struct {
	Code          *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindDTO *BindAxgResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxgResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxgResponseBody) SetCode(v string) *BindAxgResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxgResponseBody) SetMessage(v string) *BindAxgResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxgResponseBody) SetRequestId(v string) *BindAxgResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxgResponseBody) SetSecretBindDTO(v *BindAxgResponseBodySecretBindDTO) *BindAxgResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxgResponseBodySecretBindDTO struct {
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	SubsId    *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo  *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BindAxgResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxgResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxgResponseBodySecretBindDTO) SetExtension(v string) *BindAxgResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) SetSubsId(v string) *BindAxgResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxgResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

type BindAxgResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindAxgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAxgResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxgResponse) GoString() string {
	return s.String()
}

func (s *BindAxgResponse) SetHeaders(v map[string]*string) *BindAxgResponse {
	s.Headers = v
	return s
}

func (s *BindAxgResponse) SetBody(v *BindAxgResponseBody) *BindAxgResponse {
	s.Body = v
	return s
}

type QueryRecordFileDownloadUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallTime             *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
}

func (s QueryRecordFileDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordFileDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlRequest) SetOwnerId(v int64) *QueryRecordFileDownloadUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetResourceOwnerAccount(v string) *QueryRecordFileDownloadUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetResourceOwnerId(v int64) *QueryRecordFileDownloadUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetPoolKey(v string) *QueryRecordFileDownloadUrlRequest {
	s.PoolKey = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetProductType(v string) *QueryRecordFileDownloadUrlRequest {
	s.ProductType = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetCallId(v string) *QueryRecordFileDownloadUrlRequest {
	s.CallId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetCallTime(v string) *QueryRecordFileDownloadUrlRequest {
	s.CallTime = &v
	return s
}

type QueryRecordFileDownloadUrlResponseBody struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s QueryRecordFileDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordFileDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetCode(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetMessage(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetRequestId(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetDownloadUrl(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

type QueryRecordFileDownloadUrlResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRecordFileDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRecordFileDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordFileDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryRecordFileDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordFileDownloadUrlResponse) SetBody(v *QueryRecordFileDownloadUrlResponseBody) *QueryRecordFileDownloadUrlResponse {
	s.Body = v
	return s
}

type InitFaceVerifyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	MetaInfo             *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
}

func (s InitFaceVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyRequest) SetOwnerId(v int64) *InitFaceVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetResourceOwnerAccount(v string) *InitFaceVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InitFaceVerifyRequest) SetResourceOwnerId(v int64) *InitFaceVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InitFaceVerifyRequest) SetMetaInfo(v string) *InitFaceVerifyRequest {
	s.MetaInfo = &v
	return s
}

type InitFaceVerifyResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *InitFaceVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s InitFaceVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBody) SetCode(v string) *InitFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetMessage(v string) *InitFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetRequestId(v string) *InitFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetData(v *InitFaceVerifyResponseBodyData) *InitFaceVerifyResponseBody {
	s.Data = v
	return s
}

type InitFaceVerifyResponseBodyData struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
}

func (s InitFaceVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBodyData) SetCertifyId(v string) *InitFaceVerifyResponseBodyData {
	s.CertifyId = &v
	return s
}

type InitFaceVerifyResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitFaceVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponse) SetHeaders(v map[string]*string) *InitFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitFaceVerifyResponse) SetBody(v *InitFaceVerifyResponseBody) *InitFaceVerifyResponse {
	s.Body = v
	return s
}

type BindAxnRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	PhoneNoX             *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	Expiration           *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpectCity           *string `json:"ExpectCity,omitempty" xml:"ExpectCity,omitempty"`
	IsRecordingEnabled   *bool   `json:"IsRecordingEnabled,omitempty" xml:"IsRecordingEnabled,omitempty"`
	NoType               *string `json:"NoType,omitempty" xml:"NoType,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OutOrderId           *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
	CallDisplayType      *int32  `json:"CallDisplayType,omitempty" xml:"CallDisplayType,omitempty"`
	RingConfig           *string `json:"RingConfig,omitempty" xml:"RingConfig,omitempty"`
	ASRStatus            *bool   `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	ASRModelId           *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	CallTimeout          *int32  `json:"CallTimeout,omitempty" xml:"CallTimeout,omitempty"`
	CallRestrict         *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
}

func (s BindAxnRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAxnRequest) GoString() string {
	return s.String()
}

func (s *BindAxnRequest) SetOwnerId(v int64) *BindAxnRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnRequest) SetResourceOwnerAccount(v string) *BindAxnRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnRequest) SetResourceOwnerId(v int64) *BindAxnRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnRequest) SetPoolKey(v string) *BindAxnRequest {
	s.PoolKey = &v
	return s
}

func (s *BindAxnRequest) SetPhoneNoA(v string) *BindAxnRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindAxnRequest) SetPhoneNoB(v string) *BindAxnRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindAxnRequest) SetPhoneNoX(v string) *BindAxnRequest {
	s.PhoneNoX = &v
	return s
}

func (s *BindAxnRequest) SetExpiration(v string) *BindAxnRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnRequest) SetExpectCity(v string) *BindAxnRequest {
	s.ExpectCity = &v
	return s
}

func (s *BindAxnRequest) SetIsRecordingEnabled(v bool) *BindAxnRequest {
	s.IsRecordingEnabled = &v
	return s
}

func (s *BindAxnRequest) SetNoType(v string) *BindAxnRequest {
	s.NoType = &v
	return s
}

func (s *BindAxnRequest) SetOutId(v string) *BindAxnRequest {
	s.OutId = &v
	return s
}

func (s *BindAxnRequest) SetOutOrderId(v string) *BindAxnRequest {
	s.OutOrderId = &v
	return s
}

func (s *BindAxnRequest) SetCallDisplayType(v int32) *BindAxnRequest {
	s.CallDisplayType = &v
	return s
}

func (s *BindAxnRequest) SetRingConfig(v string) *BindAxnRequest {
	s.RingConfig = &v
	return s
}

func (s *BindAxnRequest) SetASRStatus(v bool) *BindAxnRequest {
	s.ASRStatus = &v
	return s
}

func (s *BindAxnRequest) SetASRModelId(v string) *BindAxnRequest {
	s.ASRModelId = &v
	return s
}

func (s *BindAxnRequest) SetCallTimeout(v int32) *BindAxnRequest {
	s.CallTimeout = &v
	return s
}

func (s *BindAxnRequest) SetCallRestrict(v string) *BindAxnRequest {
	s.CallRestrict = &v
	return s
}

type BindAxnResponseBody struct {
	Code          *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretBindDTO *BindAxnResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAxnResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnResponseBody) SetCode(v string) *BindAxnResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnResponseBody) SetMessage(v string) *BindAxnResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnResponseBody) SetRequestId(v string) *BindAxnResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnResponseBody) SetSecretBindDTO(v *BindAxnResponseBodySecretBindDTO) *BindAxnResponseBody {
	s.SecretBindDTO = v
	return s
}

type BindAxnResponseBodySecretBindDTO struct {
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	SubsId    *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
	SecretNo  *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BindAxnResponseBodySecretBindDTO) String() string {
	return tea.Prettify(s)
}

func (s BindAxnResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxnResponseBodySecretBindDTO) SetExtension(v string) *BindAxnResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) SetSubsId(v string) *BindAxnResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxnResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

type BindAxnResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindAxnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAxnResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAxnResponse) GoString() string {
	return s.String()
}

func (s *BindAxnResponse) SetHeaders(v map[string]*string) *BindAxnResponse {
	s.Headers = v
	return s
}

func (s *BindAxnResponse) SetBody(v *BindAxnResponseBody) *BindAxnResponse {
	s.Body = v
	return s
}

type AddSecretBlacklistRequest struct {
	PoolKey    *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	BlackNo    *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	BlackType  *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	WayControl *string `json:"WayControl,omitempty" xml:"WayControl,omitempty"`
}

func (s AddSecretBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSecretBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistRequest) SetPoolKey(v string) *AddSecretBlacklistRequest {
	s.PoolKey = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetBlackNo(v string) *AddSecretBlacklistRequest {
	s.BlackNo = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetRemark(v string) *AddSecretBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetBlackType(v string) *AddSecretBlacklistRequest {
	s.BlackType = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetWayControl(v string) *AddSecretBlacklistRequest {
	s.WayControl = &v
	return s
}

type AddSecretBlacklistResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecretBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSecretBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistResponseBody) SetCode(v string) *AddSecretBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) SetMessage(v string) *AddSecretBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) SetRequestId(v string) *AddSecretBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type AddSecretBlacklistResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSecretBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSecretBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSecretBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistResponse) SetHeaders(v map[string]*string) *AddSecretBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddSecretBlacklistResponse) SetBody(v *AddSecretBlacklistResponseBody) *AddSecretBlacklistResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dyplsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DeleteSecretBlacklistWithOptions(request *DeleteSecretBlacklistRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSecretBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecretBlacklist"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecretBlacklist(request *DeleteSecretBlacklistRequest) (_result *DeleteSecretBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretBlacklistResponse{}
	_body, _err := client.DeleteSecretBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAxnExtensionWithOptions(request *BindAxnExtensionRequest, runtime *util.RuntimeOptions) (_result *BindAxnExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindAxnExtensionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindAxnExtension"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAxnExtension(request *BindAxnExtensionRequest) (_result *BindAxnExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxnExtensionResponse{}
	_body, _err := client.BindAxnExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePhoneSwitchWithOptions(request *UpdatePhoneSwitchRequest, runtime *util.RuntimeOptions) (_result *UpdatePhoneSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePhoneSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePhoneSwitch"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePhoneSwitch(request *UpdatePhoneSwitchRequest) (_result *UpdatePhoneSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePhoneSwitchResponse{}
	_body, _err := client.UpdatePhoneSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTotalPublicUrlWithOptions(request *GetTotalPublicUrlRequest, runtime *util.RuntimeOptions) (_result *GetTotalPublicUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTotalPublicUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTotalPublicUrl"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTotalPublicUrl(request *GetTotalPublicUrlRequest) (_result *GetTotalPublicUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTotalPublicUrlResponse{}
	_body, _err := client.GetTotalPublicUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BuySecretNoWithOptions(request *BuySecretNoRequest, runtime *util.RuntimeOptions) (_result *BuySecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BuySecretNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BuySecretNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BuySecretNo(request *BuySecretNoRequest) (_result *BuySecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BuySecretNoResponse{}
	_body, _err := client.BuySecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscriptionWithOptions(request *CreateSubscriptionRequest, runtime *util.RuntimeOptions) (_result *CreateSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSubscription"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscription(request *CreateSubscriptionRequest) (_result *CreateSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubscriptionResponse{}
	_body, _err := client.CreateSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFaceVerifyWithOptions(request *GetFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *GetFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFaceVerify"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaceVerify(request *GetFaceVerifyRequest) (_result *GetFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFaceVerifyResponse{}
	_body, _err := client.GetFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySubscriptionDetailWithOptions(request *QuerySubscriptionDetailRequest, runtime *util.RuntimeOptions) (_result *QuerySubscriptionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySubscriptionDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySubscriptionDetail"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySubscriptionDetail(request *QuerySubscriptionDetailRequest) (_result *QuerySubscriptionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySubscriptionDetailResponse{}
	_body, _err := client.QuerySubscriptionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCallStatusWithOptions(request *QueryCallStatusRequest, runtime *util.RuntimeOptions) (_result *QueryCallStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCallStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCallStatus"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCallStatus(request *QueryCallStatusRequest) (_result *QueryCallStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallStatusResponse{}
	_body, _err := client.QueryCallStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePhoneNumberWithOptions(request *UpdatePhoneNumberRequest, runtime *util.RuntimeOptions) (_result *UpdatePhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePhoneNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePhoneNumber"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePhoneNumber(request *UpdatePhoneNumberRequest) (_result *UpdatePhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePhoneNumberResponse{}
	_body, _err := client.UpdatePhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateBlackNoWithOptions(request *OperateBlackNoRequest, runtime *util.RuntimeOptions) (_result *OperateBlackNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OperateBlackNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OperateBlackNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateBlackNo(request *OperateBlackNoRequest) (_result *OperateBlackNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateBlackNoResponse{}
	_body, _err := client.OperateBlackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretAsrDetailWithOptions(request *GetSecretAsrDetailRequest, runtime *util.RuntimeOptions) (_result *GetSecretAsrDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSecretAsrDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSecretAsrDetail"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecretAsrDetail(request *GetSecretAsrDetailRequest) (_result *GetSecretAsrDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretAsrDetailResponse{}
	_body, _err := client.GetSecretAsrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubscriptionWithOptions(request *UpdateSubscriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSubscription"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (_result *UpdateSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.UpdateSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockSecretNoWithOptions(request *LockSecretNoRequest, runtime *util.RuntimeOptions) (_result *LockSecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LockSecretNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LockSecretNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockSecretNo(request *LockSecretNoRequest) (_result *LockSecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockSecretNoResponse{}
	_body, _err := client.LockSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockSecretNoWithOptions(request *UnlockSecretNoRequest, runtime *util.RuntimeOptions) (_result *UnlockSecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnlockSecretNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnlockSecretNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockSecretNo(request *UnlockSecretNoRequest) (_result *UnlockSecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockSecretNoResponse{}
	_body, _err := client.UnlockSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseSecretNoWithOptions(request *ReleaseSecretNoRequest, runtime *util.RuntimeOptions) (_result *ReleaseSecretNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseSecretNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseSecretNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseSecretNo(request *ReleaseSecretNoRequest) (_result *ReleaseSecretNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseSecretNoResponse{}
	_body, _err := client.ReleaseSecretNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateAxgGroupWithOptions(request *OperateAxgGroupRequest, runtime *util.RuntimeOptions) (_result *OperateAxgGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OperateAxgGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OperateAxgGroup"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateAxgGroup(request *OperateAxgGroupRequest) (_result *OperateAxgGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateAxgGroupResponse{}
	_body, _err := client.OperateAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDefaultBWithOptions(request *UpdateDefaultBRequest, runtime *util.RuntimeOptions) (_result *UpdateDefaultBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDefaultBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDefaultB"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDefaultB(request *UpdateDefaultBRequest) (_result *UpdateDefaultBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDefaultBResponse{}
	_body, _err := client.UpdateDefaultBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPhoneNoAByTrackNoWithOptions(request *QueryPhoneNoAByTrackNoRequest, runtime *util.RuntimeOptions) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPhoneNoAByTrackNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPhoneNoAByTrackNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPhoneNoAByTrackNo(request *QueryPhoneNoAByTrackNoRequest) (_result *QueryPhoneNoAByTrackNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPhoneNoAByTrackNoResponse{}
	_body, _err := client.QueryPhoneNoAByTrackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAxnTrackNoWithOptions(request *AddAxnTrackNoRequest, runtime *util.RuntimeOptions) (_result *AddAxnTrackNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAxnTrackNoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAxnTrackNo"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAxnTrackNo(request *AddAxnTrackNoRequest) (_result *AddAxnTrackNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAxnTrackNoResponse{}
	_body, _err := client.AddAxnTrackNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySecretNoRemainWithOptions(request *QuerySecretNoRemainRequest, runtime *util.RuntimeOptions) (_result *QuerySecretNoRemainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySecretNoRemainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySecretNoRemain"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySecretNoRemain(request *QuerySecretNoRemainRequest) (_result *QuerySecretNoRemainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySecretNoRemainResponse{}
	_body, _err := client.QuerySecretNoRemainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAxgGroupWithOptions(request *CreateAxgGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAxgGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAxgGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAxgGroup"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAxgGroup(request *CreateAxgGroupRequest) (_result *CreateAxgGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAxgGroupResponse{}
	_body, _err := client.CreateAxgGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAxbWithOptions(request *BindAxbRequest, runtime *util.RuntimeOptions) (_result *BindAxbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindAxbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindAxb"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAxb(request *BindAxbRequest) (_result *BindAxbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxbResponse{}
	_body, _err := client.BindAxbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindSubscriptionWithOptions(request *UnbindSubscriptionRequest, runtime *util.RuntimeOptions) (_result *UnbindSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindSubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindSubscription"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindSubscription(request *UnbindSubscriptionRequest) (_result *UnbindSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindSubscriptionResponse{}
	_body, _err := client.UnbindSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySubsIdWithOptions(request *QuerySubsIdRequest, runtime *util.RuntimeOptions) (_result *QuerySubsIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySubsIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySubsId"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySubsId(request *QuerySubsIdRequest) (_result *QuerySubsIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySubsIdResponse{}
	_body, _err := client.QuerySubsIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmSendSmsWithOptions(request *ConfirmSendSmsRequest, runtime *util.RuntimeOptions) (_result *ConfirmSendSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfirmSendSmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfirmSendSms"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmSendSms(request *ConfirmSendSmsRequest) (_result *ConfirmSendSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmSendSmsResponse{}
	_body, _err := client.ConfirmSendSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubscriptionDetailWithOptions(request *GetSubscriptionDetailRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSubscriptionDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSubscriptionDetail"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubscriptionDetail(request *GetSubscriptionDetailRequest) (_result *GetSubscriptionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubscriptionDetailResponse{}
	_body, _err := client.GetSubscriptionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAxgWithOptions(request *BindAxgRequest, runtime *util.RuntimeOptions) (_result *BindAxgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindAxgResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindAxg"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAxg(request *BindAxgRequest) (_result *BindAxgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxgResponse{}
	_body, _err := client.BindAxgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordFileDownloadUrlWithOptions(request *QueryRecordFileDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRecordFileDownloadUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRecordFileDownloadUrl"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordFileDownloadUrl(request *QueryRecordFileDownloadUrlRequest) (_result *QueryRecordFileDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordFileDownloadUrlResponse{}
	_body, _err := client.QueryRecordFileDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitFaceVerifyWithOptions(request *InitFaceVerifyRequest, runtime *util.RuntimeOptions) (_result *InitFaceVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitFaceVerify"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest) (_result *InitFaceVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAxnWithOptions(request *BindAxnRequest, runtime *util.RuntimeOptions) (_result *BindAxnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindAxnResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindAxn"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAxn(request *BindAxnRequest) (_result *BindAxnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAxnResponse{}
	_body, _err := client.BindAxnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSecretBlacklistWithOptions(request *AddSecretBlacklistRequest, runtime *util.RuntimeOptions) (_result *AddSecretBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSecretBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSecretBlacklist"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSecretBlacklist(request *AddSecretBlacklistRequest) (_result *AddSecretBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSecretBlacklistResponse{}
	_body, _err := client.AddSecretBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

type BandOfferOrderRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	BandId               *string `json:"BandId,omitempty" xml:"BandId,omitempty"`
	OfferId              *string `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
}

func (s BandOfferOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s BandOfferOrderRequest) GoString() string {
	return s.String()
}

func (s *BandOfferOrderRequest) SetOwnerId(v int64) *BandOfferOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *BandOfferOrderRequest) SetResourceOwnerAccount(v string) *BandOfferOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BandOfferOrderRequest) SetResourceOwnerId(v int64) *BandOfferOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BandOfferOrderRequest) SetBandId(v string) *BandOfferOrderRequest {
	s.BandId = &v
	return s
}

func (s *BandOfferOrderRequest) SetOfferId(v string) *BandOfferOrderRequest {
	s.OfferId = &v
	return s
}

type BandOfferOrderResponseBody struct {
	ResultModule  *BandOfferOrderResponseBodyResultModule `json:"ResultModule,omitempty" xml:"ResultModule,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s BandOfferOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandOfferOrderResponseBody) GoString() string {
	return s.String()
}

func (s *BandOfferOrderResponseBody) SetResultModule(v *BandOfferOrderResponseBodyResultModule) *BandOfferOrderResponseBody {
	s.ResultModule = v
	return s
}

func (s *BandOfferOrderResponseBody) SetRequestId(v string) *BandOfferOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandOfferOrderResponseBody) SetResultMessage(v string) *BandOfferOrderResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *BandOfferOrderResponseBody) SetResultCode(v string) *BandOfferOrderResponseBody {
	s.ResultCode = &v
	return s
}

type BandOfferOrderResponseBodyResultModule struct {
	LxOrderId *int64 `json:"LxOrderId,omitempty" xml:"LxOrderId,omitempty"`
}

func (s BandOfferOrderResponseBodyResultModule) String() string {
	return tea.Prettify(s)
}

func (s BandOfferOrderResponseBodyResultModule) GoString() string {
	return s.String()
}

func (s *BandOfferOrderResponseBodyResultModule) SetLxOrderId(v int64) *BandOfferOrderResponseBodyResultModule {
	s.LxOrderId = &v
	return s
}

type BandOfferOrderResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandOfferOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandOfferOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s BandOfferOrderResponse) GoString() string {
	return s.String()
}

func (s *BandOfferOrderResponse) SetHeaders(v map[string]*string) *BandOfferOrderResponse {
	s.Headers = v
	return s
}

func (s *BandOfferOrderResponse) SetBody(v *BandOfferOrderResponseBody) *BandOfferOrderResponse {
	s.Body = v
	return s
}

type BandPrecheckRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s BandPrecheckRequest) String() string {
	return tea.Prettify(s)
}

func (s BandPrecheckRequest) GoString() string {
	return s.String()
}

func (s *BandPrecheckRequest) SetOwnerId(v int64) *BandPrecheckRequest {
	s.OwnerId = &v
	return s
}

func (s *BandPrecheckRequest) SetResourceOwnerAccount(v string) *BandPrecheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BandPrecheckRequest) SetResourceOwnerId(v int64) *BandPrecheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BandPrecheckRequest) SetIpAddress(v string) *BandPrecheckRequest {
	s.IpAddress = &v
	return s
}

func (s *BandPrecheckRequest) SetPort(v int32) *BandPrecheckRequest {
	s.Port = &v
	return s
}

type BandPrecheckResponseBody struct {
	ResultModule  *BandPrecheckResponseBodyResultModule `json:"ResultModule,omitempty" xml:"ResultModule,omitempty" type:"Struct"`
	RequestId     *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string                               `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string                               `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s BandPrecheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandPrecheckResponseBody) GoString() string {
	return s.String()
}

func (s *BandPrecheckResponseBody) SetResultModule(v *BandPrecheckResponseBodyResultModule) *BandPrecheckResponseBody {
	s.ResultModule = v
	return s
}

func (s *BandPrecheckResponseBody) SetRequestId(v string) *BandPrecheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandPrecheckResponseBody) SetResultMessage(v string) *BandPrecheckResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *BandPrecheckResponseBody) SetResultCode(v string) *BandPrecheckResponseBody {
	s.ResultCode = &v
	return s
}

type BandPrecheckResponseBodyResultModule struct {
	UploadBandwidth   *int32                                             `json:"UploadBandwidth,omitempty" xml:"UploadBandwidth,omitempty"`
	BandId            *int64                                             `json:"BandId,omitempty" xml:"BandId,omitempty"`
	BandOfferList     *BandPrecheckResponseBodyResultModuleBandOfferList `json:"BandOfferList,omitempty" xml:"BandOfferList,omitempty" type:"Struct"`
	DownloadBandwidth *int32                                             `json:"DownloadBandwidth,omitempty" xml:"DownloadBandwidth,omitempty"`
}

func (s BandPrecheckResponseBodyResultModule) String() string {
	return tea.Prettify(s)
}

func (s BandPrecheckResponseBodyResultModule) GoString() string {
	return s.String()
}

func (s *BandPrecheckResponseBodyResultModule) SetUploadBandwidth(v int32) *BandPrecheckResponseBodyResultModule {
	s.UploadBandwidth = &v
	return s
}

func (s *BandPrecheckResponseBodyResultModule) SetBandId(v int64) *BandPrecheckResponseBodyResultModule {
	s.BandId = &v
	return s
}

func (s *BandPrecheckResponseBodyResultModule) SetBandOfferList(v *BandPrecheckResponseBodyResultModuleBandOfferList) *BandPrecheckResponseBodyResultModule {
	s.BandOfferList = v
	return s
}

func (s *BandPrecheckResponseBodyResultModule) SetDownloadBandwidth(v int32) *BandPrecheckResponseBodyResultModule {
	s.DownloadBandwidth = &v
	return s
}

type BandPrecheckResponseBodyResultModuleBandOfferList struct {
	BandOfferList []*BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList `json:"BandOfferList,omitempty" xml:"BandOfferList,omitempty" type:"Repeated"`
}

func (s BandPrecheckResponseBodyResultModuleBandOfferList) String() string {
	return tea.Prettify(s)
}

func (s BandPrecheckResponseBodyResultModuleBandOfferList) GoString() string {
	return s.String()
}

func (s *BandPrecheckResponseBodyResultModuleBandOfferList) SetBandOfferList(v []*BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) *BandPrecheckResponseBodyResultModuleBandOfferList {
	s.BandOfferList = v
	return s
}

type BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	OfferId   *int64  `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
	Bandwidth *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) String() string {
	return tea.Prettify(s)
}

func (s BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) GoString() string {
	return s.String()
}

func (s *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) SetDirection(v string) *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList {
	s.Direction = &v
	return s
}

func (s *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) SetOfferId(v int64) *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList {
	s.OfferId = &v
	return s
}

func (s *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) SetBandwidth(v int32) *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList {
	s.Bandwidth = &v
	return s
}

func (s *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList) SetDuration(v int64) *BandPrecheckResponseBodyResultModuleBandOfferListBandOfferList {
	s.Duration = &v
	return s
}

type BandPrecheckResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandPrecheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandPrecheckResponse) String() string {
	return tea.Prettify(s)
}

func (s BandPrecheckResponse) GoString() string {
	return s.String()
}

func (s *BandPrecheckResponse) SetHeaders(v map[string]*string) *BandPrecheckResponse {
	s.Headers = v
	return s
}

func (s *BandPrecheckResponse) SetBody(v *BandPrecheckResponseBody) *BandPrecheckResponse {
	s.Body = v
	return s
}

type BandStartSpeedUpRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	BandId               *int64  `json:"BandId,omitempty" xml:"BandId,omitempty"`
	Direction            *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	TargetBandwidth      *int64  `json:"TargetBandwidth,omitempty" xml:"TargetBandwidth,omitempty"`
	BandScene            *string `json:"BandScene,omitempty" xml:"BandScene,omitempty"`
}

func (s BandStartSpeedUpRequest) String() string {
	return tea.Prettify(s)
}

func (s BandStartSpeedUpRequest) GoString() string {
	return s.String()
}

func (s *BandStartSpeedUpRequest) SetOwnerId(v int64) *BandStartSpeedUpRequest {
	s.OwnerId = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetResourceOwnerAccount(v string) *BandStartSpeedUpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetResourceOwnerId(v int64) *BandStartSpeedUpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetIpAddress(v string) *BandStartSpeedUpRequest {
	s.IpAddress = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetPort(v int32) *BandStartSpeedUpRequest {
	s.Port = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetBandId(v int64) *BandStartSpeedUpRequest {
	s.BandId = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetDirection(v string) *BandStartSpeedUpRequest {
	s.Direction = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetTargetBandwidth(v int64) *BandStartSpeedUpRequest {
	s.TargetBandwidth = &v
	return s
}

func (s *BandStartSpeedUpRequest) SetBandScene(v string) *BandStartSpeedUpRequest {
	s.BandScene = &v
	return s
}

type BandStartSpeedUpResponseBody struct {
	ResultModule  *bool   `json:"ResultModule,omitempty" xml:"ResultModule,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s BandStartSpeedUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandStartSpeedUpResponseBody) GoString() string {
	return s.String()
}

func (s *BandStartSpeedUpResponseBody) SetResultModule(v bool) *BandStartSpeedUpResponseBody {
	s.ResultModule = &v
	return s
}

func (s *BandStartSpeedUpResponseBody) SetRequestId(v string) *BandStartSpeedUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandStartSpeedUpResponseBody) SetResultMessage(v string) *BandStartSpeedUpResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *BandStartSpeedUpResponseBody) SetResultCode(v string) *BandStartSpeedUpResponseBody {
	s.ResultCode = &v
	return s
}

type BandStartSpeedUpResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandStartSpeedUpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandStartSpeedUpResponse) String() string {
	return tea.Prettify(s)
}

func (s BandStartSpeedUpResponse) GoString() string {
	return s.String()
}

func (s *BandStartSpeedUpResponse) SetHeaders(v map[string]*string) *BandStartSpeedUpResponse {
	s.Headers = v
	return s
}

func (s *BandStartSpeedUpResponse) SetBody(v *BandStartSpeedUpResponseBody) *BandStartSpeedUpResponse {
	s.Body = v
	return s
}

type BandStatusQueryRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	BandId               *int64  `json:"BandId,omitempty" xml:"BandId,omitempty"`
}

func (s BandStatusQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s BandStatusQueryRequest) GoString() string {
	return s.String()
}

func (s *BandStatusQueryRequest) SetOwnerId(v int64) *BandStatusQueryRequest {
	s.OwnerId = &v
	return s
}

func (s *BandStatusQueryRequest) SetResourceOwnerAccount(v string) *BandStatusQueryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BandStatusQueryRequest) SetResourceOwnerId(v int64) *BandStatusQueryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BandStatusQueryRequest) SetBandId(v int64) *BandStatusQueryRequest {
	s.BandId = &v
	return s
}

type BandStatusQueryResponseBody struct {
	ResultModule  *BandStatusQueryResponseBodyResultModule `json:"ResultModule,omitempty" xml:"ResultModule,omitempty" type:"Struct"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s BandStatusQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandStatusQueryResponseBody) GoString() string {
	return s.String()
}

func (s *BandStatusQueryResponseBody) SetResultModule(v *BandStatusQueryResponseBodyResultModule) *BandStatusQueryResponseBody {
	s.ResultModule = v
	return s
}

func (s *BandStatusQueryResponseBody) SetRequestId(v string) *BandStatusQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandStatusQueryResponseBody) SetResultMessage(v string) *BandStatusQueryResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *BandStatusQueryResponseBody) SetResultCode(v string) *BandStatusQueryResponseBody {
	s.ResultCode = &v
	return s
}

type BandStatusQueryResponseBodyResultModule struct {
	UploadTarget   *int32 `json:"UploadTarget,omitempty" xml:"UploadTarget,omitempty"`
	DownloadTarget *int32 `json:"DownloadTarget,omitempty" xml:"DownloadTarget,omitempty"`
}

func (s BandStatusQueryResponseBodyResultModule) String() string {
	return tea.Prettify(s)
}

func (s BandStatusQueryResponseBodyResultModule) GoString() string {
	return s.String()
}

func (s *BandStatusQueryResponseBodyResultModule) SetUploadTarget(v int32) *BandStatusQueryResponseBodyResultModule {
	s.UploadTarget = &v
	return s
}

func (s *BandStatusQueryResponseBodyResultModule) SetDownloadTarget(v int32) *BandStatusQueryResponseBodyResultModule {
	s.DownloadTarget = &v
	return s
}

type BandStatusQueryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandStatusQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandStatusQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s BandStatusQueryResponse) GoString() string {
	return s.String()
}

func (s *BandStatusQueryResponse) SetHeaders(v map[string]*string) *BandStatusQueryResponse {
	s.Headers = v
	return s
}

func (s *BandStatusQueryResponse) SetBody(v *BandStatusQueryResponseBody) *BandStatusQueryResponse {
	s.Body = v
	return s
}

type BandStopSpeedUpRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	BandId               *int64  `json:"BandId,omitempty" xml:"BandId,omitempty"`
	Direction            *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
}

func (s BandStopSpeedUpRequest) String() string {
	return tea.Prettify(s)
}

func (s BandStopSpeedUpRequest) GoString() string {
	return s.String()
}

func (s *BandStopSpeedUpRequest) SetOwnerId(v int64) *BandStopSpeedUpRequest {
	s.OwnerId = &v
	return s
}

func (s *BandStopSpeedUpRequest) SetResourceOwnerAccount(v string) *BandStopSpeedUpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BandStopSpeedUpRequest) SetResourceOwnerId(v int64) *BandStopSpeedUpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BandStopSpeedUpRequest) SetIpAddress(v string) *BandStopSpeedUpRequest {
	s.IpAddress = &v
	return s
}

func (s *BandStopSpeedUpRequest) SetPort(v int32) *BandStopSpeedUpRequest {
	s.Port = &v
	return s
}

func (s *BandStopSpeedUpRequest) SetBandId(v int64) *BandStopSpeedUpRequest {
	s.BandId = &v
	return s
}

func (s *BandStopSpeedUpRequest) SetDirection(v string) *BandStopSpeedUpRequest {
	s.Direction = &v
	return s
}

type BandStopSpeedUpResponseBody struct {
	ResultModule  *bool   `json:"ResultModule,omitempty" xml:"ResultModule,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s BandStopSpeedUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandStopSpeedUpResponseBody) GoString() string {
	return s.String()
}

func (s *BandStopSpeedUpResponseBody) SetResultModule(v bool) *BandStopSpeedUpResponseBody {
	s.ResultModule = &v
	return s
}

func (s *BandStopSpeedUpResponseBody) SetRequestId(v string) *BandStopSpeedUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandStopSpeedUpResponseBody) SetResultMessage(v string) *BandStopSpeedUpResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *BandStopSpeedUpResponseBody) SetResultCode(v string) *BandStopSpeedUpResponseBody {
	s.ResultCode = &v
	return s
}

type BandStopSpeedUpResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandStopSpeedUpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandStopSpeedUpResponse) String() string {
	return tea.Prettify(s)
}

func (s BandStopSpeedUpResponse) GoString() string {
	return s.String()
}

func (s *BandStopSpeedUpResponse) SetHeaders(v map[string]*string) *BandStopSpeedUpResponse {
	s.Headers = v
	return s
}

func (s *BandStopSpeedUpResponse) SetBody(v *BandStopSpeedUpResponseBody) *BandStopSpeedUpResponse {
	s.Body = v
	return s
}

type MobileStartSpeedUpRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Token                *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Ip                   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PublicIp             *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	PublicPort           *string `json:"PublicPort,omitempty" xml:"PublicPort,omitempty"`
	DestinationIpAddress *string `json:"DestinationIpAddress,omitempty" xml:"DestinationIpAddress,omitempty"`
}

func (s MobileStartSpeedUpRequest) String() string {
	return tea.Prettify(s)
}

func (s MobileStartSpeedUpRequest) GoString() string {
	return s.String()
}

func (s *MobileStartSpeedUpRequest) SetOwnerId(v int64) *MobileStartSpeedUpRequest {
	s.OwnerId = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetResourceOwnerAccount(v string) *MobileStartSpeedUpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetResourceOwnerId(v int64) *MobileStartSpeedUpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetToken(v string) *MobileStartSpeedUpRequest {
	s.Token = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetDuration(v string) *MobileStartSpeedUpRequest {
	s.Duration = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetIp(v string) *MobileStartSpeedUpRequest {
	s.Ip = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetPublicIp(v string) *MobileStartSpeedUpRequest {
	s.PublicIp = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetPublicPort(v string) *MobileStartSpeedUpRequest {
	s.PublicPort = &v
	return s
}

func (s *MobileStartSpeedUpRequest) SetDestinationIpAddress(v string) *MobileStartSpeedUpRequest {
	s.DestinationIpAddress = &v
	return s
}

type MobileStartSpeedUpResponseBody struct {
	ResultModule  *string `json:"ResultModule,omitempty" xml:"ResultModule,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s MobileStartSpeedUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MobileStartSpeedUpResponseBody) GoString() string {
	return s.String()
}

func (s *MobileStartSpeedUpResponseBody) SetResultModule(v string) *MobileStartSpeedUpResponseBody {
	s.ResultModule = &v
	return s
}

func (s *MobileStartSpeedUpResponseBody) SetRequestId(v string) *MobileStartSpeedUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileStartSpeedUpResponseBody) SetResultMessage(v string) *MobileStartSpeedUpResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *MobileStartSpeedUpResponseBody) SetResultCode(v string) *MobileStartSpeedUpResponseBody {
	s.ResultCode = &v
	return s
}

type MobileStartSpeedUpResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MobileStartSpeedUpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MobileStartSpeedUpResponse) String() string {
	return tea.Prettify(s)
}

func (s MobileStartSpeedUpResponse) GoString() string {
	return s.String()
}

func (s *MobileStartSpeedUpResponse) SetHeaders(v map[string]*string) *MobileStartSpeedUpResponse {
	s.Headers = v
	return s
}

func (s *MobileStartSpeedUpResponse) SetBody(v *MobileStartSpeedUpResponseBody) *MobileStartSpeedUpResponse {
	s.Body = v
	return s
}

type MobileStatusQueryRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CorrelationId        *string `json:"CorrelationId,omitempty" xml:"CorrelationId,omitempty"`
}

func (s MobileStatusQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MobileStatusQueryRequest) GoString() string {
	return s.String()
}

func (s *MobileStatusQueryRequest) SetOwnerId(v int64) *MobileStatusQueryRequest {
	s.OwnerId = &v
	return s
}

func (s *MobileStatusQueryRequest) SetResourceOwnerAccount(v string) *MobileStatusQueryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MobileStatusQueryRequest) SetResourceOwnerId(v int64) *MobileStatusQueryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MobileStatusQueryRequest) SetCorrelationId(v string) *MobileStatusQueryRequest {
	s.CorrelationId = &v
	return s
}

type MobileStatusQueryResponseBody struct {
	ResultModule  *bool   `json:"ResultModule,omitempty" xml:"ResultModule,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s MobileStatusQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MobileStatusQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MobileStatusQueryResponseBody) SetResultModule(v bool) *MobileStatusQueryResponseBody {
	s.ResultModule = &v
	return s
}

func (s *MobileStatusQueryResponseBody) SetRequestId(v string) *MobileStatusQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileStatusQueryResponseBody) SetResultMessage(v string) *MobileStatusQueryResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *MobileStatusQueryResponseBody) SetResultCode(v string) *MobileStatusQueryResponseBody {
	s.ResultCode = &v
	return s
}

type MobileStatusQueryResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MobileStatusQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MobileStatusQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MobileStatusQueryResponse) GoString() string {
	return s.String()
}

func (s *MobileStatusQueryResponse) SetHeaders(v map[string]*string) *MobileStatusQueryResponse {
	s.Headers = v
	return s
}

func (s *MobileStatusQueryResponse) SetBody(v *MobileStatusQueryResponseBody) *MobileStatusQueryResponse {
	s.Body = v
	return s
}

type MobileStopSpeedUpRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CorrelationId        *string `json:"CorrelationId,omitempty" xml:"CorrelationId,omitempty"`
}

func (s MobileStopSpeedUpRequest) String() string {
	return tea.Prettify(s)
}

func (s MobileStopSpeedUpRequest) GoString() string {
	return s.String()
}

func (s *MobileStopSpeedUpRequest) SetOwnerId(v int64) *MobileStopSpeedUpRequest {
	s.OwnerId = &v
	return s
}

func (s *MobileStopSpeedUpRequest) SetResourceOwnerAccount(v string) *MobileStopSpeedUpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MobileStopSpeedUpRequest) SetResourceOwnerId(v int64) *MobileStopSpeedUpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MobileStopSpeedUpRequest) SetCorrelationId(v string) *MobileStopSpeedUpRequest {
	s.CorrelationId = &v
	return s
}

type MobileStopSpeedUpResponseBody struct {
	ResultModule  *bool   `json:"ResultModule,omitempty" xml:"ResultModule,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
}

func (s MobileStopSpeedUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MobileStopSpeedUpResponseBody) GoString() string {
	return s.String()
}

func (s *MobileStopSpeedUpResponseBody) SetResultModule(v bool) *MobileStopSpeedUpResponseBody {
	s.ResultModule = &v
	return s
}

func (s *MobileStopSpeedUpResponseBody) SetRequestId(v string) *MobileStopSpeedUpResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileStopSpeedUpResponseBody) SetResultMessage(v string) *MobileStopSpeedUpResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *MobileStopSpeedUpResponseBody) SetResultCode(v string) *MobileStopSpeedUpResponseBody {
	s.ResultCode = &v
	return s
}

type MobileStopSpeedUpResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MobileStopSpeedUpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MobileStopSpeedUpResponse) String() string {
	return tea.Prettify(s)
}

func (s MobileStopSpeedUpResponse) GoString() string {
	return s.String()
}

func (s *MobileStopSpeedUpResponse) SetHeaders(v map[string]*string) *MobileStopSpeedUpResponse {
	s.Headers = v
	return s
}

func (s *MobileStopSpeedUpResponse) SetBody(v *MobileStopSpeedUpResponseBody) *MobileStopSpeedUpResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("snsuapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BandOfferOrderWithOptions(request *BandOfferOrderRequest, runtime *util.RuntimeOptions) (_result *BandOfferOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandOfferOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandOfferOrder"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandOfferOrder(request *BandOfferOrderRequest) (_result *BandOfferOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandOfferOrderResponse{}
	_body, _err := client.BandOfferOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BandPrecheckWithOptions(request *BandPrecheckRequest, runtime *util.RuntimeOptions) (_result *BandPrecheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandPrecheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandPrecheck"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandPrecheck(request *BandPrecheckRequest) (_result *BandPrecheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandPrecheckResponse{}
	_body, _err := client.BandPrecheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BandStartSpeedUpWithOptions(request *BandStartSpeedUpRequest, runtime *util.RuntimeOptions) (_result *BandStartSpeedUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandStartSpeedUpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandStartSpeedUp"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandStartSpeedUp(request *BandStartSpeedUpRequest) (_result *BandStartSpeedUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandStartSpeedUpResponse{}
	_body, _err := client.BandStartSpeedUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BandStatusQueryWithOptions(request *BandStatusQueryRequest, runtime *util.RuntimeOptions) (_result *BandStatusQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandStatusQueryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandStatusQuery"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandStatusQuery(request *BandStatusQueryRequest) (_result *BandStatusQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandStatusQueryResponse{}
	_body, _err := client.BandStatusQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BandStopSpeedUpWithOptions(request *BandStopSpeedUpRequest, runtime *util.RuntimeOptions) (_result *BandStopSpeedUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandStopSpeedUpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandStopSpeedUp"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandStopSpeedUp(request *BandStopSpeedUpRequest) (_result *BandStopSpeedUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandStopSpeedUpResponse{}
	_body, _err := client.BandStopSpeedUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MobileStartSpeedUpWithOptions(request *MobileStartSpeedUpRequest, runtime *util.RuntimeOptions) (_result *MobileStartSpeedUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MobileStartSpeedUpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MobileStartSpeedUp"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MobileStartSpeedUp(request *MobileStartSpeedUpRequest) (_result *MobileStartSpeedUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MobileStartSpeedUpResponse{}
	_body, _err := client.MobileStartSpeedUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MobileStatusQueryWithOptions(request *MobileStatusQueryRequest, runtime *util.RuntimeOptions) (_result *MobileStatusQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MobileStatusQueryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MobileStatusQuery"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MobileStatusQuery(request *MobileStatusQueryRequest) (_result *MobileStatusQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MobileStatusQueryResponse{}
	_body, _err := client.MobileStatusQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MobileStopSpeedUpWithOptions(request *MobileStopSpeedUpRequest, runtime *util.RuntimeOptions) (_result *MobileStopSpeedUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MobileStopSpeedUpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MobileStopSpeedUp"), tea.String("2018-07-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MobileStopSpeedUp(request *MobileStopSpeedUpRequest) (_result *MobileStopSpeedUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MobileStopSpeedUpResponse{}
	_body, _err := client.MobileStopSpeedUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

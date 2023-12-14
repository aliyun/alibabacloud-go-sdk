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

type ActivateAllInOneGatewayRequest struct {
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	DeviceNumber  *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Model         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s ActivateAllInOneGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateAllInOneGatewayRequest) GoString() string {
	return s.String()
}

func (s *ActivateAllInOneGatewayRequest) SetClientUUID(v string) *ActivateAllInOneGatewayRequest {
	s.ClientUUID = &v
	return s
}

func (s *ActivateAllInOneGatewayRequest) SetDeviceNumber(v string) *ActivateAllInOneGatewayRequest {
	s.DeviceNumber = &v
	return s
}

func (s *ActivateAllInOneGatewayRequest) SetGatewayId(v string) *ActivateAllInOneGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ActivateAllInOneGatewayRequest) SetModel(v string) *ActivateAllInOneGatewayRequest {
	s.Model = &v
	return s
}

func (s *ActivateAllInOneGatewayRequest) SetSecurityToken(v string) *ActivateAllInOneGatewayRequest {
	s.SecurityToken = &v
	return s
}

func (s *ActivateAllInOneGatewayRequest) SetSerialNumber(v string) *ActivateAllInOneGatewayRequest {
	s.SerialNumber = &v
	return s
}

type ActivateAllInOneGatewayResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayId      *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	LicenseContent *string `json:"LicenseContent,omitempty" xml:"LicenseContent,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateAllInOneGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateAllInOneGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateAllInOneGatewayResponseBody) SetCode(v string) *ActivateAllInOneGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateAllInOneGatewayResponseBody) SetGatewayId(v string) *ActivateAllInOneGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ActivateAllInOneGatewayResponseBody) SetLicenseContent(v string) *ActivateAllInOneGatewayResponseBody {
	s.LicenseContent = &v
	return s
}

func (s *ActivateAllInOneGatewayResponseBody) SetMessage(v string) *ActivateAllInOneGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateAllInOneGatewayResponseBody) SetRegionId(v string) *ActivateAllInOneGatewayResponseBody {
	s.RegionId = &v
	return s
}

func (s *ActivateAllInOneGatewayResponseBody) SetRequestId(v string) *ActivateAllInOneGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateAllInOneGatewayResponseBody) SetSuccess(v bool) *ActivateAllInOneGatewayResponseBody {
	s.Success = &v
	return s
}

type ActivateAllInOneGatewayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActivateAllInOneGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateAllInOneGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateAllInOneGatewayResponse) GoString() string {
	return s.String()
}

func (s *ActivateAllInOneGatewayResponse) SetHeaders(v map[string]*string) *ActivateAllInOneGatewayResponse {
	s.Headers = v
	return s
}

func (s *ActivateAllInOneGatewayResponse) SetStatusCode(v int32) *ActivateAllInOneGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateAllInOneGatewayResponse) SetBody(v *ActivateAllInOneGatewayResponseBody) *ActivateAllInOneGatewayResponse {
	s.Body = v
	return s
}

type ActivateGatewayRequest struct {
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	Model         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SerialNumber  *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ActivateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateGatewayRequest) GoString() string {
	return s.String()
}

func (s *ActivateGatewayRequest) SetCategory(v string) *ActivateGatewayRequest {
	s.Category = &v
	return s
}

func (s *ActivateGatewayRequest) SetClientUUID(v string) *ActivateGatewayRequest {
	s.ClientUUID = &v
	return s
}

func (s *ActivateGatewayRequest) SetModel(v string) *ActivateGatewayRequest {
	s.Model = &v
	return s
}

func (s *ActivateGatewayRequest) SetSecurityToken(v string) *ActivateGatewayRequest {
	s.SecurityToken = &v
	return s
}

func (s *ActivateGatewayRequest) SetSerialNumber(v string) *ActivateGatewayRequest {
	s.SerialNumber = &v
	return s
}

func (s *ActivateGatewayRequest) SetToken(v string) *ActivateGatewayRequest {
	s.Token = &v
	return s
}

func (s *ActivateGatewayRequest) SetType(v string) *ActivateGatewayRequest {
	s.Type = &v
	return s
}

type ActivateGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateGatewayResponseBody) SetCode(v string) *ActivateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateGatewayResponseBody) SetGatewayId(v string) *ActivateGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ActivateGatewayResponseBody) SetMessage(v string) *ActivateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateGatewayResponseBody) SetRegionId(v string) *ActivateGatewayResponseBody {
	s.RegionId = &v
	return s
}

func (s *ActivateGatewayResponseBody) SetRequestId(v string) *ActivateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateGatewayResponseBody) SetSuccess(v bool) *ActivateGatewayResponseBody {
	s.Success = &v
	return s
}

type ActivateGatewayResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActivateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateGatewayResponse) GoString() string {
	return s.String()
}

func (s *ActivateGatewayResponse) SetHeaders(v map[string]*string) *ActivateGatewayResponse {
	s.Headers = v
	return s
}

func (s *ActivateGatewayResponse) SetStatusCode(v int32) *ActivateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateGatewayResponse) SetBody(v *ActivateGatewayResponseBody) *ActivateGatewayResponse {
	s.Body = v
	return s
}

type AddSharesToExpressSyncRequest struct {
	ExpressSyncId *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	GatewayShares *string `json:"GatewayShares,omitempty" xml:"GatewayShares,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddSharesToExpressSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSharesToExpressSyncRequest) GoString() string {
	return s.String()
}

func (s *AddSharesToExpressSyncRequest) SetExpressSyncId(v string) *AddSharesToExpressSyncRequest {
	s.ExpressSyncId = &v
	return s
}

func (s *AddSharesToExpressSyncRequest) SetGatewayShares(v string) *AddSharesToExpressSyncRequest {
	s.GatewayShares = &v
	return s
}

func (s *AddSharesToExpressSyncRequest) SetSecurityToken(v string) *AddSharesToExpressSyncRequest {
	s.SecurityToken = &v
	return s
}

type AddSharesToExpressSyncResponseBody struct {
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MnsFullSyncDelay  *int64  `json:"MnsFullSyncDelay,omitempty" xml:"MnsFullSyncDelay,omitempty"`
	MnsInnerEndpoint  *string `json:"MnsInnerEndpoint,omitempty" xml:"MnsInnerEndpoint,omitempty"`
	MnsPublicEndpoint *string `json:"MnsPublicEndpoint,omitempty" xml:"MnsPublicEndpoint,omitempty"`
	MnsQueues         *string `json:"MnsQueues,omitempty" xml:"MnsQueues,omitempty"`
	MnsTopic          *string `json:"MnsTopic,omitempty" xml:"MnsTopic,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddSharesToExpressSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSharesToExpressSyncResponseBody) GoString() string {
	return s.String()
}

func (s *AddSharesToExpressSyncResponseBody) SetCode(v string) *AddSharesToExpressSyncResponseBody {
	s.Code = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetMessage(v string) *AddSharesToExpressSyncResponseBody {
	s.Message = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetMnsFullSyncDelay(v int64) *AddSharesToExpressSyncResponseBody {
	s.MnsFullSyncDelay = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetMnsInnerEndpoint(v string) *AddSharesToExpressSyncResponseBody {
	s.MnsInnerEndpoint = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetMnsPublicEndpoint(v string) *AddSharesToExpressSyncResponseBody {
	s.MnsPublicEndpoint = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetMnsQueues(v string) *AddSharesToExpressSyncResponseBody {
	s.MnsQueues = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetMnsTopic(v string) *AddSharesToExpressSyncResponseBody {
	s.MnsTopic = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetRequestId(v string) *AddSharesToExpressSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetSuccess(v bool) *AddSharesToExpressSyncResponseBody {
	s.Success = &v
	return s
}

func (s *AddSharesToExpressSyncResponseBody) SetTaskId(v string) *AddSharesToExpressSyncResponseBody {
	s.TaskId = &v
	return s
}

type AddSharesToExpressSyncResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSharesToExpressSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSharesToExpressSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSharesToExpressSyncResponse) GoString() string {
	return s.String()
}

func (s *AddSharesToExpressSyncResponse) SetHeaders(v map[string]*string) *AddSharesToExpressSyncResponse {
	s.Headers = v
	return s
}

func (s *AddSharesToExpressSyncResponse) SetStatusCode(v int32) *AddSharesToExpressSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSharesToExpressSyncResponse) SetBody(v *AddSharesToExpressSyncResponseBody) *AddSharesToExpressSyncResponse {
	s.Body = v
	return s
}

type AddTagsToGatewayRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddTagsToGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToGatewayRequest) GoString() string {
	return s.String()
}

func (s *AddTagsToGatewayRequest) SetGatewayId(v string) *AddTagsToGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *AddTagsToGatewayRequest) SetSecurityToken(v string) *AddTagsToGatewayRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddTagsToGatewayRequest) SetTags(v string) *AddTagsToGatewayRequest {
	s.Tags = &v
	return s
}

type AddTagsToGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTagsToGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsToGatewayResponseBody) SetCode(v string) *AddTagsToGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *AddTagsToGatewayResponseBody) SetMessage(v string) *AddTagsToGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *AddTagsToGatewayResponseBody) SetRequestId(v string) *AddTagsToGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagsToGatewayResponseBody) SetSuccess(v bool) *AddTagsToGatewayResponseBody {
	s.Success = &v
	return s
}

type AddTagsToGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTagsToGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsToGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToGatewayResponse) GoString() string {
	return s.String()
}

func (s *AddTagsToGatewayResponse) SetHeaders(v map[string]*string) *AddTagsToGatewayResponse {
	s.Headers = v
	return s
}

func (s *AddTagsToGatewayResponse) SetStatusCode(v int32) *AddTagsToGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsToGatewayResponse) SetBody(v *AddTagsToGatewayResponseBody) *AddTagsToGatewayResponse {
	s.Body = v
	return s
}

type CheckActivationKeyRequest struct {
	CryptKey      *string `json:"CryptKey,omitempty" xml:"CryptKey,omitempty"`
	CryptText     *string `json:"CryptText,omitempty" xml:"CryptText,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CheckActivationKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckActivationKeyRequest) GoString() string {
	return s.String()
}

func (s *CheckActivationKeyRequest) SetCryptKey(v string) *CheckActivationKeyRequest {
	s.CryptKey = &v
	return s
}

func (s *CheckActivationKeyRequest) SetCryptText(v string) *CheckActivationKeyRequest {
	s.CryptText = &v
	return s
}

func (s *CheckActivationKeyRequest) SetGatewayId(v string) *CheckActivationKeyRequest {
	s.GatewayId = &v
	return s
}

func (s *CheckActivationKeyRequest) SetSecurityToken(v string) *CheckActivationKeyRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckActivationKeyRequest) SetToken(v string) *CheckActivationKeyRequest {
	s.Token = &v
	return s
}

type CheckActivationKeyResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckActivationKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckActivationKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckActivationKeyResponseBody) SetCode(v string) *CheckActivationKeyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckActivationKeyResponseBody) SetMessage(v string) *CheckActivationKeyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckActivationKeyResponseBody) SetRequestId(v string) *CheckActivationKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckActivationKeyResponseBody) SetSuccess(v bool) *CheckActivationKeyResponseBody {
	s.Success = &v
	return s
}

type CheckActivationKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckActivationKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckActivationKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckActivationKeyResponse) GoString() string {
	return s.String()
}

func (s *CheckActivationKeyResponse) SetHeaders(v map[string]*string) *CheckActivationKeyResponse {
	s.Headers = v
	return s
}

func (s *CheckActivationKeyResponse) SetStatusCode(v int32) *CheckActivationKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckActivationKeyResponse) SetBody(v *CheckActivationKeyResponseBody) *CheckActivationKeyResponse {
	s.Body = v
	return s
}

type CheckBlockVolumeNameRequest struct {
	// Bucket Endpoint。
	BucketEndpoint *string `json:"BucketEndpoint,omitempty" xml:"BucketEndpoint,omitempty"`
	BucketName     *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VolumeName     *string `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
}

func (s CheckBlockVolumeNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckBlockVolumeNameRequest) GoString() string {
	return s.String()
}

func (s *CheckBlockVolumeNameRequest) SetBucketEndpoint(v string) *CheckBlockVolumeNameRequest {
	s.BucketEndpoint = &v
	return s
}

func (s *CheckBlockVolumeNameRequest) SetBucketName(v string) *CheckBlockVolumeNameRequest {
	s.BucketName = &v
	return s
}

func (s *CheckBlockVolumeNameRequest) SetSecurityToken(v string) *CheckBlockVolumeNameRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckBlockVolumeNameRequest) SetVolumeName(v string) *CheckBlockVolumeNameRequest {
	s.VolumeName = &v
	return s
}

type CheckBlockVolumeNameResponseBody struct {
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsAlreadyExist    *bool   `json:"IsAlreadyExist,omitempty" xml:"IsAlreadyExist,omitempty"`
	IsRequireRecovery *string `json:"IsRequireRecovery,omitempty" xml:"IsRequireRecovery,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckBlockVolumeNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckBlockVolumeNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBlockVolumeNameResponseBody) SetCode(v string) *CheckBlockVolumeNameResponseBody {
	s.Code = &v
	return s
}

func (s *CheckBlockVolumeNameResponseBody) SetIsAlreadyExist(v bool) *CheckBlockVolumeNameResponseBody {
	s.IsAlreadyExist = &v
	return s
}

func (s *CheckBlockVolumeNameResponseBody) SetIsRequireRecovery(v string) *CheckBlockVolumeNameResponseBody {
	s.IsRequireRecovery = &v
	return s
}

func (s *CheckBlockVolumeNameResponseBody) SetMessage(v string) *CheckBlockVolumeNameResponseBody {
	s.Message = &v
	return s
}

func (s *CheckBlockVolumeNameResponseBody) SetRequestId(v string) *CheckBlockVolumeNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBlockVolumeNameResponseBody) SetSuccess(v bool) *CheckBlockVolumeNameResponseBody {
	s.Success = &v
	return s
}

type CheckBlockVolumeNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckBlockVolumeNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckBlockVolumeNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckBlockVolumeNameResponse) GoString() string {
	return s.String()
}

func (s *CheckBlockVolumeNameResponse) SetHeaders(v map[string]*string) *CheckBlockVolumeNameResponse {
	s.Headers = v
	return s
}

func (s *CheckBlockVolumeNameResponse) SetStatusCode(v int32) *CheckBlockVolumeNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBlockVolumeNameResponse) SetBody(v *CheckBlockVolumeNameResponseBody) *CheckBlockVolumeNameResponse {
	s.Body = v
	return s
}

type CheckGatewayEssdSupportRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckGatewayEssdSupportRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGatewayEssdSupportRequest) GoString() string {
	return s.String()
}

func (s *CheckGatewayEssdSupportRequest) SetGatewayId(v string) *CheckGatewayEssdSupportRequest {
	s.GatewayId = &v
	return s
}

func (s *CheckGatewayEssdSupportRequest) SetSecurityToken(v string) *CheckGatewayEssdSupportRequest {
	s.SecurityToken = &v
	return s
}

type CheckGatewayEssdSupportResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsServiceAffect *bool   `json:"IsServiceAffect,omitempty" xml:"IsServiceAffect,omitempty"`
	IsSupportEssd   *bool   `json:"IsSupportEssd,omitempty" xml:"IsSupportEssd,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckGatewayEssdSupportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckGatewayEssdSupportResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGatewayEssdSupportResponseBody) SetCode(v string) *CheckGatewayEssdSupportResponseBody {
	s.Code = &v
	return s
}

func (s *CheckGatewayEssdSupportResponseBody) SetIsServiceAffect(v bool) *CheckGatewayEssdSupportResponseBody {
	s.IsServiceAffect = &v
	return s
}

func (s *CheckGatewayEssdSupportResponseBody) SetIsSupportEssd(v bool) *CheckGatewayEssdSupportResponseBody {
	s.IsSupportEssd = &v
	return s
}

func (s *CheckGatewayEssdSupportResponseBody) SetMessage(v string) *CheckGatewayEssdSupportResponseBody {
	s.Message = &v
	return s
}

func (s *CheckGatewayEssdSupportResponseBody) SetRequestId(v string) *CheckGatewayEssdSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGatewayEssdSupportResponseBody) SetSuccess(v bool) *CheckGatewayEssdSupportResponseBody {
	s.Success = &v
	return s
}

type CheckGatewayEssdSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckGatewayEssdSupportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckGatewayEssdSupportResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGatewayEssdSupportResponse) GoString() string {
	return s.String()
}

func (s *CheckGatewayEssdSupportResponse) SetHeaders(v map[string]*string) *CheckGatewayEssdSupportResponse {
	s.Headers = v
	return s
}

func (s *CheckGatewayEssdSupportResponse) SetStatusCode(v int32) *CheckGatewayEssdSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGatewayEssdSupportResponse) SetBody(v *CheckGatewayEssdSupportResponseBody) *CheckGatewayEssdSupportResponse {
	s.Body = v
	return s
}

type CheckMnsServiceRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckMnsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMnsServiceRequest) GoString() string {
	return s.String()
}

func (s *CheckMnsServiceRequest) SetSecurityToken(v string) *CheckMnsServiceRequest {
	s.SecurityToken = &v
	return s
}

type CheckMnsServiceResponseBody struct {
	CheckMessage *string `json:"CheckMessage,omitempty" xml:"CheckMessage,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsEnabled    *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckMnsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMnsServiceResponseBody) SetCheckMessage(v string) *CheckMnsServiceResponseBody {
	s.CheckMessage = &v
	return s
}

func (s *CheckMnsServiceResponseBody) SetCode(v string) *CheckMnsServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CheckMnsServiceResponseBody) SetIsEnabled(v bool) *CheckMnsServiceResponseBody {
	s.IsEnabled = &v
	return s
}

func (s *CheckMnsServiceResponseBody) SetMessage(v string) *CheckMnsServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CheckMnsServiceResponseBody) SetRequestId(v string) *CheckMnsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMnsServiceResponseBody) SetSuccess(v bool) *CheckMnsServiceResponseBody {
	s.Success = &v
	return s
}

type CheckMnsServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckMnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMnsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMnsServiceResponse) GoString() string {
	return s.String()
}

func (s *CheckMnsServiceResponse) SetHeaders(v map[string]*string) *CheckMnsServiceResponse {
	s.Headers = v
	return s
}

func (s *CheckMnsServiceResponse) SetStatusCode(v int32) *CheckMnsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMnsServiceResponse) SetBody(v *CheckMnsServiceResponseBody) *CheckMnsServiceResponse {
	s.Body = v
	return s
}

type CheckRoleRequest struct {
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckRoleRequest) SetRoleType(v string) *CheckRoleRequest {
	s.RoleType = &v
	return s
}

func (s *CheckRoleRequest) SetSecurityToken(v string) *CheckRoleRequest {
	s.SecurityToken = &v
	return s
}

type CheckRoleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRoleResponseBody) SetCode(v string) *CheckRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CheckRoleResponseBody) SetMessage(v string) *CheckRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CheckRoleResponseBody) SetRequestId(v string) *CheckRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRoleResponseBody) SetSuccess(v bool) *CheckRoleResponseBody {
	s.Success = &v
	return s
}

type CheckRoleResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckRoleResponse) SetHeaders(v map[string]*string) *CheckRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckRoleResponse) SetStatusCode(v int32) *CheckRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRoleResponse) SetBody(v *CheckRoleResponseBody) *CheckRoleResponse {
	s.Body = v
	return s
}

type CheckSlrRoleRequest struct {
	CreateIfNotExist *bool   `json:"CreateIfNotExist,omitempty" xml:"CreateIfNotExist,omitempty"`
	RoleName         *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckSlrRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSlrRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckSlrRoleRequest) SetCreateIfNotExist(v bool) *CheckSlrRoleRequest {
	s.CreateIfNotExist = &v
	return s
}

func (s *CheckSlrRoleRequest) SetRoleName(v string) *CheckSlrRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CheckSlrRoleRequest) SetSecurityToken(v string) *CheckSlrRoleRequest {
	s.SecurityToken = &v
	return s
}

type CheckSlrRoleResponseBody struct {
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Exist              *bool   `json:"Exist,omitempty" xml:"Exist,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequireOldWayCheck *bool   `json:"RequireOldWayCheck,omitempty" xml:"RequireOldWayCheck,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSlrRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSlrRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSlrRoleResponseBody) SetCode(v string) *CheckSlrRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CheckSlrRoleResponseBody) SetExist(v bool) *CheckSlrRoleResponseBody {
	s.Exist = &v
	return s
}

func (s *CheckSlrRoleResponseBody) SetMessage(v string) *CheckSlrRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CheckSlrRoleResponseBody) SetRequestId(v string) *CheckSlrRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSlrRoleResponseBody) SetRequireOldWayCheck(v bool) *CheckSlrRoleResponseBody {
	s.RequireOldWayCheck = &v
	return s
}

func (s *CheckSlrRoleResponseBody) SetSuccess(v bool) *CheckSlrRoleResponseBody {
	s.Success = &v
	return s
}

type CheckSlrRoleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckSlrRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckSlrRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSlrRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckSlrRoleResponse) SetHeaders(v map[string]*string) *CheckSlrRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckSlrRoleResponse) SetStatusCode(v int32) *CheckSlrRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSlrRoleResponse) SetBody(v *CheckSlrRoleResponseBody) *CheckSlrRoleResponse {
	s.Body = v
	return s
}

type CheckUpgradeVersionRequest struct {
	ClientUUID     *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	GatewayId      *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayVersion *string `json:"GatewayVersion,omitempty" xml:"GatewayVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckUpgradeVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUpgradeVersionRequest) GoString() string {
	return s.String()
}

func (s *CheckUpgradeVersionRequest) SetClientUUID(v string) *CheckUpgradeVersionRequest {
	s.ClientUUID = &v
	return s
}

func (s *CheckUpgradeVersionRequest) SetGatewayId(v string) *CheckUpgradeVersionRequest {
	s.GatewayId = &v
	return s
}

func (s *CheckUpgradeVersionRequest) SetGatewayVersion(v string) *CheckUpgradeVersionRequest {
	s.GatewayVersion = &v
	return s
}

func (s *CheckUpgradeVersionRequest) SetSecurityToken(v string) *CheckUpgradeVersionRequest {
	s.SecurityToken = &v
	return s
}

type CheckUpgradeVersionResponseBody struct {
	Code          *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	LatestVersion *string                                 `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	Message       *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Option        *string                                 `json:"Option,omitempty" xml:"Option,omitempty"`
	Patches       *CheckUpgradeVersionResponseBodyPatches `json:"Patches,omitempty" xml:"Patches,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUpgradeVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUpgradeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUpgradeVersionResponseBody) SetCode(v string) *CheckUpgradeVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUpgradeVersionResponseBody) SetLatestVersion(v string) *CheckUpgradeVersionResponseBody {
	s.LatestVersion = &v
	return s
}

func (s *CheckUpgradeVersionResponseBody) SetMessage(v string) *CheckUpgradeVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUpgradeVersionResponseBody) SetOption(v string) *CheckUpgradeVersionResponseBody {
	s.Option = &v
	return s
}

func (s *CheckUpgradeVersionResponseBody) SetPatches(v *CheckUpgradeVersionResponseBodyPatches) *CheckUpgradeVersionResponseBody {
	s.Patches = v
	return s
}

func (s *CheckUpgradeVersionResponseBody) SetRequestId(v string) *CheckUpgradeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUpgradeVersionResponseBody) SetSuccess(v bool) *CheckUpgradeVersionResponseBody {
	s.Success = &v
	return s
}

type CheckUpgradeVersionResponseBodyPatches struct {
	Patch []*CheckUpgradeVersionResponseBodyPatchesPatch `json:"Patch,omitempty" xml:"Patch,omitempty" type:"Repeated"`
}

func (s CheckUpgradeVersionResponseBodyPatches) String() string {
	return tea.Prettify(s)
}

func (s CheckUpgradeVersionResponseBodyPatches) GoString() string {
	return s.String()
}

func (s *CheckUpgradeVersionResponseBodyPatches) SetPatch(v []*CheckUpgradeVersionResponseBodyPatchesPatch) *CheckUpgradeVersionResponseBodyPatches {
	s.Patch = v
	return s
}

type CheckUpgradeVersionResponseBodyPatchesPatch struct {
	InternalUrl *string `json:"InternalUrl,omitempty" xml:"InternalUrl,omitempty"`
	MD5         *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CheckUpgradeVersionResponseBodyPatchesPatch) String() string {
	return tea.Prettify(s)
}

func (s CheckUpgradeVersionResponseBodyPatchesPatch) GoString() string {
	return s.String()
}

func (s *CheckUpgradeVersionResponseBodyPatchesPatch) SetInternalUrl(v string) *CheckUpgradeVersionResponseBodyPatchesPatch {
	s.InternalUrl = &v
	return s
}

func (s *CheckUpgradeVersionResponseBodyPatchesPatch) SetMD5(v string) *CheckUpgradeVersionResponseBodyPatchesPatch {
	s.MD5 = &v
	return s
}

func (s *CheckUpgradeVersionResponseBodyPatchesPatch) SetName(v string) *CheckUpgradeVersionResponseBodyPatchesPatch {
	s.Name = &v
	return s
}

func (s *CheckUpgradeVersionResponseBodyPatchesPatch) SetUrl(v string) *CheckUpgradeVersionResponseBodyPatchesPatch {
	s.Url = &v
	return s
}

type CheckUpgradeVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckUpgradeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUpgradeVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUpgradeVersionResponse) GoString() string {
	return s.String()
}

func (s *CheckUpgradeVersionResponse) SetHeaders(v map[string]*string) *CheckUpgradeVersionResponse {
	s.Headers = v
	return s
}

func (s *CheckUpgradeVersionResponse) SetStatusCode(v int32) *CheckUpgradeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUpgradeVersionResponse) SetBody(v *CheckUpgradeVersionResponseBody) *CheckUpgradeVersionResponse {
	s.Body = v
	return s
}

type CreateCacheRequest struct {
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SizeInGB      *int64  `json:"SizeInGB,omitempty" xml:"SizeInGB,omitempty"`
}

func (s CreateCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheRequest) SetCategory(v string) *CreateCacheRequest {
	s.Category = &v
	return s
}

func (s *CreateCacheRequest) SetGatewayId(v string) *CreateCacheRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateCacheRequest) SetSecurityToken(v string) *CreateCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateCacheRequest) SetSizeInGB(v int64) *CreateCacheRequest {
	s.SizeInGB = &v
	return s
}

type CreateCacheResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	CacheId   *string `json:"CacheId,omitempty" xml:"CacheId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheResponseBody) SetBuyURL(v string) *CreateCacheResponseBody {
	s.BuyURL = &v
	return s
}

func (s *CreateCacheResponseBody) SetCacheId(v string) *CreateCacheResponseBody {
	s.CacheId = &v
	return s
}

func (s *CreateCacheResponseBody) SetCode(v string) *CreateCacheResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCacheResponseBody) SetMessage(v string) *CreateCacheResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCacheResponseBody) SetRequestId(v string) *CreateCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCacheResponseBody) SetSuccess(v bool) *CreateCacheResponseBody {
	s.Success = &v
	return s
}

type CreateCacheResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheResponse) SetHeaders(v map[string]*string) *CreateCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheResponse) SetStatusCode(v int32) *CreateCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCacheResponse) SetBody(v *CreateCacheResponseBody) *CreateCacheResponse {
	s.Body = v
	return s
}

type CreateElasticGatewayPrivateZoneRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateElasticGatewayPrivateZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticGatewayPrivateZoneRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticGatewayPrivateZoneRequest) SetGatewayId(v string) *CreateElasticGatewayPrivateZoneRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateElasticGatewayPrivateZoneRequest) SetSecurityToken(v string) *CreateElasticGatewayPrivateZoneRequest {
	s.SecurityToken = &v
	return s
}

type CreateElasticGatewayPrivateZoneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateElasticGatewayPrivateZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticGatewayPrivateZoneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticGatewayPrivateZoneResponseBody) SetCode(v string) *CreateElasticGatewayPrivateZoneResponseBody {
	s.Code = &v
	return s
}

func (s *CreateElasticGatewayPrivateZoneResponseBody) SetMessage(v string) *CreateElasticGatewayPrivateZoneResponseBody {
	s.Message = &v
	return s
}

func (s *CreateElasticGatewayPrivateZoneResponseBody) SetRequestId(v string) *CreateElasticGatewayPrivateZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateElasticGatewayPrivateZoneResponseBody) SetSuccess(v bool) *CreateElasticGatewayPrivateZoneResponseBody {
	s.Success = &v
	return s
}

func (s *CreateElasticGatewayPrivateZoneResponseBody) SetTaskId(v string) *CreateElasticGatewayPrivateZoneResponseBody {
	s.TaskId = &v
	return s
}

type CreateElasticGatewayPrivateZoneResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateElasticGatewayPrivateZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateElasticGatewayPrivateZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticGatewayPrivateZoneResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticGatewayPrivateZoneResponse) SetHeaders(v map[string]*string) *CreateElasticGatewayPrivateZoneResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticGatewayPrivateZoneResponse) SetStatusCode(v int32) *CreateElasticGatewayPrivateZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticGatewayPrivateZoneResponse) SetBody(v *CreateElasticGatewayPrivateZoneResponseBody) *CreateElasticGatewayPrivateZoneResponse {
	s.Body = v
	return s
}

type CreateExpressSyncRequest struct {
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketPrefix  *string `json:"BucketPrefix,omitempty" xml:"BucketPrefix,omitempty"`
	BucketRegion  *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateExpressSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressSyncRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressSyncRequest) SetBucketName(v string) *CreateExpressSyncRequest {
	s.BucketName = &v
	return s
}

func (s *CreateExpressSyncRequest) SetBucketPrefix(v string) *CreateExpressSyncRequest {
	s.BucketPrefix = &v
	return s
}

func (s *CreateExpressSyncRequest) SetBucketRegion(v string) *CreateExpressSyncRequest {
	s.BucketRegion = &v
	return s
}

func (s *CreateExpressSyncRequest) SetDescription(v string) *CreateExpressSyncRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressSyncRequest) SetName(v string) *CreateExpressSyncRequest {
	s.Name = &v
	return s
}

func (s *CreateExpressSyncRequest) SetSecurityToken(v string) *CreateExpressSyncRequest {
	s.SecurityToken = &v
	return s
}

type CreateExpressSyncResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ExpressSyncId *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressSyncResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressSyncResponseBody) SetCode(v string) *CreateExpressSyncResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressSyncResponseBody) SetExpressSyncId(v string) *CreateExpressSyncResponseBody {
	s.ExpressSyncId = &v
	return s
}

func (s *CreateExpressSyncResponseBody) SetMessage(v string) *CreateExpressSyncResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressSyncResponseBody) SetRequestId(v string) *CreateExpressSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressSyncResponseBody) SetSuccess(v bool) *CreateExpressSyncResponseBody {
	s.Success = &v
	return s
}

type CreateExpressSyncResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateExpressSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateExpressSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressSyncResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressSyncResponse) SetHeaders(v map[string]*string) *CreateExpressSyncResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressSyncResponse) SetStatusCode(v int32) *CreateExpressSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressSyncResponse) SetBody(v *CreateExpressSyncResponseBody) *CreateExpressSyncResponse {
	s.Body = v
	return s
}

type CreateGatewayRequest struct {
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GatewayClass           *string `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	Location               *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PostPaid               *bool   `json:"PostPaid,omitempty" xml:"PostPaid,omitempty"`
	PublicNetworkBandwidth *int32  `json:"PublicNetworkBandwidth,omitempty" xml:"PublicNetworkBandwidth,omitempty"`
	ReleaseAfterExpiration *bool   `json:"ReleaseAfterExpiration,omitempty" xml:"ReleaseAfterExpiration,omitempty"`
	ResourceRegionId       *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	SecurityToken          *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId        *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) SetDescription(v string) *CreateGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayRequest) SetGatewayClass(v string) *CreateGatewayRequest {
	s.GatewayClass = &v
	return s
}

func (s *CreateGatewayRequest) SetLocation(v string) *CreateGatewayRequest {
	s.Location = &v
	return s
}

func (s *CreateGatewayRequest) SetName(v string) *CreateGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayRequest) SetPostPaid(v bool) *CreateGatewayRequest {
	s.PostPaid = &v
	return s
}

func (s *CreateGatewayRequest) SetPublicNetworkBandwidth(v int32) *CreateGatewayRequest {
	s.PublicNetworkBandwidth = &v
	return s
}

func (s *CreateGatewayRequest) SetReleaseAfterExpiration(v bool) *CreateGatewayRequest {
	s.ReleaseAfterExpiration = &v
	return s
}

func (s *CreateGatewayRequest) SetResourceRegionId(v string) *CreateGatewayRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *CreateGatewayRequest) SetSecurityToken(v string) *CreateGatewayRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGatewayRequest) SetStorageBundleId(v string) *CreateGatewayRequest {
	s.StorageBundleId = &v
	return s
}

func (s *CreateGatewayRequest) SetType(v string) *CreateGatewayRequest {
	s.Type = &v
	return s
}

func (s *CreateGatewayRequest) SetVSwitchId(v string) *CreateGatewayRequest {
	s.VSwitchId = &v
	return s
}

type CreateGatewayResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) SetBuyURL(v string) *CreateGatewayResponseBody {
	s.BuyURL = &v
	return s
}

func (s *CreateGatewayResponseBody) SetCode(v string) *CreateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayResponseBody) SetGatewayId(v string) *CreateGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetMessage(v string) *CreateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetSuccess(v bool) *CreateGatewayResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetStatusCode(v int32) *CreateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

type CreateGatewayBlockVolumeRequest struct {
	CacheMode      *string `json:"CacheMode,omitempty" xml:"CacheMode,omitempty"`
	ChapEnabled    *bool   `json:"ChapEnabled,omitempty" xml:"ChapEnabled,omitempty"`
	ChapInPassword *string `json:"ChapInPassword,omitempty" xml:"ChapInPassword,omitempty"`
	ChapInUser     *string `json:"ChapInUser,omitempty" xml:"ChapInUser,omitempty"`
	ChunkSize      *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	GatewayId      *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	LocalFilePath  *string `json:"LocalFilePath,omitempty" xml:"LocalFilePath,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName  *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssBucketSsl   *bool   `json:"OssBucketSsl,omitempty" xml:"OssBucketSsl,omitempty"`
	OssEndpoint    *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	Recovery       *bool   `json:"Recovery,omitempty" xml:"Recovery,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Size           *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	VolumeProtocol *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
}

func (s CreateGatewayBlockVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayBlockVolumeRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayBlockVolumeRequest) SetCacheMode(v string) *CreateGatewayBlockVolumeRequest {
	s.CacheMode = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetChapEnabled(v bool) *CreateGatewayBlockVolumeRequest {
	s.ChapEnabled = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetChapInPassword(v string) *CreateGatewayBlockVolumeRequest {
	s.ChapInPassword = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetChapInUser(v string) *CreateGatewayBlockVolumeRequest {
	s.ChapInUser = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetChunkSize(v int32) *CreateGatewayBlockVolumeRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetGatewayId(v string) *CreateGatewayBlockVolumeRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetLocalFilePath(v string) *CreateGatewayBlockVolumeRequest {
	s.LocalFilePath = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetName(v string) *CreateGatewayBlockVolumeRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetOssBucketName(v string) *CreateGatewayBlockVolumeRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetOssBucketSsl(v bool) *CreateGatewayBlockVolumeRequest {
	s.OssBucketSsl = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetOssEndpoint(v string) *CreateGatewayBlockVolumeRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetRecovery(v bool) *CreateGatewayBlockVolumeRequest {
	s.Recovery = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetSecurityToken(v string) *CreateGatewayBlockVolumeRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetSize(v int64) *CreateGatewayBlockVolumeRequest {
	s.Size = &v
	return s
}

func (s *CreateGatewayBlockVolumeRequest) SetVolumeProtocol(v string) *CreateGatewayBlockVolumeRequest {
	s.VolumeProtocol = &v
	return s
}

type CreateGatewayBlockVolumeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGatewayBlockVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayBlockVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayBlockVolumeResponseBody) SetCode(v string) *CreateGatewayBlockVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayBlockVolumeResponseBody) SetMessage(v string) *CreateGatewayBlockVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayBlockVolumeResponseBody) SetRequestId(v string) *CreateGatewayBlockVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayBlockVolumeResponseBody) SetSuccess(v bool) *CreateGatewayBlockVolumeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGatewayBlockVolumeResponseBody) SetTaskId(v string) *CreateGatewayBlockVolumeResponseBody {
	s.TaskId = &v
	return s
}

type CreateGatewayBlockVolumeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayBlockVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayBlockVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayBlockVolumeResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayBlockVolumeResponse) SetHeaders(v map[string]*string) *CreateGatewayBlockVolumeResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayBlockVolumeResponse) SetStatusCode(v int32) *CreateGatewayBlockVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayBlockVolumeResponse) SetBody(v *CreateGatewayBlockVolumeResponseBody) *CreateGatewayBlockVolumeResponse {
	s.Body = v
	return s
}

type CreateGatewayCacheDiskRequest struct {
	CacheDiskCategory *string `json:"CacheDiskCategory,omitempty" xml:"CacheDiskCategory,omitempty"`
	CacheDiskSizeInGB *int64  `json:"CacheDiskSizeInGB,omitempty" xml:"CacheDiskSizeInGB,omitempty"`
	GatewayId         *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken     *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateGatewayCacheDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayCacheDiskRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayCacheDiskRequest) SetCacheDiskCategory(v string) *CreateGatewayCacheDiskRequest {
	s.CacheDiskCategory = &v
	return s
}

func (s *CreateGatewayCacheDiskRequest) SetCacheDiskSizeInGB(v int64) *CreateGatewayCacheDiskRequest {
	s.CacheDiskSizeInGB = &v
	return s
}

func (s *CreateGatewayCacheDiskRequest) SetGatewayId(v string) *CreateGatewayCacheDiskRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayCacheDiskRequest) SetSecurityToken(v string) *CreateGatewayCacheDiskRequest {
	s.SecurityToken = &v
	return s
}

type CreateGatewayCacheDiskResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGatewayCacheDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayCacheDiskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayCacheDiskResponseBody) SetBuyURL(v string) *CreateGatewayCacheDiskResponseBody {
	s.BuyURL = &v
	return s
}

func (s *CreateGatewayCacheDiskResponseBody) SetCode(v string) *CreateGatewayCacheDiskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayCacheDiskResponseBody) SetMessage(v string) *CreateGatewayCacheDiskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayCacheDiskResponseBody) SetRequestId(v string) *CreateGatewayCacheDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayCacheDiskResponseBody) SetSuccess(v bool) *CreateGatewayCacheDiskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGatewayCacheDiskResponseBody) SetTaskId(v string) *CreateGatewayCacheDiskResponseBody {
	s.TaskId = &v
	return s
}

type CreateGatewayCacheDiskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayCacheDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayCacheDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayCacheDiskResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayCacheDiskResponse) SetHeaders(v map[string]*string) *CreateGatewayCacheDiskResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayCacheDiskResponse) SetStatusCode(v int32) *CreateGatewayCacheDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayCacheDiskResponse) SetBody(v *CreateGatewayCacheDiskResponseBody) *CreateGatewayCacheDiskResponse {
	s.Body = v
	return s
}

type CreateGatewayFileShareRequest struct {
	AccessBasedEnumeration *bool   `json:"AccessBasedEnumeration,omitempty" xml:"AccessBasedEnumeration,omitempty"`
	BackendLimit           *int32  `json:"BackendLimit,omitempty" xml:"BackendLimit,omitempty"`
	Browsable              *bool   `json:"Browsable,omitempty" xml:"Browsable,omitempty"`
	BypassCacheRead        *bool   `json:"BypassCacheRead,omitempty" xml:"BypassCacheRead,omitempty"`
	CacheMode              *string `json:"CacheMode,omitempty" xml:"CacheMode,omitempty"`
	ClientSideCmk          *string `json:"ClientSideCmk,omitempty" xml:"ClientSideCmk,omitempty"`
	ClientSideEncryption   *bool   `json:"ClientSideEncryption,omitempty" xml:"ClientSideEncryption,omitempty"`
	DirectIO               *bool   `json:"DirectIO,omitempty" xml:"DirectIO,omitempty"`
	DownloadLimit          *int32  `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	FastReclaim            *bool   `json:"FastReclaim,omitempty" xml:"FastReclaim,omitempty"`
	FrontendLimit          *int32  `json:"FrontendLimit,omitempty" xml:"FrontendLimit,omitempty"`
	GatewayId              *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IgnoreDelete           *bool   `json:"IgnoreDelete,omitempty" xml:"IgnoreDelete,omitempty"`
	InPlace                *bool   `json:"InPlace,omitempty" xml:"InPlace,omitempty"`
	KmsRotatePeriod        *int64  `json:"KmsRotatePeriod,omitempty" xml:"KmsRotatePeriod,omitempty"`
	LagPeriod              *int64  `json:"LagPeriod,omitempty" xml:"LagPeriod,omitempty"`
	LocalFilePath          *string `json:"LocalFilePath,omitempty" xml:"LocalFilePath,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NfsV4Optimization      *bool   `json:"NfsV4Optimization,omitempty" xml:"NfsV4Optimization,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssBucketSsl           *bool   `json:"OssBucketSsl,omitempty" xml:"OssBucketSsl,omitempty"`
	OssEndpoint            *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	PartialSyncPaths       *string `json:"PartialSyncPaths,omitempty" xml:"PartialSyncPaths,omitempty"`
	PathPrefix             *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	PollingInterval        *int32  `json:"PollingInterval,omitempty" xml:"PollingInterval,omitempty"`
	ReadOnlyClientList     *string `json:"ReadOnlyClientList,omitempty" xml:"ReadOnlyClientList,omitempty"`
	ReadOnlyUserList       *string `json:"ReadOnlyUserList,omitempty" xml:"ReadOnlyUserList,omitempty"`
	ReadWriteClientList    *string `json:"ReadWriteClientList,omitempty" xml:"ReadWriteClientList,omitempty"`
	ReadWriteUserList      *string `json:"ReadWriteUserList,omitempty" xml:"ReadWriteUserList,omitempty"`
	RemoteSync             *bool   `json:"RemoteSync,omitempty" xml:"RemoteSync,omitempty"`
	RemoteSyncDownload     *bool   `json:"RemoteSyncDownload,omitempty" xml:"RemoteSyncDownload,omitempty"`
	SecurityToken          *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServerSideAlgorithm    *string `json:"ServerSideAlgorithm,omitempty" xml:"ServerSideAlgorithm,omitempty"`
	ServerSideCmk          *string `json:"ServerSideCmk,omitempty" xml:"ServerSideCmk,omitempty"`
	ServerSideEncryption   *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	ShareProtocol          *string `json:"ShareProtocol,omitempty" xml:"ShareProtocol,omitempty"`
	Squash                 *string `json:"Squash,omitempty" xml:"Squash,omitempty"`
	SupportArchive         *bool   `json:"SupportArchive,omitempty" xml:"SupportArchive,omitempty"`
	TransferAcceleration   *bool   `json:"TransferAcceleration,omitempty" xml:"TransferAcceleration,omitempty"`
	WindowsAcl             *bool   `json:"WindowsAcl,omitempty" xml:"WindowsAcl,omitempty"`
}

func (s CreateGatewayFileShareRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayFileShareRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayFileShareRequest) SetAccessBasedEnumeration(v bool) *CreateGatewayFileShareRequest {
	s.AccessBasedEnumeration = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetBackendLimit(v int32) *CreateGatewayFileShareRequest {
	s.BackendLimit = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetBrowsable(v bool) *CreateGatewayFileShareRequest {
	s.Browsable = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetBypassCacheRead(v bool) *CreateGatewayFileShareRequest {
	s.BypassCacheRead = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetCacheMode(v string) *CreateGatewayFileShareRequest {
	s.CacheMode = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetClientSideCmk(v string) *CreateGatewayFileShareRequest {
	s.ClientSideCmk = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetClientSideEncryption(v bool) *CreateGatewayFileShareRequest {
	s.ClientSideEncryption = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetDirectIO(v bool) *CreateGatewayFileShareRequest {
	s.DirectIO = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetDownloadLimit(v int32) *CreateGatewayFileShareRequest {
	s.DownloadLimit = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetFastReclaim(v bool) *CreateGatewayFileShareRequest {
	s.FastReclaim = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetFrontendLimit(v int32) *CreateGatewayFileShareRequest {
	s.FrontendLimit = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetGatewayId(v string) *CreateGatewayFileShareRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetIgnoreDelete(v bool) *CreateGatewayFileShareRequest {
	s.IgnoreDelete = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetInPlace(v bool) *CreateGatewayFileShareRequest {
	s.InPlace = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetKmsRotatePeriod(v int64) *CreateGatewayFileShareRequest {
	s.KmsRotatePeriod = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetLagPeriod(v int64) *CreateGatewayFileShareRequest {
	s.LagPeriod = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetLocalFilePath(v string) *CreateGatewayFileShareRequest {
	s.LocalFilePath = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetName(v string) *CreateGatewayFileShareRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetNfsV4Optimization(v bool) *CreateGatewayFileShareRequest {
	s.NfsV4Optimization = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetOssBucketName(v string) *CreateGatewayFileShareRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetOssBucketSsl(v bool) *CreateGatewayFileShareRequest {
	s.OssBucketSsl = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetOssEndpoint(v string) *CreateGatewayFileShareRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetPartialSyncPaths(v string) *CreateGatewayFileShareRequest {
	s.PartialSyncPaths = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetPathPrefix(v string) *CreateGatewayFileShareRequest {
	s.PathPrefix = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetPollingInterval(v int32) *CreateGatewayFileShareRequest {
	s.PollingInterval = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetReadOnlyClientList(v string) *CreateGatewayFileShareRequest {
	s.ReadOnlyClientList = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetReadOnlyUserList(v string) *CreateGatewayFileShareRequest {
	s.ReadOnlyUserList = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetReadWriteClientList(v string) *CreateGatewayFileShareRequest {
	s.ReadWriteClientList = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetReadWriteUserList(v string) *CreateGatewayFileShareRequest {
	s.ReadWriteUserList = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetRemoteSync(v bool) *CreateGatewayFileShareRequest {
	s.RemoteSync = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetRemoteSyncDownload(v bool) *CreateGatewayFileShareRequest {
	s.RemoteSyncDownload = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetSecurityToken(v string) *CreateGatewayFileShareRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetServerSideAlgorithm(v string) *CreateGatewayFileShareRequest {
	s.ServerSideAlgorithm = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetServerSideCmk(v string) *CreateGatewayFileShareRequest {
	s.ServerSideCmk = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetServerSideEncryption(v bool) *CreateGatewayFileShareRequest {
	s.ServerSideEncryption = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetShareProtocol(v string) *CreateGatewayFileShareRequest {
	s.ShareProtocol = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetSquash(v string) *CreateGatewayFileShareRequest {
	s.Squash = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetSupportArchive(v bool) *CreateGatewayFileShareRequest {
	s.SupportArchive = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetTransferAcceleration(v bool) *CreateGatewayFileShareRequest {
	s.TransferAcceleration = &v
	return s
}

func (s *CreateGatewayFileShareRequest) SetWindowsAcl(v bool) *CreateGatewayFileShareRequest {
	s.WindowsAcl = &v
	return s
}

type CreateGatewayFileShareResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGatewayFileShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayFileShareResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayFileShareResponseBody) SetCode(v string) *CreateGatewayFileShareResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayFileShareResponseBody) SetMessage(v string) *CreateGatewayFileShareResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayFileShareResponseBody) SetRequestId(v string) *CreateGatewayFileShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayFileShareResponseBody) SetSuccess(v bool) *CreateGatewayFileShareResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGatewayFileShareResponseBody) SetTaskId(v string) *CreateGatewayFileShareResponseBody {
	s.TaskId = &v
	return s
}

type CreateGatewayFileShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayFileShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayFileShareResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayFileShareResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayFileShareResponse) SetHeaders(v map[string]*string) *CreateGatewayFileShareResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayFileShareResponse) SetStatusCode(v int32) *CreateGatewayFileShareResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayFileShareResponse) SetBody(v *CreateGatewayFileShareResponseBody) *CreateGatewayFileShareResponse {
	s.Body = v
	return s
}

type CreateGatewayLoggingRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SlsLogstore   *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	SlsProject    *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s CreateGatewayLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayLoggingRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayLoggingRequest) SetGatewayId(v string) *CreateGatewayLoggingRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayLoggingRequest) SetSecurityToken(v string) *CreateGatewayLoggingRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGatewayLoggingRequest) SetSlsLogstore(v string) *CreateGatewayLoggingRequest {
	s.SlsLogstore = &v
	return s
}

func (s *CreateGatewayLoggingRequest) SetSlsProject(v string) *CreateGatewayLoggingRequest {
	s.SlsProject = &v
	return s
}

type CreateGatewayLoggingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayLoggingResponseBody) SetCode(v string) *CreateGatewayLoggingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayLoggingResponseBody) SetMessage(v string) *CreateGatewayLoggingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayLoggingResponseBody) SetRequestId(v string) *CreateGatewayLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayLoggingResponseBody) SetSuccess(v bool) *CreateGatewayLoggingResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayLoggingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayLoggingResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayLoggingResponse) SetHeaders(v map[string]*string) *CreateGatewayLoggingResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayLoggingResponse) SetStatusCode(v int32) *CreateGatewayLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayLoggingResponse) SetBody(v *CreateGatewayLoggingResponseBody) *CreateGatewayLoggingResponse {
	s.Body = v
	return s
}

type CreateGatewaySMBUserRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateGatewaySMBUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewaySMBUserRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewaySMBUserRequest) SetGatewayId(v string) *CreateGatewaySMBUserRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewaySMBUserRequest) SetPassword(v string) *CreateGatewaySMBUserRequest {
	s.Password = &v
	return s
}

func (s *CreateGatewaySMBUserRequest) SetSecurityToken(v string) *CreateGatewaySMBUserRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGatewaySMBUserRequest) SetUsername(v string) *CreateGatewaySMBUserRequest {
	s.Username = &v
	return s
}

type CreateGatewaySMBUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGatewaySMBUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewaySMBUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewaySMBUserResponseBody) SetCode(v string) *CreateGatewaySMBUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewaySMBUserResponseBody) SetMessage(v string) *CreateGatewaySMBUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewaySMBUserResponseBody) SetRequestId(v string) *CreateGatewaySMBUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewaySMBUserResponseBody) SetSuccess(v bool) *CreateGatewaySMBUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGatewaySMBUserResponseBody) SetTaskId(v string) *CreateGatewaySMBUserResponseBody {
	s.TaskId = &v
	return s
}

type CreateGatewaySMBUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewaySMBUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewaySMBUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewaySMBUserResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewaySMBUserResponse) SetHeaders(v map[string]*string) *CreateGatewaySMBUserResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewaySMBUserResponse) SetStatusCode(v int32) *CreateGatewaySMBUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewaySMBUserResponse) SetBody(v *CreateGatewaySMBUserResponseBody) *CreateGatewaySMBUserResponse {
	s.Body = v
	return s
}

type CreateStorageBundleRequest struct {
	BackendBucketRegionId *string `json:"BackendBucketRegionId,omitempty" xml:"BackendBucketRegionId,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateStorageBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStorageBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageBundleRequest) SetBackendBucketRegionId(v string) *CreateStorageBundleRequest {
	s.BackendBucketRegionId = &v
	return s
}

func (s *CreateStorageBundleRequest) SetDescription(v string) *CreateStorageBundleRequest {
	s.Description = &v
	return s
}

func (s *CreateStorageBundleRequest) SetName(v string) *CreateStorageBundleRequest {
	s.Name = &v
	return s
}

func (s *CreateStorageBundleRequest) SetSecurityToken(v string) *CreateStorageBundleRequest {
	s.SecurityToken = &v
	return s
}

type CreateStorageBundleResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStorageBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStorageBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageBundleResponseBody) SetCode(v string) *CreateStorageBundleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStorageBundleResponseBody) SetMessage(v string) *CreateStorageBundleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStorageBundleResponseBody) SetRequestId(v string) *CreateStorageBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageBundleResponseBody) SetStorageBundleId(v string) *CreateStorageBundleResponseBody {
	s.StorageBundleId = &v
	return s
}

func (s *CreateStorageBundleResponseBody) SetSuccess(v bool) *CreateStorageBundleResponseBody {
	s.Success = &v
	return s
}

type CreateStorageBundleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStorageBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStorageBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStorageBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageBundleResponse) SetHeaders(v map[string]*string) *CreateStorageBundleResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageBundleResponse) SetStatusCode(v int32) *CreateStorageBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStorageBundleResponse) SetBody(v *CreateStorageBundleResponseBody) *CreateStorageBundleResponse {
	s.Body = v
	return s
}

type DeleteCSGClientsRequest struct {
	ClientIds      []*string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty" type:"Repeated"`
	ClientRegionId *string   `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	SecurityToken  *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteCSGClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCSGClientsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCSGClientsRequest) SetClientIds(v []*string) *DeleteCSGClientsRequest {
	s.ClientIds = v
	return s
}

func (s *DeleteCSGClientsRequest) SetClientRegionId(v string) *DeleteCSGClientsRequest {
	s.ClientRegionId = &v
	return s
}

func (s *DeleteCSGClientsRequest) SetSecurityToken(v string) *DeleteCSGClientsRequest {
	s.SecurityToken = &v
	return s
}

type DeleteCSGClientsShrinkRequest struct {
	ClientIdsShrink *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	ClientRegionId  *string `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteCSGClientsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCSGClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCSGClientsShrinkRequest) SetClientIdsShrink(v string) *DeleteCSGClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *DeleteCSGClientsShrinkRequest) SetClientRegionId(v string) *DeleteCSGClientsShrinkRequest {
	s.ClientRegionId = &v
	return s
}

func (s *DeleteCSGClientsShrinkRequest) SetSecurityToken(v string) *DeleteCSGClientsShrinkRequest {
	s.SecurityToken = &v
	return s
}

type DeleteCSGClientsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteCSGClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCSGClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCSGClientsResponseBody) SetCode(v string) *DeleteCSGClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCSGClientsResponseBody) SetMessage(v string) *DeleteCSGClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCSGClientsResponseBody) SetRequestId(v string) *DeleteCSGClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCSGClientsResponseBody) SetSuccess(v bool) *DeleteCSGClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCSGClientsResponseBody) SetTaskId(v string) *DeleteCSGClientsResponseBody {
	s.TaskId = &v
	return s
}

type DeleteCSGClientsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCSGClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCSGClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCSGClientsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCSGClientsResponse) SetHeaders(v map[string]*string) *DeleteCSGClientsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCSGClientsResponse) SetStatusCode(v int32) *DeleteCSGClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCSGClientsResponse) SetBody(v *DeleteCSGClientsResponseBody) *DeleteCSGClientsResponse {
	s.Body = v
	return s
}

type DeleteElasticGatewayPrivateZoneRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteElasticGatewayPrivateZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticGatewayPrivateZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticGatewayPrivateZoneRequest) SetGatewayId(v string) *DeleteElasticGatewayPrivateZoneRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneRequest) SetSecurityToken(v string) *DeleteElasticGatewayPrivateZoneRequest {
	s.SecurityToken = &v
	return s
}

type DeleteElasticGatewayPrivateZoneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteElasticGatewayPrivateZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticGatewayPrivateZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticGatewayPrivateZoneResponseBody) SetCode(v string) *DeleteElasticGatewayPrivateZoneResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneResponseBody) SetMessage(v string) *DeleteElasticGatewayPrivateZoneResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneResponseBody) SetRequestId(v string) *DeleteElasticGatewayPrivateZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneResponseBody) SetSuccess(v bool) *DeleteElasticGatewayPrivateZoneResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneResponseBody) SetTaskId(v string) *DeleteElasticGatewayPrivateZoneResponseBody {
	s.TaskId = &v
	return s
}

type DeleteElasticGatewayPrivateZoneResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteElasticGatewayPrivateZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteElasticGatewayPrivateZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticGatewayPrivateZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticGatewayPrivateZoneResponse) SetHeaders(v map[string]*string) *DeleteElasticGatewayPrivateZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneResponse) SetStatusCode(v int32) *DeleteElasticGatewayPrivateZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticGatewayPrivateZoneResponse) SetBody(v *DeleteElasticGatewayPrivateZoneResponseBody) *DeleteElasticGatewayPrivateZoneResponse {
	s.Body = v
	return s
}

type DeleteExpressSyncRequest struct {
	ExpressSyncId *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteExpressSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressSyncRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressSyncRequest) SetExpressSyncId(v string) *DeleteExpressSyncRequest {
	s.ExpressSyncId = &v
	return s
}

func (s *DeleteExpressSyncRequest) SetSecurityToken(v string) *DeleteExpressSyncRequest {
	s.SecurityToken = &v
	return s
}

type DeleteExpressSyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteExpressSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressSyncResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressSyncResponseBody) SetCode(v string) *DeleteExpressSyncResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressSyncResponseBody) SetMessage(v string) *DeleteExpressSyncResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressSyncResponseBody) SetRequestId(v string) *DeleteExpressSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressSyncResponseBody) SetSuccess(v bool) *DeleteExpressSyncResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExpressSyncResponseBody) SetTaskId(v string) *DeleteExpressSyncResponseBody {
	s.TaskId = &v
	return s
}

type DeleteExpressSyncResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteExpressSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExpressSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressSyncResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressSyncResponse) SetHeaders(v map[string]*string) *DeleteExpressSyncResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressSyncResponse) SetStatusCode(v int32) *DeleteExpressSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressSyncResponse) SetBody(v *DeleteExpressSyncResponseBody) *DeleteExpressSyncResponse {
	s.Body = v
	return s
}

type DeleteGatewayRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	ReasonDetail  *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
	ReasonType    *string `json:"ReasonType,omitempty" xml:"ReasonType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) SetGatewayId(v string) *DeleteGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayRequest) SetReasonDetail(v string) *DeleteGatewayRequest {
	s.ReasonDetail = &v
	return s
}

func (s *DeleteGatewayRequest) SetReasonType(v string) *DeleteGatewayRequest {
	s.ReasonType = &v
	return s
}

func (s *DeleteGatewayRequest) SetSecurityToken(v string) *DeleteGatewayRequest {
	s.SecurityToken = &v
	return s
}

type DeleteGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetCode(v string) *DeleteGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetMessage(v string) *DeleteGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetSuccess(v bool) *DeleteGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetTaskId(v string) *DeleteGatewayResponseBody {
	s.TaskId = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type DeleteGatewayBlockVolumesRequest struct {
	GatewayId        *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId          *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	IsSourceDeletion *bool   `json:"IsSourceDeletion,omitempty" xml:"IsSourceDeletion,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGatewayBlockVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayBlockVolumesRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayBlockVolumesRequest) SetGatewayId(v string) *DeleteGatewayBlockVolumesRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayBlockVolumesRequest) SetIndexId(v string) *DeleteGatewayBlockVolumesRequest {
	s.IndexId = &v
	return s
}

func (s *DeleteGatewayBlockVolumesRequest) SetIsSourceDeletion(v bool) *DeleteGatewayBlockVolumesRequest {
	s.IsSourceDeletion = &v
	return s
}

func (s *DeleteGatewayBlockVolumesRequest) SetSecurityToken(v string) *DeleteGatewayBlockVolumesRequest {
	s.SecurityToken = &v
	return s
}

type DeleteGatewayBlockVolumesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteGatewayBlockVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayBlockVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayBlockVolumesResponseBody) SetCode(v string) *DeleteGatewayBlockVolumesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayBlockVolumesResponseBody) SetMessage(v string) *DeleteGatewayBlockVolumesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayBlockVolumesResponseBody) SetRequestId(v string) *DeleteGatewayBlockVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayBlockVolumesResponseBody) SetSuccess(v bool) *DeleteGatewayBlockVolumesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayBlockVolumesResponseBody) SetTaskId(v string) *DeleteGatewayBlockVolumesResponseBody {
	s.TaskId = &v
	return s
}

type DeleteGatewayBlockVolumesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayBlockVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayBlockVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayBlockVolumesResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayBlockVolumesResponse) SetHeaders(v map[string]*string) *DeleteGatewayBlockVolumesResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayBlockVolumesResponse) SetStatusCode(v int32) *DeleteGatewayBlockVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayBlockVolumesResponse) SetBody(v *DeleteGatewayBlockVolumesResponseBody) *DeleteGatewayBlockVolumesResponse {
	s.Body = v
	return s
}

type DeleteGatewayCacheDiskRequest struct {
	CacheId       *string `json:"CacheId,omitempty" xml:"CacheId,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	LocalFilePath *string `json:"LocalFilePath,omitempty" xml:"LocalFilePath,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGatewayCacheDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayCacheDiskRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayCacheDiskRequest) SetCacheId(v string) *DeleteGatewayCacheDiskRequest {
	s.CacheId = &v
	return s
}

func (s *DeleteGatewayCacheDiskRequest) SetGatewayId(v string) *DeleteGatewayCacheDiskRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayCacheDiskRequest) SetLocalFilePath(v string) *DeleteGatewayCacheDiskRequest {
	s.LocalFilePath = &v
	return s
}

func (s *DeleteGatewayCacheDiskRequest) SetSecurityToken(v string) *DeleteGatewayCacheDiskRequest {
	s.SecurityToken = &v
	return s
}

type DeleteGatewayCacheDiskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteGatewayCacheDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayCacheDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayCacheDiskResponseBody) SetCode(v string) *DeleteGatewayCacheDiskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayCacheDiskResponseBody) SetMessage(v string) *DeleteGatewayCacheDiskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayCacheDiskResponseBody) SetRequestId(v string) *DeleteGatewayCacheDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayCacheDiskResponseBody) SetSuccess(v bool) *DeleteGatewayCacheDiskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayCacheDiskResponseBody) SetTaskId(v string) *DeleteGatewayCacheDiskResponseBody {
	s.TaskId = &v
	return s
}

type DeleteGatewayCacheDiskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayCacheDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayCacheDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayCacheDiskResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayCacheDiskResponse) SetHeaders(v map[string]*string) *DeleteGatewayCacheDiskResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayCacheDiskResponse) SetStatusCode(v int32) *DeleteGatewayCacheDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayCacheDiskResponse) SetBody(v *DeleteGatewayCacheDiskResponseBody) *DeleteGatewayCacheDiskResponse {
	s.Body = v
	return s
}

type DeleteGatewayFileSharesRequest struct {
	Force         *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGatewayFileSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayFileSharesRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayFileSharesRequest) SetForce(v bool) *DeleteGatewayFileSharesRequest {
	s.Force = &v
	return s
}

func (s *DeleteGatewayFileSharesRequest) SetGatewayId(v string) *DeleteGatewayFileSharesRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayFileSharesRequest) SetIndexId(v string) *DeleteGatewayFileSharesRequest {
	s.IndexId = &v
	return s
}

func (s *DeleteGatewayFileSharesRequest) SetSecurityToken(v string) *DeleteGatewayFileSharesRequest {
	s.SecurityToken = &v
	return s
}

type DeleteGatewayFileSharesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteGatewayFileSharesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayFileSharesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayFileSharesResponseBody) SetCode(v string) *DeleteGatewayFileSharesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayFileSharesResponseBody) SetMessage(v string) *DeleteGatewayFileSharesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayFileSharesResponseBody) SetRequestId(v string) *DeleteGatewayFileSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayFileSharesResponseBody) SetSuccess(v bool) *DeleteGatewayFileSharesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayFileSharesResponseBody) SetTaskId(v string) *DeleteGatewayFileSharesResponseBody {
	s.TaskId = &v
	return s
}

type DeleteGatewayFileSharesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayFileSharesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayFileSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayFileSharesResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayFileSharesResponse) SetHeaders(v map[string]*string) *DeleteGatewayFileSharesResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayFileSharesResponse) SetStatusCode(v int32) *DeleteGatewayFileSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayFileSharesResponse) SetBody(v *DeleteGatewayFileSharesResponseBody) *DeleteGatewayFileSharesResponse {
	s.Body = v
	return s
}

type DeleteGatewayLoggingRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGatewayLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayLoggingRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLoggingRequest) SetGatewayId(v string) *DeleteGatewayLoggingRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayLoggingRequest) SetSecurityToken(v string) *DeleteGatewayLoggingRequest {
	s.SecurityToken = &v
	return s
}

type DeleteGatewayLoggingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLoggingResponseBody) SetCode(v string) *DeleteGatewayLoggingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayLoggingResponseBody) SetMessage(v string) *DeleteGatewayLoggingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayLoggingResponseBody) SetRequestId(v string) *DeleteGatewayLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayLoggingResponseBody) SetSuccess(v bool) *DeleteGatewayLoggingResponseBody {
	s.Success = &v
	return s
}

type DeleteGatewayLoggingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayLoggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLoggingResponse) SetHeaders(v map[string]*string) *DeleteGatewayLoggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayLoggingResponse) SetStatusCode(v int32) *DeleteGatewayLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayLoggingResponse) SetBody(v *DeleteGatewayLoggingResponseBody) *DeleteGatewayLoggingResponse {
	s.Body = v
	return s
}

type DeleteGatewaySMBUserRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteGatewaySMBUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySMBUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySMBUserRequest) SetGatewayId(v string) *DeleteGatewaySMBUserRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewaySMBUserRequest) SetSecurityToken(v string) *DeleteGatewaySMBUserRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteGatewaySMBUserRequest) SetUsername(v string) *DeleteGatewaySMBUserRequest {
	s.Username = &v
	return s
}

type DeleteGatewaySMBUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteGatewaySMBUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySMBUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySMBUserResponseBody) SetCode(v string) *DeleteGatewaySMBUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewaySMBUserResponseBody) SetMessage(v string) *DeleteGatewaySMBUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewaySMBUserResponseBody) SetRequestId(v string) *DeleteGatewaySMBUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewaySMBUserResponseBody) SetSuccess(v bool) *DeleteGatewaySMBUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewaySMBUserResponseBody) SetTaskId(v string) *DeleteGatewaySMBUserResponseBody {
	s.TaskId = &v
	return s
}

type DeleteGatewaySMBUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewaySMBUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewaySMBUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySMBUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySMBUserResponse) SetHeaders(v map[string]*string) *DeleteGatewaySMBUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewaySMBUserResponse) SetStatusCode(v int32) *DeleteGatewaySMBUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewaySMBUserResponse) SetBody(v *DeleteGatewaySMBUserResponseBody) *DeleteGatewaySMBUserResponse {
	s.Body = v
	return s
}

type DeleteStorageBundleRequest struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
}

func (s DeleteStorageBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStorageBundleRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageBundleRequest) SetSecurityToken(v string) *DeleteStorageBundleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteStorageBundleRequest) SetStorageBundleId(v string) *DeleteStorageBundleRequest {
	s.StorageBundleId = &v
	return s
}

type DeleteStorageBundleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStorageBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStorageBundleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStorageBundleResponseBody) SetCode(v string) *DeleteStorageBundleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStorageBundleResponseBody) SetMessage(v string) *DeleteStorageBundleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStorageBundleResponseBody) SetRequestId(v string) *DeleteStorageBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStorageBundleResponseBody) SetSuccess(v bool) *DeleteStorageBundleResponseBody {
	s.Success = &v
	return s
}

type DeleteStorageBundleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStorageBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStorageBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStorageBundleResponse) GoString() string {
	return s.String()
}

func (s *DeleteStorageBundleResponse) SetHeaders(v map[string]*string) *DeleteStorageBundleResponse {
	s.Headers = v
	return s
}

func (s *DeleteStorageBundleResponse) SetStatusCode(v int32) *DeleteStorageBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStorageBundleResponse) SetBody(v *DeleteStorageBundleResponseBody) *DeleteStorageBundleResponse {
	s.Body = v
	return s
}

type DeployCSGClientsRequest struct {
	ClientRegionId *string   `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	EcsInstanceIds []*string `json:"EcsInstanceIds,omitempty" xml:"EcsInstanceIds,omitempty" type:"Repeated"`
	SecurityToken  *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId          *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeployCSGClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployCSGClientsRequest) GoString() string {
	return s.String()
}

func (s *DeployCSGClientsRequest) SetClientRegionId(v string) *DeployCSGClientsRequest {
	s.ClientRegionId = &v
	return s
}

func (s *DeployCSGClientsRequest) SetEcsInstanceIds(v []*string) *DeployCSGClientsRequest {
	s.EcsInstanceIds = v
	return s
}

func (s *DeployCSGClientsRequest) SetSecurityToken(v string) *DeployCSGClientsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployCSGClientsRequest) SetVpcId(v string) *DeployCSGClientsRequest {
	s.VpcId = &v
	return s
}

type DeployCSGClientsShrinkRequest struct {
	ClientRegionId       *string `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	EcsInstanceIdsShrink *string `json:"EcsInstanceIds,omitempty" xml:"EcsInstanceIds,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeployCSGClientsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployCSGClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeployCSGClientsShrinkRequest) SetClientRegionId(v string) *DeployCSGClientsShrinkRequest {
	s.ClientRegionId = &v
	return s
}

func (s *DeployCSGClientsShrinkRequest) SetEcsInstanceIdsShrink(v string) *DeployCSGClientsShrinkRequest {
	s.EcsInstanceIdsShrink = &v
	return s
}

func (s *DeployCSGClientsShrinkRequest) SetSecurityToken(v string) *DeployCSGClientsShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployCSGClientsShrinkRequest) SetVpcId(v string) *DeployCSGClientsShrinkRequest {
	s.VpcId = &v
	return s
}

type DeployCSGClientsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeployCSGClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployCSGClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DeployCSGClientsResponseBody) SetCode(v string) *DeployCSGClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DeployCSGClientsResponseBody) SetMessage(v string) *DeployCSGClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DeployCSGClientsResponseBody) SetRequestId(v string) *DeployCSGClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployCSGClientsResponseBody) SetSuccess(v bool) *DeployCSGClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DeployCSGClientsResponseBody) SetTaskId(v string) *DeployCSGClientsResponseBody {
	s.TaskId = &v
	return s
}

type DeployCSGClientsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployCSGClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployCSGClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployCSGClientsResponse) GoString() string {
	return s.String()
}

func (s *DeployCSGClientsResponse) SetHeaders(v map[string]*string) *DeployCSGClientsResponse {
	s.Headers = v
	return s
}

func (s *DeployCSGClientsResponse) SetStatusCode(v int32) *DeployCSGClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployCSGClientsResponse) SetBody(v *DeployCSGClientsResponseBody) *DeployCSGClientsResponse {
	s.Body = v
	return s
}

type DeployCacheDiskRequest struct {
	CacheConfig   *string `json:"CacheConfig,omitempty" xml:"CacheConfig,omitempty"`
	DiskCategory  *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SizeInGB      *int32  `json:"SizeInGB,omitempty" xml:"SizeInGB,omitempty"`
}

func (s DeployCacheDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployCacheDiskRequest) GoString() string {
	return s.String()
}

func (s *DeployCacheDiskRequest) SetCacheConfig(v string) *DeployCacheDiskRequest {
	s.CacheConfig = &v
	return s
}

func (s *DeployCacheDiskRequest) SetDiskCategory(v string) *DeployCacheDiskRequest {
	s.DiskCategory = &v
	return s
}

func (s *DeployCacheDiskRequest) SetGatewayId(v string) *DeployCacheDiskRequest {
	s.GatewayId = &v
	return s
}

func (s *DeployCacheDiskRequest) SetSecurityToken(v string) *DeployCacheDiskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployCacheDiskRequest) SetSizeInGB(v int32) *DeployCacheDiskRequest {
	s.SizeInGB = &v
	return s
}

type DeployCacheDiskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeployCacheDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployCacheDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DeployCacheDiskResponseBody) SetCode(v string) *DeployCacheDiskResponseBody {
	s.Code = &v
	return s
}

func (s *DeployCacheDiskResponseBody) SetMessage(v string) *DeployCacheDiskResponseBody {
	s.Message = &v
	return s
}

func (s *DeployCacheDiskResponseBody) SetRequestId(v string) *DeployCacheDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployCacheDiskResponseBody) SetSuccess(v bool) *DeployCacheDiskResponseBody {
	s.Success = &v
	return s
}

func (s *DeployCacheDiskResponseBody) SetTaskId(v string) *DeployCacheDiskResponseBody {
	s.TaskId = &v
	return s
}

type DeployCacheDiskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployCacheDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployCacheDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployCacheDiskResponse) GoString() string {
	return s.String()
}

func (s *DeployCacheDiskResponse) SetHeaders(v map[string]*string) *DeployCacheDiskResponse {
	s.Headers = v
	return s
}

func (s *DeployCacheDiskResponse) SetStatusCode(v int32) *DeployCacheDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployCacheDiskResponse) SetBody(v *DeployCacheDiskResponseBody) *DeployCacheDiskResponse {
	s.Body = v
	return s
}

type DeployGatewayRequest struct {
	GatewayClass  *string `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeployGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeployGatewayRequest) SetGatewayClass(v string) *DeployGatewayRequest {
	s.GatewayClass = &v
	return s
}

func (s *DeployGatewayRequest) SetGatewayId(v string) *DeployGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DeployGatewayRequest) SetSecurityToken(v string) *DeployGatewayRequest {
	s.SecurityToken = &v
	return s
}

type DeployGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeployGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeployGatewayResponseBody) SetCode(v string) *DeployGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeployGatewayResponseBody) SetMessage(v string) *DeployGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeployGatewayResponseBody) SetRequestId(v string) *DeployGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployGatewayResponseBody) SetSuccess(v bool) *DeployGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *DeployGatewayResponseBody) SetTaskId(v string) *DeployGatewayResponseBody {
	s.TaskId = &v
	return s
}

type DeployGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeployGatewayResponse) SetHeaders(v map[string]*string) *DeployGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeployGatewayResponse) SetStatusCode(v int32) *DeployGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployGatewayResponse) SetBody(v *DeployGatewayResponseBody) *DeployGatewayResponse {
	s.Body = v
	return s
}

type DescribeAccountConfigRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAccountConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountConfigRequest) SetGatewayId(v string) *DescribeAccountConfigRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeAccountConfigRequest) SetSecurityToken(v string) *DescribeAccountConfigRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAccountConfigResponseBody struct {
	Code                          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSupportClientSideEncryption *bool   `json:"IsSupportClientSideEncryption,omitempty" xml:"IsSupportClientSideEncryption,omitempty"`
	IsSupportElasticGatewayBeta   *bool   `json:"IsSupportElasticGatewayBeta,omitempty" xml:"IsSupportElasticGatewayBeta,omitempty"`
	IsSupportGatewayLogging       *bool   `json:"IsSupportGatewayLogging,omitempty" xml:"IsSupportGatewayLogging,omitempty"`
	IsSupportServerSideEncryption *bool   `json:"IsSupportServerSideEncryption,omitempty" xml:"IsSupportServerSideEncryption,omitempty"`
	Message                       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAccountConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountConfigResponseBody) SetCode(v string) *DescribeAccountConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetIsSupportClientSideEncryption(v bool) *DescribeAccountConfigResponseBody {
	s.IsSupportClientSideEncryption = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetIsSupportElasticGatewayBeta(v bool) *DescribeAccountConfigResponseBody {
	s.IsSupportElasticGatewayBeta = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetIsSupportGatewayLogging(v bool) *DescribeAccountConfigResponseBody {
	s.IsSupportGatewayLogging = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetIsSupportServerSideEncryption(v bool) *DescribeAccountConfigResponseBody {
	s.IsSupportServerSideEncryption = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetMessage(v string) *DescribeAccountConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetRequestId(v string) *DescribeAccountConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountConfigResponseBody) SetSuccess(v bool) *DescribeAccountConfigResponseBody {
	s.Success = &v
	return s
}

type DescribeAccountConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccountConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountConfigResponse) SetHeaders(v map[string]*string) *DescribeAccountConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountConfigResponse) SetStatusCode(v int32) *DescribeAccountConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountConfigResponse) SetBody(v *DescribeAccountConfigResponseBody) *DescribeAccountConfigResponse {
	s.Body = v
	return s
}

type DescribeBlockVolumeSnapshotsRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeBlockVolumeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockVolumeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlockVolumeSnapshotsRequest) SetGatewayId(v string) *DescribeBlockVolumeSnapshotsRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsRequest) SetIndexId(v string) *DescribeBlockVolumeSnapshotsRequest {
	s.IndexId = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsRequest) SetPageNumber(v int32) *DescribeBlockVolumeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsRequest) SetPageSize(v int32) *DescribeBlockVolumeSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsRequest) SetSecurityToken(v string) *DescribeBlockVolumeSnapshotsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeBlockVolumeSnapshotsResponseBody struct {
	Code       *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  *DescribeBlockVolumeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Success    *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBlockVolumeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockVolumeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetCode(v string) *DescribeBlockVolumeSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetMessage(v string) *DescribeBlockVolumeSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetPageNumber(v int32) *DescribeBlockVolumeSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetPageSize(v int32) *DescribeBlockVolumeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetRequestId(v string) *DescribeBlockVolumeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetSnapshots(v *DescribeBlockVolumeSnapshotsResponseBodySnapshots) *DescribeBlockVolumeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetSuccess(v bool) *DescribeBlockVolumeSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeBlockVolumeSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBlockVolumeSnapshotsResponseBodySnapshots struct {
	Snapshot []*DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s DescribeBlockVolumeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockVolumeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeBlockVolumeSnapshotsResponseBodySnapshots) SetSnapshot(v []*DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot) *DescribeBlockVolumeSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

type DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot struct {
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v int64) *DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot) SetSize(v int64) *DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot {
	s.Size = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotName(v string) *DescribeBlockVolumeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotName = &v
	return s
}

type DescribeBlockVolumeSnapshotsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBlockVolumeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBlockVolumeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockVolumeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlockVolumeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeBlockVolumeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponse) SetStatusCode(v int32) *DescribeBlockVolumeSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlockVolumeSnapshotsResponse) SetBody(v *DescribeBlockVolumeSnapshotsResponseBody) *DescribeBlockVolumeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeCSGClientTasksRequest struct {
	ClientId       *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientRegionId *string `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCSGClientTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientTasksRequest) SetClientId(v string) *DescribeCSGClientTasksRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeCSGClientTasksRequest) SetClientRegionId(v string) *DescribeCSGClientTasksRequest {
	s.ClientRegionId = &v
	return s
}

func (s *DescribeCSGClientTasksRequest) SetPageNumber(v int32) *DescribeCSGClientTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCSGClientTasksRequest) SetPageSize(v int32) *DescribeCSGClientTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCSGClientTasksRequest) SetSecurityToken(v string) *DescribeCSGClientTasksRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCSGClientTasksResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Tasks      []*DescribeCSGClientTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCSGClientTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientTasksResponseBody) SetCode(v string) *DescribeCSGClientTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetMessage(v string) *DescribeCSGClientTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetPageNumber(v int32) *DescribeCSGClientTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetPageSize(v int32) *DescribeCSGClientTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetRequestId(v string) *DescribeCSGClientTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetSuccess(v bool) *DescribeCSGClientTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetTasks(v []*DescribeCSGClientTasksResponseBodyTasks) *DescribeCSGClientTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeCSGClientTasksResponseBody) SetTotalCount(v int32) *DescribeCSGClientTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCSGClientTasksResponseBodyTasks struct {
	CreatedTime   *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	MessageKey    *string `json:"MessageKey,omitempty" xml:"MessageKey,omitempty"`
	MessageParams *string `json:"MessageParams,omitempty" xml:"MessageParams,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress      *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StateCode     *string `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpdatedTime   *int64  `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeCSGClientTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetCreatedTime(v int64) *DescribeCSGClientTasksResponseBodyTasks {
	s.CreatedTime = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetMessageKey(v string) *DescribeCSGClientTasksResponseBodyTasks {
	s.MessageKey = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetMessageParams(v string) *DescribeCSGClientTasksResponseBodyTasks {
	s.MessageParams = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetName(v string) *DescribeCSGClientTasksResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetProgress(v int32) *DescribeCSGClientTasksResponseBodyTasks {
	s.Progress = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetStateCode(v string) *DescribeCSGClientTasksResponseBodyTasks {
	s.StateCode = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetTaskId(v string) *DescribeCSGClientTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeCSGClientTasksResponseBodyTasks) SetUpdatedTime(v int64) *DescribeCSGClientTasksResponseBodyTasks {
	s.UpdatedTime = &v
	return s
}

type DescribeCSGClientTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCSGClientTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCSGClientTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientTasksResponse) SetHeaders(v map[string]*string) *DescribeCSGClientTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeCSGClientTasksResponse) SetStatusCode(v int32) *DescribeCSGClientTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCSGClientTasksResponse) SetBody(v *DescribeCSGClientTasksResponseBody) *DescribeCSGClientTasksResponse {
	s.Body = v
	return s
}

type DescribeCSGClientsRequest struct {
	ClientRegionId *string `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCSGClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientsRequest) SetClientRegionId(v string) *DescribeCSGClientsRequest {
	s.ClientRegionId = &v
	return s
}

func (s *DescribeCSGClientsRequest) SetPageNumber(v int32) *DescribeCSGClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCSGClientsRequest) SetPageSize(v int32) *DescribeCSGClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCSGClientsRequest) SetSecurityToken(v string) *DescribeCSGClientsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCSGClientsResponseBody struct {
	Clients    []*DescribeCSGClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCSGClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientsResponseBody) SetClients(v []*DescribeCSGClientsResponseBodyClients) *DescribeCSGClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetCode(v string) *DescribeCSGClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetMessage(v string) *DescribeCSGClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetPageNumber(v int32) *DescribeCSGClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetPageSize(v int32) *DescribeCSGClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetRequestId(v string) *DescribeCSGClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetSuccess(v bool) *DescribeCSGClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCSGClientsResponseBody) SetTotalCount(v int32) *DescribeCSGClientsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCSGClientsResponseBodyClients struct {
	ClientDeletionCommand     *string `json:"ClientDeletionCommand,omitempty" xml:"ClientDeletionCommand,omitempty"`
	ClientId                  *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientInstallationCommand *string `json:"ClientInstallationCommand,omitempty" xml:"ClientInstallationCommand,omitempty"`
	ClientVersion             *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EcsInstanceId             *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	HostInstanceId            *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	LastHeartbeatTime         *int64  `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId                     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WorkDirectory             *string `json:"WorkDirectory,omitempty" xml:"WorkDirectory,omitempty"`
}

func (s DescribeCSGClientsResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientsResponseBodyClients) SetClientDeletionCommand(v string) *DescribeCSGClientsResponseBodyClients {
	s.ClientDeletionCommand = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetClientId(v string) *DescribeCSGClientsResponseBodyClients {
	s.ClientId = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetClientInstallationCommand(v string) *DescribeCSGClientsResponseBodyClients {
	s.ClientInstallationCommand = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetClientVersion(v string) *DescribeCSGClientsResponseBodyClients {
	s.ClientVersion = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetEcsInstanceId(v string) *DescribeCSGClientsResponseBodyClients {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetHostInstanceId(v string) *DescribeCSGClientsResponseBodyClients {
	s.HostInstanceId = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetLastHeartbeatTime(v int64) *DescribeCSGClientsResponseBodyClients {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetStatus(v string) *DescribeCSGClientsResponseBodyClients {
	s.Status = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetVpcId(v string) *DescribeCSGClientsResponseBodyClients {
	s.VpcId = &v
	return s
}

func (s *DescribeCSGClientsResponseBodyClients) SetWorkDirectory(v string) *DescribeCSGClientsResponseBodyClients {
	s.WorkDirectory = &v
	return s
}

type DescribeCSGClientsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCSGClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCSGClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCSGClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCSGClientsResponse) SetHeaders(v map[string]*string) *DescribeCSGClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCSGClientsResponse) SetStatusCode(v int32) *DescribeCSGClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCSGClientsResponse) SetBody(v *DescribeCSGClientsResponseBody) *DescribeCSGClientsResponse {
	s.Body = v
	return s
}

type DescribeDashboardRequest struct {
	BackendBucketRegionId *string `json:"BackendBucketRegionId,omitempty" xml:"BackendBucketRegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDashboardRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardRequest) SetBackendBucketRegionId(v string) *DescribeDashboardRequest {
	s.BackendBucketRegionId = &v
	return s
}

func (s *DescribeDashboardRequest) SetResourceGroupId(v string) *DescribeDashboardRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDashboardRequest) SetSecurityToken(v string) *DescribeDashboardRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDashboardResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Overview  *string `json:"Overview,omitempty" xml:"Overview,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardResponseBody) SetCode(v string) *DescribeDashboardResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDashboardResponseBody) SetMessage(v string) *DescribeDashboardResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDashboardResponseBody) SetOverview(v string) *DescribeDashboardResponseBody {
	s.Overview = &v
	return s
}

func (s *DescribeDashboardResponseBody) SetRequestId(v string) *DescribeDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardResponseBody) SetSuccess(v bool) *DescribeDashboardResponseBody {
	s.Success = &v
	return s
}

type DescribeDashboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDashboardResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardResponse) SetHeaders(v map[string]*string) *DescribeDashboardResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardResponse) SetStatusCode(v int32) *DescribeDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardResponse) SetBody(v *DescribeDashboardResponseBody) *DescribeDashboardResponse {
	s.Body = v
	return s
}

type DescribeExpireCachesRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeExpireCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpireCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpireCachesRequest) SetGatewayId(v string) *DescribeExpireCachesRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeExpireCachesRequest) SetSecurityToken(v string) *DescribeExpireCachesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeExpireCachesResponseBody struct {
	CacheFilePaths *string `json:"CacheFilePaths,omitempty" xml:"CacheFilePaths,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpireCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpireCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpireCachesResponseBody) SetCacheFilePaths(v string) *DescribeExpireCachesResponseBody {
	s.CacheFilePaths = &v
	return s
}

func (s *DescribeExpireCachesResponseBody) SetCode(v string) *DescribeExpireCachesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpireCachesResponseBody) SetMessage(v string) *DescribeExpireCachesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpireCachesResponseBody) SetRequestId(v string) *DescribeExpireCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpireCachesResponseBody) SetSuccess(v bool) *DescribeExpireCachesResponseBody {
	s.Success = &v
	return s
}

type DescribeExpireCachesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExpireCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExpireCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpireCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpireCachesResponse) SetHeaders(v map[string]*string) *DescribeExpireCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpireCachesResponse) SetStatusCode(v int32) *DescribeExpireCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpireCachesResponse) SetBody(v *DescribeExpireCachesResponseBody) *DescribeExpireCachesResponse {
	s.Body = v
	return s
}

type DescribeExpressSyncSharesRequest struct {
	ExpressSyncIds *string `json:"ExpressSyncIds,omitempty" xml:"ExpressSyncIds,omitempty"`
	IsExternal     *bool   `json:"IsExternal,omitempty" xml:"IsExternal,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeExpressSyncSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncSharesRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncSharesRequest) SetExpressSyncIds(v string) *DescribeExpressSyncSharesRequest {
	s.ExpressSyncIds = &v
	return s
}

func (s *DescribeExpressSyncSharesRequest) SetIsExternal(v bool) *DescribeExpressSyncSharesRequest {
	s.IsExternal = &v
	return s
}

func (s *DescribeExpressSyncSharesRequest) SetSecurityToken(v string) *DescribeExpressSyncSharesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeExpressSyncSharesResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Shares    *DescribeExpressSyncSharesResponseBodyShares `json:"Shares,omitempty" xml:"Shares,omitempty" type:"Struct"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressSyncSharesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncSharesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncSharesResponseBody) SetCode(v string) *DescribeExpressSyncSharesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBody) SetMessage(v string) *DescribeExpressSyncSharesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBody) SetRequestId(v string) *DescribeExpressSyncSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBody) SetShares(v *DescribeExpressSyncSharesResponseBodyShares) *DescribeExpressSyncSharesResponseBody {
	s.Shares = v
	return s
}

func (s *DescribeExpressSyncSharesResponseBody) SetSuccess(v bool) *DescribeExpressSyncSharesResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressSyncSharesResponseBodyShares struct {
	Share []*DescribeExpressSyncSharesResponseBodySharesShare `json:"Share,omitempty" xml:"Share,omitempty" type:"Repeated"`
}

func (s DescribeExpressSyncSharesResponseBodyShares) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncSharesResponseBodyShares) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncSharesResponseBodyShares) SetShare(v []*DescribeExpressSyncSharesResponseBodySharesShare) *DescribeExpressSyncSharesResponseBodyShares {
	s.Share = v
	return s
}

type DescribeExpressSyncSharesResponseBodySharesShare struct {
	ExpressSyncId    *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	ExpressSyncState *string `json:"ExpressSyncState,omitempty" xml:"ExpressSyncState,omitempty"`
	GatewayId        *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayName      *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	GatewayRegion    *string `json:"GatewayRegion,omitempty" xml:"GatewayRegion,omitempty"`
	MnsQueue         *string `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	ShareName        *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	StorageBundleId  *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	SyncProgress     *int32  `json:"SyncProgress,omitempty" xml:"SyncProgress,omitempty"`
}

func (s DescribeExpressSyncSharesResponseBodySharesShare) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncSharesResponseBodySharesShare) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetExpressSyncId(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.ExpressSyncId = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetExpressSyncState(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.ExpressSyncState = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetGatewayId(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.GatewayId = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetGatewayName(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.GatewayName = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetGatewayRegion(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.GatewayRegion = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetMnsQueue(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.MnsQueue = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetShareName(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.ShareName = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetStorageBundleId(v string) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeExpressSyncSharesResponseBodySharesShare) SetSyncProgress(v int32) *DescribeExpressSyncSharesResponseBodySharesShare {
	s.SyncProgress = &v
	return s
}

type DescribeExpressSyncSharesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExpressSyncSharesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExpressSyncSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncSharesResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncSharesResponse) SetHeaders(v map[string]*string) *DescribeExpressSyncSharesResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressSyncSharesResponse) SetStatusCode(v int32) *DescribeExpressSyncSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressSyncSharesResponse) SetBody(v *DescribeExpressSyncSharesResponseBody) *DescribeExpressSyncSharesResponse {
	s.Body = v
	return s
}

type DescribeExpressSyncsRequest struct {
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketPrefix  *string `json:"BucketPrefix,omitempty" xml:"BucketPrefix,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeExpressSyncsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncsRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncsRequest) SetBucketName(v string) *DescribeExpressSyncsRequest {
	s.BucketName = &v
	return s
}

func (s *DescribeExpressSyncsRequest) SetBucketPrefix(v string) *DescribeExpressSyncsRequest {
	s.BucketPrefix = &v
	return s
}

func (s *DescribeExpressSyncsRequest) SetSecurityToken(v string) *DescribeExpressSyncsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeExpressSyncsResponseBody struct {
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ExpressSyncs *DescribeExpressSyncsResponseBodyExpressSyncs `json:"ExpressSyncs,omitempty" xml:"ExpressSyncs,omitempty" type:"Struct"`
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressSyncsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncsResponseBody) SetCode(v string) *DescribeExpressSyncsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressSyncsResponseBody) SetExpressSyncs(v *DescribeExpressSyncsResponseBodyExpressSyncs) *DescribeExpressSyncsResponseBody {
	s.ExpressSyncs = v
	return s
}

func (s *DescribeExpressSyncsResponseBody) SetMessage(v string) *DescribeExpressSyncsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressSyncsResponseBody) SetRequestId(v string) *DescribeExpressSyncsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressSyncsResponseBody) SetSuccess(v bool) *DescribeExpressSyncsResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressSyncsResponseBodyExpressSyncs struct {
	ExpressSync []*DescribeExpressSyncsResponseBodyExpressSyncsExpressSync `json:"ExpressSync,omitempty" xml:"ExpressSync,omitempty" type:"Repeated"`
}

func (s DescribeExpressSyncsResponseBodyExpressSyncs) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncsResponseBodyExpressSyncs) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncs) SetExpressSync(v []*DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) *DescribeExpressSyncsResponseBodyExpressSyncs {
	s.ExpressSync = v
	return s
}

type DescribeExpressSyncsResponseBodyExpressSyncsExpressSync struct {
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketPrefix  *string `json:"BucketPrefix,omitempty" xml:"BucketPrefix,omitempty"`
	BucketRegion  *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpressSyncId *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	MnsTopic      *string `json:"MnsTopic,omitempty" xml:"MnsTopic,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetBucketName(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.BucketName = &v
	return s
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetBucketPrefix(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.BucketPrefix = &v
	return s
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetBucketRegion(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.BucketRegion = &v
	return s
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetDescription(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.Description = &v
	return s
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetExpressSyncId(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.ExpressSyncId = &v
	return s
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetMnsTopic(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.MnsTopic = &v
	return s
}

func (s *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync) SetName(v string) *DescribeExpressSyncsResponseBodyExpressSyncsExpressSync {
	s.Name = &v
	return s
}

type DescribeExpressSyncsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExpressSyncsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExpressSyncsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressSyncsResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressSyncsResponse) SetHeaders(v map[string]*string) *DescribeExpressSyncsResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressSyncsResponse) SetStatusCode(v int32) *DescribeExpressSyncsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressSyncsResponse) SetBody(v *DescribeExpressSyncsResponseBody) *DescribeExpressSyncsResponse {
	s.Body = v
	return s
}

type DescribeGatewayRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayRequest) SetGatewayId(v string) *DescribeGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayRequest) SetSecurityToken(v string) *DescribeGatewayRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayResponseBody struct {
	ActivatedTime            *int64                                   `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	BuyURL                   *string                                  `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Capacity                 *int32                                   `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Category                 *string                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	Code                     *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	CommonBuyInstanceId      *string                                  `json:"CommonBuyInstanceId,omitempty" xml:"CommonBuyInstanceId,omitempty"`
	CreatedTime              *int64                                   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DataLoadInterval         *int32                                   `json:"DataLoadInterval,omitempty" xml:"DataLoadInterval,omitempty"`
	DataLoadType             *string                                  `json:"DataLoadType,omitempty" xml:"DataLoadType,omitempty"`
	Description              *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsInstanceId            *string                                  `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	ElasticGateway           *bool                                    `json:"ElasticGateway,omitempty" xml:"ElasticGateway,omitempty"`
	ElasticNodes             *DescribeGatewayResponseBodyElasticNodes `json:"ElasticNodes,omitempty" xml:"ElasticNodes,omitempty" type:"Struct"`
	ExpireStatus             *int32                                   `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	ExpiredTime              *int64                                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GatewayClass             *string                                  `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	GatewayId                *string                                  `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayRegionId          *string                                  `json:"GatewayRegionId,omitempty" xml:"GatewayRegionId,omitempty"`
	GatewayType              *string                                  `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	GatewayVersion           *string                                  `json:"GatewayVersion,omitempty" xml:"GatewayVersion,omitempty"`
	InnerIp                  *string                                  `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InnerIpv6Ip              *string                                  `json:"InnerIpv6Ip,omitempty" xml:"InnerIpv6Ip,omitempty"`
	Ip                       *string                                  `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IsPostPaid               *bool                                    `json:"IsPostPaid,omitempty" xml:"IsPostPaid,omitempty"`
	IsReleaseAfterExpiration *bool                                    `json:"IsReleaseAfterExpiration,omitempty" xml:"IsReleaseAfterExpiration,omitempty"`
	LastErrorKey             *string                                  `json:"LastErrorKey,omitempty" xml:"LastErrorKey,omitempty"`
	Location                 *string                                  `json:"Location,omitempty" xml:"Location,omitempty"`
	MaxThroughput            *int32                                   `json:"MaxThroughput,omitempty" xml:"MaxThroughput,omitempty"`
	Message                  *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Name                     *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	PublicNetworkBandwidth   *int32                                   `json:"PublicNetworkBandwidth,omitempty" xml:"PublicNetworkBandwidth,omitempty"`
	RenewURL                 *string                                  `json:"RenewURL,omitempty" xml:"RenewURL,omitempty"`
	RequestId                *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                   *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageBundleId          *string                                  `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	Success                  *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId                   *string                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type                     *string                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId                *string                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                    *string                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponseBody) SetActivatedTime(v int64) *DescribeGatewayResponseBody {
	s.ActivatedTime = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetBuyURL(v string) *DescribeGatewayResponseBody {
	s.BuyURL = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetCapacity(v int32) *DescribeGatewayResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetCategory(v string) *DescribeGatewayResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetCode(v string) *DescribeGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetCommonBuyInstanceId(v string) *DescribeGatewayResponseBody {
	s.CommonBuyInstanceId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetCreatedTime(v int64) *DescribeGatewayResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetDataLoadInterval(v int32) *DescribeGatewayResponseBody {
	s.DataLoadInterval = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetDataLoadType(v string) *DescribeGatewayResponseBody {
	s.DataLoadType = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetDescription(v string) *DescribeGatewayResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetEcsInstanceId(v string) *DescribeGatewayResponseBody {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetElasticGateway(v bool) *DescribeGatewayResponseBody {
	s.ElasticGateway = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetElasticNodes(v *DescribeGatewayResponseBodyElasticNodes) *DescribeGatewayResponseBody {
	s.ElasticNodes = v
	return s
}

func (s *DescribeGatewayResponseBody) SetExpireStatus(v int32) *DescribeGatewayResponseBody {
	s.ExpireStatus = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetExpiredTime(v int64) *DescribeGatewayResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayClass(v string) *DescribeGatewayResponseBody {
	s.GatewayClass = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayId(v string) *DescribeGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayRegionId(v string) *DescribeGatewayResponseBody {
	s.GatewayRegionId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayType(v string) *DescribeGatewayResponseBody {
	s.GatewayType = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayVersion(v string) *DescribeGatewayResponseBody {
	s.GatewayVersion = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInnerIp(v string) *DescribeGatewayResponseBody {
	s.InnerIp = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInnerIpv6Ip(v string) *DescribeGatewayResponseBody {
	s.InnerIpv6Ip = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIp(v string) *DescribeGatewayResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIsPostPaid(v bool) *DescribeGatewayResponseBody {
	s.IsPostPaid = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIsReleaseAfterExpiration(v bool) *DescribeGatewayResponseBody {
	s.IsReleaseAfterExpiration = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetLastErrorKey(v string) *DescribeGatewayResponseBody {
	s.LastErrorKey = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetLocation(v string) *DescribeGatewayResponseBody {
	s.Location = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetMaxThroughput(v int32) *DescribeGatewayResponseBody {
	s.MaxThroughput = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetMessage(v string) *DescribeGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetName(v string) *DescribeGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetPublicNetworkBandwidth(v int32) *DescribeGatewayResponseBody {
	s.PublicNetworkBandwidth = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetRenewURL(v string) *DescribeGatewayResponseBody {
	s.RenewURL = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetRequestId(v string) *DescribeGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetStatus(v string) *DescribeGatewayResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetStorageBundleId(v string) *DescribeGatewayResponseBody {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetSuccess(v bool) *DescribeGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetTaskId(v string) *DescribeGatewayResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetType(v string) *DescribeGatewayResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetVSwitchId(v string) *DescribeGatewayResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetVpcId(v string) *DescribeGatewayResponseBody {
	s.VpcId = &v
	return s
}

type DescribeGatewayResponseBodyElasticNodes struct {
	ElasticNode []*string `json:"ElasticNode,omitempty" xml:"ElasticNode,omitempty" type:"Repeated"`
}

func (s DescribeGatewayResponseBodyElasticNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayResponseBodyElasticNodes) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponseBodyElasticNodes) SetElasticNode(v []*string) *DescribeGatewayResponseBodyElasticNodes {
	s.ElasticNode = v
	return s
}

type DescribeGatewayResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponse) SetHeaders(v map[string]*string) *DescribeGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayResponse) SetStatusCode(v int32) *DescribeGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayResponse) SetBody(v *DescribeGatewayResponseBody) *DescribeGatewayResponse {
	s.Body = v
	return s
}

type DescribeGatewayADInfoRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayADInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayADInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayADInfoRequest) SetGatewayId(v string) *DescribeGatewayADInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayADInfoRequest) SetSecurityToken(v string) *DescribeGatewayADInfoRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayADInfoResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IsEnabled  *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerIp   *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeGatewayADInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayADInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayADInfoResponseBody) SetCode(v string) *DescribeGatewayADInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetDomainName(v string) *DescribeGatewayADInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetIsEnabled(v bool) *DescribeGatewayADInfoResponseBody {
	s.IsEnabled = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetMessage(v string) *DescribeGatewayADInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetRequestId(v string) *DescribeGatewayADInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetServerIp(v string) *DescribeGatewayADInfoResponseBody {
	s.ServerIp = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetSuccess(v bool) *DescribeGatewayADInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayADInfoResponseBody) SetUsername(v string) *DescribeGatewayADInfoResponseBody {
	s.Username = &v
	return s
}

type DescribeGatewayADInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayADInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayADInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayADInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayADInfoResponse) SetHeaders(v map[string]*string) *DescribeGatewayADInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayADInfoResponse) SetStatusCode(v int32) *DescribeGatewayADInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayADInfoResponse) SetBody(v *DescribeGatewayADInfoResponseBody) *DescribeGatewayADInfoResponse {
	s.Body = v
	return s
}

type DescribeGatewayActionsRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayActionsRequest) SetGatewayId(v string) *DescribeGatewayActionsRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayActionsRequest) SetSecurityToken(v string) *DescribeGatewayActionsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayActionsResponseBody struct {
	Actions   *DescribeGatewayActionsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayActionsResponseBody) SetActions(v *DescribeGatewayActionsResponseBodyActions) *DescribeGatewayActionsResponseBody {
	s.Actions = v
	return s
}

func (s *DescribeGatewayActionsResponseBody) SetCode(v string) *DescribeGatewayActionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayActionsResponseBody) SetMessage(v string) *DescribeGatewayActionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayActionsResponseBody) SetRequestId(v string) *DescribeGatewayActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayActionsResponseBody) SetSuccess(v bool) *DescribeGatewayActionsResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayActionsResponseBodyActions struct {
	Action []*DescribeGatewayActionsResponseBodyActionsAction `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
}

func (s DescribeGatewayActionsResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayActionsResponseBodyActions) GoString() string {
	return s.String()
}

func (s *DescribeGatewayActionsResponseBodyActions) SetAction(v []*DescribeGatewayActionsResponseBodyActionsAction) *DescribeGatewayActionsResponseBodyActions {
	s.Action = v
	return s
}

type DescribeGatewayActionsResponseBodyActionsAction struct {
	AdLdap    *string `json:"AdLdap,omitempty" xml:"AdLdap,omitempty"`
	Cache     *string `json:"Cache,omitempty" xml:"Cache,omitempty"`
	Disk      *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Monitor   *string `json:"Monitor,omitempty" xml:"Monitor,omitempty"`
	Self      *string `json:"Self,omitempty" xml:"Self,omitempty"`
	SmbUser   *string `json:"SmbUser,omitempty" xml:"SmbUser,omitempty"`
	Target    *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DescribeGatewayActionsResponseBodyActionsAction) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayActionsResponseBodyActionsAction) GoString() string {
	return s.String()
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetAdLdap(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.AdLdap = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetCache(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.Cache = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetDisk(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.Disk = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetGatewayId(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetMonitor(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.Monitor = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetSelf(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.Self = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetSmbUser(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.SmbUser = &v
	return s
}

func (s *DescribeGatewayActionsResponseBodyActionsAction) SetTarget(v string) *DescribeGatewayActionsResponseBodyActionsAction {
	s.Target = &v
	return s
}

type DescribeGatewayActionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayActionsResponse) SetHeaders(v map[string]*string) *DescribeGatewayActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayActionsResponse) SetStatusCode(v int32) *DescribeGatewayActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayActionsResponse) SetBody(v *DescribeGatewayActionsResponseBody) *DescribeGatewayActionsResponse {
	s.Body = v
	return s
}

type DescribeGatewayAuthInfoRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAuthInfoRequest) SetGatewayId(v string) *DescribeGatewayAuthInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayAuthInfoRequest) SetSecurityToken(v string) *DescribeGatewayAuthInfoRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayAuthInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Password  *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PublicIp  *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Username  *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeGatewayAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAuthInfoResponseBody) SetCode(v string) *DescribeGatewayAuthInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponseBody) SetMessage(v string) *DescribeGatewayAuthInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponseBody) SetPassword(v string) *DescribeGatewayAuthInfoResponseBody {
	s.Password = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponseBody) SetPublicIp(v string) *DescribeGatewayAuthInfoResponseBody {
	s.PublicIp = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponseBody) SetRequestId(v string) *DescribeGatewayAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponseBody) SetSuccess(v bool) *DescribeGatewayAuthInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponseBody) SetUsername(v string) *DescribeGatewayAuthInfoResponseBody {
	s.Username = &v
	return s
}

type DescribeGatewayAuthInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAuthInfoResponse) SetHeaders(v map[string]*string) *DescribeGatewayAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayAuthInfoResponse) SetStatusCode(v int32) *DescribeGatewayAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayAuthInfoResponse) SetBody(v *DescribeGatewayAuthInfoResponseBody) *DescribeGatewayAuthInfoResponse {
	s.Body = v
	return s
}

type DescribeGatewayAutoPlansRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayAutoPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoPlansRequest) SetGatewayId(v string) *DescribeGatewayAutoPlansRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayAutoPlansRequest) SetPageNumber(v int32) *DescribeGatewayAutoPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayAutoPlansRequest) SetPageSize(v int32) *DescribeGatewayAutoPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayAutoPlansRequest) SetSecurityToken(v string) *DescribeGatewayAutoPlansRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayAutoPlansResponseBody struct {
	AutoPlans  []*DescribeGatewayAutoPlansResponseBodyAutoPlans `json:"AutoPlans,omitempty" xml:"AutoPlans,omitempty" type:"Repeated"`
	Code       *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGatewayAutoPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoPlansResponseBody) SetAutoPlans(v []*DescribeGatewayAutoPlansResponseBodyAutoPlans) *DescribeGatewayAutoPlansResponseBody {
	s.AutoPlans = v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetCode(v string) *DescribeGatewayAutoPlansResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetMessage(v string) *DescribeGatewayAutoPlansResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetPageNumber(v int32) *DescribeGatewayAutoPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetPageSize(v int32) *DescribeGatewayAutoPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetRequestId(v string) *DescribeGatewayAutoPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetSuccess(v bool) *DescribeGatewayAutoPlansResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBody) SetTotalCount(v int32) *DescribeGatewayAutoPlansResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGatewayAutoPlansResponseBodyAutoPlans struct {
	Detail    *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	PlanId    *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGatewayAutoPlansResponseBodyAutoPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoPlansResponseBodyAutoPlans) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoPlansResponseBodyAutoPlans) SetDetail(v string) *DescribeGatewayAutoPlansResponseBodyAutoPlans {
	s.Detail = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBodyAutoPlans) SetEndTime(v int64) *DescribeGatewayAutoPlansResponseBodyAutoPlans {
	s.EndTime = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBodyAutoPlans) SetEvent(v string) *DescribeGatewayAutoPlansResponseBodyAutoPlans {
	s.Event = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBodyAutoPlans) SetPlanId(v string) *DescribeGatewayAutoPlansResponseBodyAutoPlans {
	s.PlanId = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBodyAutoPlans) SetStartTime(v int64) *DescribeGatewayAutoPlansResponseBodyAutoPlans {
	s.StartTime = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponseBodyAutoPlans) SetStatus(v string) *DescribeGatewayAutoPlansResponseBodyAutoPlans {
	s.Status = &v
	return s
}

type DescribeGatewayAutoPlansResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayAutoPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayAutoPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoPlansResponse) SetHeaders(v map[string]*string) *DescribeGatewayAutoPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayAutoPlansResponse) SetStatusCode(v int32) *DescribeGatewayAutoPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayAutoPlansResponse) SetBody(v *DescribeGatewayAutoPlansResponseBody) *DescribeGatewayAutoPlansResponse {
	s.Body = v
	return s
}

type DescribeGatewayAutoUpgradeConfigurationRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayAutoUpgradeConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoUpgradeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoUpgradeConfigurationRequest) SetGatewayId(v string) *DescribeGatewayAutoUpgradeConfigurationRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationRequest) SetSecurityToken(v string) *DescribeGatewayAutoUpgradeConfigurationRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayAutoUpgradeConfigurationResponseBody struct {
	AutoUpgradeEndHour   *int32  `json:"AutoUpgradeEndHour,omitempty" xml:"AutoUpgradeEndHour,omitempty"`
	AutoUpgradeStartHour *int32  `json:"AutoUpgradeStartHour,omitempty" xml:"AutoUpgradeStartHour,omitempty"`
	Code                 *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsAutoUpgrade        *bool   `json:"IsAutoUpgrade,omitempty" xml:"IsAutoUpgrade,omitempty"`
	Message              *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayAutoUpgradeConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoUpgradeConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetAutoUpgradeEndHour(v int32) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.AutoUpgradeEndHour = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetAutoUpgradeStartHour(v int32) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.AutoUpgradeStartHour = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetCode(v string) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetIsAutoUpgrade(v bool) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.IsAutoUpgrade = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetMessage(v string) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetRequestId(v string) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponseBody) SetSuccess(v bool) *DescribeGatewayAutoUpgradeConfigurationResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayAutoUpgradeConfigurationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayAutoUpgradeConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayAutoUpgradeConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayAutoUpgradeConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponse) SetHeaders(v map[string]*string) *DescribeGatewayAutoUpgradeConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponse) SetStatusCode(v int32) *DescribeGatewayAutoUpgradeConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayAutoUpgradeConfigurationResponse) SetBody(v *DescribeGatewayAutoUpgradeConfigurationResponseBody) *DescribeGatewayAutoUpgradeConfigurationResponse {
	s.Body = v
	return s
}

type DescribeGatewayBlockVolumesRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Refresh       *bool   `json:"Refresh,omitempty" xml:"Refresh,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayBlockVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBlockVolumesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBlockVolumesRequest) SetGatewayId(v string) *DescribeGatewayBlockVolumesRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayBlockVolumesRequest) SetIndexId(v string) *DescribeGatewayBlockVolumesRequest {
	s.IndexId = &v
	return s
}

func (s *DescribeGatewayBlockVolumesRequest) SetRefresh(v bool) *DescribeGatewayBlockVolumesRequest {
	s.Refresh = &v
	return s
}

func (s *DescribeGatewayBlockVolumesRequest) SetSecurityToken(v string) *DescribeGatewayBlockVolumesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayBlockVolumesResponseBody struct {
	BlockVolumes *DescribeGatewayBlockVolumesResponseBodyBlockVolumes `json:"BlockVolumes,omitempty" xml:"BlockVolumes,omitempty" type:"Struct"`
	Code         *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayBlockVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBlockVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBlockVolumesResponseBody) SetBlockVolumes(v *DescribeGatewayBlockVolumesResponseBodyBlockVolumes) *DescribeGatewayBlockVolumesResponseBody {
	s.BlockVolumes = v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBody) SetCode(v string) *DescribeGatewayBlockVolumesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBody) SetMessage(v string) *DescribeGatewayBlockVolumesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBody) SetRequestId(v string) *DescribeGatewayBlockVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBody) SetSuccess(v bool) *DescribeGatewayBlockVolumesResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayBlockVolumesResponseBodyBlockVolumes struct {
	BlockVolume []*DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume `json:"BlockVolume,omitempty" xml:"BlockVolume,omitempty" type:"Repeated"`
}

func (s DescribeGatewayBlockVolumesResponseBodyBlockVolumes) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBlockVolumesResponseBodyBlockVolumes) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumes) SetBlockVolume(v []*DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) *DescribeGatewayBlockVolumesResponseBodyBlockVolumes {
	s.BlockVolume = v
	return s
}

type DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	CacheMode   *string `json:"CacheMode,omitempty" xml:"CacheMode,omitempty"`
	ChapEnabled *bool   `json:"ChapEnabled,omitempty" xml:"ChapEnabled,omitempty"`
	ChapInUser  *string `json:"ChapInUser,omitempty" xml:"ChapInUser,omitempty"`
	ChunkSize   *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	DiskId      *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskType    *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	IndexId     *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	LocalPath   *string `json:"LocalPath,omitempty" xml:"LocalPath,omitempty"`
	// LUN ID。
	LunId         *int32  `json:"LunId,omitempty" xml:"LunId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssBucketSsl  *bool   `json:"OssBucketSsl,omitempty" xml:"OssBucketSsl,omitempty"`
	OssEndpoint   *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol      *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Size          *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Target        *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TotalDownload *int64  `json:"TotalDownload,omitempty" xml:"TotalDownload,omitempty"`
	TotalUpload   *int64  `json:"TotalUpload,omitempty" xml:"TotalUpload,omitempty"`
	VolumeState   *int32  `json:"VolumeState,omitempty" xml:"VolumeState,omitempty"`
}

func (s DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetAddress(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Address = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetCacheMode(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.CacheMode = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetChapEnabled(v bool) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.ChapEnabled = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetChapInUser(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.ChapInUser = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetChunkSize(v int32) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.ChunkSize = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetDiskId(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.DiskId = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetDiskType(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.DiskType = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetEnabled(v bool) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Enabled = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetIndexId(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.IndexId = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetLocalPath(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.LocalPath = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetLunId(v int32) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.LunId = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetName(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Name = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetOssBucketName(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.OssBucketName = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetOssBucketSsl(v bool) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.OssBucketSsl = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetOssEndpoint(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetPort(v int32) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Port = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetProtocol(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Protocol = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetSize(v int64) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Size = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetState(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.State = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetStatus(v int32) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Status = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetTarget(v string) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.Target = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetTotalDownload(v int64) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.TotalDownload = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetTotalUpload(v int64) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.TotalUpload = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume) SetVolumeState(v int32) *DescribeGatewayBlockVolumesResponseBodyBlockVolumesBlockVolume {
	s.VolumeState = &v
	return s
}

type DescribeGatewayBlockVolumesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayBlockVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayBlockVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBlockVolumesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBlockVolumesResponse) SetHeaders(v map[string]*string) *DescribeGatewayBlockVolumesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayBlockVolumesResponse) SetStatusCode(v int32) *DescribeGatewayBlockVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayBlockVolumesResponse) SetBody(v *DescribeGatewayBlockVolumesResponseBody) *DescribeGatewayBlockVolumesResponse {
	s.Body = v
	return s
}

type DescribeGatewayBucketCachesRequest struct {
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayBucketCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBucketCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBucketCachesRequest) SetBucketName(v string) *DescribeGatewayBucketCachesRequest {
	s.BucketName = &v
	return s
}

func (s *DescribeGatewayBucketCachesRequest) SetPageNumber(v int32) *DescribeGatewayBucketCachesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayBucketCachesRequest) SetPageSize(v int32) *DescribeGatewayBucketCachesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayBucketCachesRequest) SetSecurityToken(v string) *DescribeGatewayBucketCachesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayBucketCachesResponseBody struct {
	BucketCaches *DescribeGatewayBucketCachesResponseBodyBucketCaches `json:"BucketCaches,omitempty" xml:"BucketCaches,omitempty" type:"Struct"`
	Code         *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber   *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount   *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGatewayBucketCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBucketCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBucketCachesResponseBody) SetBucketCaches(v *DescribeGatewayBucketCachesResponseBodyBucketCaches) *DescribeGatewayBucketCachesResponseBody {
	s.BucketCaches = v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetCode(v string) *DescribeGatewayBucketCachesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetMessage(v string) *DescribeGatewayBucketCachesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetPageNumber(v int32) *DescribeGatewayBucketCachesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetPageSize(v int32) *DescribeGatewayBucketCachesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetRequestId(v string) *DescribeGatewayBucketCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetSuccess(v bool) *DescribeGatewayBucketCachesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBody) SetTotalCount(v int32) *DescribeGatewayBucketCachesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGatewayBucketCachesResponseBodyBucketCaches struct {
	BucketCache []*DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache `json:"BucketCache,omitempty" xml:"BucketCache,omitempty" type:"Repeated"`
}

func (s DescribeGatewayBucketCachesResponseBodyBucketCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBucketCachesResponseBodyBucketCaches) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCaches) SetBucketCache(v []*DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) *DescribeGatewayBucketCachesResponseBodyBucketCaches {
	s.BucketCache = v
	return s
}

type DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache struct {
	BucketName  *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CacheMode   *string `json:"CacheMode,omitempty" xml:"CacheMode,omitempty"`
	CacheStats  *string `json:"CacheStats,omitempty" xml:"CacheStats,omitempty"`
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	GatewayId   *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	Location    *string `json:"Location,omitempty" xml:"Location,omitempty"`
	MountPoint  *string `json:"MountPoint,omitempty" xml:"MountPoint,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareName   *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcCidr     *string `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetBucketName(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.BucketName = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetCacheMode(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.CacheMode = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetCacheStats(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.CacheStats = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetCategory(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.Category = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetGatewayId(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetGatewayName(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.GatewayName = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetLocation(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.Location = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetMountPoint(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.MountPoint = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetProtocol(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.Protocol = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetRegionId(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.RegionId = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetShareName(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.ShareName = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetType(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.Type = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetVpcCidr(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.VpcCidr = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache) SetVpcId(v string) *DescribeGatewayBucketCachesResponseBodyBucketCachesBucketCache {
	s.VpcId = &v
	return s
}

type DescribeGatewayBucketCachesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayBucketCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayBucketCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayBucketCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayBucketCachesResponse) SetHeaders(v map[string]*string) *DescribeGatewayBucketCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayBucketCachesResponse) SetStatusCode(v int32) *DescribeGatewayBucketCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayBucketCachesResponse) SetBody(v *DescribeGatewayBucketCachesResponseBody) *DescribeGatewayBucketCachesResponse {
	s.Body = v
	return s
}

type DescribeGatewayCachesRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCachesRequest) SetGatewayId(v string) *DescribeGatewayCachesRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayCachesRequest) SetSecurityToken(v string) *DescribeGatewayCachesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayCachesResponseBody struct {
	Caches    *DescribeGatewayCachesResponseBodyCaches `json:"Caches,omitempty" xml:"Caches,omitempty" type:"Struct"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCachesResponseBody) SetCaches(v *DescribeGatewayCachesResponseBodyCaches) *DescribeGatewayCachesResponseBody {
	s.Caches = v
	return s
}

func (s *DescribeGatewayCachesResponseBody) SetCode(v string) *DescribeGatewayCachesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayCachesResponseBody) SetMessage(v string) *DescribeGatewayCachesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayCachesResponseBody) SetRequestId(v string) *DescribeGatewayCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayCachesResponseBody) SetSuccess(v bool) *DescribeGatewayCachesResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayCachesResponseBodyCaches struct {
	Cache []*DescribeGatewayCachesResponseBodyCachesCache `json:"Cache,omitempty" xml:"Cache,omitempty" type:"Repeated"`
}

func (s DescribeGatewayCachesResponseBodyCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCachesResponseBodyCaches) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCachesResponseBodyCaches) SetCache(v []*DescribeGatewayCachesResponseBodyCachesCache) *DescribeGatewayCachesResponseBodyCaches {
	s.Cache = v
	return s
}

type DescribeGatewayCachesResponseBodyCachesCache struct {
	BuyURL       *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	CacheId      *string `json:"CacheId,omitempty" xml:"CacheId,omitempty"`
	CacheType    *string `json:"CacheType,omitempty" xml:"CacheType,omitempty"`
	ExpireStatus *int32  `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	ExpiredTime  *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// IOPS。
	Iops                   *int64  `json:"Iops,omitempty" xml:"Iops,omitempty"`
	IsDirectExpand         *string `json:"IsDirectExpand,omitempty" xml:"IsDirectExpand,omitempty"`
	IsNoPartition          *bool   `json:"IsNoPartition,omitempty" xml:"IsNoPartition,omitempty"`
	IsUsed                 *bool   `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	LocalFilePath          *string `json:"LocalFilePath,omitempty" xml:"LocalFilePath,omitempty"`
	PerformanceLevel       *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RenewURL               *string `json:"RenewURL,omitempty" xml:"RenewURL,omitempty"`
	SizeInGB               *int64  `json:"SizeInGB,omitempty" xml:"SizeInGB,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DescribeGatewayCachesResponseBodyCachesCache) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCachesResponseBodyCachesCache) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetBuyURL(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.BuyURL = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetCacheId(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.CacheId = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetCacheType(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.CacheType = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetExpireStatus(v int32) *DescribeGatewayCachesResponseBodyCachesCache {
	s.ExpireStatus = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetExpiredTime(v int64) *DescribeGatewayCachesResponseBodyCachesCache {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetIops(v int64) *DescribeGatewayCachesResponseBodyCachesCache {
	s.Iops = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetIsDirectExpand(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.IsDirectExpand = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetIsNoPartition(v bool) *DescribeGatewayCachesResponseBodyCachesCache {
	s.IsNoPartition = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetIsUsed(v bool) *DescribeGatewayCachesResponseBodyCachesCache {
	s.IsUsed = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetLocalFilePath(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.LocalFilePath = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetPerformanceLevel(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetRenewURL(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.RenewURL = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetSizeInGB(v int64) *DescribeGatewayCachesResponseBodyCachesCache {
	s.SizeInGB = &v
	return s
}

func (s *DescribeGatewayCachesResponseBodyCachesCache) SetSubscriptionInstanceId(v string) *DescribeGatewayCachesResponseBodyCachesCache {
	s.SubscriptionInstanceId = &v
	return s
}

type DescribeGatewayCachesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCachesResponse) SetHeaders(v map[string]*string) *DescribeGatewayCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayCachesResponse) SetStatusCode(v int32) *DescribeGatewayCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayCachesResponse) SetBody(v *DescribeGatewayCachesResponseBody) *DescribeGatewayCachesResponse {
	s.Body = v
	return s
}

type DescribeGatewayCapacityLimitRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SizeInGB      *int64  `json:"SizeInGB,omitempty" xml:"SizeInGB,omitempty"`
}

func (s DescribeGatewayCapacityLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCapacityLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCapacityLimitRequest) SetGatewayId(v string) *DescribeGatewayCapacityLimitRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayCapacityLimitRequest) SetSecurityToken(v string) *DescribeGatewayCapacityLimitRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGatewayCapacityLimitRequest) SetSizeInGB(v int64) *DescribeGatewayCapacityLimitRequest {
	s.SizeInGB = &v
	return s
}

type DescribeGatewayCapacityLimitResponseBody struct {
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FileNumber         *int64  `json:"FileNumber,omitempty" xml:"FileNumber,omitempty"`
	FileSystemSizeInTB *int64  `json:"FileSystemSizeInTB,omitempty" xml:"FileSystemSizeInTB,omitempty"`
	IsMetadataSeparate *bool   `json:"IsMetadataSeparate,omitempty" xml:"IsMetadataSeparate,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayCapacityLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCapacityLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetCode(v string) *DescribeGatewayCapacityLimitResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetFileNumber(v int64) *DescribeGatewayCapacityLimitResponseBody {
	s.FileNumber = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetFileSystemSizeInTB(v int64) *DescribeGatewayCapacityLimitResponseBody {
	s.FileSystemSizeInTB = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetIsMetadataSeparate(v bool) *DescribeGatewayCapacityLimitResponseBody {
	s.IsMetadataSeparate = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetMessage(v string) *DescribeGatewayCapacityLimitResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetRequestId(v string) *DescribeGatewayCapacityLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponseBody) SetSuccess(v bool) *DescribeGatewayCapacityLimitResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayCapacityLimitResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayCapacityLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayCapacityLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCapacityLimitResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCapacityLimitResponse) SetHeaders(v map[string]*string) *DescribeGatewayCapacityLimitResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayCapacityLimitResponse) SetStatusCode(v int32) *DescribeGatewayCapacityLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayCapacityLimitResponse) SetBody(v *DescribeGatewayCapacityLimitResponseBody) *DescribeGatewayCapacityLimitResponse {
	s.Body = v
	return s
}

type DescribeGatewayCategoriesRequest struct {
	GatewayLocation *string `json:"GatewayLocation,omitempty" xml:"GatewayLocation,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCategoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCategoriesRequest) SetGatewayLocation(v string) *DescribeGatewayCategoriesRequest {
	s.GatewayLocation = &v
	return s
}

func (s *DescribeGatewayCategoriesRequest) SetSecurityToken(v string) *DescribeGatewayCategoriesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayCategoriesResponseBody struct {
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCategoriesResponseBody) SetCategories(v string) *DescribeGatewayCategoriesResponseBody {
	s.Categories = &v
	return s
}

func (s *DescribeGatewayCategoriesResponseBody) SetCode(v string) *DescribeGatewayCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayCategoriesResponseBody) SetMessage(v string) *DescribeGatewayCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayCategoriesResponseBody) SetRequestId(v string) *DescribeGatewayCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayCategoriesResponseBody) SetSuccess(v bool) *DescribeGatewayCategoriesResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayCategoriesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCategoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCategoriesResponse) SetHeaders(v map[string]*string) *DescribeGatewayCategoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayCategoriesResponse) SetStatusCode(v int32) *DescribeGatewayCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayCategoriesResponse) SetBody(v *DescribeGatewayCategoriesResponseBody) *DescribeGatewayCategoriesResponse {
	s.Body = v
	return s
}

type DescribeGatewayClassesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayClassesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayClassesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayClassesRequest) SetSecurityToken(v string) *DescribeGatewayClassesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayClassesResponseBody struct {
	Classes   *string `json:"Classes,omitempty" xml:"Classes,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayClassesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayClassesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayClassesResponseBody) SetClasses(v string) *DescribeGatewayClassesResponseBody {
	s.Classes = &v
	return s
}

func (s *DescribeGatewayClassesResponseBody) SetCode(v string) *DescribeGatewayClassesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayClassesResponseBody) SetMessage(v string) *DescribeGatewayClassesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayClassesResponseBody) SetRequestId(v string) *DescribeGatewayClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayClassesResponseBody) SetSuccess(v bool) *DescribeGatewayClassesResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayClassesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayClassesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayClassesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayClassesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayClassesResponse) SetHeaders(v map[string]*string) *DescribeGatewayClassesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayClassesResponse) SetStatusCode(v int32) *DescribeGatewayClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayClassesResponse) SetBody(v *DescribeGatewayClassesResponseBody) *DescribeGatewayClassesResponse {
	s.Body = v
	return s
}

type DescribeGatewayCredentialRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCredentialRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCredentialRequest) SetGatewayId(v string) *DescribeGatewayCredentialRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayCredentialRequest) SetSecurityToken(v string) *DescribeGatewayCredentialRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayCredentialResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ConsolePassword *string `json:"ConsolePassword,omitempty" xml:"ConsolePassword,omitempty"`
	ConsoleUsername *string `json:"ConsoleUsername,omitempty" xml:"ConsoleUsername,omitempty"`
	EcsIp           *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
	EcsPassword     *string `json:"EcsPassword,omitempty" xml:"EcsPassword,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewayCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCredentialResponseBody) SetCode(v string) *DescribeGatewayCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetConsolePassword(v string) *DescribeGatewayCredentialResponseBody {
	s.ConsolePassword = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetConsoleUsername(v string) *DescribeGatewayCredentialResponseBody {
	s.ConsoleUsername = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetEcsIp(v string) *DescribeGatewayCredentialResponseBody {
	s.EcsIp = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetEcsPassword(v string) *DescribeGatewayCredentialResponseBody {
	s.EcsPassword = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetMessage(v string) *DescribeGatewayCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetRequestId(v string) *DescribeGatewayCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetSuccess(v bool) *DescribeGatewayCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetVSwitchId(v string) *DescribeGatewayCredentialResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeGatewayCredentialResponseBody) SetVpcId(v string) *DescribeGatewayCredentialResponseBody {
	s.VpcId = &v
	return s
}

type DescribeGatewayCredentialResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayCredentialResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayCredentialResponse) SetHeaders(v map[string]*string) *DescribeGatewayCredentialResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayCredentialResponse) SetStatusCode(v int32) *DescribeGatewayCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayCredentialResponse) SetBody(v *DescribeGatewayCredentialResponseBody) *DescribeGatewayCredentialResponse {
	s.Body = v
	return s
}

type DescribeGatewayDNSRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayDNSRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayDNSRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayDNSRequest) SetGatewayId(v string) *DescribeGatewayDNSRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayDNSRequest) SetSecurityToken(v string) *DescribeGatewayDNSRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayDNSResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DnsServer *string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayDNSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayDNSResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayDNSResponseBody) SetCode(v string) *DescribeGatewayDNSResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayDNSResponseBody) SetDnsServer(v string) *DescribeGatewayDNSResponseBody {
	s.DnsServer = &v
	return s
}

func (s *DescribeGatewayDNSResponseBody) SetMessage(v string) *DescribeGatewayDNSResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayDNSResponseBody) SetRequestId(v string) *DescribeGatewayDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayDNSResponseBody) SetSuccess(v bool) *DescribeGatewayDNSResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayDNSResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayDNSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayDNSResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayDNSResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayDNSResponse) SetHeaders(v map[string]*string) *DescribeGatewayDNSResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayDNSResponse) SetStatusCode(v int32) *DescribeGatewayDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayDNSResponse) SetBody(v *DescribeGatewayDNSResponseBody) *DescribeGatewayDNSResponse {
	s.Body = v
	return s
}

type DescribeGatewayFileSharesRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Refresh       *bool   `json:"Refresh,omitempty" xml:"Refresh,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayFileSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileSharesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileSharesRequest) SetGatewayId(v string) *DescribeGatewayFileSharesRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayFileSharesRequest) SetIndexId(v string) *DescribeGatewayFileSharesRequest {
	s.IndexId = &v
	return s
}

func (s *DescribeGatewayFileSharesRequest) SetRefresh(v bool) *DescribeGatewayFileSharesRequest {
	s.Refresh = &v
	return s
}

func (s *DescribeGatewayFileSharesRequest) SetSecurityToken(v string) *DescribeGatewayFileSharesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayFileSharesResponseBody struct {
	Code       *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	FileShares *DescribeGatewayFileSharesResponseBodyFileShares `json:"FileShares,omitempty" xml:"FileShares,omitempty" type:"Struct"`
	Message    *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayFileSharesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileSharesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileSharesResponseBody) SetCode(v string) *DescribeGatewayFileSharesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBody) SetFileShares(v *DescribeGatewayFileSharesResponseBodyFileShares) *DescribeGatewayFileSharesResponseBody {
	s.FileShares = v
	return s
}

func (s *DescribeGatewayFileSharesResponseBody) SetMessage(v string) *DescribeGatewayFileSharesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBody) SetRequestId(v string) *DescribeGatewayFileSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBody) SetSuccess(v bool) *DescribeGatewayFileSharesResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayFileSharesResponseBodyFileShares struct {
	FileShare []*DescribeGatewayFileSharesResponseBodyFileSharesFileShare `json:"FileShare,omitempty" xml:"FileShare,omitempty" type:"Repeated"`
}

func (s DescribeGatewayFileSharesResponseBodyFileShares) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileSharesResponseBodyFileShares) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileSharesResponseBodyFileShares) SetFileShare(v []*DescribeGatewayFileSharesResponseBodyFileSharesFileShare) *DescribeGatewayFileSharesResponseBodyFileShares {
	s.FileShare = v
	return s
}

type DescribeGatewayFileSharesResponseBodyFileSharesFileShare struct {
	AccessBasedEnumeration *bool   `json:"AccessBasedEnumeration,omitempty" xml:"AccessBasedEnumeration,omitempty"`
	ActiveMessages         *int64  `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	Address                *string `json:"Address,omitempty" xml:"Address,omitempty"`
	BeLimit                *int32  `json:"BeLimit,omitempty" xml:"BeLimit,omitempty"`
	Browsable              *bool   `json:"Browsable,omitempty" xml:"Browsable,omitempty"`
	BucketInfos            *string `json:"BucketInfos,omitempty" xml:"BucketInfos,omitempty"`
	BucketsStub            *bool   `json:"BucketsStub,omitempty" xml:"BucketsStub,omitempty"`
	BypassCacheRead        *bool   `json:"BypassCacheRead,omitempty" xml:"BypassCacheRead,omitempty"`
	CacheMode              *string `json:"CacheMode,omitempty" xml:"CacheMode,omitempty"`
	ClientSideCmk          *string `json:"ClientSideCmk,omitempty" xml:"ClientSideCmk,omitempty"`
	ClientSideEncryption   *bool   `json:"ClientSideEncryption,omitempty" xml:"ClientSideEncryption,omitempty"`
	DirectIO               *bool   `json:"DirectIO,omitempty" xml:"DirectIO,omitempty"`
	DiskId                 *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskType               *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DownloadLimit          *int32  `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	DownloadQueue          *int64  `json:"DownloadQueue,omitempty" xml:"DownloadQueue,omitempty"`
	DownloadRate           *int64  `json:"DownloadRate,omitempty" xml:"DownloadRate,omitempty"`
	Enabled                *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExpressSyncId          *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	FastReclaim            *bool   `json:"FastReclaim,omitempty" xml:"FastReclaim,omitempty"`
	FeLimit                *int32  `json:"FeLimit,omitempty" xml:"FeLimit,omitempty"`
	FileNumLimit           *int64  `json:"FileNumLimit,omitempty" xml:"FileNumLimit,omitempty"`
	FsSizeLimit            *int64  `json:"FsSizeLimit,omitempty" xml:"FsSizeLimit,omitempty"`
	HighWatermark          *int32  `json:"HighWatermark,omitempty" xml:"HighWatermark,omitempty"`
	IgnoreDelete           *bool   `json:"IgnoreDelete,omitempty" xml:"IgnoreDelete,omitempty"`
	InPlace                *bool   `json:"InPlace,omitempty" xml:"InPlace,omitempty"`
	InRate                 *int64  `json:"InRate,omitempty" xml:"InRate,omitempty"`
	IndexId                *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	KmsRotatePeriod        *string `json:"KmsRotatePeriod,omitempty" xml:"KmsRotatePeriod,omitempty"`
	LagPeriod              *int64  `json:"LagPeriod,omitempty" xml:"LagPeriod,omitempty"`
	LocalPath              *string `json:"LocalPath,omitempty" xml:"LocalPath,omitempty"`
	LowWatermark           *int32  `json:"LowWatermark,omitempty" xml:"LowWatermark,omitempty"`
	MnsHealth              *string `json:"MnsHealth,omitempty" xml:"MnsHealth,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NfsV4Optimization      *bool   `json:"NfsV4Optimization,omitempty" xml:"NfsV4Optimization,omitempty"`
	NoPartition            *bool   `json:"NoPartition,omitempty" xml:"NoPartition,omitempty"`
	ObsoleteBuckets        *string `json:"ObsoleteBuckets,omitempty" xml:"ObsoleteBuckets,omitempty"`
	OssBucketName          *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssBucketSsl           *bool   `json:"OssBucketSsl,omitempty" xml:"OssBucketSsl,omitempty"`
	OssEndpoint            *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssHealth              *string `json:"OssHealth,omitempty" xml:"OssHealth,omitempty"`
	OssUsed                *int64  `json:"OssUsed,omitempty" xml:"OssUsed,omitempty"`
	OutRate                *int64  `json:"OutRate,omitempty" xml:"OutRate,omitempty"`
	PartialSyncPaths       *string `json:"PartialSyncPaths,omitempty" xml:"PartialSyncPaths,omitempty"`
	// OSS Prefix。
	PathPrefix           *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	PollingInterval      *int32  `json:"PollingInterval,omitempty" xml:"PollingInterval,omitempty"`
	Protocol             *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RemainingMetaSpace   *int64  `json:"RemainingMetaSpace,omitempty" xml:"RemainingMetaSpace,omitempty"`
	RemoteSync           *bool   `json:"RemoteSync,omitempty" xml:"RemoteSync,omitempty"`
	RemoteSyncDownload   *bool   `json:"RemoteSyncDownload,omitempty" xml:"RemoteSyncDownload,omitempty"`
	RoClientList         *string `json:"RoClientList,omitempty" xml:"RoClientList,omitempty"`
	RoUserList           *string `json:"RoUserList,omitempty" xml:"RoUserList,omitempty"`
	RwClientList         *string `json:"RwClientList,omitempty" xml:"RwClientList,omitempty"`
	RwUserList           *string `json:"RwUserList,omitempty" xml:"RwUserList,omitempty"`
	ServerSideAlgorithm  *string `json:"ServerSideAlgorithm,omitempty" xml:"ServerSideAlgorithm,omitempty"`
	ServerSideCmk        *string `json:"ServerSideCmk,omitempty" xml:"ServerSideCmk,omitempty"`
	ServerSideEncryption *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	Size                 *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Squash               *string `json:"Squash,omitempty" xml:"Squash,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportArchive       *bool   `json:"SupportArchive,omitempty" xml:"SupportArchive,omitempty"`
	SyncProgress         *int32  `json:"SyncProgress,omitempty" xml:"SyncProgress,omitempty"`
	Throttling           *bool   `json:"Throttling,omitempty" xml:"Throttling,omitempty"`
	TotalDownload        *int64  `json:"TotalDownload,omitempty" xml:"TotalDownload,omitempty"`
	TotalUpload          *int64  `json:"TotalUpload,omitempty" xml:"TotalUpload,omitempty"`
	TransferAcceleration *bool   `json:"TransferAcceleration,omitempty" xml:"TransferAcceleration,omitempty"`
	UploadQueue          *int64  `json:"UploadQueue,omitempty" xml:"UploadQueue,omitempty"`
	Used                 *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	WindowsAcl           *bool   `json:"WindowsAcl,omitempty" xml:"WindowsAcl,omitempty"`
}

func (s DescribeGatewayFileSharesResponseBodyFileSharesFileShare) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileSharesResponseBodyFileSharesFileShare) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetAccessBasedEnumeration(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.AccessBasedEnumeration = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetActiveMessages(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ActiveMessages = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetAddress(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Address = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetBeLimit(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.BeLimit = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetBrowsable(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Browsable = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetBucketInfos(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.BucketInfos = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetBucketsStub(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.BucketsStub = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetBypassCacheRead(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.BypassCacheRead = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetCacheMode(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.CacheMode = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetClientSideCmk(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ClientSideCmk = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetClientSideEncryption(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ClientSideEncryption = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetDirectIO(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.DirectIO = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetDiskId(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.DiskId = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetDiskType(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.DiskType = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetDownloadLimit(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.DownloadLimit = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetDownloadQueue(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.DownloadQueue = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetDownloadRate(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.DownloadRate = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetEnabled(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Enabled = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetExpressSyncId(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ExpressSyncId = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetFastReclaim(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.FastReclaim = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetFeLimit(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.FeLimit = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetFileNumLimit(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.FileNumLimit = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetFsSizeLimit(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.FsSizeLimit = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetHighWatermark(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.HighWatermark = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetIgnoreDelete(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.IgnoreDelete = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetInPlace(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.InPlace = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetInRate(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.InRate = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetIndexId(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.IndexId = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetKmsRotatePeriod(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.KmsRotatePeriod = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetLagPeriod(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.LagPeriod = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetLocalPath(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.LocalPath = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetLowWatermark(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.LowWatermark = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetMnsHealth(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.MnsHealth = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetName(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Name = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetNfsV4Optimization(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.NfsV4Optimization = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetNoPartition(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.NoPartition = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetObsoleteBuckets(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ObsoleteBuckets = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetOssBucketName(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.OssBucketName = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetOssBucketSsl(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.OssBucketSsl = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetOssEndpoint(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetOssHealth(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.OssHealth = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetOssUsed(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.OssUsed = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetOutRate(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.OutRate = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetPartialSyncPaths(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.PartialSyncPaths = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetPathPrefix(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.PathPrefix = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetPollingInterval(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.PollingInterval = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetProtocol(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Protocol = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRemainingMetaSpace(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RemainingMetaSpace = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRemoteSync(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RemoteSync = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRemoteSyncDownload(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RemoteSyncDownload = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRoClientList(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RoClientList = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRoUserList(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RoUserList = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRwClientList(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RwClientList = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetRwUserList(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.RwUserList = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetServerSideAlgorithm(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ServerSideAlgorithm = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetServerSideCmk(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ServerSideCmk = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetServerSideEncryption(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.ServerSideEncryption = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetSize(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Size = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetSquash(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Squash = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetState(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.State = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetStatus(v string) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Status = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetSupportArchive(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.SupportArchive = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetSyncProgress(v int32) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.SyncProgress = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetThrottling(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Throttling = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetTotalDownload(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.TotalDownload = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetTotalUpload(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.TotalUpload = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetTransferAcceleration(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.TransferAcceleration = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetUploadQueue(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.UploadQueue = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetUsed(v int64) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.Used = &v
	return s
}

func (s *DescribeGatewayFileSharesResponseBodyFileSharesFileShare) SetWindowsAcl(v bool) *DescribeGatewayFileSharesResponseBodyFileSharesFileShare {
	s.WindowsAcl = &v
	return s
}

type DescribeGatewayFileSharesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayFileSharesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayFileSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileSharesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileSharesResponse) SetHeaders(v map[string]*string) *DescribeGatewayFileSharesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayFileSharesResponse) SetStatusCode(v int32) *DescribeGatewayFileSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayFileSharesResponse) SetBody(v *DescribeGatewayFileSharesResponseBody) *DescribeGatewayFileSharesResponse {
	s.Body = v
	return s
}

type DescribeGatewayFileStatusRequest struct {
	FilePath      *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayFileStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileStatusRequest) SetFilePath(v string) *DescribeGatewayFileStatusRequest {
	s.FilePath = &v
	return s
}

func (s *DescribeGatewayFileStatusRequest) SetGatewayId(v string) *DescribeGatewayFileStatusRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayFileStatusRequest) SetIndexId(v string) *DescribeGatewayFileStatusRequest {
	s.IndexId = &v
	return s
}

func (s *DescribeGatewayFileStatusRequest) SetSecurityToken(v string) *DescribeGatewayFileStatusRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayFileStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayFileStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileStatusResponseBody) SetCode(v string) *DescribeGatewayFileStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayFileStatusResponseBody) SetMessage(v string) *DescribeGatewayFileStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayFileStatusResponseBody) SetRequestId(v string) *DescribeGatewayFileStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayFileStatusResponseBody) SetStatus(v string) *DescribeGatewayFileStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGatewayFileStatusResponseBody) SetSuccess(v bool) *DescribeGatewayFileStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayFileStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayFileStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayFileStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayFileStatusResponse) SetHeaders(v map[string]*string) *DescribeGatewayFileStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayFileStatusResponse) SetStatusCode(v int32) *DescribeGatewayFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayFileStatusResponse) SetBody(v *DescribeGatewayFileStatusResponseBody) *DescribeGatewayFileStatusResponse {
	s.Body = v
	return s
}

type DescribeGatewayImagesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeGatewayImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayImagesRequest) SetSecurityToken(v string) *DescribeGatewayImagesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGatewayImagesRequest) SetType(v string) *DescribeGatewayImagesRequest {
	s.Type = &v
	return s
}

type DescribeGatewayImagesResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Images    *DescribeGatewayImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayImagesResponseBody) SetCode(v string) *DescribeGatewayImagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayImagesResponseBody) SetImages(v *DescribeGatewayImagesResponseBodyImages) *DescribeGatewayImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeGatewayImagesResponseBody) SetMessage(v string) *DescribeGatewayImagesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayImagesResponseBody) SetRequestId(v string) *DescribeGatewayImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayImagesResponseBody) SetSuccess(v bool) *DescribeGatewayImagesResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayImagesResponseBodyImages struct {
	Image []*DescribeGatewayImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeGatewayImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeGatewayImagesResponseBodyImages) SetImage(v []*DescribeGatewayImagesResponseBodyImagesImage) *DescribeGatewayImagesResponseBodyImages {
	s.Image = v
	return s
}

type DescribeGatewayImagesResponseBodyImagesImage struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MD5          *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	ModifiedDate *string `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeGatewayImagesResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetDescription(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Description = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetMD5(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.MD5 = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetModifiedDate(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.ModifiedDate = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetName(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Name = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetSize(v int64) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Size = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetTitle(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Title = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetType(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Type = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetUrl(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Url = &v
	return s
}

func (s *DescribeGatewayImagesResponseBodyImagesImage) SetVersion(v string) *DescribeGatewayImagesResponseBodyImagesImage {
	s.Version = &v
	return s
}

type DescribeGatewayImagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayImagesResponse) SetHeaders(v map[string]*string) *DescribeGatewayImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayImagesResponse) SetStatusCode(v int32) *DescribeGatewayImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayImagesResponse) SetBody(v *DescribeGatewayImagesResponseBody) *DescribeGatewayImagesResponse {
	s.Body = v
	return s
}

type DescribeGatewayInfoRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayInfoRequest) SetGatewayId(v string) *DescribeGatewayInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayInfoRequest) SetSecurityToken(v string) *DescribeGatewayInfoRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayInfoResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayInfos *DescribeGatewayInfoResponseBodyGatewayInfos `json:"GatewayInfos,omitempty" xml:"GatewayInfos,omitempty" type:"Struct"`
	Message      *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayInfoResponseBody) SetCode(v string) *DescribeGatewayInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayInfoResponseBody) SetGatewayInfos(v *DescribeGatewayInfoResponseBodyGatewayInfos) *DescribeGatewayInfoResponseBody {
	s.GatewayInfos = v
	return s
}

func (s *DescribeGatewayInfoResponseBody) SetMessage(v string) *DescribeGatewayInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayInfoResponseBody) SetRequestId(v string) *DescribeGatewayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayInfoResponseBody) SetSuccess(v bool) *DescribeGatewayInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayInfoResponseBodyGatewayInfos struct {
	GatewayInfo []*DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo `json:"GatewayInfo,omitempty" xml:"GatewayInfo,omitempty" type:"Repeated"`
}

func (s DescribeGatewayInfoResponseBodyGatewayInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayInfoResponseBodyGatewayInfos) GoString() string {
	return s.String()
}

func (s *DescribeGatewayInfoResponseBodyGatewayInfos) SetGatewayInfo(v []*DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo) *DescribeGatewayInfoResponseBodyGatewayInfos {
	s.GatewayInfo = v
	return s
}

type DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo struct {
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Time *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo) GoString() string {
	return s.String()
}

func (s *DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo) SetInfo(v string) *DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo {
	s.Info = &v
	return s
}

func (s *DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo) SetTime(v int64) *DescribeGatewayInfoResponseBodyGatewayInfosGatewayInfo {
	s.Time = &v
	return s
}

type DescribeGatewayInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayInfoResponse) SetHeaders(v map[string]*string) *DescribeGatewayInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayInfoResponse) SetStatusCode(v int32) *DescribeGatewayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayInfoResponse) SetBody(v *DescribeGatewayInfoResponseBody) *DescribeGatewayInfoResponse {
	s.Body = v
	return s
}

type DescribeGatewayLDAPInfoRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayLDAPInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLDAPInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLDAPInfoRequest) SetGatewayId(v string) *DescribeGatewayLDAPInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayLDAPInfoRequest) SetSecurityToken(v string) *DescribeGatewayLDAPInfoRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayLDAPInfoResponseBody struct {
	BaseDN    *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsEnabled *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	IsTls     *bool   `json:"IsTls,omitempty" xml:"IsTls,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootDN    *string `json:"RootDN,omitempty" xml:"RootDN,omitempty"`
	ServerIp  *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayLDAPInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLDAPInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetBaseDN(v string) *DescribeGatewayLDAPInfoResponseBody {
	s.BaseDN = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetCode(v string) *DescribeGatewayLDAPInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetIsEnabled(v bool) *DescribeGatewayLDAPInfoResponseBody {
	s.IsEnabled = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetIsTls(v bool) *DescribeGatewayLDAPInfoResponseBody {
	s.IsTls = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetMessage(v string) *DescribeGatewayLDAPInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetRequestId(v string) *DescribeGatewayLDAPInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetRootDN(v string) *DescribeGatewayLDAPInfoResponseBody {
	s.RootDN = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetServerIp(v string) *DescribeGatewayLDAPInfoResponseBody {
	s.ServerIp = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponseBody) SetSuccess(v bool) *DescribeGatewayLDAPInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayLDAPInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayLDAPInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayLDAPInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLDAPInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLDAPInfoResponse) SetHeaders(v map[string]*string) *DescribeGatewayLDAPInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayLDAPInfoResponse) SetStatusCode(v int32) *DescribeGatewayLDAPInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayLDAPInfoResponse) SetBody(v *DescribeGatewayLDAPInfoResponseBody) *DescribeGatewayLDAPInfoResponse {
	s.Body = v
	return s
}

type DescribeGatewayLocationsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayLocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLocationsRequest) SetSecurityToken(v string) *DescribeGatewayLocationsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayLocationsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Locations *string `json:"Locations,omitempty" xml:"Locations,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayLocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLocationsResponseBody) SetCode(v string) *DescribeGatewayLocationsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayLocationsResponseBody) SetLocations(v string) *DescribeGatewayLocationsResponseBody {
	s.Locations = &v
	return s
}

func (s *DescribeGatewayLocationsResponseBody) SetMessage(v string) *DescribeGatewayLocationsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayLocationsResponseBody) SetRequestId(v string) *DescribeGatewayLocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayLocationsResponseBody) SetSuccess(v bool) *DescribeGatewayLocationsResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayLocationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayLocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLocationsResponse) SetHeaders(v map[string]*string) *DescribeGatewayLocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayLocationsResponse) SetStatusCode(v int32) *DescribeGatewayLocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayLocationsResponse) SetBody(v *DescribeGatewayLocationsResponseBody) *DescribeGatewayLocationsResponse {
	s.Body = v
	return s
}

type DescribeGatewayLoggingRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLoggingRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLoggingRequest) SetGatewayId(v string) *DescribeGatewayLoggingRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayLoggingRequest) SetSecurityToken(v string) *DescribeGatewayLoggingRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayLoggingResponseBody struct {
	Code                 *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayLoggingStatus *string `json:"GatewayLoggingStatus,omitempty" xml:"GatewayLoggingStatus,omitempty"`
	Message              *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsLogstore          *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	SlsProject           *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	Success              *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLoggingResponseBody) SetCode(v string) *DescribeGatewayLoggingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayLoggingResponseBody) SetGatewayLoggingStatus(v string) *DescribeGatewayLoggingResponseBody {
	s.GatewayLoggingStatus = &v
	return s
}

func (s *DescribeGatewayLoggingResponseBody) SetMessage(v string) *DescribeGatewayLoggingResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayLoggingResponseBody) SetRequestId(v string) *DescribeGatewayLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayLoggingResponseBody) SetSlsLogstore(v string) *DescribeGatewayLoggingResponseBody {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeGatewayLoggingResponseBody) SetSlsProject(v string) *DescribeGatewayLoggingResponseBody {
	s.SlsProject = &v
	return s
}

func (s *DescribeGatewayLoggingResponseBody) SetSuccess(v bool) *DescribeGatewayLoggingResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayLoggingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLoggingResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLoggingResponse) SetHeaders(v map[string]*string) *DescribeGatewayLoggingResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayLoggingResponse) SetStatusCode(v int32) *DescribeGatewayLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayLoggingResponse) SetBody(v *DescribeGatewayLoggingResponseBody) *DescribeGatewayLoggingResponse {
	s.Body = v
	return s
}

type DescribeGatewayLogsRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	LogFilePath   *string `json:"LogFilePath,omitempty" xml:"LogFilePath,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLogsRequest) SetGatewayId(v string) *DescribeGatewayLogsRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayLogsRequest) SetLogFilePath(v string) *DescribeGatewayLogsRequest {
	s.LogFilePath = &v
	return s
}

func (s *DescribeGatewayLogsRequest) SetSecurityToken(v string) *DescribeGatewayLogsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayLogsResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogFilePaths *string `json:"LogFilePaths,omitempty" xml:"LogFilePaths,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLogsResponseBody) SetCode(v string) *DescribeGatewayLogsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayLogsResponseBody) SetLogFilePaths(v string) *DescribeGatewayLogsResponseBody {
	s.LogFilePaths = &v
	return s
}

func (s *DescribeGatewayLogsResponseBody) SetMessage(v string) *DescribeGatewayLogsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayLogsResponseBody) SetRequestId(v string) *DescribeGatewayLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayLogsResponseBody) SetSuccess(v bool) *DescribeGatewayLogsResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayLogsResponse) SetHeaders(v map[string]*string) *DescribeGatewayLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayLogsResponse) SetStatusCode(v int32) *DescribeGatewayLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayLogsResponse) SetBody(v *DescribeGatewayLogsResponseBody) *DescribeGatewayLogsResponse {
	s.Body = v
	return s
}

type DescribeGatewayModificationClassesRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayModificationClassesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayModificationClassesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayModificationClassesRequest) SetGatewayId(v string) *DescribeGatewayModificationClassesRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayModificationClassesRequest) SetSecurityToken(v string) *DescribeGatewayModificationClassesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayModificationClassesResponseBody struct {
	Code                 *string                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool                                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TargetGatewayClasses *DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses `json:"TargetGatewayClasses,omitempty" xml:"TargetGatewayClasses,omitempty" type:"Struct"`
}

func (s DescribeGatewayModificationClassesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayModificationClassesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayModificationClassesResponseBody) SetCode(v string) *DescribeGatewayModificationClassesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayModificationClassesResponseBody) SetMessage(v string) *DescribeGatewayModificationClassesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayModificationClassesResponseBody) SetRequestId(v string) *DescribeGatewayModificationClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayModificationClassesResponseBody) SetSuccess(v bool) *DescribeGatewayModificationClassesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayModificationClassesResponseBody) SetTargetGatewayClasses(v *DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses) *DescribeGatewayModificationClassesResponseBody {
	s.TargetGatewayClasses = v
	return s
}

type DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses struct {
	TargetGatewayClass []*DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass `json:"TargetGatewayClass,omitempty" xml:"TargetGatewayClass,omitempty" type:"Repeated"`
}

func (s DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses) GoString() string {
	return s.String()
}

func (s *DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses) SetTargetGatewayClass(v []*DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass) *DescribeGatewayModificationClassesResponseBodyTargetGatewayClasses {
	s.TargetGatewayClass = v
	return s
}

type DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass struct {
	GatewayClass *string `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	IsAvailable  *bool   `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
}

func (s DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass) GoString() string {
	return s.String()
}

func (s *DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass) SetGatewayClass(v string) *DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass {
	s.GatewayClass = &v
	return s
}

func (s *DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass) SetIsAvailable(v bool) *DescribeGatewayModificationClassesResponseBodyTargetGatewayClassesTargetGatewayClass {
	s.IsAvailable = &v
	return s
}

type DescribeGatewayModificationClassesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayModificationClassesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayModificationClassesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayModificationClassesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayModificationClassesResponse) SetHeaders(v map[string]*string) *DescribeGatewayModificationClassesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayModificationClassesResponse) SetStatusCode(v int32) *DescribeGatewayModificationClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayModificationClassesResponse) SetBody(v *DescribeGatewayModificationClassesResponseBody) *DescribeGatewayModificationClassesResponse {
	s.Body = v
	return s
}

type DescribeGatewayNFSClientsRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayNFSClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayNFSClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayNFSClientsRequest) SetGatewayId(v string) *DescribeGatewayNFSClientsRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayNFSClientsRequest) SetPageNumber(v int32) *DescribeGatewayNFSClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayNFSClientsRequest) SetPageSize(v int32) *DescribeGatewayNFSClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayNFSClientsRequest) SetSecurityToken(v string) *DescribeGatewayNFSClientsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayNFSClientsResponseBody struct {
	ClientInfoList   *DescribeGatewayNFSClientsResponseBodyClientInfoList `json:"ClientInfoList,omitempty" xml:"ClientInfoList,omitempty" type:"Struct"`
	Code             *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber       *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount       *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Version3Enabled  *bool                                                `json:"Version3Enabled,omitempty" xml:"Version3Enabled,omitempty"`
	Version40Enabled *bool                                                `json:"Version40Enabled,omitempty" xml:"Version40Enabled,omitempty"`
	Version41Enabled *bool                                                `json:"Version41Enabled,omitempty" xml:"Version41Enabled,omitempty"`
}

func (s DescribeGatewayNFSClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayNFSClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayNFSClientsResponseBody) SetClientInfoList(v *DescribeGatewayNFSClientsResponseBodyClientInfoList) *DescribeGatewayNFSClientsResponseBody {
	s.ClientInfoList = v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetCode(v string) *DescribeGatewayNFSClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetMessage(v string) *DescribeGatewayNFSClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetPageNumber(v int32) *DescribeGatewayNFSClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetPageSize(v int32) *DescribeGatewayNFSClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetRequestId(v string) *DescribeGatewayNFSClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetSuccess(v bool) *DescribeGatewayNFSClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetTotalCount(v int32) *DescribeGatewayNFSClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetVersion3Enabled(v bool) *DescribeGatewayNFSClientsResponseBody {
	s.Version3Enabled = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetVersion40Enabled(v bool) *DescribeGatewayNFSClientsResponseBody {
	s.Version40Enabled = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBody) SetVersion41Enabled(v bool) *DescribeGatewayNFSClientsResponseBody {
	s.Version41Enabled = &v
	return s
}

type DescribeGatewayNFSClientsResponseBodyClientInfoList struct {
	ClientInfo []*DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Repeated"`
}

func (s DescribeGatewayNFSClientsResponseBodyClientInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayNFSClientsResponseBodyClientInfoList) GoString() string {
	return s.String()
}

func (s *DescribeGatewayNFSClientsResponseBodyClientInfoList) SetClientInfo(v []*DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) *DescribeGatewayNFSClientsResponseBodyClientInfoList {
	s.ClientInfo = v
	return s
}

type DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo struct {
	ClientIpAddr *string `json:"ClientIpAddr,omitempty" xml:"ClientIpAddr,omitempty"`
	HasNFSv3     *bool   `json:"HasNFSv3,omitempty" xml:"HasNFSv3,omitempty"`
	HasNFSv40    *bool   `json:"HasNFSv40,omitempty" xml:"HasNFSv40,omitempty"`
	HasNFSv41    *bool   `json:"HasNFSv41,omitempty" xml:"HasNFSv41,omitempty"`
}

func (s DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) GoString() string {
	return s.String()
}

func (s *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) SetClientIpAddr(v string) *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo {
	s.ClientIpAddr = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) SetHasNFSv3(v bool) *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo {
	s.HasNFSv3 = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) SetHasNFSv40(v bool) *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo {
	s.HasNFSv40 = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo) SetHasNFSv41(v bool) *DescribeGatewayNFSClientsResponseBodyClientInfoListClientInfo {
	s.HasNFSv41 = &v
	return s
}

type DescribeGatewayNFSClientsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayNFSClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayNFSClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayNFSClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayNFSClientsResponse) SetHeaders(v map[string]*string) *DescribeGatewayNFSClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayNFSClientsResponse) SetStatusCode(v int32) *DescribeGatewayNFSClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayNFSClientsResponse) SetBody(v *DescribeGatewayNFSClientsResponseBody) *DescribeGatewayNFSClientsResponse {
	s.Body = v
	return s
}

type DescribeGatewaySMBUsersRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewaySMBUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySMBUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySMBUsersRequest) SetGatewayId(v string) *DescribeGatewaySMBUsersRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewaySMBUsersRequest) SetPageNumber(v int32) *DescribeGatewaySMBUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewaySMBUsersRequest) SetPageSize(v int32) *DescribeGatewaySMBUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewaySMBUsersRequest) SetSecurityToken(v string) *DescribeGatewaySMBUsersRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewaySMBUsersResponseBody struct {
	ActiveDirectory *bool                                     `json:"ActiveDirectory,omitempty" xml:"ActiveDirectory,omitempty"`
	Code            *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber      *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount      *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users           *DescribeGatewaySMBUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s DescribeGatewaySMBUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySMBUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySMBUsersResponseBody) SetActiveDirectory(v bool) *DescribeGatewaySMBUsersResponseBody {
	s.ActiveDirectory = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetCode(v string) *DescribeGatewaySMBUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetMessage(v string) *DescribeGatewaySMBUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetPageNumber(v int32) *DescribeGatewaySMBUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetPageSize(v int32) *DescribeGatewaySMBUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetRequestId(v string) *DescribeGatewaySMBUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetSuccess(v bool) *DescribeGatewaySMBUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetTotalCount(v int32) *DescribeGatewaySMBUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponseBody) SetUsers(v *DescribeGatewaySMBUsersResponseBodyUsers) *DescribeGatewaySMBUsersResponseBody {
	s.Users = v
	return s
}

type DescribeGatewaySMBUsersResponseBodyUsers struct {
	User []*DescribeGatewaySMBUsersResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeGatewaySMBUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySMBUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySMBUsersResponseBodyUsers) SetUser(v []*DescribeGatewaySMBUsersResponseBodyUsersUser) *DescribeGatewaySMBUsersResponseBodyUsers {
	s.User = v
	return s
}

type DescribeGatewaySMBUsersResponseBodyUsersUser struct {
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeGatewaySMBUsersResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySMBUsersResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySMBUsersResponseBodyUsersUser) SetUsername(v string) *DescribeGatewaySMBUsersResponseBodyUsersUser {
	s.Username = &v
	return s
}

type DescribeGatewaySMBUsersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewaySMBUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewaySMBUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySMBUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySMBUsersResponse) SetHeaders(v map[string]*string) *DescribeGatewaySMBUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewaySMBUsersResponse) SetStatusCode(v int32) *DescribeGatewaySMBUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewaySMBUsersResponse) SetBody(v *DescribeGatewaySMBUsersResponseBody) *DescribeGatewaySMBUsersResponse {
	s.Body = v
	return s
}

type DescribeGatewayStatisticsRequest struct {
	EndTimestamp    *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	GatewayCategory *string `json:"GatewayCategory,omitempty" xml:"GatewayCategory,omitempty"`
	GatewayLocation *string `json:"GatewayLocation,omitempty" xml:"GatewayLocation,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTimestamp  *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	TargetAccountId *string `json:"TargetAccountId,omitempty" xml:"TargetAccountId,omitempty"`
}

func (s DescribeGatewayStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStatisticsRequest) SetEndTimestamp(v int64) *DescribeGatewayStatisticsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeGatewayStatisticsRequest) SetGatewayCategory(v string) *DescribeGatewayStatisticsRequest {
	s.GatewayCategory = &v
	return s
}

func (s *DescribeGatewayStatisticsRequest) SetGatewayLocation(v string) *DescribeGatewayStatisticsRequest {
	s.GatewayLocation = &v
	return s
}

func (s *DescribeGatewayStatisticsRequest) SetSecurityToken(v string) *DescribeGatewayStatisticsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGatewayStatisticsRequest) SetStartTimestamp(v int64) *DescribeGatewayStatisticsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeGatewayStatisticsRequest) SetTargetAccountId(v string) *DescribeGatewayStatisticsRequest {
	s.TargetAccountId = &v
	return s
}

type DescribeGatewayStatisticsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStatisticsResponseBody) SetCode(v string) *DescribeGatewayStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayStatisticsResponseBody) SetMessage(v string) *DescribeGatewayStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayStatisticsResponseBody) SetRequestId(v string) *DescribeGatewayStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayStatisticsResponseBody) SetResult(v string) *DescribeGatewayStatisticsResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeGatewayStatisticsResponseBody) SetSuccess(v bool) *DescribeGatewayStatisticsResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStatisticsResponse) SetHeaders(v map[string]*string) *DescribeGatewayStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayStatisticsResponse) SetStatusCode(v int32) *DescribeGatewayStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayStatisticsResponse) SetBody(v *DescribeGatewayStatisticsResponseBody) *DescribeGatewayStatisticsResponse {
	s.Body = v
	return s
}

type DescribeGatewayStockRequest struct {
	GatewayRegionId *string `json:"GatewayRegionId,omitempty" xml:"GatewayRegionId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayStockRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStockRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStockRequest) SetGatewayRegionId(v string) *DescribeGatewayStockRequest {
	s.GatewayRegionId = &v
	return s
}

func (s *DescribeGatewayStockRequest) SetSecurityToken(v string) *DescribeGatewayStockRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayStockResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stocks    *DescribeGatewayStockResponseBodyStocks `json:"Stocks,omitempty" xml:"Stocks,omitempty" type:"Struct"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewayStockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStockResponseBody) SetCode(v string) *DescribeGatewayStockResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayStockResponseBody) SetMessage(v string) *DescribeGatewayStockResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayStockResponseBody) SetRequestId(v string) *DescribeGatewayStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayStockResponseBody) SetStocks(v *DescribeGatewayStockResponseBodyStocks) *DescribeGatewayStockResponseBody {
	s.Stocks = v
	return s
}

func (s *DescribeGatewayStockResponseBody) SetSuccess(v bool) *DescribeGatewayStockResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewayStockResponseBodyStocks struct {
	Stock []*DescribeGatewayStockResponseBodyStocksStock `json:"Stock,omitempty" xml:"Stock,omitempty" type:"Repeated"`
}

func (s DescribeGatewayStockResponseBodyStocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStockResponseBodyStocks) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStockResponseBodyStocks) SetStock(v []*DescribeGatewayStockResponseBodyStocksStock) *DescribeGatewayStockResponseBodyStocks {
	s.Stock = v
	return s
}

type DescribeGatewayStockResponseBodyStocksStock struct {
	StockInfo *string `json:"StockInfo,omitempty" xml:"StockInfo,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeGatewayStockResponseBodyStocksStock) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStockResponseBodyStocksStock) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStockResponseBodyStocksStock) SetStockInfo(v string) *DescribeGatewayStockResponseBodyStocksStock {
	s.StockInfo = &v
	return s
}

func (s *DescribeGatewayStockResponseBodyStocksStock) SetZoneId(v string) *DescribeGatewayStockResponseBodyStocksStock {
	s.ZoneId = &v
	return s
}

type DescribeGatewayStockResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayStockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayStockResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayStockResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayStockResponse) SetHeaders(v map[string]*string) *DescribeGatewayStockResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayStockResponse) SetStatusCode(v int32) *DescribeGatewayStockResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayStockResponse) SetBody(v *DescribeGatewayStockResponseBody) *DescribeGatewayStockResponse {
	s.Body = v
	return s
}

type DescribeGatewayTypesRequest struct {
	GatewayLocation *string `json:"GatewayLocation,omitempty" xml:"GatewayLocation,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewayTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayTypesRequest) SetGatewayLocation(v string) *DescribeGatewayTypesRequest {
	s.GatewayLocation = &v
	return s
}

func (s *DescribeGatewayTypesRequest) SetSecurityToken(v string) *DescribeGatewayTypesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewayTypesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Types     *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeGatewayTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayTypesResponseBody) SetCode(v string) *DescribeGatewayTypesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewayTypesResponseBody) SetMessage(v string) *DescribeGatewayTypesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewayTypesResponseBody) SetRequestId(v string) *DescribeGatewayTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayTypesResponseBody) SetSuccess(v bool) *DescribeGatewayTypesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewayTypesResponseBody) SetTypes(v string) *DescribeGatewayTypesResponseBody {
	s.Types = &v
	return s
}

type DescribeGatewayTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewayTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewayTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayTypesResponse) SetHeaders(v map[string]*string) *DescribeGatewayTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayTypesResponse) SetStatusCode(v int32) *DescribeGatewayTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayTypesResponse) SetBody(v *DescribeGatewayTypesResponseBody) *DescribeGatewayTypesResponse {
	s.Body = v
	return s
}

type DescribeGatewaysRequest struct {
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
}

func (s DescribeGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysRequest) SetPageNumber(v int32) *DescribeGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewaysRequest) SetPageSize(v int32) *DescribeGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewaysRequest) SetSecurityToken(v string) *DescribeGatewaysRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGatewaysRequest) SetStorageBundleId(v string) *DescribeGatewaysRequest {
	s.StorageBundleId = &v
	return s
}

type DescribeGatewaysResponseBody struct {
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Gateways   *DescribeGatewaysResponseBodyGateways `json:"Gateways,omitempty" xml:"Gateways,omitempty" type:"Struct"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysResponseBody) SetCode(v string) *DescribeGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewaysResponseBody) SetGateways(v *DescribeGatewaysResponseBodyGateways) *DescribeGatewaysResponseBody {
	s.Gateways = v
	return s
}

func (s *DescribeGatewaysResponseBody) SetMessage(v string) *DescribeGatewaysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewaysResponseBody) SetPageNumber(v int32) *DescribeGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewaysResponseBody) SetPageSize(v int32) *DescribeGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewaysResponseBody) SetRequestId(v string) *DescribeGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewaysResponseBody) SetSuccess(v bool) *DescribeGatewaysResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewaysResponseBody) SetTotalCount(v int32) *DescribeGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGatewaysResponseBodyGateways struct {
	Gateway []*DescribeGatewaysResponseBodyGatewaysGateway `json:"Gateway,omitempty" xml:"Gateway,omitempty" type:"Repeated"`
}

func (s DescribeGatewaysResponseBodyGateways) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysResponseBodyGateways) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysResponseBodyGateways) SetGateway(v []*DescribeGatewaysResponseBodyGatewaysGateway) *DescribeGatewaysResponseBodyGateways {
	s.Gateway = v
	return s
}

type DescribeGatewaysResponseBodyGatewaysGateway struct {
	ActivatedTime            *int64                                                   `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	BuyURL                   *string                                                  `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Capacity                 *int32                                                   `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Category                 *string                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	CommonBuyInstanceId      *string                                                  `json:"CommonBuyInstanceId,omitempty" xml:"CommonBuyInstanceId,omitempty"`
	CreatedTime              *int64                                                   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DataLoadInterval         *int32                                                   `json:"DataLoadInterval,omitempty" xml:"DataLoadInterval,omitempty"`
	DataLoadType             *string                                                  `json:"DataLoadType,omitempty" xml:"DataLoadType,omitempty"`
	Description              *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsInstanceId            *string                                                  `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	ElasticGateway           *bool                                                    `json:"ElasticGateway,omitempty" xml:"ElasticGateway,omitempty"`
	ElasticNodes             *DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes `json:"ElasticNodes,omitempty" xml:"ElasticNodes,omitempty" type:"Struct"`
	ExpireStatus             *int32                                                   `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	ExpiredTime              *int64                                                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GatewayClass             *string                                                  `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	GatewayId                *string                                                  `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayRegionId          *string                                                  `json:"GatewayRegionId,omitempty" xml:"GatewayRegionId,omitempty"`
	GatewayType              *string                                                  `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	GatewayVersion           *string                                                  `json:"GatewayVersion,omitempty" xml:"GatewayVersion,omitempty"`
	InnerIp                  *string                                                  `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InnerIpv6Ip              *string                                                  `json:"InnerIpv6Ip,omitempty" xml:"InnerIpv6Ip,omitempty"`
	Ip                       *string                                                  `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IsPostPaid               *bool                                                    `json:"IsPostPaid,omitempty" xml:"IsPostPaid,omitempty"`
	IsReleaseAfterExpiration *bool                                                    `json:"IsReleaseAfterExpiration,omitempty" xml:"IsReleaseAfterExpiration,omitempty"`
	LastErrorKey             *string                                                  `json:"LastErrorKey,omitempty" xml:"LastErrorKey,omitempty"`
	Location                 *string                                                  `json:"Location,omitempty" xml:"Location,omitempty"`
	MaxThroughput            *int32                                                   `json:"MaxThroughput,omitempty" xml:"MaxThroughput,omitempty"`
	Name                     *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	PublicNetworkBandwidth   *int32                                                   `json:"PublicNetworkBandwidth,omitempty" xml:"PublicNetworkBandwidth,omitempty"`
	RenewURL                 *string                                                  `json:"RenewURL,omitempty" xml:"RenewURL,omitempty"`
	Status                   *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageBundleId          *string                                                  `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	TaskId                   *string                                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type                     *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId                *string                                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                    *string                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewaysResponseBodyGatewaysGateway) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysResponseBodyGatewaysGateway) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetActivatedTime(v int64) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.ActivatedTime = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetBuyURL(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.BuyURL = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetCapacity(v int32) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Capacity = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetCategory(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Category = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetCommonBuyInstanceId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.CommonBuyInstanceId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetCreatedTime(v int64) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.CreatedTime = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetDataLoadInterval(v int32) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.DataLoadInterval = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetDataLoadType(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.DataLoadType = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetDescription(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Description = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetEcsInstanceId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetElasticGateway(v bool) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.ElasticGateway = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetElasticNodes(v *DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.ElasticNodes = v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetExpireStatus(v int32) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.ExpireStatus = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetExpiredTime(v int64) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetGatewayClass(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.GatewayClass = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetGatewayId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetGatewayRegionId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.GatewayRegionId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetGatewayType(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.GatewayType = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetGatewayVersion(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.GatewayVersion = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetInnerIp(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.InnerIp = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetInnerIpv6Ip(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.InnerIpv6Ip = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetIp(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Ip = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetIsPostPaid(v bool) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.IsPostPaid = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetIsReleaseAfterExpiration(v bool) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.IsReleaseAfterExpiration = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetLastErrorKey(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.LastErrorKey = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetLocation(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Location = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetMaxThroughput(v int32) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.MaxThroughput = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetName(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Name = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetPublicNetworkBandwidth(v int32) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.PublicNetworkBandwidth = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetRenewURL(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.RenewURL = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetStatus(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Status = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetStorageBundleId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetTaskId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.TaskId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetType(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.Type = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetVSwitchId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.VSwitchId = &v
	return s
}

func (s *DescribeGatewaysResponseBodyGatewaysGateway) SetVpcId(v string) *DescribeGatewaysResponseBodyGatewaysGateway {
	s.VpcId = &v
	return s
}

type DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes struct {
	ElasticNode []*string `json:"ElasticNode,omitempty" xml:"ElasticNode,omitempty" type:"Repeated"`
}

func (s DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes) SetElasticNode(v []*string) *DescribeGatewaysResponseBodyGatewaysGatewayElasticNodes {
	s.ElasticNode = v
	return s
}

type DescribeGatewaysResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysResponse) SetHeaders(v map[string]*string) *DescribeGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewaysResponse) SetStatusCode(v int32) *DescribeGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewaysResponse) SetBody(v *DescribeGatewaysResponseBody) *DescribeGatewaysResponse {
	s.Body = v
	return s
}

type DescribeGatewaysForCmsRequest struct {
	GatewayRegionId *string `json:"GatewayRegionId,omitempty" xml:"GatewayRegionId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGatewaysForCmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysForCmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysForCmsRequest) SetGatewayRegionId(v string) *DescribeGatewaysForCmsRequest {
	s.GatewayRegionId = &v
	return s
}

func (s *DescribeGatewaysForCmsRequest) SetPageNumber(v int32) *DescribeGatewaysForCmsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewaysForCmsRequest) SetPageSize(v int32) *DescribeGatewaysForCmsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewaysForCmsRequest) SetSecurityToken(v string) *DescribeGatewaysForCmsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeGatewaysForCmsResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Gateways   *DescribeGatewaysForCmsResponseBodyGateways `json:"Gateways,omitempty" xml:"Gateways,omitempty" type:"Struct"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGatewaysForCmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysForCmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysForCmsResponseBody) SetCode(v string) *DescribeGatewaysForCmsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetGateways(v *DescribeGatewaysForCmsResponseBodyGateways) *DescribeGatewaysForCmsResponseBody {
	s.Gateways = v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetMessage(v string) *DescribeGatewaysForCmsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetPageNumber(v int32) *DescribeGatewaysForCmsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetPageSize(v int32) *DescribeGatewaysForCmsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetRequestId(v string) *DescribeGatewaysForCmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetSuccess(v bool) *DescribeGatewaysForCmsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBody) SetTotalCount(v int32) *DescribeGatewaysForCmsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGatewaysForCmsResponseBodyGateways struct {
	Gateway []*DescribeGatewaysForCmsResponseBodyGatewaysGateway `json:"Gateway,omitempty" xml:"Gateway,omitempty" type:"Repeated"`
}

func (s DescribeGatewaysForCmsResponseBodyGateways) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysForCmsResponseBodyGateways) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysForCmsResponseBodyGateways) SetGateway(v []*DescribeGatewaysForCmsResponseBodyGatewaysGateway) *DescribeGatewaysForCmsResponseBodyGateways {
	s.Gateway = v
	return s
}

type DescribeGatewaysForCmsResponseBodyGatewaysGateway struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GatewayId   *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeGatewaysForCmsResponseBodyGatewaysGateway) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysForCmsResponseBodyGatewaysGateway) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysForCmsResponseBodyGatewaysGateway) SetDescription(v string) *DescribeGatewaysForCmsResponseBodyGatewaysGateway {
	s.Description = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBodyGatewaysGateway) SetGatewayId(v string) *DescribeGatewaysForCmsResponseBodyGatewaysGateway {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewaysForCmsResponseBodyGatewaysGateway) SetName(v string) *DescribeGatewaysForCmsResponseBodyGatewaysGateway {
	s.Name = &v
	return s
}

type DescribeGatewaysForCmsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewaysForCmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewaysForCmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysForCmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysForCmsResponse) SetHeaders(v map[string]*string) *DescribeGatewaysForCmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewaysForCmsResponse) SetStatusCode(v int32) *DescribeGatewaysForCmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewaysForCmsResponse) SetBody(v *DescribeGatewaysForCmsResponseBody) *DescribeGatewaysForCmsResponse {
	s.Body = v
	return s
}

type DescribeGatewaysTagsRequest struct {
	GatewayIds      *string `json:"GatewayIds,omitempty" xml:"GatewayIds,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	TagCategory     *string `json:"TagCategory,omitempty" xml:"TagCategory,omitempty"`
}

func (s DescribeGatewaysTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsRequest) SetGatewayIds(v string) *DescribeGatewaysTagsRequest {
	s.GatewayIds = &v
	return s
}

func (s *DescribeGatewaysTagsRequest) SetSecurityToken(v string) *DescribeGatewaysTagsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGatewaysTagsRequest) SetStorageBundleId(v string) *DescribeGatewaysTagsRequest {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeGatewaysTagsRequest) SetTagCategory(v string) *DescribeGatewaysTagsRequest {
	s.TagCategory = &v
	return s
}

type DescribeGatewaysTagsResponseBody struct {
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayTags *DescribeGatewaysTagsResponseBodyGatewayTags `json:"GatewayTags,omitempty" xml:"GatewayTags,omitempty" type:"Struct"`
	Message     *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGatewaysTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsResponseBody) SetCode(v string) *DescribeGatewaysTagsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGatewaysTagsResponseBody) SetGatewayTags(v *DescribeGatewaysTagsResponseBodyGatewayTags) *DescribeGatewaysTagsResponseBody {
	s.GatewayTags = v
	return s
}

func (s *DescribeGatewaysTagsResponseBody) SetMessage(v string) *DescribeGatewaysTagsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGatewaysTagsResponseBody) SetRequestId(v string) *DescribeGatewaysTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewaysTagsResponseBody) SetSuccess(v bool) *DescribeGatewaysTagsResponseBody {
	s.Success = &v
	return s
}

type DescribeGatewaysTagsResponseBodyGatewayTags struct {
	GatewayTag []*DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag `json:"GatewayTag,omitempty" xml:"GatewayTag,omitempty" type:"Repeated"`
}

func (s DescribeGatewaysTagsResponseBodyGatewayTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsResponseBodyGatewayTags) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsResponseBodyGatewayTags) SetGatewayTag(v []*DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag) *DescribeGatewaysTagsResponseBodyGatewayTags {
	s.GatewayTag = v
	return s
}

type DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag struct {
	GatewayId *string                                                    `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Tags      *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag) SetGatewayId(v string) *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag) SetTags(v *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags) *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTag {
	s.Tags = v
	return s
}

type DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags struct {
	Tag []*DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags) SetTag(v []*DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag) *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTags {
	s.Tag = v
	return s
}

type DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag) SetTagKey(v string) *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag) SetTagValue(v string) *DescribeGatewaysTagsResponseBodyGatewayTagsGatewayTagTagsTag {
	s.TagValue = &v
	return s
}

type DescribeGatewaysTagsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewaysTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewaysTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaysTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewaysTagsResponse) SetHeaders(v map[string]*string) *DescribeGatewaysTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewaysTagsResponse) SetStatusCode(v int32) *DescribeGatewaysTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewaysTagsResponse) SetBody(v *DescribeGatewaysTagsResponseBody) *DescribeGatewaysTagsResponse {
	s.Body = v
	return s
}

type DescribeKmsKeyRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	KmsKey        *string `json:"KmsKey,omitempty" xml:"KmsKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeKmsKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKmsKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeyRequest) SetGatewayId(v string) *DescribeKmsKeyRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeKmsKeyRequest) SetKmsKey(v string) *DescribeKmsKeyRequest {
	s.KmsKey = &v
	return s
}

func (s *DescribeKmsKeyRequest) SetSecurityToken(v string) *DescribeKmsKeyRequest {
	s.SecurityToken = &v
	return s
}

type DescribeKmsKeyResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsValid   *bool   `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeKmsKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKmsKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeyResponseBody) SetCode(v string) *DescribeKmsKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeKmsKeyResponseBody) SetIsValid(v bool) *DescribeKmsKeyResponseBody {
	s.IsValid = &v
	return s
}

func (s *DescribeKmsKeyResponseBody) SetMessage(v string) *DescribeKmsKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeKmsKeyResponseBody) SetRequestId(v string) *DescribeKmsKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKmsKeyResponseBody) SetSuccess(v bool) *DescribeKmsKeyResponseBody {
	s.Success = &v
	return s
}

type DescribeKmsKeyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeKmsKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeKmsKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKmsKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeyResponse) SetHeaders(v map[string]*string) *DescribeKmsKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeKmsKeyResponse) SetStatusCode(v int32) *DescribeKmsKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKmsKeyResponse) SetBody(v *DescribeKmsKeyResponseBody) *DescribeKmsKeyResponse {
	s.Body = v
	return s
}

type DescribeMqttConfigRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeMqttConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeMqttConfigRequest) SetGatewayId(v string) *DescribeMqttConfigRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeMqttConfigRequest) SetSecurityToken(v string) *DescribeMqttConfigRequest {
	s.SecurityToken = &v
	return s
}

type DescribeMqttConfigResponseBody struct {
	AuthType          *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BrokerUrl         *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InternalBrokerUrl *string `json:"InternalBrokerUrl,omitempty" xml:"InternalBrokerUrl,omitempty"`
	IsEnabled         *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MqttInstanceId    *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PublishTopic      *string `json:"PublishTopic,omitempty" xml:"PublishTopic,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscribeTopic    *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Username          *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeMqttConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMqttConfigResponseBody) SetAuthType(v string) *DescribeMqttConfigResponseBody {
	s.AuthType = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetBrokerUrl(v string) *DescribeMqttConfigResponseBody {
	s.BrokerUrl = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetCode(v string) *DescribeMqttConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetGroupId(v string) *DescribeMqttConfigResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetInternalBrokerUrl(v string) *DescribeMqttConfigResponseBody {
	s.InternalBrokerUrl = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetIsEnabled(v bool) *DescribeMqttConfigResponseBody {
	s.IsEnabled = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetMessage(v string) *DescribeMqttConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetMqttInstanceId(v string) *DescribeMqttConfigResponseBody {
	s.MqttInstanceId = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetPassword(v string) *DescribeMqttConfigResponseBody {
	s.Password = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetPublishTopic(v string) *DescribeMqttConfigResponseBody {
	s.PublishTopic = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetRequestId(v string) *DescribeMqttConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetSubscribeTopic(v string) *DescribeMqttConfigResponseBody {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetSuccess(v bool) *DescribeMqttConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMqttConfigResponseBody) SetUsername(v string) *DescribeMqttConfigResponseBody {
	s.Username = &v
	return s
}

type DescribeMqttConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMqttConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMqttConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMqttConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMqttConfigResponse) SetHeaders(v map[string]*string) *DescribeMqttConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMqttConfigResponse) SetStatusCode(v int32) *DescribeMqttConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMqttConfigResponse) SetBody(v *DescribeMqttConfigResponseBody) *DescribeMqttConfigResponse {
	s.Body = v
	return s
}

type DescribeOssBucketInfoRequest struct {
	BucketEndpoint *string `json:"BucketEndpoint,omitempty" xml:"BucketEndpoint,omitempty"`
	BucketName     *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	GatewayId      *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeOssBucketInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketInfoRequest) SetBucketEndpoint(v string) *DescribeOssBucketInfoRequest {
	s.BucketEndpoint = &v
	return s
}

func (s *DescribeOssBucketInfoRequest) SetBucketName(v string) *DescribeOssBucketInfoRequest {
	s.BucketName = &v
	return s
}

func (s *DescribeOssBucketInfoRequest) SetGatewayId(v string) *DescribeOssBucketInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *DescribeOssBucketInfoRequest) SetSecurityToken(v string) *DescribeOssBucketInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeOssBucketInfoRequest) SetType(v string) *DescribeOssBucketInfoRequest {
	s.Type = &v
	return s
}

type DescribeOssBucketInfoResponseBody struct {
	Code                          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsArchive                     *bool   `json:"IsArchive,omitempty" xml:"IsArchive,omitempty"`
	IsBackToResource              *bool   `json:"IsBackToResource,omitempty" xml:"IsBackToResource,omitempty"`
	IsColdArchive                 *bool   `json:"IsColdArchive,omitempty" xml:"IsColdArchive,omitempty"`
	IsFresh                       *bool   `json:"IsFresh,omitempty" xml:"IsFresh,omitempty"`
	IsSupportServerSideEncryption *bool   `json:"IsSupportServerSideEncryption,omitempty" xml:"IsSupportServerSideEncryption,omitempty"`
	IsVersioning                  *bool   `json:"IsVersioning,omitempty" xml:"IsVersioning,omitempty"`
	Message                       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PollingInterval               *int32  `json:"PollingInterval,omitempty" xml:"PollingInterval,omitempty"`
	RequestId                     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageSize                   *int64  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	Success                       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOssBucketInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketInfoResponseBody) SetCode(v string) *DescribeOssBucketInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetIsArchive(v bool) *DescribeOssBucketInfoResponseBody {
	s.IsArchive = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetIsBackToResource(v bool) *DescribeOssBucketInfoResponseBody {
	s.IsBackToResource = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetIsColdArchive(v bool) *DescribeOssBucketInfoResponseBody {
	s.IsColdArchive = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetIsFresh(v bool) *DescribeOssBucketInfoResponseBody {
	s.IsFresh = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetIsSupportServerSideEncryption(v bool) *DescribeOssBucketInfoResponseBody {
	s.IsSupportServerSideEncryption = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetIsVersioning(v bool) *DescribeOssBucketInfoResponseBody {
	s.IsVersioning = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetMessage(v string) *DescribeOssBucketInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetPollingInterval(v int32) *DescribeOssBucketInfoResponseBody {
	s.PollingInterval = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetRequestId(v string) *DescribeOssBucketInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetStorageSize(v int64) *DescribeOssBucketInfoResponseBody {
	s.StorageSize = &v
	return s
}

func (s *DescribeOssBucketInfoResponseBody) SetSuccess(v bool) *DescribeOssBucketInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeOssBucketInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOssBucketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssBucketInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketInfoResponse) SetHeaders(v map[string]*string) *DescribeOssBucketInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssBucketInfoResponse) SetStatusCode(v int32) *DescribeOssBucketInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssBucketInfoResponse) SetBody(v *DescribeOssBucketInfoResponseBody) *DescribeOssBucketInfoResponse {
	s.Body = v
	return s
}

type DescribeOssBucketsRequest struct {
	BucketEndpoint *string `json:"BucketEndpoint,omitempty" xml:"BucketEndpoint,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeOssBucketsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketsRequest) SetBucketEndpoint(v string) *DescribeOssBucketsRequest {
	s.BucketEndpoint = &v
	return s
}

func (s *DescribeOssBucketsRequest) SetSecurityToken(v string) *DescribeOssBucketsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeOssBucketsResponseBody struct {
	Buckets   *DescribeOssBucketsResponseBodyBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Struct"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOssBucketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketsResponseBody) SetBuckets(v *DescribeOssBucketsResponseBodyBuckets) *DescribeOssBucketsResponseBody {
	s.Buckets = v
	return s
}

func (s *DescribeOssBucketsResponseBody) SetCode(v string) *DescribeOssBucketsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOssBucketsResponseBody) SetMessage(v string) *DescribeOssBucketsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOssBucketsResponseBody) SetRequestId(v string) *DescribeOssBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssBucketsResponseBody) SetSuccess(v bool) *DescribeOssBucketsResponseBody {
	s.Success = &v
	return s
}

type DescribeOssBucketsResponseBodyBuckets struct {
	Bucket []*DescribeOssBucketsResponseBodyBucketsBucket `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Repeated"`
}

func (s DescribeOssBucketsResponseBodyBuckets) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketsResponseBodyBuckets) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketsResponseBodyBuckets) SetBucket(v []*DescribeOssBucketsResponseBodyBucketsBucket) *DescribeOssBucketsResponseBodyBuckets {
	s.Bucket = v
	return s
}

type DescribeOssBucketsResponseBodyBucketsBucket struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeOssBucketsResponseBodyBucketsBucket) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketsResponseBodyBucketsBucket) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketsResponseBodyBucketsBucket) SetName(v string) *DescribeOssBucketsResponseBodyBucketsBucket {
	s.Name = &v
	return s
}

type DescribeOssBucketsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOssBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssBucketsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssBucketsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssBucketsResponse) SetHeaders(v map[string]*string) *DescribeOssBucketsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssBucketsResponse) SetStatusCode(v int32) *DescribeOssBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssBucketsResponse) SetBody(v *DescribeOssBucketsResponseBody) *DescribeOssBucketsResponse {
	s.Body = v
	return s
}

type DescribePayAsYouGoPriceRequest struct {
	GatewayClass  *string `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePayAsYouGoPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePayAsYouGoPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePayAsYouGoPriceRequest) SetGatewayClass(v string) *DescribePayAsYouGoPriceRequest {
	s.GatewayClass = &v
	return s
}

func (s *DescribePayAsYouGoPriceRequest) SetRegionId(v string) *DescribePayAsYouGoPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePayAsYouGoPriceRequest) SetSecurityToken(v string) *DescribePayAsYouGoPriceRequest {
	s.SecurityToken = &v
	return s
}

type DescribePayAsYouGoPriceResponseBody struct {
	CacheCloudEfficiencySizePrice *float32 `json:"CacheCloudEfficiencySizePrice,omitempty" xml:"CacheCloudEfficiencySizePrice,omitempty"`
	CacheCloudSSDSizePrice        *float32 `json:"CacheCloudSSDSizePrice,omitempty" xml:"CacheCloudSSDSizePrice,omitempty"`
	CacheESSDPl1SizePrice         *float32 `json:"CacheESSDPl1SizePrice,omitempty" xml:"CacheESSDPl1SizePrice,omitempty"`
	Code                          *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency                      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	GatewayClassPrice             *float32 `json:"GatewayClassPrice,omitempty" xml:"GatewayClassPrice,omitempty"`
	Message                       *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                     *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                       *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePayAsYouGoPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePayAsYouGoPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePayAsYouGoPriceResponseBody) SetCacheCloudEfficiencySizePrice(v float32) *DescribePayAsYouGoPriceResponseBody {
	s.CacheCloudEfficiencySizePrice = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetCacheCloudSSDSizePrice(v float32) *DescribePayAsYouGoPriceResponseBody {
	s.CacheCloudSSDSizePrice = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetCacheESSDPl1SizePrice(v float32) *DescribePayAsYouGoPriceResponseBody {
	s.CacheESSDPl1SizePrice = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetCode(v string) *DescribePayAsYouGoPriceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetCurrency(v string) *DescribePayAsYouGoPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetGatewayClassPrice(v float32) *DescribePayAsYouGoPriceResponseBody {
	s.GatewayClassPrice = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetMessage(v string) *DescribePayAsYouGoPriceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetRequestId(v string) *DescribePayAsYouGoPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponseBody) SetSuccess(v bool) *DescribePayAsYouGoPriceResponseBody {
	s.Success = &v
	return s
}

type DescribePayAsYouGoPriceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePayAsYouGoPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePayAsYouGoPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePayAsYouGoPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePayAsYouGoPriceResponse) SetHeaders(v map[string]*string) *DescribePayAsYouGoPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePayAsYouGoPriceResponse) SetStatusCode(v int32) *DescribePayAsYouGoPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePayAsYouGoPriceResponse) SetBody(v *DescribePayAsYouGoPriceResponseBody) *DescribePayAsYouGoPriceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSharesBucketInfoForExpressSyncRequest struct {
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketRegion  *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSharesBucketInfoForExpressSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSharesBucketInfoForExpressSyncRequest) GoString() string {
	return s.String()
}

func (s *DescribeSharesBucketInfoForExpressSyncRequest) SetBucketName(v string) *DescribeSharesBucketInfoForExpressSyncRequest {
	s.BucketName = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncRequest) SetBucketRegion(v string) *DescribeSharesBucketInfoForExpressSyncRequest {
	s.BucketRegion = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncRequest) SetSecurityToken(v string) *DescribeSharesBucketInfoForExpressSyncRequest {
	s.SecurityToken = &v
	return s
}

type DescribeSharesBucketInfoForExpressSyncResponseBody struct {
	BucketInfos *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos `json:"BucketInfos,omitempty" xml:"BucketInfos,omitempty" type:"Struct"`
	Code        *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSharesBucketInfoForExpressSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSharesBucketInfoForExpressSyncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBody) SetBucketInfos(v *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos) *DescribeSharesBucketInfoForExpressSyncResponseBody {
	s.BucketInfos = v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBody) SetCode(v string) *DescribeSharesBucketInfoForExpressSyncResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBody) SetMessage(v string) *DescribeSharesBucketInfoForExpressSyncResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBody) SetRequestId(v string) *DescribeSharesBucketInfoForExpressSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBody) SetSuccess(v bool) *DescribeSharesBucketInfoForExpressSyncResponseBody {
	s.Success = &v
	return s
}

type DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos struct {
	BucketInfo []*DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo `json:"BucketInfo,omitempty" xml:"BucketInfo,omitempty" type:"Repeated"`
}

func (s DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos) GoString() string {
	return s.String()
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos) SetBucketInfo(v []*DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo) *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfos {
	s.BucketInfo = v
	return s
}

type DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo struct {
	BucketName   *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketPrefix *string `json:"BucketPrefix,omitempty" xml:"BucketPrefix,omitempty"`
	BucketRegion *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
}

func (s DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo) GoString() string {
	return s.String()
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo) SetBucketName(v string) *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo {
	s.BucketName = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo) SetBucketPrefix(v string) *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo {
	s.BucketPrefix = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo) SetBucketRegion(v string) *DescribeSharesBucketInfoForExpressSyncResponseBodyBucketInfosBucketInfo {
	s.BucketRegion = &v
	return s
}

type DescribeSharesBucketInfoForExpressSyncResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSharesBucketInfoForExpressSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSharesBucketInfoForExpressSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSharesBucketInfoForExpressSyncResponse) GoString() string {
	return s.String()
}

func (s *DescribeSharesBucketInfoForExpressSyncResponse) SetHeaders(v map[string]*string) *DescribeSharesBucketInfoForExpressSyncResponse {
	s.Headers = v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponse) SetStatusCode(v int32) *DescribeSharesBucketInfoForExpressSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSharesBucketInfoForExpressSyncResponse) SetBody(v *DescribeSharesBucketInfoForExpressSyncResponseBody) *DescribeSharesBucketInfoForExpressSyncResponse {
	s.Body = v
	return s
}

type DescribeStorageBundleRequest struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
}

func (s DescribeStorageBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundleRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundleRequest) SetSecurityToken(v string) *DescribeStorageBundleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeStorageBundleRequest) SetStorageBundleId(v string) *DescribeStorageBundleRequest {
	s.StorageBundleId = &v
	return s
}

type DescribeStorageBundleResponseBody struct {
	BackendBucketRegionId *string `json:"BackendBucketRegionId,omitempty" xml:"BackendBucketRegionId,omitempty"`
	Code                  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreatedTime           *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message               *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageBundleId       *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	Success               *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStorageBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundleResponseBody) SetBackendBucketRegionId(v string) *DescribeStorageBundleResponseBody {
	s.BackendBucketRegionId = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetCode(v string) *DescribeStorageBundleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetCreatedTime(v int64) *DescribeStorageBundleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetDescription(v string) *DescribeStorageBundleResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetMessage(v string) *DescribeStorageBundleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetName(v string) *DescribeStorageBundleResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetRequestId(v string) *DescribeStorageBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetStorageBundleId(v string) *DescribeStorageBundleResponseBody {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeStorageBundleResponseBody) SetSuccess(v bool) *DescribeStorageBundleResponseBody {
	s.Success = &v
	return s
}

type DescribeStorageBundleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStorageBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStorageBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundleResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundleResponse) SetHeaders(v map[string]*string) *DescribeStorageBundleResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageBundleResponse) SetStatusCode(v int32) *DescribeStorageBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageBundleResponse) SetBody(v *DescribeStorageBundleResponseBody) *DescribeStorageBundleResponse {
	s.Body = v
	return s
}

type DescribeStorageBundlesRequest struct {
	BackendBucketRegionId *string `json:"BackendBucketRegionId,omitempty" xml:"BackendBucketRegionId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeStorageBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundlesRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundlesRequest) SetBackendBucketRegionId(v string) *DescribeStorageBundlesRequest {
	s.BackendBucketRegionId = &v
	return s
}

func (s *DescribeStorageBundlesRequest) SetPageNumber(v int32) *DescribeStorageBundlesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageBundlesRequest) SetPageSize(v int32) *DescribeStorageBundlesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageBundlesRequest) SetSecurityToken(v string) *DescribeStorageBundlesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeStorageBundlesResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageBundles *DescribeStorageBundlesResponseBodyStorageBundles `json:"StorageBundles,omitempty" xml:"StorageBundles,omitempty" type:"Struct"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStorageBundlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundlesResponseBody) SetCode(v string) *DescribeStorageBundlesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetMessage(v string) *DescribeStorageBundlesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetPageNumber(v int32) *DescribeStorageBundlesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetPageSize(v int32) *DescribeStorageBundlesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetRequestId(v string) *DescribeStorageBundlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetStorageBundles(v *DescribeStorageBundlesResponseBodyStorageBundles) *DescribeStorageBundlesResponseBody {
	s.StorageBundles = v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetSuccess(v bool) *DescribeStorageBundlesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeStorageBundlesResponseBody) SetTotalCount(v int32) *DescribeStorageBundlesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeStorageBundlesResponseBodyStorageBundles struct {
	StorageBundle []*DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle `json:"StorageBundle,omitempty" xml:"StorageBundle,omitempty" type:"Repeated"`
}

func (s DescribeStorageBundlesResponseBodyStorageBundles) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundlesResponseBodyStorageBundles) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundlesResponseBodyStorageBundles) SetStorageBundle(v []*DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) *DescribeStorageBundlesResponseBodyStorageBundles {
	s.StorageBundle = v
	return s
}

type DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle struct {
	BackendBucketRegionId *string `json:"BackendBucketRegionId,omitempty" xml:"BackendBucketRegionId,omitempty"`
	CreatedTime           *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StorageBundleId       *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
}

func (s DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) SetBackendBucketRegionId(v string) *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle {
	s.BackendBucketRegionId = &v
	return s
}

func (s *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) SetCreatedTime(v int64) *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) SetDescription(v string) *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle {
	s.Description = &v
	return s
}

func (s *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) SetName(v string) *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle {
	s.Name = &v
	return s
}

func (s *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle) SetStorageBundleId(v string) *DescribeStorageBundlesResponseBodyStorageBundlesStorageBundle {
	s.StorageBundleId = &v
	return s
}

type DescribeStorageBundlesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStorageBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStorageBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageBundlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageBundlesResponse) SetHeaders(v map[string]*string) *DescribeStorageBundlesResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageBundlesResponse) SetStatusCode(v int32) *DescribeStorageBundlesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageBundlesResponse) SetBody(v *DescribeStorageBundlesResponseBody) *DescribeStorageBundlesResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionPriceRequest struct {
	CacheCloudEfficiencySize *int64  `json:"CacheCloudEfficiencySize,omitempty" xml:"CacheCloudEfficiencySize,omitempty"`
	CacheESSDPl1Size         *int64  `json:"CacheESSDPl1Size,omitempty" xml:"CacheESSDPl1Size,omitempty"`
	CacheSSDSize             *int64  `json:"CacheSSDSize,omitempty" xml:"CacheSSDSize,omitempty"`
	GatewayClass             *string `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	PeriodQuantity           *int32  `json:"PeriodQuantity,omitempty" xml:"PeriodQuantity,omitempty"`
	PeriodUnit               *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken            *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSubscriptionPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPriceRequest) SetCacheCloudEfficiencySize(v int64) *DescribeSubscriptionPriceRequest {
	s.CacheCloudEfficiencySize = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetCacheESSDPl1Size(v int64) *DescribeSubscriptionPriceRequest {
	s.CacheESSDPl1Size = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetCacheSSDSize(v int64) *DescribeSubscriptionPriceRequest {
	s.CacheSSDSize = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetGatewayClass(v string) *DescribeSubscriptionPriceRequest {
	s.GatewayClass = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetPeriodQuantity(v int32) *DescribeSubscriptionPriceRequest {
	s.PeriodQuantity = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetPeriodUnit(v string) *DescribeSubscriptionPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetRegionId(v string) *DescribeSubscriptionPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionPriceRequest) SetSecurityToken(v string) *DescribeSubscriptionPriceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeSubscriptionPriceResponseBody struct {
	Code       *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency   *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	Message    *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeSubscriptionPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPriceResponseBody) SetCode(v string) *DescribeSubscriptionPriceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSubscriptionPriceResponseBody) SetCurrency(v string) *DescribeSubscriptionPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribeSubscriptionPriceResponseBody) SetMessage(v string) *DescribeSubscriptionPriceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSubscriptionPriceResponseBody) SetRequestId(v string) *DescribeSubscriptionPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionPriceResponseBody) SetSuccess(v bool) *DescribeSubscriptionPriceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionPriceResponseBody) SetTradePrice(v float32) *DescribeSubscriptionPriceResponseBody {
	s.TradePrice = &v
	return s
}

type DescribeSubscriptionPriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSubscriptionPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubscriptionPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPriceResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionPriceResponse) SetStatusCode(v int32) *DescribeSubscriptionPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionPriceResponse) SetBody(v *DescribeSubscriptionPriceResponseBody) *DescribeSubscriptionPriceResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TargetId      *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetSecurityToken(v string) *DescribeTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTasksRequest) SetTargetId(v string) *DescribeTasksRequest {
	s.TargetId = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskId(v string) *DescribeTasksRequest {
	s.TaskId = &v
	return s
}

type DescribeTasksResponseBody struct {
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Tasks      *DescribeTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetCode(v string) *DescribeTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTasksResponseBody) SetMessage(v string) *DescribeTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageSize(v int32) *DescribeTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetSuccess(v bool) *DescribeTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTasks(v *DescribeTasksResponseBodyTasks) *DescribeTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalCount(v int32) *DescribeTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTasksResponseBodyTasks struct {
	SimpleTask []*DescribeTasksResponseBodyTasksSimpleTask `json:"SimpleTask,omitempty" xml:"SimpleTask,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTasks) SetSimpleTask(v []*DescribeTasksResponseBodyTasksSimpleTask) *DescribeTasksResponseBodyTasks {
	s.SimpleTask = v
	return s
}

type DescribeTasksResponseBodyTasksSimpleTask struct {
	CreatedTime       *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	MessageKey        *string `json:"MessageKey,omitempty" xml:"MessageKey,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress          *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RelatedResourceId *string `json:"RelatedResourceId,omitempty" xml:"RelatedResourceId,omitempty"`
	StateCode         *string `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpdatedTime       *int64  `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeTasksResponseBodyTasksSimpleTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyTasksSimpleTask) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetCreatedTime(v int64) *DescribeTasksResponseBodyTasksSimpleTask {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetMessageKey(v string) *DescribeTasksResponseBodyTasksSimpleTask {
	s.MessageKey = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetName(v string) *DescribeTasksResponseBodyTasksSimpleTask {
	s.Name = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetProgress(v int32) *DescribeTasksResponseBodyTasksSimpleTask {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetRelatedResourceId(v string) *DescribeTasksResponseBodyTasksSimpleTask {
	s.RelatedResourceId = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetStateCode(v string) *DescribeTasksResponseBodyTasksSimpleTask {
	s.StateCode = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetTaskId(v string) *DescribeTasksResponseBodyTasksSimpleTask {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksSimpleTask) SetUpdatedTime(v int64) *DescribeTasksResponseBodyTasksSimpleTask {
	s.UpdatedTime = &v
	return s
}

type DescribeTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetStatusCode(v int32) *DescribeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DescribeUserBusinessStatusRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserBusinessStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBusinessStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBusinessStatusRequest) SetSecurityToken(v string) *DescribeUserBusinessStatusRequest {
	s.SecurityToken = &v
	return s
}

type DescribeUserBusinessStatusResponseBody struct {
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsEnabled         *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	IsIndebted        *bool   `json:"IsIndebted,omitempty" xml:"IsIndebted,omitempty"`
	IsIndebtedOverdue *bool   `json:"IsIndebtedOverdue,omitempty" xml:"IsIndebtedOverdue,omitempty"`
	IsRiskControl     *bool   `json:"IsRiskControl,omitempty" xml:"IsRiskControl,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeUserBusinessStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBusinessStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBusinessStatusResponseBody) SetCode(v string) *DescribeUserBusinessStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetIsEnabled(v bool) *DescribeUserBusinessStatusResponseBody {
	s.IsEnabled = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetIsIndebted(v bool) *DescribeUserBusinessStatusResponseBody {
	s.IsIndebted = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetIsIndebtedOverdue(v bool) *DescribeUserBusinessStatusResponseBody {
	s.IsIndebtedOverdue = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetIsRiskControl(v bool) *DescribeUserBusinessStatusResponseBody {
	s.IsRiskControl = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetMessage(v string) *DescribeUserBusinessStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetRequestId(v string) *DescribeUserBusinessStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBusinessStatusResponseBody) SetSuccess(v bool) *DescribeUserBusinessStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeUserBusinessStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserBusinessStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserBusinessStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBusinessStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBusinessStatusResponse) SetHeaders(v map[string]*string) *DescribeUserBusinessStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBusinessStatusResponse) SetStatusCode(v int32) *DescribeUserBusinessStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBusinessStatusResponse) SetBody(v *DescribeUserBusinessStatusResponseBody) *DescribeUserBusinessStatusResponse {
	s.Body = v
	return s
}

type DescribeVSwitchesRequest struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId  *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) SetName(v string) *DescribeVSwitchesRequest {
	s.Name = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int32) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int32) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceRegionId(v string) *DescribeVSwitchesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetSecurityToken(v string) *DescribeVSwitchesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetStorageBundleId(v string) *DescribeVSwitchesRequest {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

type DescribeVSwitchesResponseBody struct {
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitches  *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) SetCode(v string) *DescribeVSwitchesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetMessage(v string) *DescribeVSwitchesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageNumber(v int32) *DescribeVSwitchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageSize(v int32) *DescribeVSwitchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetSuccess(v bool) *DescribeVSwitchesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	VSwitch []*DescribeVSwitchesResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitch) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitch struct {
	AvailableSelectionInfo *string `json:"AvailableSelectionInfo,omitempty" xml:"AvailableSelectionInfo,omitempty"`
	Id                     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDefault              *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetAvailableSelectionInfo(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.AvailableSelectionInfo = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Id = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetName(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Name = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetZoneId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.ZoneId = &v
	return s
}

type DescribeVSwitchesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponse) SetHeaders(v map[string]*string) *DescribeVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchesResponse) SetStatusCode(v int32) *DescribeVSwitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse {
	s.Body = v
	return s
}

type DescribeVpcsRequest struct {
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId  *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequest) SetName(v string) *DescribeVpcsRequest {
	s.Name = &v
	return s
}

func (s *DescribeVpcsRequest) SetPageNumber(v int32) *DescribeVpcsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsRequest) SetPageSize(v int32) *DescribeVpcsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceRegionId(v string) *DescribeVpcsRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeVpcsRequest) SetSecurityToken(v string) *DescribeVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVpcsRequest) SetStorageBundleId(v string) *DescribeVpcsRequest {
	s.StorageBundleId = &v
	return s
}

func (s *DescribeVpcsRequest) SetVpcId(v string) *DescribeVpcsRequest {
	s.VpcId = &v
	return s
}

type DescribeVpcsResponseBody struct {
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vpcs       *DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) SetCode(v string) *DescribeVpcsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetMessage(v string) *DescribeVpcsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetPageNumber(v int32) *DescribeVpcsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetPageSize(v int32) *DescribeVpcsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetSuccess(v bool) *DescribeVpcsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetTotalCount(v int32) *DescribeVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

type DescribeVpcsResponseBodyVpcs struct {
	Vpc []*DescribeVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpc(v []*DescribeVpcsResponseBodyVpcsVpc) *DescribeVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

type DescribeVpcsResponseBodyVpcsVpc struct {
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDefault *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Id = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetName(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Name = &v
	return s
}

type DescribeVpcsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponse) SetHeaders(v map[string]*string) *DescribeVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcsResponse) SetStatusCode(v int32) *DescribeVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcsResponse) SetBody(v *DescribeVpcsResponseBody) *DescribeVpcsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetSecurityToken(v string) *DescribeZonesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeZonesResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetCode(v string) *DescribeZonesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeZonesResponseBody) SetMessage(v string) *DescribeZonesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetSuccess(v bool) *DescribeZonesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DisableGatewayLoggingRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DisableGatewayLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableGatewayLoggingRequest) GoString() string {
	return s.String()
}

func (s *DisableGatewayLoggingRequest) SetGatewayId(v string) *DisableGatewayLoggingRequest {
	s.GatewayId = &v
	return s
}

func (s *DisableGatewayLoggingRequest) SetSecurityToken(v string) *DisableGatewayLoggingRequest {
	s.SecurityToken = &v
	return s
}

type DisableGatewayLoggingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableGatewayLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableGatewayLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *DisableGatewayLoggingResponseBody) SetCode(v string) *DisableGatewayLoggingResponseBody {
	s.Code = &v
	return s
}

func (s *DisableGatewayLoggingResponseBody) SetMessage(v string) *DisableGatewayLoggingResponseBody {
	s.Message = &v
	return s
}

func (s *DisableGatewayLoggingResponseBody) SetRequestId(v string) *DisableGatewayLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableGatewayLoggingResponseBody) SetSuccess(v bool) *DisableGatewayLoggingResponseBody {
	s.Success = &v
	return s
}

type DisableGatewayLoggingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableGatewayLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableGatewayLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableGatewayLoggingResponse) GoString() string {
	return s.String()
}

func (s *DisableGatewayLoggingResponse) SetHeaders(v map[string]*string) *DisableGatewayLoggingResponse {
	s.Headers = v
	return s
}

func (s *DisableGatewayLoggingResponse) SetStatusCode(v int32) *DisableGatewayLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableGatewayLoggingResponse) SetBody(v *DisableGatewayLoggingResponseBody) *DisableGatewayLoggingResponse {
	s.Body = v
	return s
}

type DisableGatewayNFSVersionRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	NFSVersion    *string `json:"NFSVersion,omitempty" xml:"NFSVersion,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DisableGatewayNFSVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableGatewayNFSVersionRequest) GoString() string {
	return s.String()
}

func (s *DisableGatewayNFSVersionRequest) SetGatewayId(v string) *DisableGatewayNFSVersionRequest {
	s.GatewayId = &v
	return s
}

func (s *DisableGatewayNFSVersionRequest) SetNFSVersion(v string) *DisableGatewayNFSVersionRequest {
	s.NFSVersion = &v
	return s
}

func (s *DisableGatewayNFSVersionRequest) SetSecurityToken(v string) *DisableGatewayNFSVersionRequest {
	s.SecurityToken = &v
	return s
}

type DisableGatewayNFSVersionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DisableGatewayNFSVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableGatewayNFSVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableGatewayNFSVersionResponseBody) SetCode(v string) *DisableGatewayNFSVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DisableGatewayNFSVersionResponseBody) SetMessage(v string) *DisableGatewayNFSVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DisableGatewayNFSVersionResponseBody) SetRequestId(v string) *DisableGatewayNFSVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableGatewayNFSVersionResponseBody) SetSuccess(v bool) *DisableGatewayNFSVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DisableGatewayNFSVersionResponseBody) SetTaskId(v string) *DisableGatewayNFSVersionResponseBody {
	s.TaskId = &v
	return s
}

type DisableGatewayNFSVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableGatewayNFSVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableGatewayNFSVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableGatewayNFSVersionResponse) GoString() string {
	return s.String()
}

func (s *DisableGatewayNFSVersionResponse) SetHeaders(v map[string]*string) *DisableGatewayNFSVersionResponse {
	s.Headers = v
	return s
}

func (s *DisableGatewayNFSVersionResponse) SetStatusCode(v int32) *DisableGatewayNFSVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableGatewayNFSVersionResponse) SetBody(v *DisableGatewayNFSVersionResponseBody) *DisableGatewayNFSVersionResponse {
	s.Body = v
	return s
}

type EnableGatewayIpv6Request struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s EnableGatewayIpv6Request) String() string {
	return tea.Prettify(s)
}

func (s EnableGatewayIpv6Request) GoString() string {
	return s.String()
}

func (s *EnableGatewayIpv6Request) SetGatewayId(v string) *EnableGatewayIpv6Request {
	s.GatewayId = &v
	return s
}

func (s *EnableGatewayIpv6Request) SetSecurityToken(v string) *EnableGatewayIpv6Request {
	s.SecurityToken = &v
	return s
}

type EnableGatewayIpv6ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EnableGatewayIpv6ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableGatewayIpv6ResponseBody) GoString() string {
	return s.String()
}

func (s *EnableGatewayIpv6ResponseBody) SetCode(v string) *EnableGatewayIpv6ResponseBody {
	s.Code = &v
	return s
}

func (s *EnableGatewayIpv6ResponseBody) SetMessage(v string) *EnableGatewayIpv6ResponseBody {
	s.Message = &v
	return s
}

func (s *EnableGatewayIpv6ResponseBody) SetRequestId(v string) *EnableGatewayIpv6ResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableGatewayIpv6ResponseBody) SetSuccess(v bool) *EnableGatewayIpv6ResponseBody {
	s.Success = &v
	return s
}

func (s *EnableGatewayIpv6ResponseBody) SetTaskId(v string) *EnableGatewayIpv6ResponseBody {
	s.TaskId = &v
	return s
}

type EnableGatewayIpv6Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableGatewayIpv6ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableGatewayIpv6Response) String() string {
	return tea.Prettify(s)
}

func (s EnableGatewayIpv6Response) GoString() string {
	return s.String()
}

func (s *EnableGatewayIpv6Response) SetHeaders(v map[string]*string) *EnableGatewayIpv6Response {
	s.Headers = v
	return s
}

func (s *EnableGatewayIpv6Response) SetStatusCode(v int32) *EnableGatewayIpv6Response {
	s.StatusCode = &v
	return s
}

func (s *EnableGatewayIpv6Response) SetBody(v *EnableGatewayIpv6ResponseBody) *EnableGatewayIpv6Response {
	s.Body = v
	return s
}

type EnableGatewayLoggingRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s EnableGatewayLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableGatewayLoggingRequest) GoString() string {
	return s.String()
}

func (s *EnableGatewayLoggingRequest) SetGatewayId(v string) *EnableGatewayLoggingRequest {
	s.GatewayId = &v
	return s
}

func (s *EnableGatewayLoggingRequest) SetSecurityToken(v string) *EnableGatewayLoggingRequest {
	s.SecurityToken = &v
	return s
}

type EnableGatewayLoggingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableGatewayLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableGatewayLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *EnableGatewayLoggingResponseBody) SetCode(v string) *EnableGatewayLoggingResponseBody {
	s.Code = &v
	return s
}

func (s *EnableGatewayLoggingResponseBody) SetMessage(v string) *EnableGatewayLoggingResponseBody {
	s.Message = &v
	return s
}

func (s *EnableGatewayLoggingResponseBody) SetRequestId(v string) *EnableGatewayLoggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableGatewayLoggingResponseBody) SetSuccess(v bool) *EnableGatewayLoggingResponseBody {
	s.Success = &v
	return s
}

type EnableGatewayLoggingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableGatewayLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableGatewayLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableGatewayLoggingResponse) GoString() string {
	return s.String()
}

func (s *EnableGatewayLoggingResponse) SetHeaders(v map[string]*string) *EnableGatewayLoggingResponse {
	s.Headers = v
	return s
}

func (s *EnableGatewayLoggingResponse) SetStatusCode(v int32) *EnableGatewayLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableGatewayLoggingResponse) SetBody(v *EnableGatewayLoggingResponseBody) *EnableGatewayLoggingResponse {
	s.Body = v
	return s
}

type ExpandCacheDiskRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	LocalFilePath *string `json:"LocalFilePath,omitempty" xml:"LocalFilePath,omitempty"`
	NewSizeInGB   *int32  `json:"NewSizeInGB,omitempty" xml:"NewSizeInGB,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ExpandCacheDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpandCacheDiskRequest) GoString() string {
	return s.String()
}

func (s *ExpandCacheDiskRequest) SetGatewayId(v string) *ExpandCacheDiskRequest {
	s.GatewayId = &v
	return s
}

func (s *ExpandCacheDiskRequest) SetLocalFilePath(v string) *ExpandCacheDiskRequest {
	s.LocalFilePath = &v
	return s
}

func (s *ExpandCacheDiskRequest) SetNewSizeInGB(v int32) *ExpandCacheDiskRequest {
	s.NewSizeInGB = &v
	return s
}

func (s *ExpandCacheDiskRequest) SetSecurityToken(v string) *ExpandCacheDiskRequest {
	s.SecurityToken = &v
	return s
}

type ExpandCacheDiskResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExpandCacheDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExpandCacheDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ExpandCacheDiskResponseBody) SetBuyURL(v string) *ExpandCacheDiskResponseBody {
	s.BuyURL = &v
	return s
}

func (s *ExpandCacheDiskResponseBody) SetCode(v string) *ExpandCacheDiskResponseBody {
	s.Code = &v
	return s
}

func (s *ExpandCacheDiskResponseBody) SetMessage(v string) *ExpandCacheDiskResponseBody {
	s.Message = &v
	return s
}

func (s *ExpandCacheDiskResponseBody) SetRequestId(v string) *ExpandCacheDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExpandCacheDiskResponseBody) SetSuccess(v bool) *ExpandCacheDiskResponseBody {
	s.Success = &v
	return s
}

func (s *ExpandCacheDiskResponseBody) SetTaskId(v string) *ExpandCacheDiskResponseBody {
	s.TaskId = &v
	return s
}

type ExpandCacheDiskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExpandCacheDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExpandCacheDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpandCacheDiskResponse) GoString() string {
	return s.String()
}

func (s *ExpandCacheDiskResponse) SetHeaders(v map[string]*string) *ExpandCacheDiskResponse {
	s.Headers = v
	return s
}

func (s *ExpandCacheDiskResponse) SetStatusCode(v int32) *ExpandCacheDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpandCacheDiskResponse) SetBody(v *ExpandCacheDiskResponseBody) *ExpandCacheDiskResponse {
	s.Body = v
	return s
}

type ExpandGatewayFileShareRequest struct {
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId   *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s ExpandGatewayFileShareRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpandGatewayFileShareRequest) GoString() string {
	return s.String()
}

func (s *ExpandGatewayFileShareRequest) SetGatewayId(v string) *ExpandGatewayFileShareRequest {
	s.GatewayId = &v
	return s
}

func (s *ExpandGatewayFileShareRequest) SetIndexId(v string) *ExpandGatewayFileShareRequest {
	s.IndexId = &v
	return s
}

type ExpandGatewayFileShareResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExpandGatewayFileShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExpandGatewayFileShareResponseBody) GoString() string {
	return s.String()
}

func (s *ExpandGatewayFileShareResponseBody) SetCode(v string) *ExpandGatewayFileShareResponseBody {
	s.Code = &v
	return s
}

func (s *ExpandGatewayFileShareResponseBody) SetMessage(v string) *ExpandGatewayFileShareResponseBody {
	s.Message = &v
	return s
}

func (s *ExpandGatewayFileShareResponseBody) SetRequestId(v string) *ExpandGatewayFileShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExpandGatewayFileShareResponseBody) SetSuccess(v bool) *ExpandGatewayFileShareResponseBody {
	s.Success = &v
	return s
}

func (s *ExpandGatewayFileShareResponseBody) SetTaskId(v string) *ExpandGatewayFileShareResponseBody {
	s.TaskId = &v
	return s
}

type ExpandGatewayFileShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExpandGatewayFileShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExpandGatewayFileShareResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpandGatewayFileShareResponse) GoString() string {
	return s.String()
}

func (s *ExpandGatewayFileShareResponse) SetHeaders(v map[string]*string) *ExpandGatewayFileShareResponse {
	s.Headers = v
	return s
}

func (s *ExpandGatewayFileShareResponse) SetStatusCode(v int32) *ExpandGatewayFileShareResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpandGatewayFileShareResponse) SetBody(v *ExpandGatewayFileShareResponseBody) *ExpandGatewayFileShareResponse {
	s.Body = v
	return s
}

type ExpandGatewayNetworkBandwidthRequest struct {
	GatewayId           *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	NewNetworkBandwidth *int32  `json:"NewNetworkBandwidth,omitempty" xml:"NewNetworkBandwidth,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ExpandGatewayNetworkBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpandGatewayNetworkBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ExpandGatewayNetworkBandwidthRequest) SetGatewayId(v string) *ExpandGatewayNetworkBandwidthRequest {
	s.GatewayId = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthRequest) SetNewNetworkBandwidth(v int32) *ExpandGatewayNetworkBandwidthRequest {
	s.NewNetworkBandwidth = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthRequest) SetSecurityToken(v string) *ExpandGatewayNetworkBandwidthRequest {
	s.SecurityToken = &v
	return s
}

type ExpandGatewayNetworkBandwidthResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExpandGatewayNetworkBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExpandGatewayNetworkBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ExpandGatewayNetworkBandwidthResponseBody) SetBuyURL(v string) *ExpandGatewayNetworkBandwidthResponseBody {
	s.BuyURL = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponseBody) SetCode(v string) *ExpandGatewayNetworkBandwidthResponseBody {
	s.Code = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponseBody) SetMessage(v string) *ExpandGatewayNetworkBandwidthResponseBody {
	s.Message = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponseBody) SetRequestId(v string) *ExpandGatewayNetworkBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponseBody) SetSuccess(v bool) *ExpandGatewayNetworkBandwidthResponseBody {
	s.Success = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponseBody) SetTaskId(v string) *ExpandGatewayNetworkBandwidthResponseBody {
	s.TaskId = &v
	return s
}

type ExpandGatewayNetworkBandwidthResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExpandGatewayNetworkBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExpandGatewayNetworkBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpandGatewayNetworkBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ExpandGatewayNetworkBandwidthResponse) SetHeaders(v map[string]*string) *ExpandGatewayNetworkBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponse) SetStatusCode(v int32) *ExpandGatewayNetworkBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpandGatewayNetworkBandwidthResponse) SetBody(v *ExpandGatewayNetworkBandwidthResponseBody) *ExpandGatewayNetworkBandwidthResponse {
	s.Body = v
	return s
}

type GenerateGatewayTokenRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GenerateGatewayTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateGatewayTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateGatewayTokenRequest) SetGatewayId(v string) *GenerateGatewayTokenRequest {
	s.GatewayId = &v
	return s
}

func (s *GenerateGatewayTokenRequest) SetSecurityToken(v string) *GenerateGatewayTokenRequest {
	s.SecurityToken = &v
	return s
}

type GenerateGatewayTokenResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GenerateGatewayTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateGatewayTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateGatewayTokenResponseBody) SetCode(v string) *GenerateGatewayTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateGatewayTokenResponseBody) SetMessage(v string) *GenerateGatewayTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateGatewayTokenResponseBody) SetRequestId(v string) *GenerateGatewayTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateGatewayTokenResponseBody) SetSuccess(v bool) *GenerateGatewayTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateGatewayTokenResponseBody) SetToken(v string) *GenerateGatewayTokenResponseBody {
	s.Token = &v
	return s
}

type GenerateGatewayTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateGatewayTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateGatewayTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateGatewayTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateGatewayTokenResponse) SetHeaders(v map[string]*string) *GenerateGatewayTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateGatewayTokenResponse) SetStatusCode(v int32) *GenerateGatewayTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateGatewayTokenResponse) SetBody(v *GenerateGatewayTokenResponseBody) *GenerateGatewayTokenResponse {
	s.Body = v
	return s
}

type GenerateMqttTokenRequest struct {
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GenerateMqttTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateMqttTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateMqttTokenRequest) SetClientUUID(v string) *GenerateMqttTokenRequest {
	s.ClientUUID = &v
	return s
}

func (s *GenerateMqttTokenRequest) SetGatewayId(v string) *GenerateMqttTokenRequest {
	s.GatewayId = &v
	return s
}

func (s *GenerateMqttTokenRequest) SetSecurityToken(v string) *GenerateMqttTokenRequest {
	s.SecurityToken = &v
	return s
}

type GenerateMqttTokenResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MqttToken *string `json:"MqttToken,omitempty" xml:"MqttToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateMqttTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateMqttTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMqttTokenResponseBody) SetCode(v string) *GenerateMqttTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateMqttTokenResponseBody) SetMessage(v string) *GenerateMqttTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateMqttTokenResponseBody) SetMqttToken(v string) *GenerateMqttTokenResponseBody {
	s.MqttToken = &v
	return s
}

func (s *GenerateMqttTokenResponseBody) SetRequestId(v string) *GenerateMqttTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMqttTokenResponseBody) SetSuccess(v bool) *GenerateMqttTokenResponseBody {
	s.Success = &v
	return s
}

type GenerateMqttTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateMqttTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateMqttTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateMqttTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateMqttTokenResponse) SetHeaders(v map[string]*string) *GenerateMqttTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateMqttTokenResponse) SetStatusCode(v int32) *GenerateMqttTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateMqttTokenResponse) SetBody(v *GenerateMqttTokenResponseBody) *GenerateMqttTokenResponse {
	s.Body = v
	return s
}

type GenerateStsTokenRequest struct {
	ExpireInSeconds *int64  `json:"ExpireInSeconds,omitempty" xml:"ExpireInSeconds,omitempty"`
	GatewayId       *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TokenType       *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s GenerateStsTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateStsTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateStsTokenRequest) SetExpireInSeconds(v int64) *GenerateStsTokenRequest {
	s.ExpireInSeconds = &v
	return s
}

func (s *GenerateStsTokenRequest) SetGatewayId(v string) *GenerateStsTokenRequest {
	s.GatewayId = &v
	return s
}

func (s *GenerateStsTokenRequest) SetSecurityToken(v string) *GenerateStsTokenRequest {
	s.SecurityToken = &v
	return s
}

func (s *GenerateStsTokenRequest) SetTokenType(v string) *GenerateStsTokenRequest {
	s.TokenType = &v
	return s
}

type GenerateStsTokenResponseBody struct {
	AccessKeyId         *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret     *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Code                *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Environment         *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	Expiration          *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Success             *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SupportBundleTarget *string `json:"SupportBundleTarget,omitempty" xml:"SupportBundleTarget,omitempty"`
}

func (s GenerateStsTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateStsTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateStsTokenResponseBody) SetAccessKeyId(v string) *GenerateStsTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetAccessKeySecret(v string) *GenerateStsTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetCode(v string) *GenerateStsTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetEnvironment(v string) *GenerateStsTokenResponseBody {
	s.Environment = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetExpiration(v string) *GenerateStsTokenResponseBody {
	s.Expiration = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetMessage(v string) *GenerateStsTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetRequestId(v string) *GenerateStsTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetSecurityToken(v string) *GenerateStsTokenResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetSuccess(v bool) *GenerateStsTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateStsTokenResponseBody) SetSupportBundleTarget(v string) *GenerateStsTokenResponseBody {
	s.SupportBundleTarget = &v
	return s
}

type GenerateStsTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateStsTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateStsTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateStsTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateStsTokenResponse) SetHeaders(v map[string]*string) *GenerateStsTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateStsTokenResponse) SetStatusCode(v int32) *GenerateStsTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateStsTokenResponse) SetBody(v *GenerateStsTokenResponseBody) *GenerateStsTokenResponse {
	s.Body = v
	return s
}

type HandleGatewayAutoPlanRequest struct {
	Cancel        *bool   `json:"Cancel,omitempty" xml:"Cancel,omitempty"`
	DelayHours    *int32  `json:"DelayHours,omitempty" xml:"DelayHours,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	PlanId        *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s HandleGatewayAutoPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleGatewayAutoPlanRequest) GoString() string {
	return s.String()
}

func (s *HandleGatewayAutoPlanRequest) SetCancel(v bool) *HandleGatewayAutoPlanRequest {
	s.Cancel = &v
	return s
}

func (s *HandleGatewayAutoPlanRequest) SetDelayHours(v int32) *HandleGatewayAutoPlanRequest {
	s.DelayHours = &v
	return s
}

func (s *HandleGatewayAutoPlanRequest) SetGatewayId(v string) *HandleGatewayAutoPlanRequest {
	s.GatewayId = &v
	return s
}

func (s *HandleGatewayAutoPlanRequest) SetPlanId(v string) *HandleGatewayAutoPlanRequest {
	s.PlanId = &v
	return s
}

func (s *HandleGatewayAutoPlanRequest) SetSecurityToken(v string) *HandleGatewayAutoPlanRequest {
	s.SecurityToken = &v
	return s
}

type HandleGatewayAutoPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HandleGatewayAutoPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleGatewayAutoPlanResponseBody) GoString() string {
	return s.String()
}

func (s *HandleGatewayAutoPlanResponseBody) SetCode(v string) *HandleGatewayAutoPlanResponseBody {
	s.Code = &v
	return s
}

func (s *HandleGatewayAutoPlanResponseBody) SetMessage(v string) *HandleGatewayAutoPlanResponseBody {
	s.Message = &v
	return s
}

func (s *HandleGatewayAutoPlanResponseBody) SetRequestId(v string) *HandleGatewayAutoPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleGatewayAutoPlanResponseBody) SetSuccess(v bool) *HandleGatewayAutoPlanResponseBody {
	s.Success = &v
	return s
}

type HandleGatewayAutoPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HandleGatewayAutoPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandleGatewayAutoPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleGatewayAutoPlanResponse) GoString() string {
	return s.String()
}

func (s *HandleGatewayAutoPlanResponse) SetHeaders(v map[string]*string) *HandleGatewayAutoPlanResponse {
	s.Headers = v
	return s
}

func (s *HandleGatewayAutoPlanResponse) SetStatusCode(v int32) *HandleGatewayAutoPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleGatewayAutoPlanResponse) SetBody(v *HandleGatewayAutoPlanResponseBody) *HandleGatewayAutoPlanResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken        *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId         *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId       []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceRegionId *string                       `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	ResourceType     *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken    *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag              []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceRegionId(v string) *ListTagResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetSecurityToken(v string) *ListTagResourcesRequest {
	s.SecurityToken = &v
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyGatewayRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayRequest) SetDescription(v string) *ModifyGatewayRequest {
	s.Description = &v
	return s
}

func (s *ModifyGatewayRequest) SetGatewayId(v string) *ModifyGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayRequest) SetName(v string) *ModifyGatewayRequest {
	s.Name = &v
	return s
}

func (s *ModifyGatewayRequest) SetSecurityToken(v string) *ModifyGatewayRequest {
	s.SecurityToken = &v
	return s
}

type ModifyGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayResponseBody) SetCode(v string) *ModifyGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetGatewayId(v string) *ModifyGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetMessage(v string) *ModifyGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetRequestId(v string) *ModifyGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayResponseBody) SetSuccess(v bool) *ModifyGatewayResponseBody {
	s.Success = &v
	return s
}

type ModifyGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayResponse) SetHeaders(v map[string]*string) *ModifyGatewayResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayResponse) SetStatusCode(v int32) *ModifyGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayResponse) SetBody(v *ModifyGatewayResponseBody) *ModifyGatewayResponse {
	s.Body = v
	return s
}

type ModifyGatewayBlockVolumeRequest struct {
	CacheConfig   *string `json:"CacheConfig,omitempty" xml:"CacheConfig,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGatewayBlockVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayBlockVolumeRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayBlockVolumeRequest) SetCacheConfig(v string) *ModifyGatewayBlockVolumeRequest {
	s.CacheConfig = &v
	return s
}

func (s *ModifyGatewayBlockVolumeRequest) SetGatewayId(v string) *ModifyGatewayBlockVolumeRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayBlockVolumeRequest) SetIndexId(v string) *ModifyGatewayBlockVolumeRequest {
	s.IndexId = &v
	return s
}

func (s *ModifyGatewayBlockVolumeRequest) SetSecurityToken(v string) *ModifyGatewayBlockVolumeRequest {
	s.SecurityToken = &v
	return s
}

type ModifyGatewayBlockVolumeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyGatewayBlockVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayBlockVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayBlockVolumeResponseBody) SetCode(v string) *ModifyGatewayBlockVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayBlockVolumeResponseBody) SetMessage(v string) *ModifyGatewayBlockVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGatewayBlockVolumeResponseBody) SetRequestId(v string) *ModifyGatewayBlockVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayBlockVolumeResponseBody) SetSuccess(v bool) *ModifyGatewayBlockVolumeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGatewayBlockVolumeResponseBody) SetTaskId(v string) *ModifyGatewayBlockVolumeResponseBody {
	s.TaskId = &v
	return s
}

type ModifyGatewayBlockVolumeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyGatewayBlockVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGatewayBlockVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayBlockVolumeResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayBlockVolumeResponse) SetHeaders(v map[string]*string) *ModifyGatewayBlockVolumeResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayBlockVolumeResponse) SetStatusCode(v int32) *ModifyGatewayBlockVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayBlockVolumeResponse) SetBody(v *ModifyGatewayBlockVolumeResponseBody) *ModifyGatewayBlockVolumeResponse {
	s.Body = v
	return s
}

type ModifyGatewayClassRequest struct {
	GatewayClass  *string `json:"GatewayClass,omitempty" xml:"GatewayClass,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGatewayClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayClassRequest) SetGatewayClass(v string) *ModifyGatewayClassRequest {
	s.GatewayClass = &v
	return s
}

func (s *ModifyGatewayClassRequest) SetGatewayId(v string) *ModifyGatewayClassRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayClassRequest) SetSecurityToken(v string) *ModifyGatewayClassRequest {
	s.SecurityToken = &v
	return s
}

type ModifyGatewayClassResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyGatewayClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayClassResponseBody) SetBuyURL(v string) *ModifyGatewayClassResponseBody {
	s.BuyURL = &v
	return s
}

func (s *ModifyGatewayClassResponseBody) SetCode(v string) *ModifyGatewayClassResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayClassResponseBody) SetMessage(v string) *ModifyGatewayClassResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGatewayClassResponseBody) SetRequestId(v string) *ModifyGatewayClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayClassResponseBody) SetSuccess(v bool) *ModifyGatewayClassResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGatewayClassResponseBody) SetTaskId(v string) *ModifyGatewayClassResponseBody {
	s.TaskId = &v
	return s
}

type ModifyGatewayClassResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyGatewayClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGatewayClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayClassResponse) SetHeaders(v map[string]*string) *ModifyGatewayClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayClassResponse) SetStatusCode(v int32) *ModifyGatewayClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayClassResponse) SetBody(v *ModifyGatewayClassResponseBody) *ModifyGatewayClassResponse {
	s.Body = v
	return s
}

type ModifyGatewayFileShareRequest struct {
	CacheConfig   *string `json:"CacheConfig,omitempty" xml:"CacheConfig,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGatewayFileShareRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayFileShareRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayFileShareRequest) SetCacheConfig(v string) *ModifyGatewayFileShareRequest {
	s.CacheConfig = &v
	return s
}

func (s *ModifyGatewayFileShareRequest) SetGatewayId(v string) *ModifyGatewayFileShareRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayFileShareRequest) SetIndexId(v string) *ModifyGatewayFileShareRequest {
	s.IndexId = &v
	return s
}

func (s *ModifyGatewayFileShareRequest) SetSecurityToken(v string) *ModifyGatewayFileShareRequest {
	s.SecurityToken = &v
	return s
}

type ModifyGatewayFileShareResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyGatewayFileShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayFileShareResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayFileShareResponseBody) SetCode(v string) *ModifyGatewayFileShareResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayFileShareResponseBody) SetMessage(v string) *ModifyGatewayFileShareResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGatewayFileShareResponseBody) SetRequestId(v string) *ModifyGatewayFileShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayFileShareResponseBody) SetSuccess(v bool) *ModifyGatewayFileShareResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGatewayFileShareResponseBody) SetTaskId(v string) *ModifyGatewayFileShareResponseBody {
	s.TaskId = &v
	return s
}

type ModifyGatewayFileShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyGatewayFileShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGatewayFileShareResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayFileShareResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayFileShareResponse) SetHeaders(v map[string]*string) *ModifyGatewayFileShareResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayFileShareResponse) SetStatusCode(v int32) *ModifyGatewayFileShareResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayFileShareResponse) SetBody(v *ModifyGatewayFileShareResponseBody) *ModifyGatewayFileShareResponse {
	s.Body = v
	return s
}

type ModifyGatewayFileShareWatermarkRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Watermark     *int32  `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s ModifyGatewayFileShareWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayFileShareWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyGatewayFileShareWatermarkRequest) SetGatewayId(v string) *ModifyGatewayFileShareWatermarkRequest {
	s.GatewayId = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkRequest) SetIndexId(v string) *ModifyGatewayFileShareWatermarkRequest {
	s.IndexId = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkRequest) SetSecurityToken(v string) *ModifyGatewayFileShareWatermarkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkRequest) SetWatermark(v int32) *ModifyGatewayFileShareWatermarkRequest {
	s.Watermark = &v
	return s
}

type ModifyGatewayFileShareWatermarkResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyGatewayFileShareWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayFileShareWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGatewayFileShareWatermarkResponseBody) SetCode(v string) *ModifyGatewayFileShareWatermarkResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkResponseBody) SetMessage(v string) *ModifyGatewayFileShareWatermarkResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkResponseBody) SetRequestId(v string) *ModifyGatewayFileShareWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkResponseBody) SetSuccess(v bool) *ModifyGatewayFileShareWatermarkResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkResponseBody) SetTaskId(v string) *ModifyGatewayFileShareWatermarkResponseBody {
	s.TaskId = &v
	return s
}

type ModifyGatewayFileShareWatermarkResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyGatewayFileShareWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGatewayFileShareWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGatewayFileShareWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyGatewayFileShareWatermarkResponse) SetHeaders(v map[string]*string) *ModifyGatewayFileShareWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyGatewayFileShareWatermarkResponse) SetStatusCode(v int32) *ModifyGatewayFileShareWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGatewayFileShareWatermarkResponse) SetBody(v *ModifyGatewayFileShareWatermarkResponseBody) *ModifyGatewayFileShareWatermarkResponse {
	s.Body = v
	return s
}

type ModifyStorageBundleRequest struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
}

func (s ModifyStorageBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyStorageBundleRequest) GoString() string {
	return s.String()
}

func (s *ModifyStorageBundleRequest) SetDescription(v string) *ModifyStorageBundleRequest {
	s.Description = &v
	return s
}

func (s *ModifyStorageBundleRequest) SetName(v string) *ModifyStorageBundleRequest {
	s.Name = &v
	return s
}

func (s *ModifyStorageBundleRequest) SetSecurityToken(v string) *ModifyStorageBundleRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyStorageBundleRequest) SetStorageBundleId(v string) *ModifyStorageBundleRequest {
	s.StorageBundleId = &v
	return s
}

type ModifyStorageBundleResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyStorageBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyStorageBundleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStorageBundleResponseBody) SetCode(v string) *ModifyStorageBundleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyStorageBundleResponseBody) SetMessage(v string) *ModifyStorageBundleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyStorageBundleResponseBody) SetRequestId(v string) *ModifyStorageBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStorageBundleResponseBody) SetStorageBundleId(v string) *ModifyStorageBundleResponseBody {
	s.StorageBundleId = &v
	return s
}

func (s *ModifyStorageBundleResponseBody) SetSuccess(v bool) *ModifyStorageBundleResponseBody {
	s.Success = &v
	return s
}

type ModifyStorageBundleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyStorageBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyStorageBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyStorageBundleResponse) GoString() string {
	return s.String()
}

func (s *ModifyStorageBundleResponse) SetHeaders(v map[string]*string) *ModifyStorageBundleResponse {
	s.Headers = v
	return s
}

func (s *ModifyStorageBundleResponse) SetStatusCode(v int32) *ModifyStorageBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStorageBundleResponse) SetBody(v *ModifyStorageBundleResponseBody) *ModifyStorageBundleResponse {
	s.Body = v
	return s
}

type OpenSgwServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenSgwServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenSgwServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSgwServiceResponseBody) SetOrderId(v string) *OpenSgwServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenSgwServiceResponseBody) SetRequestId(v string) *OpenSgwServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenSgwServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenSgwServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenSgwServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenSgwServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenSgwServiceResponse) SetHeaders(v map[string]*string) *OpenSgwServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenSgwServiceResponse) SetStatusCode(v int32) *OpenSgwServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSgwServiceResponse) SetBody(v *OpenSgwServiceResponseBody) *OpenSgwServiceResponse {
	s.Body = v
	return s
}

type OperateGatewayRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	OperateAction *string `json:"OperateAction,omitempty" xml:"OperateAction,omitempty"`
	Params        *string `json:"Params,omitempty" xml:"Params,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OperateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateGatewayRequest) GoString() string {
	return s.String()
}

func (s *OperateGatewayRequest) SetGatewayId(v string) *OperateGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *OperateGatewayRequest) SetOperateAction(v string) *OperateGatewayRequest {
	s.OperateAction = &v
	return s
}

func (s *OperateGatewayRequest) SetParams(v string) *OperateGatewayRequest {
	s.Params = &v
	return s
}

func (s *OperateGatewayRequest) SetSecurityToken(v string) *OperateGatewayRequest {
	s.SecurityToken = &v
	return s
}

type OperateGatewayResponseBody struct {
	BuyURL    *string `json:"BuyURL,omitempty" xml:"BuyURL,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *OperateGatewayResponseBody) SetBuyURL(v string) *OperateGatewayResponseBody {
	s.BuyURL = &v
	return s
}

func (s *OperateGatewayResponseBody) SetCode(v string) *OperateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *OperateGatewayResponseBody) SetMessage(v string) *OperateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *OperateGatewayResponseBody) SetRequestId(v string) *OperateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateGatewayResponseBody) SetSuccess(v bool) *OperateGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *OperateGatewayResponseBody) SetTaskId(v string) *OperateGatewayResponseBody {
	s.TaskId = &v
	return s
}

type OperateGatewayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OperateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateGatewayResponse) GoString() string {
	return s.String()
}

func (s *OperateGatewayResponse) SetHeaders(v map[string]*string) *OperateGatewayResponse {
	s.Headers = v
	return s
}

func (s *OperateGatewayResponse) SetStatusCode(v int32) *OperateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateGatewayResponse) SetBody(v *OperateGatewayResponseBody) *OperateGatewayResponse {
	s.Body = v
	return s
}

type ReleaseServiceRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReleaseServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseServiceRequest) SetSecurityToken(v string) *ReleaseServiceRequest {
	s.SecurityToken = &v
	return s
}

type ReleaseServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponseBody) SetCode(v string) *ReleaseServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseServiceResponseBody) SetMessage(v string) *ReleaseServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseServiceResponseBody) SetRequestId(v string) *ReleaseServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseServiceResponseBody) SetSuccess(v bool) *ReleaseServiceResponseBody {
	s.Success = &v
	return s
}

type ReleaseServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponse) SetHeaders(v map[string]*string) *ReleaseServiceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseServiceResponse) SetStatusCode(v int32) *ReleaseServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseServiceResponse) SetBody(v *ReleaseServiceResponseBody) *ReleaseServiceResponse {
	s.Body = v
	return s
}

type RemoveSharesFromExpressSyncRequest struct {
	ExpressSyncId *string `json:"ExpressSyncId,omitempty" xml:"ExpressSyncId,omitempty"`
	GatewayShares *string `json:"GatewayShares,omitempty" xml:"GatewayShares,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveSharesFromExpressSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSharesFromExpressSyncRequest) GoString() string {
	return s.String()
}

func (s *RemoveSharesFromExpressSyncRequest) SetExpressSyncId(v string) *RemoveSharesFromExpressSyncRequest {
	s.ExpressSyncId = &v
	return s
}

func (s *RemoveSharesFromExpressSyncRequest) SetGatewayShares(v string) *RemoveSharesFromExpressSyncRequest {
	s.GatewayShares = &v
	return s
}

func (s *RemoveSharesFromExpressSyncRequest) SetSecurityToken(v string) *RemoveSharesFromExpressSyncRequest {
	s.SecurityToken = &v
	return s
}

type RemoveSharesFromExpressSyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RemoveSharesFromExpressSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSharesFromExpressSyncResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSharesFromExpressSyncResponseBody) SetCode(v string) *RemoveSharesFromExpressSyncResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSharesFromExpressSyncResponseBody) SetMessage(v string) *RemoveSharesFromExpressSyncResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSharesFromExpressSyncResponseBody) SetRequestId(v string) *RemoveSharesFromExpressSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSharesFromExpressSyncResponseBody) SetSuccess(v bool) *RemoveSharesFromExpressSyncResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveSharesFromExpressSyncResponseBody) SetTaskId(v string) *RemoveSharesFromExpressSyncResponseBody {
	s.TaskId = &v
	return s
}

type RemoveSharesFromExpressSyncResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveSharesFromExpressSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSharesFromExpressSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSharesFromExpressSyncResponse) GoString() string {
	return s.String()
}

func (s *RemoveSharesFromExpressSyncResponse) SetHeaders(v map[string]*string) *RemoveSharesFromExpressSyncResponse {
	s.Headers = v
	return s
}

func (s *RemoveSharesFromExpressSyncResponse) SetStatusCode(v int32) *RemoveSharesFromExpressSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSharesFromExpressSyncResponse) SetBody(v *RemoveSharesFromExpressSyncResponseBody) *RemoveSharesFromExpressSyncResponse {
	s.Body = v
	return s
}

type RemoveTagsFromGatewayRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tags          *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s RemoveTagsFromGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsFromGatewayRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromGatewayRequest) SetGatewayId(v string) *RemoveTagsFromGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *RemoveTagsFromGatewayRequest) SetSecurityToken(v string) *RemoveTagsFromGatewayRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveTagsFromGatewayRequest) SetTags(v string) *RemoveTagsFromGatewayRequest {
	s.Tags = &v
	return s
}

type RemoveTagsFromGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveTagsFromGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsFromGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromGatewayResponseBody) SetCode(v string) *RemoveTagsFromGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveTagsFromGatewayResponseBody) SetMessage(v string) *RemoveTagsFromGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveTagsFromGatewayResponseBody) SetRequestId(v string) *RemoveTagsFromGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTagsFromGatewayResponseBody) SetSuccess(v bool) *RemoveTagsFromGatewayResponseBody {
	s.Success = &v
	return s
}

type RemoveTagsFromGatewayResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveTagsFromGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTagsFromGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsFromGatewayResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromGatewayResponse) SetHeaders(v map[string]*string) *RemoveTagsFromGatewayResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsFromGatewayResponse) SetStatusCode(v int32) *RemoveTagsFromGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagsFromGatewayResponse) SetBody(v *RemoveTagsFromGatewayResponseBody) *RemoveTagsFromGatewayResponse {
	s.Body = v
	return s
}

type ReportBlockVolumesRequest struct {
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Info          *string `json:"Info,omitempty" xml:"Info,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReportBlockVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportBlockVolumesRequest) GoString() string {
	return s.String()
}

func (s *ReportBlockVolumesRequest) SetClientUUID(v string) *ReportBlockVolumesRequest {
	s.ClientUUID = &v
	return s
}

func (s *ReportBlockVolumesRequest) SetGatewayId(v string) *ReportBlockVolumesRequest {
	s.GatewayId = &v
	return s
}

func (s *ReportBlockVolumesRequest) SetInfo(v string) *ReportBlockVolumesRequest {
	s.Info = &v
	return s
}

func (s *ReportBlockVolumesRequest) SetSecurityToken(v string) *ReportBlockVolumesRequest {
	s.SecurityToken = &v
	return s
}

type ReportBlockVolumesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportBlockVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportBlockVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ReportBlockVolumesResponseBody) SetCode(v string) *ReportBlockVolumesResponseBody {
	s.Code = &v
	return s
}

func (s *ReportBlockVolumesResponseBody) SetMessage(v string) *ReportBlockVolumesResponseBody {
	s.Message = &v
	return s
}

func (s *ReportBlockVolumesResponseBody) SetRequestId(v string) *ReportBlockVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportBlockVolumesResponseBody) SetSuccess(v bool) *ReportBlockVolumesResponseBody {
	s.Success = &v
	return s
}

type ReportBlockVolumesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportBlockVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportBlockVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportBlockVolumesResponse) GoString() string {
	return s.String()
}

func (s *ReportBlockVolumesResponse) SetHeaders(v map[string]*string) *ReportBlockVolumesResponse {
	s.Headers = v
	return s
}

func (s *ReportBlockVolumesResponse) SetStatusCode(v int32) *ReportBlockVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportBlockVolumesResponse) SetBody(v *ReportBlockVolumesResponseBody) *ReportBlockVolumesResponse {
	s.Body = v
	return s
}

type ReportFileSharesRequest struct {
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Info          *string `json:"Info,omitempty" xml:"Info,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReportFileSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportFileSharesRequest) GoString() string {
	return s.String()
}

func (s *ReportFileSharesRequest) SetClientUUID(v string) *ReportFileSharesRequest {
	s.ClientUUID = &v
	return s
}

func (s *ReportFileSharesRequest) SetGatewayId(v string) *ReportFileSharesRequest {
	s.GatewayId = &v
	return s
}

func (s *ReportFileSharesRequest) SetInfo(v string) *ReportFileSharesRequest {
	s.Info = &v
	return s
}

func (s *ReportFileSharesRequest) SetSecurityToken(v string) *ReportFileSharesRequest {
	s.SecurityToken = &v
	return s
}

type ReportFileSharesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportFileSharesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportFileSharesResponseBody) GoString() string {
	return s.String()
}

func (s *ReportFileSharesResponseBody) SetCode(v string) *ReportFileSharesResponseBody {
	s.Code = &v
	return s
}

func (s *ReportFileSharesResponseBody) SetMessage(v string) *ReportFileSharesResponseBody {
	s.Message = &v
	return s
}

func (s *ReportFileSharesResponseBody) SetRequestId(v string) *ReportFileSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportFileSharesResponseBody) SetSuccess(v bool) *ReportFileSharesResponseBody {
	s.Success = &v
	return s
}

type ReportFileSharesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportFileSharesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportFileSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportFileSharesResponse) GoString() string {
	return s.String()
}

func (s *ReportFileSharesResponse) SetHeaders(v map[string]*string) *ReportFileSharesResponse {
	s.Headers = v
	return s
}

func (s *ReportFileSharesResponse) SetStatusCode(v int32) *ReportFileSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportFileSharesResponse) SetBody(v *ReportFileSharesResponseBody) *ReportFileSharesResponse {
	s.Body = v
	return s
}

type ReportGatewayInfoRequest struct {
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayStatus *string `json:"GatewayStatus,omitempty" xml:"GatewayStatus,omitempty"`
	Info          *string `json:"Info,omitempty" xml:"Info,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Time          *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ReportGatewayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportGatewayInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportGatewayInfoRequest) SetClientUUID(v string) *ReportGatewayInfoRequest {
	s.ClientUUID = &v
	return s
}

func (s *ReportGatewayInfoRequest) SetGatewayId(v string) *ReportGatewayInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *ReportGatewayInfoRequest) SetGatewayStatus(v string) *ReportGatewayInfoRequest {
	s.GatewayStatus = &v
	return s
}

func (s *ReportGatewayInfoRequest) SetInfo(v string) *ReportGatewayInfoRequest {
	s.Info = &v
	return s
}

func (s *ReportGatewayInfoRequest) SetSecurityToken(v string) *ReportGatewayInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReportGatewayInfoRequest) SetTime(v int64) *ReportGatewayInfoRequest {
	s.Time = &v
	return s
}

type ReportGatewayInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportGatewayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportGatewayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportGatewayInfoResponseBody) SetCode(v string) *ReportGatewayInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportGatewayInfoResponseBody) SetMessage(v string) *ReportGatewayInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportGatewayInfoResponseBody) SetRequestId(v string) *ReportGatewayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportGatewayInfoResponseBody) SetSuccess(v bool) *ReportGatewayInfoResponseBody {
	s.Success = &v
	return s
}

type ReportGatewayInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportGatewayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportGatewayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportGatewayInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportGatewayInfoResponse) SetHeaders(v map[string]*string) *ReportGatewayInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportGatewayInfoResponse) SetStatusCode(v int32) *ReportGatewayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportGatewayInfoResponse) SetBody(v *ReportGatewayInfoResponseBody) *ReportGatewayInfoResponse {
	s.Body = v
	return s
}

type ReportGatewayUsageRequest struct {
	ClientUUID    *string `json:"ClientUUID,omitempty" xml:"ClientUUID,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Usage         *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ReportGatewayUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportGatewayUsageRequest) GoString() string {
	return s.String()
}

func (s *ReportGatewayUsageRequest) SetClientUUID(v string) *ReportGatewayUsageRequest {
	s.ClientUUID = &v
	return s
}

func (s *ReportGatewayUsageRequest) SetGatewayId(v string) *ReportGatewayUsageRequest {
	s.GatewayId = &v
	return s
}

func (s *ReportGatewayUsageRequest) SetSecurityToken(v string) *ReportGatewayUsageRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReportGatewayUsageRequest) SetUsage(v string) *ReportGatewayUsageRequest {
	s.Usage = &v
	return s
}

type ReportGatewayUsageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportGatewayUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportGatewayUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ReportGatewayUsageResponseBody) SetCode(v string) *ReportGatewayUsageResponseBody {
	s.Code = &v
	return s
}

func (s *ReportGatewayUsageResponseBody) SetMessage(v string) *ReportGatewayUsageResponseBody {
	s.Message = &v
	return s
}

func (s *ReportGatewayUsageResponseBody) SetRequestId(v string) *ReportGatewayUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportGatewayUsageResponseBody) SetSuccess(v bool) *ReportGatewayUsageResponseBody {
	s.Success = &v
	return s
}

type ReportGatewayUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportGatewayUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportGatewayUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportGatewayUsageResponse) GoString() string {
	return s.String()
}

func (s *ReportGatewayUsageResponse) SetHeaders(v map[string]*string) *ReportGatewayUsageResponse {
	s.Headers = v
	return s
}

func (s *ReportGatewayUsageResponse) SetStatusCode(v int32) *ReportGatewayUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportGatewayUsageResponse) SetBody(v *ReportGatewayUsageResponseBody) *ReportGatewayUsageResponse {
	s.Body = v
	return s
}

type ResetGatewayPasswordRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ResetGatewayPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetGatewayPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetGatewayPasswordRequest) SetGatewayId(v string) *ResetGatewayPasswordRequest {
	s.GatewayId = &v
	return s
}

func (s *ResetGatewayPasswordRequest) SetPassword(v string) *ResetGatewayPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetGatewayPasswordRequest) SetSecurityToken(v string) *ResetGatewayPasswordRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetGatewayPasswordRequest) SetUsername(v string) *ResetGatewayPasswordRequest {
	s.Username = &v
	return s
}

type ResetGatewayPasswordResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetGatewayPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetGatewayPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetGatewayPasswordResponseBody) SetCode(v string) *ResetGatewayPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetGatewayPasswordResponseBody) SetMessage(v string) *ResetGatewayPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetGatewayPasswordResponseBody) SetRequestId(v string) *ResetGatewayPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetGatewayPasswordResponseBody) SetSuccess(v bool) *ResetGatewayPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ResetGatewayPasswordResponseBody) SetTaskId(v string) *ResetGatewayPasswordResponseBody {
	s.TaskId = &v
	return s
}

type ResetGatewayPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetGatewayPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetGatewayPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetGatewayPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetGatewayPasswordResponse) SetHeaders(v map[string]*string) *ResetGatewayPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetGatewayPasswordResponse) SetStatusCode(v int32) *ResetGatewayPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetGatewayPasswordResponse) SetBody(v *ResetGatewayPasswordResponseBody) *ResetGatewayPasswordResponse {
	s.Body = v
	return s
}

type RestartFileSharesRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ShareProtocol *string `json:"ShareProtocol,omitempty" xml:"ShareProtocol,omitempty"`
}

func (s RestartFileSharesRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartFileSharesRequest) GoString() string {
	return s.String()
}

func (s *RestartFileSharesRequest) SetGatewayId(v string) *RestartFileSharesRequest {
	s.GatewayId = &v
	return s
}

func (s *RestartFileSharesRequest) SetSecurityToken(v string) *RestartFileSharesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartFileSharesRequest) SetShareProtocol(v string) *RestartFileSharesRequest {
	s.ShareProtocol = &v
	return s
}

type RestartFileSharesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartFileSharesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartFileSharesResponseBody) GoString() string {
	return s.String()
}

func (s *RestartFileSharesResponseBody) SetCode(v string) *RestartFileSharesResponseBody {
	s.Code = &v
	return s
}

func (s *RestartFileSharesResponseBody) SetMessage(v string) *RestartFileSharesResponseBody {
	s.Message = &v
	return s
}

func (s *RestartFileSharesResponseBody) SetRequestId(v string) *RestartFileSharesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartFileSharesResponseBody) SetSuccess(v bool) *RestartFileSharesResponseBody {
	s.Success = &v
	return s
}

func (s *RestartFileSharesResponseBody) SetTaskId(v string) *RestartFileSharesResponseBody {
	s.TaskId = &v
	return s
}

type RestartFileSharesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartFileSharesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartFileSharesResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartFileSharesResponse) GoString() string {
	return s.String()
}

func (s *RestartFileSharesResponse) SetHeaders(v map[string]*string) *RestartFileSharesResponse {
	s.Headers = v
	return s
}

func (s *RestartFileSharesResponse) SetStatusCode(v int32) *RestartFileSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartFileSharesResponse) SetBody(v *RestartFileSharesResponseBody) *RestartFileSharesResponse {
	s.Body = v
	return s
}

type SetGatewayADInfoRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IsEnabled     *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s SetGatewayADInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayADInfoRequest) GoString() string {
	return s.String()
}

func (s *SetGatewayADInfoRequest) SetGatewayId(v string) *SetGatewayADInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *SetGatewayADInfoRequest) SetIsEnabled(v bool) *SetGatewayADInfoRequest {
	s.IsEnabled = &v
	return s
}

func (s *SetGatewayADInfoRequest) SetPassword(v string) *SetGatewayADInfoRequest {
	s.Password = &v
	return s
}

func (s *SetGatewayADInfoRequest) SetSecurityToken(v string) *SetGatewayADInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetGatewayADInfoRequest) SetServerIp(v string) *SetGatewayADInfoRequest {
	s.ServerIp = &v
	return s
}

func (s *SetGatewayADInfoRequest) SetUsername(v string) *SetGatewayADInfoRequest {
	s.Username = &v
	return s
}

type SetGatewayADInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetGatewayADInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayADInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetGatewayADInfoResponseBody) SetCode(v string) *SetGatewayADInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SetGatewayADInfoResponseBody) SetMessage(v string) *SetGatewayADInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetGatewayADInfoResponseBody) SetRequestId(v string) *SetGatewayADInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGatewayADInfoResponseBody) SetSuccess(v bool) *SetGatewayADInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SetGatewayADInfoResponseBody) SetTaskId(v string) *SetGatewayADInfoResponseBody {
	s.TaskId = &v
	return s
}

type SetGatewayADInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGatewayADInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGatewayADInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayADInfoResponse) GoString() string {
	return s.String()
}

func (s *SetGatewayADInfoResponse) SetHeaders(v map[string]*string) *SetGatewayADInfoResponse {
	s.Headers = v
	return s
}

func (s *SetGatewayADInfoResponse) SetStatusCode(v int32) *SetGatewayADInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGatewayADInfoResponse) SetBody(v *SetGatewayADInfoResponseBody) *SetGatewayADInfoResponse {
	s.Body = v
	return s
}

type SetGatewayAutoUpgradeConfigurationRequest struct {
	AutoUpgradeEndHour   *int32  `json:"AutoUpgradeEndHour,omitempty" xml:"AutoUpgradeEndHour,omitempty"`
	AutoUpgradeStartHour *int32  `json:"AutoUpgradeStartHour,omitempty" xml:"AutoUpgradeStartHour,omitempty"`
	GatewayId            *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IsAutoUpgrade        *bool   `json:"IsAutoUpgrade,omitempty" xml:"IsAutoUpgrade,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetGatewayAutoUpgradeConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayAutoUpgradeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetGatewayAutoUpgradeConfigurationRequest) SetAutoUpgradeEndHour(v int32) *SetGatewayAutoUpgradeConfigurationRequest {
	s.AutoUpgradeEndHour = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationRequest) SetAutoUpgradeStartHour(v int32) *SetGatewayAutoUpgradeConfigurationRequest {
	s.AutoUpgradeStartHour = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationRequest) SetGatewayId(v string) *SetGatewayAutoUpgradeConfigurationRequest {
	s.GatewayId = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationRequest) SetIsAutoUpgrade(v bool) *SetGatewayAutoUpgradeConfigurationRequest {
	s.IsAutoUpgrade = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationRequest) SetSecurityToken(v string) *SetGatewayAutoUpgradeConfigurationRequest {
	s.SecurityToken = &v
	return s
}

type SetGatewayAutoUpgradeConfigurationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetGatewayAutoUpgradeConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayAutoUpgradeConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetGatewayAutoUpgradeConfigurationResponseBody) SetCode(v string) *SetGatewayAutoUpgradeConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationResponseBody) SetMessage(v string) *SetGatewayAutoUpgradeConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationResponseBody) SetRequestId(v string) *SetGatewayAutoUpgradeConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationResponseBody) SetSuccess(v bool) *SetGatewayAutoUpgradeConfigurationResponseBody {
	s.Success = &v
	return s
}

type SetGatewayAutoUpgradeConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGatewayAutoUpgradeConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGatewayAutoUpgradeConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayAutoUpgradeConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetGatewayAutoUpgradeConfigurationResponse) SetHeaders(v map[string]*string) *SetGatewayAutoUpgradeConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationResponse) SetStatusCode(v int32) *SetGatewayAutoUpgradeConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGatewayAutoUpgradeConfigurationResponse) SetBody(v *SetGatewayAutoUpgradeConfigurationResponseBody) *SetGatewayAutoUpgradeConfigurationResponse {
	s.Body = v
	return s
}

type SetGatewayDNSRequest struct {
	DnsServer     *string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetGatewayDNSRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayDNSRequest) GoString() string {
	return s.String()
}

func (s *SetGatewayDNSRequest) SetDnsServer(v string) *SetGatewayDNSRequest {
	s.DnsServer = &v
	return s
}

func (s *SetGatewayDNSRequest) SetGatewayId(v string) *SetGatewayDNSRequest {
	s.GatewayId = &v
	return s
}

func (s *SetGatewayDNSRequest) SetSecurityToken(v string) *SetGatewayDNSRequest {
	s.SecurityToken = &v
	return s
}

type SetGatewayDNSResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetGatewayDNSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayDNSResponseBody) GoString() string {
	return s.String()
}

func (s *SetGatewayDNSResponseBody) SetCode(v string) *SetGatewayDNSResponseBody {
	s.Code = &v
	return s
}

func (s *SetGatewayDNSResponseBody) SetMessage(v string) *SetGatewayDNSResponseBody {
	s.Message = &v
	return s
}

func (s *SetGatewayDNSResponseBody) SetRequestId(v string) *SetGatewayDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGatewayDNSResponseBody) SetSuccess(v bool) *SetGatewayDNSResponseBody {
	s.Success = &v
	return s
}

func (s *SetGatewayDNSResponseBody) SetTaskId(v string) *SetGatewayDNSResponseBody {
	s.TaskId = &v
	return s
}

type SetGatewayDNSResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGatewayDNSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGatewayDNSResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayDNSResponse) GoString() string {
	return s.String()
}

func (s *SetGatewayDNSResponse) SetHeaders(v map[string]*string) *SetGatewayDNSResponse {
	s.Headers = v
	return s
}

func (s *SetGatewayDNSResponse) SetStatusCode(v int32) *SetGatewayDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGatewayDNSResponse) SetBody(v *SetGatewayDNSResponseBody) *SetGatewayDNSResponse {
	s.Body = v
	return s
}

type SetGatewayLDAPInfoRequest struct {
	BaseDN        *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IsEnabled     *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	IsTls         *bool   `json:"IsTls,omitempty" xml:"IsTls,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RootDN        *string `json:"RootDN,omitempty" xml:"RootDN,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
}

func (s SetGatewayLDAPInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayLDAPInfoRequest) GoString() string {
	return s.String()
}

func (s *SetGatewayLDAPInfoRequest) SetBaseDN(v string) *SetGatewayLDAPInfoRequest {
	s.BaseDN = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetGatewayId(v string) *SetGatewayLDAPInfoRequest {
	s.GatewayId = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetIsEnabled(v bool) *SetGatewayLDAPInfoRequest {
	s.IsEnabled = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetIsTls(v bool) *SetGatewayLDAPInfoRequest {
	s.IsTls = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetPassword(v string) *SetGatewayLDAPInfoRequest {
	s.Password = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetRootDN(v string) *SetGatewayLDAPInfoRequest {
	s.RootDN = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetSecurityToken(v string) *SetGatewayLDAPInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetGatewayLDAPInfoRequest) SetServerIp(v string) *SetGatewayLDAPInfoRequest {
	s.ServerIp = &v
	return s
}

type SetGatewayLDAPInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetGatewayLDAPInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayLDAPInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetGatewayLDAPInfoResponseBody) SetCode(v string) *SetGatewayLDAPInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SetGatewayLDAPInfoResponseBody) SetMessage(v string) *SetGatewayLDAPInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetGatewayLDAPInfoResponseBody) SetRequestId(v string) *SetGatewayLDAPInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGatewayLDAPInfoResponseBody) SetSuccess(v bool) *SetGatewayLDAPInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SetGatewayLDAPInfoResponseBody) SetTaskId(v string) *SetGatewayLDAPInfoResponseBody {
	s.TaskId = &v
	return s
}

type SetGatewayLDAPInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGatewayLDAPInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGatewayLDAPInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGatewayLDAPInfoResponse) GoString() string {
	return s.String()
}

func (s *SetGatewayLDAPInfoResponse) SetHeaders(v map[string]*string) *SetGatewayLDAPInfoResponse {
	s.Headers = v
	return s
}

func (s *SetGatewayLDAPInfoResponse) SetStatusCode(v int32) *SetGatewayLDAPInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGatewayLDAPInfoResponse) SetBody(v *SetGatewayLDAPInfoResponseBody) *SetGatewayLDAPInfoResponse {
	s.Body = v
	return s
}

type SwitchCSGClientsReverseSyncConfigurationRequest struct {
	ClientIds                 []*string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty" type:"Repeated"`
	ClientRegionId            *string   `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	IsReverseSync             *bool     `json:"IsReverseSync,omitempty" xml:"IsReverseSync,omitempty"`
	ReverseSyncInternalSecond *int32    `json:"ReverseSyncInternalSecond,omitempty" xml:"ReverseSyncInternalSecond,omitempty"`
	SecurityToken             *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchCSGClientsReverseSyncConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchCSGClientsReverseSyncConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SwitchCSGClientsReverseSyncConfigurationRequest) SetClientIds(v []*string) *SwitchCSGClientsReverseSyncConfigurationRequest {
	s.ClientIds = v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationRequest) SetClientRegionId(v string) *SwitchCSGClientsReverseSyncConfigurationRequest {
	s.ClientRegionId = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationRequest) SetIsReverseSync(v bool) *SwitchCSGClientsReverseSyncConfigurationRequest {
	s.IsReverseSync = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationRequest) SetReverseSyncInternalSecond(v int32) *SwitchCSGClientsReverseSyncConfigurationRequest {
	s.ReverseSyncInternalSecond = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationRequest) SetSecurityToken(v string) *SwitchCSGClientsReverseSyncConfigurationRequest {
	s.SecurityToken = &v
	return s
}

type SwitchCSGClientsReverseSyncConfigurationShrinkRequest struct {
	ClientIdsShrink           *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	ClientRegionId            *string `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	IsReverseSync             *bool   `json:"IsReverseSync,omitempty" xml:"IsReverseSync,omitempty"`
	ReverseSyncInternalSecond *int32  `json:"ReverseSyncInternalSecond,omitempty" xml:"ReverseSyncInternalSecond,omitempty"`
	SecurityToken             *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchCSGClientsReverseSyncConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchCSGClientsReverseSyncConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *SwitchCSGClientsReverseSyncConfigurationShrinkRequest) SetClientIdsShrink(v string) *SwitchCSGClientsReverseSyncConfigurationShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationShrinkRequest) SetClientRegionId(v string) *SwitchCSGClientsReverseSyncConfigurationShrinkRequest {
	s.ClientRegionId = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationShrinkRequest) SetIsReverseSync(v bool) *SwitchCSGClientsReverseSyncConfigurationShrinkRequest {
	s.IsReverseSync = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationShrinkRequest) SetReverseSyncInternalSecond(v int32) *SwitchCSGClientsReverseSyncConfigurationShrinkRequest {
	s.ReverseSyncInternalSecond = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationShrinkRequest) SetSecurityToken(v string) *SwitchCSGClientsReverseSyncConfigurationShrinkRequest {
	s.SecurityToken = &v
	return s
}

type SwitchCSGClientsReverseSyncConfigurationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SwitchCSGClientsReverseSyncConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchCSGClientsReverseSyncConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponseBody) SetCode(v string) *SwitchCSGClientsReverseSyncConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponseBody) SetMessage(v string) *SwitchCSGClientsReverseSyncConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponseBody) SetRequestId(v string) *SwitchCSGClientsReverseSyncConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponseBody) SetSuccess(v bool) *SwitchCSGClientsReverseSyncConfigurationResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponseBody) SetTaskId(v string) *SwitchCSGClientsReverseSyncConfigurationResponseBody {
	s.TaskId = &v
	return s
}

type SwitchCSGClientsReverseSyncConfigurationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchCSGClientsReverseSyncConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchCSGClientsReverseSyncConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchCSGClientsReverseSyncConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponse) SetHeaders(v map[string]*string) *SwitchCSGClientsReverseSyncConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponse) SetStatusCode(v int32) *SwitchCSGClientsReverseSyncConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchCSGClientsReverseSyncConfigurationResponse) SetBody(v *SwitchCSGClientsReverseSyncConfigurationResponseBody) *SwitchCSGClientsReverseSyncConfigurationResponse {
	s.Body = v
	return s
}

type SwitchGatewayExpirationPolicyRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchGatewayExpirationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchGatewayExpirationPolicyRequest) GoString() string {
	return s.String()
}

func (s *SwitchGatewayExpirationPolicyRequest) SetGatewayId(v string) *SwitchGatewayExpirationPolicyRequest {
	s.GatewayId = &v
	return s
}

func (s *SwitchGatewayExpirationPolicyRequest) SetSecurityToken(v string) *SwitchGatewayExpirationPolicyRequest {
	s.SecurityToken = &v
	return s
}

type SwitchGatewayExpirationPolicyResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchGatewayExpirationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchGatewayExpirationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchGatewayExpirationPolicyResponseBody) SetCode(v string) *SwitchGatewayExpirationPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchGatewayExpirationPolicyResponseBody) SetMessage(v string) *SwitchGatewayExpirationPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchGatewayExpirationPolicyResponseBody) SetRequestId(v string) *SwitchGatewayExpirationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchGatewayExpirationPolicyResponseBody) SetSuccess(v bool) *SwitchGatewayExpirationPolicyResponseBody {
	s.Success = &v
	return s
}

type SwitchGatewayExpirationPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchGatewayExpirationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchGatewayExpirationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchGatewayExpirationPolicyResponse) GoString() string {
	return s.String()
}

func (s *SwitchGatewayExpirationPolicyResponse) SetHeaders(v map[string]*string) *SwitchGatewayExpirationPolicyResponse {
	s.Headers = v
	return s
}

func (s *SwitchGatewayExpirationPolicyResponse) SetStatusCode(v int32) *SwitchGatewayExpirationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchGatewayExpirationPolicyResponse) SetBody(v *SwitchGatewayExpirationPolicyResponseBody) *SwitchGatewayExpirationPolicyResponse {
	s.Body = v
	return s
}

type SwitchToSubscriptionRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchToSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchToSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *SwitchToSubscriptionRequest) SetGatewayId(v string) *SwitchToSubscriptionRequest {
	s.GatewayId = &v
	return s
}

func (s *SwitchToSubscriptionRequest) SetSecurityToken(v string) *SwitchToSubscriptionRequest {
	s.SecurityToken = &v
	return s
}

type SwitchToSubscriptionResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionURL *string `json:"SubscriptionURL,omitempty" xml:"SubscriptionURL,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchToSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchToSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchToSubscriptionResponseBody) SetCode(v string) *SwitchToSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *SwitchToSubscriptionResponseBody) SetMessage(v string) *SwitchToSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchToSubscriptionResponseBody) SetRequestId(v string) *SwitchToSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchToSubscriptionResponseBody) SetSubscriptionURL(v string) *SwitchToSubscriptionResponseBody {
	s.SubscriptionURL = &v
	return s
}

func (s *SwitchToSubscriptionResponseBody) SetSuccess(v bool) *SwitchToSubscriptionResponseBody {
	s.Success = &v
	return s
}

type SwitchToSubscriptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchToSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchToSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchToSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *SwitchToSubscriptionResponse) SetHeaders(v map[string]*string) *SwitchToSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *SwitchToSubscriptionResponse) SetStatusCode(v int32) *SwitchToSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchToSubscriptionResponse) SetBody(v *SwitchToSubscriptionResponseBody) *SwitchToSubscriptionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId         *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId       []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceRegionId *string                   `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	ResourceType     *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken    *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag              []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceRegionId(v string) *TagResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetSecurityToken(v string) *TagResourcesRequest {
	s.SecurityToken = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TriggerGatewayRemoteSyncRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId       *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Path          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s TriggerGatewayRemoteSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerGatewayRemoteSyncRequest) GoString() string {
	return s.String()
}

func (s *TriggerGatewayRemoteSyncRequest) SetGatewayId(v string) *TriggerGatewayRemoteSyncRequest {
	s.GatewayId = &v
	return s
}

func (s *TriggerGatewayRemoteSyncRequest) SetIndexId(v string) *TriggerGatewayRemoteSyncRequest {
	s.IndexId = &v
	return s
}

func (s *TriggerGatewayRemoteSyncRequest) SetPath(v string) *TriggerGatewayRemoteSyncRequest {
	s.Path = &v
	return s
}

func (s *TriggerGatewayRemoteSyncRequest) SetSecurityToken(v string) *TriggerGatewayRemoteSyncRequest {
	s.SecurityToken = &v
	return s
}

type TriggerGatewayRemoteSyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TriggerGatewayRemoteSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerGatewayRemoteSyncResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerGatewayRemoteSyncResponseBody) SetCode(v string) *TriggerGatewayRemoteSyncResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerGatewayRemoteSyncResponseBody) SetMessage(v string) *TriggerGatewayRemoteSyncResponseBody {
	s.Message = &v
	return s
}

func (s *TriggerGatewayRemoteSyncResponseBody) SetRequestId(v string) *TriggerGatewayRemoteSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerGatewayRemoteSyncResponseBody) SetSuccess(v bool) *TriggerGatewayRemoteSyncResponseBody {
	s.Success = &v
	return s
}

func (s *TriggerGatewayRemoteSyncResponseBody) SetTaskId(v string) *TriggerGatewayRemoteSyncResponseBody {
	s.TaskId = &v
	return s
}

type TriggerGatewayRemoteSyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerGatewayRemoteSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerGatewayRemoteSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerGatewayRemoteSyncResponse) GoString() string {
	return s.String()
}

func (s *TriggerGatewayRemoteSyncResponse) SetHeaders(v map[string]*string) *TriggerGatewayRemoteSyncResponse {
	s.Headers = v
	return s
}

func (s *TriggerGatewayRemoteSyncResponse) SetStatusCode(v int32) *TriggerGatewayRemoteSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerGatewayRemoteSyncResponse) SetBody(v *TriggerGatewayRemoteSyncResponseBody) *TriggerGatewayRemoteSyncResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All              *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId       []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceRegionId *string   `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	ResourceType     *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken    *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TagKey           []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceRegionId(v string) *UntagResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetSecurityToken(v string) *UntagResourcesRequest {
	s.SecurityToken = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateGatewayBlockVolumeRequest struct {
	ChapEnabled    *bool   `json:"ChapEnabled,omitempty" xml:"ChapEnabled,omitempty"`
	ChapInPassword *string `json:"ChapInPassword,omitempty" xml:"ChapInPassword,omitempty"`
	ChapInUser     *string `json:"ChapInUser,omitempty" xml:"ChapInUser,omitempty"`
	GatewayId      *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IndexId        *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Size           *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s UpdateGatewayBlockVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayBlockVolumeRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayBlockVolumeRequest) SetChapEnabled(v bool) *UpdateGatewayBlockVolumeRequest {
	s.ChapEnabled = &v
	return s
}

func (s *UpdateGatewayBlockVolumeRequest) SetChapInPassword(v string) *UpdateGatewayBlockVolumeRequest {
	s.ChapInPassword = &v
	return s
}

func (s *UpdateGatewayBlockVolumeRequest) SetChapInUser(v string) *UpdateGatewayBlockVolumeRequest {
	s.ChapInUser = &v
	return s
}

func (s *UpdateGatewayBlockVolumeRequest) SetGatewayId(v string) *UpdateGatewayBlockVolumeRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayBlockVolumeRequest) SetIndexId(v string) *UpdateGatewayBlockVolumeRequest {
	s.IndexId = &v
	return s
}

func (s *UpdateGatewayBlockVolumeRequest) SetSecurityToken(v string) *UpdateGatewayBlockVolumeRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateGatewayBlockVolumeRequest) SetSize(v int64) *UpdateGatewayBlockVolumeRequest {
	s.Size = &v
	return s
}

type UpdateGatewayBlockVolumeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateGatewayBlockVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayBlockVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayBlockVolumeResponseBody) SetCode(v string) *UpdateGatewayBlockVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayBlockVolumeResponseBody) SetMessage(v string) *UpdateGatewayBlockVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayBlockVolumeResponseBody) SetRequestId(v string) *UpdateGatewayBlockVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayBlockVolumeResponseBody) SetSuccess(v bool) *UpdateGatewayBlockVolumeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayBlockVolumeResponseBody) SetTaskId(v string) *UpdateGatewayBlockVolumeResponseBody {
	s.TaskId = &v
	return s
}

type UpdateGatewayBlockVolumeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGatewayBlockVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayBlockVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayBlockVolumeResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayBlockVolumeResponse) SetHeaders(v map[string]*string) *UpdateGatewayBlockVolumeResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayBlockVolumeResponse) SetStatusCode(v int32) *UpdateGatewayBlockVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayBlockVolumeResponse) SetBody(v *UpdateGatewayBlockVolumeResponseBody) *UpdateGatewayBlockVolumeResponse {
	s.Body = v
	return s
}

type UpdateGatewayFileShareRequest struct {
	AccessBasedEnumeration *bool   `json:"AccessBasedEnumeration,omitempty" xml:"AccessBasedEnumeration,omitempty"`
	BackendLimit           *int32  `json:"BackendLimit,omitempty" xml:"BackendLimit,omitempty"`
	Browsable              *bool   `json:"Browsable,omitempty" xml:"Browsable,omitempty"`
	BypassCacheRead        *bool   `json:"BypassCacheRead,omitempty" xml:"BypassCacheRead,omitempty"`
	CacheMode              *string `json:"CacheMode,omitempty" xml:"CacheMode,omitempty"`
	ClientSideCmk          *string `json:"ClientSideCmk,omitempty" xml:"ClientSideCmk,omitempty"`
	ClientSideEncryption   *bool   `json:"ClientSideEncryption,omitempty" xml:"ClientSideEncryption,omitempty"`
	DirectIO               *bool   `json:"DirectIO,omitempty" xml:"DirectIO,omitempty"`
	DownloadLimit          *int32  `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	FastReclaim            *bool   `json:"FastReclaim,omitempty" xml:"FastReclaim,omitempty"`
	FrontendLimit          *int32  `json:"FrontendLimit,omitempty" xml:"FrontendLimit,omitempty"`
	GatewayId              *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	IgnoreDelete           *bool   `json:"IgnoreDelete,omitempty" xml:"IgnoreDelete,omitempty"`
	InPlace                *bool   `json:"InPlace,omitempty" xml:"InPlace,omitempty"`
	IndexId                *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	KmsRotatePeriod        *int64  `json:"KmsRotatePeriod,omitempty" xml:"KmsRotatePeriod,omitempty"`
	LagPeriod              *int64  `json:"LagPeriod,omitempty" xml:"LagPeriod,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NfsV4Optimization      *bool   `json:"NfsV4Optimization,omitempty" xml:"NfsV4Optimization,omitempty"`
	PollingInterval        *int32  `json:"PollingInterval,omitempty" xml:"PollingInterval,omitempty"`
	ReadOnlyClientList     *string `json:"ReadOnlyClientList,omitempty" xml:"ReadOnlyClientList,omitempty"`
	ReadOnlyUserList       *string `json:"ReadOnlyUserList,omitempty" xml:"ReadOnlyUserList,omitempty"`
	ReadWriteClientList    *string `json:"ReadWriteClientList,omitempty" xml:"ReadWriteClientList,omitempty"`
	ReadWriteUserList      *string `json:"ReadWriteUserList,omitempty" xml:"ReadWriteUserList,omitempty"`
	RemoteSync             *bool   `json:"RemoteSync,omitempty" xml:"RemoteSync,omitempty"`
	RemoteSyncDownload     *bool   `json:"RemoteSyncDownload,omitempty" xml:"RemoteSyncDownload,omitempty"`
	SecurityToken          *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServerSideCmk          *string `json:"ServerSideCmk,omitempty" xml:"ServerSideCmk,omitempty"`
	ServerSideEncryption   *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	Squash                 *string `json:"Squash,omitempty" xml:"Squash,omitempty"`
	TransferAcceleration   *bool   `json:"TransferAcceleration,omitempty" xml:"TransferAcceleration,omitempty"`
	WindowsAcl             *bool   `json:"WindowsAcl,omitempty" xml:"WindowsAcl,omitempty"`
}

func (s UpdateGatewayFileShareRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayFileShareRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFileShareRequest) SetAccessBasedEnumeration(v bool) *UpdateGatewayFileShareRequest {
	s.AccessBasedEnumeration = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetBackendLimit(v int32) *UpdateGatewayFileShareRequest {
	s.BackendLimit = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetBrowsable(v bool) *UpdateGatewayFileShareRequest {
	s.Browsable = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetBypassCacheRead(v bool) *UpdateGatewayFileShareRequest {
	s.BypassCacheRead = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetCacheMode(v string) *UpdateGatewayFileShareRequest {
	s.CacheMode = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetClientSideCmk(v string) *UpdateGatewayFileShareRequest {
	s.ClientSideCmk = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetClientSideEncryption(v bool) *UpdateGatewayFileShareRequest {
	s.ClientSideEncryption = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetDirectIO(v bool) *UpdateGatewayFileShareRequest {
	s.DirectIO = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetDownloadLimit(v int32) *UpdateGatewayFileShareRequest {
	s.DownloadLimit = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetFastReclaim(v bool) *UpdateGatewayFileShareRequest {
	s.FastReclaim = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetFrontendLimit(v int32) *UpdateGatewayFileShareRequest {
	s.FrontendLimit = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetGatewayId(v string) *UpdateGatewayFileShareRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetIgnoreDelete(v bool) *UpdateGatewayFileShareRequest {
	s.IgnoreDelete = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetInPlace(v bool) *UpdateGatewayFileShareRequest {
	s.InPlace = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetIndexId(v string) *UpdateGatewayFileShareRequest {
	s.IndexId = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetKmsRotatePeriod(v int64) *UpdateGatewayFileShareRequest {
	s.KmsRotatePeriod = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetLagPeriod(v int64) *UpdateGatewayFileShareRequest {
	s.LagPeriod = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetName(v string) *UpdateGatewayFileShareRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetNfsV4Optimization(v bool) *UpdateGatewayFileShareRequest {
	s.NfsV4Optimization = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetPollingInterval(v int32) *UpdateGatewayFileShareRequest {
	s.PollingInterval = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetReadOnlyClientList(v string) *UpdateGatewayFileShareRequest {
	s.ReadOnlyClientList = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetReadOnlyUserList(v string) *UpdateGatewayFileShareRequest {
	s.ReadOnlyUserList = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetReadWriteClientList(v string) *UpdateGatewayFileShareRequest {
	s.ReadWriteClientList = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetReadWriteUserList(v string) *UpdateGatewayFileShareRequest {
	s.ReadWriteUserList = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetRemoteSync(v bool) *UpdateGatewayFileShareRequest {
	s.RemoteSync = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetRemoteSyncDownload(v bool) *UpdateGatewayFileShareRequest {
	s.RemoteSyncDownload = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetSecurityToken(v string) *UpdateGatewayFileShareRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetServerSideCmk(v string) *UpdateGatewayFileShareRequest {
	s.ServerSideCmk = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetServerSideEncryption(v bool) *UpdateGatewayFileShareRequest {
	s.ServerSideEncryption = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetSquash(v string) *UpdateGatewayFileShareRequest {
	s.Squash = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetTransferAcceleration(v bool) *UpdateGatewayFileShareRequest {
	s.TransferAcceleration = &v
	return s
}

func (s *UpdateGatewayFileShareRequest) SetWindowsAcl(v bool) *UpdateGatewayFileShareRequest {
	s.WindowsAcl = &v
	return s
}

type UpdateGatewayFileShareResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateGatewayFileShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayFileShareResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFileShareResponseBody) SetCode(v string) *UpdateGatewayFileShareResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayFileShareResponseBody) SetMessage(v string) *UpdateGatewayFileShareResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayFileShareResponseBody) SetRequestId(v string) *UpdateGatewayFileShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayFileShareResponseBody) SetSuccess(v bool) *UpdateGatewayFileShareResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayFileShareResponseBody) SetTaskId(v string) *UpdateGatewayFileShareResponseBody {
	s.TaskId = &v
	return s
}

type UpdateGatewayFileShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGatewayFileShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayFileShareResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayFileShareResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFileShareResponse) SetHeaders(v map[string]*string) *UpdateGatewayFileShareResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayFileShareResponse) SetStatusCode(v int32) *UpdateGatewayFileShareResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayFileShareResponse) SetBody(v *UpdateGatewayFileShareResponseBody) *UpdateGatewayFileShareResponse {
	s.Body = v
	return s
}

type UpgradeGatewayRequest struct {
	GatewayId     *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpgradeGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayRequest) SetGatewayId(v string) *UpgradeGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *UpgradeGatewayRequest) SetSecurityToken(v string) *UpgradeGatewayRequest {
	s.SecurityToken = &v
	return s
}

type UpgradeGatewayResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayResponseBody) SetCode(v string) *UpgradeGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetMessage(v string) *UpgradeGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetRequestId(v string) *UpgradeGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetSuccess(v bool) *UpgradeGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetTaskId(v string) *UpgradeGatewayResponseBody {
	s.TaskId = &v
	return s
}

type UpgradeGatewayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayResponse) SetHeaders(v map[string]*string) *UpgradeGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpgradeGatewayResponse) SetStatusCode(v int32) *UpgradeGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeGatewayResponse) SetBody(v *UpgradeGatewayResponseBody) *UpgradeGatewayResponse {
	s.Body = v
	return s
}

type UploadCSGClientLogRequest struct {
	ClientId       *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientRegionId *string `json:"ClientRegionId,omitempty" xml:"ClientRegionId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UploadCSGClientLogRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCSGClientLogRequest) GoString() string {
	return s.String()
}

func (s *UploadCSGClientLogRequest) SetClientId(v string) *UploadCSGClientLogRequest {
	s.ClientId = &v
	return s
}

func (s *UploadCSGClientLogRequest) SetClientRegionId(v string) *UploadCSGClientLogRequest {
	s.ClientRegionId = &v
	return s
}

func (s *UploadCSGClientLogRequest) SetSecurityToken(v string) *UploadCSGClientLogRequest {
	s.SecurityToken = &v
	return s
}

type UploadCSGClientLogResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UploadCSGClientLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCSGClientLogResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCSGClientLogResponseBody) SetCode(v string) *UploadCSGClientLogResponseBody {
	s.Code = &v
	return s
}

func (s *UploadCSGClientLogResponseBody) SetMessage(v string) *UploadCSGClientLogResponseBody {
	s.Message = &v
	return s
}

func (s *UploadCSGClientLogResponseBody) SetRequestId(v string) *UploadCSGClientLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCSGClientLogResponseBody) SetSuccess(v bool) *UploadCSGClientLogResponseBody {
	s.Success = &v
	return s
}

func (s *UploadCSGClientLogResponseBody) SetTaskId(v string) *UploadCSGClientLogResponseBody {
	s.TaskId = &v
	return s
}

type UploadCSGClientLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadCSGClientLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadCSGClientLogResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCSGClientLogResponse) GoString() string {
	return s.String()
}

func (s *UploadCSGClientLogResponse) SetHeaders(v map[string]*string) *UploadCSGClientLogResponse {
	s.Headers = v
	return s
}

func (s *UploadCSGClientLogResponse) SetStatusCode(v int32) *UploadCSGClientLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCSGClientLogResponse) SetBody(v *UploadCSGClientLogResponseBody) *UploadCSGClientLogResponse {
	s.Body = v
	return s
}

type UploadGatewayLogRequest struct {
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s UploadGatewayLogRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadGatewayLogRequest) GoString() string {
	return s.String()
}

func (s *UploadGatewayLogRequest) SetGatewayId(v string) *UploadGatewayLogRequest {
	s.GatewayId = &v
	return s
}

type UploadGatewayLogResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UploadGatewayLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadGatewayLogResponseBody) GoString() string {
	return s.String()
}

func (s *UploadGatewayLogResponseBody) SetCode(v string) *UploadGatewayLogResponseBody {
	s.Code = &v
	return s
}

func (s *UploadGatewayLogResponseBody) SetMessage(v string) *UploadGatewayLogResponseBody {
	s.Message = &v
	return s
}

func (s *UploadGatewayLogResponseBody) SetRequestId(v string) *UploadGatewayLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadGatewayLogResponseBody) SetSuccess(v bool) *UploadGatewayLogResponseBody {
	s.Success = &v
	return s
}

func (s *UploadGatewayLogResponseBody) SetTaskId(v string) *UploadGatewayLogResponseBody {
	s.TaskId = &v
	return s
}

type UploadGatewayLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadGatewayLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadGatewayLogResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadGatewayLogResponse) GoString() string {
	return s.String()
}

func (s *UploadGatewayLogResponse) SetHeaders(v map[string]*string) *UploadGatewayLogResponse {
	s.Headers = v
	return s
}

func (s *UploadGatewayLogResponse) SetStatusCode(v int32) *UploadGatewayLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadGatewayLogResponse) SetBody(v *UploadGatewayLogResponseBody) *UploadGatewayLogResponse {
	s.Body = v
	return s
}

type ValidateExpressSyncConfigRequest struct {
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketPrefix  *string `json:"BucketPrefix,omitempty" xml:"BucketPrefix,omitempty"`
	BucketRegion  *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ValidateExpressSyncConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateExpressSyncConfigRequest) GoString() string {
	return s.String()
}

func (s *ValidateExpressSyncConfigRequest) SetBucketName(v string) *ValidateExpressSyncConfigRequest {
	s.BucketName = &v
	return s
}

func (s *ValidateExpressSyncConfigRequest) SetBucketPrefix(v string) *ValidateExpressSyncConfigRequest {
	s.BucketPrefix = &v
	return s
}

func (s *ValidateExpressSyncConfigRequest) SetBucketRegion(v string) *ValidateExpressSyncConfigRequest {
	s.BucketRegion = &v
	return s
}

func (s *ValidateExpressSyncConfigRequest) SetName(v string) *ValidateExpressSyncConfigRequest {
	s.Name = &v
	return s
}

func (s *ValidateExpressSyncConfigRequest) SetSecurityToken(v string) *ValidateExpressSyncConfigRequest {
	s.SecurityToken = &v
	return s
}

type ValidateExpressSyncConfigResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsValid         *bool   `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ValidateMessage *string `json:"ValidateMessage,omitempty" xml:"ValidateMessage,omitempty"`
}

func (s ValidateExpressSyncConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateExpressSyncConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateExpressSyncConfigResponseBody) SetCode(v string) *ValidateExpressSyncConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateExpressSyncConfigResponseBody) SetIsValid(v bool) *ValidateExpressSyncConfigResponseBody {
	s.IsValid = &v
	return s
}

func (s *ValidateExpressSyncConfigResponseBody) SetMessage(v string) *ValidateExpressSyncConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateExpressSyncConfigResponseBody) SetRequestId(v string) *ValidateExpressSyncConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateExpressSyncConfigResponseBody) SetSuccess(v bool) *ValidateExpressSyncConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateExpressSyncConfigResponseBody) SetValidateMessage(v string) *ValidateExpressSyncConfigResponseBody {
	s.ValidateMessage = &v
	return s
}

type ValidateExpressSyncConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateExpressSyncConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateExpressSyncConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateExpressSyncConfigResponse) GoString() string {
	return s.String()
}

func (s *ValidateExpressSyncConfigResponse) SetHeaders(v map[string]*string) *ValidateExpressSyncConfigResponse {
	s.Headers = v
	return s
}

func (s *ValidateExpressSyncConfigResponse) SetStatusCode(v int32) *ValidateExpressSyncConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateExpressSyncConfigResponse) SetBody(v *ValidateExpressSyncConfigResponseBody) *ValidateExpressSyncConfigResponse {
	s.Body = v
	return s
}

type ValidateGatewayNameRequest struct {
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StorageBundleId *string `json:"StorageBundleId,omitempty" xml:"StorageBundleId,omitempty"`
}

func (s ValidateGatewayNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateGatewayNameRequest) GoString() string {
	return s.String()
}

func (s *ValidateGatewayNameRequest) SetName(v string) *ValidateGatewayNameRequest {
	s.Name = &v
	return s
}

func (s *ValidateGatewayNameRequest) SetStorageBundleId(v string) *ValidateGatewayNameRequest {
	s.StorageBundleId = &v
	return s
}

type ValidateGatewayNameResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsValid   *bool   `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateGatewayNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateGatewayNameResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateGatewayNameResponseBody) SetCode(v string) *ValidateGatewayNameResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateGatewayNameResponseBody) SetIsValid(v bool) *ValidateGatewayNameResponseBody {
	s.IsValid = &v
	return s
}

func (s *ValidateGatewayNameResponseBody) SetMessage(v string) *ValidateGatewayNameResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateGatewayNameResponseBody) SetRequestId(v string) *ValidateGatewayNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateGatewayNameResponseBody) SetSuccess(v bool) *ValidateGatewayNameResponseBody {
	s.Success = &v
	return s
}

type ValidateGatewayNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateGatewayNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateGatewayNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateGatewayNameResponse) GoString() string {
	return s.String()
}

func (s *ValidateGatewayNameResponse) SetHeaders(v map[string]*string) *ValidateGatewayNameResponse {
	s.Headers = v
	return s
}

func (s *ValidateGatewayNameResponse) SetStatusCode(v int32) *ValidateGatewayNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateGatewayNameResponse) SetBody(v *ValidateGatewayNameResponseBody) *ValidateGatewayNameResponse {
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
		"cn-qingdao":     tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-beijing":     tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-chengdu":     tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-huhehaote":   tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou":    tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":    tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"cn-hongkong":    tea.String("sgw.cn-shanghai.aliyuncs.com"),
		"us-east-1":      tea.String("sgw.us-west-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sgw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ActivateAllInOneGatewayWithOptions(request *ActivateAllInOneGatewayRequest, runtime *util.RuntimeOptions) (_result *ActivateAllInOneGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceNumber)) {
		query["DeviceNumber"] = request.DeviceNumber
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateAllInOneGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateAllInOneGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateAllInOneGateway(request *ActivateAllInOneGatewayRequest) (_result *ActivateAllInOneGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateAllInOneGatewayResponse{}
	_body, _err := client.ActivateAllInOneGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivateGatewayWithOptions(request *ActivateGatewayRequest, runtime *util.RuntimeOptions) (_result *ActivateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		query["Model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateGateway(request *ActivateGatewayRequest) (_result *ActivateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateGatewayResponse{}
	_body, _err := client.ActivateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSharesToExpressSyncWithOptions(request *AddSharesToExpressSyncRequest, runtime *util.RuntimeOptions) (_result *AddSharesToExpressSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpressSyncId)) {
		query["ExpressSyncId"] = request.ExpressSyncId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayShares)) {
		query["GatewayShares"] = request.GatewayShares
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSharesToExpressSync"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSharesToExpressSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSharesToExpressSync(request *AddSharesToExpressSyncRequest) (_result *AddSharesToExpressSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSharesToExpressSyncResponse{}
	_body, _err := client.AddSharesToExpressSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTagsToGatewayWithOptions(request *AddTagsToGatewayRequest, runtime *util.RuntimeOptions) (_result *AddTagsToGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTagsToGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTagsToGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTagsToGateway(request *AddTagsToGatewayRequest) (_result *AddTagsToGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsToGatewayResponse{}
	_body, _err := client.AddTagsToGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckActivationKeyWithOptions(request *CheckActivationKeyRequest, runtime *util.RuntimeOptions) (_result *CheckActivationKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CryptKey)) {
		query["CryptKey"] = request.CryptKey
	}

	if !tea.BoolValue(util.IsUnset(request.CryptText)) {
		query["CryptText"] = request.CryptText
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckActivationKey"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckActivationKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckActivationKey(request *CheckActivationKeyRequest) (_result *CheckActivationKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckActivationKeyResponse{}
	_body, _err := client.CheckActivationKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckBlockVolumeNameWithOptions(request *CheckBlockVolumeNameRequest, runtime *util.RuntimeOptions) (_result *CheckBlockVolumeNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketEndpoint)) {
		query["BucketEndpoint"] = request.BucketEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VolumeName)) {
		query["VolumeName"] = request.VolumeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckBlockVolumeName"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckBlockVolumeNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckBlockVolumeName(request *CheckBlockVolumeNameRequest) (_result *CheckBlockVolumeNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckBlockVolumeNameResponse{}
	_body, _err := client.CheckBlockVolumeNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckGatewayEssdSupportWithOptions(request *CheckGatewayEssdSupportRequest, runtime *util.RuntimeOptions) (_result *CheckGatewayEssdSupportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckGatewayEssdSupport"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckGatewayEssdSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckGatewayEssdSupport(request *CheckGatewayEssdSupportRequest) (_result *CheckGatewayEssdSupportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckGatewayEssdSupportResponse{}
	_body, _err := client.CheckGatewayEssdSupportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMnsServiceWithOptions(request *CheckMnsServiceRequest, runtime *util.RuntimeOptions) (_result *CheckMnsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckMnsService"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMnsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMnsService(request *CheckMnsServiceRequest) (_result *CheckMnsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMnsServiceResponse{}
	_body, _err := client.CheckMnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckRoleWithOptions(request *CheckRoleRequest, runtime *util.RuntimeOptions) (_result *CheckRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckRole"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckRole(request *CheckRoleRequest) (_result *CheckRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckRoleResponse{}
	_body, _err := client.CheckRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckSlrRoleWithOptions(request *CheckSlrRoleRequest, runtime *util.RuntimeOptions) (_result *CheckSlrRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateIfNotExist)) {
		query["CreateIfNotExist"] = request.CreateIfNotExist
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSlrRole"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSlrRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSlrRole(request *CheckSlrRoleRequest) (_result *CheckSlrRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSlrRoleResponse{}
	_body, _err := client.CheckSlrRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUpgradeVersionWithOptions(request *CheckUpgradeVersionRequest, runtime *util.RuntimeOptions) (_result *CheckUpgradeVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayVersion)) {
		query["GatewayVersion"] = request.GatewayVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUpgradeVersion"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUpgradeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUpgradeVersion(request *CheckUpgradeVersionRequest) (_result *CheckUpgradeVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUpgradeVersionResponse{}
	_body, _err := client.CheckUpgradeVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCacheWithOptions(request *CreateCacheRequest, runtime *util.RuntimeOptions) (_result *CreateCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SizeInGB)) {
		query["SizeInGB"] = request.SizeInGB
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCache"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCache(request *CreateCacheRequest) (_result *CreateCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCacheResponse{}
	_body, _err := client.CreateCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateElasticGatewayPrivateZoneWithOptions(request *CreateElasticGatewayPrivateZoneRequest, runtime *util.RuntimeOptions) (_result *CreateElasticGatewayPrivateZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateElasticGatewayPrivateZone"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateElasticGatewayPrivateZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateElasticGatewayPrivateZone(request *CreateElasticGatewayPrivateZoneRequest) (_result *CreateElasticGatewayPrivateZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateElasticGatewayPrivateZoneResponse{}
	_body, _err := client.CreateElasticGatewayPrivateZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExpressSyncWithOptions(request *CreateExpressSyncRequest, runtime *util.RuntimeOptions) (_result *CreateExpressSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketPrefix)) {
		query["BucketPrefix"] = request.BucketPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.BucketRegion)) {
		query["BucketRegion"] = request.BucketRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExpressSync"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExpressSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExpressSync(request *CreateExpressSyncRequest) (_result *CreateExpressSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExpressSyncResponse{}
	_body, _err := client.CreateExpressSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayClass)) {
		query["GatewayClass"] = request.GatewayClass
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		query["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PostPaid)) {
		query["PostPaid"] = request.PostPaid
	}

	if !tea.BoolValue(util.IsUnset(request.PublicNetworkBandwidth)) {
		query["PublicNetworkBandwidth"] = request.PublicNetworkBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseAfterExpiration)) {
		query["ReleaseAfterExpiration"] = request.ReleaseAfterExpiration
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayBlockVolumeWithOptions(request *CreateGatewayBlockVolumeRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayBlockVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheMode)) {
		query["CacheMode"] = request.CacheMode
	}

	if !tea.BoolValue(util.IsUnset(request.ChapEnabled)) {
		query["ChapEnabled"] = request.ChapEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ChapInPassword)) {
		query["ChapInPassword"] = request.ChapInPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ChapInUser)) {
		query["ChapInUser"] = request.ChapInUser
	}

	if !tea.BoolValue(util.IsUnset(request.ChunkSize)) {
		query["ChunkSize"] = request.ChunkSize
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalFilePath)) {
		query["LocalFilePath"] = request.LocalFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketSsl)) {
		query["OssBucketSsl"] = request.OssBucketSsl
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Recovery)) {
		query["Recovery"] = request.Recovery
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.VolumeProtocol)) {
		query["VolumeProtocol"] = request.VolumeProtocol
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayBlockVolume"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayBlockVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewayBlockVolume(request *CreateGatewayBlockVolumeRequest) (_result *CreateGatewayBlockVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayBlockVolumeResponse{}
	_body, _err := client.CreateGatewayBlockVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayCacheDiskWithOptions(request *CreateGatewayCacheDiskRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayCacheDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheDiskCategory)) {
		query["CacheDiskCategory"] = request.CacheDiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.CacheDiskSizeInGB)) {
		query["CacheDiskSizeInGB"] = request.CacheDiskSizeInGB
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayCacheDisk"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayCacheDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewayCacheDisk(request *CreateGatewayCacheDiskRequest) (_result *CreateGatewayCacheDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayCacheDiskResponse{}
	_body, _err := client.CreateGatewayCacheDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayFileShareWithOptions(request *CreateGatewayFileShareRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayFileShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessBasedEnumeration)) {
		query["AccessBasedEnumeration"] = request.AccessBasedEnumeration
	}

	if !tea.BoolValue(util.IsUnset(request.BackendLimit)) {
		query["BackendLimit"] = request.BackendLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Browsable)) {
		query["Browsable"] = request.Browsable
	}

	if !tea.BoolValue(util.IsUnset(request.BypassCacheRead)) {
		query["BypassCacheRead"] = request.BypassCacheRead
	}

	if !tea.BoolValue(util.IsUnset(request.CacheMode)) {
		query["CacheMode"] = request.CacheMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSideCmk)) {
		query["ClientSideCmk"] = request.ClientSideCmk
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSideEncryption)) {
		query["ClientSideEncryption"] = request.ClientSideEncryption
	}

	if !tea.BoolValue(util.IsUnset(request.DirectIO)) {
		query["DirectIO"] = request.DirectIO
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadLimit)) {
		query["DownloadLimit"] = request.DownloadLimit
	}

	if !tea.BoolValue(util.IsUnset(request.FastReclaim)) {
		query["FastReclaim"] = request.FastReclaim
	}

	if !tea.BoolValue(util.IsUnset(request.FrontendLimit)) {
		query["FrontendLimit"] = request.FrontendLimit
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreDelete)) {
		query["IgnoreDelete"] = request.IgnoreDelete
	}

	if !tea.BoolValue(util.IsUnset(request.InPlace)) {
		query["InPlace"] = request.InPlace
	}

	if !tea.BoolValue(util.IsUnset(request.KmsRotatePeriod)) {
		query["KmsRotatePeriod"] = request.KmsRotatePeriod
	}

	if !tea.BoolValue(util.IsUnset(request.LagPeriod)) {
		query["LagPeriod"] = request.LagPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.LocalFilePath)) {
		query["LocalFilePath"] = request.LocalFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NfsV4Optimization)) {
		query["NfsV4Optimization"] = request.NfsV4Optimization
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketSsl)) {
		query["OssBucketSsl"] = request.OssBucketSsl
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.PartialSyncPaths)) {
		query["PartialSyncPaths"] = request.PartialSyncPaths
	}

	if !tea.BoolValue(util.IsUnset(request.PathPrefix)) {
		query["PathPrefix"] = request.PathPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.PollingInterval)) {
		query["PollingInterval"] = request.PollingInterval
	}

	if !tea.BoolValue(util.IsUnset(request.ReadOnlyClientList)) {
		query["ReadOnlyClientList"] = request.ReadOnlyClientList
	}

	if !tea.BoolValue(util.IsUnset(request.ReadOnlyUserList)) {
		query["ReadOnlyUserList"] = request.ReadOnlyUserList
	}

	if !tea.BoolValue(util.IsUnset(request.ReadWriteClientList)) {
		query["ReadWriteClientList"] = request.ReadWriteClientList
	}

	if !tea.BoolValue(util.IsUnset(request.ReadWriteUserList)) {
		query["ReadWriteUserList"] = request.ReadWriteUserList
	}

	if !tea.BoolValue(util.IsUnset(request.RemoteSync)) {
		query["RemoteSync"] = request.RemoteSync
	}

	if !tea.BoolValue(util.IsUnset(request.RemoteSyncDownload)) {
		query["RemoteSyncDownload"] = request.RemoteSyncDownload
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerSideAlgorithm)) {
		query["ServerSideAlgorithm"] = request.ServerSideAlgorithm
	}

	if !tea.BoolValue(util.IsUnset(request.ServerSideCmk)) {
		query["ServerSideCmk"] = request.ServerSideCmk
	}

	if !tea.BoolValue(util.IsUnset(request.ServerSideEncryption)) {
		query["ServerSideEncryption"] = request.ServerSideEncryption
	}

	if !tea.BoolValue(util.IsUnset(request.ShareProtocol)) {
		query["ShareProtocol"] = request.ShareProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.Squash)) {
		query["Squash"] = request.Squash
	}

	if !tea.BoolValue(util.IsUnset(request.SupportArchive)) {
		query["SupportArchive"] = request.SupportArchive
	}

	if !tea.BoolValue(util.IsUnset(request.TransferAcceleration)) {
		query["TransferAcceleration"] = request.TransferAcceleration
	}

	if !tea.BoolValue(util.IsUnset(request.WindowsAcl)) {
		query["WindowsAcl"] = request.WindowsAcl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayFileShare"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayFileShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewayFileShare(request *CreateGatewayFileShareRequest) (_result *CreateGatewayFileShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayFileShareResponse{}
	_body, _err := client.CreateGatewayFileShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayLoggingWithOptions(request *CreateGatewayLoggingRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogstore)) {
		query["SlsLogstore"] = request.SlsLogstore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayLogging"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewayLogging(request *CreateGatewayLoggingRequest) (_result *CreateGatewayLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayLoggingResponse{}
	_body, _err := client.CreateGatewayLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewaySMBUserWithOptions(request *CreateGatewaySMBUserRequest, runtime *util.RuntimeOptions) (_result *CreateGatewaySMBUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewaySMBUser"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewaySMBUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewaySMBUser(request *CreateGatewaySMBUserRequest) (_result *CreateGatewaySMBUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewaySMBUserResponse{}
	_body, _err := client.CreateGatewaySMBUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStorageBundleWithOptions(request *CreateStorageBundleRequest, runtime *util.RuntimeOptions) (_result *CreateStorageBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendBucketRegionId)) {
		query["BackendBucketRegionId"] = request.BackendBucketRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStorageBundle"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStorageBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStorageBundle(request *CreateStorageBundleRequest) (_result *CreateStorageBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStorageBundleResponse{}
	_body, _err := client.CreateStorageBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCSGClientsWithOptions(tmpReq *DeleteCSGClientsRequest, runtime *util.RuntimeOptions) (_result *DeleteCSGClientsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteCSGClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClientIds)) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, tea.String("ClientIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdsShrink)) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientRegionId)) {
		query["ClientRegionId"] = request.ClientRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCSGClients"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCSGClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCSGClients(request *DeleteCSGClientsRequest) (_result *DeleteCSGClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCSGClientsResponse{}
	_body, _err := client.DeleteCSGClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteElasticGatewayPrivateZoneWithOptions(request *DeleteElasticGatewayPrivateZoneRequest, runtime *util.RuntimeOptions) (_result *DeleteElasticGatewayPrivateZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteElasticGatewayPrivateZone"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteElasticGatewayPrivateZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteElasticGatewayPrivateZone(request *DeleteElasticGatewayPrivateZoneRequest) (_result *DeleteElasticGatewayPrivateZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteElasticGatewayPrivateZoneResponse{}
	_body, _err := client.DeleteElasticGatewayPrivateZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExpressSyncWithOptions(request *DeleteExpressSyncRequest, runtime *util.RuntimeOptions) (_result *DeleteExpressSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpressSyncId)) {
		query["ExpressSyncId"] = request.ExpressSyncId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExpressSync"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExpressSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExpressSync(request *DeleteExpressSyncRequest) (_result *DeleteExpressSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExpressSyncResponse{}
	_body, _err := client.DeleteExpressSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonDetail)) {
		query["ReasonDetail"] = request.ReasonDetail
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonType)) {
		query["ReasonType"] = request.ReasonType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayBlockVolumesWithOptions(request *DeleteGatewayBlockVolumesRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayBlockVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSourceDeletion)) {
		query["IsSourceDeletion"] = request.IsSourceDeletion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayBlockVolumes"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayBlockVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewayBlockVolumes(request *DeleteGatewayBlockVolumesRequest) (_result *DeleteGatewayBlockVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayBlockVolumesResponse{}
	_body, _err := client.DeleteGatewayBlockVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayCacheDiskWithOptions(request *DeleteGatewayCacheDiskRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayCacheDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheId)) {
		query["CacheId"] = request.CacheId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalFilePath)) {
		query["LocalFilePath"] = request.LocalFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayCacheDisk"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayCacheDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewayCacheDisk(request *DeleteGatewayCacheDiskRequest) (_result *DeleteGatewayCacheDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayCacheDiskResponse{}
	_body, _err := client.DeleteGatewayCacheDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayFileSharesWithOptions(request *DeleteGatewayFileSharesRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayFileSharesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayFileShares"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayFileSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewayFileShares(request *DeleteGatewayFileSharesRequest) (_result *DeleteGatewayFileSharesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayFileSharesResponse{}
	_body, _err := client.DeleteGatewayFileSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayLoggingWithOptions(request *DeleteGatewayLoggingRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayLogging"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewayLogging(request *DeleteGatewayLoggingRequest) (_result *DeleteGatewayLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayLoggingResponse{}
	_body, _err := client.DeleteGatewayLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewaySMBUserWithOptions(request *DeleteGatewaySMBUserRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewaySMBUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewaySMBUser"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewaySMBUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewaySMBUser(request *DeleteGatewaySMBUserRequest) (_result *DeleteGatewaySMBUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewaySMBUserResponse{}
	_body, _err := client.DeleteGatewaySMBUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStorageBundleWithOptions(request *DeleteStorageBundleRequest, runtime *util.RuntimeOptions) (_result *DeleteStorageBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStorageBundle"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStorageBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStorageBundle(request *DeleteStorageBundleRequest) (_result *DeleteStorageBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStorageBundleResponse{}
	_body, _err := client.DeleteStorageBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployCSGClientsWithOptions(tmpReq *DeployCSGClientsRequest, runtime *util.RuntimeOptions) (_result *DeployCSGClientsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeployCSGClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcsInstanceIds)) {
		request.EcsInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcsInstanceIds, tea.String("EcsInstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientRegionId)) {
		query["ClientRegionId"] = request.ClientRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceIdsShrink)) {
		query["EcsInstanceIds"] = request.EcsInstanceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployCSGClients"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployCSGClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployCSGClients(request *DeployCSGClientsRequest) (_result *DeployCSGClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployCSGClientsResponse{}
	_body, _err := client.DeployCSGClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployCacheDiskWithOptions(request *DeployCacheDiskRequest, runtime *util.RuntimeOptions) (_result *DeployCacheDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheConfig)) {
		query["CacheConfig"] = request.CacheConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DiskCategory)) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SizeInGB)) {
		query["SizeInGB"] = request.SizeInGB
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployCacheDisk"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployCacheDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployCacheDisk(request *DeployCacheDiskRequest) (_result *DeployCacheDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployCacheDiskResponse{}
	_body, _err := client.DeployCacheDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployGatewayWithOptions(request *DeployGatewayRequest, runtime *util.RuntimeOptions) (_result *DeployGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayClass)) {
		query["GatewayClass"] = request.GatewayClass
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployGateway(request *DeployGatewayRequest) (_result *DeployGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployGatewayResponse{}
	_body, _err := client.DeployGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountConfigWithOptions(request *DescribeAccountConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountConfig"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountConfig(request *DescribeAccountConfigRequest) (_result *DescribeAccountConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountConfigResponse{}
	_body, _err := client.DescribeAccountConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBlockVolumeSnapshotsWithOptions(request *DescribeBlockVolumeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeBlockVolumeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBlockVolumeSnapshots"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBlockVolumeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBlockVolumeSnapshots(request *DescribeBlockVolumeSnapshotsRequest) (_result *DescribeBlockVolumeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBlockVolumeSnapshotsResponse{}
	_body, _err := client.DescribeBlockVolumeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCSGClientTasksWithOptions(request *DescribeCSGClientTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeCSGClientTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientRegionId)) {
		query["ClientRegionId"] = request.ClientRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCSGClientTasks"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCSGClientTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCSGClientTasks(request *DescribeCSGClientTasksRequest) (_result *DescribeCSGClientTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCSGClientTasksResponse{}
	_body, _err := client.DescribeCSGClientTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCSGClientsWithOptions(request *DescribeCSGClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeCSGClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientRegionId)) {
		query["ClientRegionId"] = request.ClientRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCSGClients"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCSGClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCSGClients(request *DescribeCSGClientsRequest) (_result *DescribeCSGClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCSGClientsResponse{}
	_body, _err := client.DescribeCSGClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDashboardWithOptions(request *DescribeDashboardRequest, runtime *util.RuntimeOptions) (_result *DescribeDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendBucketRegionId)) {
		query["BackendBucketRegionId"] = request.BackendBucketRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDashboard"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDashboard(request *DescribeDashboardRequest) (_result *DescribeDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDashboardResponse{}
	_body, _err := client.DescribeDashboardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpireCachesWithOptions(request *DescribeExpireCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeExpireCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpireCaches"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpireCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpireCaches(request *DescribeExpireCachesRequest) (_result *DescribeExpireCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpireCachesResponse{}
	_body, _err := client.DescribeExpireCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressSyncSharesWithOptions(request *DescribeExpressSyncSharesRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressSyncSharesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpressSyncIds)) {
		query["ExpressSyncIds"] = request.ExpressSyncIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsExternal)) {
		query["IsExternal"] = request.IsExternal
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressSyncShares"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressSyncSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressSyncShares(request *DescribeExpressSyncSharesRequest) (_result *DescribeExpressSyncSharesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressSyncSharesResponse{}
	_body, _err := client.DescribeExpressSyncSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressSyncsWithOptions(request *DescribeExpressSyncsRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressSyncsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketPrefix)) {
		query["BucketPrefix"] = request.BucketPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressSyncs"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressSyncsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressSyncs(request *DescribeExpressSyncsRequest) (_result *DescribeExpressSyncsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressSyncsResponse{}
	_body, _err := client.DescribeExpressSyncsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayWithOptions(request *DescribeGatewayRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGateway(request *DescribeGatewayRequest) (_result *DescribeGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayResponse{}
	_body, _err := client.DescribeGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayADInfoWithOptions(request *DescribeGatewayADInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayADInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayADInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayADInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayADInfo(request *DescribeGatewayADInfoRequest) (_result *DescribeGatewayADInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayADInfoResponse{}
	_body, _err := client.DescribeGatewayADInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayActionsWithOptions(request *DescribeGatewayActionsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayActions"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayActions(request *DescribeGatewayActionsRequest) (_result *DescribeGatewayActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayActionsResponse{}
	_body, _err := client.DescribeGatewayActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayAuthInfoWithOptions(request *DescribeGatewayAuthInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayAuthInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayAuthInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayAuthInfo(request *DescribeGatewayAuthInfoRequest) (_result *DescribeGatewayAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayAuthInfoResponse{}
	_body, _err := client.DescribeGatewayAuthInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayAutoPlansWithOptions(request *DescribeGatewayAutoPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayAutoPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayAutoPlans"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayAutoPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayAutoPlans(request *DescribeGatewayAutoPlansRequest) (_result *DescribeGatewayAutoPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayAutoPlansResponse{}
	_body, _err := client.DescribeGatewayAutoPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayAutoUpgradeConfigurationWithOptions(request *DescribeGatewayAutoUpgradeConfigurationRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayAutoUpgradeConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayAutoUpgradeConfiguration"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayAutoUpgradeConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayAutoUpgradeConfiguration(request *DescribeGatewayAutoUpgradeConfigurationRequest) (_result *DescribeGatewayAutoUpgradeConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayAutoUpgradeConfigurationResponse{}
	_body, _err := client.DescribeGatewayAutoUpgradeConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribeGatewayBlockVolumesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeGatewayBlockVolumesResponse
 */
func (client *Client) DescribeGatewayBlockVolumesWithOptions(request *DescribeGatewayBlockVolumesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayBlockVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.Refresh)) {
		query["Refresh"] = request.Refresh
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayBlockVolumes"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayBlockVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribeGatewayBlockVolumesRequest
 * @return DescribeGatewayBlockVolumesResponse
 */
func (client *Client) DescribeGatewayBlockVolumes(request *DescribeGatewayBlockVolumesRequest) (_result *DescribeGatewayBlockVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayBlockVolumesResponse{}
	_body, _err := client.DescribeGatewayBlockVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayBucketCachesWithOptions(request *DescribeGatewayBucketCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayBucketCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayBucketCaches"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayBucketCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayBucketCaches(request *DescribeGatewayBucketCachesRequest) (_result *DescribeGatewayBucketCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayBucketCachesResponse{}
	_body, _err := client.DescribeGatewayBucketCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayCachesWithOptions(request *DescribeGatewayCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayCaches"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayCaches(request *DescribeGatewayCachesRequest) (_result *DescribeGatewayCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayCachesResponse{}
	_body, _err := client.DescribeGatewayCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayCapacityLimitWithOptions(request *DescribeGatewayCapacityLimitRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayCapacityLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SizeInGB)) {
		query["SizeInGB"] = request.SizeInGB
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayCapacityLimit"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayCapacityLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayCapacityLimit(request *DescribeGatewayCapacityLimitRequest) (_result *DescribeGatewayCapacityLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayCapacityLimitResponse{}
	_body, _err := client.DescribeGatewayCapacityLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayCategoriesWithOptions(request *DescribeGatewayCategoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayLocation)) {
		query["GatewayLocation"] = request.GatewayLocation
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayCategories"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayCategories(request *DescribeGatewayCategoriesRequest) (_result *DescribeGatewayCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayCategoriesResponse{}
	_body, _err := client.DescribeGatewayCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayClassesWithOptions(request *DescribeGatewayClassesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayClassesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayClasses"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayClassesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayClasses(request *DescribeGatewayClassesRequest) (_result *DescribeGatewayClassesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayClassesResponse{}
	_body, _err := client.DescribeGatewayClassesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayCredentialWithOptions(request *DescribeGatewayCredentialRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayCredential"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayCredential(request *DescribeGatewayCredentialRequest) (_result *DescribeGatewayCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayCredentialResponse{}
	_body, _err := client.DescribeGatewayCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayDNSWithOptions(request *DescribeGatewayDNSRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayDNSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayDNS"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayDNS(request *DescribeGatewayDNSRequest) (_result *DescribeGatewayDNSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayDNSResponse{}
	_body, _err := client.DescribeGatewayDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayFileSharesWithOptions(request *DescribeGatewayFileSharesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayFileSharesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.Refresh)) {
		query["Refresh"] = request.Refresh
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayFileShares"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayFileSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayFileShares(request *DescribeGatewayFileSharesRequest) (_result *DescribeGatewayFileSharesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayFileSharesResponse{}
	_body, _err := client.DescribeGatewayFileSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayFileStatusWithOptions(request *DescribeGatewayFileStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayFileStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayFileStatus"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayFileStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayFileStatus(request *DescribeGatewayFileStatusRequest) (_result *DescribeGatewayFileStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayFileStatusResponse{}
	_body, _err := client.DescribeGatewayFileStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayImagesWithOptions(request *DescribeGatewayImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayImages"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayImages(request *DescribeGatewayImagesRequest) (_result *DescribeGatewayImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayImagesResponse{}
	_body, _err := client.DescribeGatewayImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayInfoWithOptions(request *DescribeGatewayInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayInfo(request *DescribeGatewayInfoRequest) (_result *DescribeGatewayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayInfoResponse{}
	_body, _err := client.DescribeGatewayInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayLDAPInfoWithOptions(request *DescribeGatewayLDAPInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayLDAPInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayLDAPInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayLDAPInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayLDAPInfo(request *DescribeGatewayLDAPInfoRequest) (_result *DescribeGatewayLDAPInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayLDAPInfoResponse{}
	_body, _err := client.DescribeGatewayLDAPInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayLocationsWithOptions(request *DescribeGatewayLocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayLocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayLocations"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayLocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayLocations(request *DescribeGatewayLocationsRequest) (_result *DescribeGatewayLocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayLocationsResponse{}
	_body, _err := client.DescribeGatewayLocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayLoggingWithOptions(request *DescribeGatewayLoggingRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayLogging"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayLogging(request *DescribeGatewayLoggingRequest) (_result *DescribeGatewayLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayLoggingResponse{}
	_body, _err := client.DescribeGatewayLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayLogsWithOptions(request *DescribeGatewayLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.LogFilePath)) {
		query["LogFilePath"] = request.LogFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayLogs"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayLogs(request *DescribeGatewayLogsRequest) (_result *DescribeGatewayLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayLogsResponse{}
	_body, _err := client.DescribeGatewayLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayModificationClassesWithOptions(request *DescribeGatewayModificationClassesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayModificationClassesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayModificationClasses"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayModificationClassesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayModificationClasses(request *DescribeGatewayModificationClassesRequest) (_result *DescribeGatewayModificationClassesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayModificationClassesResponse{}
	_body, _err := client.DescribeGatewayModificationClassesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayNFSClientsWithOptions(request *DescribeGatewayNFSClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayNFSClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayNFSClients"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayNFSClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayNFSClients(request *DescribeGatewayNFSClientsRequest) (_result *DescribeGatewayNFSClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayNFSClientsResponse{}
	_body, _err := client.DescribeGatewayNFSClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewaySMBUsersWithOptions(request *DescribeGatewaySMBUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewaySMBUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewaySMBUsers"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewaySMBUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewaySMBUsers(request *DescribeGatewaySMBUsersRequest) (_result *DescribeGatewaySMBUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewaySMBUsersResponse{}
	_body, _err := client.DescribeGatewaySMBUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayStatisticsWithOptions(request *DescribeGatewayStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayCategory)) {
		query["GatewayCategory"] = request.GatewayCategory
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayLocation)) {
		query["GatewayLocation"] = request.GatewayLocation
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAccountId)) {
		query["TargetAccountId"] = request.TargetAccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayStatistics"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayStatistics(request *DescribeGatewayStatisticsRequest) (_result *DescribeGatewayStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayStatisticsResponse{}
	_body, _err := client.DescribeGatewayStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayStockWithOptions(request *DescribeGatewayStockRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayStockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayRegionId)) {
		query["GatewayRegionId"] = request.GatewayRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayStock"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayStockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayStock(request *DescribeGatewayStockRequest) (_result *DescribeGatewayStockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayStockResponse{}
	_body, _err := client.DescribeGatewayStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewayTypesWithOptions(request *DescribeGatewayTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewayTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayLocation)) {
		query["GatewayLocation"] = request.GatewayLocation
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewayTypes"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewayTypes(request *DescribeGatewayTypesRequest) (_result *DescribeGatewayTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewayTypesResponse{}
	_body, _err := client.DescribeGatewayTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewaysWithOptions(request *DescribeGatewaysRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewaysResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGateways"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGateways(request *DescribeGatewaysRequest) (_result *DescribeGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewaysResponse{}
	_body, _err := client.DescribeGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewaysForCmsWithOptions(request *DescribeGatewaysForCmsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewaysForCmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayRegionId)) {
		query["GatewayRegionId"] = request.GatewayRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewaysForCms"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewaysForCmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewaysForCms(request *DescribeGatewaysForCmsRequest) (_result *DescribeGatewaysForCmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewaysForCmsResponse{}
	_body, _err := client.DescribeGatewaysForCmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewaysTagsWithOptions(request *DescribeGatewaysTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewaysTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayIds)) {
		query["GatewayIds"] = request.GatewayIds
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.TagCategory)) {
		query["TagCategory"] = request.TagCategory
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewaysTags"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewaysTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewaysTags(request *DescribeGatewaysTagsRequest) (_result *DescribeGatewaysTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewaysTagsResponse{}
	_body, _err := client.DescribeGatewaysTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKmsKeyWithOptions(request *DescribeKmsKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeKmsKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.KmsKey)) {
		query["KmsKey"] = request.KmsKey
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKmsKey"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeKmsKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKmsKey(request *DescribeKmsKeyRequest) (_result *DescribeKmsKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKmsKeyResponse{}
	_body, _err := client.DescribeKmsKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMqttConfigWithOptions(request *DescribeMqttConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeMqttConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMqttConfig"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMqttConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMqttConfig(request *DescribeMqttConfigRequest) (_result *DescribeMqttConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMqttConfigResponse{}
	_body, _err := client.DescribeMqttConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssBucketInfoWithOptions(request *DescribeOssBucketInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeOssBucketInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketEndpoint)) {
		query["BucketEndpoint"] = request.BucketEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssBucketInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssBucketInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssBucketInfo(request *DescribeOssBucketInfoRequest) (_result *DescribeOssBucketInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssBucketInfoResponse{}
	_body, _err := client.DescribeOssBucketInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssBucketsWithOptions(request *DescribeOssBucketsRequest, runtime *util.RuntimeOptions) (_result *DescribeOssBucketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketEndpoint)) {
		query["BucketEndpoint"] = request.BucketEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssBuckets"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssBucketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssBuckets(request *DescribeOssBucketsRequest) (_result *DescribeOssBucketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssBucketsResponse{}
	_body, _err := client.DescribeOssBucketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePayAsYouGoPriceWithOptions(request *DescribePayAsYouGoPriceRequest, runtime *util.RuntimeOptions) (_result *DescribePayAsYouGoPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayClass)) {
		query["GatewayClass"] = request.GatewayClass
	}

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
		Action:      tea.String("DescribePayAsYouGoPrice"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePayAsYouGoPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePayAsYouGoPrice(request *DescribePayAsYouGoPriceRequest) (_result *DescribePayAsYouGoPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePayAsYouGoPriceResponse{}
	_body, _err := client.DescribePayAsYouGoPriceWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeSharesBucketInfoForExpressSyncWithOptions(request *DescribeSharesBucketInfoForExpressSyncRequest, runtime *util.RuntimeOptions) (_result *DescribeSharesBucketInfoForExpressSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketRegion)) {
		query["BucketRegion"] = request.BucketRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSharesBucketInfoForExpressSync"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSharesBucketInfoForExpressSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSharesBucketInfoForExpressSync(request *DescribeSharesBucketInfoForExpressSyncRequest) (_result *DescribeSharesBucketInfoForExpressSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSharesBucketInfoForExpressSyncResponse{}
	_body, _err := client.DescribeSharesBucketInfoForExpressSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStorageBundleWithOptions(request *DescribeStorageBundleRequest, runtime *util.RuntimeOptions) (_result *DescribeStorageBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStorageBundle"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStorageBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStorageBundle(request *DescribeStorageBundleRequest) (_result *DescribeStorageBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStorageBundleResponse{}
	_body, _err := client.DescribeStorageBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStorageBundlesWithOptions(request *DescribeStorageBundlesRequest, runtime *util.RuntimeOptions) (_result *DescribeStorageBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendBucketRegionId)) {
		query["BackendBucketRegionId"] = request.BackendBucketRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStorageBundles"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStorageBundlesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStorageBundles(request *DescribeStorageBundlesRequest) (_result *DescribeStorageBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStorageBundlesResponse{}
	_body, _err := client.DescribeStorageBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionPriceWithOptions(request *DescribeSubscriptionPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheCloudEfficiencySize)) {
		query["CacheCloudEfficiencySize"] = request.CacheCloudEfficiencySize
	}

	if !tea.BoolValue(util.IsUnset(request.CacheESSDPl1Size)) {
		query["CacheESSDPl1Size"] = request.CacheESSDPl1Size
	}

	if !tea.BoolValue(util.IsUnset(request.CacheSSDSize)) {
		query["CacheSSDSize"] = request.CacheSSDSize
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayClass)) {
		query["GatewayClass"] = request.GatewayClass
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodQuantity)) {
		query["PeriodQuantity"] = request.PeriodQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

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
		Action:      tea.String("DescribeSubscriptionPrice"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionPrice(request *DescribeSubscriptionPriceRequest) (_result *DescribeSubscriptionPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionPriceResponse{}
	_body, _err := client.DescribeSubscriptionPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTasks"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserBusinessStatusWithOptions(request *DescribeUserBusinessStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserBusinessStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserBusinessStatus"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserBusinessStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserBusinessStatus(request *DescribeUserBusinessStatusRequest) (_result *DescribeUserBusinessStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserBusinessStatusResponse{}
	_body, _err := client.DescribeUserBusinessStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *util.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVSwitches"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcsWithOptions(request *DescribeVpcsRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcs"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcs(request *DescribeVpcsRequest) (_result *DescribeVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.DescribeVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
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
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableGatewayLoggingWithOptions(request *DisableGatewayLoggingRequest, runtime *util.RuntimeOptions) (_result *DisableGatewayLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableGatewayLogging"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableGatewayLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableGatewayLogging(request *DisableGatewayLoggingRequest) (_result *DisableGatewayLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableGatewayLoggingResponse{}
	_body, _err := client.DisableGatewayLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableGatewayNFSVersionWithOptions(request *DisableGatewayNFSVersionRequest, runtime *util.RuntimeOptions) (_result *DisableGatewayNFSVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.NFSVersion)) {
		query["NFSVersion"] = request.NFSVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableGatewayNFSVersion"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableGatewayNFSVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableGatewayNFSVersion(request *DisableGatewayNFSVersionRequest) (_result *DisableGatewayNFSVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableGatewayNFSVersionResponse{}
	_body, _err := client.DisableGatewayNFSVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableGatewayIpv6WithOptions(request *EnableGatewayIpv6Request, runtime *util.RuntimeOptions) (_result *EnableGatewayIpv6Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableGatewayIpv6"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableGatewayIpv6Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableGatewayIpv6(request *EnableGatewayIpv6Request) (_result *EnableGatewayIpv6Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableGatewayIpv6Response{}
	_body, _err := client.EnableGatewayIpv6WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableGatewayLoggingWithOptions(request *EnableGatewayLoggingRequest, runtime *util.RuntimeOptions) (_result *EnableGatewayLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableGatewayLogging"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableGatewayLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableGatewayLogging(request *EnableGatewayLoggingRequest) (_result *EnableGatewayLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableGatewayLoggingResponse{}
	_body, _err := client.EnableGatewayLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpandCacheDiskWithOptions(request *ExpandCacheDiskRequest, runtime *util.RuntimeOptions) (_result *ExpandCacheDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalFilePath)) {
		query["LocalFilePath"] = request.LocalFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.NewSizeInGB)) {
		query["NewSizeInGB"] = request.NewSizeInGB
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpandCacheDisk"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpandCacheDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpandCacheDisk(request *ExpandCacheDiskRequest) (_result *ExpandCacheDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExpandCacheDiskResponse{}
	_body, _err := client.ExpandCacheDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpandGatewayFileShareWithOptions(request *ExpandGatewayFileShareRequest, runtime *util.RuntimeOptions) (_result *ExpandGatewayFileShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpandGatewayFileShare"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpandGatewayFileShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpandGatewayFileShare(request *ExpandGatewayFileShareRequest) (_result *ExpandGatewayFileShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExpandGatewayFileShareResponse{}
	_body, _err := client.ExpandGatewayFileShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpandGatewayNetworkBandwidthWithOptions(request *ExpandGatewayNetworkBandwidthRequest, runtime *util.RuntimeOptions) (_result *ExpandGatewayNetworkBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.NewNetworkBandwidth)) {
		query["NewNetworkBandwidth"] = request.NewNetworkBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpandGatewayNetworkBandwidth"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpandGatewayNetworkBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpandGatewayNetworkBandwidth(request *ExpandGatewayNetworkBandwidthRequest) (_result *ExpandGatewayNetworkBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExpandGatewayNetworkBandwidthResponse{}
	_body, _err := client.ExpandGatewayNetworkBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateGatewayTokenWithOptions(request *GenerateGatewayTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateGatewayTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateGatewayToken"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateGatewayTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateGatewayToken(request *GenerateGatewayTokenRequest) (_result *GenerateGatewayTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateGatewayTokenResponse{}
	_body, _err := client.GenerateGatewayTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateMqttTokenWithOptions(request *GenerateMqttTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateMqttTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateMqttToken"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateMqttTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateMqttToken(request *GenerateMqttTokenRequest) (_result *GenerateMqttTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateMqttTokenResponse{}
	_body, _err := client.GenerateMqttTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateStsTokenWithOptions(request *GenerateStsTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateStsTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireInSeconds)) {
		query["ExpireInSeconds"] = request.ExpireInSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TokenType)) {
		query["TokenType"] = request.TokenType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateStsToken"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateStsTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateStsToken(request *GenerateStsTokenRequest) (_result *GenerateStsTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateStsTokenResponse{}
	_body, _err := client.GenerateStsTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleGatewayAutoPlanWithOptions(request *HandleGatewayAutoPlanRequest, runtime *util.RuntimeOptions) (_result *HandleGatewayAutoPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cancel)) {
		query["Cancel"] = request.Cancel
	}

	if !tea.BoolValue(util.IsUnset(request.DelayHours)) {
		query["DelayHours"] = request.DelayHours
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HandleGatewayAutoPlan"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HandleGatewayAutoPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleGatewayAutoPlan(request *HandleGatewayAutoPlanRequest) (_result *HandleGatewayAutoPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleGatewayAutoPlanResponse{}
	_body, _err := client.HandleGatewayAutoPlanWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyGatewayWithOptions(request *ModifyGatewayRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGateway(request *ModifyGatewayRequest) (_result *ModifyGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayResponse{}
	_body, _err := client.ModifyGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGatewayBlockVolumeWithOptions(request *ModifyGatewayBlockVolumeRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayBlockVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheConfig)) {
		query["CacheConfig"] = request.CacheConfig
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGatewayBlockVolume"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayBlockVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGatewayBlockVolume(request *ModifyGatewayBlockVolumeRequest) (_result *ModifyGatewayBlockVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayBlockVolumeResponse{}
	_body, _err := client.ModifyGatewayBlockVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGatewayClassWithOptions(request *ModifyGatewayClassRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayClass)) {
		query["GatewayClass"] = request.GatewayClass
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGatewayClass"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGatewayClass(request *ModifyGatewayClassRequest) (_result *ModifyGatewayClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayClassResponse{}
	_body, _err := client.ModifyGatewayClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGatewayFileShareWithOptions(request *ModifyGatewayFileShareRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayFileShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CacheConfig)) {
		query["CacheConfig"] = request.CacheConfig
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGatewayFileShare"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayFileShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGatewayFileShare(request *ModifyGatewayFileShareRequest) (_result *ModifyGatewayFileShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayFileShareResponse{}
	_body, _err := client.ModifyGatewayFileShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGatewayFileShareWatermarkWithOptions(request *ModifyGatewayFileShareWatermarkRequest, runtime *util.RuntimeOptions) (_result *ModifyGatewayFileShareWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Watermark)) {
		query["Watermark"] = request.Watermark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGatewayFileShareWatermark"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGatewayFileShareWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGatewayFileShareWatermark(request *ModifyGatewayFileShareWatermarkRequest) (_result *ModifyGatewayFileShareWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGatewayFileShareWatermarkResponse{}
	_body, _err := client.ModifyGatewayFileShareWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyStorageBundleWithOptions(request *ModifyStorageBundleRequest, runtime *util.RuntimeOptions) (_result *ModifyStorageBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyStorageBundle"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyStorageBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyStorageBundle(request *ModifyStorageBundleRequest) (_result *ModifyStorageBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyStorageBundleResponse{}
	_body, _err := client.ModifyStorageBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenSgwServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenSgwServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenSgwService"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenSgwServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenSgwService() (_result *OpenSgwServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenSgwServiceResponse{}
	_body, _err := client.OpenSgwServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateGatewayWithOptions(request *OperateGatewayRequest, runtime *util.RuntimeOptions) (_result *OperateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateAction)) {
		query["OperateAction"] = request.OperateAction
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateGateway(request *OperateGatewayRequest) (_result *OperateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateGatewayResponse{}
	_body, _err := client.OperateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseServiceWithOptions(request *ReleaseServiceRequest, runtime *util.RuntimeOptions) (_result *ReleaseServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseService"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseService(request *ReleaseServiceRequest) (_result *ReleaseServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseServiceResponse{}
	_body, _err := client.ReleaseServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSharesFromExpressSyncWithOptions(request *RemoveSharesFromExpressSyncRequest, runtime *util.RuntimeOptions) (_result *RemoveSharesFromExpressSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpressSyncId)) {
		query["ExpressSyncId"] = request.ExpressSyncId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayShares)) {
		query["GatewayShares"] = request.GatewayShares
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSharesFromExpressSync"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSharesFromExpressSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSharesFromExpressSync(request *RemoveSharesFromExpressSyncRequest) (_result *RemoveSharesFromExpressSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSharesFromExpressSyncResponse{}
	_body, _err := client.RemoveSharesFromExpressSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTagsFromGatewayWithOptions(request *RemoveTagsFromGatewayRequest, runtime *util.RuntimeOptions) (_result *RemoveTagsFromGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTagsFromGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTagsFromGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTagsFromGateway(request *RemoveTagsFromGatewayRequest) (_result *RemoveTagsFromGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTagsFromGatewayResponse{}
	_body, _err := client.RemoveTagsFromGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportBlockVolumesWithOptions(request *ReportBlockVolumesRequest, runtime *util.RuntimeOptions) (_result *ReportBlockVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Info)) {
		query["Info"] = request.Info
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportBlockVolumes"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportBlockVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportBlockVolumes(request *ReportBlockVolumesRequest) (_result *ReportBlockVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportBlockVolumesResponse{}
	_body, _err := client.ReportBlockVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportFileSharesWithOptions(request *ReportFileSharesRequest, runtime *util.RuntimeOptions) (_result *ReportFileSharesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Info)) {
		query["Info"] = request.Info
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportFileShares"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportFileSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportFileShares(request *ReportFileSharesRequest) (_result *ReportFileSharesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportFileSharesResponse{}
	_body, _err := client.ReportFileSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportGatewayInfoWithOptions(request *ReportGatewayInfoRequest, runtime *util.RuntimeOptions) (_result *ReportGatewayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayStatus)) {
		query["GatewayStatus"] = request.GatewayStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Info)) {
		query["Info"] = request.Info
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportGatewayInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportGatewayInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportGatewayInfo(request *ReportGatewayInfoRequest) (_result *ReportGatewayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportGatewayInfoResponse{}
	_body, _err := client.ReportGatewayInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportGatewayUsageWithOptions(request *ReportGatewayUsageRequest, runtime *util.RuntimeOptions) (_result *ReportGatewayUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientUUID)) {
		query["ClientUUID"] = request.ClientUUID
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Usage)) {
		query["Usage"] = request.Usage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportGatewayUsage"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportGatewayUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportGatewayUsage(request *ReportGatewayUsageRequest) (_result *ReportGatewayUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportGatewayUsageResponse{}
	_body, _err := client.ReportGatewayUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetGatewayPasswordWithOptions(request *ResetGatewayPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetGatewayPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetGatewayPassword"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetGatewayPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetGatewayPassword(request *ResetGatewayPasswordRequest) (_result *ResetGatewayPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetGatewayPasswordResponse{}
	_body, _err := client.ResetGatewayPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartFileSharesWithOptions(request *RestartFileSharesRequest, runtime *util.RuntimeOptions) (_result *RestartFileSharesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ShareProtocol)) {
		query["ShareProtocol"] = request.ShareProtocol
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartFileShares"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartFileSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartFileShares(request *RestartFileSharesRequest) (_result *RestartFileSharesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartFileSharesResponse{}
	_body, _err := client.RestartFileSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGatewayADInfoWithOptions(request *SetGatewayADInfoRequest, runtime *util.RuntimeOptions) (_result *SetGatewayADInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnabled)) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIp)) {
		query["ServerIp"] = request.ServerIp
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGatewayADInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGatewayADInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGatewayADInfo(request *SetGatewayADInfoRequest) (_result *SetGatewayADInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGatewayADInfoResponse{}
	_body, _err := client.SetGatewayADInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGatewayAutoUpgradeConfigurationWithOptions(request *SetGatewayAutoUpgradeConfigurationRequest, runtime *util.RuntimeOptions) (_result *SetGatewayAutoUpgradeConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoUpgradeEndHour)) {
		query["AutoUpgradeEndHour"] = request.AutoUpgradeEndHour
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUpgradeStartHour)) {
		query["AutoUpgradeStartHour"] = request.AutoUpgradeStartHour
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoUpgrade)) {
		query["IsAutoUpgrade"] = request.IsAutoUpgrade
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGatewayAutoUpgradeConfiguration"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGatewayAutoUpgradeConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGatewayAutoUpgradeConfiguration(request *SetGatewayAutoUpgradeConfigurationRequest) (_result *SetGatewayAutoUpgradeConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGatewayAutoUpgradeConfigurationResponse{}
	_body, _err := client.SetGatewayAutoUpgradeConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGatewayDNSWithOptions(request *SetGatewayDNSRequest, runtime *util.RuntimeOptions) (_result *SetGatewayDNSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DnsServer)) {
		query["DnsServer"] = request.DnsServer
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGatewayDNS"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGatewayDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGatewayDNS(request *SetGatewayDNSRequest) (_result *SetGatewayDNSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGatewayDNSResponse{}
	_body, _err := client.SetGatewayDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGatewayLDAPInfoWithOptions(request *SetGatewayLDAPInfoRequest, runtime *util.RuntimeOptions) (_result *SetGatewayLDAPInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseDN)) {
		query["BaseDN"] = request.BaseDN
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnabled)) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.IsTls)) {
		query["IsTls"] = request.IsTls
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RootDN)) {
		query["RootDN"] = request.RootDN
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIp)) {
		query["ServerIp"] = request.ServerIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGatewayLDAPInfo"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGatewayLDAPInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGatewayLDAPInfo(request *SetGatewayLDAPInfoRequest) (_result *SetGatewayLDAPInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGatewayLDAPInfoResponse{}
	_body, _err := client.SetGatewayLDAPInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchCSGClientsReverseSyncConfigurationWithOptions(tmpReq *SwitchCSGClientsReverseSyncConfigurationRequest, runtime *util.RuntimeOptions) (_result *SwitchCSGClientsReverseSyncConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SwitchCSGClientsReverseSyncConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClientIds)) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, tea.String("ClientIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdsShrink)) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientRegionId)) {
		query["ClientRegionId"] = request.ClientRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.IsReverseSync)) {
		query["IsReverseSync"] = request.IsReverseSync
	}

	if !tea.BoolValue(util.IsUnset(request.ReverseSyncInternalSecond)) {
		query["ReverseSyncInternalSecond"] = request.ReverseSyncInternalSecond
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchCSGClientsReverseSyncConfiguration"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchCSGClientsReverseSyncConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchCSGClientsReverseSyncConfiguration(request *SwitchCSGClientsReverseSyncConfigurationRequest) (_result *SwitchCSGClientsReverseSyncConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchCSGClientsReverseSyncConfigurationResponse{}
	_body, _err := client.SwitchCSGClientsReverseSyncConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchGatewayExpirationPolicyWithOptions(request *SwitchGatewayExpirationPolicyRequest, runtime *util.RuntimeOptions) (_result *SwitchGatewayExpirationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchGatewayExpirationPolicy"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchGatewayExpirationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchGatewayExpirationPolicy(request *SwitchGatewayExpirationPolicyRequest) (_result *SwitchGatewayExpirationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchGatewayExpirationPolicyResponse{}
	_body, _err := client.SwitchGatewayExpirationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchToSubscriptionWithOptions(request *SwitchToSubscriptionRequest, runtime *util.RuntimeOptions) (_result *SwitchToSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchToSubscription"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchToSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchToSubscription(request *SwitchToSubscriptionRequest) (_result *SwitchToSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchToSubscriptionResponse{}
	_body, _err := client.SwitchToSubscriptionWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) TriggerGatewayRemoteSyncWithOptions(request *TriggerGatewayRemoteSyncRequest, runtime *util.RuntimeOptions) (_result *TriggerGatewayRemoteSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerGatewayRemoteSync"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerGatewayRemoteSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerGatewayRemoteSync(request *TriggerGatewayRemoteSyncRequest) (_result *TriggerGatewayRemoteSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerGatewayRemoteSyncResponse{}
	_body, _err := client.TriggerGatewayRemoteSyncWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateGatewayBlockVolumeWithOptions(request *UpdateGatewayBlockVolumeRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayBlockVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChapEnabled)) {
		query["ChapEnabled"] = request.ChapEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ChapInPassword)) {
		query["ChapInPassword"] = request.ChapInPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ChapInUser)) {
		query["ChapInUser"] = request.ChapInUser
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayBlockVolume"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayBlockVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGatewayBlockVolume(request *UpdateGatewayBlockVolumeRequest) (_result *UpdateGatewayBlockVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayBlockVolumeResponse{}
	_body, _err := client.UpdateGatewayBlockVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayFileShareWithOptions(request *UpdateGatewayFileShareRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayFileShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessBasedEnumeration)) {
		query["AccessBasedEnumeration"] = request.AccessBasedEnumeration
	}

	if !tea.BoolValue(util.IsUnset(request.BackendLimit)) {
		query["BackendLimit"] = request.BackendLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Browsable)) {
		query["Browsable"] = request.Browsable
	}

	if !tea.BoolValue(util.IsUnset(request.BypassCacheRead)) {
		query["BypassCacheRead"] = request.BypassCacheRead
	}

	if !tea.BoolValue(util.IsUnset(request.CacheMode)) {
		query["CacheMode"] = request.CacheMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSideCmk)) {
		query["ClientSideCmk"] = request.ClientSideCmk
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSideEncryption)) {
		query["ClientSideEncryption"] = request.ClientSideEncryption
	}

	if !tea.BoolValue(util.IsUnset(request.DirectIO)) {
		query["DirectIO"] = request.DirectIO
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadLimit)) {
		query["DownloadLimit"] = request.DownloadLimit
	}

	if !tea.BoolValue(util.IsUnset(request.FastReclaim)) {
		query["FastReclaim"] = request.FastReclaim
	}

	if !tea.BoolValue(util.IsUnset(request.FrontendLimit)) {
		query["FrontendLimit"] = request.FrontendLimit
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreDelete)) {
		query["IgnoreDelete"] = request.IgnoreDelete
	}

	if !tea.BoolValue(util.IsUnset(request.InPlace)) {
		query["InPlace"] = request.InPlace
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.KmsRotatePeriod)) {
		query["KmsRotatePeriod"] = request.KmsRotatePeriod
	}

	if !tea.BoolValue(util.IsUnset(request.LagPeriod)) {
		query["LagPeriod"] = request.LagPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NfsV4Optimization)) {
		query["NfsV4Optimization"] = request.NfsV4Optimization
	}

	if !tea.BoolValue(util.IsUnset(request.PollingInterval)) {
		query["PollingInterval"] = request.PollingInterval
	}

	if !tea.BoolValue(util.IsUnset(request.ReadOnlyClientList)) {
		query["ReadOnlyClientList"] = request.ReadOnlyClientList
	}

	if !tea.BoolValue(util.IsUnset(request.ReadOnlyUserList)) {
		query["ReadOnlyUserList"] = request.ReadOnlyUserList
	}

	if !tea.BoolValue(util.IsUnset(request.ReadWriteClientList)) {
		query["ReadWriteClientList"] = request.ReadWriteClientList
	}

	if !tea.BoolValue(util.IsUnset(request.ReadWriteUserList)) {
		query["ReadWriteUserList"] = request.ReadWriteUserList
	}

	if !tea.BoolValue(util.IsUnset(request.RemoteSync)) {
		query["RemoteSync"] = request.RemoteSync
	}

	if !tea.BoolValue(util.IsUnset(request.RemoteSyncDownload)) {
		query["RemoteSyncDownload"] = request.RemoteSyncDownload
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerSideCmk)) {
		query["ServerSideCmk"] = request.ServerSideCmk
	}

	if !tea.BoolValue(util.IsUnset(request.ServerSideEncryption)) {
		query["ServerSideEncryption"] = request.ServerSideEncryption
	}

	if !tea.BoolValue(util.IsUnset(request.Squash)) {
		query["Squash"] = request.Squash
	}

	if !tea.BoolValue(util.IsUnset(request.TransferAcceleration)) {
		query["TransferAcceleration"] = request.TransferAcceleration
	}

	if !tea.BoolValue(util.IsUnset(request.WindowsAcl)) {
		query["WindowsAcl"] = request.WindowsAcl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayFileShare"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayFileShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGatewayFileShare(request *UpdateGatewayFileShareRequest) (_result *UpdateGatewayFileShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayFileShareResponse{}
	_body, _err := client.UpdateGatewayFileShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeGatewayWithOptions(request *UpgradeGatewayRequest, runtime *util.RuntimeOptions) (_result *UpgradeGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeGateway"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeGateway(request *UpgradeGatewayRequest) (_result *UpgradeGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeGatewayResponse{}
	_body, _err := client.UpgradeGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadCSGClientLogWithOptions(request *UploadCSGClientLogRequest, runtime *util.RuntimeOptions) (_result *UploadCSGClientLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientRegionId)) {
		query["ClientRegionId"] = request.ClientRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadCSGClientLog"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadCSGClientLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadCSGClientLog(request *UploadCSGClientLogRequest) (_result *UploadCSGClientLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadCSGClientLogResponse{}
	_body, _err := client.UploadCSGClientLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadGatewayLogWithOptions(request *UploadGatewayLogRequest, runtime *util.RuntimeOptions) (_result *UploadGatewayLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadGatewayLog"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadGatewayLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadGatewayLog(request *UploadGatewayLogRequest) (_result *UploadGatewayLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadGatewayLogResponse{}
	_body, _err := client.UploadGatewayLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateExpressSyncConfigWithOptions(request *ValidateExpressSyncConfigRequest, runtime *util.RuntimeOptions) (_result *ValidateExpressSyncConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketPrefix)) {
		query["BucketPrefix"] = request.BucketPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.BucketRegion)) {
		query["BucketRegion"] = request.BucketRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateExpressSyncConfig"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateExpressSyncConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateExpressSyncConfig(request *ValidateExpressSyncConfigRequest) (_result *ValidateExpressSyncConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateExpressSyncConfigResponse{}
	_body, _err := client.ValidateExpressSyncConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateGatewayNameWithOptions(request *ValidateGatewayNameRequest, runtime *util.RuntimeOptions) (_result *ValidateGatewayNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StorageBundleId)) {
		query["StorageBundleId"] = request.StorageBundleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateGatewayName"),
		Version:     tea.String("2018-05-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateGatewayNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateGatewayName(request *ValidateGatewayNameRequest) (_result *ValidateGatewayNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateGatewayNameResponse{}
	_body, _err := client.ValidateGatewayNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

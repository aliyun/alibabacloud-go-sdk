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

type ApplyDataModelConfigInfoRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	DataModelCode *string `json:"DataModelCode,omitempty" xml:"DataModelCode,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ApplyDataModelConfigInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataModelConfigInfoRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataModelConfigInfoRequest) SetApiVersion(v string) *ApplyDataModelConfigInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *ApplyDataModelConfigInfoRequest) SetConfiguration(v string) *ApplyDataModelConfigInfoRequest {
	s.Configuration = &v
	return s
}

func (s *ApplyDataModelConfigInfoRequest) SetDataModelCode(v string) *ApplyDataModelConfigInfoRequest {
	s.DataModelCode = &v
	return s
}

func (s *ApplyDataModelConfigInfoRequest) SetProductKey(v string) *ApplyDataModelConfigInfoRequest {
	s.ProductKey = &v
	return s
}

type ApplyDataModelConfigInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyDataModelConfigInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataModelConfigInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyDataModelConfigInfoResponseBody) SetCode(v int32) *ApplyDataModelConfigInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyDataModelConfigInfoResponseBody) SetData(v string) *ApplyDataModelConfigInfoResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyDataModelConfigInfoResponseBody) SetMessage(v string) *ApplyDataModelConfigInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyDataModelConfigInfoResponseBody) SetRequestId(v string) *ApplyDataModelConfigInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyDataModelConfigInfoResponseBody) SetSuccess(v bool) *ApplyDataModelConfigInfoResponseBody {
	s.Success = &v
	return s
}

type ApplyDataModelConfigInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyDataModelConfigInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyDataModelConfigInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataModelConfigInfoResponse) GoString() string {
	return s.String()
}

func (s *ApplyDataModelConfigInfoResponse) SetHeaders(v map[string]*string) *ApplyDataModelConfigInfoResponse {
	s.Headers = v
	return s
}

func (s *ApplyDataModelConfigInfoResponse) SetStatusCode(v int32) *ApplyDataModelConfigInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyDataModelConfigInfoResponse) SetBody(v *ApplyDataModelConfigInfoResponseBody) *ApplyDataModelConfigInfoResponse {
	s.Body = v
	return s
}

type AttachDataRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AttachDataRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDataRequest) GoString() string {
	return s.String()
}

func (s *AttachDataRequest) SetApiVersion(v string) *AttachDataRequest {
	s.ApiVersion = &v
	return s
}

func (s *AttachDataRequest) SetBusinessId(v string) *AttachDataRequest {
	s.BusinessId = &v
	return s
}

func (s *AttachDataRequest) SetKey(v string) *AttachDataRequest {
	s.Key = &v
	return s
}

func (s *AttachDataRequest) SetProductKey(v string) *AttachDataRequest {
	s.ProductKey = &v
	return s
}

func (s *AttachDataRequest) SetValue(v string) *AttachDataRequest {
	s.Value = &v
	return s
}

type AttachDataResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDataResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDataResponseBody) SetCode(v int32) *AttachDataResponseBody {
	s.Code = &v
	return s
}

func (s *AttachDataResponseBody) SetData(v string) *AttachDataResponseBody {
	s.Data = &v
	return s
}

func (s *AttachDataResponseBody) SetMessage(v string) *AttachDataResponseBody {
	s.Message = &v
	return s
}

func (s *AttachDataResponseBody) SetRequestId(v string) *AttachDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDataResponseBody) SetSuccess(v bool) *AttachDataResponseBody {
	s.Success = &v
	return s
}

type AttachDataResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDataResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDataResponse) GoString() string {
	return s.String()
}

func (s *AttachDataResponse) SetHeaders(v map[string]*string) *AttachDataResponse {
	s.Headers = v
	return s
}

func (s *AttachDataResponse) SetStatusCode(v int32) *AttachDataResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDataResponse) SetBody(v *AttachDataResponseBody) *AttachDataResponse {
	s.Body = v
	return s
}

type AttachDataWithSignatureRequest struct {
	ApiVersion           *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BusinessId           *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	IotAuthType          *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey           *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AttachDataWithSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDataWithSignatureRequest) GoString() string {
	return s.String()
}

func (s *AttachDataWithSignatureRequest) SetApiVersion(v string) *AttachDataWithSignatureRequest {
	s.ApiVersion = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetBusinessId(v string) *AttachDataWithSignatureRequest {
	s.BusinessId = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetIotAuthType(v string) *AttachDataWithSignatureRequest {
	s.IotAuthType = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetIotDataDigest(v string) *AttachDataWithSignatureRequest {
	s.IotDataDigest = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetIotId(v string) *AttachDataWithSignatureRequest {
	s.IotId = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetIotIdServiceProvider(v string) *AttachDataWithSignatureRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetIotIdSource(v string) *AttachDataWithSignatureRequest {
	s.IotIdSource = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetIotSignature(v string) *AttachDataWithSignatureRequest {
	s.IotSignature = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetKey(v string) *AttachDataWithSignatureRequest {
	s.Key = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetProductKey(v string) *AttachDataWithSignatureRequest {
	s.ProductKey = &v
	return s
}

func (s *AttachDataWithSignatureRequest) SetValue(v string) *AttachDataWithSignatureRequest {
	s.Value = &v
	return s
}

type AttachDataWithSignatureResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachDataWithSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDataWithSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDataWithSignatureResponseBody) SetCode(v int32) *AttachDataWithSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *AttachDataWithSignatureResponseBody) SetData(v string) *AttachDataWithSignatureResponseBody {
	s.Data = &v
	return s
}

func (s *AttachDataWithSignatureResponseBody) SetMessage(v string) *AttachDataWithSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *AttachDataWithSignatureResponseBody) SetRequestId(v string) *AttachDataWithSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDataWithSignatureResponseBody) SetSuccess(v bool) *AttachDataWithSignatureResponseBody {
	s.Success = &v
	return s
}

type AttachDataWithSignatureResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachDataWithSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDataWithSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDataWithSignatureResponse) GoString() string {
	return s.String()
}

func (s *AttachDataWithSignatureResponse) SetHeaders(v map[string]*string) *AttachDataWithSignatureResponse {
	s.Headers = v
	return s
}

func (s *AttachDataWithSignatureResponse) SetStatusCode(v int32) *AttachDataWithSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDataWithSignatureResponse) SetBody(v *AttachDataWithSignatureResponseBody) *AttachDataWithSignatureResponse {
	s.Body = v
	return s
}

type AuthorizeDeviceRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s AuthorizeDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeDeviceRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceRequest) SetApiVersion(v string) *AuthorizeDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *AuthorizeDeviceRequest) SetBizChainId(v string) *AuthorizeDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *AuthorizeDeviceRequest) SetDeviceGroupId(v string) *AuthorizeDeviceRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AuthorizeDeviceRequest) SetDeviceId(v string) *AuthorizeDeviceRequest {
	s.DeviceId = &v
	return s
}

type AuthorizeDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceResponseBody) SetCode(v int32) *AuthorizeDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeDeviceResponseBody) SetData(v string) *AuthorizeDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *AuthorizeDeviceResponseBody) SetMessage(v string) *AuthorizeDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeDeviceResponseBody) SetRequestId(v string) *AuthorizeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeDeviceResponseBody) SetSuccess(v bool) *AuthorizeDeviceResponseBody {
	s.Success = &v
	return s
}

type AuthorizeDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthorizeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeDeviceResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceResponse) SetHeaders(v map[string]*string) *AuthorizeDeviceResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeDeviceResponse) SetStatusCode(v int32) *AuthorizeDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeDeviceResponse) SetBody(v *AuthorizeDeviceResponseBody) *AuthorizeDeviceResponse {
	s.Body = v
	return s
}

type AuthorizeDeviceGroupRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s AuthorizeDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceGroupRequest) SetApiVersion(v string) *AuthorizeDeviceGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *AuthorizeDeviceGroupRequest) SetBizChainId(v string) *AuthorizeDeviceGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *AuthorizeDeviceGroupRequest) SetDeviceGroupId(v string) *AuthorizeDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

type AuthorizeDeviceGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceGroupResponseBody) SetCode(v int32) *AuthorizeDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeDeviceGroupResponseBody) SetData(v string) *AuthorizeDeviceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *AuthorizeDeviceGroupResponseBody) SetMessage(v string) *AuthorizeDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeDeviceGroupResponseBody) SetRequestId(v string) *AuthorizeDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeDeviceGroupResponseBody) SetSuccess(v bool) *AuthorizeDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type AuthorizeDeviceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthorizeDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeDeviceGroupResponse) SetStatusCode(v int32) *AuthorizeDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeDeviceGroupResponse) SetBody(v *AuthorizeDeviceGroupResponseBody) *AuthorizeDeviceGroupResponse {
	s.Body = v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoRequest struct {
	ApiVersion    *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	PhaseDataList map[string]interface{} `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId  *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId       *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseDigestInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseDigestInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseDigestInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoRequest) SetPhaseDataList(v map[string]interface{}) *BatchUploadMPCoSPhaseDigestInfoRequest {
	s.PhaseDataList = v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseDigestInfoRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseDigestInfoRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoShrinkRequest struct {
	ApiVersion          *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId          *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	PhaseDataListShrink *string `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId        *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId             *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseDigestInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoShrinkRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseDigestInfoShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoShrinkRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseDigestInfoShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoShrinkRequest) SetPhaseDataListShrink(v string) *BatchUploadMPCoSPhaseDigestInfoShrinkRequest {
	s.PhaseDataListShrink = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoShrinkRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseDigestInfoShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoShrinkRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseDigestInfoShrinkRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUploadMPCoSPhaseDigestInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponseBody) SetCode(v int32) *BatchUploadMPCoSPhaseDigestInfoResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponseBody) SetData(v string) *BatchUploadMPCoSPhaseDigestInfoResponseBody {
	s.Data = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponseBody) SetMessage(v string) *BatchUploadMPCoSPhaseDigestInfoResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponseBody) SetRequestId(v string) *BatchUploadMPCoSPhaseDigestInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponseBody) SetSuccess(v bool) *BatchUploadMPCoSPhaseDigestInfoResponseBody {
	s.Success = &v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUploadMPCoSPhaseDigestInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUploadMPCoSPhaseDigestInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponse) SetHeaders(v map[string]*string) *BatchUploadMPCoSPhaseDigestInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponse) SetStatusCode(v int32) *BatchUploadMPCoSPhaseDigestInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoResponse) SetBody(v *BatchUploadMPCoSPhaseDigestInfoResponseBody) *BatchUploadMPCoSPhaseDigestInfoResponse {
	s.Body = v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoByDeviceRequest struct {
	ApiVersion           *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId           *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	IotAuthType          *string                `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string                `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string                `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string                `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string                `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string                `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseDataList        map[string]interface{} `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId         *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId              *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotAuthType(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotAuthType = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotDataDigest(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotDataDigest = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotIdServiceProvider(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotIdSource(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotIdSource = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotSignature(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotSignature = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetPhaseDataList(v map[string]interface{}) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.PhaseDataList = v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest struct {
	ApiVersion           *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId           *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	IotAuthType          *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseDataListShrink  *string `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId         *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId              *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotAuthType(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotAuthType = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotDataDigest(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotDataDigest = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotIdServiceProvider(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotIdSource(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotIdSource = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotSignature(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotSignature = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetPhaseDataListShrink(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.PhaseDataListShrink = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetCode(v int32) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetData(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetMessage(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetRequestId(v string) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetSuccess(v bool) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Success = &v
	return s
}

type BatchUploadMPCoSPhaseDigestInfoByDeviceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseDigestInfoByDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse) SetHeaders(v map[string]*string) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse) SetStatusCode(v int32) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse) SetBody(v *BatchUploadMPCoSPhaseDigestInfoByDeviceResponseBody) *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse {
	s.Body = v
	return s
}

type BatchUploadMPCoSPhaseTextInfoRequest struct {
	ApiVersion    *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	PhaseDataList map[string]interface{} `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId  *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId       *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseTextInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseTextInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseTextInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoRequest) SetPhaseDataList(v map[string]interface{}) *BatchUploadMPCoSPhaseTextInfoRequest {
	s.PhaseDataList = v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseTextInfoRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseTextInfoRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseTextInfoShrinkRequest struct {
	ApiVersion          *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId          *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	PhaseDataListShrink *string `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId        *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId             *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseTextInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoShrinkRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseTextInfoShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoShrinkRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseTextInfoShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoShrinkRequest) SetPhaseDataListShrink(v string) *BatchUploadMPCoSPhaseTextInfoShrinkRequest {
	s.PhaseDataListShrink = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoShrinkRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseTextInfoShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoShrinkRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseTextInfoShrinkRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseTextInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUploadMPCoSPhaseTextInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoResponseBody) SetCode(v int32) *BatchUploadMPCoSPhaseTextInfoResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoResponseBody) SetData(v string) *BatchUploadMPCoSPhaseTextInfoResponseBody {
	s.Data = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoResponseBody) SetMessage(v string) *BatchUploadMPCoSPhaseTextInfoResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoResponseBody) SetRequestId(v string) *BatchUploadMPCoSPhaseTextInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoResponseBody) SetSuccess(v bool) *BatchUploadMPCoSPhaseTextInfoResponseBody {
	s.Success = &v
	return s
}

type BatchUploadMPCoSPhaseTextInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUploadMPCoSPhaseTextInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUploadMPCoSPhaseTextInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoResponse) SetHeaders(v map[string]*string) *BatchUploadMPCoSPhaseTextInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoResponse) SetStatusCode(v int32) *BatchUploadMPCoSPhaseTextInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoResponse) SetBody(v *BatchUploadMPCoSPhaseTextInfoResponseBody) *BatchUploadMPCoSPhaseTextInfoResponse {
	s.Body = v
	return s
}

type BatchUploadMPCoSPhaseTextInfoByDeviceRequest struct {
	ApiVersion           *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId           *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	IotAuthType          *string                `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string                `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string                `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string                `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string                `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string                `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseDataList        map[string]interface{} `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId         *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId              *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetIotAuthType(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotAuthType = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetIotDataDigest(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotDataDigest = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetIotId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetIotIdServiceProvider(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetIotIdSource(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotIdSource = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetIotSignature(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotSignature = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetPhaseDataList(v map[string]interface{}) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.PhaseDataList = v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest struct {
	ApiVersion           *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId           *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	IotAuthType          *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseDataListShrink  *string `json:"PhaseDataList,omitempty" xml:"PhaseDataList,omitempty"`
	PhaseGroupId         *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId              *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetApiVersion(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetBizChainId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotAuthType(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotAuthType = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotDataDigest(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotDataDigest = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotIdServiceProvider(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotIdSource(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotIdSource = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotSignature(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotSignature = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetPhaseDataListShrink(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.PhaseDataListShrink = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetPhaseGroupId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetPhaseId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.PhaseId = &v
	return s
}

type BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) SetCode(v int32) *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) SetData(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) SetMessage(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) SetRequestId(v string) *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) SetSuccess(v bool) *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Success = &v
	return s
}

type BatchUploadMPCoSPhaseTextInfoByDeviceResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUploadMPCoSPhaseTextInfoByDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponse) SetHeaders(v map[string]*string) *BatchUploadMPCoSPhaseTextInfoByDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponse) SetStatusCode(v int32) *BatchUploadMPCoSPhaseTextInfoByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUploadMPCoSPhaseTextInfoByDeviceResponse) SetBody(v *BatchUploadMPCoSPhaseTextInfoByDeviceResponseBody) *BatchUploadMPCoSPhaseTextInfoByDeviceResponse {
	s.Body = v
	return s
}

type CreateMPCoSPhaseRequest struct {
	ApiVersion   *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateMPCoSPhaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMPCoSPhaseRequest) GoString() string {
	return s.String()
}

func (s *CreateMPCoSPhaseRequest) SetApiVersion(v string) *CreateMPCoSPhaseRequest {
	s.ApiVersion = &v
	return s
}

func (s *CreateMPCoSPhaseRequest) SetBizChainId(v string) *CreateMPCoSPhaseRequest {
	s.BizChainId = &v
	return s
}

func (s *CreateMPCoSPhaseRequest) SetName(v string) *CreateMPCoSPhaseRequest {
	s.Name = &v
	return s
}

func (s *CreateMPCoSPhaseRequest) SetPhaseGroupId(v string) *CreateMPCoSPhaseRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *CreateMPCoSPhaseRequest) SetRemark(v string) *CreateMPCoSPhaseRequest {
	s.Remark = &v
	return s
}

type CreateMPCoSPhaseResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMPCoSPhaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMPCoSPhaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMPCoSPhaseResponseBody) SetCode(v int32) *CreateMPCoSPhaseResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMPCoSPhaseResponseBody) SetData(v string) *CreateMPCoSPhaseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateMPCoSPhaseResponseBody) SetMessage(v string) *CreateMPCoSPhaseResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMPCoSPhaseResponseBody) SetRequestId(v string) *CreateMPCoSPhaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMPCoSPhaseResponseBody) SetSuccess(v bool) *CreateMPCoSPhaseResponseBody {
	s.Success = &v
	return s
}

type CreateMPCoSPhaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMPCoSPhaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMPCoSPhaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMPCoSPhaseResponse) GoString() string {
	return s.String()
}

func (s *CreateMPCoSPhaseResponse) SetHeaders(v map[string]*string) *CreateMPCoSPhaseResponse {
	s.Headers = v
	return s
}

func (s *CreateMPCoSPhaseResponse) SetStatusCode(v int32) *CreateMPCoSPhaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMPCoSPhaseResponse) SetBody(v *CreateMPCoSPhaseResponseBody) *CreateMPCoSPhaseResponse {
	s.Body = v
	return s
}

type CreateMPCoSPhaseGroupRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateMPCoSPhaseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMPCoSPhaseGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMPCoSPhaseGroupRequest) SetApiVersion(v string) *CreateMPCoSPhaseGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *CreateMPCoSPhaseGroupRequest) SetBizChainId(v string) *CreateMPCoSPhaseGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *CreateMPCoSPhaseGroupRequest) SetName(v string) *CreateMPCoSPhaseGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateMPCoSPhaseGroupRequest) SetRemark(v string) *CreateMPCoSPhaseGroupRequest {
	s.Remark = &v
	return s
}

type CreateMPCoSPhaseGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMPCoSPhaseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMPCoSPhaseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMPCoSPhaseGroupResponseBody) SetCode(v int32) *CreateMPCoSPhaseGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMPCoSPhaseGroupResponseBody) SetData(v string) *CreateMPCoSPhaseGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateMPCoSPhaseGroupResponseBody) SetMessage(v string) *CreateMPCoSPhaseGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMPCoSPhaseGroupResponseBody) SetRequestId(v string) *CreateMPCoSPhaseGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMPCoSPhaseGroupResponseBody) SetSuccess(v bool) *CreateMPCoSPhaseGroupResponseBody {
	s.Success = &v
	return s
}

type CreateMPCoSPhaseGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMPCoSPhaseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMPCoSPhaseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMPCoSPhaseGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMPCoSPhaseGroupResponse) SetHeaders(v map[string]*string) *CreateMPCoSPhaseGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMPCoSPhaseGroupResponse) SetStatusCode(v int32) *CreateMPCoSPhaseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMPCoSPhaseGroupResponse) SetBody(v *CreateMPCoSPhaseGroupResponseBody) *CreateMPCoSPhaseGroupResponse {
	s.Body = v
	return s
}

type CreateMemberRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberContact *string `json:"MemberContact,omitempty" xml:"MemberContact,omitempty"`
	MemberName    *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberPhone   *string `json:"MemberPhone,omitempty" xml:"MemberPhone,omitempty"`
	MemberUid     *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberRequest) SetApiVersion(v string) *CreateMemberRequest {
	s.ApiVersion = &v
	return s
}

func (s *CreateMemberRequest) SetBizChainId(v string) *CreateMemberRequest {
	s.BizChainId = &v
	return s
}

func (s *CreateMemberRequest) SetMemberContact(v string) *CreateMemberRequest {
	s.MemberContact = &v
	return s
}

func (s *CreateMemberRequest) SetMemberName(v string) *CreateMemberRequest {
	s.MemberName = &v
	return s
}

func (s *CreateMemberRequest) SetMemberPhone(v string) *CreateMemberRequest {
	s.MemberPhone = &v
	return s
}

func (s *CreateMemberRequest) SetMemberUid(v string) *CreateMemberRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateMemberRequest) SetRemark(v string) *CreateMemberRequest {
	s.Remark = &v
	return s
}

type CreateMemberResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBody) SetCode(v int32) *CreateMemberResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemberResponseBody) SetData(v string) *CreateMemberResponseBody {
	s.Data = &v
	return s
}

func (s *CreateMemberResponseBody) SetMessage(v string) *CreateMemberResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMemberResponseBody) SetRequestId(v string) *CreateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemberResponseBody) SetSuccess(v bool) *CreateMemberResponseBody {
	s.Success = &v
	return s
}

type CreateMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateMemberResponse) SetHeaders(v map[string]*string) *CreateMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateMemberResponse) SetStatusCode(v int32) *CreateMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemberResponse) SetBody(v *CreateMemberResponseBody) *CreateMemberResponse {
	s.Body = v
	return s
}

type DescribeCapacityInfoRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
}

func (s DescribeCapacityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapacityInfoRequest) SetApiVersion(v string) *DescribeCapacityInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeCapacityInfoRequest) SetBizChainId(v string) *DescribeCapacityInfoRequest {
	s.BizChainId = &v
	return s
}

type DescribeCapacityInfoResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCapacityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCapacityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapacityInfoResponseBody) SetCode(v int32) *DescribeCapacityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCapacityInfoResponseBody) SetData(v *DescribeCapacityInfoResponseBodyData) *DescribeCapacityInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCapacityInfoResponseBody) SetMessage(v string) *DescribeCapacityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCapacityInfoResponseBody) SetRequestId(v string) *DescribeCapacityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapacityInfoResponseBody) SetSuccess(v bool) *DescribeCapacityInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeCapacityInfoResponseBodyData struct {
	CapacityQuota      *int64 `json:"CapacityQuota,omitempty" xml:"CapacityQuota,omitempty"`
	CountQuota         *int64 `json:"CountQuota,omitempty" xml:"CountQuota,omitempty"`
	MemberUsedCapacity *int64 `json:"MemberUsedCapacity,omitempty" xml:"MemberUsedCapacity,omitempty"`
	MemberUsedCount    *int64 `json:"MemberUsedCount,omitempty" xml:"MemberUsedCount,omitempty"`
	UsedCapacity       *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	UsedCount          *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeCapacityInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCapacityInfoResponseBodyData) SetCapacityQuota(v int64) *DescribeCapacityInfoResponseBodyData {
	s.CapacityQuota = &v
	return s
}

func (s *DescribeCapacityInfoResponseBodyData) SetCountQuota(v int64) *DescribeCapacityInfoResponseBodyData {
	s.CountQuota = &v
	return s
}

func (s *DescribeCapacityInfoResponseBodyData) SetMemberUsedCapacity(v int64) *DescribeCapacityInfoResponseBodyData {
	s.MemberUsedCapacity = &v
	return s
}

func (s *DescribeCapacityInfoResponseBodyData) SetMemberUsedCount(v int64) *DescribeCapacityInfoResponseBodyData {
	s.MemberUsedCount = &v
	return s
}

func (s *DescribeCapacityInfoResponseBodyData) SetUsedCapacity(v int64) *DescribeCapacityInfoResponseBodyData {
	s.UsedCapacity = &v
	return s
}

func (s *DescribeCapacityInfoResponseBodyData) SetUsedCount(v int64) *DescribeCapacityInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

type DescribeCapacityInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCapacityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCapacityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapacityInfoResponse) SetHeaders(v map[string]*string) *DescribeCapacityInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapacityInfoResponse) SetStatusCode(v int32) *DescribeCapacityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapacityInfoResponse) SetBody(v *DescribeCapacityInfoResponseBody) *DescribeCapacityInfoResponse {
	s.Body = v
	return s
}

type DescribeMPCoSAuthorizedInfoRequest struct {
	ApiVersion   *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberId     *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
}

func (s DescribeMPCoSAuthorizedInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSAuthorizedInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSAuthorizedInfoRequest) SetApiVersion(v string) *DescribeMPCoSAuthorizedInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoRequest) SetBizChainId(v string) *DescribeMPCoSAuthorizedInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoRequest) SetMemberId(v string) *DescribeMPCoSAuthorizedInfoRequest {
	s.MemberId = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoRequest) SetPhaseGroupId(v string) *DescribeMPCoSAuthorizedInfoRequest {
	s.PhaseGroupId = &v
	return s
}

type DescribeMPCoSAuthorizedInfoResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeMPCoSAuthorizedInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMPCoSAuthorizedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSAuthorizedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSAuthorizedInfoResponseBody) SetCode(v int32) *DescribeMPCoSAuthorizedInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBody) SetData(v *DescribeMPCoSAuthorizedInfoResponseBodyData) *DescribeMPCoSAuthorizedInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBody) SetMessage(v string) *DescribeMPCoSAuthorizedInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBody) SetRequestId(v string) *DescribeMPCoSAuthorizedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBody) SetSuccess(v bool) *DescribeMPCoSAuthorizedInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeMPCoSAuthorizedInfoResponseBodyData struct {
	AuthorizedPhaseList   []*DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList   `json:"AuthorizedPhaseList,omitempty" xml:"AuthorizedPhaseList,omitempty" type:"Repeated"`
	UnAuthorizedPhaseList []*DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList `json:"UnAuthorizedPhaseList,omitempty" xml:"UnAuthorizedPhaseList,omitempty" type:"Repeated"`
}

func (s DescribeMPCoSAuthorizedInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSAuthorizedInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSAuthorizedInfoResponseBodyData) SetAuthorizedPhaseList(v []*DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList) *DescribeMPCoSAuthorizedInfoResponseBodyData {
	s.AuthorizedPhaseList = v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBodyData) SetUnAuthorizedPhaseList(v []*DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList) *DescribeMPCoSAuthorizedInfoResponseBodyData {
	s.UnAuthorizedPhaseList = v
	return s
}

type DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList struct {
	PhaseId   *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	PhaseName *string `json:"PhaseName,omitempty" xml:"PhaseName,omitempty"`
}

func (s DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList) SetPhaseId(v string) *DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList {
	s.PhaseId = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList) SetPhaseName(v string) *DescribeMPCoSAuthorizedInfoResponseBodyDataAuthorizedPhaseList {
	s.PhaseName = &v
	return s
}

type DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList struct {
	PhaseId   *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	PhaseName *string `json:"PhaseName,omitempty" xml:"PhaseName,omitempty"`
}

func (s DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList) SetPhaseId(v string) *DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList {
	s.PhaseId = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList) SetPhaseName(v string) *DescribeMPCoSAuthorizedInfoResponseBodyDataUnAuthorizedPhaseList {
	s.PhaseName = &v
	return s
}

type DescribeMPCoSAuthorizedInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMPCoSAuthorizedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPCoSAuthorizedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSAuthorizedInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSAuthorizedInfoResponse) SetHeaders(v map[string]*string) *DescribeMPCoSAuthorizedInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponse) SetStatusCode(v int32) *DescribeMPCoSAuthorizedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMPCoSAuthorizedInfoResponse) SetBody(v *DescribeMPCoSAuthorizedInfoResponseBody) *DescribeMPCoSAuthorizedInfoResponse {
	s.Body = v
	return s
}

type DescribeMPCoSPhaseInfoRequest struct {
	ApiVersion   *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey      *string `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq      *string `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId      *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
}

func (s DescribeMPCoSPhaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSPhaseInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSPhaseInfoRequest) SetApiVersion(v string) *DescribeMPCoSPhaseInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoRequest) SetBizChainId(v string) *DescribeMPCoSPhaseInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoRequest) SetDataKey(v string) *DescribeMPCoSPhaseInfoRequest {
	s.DataKey = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoRequest) SetDataSeq(v string) *DescribeMPCoSPhaseInfoRequest {
	s.DataSeq = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoRequest) SetPhaseGroupId(v string) *DescribeMPCoSPhaseInfoRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoRequest) SetPhaseId(v string) *DescribeMPCoSPhaseInfoRequest {
	s.PhaseId = &v
	return s
}

type DescribeMPCoSPhaseInfoResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeMPCoSPhaseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMPCoSPhaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSPhaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSPhaseInfoResponseBody) SetCode(v int32) *DescribeMPCoSPhaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBody) SetData(v *DescribeMPCoSPhaseInfoResponseBodyData) *DescribeMPCoSPhaseInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBody) SetMessage(v string) *DescribeMPCoSPhaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBody) SetRequestId(v string) *DescribeMPCoSPhaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBody) SetSuccess(v bool) *DescribeMPCoSPhaseInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeMPCoSPhaseInfoResponseBodyData struct {
	BlockHash       *string                                                  `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	BlockNumber     *int64                                                   `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	DataHash        *string                                                  `json:"DataHash,omitempty" xml:"DataHash,omitempty"`
	DataValue       *string                                                  `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
	IotId           *string                                                  `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PreviousHash    *string                                                  `json:"PreviousHash,omitempty" xml:"PreviousHash,omitempty"`
	ProductKey      *string                                                  `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RelatedDataList []*DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty" type:"Repeated"`
	Timestamp       *int64                                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TransactionHash *string                                                  `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
}

func (s DescribeMPCoSPhaseInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSPhaseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetBlockHash(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.BlockHash = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetBlockNumber(v int64) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.BlockNumber = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetDataHash(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.DataHash = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetDataValue(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.DataValue = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetIotId(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.IotId = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetPreviousHash(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.PreviousHash = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetProductKey(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetRelatedDataList(v []*DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.RelatedDataList = v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetTimestamp(v int64) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyData) SetTransactionHash(v string) *DescribeMPCoSPhaseInfoResponseBodyData {
	s.TransactionHash = &v
	return s
}

type DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList struct {
	RelatedDataKey       *string `json:"RelatedDataKey,omitempty" xml:"RelatedDataKey,omitempty"`
	RelatedDataSeq       *string `json:"RelatedDataSeq,omitempty" xml:"RelatedDataSeq,omitempty"`
	RelatedPhaseDataHash *string `json:"RelatedPhaseDataHash,omitempty" xml:"RelatedPhaseDataHash,omitempty"`
	RelatedPhaseId       *string `json:"RelatedPhaseId,omitempty" xml:"RelatedPhaseId,omitempty"`
	RelatedPhaseName     *string `json:"RelatedPhaseName,omitempty" xml:"RelatedPhaseName,omitempty"`
}

func (s DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) SetRelatedDataKey(v string) *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList {
	s.RelatedDataKey = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) SetRelatedDataSeq(v string) *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList {
	s.RelatedDataSeq = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) SetRelatedPhaseDataHash(v string) *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList {
	s.RelatedPhaseDataHash = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) SetRelatedPhaseId(v string) *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList {
	s.RelatedPhaseId = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList) SetRelatedPhaseName(v string) *DescribeMPCoSPhaseInfoResponseBodyDataRelatedDataList {
	s.RelatedPhaseName = &v
	return s
}

type DescribeMPCoSPhaseInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMPCoSPhaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPCoSPhaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSPhaseInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSPhaseInfoResponse) SetHeaders(v map[string]*string) *DescribeMPCoSPhaseInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponse) SetStatusCode(v int32) *DescribeMPCoSPhaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMPCoSPhaseInfoResponse) SetBody(v *DescribeMPCoSPhaseInfoResponseBody) *DescribeMPCoSPhaseInfoResponse {
	s.Body = v
	return s
}

type DescribeMPCoSResourceInfoRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
}

func (s DescribeMPCoSResourceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSResourceInfoRequest) SetApiVersion(v string) *DescribeMPCoSResourceInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeMPCoSResourceInfoRequest) SetBizChainId(v string) *DescribeMPCoSResourceInfoRequest {
	s.BizChainId = &v
	return s
}

type DescribeMPCoSResourceInfoResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeMPCoSResourceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMPCoSResourceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSResourceInfoResponseBody) SetCode(v int32) *DescribeMPCoSResourceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBody) SetData(v *DescribeMPCoSResourceInfoResponseBodyData) *DescribeMPCoSResourceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBody) SetMessage(v string) *DescribeMPCoSResourceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBody) SetRequestId(v string) *DescribeMPCoSResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBody) SetSuccess(v bool) *DescribeMPCoSResourceInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeMPCoSResourceInfoResponseBodyData struct {
	MemberQuota        *int64                                                         `json:"MemberQuota,omitempty" xml:"MemberQuota,omitempty"`
	PhaseGroupQuota    *int64                                                         `json:"PhaseGroupQuota,omitempty" xml:"PhaseGroupQuota,omitempty"`
	PhaseQuotaInfoList []*DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList `json:"PhaseQuotaInfoList,omitempty" xml:"PhaseQuotaInfoList,omitempty" type:"Repeated"`
	UsedMember         *int64                                                         `json:"UsedMember,omitempty" xml:"UsedMember,omitempty"`
	UsedPhaseGroup     *int64                                                         `json:"UsedPhaseGroup,omitempty" xml:"UsedPhaseGroup,omitempty"`
}

func (s DescribeMPCoSResourceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSResourceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSResourceInfoResponseBodyData) SetMemberQuota(v int64) *DescribeMPCoSResourceInfoResponseBodyData {
	s.MemberQuota = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyData) SetPhaseGroupQuota(v int64) *DescribeMPCoSResourceInfoResponseBodyData {
	s.PhaseGroupQuota = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyData) SetPhaseQuotaInfoList(v []*DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) *DescribeMPCoSResourceInfoResponseBodyData {
	s.PhaseQuotaInfoList = v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyData) SetUsedMember(v int64) *DescribeMPCoSResourceInfoResponseBodyData {
	s.UsedMember = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyData) SetUsedPhaseGroup(v int64) *DescribeMPCoSResourceInfoResponseBodyData {
	s.UsedPhaseGroup = &v
	return s
}

type DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList struct {
	PhaseGroupId   *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseGroupName *string `json:"PhaseGroupName,omitempty" xml:"PhaseGroupName,omitempty"`
	PhaseQuota     *int64  `json:"PhaseQuota,omitempty" xml:"PhaseQuota,omitempty"`
	UsedPhase      *int64  `json:"UsedPhase,omitempty" xml:"UsedPhase,omitempty"`
}

func (s DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) SetPhaseGroupId(v string) *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList {
	s.PhaseGroupId = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) SetPhaseGroupName(v string) *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList {
	s.PhaseGroupName = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) SetPhaseQuota(v int64) *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList {
	s.PhaseQuota = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList) SetUsedPhase(v int64) *DescribeMPCoSResourceInfoResponseBodyDataPhaseQuotaInfoList {
	s.UsedPhase = &v
	return s
}

type DescribeMPCoSResourceInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMPCoSResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPCoSResourceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPCoSResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPCoSResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeMPCoSResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPCoSResourceInfoResponse) SetStatusCode(v int32) *DescribeMPCoSResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMPCoSResourceInfoResponse) SetBody(v *DescribeMPCoSResourceInfoResponseBody) *DescribeMPCoSResourceInfoResponse {
	s.Body = v
	return s
}

type DescribeMemberCapacityInfoRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
}

func (s DescribeMemberCapacityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberCapacityInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberCapacityInfoRequest) SetApiVersion(v string) *DescribeMemberCapacityInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeMemberCapacityInfoRequest) SetBizChainId(v string) *DescribeMemberCapacityInfoRequest {
	s.BizChainId = &v
	return s
}

type DescribeMemberCapacityInfoResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeMemberCapacityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMemberCapacityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberCapacityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberCapacityInfoResponseBody) SetCode(v int32) *DescribeMemberCapacityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBody) SetData(v []*DescribeMemberCapacityInfoResponseBodyData) *DescribeMemberCapacityInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBody) SetMessage(v string) *DescribeMemberCapacityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBody) SetRequestId(v string) *DescribeMemberCapacityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBody) SetSuccess(v bool) *DescribeMemberCapacityInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeMemberCapacityInfoResponseBodyData struct {
	MemberId     *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberUid    *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	UsedCapacity *string `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	UsedCount    *string `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeMemberCapacityInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberCapacityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMemberCapacityInfoResponseBodyData) SetMemberId(v string) *DescribeMemberCapacityInfoResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBodyData) SetMemberName(v string) *DescribeMemberCapacityInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBodyData) SetMemberUid(v string) *DescribeMemberCapacityInfoResponseBodyData {
	s.MemberUid = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBodyData) SetUsedCapacity(v string) *DescribeMemberCapacityInfoResponseBodyData {
	s.UsedCapacity = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponseBodyData) SetUsedCount(v string) *DescribeMemberCapacityInfoResponseBodyData {
	s.UsedCount = &v
	return s
}

type DescribeMemberCapacityInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMemberCapacityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMemberCapacityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberCapacityInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberCapacityInfoResponse) SetHeaders(v map[string]*string) *DescribeMemberCapacityInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberCapacityInfoResponse) SetStatusCode(v int32) *DescribeMemberCapacityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberCapacityInfoResponse) SetBody(v *DescribeMemberCapacityInfoResponseBody) *DescribeMemberCapacityInfoResponse {
	s.Body = v
	return s
}

type DescribeResourceInfoRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
}

func (s DescribeResourceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceInfoRequest) SetApiVersion(v string) *DescribeResourceInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeResourceInfoRequest) SetBizChainId(v string) *DescribeResourceInfoRequest {
	s.BizChainId = &v
	return s
}

type DescribeResourceInfoResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeResourceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeResourceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceInfoResponseBody) SetCode(v int32) *DescribeResourceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourceInfoResponseBody) SetData(v *DescribeResourceInfoResponseBodyData) *DescribeResourceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourceInfoResponseBody) SetMessage(v string) *DescribeResourceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceInfoResponseBody) SetRequestId(v string) *DescribeResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceInfoResponseBody) SetSuccess(v bool) *DescribeResourceInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeResourceInfoResponseBodyData struct {
	AuthorizeType *string `json:"AuthorizeType,omitempty" xml:"AuthorizeType,omitempty"`
	EffectiveTime *int64  `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiredTime   *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourceInfoResponseBodyData) SetAuthorizeType(v string) *DescribeResourceInfoResponseBodyData {
	s.AuthorizeType = &v
	return s
}

func (s *DescribeResourceInfoResponseBodyData) SetEffectiveTime(v int64) *DescribeResourceInfoResponseBodyData {
	s.EffectiveTime = &v
	return s
}

func (s *DescribeResourceInfoResponseBodyData) SetExpiredTime(v int64) *DescribeResourceInfoResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeResourceInfoResponseBodyData) SetRegion(v string) *DescribeResourceInfoResponseBodyData {
	s.Region = &v
	return s
}

func (s *DescribeResourceInfoResponseBodyData) SetStatus(v string) *DescribeResourceInfoResponseBodyData {
	s.Status = &v
	return s
}

type DescribeResourceInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceInfoResponse) SetStatusCode(v int32) *DescribeResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceInfoResponse) SetBody(v *DescribeResourceInfoResponseBody) *DescribeResourceInfoResponse {
	s.Body = v
	return s
}

type GetBlockChainInfoRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s GetBlockChainInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBlockChainInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBlockChainInfoRequest) SetApiVersion(v string) *GetBlockChainInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetBlockChainInfoRequest) SetBusinessId(v string) *GetBlockChainInfoRequest {
	s.BusinessId = &v
	return s
}

func (s *GetBlockChainInfoRequest) SetKey(v string) *GetBlockChainInfoRequest {
	s.Key = &v
	return s
}

func (s *GetBlockChainInfoRequest) SetProductKey(v string) *GetBlockChainInfoRequest {
	s.ProductKey = &v
	return s
}

type GetBlockChainInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBlockChainInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBlockChainInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBlockChainInfoResponseBody) SetCode(v int32) *GetBlockChainInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetBlockChainInfoResponseBody) SetData(v string) *GetBlockChainInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetBlockChainInfoResponseBody) SetMessage(v string) *GetBlockChainInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetBlockChainInfoResponseBody) SetRequestId(v string) *GetBlockChainInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBlockChainInfoResponseBody) SetSuccess(v bool) *GetBlockChainInfoResponseBody {
	s.Success = &v
	return s
}

type GetBlockChainInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBlockChainInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBlockChainInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBlockChainInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBlockChainInfoResponse) SetHeaders(v map[string]*string) *GetBlockChainInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBlockChainInfoResponse) SetStatusCode(v int32) *GetBlockChainInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBlockChainInfoResponse) SetBody(v *GetBlockChainInfoResponseBody) *GetBlockChainInfoResponse {
	s.Body = v
	return s
}

type GetDataRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s GetDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataRequest) GoString() string {
	return s.String()
}

func (s *GetDataRequest) SetApiVersion(v string) *GetDataRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetDataRequest) SetBusinessId(v string) *GetDataRequest {
	s.BusinessId = &v
	return s
}

func (s *GetDataRequest) SetKey(v string) *GetDataRequest {
	s.Key = &v
	return s
}

func (s *GetDataRequest) SetProductKey(v string) *GetDataRequest {
	s.ProductKey = &v
	return s
}

type GetDataResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataResponseBody) SetCode(v int32) *GetDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataResponseBody) SetData(v string) *GetDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetDataResponseBody) SetMessage(v string) *GetDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataResponseBody) SetRequestId(v string) *GetDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataResponseBody) SetSuccess(v bool) *GetDataResponseBody {
	s.Success = &v
	return s
}

type GetDataResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataResponse) GoString() string {
	return s.String()
}

func (s *GetDataResponse) SetHeaders(v map[string]*string) *GetDataResponse {
	s.Headers = v
	return s
}

func (s *GetDataResponse) SetStatusCode(v int32) *GetDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataResponse) SetBody(v *GetDataResponseBody) *GetDataResponse {
	s.Body = v
	return s
}

type GetDataModelConfigInfoRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	DataModelCode *string `json:"DataModelCode,omitempty" xml:"DataModelCode,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s GetDataModelConfigInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataModelConfigInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDataModelConfigInfoRequest) SetApiVersion(v string) *GetDataModelConfigInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetDataModelConfigInfoRequest) SetDataModelCode(v string) *GetDataModelConfigInfoRequest {
	s.DataModelCode = &v
	return s
}

func (s *GetDataModelConfigInfoRequest) SetProductKey(v string) *GetDataModelConfigInfoRequest {
	s.ProductKey = &v
	return s
}

type GetDataModelConfigInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataModelConfigInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataModelConfigInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataModelConfigInfoResponseBody) SetCode(v int32) *GetDataModelConfigInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataModelConfigInfoResponseBody) SetData(v string) *GetDataModelConfigInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetDataModelConfigInfoResponseBody) SetMessage(v string) *GetDataModelConfigInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataModelConfigInfoResponseBody) SetRequestId(v string) *GetDataModelConfigInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataModelConfigInfoResponseBody) SetSuccess(v bool) *GetDataModelConfigInfoResponseBody {
	s.Success = &v
	return s
}

type GetDataModelConfigInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDataModelConfigInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataModelConfigInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataModelConfigInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDataModelConfigInfoResponse) SetHeaders(v map[string]*string) *GetDataModelConfigInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDataModelConfigInfoResponse) SetStatusCode(v int32) *GetDataModelConfigInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataModelConfigInfoResponse) SetBody(v *GetDataModelConfigInfoResponseBody) *GetDataModelConfigInfoResponse {
	s.Body = v
	return s
}

type GetHistoryDataCountRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHistoryDataCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataCountRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryDataCountRequest) SetApiVersion(v string) *GetHistoryDataCountRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetHistoryDataCountRequest) SetEndTime(v int64) *GetHistoryDataCountRequest {
	s.EndTime = &v
	return s
}

func (s *GetHistoryDataCountRequest) SetKey(v string) *GetHistoryDataCountRequest {
	s.Key = &v
	return s
}

func (s *GetHistoryDataCountRequest) SetProductKey(v string) *GetHistoryDataCountRequest {
	s.ProductKey = &v
	return s
}

func (s *GetHistoryDataCountRequest) SetStartTime(v int64) *GetHistoryDataCountRequest {
	s.StartTime = &v
	return s
}

type GetHistoryDataCountResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHistoryDataCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryDataCountResponseBody) SetCode(v int32) *GetHistoryDataCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoryDataCountResponseBody) SetData(v string) *GetHistoryDataCountResponseBody {
	s.Data = &v
	return s
}

func (s *GetHistoryDataCountResponseBody) SetMessage(v string) *GetHistoryDataCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoryDataCountResponseBody) SetRequestId(v string) *GetHistoryDataCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoryDataCountResponseBody) SetSuccess(v bool) *GetHistoryDataCountResponseBody {
	s.Success = &v
	return s
}

type GetHistoryDataCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHistoryDataCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHistoryDataCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataCountResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryDataCountResponse) SetHeaders(v map[string]*string) *GetHistoryDataCountResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryDataCountResponse) SetStatusCode(v int32) *GetHistoryDataCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryDataCountResponse) SetBody(v *GetHistoryDataCountResponseBody) *GetHistoryDataCountResponse {
	s.Body = v
	return s
}

type GetHistoryDataListRequest struct {
	ApiVersion  *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey  *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHistoryDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataListRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryDataListRequest) SetApiVersion(v string) *GetHistoryDataListRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetHistoryDataListRequest) SetCurrentPage(v int32) *GetHistoryDataListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHistoryDataListRequest) SetEndTime(v int64) *GetHistoryDataListRequest {
	s.EndTime = &v
	return s
}

func (s *GetHistoryDataListRequest) SetKey(v string) *GetHistoryDataListRequest {
	s.Key = &v
	return s
}

func (s *GetHistoryDataListRequest) SetPageSize(v int32) *GetHistoryDataListRequest {
	s.PageSize = &v
	return s
}

func (s *GetHistoryDataListRequest) SetProductKey(v string) *GetHistoryDataListRequest {
	s.ProductKey = &v
	return s
}

func (s *GetHistoryDataListRequest) SetStartTime(v int64) *GetHistoryDataListRequest {
	s.StartTime = &v
	return s
}

type GetHistoryDataListResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHistoryDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHistoryDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryDataListResponseBody) SetCode(v int32) *GetHistoryDataListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoryDataListResponseBody) SetData(v *GetHistoryDataListResponseBodyData) *GetHistoryDataListResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoryDataListResponseBody) SetMessage(v string) *GetHistoryDataListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoryDataListResponseBody) SetRequestId(v string) *GetHistoryDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoryDataListResponseBody) SetSuccess(v bool) *GetHistoryDataListResponseBody {
	s.Success = &v
	return s
}

type GetHistoryDataListResponseBodyData struct {
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetHistoryDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoryDataListResponseBodyData) SetData(v []map[string]interface{}) *GetHistoryDataListResponseBodyData {
	s.Data = v
	return s
}

type GetHistoryDataListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHistoryDataListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHistoryDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryDataListResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryDataListResponse) SetHeaders(v map[string]*string) *GetHistoryDataListResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryDataListResponse) SetStatusCode(v int32) *GetHistoryDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryDataListResponse) SetBody(v *GetHistoryDataListResponseBody) *GetHistoryDataListResponse {
	s.Body = v
	return s
}

type ListDependentDataModelsRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ListDependentDataModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDependentDataModelsRequest) GoString() string {
	return s.String()
}

func (s *ListDependentDataModelsRequest) SetApiVersion(v string) *ListDependentDataModelsRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListDependentDataModelsRequest) SetProductKey(v string) *ListDependentDataModelsRequest {
	s.ProductKey = &v
	return s
}

type ListDependentDataModelsResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDependentDataModelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDependentDataModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDependentDataModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDependentDataModelsResponseBody) SetCode(v int32) *ListDependentDataModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDependentDataModelsResponseBody) SetData(v *ListDependentDataModelsResponseBodyData) *ListDependentDataModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListDependentDataModelsResponseBody) SetMessage(v string) *ListDependentDataModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDependentDataModelsResponseBody) SetRequestId(v string) *ListDependentDataModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDependentDataModelsResponseBody) SetSuccess(v bool) *ListDependentDataModelsResponseBody {
	s.Success = &v
	return s
}

type ListDependentDataModelsResponseBodyData struct {
	DataModelInfo []*ListDependentDataModelsResponseBodyDataDataModelInfo `json:"DataModelInfo,omitempty" xml:"DataModelInfo,omitempty" type:"Repeated"`
}

func (s ListDependentDataModelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDependentDataModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDependentDataModelsResponseBodyData) SetDataModelInfo(v []*ListDependentDataModelsResponseBodyDataDataModelInfo) *ListDependentDataModelsResponseBodyData {
	s.DataModelInfo = v
	return s
}

type ListDependentDataModelsResponseBodyDataDataModelInfo struct {
	DataModelCode *string `json:"DataModelCode,omitempty" xml:"DataModelCode,omitempty"`
	DataModelName *string `json:"DataModelName,omitempty" xml:"DataModelName,omitempty"`
}

func (s ListDependentDataModelsResponseBodyDataDataModelInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDependentDataModelsResponseBodyDataDataModelInfo) GoString() string {
	return s.String()
}

func (s *ListDependentDataModelsResponseBodyDataDataModelInfo) SetDataModelCode(v string) *ListDependentDataModelsResponseBodyDataDataModelInfo {
	s.DataModelCode = &v
	return s
}

func (s *ListDependentDataModelsResponseBodyDataDataModelInfo) SetDataModelName(v string) *ListDependentDataModelsResponseBodyDataDataModelInfo {
	s.DataModelName = &v
	return s
}

type ListDependentDataModelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDependentDataModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDependentDataModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDependentDataModelsResponse) GoString() string {
	return s.String()
}

func (s *ListDependentDataModelsResponse) SetHeaders(v map[string]*string) *ListDependentDataModelsResponse {
	s.Headers = v
	return s
}

func (s *ListDependentDataModelsResponse) SetStatusCode(v int32) *ListDependentDataModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDependentDataModelsResponse) SetBody(v *ListDependentDataModelsResponseBody) *ListDependentDataModelsResponse {
	s.Body = v
	return s
}

type ListDeviceRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	Num           *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Size          *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceRequest) SetApiVersion(v string) *ListDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListDeviceRequest) SetBizChainId(v string) *ListDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *ListDeviceRequest) SetDeviceGroupId(v string) *ListDeviceRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDeviceRequest) SetIotId(v string) *ListDeviceRequest {
	s.IotId = &v
	return s
}

func (s *ListDeviceRequest) SetNum(v int32) *ListDeviceRequest {
	s.Num = &v
	return s
}

func (s *ListDeviceRequest) SetSize(v int32) *ListDeviceRequest {
	s.Size = &v
	return s
}

type ListDeviceResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceResponseBody) SetCode(v int32) *ListDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceResponseBody) SetData(v *ListDeviceResponseBodyData) *ListDeviceResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceResponseBody) SetMessage(v string) *ListDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceResponseBody) SetRequestId(v string) *ListDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceResponseBody) SetSuccess(v bool) *ListDeviceResponseBody {
	s.Success = &v
	return s
}

type ListDeviceResponseBodyData struct {
	Num      *int32                                `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListDeviceResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceResponseBodyData) SetNum(v int32) *ListDeviceResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListDeviceResponseBodyData) SetPageData(v []*ListDeviceResponseBodyDataPageData) *ListDeviceResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListDeviceResponseBodyData) SetSize(v int32) *ListDeviceResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListDeviceResponseBodyData) SetTotal(v int32) *ListDeviceResponseBodyData {
	s.Total = &v
	return s
}

type ListDeviceResponseBodyDataPageData struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	LastSaveTime *int64  `json:"LastSaveTime,omitempty" xml:"LastSaveTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeviceResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListDeviceResponseBodyDataPageData) SetDeviceId(v string) *ListDeviceResponseBodyDataPageData {
	s.DeviceId = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetIotId(v string) *ListDeviceResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetLastSaveTime(v int64) *ListDeviceResponseBodyDataPageData {
	s.LastSaveTime = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetStatus(v string) *ListDeviceResponseBodyDataPageData {
	s.Status = &v
	return s
}

type ListDeviceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceResponse) SetHeaders(v map[string]*string) *ListDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceResponse) SetStatusCode(v int32) *ListDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceResponse) SetBody(v *ListDeviceResponseBody) *ListDeviceResponse {
	s.Body = v
	return s
}

type ListDeviceGroupRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupRequest) SetApiVersion(v string) *ListDeviceGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListDeviceGroupRequest) SetBizChainId(v string) *ListDeviceGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *ListDeviceGroupRequest) SetNum(v int32) *ListDeviceGroupRequest {
	s.Num = &v
	return s
}

func (s *ListDeviceGroupRequest) SetProductKey(v string) *ListDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *ListDeviceGroupRequest) SetSize(v int32) *ListDeviceGroupRequest {
	s.Size = &v
	return s
}

type ListDeviceGroupResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponseBody) SetCode(v int32) *ListDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetData(v *ListDeviceGroupResponseBodyData) *ListDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceGroupResponseBody) SetMessage(v string) *ListDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetRequestId(v string) *ListDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetSuccess(v bool) *ListDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type ListDeviceGroupResponseBodyData struct {
	Num      *int32                                     `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListDeviceGroupResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                     `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponseBodyData) SetNum(v int32) *ListDeviceGroupResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListDeviceGroupResponseBodyData) SetPageData(v []*ListDeviceGroupResponseBodyDataPageData) *ListDeviceGroupResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListDeviceGroupResponseBodyData) SetSize(v int32) *ListDeviceGroupResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListDeviceGroupResponseBodyData) SetTotal(v int32) *ListDeviceGroupResponseBodyData {
	s.Total = &v
	return s
}

type ListDeviceGroupResponseBodyDataPageData struct {
	AuthorizeType *string `json:"AuthorizeType,omitempty" xml:"AuthorizeType,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	OwnerName     *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerUid      *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeviceGroupResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetAuthorizeType(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.AuthorizeType = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetDeviceGroupId(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetOwnerName(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.OwnerName = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetOwnerUid(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.OwnerUid = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetProductKey(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.ProductKey = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetRemark(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetStatus(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.Status = &v
	return s
}

type ListDeviceGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponse) SetHeaders(v map[string]*string) *ListDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceGroupResponse) SetStatusCode(v int32) *ListDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceGroupResponse) SetBody(v *ListDeviceGroupResponseBody) *ListDeviceGroupResponse {
	s.Body = v
	return s
}

type ListMPCoSPhaseRequest struct {
	ApiVersion   *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListMPCoSPhaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseRequest) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseRequest) SetApiVersion(v string) *ListMPCoSPhaseRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListMPCoSPhaseRequest) SetBizChainId(v string) *ListMPCoSPhaseRequest {
	s.BizChainId = &v
	return s
}

func (s *ListMPCoSPhaseRequest) SetName(v string) *ListMPCoSPhaseRequest {
	s.Name = &v
	return s
}

func (s *ListMPCoSPhaseRequest) SetNum(v int32) *ListMPCoSPhaseRequest {
	s.Num = &v
	return s
}

func (s *ListMPCoSPhaseRequest) SetPhaseGroupId(v string) *ListMPCoSPhaseRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *ListMPCoSPhaseRequest) SetSize(v int32) *ListMPCoSPhaseRequest {
	s.Size = &v
	return s
}

type ListMPCoSPhaseResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListMPCoSPhaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMPCoSPhaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseResponseBody) SetCode(v int32) *ListMPCoSPhaseResponseBody {
	s.Code = &v
	return s
}

func (s *ListMPCoSPhaseResponseBody) SetData(v *ListMPCoSPhaseResponseBodyData) *ListMPCoSPhaseResponseBody {
	s.Data = v
	return s
}

func (s *ListMPCoSPhaseResponseBody) SetMessage(v string) *ListMPCoSPhaseResponseBody {
	s.Message = &v
	return s
}

func (s *ListMPCoSPhaseResponseBody) SetRequestId(v string) *ListMPCoSPhaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMPCoSPhaseResponseBody) SetSuccess(v bool) *ListMPCoSPhaseResponseBody {
	s.Success = &v
	return s
}

type ListMPCoSPhaseResponseBodyData struct {
	Num      *int32                                    `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMPCoSPhaseResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMPCoSPhaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseResponseBodyData) SetNum(v int32) *ListMPCoSPhaseResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMPCoSPhaseResponseBodyData) SetPageData(v []*ListMPCoSPhaseResponseBodyDataPageData) *ListMPCoSPhaseResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMPCoSPhaseResponseBodyData) SetSize(v int32) *ListMPCoSPhaseResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMPCoSPhaseResponseBodyData) SetTotal(v int32) *ListMPCoSPhaseResponseBodyData {
	s.Total = &v
	return s
}

type ListMPCoSPhaseResponseBodyDataPageData struct {
	AccessPermission *int32  `json:"AccessPermission,omitempty" xml:"AccessPermission,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhaseId          *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	Remark           *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ListMPCoSPhaseResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseResponseBodyDataPageData) SetAccessPermission(v int32) *ListMPCoSPhaseResponseBodyDataPageData {
	s.AccessPermission = &v
	return s
}

func (s *ListMPCoSPhaseResponseBodyDataPageData) SetName(v string) *ListMPCoSPhaseResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListMPCoSPhaseResponseBodyDataPageData) SetPhaseId(v string) *ListMPCoSPhaseResponseBodyDataPageData {
	s.PhaseId = &v
	return s
}

func (s *ListMPCoSPhaseResponseBodyDataPageData) SetRemark(v string) *ListMPCoSPhaseResponseBodyDataPageData {
	s.Remark = &v
	return s
}

type ListMPCoSPhaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMPCoSPhaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMPCoSPhaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseResponse) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseResponse) SetHeaders(v map[string]*string) *ListMPCoSPhaseResponse {
	s.Headers = v
	return s
}

func (s *ListMPCoSPhaseResponse) SetStatusCode(v int32) *ListMPCoSPhaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMPCoSPhaseResponse) SetBody(v *ListMPCoSPhaseResponseBody) *ListMPCoSPhaseResponse {
	s.Body = v
	return s
}

type ListMPCoSPhaseGroupRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListMPCoSPhaseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseGroupRequest) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseGroupRequest) SetApiVersion(v string) *ListMPCoSPhaseGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListMPCoSPhaseGroupRequest) SetBizChainId(v string) *ListMPCoSPhaseGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *ListMPCoSPhaseGroupRequest) SetName(v string) *ListMPCoSPhaseGroupRequest {
	s.Name = &v
	return s
}

func (s *ListMPCoSPhaseGroupRequest) SetNum(v int32) *ListMPCoSPhaseGroupRequest {
	s.Num = &v
	return s
}

func (s *ListMPCoSPhaseGroupRequest) SetSize(v int32) *ListMPCoSPhaseGroupRequest {
	s.Size = &v
	return s
}

type ListMPCoSPhaseGroupResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListMPCoSPhaseGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMPCoSPhaseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseGroupResponseBody) SetCode(v int32) *ListMPCoSPhaseGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBody) SetData(v *ListMPCoSPhaseGroupResponseBodyData) *ListMPCoSPhaseGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBody) SetMessage(v string) *ListMPCoSPhaseGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBody) SetRequestId(v string) *ListMPCoSPhaseGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBody) SetSuccess(v bool) *ListMPCoSPhaseGroupResponseBody {
	s.Success = &v
	return s
}

type ListMPCoSPhaseGroupResponseBodyData struct {
	Num      *int32                                         `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMPCoSPhaseGroupResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMPCoSPhaseGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseGroupResponseBodyData) SetNum(v int32) *ListMPCoSPhaseGroupResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBodyData) SetPageData(v []*ListMPCoSPhaseGroupResponseBodyDataPageData) *ListMPCoSPhaseGroupResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBodyData) SetSize(v int32) *ListMPCoSPhaseGroupResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBodyData) SetTotal(v int32) *ListMPCoSPhaseGroupResponseBodyData {
	s.Total = &v
	return s
}

type ListMPCoSPhaseGroupResponseBodyDataPageData struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ListMPCoSPhaseGroupResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseGroupResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseGroupResponseBodyDataPageData) SetName(v string) *ListMPCoSPhaseGroupResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBodyDataPageData) SetPhaseGroupId(v string) *ListMPCoSPhaseGroupResponseBodyDataPageData {
	s.PhaseGroupId = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponseBodyDataPageData) SetRemark(v string) *ListMPCoSPhaseGroupResponseBodyDataPageData {
	s.Remark = &v
	return s
}

type ListMPCoSPhaseGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMPCoSPhaseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMPCoSPhaseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseGroupResponse) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseGroupResponse) SetHeaders(v map[string]*string) *ListMPCoSPhaseGroupResponse {
	s.Headers = v
	return s
}

func (s *ListMPCoSPhaseGroupResponse) SetStatusCode(v int32) *ListMPCoSPhaseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMPCoSPhaseGroupResponse) SetBody(v *ListMPCoSPhaseGroupResponseBody) *ListMPCoSPhaseGroupResponse {
	s.Body = v
	return s
}

type ListMPCoSPhaseHistoryRequest struct {
	ApiVersion   *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey      *string `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Num          *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId      *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListMPCoSPhaseHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseHistoryRequest) SetApiVersion(v string) *ListMPCoSPhaseHistoryRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetBizChainId(v string) *ListMPCoSPhaseHistoryRequest {
	s.BizChainId = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetDataKey(v string) *ListMPCoSPhaseHistoryRequest {
	s.DataKey = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetEndTime(v int64) *ListMPCoSPhaseHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetNum(v int32) *ListMPCoSPhaseHistoryRequest {
	s.Num = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetPhaseGroupId(v string) *ListMPCoSPhaseHistoryRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetPhaseId(v string) *ListMPCoSPhaseHistoryRequest {
	s.PhaseId = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetSize(v int32) *ListMPCoSPhaseHistoryRequest {
	s.Size = &v
	return s
}

func (s *ListMPCoSPhaseHistoryRequest) SetStartTime(v int64) *ListMPCoSPhaseHistoryRequest {
	s.StartTime = &v
	return s
}

type ListMPCoSPhaseHistoryResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListMPCoSPhaseHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMPCoSPhaseHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseHistoryResponseBody) SetCode(v int32) *ListMPCoSPhaseHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBody) SetData(v *ListMPCoSPhaseHistoryResponseBodyData) *ListMPCoSPhaseHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBody) SetMessage(v string) *ListMPCoSPhaseHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBody) SetRequestId(v string) *ListMPCoSPhaseHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBody) SetSuccess(v bool) *ListMPCoSPhaseHistoryResponseBody {
	s.Success = &v
	return s
}

type ListMPCoSPhaseHistoryResponseBodyData struct {
	Num      *int32                                           `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMPCoSPhaseHistoryResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMPCoSPhaseHistoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseHistoryResponseBodyData) SetNum(v int32) *ListMPCoSPhaseHistoryResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyData) SetPageData(v []*ListMPCoSPhaseHistoryResponseBodyDataPageData) *ListMPCoSPhaseHistoryResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyData) SetSize(v int32) *ListMPCoSPhaseHistoryResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyData) SetTotal(v int32) *ListMPCoSPhaseHistoryResponseBodyData {
	s.Total = &v
	return s
}

type ListMPCoSPhaseHistoryResponseBodyDataPageData struct {
	BlockHash       *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	BlockNumber     *int64  `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	DataHash        *string `json:"DataHash,omitempty" xml:"DataHash,omitempty"`
	DataSeq         *string `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	DataValue       *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
	IotId           *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PreviousHash    *string `json:"PreviousHash,omitempty" xml:"PreviousHash,omitempty"`
	ProductKey      *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Timestamp       *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TransactionHash *string `json:"TransactionHash,omitempty" xml:"TransactionHash,omitempty"`
}

func (s ListMPCoSPhaseHistoryResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseHistoryResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetBlockHash(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.BlockHash = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetBlockNumber(v int64) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.BlockNumber = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetDataHash(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.DataHash = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetDataSeq(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.DataSeq = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetDataValue(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.DataValue = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetIotId(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetPreviousHash(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.PreviousHash = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetProductKey(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.ProductKey = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetTimestamp(v int64) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.Timestamp = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponseBodyDataPageData) SetTransactionHash(v string) *ListMPCoSPhaseHistoryResponseBodyDataPageData {
	s.TransactionHash = &v
	return s
}

type ListMPCoSPhaseHistoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMPCoSPhaseHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMPCoSPhaseHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMPCoSPhaseHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListMPCoSPhaseHistoryResponse) SetHeaders(v map[string]*string) *ListMPCoSPhaseHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListMPCoSPhaseHistoryResponse) SetStatusCode(v int32) *ListMPCoSPhaseHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMPCoSPhaseHistoryResponse) SetBody(v *ListMPCoSPhaseHistoryResponseBody) *ListMPCoSPhaseHistoryResponse {
	s.Body = v
	return s
}

type ListMemberRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberUid  *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMemberRequest) GoString() string {
	return s.String()
}

func (s *ListMemberRequest) SetApiVersion(v string) *ListMemberRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListMemberRequest) SetBizChainId(v string) *ListMemberRequest {
	s.BizChainId = &v
	return s
}

func (s *ListMemberRequest) SetMemberUid(v string) *ListMemberRequest {
	s.MemberUid = &v
	return s
}

func (s *ListMemberRequest) SetNum(v int32) *ListMemberRequest {
	s.Num = &v
	return s
}

func (s *ListMemberRequest) SetSize(v int32) *ListMemberRequest {
	s.Size = &v
	return s
}

type ListMemberResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemberResponseBody) SetCode(v int32) *ListMemberResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemberResponseBody) SetData(v *ListMemberResponseBodyData) *ListMemberResponseBody {
	s.Data = v
	return s
}

func (s *ListMemberResponseBody) SetMessage(v string) *ListMemberResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemberResponseBody) SetRequestId(v string) *ListMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemberResponseBody) SetSuccess(v bool) *ListMemberResponseBody {
	s.Success = &v
	return s
}

type ListMemberResponseBodyData struct {
	Num      *int32                                `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMemberResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMemberResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemberResponseBodyData) SetNum(v int32) *ListMemberResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMemberResponseBodyData) SetPageData(v []*ListMemberResponseBodyDataPageData) *ListMemberResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMemberResponseBodyData) SetSize(v int32) *ListMemberResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMemberResponseBodyData) SetTotal(v int32) *ListMemberResponseBodyData {
	s.Total = &v
	return s
}

type ListMemberResponseBodyDataPageData struct {
	MemberContact *string `json:"MemberContact,omitempty" xml:"MemberContact,omitempty"`
	MemberId      *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName    *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberPhone   *string `json:"MemberPhone,omitempty" xml:"MemberPhone,omitempty"`
	MemberUid     *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMemberResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListMemberResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMemberResponseBodyDataPageData) SetMemberContact(v string) *ListMemberResponseBodyDataPageData {
	s.MemberContact = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetMemberId(v string) *ListMemberResponseBodyDataPageData {
	s.MemberId = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetMemberName(v string) *ListMemberResponseBodyDataPageData {
	s.MemberName = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetMemberPhone(v string) *ListMemberResponseBodyDataPageData {
	s.MemberPhone = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetMemberUid(v string) *ListMemberResponseBodyDataPageData {
	s.MemberUid = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetRemark(v string) *ListMemberResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetStatus(v string) *ListMemberResponseBodyDataPageData {
	s.Status = &v
	return s
}

type ListMemberResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMemberResponse) GoString() string {
	return s.String()
}

func (s *ListMemberResponse) SetHeaders(v map[string]*string) *ListMemberResponse {
	s.Headers = v
	return s
}

func (s *ListMemberResponse) SetStatusCode(v int32) *ListMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemberResponse) SetBody(v *ListMemberResponseBody) *ListMemberResponse {
	s.Body = v
	return s
}

type ListMultiPartyCollaborationChainRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListMultiPartyCollaborationChainRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiPartyCollaborationChainRequest) GoString() string {
	return s.String()
}

func (s *ListMultiPartyCollaborationChainRequest) SetApiVersion(v string) *ListMultiPartyCollaborationChainRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListMultiPartyCollaborationChainRequest) SetName(v string) *ListMultiPartyCollaborationChainRequest {
	s.Name = &v
	return s
}

func (s *ListMultiPartyCollaborationChainRequest) SetNum(v int32) *ListMultiPartyCollaborationChainRequest {
	s.Num = &v
	return s
}

func (s *ListMultiPartyCollaborationChainRequest) SetSize(v int32) *ListMultiPartyCollaborationChainRequest {
	s.Size = &v
	return s
}

type ListMultiPartyCollaborationChainResponseBody struct {
	Code      *int32                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListMultiPartyCollaborationChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMultiPartyCollaborationChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiPartyCollaborationChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiPartyCollaborationChainResponseBody) SetCode(v int32) *ListMultiPartyCollaborationChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBody) SetData(v *ListMultiPartyCollaborationChainResponseBodyData) *ListMultiPartyCollaborationChainResponseBody {
	s.Data = v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBody) SetMessage(v string) *ListMultiPartyCollaborationChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBody) SetRequestId(v string) *ListMultiPartyCollaborationChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBody) SetSuccess(v bool) *ListMultiPartyCollaborationChainResponseBody {
	s.Success = &v
	return s
}

type ListMultiPartyCollaborationChainResponseBodyData struct {
	Num      *int32                                                      `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMultiPartyCollaborationChainResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                                      `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMultiPartyCollaborationChainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMultiPartyCollaborationChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMultiPartyCollaborationChainResponseBodyData) SetNum(v int32) *ListMultiPartyCollaborationChainResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBodyData) SetPageData(v []*ListMultiPartyCollaborationChainResponseBodyDataPageData) *ListMultiPartyCollaborationChainResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBodyData) SetSize(v int32) *ListMultiPartyCollaborationChainResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBodyData) SetTotal(v int32) *ListMultiPartyCollaborationChainResponseBodyData {
	s.Total = &v
	return s
}

type ListMultiPartyCollaborationChainResponseBodyDataPageData struct {
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RoleType   *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListMultiPartyCollaborationChainResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListMultiPartyCollaborationChainResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMultiPartyCollaborationChainResponseBodyDataPageData) SetBizChainId(v string) *ListMultiPartyCollaborationChainResponseBodyDataPageData {
	s.BizChainId = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBodyDataPageData) SetName(v string) *ListMultiPartyCollaborationChainResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBodyDataPageData) SetRemark(v string) *ListMultiPartyCollaborationChainResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponseBodyDataPageData) SetRoleType(v string) *ListMultiPartyCollaborationChainResponseBodyDataPageData {
	s.RoleType = &v
	return s
}

type ListMultiPartyCollaborationChainResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMultiPartyCollaborationChainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMultiPartyCollaborationChainResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiPartyCollaborationChainResponse) GoString() string {
	return s.String()
}

func (s *ListMultiPartyCollaborationChainResponse) SetHeaders(v map[string]*string) *ListMultiPartyCollaborationChainResponse {
	s.Headers = v
	return s
}

func (s *ListMultiPartyCollaborationChainResponse) SetStatusCode(v int32) *ListMultiPartyCollaborationChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiPartyCollaborationChainResponse) SetBody(v *ListMultiPartyCollaborationChainResponseBody) *ListMultiPartyCollaborationChainResponse {
	s.Body = v
	return s
}

type ListPSMemberDataTypeCodeRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberUid  *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListPSMemberDataTypeCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPSMemberDataTypeCodeRequest) GoString() string {
	return s.String()
}

func (s *ListPSMemberDataTypeCodeRequest) SetApiVersion(v string) *ListPSMemberDataTypeCodeRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListPSMemberDataTypeCodeRequest) SetBizChainId(v string) *ListPSMemberDataTypeCodeRequest {
	s.BizChainId = &v
	return s
}

func (s *ListPSMemberDataTypeCodeRequest) SetMemberUid(v string) *ListPSMemberDataTypeCodeRequest {
	s.MemberUid = &v
	return s
}

func (s *ListPSMemberDataTypeCodeRequest) SetNum(v int32) *ListPSMemberDataTypeCodeRequest {
	s.Num = &v
	return s
}

func (s *ListPSMemberDataTypeCodeRequest) SetSize(v int32) *ListPSMemberDataTypeCodeRequest {
	s.Size = &v
	return s
}

type ListPSMemberDataTypeCodeResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListPSMemberDataTypeCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPSMemberDataTypeCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPSMemberDataTypeCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListPSMemberDataTypeCodeResponseBody) SetCode(v int32) *ListPSMemberDataTypeCodeResponseBody {
	s.Code = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBody) SetData(v *ListPSMemberDataTypeCodeResponseBodyData) *ListPSMemberDataTypeCodeResponseBody {
	s.Data = v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBody) SetMessage(v string) *ListPSMemberDataTypeCodeResponseBody {
	s.Message = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBody) SetRequestId(v string) *ListPSMemberDataTypeCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBody) SetSuccess(v bool) *ListPSMemberDataTypeCodeResponseBody {
	s.Success = &v
	return s
}

type ListPSMemberDataTypeCodeResponseBodyData struct {
	Num      *int32                                              `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListPSMemberDataTypeCodeResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                              `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPSMemberDataTypeCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPSMemberDataTypeCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPSMemberDataTypeCodeResponseBodyData) SetNum(v int32) *ListPSMemberDataTypeCodeResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBodyData) SetPageData(v []*ListPSMemberDataTypeCodeResponseBodyDataPageData) *ListPSMemberDataTypeCodeResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBodyData) SetSize(v int32) *ListPSMemberDataTypeCodeResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBodyData) SetTotal(v int32) *ListPSMemberDataTypeCodeResponseBodyData {
	s.Total = &v
	return s
}

type ListPSMemberDataTypeCodeResponseBodyDataPageData struct {
	DataTypeCode *string `json:"DataTypeCode,omitempty" xml:"DataTypeCode,omitempty"`
	MemberId     *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberUid    *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s ListPSMemberDataTypeCodeResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListPSMemberDataTypeCodeResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListPSMemberDataTypeCodeResponseBodyDataPageData) SetDataTypeCode(v string) *ListPSMemberDataTypeCodeResponseBodyDataPageData {
	s.DataTypeCode = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBodyDataPageData) SetMemberId(v string) *ListPSMemberDataTypeCodeResponseBodyDataPageData {
	s.MemberId = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBodyDataPageData) SetMemberName(v string) *ListPSMemberDataTypeCodeResponseBodyDataPageData {
	s.MemberName = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponseBodyDataPageData) SetMemberUid(v string) *ListPSMemberDataTypeCodeResponseBodyDataPageData {
	s.MemberUid = &v
	return s
}

type ListPSMemberDataTypeCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPSMemberDataTypeCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPSMemberDataTypeCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPSMemberDataTypeCodeResponse) GoString() string {
	return s.String()
}

func (s *ListPSMemberDataTypeCodeResponse) SetHeaders(v map[string]*string) *ListPSMemberDataTypeCodeResponse {
	s.Headers = v
	return s
}

func (s *ListPSMemberDataTypeCodeResponse) SetStatusCode(v int32) *ListPSMemberDataTypeCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPSMemberDataTypeCodeResponse) SetBody(v *ListPSMemberDataTypeCodeResponseBody) *ListPSMemberDataTypeCodeResponse {
	s.Body = v
	return s
}

type ListProofChainRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListProofChainRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProofChainRequest) GoString() string {
	return s.String()
}

func (s *ListProofChainRequest) SetApiVersion(v string) *ListProofChainRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListProofChainRequest) SetName(v string) *ListProofChainRequest {
	s.Name = &v
	return s
}

func (s *ListProofChainRequest) SetNum(v int32) *ListProofChainRequest {
	s.Num = &v
	return s
}

func (s *ListProofChainRequest) SetSize(v int32) *ListProofChainRequest {
	s.Size = &v
	return s
}

type ListProofChainResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListProofChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProofChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProofChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListProofChainResponseBody) SetCode(v int32) *ListProofChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListProofChainResponseBody) SetData(v *ListProofChainResponseBodyData) *ListProofChainResponseBody {
	s.Data = v
	return s
}

func (s *ListProofChainResponseBody) SetMessage(v string) *ListProofChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListProofChainResponseBody) SetRequestId(v string) *ListProofChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProofChainResponseBody) SetSuccess(v bool) *ListProofChainResponseBody {
	s.Success = &v
	return s
}

type ListProofChainResponseBodyData struct {
	Num      *int32                                    `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListProofChainResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListProofChainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProofChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProofChainResponseBodyData) SetNum(v int32) *ListProofChainResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListProofChainResponseBodyData) SetPageData(v []*ListProofChainResponseBodyDataPageData) *ListProofChainResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListProofChainResponseBodyData) SetSize(v int32) *ListProofChainResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListProofChainResponseBodyData) SetTotal(v int32) *ListProofChainResponseBodyData {
	s.Total = &v
	return s
}

type ListProofChainResponseBodyDataPageData struct {
	BizChainCode *string `json:"BizChainCode,omitempty" xml:"BizChainCode,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataTypeCode *string `json:"DataTypeCode,omitempty" xml:"DataTypeCode,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RoleType     *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListProofChainResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListProofChainResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListProofChainResponseBodyDataPageData) SetBizChainCode(v string) *ListProofChainResponseBodyDataPageData {
	s.BizChainCode = &v
	return s
}

func (s *ListProofChainResponseBodyDataPageData) SetBizChainId(v string) *ListProofChainResponseBodyDataPageData {
	s.BizChainId = &v
	return s
}

func (s *ListProofChainResponseBodyDataPageData) SetDataTypeCode(v string) *ListProofChainResponseBodyDataPageData {
	s.DataTypeCode = &v
	return s
}

func (s *ListProofChainResponseBodyDataPageData) SetName(v string) *ListProofChainResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListProofChainResponseBodyDataPageData) SetRemark(v string) *ListProofChainResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListProofChainResponseBodyDataPageData) SetRoleType(v string) *ListProofChainResponseBodyDataPageData {
	s.RoleType = &v
	return s
}

type ListProofChainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProofChainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProofChainResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProofChainResponse) GoString() string {
	return s.String()
}

func (s *ListProofChainResponse) SetHeaders(v map[string]*string) *ListProofChainResponse {
	s.Headers = v
	return s
}

func (s *ListProofChainResponse) SetStatusCode(v int32) *ListProofChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProofChainResponse) SetBody(v *ListProofChainResponseBody) *ListProofChainResponse {
	s.Body = v
	return s
}

type LockMemberRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s LockMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s LockMemberRequest) GoString() string {
	return s.String()
}

func (s *LockMemberRequest) SetApiVersion(v string) *LockMemberRequest {
	s.ApiVersion = &v
	return s
}

func (s *LockMemberRequest) SetBizChainId(v string) *LockMemberRequest {
	s.BizChainId = &v
	return s
}

func (s *LockMemberRequest) SetMemberId(v string) *LockMemberRequest {
	s.MemberId = &v
	return s
}

type LockMemberResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LockMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockMemberResponseBody) GoString() string {
	return s.String()
}

func (s *LockMemberResponseBody) SetCode(v int32) *LockMemberResponseBody {
	s.Code = &v
	return s
}

func (s *LockMemberResponseBody) SetData(v string) *LockMemberResponseBody {
	s.Data = &v
	return s
}

func (s *LockMemberResponseBody) SetMessage(v string) *LockMemberResponseBody {
	s.Message = &v
	return s
}

func (s *LockMemberResponseBody) SetRequestId(v string) *LockMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockMemberResponseBody) SetSuccess(v bool) *LockMemberResponseBody {
	s.Success = &v
	return s
}

type LockMemberResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LockMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s LockMemberResponse) GoString() string {
	return s.String()
}

func (s *LockMemberResponse) SetHeaders(v map[string]*string) *LockMemberResponse {
	s.Headers = v
	return s
}

func (s *LockMemberResponse) SetStatusCode(v int32) *LockMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *LockMemberResponse) SetBody(v *LockMemberResponseBody) *LockMemberResponse {
	s.Body = v
	return s
}

type ModifyMPCoSPhaseRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhaseId    *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyMPCoSPhaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPCoSPhaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyMPCoSPhaseRequest) SetApiVersion(v string) *ModifyMPCoSPhaseRequest {
	s.ApiVersion = &v
	return s
}

func (s *ModifyMPCoSPhaseRequest) SetBizChainId(v string) *ModifyMPCoSPhaseRequest {
	s.BizChainId = &v
	return s
}

func (s *ModifyMPCoSPhaseRequest) SetName(v string) *ModifyMPCoSPhaseRequest {
	s.Name = &v
	return s
}

func (s *ModifyMPCoSPhaseRequest) SetPhaseId(v string) *ModifyMPCoSPhaseRequest {
	s.PhaseId = &v
	return s
}

func (s *ModifyMPCoSPhaseRequest) SetRemark(v string) *ModifyMPCoSPhaseRequest {
	s.Remark = &v
	return s
}

type ModifyMPCoSPhaseResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMPCoSPhaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPCoSPhaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMPCoSPhaseResponseBody) SetCode(v int32) *ModifyMPCoSPhaseResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMPCoSPhaseResponseBody) SetData(v string) *ModifyMPCoSPhaseResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyMPCoSPhaseResponseBody) SetMessage(v string) *ModifyMPCoSPhaseResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMPCoSPhaseResponseBody) SetRequestId(v string) *ModifyMPCoSPhaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMPCoSPhaseResponseBody) SetSuccess(v bool) *ModifyMPCoSPhaseResponseBody {
	s.Success = &v
	return s
}

type ModifyMPCoSPhaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMPCoSPhaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMPCoSPhaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPCoSPhaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyMPCoSPhaseResponse) SetHeaders(v map[string]*string) *ModifyMPCoSPhaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyMPCoSPhaseResponse) SetStatusCode(v int32) *ModifyMPCoSPhaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMPCoSPhaseResponse) SetBody(v *ModifyMPCoSPhaseResponseBody) *ModifyMPCoSPhaseResponse {
	s.Body = v
	return s
}

type ModifyMPCoSPhaseGroupRequest struct {
	ApiVersion   *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhaseGroupId *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyMPCoSPhaseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPCoSPhaseGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyMPCoSPhaseGroupRequest) SetApiVersion(v string) *ModifyMPCoSPhaseGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupRequest) SetBizChainId(v string) *ModifyMPCoSPhaseGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupRequest) SetName(v string) *ModifyMPCoSPhaseGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupRequest) SetPhaseGroupId(v string) *ModifyMPCoSPhaseGroupRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupRequest) SetRemark(v string) *ModifyMPCoSPhaseGroupRequest {
	s.Remark = &v
	return s
}

type ModifyMPCoSPhaseGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMPCoSPhaseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPCoSPhaseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMPCoSPhaseGroupResponseBody) SetCode(v int32) *ModifyMPCoSPhaseGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupResponseBody) SetData(v string) *ModifyMPCoSPhaseGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupResponseBody) SetMessage(v string) *ModifyMPCoSPhaseGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupResponseBody) SetRequestId(v string) *ModifyMPCoSPhaseGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupResponseBody) SetSuccess(v bool) *ModifyMPCoSPhaseGroupResponseBody {
	s.Success = &v
	return s
}

type ModifyMPCoSPhaseGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMPCoSPhaseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMPCoSPhaseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPCoSPhaseGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyMPCoSPhaseGroupResponse) SetHeaders(v map[string]*string) *ModifyMPCoSPhaseGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyMPCoSPhaseGroupResponse) SetStatusCode(v int32) *ModifyMPCoSPhaseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMPCoSPhaseGroupResponse) SetBody(v *ModifyMPCoSPhaseGroupResponseBody) *ModifyMPCoSPhaseGroupResponse {
	s.Body = v
	return s
}

type ModifyMemberRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberContact *string `json:"MemberContact,omitempty" xml:"MemberContact,omitempty"`
	MemberId      *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName    *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberPhone   *string `json:"MemberPhone,omitempty" xml:"MemberPhone,omitempty"`
	MemberUid     *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberRequest) GoString() string {
	return s.String()
}

func (s *ModifyMemberRequest) SetApiVersion(v string) *ModifyMemberRequest {
	s.ApiVersion = &v
	return s
}

func (s *ModifyMemberRequest) SetBizChainId(v string) *ModifyMemberRequest {
	s.BizChainId = &v
	return s
}

func (s *ModifyMemberRequest) SetMemberContact(v string) *ModifyMemberRequest {
	s.MemberContact = &v
	return s
}

func (s *ModifyMemberRequest) SetMemberId(v string) *ModifyMemberRequest {
	s.MemberId = &v
	return s
}

func (s *ModifyMemberRequest) SetMemberName(v string) *ModifyMemberRequest {
	s.MemberName = &v
	return s
}

func (s *ModifyMemberRequest) SetMemberPhone(v string) *ModifyMemberRequest {
	s.MemberPhone = &v
	return s
}

func (s *ModifyMemberRequest) SetMemberUid(v string) *ModifyMemberRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyMemberRequest) SetRemark(v string) *ModifyMemberRequest {
	s.Remark = &v
	return s
}

type ModifyMemberResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMemberResponseBody) SetCode(v int32) *ModifyMemberResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMemberResponseBody) SetData(v string) *ModifyMemberResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyMemberResponseBody) SetMessage(v string) *ModifyMemberResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMemberResponseBody) SetRequestId(v string) *ModifyMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMemberResponseBody) SetSuccess(v bool) *ModifyMemberResponseBody {
	s.Success = &v
	return s
}

type ModifyMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberResponse) GoString() string {
	return s.String()
}

func (s *ModifyMemberResponse) SetHeaders(v map[string]*string) *ModifyMemberResponse {
	s.Headers = v
	return s
}

func (s *ModifyMemberResponse) SetStatusCode(v int32) *ModifyMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMemberResponse) SetBody(v *ModifyMemberResponseBody) *ModifyMemberResponse {
	s.Body = v
	return s
}

type RegisterDeviceGroupRequest struct {
	ApiVersion      *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	AuthorizeType   *string `json:"AuthorizeType,omitempty" xml:"AuthorizeType,omitempty"`
	BizChainId      *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	ProductKey      *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s RegisterDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceGroupRequest) SetApiVersion(v string) *RegisterDeviceGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *RegisterDeviceGroupRequest) SetAuthorizeType(v string) *RegisterDeviceGroupRequest {
	s.AuthorizeType = &v
	return s
}

func (s *RegisterDeviceGroupRequest) SetBizChainId(v string) *RegisterDeviceGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *RegisterDeviceGroupRequest) SetDeviceGroupName(v string) *RegisterDeviceGroupRequest {
	s.DeviceGroupName = &v
	return s
}

func (s *RegisterDeviceGroupRequest) SetProductKey(v string) *RegisterDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *RegisterDeviceGroupRequest) SetRemark(v string) *RegisterDeviceGroupRequest {
	s.Remark = &v
	return s
}

type RegisterDeviceGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceGroupResponseBody) SetCode(v int32) *RegisterDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterDeviceGroupResponseBody) SetData(v string) *RegisterDeviceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *RegisterDeviceGroupResponseBody) SetMessage(v string) *RegisterDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceGroupResponseBody) SetRequestId(v string) *RegisterDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceGroupResponseBody) SetSuccess(v bool) *RegisterDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type RegisterDeviceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceGroupResponse) SetHeaders(v map[string]*string) *RegisterDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceGroupResponse) SetStatusCode(v int32) *RegisterDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceGroupResponse) SetBody(v *RegisterDeviceGroupResponseBody) *RegisterDeviceGroupResponse {
	s.Body = v
	return s
}

type SetDataRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataRequest) GoString() string {
	return s.String()
}

func (s *SetDataRequest) SetApiVersion(v string) *SetDataRequest {
	s.ApiVersion = &v
	return s
}

func (s *SetDataRequest) SetKey(v string) *SetDataRequest {
	s.Key = &v
	return s
}

func (s *SetDataRequest) SetProductKey(v string) *SetDataRequest {
	s.ProductKey = &v
	return s
}

func (s *SetDataRequest) SetValue(v string) *SetDataRequest {
	s.Value = &v
	return s
}

type SetDataResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataResponseBody) SetCode(v int32) *SetDataResponseBody {
	s.Code = &v
	return s
}

func (s *SetDataResponseBody) SetData(v string) *SetDataResponseBody {
	s.Data = &v
	return s
}

func (s *SetDataResponseBody) SetMessage(v string) *SetDataResponseBody {
	s.Message = &v
	return s
}

func (s *SetDataResponseBody) SetRequestId(v string) *SetDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataResponseBody) SetSuccess(v bool) *SetDataResponseBody {
	s.Success = &v
	return s
}

type SetDataResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataResponse) GoString() string {
	return s.String()
}

func (s *SetDataResponse) SetHeaders(v map[string]*string) *SetDataResponse {
	s.Headers = v
	return s
}

func (s *SetDataResponse) SetStatusCode(v int32) *SetDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataResponse) SetBody(v *SetDataResponseBody) *SetDataResponse {
	s.Body = v
	return s
}

type SetDataWithSignatureRequest struct {
	ApiVersion           *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	IotAuthType          *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProductKey           *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetDataWithSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataWithSignatureRequest) GoString() string {
	return s.String()
}

func (s *SetDataWithSignatureRequest) SetApiVersion(v string) *SetDataWithSignatureRequest {
	s.ApiVersion = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetIotAuthType(v string) *SetDataWithSignatureRequest {
	s.IotAuthType = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetIotDataDigest(v string) *SetDataWithSignatureRequest {
	s.IotDataDigest = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetIotId(v string) *SetDataWithSignatureRequest {
	s.IotId = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetIotIdServiceProvider(v string) *SetDataWithSignatureRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetIotIdSource(v string) *SetDataWithSignatureRequest {
	s.IotIdSource = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetIotSignature(v string) *SetDataWithSignatureRequest {
	s.IotSignature = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetKey(v string) *SetDataWithSignatureRequest {
	s.Key = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetProductKey(v string) *SetDataWithSignatureRequest {
	s.ProductKey = &v
	return s
}

func (s *SetDataWithSignatureRequest) SetValue(v string) *SetDataWithSignatureRequest {
	s.Value = &v
	return s
}

type SetDataWithSignatureResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataWithSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataWithSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataWithSignatureResponseBody) SetCode(v int32) *SetDataWithSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *SetDataWithSignatureResponseBody) SetData(v string) *SetDataWithSignatureResponseBody {
	s.Data = &v
	return s
}

func (s *SetDataWithSignatureResponseBody) SetMessage(v string) *SetDataWithSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *SetDataWithSignatureResponseBody) SetRequestId(v string) *SetDataWithSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataWithSignatureResponseBody) SetSuccess(v bool) *SetDataWithSignatureResponseBody {
	s.Success = &v
	return s
}

type SetDataWithSignatureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataWithSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataWithSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataWithSignatureResponse) GoString() string {
	return s.String()
}

func (s *SetDataWithSignatureResponse) SetHeaders(v map[string]*string) *SetDataWithSignatureResponse {
	s.Headers = v
	return s
}

func (s *SetDataWithSignatureResponse) SetStatusCode(v int32) *SetDataWithSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataWithSignatureResponse) SetBody(v *SetDataWithSignatureResponseBody) *SetDataWithSignatureResponse {
	s.Body = v
	return s
}

type UnAuthorizeDeviceRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s UnAuthorizeDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAuthorizeDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnAuthorizeDeviceRequest) SetApiVersion(v string) *UnAuthorizeDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *UnAuthorizeDeviceRequest) SetBizChainId(v string) *UnAuthorizeDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *UnAuthorizeDeviceRequest) SetDeviceGroupId(v string) *UnAuthorizeDeviceRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *UnAuthorizeDeviceRequest) SetDeviceId(v string) *UnAuthorizeDeviceRequest {
	s.DeviceId = &v
	return s
}

type UnAuthorizeDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnAuthorizeDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAuthorizeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnAuthorizeDeviceResponseBody) SetCode(v int32) *UnAuthorizeDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnAuthorizeDeviceResponseBody) SetData(v string) *UnAuthorizeDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *UnAuthorizeDeviceResponseBody) SetMessage(v string) *UnAuthorizeDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnAuthorizeDeviceResponseBody) SetRequestId(v string) *UnAuthorizeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnAuthorizeDeviceResponseBody) SetSuccess(v bool) *UnAuthorizeDeviceResponseBody {
	s.Success = &v
	return s
}

type UnAuthorizeDeviceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnAuthorizeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnAuthorizeDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAuthorizeDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnAuthorizeDeviceResponse) SetHeaders(v map[string]*string) *UnAuthorizeDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnAuthorizeDeviceResponse) SetStatusCode(v int32) *UnAuthorizeDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAuthorizeDeviceResponse) SetBody(v *UnAuthorizeDeviceResponseBody) *UnAuthorizeDeviceResponse {
	s.Body = v
	return s
}

type UnAuthorizeDeviceGroupRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId    *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s UnAuthorizeDeviceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAuthorizeDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *UnAuthorizeDeviceGroupRequest) SetApiVersion(v string) *UnAuthorizeDeviceGroupRequest {
	s.ApiVersion = &v
	return s
}

func (s *UnAuthorizeDeviceGroupRequest) SetBizChainId(v string) *UnAuthorizeDeviceGroupRequest {
	s.BizChainId = &v
	return s
}

func (s *UnAuthorizeDeviceGroupRequest) SetDeviceGroupId(v string) *UnAuthorizeDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

type UnAuthorizeDeviceGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnAuthorizeDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAuthorizeDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnAuthorizeDeviceGroupResponseBody) SetCode(v int32) *UnAuthorizeDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UnAuthorizeDeviceGroupResponseBody) SetData(v string) *UnAuthorizeDeviceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *UnAuthorizeDeviceGroupResponseBody) SetMessage(v string) *UnAuthorizeDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UnAuthorizeDeviceGroupResponseBody) SetRequestId(v string) *UnAuthorizeDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnAuthorizeDeviceGroupResponseBody) SetSuccess(v bool) *UnAuthorizeDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type UnAuthorizeDeviceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnAuthorizeDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnAuthorizeDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAuthorizeDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *UnAuthorizeDeviceGroupResponse) SetHeaders(v map[string]*string) *UnAuthorizeDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *UnAuthorizeDeviceGroupResponse) SetStatusCode(v int32) *UnAuthorizeDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAuthorizeDeviceGroupResponse) SetBody(v *UnAuthorizeDeviceGroupResponseBody) *UnAuthorizeDeviceGroupResponse {
	s.Body = v
	return s
}

type UnLockMemberRequest struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s UnLockMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UnLockMemberRequest) GoString() string {
	return s.String()
}

func (s *UnLockMemberRequest) SetApiVersion(v string) *UnLockMemberRequest {
	s.ApiVersion = &v
	return s
}

func (s *UnLockMemberRequest) SetBizChainId(v string) *UnLockMemberRequest {
	s.BizChainId = &v
	return s
}

func (s *UnLockMemberRequest) SetMemberId(v string) *UnLockMemberRequest {
	s.MemberId = &v
	return s
}

type UnLockMemberResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnLockMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnLockMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UnLockMemberResponseBody) SetCode(v int32) *UnLockMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UnLockMemberResponseBody) SetData(v string) *UnLockMemberResponseBody {
	s.Data = &v
	return s
}

func (s *UnLockMemberResponseBody) SetMessage(v string) *UnLockMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UnLockMemberResponseBody) SetRequestId(v string) *UnLockMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnLockMemberResponseBody) SetSuccess(v bool) *UnLockMemberResponseBody {
	s.Success = &v
	return s
}

type UnLockMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnLockMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnLockMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UnLockMemberResponse) GoString() string {
	return s.String()
}

func (s *UnLockMemberResponse) SetHeaders(v map[string]*string) *UnLockMemberResponse {
	s.Headers = v
	return s
}

func (s *UnLockMemberResponse) SetStatusCode(v int32) *UnLockMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UnLockMemberResponse) SetBody(v *UnLockMemberResponseBody) *UnLockMemberResponse {
	s.Body = v
	return s
}

type UpdateMPCoSAuthorizedInfoRequest struct {
	ApiVersion          *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	AuthorizedPhaseList map[string]interface{} `json:"AuthorizedPhaseList,omitempty" xml:"AuthorizedPhaseList,omitempty"`
	BizChainId          *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberId            *string                `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PhaseGroupId        *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
}

func (s UpdateMPCoSAuthorizedInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPCoSAuthorizedInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateMPCoSAuthorizedInfoRequest) SetApiVersion(v string) *UpdateMPCoSAuthorizedInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoRequest) SetAuthorizedPhaseList(v map[string]interface{}) *UpdateMPCoSAuthorizedInfoRequest {
	s.AuthorizedPhaseList = v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoRequest) SetBizChainId(v string) *UpdateMPCoSAuthorizedInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoRequest) SetMemberId(v string) *UpdateMPCoSAuthorizedInfoRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoRequest) SetPhaseGroupId(v string) *UpdateMPCoSAuthorizedInfoRequest {
	s.PhaseGroupId = &v
	return s
}

type UpdateMPCoSAuthorizedInfoShrinkRequest struct {
	ApiVersion                *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	AuthorizedPhaseListShrink *string `json:"AuthorizedPhaseList,omitempty" xml:"AuthorizedPhaseList,omitempty"`
	BizChainId                *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	MemberId                  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PhaseGroupId              *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
}

func (s UpdateMPCoSAuthorizedInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPCoSAuthorizedInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMPCoSAuthorizedInfoShrinkRequest) SetApiVersion(v string) *UpdateMPCoSAuthorizedInfoShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoShrinkRequest) SetAuthorizedPhaseListShrink(v string) *UpdateMPCoSAuthorizedInfoShrinkRequest {
	s.AuthorizedPhaseListShrink = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoShrinkRequest) SetBizChainId(v string) *UpdateMPCoSAuthorizedInfoShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoShrinkRequest) SetMemberId(v string) *UpdateMPCoSAuthorizedInfoShrinkRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoShrinkRequest) SetPhaseGroupId(v string) *UpdateMPCoSAuthorizedInfoShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

type UpdateMPCoSAuthorizedInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMPCoSAuthorizedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPCoSAuthorizedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMPCoSAuthorizedInfoResponseBody) SetCode(v int32) *UpdateMPCoSAuthorizedInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoResponseBody) SetData(v string) *UpdateMPCoSAuthorizedInfoResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoResponseBody) SetMessage(v string) *UpdateMPCoSAuthorizedInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoResponseBody) SetRequestId(v string) *UpdateMPCoSAuthorizedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoResponseBody) SetSuccess(v bool) *UpdateMPCoSAuthorizedInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateMPCoSAuthorizedInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMPCoSAuthorizedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMPCoSAuthorizedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPCoSAuthorizedInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateMPCoSAuthorizedInfoResponse) SetHeaders(v map[string]*string) *UpdateMPCoSAuthorizedInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoResponse) SetStatusCode(v int32) *UpdateMPCoSAuthorizedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMPCoSAuthorizedInfoResponse) SetBody(v *UpdateMPCoSAuthorizedInfoResponseBody) *UpdateMPCoSAuthorizedInfoResponse {
	s.Body = v
	return s
}

type UploadMPCoSPhaseDigestInfoRequest struct {
	ApiVersion      *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId      *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey         *string                `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq         *string                `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	PhaseData       *string                `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId    *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId         *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataList map[string]interface{} `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseDigestInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetApiVersion(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetBizChainId(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetDataKey(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetDataSeq(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetPhaseData(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetPhaseId(v string) *UploadMPCoSPhaseDigestInfoRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoRequest) SetRelatedDataList(v map[string]interface{}) *UploadMPCoSPhaseDigestInfoRequest {
	s.RelatedDataList = v
	return s
}

type UploadMPCoSPhaseDigestInfoShrinkRequest struct {
	ApiVersion            *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId            *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey               *string `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq               *string `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	PhaseData             *string `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId          *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId               *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataListShrink *string `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseDigestInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetApiVersion(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetBizChainId(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetDataKey(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetDataSeq(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetPhaseData(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetPhaseId(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoShrinkRequest) SetRelatedDataListShrink(v string) *UploadMPCoSPhaseDigestInfoShrinkRequest {
	s.RelatedDataListShrink = &v
	return s
}

type UploadMPCoSPhaseDigestInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadMPCoSPhaseDigestInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoResponseBody) SetCode(v int32) *UploadMPCoSPhaseDigestInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoResponseBody) SetData(v string) *UploadMPCoSPhaseDigestInfoResponseBody {
	s.Data = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoResponseBody) SetMessage(v string) *UploadMPCoSPhaseDigestInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoResponseBody) SetRequestId(v string) *UploadMPCoSPhaseDigestInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoResponseBody) SetSuccess(v bool) *UploadMPCoSPhaseDigestInfoResponseBody {
	s.Success = &v
	return s
}

type UploadMPCoSPhaseDigestInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadMPCoSPhaseDigestInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadMPCoSPhaseDigestInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoResponse) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoResponse) SetHeaders(v map[string]*string) *UploadMPCoSPhaseDigestInfoResponse {
	s.Headers = v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoResponse) SetStatusCode(v int32) *UploadMPCoSPhaseDigestInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoResponse) SetBody(v *UploadMPCoSPhaseDigestInfoResponseBody) *UploadMPCoSPhaseDigestInfoResponse {
	s.Body = v
	return s
}

type UploadMPCoSPhaseDigestInfoByDeviceRequest struct {
	ApiVersion           *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId           *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey              *string                `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq              *string                `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	IotAuthType          *string                `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string                `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string                `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string                `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string                `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string                `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseData            *string                `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId         *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId              *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataList      map[string]interface{} `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseDigestInfoByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoByDeviceRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetApiVersion(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetBizChainId(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetDataKey(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetDataSeq(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotAuthType(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotAuthType = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotDataDigest(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotDataDigest = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotId(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotIdServiceProvider(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotIdSource(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotIdSource = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetIotSignature(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.IotSignature = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetPhaseData(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetPhaseId(v string) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceRequest) SetRelatedDataList(v map[string]interface{}) *UploadMPCoSPhaseDigestInfoByDeviceRequest {
	s.RelatedDataList = v
	return s
}

type UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest struct {
	ApiVersion            *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId            *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey               *string `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq               *string `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	IotAuthType           *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest         *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                 *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider  *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource           *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature          *string `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseData             *string `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId          *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId               *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataListShrink *string `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetApiVersion(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetBizChainId(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetDataKey(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetDataSeq(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotAuthType(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotAuthType = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotDataDigest(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotDataDigest = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotId(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotIdServiceProvider(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotIdSource(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotIdSource = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetIotSignature(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.IotSignature = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetPhaseData(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetPhaseId(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest) SetRelatedDataListShrink(v string) *UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest {
	s.RelatedDataListShrink = &v
	return s
}

type UploadMPCoSPhaseDigestInfoByDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadMPCoSPhaseDigestInfoByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetCode(v int32) *UploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetData(v string) *UploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetMessage(v string) *UploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetRequestId(v string) *UploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponseBody) SetSuccess(v bool) *UploadMPCoSPhaseDigestInfoByDeviceResponseBody {
	s.Success = &v
	return s
}

type UploadMPCoSPhaseDigestInfoByDeviceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadMPCoSPhaseDigestInfoByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadMPCoSPhaseDigestInfoByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseDigestInfoByDeviceResponse) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponse) SetHeaders(v map[string]*string) *UploadMPCoSPhaseDigestInfoByDeviceResponse {
	s.Headers = v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponse) SetStatusCode(v int32) *UploadMPCoSPhaseDigestInfoByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMPCoSPhaseDigestInfoByDeviceResponse) SetBody(v *UploadMPCoSPhaseDigestInfoByDeviceResponseBody) *UploadMPCoSPhaseDigestInfoByDeviceResponse {
	s.Body = v
	return s
}

type UploadMPCoSPhaseTextInfoRequest struct {
	ApiVersion      *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId      *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey         *string                `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq         *string                `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	PhaseData       *string                `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId    *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId         *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataList map[string]interface{} `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseTextInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetApiVersion(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetBizChainId(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetDataKey(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetDataSeq(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetPhaseData(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetPhaseId(v string) *UploadMPCoSPhaseTextInfoRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoRequest) SetRelatedDataList(v map[string]interface{}) *UploadMPCoSPhaseTextInfoRequest {
	s.RelatedDataList = v
	return s
}

type UploadMPCoSPhaseTextInfoShrinkRequest struct {
	ApiVersion            *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId            *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey               *string `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq               *string `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	PhaseData             *string `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId          *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId               *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataListShrink *string `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseTextInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetApiVersion(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetBizChainId(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetDataKey(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetDataSeq(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetPhaseData(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetPhaseId(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoShrinkRequest) SetRelatedDataListShrink(v string) *UploadMPCoSPhaseTextInfoShrinkRequest {
	s.RelatedDataListShrink = &v
	return s
}

type UploadMPCoSPhaseTextInfoResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadMPCoSPhaseTextInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoResponseBody) SetCode(v int32) *UploadMPCoSPhaseTextInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoResponseBody) SetData(v string) *UploadMPCoSPhaseTextInfoResponseBody {
	s.Data = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoResponseBody) SetMessage(v string) *UploadMPCoSPhaseTextInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoResponseBody) SetRequestId(v string) *UploadMPCoSPhaseTextInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoResponseBody) SetSuccess(v bool) *UploadMPCoSPhaseTextInfoResponseBody {
	s.Success = &v
	return s
}

type UploadMPCoSPhaseTextInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadMPCoSPhaseTextInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadMPCoSPhaseTextInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoResponse) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoResponse) SetHeaders(v map[string]*string) *UploadMPCoSPhaseTextInfoResponse {
	s.Headers = v
	return s
}

func (s *UploadMPCoSPhaseTextInfoResponse) SetStatusCode(v int32) *UploadMPCoSPhaseTextInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoResponse) SetBody(v *UploadMPCoSPhaseTextInfoResponseBody) *UploadMPCoSPhaseTextInfoResponse {
	s.Body = v
	return s
}

type UploadMPCoSPhaseTextInfoByDeviceRequest struct {
	ApiVersion           *string                `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId           *string                `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey              *string                `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq              *string                `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	IotAuthType          *string                `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest        *string                `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                *string                `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider *string                `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource          *string                `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature         *string                `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseData            *string                `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId         *string                `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId              *string                `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataList      map[string]interface{} `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseTextInfoByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoByDeviceRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetApiVersion(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetBizChainId(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetDataKey(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetDataSeq(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetIotAuthType(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotAuthType = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetIotDataDigest(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotDataDigest = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetIotId(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetIotIdServiceProvider(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetIotIdSource(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotIdSource = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetIotSignature(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.IotSignature = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetPhaseData(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetPhaseId(v string) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceRequest) SetRelatedDataList(v map[string]interface{}) *UploadMPCoSPhaseTextInfoByDeviceRequest {
	s.RelatedDataList = v
	return s
}

type UploadMPCoSPhaseTextInfoByDeviceShrinkRequest struct {
	ApiVersion            *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	BizChainId            *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	DataKey               *string `json:"DataKey,omitempty" xml:"DataKey,omitempty"`
	DataSeq               *string `json:"DataSeq,omitempty" xml:"DataSeq,omitempty"`
	IotAuthType           *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	IotDataDigest         *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	IotId                 *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IotIdServiceProvider  *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	IotIdSource           *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	IotSignature          *string `json:"IotSignature,omitempty" xml:"IotSignature,omitempty"`
	PhaseData             *string `json:"PhaseData,omitempty" xml:"PhaseData,omitempty"`
	PhaseGroupId          *string `json:"PhaseGroupId,omitempty" xml:"PhaseGroupId,omitempty"`
	PhaseId               *string `json:"PhaseId,omitempty" xml:"PhaseId,omitempty"`
	RelatedDataListShrink *string `json:"RelatedDataList,omitempty" xml:"RelatedDataList,omitempty"`
}

func (s UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetApiVersion(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetBizChainId(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.BizChainId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetDataKey(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.DataKey = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetDataSeq(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.DataSeq = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotAuthType(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotAuthType = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotDataDigest(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotDataDigest = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotId(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotIdServiceProvider(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotIdSource(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotIdSource = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetIotSignature(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.IotSignature = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetPhaseData(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.PhaseData = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetPhaseGroupId(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.PhaseGroupId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetPhaseId(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.PhaseId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest) SetRelatedDataListShrink(v string) *UploadMPCoSPhaseTextInfoByDeviceShrinkRequest {
	s.RelatedDataListShrink = &v
	return s
}

type UploadMPCoSPhaseTextInfoByDeviceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadMPCoSPhaseTextInfoByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponseBody) SetCode(v int32) *UploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponseBody) SetData(v string) *UploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Data = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponseBody) SetMessage(v string) *UploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponseBody) SetRequestId(v string) *UploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponseBody) SetSuccess(v bool) *UploadMPCoSPhaseTextInfoByDeviceResponseBody {
	s.Success = &v
	return s
}

type UploadMPCoSPhaseTextInfoByDeviceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadMPCoSPhaseTextInfoByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadMPCoSPhaseTextInfoByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadMPCoSPhaseTextInfoByDeviceResponse) GoString() string {
	return s.String()
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponse) SetHeaders(v map[string]*string) *UploadMPCoSPhaseTextInfoByDeviceResponse {
	s.Headers = v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponse) SetStatusCode(v int32) *UploadMPCoSPhaseTextInfoByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMPCoSPhaseTextInfoByDeviceResponse) SetBody(v *UploadMPCoSPhaseTextInfoByDeviceResponseBody) *UploadMPCoSPhaseTextInfoByDeviceResponse {
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
		"ap-northeast-1":              tea.String("ltl.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("ltl.aliyuncs.com"),
		"ap-south-1":                  tea.String("ltl.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ltl.aliyuncs.com"),
		"ap-southeast-2":              tea.String("ltl.aliyuncs.com"),
		"ap-southeast-3":              tea.String("ltl.aliyuncs.com"),
		"ap-southeast-5":              tea.String("ltl.aliyuncs.com"),
		"cn-beijing":                  tea.String("ltl.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ltl.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ltl.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ltl.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ltl.aliyuncs.com"),
		"cn-chengdu":                  tea.String("ltl.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ltl.aliyuncs.com"),
		"cn-fujian":                   tea.String("ltl.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ltl.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ltl.aliyuncs.com"),
		"cn-hongkong":                 tea.String("ltl.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ltl.aliyuncs.com"),
		"cn-huhehaote":                tea.String("ltl.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ltl.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("ltl.aliyuncs.com"),
		"cn-qingdao":                  tea.String("ltl.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ltl.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ltl.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ltl.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ltl.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ltl.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ltl.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("ltl.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ltl.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ltl.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ltl.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ltl.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ltl.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("ltl.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ltl.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ltl.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ltl.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("ltl.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ltl.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ltl.aliyuncs.com"),
		"eu-central-1":                tea.String("ltl.aliyuncs.com"),
		"eu-west-1":                   tea.String("ltl.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ltl.aliyuncs.com"),
		"me-east-1":                   tea.String("ltl.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ltl.aliyuncs.com"),
		"us-east-1":                   tea.String("ltl.aliyuncs.com"),
		"us-west-1":                   tea.String("ltl.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ltl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ApplyDataModelConfigInfoWithOptions(request *ApplyDataModelConfigInfoRequest, runtime *util.RuntimeOptions) (_result *ApplyDataModelConfigInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		query["Configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.DataModelCode)) {
		query["DataModelCode"] = request.DataModelCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyDataModelConfigInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyDataModelConfigInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyDataModelConfigInfo(request *ApplyDataModelConfigInfoRequest) (_result *ApplyDataModelConfigInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyDataModelConfigInfoResponse{}
	_body, _err := client.ApplyDataModelConfigInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDataWithOptions(request *AttachDataRequest, runtime *util.RuntimeOptions) (_result *AttachDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachData"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachData(request *AttachDataRequest) (_result *AttachDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachDataResponse{}
	_body, _err := client.AttachDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDataWithSignatureWithOptions(request *AttachDataWithSignatureRequest, runtime *util.RuntimeOptions) (_result *AttachDataWithSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.IotAuthType)) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !tea.BoolValue(util.IsUnset(request.IotDataDigest)) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdServiceProvider)) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdSource)) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !tea.BoolValue(util.IsUnset(request.IotSignature)) {
		query["IotSignature"] = request.IotSignature
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachDataWithSignature"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachDataWithSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachDataWithSignature(request *AttachDataWithSignatureRequest) (_result *AttachDataWithSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachDataWithSignatureResponse{}
	_body, _err := client.AttachDataWithSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeDeviceWithOptions(request *AuthorizeDeviceRequest, runtime *util.RuntimeOptions) (_result *AuthorizeDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeDevice(request *AuthorizeDeviceRequest) (_result *AuthorizeDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeDeviceResponse{}
	_body, _err := client.AuthorizeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeDeviceGroupWithOptions(request *AuthorizeDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeDeviceGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeDeviceGroup(request *AuthorizeDeviceGroupRequest) (_result *AuthorizeDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeDeviceGroupResponse{}
	_body, _err := client.AuthorizeDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseDigestInfoWithOptions(tmpReq *BatchUploadMPCoSPhaseDigestInfoRequest, runtime *util.RuntimeOptions) (_result *BatchUploadMPCoSPhaseDigestInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchUploadMPCoSPhaseDigestInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PhaseDataList)) {
		request.PhaseDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhaseDataList, tea.String("PhaseDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseDataListShrink)) {
		query["PhaseDataList"] = request.PhaseDataListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUploadMPCoSPhaseDigestInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUploadMPCoSPhaseDigestInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseDigestInfo(request *BatchUploadMPCoSPhaseDigestInfoRequest) (_result *BatchUploadMPCoSPhaseDigestInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUploadMPCoSPhaseDigestInfoResponse{}
	_body, _err := client.BatchUploadMPCoSPhaseDigestInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseDigestInfoByDeviceWithOptions(tmpReq *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest, runtime *util.RuntimeOptions) (_result *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchUploadMPCoSPhaseDigestInfoByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PhaseDataList)) {
		request.PhaseDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhaseDataList, tea.String("PhaseDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.IotAuthType)) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !tea.BoolValue(util.IsUnset(request.IotDataDigest)) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdServiceProvider)) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdSource)) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !tea.BoolValue(util.IsUnset(request.IotSignature)) {
		query["IotSignature"] = request.IotSignature
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseDataListShrink)) {
		query["PhaseDataList"] = request.PhaseDataListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUploadMPCoSPhaseDigestInfoByDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUploadMPCoSPhaseDigestInfoByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseDigestInfoByDevice(request *BatchUploadMPCoSPhaseDigestInfoByDeviceRequest) (_result *BatchUploadMPCoSPhaseDigestInfoByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUploadMPCoSPhaseDigestInfoByDeviceResponse{}
	_body, _err := client.BatchUploadMPCoSPhaseDigestInfoByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseTextInfoWithOptions(tmpReq *BatchUploadMPCoSPhaseTextInfoRequest, runtime *util.RuntimeOptions) (_result *BatchUploadMPCoSPhaseTextInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchUploadMPCoSPhaseTextInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PhaseDataList)) {
		request.PhaseDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhaseDataList, tea.String("PhaseDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseDataListShrink)) {
		query["PhaseDataList"] = request.PhaseDataListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUploadMPCoSPhaseTextInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUploadMPCoSPhaseTextInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseTextInfo(request *BatchUploadMPCoSPhaseTextInfoRequest) (_result *BatchUploadMPCoSPhaseTextInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUploadMPCoSPhaseTextInfoResponse{}
	_body, _err := client.BatchUploadMPCoSPhaseTextInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseTextInfoByDeviceWithOptions(tmpReq *BatchUploadMPCoSPhaseTextInfoByDeviceRequest, runtime *util.RuntimeOptions) (_result *BatchUploadMPCoSPhaseTextInfoByDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchUploadMPCoSPhaseTextInfoByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PhaseDataList)) {
		request.PhaseDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhaseDataList, tea.String("PhaseDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.IotAuthType)) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !tea.BoolValue(util.IsUnset(request.IotDataDigest)) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdServiceProvider)) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdSource)) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !tea.BoolValue(util.IsUnset(request.IotSignature)) {
		query["IotSignature"] = request.IotSignature
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseDataListShrink)) {
		query["PhaseDataList"] = request.PhaseDataListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUploadMPCoSPhaseTextInfoByDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUploadMPCoSPhaseTextInfoByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUploadMPCoSPhaseTextInfoByDevice(request *BatchUploadMPCoSPhaseTextInfoByDeviceRequest) (_result *BatchUploadMPCoSPhaseTextInfoByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUploadMPCoSPhaseTextInfoByDeviceResponse{}
	_body, _err := client.BatchUploadMPCoSPhaseTextInfoByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMPCoSPhaseWithOptions(request *CreateMPCoSPhaseRequest, runtime *util.RuntimeOptions) (_result *CreateMPCoSPhaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMPCoSPhase"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMPCoSPhaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMPCoSPhase(request *CreateMPCoSPhaseRequest) (_result *CreateMPCoSPhaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMPCoSPhaseResponse{}
	_body, _err := client.CreateMPCoSPhaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMPCoSPhaseGroupWithOptions(request *CreateMPCoSPhaseGroupRequest, runtime *util.RuntimeOptions) (_result *CreateMPCoSPhaseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMPCoSPhaseGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMPCoSPhaseGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMPCoSPhaseGroup(request *CreateMPCoSPhaseGroupRequest) (_result *CreateMPCoSPhaseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMPCoSPhaseGroupResponse{}
	_body, _err := client.CreateMPCoSPhaseGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMemberWithOptions(request *CreateMemberRequest, runtime *util.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberContact)) {
		query["MemberContact"] = request.MemberContact
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		query["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.MemberPhone)) {
		query["MemberPhone"] = request.MemberPhone
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMember"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMember(request *CreateMemberRequest) (_result *CreateMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMemberResponse{}
	_body, _err := client.CreateMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCapacityInfoWithOptions(request *DescribeCapacityInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeCapacityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCapacityInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCapacityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCapacityInfo(request *DescribeCapacityInfoRequest) (_result *DescribeCapacityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCapacityInfoResponse{}
	_body, _err := client.DescribeCapacityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMPCoSAuthorizedInfoWithOptions(request *DescribeMPCoSAuthorizedInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeMPCoSAuthorizedInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMPCoSAuthorizedInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMPCoSAuthorizedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPCoSAuthorizedInfo(request *DescribeMPCoSAuthorizedInfoRequest) (_result *DescribeMPCoSAuthorizedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPCoSAuthorizedInfoResponse{}
	_body, _err := client.DescribeMPCoSAuthorizedInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMPCoSPhaseInfoWithOptions(request *DescribeMPCoSPhaseInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeMPCoSPhaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DataKey)) {
		query["DataKey"] = request.DataKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataSeq)) {
		query["DataSeq"] = request.DataSeq
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMPCoSPhaseInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMPCoSPhaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPCoSPhaseInfo(request *DescribeMPCoSPhaseInfoRequest) (_result *DescribeMPCoSPhaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPCoSPhaseInfoResponse{}
	_body, _err := client.DescribeMPCoSPhaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMPCoSResourceInfoWithOptions(request *DescribeMPCoSResourceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeMPCoSResourceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMPCoSResourceInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMPCoSResourceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPCoSResourceInfo(request *DescribeMPCoSResourceInfoRequest) (_result *DescribeMPCoSResourceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPCoSResourceInfoResponse{}
	_body, _err := client.DescribeMPCoSResourceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMemberCapacityInfoWithOptions(request *DescribeMemberCapacityInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeMemberCapacityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMemberCapacityInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMemberCapacityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMemberCapacityInfo(request *DescribeMemberCapacityInfoRequest) (_result *DescribeMemberCapacityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMemberCapacityInfoResponse{}
	_body, _err := client.DescribeMemberCapacityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceInfoWithOptions(request *DescribeResourceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceInfo(request *DescribeResourceInfoRequest) (_result *DescribeResourceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceInfoResponse{}
	_body, _err := client.DescribeResourceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBlockChainInfoWithOptions(request *GetBlockChainInfoRequest, runtime *util.RuntimeOptions) (_result *GetBlockChainInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBlockChainInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBlockChainInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBlockChainInfo(request *GetBlockChainInfoRequest) (_result *GetBlockChainInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBlockChainInfoResponse{}
	_body, _err := client.GetBlockChainInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataWithOptions(request *GetDataRequest, runtime *util.RuntimeOptions) (_result *GetDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetData"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetData(request *GetDataRequest) (_result *GetDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataResponse{}
	_body, _err := client.GetDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataModelConfigInfoWithOptions(request *GetDataModelConfigInfoRequest, runtime *util.RuntimeOptions) (_result *GetDataModelConfigInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataModelCode)) {
		query["DataModelCode"] = request.DataModelCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataModelConfigInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataModelConfigInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataModelConfigInfo(request *GetDataModelConfigInfoRequest) (_result *GetDataModelConfigInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataModelConfigInfoResponse{}
	_body, _err := client.GetDataModelConfigInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoryDataCountWithOptions(request *GetHistoryDataCountRequest, runtime *util.RuntimeOptions) (_result *GetHistoryDataCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoryDataCount"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoryDataCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoryDataCount(request *GetHistoryDataCountRequest) (_result *GetHistoryDataCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoryDataCountResponse{}
	_body, _err := client.GetHistoryDataCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoryDataListWithOptions(request *GetHistoryDataListRequest, runtime *util.RuntimeOptions) (_result *GetHistoryDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoryDataList"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoryDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoryDataList(request *GetHistoryDataListRequest) (_result *GetHistoryDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoryDataListResponse{}
	_body, _err := client.GetHistoryDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDependentDataModelsWithOptions(request *ListDependentDataModelsRequest, runtime *util.RuntimeOptions) (_result *ListDependentDataModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDependentDataModels"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDependentDataModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDependentDataModels(request *ListDependentDataModelsRequest) (_result *ListDependentDataModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDependentDataModelsResponse{}
	_body, _err := client.ListDependentDataModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceWithOptions(request *ListDeviceRequest, runtime *util.RuntimeOptions) (_result *ListDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevice(request *ListDeviceRequest) (_result *ListDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceResponse{}
	_body, _err := client.ListDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceGroupWithOptions(request *ListDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *ListDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceGroup(request *ListDeviceGroupRequest) (_result *ListDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceGroupResponse{}
	_body, _err := client.ListDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMPCoSPhaseWithOptions(request *ListMPCoSPhaseRequest, runtime *util.RuntimeOptions) (_result *ListMPCoSPhaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMPCoSPhase"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMPCoSPhaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMPCoSPhase(request *ListMPCoSPhaseRequest) (_result *ListMPCoSPhaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMPCoSPhaseResponse{}
	_body, _err := client.ListMPCoSPhaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMPCoSPhaseGroupWithOptions(request *ListMPCoSPhaseGroupRequest, runtime *util.RuntimeOptions) (_result *ListMPCoSPhaseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMPCoSPhaseGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMPCoSPhaseGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMPCoSPhaseGroup(request *ListMPCoSPhaseGroupRequest) (_result *ListMPCoSPhaseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMPCoSPhaseGroupResponse{}
	_body, _err := client.ListMPCoSPhaseGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMPCoSPhaseHistoryWithOptions(request *ListMPCoSPhaseHistoryRequest, runtime *util.RuntimeOptions) (_result *ListMPCoSPhaseHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DataKey)) {
		query["DataKey"] = request.DataKey
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMPCoSPhaseHistory"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMPCoSPhaseHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMPCoSPhaseHistory(request *ListMPCoSPhaseHistoryRequest) (_result *ListMPCoSPhaseHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMPCoSPhaseHistoryResponse{}
	_body, _err := client.ListMPCoSPhaseHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMemberWithOptions(request *ListMemberRequest, runtime *util.RuntimeOptions) (_result *ListMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMember"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMember(request *ListMemberRequest) (_result *ListMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMemberResponse{}
	_body, _err := client.ListMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMultiPartyCollaborationChainWithOptions(request *ListMultiPartyCollaborationChainRequest, runtime *util.RuntimeOptions) (_result *ListMultiPartyCollaborationChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiPartyCollaborationChain"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiPartyCollaborationChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMultiPartyCollaborationChain(request *ListMultiPartyCollaborationChainRequest) (_result *ListMultiPartyCollaborationChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiPartyCollaborationChainResponse{}
	_body, _err := client.ListMultiPartyCollaborationChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPSMemberDataTypeCodeWithOptions(request *ListPSMemberDataTypeCodeRequest, runtime *util.RuntimeOptions) (_result *ListPSMemberDataTypeCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPSMemberDataTypeCode"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPSMemberDataTypeCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPSMemberDataTypeCode(request *ListPSMemberDataTypeCodeRequest) (_result *ListPSMemberDataTypeCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPSMemberDataTypeCodeResponse{}
	_body, _err := client.ListPSMemberDataTypeCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProofChainWithOptions(request *ListProofChainRequest, runtime *util.RuntimeOptions) (_result *ListProofChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProofChain"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProofChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProofChain(request *ListProofChainRequest) (_result *ListProofChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProofChainResponse{}
	_body, _err := client.ListProofChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockMemberWithOptions(request *LockMemberRequest, runtime *util.RuntimeOptions) (_result *LockMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockMember"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockMember(request *LockMemberRequest) (_result *LockMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockMemberResponse{}
	_body, _err := client.LockMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMPCoSPhaseWithOptions(request *ModifyMPCoSPhaseRequest, runtime *util.RuntimeOptions) (_result *ModifyMPCoSPhaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMPCoSPhase"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMPCoSPhaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMPCoSPhase(request *ModifyMPCoSPhaseRequest) (_result *ModifyMPCoSPhaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMPCoSPhaseResponse{}
	_body, _err := client.ModifyMPCoSPhaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMPCoSPhaseGroupWithOptions(request *ModifyMPCoSPhaseGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyMPCoSPhaseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMPCoSPhaseGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMPCoSPhaseGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMPCoSPhaseGroup(request *ModifyMPCoSPhaseGroupRequest) (_result *ModifyMPCoSPhaseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMPCoSPhaseGroupResponse{}
	_body, _err := client.ModifyMPCoSPhaseGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMemberWithOptions(request *ModifyMemberRequest, runtime *util.RuntimeOptions) (_result *ModifyMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberContact)) {
		query["MemberContact"] = request.MemberContact
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		query["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.MemberPhone)) {
		query["MemberPhone"] = request.MemberPhone
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMember"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMember(request *ModifyMemberRequest) (_result *ModifyMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMemberResponse{}
	_body, _err := client.ModifyMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDeviceGroupWithOptions(request *RegisterDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeType)) {
		query["AuthorizeType"] = request.AuthorizeType
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupName)) {
		query["DeviceGroupName"] = request.DeviceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDeviceGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDeviceGroup(request *RegisterDeviceGroupRequest) (_result *RegisterDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceGroupResponse{}
	_body, _err := client.RegisterDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDataWithOptions(request *SetDataRequest, runtime *util.RuntimeOptions) (_result *SetDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetData"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetData(request *SetDataRequest) (_result *SetDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataResponse{}
	_body, _err := client.SetDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDataWithSignatureWithOptions(request *SetDataWithSignatureRequest, runtime *util.RuntimeOptions) (_result *SetDataWithSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IotAuthType)) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !tea.BoolValue(util.IsUnset(request.IotDataDigest)) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdServiceProvider)) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdSource)) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !tea.BoolValue(util.IsUnset(request.IotSignature)) {
		query["IotSignature"] = request.IotSignature
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDataWithSignature"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataWithSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDataWithSignature(request *SetDataWithSignatureRequest) (_result *SetDataWithSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataWithSignatureResponse{}
	_body, _err := client.SetDataWithSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnAuthorizeDeviceWithOptions(request *UnAuthorizeDeviceRequest, runtime *util.RuntimeOptions) (_result *UnAuthorizeDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnAuthorizeDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnAuthorizeDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnAuthorizeDevice(request *UnAuthorizeDeviceRequest) (_result *UnAuthorizeDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnAuthorizeDeviceResponse{}
	_body, _err := client.UnAuthorizeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnAuthorizeDeviceGroupWithOptions(request *UnAuthorizeDeviceGroupRequest, runtime *util.RuntimeOptions) (_result *UnAuthorizeDeviceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceGroupId)) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnAuthorizeDeviceGroup"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnAuthorizeDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnAuthorizeDeviceGroup(request *UnAuthorizeDeviceGroupRequest) (_result *UnAuthorizeDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnAuthorizeDeviceGroupResponse{}
	_body, _err := client.UnAuthorizeDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnLockMemberWithOptions(request *UnLockMemberRequest, runtime *util.RuntimeOptions) (_result *UnLockMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnLockMember"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnLockMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnLockMember(request *UnLockMemberRequest) (_result *UnLockMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnLockMemberResponse{}
	_body, _err := client.UnLockMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMPCoSAuthorizedInfoWithOptions(tmpReq *UpdateMPCoSAuthorizedInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateMPCoSAuthorizedInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMPCoSAuthorizedInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuthorizedPhaseList)) {
		request.AuthorizedPhaseListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthorizedPhaseList, tea.String("AuthorizedPhaseList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizedPhaseListShrink)) {
		query["AuthorizedPhaseList"] = request.AuthorizedPhaseListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMPCoSAuthorizedInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMPCoSAuthorizedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMPCoSAuthorizedInfo(request *UpdateMPCoSAuthorizedInfoRequest) (_result *UpdateMPCoSAuthorizedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMPCoSAuthorizedInfoResponse{}
	_body, _err := client.UpdateMPCoSAuthorizedInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseDigestInfoWithOptions(tmpReq *UploadMPCoSPhaseDigestInfoRequest, runtime *util.RuntimeOptions) (_result *UploadMPCoSPhaseDigestInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UploadMPCoSPhaseDigestInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedDataList)) {
		request.RelatedDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedDataList, tea.String("RelatedDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DataKey)) {
		query["DataKey"] = request.DataKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataSeq)) {
		query["DataSeq"] = request.DataSeq
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseData)) {
		query["PhaseData"] = request.PhaseData
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedDataListShrink)) {
		query["RelatedDataList"] = request.RelatedDataListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadMPCoSPhaseDigestInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadMPCoSPhaseDigestInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseDigestInfo(request *UploadMPCoSPhaseDigestInfoRequest) (_result *UploadMPCoSPhaseDigestInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadMPCoSPhaseDigestInfoResponse{}
	_body, _err := client.UploadMPCoSPhaseDigestInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseDigestInfoByDeviceWithOptions(tmpReq *UploadMPCoSPhaseDigestInfoByDeviceRequest, runtime *util.RuntimeOptions) (_result *UploadMPCoSPhaseDigestInfoByDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UploadMPCoSPhaseDigestInfoByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedDataList)) {
		request.RelatedDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedDataList, tea.String("RelatedDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DataKey)) {
		query["DataKey"] = request.DataKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataSeq)) {
		query["DataSeq"] = request.DataSeq
	}

	if !tea.BoolValue(util.IsUnset(request.IotAuthType)) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !tea.BoolValue(util.IsUnset(request.IotDataDigest)) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdServiceProvider)) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdSource)) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !tea.BoolValue(util.IsUnset(request.IotSignature)) {
		query["IotSignature"] = request.IotSignature
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseData)) {
		query["PhaseData"] = request.PhaseData
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedDataListShrink)) {
		query["RelatedDataList"] = request.RelatedDataListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadMPCoSPhaseDigestInfoByDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadMPCoSPhaseDigestInfoByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseDigestInfoByDevice(request *UploadMPCoSPhaseDigestInfoByDeviceRequest) (_result *UploadMPCoSPhaseDigestInfoByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadMPCoSPhaseDigestInfoByDeviceResponse{}
	_body, _err := client.UploadMPCoSPhaseDigestInfoByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseTextInfoWithOptions(tmpReq *UploadMPCoSPhaseTextInfoRequest, runtime *util.RuntimeOptions) (_result *UploadMPCoSPhaseTextInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UploadMPCoSPhaseTextInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedDataList)) {
		request.RelatedDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedDataList, tea.String("RelatedDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DataKey)) {
		query["DataKey"] = request.DataKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataSeq)) {
		query["DataSeq"] = request.DataSeq
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseData)) {
		query["PhaseData"] = request.PhaseData
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedDataListShrink)) {
		query["RelatedDataList"] = request.RelatedDataListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadMPCoSPhaseTextInfo"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadMPCoSPhaseTextInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseTextInfo(request *UploadMPCoSPhaseTextInfoRequest) (_result *UploadMPCoSPhaseTextInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadMPCoSPhaseTextInfoResponse{}
	_body, _err := client.UploadMPCoSPhaseTextInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseTextInfoByDeviceWithOptions(tmpReq *UploadMPCoSPhaseTextInfoByDeviceRequest, runtime *util.RuntimeOptions) (_result *UploadMPCoSPhaseTextInfoByDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UploadMPCoSPhaseTextInfoByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedDataList)) {
		request.RelatedDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedDataList, tea.String("RelatedDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiVersion)) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizChainId)) {
		query["BizChainId"] = request.BizChainId
	}

	if !tea.BoolValue(util.IsUnset(request.DataKey)) {
		query["DataKey"] = request.DataKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataSeq)) {
		query["DataSeq"] = request.DataSeq
	}

	if !tea.BoolValue(util.IsUnset(request.IotAuthType)) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !tea.BoolValue(util.IsUnset(request.IotDataDigest)) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdServiceProvider)) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.IotIdSource)) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !tea.BoolValue(util.IsUnset(request.IotSignature)) {
		query["IotSignature"] = request.IotSignature
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseData)) {
		query["PhaseData"] = request.PhaseData
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseGroupId)) {
		query["PhaseGroupId"] = request.PhaseGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PhaseId)) {
		query["PhaseId"] = request.PhaseId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedDataListShrink)) {
		query["RelatedDataList"] = request.RelatedDataListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadMPCoSPhaseTextInfoByDevice"),
		Version:     tea.String("2019-05-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadMPCoSPhaseTextInfoByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadMPCoSPhaseTextInfoByDevice(request *UploadMPCoSPhaseTextInfoByDeviceRequest) (_result *UploadMPCoSPhaseTextInfoByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadMPCoSPhaseTextInfoByDeviceResponse{}
	_body, _err := client.UploadMPCoSPhaseTextInfoByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

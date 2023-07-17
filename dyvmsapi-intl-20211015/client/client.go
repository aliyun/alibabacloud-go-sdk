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

type BackendCallGroupRequest struct {
	CalledNumber         []*string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty" type:"Repeated"`
	CallerIdNumber       *string   `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	CountryId            *string   `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	OutId                *string   `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int64    `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SendType             *int64    `json:"SendType,omitempty" xml:"SendType,omitempty"`
	Speed                *int64    `json:"Speed,omitempty" xml:"Speed,omitempty"`
	TaskName             *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TimingStart          *string   `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	TtsCode              *string   `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	VoiceCode            *string   `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	Volume               *int64    `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s BackendCallGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupRequest) GoString() string {
	return s.String()
}

func (s *BackendCallGroupRequest) SetCalledNumber(v []*string) *BackendCallGroupRequest {
	s.CalledNumber = v
	return s
}

func (s *BackendCallGroupRequest) SetCallerIdNumber(v string) *BackendCallGroupRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *BackendCallGroupRequest) SetCountryId(v string) *BackendCallGroupRequest {
	s.CountryId = &v
	return s
}

func (s *BackendCallGroupRequest) SetOutId(v string) *BackendCallGroupRequest {
	s.OutId = &v
	return s
}

func (s *BackendCallGroupRequest) SetOwnerId(v int64) *BackendCallGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *BackendCallGroupRequest) SetPlayTimes(v int64) *BackendCallGroupRequest {
	s.PlayTimes = &v
	return s
}

func (s *BackendCallGroupRequest) SetResourceOwnerAccount(v string) *BackendCallGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BackendCallGroupRequest) SetResourceOwnerId(v int64) *BackendCallGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BackendCallGroupRequest) SetSendType(v int64) *BackendCallGroupRequest {
	s.SendType = &v
	return s
}

func (s *BackendCallGroupRequest) SetSpeed(v int64) *BackendCallGroupRequest {
	s.Speed = &v
	return s
}

func (s *BackendCallGroupRequest) SetTaskName(v string) *BackendCallGroupRequest {
	s.TaskName = &v
	return s
}

func (s *BackendCallGroupRequest) SetTimingStart(v string) *BackendCallGroupRequest {
	s.TimingStart = &v
	return s
}

func (s *BackendCallGroupRequest) SetTtsCode(v string) *BackendCallGroupRequest {
	s.TtsCode = &v
	return s
}

func (s *BackendCallGroupRequest) SetVoiceCode(v string) *BackendCallGroupRequest {
	s.VoiceCode = &v
	return s
}

func (s *BackendCallGroupRequest) SetVolume(v int64) *BackendCallGroupRequest {
	s.Volume = &v
	return s
}

type BackendCallGroupShrinkRequest struct {
	CalledNumberShrink   *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallerIdNumber       *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SendType             *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	Speed                *int64  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TimingStart          *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	TtsCode              *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	Volume               *int64  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s BackendCallGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *BackendCallGroupShrinkRequest) SetCalledNumberShrink(v string) *BackendCallGroupShrinkRequest {
	s.CalledNumberShrink = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetCallerIdNumber(v string) *BackendCallGroupShrinkRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetCountryId(v string) *BackendCallGroupShrinkRequest {
	s.CountryId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetOutId(v string) *BackendCallGroupShrinkRequest {
	s.OutId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetOwnerId(v int64) *BackendCallGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetPlayTimes(v int64) *BackendCallGroupShrinkRequest {
	s.PlayTimes = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetResourceOwnerAccount(v string) *BackendCallGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetResourceOwnerId(v int64) *BackendCallGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetSendType(v int64) *BackendCallGroupShrinkRequest {
	s.SendType = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetSpeed(v int64) *BackendCallGroupShrinkRequest {
	s.Speed = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetTaskName(v string) *BackendCallGroupShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetTimingStart(v string) *BackendCallGroupShrinkRequest {
	s.TimingStart = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetTtsCode(v string) *BackendCallGroupShrinkRequest {
	s.TtsCode = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetVoiceCode(v string) *BackendCallGroupShrinkRequest {
	s.VoiceCode = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetVolume(v int64) *BackendCallGroupShrinkRequest {
	s.Volume = &v
	return s
}

type BackendCallGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackendCallGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BackendCallGroupResponseBody) SetCode(v string) *BackendCallGroupResponseBody {
	s.Code = &v
	return s
}

func (s *BackendCallGroupResponseBody) SetMessage(v string) *BackendCallGroupResponseBody {
	s.Message = &v
	return s
}

func (s *BackendCallGroupResponseBody) SetRequestId(v string) *BackendCallGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackendCallGroupResponseBody) SetTaskId(v string) *BackendCallGroupResponseBody {
	s.TaskId = &v
	return s
}

type BackendCallGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BackendCallGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BackendCallGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupResponse) GoString() string {
	return s.String()
}

func (s *BackendCallGroupResponse) SetHeaders(v map[string]*string) *BackendCallGroupResponse {
	s.Headers = v
	return s
}

func (s *BackendCallGroupResponse) SetStatusCode(v int32) *BackendCallGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *BackendCallGroupResponse) SetBody(v *BackendCallGroupResponseBody) *BackendCallGroupResponse {
	s.Body = v
	return s
}

type BackendCallSignalRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallerIdNumber       *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Speed                *int64  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	TtsCode              *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	Volume               *int64  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s BackendCallSignalRequest) String() string {
	return tea.Prettify(s)
}

func (s BackendCallSignalRequest) GoString() string {
	return s.String()
}

func (s *BackendCallSignalRequest) SetCalledNumber(v string) *BackendCallSignalRequest {
	s.CalledNumber = &v
	return s
}

func (s *BackendCallSignalRequest) SetCallerIdNumber(v string) *BackendCallSignalRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *BackendCallSignalRequest) SetCountryId(v string) *BackendCallSignalRequest {
	s.CountryId = &v
	return s
}

func (s *BackendCallSignalRequest) SetOutId(v string) *BackendCallSignalRequest {
	s.OutId = &v
	return s
}

func (s *BackendCallSignalRequest) SetOwnerId(v int64) *BackendCallSignalRequest {
	s.OwnerId = &v
	return s
}

func (s *BackendCallSignalRequest) SetPlayTimes(v int64) *BackendCallSignalRequest {
	s.PlayTimes = &v
	return s
}

func (s *BackendCallSignalRequest) SetResourceOwnerAccount(v string) *BackendCallSignalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BackendCallSignalRequest) SetResourceOwnerId(v int64) *BackendCallSignalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BackendCallSignalRequest) SetSpeed(v int64) *BackendCallSignalRequest {
	s.Speed = &v
	return s
}

func (s *BackendCallSignalRequest) SetTtsCode(v string) *BackendCallSignalRequest {
	s.TtsCode = &v
	return s
}

func (s *BackendCallSignalRequest) SetTtsParam(v string) *BackendCallSignalRequest {
	s.TtsParam = &v
	return s
}

func (s *BackendCallSignalRequest) SetVolume(v int64) *BackendCallSignalRequest {
	s.Volume = &v
	return s
}

type BackendCallSignalResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BackendCallSignalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BackendCallSignalResponseBody) GoString() string {
	return s.String()
}

func (s *BackendCallSignalResponseBody) SetCallId(v string) *BackendCallSignalResponseBody {
	s.CallId = &v
	return s
}

func (s *BackendCallSignalResponseBody) SetCode(v string) *BackendCallSignalResponseBody {
	s.Code = &v
	return s
}

func (s *BackendCallSignalResponseBody) SetMessage(v string) *BackendCallSignalResponseBody {
	s.Message = &v
	return s
}

func (s *BackendCallSignalResponseBody) SetRequestId(v string) *BackendCallSignalResponseBody {
	s.RequestId = &v
	return s
}

type BackendCallSignalResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BackendCallSignalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BackendCallSignalResponse) String() string {
	return tea.Prettify(s)
}

func (s BackendCallSignalResponse) GoString() string {
	return s.String()
}

func (s *BackendCallSignalResponse) SetHeaders(v map[string]*string) *BackendCallSignalResponse {
	s.Headers = v
	return s
}

func (s *BackendCallSignalResponse) SetStatusCode(v int32) *BackendCallSignalResponse {
	s.StatusCode = &v
	return s
}

func (s *BackendCallSignalResponse) SetBody(v *BackendCallSignalResponseBody) *BackendCallSignalResponse {
	s.Body = v
	return s
}

type CancleGroupCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancleGroupCallRequest) String() string {
	return tea.Prettify(s)
}

func (s CancleGroupCallRequest) GoString() string {
	return s.String()
}

func (s *CancleGroupCallRequest) SetOwnerId(v int64) *CancleGroupCallRequest {
	s.OwnerId = &v
	return s
}

func (s *CancleGroupCallRequest) SetResourceOwnerAccount(v string) *CancleGroupCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancleGroupCallRequest) SetResourceOwnerId(v int64) *CancleGroupCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancleGroupCallRequest) SetTaskId(v string) *CancleGroupCallRequest {
	s.TaskId = &v
	return s
}

type CancleGroupCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancleGroupCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancleGroupCallResponseBody) GoString() string {
	return s.String()
}

func (s *CancleGroupCallResponseBody) SetCode(v string) *CancleGroupCallResponseBody {
	s.Code = &v
	return s
}

func (s *CancleGroupCallResponseBody) SetRequestId(v string) *CancleGroupCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancleGroupCallResponseBody) SetTaskId(v string) *CancleGroupCallResponseBody {
	s.TaskId = &v
	return s
}

type CancleGroupCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancleGroupCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancleGroupCallResponse) String() string {
	return tea.Prettify(s)
}

func (s CancleGroupCallResponse) GoString() string {
	return s.String()
}

func (s *CancleGroupCallResponse) SetHeaders(v map[string]*string) *CancleGroupCallResponse {
	s.Headers = v
	return s
}

func (s *CancleGroupCallResponse) SetStatusCode(v int32) *CancleGroupCallResponse {
	s.StatusCode = &v
	return s
}

func (s *CancleGroupCallResponse) SetBody(v *CancleGroupCallResponseBody) *CancleGroupCallResponse {
	s.Body = v
	return s
}

type DeleteApplyNumberRecordRequest struct {
	ApplyId              *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteApplyNumberRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplyNumberRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplyNumberRecordRequest) SetApplyId(v string) *DeleteApplyNumberRecordRequest {
	s.ApplyId = &v
	return s
}

func (s *DeleteApplyNumberRecordRequest) SetOwnerId(v int64) *DeleteApplyNumberRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteApplyNumberRecordRequest) SetResourceOwnerAccount(v string) *DeleteApplyNumberRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteApplyNumberRecordRequest) SetResourceOwnerId(v int64) *DeleteApplyNumberRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteApplyNumberRecordResponseBody struct {
	ApplyId   *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplyNumberRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplyNumberRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplyNumberRecordResponseBody) SetApplyId(v string) *DeleteApplyNumberRecordResponseBody {
	s.ApplyId = &v
	return s
}

func (s *DeleteApplyNumberRecordResponseBody) SetCode(v string) *DeleteApplyNumberRecordResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplyNumberRecordResponseBody) SetMessage(v string) *DeleteApplyNumberRecordResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplyNumberRecordResponseBody) SetRequestId(v string) *DeleteApplyNumberRecordResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplyNumberRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApplyNumberRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplyNumberRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplyNumberRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplyNumberRecordResponse) SetHeaders(v map[string]*string) *DeleteApplyNumberRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplyNumberRecordResponse) SetStatusCode(v int32) *DeleteApplyNumberRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplyNumberRecordResponse) SetBody(v *DeleteApplyNumberRecordResponseBody) *DeleteApplyNumberRecordResponse {
	s.Body = v
	return s
}

type DeleteQualificationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteQualificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualificationRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualificationRequest) SetOwnerId(v int64) *DeleteQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteQualificationRequest) SetQualificationId(v string) *DeleteQualificationRequest {
	s.QualificationId = &v
	return s
}

func (s *DeleteQualificationRequest) SetResourceOwnerAccount(v string) *DeleteQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteQualificationRequest) SetResourceOwnerId(v int64) *DeleteQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteQualificationResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQualificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualificationResponseBody) SetCode(v string) *DeleteQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualificationResponseBody) SetMessage(v string) *DeleteQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualificationResponseBody) SetQualificationId(v string) *DeleteQualificationResponseBody {
	s.QualificationId = &v
	return s
}

func (s *DeleteQualificationResponseBody) SetRequestId(v string) *DeleteQualificationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQualificationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQualificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQualificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualificationResponse) SetHeaders(v map[string]*string) *DeleteQualificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualificationResponse) SetStatusCode(v int32) *DeleteQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualificationResponse) SetBody(v *DeleteQualificationResponseBody) *DeleteQualificationResponse {
	s.Body = v
	return s
}

type DeleteSipTrunkApplyRequest struct {
	ApplyId              *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteSipTrunkApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSipTrunkApplyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSipTrunkApplyRequest) SetApplyId(v string) *DeleteSipTrunkApplyRequest {
	s.ApplyId = &v
	return s
}

func (s *DeleteSipTrunkApplyRequest) SetOwnerId(v int64) *DeleteSipTrunkApplyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSipTrunkApplyRequest) SetResourceOwnerAccount(v string) *DeleteSipTrunkApplyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSipTrunkApplyRequest) SetResourceOwnerId(v int64) *DeleteSipTrunkApplyRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteSipTrunkApplyResponseBody struct {
	ApplyId   *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSipTrunkApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSipTrunkApplyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSipTrunkApplyResponseBody) SetApplyId(v string) *DeleteSipTrunkApplyResponseBody {
	s.ApplyId = &v
	return s
}

func (s *DeleteSipTrunkApplyResponseBody) SetCode(v string) *DeleteSipTrunkApplyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSipTrunkApplyResponseBody) SetMessage(v string) *DeleteSipTrunkApplyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSipTrunkApplyResponseBody) SetRequestId(v string) *DeleteSipTrunkApplyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSipTrunkApplyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSipTrunkApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSipTrunkApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSipTrunkApplyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSipTrunkApplyResponse) SetHeaders(v map[string]*string) *DeleteSipTrunkApplyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSipTrunkApplyResponse) SetStatusCode(v int32) *DeleteSipTrunkApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSipTrunkApplyResponse) SetBody(v *DeleteSipTrunkApplyResponseBody) *DeleteSipTrunkApplyResponse {
	s.Body = v
	return s
}

type DeleteVoiceFileRequest struct {
	FileId               *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteVoiceFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVoiceFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteVoiceFileRequest) SetFileId(v string) *DeleteVoiceFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteVoiceFileRequest) SetOwnerId(v int64) *DeleteVoiceFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVoiceFileRequest) SetResourceOwnerAccount(v string) *DeleteVoiceFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVoiceFileRequest) SetResourceOwnerId(v int64) *DeleteVoiceFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteVoiceFileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FileId    *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVoiceFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVoiceFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVoiceFileResponseBody) SetCode(v string) *DeleteVoiceFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVoiceFileResponseBody) SetFileId(v string) *DeleteVoiceFileResponseBody {
	s.FileId = &v
	return s
}

func (s *DeleteVoiceFileResponseBody) SetMessage(v string) *DeleteVoiceFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVoiceFileResponseBody) SetRequestId(v string) *DeleteVoiceFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVoiceFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVoiceFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVoiceFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVoiceFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteVoiceFileResponse) SetHeaders(v map[string]*string) *DeleteVoiceFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteVoiceFileResponse) SetStatusCode(v int32) *DeleteVoiceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVoiceFileResponse) SetBody(v *DeleteVoiceFileResponseBody) *DeleteVoiceFileResponse {
	s.Body = v
	return s
}

type DeleteVoiceTtsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteVoiceTtsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVoiceTtsRequest) GoString() string {
	return s.String()
}

func (s *DeleteVoiceTtsRequest) SetOwnerId(v int64) *DeleteVoiceTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVoiceTtsRequest) SetResourceOwnerAccount(v string) *DeleteVoiceTtsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVoiceTtsRequest) SetResourceOwnerId(v int64) *DeleteVoiceTtsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVoiceTtsRequest) SetTemplateId(v string) *DeleteVoiceTtsRequest {
	s.TemplateId = &v
	return s
}

type DeleteVoiceTtsResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteVoiceTtsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVoiceTtsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVoiceTtsResponseBody) SetCode(v string) *DeleteVoiceTtsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVoiceTtsResponseBody) SetMessage(v string) *DeleteVoiceTtsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVoiceTtsResponseBody) SetRequestId(v string) *DeleteVoiceTtsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVoiceTtsResponseBody) SetTemplateId(v string) *DeleteVoiceTtsResponseBody {
	s.TemplateId = &v
	return s
}

type DeleteVoiceTtsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVoiceTtsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVoiceTtsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVoiceTtsResponse) GoString() string {
	return s.String()
}

func (s *DeleteVoiceTtsResponse) SetHeaders(v map[string]*string) *DeleteVoiceTtsResponse {
	s.Headers = v
	return s
}

func (s *DeleteVoiceTtsResponse) SetStatusCode(v int32) *DeleteVoiceTtsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVoiceTtsResponse) SetBody(v *DeleteVoiceTtsResponseBody) *DeleteVoiceTtsResponse {
	s.Body = v
	return s
}

type DownloadTemplateFileRequest struct {
	FileType             *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DownloadTemplateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadTemplateFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadTemplateFileRequest) SetFileType(v string) *DownloadTemplateFileRequest {
	s.FileType = &v
	return s
}

func (s *DownloadTemplateFileRequest) SetOwnerId(v int64) *DownloadTemplateFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DownloadTemplateFileRequest) SetResourceOwnerAccount(v string) *DownloadTemplateFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DownloadTemplateFileRequest) SetResourceOwnerId(v int64) *DownloadTemplateFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type DownloadTemplateFileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadTemplateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadTemplateFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadTemplateFileResponseBody) SetCode(v string) *DownloadTemplateFileResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadTemplateFileResponseBody) SetMessage(v string) *DownloadTemplateFileResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadTemplateFileResponseBody) SetRequestId(v string) *DownloadTemplateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadTemplateFileResponseBody) SetUrl(v string) *DownloadTemplateFileResponseBody {
	s.Url = &v
	return s
}

type DownloadTemplateFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadTemplateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadTemplateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadTemplateFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadTemplateFileResponse) SetHeaders(v map[string]*string) *DownloadTemplateFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadTemplateFileResponse) SetStatusCode(v int32) *DownloadTemplateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadTemplateFileResponse) SetBody(v *DownloadTemplateFileResponseBody) *DownloadTemplateFileResponse {
	s.Body = v
	return s
}

type GetIntlVoiceOpenStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetIntlVoiceOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIntlVoiceOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *GetIntlVoiceOpenStatusRequest) SetOwnerId(v int64) *GetIntlVoiceOpenStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetIntlVoiceOpenStatusRequest) SetResourceOwnerAccount(v string) *GetIntlVoiceOpenStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetIntlVoiceOpenStatusRequest) SetResourceOwnerId(v int64) *GetIntlVoiceOpenStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetIntlVoiceOpenStatusResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OpenStatus *bool   `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIntlVoiceOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntlVoiceOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntlVoiceOpenStatusResponseBody) SetCode(v string) *GetIntlVoiceOpenStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetIntlVoiceOpenStatusResponseBody) SetMessage(v string) *GetIntlVoiceOpenStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetIntlVoiceOpenStatusResponseBody) SetOpenStatus(v bool) *GetIntlVoiceOpenStatusResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *GetIntlVoiceOpenStatusResponseBody) SetRequestId(v string) *GetIntlVoiceOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetIntlVoiceOpenStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIntlVoiceOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIntlVoiceOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntlVoiceOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetIntlVoiceOpenStatusResponse) SetHeaders(v map[string]*string) *GetIntlVoiceOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetIntlVoiceOpenStatusResponse) SetStatusCode(v int32) *GetIntlVoiceOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntlVoiceOpenStatusResponse) SetBody(v *GetIntlVoiceOpenStatusResponseBody) *GetIntlVoiceOpenStatusResponse {
	s.Body = v
	return s
}

type GetOssInfoForUploadFileRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetOssInfoForUploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssInfoForUploadFileRequest) GoString() string {
	return s.String()
}

func (s *GetOssInfoForUploadFileRequest) SetBizType(v string) *GetOssInfoForUploadFileRequest {
	s.BizType = &v
	return s
}

func (s *GetOssInfoForUploadFileRequest) SetOwnerId(v int64) *GetOssInfoForUploadFileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOssInfoForUploadFileRequest) SetResourceOwnerAccount(v string) *GetOssInfoForUploadFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOssInfoForUploadFileRequest) SetResourceOwnerId(v int64) *GetOssInfoForUploadFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetOssInfoForUploadFileResponseBody struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	StartWith   *string `json:"StartWith,omitempty" xml:"StartWith,omitempty"`
}

func (s GetOssInfoForUploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssInfoForUploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssInfoForUploadFileResponseBody) SetAccessKeyId(v string) *GetOssInfoForUploadFileResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssInfoForUploadFileResponseBody) SetExpireTime(v string) *GetOssInfoForUploadFileResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetOssInfoForUploadFileResponseBody) SetHost(v string) *GetOssInfoForUploadFileResponseBody {
	s.Host = &v
	return s
}

func (s *GetOssInfoForUploadFileResponseBody) SetPolicy(v string) *GetOssInfoForUploadFileResponseBody {
	s.Policy = &v
	return s
}

func (s *GetOssInfoForUploadFileResponseBody) SetRequestId(v string) *GetOssInfoForUploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssInfoForUploadFileResponseBody) SetSignature(v string) *GetOssInfoForUploadFileResponseBody {
	s.Signature = &v
	return s
}

func (s *GetOssInfoForUploadFileResponseBody) SetStartWith(v string) *GetOssInfoForUploadFileResponseBody {
	s.StartWith = &v
	return s
}

type GetOssInfoForUploadFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOssInfoForUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssInfoForUploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssInfoForUploadFileResponse) GoString() string {
	return s.String()
}

func (s *GetOssInfoForUploadFileResponse) SetHeaders(v map[string]*string) *GetOssInfoForUploadFileResponse {
	s.Headers = v
	return s
}

func (s *GetOssInfoForUploadFileResponse) SetStatusCode(v int32) *GetOssInfoForUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssInfoForUploadFileResponse) SetBody(v *GetOssInfoForUploadFileResponseBody) *GetOssInfoForUploadFileResponse {
	s.Body = v
	return s
}

type HomeStartRequest struct {
	BusinessType         *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s HomeStartRequest) String() string {
	return tea.Prettify(s)
}

func (s HomeStartRequest) GoString() string {
	return s.String()
}

func (s *HomeStartRequest) SetBusinessType(v int64) *HomeStartRequest {
	s.BusinessType = &v
	return s
}

func (s *HomeStartRequest) SetEndTime(v string) *HomeStartRequest {
	s.EndTime = &v
	return s
}

func (s *HomeStartRequest) SetOwnerId(v int64) *HomeStartRequest {
	s.OwnerId = &v
	return s
}

func (s *HomeStartRequest) SetResourceOwnerAccount(v string) *HomeStartRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *HomeStartRequest) SetResourceOwnerId(v int64) *HomeStartRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *HomeStartRequest) SetStartTime(v string) *HomeStartRequest {
	s.StartTime = &v
	return s
}

type HomeStartResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model          *HomeStartResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HomeStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HomeStartResponseBody) GoString() string {
	return s.String()
}

func (s *HomeStartResponseBody) SetCode(v string) *HomeStartResponseBody {
	s.Code = &v
	return s
}

func (s *HomeStartResponseBody) SetHttpStatusCode(v int64) *HomeStartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HomeStartResponseBody) SetMessage(v string) *HomeStartResponseBody {
	s.Message = &v
	return s
}

func (s *HomeStartResponseBody) SetModel(v *HomeStartResponseBodyModel) *HomeStartResponseBody {
	s.Model = v
	return s
}

func (s *HomeStartResponseBody) SetRequestId(v string) *HomeStartResponseBody {
	s.RequestId = &v
	return s
}

type HomeStartResponseBodyModel struct {
	CdrDrUrl *string                           `json:"CdrDrUrl,omitempty" xml:"CdrDrUrl,omitempty"`
	List     []*HomeStartResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s HomeStartResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s HomeStartResponseBodyModel) GoString() string {
	return s.String()
}

func (s *HomeStartResponseBodyModel) SetCdrDrUrl(v string) *HomeStartResponseBodyModel {
	s.CdrDrUrl = &v
	return s
}

func (s *HomeStartResponseBodyModel) SetList(v []*HomeStartResponseBodyModelList) *HomeStartResponseBodyModel {
	s.List = v
	return s
}

type HomeStartResponseBodyModelList struct {
	CallFailed  *int64  `json:"CallFailed,omitempty" xml:"CallFailed,omitempty"`
	CallSuccess *int64  `json:"CallSuccess,omitempty" xml:"CallSuccess,omitempty"`
	CallTotal   *int64  `json:"CallTotal,omitempty" xml:"CallTotal,omitempty"`
	Date        *string `json:"Date,omitempty" xml:"Date,omitempty"`
	SuccessRate *int64  `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
}

func (s HomeStartResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s HomeStartResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *HomeStartResponseBodyModelList) SetCallFailed(v int64) *HomeStartResponseBodyModelList {
	s.CallFailed = &v
	return s
}

func (s *HomeStartResponseBodyModelList) SetCallSuccess(v int64) *HomeStartResponseBodyModelList {
	s.CallSuccess = &v
	return s
}

func (s *HomeStartResponseBodyModelList) SetCallTotal(v int64) *HomeStartResponseBodyModelList {
	s.CallTotal = &v
	return s
}

func (s *HomeStartResponseBodyModelList) SetDate(v string) *HomeStartResponseBodyModelList {
	s.Date = &v
	return s
}

func (s *HomeStartResponseBodyModelList) SetSuccessRate(v int64) *HomeStartResponseBodyModelList {
	s.SuccessRate = &v
	return s
}

type HomeStartResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HomeStartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HomeStartResponse) String() string {
	return tea.Prettify(s)
}

func (s HomeStartResponse) GoString() string {
	return s.String()
}

func (s *HomeStartResponse) SetHeaders(v map[string]*string) *HomeStartResponse {
	s.Headers = v
	return s
}

func (s *HomeStartResponse) SetStatusCode(v int32) *HomeStartResponse {
	s.StatusCode = &v
	return s
}

func (s *HomeStartResponse) SetBody(v *HomeStartResponseBody) *HomeStartResponse {
	s.Body = v
	return s
}

type ListApplyNumberRecordRequest struct {
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplyNumberRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplyNumberRecordRequest) GoString() string {
	return s.String()
}

func (s *ListApplyNumberRecordRequest) SetCountryId(v string) *ListApplyNumberRecordRequest {
	s.CountryId = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetEndTime(v string) *ListApplyNumberRecordRequest {
	s.EndTime = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetOwnerId(v int64) *ListApplyNumberRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetPageNumber(v int64) *ListApplyNumberRecordRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetPageSize(v int64) *ListApplyNumberRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetResourceOwnerAccount(v string) *ListApplyNumberRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetResourceOwnerId(v int64) *ListApplyNumberRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetStartTime(v string) *ListApplyNumberRecordRequest {
	s.StartTime = &v
	return s
}

func (s *ListApplyNumberRecordRequest) SetStatus(v int64) *ListApplyNumberRecordRequest {
	s.Status = &v
	return s
}

type ListApplyNumberRecordResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListApplyNumberRecordResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListApplyNumberRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplyNumberRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplyNumberRecordResponseBody) SetCode(v string) *ListApplyNumberRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplyNumberRecordResponseBody) SetList(v []*ListApplyNumberRecordResponseBodyList) *ListApplyNumberRecordResponseBody {
	s.List = v
	return s
}

func (s *ListApplyNumberRecordResponseBody) SetMessage(v string) *ListApplyNumberRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplyNumberRecordResponseBody) SetRequestId(v string) *ListApplyNumberRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplyNumberRecordResponseBody) SetTotal(v int64) *ListApplyNumberRecordResponseBody {
	s.Total = &v
	return s
}

type ListApplyNumberRecordResponseBodyList struct {
	Amount              *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	ApplyId             *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	ApplyNote           *string `json:"ApplyNote,omitempty" xml:"ApplyNote,omitempty"`
	AuditNote           *string `json:"AuditNote,omitempty" xml:"AuditNote,omitempty"`
	AuditTs             *string `json:"AuditTs,omitempty" xml:"AuditTs,omitempty"`
	BatchIndex          *int64  `json:"BatchIndex,omitempty" xml:"BatchIndex,omitempty"`
	CommitTs            *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CountryId           *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	InboundConcurrency  *int64  `json:"InboundConcurrency,omitempty" xml:"InboundConcurrency,omitempty"`
	OutboundConcurrency *int64  `json:"OutboundConcurrency,omitempty" xml:"OutboundConcurrency,omitempty"`
	PhoneType           *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	QualificationId     *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	Scene               *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Status              *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Support             *int64  `json:"Support,omitempty" xml:"Support,omitempty"`
	UpdateTs            *string `json:"UpdateTs,omitempty" xml:"UpdateTs,omitempty"`
}

func (s ListApplyNumberRecordResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListApplyNumberRecordResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListApplyNumberRecordResponseBodyList) SetAmount(v string) *ListApplyNumberRecordResponseBodyList {
	s.Amount = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetApplyId(v string) *ListApplyNumberRecordResponseBodyList {
	s.ApplyId = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetApplyNote(v string) *ListApplyNumberRecordResponseBodyList {
	s.ApplyNote = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetAuditNote(v string) *ListApplyNumberRecordResponseBodyList {
	s.AuditNote = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetAuditTs(v string) *ListApplyNumberRecordResponseBodyList {
	s.AuditTs = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetBatchIndex(v int64) *ListApplyNumberRecordResponseBodyList {
	s.BatchIndex = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetCommitTs(v string) *ListApplyNumberRecordResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetCountryId(v string) *ListApplyNumberRecordResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetInboundConcurrency(v int64) *ListApplyNumberRecordResponseBodyList {
	s.InboundConcurrency = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetOutboundConcurrency(v int64) *ListApplyNumberRecordResponseBodyList {
	s.OutboundConcurrency = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetPhoneType(v int64) *ListApplyNumberRecordResponseBodyList {
	s.PhoneType = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetQualificationId(v string) *ListApplyNumberRecordResponseBodyList {
	s.QualificationId = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetScene(v string) *ListApplyNumberRecordResponseBodyList {
	s.Scene = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetStatus(v int64) *ListApplyNumberRecordResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetSupport(v int64) *ListApplyNumberRecordResponseBodyList {
	s.Support = &v
	return s
}

func (s *ListApplyNumberRecordResponseBodyList) SetUpdateTs(v string) *ListApplyNumberRecordResponseBodyList {
	s.UpdateTs = &v
	return s
}

type ListApplyNumberRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplyNumberRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplyNumberRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplyNumberRecordResponse) GoString() string {
	return s.String()
}

func (s *ListApplyNumberRecordResponse) SetHeaders(v map[string]*string) *ListApplyNumberRecordResponse {
	s.Headers = v
	return s
}

func (s *ListApplyNumberRecordResponse) SetStatusCode(v int32) *ListApplyNumberRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplyNumberRecordResponse) SetBody(v *ListApplyNumberRecordResponseBody) *ListApplyNumberRecordResponse {
	s.Body = v
	return s
}

type ListNumberRequest struct {
	ApplyId              *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	NumberName           *string `json:"NumberName,omitempty" xml:"NumberName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneType            *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNumberRequest) GoString() string {
	return s.String()
}

func (s *ListNumberRequest) SetApplyId(v string) *ListNumberRequest {
	s.ApplyId = &v
	return s
}

func (s *ListNumberRequest) SetCountryId(v string) *ListNumberRequest {
	s.CountryId = &v
	return s
}

func (s *ListNumberRequest) SetNumber(v string) *ListNumberRequest {
	s.Number = &v
	return s
}

func (s *ListNumberRequest) SetNumberName(v string) *ListNumberRequest {
	s.NumberName = &v
	return s
}

func (s *ListNumberRequest) SetOwnerId(v int64) *ListNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ListNumberRequest) SetPageNumber(v int64) *ListNumberRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNumberRequest) SetPageSize(v int64) *ListNumberRequest {
	s.PageSize = &v
	return s
}

func (s *ListNumberRequest) SetPhoneType(v int64) *ListNumberRequest {
	s.PhoneType = &v
	return s
}

func (s *ListNumberRequest) SetResourceOwnerAccount(v string) *ListNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListNumberRequest) SetResourceOwnerId(v int64) *ListNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListNumberResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListNumberResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListNumberResponseBody) SetCode(v string) *ListNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListNumberResponseBody) SetList(v []*ListNumberResponseBodyList) *ListNumberResponseBody {
	s.List = v
	return s
}

func (s *ListNumberResponseBody) SetMessage(v string) *ListNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListNumberResponseBody) SetRequestId(v string) *ListNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNumberResponseBody) SetTotal(v int64) *ListNumberResponseBody {
	s.Total = &v
	return s
}

type ListNumberResponseBodyList struct {
	ApplyId             *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	CountryId           *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	DisableTs           *string `json:"DisableTs,omitempty" xml:"DisableTs,omitempty"`
	InboundConcurrency  *int64  `json:"InboundConcurrency,omitempty" xml:"InboundConcurrency,omitempty"`
	Number              *string `json:"Number,omitempty" xml:"Number,omitempty"`
	NumberName          *string `json:"NumberName,omitempty" xml:"NumberName,omitempty"`
	OutboundConcurrency *int64  `json:"OutboundConcurrency,omitempty" xml:"OutboundConcurrency,omitempty"`
	PhoneType           *int64  `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	QualificationId     *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceId          *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status              *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierId          *string `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
	Support             *int64  `json:"Support,omitempty" xml:"Support,omitempty"`
	UpdateTs            *string `json:"UpdateTs,omitempty" xml:"UpdateTs,omitempty"`
}

func (s ListNumberResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListNumberResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListNumberResponseBodyList) SetApplyId(v string) *ListNumberResponseBodyList {
	s.ApplyId = &v
	return s
}

func (s *ListNumberResponseBodyList) SetCountryId(v string) *ListNumberResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListNumberResponseBodyList) SetDisableTs(v string) *ListNumberResponseBodyList {
	s.DisableTs = &v
	return s
}

func (s *ListNumberResponseBodyList) SetInboundConcurrency(v int64) *ListNumberResponseBodyList {
	s.InboundConcurrency = &v
	return s
}

func (s *ListNumberResponseBodyList) SetNumber(v string) *ListNumberResponseBodyList {
	s.Number = &v
	return s
}

func (s *ListNumberResponseBodyList) SetNumberName(v string) *ListNumberResponseBodyList {
	s.NumberName = &v
	return s
}

func (s *ListNumberResponseBodyList) SetOutboundConcurrency(v int64) *ListNumberResponseBodyList {
	s.OutboundConcurrency = &v
	return s
}

func (s *ListNumberResponseBodyList) SetPhoneType(v int64) *ListNumberResponseBodyList {
	s.PhoneType = &v
	return s
}

func (s *ListNumberResponseBodyList) SetQualificationId(v string) *ListNumberResponseBodyList {
	s.QualificationId = &v
	return s
}

func (s *ListNumberResponseBodyList) SetResourceId(v string) *ListNumberResponseBodyList {
	s.ResourceId = &v
	return s
}

func (s *ListNumberResponseBodyList) SetStatus(v int64) *ListNumberResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListNumberResponseBodyList) SetSupplierId(v string) *ListNumberResponseBodyList {
	s.SupplierId = &v
	return s
}

func (s *ListNumberResponseBodyList) SetSupport(v int64) *ListNumberResponseBodyList {
	s.Support = &v
	return s
}

func (s *ListNumberResponseBodyList) SetUpdateTs(v string) *ListNumberResponseBodyList {
	s.UpdateTs = &v
	return s
}

type ListNumberResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNumberResponse) GoString() string {
	return s.String()
}

func (s *ListNumberResponse) SetHeaders(v map[string]*string) *ListNumberResponse {
	s.Headers = v
	return s
}

func (s *ListNumberResponse) SetStatusCode(v int32) *ListNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNumberResponse) SetBody(v *ListNumberResponseBody) *ListNumberResponse {
	s.Body = v
	return s
}

type ListQualificationRequest struct {
	CompanyName          *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	ContactPhone         *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListQualificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQualificationRequest) GoString() string {
	return s.String()
}

func (s *ListQualificationRequest) SetCompanyName(v string) *ListQualificationRequest {
	s.CompanyName = &v
	return s
}

func (s *ListQualificationRequest) SetContactPhone(v string) *ListQualificationRequest {
	s.ContactPhone = &v
	return s
}

func (s *ListQualificationRequest) SetCountryId(v string) *ListQualificationRequest {
	s.CountryId = &v
	return s
}

func (s *ListQualificationRequest) SetEndTime(v string) *ListQualificationRequest {
	s.EndTime = &v
	return s
}

func (s *ListQualificationRequest) SetOwnerId(v int64) *ListQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ListQualificationRequest) SetPageNumber(v int64) *ListQualificationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQualificationRequest) SetPageSize(v int64) *ListQualificationRequest {
	s.PageSize = &v
	return s
}

func (s *ListQualificationRequest) SetResourceOwnerAccount(v string) *ListQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListQualificationRequest) SetResourceOwnerId(v int64) *ListQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListQualificationRequest) SetStartTime(v string) *ListQualificationRequest {
	s.StartTime = &v
	return s
}

func (s *ListQualificationRequest) SetStatus(v int64) *ListQualificationRequest {
	s.Status = &v
	return s
}

type ListQualificationResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListQualificationResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListQualificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualificationResponseBody) SetCode(v string) *ListQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualificationResponseBody) SetList(v []*ListQualificationResponseBodyList) *ListQualificationResponseBody {
	s.List = v
	return s
}

func (s *ListQualificationResponseBody) SetMessage(v string) *ListQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualificationResponseBody) SetRequestId(v string) *ListQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualificationResponseBody) SetTotal(v int64) *ListQualificationResponseBody {
	s.Total = &v
	return s
}

type ListQualificationResponseBodyList struct {
	Address                *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AuditRemark            *string `json:"AuditRemark,omitempty" xml:"AuditRemark,omitempty"`
	AuditTs                *string `json:"AuditTs,omitempty" xml:"AuditTs,omitempty"`
	BusinessLicenseFileKey *string `json:"BusinessLicenseFileKey,omitempty" xml:"BusinessLicenseFileKey,omitempty"`
	CommitTs               *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CompanyName            *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	ContactEmail           *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName            *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactPhone           *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	CountryId              *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	LegalId                *string `json:"LegalId,omitempty" xml:"LegalId,omitempty"`
	LegalLicenseFileKey    *string `json:"LegalLicenseFileKey,omitempty" xml:"LegalLicenseFileKey,omitempty"`
	LegalName              *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	NetworkAccessFileKey   *string `json:"NetworkAccessFileKey,omitempty" xml:"NetworkAccessFileKey,omitempty"`
	QualificationId        *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	Status                 *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	UnifiedCode            *string `json:"UnifiedCode,omitempty" xml:"UnifiedCode,omitempty"`
	UpdateTs               *string `json:"UpdateTs,omitempty" xml:"UpdateTs,omitempty"`
}

func (s ListQualificationResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListQualificationResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListQualificationResponseBodyList) SetAddress(v string) *ListQualificationResponseBodyList {
	s.Address = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetAuditRemark(v string) *ListQualificationResponseBodyList {
	s.AuditRemark = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetAuditTs(v string) *ListQualificationResponseBodyList {
	s.AuditTs = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetBusinessLicenseFileKey(v string) *ListQualificationResponseBodyList {
	s.BusinessLicenseFileKey = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetCommitTs(v string) *ListQualificationResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetCompanyName(v string) *ListQualificationResponseBodyList {
	s.CompanyName = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetContactEmail(v string) *ListQualificationResponseBodyList {
	s.ContactEmail = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetContactName(v string) *ListQualificationResponseBodyList {
	s.ContactName = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetContactPhone(v string) *ListQualificationResponseBodyList {
	s.ContactPhone = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetCountryId(v string) *ListQualificationResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetLegalId(v string) *ListQualificationResponseBodyList {
	s.LegalId = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetLegalLicenseFileKey(v string) *ListQualificationResponseBodyList {
	s.LegalLicenseFileKey = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetLegalName(v string) *ListQualificationResponseBodyList {
	s.LegalName = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetNetworkAccessFileKey(v string) *ListQualificationResponseBodyList {
	s.NetworkAccessFileKey = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetQualificationId(v string) *ListQualificationResponseBodyList {
	s.QualificationId = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetStatus(v int64) *ListQualificationResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetUnifiedCode(v string) *ListQualificationResponseBodyList {
	s.UnifiedCode = &v
	return s
}

func (s *ListQualificationResponseBodyList) SetUpdateTs(v string) *ListQualificationResponseBodyList {
	s.UpdateTs = &v
	return s
}

type ListQualificationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQualificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQualificationResponse) GoString() string {
	return s.String()
}

func (s *ListQualificationResponse) SetHeaders(v map[string]*string) *ListQualificationResponse {
	s.Headers = v
	return s
}

func (s *ListQualificationResponse) SetStatusCode(v int32) *ListQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualificationResponse) SetBody(v *ListQualificationResponseBody) *ListQualificationResponse {
	s.Body = v
	return s
}

type ListReceiptRequest struct {
	BusinessType         *int32  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReceiptRequest) GoString() string {
	return s.String()
}

func (s *ListReceiptRequest) SetBusinessType(v int32) *ListReceiptRequest {
	s.BusinessType = &v
	return s
}

func (s *ListReceiptRequest) SetEndTime(v string) *ListReceiptRequest {
	s.EndTime = &v
	return s
}

func (s *ListReceiptRequest) SetOwnerId(v int64) *ListReceiptRequest {
	s.OwnerId = &v
	return s
}

func (s *ListReceiptRequest) SetResourceOwnerAccount(v string) *ListReceiptRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListReceiptRequest) SetResourceOwnerId(v int64) *ListReceiptRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListReceiptRequest) SetStartTime(v string) *ListReceiptRequest {
	s.StartTime = &v
	return s
}

type ListReceiptResponseBody struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	List       []*ListReceiptResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	ReceiptUrl *string                        `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *ListReceiptResponseBody) SetCode(v string) *ListReceiptResponseBody {
	s.Code = &v
	return s
}

func (s *ListReceiptResponseBody) SetList(v []*ListReceiptResponseBodyList) *ListReceiptResponseBody {
	s.List = v
	return s
}

func (s *ListReceiptResponseBody) SetMessage(v string) *ListReceiptResponseBody {
	s.Message = &v
	return s
}

func (s *ListReceiptResponseBody) SetReceiptUrl(v string) *ListReceiptResponseBody {
	s.ReceiptUrl = &v
	return s
}

func (s *ListReceiptResponseBody) SetRequestId(v string) *ListReceiptResponseBody {
	s.RequestId = &v
	return s
}

type ListReceiptResponseBodyList struct {
	CallFailedCount  *int32  `json:"CallFailedCount,omitempty" xml:"CallFailedCount,omitempty"`
	CallSuccessCount *int32  `json:"CallSuccessCount,omitempty" xml:"CallSuccessCount,omitempty"`
	CallTotalCount   *int32  `json:"CallTotalCount,omitempty" xml:"CallTotalCount,omitempty"`
	Date             *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s ListReceiptResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListReceiptResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListReceiptResponseBodyList) SetCallFailedCount(v int32) *ListReceiptResponseBodyList {
	s.CallFailedCount = &v
	return s
}

func (s *ListReceiptResponseBodyList) SetCallSuccessCount(v int32) *ListReceiptResponseBodyList {
	s.CallSuccessCount = &v
	return s
}

func (s *ListReceiptResponseBodyList) SetCallTotalCount(v int32) *ListReceiptResponseBodyList {
	s.CallTotalCount = &v
	return s
}

func (s *ListReceiptResponseBodyList) SetDate(v string) *ListReceiptResponseBodyList {
	s.Date = &v
	return s
}

type ListReceiptResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReceiptResponse) GoString() string {
	return s.String()
}

func (s *ListReceiptResponse) SetHeaders(v map[string]*string) *ListReceiptResponse {
	s.Headers = v
	return s
}

func (s *ListReceiptResponse) SetStatusCode(v int32) *ListReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReceiptResponse) SetBody(v *ListReceiptResponseBody) *ListReceiptResponse {
	s.Body = v
	return s
}

type ListSipTrunkRequest struct {
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSipTrunkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkRequest) GoString() string {
	return s.String()
}

func (s *ListSipTrunkRequest) SetEndTime(v string) *ListSipTrunkRequest {
	s.EndTime = &v
	return s
}

func (s *ListSipTrunkRequest) SetOwnerId(v int64) *ListSipTrunkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSipTrunkRequest) SetPageNumber(v int64) *ListSipTrunkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSipTrunkRequest) SetPageSize(v int64) *ListSipTrunkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSipTrunkRequest) SetResourceOwnerAccount(v string) *ListSipTrunkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSipTrunkRequest) SetResourceOwnerId(v int64) *ListSipTrunkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSipTrunkRequest) SetStartTime(v string) *ListSipTrunkRequest {
	s.StartTime = &v
	return s
}

func (s *ListSipTrunkRequest) SetStatus(v int64) *ListSipTrunkRequest {
	s.Status = &v
	return s
}

type ListSipTrunkResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListSipTrunkResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *string                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSipTrunkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkResponseBody) GoString() string {
	return s.String()
}

func (s *ListSipTrunkResponseBody) SetCode(v string) *ListSipTrunkResponseBody {
	s.Code = &v
	return s
}

func (s *ListSipTrunkResponseBody) SetList(v []*ListSipTrunkResponseBodyList) *ListSipTrunkResponseBody {
	s.List = v
	return s
}

func (s *ListSipTrunkResponseBody) SetMessage(v string) *ListSipTrunkResponseBody {
	s.Message = &v
	return s
}

func (s *ListSipTrunkResponseBody) SetRequestId(v string) *ListSipTrunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSipTrunkResponseBody) SetTotal(v string) *ListSipTrunkResponseBody {
	s.Total = &v
	return s
}

type ListSipTrunkResponseBodyList struct {
	ApplyId         *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	ApplyNote       *string `json:"ApplyNote,omitempty" xml:"ApplyNote,omitempty"`
	AuditNote       *string `json:"AuditNote,omitempty" xml:"AuditNote,omitempty"`
	AuditTs         *string `json:"AuditTs,omitempty" xml:"AuditTs,omitempty"`
	CommitTs        *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CountryId       *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	InboundIpPorts  *string `json:"InboundIpPorts,omitempty" xml:"InboundIpPorts,omitempty"`
	OutboundIps     *string `json:"OutboundIps,omitempty" xml:"OutboundIps,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	Scene           *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTs        *string `json:"UpdateTs,omitempty" xml:"UpdateTs,omitempty"`
	UserUuid        *string `json:"UserUuid,omitempty" xml:"UserUuid,omitempty"`
}

func (s ListSipTrunkResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSipTrunkResponseBodyList) SetApplyId(v string) *ListSipTrunkResponseBodyList {
	s.ApplyId = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetApplyNote(v string) *ListSipTrunkResponseBodyList {
	s.ApplyNote = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetAuditNote(v string) *ListSipTrunkResponseBodyList {
	s.AuditNote = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetAuditTs(v string) *ListSipTrunkResponseBodyList {
	s.AuditTs = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetCommitTs(v string) *ListSipTrunkResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetCountryId(v string) *ListSipTrunkResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetInboundIpPorts(v string) *ListSipTrunkResponseBodyList {
	s.InboundIpPorts = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetOutboundIps(v string) *ListSipTrunkResponseBodyList {
	s.OutboundIps = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetQualificationId(v string) *ListSipTrunkResponseBodyList {
	s.QualificationId = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetScene(v string) *ListSipTrunkResponseBodyList {
	s.Scene = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetStatus(v string) *ListSipTrunkResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetUpdateTs(v string) *ListSipTrunkResponseBodyList {
	s.UpdateTs = &v
	return s
}

func (s *ListSipTrunkResponseBodyList) SetUserUuid(v string) *ListSipTrunkResponseBodyList {
	s.UserUuid = &v
	return s
}

type ListSipTrunkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSipTrunkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSipTrunkResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkResponse) GoString() string {
	return s.String()
}

func (s *ListSipTrunkResponse) SetHeaders(v map[string]*string) *ListSipTrunkResponse {
	s.Headers = v
	return s
}

func (s *ListSipTrunkResponse) SetStatusCode(v int32) *ListSipTrunkResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSipTrunkResponse) SetBody(v *ListSipTrunkResponseBody) *ListSipTrunkResponse {
	s.Body = v
	return s
}

type ListSipTrunkDetailRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSipTrunkDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkDetailRequest) GoString() string {
	return s.String()
}

func (s *ListSipTrunkDetailRequest) SetCalledNumber(v string) *ListSipTrunkDetailRequest {
	s.CalledNumber = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetCallingNumber(v string) *ListSipTrunkDetailRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetCountryId(v string) *ListSipTrunkDetailRequest {
	s.CountryId = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetEndTime(v string) *ListSipTrunkDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetOwnerId(v int64) *ListSipTrunkDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetPageNumber(v int64) *ListSipTrunkDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetPageSize(v int64) *ListSipTrunkDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetResourceOwnerAccount(v string) *ListSipTrunkDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetResourceOwnerId(v int64) *ListSipTrunkDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSipTrunkDetailRequest) SetStartTime(v string) *ListSipTrunkDetailRequest {
	s.StartTime = &v
	return s
}

type ListSipTrunkDetailResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListSipTrunkDetailResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSipTrunkDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListSipTrunkDetailResponseBody) SetCode(v string) *ListSipTrunkDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListSipTrunkDetailResponseBody) SetList(v []*ListSipTrunkDetailResponseBodyList) *ListSipTrunkDetailResponseBody {
	s.List = v
	return s
}

func (s *ListSipTrunkDetailResponseBody) SetMessage(v string) *ListSipTrunkDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListSipTrunkDetailResponseBody) SetRequestId(v string) *ListSipTrunkDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSipTrunkDetailResponseBody) SetTotal(v int64) *ListSipTrunkDetailResponseBody {
	s.Total = &v
	return s
}

type ListSipTrunkDetailResponseBodyList struct {
	AnswerTime      *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	Answered        *int64  `json:"Answered,omitempty" xml:"Answered,omitempty"`
	BusiType        *string `json:"BusiType,omitempty" xml:"BusiType,omitempty"`
	CallEndTime     *string `json:"CallEndTime,omitempty" xml:"CallEndTime,omitempty"`
	CallType        *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CalledIp        *string `json:"CalledIp,omitempty" xml:"CalledIp,omitempty"`
	CalledNum       *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	CalledNumRegion *string `json:"CalledNumRegion,omitempty" xml:"CalledNumRegion,omitempty"`
	CalledNumType   *string `json:"CalledNumType,omitempty" xml:"CalledNumType,omitempty"`
	CallerDisplay   *string `json:"CallerDisplay,omitempty" xml:"CallerDisplay,omitempty"`
	CallerIp        *string `json:"CallerIp,omitempty" xml:"CallerIp,omitempty"`
	CallerNum       *string `json:"CallerNum,omitempty" xml:"CallerNum,omitempty"`
	CallerNumRegion *string `json:"CallerNumRegion,omitempty" xml:"CallerNumRegion,omitempty"`
	CallerNumType   *string `json:"CallerNumType,omitempty" xml:"CallerNumType,omitempty"`
	Duration        *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Hangup          *int64  `json:"Hangup,omitempty" xml:"Hangup,omitempty"`
	HangupType      *int64  `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	IvnCliType      *string `json:"IvnCliType,omitempty" xml:"IvnCliType,omitempty"`
	RecordUrl       *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserUuid        *string `json:"UserUuid,omitempty" xml:"UserUuid,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListSipTrunkDetailResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkDetailResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSipTrunkDetailResponseBodyList) SetAnswerTime(v string) *ListSipTrunkDetailResponseBodyList {
	s.AnswerTime = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetAnswered(v int64) *ListSipTrunkDetailResponseBodyList {
	s.Answered = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetBusiType(v string) *ListSipTrunkDetailResponseBodyList {
	s.BusiType = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallEndTime(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallEndTime = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallType(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallType = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCalledIp(v string) *ListSipTrunkDetailResponseBodyList {
	s.CalledIp = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCalledNum(v string) *ListSipTrunkDetailResponseBodyList {
	s.CalledNum = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCalledNumRegion(v string) *ListSipTrunkDetailResponseBodyList {
	s.CalledNumRegion = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCalledNumType(v string) *ListSipTrunkDetailResponseBodyList {
	s.CalledNumType = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallerDisplay(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallerDisplay = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallerIp(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallerIp = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallerNum(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallerNum = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallerNumRegion(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallerNumRegion = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetCallerNumType(v string) *ListSipTrunkDetailResponseBodyList {
	s.CallerNumType = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetDuration(v int64) *ListSipTrunkDetailResponseBodyList {
	s.Duration = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetHangup(v int64) *ListSipTrunkDetailResponseBodyList {
	s.Hangup = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetHangupType(v int64) *ListSipTrunkDetailResponseBodyList {
	s.HangupType = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetIvnCliType(v string) *ListSipTrunkDetailResponseBodyList {
	s.IvnCliType = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetRecordUrl(v string) *ListSipTrunkDetailResponseBodyList {
	s.RecordUrl = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetStartTime(v string) *ListSipTrunkDetailResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetUserUuid(v string) *ListSipTrunkDetailResponseBodyList {
	s.UserUuid = &v
	return s
}

func (s *ListSipTrunkDetailResponseBodyList) SetUuid(v string) *ListSipTrunkDetailResponseBodyList {
	s.Uuid = &v
	return s
}

type ListSipTrunkDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSipTrunkDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSipTrunkDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkDetailResponse) GoString() string {
	return s.String()
}

func (s *ListSipTrunkDetailResponse) SetHeaders(v map[string]*string) *ListSipTrunkDetailResponse {
	s.Headers = v
	return s
}

func (s *ListSipTrunkDetailResponse) SetStatusCode(v int32) *ListSipTrunkDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSipTrunkDetailResponse) SetBody(v *ListSipTrunkDetailResponseBody) *ListSipTrunkDetailResponse {
	s.Body = v
	return s
}

type ListSipTrunkStatRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSipTrunkStatRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkStatRequest) GoString() string {
	return s.String()
}

func (s *ListSipTrunkStatRequest) SetCalledNumber(v string) *ListSipTrunkStatRequest {
	s.CalledNumber = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetCallingNumber(v string) *ListSipTrunkStatRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetEndTime(v string) *ListSipTrunkStatRequest {
	s.EndTime = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetOwnerId(v int64) *ListSipTrunkStatRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetPageNumber(v int64) *ListSipTrunkStatRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetPageSize(v int64) *ListSipTrunkStatRequest {
	s.PageSize = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetResourceOwnerAccount(v string) *ListSipTrunkStatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetResourceOwnerId(v int64) *ListSipTrunkStatRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSipTrunkStatRequest) SetStartTime(v string) *ListSipTrunkStatRequest {
	s.StartTime = &v
	return s
}

type ListSipTrunkStatResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListSipTrunkStatResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *string                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSipTrunkStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkStatResponseBody) GoString() string {
	return s.String()
}

func (s *ListSipTrunkStatResponseBody) SetCode(v string) *ListSipTrunkStatResponseBody {
	s.Code = &v
	return s
}

func (s *ListSipTrunkStatResponseBody) SetList(v []*ListSipTrunkStatResponseBodyList) *ListSipTrunkStatResponseBody {
	s.List = v
	return s
}

func (s *ListSipTrunkStatResponseBody) SetMessage(v string) *ListSipTrunkStatResponseBody {
	s.Message = &v
	return s
}

func (s *ListSipTrunkStatResponseBody) SetRequestId(v string) *ListSipTrunkStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSipTrunkStatResponseBody) SetTotal(v string) *ListSipTrunkStatResponseBody {
	s.Total = &v
	return s
}

type ListSipTrunkStatResponseBodyList struct {
	AnsweredCalls   *int64  `json:"AnsweredCalls,omitempty" xml:"AnsweredCalls,omitempty"`
	AnsweredRate    *int64  `json:"AnsweredRate,omitempty" xml:"AnsweredRate,omitempty"`
	AverageDuration *int64  `json:"AverageDuration,omitempty" xml:"AverageDuration,omitempty"`
	Billing         *int64  `json:"Billing,omitempty" xml:"Billing,omitempty"`
	CalledNumber    *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber   *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	FailedCalls     *int64  `json:"FailedCalls,omitempty" xml:"FailedCalls,omitempty"`
	TotalCalls      *int64  `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	TotalDuration   *int64  `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
}

func (s ListSipTrunkStatResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkStatResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSipTrunkStatResponseBodyList) SetAnsweredCalls(v int64) *ListSipTrunkStatResponseBodyList {
	s.AnsweredCalls = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetAnsweredRate(v int64) *ListSipTrunkStatResponseBodyList {
	s.AnsweredRate = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetAverageDuration(v int64) *ListSipTrunkStatResponseBodyList {
	s.AverageDuration = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetBilling(v int64) *ListSipTrunkStatResponseBodyList {
	s.Billing = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetCalledNumber(v string) *ListSipTrunkStatResponseBodyList {
	s.CalledNumber = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetCallingNumber(v string) *ListSipTrunkStatResponseBodyList {
	s.CallingNumber = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetFailedCalls(v int64) *ListSipTrunkStatResponseBodyList {
	s.FailedCalls = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetTotalCalls(v int64) *ListSipTrunkStatResponseBodyList {
	s.TotalCalls = &v
	return s
}

func (s *ListSipTrunkStatResponseBodyList) SetTotalDuration(v int64) *ListSipTrunkStatResponseBodyList {
	s.TotalDuration = &v
	return s
}

type ListSipTrunkStatResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSipTrunkStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSipTrunkStatResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSipTrunkStatResponse) GoString() string {
	return s.String()
}

func (s *ListSipTrunkStatResponse) SetHeaders(v map[string]*string) *ListSipTrunkStatResponse {
	s.Headers = v
	return s
}

func (s *ListSipTrunkStatResponse) SetStatusCode(v int32) *ListSipTrunkStatResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSipTrunkStatResponse) SetBody(v *ListSipTrunkStatResponseBody) *ListSipTrunkStatResponse {
	s.Body = v
	return s
}

type ListVoiceCallRequest struct {
	BusinessType         *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SendType             *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListVoiceCallRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceCallRequest) SetBusinessType(v int64) *ListVoiceCallRequest {
	s.BusinessType = &v
	return s
}

func (s *ListVoiceCallRequest) SetCallingNumber(v string) *ListVoiceCallRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListVoiceCallRequest) SetCountryId(v string) *ListVoiceCallRequest {
	s.CountryId = &v
	return s
}

func (s *ListVoiceCallRequest) SetEndTime(v string) *ListVoiceCallRequest {
	s.EndTime = &v
	return s
}

func (s *ListVoiceCallRequest) SetOwnerId(v int64) *ListVoiceCallRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVoiceCallRequest) SetPageNumber(v int64) *ListVoiceCallRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceCallRequest) SetPageSize(v int64) *ListVoiceCallRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceCallRequest) SetResourceOwnerAccount(v string) *ListVoiceCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVoiceCallRequest) SetResourceOwnerId(v int64) *ListVoiceCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVoiceCallRequest) SetSendType(v int64) *ListVoiceCallRequest {
	s.SendType = &v
	return s
}

func (s *ListVoiceCallRequest) SetStartTime(v string) *ListVoiceCallRequest {
	s.StartTime = &v
	return s
}

func (s *ListVoiceCallRequest) SetStatus(v int64) *ListVoiceCallRequest {
	s.Status = &v
	return s
}

func (s *ListVoiceCallRequest) SetTaskId(v string) *ListVoiceCallRequest {
	s.TaskId = &v
	return s
}

func (s *ListVoiceCallRequest) SetTaskName(v string) *ListVoiceCallRequest {
	s.TaskName = &v
	return s
}

type ListVoiceCallResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListVoiceCallResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVoiceCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceCallResponseBody) SetCode(v string) *ListVoiceCallResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceCallResponseBody) SetList(v []*ListVoiceCallResponseBodyList) *ListVoiceCallResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceCallResponseBody) SetMessage(v string) *ListVoiceCallResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceCallResponseBody) SetRequestId(v string) *ListVoiceCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceCallResponseBody) SetTotal(v int64) *ListVoiceCallResponseBody {
	s.Total = &v
	return s
}

type ListVoiceCallResponseBodyList struct {
	AnsweredCalls   *int64  `json:"AnsweredCalls,omitempty" xml:"AnsweredCalls,omitempty"`
	BusinessType    *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CalledCalls     *int64  `json:"CalledCalls,omitempty" xml:"CalledCalls,omitempty"`
	CallingNumber   *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CommitTs        *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CountryId       *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	FailedCalls     *int64  `json:"FailedCalls,omitempty" xml:"FailedCalls,omitempty"`
	GroupCallType   *int64  `json:"GroupCallType,omitempty" xml:"GroupCallType,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SendType        *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TickerTs        *string `json:"TickerTs,omitempty" xml:"TickerTs,omitempty"`
	TotalCalls      *int64  `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	TotalDuration   *int64  `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	UserUuid        *string `json:"UserUuid,omitempty" xml:"UserUuid,omitempty"`
}

func (s ListVoiceCallResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVoiceCallResponseBodyList) SetAnsweredCalls(v int64) *ListVoiceCallResponseBodyList {
	s.AnsweredCalls = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetBusinessType(v int64) *ListVoiceCallResponseBodyList {
	s.BusinessType = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetCalledCalls(v int64) *ListVoiceCallResponseBodyList {
	s.CalledCalls = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetCallingNumber(v string) *ListVoiceCallResponseBodyList {
	s.CallingNumber = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetCommitTs(v string) *ListVoiceCallResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetCountryId(v string) *ListVoiceCallResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetFailedCalls(v int64) *ListVoiceCallResponseBodyList {
	s.FailedCalls = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetGroupCallType(v int64) *ListVoiceCallResponseBodyList {
	s.GroupCallType = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetRequestId(v string) *ListVoiceCallResponseBodyList {
	s.RequestId = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetSendType(v int64) *ListVoiceCallResponseBodyList {
	s.SendType = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetStatus(v string) *ListVoiceCallResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTaskId(v string) *ListVoiceCallResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTaskName(v string) *ListVoiceCallResponseBodyList {
	s.TaskName = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTemplateContent(v string) *ListVoiceCallResponseBodyList {
	s.TemplateContent = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTemplateId(v string) *ListVoiceCallResponseBodyList {
	s.TemplateId = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTickerTs(v string) *ListVoiceCallResponseBodyList {
	s.TickerTs = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTotalCalls(v int64) *ListVoiceCallResponseBodyList {
	s.TotalCalls = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetTotalDuration(v int64) *ListVoiceCallResponseBodyList {
	s.TotalDuration = &v
	return s
}

func (s *ListVoiceCallResponseBodyList) SetUserUuid(v string) *ListVoiceCallResponseBodyList {
	s.UserUuid = &v
	return s
}

type ListVoiceCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVoiceCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVoiceCallResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceCallResponse) SetHeaders(v map[string]*string) *ListVoiceCallResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceCallResponse) SetStatusCode(v int32) *ListVoiceCallResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceCallResponse) SetBody(v *ListVoiceCallResponseBody) *ListVoiceCallResponse {
	s.Body = v
	return s
}

type ListVoiceCallDetailRequest struct {
	BusinessType         *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HangupType           *int64  `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListVoiceCallDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallDetailRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceCallDetailRequest) SetBusinessType(v int64) *ListVoiceCallDetailRequest {
	s.BusinessType = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetCalledNumber(v string) *ListVoiceCallDetailRequest {
	s.CalledNumber = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetCallingNumber(v string) *ListVoiceCallDetailRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetCountryId(v string) *ListVoiceCallDetailRequest {
	s.CountryId = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetEndTime(v string) *ListVoiceCallDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetHangupType(v int64) *ListVoiceCallDetailRequest {
	s.HangupType = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetOwnerId(v int64) *ListVoiceCallDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetPageNumber(v int64) *ListVoiceCallDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetPageSize(v int64) *ListVoiceCallDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetResourceOwnerAccount(v string) *ListVoiceCallDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetResourceOwnerId(v int64) *ListVoiceCallDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetStartTime(v string) *ListVoiceCallDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetStatus(v int64) *ListVoiceCallDetailRequest {
	s.Status = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetTaskId(v string) *ListVoiceCallDetailRequest {
	s.TaskId = &v
	return s
}

func (s *ListVoiceCallDetailRequest) SetTaskName(v string) *ListVoiceCallDetailRequest {
	s.TaskName = &v
	return s
}

type ListVoiceCallDetailResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListVoiceCallDetailResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVoiceCallDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceCallDetailResponseBody) SetCode(v string) *ListVoiceCallDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceCallDetailResponseBody) SetList(v []*ListVoiceCallDetailResponseBodyList) *ListVoiceCallDetailResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceCallDetailResponseBody) SetMessage(v string) *ListVoiceCallDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceCallDetailResponseBody) SetRequestId(v string) *ListVoiceCallDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceCallDetailResponseBody) SetTotal(v int64) *ListVoiceCallDetailResponseBody {
	s.Total = &v
	return s
}

type ListVoiceCallDetailResponseBodyList struct {
	Billing       *int64  `json:"Billing,omitempty" xml:"Billing,omitempty"`
	BusinessType  *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CalledNumber  *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CommitTs      *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CountryId     *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTs         *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	HangupType    *int64  `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	SendType      *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	StartTs       *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	Status        *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListVoiceCallDetailResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallDetailResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVoiceCallDetailResponseBodyList) SetBilling(v int64) *ListVoiceCallDetailResponseBodyList {
	s.Billing = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetBusinessType(v int64) *ListVoiceCallDetailResponseBodyList {
	s.BusinessType = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetCalledNumber(v string) *ListVoiceCallDetailResponseBodyList {
	s.CalledNumber = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetCallingNumber(v string) *ListVoiceCallDetailResponseBodyList {
	s.CallingNumber = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetCommitTs(v string) *ListVoiceCallDetailResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetCountryId(v string) *ListVoiceCallDetailResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetDuration(v int64) *ListVoiceCallDetailResponseBodyList {
	s.Duration = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetEndTs(v string) *ListVoiceCallDetailResponseBodyList {
	s.EndTs = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetHangupType(v int64) *ListVoiceCallDetailResponseBodyList {
	s.HangupType = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetSendType(v int64) *ListVoiceCallDetailResponseBodyList {
	s.SendType = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetStartTs(v string) *ListVoiceCallDetailResponseBodyList {
	s.StartTs = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetStatus(v int64) *ListVoiceCallDetailResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetTaskId(v string) *ListVoiceCallDetailResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *ListVoiceCallDetailResponseBodyList) SetTaskName(v string) *ListVoiceCallDetailResponseBodyList {
	s.TaskName = &v
	return s
}

type ListVoiceCallDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVoiceCallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVoiceCallDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallDetailResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceCallDetailResponse) SetHeaders(v map[string]*string) *ListVoiceCallDetailResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceCallDetailResponse) SetStatusCode(v int32) *ListVoiceCallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceCallDetailResponse) SetBody(v *ListVoiceCallDetailResponseBody) *ListVoiceCallDetailResponse {
	s.Body = v
	return s
}

type ListVoiceCallStatRequest struct {
	BusinessType         *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListVoiceCallStatRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallStatRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceCallStatRequest) SetBusinessType(v int64) *ListVoiceCallStatRequest {
	s.BusinessType = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetEndTime(v string) *ListVoiceCallStatRequest {
	s.EndTime = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetOwnerId(v int64) *ListVoiceCallStatRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetPageNumber(v int64) *ListVoiceCallStatRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetPageSize(v int64) *ListVoiceCallStatRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetResourceOwnerAccount(v string) *ListVoiceCallStatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetResourceOwnerId(v int64) *ListVoiceCallStatRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetStartTime(v string) *ListVoiceCallStatRequest {
	s.StartTime = &v
	return s
}

func (s *ListVoiceCallStatRequest) SetTaskName(v string) *ListVoiceCallStatRequest {
	s.TaskName = &v
	return s
}

type ListVoiceCallStatResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListVoiceCallStatResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVoiceCallStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallStatResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceCallStatResponseBody) SetCode(v string) *ListVoiceCallStatResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceCallStatResponseBody) SetList(v []*ListVoiceCallStatResponseBodyList) *ListVoiceCallStatResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceCallStatResponseBody) SetMessage(v string) *ListVoiceCallStatResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceCallStatResponseBody) SetRequestId(v string) *ListVoiceCallStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceCallStatResponseBody) SetTotal(v int64) *ListVoiceCallStatResponseBody {
	s.Total = &v
	return s
}

type ListVoiceCallStatResponseBodyList struct {
	AnsweredCalls   *int64 `json:"AnsweredCalls,omitempty" xml:"AnsweredCalls,omitempty"`
	AnsweredRate    *int64 `json:"AnsweredRate,omitempty" xml:"AnsweredRate,omitempty"`
	AverageDuration *int64 `json:"AverageDuration,omitempty" xml:"AverageDuration,omitempty"`
	Billing         *int64 `json:"Billing,omitempty" xml:"Billing,omitempty"`
	BusinessType    *int64 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	FailedCalls     *int64 `json:"FailedCalls,omitempty" xml:"FailedCalls,omitempty"`
	TotalCalls      *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	TotalDuration   *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
}

func (s ListVoiceCallStatResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallStatResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVoiceCallStatResponseBodyList) SetAnsweredCalls(v int64) *ListVoiceCallStatResponseBodyList {
	s.AnsweredCalls = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetAnsweredRate(v int64) *ListVoiceCallStatResponseBodyList {
	s.AnsweredRate = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetAverageDuration(v int64) *ListVoiceCallStatResponseBodyList {
	s.AverageDuration = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetBilling(v int64) *ListVoiceCallStatResponseBodyList {
	s.Billing = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetBusinessType(v int64) *ListVoiceCallStatResponseBodyList {
	s.BusinessType = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetFailedCalls(v int64) *ListVoiceCallStatResponseBodyList {
	s.FailedCalls = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetTotalCalls(v int64) *ListVoiceCallStatResponseBodyList {
	s.TotalCalls = &v
	return s
}

func (s *ListVoiceCallStatResponseBodyList) SetTotalDuration(v int64) *ListVoiceCallStatResponseBodyList {
	s.TotalDuration = &v
	return s
}

type ListVoiceCallStatResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVoiceCallStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVoiceCallStatResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceCallStatResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceCallStatResponse) SetHeaders(v map[string]*string) *ListVoiceCallStatResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceCallStatResponse) SetStatusCode(v int32) *ListVoiceCallStatResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceCallStatResponse) SetBody(v *ListVoiceCallStatResponseBody) *ListVoiceCallStatResponse {
	s.Body = v
	return s
}

type ListVoiceFileRequest struct {
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	FileName             *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListVoiceFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceFileRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceFileRequest) SetCountryId(v string) *ListVoiceFileRequest {
	s.CountryId = &v
	return s
}

func (s *ListVoiceFileRequest) SetFileName(v string) *ListVoiceFileRequest {
	s.FileName = &v
	return s
}

func (s *ListVoiceFileRequest) SetOwnerId(v int64) *ListVoiceFileRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVoiceFileRequest) SetPageNumber(v int64) *ListVoiceFileRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceFileRequest) SetPageSize(v int64) *ListVoiceFileRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceFileRequest) SetResourceOwnerAccount(v string) *ListVoiceFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVoiceFileRequest) SetResourceOwnerId(v int64) *ListVoiceFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVoiceFileRequest) SetStatus(v int64) *ListVoiceFileRequest {
	s.Status = &v
	return s
}

type ListVoiceFileResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListVoiceFileResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVoiceFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceFileResponseBody) SetCode(v string) *ListVoiceFileResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceFileResponseBody) SetList(v []*ListVoiceFileResponseBodyList) *ListVoiceFileResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceFileResponseBody) SetMessage(v string) *ListVoiceFileResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceFileResponseBody) SetRequestId(v string) *ListVoiceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceFileResponseBody) SetTotal(v int64) *ListVoiceFileResponseBody {
	s.Total = &v
	return s
}

type ListVoiceFileResponseBodyList struct {
	AuditRemark     *string `json:"AuditRemark,omitempty" xml:"AuditRemark,omitempty"`
	CommitTs        *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CountryId       *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileKey         *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	Status          *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTs        *string `json:"UpdateTs,omitempty" xml:"UpdateTs,omitempty"`
}

func (s ListVoiceFileResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceFileResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVoiceFileResponseBodyList) SetAuditRemark(v string) *ListVoiceFileResponseBodyList {
	s.AuditRemark = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetCommitTs(v string) *ListVoiceFileResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetCountryId(v string) *ListVoiceFileResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetFileId(v string) *ListVoiceFileResponseBodyList {
	s.FileId = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetFileKey(v string) *ListVoiceFileResponseBodyList {
	s.FileKey = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetFileName(v string) *ListVoiceFileResponseBodyList {
	s.FileName = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetFileUrl(v string) *ListVoiceFileResponseBodyList {
	s.FileUrl = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetQualificationId(v string) *ListVoiceFileResponseBodyList {
	s.QualificationId = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetStatus(v int64) *ListVoiceFileResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListVoiceFileResponseBodyList) SetUpdateTs(v string) *ListVoiceFileResponseBodyList {
	s.UpdateTs = &v
	return s
}

type ListVoiceFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVoiceFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVoiceFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceFileResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceFileResponse) SetHeaders(v map[string]*string) *ListVoiceFileResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceFileResponse) SetStatusCode(v int32) *ListVoiceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceFileResponse) SetBody(v *ListVoiceFileResponseBody) *ListVoiceFileResponse {
	s.Body = v
	return s
}

type ListVoiceTtsRequest struct {
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListVoiceTtsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceTtsRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceTtsRequest) SetCountryId(v string) *ListVoiceTtsRequest {
	s.CountryId = &v
	return s
}

func (s *ListVoiceTtsRequest) SetOwnerId(v int64) *ListVoiceTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVoiceTtsRequest) SetPageNumber(v int64) *ListVoiceTtsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceTtsRequest) SetPageSize(v int64) *ListVoiceTtsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceTtsRequest) SetResourceOwnerAccount(v string) *ListVoiceTtsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVoiceTtsRequest) SetResourceOwnerId(v int64) *ListVoiceTtsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVoiceTtsRequest) SetStatus(v int64) *ListVoiceTtsRequest {
	s.Status = &v
	return s
}

func (s *ListVoiceTtsRequest) SetTemplateName(v string) *ListVoiceTtsRequest {
	s.TemplateName = &v
	return s
}

type ListVoiceTtsResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*ListVoiceTtsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVoiceTtsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceTtsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceTtsResponseBody) SetCode(v string) *ListVoiceTtsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceTtsResponseBody) SetList(v []*ListVoiceTtsResponseBodyList) *ListVoiceTtsResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceTtsResponseBody) SetMessage(v string) *ListVoiceTtsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceTtsResponseBody) SetRequestId(v string) *ListVoiceTtsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceTtsResponseBody) SetTotal(v int64) *ListVoiceTtsResponseBody {
	s.Total = &v
	return s
}

type ListVoiceTtsResponseBodyList struct {
	AuditBy         *string `json:"AuditBy,omitempty" xml:"AuditBy,omitempty"`
	AuditRemark     *string `json:"AuditRemark,omitempty" xml:"AuditRemark,omitempty"`
	AuditTs         *string `json:"AuditTs,omitempty" xml:"AuditTs,omitempty"`
	CommitTs        *string `json:"CommitTs,omitempty" xml:"CommitTs,omitempty"`
	CountryId       *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	Speed           *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateText    *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	UpdateTs        *string `json:"UpdateTs,omitempty" xml:"UpdateTs,omitempty"`
	UserUuid        *string `json:"UserUuid,omitempty" xml:"UserUuid,omitempty"`
}

func (s ListVoiceTtsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceTtsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVoiceTtsResponseBodyList) SetAuditBy(v string) *ListVoiceTtsResponseBodyList {
	s.AuditBy = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetAuditRemark(v string) *ListVoiceTtsResponseBodyList {
	s.AuditRemark = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetAuditTs(v string) *ListVoiceTtsResponseBodyList {
	s.AuditTs = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetCommitTs(v string) *ListVoiceTtsResponseBodyList {
	s.CommitTs = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetCountryId(v string) *ListVoiceTtsResponseBodyList {
	s.CountryId = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetLanguage(v string) *ListVoiceTtsResponseBodyList {
	s.Language = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetQualificationId(v string) *ListVoiceTtsResponseBodyList {
	s.QualificationId = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetSpeed(v int32) *ListVoiceTtsResponseBodyList {
	s.Speed = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetStatus(v int32) *ListVoiceTtsResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetTemplateId(v string) *ListVoiceTtsResponseBodyList {
	s.TemplateId = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetTemplateName(v string) *ListVoiceTtsResponseBodyList {
	s.TemplateName = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetTemplateText(v string) *ListVoiceTtsResponseBodyList {
	s.TemplateText = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetUpdateTs(v string) *ListVoiceTtsResponseBodyList {
	s.UpdateTs = &v
	return s
}

func (s *ListVoiceTtsResponseBodyList) SetUserUuid(v string) *ListVoiceTtsResponseBodyList {
	s.UserUuid = &v
	return s
}

type ListVoiceTtsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVoiceTtsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVoiceTtsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceTtsResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceTtsResponse) SetHeaders(v map[string]*string) *ListVoiceTtsResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceTtsResponse) SetStatusCode(v int32) *ListVoiceTtsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceTtsResponse) SetBody(v *ListVoiceTtsResponseBody) *ListVoiceTtsResponse {
	s.Body = v
	return s
}

type NumberEnableRequest struct {
	Enable               *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s NumberEnableRequest) String() string {
	return tea.Prettify(s)
}

func (s NumberEnableRequest) GoString() string {
	return s.String()
}

func (s *NumberEnableRequest) SetEnable(v bool) *NumberEnableRequest {
	s.Enable = &v
	return s
}

func (s *NumberEnableRequest) SetNumber(v string) *NumberEnableRequest {
	s.Number = &v
	return s
}

func (s *NumberEnableRequest) SetOwnerId(v int64) *NumberEnableRequest {
	s.OwnerId = &v
	return s
}

func (s *NumberEnableRequest) SetResourceOwnerAccount(v string) *NumberEnableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *NumberEnableRequest) SetResourceOwnerId(v int64) *NumberEnableRequest {
	s.ResourceOwnerId = &v
	return s
}

type NumberEnableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NumberEnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NumberEnableResponseBody) GoString() string {
	return s.String()
}

func (s *NumberEnableResponseBody) SetCode(v string) *NumberEnableResponseBody {
	s.Code = &v
	return s
}

func (s *NumberEnableResponseBody) SetMessage(v string) *NumberEnableResponseBody {
	s.Message = &v
	return s
}

func (s *NumberEnableResponseBody) SetRequestId(v string) *NumberEnableResponseBody {
	s.RequestId = &v
	return s
}

type NumberEnableResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *NumberEnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NumberEnableResponse) String() string {
	return tea.Prettify(s)
}

func (s NumberEnableResponse) GoString() string {
	return s.String()
}

func (s *NumberEnableResponse) SetHeaders(v map[string]*string) *NumberEnableResponse {
	s.Headers = v
	return s
}

func (s *NumberEnableResponse) SetStatusCode(v int32) *NumberEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *NumberEnableResponse) SetBody(v *NumberEnableResponseBody) *NumberEnableResponse {
	s.Body = v
	return s
}

type OpenIntlVoiceServiceRequest struct {
	Env                  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenIntlVoiceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenIntlVoiceServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenIntlVoiceServiceRequest) SetEnv(v string) *OpenIntlVoiceServiceRequest {
	s.Env = &v
	return s
}

func (s *OpenIntlVoiceServiceRequest) SetOwnerId(v int64) *OpenIntlVoiceServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenIntlVoiceServiceRequest) SetResourceOwnerAccount(v string) *OpenIntlVoiceServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenIntlVoiceServiceRequest) SetResourceOwnerId(v int64) *OpenIntlVoiceServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

type OpenIntlVoiceServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model     *string `json:"Model,omitempty" xml:"Model,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenIntlVoiceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenIntlVoiceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenIntlVoiceServiceResponseBody) SetCode(v string) *OpenIntlVoiceServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenIntlVoiceServiceResponseBody) SetMessage(v string) *OpenIntlVoiceServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenIntlVoiceServiceResponseBody) SetModel(v string) *OpenIntlVoiceServiceResponseBody {
	s.Model = &v
	return s
}

func (s *OpenIntlVoiceServiceResponseBody) SetRequestId(v string) *OpenIntlVoiceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenIntlVoiceServiceResponseBody) SetSuccess(v bool) *OpenIntlVoiceServiceResponseBody {
	s.Success = &v
	return s
}

type OpenIntlVoiceServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenIntlVoiceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenIntlVoiceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenIntlVoiceServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenIntlVoiceServiceResponse) SetHeaders(v map[string]*string) *OpenIntlVoiceServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenIntlVoiceServiceResponse) SetStatusCode(v int32) *OpenIntlVoiceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenIntlVoiceServiceResponse) SetBody(v *OpenIntlVoiceServiceResponseBody) *OpenIntlVoiceServiceResponse {
	s.Body = v
	return s
}

type OswTest1Request struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OswTest1Request) String() string {
	return tea.Prettify(s)
}

func (s OswTest1Request) GoString() string {
	return s.String()
}

func (s *OswTest1Request) SetOwnerId(v int64) *OswTest1Request {
	s.OwnerId = &v
	return s
}

func (s *OswTest1Request) SetResourceOwnerAccount(v string) *OswTest1Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OswTest1Request) SetResourceOwnerId(v int64) *OswTest1Request {
	s.ResourceOwnerId = &v
	return s
}

type OswTest1ResponseBody struct {
}

func (s OswTest1ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OswTest1ResponseBody) GoString() string {
	return s.String()
}

type OswTest1Response struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OswTest1ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OswTest1Response) String() string {
	return tea.Prettify(s)
}

func (s OswTest1Response) GoString() string {
	return s.String()
}

func (s *OswTest1Response) SetHeaders(v map[string]*string) *OswTest1Response {
	s.Headers = v
	return s
}

func (s *OswTest1Response) SetStatusCode(v int32) *OswTest1Response {
	s.StatusCode = &v
	return s
}

func (s *OswTest1Response) SetBody(v *OswTest1ResponseBody) *OswTest1Response {
	s.Body = v
	return s
}

type QueryFileOssUrlRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FileKey              *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryFileOssUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFileOssUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryFileOssUrlRequest) SetBizType(v string) *QueryFileOssUrlRequest {
	s.BizType = &v
	return s
}

func (s *QueryFileOssUrlRequest) SetFileKey(v string) *QueryFileOssUrlRequest {
	s.FileKey = &v
	return s
}

func (s *QueryFileOssUrlRequest) SetOwnerId(v int64) *QueryFileOssUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryFileOssUrlRequest) SetResourceOwnerAccount(v string) *QueryFileOssUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryFileOssUrlRequest) SetResourceOwnerId(v int64) *QueryFileOssUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryFileOssUrlResponseBody struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryFileOssUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFileOssUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFileOssUrlResponseBody) SetAccessKeyId(v string) *QueryFileOssUrlResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *QueryFileOssUrlResponseBody) SetRequestId(v string) *QueryFileOssUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFileOssUrlResponseBody) SetUrl(v string) *QueryFileOssUrlResponseBody {
	s.Url = &v
	return s
}

type QueryFileOssUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFileOssUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFileOssUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFileOssUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryFileOssUrlResponse) SetHeaders(v map[string]*string) *QueryFileOssUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryFileOssUrlResponse) SetStatusCode(v int32) *QueryFileOssUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFileOssUrlResponse) SetBody(v *QueryFileOssUrlResponseBody) *QueryFileOssUrlResponse {
	s.Body = v
	return s
}

type QueryHomeStatRequest struct {
	BusinessType         *int64  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryHomeStatRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHomeStatRequest) GoString() string {
	return s.String()
}

func (s *QueryHomeStatRequest) SetBusinessType(v int64) *QueryHomeStatRequest {
	s.BusinessType = &v
	return s
}

func (s *QueryHomeStatRequest) SetEndTime(v string) *QueryHomeStatRequest {
	s.EndTime = &v
	return s
}

func (s *QueryHomeStatRequest) SetOwnerId(v int64) *QueryHomeStatRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHomeStatRequest) SetResourceOwnerAccount(v string) *QueryHomeStatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryHomeStatRequest) SetResourceOwnerId(v int64) *QueryHomeStatRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryHomeStatRequest) SetStartTime(v string) *QueryHomeStatRequest {
	s.StartTime = &v
	return s
}

type QueryHomeStatResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model          *QueryHomeStatResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryHomeStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHomeStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHomeStatResponseBody) SetCode(v string) *QueryHomeStatResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHomeStatResponseBody) SetHttpStatusCode(v int64) *QueryHomeStatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryHomeStatResponseBody) SetMessage(v string) *QueryHomeStatResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHomeStatResponseBody) SetModel(v *QueryHomeStatResponseBodyModel) *QueryHomeStatResponseBody {
	s.Model = v
	return s
}

func (s *QueryHomeStatResponseBody) SetRequestId(v string) *QueryHomeStatResponseBody {
	s.RequestId = &v
	return s
}

type QueryHomeStatResponseBodyModel struct {
	CdrDrUrl *string                               `json:"CdrDrUrl,omitempty" xml:"CdrDrUrl,omitempty"`
	List     []*QueryHomeStatResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryHomeStatResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryHomeStatResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryHomeStatResponseBodyModel) SetCdrDrUrl(v string) *QueryHomeStatResponseBodyModel {
	s.CdrDrUrl = &v
	return s
}

func (s *QueryHomeStatResponseBodyModel) SetList(v []*QueryHomeStatResponseBodyModelList) *QueryHomeStatResponseBodyModel {
	s.List = v
	return s
}

type QueryHomeStatResponseBodyModelList struct {
	CallFailed  *int64  `json:"CallFailed,omitempty" xml:"CallFailed,omitempty"`
	CallSuccess *int64  `json:"CallSuccess,omitempty" xml:"CallSuccess,omitempty"`
	CallTotal   *int64  `json:"CallTotal,omitempty" xml:"CallTotal,omitempty"`
	Date        *string `json:"Date,omitempty" xml:"Date,omitempty"`
	SuccessRate *int64  `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
}

func (s QueryHomeStatResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryHomeStatResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *QueryHomeStatResponseBodyModelList) SetCallFailed(v int64) *QueryHomeStatResponseBodyModelList {
	s.CallFailed = &v
	return s
}

func (s *QueryHomeStatResponseBodyModelList) SetCallSuccess(v int64) *QueryHomeStatResponseBodyModelList {
	s.CallSuccess = &v
	return s
}

func (s *QueryHomeStatResponseBodyModelList) SetCallTotal(v int64) *QueryHomeStatResponseBodyModelList {
	s.CallTotal = &v
	return s
}

func (s *QueryHomeStatResponseBodyModelList) SetDate(v string) *QueryHomeStatResponseBodyModelList {
	s.Date = &v
	return s
}

func (s *QueryHomeStatResponseBodyModelList) SetSuccessRate(v int64) *QueryHomeStatResponseBodyModelList {
	s.SuccessRate = &v
	return s
}

type QueryHomeStatResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHomeStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHomeStatResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHomeStatResponse) GoString() string {
	return s.String()
}

func (s *QueryHomeStatResponse) SetHeaders(v map[string]*string) *QueryHomeStatResponse {
	s.Headers = v
	return s
}

func (s *QueryHomeStatResponse) SetStatusCode(v int32) *QueryHomeStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHomeStatResponse) SetBody(v *QueryHomeStatResponseBody) *QueryHomeStatResponse {
	s.Body = v
	return s
}

type QueryRecordingEnableRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s QueryRecordingEnableRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordingEnableRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordingEnableRequest) SetOwnerId(v int64) *QueryRecordingEnableRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRecordingEnableRequest) SetResourceOwnerAccount(v string) *QueryRecordingEnableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type QueryRecordingEnableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Enable    *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRecordingEnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordingEnableResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordingEnableResponseBody) SetCode(v string) *QueryRecordingEnableResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordingEnableResponseBody) SetEnable(v bool) *QueryRecordingEnableResponseBody {
	s.Enable = &v
	return s
}

func (s *QueryRecordingEnableResponseBody) SetMessage(v string) *QueryRecordingEnableResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRecordingEnableResponseBody) SetRequestId(v string) *QueryRecordingEnableResponseBody {
	s.RequestId = &v
	return s
}

type QueryRecordingEnableResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRecordingEnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRecordingEnableResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordingEnableResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordingEnableResponse) SetHeaders(v map[string]*string) *QueryRecordingEnableResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordingEnableResponse) SetStatusCode(v int32) *QueryRecordingEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordingEnableResponse) SetBody(v *QueryRecordingEnableResponseBody) *QueryRecordingEnableResponse {
	s.Body = v
	return s
}

type QueryServiceEnableRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryServiceEnableRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceEnableRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceEnableRequest) SetOwnerId(v int64) *QueryServiceEnableRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryServiceEnableRequest) SetResourceOwnerAccount(v string) *QueryServiceEnableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryServiceEnableRequest) SetResourceOwnerId(v int64) *QueryServiceEnableRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryServiceEnableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Enable    *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryServiceEnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceEnableResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceEnableResponseBody) SetCode(v string) *QueryServiceEnableResponseBody {
	s.Code = &v
	return s
}

func (s *QueryServiceEnableResponseBody) SetEnable(v bool) *QueryServiceEnableResponseBody {
	s.Enable = &v
	return s
}

func (s *QueryServiceEnableResponseBody) SetMessage(v string) *QueryServiceEnableResponseBody {
	s.Message = &v
	return s
}

func (s *QueryServiceEnableResponseBody) SetRequestId(v string) *QueryServiceEnableResponseBody {
	s.RequestId = &v
	return s
}

type QueryServiceEnableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryServiceEnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceEnableResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceEnableResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceEnableResponse) SetHeaders(v map[string]*string) *QueryServiceEnableResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceEnableResponse) SetStatusCode(v int32) *QueryServiceEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServiceEnableResponse) SetBody(v *QueryServiceEnableResponseBody) *QueryServiceEnableResponse {
	s.Body = v
	return s
}

type RecordingEnableRequest struct {
	Enable               *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RecordingEnableRequest) String() string {
	return tea.Prettify(s)
}

func (s RecordingEnableRequest) GoString() string {
	return s.String()
}

func (s *RecordingEnableRequest) SetEnable(v bool) *RecordingEnableRequest {
	s.Enable = &v
	return s
}

func (s *RecordingEnableRequest) SetOwnerId(v int64) *RecordingEnableRequest {
	s.OwnerId = &v
	return s
}

func (s *RecordingEnableRequest) SetResourceOwnerAccount(v string) *RecordingEnableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecordingEnableRequest) SetResourceOwnerId(v int64) *RecordingEnableRequest {
	s.ResourceOwnerId = &v
	return s
}

type RecordingEnableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecordingEnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecordingEnableResponseBody) GoString() string {
	return s.String()
}

func (s *RecordingEnableResponseBody) SetCode(v string) *RecordingEnableResponseBody {
	s.Code = &v
	return s
}

func (s *RecordingEnableResponseBody) SetMessage(v string) *RecordingEnableResponseBody {
	s.Message = &v
	return s
}

func (s *RecordingEnableResponseBody) SetRequestId(v string) *RecordingEnableResponseBody {
	s.RequestId = &v
	return s
}

type RecordingEnableResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecordingEnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecordingEnableResponse) String() string {
	return tea.Prettify(s)
}

func (s RecordingEnableResponse) GoString() string {
	return s.String()
}

func (s *RecordingEnableResponse) SetHeaders(v map[string]*string) *RecordingEnableResponse {
	s.Headers = v
	return s
}

func (s *RecordingEnableResponse) SetStatusCode(v int32) *RecordingEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordingEnableResponse) SetBody(v *RecordingEnableResponseBody) *RecordingEnableResponse {
	s.Body = v
	return s
}

type SetReceiptUrlRequest struct {
	CdrDrUrl             *string `json:"CdrDrUrl,omitempty" xml:"CdrDrUrl,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetReceiptUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s SetReceiptUrlRequest) GoString() string {
	return s.String()
}

func (s *SetReceiptUrlRequest) SetCdrDrUrl(v string) *SetReceiptUrlRequest {
	s.CdrDrUrl = &v
	return s
}

func (s *SetReceiptUrlRequest) SetOwnerId(v int64) *SetReceiptUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *SetReceiptUrlRequest) SetResourceOwnerAccount(v string) *SetReceiptUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetReceiptUrlRequest) SetResourceOwnerId(v int64) *SetReceiptUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetReceiptUrlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetReceiptUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetReceiptUrlResponseBody) GoString() string {
	return s.String()
}

func (s *SetReceiptUrlResponseBody) SetCode(v string) *SetReceiptUrlResponseBody {
	s.Code = &v
	return s
}

func (s *SetReceiptUrlResponseBody) SetMessage(v string) *SetReceiptUrlResponseBody {
	s.Message = &v
	return s
}

func (s *SetReceiptUrlResponseBody) SetRequestId(v string) *SetReceiptUrlResponseBody {
	s.RequestId = &v
	return s
}

type SetReceiptUrlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetReceiptUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetReceiptUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s SetReceiptUrlResponse) GoString() string {
	return s.String()
}

func (s *SetReceiptUrlResponse) SetHeaders(v map[string]*string) *SetReceiptUrlResponse {
	s.Headers = v
	return s
}

func (s *SetReceiptUrlResponse) SetStatusCode(v int32) *SetReceiptUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *SetReceiptUrlResponse) SetBody(v *SetReceiptUrlResponseBody) *SetReceiptUrlResponse {
	s.Body = v
	return s
}

type SipTrunkDetailRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SipTrunkDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s SipTrunkDetailRequest) GoString() string {
	return s.String()
}

func (s *SipTrunkDetailRequest) SetCalledNumber(v string) *SipTrunkDetailRequest {
	s.CalledNumber = &v
	return s
}

func (s *SipTrunkDetailRequest) SetCallingNumber(v string) *SipTrunkDetailRequest {
	s.CallingNumber = &v
	return s
}

func (s *SipTrunkDetailRequest) SetEndTime(v string) *SipTrunkDetailRequest {
	s.EndTime = &v
	return s
}

func (s *SipTrunkDetailRequest) SetOwnerId(v int64) *SipTrunkDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *SipTrunkDetailRequest) SetPageNumber(v int64) *SipTrunkDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *SipTrunkDetailRequest) SetPageSize(v int64) *SipTrunkDetailRequest {
	s.PageSize = &v
	return s
}

func (s *SipTrunkDetailRequest) SetRegionId(v string) *SipTrunkDetailRequest {
	s.RegionId = &v
	return s
}

func (s *SipTrunkDetailRequest) SetResourceOwnerAccount(v string) *SipTrunkDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SipTrunkDetailRequest) SetResourceOwnerId(v int64) *SipTrunkDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SipTrunkDetailRequest) SetStartTime(v string) *SipTrunkDetailRequest {
	s.StartTime = &v
	return s
}

type SipTrunkDetailResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model          *SipTrunkDetailResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SipTrunkDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SipTrunkDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SipTrunkDetailResponseBody) SetCode(v string) *SipTrunkDetailResponseBody {
	s.Code = &v
	return s
}

func (s *SipTrunkDetailResponseBody) SetHttpStatusCode(v int64) *SipTrunkDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SipTrunkDetailResponseBody) SetMessage(v string) *SipTrunkDetailResponseBody {
	s.Message = &v
	return s
}

func (s *SipTrunkDetailResponseBody) SetModel(v *SipTrunkDetailResponseBodyModel) *SipTrunkDetailResponseBody {
	s.Model = v
	return s
}

func (s *SipTrunkDetailResponseBody) SetRequestId(v string) *SipTrunkDetailResponseBody {
	s.RequestId = &v
	return s
}

type SipTrunkDetailResponseBodyModel struct {
	List  []*SipTrunkDetailResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SipTrunkDetailResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s SipTrunkDetailResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SipTrunkDetailResponseBodyModel) SetList(v []*SipTrunkDetailResponseBodyModelList) *SipTrunkDetailResponseBodyModel {
	s.List = v
	return s
}

func (s *SipTrunkDetailResponseBodyModel) SetTotal(v int64) *SipTrunkDetailResponseBodyModel {
	s.Total = &v
	return s
}

type SipTrunkDetailResponseBodyModelList struct {
	AnswerTime      *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	Answered        *int64  `json:"Answered,omitempty" xml:"Answered,omitempty"`
	BusiType        *string `json:"BusiType,omitempty" xml:"BusiType,omitempty"`
	CallEndTime     *string `json:"CallEndTime,omitempty" xml:"CallEndTime,omitempty"`
	CallType        *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CalledIp        *string `json:"CalledIp,omitempty" xml:"CalledIp,omitempty"`
	CalledNum       *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	CalledNumRegion *string `json:"CalledNumRegion,omitempty" xml:"CalledNumRegion,omitempty"`
	CalledNumType   *string `json:"CalledNumType,omitempty" xml:"CalledNumType,omitempty"`
	CallerDisplay   *string `json:"CallerDisplay,omitempty" xml:"CallerDisplay,omitempty"`
	CallerIp        *string `json:"CallerIp,omitempty" xml:"CallerIp,omitempty"`
	CallerNum       *string `json:"CallerNum,omitempty" xml:"CallerNum,omitempty"`
	CallerNumRegion *string `json:"CallerNumRegion,omitempty" xml:"CallerNumRegion,omitempty"`
	CallerNumType   *string `json:"CallerNumType,omitempty" xml:"CallerNumType,omitempty"`
	Duration        *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Hangup          *int64  `json:"Hangup,omitempty" xml:"Hangup,omitempty"`
	HangupType      *int64  `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	LvnCliType      *string `json:"LvnCliType,omitempty" xml:"LvnCliType,omitempty"`
	RecordUrl       *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserUuid        *string `json:"UserUuid,omitempty" xml:"UserUuid,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s SipTrunkDetailResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s SipTrunkDetailResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *SipTrunkDetailResponseBodyModelList) SetAnswerTime(v string) *SipTrunkDetailResponseBodyModelList {
	s.AnswerTime = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetAnswered(v int64) *SipTrunkDetailResponseBodyModelList {
	s.Answered = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetBusiType(v string) *SipTrunkDetailResponseBodyModelList {
	s.BusiType = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallEndTime(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallEndTime = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallType(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallType = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCalledIp(v string) *SipTrunkDetailResponseBodyModelList {
	s.CalledIp = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCalledNum(v string) *SipTrunkDetailResponseBodyModelList {
	s.CalledNum = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCalledNumRegion(v string) *SipTrunkDetailResponseBodyModelList {
	s.CalledNumRegion = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCalledNumType(v string) *SipTrunkDetailResponseBodyModelList {
	s.CalledNumType = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallerDisplay(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallerDisplay = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallerIp(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallerIp = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallerNum(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallerNum = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallerNumRegion(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallerNumRegion = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetCallerNumType(v string) *SipTrunkDetailResponseBodyModelList {
	s.CallerNumType = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetDuration(v int64) *SipTrunkDetailResponseBodyModelList {
	s.Duration = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetHangup(v int64) *SipTrunkDetailResponseBodyModelList {
	s.Hangup = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetHangupType(v int64) *SipTrunkDetailResponseBodyModelList {
	s.HangupType = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetLvnCliType(v string) *SipTrunkDetailResponseBodyModelList {
	s.LvnCliType = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetRecordUrl(v string) *SipTrunkDetailResponseBodyModelList {
	s.RecordUrl = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetStartTime(v string) *SipTrunkDetailResponseBodyModelList {
	s.StartTime = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetUserUuid(v string) *SipTrunkDetailResponseBodyModelList {
	s.UserUuid = &v
	return s
}

func (s *SipTrunkDetailResponseBodyModelList) SetUuid(v string) *SipTrunkDetailResponseBodyModelList {
	s.Uuid = &v
	return s
}

type SipTrunkDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SipTrunkDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SipTrunkDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s SipTrunkDetailResponse) GoString() string {
	return s.String()
}

func (s *SipTrunkDetailResponse) SetHeaders(v map[string]*string) *SipTrunkDetailResponse {
	s.Headers = v
	return s
}

func (s *SipTrunkDetailResponse) SetStatusCode(v int32) *SipTrunkDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *SipTrunkDetailResponse) SetBody(v *SipTrunkDetailResponseBody) *SipTrunkDetailResponse {
	s.Body = v
	return s
}

type SubmitGroupCallRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallerIdNumber       *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	FileKey              *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileName             *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	GroupCallType        *int64  `json:"GroupCallType,omitempty" xml:"GroupCallType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SendType             *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateId           *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TimingStart          *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
}

func (s SubmitGroupCallRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitGroupCallRequest) GoString() string {
	return s.String()
}

func (s *SubmitGroupCallRequest) SetBizType(v string) *SubmitGroupCallRequest {
	s.BizType = &v
	return s
}

func (s *SubmitGroupCallRequest) SetCallerIdNumber(v string) *SubmitGroupCallRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *SubmitGroupCallRequest) SetCountryId(v string) *SubmitGroupCallRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitGroupCallRequest) SetFileKey(v string) *SubmitGroupCallRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitGroupCallRequest) SetFileName(v string) *SubmitGroupCallRequest {
	s.FileName = &v
	return s
}

func (s *SubmitGroupCallRequest) SetGroupCallType(v int64) *SubmitGroupCallRequest {
	s.GroupCallType = &v
	return s
}

func (s *SubmitGroupCallRequest) SetOwnerId(v int64) *SubmitGroupCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitGroupCallRequest) SetProductCode(v string) *SubmitGroupCallRequest {
	s.ProductCode = &v
	return s
}

func (s *SubmitGroupCallRequest) SetResourceOwnerAccount(v string) *SubmitGroupCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitGroupCallRequest) SetResourceOwnerId(v int64) *SubmitGroupCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitGroupCallRequest) SetSendType(v int64) *SubmitGroupCallRequest {
	s.SendType = &v
	return s
}

func (s *SubmitGroupCallRequest) SetTaskName(v string) *SubmitGroupCallRequest {
	s.TaskName = &v
	return s
}

func (s *SubmitGroupCallRequest) SetTemplateId(v string) *SubmitGroupCallRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitGroupCallRequest) SetTimingStart(v string) *SubmitGroupCallRequest {
	s.TimingStart = &v
	return s
}

type SubmitGroupCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitGroupCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitGroupCallResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitGroupCallResponseBody) SetCode(v string) *SubmitGroupCallResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitGroupCallResponseBody) SetRequestId(v string) *SubmitGroupCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitGroupCallResponseBody) SetTaskId(v string) *SubmitGroupCallResponseBody {
	s.TaskId = &v
	return s
}

type SubmitGroupCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitGroupCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitGroupCallResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitGroupCallResponse) GoString() string {
	return s.String()
}

func (s *SubmitGroupCallResponse) SetHeaders(v map[string]*string) *SubmitGroupCallResponse {
	s.Headers = v
	return s
}

func (s *SubmitGroupCallResponse) SetStatusCode(v int32) *SubmitGroupCallResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitGroupCallResponse) SetBody(v *SubmitGroupCallResponseBody) *SubmitGroupCallResponse {
	s.Body = v
	return s
}

type SubmitNumberRequest struct {
	ApplyNote            *string                          `json:"ApplyNote,omitempty" xml:"ApplyNote,omitempty"`
	CountryId            *string                          `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	NumberList           []*SubmitNumberRequestNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId      *string                          `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scene                *string                          `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s SubmitNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitNumberRequest) GoString() string {
	return s.String()
}

func (s *SubmitNumberRequest) SetApplyNote(v string) *SubmitNumberRequest {
	s.ApplyNote = &v
	return s
}

func (s *SubmitNumberRequest) SetCountryId(v string) *SubmitNumberRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitNumberRequest) SetNumberList(v []*SubmitNumberRequestNumberList) *SubmitNumberRequest {
	s.NumberList = v
	return s
}

func (s *SubmitNumberRequest) SetOwnerId(v int64) *SubmitNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitNumberRequest) SetQualificationId(v string) *SubmitNumberRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitNumberRequest) SetResourceOwnerAccount(v string) *SubmitNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitNumberRequest) SetResourceOwnerId(v int64) *SubmitNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitNumberRequest) SetScene(v string) *SubmitNumberRequest {
	s.Scene = &v
	return s
}

type SubmitNumberRequestNumberList struct {
	Amount              *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	InboundConcurrency  *int64 `json:"InboundConcurrency,omitempty" xml:"InboundConcurrency,omitempty"`
	OutboundConcurrency *int64 `json:"OutboundConcurrency,omitempty" xml:"OutboundConcurrency,omitempty"`
	PhoneType           *int64 `json:"PhoneType,omitempty" xml:"PhoneType,omitempty"`
	Support             *int64 `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s SubmitNumberRequestNumberList) String() string {
	return tea.Prettify(s)
}

func (s SubmitNumberRequestNumberList) GoString() string {
	return s.String()
}

func (s *SubmitNumberRequestNumberList) SetAmount(v int64) *SubmitNumberRequestNumberList {
	s.Amount = &v
	return s
}

func (s *SubmitNumberRequestNumberList) SetInboundConcurrency(v int64) *SubmitNumberRequestNumberList {
	s.InboundConcurrency = &v
	return s
}

func (s *SubmitNumberRequestNumberList) SetOutboundConcurrency(v int64) *SubmitNumberRequestNumberList {
	s.OutboundConcurrency = &v
	return s
}

func (s *SubmitNumberRequestNumberList) SetPhoneType(v int64) *SubmitNumberRequestNumberList {
	s.PhoneType = &v
	return s
}

func (s *SubmitNumberRequestNumberList) SetSupport(v int64) *SubmitNumberRequestNumberList {
	s.Support = &v
	return s
}

type SubmitNumberShrinkRequest struct {
	ApplyNote            *string `json:"ApplyNote,omitempty" xml:"ApplyNote,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	NumberListShrink     *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scene                *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s SubmitNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitNumberShrinkRequest) SetApplyNote(v string) *SubmitNumberShrinkRequest {
	s.ApplyNote = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetCountryId(v string) *SubmitNumberShrinkRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetNumberListShrink(v string) *SubmitNumberShrinkRequest {
	s.NumberListShrink = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetOwnerId(v int64) *SubmitNumberShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetQualificationId(v string) *SubmitNumberShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetResourceOwnerAccount(v string) *SubmitNumberShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetResourceOwnerId(v int64) *SubmitNumberShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitNumberShrinkRequest) SetScene(v string) *SubmitNumberShrinkRequest {
	s.Scene = &v
	return s
}

type SubmitNumberResponseBody struct {
	ApplyId   *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitNumberResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitNumberResponseBody) SetApplyId(v string) *SubmitNumberResponseBody {
	s.ApplyId = &v
	return s
}

func (s *SubmitNumberResponseBody) SetCode(v string) *SubmitNumberResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitNumberResponseBody) SetRequestId(v string) *SubmitNumberResponseBody {
	s.RequestId = &v
	return s
}

type SubmitNumberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitNumberResponse) GoString() string {
	return s.String()
}

func (s *SubmitNumberResponse) SetHeaders(v map[string]*string) *SubmitNumberResponse {
	s.Headers = v
	return s
}

func (s *SubmitNumberResponse) SetStatusCode(v int32) *SubmitNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitNumberResponse) SetBody(v *SubmitNumberResponseBody) *SubmitNumberResponse {
	s.Body = v
	return s
}

type SubmitQualificationRequest struct {
	Address                *string `json:"Address,omitempty" xml:"Address,omitempty"`
	BusinessLicenseFileKey *string `json:"BusinessLicenseFileKey,omitempty" xml:"BusinessLicenseFileKey,omitempty"`
	CompanyName            *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	ContactEmail           *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName            *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactPhone           *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	CountryId              *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	LegalId                *string `json:"LegalId,omitempty" xml:"LegalId,omitempty"`
	LegalLicenseFileKey    *string `json:"LegalLicenseFileKey,omitempty" xml:"LegalLicenseFileKey,omitempty"`
	LegalName              *string `json:"LegalName,omitempty" xml:"LegalName,omitempty"`
	NetworkAccessFileKey   *string `json:"NetworkAccessFileKey,omitempty" xml:"NetworkAccessFileKey,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UnifiedCode            *string `json:"UnifiedCode,omitempty" xml:"UnifiedCode,omitempty"`
}

func (s SubmitQualificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualificationRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualificationRequest) SetAddress(v string) *SubmitQualificationRequest {
	s.Address = &v
	return s
}

func (s *SubmitQualificationRequest) SetBusinessLicenseFileKey(v string) *SubmitQualificationRequest {
	s.BusinessLicenseFileKey = &v
	return s
}

func (s *SubmitQualificationRequest) SetCompanyName(v string) *SubmitQualificationRequest {
	s.CompanyName = &v
	return s
}

func (s *SubmitQualificationRequest) SetContactEmail(v string) *SubmitQualificationRequest {
	s.ContactEmail = &v
	return s
}

func (s *SubmitQualificationRequest) SetContactName(v string) *SubmitQualificationRequest {
	s.ContactName = &v
	return s
}

func (s *SubmitQualificationRequest) SetContactPhone(v string) *SubmitQualificationRequest {
	s.ContactPhone = &v
	return s
}

func (s *SubmitQualificationRequest) SetCountryId(v string) *SubmitQualificationRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitQualificationRequest) SetLegalId(v string) *SubmitQualificationRequest {
	s.LegalId = &v
	return s
}

func (s *SubmitQualificationRequest) SetLegalLicenseFileKey(v string) *SubmitQualificationRequest {
	s.LegalLicenseFileKey = &v
	return s
}

func (s *SubmitQualificationRequest) SetLegalName(v string) *SubmitQualificationRequest {
	s.LegalName = &v
	return s
}

func (s *SubmitQualificationRequest) SetNetworkAccessFileKey(v string) *SubmitQualificationRequest {
	s.NetworkAccessFileKey = &v
	return s
}

func (s *SubmitQualificationRequest) SetOwnerId(v int64) *SubmitQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitQualificationRequest) SetResourceOwnerAccount(v string) *SubmitQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitQualificationRequest) SetResourceOwnerId(v int64) *SubmitQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitQualificationRequest) SetUnifiedCode(v string) *SubmitQualificationRequest {
	s.UnifiedCode = &v
	return s
}

type SubmitQualificationResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitQualificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQualificationResponseBody) SetCode(v string) *SubmitQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitQualificationResponseBody) SetMessage(v string) *SubmitQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitQualificationResponseBody) SetQualificationId(v string) *SubmitQualificationResponseBody {
	s.QualificationId = &v
	return s
}

func (s *SubmitQualificationResponseBody) SetRequestId(v string) *SubmitQualificationResponseBody {
	s.RequestId = &v
	return s
}

type SubmitQualificationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitQualificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualificationResponse) GoString() string {
	return s.String()
}

func (s *SubmitQualificationResponse) SetHeaders(v map[string]*string) *SubmitQualificationResponse {
	s.Headers = v
	return s
}

func (s *SubmitQualificationResponse) SetStatusCode(v int32) *SubmitQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQualificationResponse) SetBody(v *SubmitQualificationResponseBody) *SubmitQualificationResponse {
	s.Body = v
	return s
}

type SubmitSipTrunkRequest struct {
	ApplyNote            *string `json:"ApplyNote,omitempty" xml:"ApplyNote,omitempty"`
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	InboundIpPorts       *string `json:"InboundIpPorts,omitempty" xml:"InboundIpPorts,omitempty"`
	OutboundIps          *string `json:"OutboundIps,omitempty" xml:"OutboundIps,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scene                *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s SubmitSipTrunkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSipTrunkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSipTrunkRequest) SetApplyNote(v string) *SubmitSipTrunkRequest {
	s.ApplyNote = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetCountryId(v string) *SubmitSipTrunkRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetInboundIpPorts(v string) *SubmitSipTrunkRequest {
	s.InboundIpPorts = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetOutboundIps(v string) *SubmitSipTrunkRequest {
	s.OutboundIps = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetOwnerId(v int64) *SubmitSipTrunkRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetQualificationId(v string) *SubmitSipTrunkRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetResourceOwnerAccount(v string) *SubmitSipTrunkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetResourceOwnerId(v int64) *SubmitSipTrunkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitSipTrunkRequest) SetScene(v string) *SubmitSipTrunkRequest {
	s.Scene = &v
	return s
}

type SubmitSipTrunkResponseBody struct {
	ApplyId   *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSipTrunkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSipTrunkResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSipTrunkResponseBody) SetApplyId(v string) *SubmitSipTrunkResponseBody {
	s.ApplyId = &v
	return s
}

func (s *SubmitSipTrunkResponseBody) SetCode(v string) *SubmitSipTrunkResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSipTrunkResponseBody) SetMessage(v string) *SubmitSipTrunkResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSipTrunkResponseBody) SetRequestId(v string) *SubmitSipTrunkResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSipTrunkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSipTrunkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSipTrunkResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSipTrunkResponse) GoString() string {
	return s.String()
}

func (s *SubmitSipTrunkResponse) SetHeaders(v map[string]*string) *SubmitSipTrunkResponse {
	s.Headers = v
	return s
}

func (s *SubmitSipTrunkResponse) SetStatusCode(v int32) *SubmitSipTrunkResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSipTrunkResponse) SetBody(v *SubmitSipTrunkResponseBody) *SubmitSipTrunkResponse {
	s.Body = v
	return s
}

type SubmitVoiceFileRequest struct {
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	FileKey              *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileName             *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl              *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SubmitVoiceFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitVoiceFileRequest) GoString() string {
	return s.String()
}

func (s *SubmitVoiceFileRequest) SetCountryId(v string) *SubmitVoiceFileRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetFileKey(v string) *SubmitVoiceFileRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetFileName(v string) *SubmitVoiceFileRequest {
	s.FileName = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetFileUrl(v string) *SubmitVoiceFileRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetOwnerId(v int64) *SubmitVoiceFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetQualificationId(v string) *SubmitVoiceFileRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetResourceOwnerAccount(v string) *SubmitVoiceFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitVoiceFileRequest) SetResourceOwnerId(v int64) *SubmitVoiceFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type SubmitVoiceFileResponseBody struct {
	ApplyId   *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVoiceFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitVoiceFileResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVoiceFileResponseBody) SetApplyId(v string) *SubmitVoiceFileResponseBody {
	s.ApplyId = &v
	return s
}

func (s *SubmitVoiceFileResponseBody) SetCode(v string) *SubmitVoiceFileResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVoiceFileResponseBody) SetMessage(v string) *SubmitVoiceFileResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVoiceFileResponseBody) SetRequestId(v string) *SubmitVoiceFileResponseBody {
	s.RequestId = &v
	return s
}

type SubmitVoiceFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitVoiceFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitVoiceFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitVoiceFileResponse) GoString() string {
	return s.String()
}

func (s *SubmitVoiceFileResponse) SetHeaders(v map[string]*string) *SubmitVoiceFileResponse {
	s.Headers = v
	return s
}

func (s *SubmitVoiceFileResponse) SetStatusCode(v int32) *SubmitVoiceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVoiceFileResponse) SetBody(v *SubmitVoiceFileResponseBody) *SubmitVoiceFileResponse {
	s.Body = v
	return s
}

type SubmitVoiceTtsRequest struct {
	CountryId            *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Speed                *int64  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	Status               *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId           *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateText         *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
}

func (s SubmitVoiceTtsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitVoiceTtsRequest) GoString() string {
	return s.String()
}

func (s *SubmitVoiceTtsRequest) SetCountryId(v string) *SubmitVoiceTtsRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetLanguage(v string) *SubmitVoiceTtsRequest {
	s.Language = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetOwnerId(v int64) *SubmitVoiceTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetQualificationId(v string) *SubmitVoiceTtsRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetResourceOwnerAccount(v string) *SubmitVoiceTtsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetResourceOwnerId(v int64) *SubmitVoiceTtsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetSpeed(v int64) *SubmitVoiceTtsRequest {
	s.Speed = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetStatus(v int64) *SubmitVoiceTtsRequest {
	s.Status = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetTemplateId(v string) *SubmitVoiceTtsRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetTemplateName(v string) *SubmitVoiceTtsRequest {
	s.TemplateName = &v
	return s
}

func (s *SubmitVoiceTtsRequest) SetTemplateText(v string) *SubmitVoiceTtsRequest {
	s.TemplateText = &v
	return s
}

type SubmitVoiceTtsResponseBody struct {
	ApplyId   *string `json:"ApplyId,omitempty" xml:"ApplyId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVoiceTtsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitVoiceTtsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVoiceTtsResponseBody) SetApplyId(v string) *SubmitVoiceTtsResponseBody {
	s.ApplyId = &v
	return s
}

func (s *SubmitVoiceTtsResponseBody) SetCode(v string) *SubmitVoiceTtsResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVoiceTtsResponseBody) SetMessage(v string) *SubmitVoiceTtsResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVoiceTtsResponseBody) SetRequestId(v string) *SubmitVoiceTtsResponseBody {
	s.RequestId = &v
	return s
}

type SubmitVoiceTtsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitVoiceTtsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitVoiceTtsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitVoiceTtsResponse) GoString() string {
	return s.String()
}

func (s *SubmitVoiceTtsResponse) SetHeaders(v map[string]*string) *SubmitVoiceTtsResponse {
	s.Headers = v
	return s
}

func (s *SubmitVoiceTtsResponse) SetStatusCode(v int32) *SubmitVoiceTtsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVoiceTtsResponse) SetBody(v *SubmitVoiceTtsResponseBody) *SubmitVoiceTtsResponse {
	s.Body = v
	return s
}

type UpdateNumberRequest struct {
	Note                 *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNumberRequest) GoString() string {
	return s.String()
}

func (s *UpdateNumberRequest) SetNote(v string) *UpdateNumberRequest {
	s.Note = &v
	return s
}

func (s *UpdateNumberRequest) SetNumber(v string) *UpdateNumberRequest {
	s.Number = &v
	return s
}

func (s *UpdateNumberRequest) SetOwnerId(v int64) *UpdateNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateNumberRequest) SetResourceOwnerAccount(v string) *UpdateNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateNumberRequest) SetResourceOwnerId(v int64) *UpdateNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateNumberResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNumberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNumberResponseBody) SetCode(v string) *UpdateNumberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNumberResponseBody) SetMessage(v string) *UpdateNumberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNumberResponseBody) SetRequestId(v string) *UpdateNumberResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNumberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNumberResponse) GoString() string {
	return s.String()
}

func (s *UpdateNumberResponse) SetHeaders(v map[string]*string) *UpdateNumberResponse {
	s.Headers = v
	return s
}

func (s *UpdateNumberResponse) SetStatusCode(v int32) *UpdateNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNumberResponse) SetBody(v *UpdateNumberResponseBody) *UpdateNumberResponse {
	s.Body = v
	return s
}

type Test02Request struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s Test02Request) String() string {
	return tea.Prettify(s)
}

func (s Test02Request) GoString() string {
	return s.String()
}

func (s *Test02Request) SetOwnerId(v int64) *Test02Request {
	s.OwnerId = &v
	return s
}

func (s *Test02Request) SetResourceOwnerAccount(v string) *Test02Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *Test02Request) SetResourceOwnerId(v int64) *Test02Request {
	s.ResourceOwnerId = &v
	return s
}

type Test02ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s Test02ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Test02ResponseBody) GoString() string {
	return s.String()
}

func (s *Test02ResponseBody) SetCode(v string) *Test02ResponseBody {
	s.Code = &v
	return s
}

func (s *Test02ResponseBody) SetMessage(v string) *Test02ResponseBody {
	s.Message = &v
	return s
}

func (s *Test02ResponseBody) SetRequestId(v string) *Test02ResponseBody {
	s.RequestId = &v
	return s
}

type Test02Response struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Test02ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Test02Response) String() string {
	return tea.Prettify(s)
}

func (s Test02Response) GoString() string {
	return s.String()
}

func (s *Test02Response) SetHeaders(v map[string]*string) *Test02Response {
	s.Headers = v
	return s
}

func (s *Test02Response) SetStatusCode(v int32) *Test02Response {
	s.StatusCode = &v
	return s
}

func (s *Test02Response) SetBody(v *Test02ResponseBody) *Test02Response {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dyvmsapi-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BackendCallGroupWithOptions(tmpReq *BackendCallGroupRequest, runtime *util.RuntimeOptions) (_result *BackendCallGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BackendCallGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CalledNumber)) {
		request.CalledNumberShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CalledNumber, tea.String("CalledNumber"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumberShrink)) {
		query["CalledNumber"] = request.CalledNumberShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		query["SendType"] = request.SendType
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TimingStart)) {
		query["TimingStart"] = request.TimingStart
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BackendCallGroup"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BackendCallGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BackendCallGroup(request *BackendCallGroupRequest) (_result *BackendCallGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BackendCallGroupResponse{}
	_body, _err := client.BackendCallGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BackendCallSignalWithOptions(request *BackendCallSignalRequest, runtime *util.RuntimeOptions) (_result *BackendCallSignalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BackendCallSignal"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BackendCallSignalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BackendCallSignal(request *BackendCallSignalRequest) (_result *BackendCallSignalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BackendCallSignalResponse{}
	_body, _err := client.BackendCallSignalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancleGroupCallWithOptions(request *CancleGroupCallRequest, runtime *util.RuntimeOptions) (_result *CancleGroupCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancleGroupCall"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancleGroupCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancleGroupCall(request *CancleGroupCallRequest) (_result *CancleGroupCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancleGroupCallResponse{}
	_body, _err := client.CancleGroupCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplyNumberRecordWithOptions(request *DeleteApplyNumberRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteApplyNumberRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplyNumberRecord"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplyNumberRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplyNumberRecord(request *DeleteApplyNumberRecordRequest) (_result *DeleteApplyNumberRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplyNumberRecordResponse{}
	_body, _err := client.DeleteApplyNumberRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQualificationWithOptions(request *DeleteQualificationRequest, runtime *util.RuntimeOptions) (_result *DeleteQualificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
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
		Action:      tea.String("DeleteQualification"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQualification(request *DeleteQualificationRequest) (_result *DeleteQualificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQualificationResponse{}
	_body, _err := client.DeleteQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSipTrunkApplyWithOptions(request *DeleteSipTrunkApplyRequest, runtime *util.RuntimeOptions) (_result *DeleteSipTrunkApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["ApplyId"] = request.ApplyId
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
		Action:      tea.String("DeleteSipTrunkApply"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSipTrunkApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSipTrunkApply(request *DeleteSipTrunkApplyRequest) (_result *DeleteSipTrunkApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSipTrunkApplyResponse{}
	_body, _err := client.DeleteSipTrunkApplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVoiceFileWithOptions(request *DeleteVoiceFileRequest, runtime *util.RuntimeOptions) (_result *DeleteVoiceFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
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
		Action:      tea.String("DeleteVoiceFile"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVoiceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVoiceFile(request *DeleteVoiceFileRequest) (_result *DeleteVoiceFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVoiceFileResponse{}
	_body, _err := client.DeleteVoiceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVoiceTtsWithOptions(request *DeleteVoiceTtsRequest, runtime *util.RuntimeOptions) (_result *DeleteVoiceTtsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVoiceTts"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVoiceTtsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVoiceTts(request *DeleteVoiceTtsRequest) (_result *DeleteVoiceTtsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVoiceTtsResponse{}
	_body, _err := client.DeleteVoiceTtsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadTemplateFileWithOptions(request *DownloadTemplateFileRequest, runtime *util.RuntimeOptions) (_result *DownloadTemplateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
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
		Action:      tea.String("DownloadTemplateFile"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadTemplateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadTemplateFile(request *DownloadTemplateFileRequest) (_result *DownloadTemplateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadTemplateFileResponse{}
	_body, _err := client.DownloadTemplateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIntlVoiceOpenStatusWithOptions(request *GetIntlVoiceOpenStatusRequest, runtime *util.RuntimeOptions) (_result *GetIntlVoiceOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("GetIntlVoiceOpenStatus"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIntlVoiceOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIntlVoiceOpenStatus(request *GetIntlVoiceOpenStatusRequest) (_result *GetIntlVoiceOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIntlVoiceOpenStatusResponse{}
	_body, _err := client.GetIntlVoiceOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssInfoForUploadFileWithOptions(request *GetOssInfoForUploadFileRequest, runtime *util.RuntimeOptions) (_result *GetOssInfoForUploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
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
		Action:      tea.String("GetOssInfoForUploadFile"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssInfoForUploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssInfoForUploadFile(request *GetOssInfoForUploadFileRequest) (_result *GetOssInfoForUploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssInfoForUploadFileResponse{}
	_body, _err := client.GetOssInfoForUploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HomeStartWithOptions(request *HomeStartRequest, runtime *util.RuntimeOptions) (_result *HomeStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
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
		Action:      tea.String("HomeStart"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HomeStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HomeStart(request *HomeStartRequest) (_result *HomeStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HomeStartResponse{}
	_body, _err := client.HomeStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplyNumberRecordWithOptions(request *ListApplyNumberRecordRequest, runtime *util.RuntimeOptions) (_result *ListApplyNumberRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplyNumberRecord"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplyNumberRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplyNumberRecord(request *ListApplyNumberRecordRequest) (_result *ListApplyNumberRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplyNumberRecordResponse{}
	_body, _err := client.ListApplyNumberRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNumberWithOptions(request *ListNumberRequest, runtime *util.RuntimeOptions) (_result *ListNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyId)) {
		query["ApplyId"] = request.ApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.NumberName)) {
		query["NumberName"] = request.NumberName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneType)) {
		query["PhoneType"] = request.PhoneType
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
		Action:      tea.String("ListNumber"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNumber(request *ListNumberRequest) (_result *ListNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNumberResponse{}
	_body, _err := client.ListNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQualificationWithOptions(request *ListQualificationRequest, runtime *util.RuntimeOptions) (_result *ListQualificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyName)) {
		query["CompanyName"] = request.CompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPhone)) {
		query["ContactPhone"] = request.ContactPhone
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQualification"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQualification(request *ListQualificationRequest) (_result *ListQualificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQualificationResponse{}
	_body, _err := client.ListQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListReceiptWithOptions(request *ListReceiptRequest, runtime *util.RuntimeOptions) (_result *ListReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
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
		Action:      tea.String("ListReceipt"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReceiptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListReceipt(request *ListReceiptRequest) (_result *ListReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListReceiptResponse{}
	_body, _err := client.ListReceiptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSipTrunkWithOptions(request *ListSipTrunkRequest, runtime *util.RuntimeOptions) (_result *ListSipTrunkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSipTrunk"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSipTrunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSipTrunk(request *ListSipTrunkRequest) (_result *ListSipTrunkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSipTrunkResponse{}
	_body, _err := client.ListSipTrunkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSipTrunkDetailWithOptions(request *ListSipTrunkDetailRequest, runtime *util.RuntimeOptions) (_result *ListSipTrunkDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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
		Action:      tea.String("ListSipTrunkDetail"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSipTrunkDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSipTrunkDetail(request *ListSipTrunkDetailRequest) (_result *ListSipTrunkDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSipTrunkDetailResponse{}
	_body, _err := client.ListSipTrunkDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSipTrunkStatWithOptions(request *ListSipTrunkStatRequest, runtime *util.RuntimeOptions) (_result *ListSipTrunkStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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
		Action:      tea.String("ListSipTrunkStat"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSipTrunkStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSipTrunkStat(request *ListSipTrunkStatRequest) (_result *ListSipTrunkStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSipTrunkStatResponse{}
	_body, _err := client.ListSipTrunkStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVoiceCallWithOptions(request *ListVoiceCallRequest, runtime *util.RuntimeOptions) (_result *ListVoiceCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		query["SendType"] = request.SendType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVoiceCall"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVoiceCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVoiceCall(request *ListVoiceCallRequest) (_result *ListVoiceCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVoiceCallResponse{}
	_body, _err := client.ListVoiceCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVoiceCallDetailWithOptions(request *ListVoiceCallDetailRequest, runtime *util.RuntimeOptions) (_result *ListVoiceCallDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.HangupType)) {
		query["HangupType"] = request.HangupType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVoiceCallDetail"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVoiceCallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVoiceCallDetail(request *ListVoiceCallDetailRequest) (_result *ListVoiceCallDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVoiceCallDetailResponse{}
	_body, _err := client.ListVoiceCallDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVoiceCallStatWithOptions(request *ListVoiceCallStatRequest, runtime *util.RuntimeOptions) (_result *ListVoiceCallStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVoiceCallStat"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVoiceCallStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVoiceCallStat(request *ListVoiceCallStatRequest) (_result *ListVoiceCallStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVoiceCallStatResponse{}
	_body, _err := client.ListVoiceCallStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVoiceFileWithOptions(request *ListVoiceFileRequest, runtime *util.RuntimeOptions) (_result *ListVoiceFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVoiceFile"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVoiceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVoiceFile(request *ListVoiceFileRequest) (_result *ListVoiceFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVoiceFileResponse{}
	_body, _err := client.ListVoiceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVoiceTtsWithOptions(request *ListVoiceTtsRequest, runtime *util.RuntimeOptions) (_result *ListVoiceTtsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVoiceTts"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVoiceTtsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVoiceTts(request *ListVoiceTtsRequest) (_result *ListVoiceTtsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVoiceTtsResponse{}
	_body, _err := client.ListVoiceTtsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NumberEnableWithOptions(request *NumberEnableRequest, runtime *util.RuntimeOptions) (_result *NumberEnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NumberEnable"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NumberEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NumberEnable(request *NumberEnableRequest) (_result *NumberEnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NumberEnableResponse{}
	_body, _err := client.NumberEnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenIntlVoiceServiceWithOptions(request *OpenIntlVoiceServiceRequest, runtime *util.RuntimeOptions) (_result *OpenIntlVoiceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
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
		Action:      tea.String("OpenIntlVoiceService"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenIntlVoiceServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenIntlVoiceService(request *OpenIntlVoiceServiceRequest) (_result *OpenIntlVoiceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenIntlVoiceServiceResponse{}
	_body, _err := client.OpenIntlVoiceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OswTest1WithOptions(request *OswTest1Request, runtime *util.RuntimeOptions) (_result *OswTest1Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("OswTest1"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OswTest1Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OswTest1(request *OswTest1Request) (_result *OswTest1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OswTest1Response{}
	_body, _err := client.OswTest1WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFileOssUrlWithOptions(request *QueryFileOssUrlRequest, runtime *util.RuntimeOptions) (_result *QueryFileOssUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		query["FileKey"] = request.FileKey
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
		Action:      tea.String("QueryFileOssUrl"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFileOssUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFileOssUrl(request *QueryFileOssUrlRequest) (_result *QueryFileOssUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFileOssUrlResponse{}
	_body, _err := client.QueryFileOssUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHomeStatWithOptions(request *QueryHomeStatRequest, runtime *util.RuntimeOptions) (_result *QueryHomeStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
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
		Action:      tea.String("QueryHomeStat"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHomeStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHomeStat(request *QueryHomeStatRequest) (_result *QueryHomeStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHomeStatResponse{}
	_body, _err := client.QueryHomeStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordingEnableWithOptions(request *QueryRecordingEnableRequest, runtime *util.RuntimeOptions) (_result *QueryRecordingEnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordingEnable"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordingEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordingEnable(request *QueryRecordingEnableRequest) (_result *QueryRecordingEnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordingEnableResponse{}
	_body, _err := client.QueryRecordingEnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceEnableWithOptions(request *QueryServiceEnableRequest, runtime *util.RuntimeOptions) (_result *QueryServiceEnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryServiceEnable"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServiceEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceEnable(request *QueryServiceEnableRequest) (_result *QueryServiceEnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServiceEnableResponse{}
	_body, _err := client.QueryServiceEnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecordingEnableWithOptions(request *RecordingEnableRequest, runtime *util.RuntimeOptions) (_result *RecordingEnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
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
		Action:      tea.String("RecordingEnable"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecordingEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecordingEnable(request *RecordingEnableRequest) (_result *RecordingEnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecordingEnableResponse{}
	_body, _err := client.RecordingEnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetReceiptUrlWithOptions(request *SetReceiptUrlRequest, runtime *util.RuntimeOptions) (_result *SetReceiptUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CdrDrUrl)) {
		query["CdrDrUrl"] = request.CdrDrUrl
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
		Action:      tea.String("SetReceiptUrl"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetReceiptUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetReceiptUrl(request *SetReceiptUrlRequest) (_result *SetReceiptUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetReceiptUrlResponse{}
	_body, _err := client.SetReceiptUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SipTrunkDetailWithOptions(request *SipTrunkDetailRequest, runtime *util.RuntimeOptions) (_result *SipTrunkDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Action:      tea.String("SipTrunkDetail"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SipTrunkDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SipTrunkDetail(request *SipTrunkDetailRequest) (_result *SipTrunkDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SipTrunkDetailResponse{}
	_body, _err := client.SipTrunkDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitGroupCallWithOptions(request *SubmitGroupCallRequest, runtime *util.RuntimeOptions) (_result *SubmitGroupCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		query["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupCallType)) {
		query["GroupCallType"] = request.GroupCallType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		query["SendType"] = request.SendType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TimingStart)) {
		query["TimingStart"] = request.TimingStart
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitGroupCall"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitGroupCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitGroupCall(request *SubmitGroupCallRequest) (_result *SubmitGroupCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitGroupCallResponse{}
	_body, _err := client.SubmitGroupCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitNumberWithOptions(tmpReq *SubmitNumberRequest, runtime *util.RuntimeOptions) (_result *SubmitNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NumberList)) {
		request.NumberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NumberList, tea.String("NumberList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyNote)) {
		query["ApplyNote"] = request.ApplyNote
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.NumberListShrink)) {
		query["NumberList"] = request.NumberListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitNumber"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitNumber(request *SubmitNumberRequest) (_result *SubmitNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitNumberResponse{}
	_body, _err := client.SubmitNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitQualificationWithOptions(request *SubmitQualificationRequest, runtime *util.RuntimeOptions) (_result *SubmitQualificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessLicenseFileKey)) {
		query["BusinessLicenseFileKey"] = request.BusinessLicenseFileKey
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyName)) {
		query["CompanyName"] = request.CompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactEmail)) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPhone)) {
		query["ContactPhone"] = request.ContactPhone
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalId)) {
		query["LegalId"] = request.LegalId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalLicenseFileKey)) {
		query["LegalLicenseFileKey"] = request.LegalLicenseFileKey
	}

	if !tea.BoolValue(util.IsUnset(request.LegalName)) {
		query["LegalName"] = request.LegalName
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessFileKey)) {
		query["NetworkAccessFileKey"] = request.NetworkAccessFileKey
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

	if !tea.BoolValue(util.IsUnset(request.UnifiedCode)) {
		query["UnifiedCode"] = request.UnifiedCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitQualification"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitQualification(request *SubmitQualificationRequest) (_result *SubmitQualificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitQualificationResponse{}
	_body, _err := client.SubmitQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSipTrunkWithOptions(request *SubmitSipTrunkRequest, runtime *util.RuntimeOptions) (_result *SubmitSipTrunkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyNote)) {
		query["ApplyNote"] = request.ApplyNote
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.InboundIpPorts)) {
		query["InboundIpPorts"] = request.InboundIpPorts
	}

	if !tea.BoolValue(util.IsUnset(request.OutboundIps)) {
		query["OutboundIps"] = request.OutboundIps
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSipTrunk"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSipTrunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSipTrunk(request *SubmitSipTrunkRequest) (_result *SubmitSipTrunkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSipTrunkResponse{}
	_body, _err := client.SubmitSipTrunkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitVoiceFileWithOptions(request *SubmitVoiceFileRequest, runtime *util.RuntimeOptions) (_result *SubmitVoiceFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		query["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
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
		Action:      tea.String("SubmitVoiceFile"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitVoiceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitVoiceFile(request *SubmitVoiceFileRequest) (_result *SubmitVoiceFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitVoiceFileResponse{}
	_body, _err := client.SubmitVoiceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitVoiceTtsWithOptions(request *SubmitVoiceTtsRequest, runtime *util.RuntimeOptions) (_result *SubmitVoiceTtsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateText)) {
		query["TemplateText"] = request.TemplateText
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitVoiceTts"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitVoiceTtsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitVoiceTts(request *SubmitVoiceTtsRequest) (_result *SubmitVoiceTtsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitVoiceTtsResponse{}
	_body, _err := client.SubmitVoiceTtsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNumberWithOptions(request *UpdateNumberRequest, runtime *util.RuntimeOptions) (_result *UpdateNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
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
		Action:      tea.String("UpdateNumber"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNumber(request *UpdateNumberRequest) (_result *UpdateNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNumberResponse{}
	_body, _err := client.UpdateNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Test02WithOptions(request *Test02Request, runtime *util.RuntimeOptions) (_result *Test02Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("test02"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Test02Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Test02(request *Test02Request) (_result *Test02Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Test02Response{}
	_body, _err := client.Test02WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

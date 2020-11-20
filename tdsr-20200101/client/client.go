// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SaveHotspotConfigRequest struct {
	ParamTag     *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s SaveHotspotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigRequest) SetParamTag(v string) *SaveHotspotConfigRequest {
	s.ParamTag = &v
	return s
}

func (s *SaveHotspotConfigRequest) SetPreviewToken(v string) *SaveHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

type SaveHotspotConfigResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
}

func (s SaveHotspotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponse) SetRequestId(v string) *SaveHotspotConfigResponse {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotConfigResponse) SetSuccess(v bool) *SaveHotspotConfigResponse {
	s.Success = &v
	return s
}

func (s *SaveHotspotConfigResponse) SetErrMessage(v string) *SaveHotspotConfigResponse {
	s.ErrMessage = &v
	return s
}

type GetHotspotConfigRequest struct {
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s GetHotspotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigRequest) SetPreviewToken(v string) *GetHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

type GetHotspotConfigResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s GetHotspotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponse) SetRequestId(v string) *GetHotspotConfigResponse {
	s.RequestId = &v
	return s
}

func (s *GetHotspotConfigResponse) SetSuccess(v bool) *GetHotspotConfigResponse {
	s.Success = &v
	return s
}

func (s *GetHotspotConfigResponse) SetErrMessage(v string) *GetHotspotConfigResponse {
	s.ErrMessage = &v
	return s
}

func (s *GetHotspotConfigResponse) SetData(v string) *GetHotspotConfigResponse {
	s.Data = &v
	return s
}

func (s *GetHotspotConfigResponse) SetObjectString(v string) *GetHotspotConfigResponse {
	s.ObjectString = &v
	return s
}

type ListMainScenesRequest struct {
	QueryName *string `json:"QueryName,omitempty" xml:"QueryName,omitempty"`
}

func (s ListMainScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMainScenesRequest) GoString() string {
	return s.String()
}

func (s *ListMainScenesRequest) SetQueryName(v string) *ListMainScenesRequest {
	s.QueryName = &v
	return s
}

type ListMainScenesResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s ListMainScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMainScenesResponse) GoString() string {
	return s.String()
}

func (s *ListMainScenesResponse) SetRequestId(v string) *ListMainScenesResponse {
	s.RequestId = &v
	return s
}

func (s *ListMainScenesResponse) SetSuccess(v bool) *ListMainScenesResponse {
	s.Success = &v
	return s
}

func (s *ListMainScenesResponse) SetErrMessage(v string) *ListMainScenesResponse {
	s.ErrMessage = &v
	return s
}

func (s *ListMainScenesResponse) SetData(v string) *ListMainScenesResponse {
	s.Data = &v
	return s
}

func (s *ListMainScenesResponse) SetObjectString(v string) *ListMainScenesResponse {
	s.ObjectString = &v
	return s
}

type SaveHotspotTagRequest struct {
	ParamTag     *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s SaveHotspotTagRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagRequest) SetParamTag(v string) *SaveHotspotTagRequest {
	s.ParamTag = &v
	return s
}

func (s *SaveHotspotTagRequest) SetSubSceneUuid(v string) *SaveHotspotTagRequest {
	s.SubSceneUuid = &v
	return s
}

type SaveHotspotTagResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
}

func (s SaveHotspotTagResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponse) SetRequestId(v string) *SaveHotspotTagResponse {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotTagResponse) SetSuccess(v bool) *SaveHotspotTagResponse {
	s.Success = &v
	return s
}

func (s *SaveHotspotTagResponse) SetErrMessage(v string) *SaveHotspotTagResponse {
	s.ErrMessage = &v
	return s
}

type GetSceneListRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetSceneListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneListRequest) GoString() string {
	return s.String()
}

func (s *GetSceneListRequest) SetAccountId(v string) *GetSceneListRequest {
	s.AccountId = &v
	return s
}

type GetSceneListResponse struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data       map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s GetSceneListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneListResponse) GoString() string {
	return s.String()
}

func (s *GetSceneListResponse) SetRequestId(v string) *GetSceneListResponse {
	s.RequestId = &v
	return s
}

func (s *GetSceneListResponse) SetSuccess(v bool) *GetSceneListResponse {
	s.Success = &v
	return s
}

func (s *GetSceneListResponse) SetErrMessage(v string) *GetSceneListResponse {
	s.ErrMessage = &v
	return s
}

func (s *GetSceneListResponse) SetData(v map[string]interface{}) *GetSceneListResponse {
	s.Data = v
	return s
}

type SaveFileRequest struct {
	ParamFile    *string `json:"ParamFile,omitempty" xml:"ParamFile,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s SaveFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFileRequest) GoString() string {
	return s.String()
}

func (s *SaveFileRequest) SetParamFile(v string) *SaveFileRequest {
	s.ParamFile = &v
	return s
}

func (s *SaveFileRequest) SetSubSceneUuid(v string) *SaveFileRequest {
	s.SubSceneUuid = &v
	return s
}

type SaveFileResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s SaveFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFileResponse) GoString() string {
	return s.String()
}

func (s *SaveFileResponse) SetRequestId(v string) *SaveFileResponse {
	s.RequestId = &v
	return s
}

func (s *SaveFileResponse) SetSuccess(v bool) *SaveFileResponse {
	s.Success = &v
	return s
}

func (s *SaveFileResponse) SetErrMessage(v string) *SaveFileResponse {
	s.ErrMessage = &v
	return s
}

func (s *SaveFileResponse) SetData(v string) *SaveFileResponse {
	s.Data = &v
	return s
}

func (s *SaveFileResponse) SetObjectString(v string) *SaveFileResponse {
	s.ObjectString = &v
	return s
}

type DeleteFileRequest struct {
	ParamFile    *string `json:"ParamFile,omitempty" xml:"ParamFile,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetParamFile(v string) *DeleteFileRequest {
	s.ParamFile = &v
	return s
}

func (s *DeleteFileRequest) SetSubSceneUuid(v string) *DeleteFileRequest {
	s.SubSceneUuid = &v
	return s
}

type DeleteFileResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetRequestId(v string) *DeleteFileResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponse) SetSuccess(v bool) *DeleteFileResponse {
	s.Success = &v
	return s
}

func (s *DeleteFileResponse) SetErrMessage(v string) *DeleteFileResponse {
	s.ErrMessage = &v
	return s
}

type GetHotspotTagRequest struct {
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHotspotTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotTagRequest) SetPreviewToken(v string) *GetHotspotTagRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotTagRequest) SetSubSceneUuid(v string) *GetHotspotTagRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *GetHotspotTagRequest) SetType(v string) *GetHotspotTagRequest {
	s.Type = &v
	return s
}

type GetHotspotTagResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s GetHotspotTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponse) SetRequestId(v string) *GetHotspotTagResponse {
	s.RequestId = &v
	return s
}

func (s *GetHotspotTagResponse) SetSuccess(v bool) *GetHotspotTagResponse {
	s.Success = &v
	return s
}

func (s *GetHotspotTagResponse) SetErrMessage(v string) *GetHotspotTagResponse {
	s.ErrMessage = &v
	return s
}

func (s *GetHotspotTagResponse) SetData(v string) *GetHotspotTagResponse {
	s.Data = &v
	return s
}

func (s *GetHotspotTagResponse) SetObjectString(v string) *GetHotspotTagResponse {
	s.ObjectString = &v
	return s
}

type GetPolicyRequest struct {
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) SetSubSceneUuid(v string) *GetPolicyRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *GetPolicyRequest) SetType(v string) *GetPolicyRequest {
	s.Type = &v
	return s
}

type GetPolicyResponse struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string                `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s GetPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) SetRequestId(v string) *GetPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponse) SetSuccess(v bool) *GetPolicyResponse {
	s.Success = &v
	return s
}

func (s *GetPolicyResponse) SetErrMessage(v string) *GetPolicyResponse {
	s.ErrMessage = &v
	return s
}

func (s *GetPolicyResponse) SetData(v map[string]interface{}) *GetPolicyResponse {
	s.Data = v
	return s
}

func (s *GetPolicyResponse) SetObjectString(v string) *GetPolicyResponse {
	s.ObjectString = &v
	return s
}

type PublishHotspotRequest struct {
	ParamTag     *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s PublishHotspotRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotRequest) GoString() string {
	return s.String()
}

func (s *PublishHotspotRequest) SetParamTag(v string) *PublishHotspotRequest {
	s.ParamTag = &v
	return s
}

func (s *PublishHotspotRequest) SetSubSceneUuid(v string) *PublishHotspotRequest {
	s.SubSceneUuid = &v
	return s
}

type PublishHotspotResponse struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data       map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s PublishHotspotResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotResponse) GoString() string {
	return s.String()
}

func (s *PublishHotspotResponse) SetRequestId(v string) *PublishHotspotResponse {
	s.RequestId = &v
	return s
}

func (s *PublishHotspotResponse) SetSuccess(v bool) *PublishHotspotResponse {
	s.Success = &v
	return s
}

func (s *PublishHotspotResponse) SetErrMessage(v string) *PublishHotspotResponse {
	s.ErrMessage = &v
	return s
}

func (s *PublishHotspotResponse) SetData(v map[string]interface{}) *PublishHotspotResponse {
	s.Data = v
	return s
}

type GetWindowConfigRequest struct {
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s GetWindowConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWindowConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWindowConfigRequest) SetPreviewToken(v string) *GetWindowConfigRequest {
	s.PreviewToken = &v
	return s
}

type GetWindowConfigResponse struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string                `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s GetWindowConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWindowConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWindowConfigResponse) SetRequestId(v string) *GetWindowConfigResponse {
	s.RequestId = &v
	return s
}

func (s *GetWindowConfigResponse) SetSuccess(v bool) *GetWindowConfigResponse {
	s.Success = &v
	return s
}

func (s *GetWindowConfigResponse) SetErrMessage(v string) *GetWindowConfigResponse {
	s.ErrMessage = &v
	return s
}

func (s *GetWindowConfigResponse) SetData(v map[string]interface{}) *GetWindowConfigResponse {
	s.Data = v
	return s
}

func (s *GetWindowConfigResponse) SetObjectString(v string) *GetWindowConfigResponse {
	s.ObjectString = &v
	return s
}

type GetSceneDataRequest struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetSceneDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneDataRequest) GoString() string {
	return s.String()
}

func (s *GetSceneDataRequest) SetToken(v string) *GetSceneDataRequest {
	s.Token = &v
	return s
}

type GetSceneDataResponse struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ObjectString *string                `json:"ObjectString,omitempty" xml:"ObjectString,omitempty" require:"true"`
}

func (s GetSceneDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneDataResponse) GoString() string {
	return s.String()
}

func (s *GetSceneDataResponse) SetRequestId(v string) *GetSceneDataResponse {
	s.RequestId = &v
	return s
}

func (s *GetSceneDataResponse) SetSuccess(v bool) *GetSceneDataResponse {
	s.Success = &v
	return s
}

func (s *GetSceneDataResponse) SetErrMessage(v string) *GetSceneDataResponse {
	s.ErrMessage = &v
	return s
}

func (s *GetSceneDataResponse) SetData(v map[string]interface{}) *GetSceneDataResponse {
	s.Data = v
	return s
}

func (s *GetSceneDataResponse) SetObjectString(v string) *GetSceneDataResponse {
	s.ObjectString = &v
	return s
}

type CheckPermissionRequest struct {
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s CheckPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckPermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckPermissionRequest) SetAliyunId(v string) *CheckPermissionRequest {
	s.AliyunId = &v
	return s
}

type CheckPermissionResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
}

func (s CheckPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckPermissionResponse) GoString() string {
	return s.String()
}

func (s *CheckPermissionResponse) SetRequestId(v string) *CheckPermissionResponse {
	s.RequestId = &v
	return s
}

func (s *CheckPermissionResponse) SetSuccess(v bool) *CheckPermissionResponse {
	s.Success = &v
	return s
}

func (s *CheckPermissionResponse) SetErrMessage(v string) *CheckPermissionResponse {
	s.ErrMessage = &v
	return s
}

type CheckResourceRequest struct {
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty" require:"true"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty" require:"true"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty" require:"true"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty" require:"true"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty" require:"true"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty" require:"true"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty" require:"true"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceRequest) GoString() string {
	return s.String()
}

func (s *CheckResourceRequest) SetCountry(v string) *CheckResourceRequest {
	s.Country = &v
	return s
}

func (s *CheckResourceRequest) SetInterrupt(v bool) *CheckResourceRequest {
	s.Interrupt = &v
	return s
}

func (s *CheckResourceRequest) SetInvoker(v string) *CheckResourceRequest {
	s.Invoker = &v
	return s
}

func (s *CheckResourceRequest) SetPk(v string) *CheckResourceRequest {
	s.Pk = &v
	return s
}

func (s *CheckResourceRequest) SetBid(v string) *CheckResourceRequest {
	s.Bid = &v
	return s
}

func (s *CheckResourceRequest) SetHid(v int64) *CheckResourceRequest {
	s.Hid = &v
	return s
}

func (s *CheckResourceRequest) SetTaskIdentifier(v string) *CheckResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *CheckResourceRequest) SetTaskExtraData(v string) *CheckResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *CheckResourceRequest) SetGmtWakeup(v string) *CheckResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *CheckResourceRequest) SetSuccess(v bool) *CheckResourceRequest {
	s.Success = &v
	return s
}

func (s *CheckResourceRequest) SetMessage(v string) *CheckResourceRequest {
	s.Message = &v
	return s
}

func (s *CheckResourceRequest) SetLevel(v int64) *CheckResourceRequest {
	s.Level = &v
	return s
}

func (s *CheckResourceRequest) SetUrl(v string) *CheckResourceRequest {
	s.Url = &v
	return s
}

func (s *CheckResourceRequest) SetPrompt(v string) *CheckResourceRequest {
	s.Prompt = &v
	return s
}

type CheckResourceResponse struct {
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty" require:"true"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty" require:"true"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty" require:"true"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty" require:"true"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty" require:"true"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty" require:"true"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty" require:"true"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty" require:"true"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceResponse) GoString() string {
	return s.String()
}

func (s *CheckResourceResponse) SetInterrupt(v bool) *CheckResourceResponse {
	s.Interrupt = &v
	return s
}

func (s *CheckResourceResponse) SetInvoker(v string) *CheckResourceResponse {
	s.Invoker = &v
	return s
}

func (s *CheckResourceResponse) SetPk(v string) *CheckResourceResponse {
	s.Pk = &v
	return s
}

func (s *CheckResourceResponse) SetBid(v string) *CheckResourceResponse {
	s.Bid = &v
	return s
}

func (s *CheckResourceResponse) SetHid(v int64) *CheckResourceResponse {
	s.Hid = &v
	return s
}

func (s *CheckResourceResponse) SetCountry(v string) *CheckResourceResponse {
	s.Country = &v
	return s
}

func (s *CheckResourceResponse) SetTaskIdentifier(v string) *CheckResourceResponse {
	s.TaskIdentifier = &v
	return s
}

func (s *CheckResourceResponse) SetTaskExtraData(v string) *CheckResourceResponse {
	s.TaskExtraData = &v
	return s
}

func (s *CheckResourceResponse) SetGmtWakeup(v string) *CheckResourceResponse {
	s.GmtWakeup = &v
	return s
}

func (s *CheckResourceResponse) SetSuccess(v bool) *CheckResourceResponse {
	s.Success = &v
	return s
}

func (s *CheckResourceResponse) SetMessage(v string) *CheckResourceResponse {
	s.Message = &v
	return s
}

func (s *CheckResourceResponse) SetLevel(v int64) *CheckResourceResponse {
	s.Level = &v
	return s
}

func (s *CheckResourceResponse) SetUrl(v string) *CheckResourceResponse {
	s.Url = &v
	return s
}

func (s *CheckResourceResponse) SetPrompt(v string) *CheckResourceResponse {
	s.Prompt = &v
	return s
}

func (s *CheckResourceResponse) SetRequestId(v string) *CheckResourceResponse {
	s.RequestId = &v
	return s
}

type CreateSceneRequest struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneRequest) SetName(v string) *CreateSceneRequest {
	s.Name = &v
	return s
}

func (s *CreateSceneRequest) SetProjectId(v string) *CreateSceneRequest {
	s.ProjectId = &v
	return s
}

type CreateSceneResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SceneId      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty" require:"true"`
}

func (s CreateSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneResponse) SetRequestId(v string) *CreateSceneResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSceneResponse) SetSceneId(v int64) *CreateSceneResponse {
	s.SceneId = &v
	return s
}

func (s *CreateSceneResponse) SetSuccess(v bool) *CreateSceneResponse {
	s.Success = &v
	return s
}

func (s *CreateSceneResponse) SetErrMessage(v string) *CreateSceneResponse {
	s.ErrMessage = &v
	return s
}

func (s *CreateSceneResponse) SetPreviewToken(v string) *CreateSceneResponse {
	s.PreviewToken = &v
	return s
}

type CreateProjectRequest struct {
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BusinessId         *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessUserIdList *string `json:"BusinessUserIdList,omitempty" xml:"BusinessUserIdList,omitempty"`
	GatherUserIdList   *string `json:"GatherUserIdList,omitempty" xml:"GatherUserIdList,omitempty"`
	BuilderUserIdList  *string `json:"BuilderUserIdList,omitempty" xml:"BuilderUserIdList,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetBusinessId(v string) *CreateProjectRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateProjectRequest) SetBusinessUserIdList(v string) *CreateProjectRequest {
	s.BusinessUserIdList = &v
	return s
}

func (s *CreateProjectRequest) SetGatherUserIdList(v string) *CreateProjectRequest {
	s.GatherUserIdList = &v
	return s
}

func (s *CreateProjectRequest) SetBuilderUserIdList(v string) *CreateProjectRequest {
	s.BuilderUserIdList = &v
	return s
}

type CreateProjectResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetRequestId(v string) *CreateProjectResponse {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponse) SetId(v int64) *CreateProjectResponse {
	s.Id = &v
	return s
}

func (s *CreateProjectResponse) SetName(v string) *CreateProjectResponse {
	s.Name = &v
	return s
}

func (s *CreateProjectResponse) SetSuccess(v bool) *CreateProjectResponse {
	s.Success = &v
	return s
}

func (s *CreateProjectResponse) SetErrMessage(v string) *CreateProjectResponse {
	s.ErrMessage = &v
	return s
}

type DeleteProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetRequestId(v string) *DeleteProjectResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponse) SetSuccess(v bool) *DeleteProjectResponse {
	s.Success = &v
	return s
}

func (s *DeleteProjectResponse) SetErrMessage(v string) *DeleteProjectResponse {
	s.ErrMessage = &v
	return s
}

type ListScenesRequest struct {
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IsPublishQuery *bool   `json:"IsPublishQuery,omitempty" xml:"IsPublishQuery,omitempty"`
}

func (s ListScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenesRequest) GoString() string {
	return s.String()
}

func (s *ListScenesRequest) SetProjectId(v string) *ListScenesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListScenesRequest) SetIsPublishQuery(v bool) *ListScenesRequest {
	s.IsPublishQuery = &v
	return s
}

type ListScenesResponse struct {
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool                     `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrMessage *string                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty" require:"true"`
	Data       []*ListScenesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponse) GoString() string {
	return s.String()
}

func (s *ListScenesResponse) SetRequestId(v string) *ListScenesResponse {
	s.RequestId = &v
	return s
}

func (s *ListScenesResponse) SetSuccess(v bool) *ListScenesResponse {
	s.Success = &v
	return s
}

func (s *ListScenesResponse) SetErrMessage(v string) *ListScenesResponse {
	s.ErrMessage = &v
	return s
}

func (s *ListScenesResponse) SetData(v []*ListScenesResponseData) *ListScenesResponse {
	s.Data = v
	return s
}

type ListScenesResponseData struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
}

func (s ListScenesResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseData) GoString() string {
	return s.String()
}

func (s *ListScenesResponseData) SetSceneId(v string) *ListScenesResponseData {
	s.SceneId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": tea.String("lyj.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tdsr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) SaveHotspotConfigWithOptions(request *SaveHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveHotspotConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SaveHotspotConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveHotspotConfig(request *SaveHotspotConfigRequest) (_result *SaveHotspotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveHotspotConfigResponse{}
	_body, _err := client.SaveHotspotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotspotConfigWithOptions(request *GetHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *GetHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetHotspotConfigResponse{}
	_body, _err := client.DoRequest(tea.String("GetHotspotConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotspotConfig(request *GetHotspotConfigRequest) (_result *GetHotspotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotspotConfigResponse{}
	_body, _err := client.GetHotspotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMainScenesWithOptions(request *ListMainScenesRequest, runtime *util.RuntimeOptions) (_result *ListMainScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListMainScenesResponse{}
	_body, _err := client.DoRequest(tea.String("ListMainScenes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMainScenes(request *ListMainScenesRequest) (_result *ListMainScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMainScenesResponse{}
	_body, _err := client.ListMainScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveHotspotTagWithOptions(request *SaveHotspotTagRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveHotspotTagResponse{}
	_body, _err := client.DoRequest(tea.String("SaveHotspotTag"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveHotspotTag(request *SaveHotspotTagRequest) (_result *SaveHotspotTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveHotspotTagResponse{}
	_body, _err := client.SaveHotspotTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSceneListWithOptions(request *GetSceneListRequest, runtime *util.RuntimeOptions) (_result *GetSceneListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSceneListResponse{}
	_body, _err := client.DoRequest(tea.String("GetSceneList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSceneList(request *GetSceneListRequest) (_result *GetSceneListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSceneListResponse{}
	_body, _err := client.GetSceneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveFileWithOptions(request *SaveFileRequest, runtime *util.RuntimeOptions) (_result *SaveFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveFileResponse{}
	_body, _err := client.DoRequest(tea.String("SaveFile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveFile(request *SaveFileRequest) (_result *SaveFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveFileResponse{}
	_body, _err := client.SaveFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteFile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotspotTagWithOptions(request *GetHotspotTagRequest, runtime *util.RuntimeOptions) (_result *GetHotspotTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetHotspotTagResponse{}
	_body, _err := client.DoRequest(tea.String("GetHotspotTag"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotspotTag(request *GetHotspotTagRequest) (_result *GetHotspotTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotspotTagResponse{}
	_body, _err := client.GetHotspotTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("GetPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishHotspotWithOptions(request *PublishHotspotRequest, runtime *util.RuntimeOptions) (_result *PublishHotspotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PublishHotspotResponse{}
	_body, _err := client.DoRequest(tea.String("PublishHotspot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishHotspot(request *PublishHotspotRequest) (_result *PublishHotspotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishHotspotResponse{}
	_body, _err := client.PublishHotspotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWindowConfigWithOptions(request *GetWindowConfigRequest, runtime *util.RuntimeOptions) (_result *GetWindowConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetWindowConfigResponse{}
	_body, _err := client.DoRequest(tea.String("GetWindowConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWindowConfig(request *GetWindowConfigRequest) (_result *GetWindowConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWindowConfigResponse{}
	_body, _err := client.GetWindowConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSceneDataWithOptions(request *GetSceneDataRequest, runtime *util.RuntimeOptions) (_result *GetSceneDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSceneDataResponse{}
	_body, _err := client.DoRequest(tea.String("GetSceneData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSceneData(request *GetSceneDataRequest) (_result *GetSceneDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSceneDataResponse{}
	_body, _err := client.GetSceneDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckPermissionWithOptions(request *CheckPermissionRequest, runtime *util.RuntimeOptions) (_result *CheckPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckPermissionResponse{}
	_body, _err := client.DoRequest(tea.String("CheckPermission"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckPermission(request *CheckPermissionRequest) (_result *CheckPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckPermissionResponse{}
	_body, _err := client.CheckPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckResourceWithOptions(request *CheckResourceRequest, runtime *util.RuntimeOptions) (_result *CheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckResourceResponse{}
	_body, _err := client.DoRequest(tea.String("CheckResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckResource(request *CheckResourceRequest) (_result *CheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResourceResponse{}
	_body, _err := client.CheckResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSceneWithOptions(request *CreateSceneRequest, runtime *util.RuntimeOptions) (_result *CreateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSceneResponse{}
	_body, _err := client.DoRequest(tea.String("CreateScene"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScene(request *CreateSceneRequest) (_result *CreateSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSceneResponse{}
	_body, _err := client.CreateSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRequest(tea.String("CreateProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScenesWithOptions(request *ListScenesRequest, runtime *util.RuntimeOptions) (_result *ListScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListScenesResponse{}
	_body, _err := client.DoRequest(tea.String("ListScenes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-01-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScenes(request *ListScenesRequest) (_result *ListScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScenesResponse{}
	_body, _err := client.ListScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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

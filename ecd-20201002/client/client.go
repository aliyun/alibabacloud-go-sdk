// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OssUploadCredential struct {
	AccessKeyId  *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Endpoint     *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FilePath     *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	OssPolicy    *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	StsToken     *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s OssUploadCredential) String() string {
	return tea.Prettify(s)
}

func (s OssUploadCredential) GoString() string {
	return s.String()
}

func (s *OssUploadCredential) SetAccessKeyId(v string) *OssUploadCredential {
	s.AccessKeyId = &v
	return s
}

func (s *OssUploadCredential) SetEndpoint(v string) *OssUploadCredential {
	s.Endpoint = &v
	return s
}

func (s *OssUploadCredential) SetFilePath(v string) *OssUploadCredential {
	s.FilePath = &v
	return s
}

func (s *OssUploadCredential) SetOssPolicy(v string) *OssUploadCredential {
	s.OssPolicy = &v
	return s
}

func (s *OssUploadCredential) SetOssSignature(v string) *OssUploadCredential {
	s.OssSignature = &v
	return s
}

func (s *OssUploadCredential) SetStsToken(v string) *OssUploadCredential {
	s.StsToken = &v
	return s
}

type ApproveFotaUpdateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1-D-20220513.14****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d4452bd7-88df-4b90-a409-806da674****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// ecd-138dsptkrt00u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18390c954ce59e2915ef16663205371721e0db9a46179ac89249a203320459523cb54ad08efe324784dd0eba25950****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 4771b873-c757-4893-973c-7f5beejh****
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TargetStatus *string `json:"TargetStatus,omitempty" xml:"TargetStatus,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApproveFotaUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateRequest) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateRequest) SetAppVersion(v string) *ApproveFotaUpdateRequest {
	s.AppVersion = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetClientId(v string) *ApproveFotaUpdateRequest {
	s.ClientId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetDesktopId(v string) *ApproveFotaUpdateRequest {
	s.DesktopId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetLoginToken(v string) *ApproveFotaUpdateRequest {
	s.LoginToken = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetRegionId(v string) *ApproveFotaUpdateRequest {
	s.RegionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetSessionId(v string) *ApproveFotaUpdateRequest {
	s.SessionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetTargetStatus(v string) *ApproveFotaUpdateRequest {
	s.TargetStatus = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetUuid(v string) *ApproveFotaUpdateRequest {
	s.Uuid = &v
	return s
}

type ApproveFotaUpdateResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveFotaUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponseBody) SetRequestId(v string) *ApproveFotaUpdateResponseBody {
	s.RequestId = &v
	return s
}

type ApproveFotaUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveFotaUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveFotaUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateResponse) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponse) SetHeaders(v map[string]*string) *ApproveFotaUpdateResponse {
	s.Headers = v
	return s
}

func (s *ApproveFotaUpdateResponse) SetStatusCode(v int32) *ApproveFotaUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveFotaUpdateResponse) SetBody(v *ApproveFotaUpdateResponseBody) *ApproveFotaUpdateResponse {
	s.Body = v
	return s
}

type ChangePasswordRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 42f6645a-9c3c-4772-be2a-cc5f5732****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The new password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67436290
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+dir-227468****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The current password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 1
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ChangePasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) SetClientId(v string) *ChangePasswordRequest {
	s.ClientId = &v
	return s
}

func (s *ChangePasswordRequest) SetEndUserId(v string) *ChangePasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ChangePasswordRequest) SetLoginToken(v string) *ChangePasswordRequest {
	s.LoginToken = &v
	return s
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetOfficeSiteId(v string) *ChangePasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetRegionId(v string) *ChangePasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ChangePasswordRequest) SetSessionId(v string) *ChangePasswordRequest {
	s.SessionId = &v
	return s
}

type ChangePasswordResponseBody struct {
	// The logon token.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 484256DA-D816-44D2-9D86-B6EE4D5B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangePasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponseBody) SetLoginToken(v string) *ChangePasswordResponseBody {
	s.LoginToken = &v
	return s
}

func (s *ChangePasswordResponseBody) SetRequestId(v string) *ChangePasswordResponseBody {
	s.RequestId = &v
	return s
}

type ChangePasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangePasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponse) SetHeaders(v map[string]*string) *ChangePasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangePasswordResponse) SetStatusCode(v int32) *ChangePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangePasswordResponse) SetBody(v *ChangePasswordResponseBody) *ChangePasswordResponse {
	s.Body = v
	return s
}

type DeleteFingerPrintTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 59e86b39-ccac-4dfa-93d7-1f724052****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 40401e62-5caf-4508-8de7-bf98af12****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1c0436c721786529914f16516396228454fa6284c9b80f9917f25ebbec2aa30c10343e3f6f9aff64500ce13808aef****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6df06330-3b75-4768-b334-41a73a64****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteFingerPrintTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFingerPrintTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateRequest) SetClientId(v string) *DeleteFingerPrintTemplateRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetClientToken(v string) *DeleteFingerPrintTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetIndex(v int32) *DeleteFingerPrintTemplateRequest {
	s.Index = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetLoginToken(v string) *DeleteFingerPrintTemplateRequest {
	s.LoginToken = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetRegionId(v string) *DeleteFingerPrintTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetSessionId(v string) *DeleteFingerPrintTemplateRequest {
	s.SessionId = &v
	return s
}

type DeleteFingerPrintTemplateResponseBody struct {
	// example:
	//
	// 134BD0B2-B848-5743-9CE2-C1FD3D5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFingerPrintTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFingerPrintTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateResponseBody) SetRequestId(v string) *DeleteFingerPrintTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFingerPrintTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFingerPrintTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFingerPrintTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFingerPrintTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateResponse) SetHeaders(v map[string]*string) *DeleteFingerPrintTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) SetStatusCode(v int32) *DeleteFingerPrintTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) SetBody(v *DeleteFingerPrintTemplateResponseBody) *DeleteFingerPrintTemplateResponse {
	s.Body = v
	return s
}

type DescribeDirectoriesRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54c17e1d-2d72-4b87-aa33-25f3b3f2****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The directory IDs.
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetClientId(v string) *DescribeDirectoriesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

type DescribeDirectoriesResponseBody struct {
	// The directories.
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F369A091-002F-49C8-AD55-02A77629****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	// The connection method.
	//
	// Valid values:
	//
	// 	- VPC: End users connect to cloud computers over an enterprise virtual private cloud (VPC).
	//
	// 	- INTERNET: End users connect to cloud computers over the Internet.
	//
	// 	- ANY: End users connect to cloud computers over the Internet or an enterprise VPC.
	//
	// example:
	//
	// INTERNET
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The directory type.
	//
	// example:
	//
	// AD_CONNECTOR
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The provider ID.
	//
	// example:
	//
	// 26842
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// The URL of the SSO service.
	//
	// example:
	//
	// https://eds-cn-shanghai-67726****
	SsoServiceUrl *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetProviderId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.ProviderId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSsoServiceUrl(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.SsoServiceUrl = &v
	return s
}

type DescribeDirectoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetStatusCode(v int32) *DescribeDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

type DescribeFingerPrintTemplatesRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61e39dc6-0450-45f6-a372-2a09e938****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v189646d6f329e4dfcbf51653542202890570fec26e4f9ee26427c5920fcd93871f017d2190199c4c7d0c0bf00f573****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a5062d68-e550-4d09-8288-67c8ba9e****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeFingerPrintTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesRequest) SetClientId(v string) *DescribeFingerPrintTemplatesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetLoginToken(v string) *DescribeFingerPrintTemplatesRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetRegionId(v string) *DescribeFingerPrintTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetSessionId(v string) *DescribeFingerPrintTemplatesRequest {
	s.SessionId = &v
	return s
}

type DescribeFingerPrintTemplatesResponseBody struct {
	// The fingerprint templates.
	FingerPrintTemplates []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates `json:"FingerPrintTemplates,omitempty" xml:"FingerPrintTemplates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9C1D3FBE-84E1-5ABB-AD98-2003AC71****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponseBody) SetFingerPrintTemplates(v []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) *DescribeFingerPrintTemplatesResponseBody {
	s.FingerPrintTemplates = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBody) SetRequestId(v string) *DescribeFingerPrintTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// example:
	//
	// 2258a3d5-b8f8-4d79-a221-eaecf211****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2022-03-13T13:26:29Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Finger 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user ID.
	//
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The index.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The logon time.
	//
	// example:
	//
	// 2022-03-13T13:26:29Z
	LoginTime *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-074949****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetClientId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.ClientId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetCreationTime(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.CreationTime = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetDescription(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.Description = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetEndUserId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.EndUserId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetIndex(v int64) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.Index = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetLoginTime(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.LoginTime = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetOfficeSiteId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.OfficeSiteId = &v
	return s
}

type DescribeFingerPrintTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFingerPrintTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponse) SetHeaders(v map[string]*string) *DescribeFingerPrintTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) SetStatusCode(v int32) *DescribeFingerPrintTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) SetBody(v *DescribeFingerPrintTemplatesResponseBody) *DescribeFingerPrintTemplatesResponse {
	s.Body = v
	return s
}

type DescribeGlobalDesktopsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c213150d-7ac3-432c-b749-6e1e090b****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// INTERNET
	DesktopAccessType *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DesktopId         []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// example:
	//
	// DesktopTest
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-880841****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// 关键字。支持模糊搜索桌面ID、云桌面名称和终端用户自定义的桌面名称。
	//
	// example:
	//
	// ecd
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v147c9114a180489f89691663893169****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eyJkZWZhdWx0IjpbIjk2MjEy****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-880841****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// AssignTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// true
	QueryFotaUpdate *bool `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// example:
	//
	// 5c456a41-1e65-4e72-ab4d-5dcfff52****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// example:
	//
	// true
	WithoutLatency *bool `json:"WithoutLatency,omitempty" xml:"WithoutLatency,omitempty"`
}

func (s DescribeGlobalDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsRequest) SetClientId(v string) *DescribeGlobalDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopAccessType(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopId(v []*string) *DescribeGlobalDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopName(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopStatus(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDirectoryId(v string) *DescribeGlobalDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetKeyword(v string) *DescribeGlobalDesktopsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLanguage(v string) *DescribeGlobalDesktopsRequest {
	s.Language = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLoginRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.LoginRegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLoginToken(v string) *DescribeGlobalDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetMaxResults(v int32) *DescribeGlobalDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetNextToken(v string) *DescribeGlobalDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetOfficeSiteId(v string) *DescribeGlobalDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetOrderBy(v string) *DescribeGlobalDesktopsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetQueryFotaUpdate(v bool) *DescribeGlobalDesktopsRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSearchRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSessionId(v string) *DescribeGlobalDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSortType(v string) *DescribeGlobalDesktopsRequest {
	s.SortType = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetWithoutLatency(v bool) *DescribeGlobalDesktopsRequest {
	s.WithoutLatency = &v
	return s
}

type DescribeGlobalDesktopsResponseBody struct {
	Desktops []*DescribeGlobalDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	// example:
	//
	// eyJkZWZhdWx0IjpbIjIwMjItMDgtMTdUM****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 4686A731-D601-548C-83E2-4CB6371E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBody) SetDesktops(v []*DescribeGlobalDesktopsResponseBodyDesktops) *DescribeGlobalDesktopsResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) SetNextToken(v string) *DescribeGlobalDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) SetRequestId(v string) *DescribeGlobalDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktops struct {
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 支持的客户端信息
	Clients []*DescribeGlobalDesktopsResponseBodyDesktopsClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 2020-11-06T08:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// dg-3uiojcc0j4kh7****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Running
	DesktopStatus *string                                                    `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopTimers []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// example:
	//
	// ecd.basic.large
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string                                            `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Disks       []*DescribeGlobalDesktopsResponseBodyDesktopsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// example:
	//
	// User1
	EndUserId  *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-31T15:59Z
	ExpiredTime *string                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FotaUpdate  *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	// example:
	//
	// 2048
	GpuMemory       *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	HibernationBeta *bool  `json:"HibernationBeta,omitempty" xml:"HibernationBeta,omitempty"`
	// example:
	//
	// testName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 2021-07-13T15:59Z
	LastStartTime   *string   `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	LocalName       *string   `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ManagementFlags []*string `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	// example:
	//
	// 4096
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 192.168.xx.xx
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId  *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Os            *string `json:"Os,omitempty" xml:"Os,omitempty"`
	OsDescription *string `json:"OsDescription,omitempty" xml:"OsDescription,omitempty"`
	// example:
	//
	// Windows
	OsType   *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// pg-9cktlowtxfl6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// ecd-gx2x1dhsm****
	RealDesktopId *string `json:"RealDesktopId,omitempty" xml:"RealDesktopId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId       *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionLocation *string                                               `json:"RegionLocation,omitempty" xml:"RegionLocation,omitempty"`
	SessionType    *string                                               `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	Sessions       []*DescribeGlobalDesktopsResponseBodyDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SupportHibernation *bool `json:"SupportHibernation,omitempty" xml:"SupportHibernation,omitempty"`
	// example:
	//
	// testDesktop
	UserCustomName *string `json:"UserCustomName,omitempty" xml:"UserCustomName,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetChargeType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetClients(v []*DescribeGlobalDesktopsResponseBodyDesktopsClients) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Clients = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopTimers(v []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopTimers = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDisks(v []*DescribeGlobalDesktopsResponseBodyDesktopsDisks) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetEndUserId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetFotaUpdate(v *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.FotaUpdate = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetGpuMemory(v int32) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.GpuMemory = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetHibernationBeta(v bool) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.HibernationBeta = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetHostName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.HostName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetLastStartTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.LastStartTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetLocalName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.LocalName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetManagementFlags(v []*string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetNetworkInterfaceIp(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOs(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Os = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOsDescription(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OsDescription = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOsType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetPlatform(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Platform = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetProtocolType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRealDesktopId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RealDesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRegionId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRegionLocation(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RegionLocation = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSessionType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.SessionType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSessions(v []*DescribeGlobalDesktopsResponseBodyDesktopsSessions) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSupportHibernation(v bool) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.SupportHibernation = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetUserCustomName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.UserCustomName = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsClients struct {
	// 客户端类型，取值：
	//
	// - macos：Mac客户端
	//
	// - ios：IOS客户端
	//
	// - android：Android客户端
	//
	// - html5：Web客户端
	//
	// - windows：Windows客户端
	//
	// - linux：Linux客户端
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// 客户端状态，取值：
	//
	// - ON：允许登录
	//
	// - OFF：不允许登录
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsClients) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) SetClientType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsClients {
	s.ClientType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) SetStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktopsClients {
	s.Status = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers struct {
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// example:
	//
	// 0 0 0 ? 	- 1
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// false
	Enforce       *bool   `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// 60
	Interval      *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// RESET_TYPE_BOTH
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// example:
	//
	// NoConnectShutdown
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetAllowClientSetting(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetCronExpression(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetEnforce(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetExecutionTime(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetInterval(v int32) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.Interval = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetOperationType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetResetType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetTimerType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.TimerType = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsDisks struct {
	// example:
	//
	// d-jedbpr4sl9l37****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// example:
	//
	// 80
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate struct {
	// example:
	//
	// Enterprise
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 0.0.0-D-20220102.xxxx
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// 0.0.0-R-20220307.xxxx
	NewAppVersion *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	NewDcdVersion *string `json:"NewDcdVersion,omitempty" xml:"NewDcdVersion,omitempty"`
	// example:
	//
	// testProject
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ReleaseNote   *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	ReleaseNoteJp *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	// example:
	//
	// 100
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetChannel(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Channel = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetCurrentAppVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetForce(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Force = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetNewAppVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetNewDcdVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewDcdVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetProject(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Project = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNote(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteEn(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteJp(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteJp = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetSize(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Size = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsSessions struct {
	// example:
	//
	// User1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// 2021-03-07T08:23Z
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) SetEndUserId(v string) *DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) SetEstablishmentTime(v string) *DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

type DescribeGlobalDesktopsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponse) SetHeaders(v map[string]*string) *DescribeGlobalDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDesktopsResponse) SetStatusCode(v int32) *DescribeGlobalDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDesktopsResponse) SetBody(v *DescribeGlobalDesktopsResponseBody) *DescribeGlobalDesktopsResponse {
	s.Body = v
	return s
}

type DescribeOfficeSitesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 42f6645a-9c3c-4772-be2a-cc5f5732****
	ClientId     *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) SetClientId(v string) *DescribeOfficeSitesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

type DescribeOfficeSitesResponseBody struct {
	OfficeSites []*DescribeOfficeSitesResponseBodyOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOfficeSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBody) SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody {
	s.OfficeSites = v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetRequestId(v string) *DescribeOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSites struct {
	// example:
	//
	// VPC
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// example:
	//
	// http://ep-bp1s2vmbj55r5rzc****.epsrv-bp1pcfhpwvlpny01****.cn-hangzhou.privatelink.aliyuncs.com
	DesktopVpcEndpoint *string `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// example:
	//
	// 268****
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// https://eds-cn-shanghai-67****
	SsoServiceUrl *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetProviderId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ProviderId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoServiceUrl(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoServiceUrl = &v
	return s
}

type DescribeOfficeSitesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponse) SetHeaders(v map[string]*string) *DescribeOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfficeSitesResponse) SetStatusCode(v int32) *DescribeOfficeSitesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOfficeSitesResponse) SetBody(v *DescribeOfficeSitesResponseBody) *DescribeOfficeSitesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac4a73ad-789a-449a-a88f-d18571d6****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetClientId(v string) *DescribeRegionsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The region endpoint.
	//
	// example:
	//
	// ecd.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeSnapshotsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 16dad2b6-3c6d-4e4c-b057-78ecb13c****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v16abfb945208fc5745061668654680853da4a25202d1a394fcad57bba484e9827ad43ea7d10fb6bf13d44a4adc0e9****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// a99b9aca-bac5-4da2-819e-400ce98f****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// s-2ze81owrnv9pity4****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetClientId(v string) *DescribeSnapshotsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetLoginToken(v string) *DescribeSnapshotsRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSessionId(v string) *DescribeSnapshotsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

type DescribeSnapshotsResponseBody struct {
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) SetNextToken(v string) *DescribeSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseBodySnapshots struct {
	// example:
	//
	// 2020-12-20T14:52:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ecd-g03l3tlm8djoj****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 30
	RemainTime       *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	RestorePointId   *string `json:"RestorePointId,omitempty" xml:"RestorePointId,omitempty"`
	RestorePointName *string `json:"RestorePointName,omitempty" xml:"RestorePointName,omitempty"`
	// example:
	//
	// s-2zeipxmnhej803x7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// example:
	//
	// USER
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// example:
	//
	// 150
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	// example:
	//
	// SYSTEM
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// example:
	//
	// ACCOMPLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRestorePointId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.RestorePointId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRestorePointName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.RestorePointName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetStatusCode(v int32) *DescribeSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeUserResourcesRequest struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// false
	AutoRefresh *bool `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty"`
	// example:
	//
	// 0
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 54c17e1d-2d72-4b87-aa33-25f3b3f2****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 7.6.0-R-20241112.222305
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// false
	DualCenterForward *bool `json:"DualCenterForward,omitempty" xml:"DualCenterForward,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteIds []*string `json:"OfficeSiteIds,omitempty" xml:"OfficeSiteIds,omitempty" type:"Repeated"`
	// example:
	//
	// AssignTime
	OrderBy      *string   `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	ProductTypes []*string `json:"ProductTypes,omitempty" xml:"ProductTypes,omitempty" type:"Repeated"`
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// false
	QueryFotaUpdate *bool `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	// example:
	//
	// false
	RefreshFotaUpdate *bool     `json:"RefreshFotaUpdate,omitempty" xml:"RefreshFotaUpdate,omitempty"`
	ResourceIds       []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// testName
	ResourceName  *string   `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// desktop
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s DescribeUserResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesRequest) SetAccessType(v string) *DescribeUserResourcesRequest {
	s.AccessType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetAutoRefresh(v bool) *DescribeUserResourcesRequest {
	s.AutoRefresh = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetCategoryId(v int32) *DescribeUserResourcesRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetCategoryType(v int32) *DescribeUserResourcesRequest {
	s.CategoryType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetClientId(v string) *DescribeUserResourcesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetClientType(v string) *DescribeUserResourcesRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetClientVersion(v string) *DescribeUserResourcesRequest {
	s.ClientVersion = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetDualCenterForward(v bool) *DescribeUserResourcesRequest {
	s.DualCenterForward = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetLanguage(v string) *DescribeUserResourcesRequest {
	s.Language = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetLoginRegionId(v string) *DescribeUserResourcesRequest {
	s.LoginRegionId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetLoginToken(v string) *DescribeUserResourcesRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetMaxResults(v int32) *DescribeUserResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetNextToken(v string) *DescribeUserResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetOfficeSiteIds(v []*string) *DescribeUserResourcesRequest {
	s.OfficeSiteIds = v
	return s
}

func (s *DescribeUserResourcesRequest) SetOrderBy(v string) *DescribeUserResourcesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetProductTypes(v []*string) *DescribeUserResourcesRequest {
	s.ProductTypes = v
	return s
}

func (s *DescribeUserResourcesRequest) SetProtocolType(v string) *DescribeUserResourcesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetQueryFotaUpdate(v bool) *DescribeUserResourcesRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetRefreshFotaUpdate(v bool) *DescribeUserResourcesRequest {
	s.RefreshFotaUpdate = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetResourceIds(v []*string) *DescribeUserResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeUserResourcesRequest) SetResourceName(v string) *DescribeUserResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetResourceTypes(v []*string) *DescribeUserResourcesRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribeUserResourcesRequest) SetScene(v string) *DescribeUserResourcesRequest {
	s.Scene = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetSearchRegionId(v string) *DescribeUserResourcesRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetSessionId(v string) *DescribeUserResourcesRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetSortType(v string) *DescribeUserResourcesRequest {
	s.SortType = &v
	return s
}

type DescribeUserResourcesResponseBody struct {
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryFailedResourceTypes []*string `json:"QueryFailedResourceTypes,omitempty" xml:"QueryFailedResourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1732869815062
	RankVersion *int64 `json:"RankVersion,omitempty" xml:"RankVersion,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*DescribeUserResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeUserResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBody) SetNextToken(v string) *DescribeUserResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetQueryFailedResourceTypes(v []*string) *DescribeUserResourcesResponseBody {
	s.QueryFailedResourceTypes = v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetRankVersion(v int64) *DescribeUserResourcesResponseBody {
	s.RankVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetRequestId(v string) *DescribeUserResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetResources(v []*DescribeUserResourcesResponseBodyResources) *DescribeUserResourcesResponseBody {
	s.Resources = v
	return s
}

type DescribeUserResourcesResponseBodyResources struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 194101959****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// app-0001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// example:
	//
	// 0
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// cn-shanghai+cds-695277****
	CdsName *string `json:"CdsName,omitempty" xml:"CdsName,omitempty"`
	// example:
	//
	// ecds-0****
	CenterResourceId *string `json:"CenterResourceId,omitempty" xml:"CenterResourceId,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string                                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Clients    []*DescribeUserResourcesResponseBodyResourcesClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// example:
	//
	// {"authMode":"App"}
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// 2024-12-11T07:12:12Z
	CreateTime          *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DesktopDurationList []*DescribeUserResourcesResponseBodyResourcesDesktopDurationList `json:"DesktopDurationList,omitempty" xml:"DesktopDurationList,omitempty" type:"Repeated"`
	DesktopTimers       []*DescribeUserResourcesResponseBodyResourcesDesktopTimers       `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-02-22T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// stg114510
	ExternalDomainId *string `json:"ExternalDomainId,omitempty" xml:"ExternalDomainId,omitempty"`
	// example:
	//
	// test001
	ExternalUserId *string                                               `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	FotaUpdate     *DescribeUserResourcesResponseBodyResourcesFotaUpdate `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	// example:
	//
	// true
	GlobalStatus *bool `json:"GlobalStatus,omitempty" xml:"GlobalStatus,omitempty"`
	HasUpgrade   *bool `json:"HasUpgrade,omitempty" xml:"HasUpgrade,omitempty"`
	// example:
	//
	// false
	HibernationBeta *bool `json:"HibernationBeta,omitempty" xml:"HibernationBeta,omitempty"`
	// example:
	//
	// http://example.com/icon.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 2025-01-24T03:12:04Z
	LastStartTime      *string   `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	LocalName          *string   `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ManagementStatuses []*string `json:"ManagementStatuses,omitempty" xml:"ManagementStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// cn-shanghai+dir-3367****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// Normal
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// Windows Server 2022
	Os            *string `json:"Os,omitempty" xml:"Os,omitempty"`
	OsDescription *string `json:"OsDescription,omitempty" xml:"OsDescription,omitempty"`
	// example:
	//
	// Windows
	OsType   *string                                             `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OsUpdate *DescribeUserResourcesResponseBodyResourcesOsUpdate `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty" type:"Struct"`
	// example:
	//
	// AndroidCloud
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// ecd-0001
	RealDesktopId *string `json:"RealDesktopId,omitempty" xml:"RealDesktopId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Mainland
	RegionLocation *string `json:"RegionLocation,omitempty" xml:"RegionLocation,omitempty"`
	// example:
	//
	// dg-0****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ecd-d19tya8zi4****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Center
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// example:
	//
	// testName01
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// Connected
	ResourceSessionStatus *string `json:"ResourceSessionStatus,omitempty" xml:"ResourceSessionStatus,omitempty"`
	// example:
	//
	// Running
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// SINGLE_SESSION
	SessionType *string                                               `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	Sessions    []*DescribeUserResourcesResponseBodyResourcesSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// example:
	//
	// PrePaid
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// example:
	//
	// true
	SupportHibernation *bool     `json:"SupportHibernation,omitempty" xml:"SupportHibernation,omitempty"`
	SupportedActions   []*string `json:"SupportedActions,omitempty" xml:"SupportedActions,omitempty" type:"Repeated"`
	// example:
	//
	// #FFFFFF
	ThemeColor     *string `json:"ThemeColor,omitempty" xml:"ThemeColor,omitempty"`
	UserCustomName *string `json:"UserCustomName,omitempty" xml:"UserCustomName,omitempty"`
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResources) SetAccessType(v string) *DescribeUserResourcesResponseBodyResources {
	s.AccessType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetAliUid(v int64) *DescribeUserResourcesResponseBodyResources {
	s.AliUid = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetAppId(v string) *DescribeUserResourcesResponseBodyResources {
	s.AppId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetAuthMode(v string) *DescribeUserResourcesResponseBodyResources {
	s.AuthMode = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCategoryId(v int32) *DescribeUserResourcesResponseBodyResources {
	s.CategoryId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCategoryType(v int32) *DescribeUserResourcesResponseBodyResources {
	s.CategoryType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCdsName(v string) *DescribeUserResourcesResponseBodyResources {
	s.CdsName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCenterResourceId(v string) *DescribeUserResourcesResponseBodyResources {
	s.CenterResourceId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetChargeType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ChargeType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetClients(v []*DescribeUserResourcesResponseBodyResourcesClients) *DescribeUserResourcesResponseBodyResources {
	s.Clients = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetConnectionProperties(v string) *DescribeUserResourcesResponseBodyResources {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCreateTime(v string) *DescribeUserResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetDesktopDurationList(v []*DescribeUserResourcesResponseBodyResourcesDesktopDurationList) *DescribeUserResourcesResponseBodyResources {
	s.DesktopDurationList = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetDesktopTimers(v []*DescribeUserResourcesResponseBodyResourcesDesktopTimers) *DescribeUserResourcesResponseBodyResources {
	s.DesktopTimers = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetExpiredTime(v string) *DescribeUserResourcesResponseBodyResources {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetExternalDomainId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ExternalDomainId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetExternalUserId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ExternalUserId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetFotaUpdate(v *DescribeUserResourcesResponseBodyResourcesFotaUpdate) *DescribeUserResourcesResponseBodyResources {
	s.FotaUpdate = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetGlobalStatus(v bool) *DescribeUserResourcesResponseBodyResources {
	s.GlobalStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetHasUpgrade(v bool) *DescribeUserResourcesResponseBodyResources {
	s.HasUpgrade = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetHibernationBeta(v bool) *DescribeUserResourcesResponseBodyResources {
	s.HibernationBeta = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetIcon(v string) *DescribeUserResourcesResponseBodyResources {
	s.Icon = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetLastStartTime(v string) *DescribeUserResourcesResponseBodyResources {
	s.LastStartTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetLocalName(v string) *DescribeUserResourcesResponseBodyResources {
	s.LocalName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetManagementStatuses(v []*string) *DescribeUserResourcesResponseBodyResources {
	s.ManagementStatuses = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOfficeSiteId(v string) *DescribeUserResourcesResponseBodyResources {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOrderStatus(v string) *DescribeUserResourcesResponseBodyResources {
	s.OrderStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOs(v string) *DescribeUserResourcesResponseBodyResources {
	s.Os = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOsDescription(v string) *DescribeUserResourcesResponseBodyResources {
	s.OsDescription = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOsType(v string) *DescribeUserResourcesResponseBodyResources {
	s.OsType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOsUpdate(v *DescribeUserResourcesResponseBodyResourcesOsUpdate) *DescribeUserResourcesResponseBodyResources {
	s.OsUpdate = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetProductType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetProtocolType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ProtocolType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetRealDesktopId(v string) *DescribeUserResourcesResponseBodyResources {
	s.RealDesktopId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetRegionId(v string) *DescribeUserResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetRegionLocation(v string) *DescribeUserResourcesResponseBodyResources {
	s.RegionLocation = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceGroupId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceLevel(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceLevel = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceName(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceSessionStatus(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceSessionStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceStatus(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSessionType(v string) *DescribeUserResourcesResponseBodyResources {
	s.SessionType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSessions(v []*DescribeUserResourcesResponseBodyResourcesSessions) *DescribeUserResourcesResponseBodyResources {
	s.Sessions = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSubPayType(v string) *DescribeUserResourcesResponseBodyResources {
	s.SubPayType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSupportHibernation(v bool) *DescribeUserResourcesResponseBodyResources {
	s.SupportHibernation = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSupportedActions(v []*string) *DescribeUserResourcesResponseBodyResources {
	s.SupportedActions = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetThemeColor(v string) *DescribeUserResourcesResponseBodyResources {
	s.ThemeColor = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetUserCustomName(v string) *DescribeUserResourcesResponseBodyResources {
	s.UserCustomName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetVersion(v string) *DescribeUserResourcesResponseBodyResources {
	s.Version = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesClients struct {
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesClients) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) SetClientType(v string) *DescribeUserResourcesResponseBodyResourcesClients {
	s.ClientType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) SetStatus(v string) *DescribeUserResourcesResponseBodyResourcesClients {
	s.Status = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesDesktopDurationList struct {
	// example:
	//
	// mdp-0bxls4qpi6bl6****
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// example:
	//
	// 2025-01-17T07:01Z
	PackageCreationTime *string `json:"PackageCreationTime,omitempty" xml:"PackageCreationTime,omitempty"`
	// example:
	//
	// 2025-02-17T15:59Z
	PackageExpiredTime *string `json:"PackageExpiredTime,omitempty" xml:"PackageExpiredTime,omitempty"`
	// example:
	//
	// mdp-0bxls4qpi6bl6****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// Available
	PackageStatus *string `json:"PackageStatus,omitempty" xml:"PackageStatus,omitempty"`
	// example:
	//
	// NORMAL_PACKAGE
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// Postpaid
	PackageUsedUpStrategy *string `json:"PackageUsedUpStrategy,omitempty" xml:"PackageUsedUpStrategy,omitempty"`
	// example:
	//
	// 2025-02-17T15:59Z
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	// example:
	//
	// 2025-01-17T07:01Z
	PeriodStartTime *string `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	// example:
	//
	// 199
	PostPaidLimitFee *float32 `json:"PostPaidLimitFee,omitempty" xml:"PostPaidLimitFee,omitempty"`
	// example:
	//
	// 432000
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// example:
	//
	// 16850
	UsedDuration *int64 `json:"UsedDuration,omitempty" xml:"UsedDuration,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopDurationList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetOrderInstanceId(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageCreationTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageCreationTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageExpiredTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageExpiredTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageId(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageStatus(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageUsedUpStrategy(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageUsedUpStrategy = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPeriodEndTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PeriodEndTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPeriodStartTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PeriodStartTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPostPaidLimitFee(v float32) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PostPaidLimitFee = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetTotalDuration(v int64) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.TotalDuration = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetUsedDuration(v int64) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.UsedDuration = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesDesktopTimers struct {
	// example:
	//
	// false
	AllowClientSetting *string `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// example:
	//
	// 0 30 13 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// false
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// example:
	//
	// 2025-01-21T11:37Z
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// Hibernate
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// example:
	//
	// TimerBoot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopTimers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopTimers) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetAllowClientSetting(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetCronExpression(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetEnforce(v bool) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetExecutionTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetInterval(v int32) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.Interval = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetOperationType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetResetType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetTimerType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.TimerType = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesFotaUpdate struct {
	// example:
	//
	// aliyun
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 2.7.0-R-20250122.154826
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// 2.7.0-R-20250125.154826
	NewAppVersion *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	// example:
	//
	// 2.6.9-R-20250123.153415
	NewDcdVersion *string `json:"NewDcdVersion,omitempty" xml:"NewDcdVersion,omitempty"`
	// example:
	//
	// wuying-asp_single_session_desktop_win_x64
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// example:
	//
	// up
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// example:
	//
	// up
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	// example:
	//
	// up
	ReleaseNoteJp *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	// example:
	//
	// 474981930
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesFotaUpdate) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetChannel(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Channel = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetCurrentAppVersion(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetForce(v bool) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Force = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetNewAppVersion(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetNewDcdVersion(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.NewDcdVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetProject(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Project = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetReleaseNote(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetReleaseNoteEn(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetReleaseNoteJp(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.ReleaseNoteJp = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetSize(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Size = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesOsUpdate struct {
	CheckId          *string                                                       `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	KbListString     *string                                                       `json:"KbListString,omitempty" xml:"KbListString,omitempty"`
	PackageCount     *int32                                                        `json:"PackageCount,omitempty" xml:"PackageCount,omitempty"`
	Packages         []*DescribeUserResourcesResponseBodyResourcesOsUpdatePackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	UpdateCatalogUrl *string                                                       `json:"UpdateCatalogUrl,omitempty" xml:"UpdateCatalogUrl,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdate) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdate) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetCheckId(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.CheckId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetKbListString(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.KbListString = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetPackageCount(v int32) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.PackageCount = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetPackages(v []*DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.Packages = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetUpdateCatalogUrl(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.UpdateCatalogUrl = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesOsUpdatePackages struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Kb          *string `json:"Kb,omitempty" xml:"Kb,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) SetDescription(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	s.Description = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) SetKb(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	s.Kb = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) SetTitle(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	s.Title = &v
	return s
}

type DescribeUserResourcesResponseBodyResourcesSessions struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 2025-01-22T11:03:36Z
	ResourceSessionStartTime *string `json:"ResourceSessionStartTime,omitempty" xml:"ResourceSessionStartTime,omitempty"`
	// example:
	//
	// user001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// benchmark_test@test.shenzhen
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesSessions) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetNickName(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.NickName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetResourceSessionStartTime(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.ResourceSessionStartTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetUserId(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.UserId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetUserPrincipalName(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.UserPrincipalName = &v
	return s
}

type DescribeUserResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponse) SetHeaders(v map[string]*string) *DescribeUserResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserResourcesResponse) SetStatusCode(v int32) *DescribeUserResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserResourcesResponse) SetBody(v *DescribeUserResourcesResponseBody) *DescribeUserResourcesResponse {
	s.Body = v
	return s
}

type EncryptPasswordRequest struct {
	// The ID of the client. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1d40776f-e9cb-4e2b-a8da-308d10e8****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// cn-beijing+dir-131196****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1b16dcff3ab21a6c5ec01652238375511cff5a1db59fd4dc49afb37e2ea7a626af6f38109fd0498b6abd9de1af7743****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-beijing+dir-131196****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The password that you want to encrypt.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ab123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c78e2e52-23d9-4401-a648-e67ac6ff****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EncryptPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptPasswordRequest) GoString() string {
	return s.String()
}

func (s *EncryptPasswordRequest) SetClientId(v string) *EncryptPasswordRequest {
	s.ClientId = &v
	return s
}

func (s *EncryptPasswordRequest) SetDirectoryId(v string) *EncryptPasswordRequest {
	s.DirectoryId = &v
	return s
}

func (s *EncryptPasswordRequest) SetLoginToken(v string) *EncryptPasswordRequest {
	s.LoginToken = &v
	return s
}

func (s *EncryptPasswordRequest) SetOfficeSiteId(v string) *EncryptPasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *EncryptPasswordRequest) SetPassword(v string) *EncryptPasswordRequest {
	s.Password = &v
	return s
}

func (s *EncryptPasswordRequest) SetRegionId(v string) *EncryptPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *EncryptPasswordRequest) SetSessionId(v string) *EncryptPasswordRequest {
	s.SessionId = &v
	return s
}

type EncryptPasswordResponseBody struct {
	// The encrypted password.
	//
	// example:
	//
	// d34601bc-e6b1-4433-b0cc-8f6c5e52;n4apvGub3OBoj4Grwg==;thhO4UEomJfdvwnwlA==
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AF538DA8-FFC6-52DA-8FF8-7B92579F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncryptPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *EncryptPasswordResponseBody) SetEncryptedPassword(v string) *EncryptPasswordResponseBody {
	s.EncryptedPassword = &v
	return s
}

func (s *EncryptPasswordResponseBody) SetRequestId(v string) *EncryptPasswordResponseBody {
	s.RequestId = &v
	return s
}

type EncryptPasswordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EncryptPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptPasswordResponse) GoString() string {
	return s.String()
}

func (s *EncryptPasswordResponse) SetHeaders(v map[string]*string) *EncryptPasswordResponse {
	s.Headers = v
	return s
}

func (s *EncryptPasswordResponse) SetStatusCode(v int32) *EncryptPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *EncryptPasswordResponse) SetBody(v *EncryptPasswordResponseBody) *EncryptPasswordResponse {
	s.Body = v
	return s
}

type GetCloudDriveServiceMountTokenRequest struct {
	// example:
	//
	// 00e122c3-13fb-4fc3-bc7a-5d9acb89****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// v1972cd3446f0e523598916520951742474e6624fcdea6652994d47bc6157d27f7cc900c339db67882j3no4nh5bk3b4****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-7186763****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 14e1fe41-ce9b-491d-aa8c-345jk2n4bk****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetCloudDriveServiceMountTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenRequest) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenRequest) SetClientId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetLoginToken(v string) *GetCloudDriveServiceMountTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetOfficeSiteId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetRegionId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetSessionId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.SessionId = &v
	return s
}

type GetCloudDriveServiceMountTokenResponseBody struct {
	// example:
	//
	// DC27288A-F9E1-5092-9B5B-71C27D15****
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *GetCloudDriveServiceMountTokenResponseBodyToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s GetCloudDriveServiceMountTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponseBody) SetRequestId(v string) *GetCloudDriveServiceMountTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBody) SetToken(v *GetCloudDriveServiceMountTokenResponseBodyToken) *GetCloudDriveServiceMountTokenResponseBody {
	s.Token = v
	return s
}

type GetCloudDriveServiceMountTokenResponseBodyToken struct {
	// example:
	//
	// h****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// 2022-10-10T04:41:35Z
	ExpiredAfter *string `json:"ExpiredAfter,omitempty" xml:"ExpiredAfter,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 7836fa6eced7dc8d54c775k34iu3h4i2kh534f****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 6050416754750
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// example:
	//
	// 605089
	UsedSize *int64 `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
}

func (s GetCloudDriveServiceMountTokenResponseBodyToken) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponseBodyToken) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetDomainId(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.DomainId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetExpiredAfter(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.ExpiredAfter = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetStatus(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.Status = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetToken(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.Token = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetTotalSize(v int64) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.TotalSize = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetUsedSize(v int64) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.UsedSize = &v
	return s
}

type GetCloudDriveServiceMountTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudDriveServiceMountTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudDriveServiceMountTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponse) SetHeaders(v map[string]*string) *GetCloudDriveServiceMountTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) SetStatusCode(v int32) *GetCloudDriveServiceMountTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) SetBody(v *GetCloudDriveServiceMountTokenResponseBody) *GetCloudDriveServiceMountTokenResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion  *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string                          `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Tag       []*GetConnectionTicketRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c638****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uuid   *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetClientId(v string) *GetConnectionTicketRequest {
	s.ClientId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientType(v string) *GetConnectionTicketRequest {
	s.ClientType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetCommandContent(v string) *GetConnectionTicketRequest {
	s.CommandContent = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginToken(v string) *GetConnectionTicketRequest {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerAccount(v string) *GetConnectionTicketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerId(v int64) *GetConnectionTicketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetSessionId(v string) *GetConnectionTicketRequest {
	s.SessionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTag(v []*GetConnectionTicketRequestTag) *GetConnectionTicketRequest {
	s.Tag = v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

type GetConnectionTicketRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionTicketRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequestTag) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequestTag) SetKey(v string) *GetConnectionTicketRequestTag {
	s.Key = &v
	return s
}

func (s *GetConnectionTicketRequestTag) SetValue(v string) *GetConnectionTicketRequestTag {
	s.Value = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskCode  *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c63862da
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// example:
	//
	// FINISHED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// W0VuY29kaW5nXQ0KSW5wdXRFbmNvZGluZz1V********
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskCode(v string) *GetConnectionTicketResponseBody {
	s.TaskCode = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskMessage(v string) *GetConnectionTicketResponseBody {
	s.TaskMessage = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetStatusCode(v int32) *GetConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetLoginTokenRequest struct {
	// The verification code that is generated by the virtual MFA device. This parameter is required if you set `CurrentStage` to `MFAVerify`.
	//
	// example:
	//
	// 47****
	AuthenticationCode *string `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	// The ID of the Alibaba Cloud Workspace client. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that runs an Alibaba Cloud Workspace client.
	//
	// example:
	//
	// Windows_NT 10.0.1***	- x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The type of Alibaba Cloud Workspace clients.
	//
	// Valid values:
	//
	// 	- HTML5: web client.
	//
	// 	- WINDOWS: Windows client.
	//
	// 	- MACOS: macOS client.
	//
	// 	- IOS: iOS client.
	//
	// 	- ANDROID: Android client.
	//
	// example:
	//
	// Windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The version of the client. When you use an Alibaba Cloud Workspace client, you can view the client version in the **About*	- dialog box on the client logon page.
	//
	// example:
	//
	// 2.1.0-R-20210731.1****
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The logon authentication stage. Valid values:
	//
	// 	- `ADPassword`: the stage to verify the identity of the Active Directory (AD) user. You must specify the value when the system verifies the identity of a convenience account or an AD account.
	//
	// 	- `MFABind: the stage to bind a virtual multi-factor authentication (MFA) device.`
	//
	// 	- `MFAVerify: the stage to verify the verification code that is generated by the virtual MFA device.`
	//
	// 	- `TokenVerify`: the stage to perform two-factor authentication on an Alibaba Cloud Workspace client (hereinafter referred to as the client).
	//
	// 	- `ChangePassword`: the stage to change the password of the user.
	//
	// 	- `KeepAliveVerify`: the stage to obtain LoginToken if KeepAliveToken is valid.
	//
	// example:
	//
	// ADPassword
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The office network ID. This parameter has the same meaning as `OfficeSiteId`. We recommend that you replace `DirectoryId` with `OfficeSiteId`. You can specify a value for `DirectoryId` or `OfficeSiteId`.
	//
	// example:
	//
	// cn-hangzhou+dir-885351****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the convenience user or the AD user. This parameter is required if you set `CurrentStage` to `ADPassword`.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Specifies whether to keep the user logged on to the client.
	//
	// Valid values:
	//
	// 	- null: Default value. Do not keep the user logged on to the client.
	//
	// 	- true: Keep the user logged on to the client.
	//
	// 	- false:  Do not keep the user logged on to the client.
	//
	// example:
	//
	// false
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The token to keep logging on to an Alibaba Cloud Workspace client. When an end user logs on to the Alibaba Cloud Workspace client and select Auto Sign-in, `KeepAliveToken` is returned after you call this operation. Within the valid period of the returned token``, you can call the `GetLoginToken` operation and set `CurrentStage` to `KeepAliveVerify`. Then, you can obtain LoginToken. If you set `CurrentStage` to `KeepAliveVerify`, `KeepAliveToken` is required.
	//
	// example:
	//
	// hide
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// The new password. This parameter is required if you set `CurrentStage` to `ChangePassword`.
	//
	// example:
	//
	// NewPassword
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The current password. This parameter is required if you set `CurrentStage` to `ChangePassword`.
	//
	// example:
	//
	// OldPassword
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	// The password of the convenience user or the AD user. This parameter is required if you set `CurrentStage` to `ADPassword`.
	//
	// example:
	//
	// Password1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by EDS.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the session.
	//
	// 	- If the virtual multi-factor authentication (MFA) device is not bound or two-factor authentication is not enabled for the client, you do not need to specify a value for `SessionId`.
	//
	// 	- If the virtual MFA device is not bound or two-factor authentication is enabled for the client, you must specify a value for `SessionId` to verify the user identity after you specify a value for `ADPassword`. The value of the `SessionId` parameter is returned only if the CurrentStage parameter is set to `ADPassword` when you call the `GetLoginToken` operation.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// If two-factor authentication is enabled for Alibaba Cloud Workspace terminals in the EDS console and the system detects that the current logon user is exposed to risks, the system sends a verification code to the email address of the user. This parameter is required if you set `CurrentStage` to `TokenVerify`.
	//
	// example:
	//
	// 63****
	TokenCode *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
	// The unique identifier of the client. When you use an Alibaba Cloud Workspace client, you can view the client version in the **About*	- dialog box on the client logon page.
	//
	// example:
	//
	// C78CA9E99315687575DD2844C1F3****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequest) SetAuthenticationCode(v string) *GetLoginTokenRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientId(v string) *GetLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientOS(v string) *GetLoginTokenRequest {
	s.ClientOS = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientType(v string) *GetLoginTokenRequest {
	s.ClientType = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientVersion(v string) *GetLoginTokenRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetLoginTokenRequest) SetCurrentStage(v string) *GetLoginTokenRequest {
	s.CurrentStage = &v
	return s
}

func (s *GetLoginTokenRequest) SetDirectoryId(v string) *GetLoginTokenRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetLoginTokenRequest) SetEndUserId(v string) *GetLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenRequest) SetKeepAlive(v bool) *GetLoginTokenRequest {
	s.KeepAlive = &v
	return s
}

func (s *GetLoginTokenRequest) SetKeepAliveToken(v string) *GetLoginTokenRequest {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenRequest) SetNewPassword(v string) *GetLoginTokenRequest {
	s.NewPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetOfficeSiteId(v string) *GetLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetLoginTokenRequest) SetOldPassword(v string) *GetLoginTokenRequest {
	s.OldPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetPassword(v string) *GetLoginTokenRequest {
	s.Password = &v
	return s
}

func (s *GetLoginTokenRequest) SetRegionId(v string) *GetLoginTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetLoginTokenRequest) SetSessionId(v string) *GetLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenRequest) SetTokenCode(v string) *GetLoginTokenRequest {
	s.TokenCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetUuid(v string) *GetLoginTokenRequest {
	s.Uuid = &v
	return s
}

type GetLoginTokenResponseBody struct {
	// The email address of the user. The system returns the email address in the return value of the LoginToken parameter after the user logs on to the client.
	//
	// 	- For a convenience user, the return value is the email address specified when the administrator creates the convenience user.
	//
	// 	- For an AD user, the return value is in the following format: `Username@Name of the AD domain`.
	//
	// example:
	//
	// alice
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The account of the convenience user or the AD user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// edu
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// The token used to keep the user logged on. After the user logs on to the client and select the Keep Logon option, `KeepAliveToken` is returned when you call the operation. If the user does not select the Keep Logon option, null is returned.
	//
	// example:
	//
	// 006YwvYMsesWWsDBZnVB+Wq9AvJDVIqOY3YCktvtb7+KxMb3ClnNlV8+l/knhZYrXUmeP06IzkjF+IgcZ3vZKOyMprDyFHjCy1r27FRE/U7+geWCl8iQ+yF8GaCRHfJEkC2+ROs93HkT4tfHxyY1J8W7O7ZQGUC/cdCvm+cCP6FIy73IUuPuVR6PcKYXIpEZPW
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// The attribute of the convenience user. For an AD user, null is returned.
	//
	// example:
	//
	// test:sample
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The logon token.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The next stage that is expected to enter. For example, an administrator enables MFA in the EDS console. When an end user enters the password, that is, the end user completes the `ADPassword` stage, this parameter returns `MFAVerify`. This indicates that MFA is required.
	//
	// >  For more information about the authentication stages, see the `CurrentStage` parameter.
	//
	// example:
	//
	// MFAVerify
	NextStage *string `json:"NextStage,omitempty" xml:"NextStage,omitempty"`
	NickName  *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// > This is a parameter only for internal use.
	PasswordStrategy *GetLoginTokenResponseBodyPasswordStrategy `json:"PasswordStrategy,omitempty" xml:"PasswordStrategy,omitempty" type:"Struct"`
	// Enter the mobile number of the convenience user. For an AD user, null is returned.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// > This is a parameter only for internal use.
	Props map[string]*string `json:"Props,omitempty" xml:"Props,omitempty"`
	// The QR code that is generated when the virtual MFA device is bound. The value is encoded in Base64. This parameter can be empty. This parameter is required only when the CurrentStage parameter is set to `MFABind`.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	//
	// example:
	//
	// 5OCLLKKOJU5HPBX66H3QCTWY******
	QrCodePng *string `json:"QrCodePng,omitempty" xml:"QrCodePng,omitempty"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// null
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Risk identification information regarding the signin process.
	RiskVerifyInfo *GetLoginTokenResponseBodyRiskVerifyInfo `json:"RiskVerifyInfo,omitempty" xml:"RiskVerifyInfo,omitempty" type:"Struct"`
	// The key that is generated when you bind the virtual MFA device. This parameter is required when the CurrentStage parameter is set to `MFABind`.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	//
	// example:
	//
	// 5OCLLKKOJU5HPBX66H3QCTWYI7MH****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The ID of the session. The ID is returned the first time you call the `GetLoginToken` operation in the session. If MFA is required, you must specify this parameter in subsequent stages.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	//
	// example:
	//
	// d6ec166d-ab93-4286-bf7f-a18bb929****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The ID of the Alibaba Cloud account. The ID is used for hardware client authentication.
	//
	// example:
	//
	// 166353906220****
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// mode
	WindowDisplayMode *string `json:"WindowDisplayMode,omitempty" xml:"WindowDisplayMode,omitempty"`
}

func (s GetLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) SetEmail(v string) *GetLoginTokenResponseBody {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEndUserId(v string) *GetLoginTokenResponseBody {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetIndustry(v string) *GetLoginTokenResponseBody {
	s.Industry = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetKeepAliveToken(v string) *GetLoginTokenResponseBody {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLabel(v string) *GetLoginTokenResponseBody {
	s.Label = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLoginToken(v string) *GetLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetNextStage(v string) *GetLoginTokenResponseBody {
	s.NextStage = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetNickName(v string) *GetLoginTokenResponseBody {
	s.NickName = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetPasswordStrategy(v *GetLoginTokenResponseBodyPasswordStrategy) *GetLoginTokenResponseBody {
	s.PasswordStrategy = v
	return s
}

func (s *GetLoginTokenResponseBody) SetPhone(v string) *GetLoginTokenResponseBody {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetProps(v map[string]*string) *GetLoginTokenResponseBody {
	s.Props = v
	return s
}

func (s *GetLoginTokenResponseBody) SetQrCodePng(v string) *GetLoginTokenResponseBody {
	s.QrCodePng = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetReason(v string) *GetLoginTokenResponseBody {
	s.Reason = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRequestId(v string) *GetLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRiskVerifyInfo(v *GetLoginTokenResponseBodyRiskVerifyInfo) *GetLoginTokenResponseBody {
	s.RiskVerifyInfo = v
	return s
}

func (s *GetLoginTokenResponseBody) SetSecret(v string) *GetLoginTokenResponseBody {
	s.Secret = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetSessionId(v string) *GetLoginTokenResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantId(v int64) *GetLoginTokenResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWindowDisplayMode(v string) *GetLoginTokenResponseBody {
	s.WindowDisplayMode = &v
	return s
}

type GetLoginTokenResponseBodyPasswordStrategy struct {
	// > This is a parameter only for internal use.
	TenantAlternativeChars []*string `json:"TenantAlternativeChars,omitempty" xml:"TenantAlternativeChars,omitempty" type:"Repeated"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// null
	TenantPasswordLength *string `json:"TenantPasswordLength,omitempty" xml:"TenantPasswordLength,omitempty"`
}

func (s GetLoginTokenResponseBodyPasswordStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyPasswordStrategy) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantAlternativeChars(v []*string) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantAlternativeChars = v
	return s
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantPasswordLength(v string) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantPasswordLength = &v
	return s
}

type GetLoginTokenResponseBodyRiskVerifyInfo struct {
	// The email used for authentication.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The duration of the lock.
	//
	// example:
	//
	// 1713749778
	LastLockDuration *int64 `json:"LastLockDuration,omitempty" xml:"LastLockDuration,omitempty"`
	// Whether the account is locked or not.
	//
	// example:
	//
	// true
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The mobile number used for authentication.
	//
	// example:
	//
	// 1388888****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetEmail(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLastLockDuration(v int64) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.LastLockDuration = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLocked(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Locked = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetPhone(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Phone = &v
	return s
}

type GetLoginTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponse) SetHeaders(v map[string]*string) *GetLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLoginTokenResponse) SetStatusCode(v int32) *GetLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginTokenResponse) SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse {
	s.Body = v
	return s
}

type IsKeepAliveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s IsKeepAliveRequest) String() string {
	return tea.Prettify(s)
}

func (s IsKeepAliveRequest) GoString() string {
	return s.String()
}

func (s *IsKeepAliveRequest) SetClientId(v string) *IsKeepAliveRequest {
	s.ClientId = &v
	return s
}

func (s *IsKeepAliveRequest) SetOfficeSiteId(v string) *IsKeepAliveRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *IsKeepAliveRequest) SetRegionId(v string) *IsKeepAliveRequest {
	s.RegionId = &v
	return s
}

type IsKeepAliveResponseBody struct {
	// example:
	//
	// True
	IsKeepAlive *bool `json:"IsKeepAlive,omitempty" xml:"IsKeepAlive,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 141631846826****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s IsKeepAliveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsKeepAliveResponseBody) GoString() string {
	return s.String()
}

func (s *IsKeepAliveResponseBody) SetIsKeepAlive(v bool) *IsKeepAliveResponseBody {
	s.IsKeepAlive = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetOfficeSiteId(v string) *IsKeepAliveResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetRequestId(v string) *IsKeepAliveResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetTenantId(v string) *IsKeepAliveResponseBody {
	s.TenantId = &v
	return s
}

type IsKeepAliveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsKeepAliveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsKeepAliveResponse) String() string {
	return tea.Prettify(s)
}

func (s IsKeepAliveResponse) GoString() string {
	return s.String()
}

func (s *IsKeepAliveResponse) SetHeaders(v map[string]*string) *IsKeepAliveResponse {
	s.Headers = v
	return s
}

func (s *IsKeepAliveResponse) SetStatusCode(v int32) *IsKeepAliveResponse {
	s.StatusCode = &v
	return s
}

func (s *IsKeepAliveResponse) SetBody(v *IsKeepAliveResponseBody) *IsKeepAliveResponse {
	s.Body = v
	return s
}

type QueryEdsAgentReportConfigRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
}

func (s QueryEdsAgentReportConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigRequest) SetAliUid(v int64) *QueryEdsAgentReportConfigRequest {
	s.AliUid = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) SetDesktopId(v string) *QueryEdsAgentReportConfigRequest {
	s.DesktopId = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) SetEcsInstanceId(v string) *QueryEdsAgentReportConfigRequest {
	s.EcsInstanceId = &v
	return s
}

type QueryEdsAgentReportConfigResponseBody struct {
	Data      *QueryEdsAgentReportConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEdsAgentReportConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponseBody) SetData(v *QueryEdsAgentReportConfigResponseBodyData) *QueryEdsAgentReportConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryEdsAgentReportConfigResponseBody) SetRequestId(v string) *QueryEdsAgentReportConfigResponseBody {
	s.RequestId = &v
	return s
}

type QueryEdsAgentReportConfigResponseBodyData struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s QueryEdsAgentReportConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponseBodyData) SetConfig(v string) *QueryEdsAgentReportConfigResponseBodyData {
	s.Config = &v
	return s
}

type QueryEdsAgentReportConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEdsAgentReportConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEdsAgentReportConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponse) SetHeaders(v map[string]*string) *QueryEdsAgentReportConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) SetStatusCode(v int32) *QueryEdsAgentReportConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) SetBody(v *QueryEdsAgentReportConfigResponseBody) *QueryEdsAgentReportConfigResponse {
	s.Body = v
	return s
}

type RebootDesktopsRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that runs the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client).
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence of a request?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 40401e62-5caf-4508-8de7-bf98af12****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can view the client version in the **About*	- dialog box on the client logon page.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OsUpdate   *bool   `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The logon token.
	//
	// example:
	//
	// 04b7b80a0b020715c5c1b4175fc4771698****9e2a759557a4624665fd53ae40
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	// The UUID of the client.
	//
	// example:
	//
	// 91761ED27169E2FC564F29388E2D****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetClientId(v string) *RebootDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientOS(v string) *RebootDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientToken(v string) *RebootDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientVersion(v string) *RebootDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebootDesktopsRequest) SetLoginToken(v string) *RebootDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *RebootDesktopsRequest) SetOsUpdate(v bool) *RebootDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetSessionId(v string) *RebootDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetSessionToken(v string) *RebootDesktopsRequest {
	s.SessionToken = &v
	return s
}

func (s *RebootDesktopsRequest) SetUuid(v string) *RebootDesktopsRequest {
	s.Uuid = &v
	return s
}

type RebootDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponseBody) SetRequestId(v string) *RebootDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RebootDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetHeaders(v map[string]*string) *RebootDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebootDesktopsResponse) SetStatusCode(v int32) *RebootDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootDesktopsResponse) SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse {
	s.Body = v
	return s
}

type RefreshLoginTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-shanghai+dir-238191****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RefreshLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenRequest) SetClientId(v string) *RefreshLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetDirectoryId(v string) *RefreshLoginTokenRequest {
	s.DirectoryId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetEndUserId(v string) *RefreshLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginToken(v string) *RefreshLoginTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetOfficeSiteId(v string) *RefreshLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetRegionId(v string) *RefreshLoginTokenRequest {
	s.RegionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetSessionId(v string) *RefreshLoginTokenRequest {
	s.SessionId = &v
	return s
}

type RefreshLoginTokenResponseBody struct {
	// example:
	//
	// v1c27bab6c205b2fdfac916434306375722776d6aa89e30b7836d18c95ade9137f0f5ac4325260782184e96ee2b3f0****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponseBody) SetLoginToken(v string) *RefreshLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenResponseBody) SetRequestId(v string) *RefreshLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshLoginTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponse) SetHeaders(v map[string]*string) *RefreshLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLoginTokenResponse) SetStatusCode(v int32) *RefreshLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLoginTokenResponse) SetBody(v *RefreshLoginTokenResponseBody) *RefreshLoginTokenResponse {
	s.Body = v
	return s
}

type ReportEdsAgentInfoRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EdsAgentInfo  *string `json:"EdsAgentInfo,omitempty" xml:"EdsAgentInfo,omitempty"`
}

func (s ReportEdsAgentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportEdsAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoRequest) SetAliUid(v int64) *ReportEdsAgentInfoRequest {
	s.AliUid = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetDesktopId(v string) *ReportEdsAgentInfoRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetEcsInstanceId(v string) *ReportEdsAgentInfoRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetEdsAgentInfo(v string) *ReportEdsAgentInfoRequest {
	s.EdsAgentInfo = &v
	return s
}

type ReportEdsAgentInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportEdsAgentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportEdsAgentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoResponseBody) SetRequestId(v string) *ReportEdsAgentInfoResponseBody {
	s.RequestId = &v
	return s
}

type ReportEdsAgentInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportEdsAgentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportEdsAgentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportEdsAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoResponse) SetHeaders(v map[string]*string) *ReportEdsAgentInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportEdsAgentInfoResponse) SetStatusCode(v int32) *ReportEdsAgentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportEdsAgentInfoResponse) SetBody(v *ReportEdsAgentInfoResponseBody) *ReportEdsAgentInfoResponse {
	s.Body = v
	return s
}

type ReportSessionStatusRequest struct {
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp167fcodoa90ixn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1642909143781
	SessionChangeTime *int64 `json:"SessionChangeTime,omitempty" xml:"SessionChangeTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SessionLogOn
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
}

func (s ReportSessionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportSessionStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusRequest) SetEndUserId(v string) *ReportSessionStatusRequest {
	s.EndUserId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetInstanceId(v string) *ReportSessionStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetRegionId(v string) *ReportSessionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionChangeTime(v int64) *ReportSessionStatusRequest {
	s.SessionChangeTime = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionId(v string) *ReportSessionStatusRequest {
	s.SessionId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionStatus(v string) *ReportSessionStatusRequest {
	s.SessionStatus = &v
	return s
}

type ReportSessionStatusResponseBody struct {
	// example:
	//
	// 0EE5DE20-25F4-5870-9D56-C259A45B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportSessionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportSessionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusResponseBody) SetRequestId(v string) *ReportSessionStatusResponseBody {
	s.RequestId = &v
	return s
}

type ReportSessionStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportSessionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportSessionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportSessionStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusResponse) SetHeaders(v map[string]*string) *ReportSessionStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportSessionStatusResponse) SetStatusCode(v int32) *ReportSessionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportSessionStatusResponse) SetBody(v *ReportSessionStatusResponseBody) *ReportSessionStatusResponse {
	s.Body = v
	return s
}

type ResetPasswordRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95e41934-383e-4c9f-824f-3b93b19b****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 2f00ab32-a473-4c90-9aae-dd8842ae****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// a***@example.edu
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-899235****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The phone number of the user.
	//
	// example:
	//
	// 1827912****
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s ResetPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetPasswordRequest) SetClientId(v string) *ResetPasswordRequest {
	s.ClientId = &v
	return s
}

func (s *ResetPasswordRequest) SetClientToken(v string) *ResetPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetPasswordRequest) SetEmail(v string) *ResetPasswordRequest {
	s.Email = &v
	return s
}

func (s *ResetPasswordRequest) SetEndUserId(v string) *ResetPasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ResetPasswordRequest) SetOfficeSiteId(v string) *ResetPasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ResetPasswordRequest) SetRegionId(v string) *ResetPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetPasswordRequest) SetPhone(v string) *ResetPasswordRequest {
	s.Phone = &v
	return s
}

type ResetPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A00477A5-167F-56D2-A315-EA77E4BD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponseBody) SetRequestId(v string) *ResetPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetPasswordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponse) SetHeaders(v map[string]*string) *ResetPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetPasswordResponse) SetStatusCode(v int32) *ResetPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetPasswordResponse) SetBody(v *ResetPasswordResponseBody) *ResetPasswordResponse {
	s.Body = v
	return s
}

type ResetSnapshotRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// b9d8ddfd-65d4-4857-9e97-56477d1f****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-e964cr92klwqb****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1fdef51b727aa91d6c881658978508114d3f5680fa99a66b2a631d17d5bb4860cccf1173be24d77d5ef1423c83aea****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 05182b8c-bb0d-49d3-963c-ee63a507****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-2zeipxmnhej803x7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// Specifies whether to stop the cloud computer.
	//
	// example:
	//
	// true
	StopDesktop *bool `json:"StopDesktop,omitempty" xml:"StopDesktop,omitempty"`
}

func (s ResetSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) SetClientId(v string) *ResetSnapshotRequest {
	s.ClientId = &v
	return s
}

func (s *ResetSnapshotRequest) SetDesktopId(v string) *ResetSnapshotRequest {
	s.DesktopId = &v
	return s
}

func (s *ResetSnapshotRequest) SetLoginToken(v string) *ResetSnapshotRequest {
	s.LoginToken = &v
	return s
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSessionId(v string) *ResetSnapshotRequest {
	s.SessionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetSnapshotRequest) SetStopDesktop(v bool) *ResetSnapshotRequest {
	s.StopDesktop = &v
	return s
}

type ResetSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponseBody) SetRequestId(v string) *ResetSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type ResetSnapshotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) SetHeaders(v map[string]*string) *ResetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ResetSnapshotResponse) SetStatusCode(v int32) *ResetSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSnapshotResponse) SetBody(v *ResetSnapshotResponseBody) *ResetSnapshotResponse {
	s.Body = v
	return s
}

type SendTokenCodeRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system on which the client runs.
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client version. If you use an Alibaba Cloud Workspace client, you can view the client version in the "About" dialog box on the client logon page.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The logon token.
	//
	// example:
	//
	// v28101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-2925105532
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// If two-factor authentication is enabled for clients in the Elastic Desktop Service (EDS) Enterprise console, the system will send a verification code to the user\\"s email address if it detects that the current logged-on user is at risk. This parameter is required if you set `CurrentStage` to `TokenVerify`.
	//
	// example:
	//
	// 63****
	TokenCode *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
}

func (s SendTokenCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTokenCodeRequest) GoString() string {
	return s.String()
}

func (s *SendTokenCodeRequest) SetClientId(v string) *SendTokenCodeRequest {
	s.ClientId = &v
	return s
}

func (s *SendTokenCodeRequest) SetClientOS(v string) *SendTokenCodeRequest {
	s.ClientOS = &v
	return s
}

func (s *SendTokenCodeRequest) SetClientVersion(v string) *SendTokenCodeRequest {
	s.ClientVersion = &v
	return s
}

func (s *SendTokenCodeRequest) SetEndUserId(v string) *SendTokenCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *SendTokenCodeRequest) SetLoginToken(v string) *SendTokenCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *SendTokenCodeRequest) SetOfficeSiteId(v string) *SendTokenCodeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SendTokenCodeRequest) SetSessionId(v string) *SendTokenCodeRequest {
	s.SessionId = &v
	return s
}

func (s *SendTokenCodeRequest) SetTokenCode(v string) *SendTokenCodeRequest {
	s.TokenCode = &v
	return s
}

type SendTokenCodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 134BD0B2-B848-5743-9CE2-C1FD3D5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTokenCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTokenCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendTokenCodeResponseBody) SetRequestId(v string) *SendTokenCodeResponseBody {
	s.RequestId = &v
	return s
}

type SendTokenCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTokenCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTokenCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTokenCodeResponse) GoString() string {
	return s.String()
}

func (s *SendTokenCodeResponse) SetHeaders(v map[string]*string) *SendTokenCodeResponse {
	s.Headers = v
	return s
}

func (s *SendTokenCodeResponse) SetStatusCode(v int32) *SendTokenCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTokenCodeResponse) SetBody(v *SendTokenCodeResponseBody) *SendTokenCodeResponse {
	s.Body = v
	return s
}

type SetFingerPrintTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 347431a9-90f6-448e-82c4-42bc84a9****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// AAAAAAAAAAAAAA
	EncryptedFingerPrintTemplate *string `json:"EncryptedFingerPrintTemplate,omitempty" xml:"EncryptedFingerPrintTemplate,omitempty"`
	// example:
	//
	// drjfspchj
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// example:
	//
	// goG3gG8AAABhujtscn
	FingerPrintTemplate *string `json:"FingerPrintTemplate,omitempty" xml:"FingerPrintTemplate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v11c73e7af0cb43ff39301651142485099ffb447085d76c4147519dbaa21c3bd90d53045e327c1f525ee6331c52556****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// As53328794
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8b42538a-246e-45a1-95ea-e5c65b09****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SetFingerPrintTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateRequest) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateRequest) SetClientId(v string) *SetFingerPrintTemplateRequest {
	s.ClientId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetClientToken(v string) *SetFingerPrintTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetDescription(v string) *SetFingerPrintTemplateRequest {
	s.Description = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetEncryptedFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest {
	s.EncryptedFingerPrintTemplate = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetEncryptedKey(v string) *SetFingerPrintTemplateRequest {
	s.EncryptedKey = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest {
	s.FingerPrintTemplate = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetLoginToken(v string) *SetFingerPrintTemplateRequest {
	s.LoginToken = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetPassword(v string) *SetFingerPrintTemplateRequest {
	s.Password = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetRegionId(v string) *SetFingerPrintTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetSessionId(v string) *SetFingerPrintTemplateRequest {
	s.SessionId = &v
	return s
}

type SetFingerPrintTemplateResponseBody struct {
	// example:
	//
	// 0711abb9-4cf8-41b2-9d0e-b51209468631;da4VFPNxwY3CZegFjOrCNw==;iHp2l9/qGcfD4tWx7jZIZQ==
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// CDE666EA-4FCD-5024-895C-8698E3D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFingerPrintTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateResponseBody) SetEncryptedPassword(v string) *SetFingerPrintTemplateResponseBody {
	s.EncryptedPassword = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) SetIndex(v int32) *SetFingerPrintTemplateResponseBody {
	s.Index = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) SetRequestId(v string) *SetFingerPrintTemplateResponseBody {
	s.RequestId = &v
	return s
}

type SetFingerPrintTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFingerPrintTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFingerPrintTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateResponse) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateResponse) SetHeaders(v map[string]*string) *SetFingerPrintTemplateResponse {
	s.Headers = v
	return s
}

func (s *SetFingerPrintTemplateResponse) SetStatusCode(v int32) *SetFingerPrintTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFingerPrintTemplateResponse) SetBody(v *SetFingerPrintTemplateResponseBody) *SetFingerPrintTemplateResponse {
	s.Body = v
	return s
}

type SetFingerPrintTemplateDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0764064c-1609-4d3c-8cb7-ab8d3feg****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 40401e62-5caf-4508-8de7-bf98af12****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Finger 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v14e5a2404c495249f7541646535779667ea0b5d87754b5d2d2a3099bda774f3832e24756ef3e66eb574b1f3e99078****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d28520d4-da0b-4a97-981d-683db865****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionRequest) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetClientId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.ClientId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetClientToken(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetDescription(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.Description = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetIndex(v int32) *SetFingerPrintTemplateDescriptionRequest {
	s.Index = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetLoginToken(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.LoginToken = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetRegionId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetSessionId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.SessionId = &v
	return s
}

type SetFingerPrintTemplateDescriptionResponseBody struct {
	// example:
	//
	// BBD7DFD1-A5DE-51D9-8FD6-3BF54EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionResponseBody) SetRequestId(v string) *SetFingerPrintTemplateDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type SetFingerPrintTemplateDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFingerPrintTemplateDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionResponse) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetHeaders(v map[string]*string) *SetFingerPrintTemplateDescriptionResponse {
	s.Headers = v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetStatusCode(v int32) *SetFingerPrintTemplateDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetBody(v *SetFingerPrintTemplateDescriptionResponseBody) *SetFingerPrintTemplateDescriptionResponse {
	s.Body = v
	return s
}

type StartDesktopsRequest struct {
	// The ID of the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client). The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that run the client.
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 21e7be12-aa4f-4389-b3e1-82f4a1b5****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can click **About*	- on the client logon page to view the version of the client.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-cg27ufmapab08****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The UUID of the client.
	//
	// example:
	//
	// 71F6A700735E74A61161A53F0C47****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetClientId(v string) *StartDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *StartDesktopsRequest) SetClientOS(v string) *StartDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StartDesktopsRequest) SetClientToken(v string) *StartDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDesktopsRequest) SetClientVersion(v string) *StartDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StartDesktopsRequest) SetLoginToken(v string) *StartDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) SetSessionId(v string) *StartDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *StartDesktopsRequest) SetUuid(v string) *StartDesktopsRequest {
	s.Uuid = &v
	return s
}

type StartDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponseBody) SetRequestId(v string) *StartDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StartDesktopsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetHeaders(v map[string]*string) *StartDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopsResponse) SetStatusCode(v int32) *StartDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDesktopsResponse) SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse {
	s.Body = v
	return s
}

type StartRecordContentRequest struct {
	// This parameter is required.
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StartRecordContentRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRecordContentRequest) GoString() string {
	return s.String()
}

func (s *StartRecordContentRequest) SetClientId(v string) *StartRecordContentRequest {
	s.ClientId = &v
	return s
}

func (s *StartRecordContentRequest) SetClientOS(v string) *StartRecordContentRequest {
	s.ClientOS = &v
	return s
}

func (s *StartRecordContentRequest) SetClientVersion(v string) *StartRecordContentRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartRecordContentRequest) SetDesktopId(v string) *StartRecordContentRequest {
	s.DesktopId = &v
	return s
}

func (s *StartRecordContentRequest) SetFilePath(v string) *StartRecordContentRequest {
	s.FilePath = &v
	return s
}

func (s *StartRecordContentRequest) SetLoginToken(v string) *StartRecordContentRequest {
	s.LoginToken = &v
	return s
}

func (s *StartRecordContentRequest) SetRegionId(v string) *StartRecordContentRequest {
	s.RegionId = &v
	return s
}

func (s *StartRecordContentRequest) SetSessionId(v string) *StartRecordContentRequest {
	s.SessionId = &v
	return s
}

type StartRecordContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordContentResponseBody) SetRequestId(v string) *StartRecordContentResponseBody {
	s.RequestId = &v
	return s
}

type StartRecordContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRecordContentResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRecordContentResponse) GoString() string {
	return s.String()
}

func (s *StartRecordContentResponse) SetHeaders(v map[string]*string) *StartRecordContentResponse {
	s.Headers = v
	return s
}

func (s *StartRecordContentResponse) SetStatusCode(v int32) *StartRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRecordContentResponse) SetBody(v *StartRecordContentResponseBody) *StartRecordContentResponse {
	s.Body = v
	return s
}

type StopDesktopsRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that runs the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client).
	//
	// example:
	//
	// Windows_NT 10.0.18363 x64
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence of a request?](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 6ce412a8-399f-49f9-9518-66ee028a****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can view the client version in the **About*	- dialog box on the client logon page.
	//
	// example:
	//
	// 2.1.0-R-20210731.151756
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OsUpdate   *bool   `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The logon token.
	//
	// example:
	//
	// 04b7b80a0b020715c5c1b4175fc4771698****9e2a759557a4624665fd53ae40
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetClientId(v string) *StopDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *StopDesktopsRequest) SetClientOS(v string) *StopDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StopDesktopsRequest) SetClientToken(v string) *StopDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDesktopsRequest) SetClientVersion(v string) *StopDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetLoginToken(v string) *StopDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *StopDesktopsRequest) SetOsUpdate(v bool) *StopDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetSessionId(v string) *StopDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *StopDesktopsRequest) SetSessionToken(v string) *StopDesktopsRequest {
	s.SessionToken = &v
	return s
}

func (s *StopDesktopsRequest) SetUuid(v string) *StopDesktopsRequest {
	s.Uuid = &v
	return s
}

type StopDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponseBody) SetRequestId(v string) *StopDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StopDesktopsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetHeaders(v map[string]*string) *StopDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopsResponse) SetStatusCode(v int32) *StopDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDesktopsResponse) SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse {
	s.Body = v
	return s
}

type StopRecordContentRequest struct {
	// This parameter is required.
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StopRecordContentRequest) String() string {
	return tea.Prettify(s)
}

func (s StopRecordContentRequest) GoString() string {
	return s.String()
}

func (s *StopRecordContentRequest) SetClientId(v string) *StopRecordContentRequest {
	s.ClientId = &v
	return s
}

func (s *StopRecordContentRequest) SetClientOS(v string) *StopRecordContentRequest {
	s.ClientOS = &v
	return s
}

func (s *StopRecordContentRequest) SetClientVersion(v string) *StopRecordContentRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopRecordContentRequest) SetDesktopId(v string) *StopRecordContentRequest {
	s.DesktopId = &v
	return s
}

func (s *StopRecordContentRequest) SetLoginToken(v string) *StopRecordContentRequest {
	s.LoginToken = &v
	return s
}

func (s *StopRecordContentRequest) SetRegionId(v string) *StopRecordContentRequest {
	s.RegionId = &v
	return s
}

func (s *StopRecordContentRequest) SetSessionId(v string) *StopRecordContentRequest {
	s.SessionId = &v
	return s
}

type StopRecordContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordContentResponseBody) SetRequestId(v string) *StopRecordContentResponseBody {
	s.RequestId = &v
	return s
}

type StopRecordContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRecordContentResponse) String() string {
	return tea.Prettify(s)
}

func (s StopRecordContentResponse) GoString() string {
	return s.String()
}

func (s *StopRecordContentResponse) SetHeaders(v map[string]*string) *StopRecordContentResponse {
	s.Headers = v
	return s
}

func (s *StopRecordContentResponse) SetStatusCode(v int32) *StopRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRecordContentResponse) SetBody(v *StopRecordContentResponseBody) *StopRecordContentResponse {
	s.Body = v
	return s
}

type UnbindUserDesktopRequest struct {
	// The client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58f96f67-7944-4f97-9342-****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client type.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether to enable forced unbinding.
	//
	// Valid values:
	//
	// 	- true: Even when end users connect to cloud computers, the forced unbinding is still enforced.
	//
	// 	- false: Forced unbinding is only enforced when end users are disconnected from cloud computers.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v12307f5e0****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b053331-dc98-43d8-b247-****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The cloud computer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ud-sdfs****
	UserDesktopId *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
}

func (s UnbindUserDesktopRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserDesktopRequest) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopRequest) SetClientId(v string) *UnbindUserDesktopRequest {
	s.ClientId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetClientType(v string) *UnbindUserDesktopRequest {
	s.ClientType = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetForce(v bool) *UnbindUserDesktopRequest {
	s.Force = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetLoginToken(v string) *UnbindUserDesktopRequest {
	s.LoginToken = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetRegionId(v string) *UnbindUserDesktopRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetSessionId(v string) *UnbindUserDesktopRequest {
	s.SessionId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetUserDesktopId(v string) *UnbindUserDesktopRequest {
	s.UserDesktopId = &v
	return s
}

type UnbindUserDesktopResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D5B7CF35-E078-5EBF-A010-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindUserDesktopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserDesktopResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopResponseBody) SetRequestId(v string) *UnbindUserDesktopResponseBody {
	s.RequestId = &v
	return s
}

type UnbindUserDesktopResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindUserDesktopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindUserDesktopResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserDesktopResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopResponse) SetHeaders(v map[string]*string) *UnbindUserDesktopResponse {
	s.Headers = v
	return s
}

func (s *UnbindUserDesktopResponse) SetStatusCode(v int32) *UnbindUserDesktopResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindUserDesktopResponse) SetBody(v *UnbindUserDesktopResponseBody) *UnbindUserDesktopResponse {
	s.Body = v
	return s
}

type VerifyCredentialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d0b95762-0541-4b53-a0e4-7ed09f39****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456789cyG
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// example:
	//
	// Password
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// example:
	//
	// drjfs****
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1f5772a1c60dbea9fd8e1648567079018086448d234b5bc8e30bec0ba6e80c41c767c4dd0db51e9e5c4e0f111431a****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+dir-227468****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// eb17af2e-1dd6-4cc4-a3ee-3a14d0d7****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s VerifyCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyCredentialRequest) GoString() string {
	return s.String()
}

func (s *VerifyCredentialRequest) SetClientId(v string) *VerifyCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *VerifyCredentialRequest) SetCredential(v string) *VerifyCredentialRequest {
	s.Credential = &v
	return s
}

func (s *VerifyCredentialRequest) SetCredentialType(v string) *VerifyCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *VerifyCredentialRequest) SetEncryptedKey(v string) *VerifyCredentialRequest {
	s.EncryptedKey = &v
	return s
}

func (s *VerifyCredentialRequest) SetLoginToken(v string) *VerifyCredentialRequest {
	s.LoginToken = &v
	return s
}

func (s *VerifyCredentialRequest) SetOfficeSiteId(v string) *VerifyCredentialRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *VerifyCredentialRequest) SetRegionId(v string) *VerifyCredentialRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyCredentialRequest) SetSessionId(v string) *VerifyCredentialRequest {
	s.SessionId = &v
	return s
}

type VerifyCredentialResponseBody struct {
	// example:
	//
	// D5F0BDFB-A229-5F1D-B790-33709D43****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCredentialResponseBody) SetRequestId(v string) *VerifyCredentialResponseBody {
	s.RequestId = &v
	return s
}

type VerifyCredentialResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyCredentialResponse) GoString() string {
	return s.String()
}

func (s *VerifyCredentialResponse) SetHeaders(v map[string]*string) *VerifyCredentialResponse {
	s.Headers = v
	return s
}

func (s *VerifyCredentialResponse) SetStatusCode(v int32) *VerifyCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCredentialResponse) SetBody(v *VerifyCredentialResponseBody) *VerifyCredentialResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 允许桌面FOTA升级
//
// @param request - ApproveFotaUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveFotaUpdateResponse
func (client *Client) ApproveFotaUpdateWithOptions(request *ApproveFotaUpdateRequest, runtime *util.RuntimeOptions) (_result *ApproveFotaUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetStatus)) {
		query["TargetStatus"] = request.TargetStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveFotaUpdate"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 允许桌面FOTA升级
//
// @param request - ApproveFotaUpdateRequest
//
// @return ApproveFotaUpdateResponse
func (client *Client) ApproveFotaUpdate(request *ApproveFotaUpdateRequest) (_result *ApproveFotaUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.ApproveFotaUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password of a user account.
//
// @param request - ChangePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangePasswordResponse
func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *util.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangePassword"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password of a user account.
//
// @param request - ChangePasswordRequest
//
// @return ChangePasswordResponse
func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFingerPrintTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFingerPrintTemplateResponse
func (client *Client) DeleteFingerPrintTemplateWithOptions(request *DeleteFingerPrintTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteFingerPrintTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Index)) {
		query["Index"] = request.Index
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFingerPrintTemplate"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFingerPrintTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFingerPrintTemplateRequest
//
// @return DeleteFingerPrintTemplateResponse
func (client *Client) DeleteFingerPrintTemplate(request *DeleteFingerPrintTemplateRequest) (_result *DeleteFingerPrintTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFingerPrintTemplateResponse{}
	_body, _err := client.DeleteFingerPrintTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries directory details.
//
// @param request - DescribeDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirectories"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries directory details.
//
// @param request - DescribeDirectoriesRequest
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries fingerprint templates.
//
// @param request - DescribeFingerPrintTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFingerPrintTemplatesResponse
func (client *Client) DescribeFingerPrintTemplatesWithOptions(request *DescribeFingerPrintTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeFingerPrintTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFingerPrintTemplates"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFingerPrintTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries fingerprint templates.
//
// @param request - DescribeFingerPrintTemplatesRequest
//
// @return DescribeFingerPrintTemplatesResponse
func (client *Client) DescribeFingerPrintTemplates(request *DescribeFingerPrintTemplatesRequest) (_result *DescribeFingerPrintTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFingerPrintTemplatesResponse{}
	_body, _err := client.DescribeFingerPrintTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeGlobalDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDesktopsResponse
func (client *Client) DescribeGlobalDesktopsWithOptions(request *DescribeGlobalDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeGlobalDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopStatus)) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.QueryFotaUpdate)) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRegionId)) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.WithoutLatency)) {
		query["WithoutLatency"] = request.WithoutLatency
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGlobalDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGlobalDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeGlobalDesktopsRequest
//
// @return DescribeGlobalDesktopsResponse
func (client *Client) DescribeGlobalDesktops(request *DescribeGlobalDesktopsRequest) (_result *DescribeGlobalDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGlobalDesktopsResponse{}
	_body, _err := client.DescribeGlobalDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeOfficeSitesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOfficeSitesResponse
func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOfficeSites"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeOfficeSitesRequest
//
// @return DescribeOfficeSitesResponse
func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
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

// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// 列举快照
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnapshots"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举快照
//
// @param request - DescribeSnapshotsRequest
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户资源列表
//
// @param request - DescribeUserResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserResourcesResponse
func (client *Client) DescribeUserResourcesWithOptions(request *DescribeUserResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeUserResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefresh)) {
		query["AutoRefresh"] = request.AutoRefresh
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryType)) {
		query["CategoryType"] = request.CategoryType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DualCenterForward)) {
		query["DualCenterForward"] = request.DualCenterForward
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteIds)) {
		query["OfficeSiteIds"] = request.OfficeSiteIds
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.ProductTypes)) {
		query["ProductTypes"] = request.ProductTypes
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryFotaUpdate)) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshFotaUpdate)) {
		query["RefreshFotaUpdate"] = request.RefreshFotaUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRegionId)) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserResources"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户资源列表
//
// @param request - DescribeUserResourcesRequest
//
// @return DescribeUserResourcesResponse
func (client *Client) DescribeUserResources(request *DescribeUserResourcesRequest) (_result *DescribeUserResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserResourcesResponse{}
	_body, _err := client.DescribeUserResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts a password.
//
// @param request - EncryptPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptPasswordResponse
func (client *Client) EncryptPasswordWithOptions(request *EncryptPasswordRequest, runtime *util.RuntimeOptions) (_result *EncryptPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EncryptPassword"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EncryptPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts a password.
//
// @param request - EncryptPasswordRequest
//
// @return EncryptPasswordResponse
func (client *Client) EncryptPassword(request *EncryptPasswordRequest) (_result *EncryptPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncryptPasswordResponse{}
	_body, _err := client.EncryptPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取无影云盘的免密token
//
// @param request - GetCloudDriveServiceMountTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudDriveServiceMountTokenResponse
func (client *Client) GetCloudDriveServiceMountTokenWithOptions(request *GetCloudDriveServiceMountTokenRequest, runtime *util.RuntimeOptions) (_result *GetCloudDriveServiceMountTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCloudDriveServiceMountToken"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCloudDriveServiceMountTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取无影云盘的免密token
//
// @param request - GetCloudDriveServiceMountTokenRequest
//
// @return GetCloudDriveServiceMountTokenResponse
func (client *Client) GetCloudDriveServiceMountToken(request *GetCloudDriveServiceMountTokenRequest) (_result *GetCloudDriveServiceMountTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudDriveServiceMountTokenResponse{}
	_body, _err := client.GetCloudDriveServiceMountTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetConnectionTicketRequest
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains logon credentials.
//
// @param request - GetLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginTokenWithOptions(request *GetLoginTokenRequest, runtime *util.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode)) {
		query["AuthenticationCode"] = request.AuthenticationCode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentStage)) {
		query["CurrentStage"] = request.CurrentStage
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAlive)) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAliveToken)) {
		query["KeepAliveToken"] = request.KeepAliveToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TokenCode)) {
		query["TokenCode"] = request.TokenCode
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoginToken"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains logon credentials.
//
// @param request - GetLoginTokenRequest
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginToken(request *GetLoginTokenRequest) (_result *GetLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.GetLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 是否保持登录判断接口
//
// @param request - IsKeepAliveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsKeepAliveResponse
func (client *Client) IsKeepAliveWithOptions(request *IsKeepAliveRequest, runtime *util.RuntimeOptions) (_result *IsKeepAliveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IsKeepAlive"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IsKeepAliveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 是否保持登录判断接口
//
// @param request - IsKeepAliveRequest
//
// @return IsKeepAliveResponse
func (client *Client) IsKeepAlive(request *IsKeepAliveRequest) (_result *IsKeepAliveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsKeepAliveResponse{}
	_body, _err := client.IsKeepAliveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Agent需要上报的配置信息
//
// @param request - QueryEdsAgentReportConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEdsAgentReportConfigResponse
func (client *Client) QueryEdsAgentReportConfigWithOptions(request *QueryEdsAgentReportConfigRequest, runtime *util.RuntimeOptions) (_result *QueryEdsAgentReportConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEdsAgentReportConfig"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEdsAgentReportConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Agent需要上报的配置信息
//
// @param request - QueryEdsAgentReportConfigRequest
//
// @return QueryEdsAgentReportConfigResponse
func (client *Client) QueryEdsAgentReportConfig(request *QueryEdsAgentReportConfigRequest) (_result *QueryEdsAgentReportConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEdsAgentReportConfigResponse{}
	_body, _err := client.QueryEdsAgentReportConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restart cloud computers.
//
// @param request - RebootDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootDesktopsResponse
func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OsUpdate)) {
		query["OsUpdate"] = request.OsUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionToken)) {
		query["SessionToken"] = request.SessionToken
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restart cloud computers.
//
// @param request - RebootDesktopsRequest
//
// @return RebootDesktopsResponse
func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginTokenWithOptions(request *RefreshLoginTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshLoginToken"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginToken(request *RefreshLoginTokenRequest) (_result *RefreshLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.RefreshLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报edsAgent的信息
//
// @param request - ReportEdsAgentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportEdsAgentInfoResponse
func (client *Client) ReportEdsAgentInfoWithOptions(request *ReportEdsAgentInfoRequest, runtime *util.RuntimeOptions) (_result *ReportEdsAgentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EdsAgentInfo)) {
		query["EdsAgentInfo"] = request.EdsAgentInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportEdsAgentInfo"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportEdsAgentInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报edsAgent的信息
//
// @param request - ReportEdsAgentInfoRequest
//
// @return ReportEdsAgentInfoResponse
func (client *Client) ReportEdsAgentInfo(request *ReportEdsAgentInfoRequest) (_result *ReportEdsAgentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportEdsAgentInfoResponse{}
	_body, _err := client.ReportEdsAgentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReportSessionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportSessionStatusResponse
func (client *Client) ReportSessionStatusWithOptions(request *ReportSessionStatusRequest, runtime *util.RuntimeOptions) (_result *ReportSessionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionChangeTime)) {
		query["SessionChangeTime"] = request.SessionChangeTime
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionStatus)) {
		query["SessionStatus"] = request.SessionStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportSessionStatus"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportSessionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReportSessionStatusRequest
//
// @return ReportSessionStatusResponse
func (client *Client) ReportSessionStatus(request *ReportSessionStatusRequest) (_result *ReportSessionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportSessionStatusResponse{}
	_body, _err := client.ReportSessionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a password.
//
// @param request - ResetPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetPasswordResponse
func (client *Client) ResetPasswordWithOptions(request *ResetPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetPassword"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a password.
//
// @param request - ResetPasswordRequest
//
// @return ResetPasswordResponse
func (client *Client) ResetPassword(request *ResetPasswordRequest) (_result *ResetPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetPasswordResponse{}
	_body, _err := client.ResetPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores the data of a disk from a snapshot.
//
// @param request - ResetSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSnapshotResponse
func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *util.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.StopDesktop)) {
		query["StopDesktop"] = request.StopDesktop
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSnapshot"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores the data of a disk from a snapshot.
//
// @param request - ResetSnapshotRequest
//
// @return ResetSnapshotResponse
func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a logon verification code.
//
// @param request - SendTokenCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTokenCodeResponse
func (client *Client) SendTokenCodeWithOptions(request *SendTokenCodeRequest, runtime *util.RuntimeOptions) (_result *SendTokenCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TokenCode)) {
		query["TokenCode"] = request.TokenCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendTokenCode"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTokenCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a logon verification code.
//
// @param request - SendTokenCodeRequest
//
// @return SendTokenCodeResponse
func (client *Client) SendTokenCode(request *SendTokenCodeRequest) (_result *SendTokenCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendTokenCodeResponse{}
	_body, _err := client.SendTokenCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetFingerPrintTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFingerPrintTemplateResponse
func (client *Client) SetFingerPrintTemplateWithOptions(request *SetFingerPrintTemplateRequest, runtime *util.RuntimeOptions) (_result *SetFingerPrintTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedFingerPrintTemplate)) {
		query["EncryptedFingerPrintTemplate"] = request.EncryptedFingerPrintTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedKey)) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !tea.BoolValue(util.IsUnset(request.FingerPrintTemplate)) {
		query["FingerPrintTemplate"] = request.FingerPrintTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFingerPrintTemplate"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFingerPrintTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetFingerPrintTemplateRequest
//
// @return SetFingerPrintTemplateResponse
func (client *Client) SetFingerPrintTemplate(request *SetFingerPrintTemplateRequest) (_result *SetFingerPrintTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFingerPrintTemplateResponse{}
	_body, _err := client.SetFingerPrintTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetFingerPrintTemplateDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFingerPrintTemplateDescriptionResponse
func (client *Client) SetFingerPrintTemplateDescriptionWithOptions(request *SetFingerPrintTemplateDescriptionRequest, runtime *util.RuntimeOptions) (_result *SetFingerPrintTemplateDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Index)) {
		query["Index"] = request.Index
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFingerPrintTemplateDescription"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFingerPrintTemplateDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetFingerPrintTemplateDescriptionRequest
//
// @return SetFingerPrintTemplateDescriptionResponse
func (client *Client) SetFingerPrintTemplateDescription(request *SetFingerPrintTemplateDescriptionRequest) (_result *SetFingerPrintTemplateDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFingerPrintTemplateDescriptionResponse{}
	_body, _err := client.SetFingerPrintTemplateDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start cloud computers.
//
// Description:
//
// The cloud computers that you want to start must be in the Stopped state. After you call this operation, the cloud computers enter the Running state.
//
// @param request - StartDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDesktopsResponse
func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start cloud computers.
//
// Description:
//
// The cloud computers that you want to start must be in the Stopped state. After you call this operation, the cloud computers enter the Running state.
//
// @param request - StartDesktopsRequest
//
// @return StartDesktopsResponse
func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartRecordContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordContentResponse
func (client *Client) StartRecordContentWithOptions(request *StartRecordContentRequest, runtime *util.RuntimeOptions) (_result *StartRecordContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRecordContent"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRecordContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartRecordContentRequest
//
// @return StartRecordContentResponse
func (client *Client) StartRecordContent(request *StartRecordContentRequest) (_result *StartRecordContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRecordContentResponse{}
	_body, _err := client.StartRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops cloud computers.
//
// Description:
//
// The cloud computers that you want to stop must be in the Running state. After you call this operation, the cloud computers enter the Stopped state.
//
// @param request - StopDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDesktopsResponse
func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OsUpdate)) {
		query["OsUpdate"] = request.OsUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionToken)) {
		query["SessionToken"] = request.SessionToken
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops cloud computers.
//
// Description:
//
// The cloud computers that you want to stop must be in the Running state. After you call this operation, the cloud computers enter the Stopped state.
//
// @param request - StopDesktopsRequest
//
// @return StopDesktopsResponse
func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StopRecordContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordContentResponse
func (client *Client) StopRecordContentWithOptions(request *StopRecordContentRequest, runtime *util.RuntimeOptions) (_result *StopRecordContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopRecordContent"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopRecordContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopRecordContentRequest
//
// @return StopRecordContentResponse
func (client *Client) StopRecordContent(request *StopRecordContentRequest) (_result *StopRecordContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopRecordContentResponse{}
	_body, _err := client.StopRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds end users from cloud computers.
//
// @param request - UnbindUserDesktopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindUserDesktopResponse
func (client *Client) UnbindUserDesktopWithOptions(request *UnbindUserDesktopRequest, runtime *util.RuntimeOptions) (_result *UnbindUserDesktopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserDesktopId)) {
		query["UserDesktopId"] = request.UserDesktopId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindUserDesktop"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds end users from cloud computers.
//
// @param request - UnbindUserDesktopRequest
//
// @return UnbindUserDesktopResponse
func (client *Client) UnbindUserDesktop(request *UnbindUserDesktopRequest) (_result *UnbindUserDesktopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.UnbindUserDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifyCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyCredentialResponse
func (client *Client) VerifyCredentialWithOptions(request *VerifyCredentialRequest, runtime *util.RuntimeOptions) (_result *VerifyCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.Credential)) {
		query["Credential"] = request.Credential
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialType)) {
		query["CredentialType"] = request.CredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedKey)) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyCredential"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifyCredentialRequest
//
// @return VerifyCredentialResponse
func (client *Client) VerifyCredential(request *VerifyCredentialRequest) (_result *VerifyCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyCredentialResponse{}
	_body, _err := client.VerifyCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

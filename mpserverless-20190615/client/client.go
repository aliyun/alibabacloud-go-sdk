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

type AddCorsDomainRequest struct {
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s AddCorsDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCorsDomainRequest) GoString() string {
	return s.String()
}

func (s *AddCorsDomainRequest) SetDomain(v string) *AddCorsDomainRequest {
	s.Domain = &v
	return s
}

func (s *AddCorsDomainRequest) SetSpaceId(v string) *AddCorsDomainRequest {
	s.SpaceId = &v
	return s
}

type AddCorsDomainResponseBody struct {
	DomainId  *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCorsDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCorsDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddCorsDomainResponseBody) SetDomainId(v string) *AddCorsDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *AddCorsDomainResponseBody) SetRequestId(v string) *AddCorsDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddCorsDomainResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCorsDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCorsDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCorsDomainResponse) GoString() string {
	return s.String()
}

func (s *AddCorsDomainResponse) SetHeaders(v map[string]*string) *AddCorsDomainResponse {
	s.Headers = v
	return s
}

func (s *AddCorsDomainResponse) SetBody(v *AddCorsDomainResponseBody) *AddCorsDomainResponse {
	s.Body = v
	return s
}

type AddDingtalkOpenPlatformConfigRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s AddDingtalkOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDingtalkOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *AddDingtalkOpenPlatformConfigRequest) SetAppId(v string) *AddDingtalkOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *AddDingtalkOpenPlatformConfigRequest) SetAppSecret(v string) *AddDingtalkOpenPlatformConfigRequest {
	s.AppSecret = &v
	return s
}

func (s *AddDingtalkOpenPlatformConfigRequest) SetSpaceId(v string) *AddDingtalkOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type AddDingtalkOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDingtalkOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDingtalkOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddDingtalkOpenPlatformConfigResponseBody) SetRequestId(v string) *AddDingtalkOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type AddDingtalkOpenPlatformConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDingtalkOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDingtalkOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDingtalkOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *AddDingtalkOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *AddDingtalkOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *AddDingtalkOpenPlatformConfigResponse) SetBody(v *AddDingtalkOpenPlatformConfigResponseBody) *AddDingtalkOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type AttachWebHostingCertificateRequest struct {
	CertName          *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType          *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	Domain            *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PrivateKey        *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	SpaceId           *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s AttachWebHostingCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachWebHostingCertificateRequest) GoString() string {
	return s.String()
}

func (s *AttachWebHostingCertificateRequest) SetCertName(v string) *AttachWebHostingCertificateRequest {
	s.CertName = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetCertType(v string) *AttachWebHostingCertificateRequest {
	s.CertType = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetDomain(v string) *AttachWebHostingCertificateRequest {
	s.Domain = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetPrivateKey(v string) *AttachWebHostingCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetServerCertificate(v string) *AttachWebHostingCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetSpaceId(v string) *AttachWebHostingCertificateRequest {
	s.SpaceId = &v
	return s
}

type AttachWebHostingCertificateResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachWebHostingCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachWebHostingCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *AttachWebHostingCertificateResponseBody) SetData(v bool) *AttachWebHostingCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *AttachWebHostingCertificateResponseBody) SetRequestId(v string) *AttachWebHostingCertificateResponseBody {
	s.RequestId = &v
	return s
}

type AttachWebHostingCertificateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachWebHostingCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachWebHostingCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachWebHostingCertificateResponse) GoString() string {
	return s.String()
}

func (s *AttachWebHostingCertificateResponse) SetHeaders(v map[string]*string) *AttachWebHostingCertificateResponse {
	s.Headers = v
	return s
}

func (s *AttachWebHostingCertificateResponse) SetBody(v *AttachWebHostingCertificateResponseBody) *AttachWebHostingCertificateResponse {
	s.Body = v
	return s
}

type BatchDeleteWebHostingFilesRequest struct {
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	SpaceId   *string   `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s BatchDeleteWebHostingFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteWebHostingFilesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteWebHostingFilesRequest) SetFilePaths(v []*string) *BatchDeleteWebHostingFilesRequest {
	s.FilePaths = v
	return s
}

func (s *BatchDeleteWebHostingFilesRequest) SetSpaceId(v string) *BatchDeleteWebHostingFilesRequest {
	s.SpaceId = &v
	return s
}

type BatchDeleteWebHostingFilesResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteWebHostingFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteWebHostingFilesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteWebHostingFilesResponseBody) SetData(v bool) *BatchDeleteWebHostingFilesResponseBody {
	s.Data = &v
	return s
}

func (s *BatchDeleteWebHostingFilesResponseBody) SetRequestId(v string) *BatchDeleteWebHostingFilesResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteWebHostingFilesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteWebHostingFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteWebHostingFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteWebHostingFilesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteWebHostingFilesResponse) SetHeaders(v map[string]*string) *BatchDeleteWebHostingFilesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteWebHostingFilesResponse) SetBody(v *BatchDeleteWebHostingFilesResponseBody) *BatchDeleteWebHostingFilesResponse {
	s.Body = v
	return s
}

type BindWebHostingCustomDomainRequest struct {
	CustomDomain *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s BindWebHostingCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BindWebHostingCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *BindWebHostingCustomDomainRequest) SetCustomDomain(v string) *BindWebHostingCustomDomainRequest {
	s.CustomDomain = &v
	return s
}

func (s *BindWebHostingCustomDomainRequest) SetSpaceId(v string) *BindWebHostingCustomDomainRequest {
	s.SpaceId = &v
	return s
}

type BindWebHostingCustomDomainResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindWebHostingCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindWebHostingCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BindWebHostingCustomDomainResponseBody) SetData(v bool) *BindWebHostingCustomDomainResponseBody {
	s.Data = &v
	return s
}

func (s *BindWebHostingCustomDomainResponseBody) SetRequestId(v string) *BindWebHostingCustomDomainResponseBody {
	s.RequestId = &v
	return s
}

type BindWebHostingCustomDomainResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindWebHostingCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindWebHostingCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BindWebHostingCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *BindWebHostingCustomDomainResponse) SetHeaders(v map[string]*string) *BindWebHostingCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *BindWebHostingCustomDomainResponse) SetBody(v *BindWebHostingCustomDomainResponseBody) *BindWebHostingCustomDomainResponse {
	s.Body = v
	return s
}

type CheckMpServerlessRoleExistsRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CheckMpServerlessRoleExistsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMpServerlessRoleExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckMpServerlessRoleExistsRequest) SetRoleName(v string) *CheckMpServerlessRoleExistsRequest {
	s.RoleName = &v
	return s
}

type CheckMpServerlessRoleExistsResponseBody struct {
	Exists    *bool   `json:"Exists,omitempty" xml:"Exists,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMpServerlessRoleExistsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMpServerlessRoleExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMpServerlessRoleExistsResponseBody) SetExists(v bool) *CheckMpServerlessRoleExistsResponseBody {
	s.Exists = &v
	return s
}

func (s *CheckMpServerlessRoleExistsResponseBody) SetRequestId(v string) *CheckMpServerlessRoleExistsResponseBody {
	s.RequestId = &v
	return s
}

type CheckMpServerlessRoleExistsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckMpServerlessRoleExistsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMpServerlessRoleExistsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMpServerlessRoleExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckMpServerlessRoleExistsResponse) SetHeaders(v map[string]*string) *CheckMpServerlessRoleExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckMpServerlessRoleExistsResponse) SetBody(v *CheckMpServerlessRoleExistsResponseBody) *CheckMpServerlessRoleExistsResponse {
	s.Body = v
	return s
}

type CreateDBExportTaskRequest struct {
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	Fields     *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	FileType   *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateDBExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDBExportTaskRequest) SetCollection(v string) *CreateDBExportTaskRequest {
	s.Collection = &v
	return s
}

func (s *CreateDBExportTaskRequest) SetFields(v string) *CreateDBExportTaskRequest {
	s.Fields = &v
	return s
}

func (s *CreateDBExportTaskRequest) SetFileType(v string) *CreateDBExportTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateDBExportTaskRequest) SetSpaceId(v string) *CreateDBExportTaskRequest {
	s.SpaceId = &v
	return s
}

type CreateDBExportTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBExportTaskResponseBody) SetRequestId(v string) *CreateDBExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBExportTaskResponseBody) SetTaskId(v string) *CreateDBExportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDBExportTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDBExportTaskResponse) SetHeaders(v map[string]*string) *CreateDBExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDBExportTaskResponse) SetBody(v *CreateDBExportTaskResponseBody) *CreateDBExportTaskResponse {
	s.Body = v
	return s
}

type CreateDBImportTaskRequest struct {
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	FileType   *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateDBImportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDBImportTaskRequest) SetCollection(v string) *CreateDBImportTaskRequest {
	s.Collection = &v
	return s
}

func (s *CreateDBImportTaskRequest) SetFileType(v string) *CreateDBImportTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateDBImportTaskRequest) SetMode(v string) *CreateDBImportTaskRequest {
	s.Mode = &v
	return s
}

func (s *CreateDBImportTaskRequest) SetSpaceId(v string) *CreateDBImportTaskRequest {
	s.SpaceId = &v
	return s
}

type CreateDBImportTaskResponseBody struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileKey     *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBImportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBImportTaskResponseBody) SetAccessKeyId(v string) *CreateDBImportTaskResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetExpireTime(v string) *CreateDBImportTaskResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetFileKey(v string) *CreateDBImportTaskResponseBody {
	s.FileKey = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetHost(v string) *CreateDBImportTaskResponseBody {
	s.Host = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetPolicy(v string) *CreateDBImportTaskResponseBody {
	s.Policy = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetRequestId(v string) *CreateDBImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetSignature(v string) *CreateDBImportTaskResponseBody {
	s.Signature = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetTaskId(v string) *CreateDBImportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDBImportTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBImportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBImportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDBImportTaskResponse) SetHeaders(v map[string]*string) *CreateDBImportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDBImportTaskResponse) SetBody(v *CreateDBImportTaskResponseBody) *CreateDBImportTaskResponse {
	s.Body = v
	return s
}

type CreateDBRestoreTaskRequest struct {
	BackupId          *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	NewCollections    *string `json:"NewCollections,omitempty" xml:"NewCollections,omitempty"`
	OriginCollections *string `json:"OriginCollections,omitempty" xml:"OriginCollections,omitempty"`
	SpaceId           *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateDBRestoreTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRestoreTaskRequest) SetBackupId(v string) *CreateDBRestoreTaskRequest {
	s.BackupId = &v
	return s
}

func (s *CreateDBRestoreTaskRequest) SetNewCollections(v string) *CreateDBRestoreTaskRequest {
	s.NewCollections = &v
	return s
}

func (s *CreateDBRestoreTaskRequest) SetOriginCollections(v string) *CreateDBRestoreTaskRequest {
	s.OriginCollections = &v
	return s
}

func (s *CreateDBRestoreTaskRequest) SetSpaceId(v string) *CreateDBRestoreTaskRequest {
	s.SpaceId = &v
	return s
}

type CreateDBRestoreTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBRestoreTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBRestoreTaskResponseBody) SetRequestId(v string) *CreateDBRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBRestoreTaskResponseBody) SetTaskId(v string) *CreateDBRestoreTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDBRestoreTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBRestoreTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBRestoreTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRestoreTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDBRestoreTaskResponse) SetHeaders(v map[string]*string) *CreateDBRestoreTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDBRestoreTaskResponse) SetBody(v *CreateDBRestoreTaskResponseBody) *CreateDBRestoreTaskResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Runtime *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetDesc(v string) *CreateFunctionRequest {
	s.Desc = &v
	return s
}

func (s *CreateFunctionRequest) SetName(v string) *CreateFunctionRequest {
	s.Name = &v
	return s
}

func (s *CreateFunctionRequest) SetRuntime(v string) *CreateFunctionRequest {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionRequest) SetSpaceId(v string) *CreateFunctionRequest {
	s.SpaceId = &v
	return s
}

type CreateFunctionResponseBody struct {
	CreatedAt  *string                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Desc       *string                         `json:"Desc,omitempty" xml:"Desc,omitempty"`
	ModifiedAt *string                         `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Name       *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spec       *CreateFunctionResponseBodySpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetCreatedAt(v string) *CreateFunctionResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateFunctionResponseBody) SetDesc(v string) *CreateFunctionResponseBody {
	s.Desc = &v
	return s
}

func (s *CreateFunctionResponseBody) SetModifiedAt(v string) *CreateFunctionResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *CreateFunctionResponseBody) SetName(v string) *CreateFunctionResponseBody {
	s.Name = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionResponseBody) SetSpec(v *CreateFunctionResponseBodySpec) *CreateFunctionResponseBody {
	s.Spec = v
	return s
}

type CreateFunctionResponseBodySpec struct {
	InstanceConcurrency *string `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateFunctionResponseBodySpec) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBodySpec) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBodySpec) SetInstanceConcurrency(v string) *CreateFunctionResponseBodySpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateFunctionResponseBodySpec) SetMemory(v string) *CreateFunctionResponseBodySpec {
	s.Memory = &v
	return s
}

func (s *CreateFunctionResponseBodySpec) SetRuntime(v string) *CreateFunctionResponseBodySpec {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionResponseBodySpec) SetTimeout(v string) *CreateFunctionResponseBodySpec {
	s.Timeout = &v
	return s
}

type CreateFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateFunctionDeploymentRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateFunctionDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionDeploymentRequest) SetName(v string) *CreateFunctionDeploymentRequest {
	s.Name = &v
	return s
}

func (s *CreateFunctionDeploymentRequest) SetSpaceId(v string) *CreateFunctionDeploymentRequest {
	s.SpaceId = &v
	return s
}

type CreateFunctionDeploymentResponseBody struct {
	DeploymentId    *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadSignedUrl *string `json:"UploadSignedUrl,omitempty" xml:"UploadSignedUrl,omitempty"`
}

func (s CreateFunctionDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionDeploymentResponseBody) SetDeploymentId(v string) *CreateFunctionDeploymentResponseBody {
	s.DeploymentId = &v
	return s
}

func (s *CreateFunctionDeploymentResponseBody) SetRequestId(v string) *CreateFunctionDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionDeploymentResponseBody) SetUploadSignedUrl(v string) *CreateFunctionDeploymentResponseBody {
	s.UploadSignedUrl = &v
	return s
}

type CreateFunctionDeploymentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionDeploymentResponse) SetHeaders(v map[string]*string) *CreateFunctionDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionDeploymentResponse) SetBody(v *CreateFunctionDeploymentResponseBody) *CreateFunctionDeploymentResponse {
	s.Body = v
	return s
}

type CreateSpaceRequest struct {
	Desc        *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkspaceId *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateSpaceRequest) SetDesc(v string) *CreateSpaceRequest {
	s.Desc = &v
	return s
}

func (s *CreateSpaceRequest) SetName(v string) *CreateSpaceRequest {
	s.Name = &v
	return s
}

func (s *CreateSpaceRequest) SetWorkspaceId(v int64) *CreateSpaceRequest {
	s.WorkspaceId = &v
	return s
}

type CreateSpaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponseBody) SetRequestId(v string) *CreateSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSpaceResponseBody) SetSpaceId(v string) *CreateSpaceResponseBody {
	s.SpaceId = &v
	return s
}

type CreateSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponse) SetHeaders(v map[string]*string) *CreateSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateSpaceResponse) SetBody(v *CreateSpaceResponseBody) *CreateSpaceResponse {
	s.Body = v
	return s
}

type DeleteAntOpenPlatformConfigRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteAntOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntOpenPlatformConfigRequest) SetAppId(v string) *DeleteAntOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAntOpenPlatformConfigRequest) SetSpaceId(v string) *DeleteAntOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type DeleteAntOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAntOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntOpenPlatformConfigResponseBody) SetRequestId(v string) *DeleteAntOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAntOpenPlatformConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAntOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAntOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *DeleteAntOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntOpenPlatformConfigResponse) SetBody(v *DeleteAntOpenPlatformConfigResponseBody) *DeleteAntOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type DeleteCorsDomainRequest struct {
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteCorsDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorsDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteCorsDomainRequest) SetDomainId(v string) *DeleteCorsDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteCorsDomainRequest) SetSpaceId(v string) *DeleteCorsDomainRequest {
	s.SpaceId = &v
	return s
}

type DeleteCorsDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCorsDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorsDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCorsDomainResponseBody) SetRequestId(v string) *DeleteCorsDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCorsDomainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCorsDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCorsDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorsDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCorsDomainResponse) SetHeaders(v map[string]*string) *DeleteCorsDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteCorsDomainResponse) SetBody(v *DeleteCorsDomainResponseBody) *DeleteCorsDomainResponse {
	s.Body = v
	return s
}

type DeleteDBCollectionRequest struct {
	Body    *string `json:"Body,omitempty" xml:"Body,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteDBCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBCollectionRequest) SetBody(v string) *DeleteDBCollectionRequest {
	s.Body = &v
	return s
}

func (s *DeleteDBCollectionRequest) SetSpaceId(v string) *DeleteDBCollectionRequest {
	s.SpaceId = &v
	return s
}

type DeleteDBCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBCollectionResponseBody) SetRequestId(v string) *DeleteDBCollectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBCollectionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBCollectionResponse) SetHeaders(v map[string]*string) *DeleteDBCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBCollectionResponse) SetBody(v *DeleteDBCollectionResponseBody) *DeleteDBCollectionResponse {
	s.Body = v
	return s
}

type DeleteDingtalkOpenPlatformConfigRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteDingtalkOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDingtalkOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDingtalkOpenPlatformConfigRequest) SetAppId(v string) *DeleteDingtalkOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *DeleteDingtalkOpenPlatformConfigRequest) SetSpaceId(v string) *DeleteDingtalkOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type DeleteDingtalkOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDingtalkOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDingtalkOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDingtalkOpenPlatformConfigResponseBody) SetRequestId(v string) *DeleteDingtalkOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDingtalkOpenPlatformConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDingtalkOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDingtalkOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDingtalkOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDingtalkOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *DeleteDingtalkOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDingtalkOpenPlatformConfigResponse) SetBody(v *DeleteDingtalkOpenPlatformConfigResponseBody) *DeleteDingtalkOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetId(v string) *DeleteFileRequest {
	s.Id = &v
	return s
}

func (s *DeleteFileRequest) SetSpaceId(v string) *DeleteFileRequest {
	s.SpaceId = &v
	return s
}

type DeleteFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DeleteFunctionRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) SetName(v string) *DeleteFunctionRequest {
	s.Name = &v
	return s
}

func (s *DeleteFunctionRequest) SetSpaceId(v string) *DeleteFunctionRequest {
	s.SpaceId = &v
	return s
}

type DeleteFunctionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetBody(v *DeleteFunctionResponseBody) *DeleteFunctionResponse {
	s.Body = v
	return s
}

type DeleteSpaceRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpaceRequest) SetSpaceId(v string) *DeleteSpaceRequest {
	s.SpaceId = &v
	return s
}

type DeleteSpaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponseBody) SetRequestId(v string) *DeleteSpaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponse) SetHeaders(v map[string]*string) *DeleteSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpaceResponse) SetBody(v *DeleteSpaceResponseBody) *DeleteSpaceResponse {
	s.Body = v
	return s
}

type DeleteWebHostingCertificateRequest struct {
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteWebHostingCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingCertificateRequest) SetDomain(v string) *DeleteWebHostingCertificateRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebHostingCertificateRequest) SetSpaceId(v string) *DeleteWebHostingCertificateRequest {
	s.SpaceId = &v
	return s
}

type DeleteWebHostingCertificateResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebHostingCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingCertificateResponseBody) SetData(v bool) *DeleteWebHostingCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWebHostingCertificateResponseBody) SetRequestId(v string) *DeleteWebHostingCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebHostingCertificateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebHostingCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebHostingCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingCertificateResponse) SetHeaders(v map[string]*string) *DeleteWebHostingCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebHostingCertificateResponse) SetBody(v *DeleteWebHostingCertificateResponseBody) *DeleteWebHostingCertificateResponse {
	s.Body = v
	return s
}

type DeleteWebHostingFileRequest struct {
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteWebHostingFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingFileRequest) SetFilePath(v string) *DeleteWebHostingFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteWebHostingFileRequest) SetSpaceId(v string) *DeleteWebHostingFileRequest {
	s.SpaceId = &v
	return s
}

type DeleteWebHostingFileResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebHostingFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingFileResponseBody) SetData(v bool) *DeleteWebHostingFileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWebHostingFileResponseBody) SetRequestId(v string) *DeleteWebHostingFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebHostingFileResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebHostingFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebHostingFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingFileResponse) SetHeaders(v map[string]*string) *DeleteWebHostingFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebHostingFileResponse) SetBody(v *DeleteWebHostingFileResponseBody) *DeleteWebHostingFileResponse {
	s.Body = v
	return s
}

type DeleteWechatOpenPlatformConfigRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteWechatOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWechatOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWechatOpenPlatformConfigRequest) SetAppId(v string) *DeleteWechatOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *DeleteWechatOpenPlatformConfigRequest) SetSpaceId(v string) *DeleteWechatOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type DeleteWechatOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWechatOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWechatOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWechatOpenPlatformConfigResponseBody) SetRequestId(v string) *DeleteWechatOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWechatOpenPlatformConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWechatOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWechatOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWechatOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWechatOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *DeleteWechatOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWechatOpenPlatformConfigResponse) SetBody(v *DeleteWechatOpenPlatformConfigResponseBody) *DeleteWechatOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type DeployFunctionRequest struct {
	DeploymentId *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeployFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeployFunctionRequest) SetDeploymentId(v string) *DeployFunctionRequest {
	s.DeploymentId = &v
	return s
}

func (s *DeployFunctionRequest) SetSpaceId(v string) *DeployFunctionRequest {
	s.SpaceId = &v
	return s
}

type DeployFunctionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeployFunctionResponseBody) SetRequestId(v string) *DeployFunctionResponseBody {
	s.RequestId = &v
	return s
}

type DeployFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeployFunctionResponse) SetHeaders(v map[string]*string) *DeployFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeployFunctionResponse) SetBody(v *DeployFunctionResponseBody) *DeployFunctionResponse {
	s.Body = v
	return s
}

type DescribeFCOpenStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFCOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFCOpenStatusResponseBody) SetRequestId(v string) *DescribeFCOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFCOpenStatusResponseBody) SetStatus(v string) *DescribeFCOpenStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeFCOpenStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFCOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFCOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeFCOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeFCOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeFCOpenStatusResponse) SetBody(v *DescribeFCOpenStatusResponseBody) *DescribeFCOpenStatusResponse {
	s.Body = v
	return s
}

type DescribeFileUploadSignedUrlRequest struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Filename    *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeFileUploadSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileUploadSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignedUrlRequest) SetContentType(v string) *DescribeFileUploadSignedUrlRequest {
	s.ContentType = &v
	return s
}

func (s *DescribeFileUploadSignedUrlRequest) SetFilename(v string) *DescribeFileUploadSignedUrlRequest {
	s.Filename = &v
	return s
}

func (s *DescribeFileUploadSignedUrlRequest) SetSize(v int64) *DescribeFileUploadSignedUrlRequest {
	s.Size = &v
	return s
}

func (s *DescribeFileUploadSignedUrlRequest) SetSpaceId(v string) *DescribeFileUploadSignedUrlRequest {
	s.SpaceId = &v
	return s
}

type DescribeFileUploadSignedUrlResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignUrl   *string `json:"SignUrl,omitempty" xml:"SignUrl,omitempty"`
}

func (s DescribeFileUploadSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileUploadSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignedUrlResponseBody) SetId(v string) *DescribeFileUploadSignedUrlResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFileUploadSignedUrlResponseBody) SetRequestId(v string) *DescribeFileUploadSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileUploadSignedUrlResponseBody) SetSignUrl(v string) *DescribeFileUploadSignedUrlResponseBody {
	s.SignUrl = &v
	return s
}

type DescribeFileUploadSignedUrlResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFileUploadSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileUploadSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileUploadSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignedUrlResponse) SetHeaders(v map[string]*string) *DescribeFileUploadSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileUploadSignedUrlResponse) SetBody(v *DescribeFileUploadSignedUrlResponseBody) *DescribeFileUploadSignedUrlResponse {
	s.Body = v
	return s
}

type DescribeFunctionRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionRequest) GoString() string {
	return s.String()
}

func (s *DescribeFunctionRequest) SetName(v string) *DescribeFunctionRequest {
	s.Name = &v
	return s
}

func (s *DescribeFunctionRequest) SetSpaceId(v string) *DescribeFunctionRequest {
	s.SpaceId = &v
	return s
}

type DescribeFunctionResponseBody struct {
	Deployment *DescribeFunctionResponseBodyDeployment `json:"Deployment,omitempty" xml:"Deployment,omitempty" type:"Struct"`
	Function   *DescribeFunctionResponseBodyFunction   `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBody) SetDeployment(v *DescribeFunctionResponseBodyDeployment) *DescribeFunctionResponseBody {
	s.Deployment = v
	return s
}

func (s *DescribeFunctionResponseBody) SetFunction(v *DescribeFunctionResponseBodyFunction) *DescribeFunctionResponseBody {
	s.Function = v
	return s
}

func (s *DescribeFunctionResponseBody) SetRequestId(v string) *DescribeFunctionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFunctionResponseBodyDeployment struct {
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DeploymentId      *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	DownloadSignedUrl *string `json:"DownloadSignedUrl,omitempty" xml:"DownloadSignedUrl,omitempty"`
	ModifiedAt        *string `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	VersionNo         *string `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s DescribeFunctionResponseBodyDeployment) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBodyDeployment) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBodyDeployment) SetCreatedAt(v string) *DescribeFunctionResponseBodyDeployment {
	s.CreatedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetDeploymentId(v string) *DescribeFunctionResponseBodyDeployment {
	s.DeploymentId = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetDownloadSignedUrl(v string) *DescribeFunctionResponseBodyDeployment {
	s.DownloadSignedUrl = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetModifiedAt(v string) *DescribeFunctionResponseBodyDeployment {
	s.ModifiedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetVersionNo(v string) *DescribeFunctionResponseBodyDeployment {
	s.VersionNo = &v
	return s
}

type DescribeFunctionResponseBodyFunction struct {
	CreatedAt           *string                                   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Desc                *string                                   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	HttpTriggerPath     *string                                   `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	ModifiedAt          *string                                   `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Name                *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Spec                *DescribeFunctionResponseBodyFunctionSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	TimingTriggerConfig *string                                   `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
}

func (s DescribeFunctionResponseBodyFunction) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBodyFunction) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBodyFunction) SetCreatedAt(v string) *DescribeFunctionResponseBodyFunction {
	s.CreatedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetDesc(v string) *DescribeFunctionResponseBodyFunction {
	s.Desc = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetHttpTriggerPath(v string) *DescribeFunctionResponseBodyFunction {
	s.HttpTriggerPath = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetModifiedAt(v string) *DescribeFunctionResponseBodyFunction {
	s.ModifiedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetName(v string) *DescribeFunctionResponseBodyFunction {
	s.Name = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetSpec(v *DescribeFunctionResponseBodyFunctionSpec) *DescribeFunctionResponseBodyFunction {
	s.Spec = v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetTimingTriggerConfig(v string) *DescribeFunctionResponseBodyFunction {
	s.TimingTriggerConfig = &v
	return s
}

type DescribeFunctionResponseBodyFunctionSpec struct {
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DescribeFunctionResponseBodyFunctionSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBodyFunctionSpec) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetInstanceConcurrency(v int32) *DescribeFunctionResponseBodyFunctionSpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetMemory(v string) *DescribeFunctionResponseBodyFunctionSpec {
	s.Memory = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetRuntime(v string) *DescribeFunctionResponseBodyFunctionSpec {
	s.Runtime = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetTimeout(v string) *DescribeFunctionResponseBodyFunctionSpec {
	s.Timeout = &v
	return s
}

type DescribeFunctionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponse) SetHeaders(v map[string]*string) *DescribeFunctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeFunctionResponse) SetBody(v *DescribeFunctionResponseBody) *DescribeFunctionResponse {
	s.Body = v
	return s
}

type DescribeHttpTriggerConfigRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeHttpTriggerConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpTriggerConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpTriggerConfigRequest) SetSpaceId(v string) *DescribeHttpTriggerConfigRequest {
	s.SpaceId = &v
	return s
}

type DescribeHttpTriggerConfigResponseBody struct {
	CustomDomain                *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	CustomDomainCertificateInfo *string `json:"CustomDomainCertificateInfo,omitempty" xml:"CustomDomainCertificateInfo,omitempty"`
	CustomDomainCname           *string `json:"CustomDomainCname,omitempty" xml:"CustomDomainCname,omitempty"`
	DefaultEndpoint             *string `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty"`
	EnableService               *bool   `json:"EnableService,omitempty" xml:"EnableService,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHttpTriggerConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpTriggerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpTriggerConfigResponseBody) SetCustomDomain(v string) *DescribeHttpTriggerConfigResponseBody {
	s.CustomDomain = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetCustomDomainCertificateInfo(v string) *DescribeHttpTriggerConfigResponseBody {
	s.CustomDomainCertificateInfo = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetCustomDomainCname(v string) *DescribeHttpTriggerConfigResponseBody {
	s.CustomDomainCname = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetDefaultEndpoint(v string) *DescribeHttpTriggerConfigResponseBody {
	s.DefaultEndpoint = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetEnableService(v bool) *DescribeHttpTriggerConfigResponseBody {
	s.EnableService = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetRequestId(v string) *DescribeHttpTriggerConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHttpTriggerConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHttpTriggerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHttpTriggerConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpTriggerConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpTriggerConfigResponse) SetHeaders(v map[string]*string) *DescribeHttpTriggerConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpTriggerConfigResponse) SetBody(v *DescribeHttpTriggerConfigResponseBody) *DescribeHttpTriggerConfigResponse {
	s.Body = v
	return s
}

type DescribeResourceQuotaRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeResourceQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceQuotaRequest) SetSpaceId(v string) *DescribeResourceQuotaRequest {
	s.SpaceId = &v
	return s
}

type DescribeResourceQuotaResponseBody struct {
	CloudStorageDataSizeQuota *float64 `json:"CloudStorageDataSizeQuota,omitempty" xml:"CloudStorageDataSizeQuota,omitempty"`
	RequestId                 *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StaticWebDataSizeQuota    *float64 `json:"StaticWebDataSizeQuota,omitempty" xml:"StaticWebDataSizeQuota,omitempty"`
}

func (s DescribeResourceQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceQuotaResponseBody) SetCloudStorageDataSizeQuota(v float64) *DescribeResourceQuotaResponseBody {
	s.CloudStorageDataSizeQuota = &v
	return s
}

func (s *DescribeResourceQuotaResponseBody) SetRequestId(v string) *DescribeResourceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceQuotaResponseBody) SetStaticWebDataSizeQuota(v float64) *DescribeResourceQuotaResponseBody {
	s.StaticWebDataSizeQuota = &v
	return s
}

type DescribeResourceQuotaResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceQuotaResponse) SetHeaders(v map[string]*string) *DescribeResourceQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceQuotaResponse) SetBody(v *DescribeResourceQuotaResponseBody) *DescribeResourceQuotaResponse {
	s.Body = v
	return s
}

type DescribeResourceUsageRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Format     *string `json:"Format,omitempty" xml:"Format,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageRequest) SetEndTime(v string) *DescribeResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetFormat(v string) *DescribeResourceUsageRequest {
	s.Format = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetPageNumber(v int64) *DescribeResourceUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetPageSize(v int64) *DescribeResourceUsageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetSpaceId(v string) *DescribeResourceUsageRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeResourceUsageRequest) SetStartTime(v string) *DescribeResourceUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeResourceUsageResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DataList       []*DescribeResourceUsageResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	HttpStatusCode *string                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Paginator      *DescribeResourceUsageResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBody) SetCode(v string) *DescribeResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDataList(v []*DescribeResourceUsageResponseBodyDataList) *DescribeResourceUsageResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetHttpStatusCode(v string) *DescribeResourceUsageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetMessage(v string) *DescribeResourceUsageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetPaginator(v *DescribeResourceUsageResponseBodyPaginator) *DescribeResourceUsageResponseBody {
	s.Paginator = v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetRequestId(v string) *DescribeResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetSuccess(v bool) *DescribeResourceUsageResponseBody {
	s.Success = &v
	return s
}

type DescribeResourceUsageResponseBodyDataList struct {
	CloudDB       *DescribeResourceUsageResponseBodyDataListCloudDB       `json:"CloudDB,omitempty" xml:"CloudDB,omitempty" type:"Struct"`
	CloudFunction *DescribeResourceUsageResponseBodyDataListCloudFunction `json:"CloudFunction,omitempty" xml:"CloudFunction,omitempty" type:"Struct"`
	CloudStorage  *DescribeResourceUsageResponseBodyDataListCloudStorage  `json:"CloudStorage,omitempty" xml:"CloudStorage,omitempty" type:"Struct"`
	EndTime       *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime     *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StaticWeb     *DescribeResourceUsageResponseBodyDataListStaticWeb     `json:"StaticWeb,omitempty" xml:"StaticWeb,omitempty" type:"Struct"`
}

func (s DescribeResourceUsageResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBodyDataList) SetCloudDB(v *DescribeResourceUsageResponseBodyDataListCloudDB) *DescribeResourceUsageResponseBodyDataList {
	s.CloudDB = v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataList) SetCloudFunction(v *DescribeResourceUsageResponseBodyDataListCloudFunction) *DescribeResourceUsageResponseBodyDataList {
	s.CloudFunction = v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataList) SetCloudStorage(v *DescribeResourceUsageResponseBodyDataListCloudStorage) *DescribeResourceUsageResponseBodyDataList {
	s.CloudStorage = v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataList) SetEndTime(v string) *DescribeResourceUsageResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataList) SetStartTime(v string) *DescribeResourceUsageResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataList) SetStaticWeb(v *DescribeResourceUsageResponseBodyDataListStaticWeb) *DescribeResourceUsageResponseBodyDataList {
	s.StaticWeb = v
	return s
}

type DescribeResourceUsageResponseBodyDataListCloudDB struct {
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Read     *int64 `json:"Read,omitempty" xml:"Read,omitempty"`
	Write    *int64 `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s DescribeResourceUsageResponseBodyDataListCloudDB) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBodyDataListCloudDB) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBodyDataListCloudDB) SetDataSize(v int64) *DescribeResourceUsageResponseBodyDataListCloudDB {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudDB) SetRead(v int64) *DescribeResourceUsageResponseBodyDataListCloudDB {
	s.Read = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudDB) SetWrite(v int64) *DescribeResourceUsageResponseBodyDataListCloudDB {
	s.Write = &v
	return s
}

type DescribeResourceUsageResponseBodyDataListCloudFunction struct {
	Compute *int64 `json:"Compute,omitempty" xml:"Compute,omitempty"`
	Count   *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	Traffic *int64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeResourceUsageResponseBodyDataListCloudFunction) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBodyDataListCloudFunction) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBodyDataListCloudFunction) SetCompute(v int64) *DescribeResourceUsageResponseBodyDataListCloudFunction {
	s.Compute = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudFunction) SetCount(v int64) *DescribeResourceUsageResponseBodyDataListCloudFunction {
	s.Count = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudFunction) SetTraffic(v int64) *DescribeResourceUsageResponseBodyDataListCloudFunction {
	s.Traffic = &v
	return s
}

type DescribeResourceUsageResponseBodyDataListCloudStorage struct {
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Download *int64 `json:"Download,omitempty" xml:"Download,omitempty"`
	Traffic  *int64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
	Upload   *int64 `json:"Upload,omitempty" xml:"Upload,omitempty"`
}

func (s DescribeResourceUsageResponseBodyDataListCloudStorage) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBodyDataListCloudStorage) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBodyDataListCloudStorage) SetDataSize(v int64) *DescribeResourceUsageResponseBodyDataListCloudStorage {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudStorage) SetDownload(v int64) *DescribeResourceUsageResponseBodyDataListCloudStorage {
	s.Download = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudStorage) SetTraffic(v int64) *DescribeResourceUsageResponseBodyDataListCloudStorage {
	s.Traffic = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListCloudStorage) SetUpload(v int64) *DescribeResourceUsageResponseBodyDataListCloudStorage {
	s.Upload = &v
	return s
}

type DescribeResourceUsageResponseBodyDataListStaticWeb struct {
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Traffic  *int64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeResourceUsageResponseBodyDataListStaticWeb) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBodyDataListStaticWeb) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBodyDataListStaticWeb) SetDataSize(v int64) *DescribeResourceUsageResponseBodyDataListStaticWeb {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyDataListStaticWeb) SetTraffic(v int64) *DescribeResourceUsageResponseBodyDataListStaticWeb {
	s.Traffic = &v
	return s
}

type DescribeResourceUsageResponseBodyPaginator struct {
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum   *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeResourceUsageResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBodyPaginator) SetPageCount(v int64) *DescribeResourceUsageResponseBodyPaginator {
	s.PageCount = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyPaginator) SetPageNum(v int64) *DescribeResourceUsageResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyPaginator) SetPageSize(v int64) *DescribeResourceUsageResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBodyPaginator) SetTotal(v int64) *DescribeResourceUsageResponseBodyPaginator {
	s.Total = &v
	return s
}

type DescribeResourceUsageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceUsageResponse) SetBody(v *DescribeResourceUsageResponseBody) *DescribeResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeServicePolicyRequest struct {
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeServicePolicyRequest) SetCollectionName(v string) *DescribeServicePolicyRequest {
	s.CollectionName = &v
	return s
}

func (s *DescribeServicePolicyRequest) SetServiceName(v string) *DescribeServicePolicyRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeServicePolicyRequest) SetSpaceId(v string) *DescribeServicePolicyRequest {
	s.SpaceId = &v
	return s
}

type DescribeServicePolicyResponseBody struct {
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServicePolicyResponseBody) SetCollectionName(v string) *DescribeServicePolicyResponseBody {
	s.CollectionName = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetPolicy(v string) *DescribeServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetPolicyName(v string) *DescribeServicePolicyResponseBody {
	s.PolicyName = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetRequestId(v string) *DescribeServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetServiceName(v string) *DescribeServicePolicyResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetSpaceId(v string) *DescribeServicePolicyResponseBody {
	s.SpaceId = &v
	return s
}

type DescribeServicePolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeServicePolicyResponse) SetHeaders(v map[string]*string) *DescribeServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeServicePolicyResponse) SetBody(v *DescribeServicePolicyResponseBody) *DescribeServicePolicyResponse {
	s.Body = v
	return s
}

type DescribeSpaceClientConfigRequest struct {
	Detail      *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	WorkspaceId *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeSpaceClientConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceClientConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpaceClientConfigRequest) SetDetail(v string) *DescribeSpaceClientConfigRequest {
	s.Detail = &v
	return s
}

func (s *DescribeSpaceClientConfigRequest) SetSpaceId(v string) *DescribeSpaceClientConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeSpaceClientConfigRequest) SetWorkspaceId(v int64) *DescribeSpaceClientConfigRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeSpaceClientConfigResponseBody struct {
	ApiKey             *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FileUploadEndpoint *string `json:"FileUploadEndpoint,omitempty" xml:"FileUploadEndpoint,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrivateKey         *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpaceId            *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeSpaceClientConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceClientConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpaceClientConfigResponseBody) SetApiKey(v string) *DescribeSpaceClientConfigResponseBody {
	s.ApiKey = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetEndpoint(v string) *DescribeSpaceClientConfigResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetFileUploadEndpoint(v string) *DescribeSpaceClientConfigResponseBody {
	s.FileUploadEndpoint = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetName(v string) *DescribeSpaceClientConfigResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetPrivateKey(v string) *DescribeSpaceClientConfigResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetRequestId(v string) *DescribeSpaceClientConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetSpaceId(v string) *DescribeSpaceClientConfigResponseBody {
	s.SpaceId = &v
	return s
}

type DescribeSpaceClientConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSpaceClientConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSpaceClientConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceClientConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpaceClientConfigResponse) SetHeaders(v map[string]*string) *DescribeSpaceClientConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpaceClientConfigResponse) SetBody(v *DescribeSpaceClientConfigResponseBody) *DescribeSpaceClientConfigResponse {
	s.Body = v
	return s
}

type DescribeWebHostingFileRequest struct {
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeWebHostingFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileRequest) SetFilePath(v string) *DescribeWebHostingFileRequest {
	s.FilePath = &v
	return s
}

func (s *DescribeWebHostingFileRequest) SetSpaceId(v string) *DescribeWebHostingFileRequest {
	s.SpaceId = &v
	return s
}

type DescribeWebHostingFileResponseBody struct {
	Data      *DescribeWebHostingFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebHostingFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileResponseBody) SetData(v *DescribeWebHostingFileResponseBodyData) *DescribeWebHostingFileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWebHostingFileResponseBody) SetRequestId(v string) *DescribeWebHostingFileResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWebHostingFileResponseBodyData struct {
	ContentType      *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ETag             *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	Exists           *bool   `json:"Exists,omitempty" xml:"Exists,omitempty"`
	FilePath         *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	LastModifiedTime *int64  `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	SignedUrl        *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeWebHostingFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileResponseBodyData) SetContentType(v string) *DescribeWebHostingFileResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetETag(v string) *DescribeWebHostingFileResponseBodyData {
	s.ETag = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetExists(v bool) *DescribeWebHostingFileResponseBodyData {
	s.Exists = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetFilePath(v string) *DescribeWebHostingFileResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetLastModifiedTime(v int64) *DescribeWebHostingFileResponseBodyData {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetSignedUrl(v string) *DescribeWebHostingFileResponseBodyData {
	s.SignedUrl = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetSize(v int64) *DescribeWebHostingFileResponseBodyData {
	s.Size = &v
	return s
}

type DescribeWebHostingFileResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebHostingFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebHostingFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileResponse) SetHeaders(v map[string]*string) *DescribeWebHostingFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebHostingFileResponse) SetBody(v *DescribeWebHostingFileResponseBody) *DescribeWebHostingFileResponse {
	s.Body = v
	return s
}

type EnableExtensionRequest struct {
	ExtensionId *string `json:"ExtensionId,omitempty" xml:"ExtensionId,omitempty"`
}

func (s EnableExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableExtensionRequest) GoString() string {
	return s.String()
}

func (s *EnableExtensionRequest) SetExtensionId(v string) *EnableExtensionRequest {
	s.ExtensionId = &v
	return s
}

type EnableExtensionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableExtensionResponseBody) SetRequestId(v string) *EnableExtensionResponseBody {
	s.RequestId = &v
	return s
}

type EnableExtensionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableExtensionResponse) GoString() string {
	return s.String()
}

func (s *EnableExtensionResponse) SetHeaders(v map[string]*string) *EnableExtensionResponse {
	s.Headers = v
	return s
}

func (s *EnableExtensionResponse) SetBody(v *EnableExtensionResponseBody) *EnableExtensionResponse {
	s.Body = v
	return s
}

type GetWebHostingCertificateDetailRequest struct {
	CustomDomain *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailRequest) SetCustomDomain(v string) *GetWebHostingCertificateDetailRequest {
	s.CustomDomain = &v
	return s
}

func (s *GetWebHostingCertificateDetailRequest) SetSpaceId(v string) *GetWebHostingCertificateDetailRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingCertificateDetailResponseBody struct {
	Data      *GetWebHostingCertificateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWebHostingCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailResponseBody) SetData(v *GetWebHostingCertificateDetailResponseBodyData) *GetWebHostingCertificateDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBody) SetRequestId(v string) *GetWebHostingCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetWebHostingCertificateDetailResponseBodyData struct {
	CertDomainName          *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpiredTime         *int64  `json:"CertExpiredTime,omitempty" xml:"CertExpiredTime,omitempty"`
	CertLife                *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
}

func (s GetWebHostingCertificateDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertDomainName(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertDomainName = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertExpiredTime(v int64) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertExpiredTime = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertLife(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertLife = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertName(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertName = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertType(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertType = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetServerCertificateStatus(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.ServerCertificateStatus = &v
	return s
}

type GetWebHostingCertificateDetailResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailResponse) SetHeaders(v map[string]*string) *GetWebHostingCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingCertificateDetailResponse) SetBody(v *GetWebHostingCertificateDetailResponseBody) *GetWebHostingCertificateDetailResponse {
	s.Body = v
	return s
}

type GetWebHostingConfigRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigRequest) SetSpaceId(v string) *GetWebHostingConfigRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingConfigResponseBody struct {
	Data      *GetWebHostingConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWebHostingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigResponseBody) SetData(v *GetWebHostingConfigResponseBodyData) *GetWebHostingConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetWebHostingConfigResponseBody) SetRequestId(v string) *GetWebHostingConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetWebHostingConfigResponseBodyData struct {
	AllowedIps      *string `json:"AllowedIps,omitempty" xml:"AllowedIps,omitempty"`
	DefaultDomain   *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	ErrorPath       *string `json:"ErrorPath,omitempty" xml:"ErrorPath,omitempty"`
	HistoryModePath *string `json:"HistoryModePath,omitempty" xml:"HistoryModePath,omitempty"`
	IndexPath       *string `json:"IndexPath,omitempty" xml:"IndexPath,omitempty"`
	SpaceId         *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigResponseBodyData) SetAllowedIps(v string) *GetWebHostingConfigResponseBodyData {
	s.AllowedIps = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetDefaultDomain(v string) *GetWebHostingConfigResponseBodyData {
	s.DefaultDomain = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetErrorPath(v string) *GetWebHostingConfigResponseBodyData {
	s.ErrorPath = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetHistoryModePath(v string) *GetWebHostingConfigResponseBodyData {
	s.HistoryModePath = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetIndexPath(v string) *GetWebHostingConfigResponseBodyData {
	s.IndexPath = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetSpaceId(v string) *GetWebHostingConfigResponseBodyData {
	s.SpaceId = &v
	return s
}

type GetWebHostingConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigResponse) SetHeaders(v map[string]*string) *GetWebHostingConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingConfigResponse) SetBody(v *GetWebHostingConfigResponseBody) *GetWebHostingConfigResponse {
	s.Body = v
	return s
}

type GetWebHostingDomainVerificationContentRequest struct {
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingDomainVerificationContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentRequest) SetDomain(v string) *GetWebHostingDomainVerificationContentRequest {
	s.Domain = &v
	return s
}

func (s *GetWebHostingDomainVerificationContentRequest) SetSpaceId(v string) *GetWebHostingDomainVerificationContentRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingDomainVerificationContentResponseBody struct {
	Data      *GetWebHostingDomainVerificationContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWebHostingDomainVerificationContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentResponseBody) SetData(v *GetWebHostingDomainVerificationContentResponseBodyData) *GetWebHostingDomainVerificationContentResponseBody {
	s.Data = v
	return s
}

func (s *GetWebHostingDomainVerificationContentResponseBody) SetRequestId(v string) *GetWebHostingDomainVerificationContentResponseBody {
	s.RequestId = &v
	return s
}

type GetWebHostingDomainVerificationContentResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetWebHostingDomainVerificationContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentResponseBodyData) SetContent(v string) *GetWebHostingDomainVerificationContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetWebHostingDomainVerificationContentResponseBodyData) SetDomain(v string) *GetWebHostingDomainVerificationContentResponseBodyData {
	s.Domain = &v
	return s
}

type GetWebHostingDomainVerificationContentResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingDomainVerificationContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingDomainVerificationContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentResponse) SetHeaders(v map[string]*string) *GetWebHostingDomainVerificationContentResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingDomainVerificationContentResponse) SetBody(v *GetWebHostingDomainVerificationContentResponseBody) *GetWebHostingDomainVerificationContentResponse {
	s.Body = v
	return s
}

type GetWebHostingStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusRequest) SetSpaceId(v string) *GetWebHostingStatusRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingStatusResponseBody struct {
	Data      *GetWebHostingStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWebHostingStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusResponseBody) SetData(v *GetWebHostingStatusResponseBodyData) *GetWebHostingStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetWebHostingStatusResponseBody) SetRequestId(v string) *GetWebHostingStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetWebHostingStatusResponseBodyData struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWebHostingStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusResponseBodyData) SetSpaceId(v string) *GetWebHostingStatusResponseBodyData {
	s.SpaceId = &v
	return s
}

func (s *GetWebHostingStatusResponseBodyData) SetStatus(v string) *GetWebHostingStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetWebHostingStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusResponse) SetHeaders(v map[string]*string) *GetWebHostingStatusResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingStatusResponse) SetBody(v *GetWebHostingStatusResponseBody) *GetWebHostingStatusResponse {
	s.Body = v
	return s
}

type GetWebHostingUploadCredentialRequest struct {
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingUploadCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialRequest) SetFilePath(v string) *GetWebHostingUploadCredentialRequest {
	s.FilePath = &v
	return s
}

func (s *GetWebHostingUploadCredentialRequest) SetSpaceId(v string) *GetWebHostingUploadCredentialRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingUploadCredentialResponseBody struct {
	Data      *GetWebHostingUploadCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWebHostingUploadCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialResponseBody) SetData(v *GetWebHostingUploadCredentialResponseBodyData) *GetWebHostingUploadCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBody) SetRequestId(v string) *GetWebHostingUploadCredentialResponseBody {
	s.RequestId = &v
	return s
}

type GetWebHostingUploadCredentialResponseBodyData struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Endpoint      *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ExpiredTime   *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FilePath      *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetWebHostingUploadCredentialResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetAccessKeyId(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetEndpoint(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetExpiredTime(v int64) *GetWebHostingUploadCredentialResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetFilePath(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetPolicy(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetSecurityToken(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetSignature(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.Signature = &v
	return s
}

type GetWebHostingUploadCredentialResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingUploadCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingUploadCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialResponse) SetHeaders(v map[string]*string) *GetWebHostingUploadCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingUploadCredentialResponse) SetBody(v *GetWebHostingUploadCredentialResponseBody) *GetWebHostingUploadCredentialResponse {
	s.Body = v
	return s
}

type ListAvailableCertificatesRequest struct {
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListAvailableCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesRequest) SetDomain(v string) *ListAvailableCertificatesRequest {
	s.Domain = &v
	return s
}

func (s *ListAvailableCertificatesRequest) SetSpaceId(v string) *ListAvailableCertificatesRequest {
	s.SpaceId = &v
	return s
}

type ListAvailableCertificatesResponseBody struct {
	Data      []*ListAvailableCertificatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesResponseBody) SetData(v []*ListAvailableCertificatesResponseBodyData) *ListAvailableCertificatesResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableCertificatesResponseBody) SetRequestId(v string) *ListAvailableCertificatesResponseBody {
	s.RequestId = &v
	return s
}

type ListAvailableCertificatesResponseBodyData struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListAvailableCertificatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesResponseBodyData) SetId(v string) *ListAvailableCertificatesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAvailableCertificatesResponseBodyData) SetName(v string) *ListAvailableCertificatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAvailableCertificatesResponseBodyData) SetStatusCode(v string) *ListAvailableCertificatesResponseBodyData {
	s.StatusCode = &v
	return s
}

type ListAvailableCertificatesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesResponse) SetHeaders(v map[string]*string) *ListAvailableCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableCertificatesResponse) SetBody(v *ListAvailableCertificatesResponseBody) *ListAvailableCertificatesResponse {
	s.Body = v
	return s
}

type ListCorsDomainsRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListCorsDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsRequest) SetSpaceId(v string) *ListCorsDomainsRequest {
	s.SpaceId = &v
	return s
}

type ListCorsDomainsResponseBody struct {
	Domains   []*ListCorsDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCorsDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsResponseBody) SetDomains(v []*ListCorsDomainsResponseBodyDomains) *ListCorsDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListCorsDomainsResponseBody) SetRequestId(v string) *ListCorsDomainsResponseBody {
	s.RequestId = &v
	return s
}

type ListCorsDomainsResponseBodyDomains struct {
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s ListCorsDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsResponseBodyDomains) SetDomain(v string) *ListCorsDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *ListCorsDomainsResponseBodyDomains) SetDomainId(v string) *ListCorsDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

type ListCorsDomainsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorsDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorsDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsResponse) SetHeaders(v map[string]*string) *ListCorsDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListCorsDomainsResponse) SetBody(v *ListCorsDomainsResponseBody) *ListCorsDomainsResponse {
	s.Body = v
	return s
}

type ListDingtalkOpenPlatformConfigsRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListDingtalkOpenPlatformConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsRequest) SetSpaceId(v string) *ListDingtalkOpenPlatformConfigsRequest {
	s.SpaceId = &v
	return s
}

type ListDingtalkOpenPlatformConfigsResponseBody struct {
	Configs   []*ListDingtalkOpenPlatformConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDingtalkOpenPlatformConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsResponseBody) SetConfigs(v []*ListDingtalkOpenPlatformConfigsResponseBodyConfigs) *ListDingtalkOpenPlatformConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBody) SetRequestId(v string) *ListDingtalkOpenPlatformConfigsResponseBody {
	s.RequestId = &v
	return s
}

type ListDingtalkOpenPlatformConfigsResponseBodyConfigs struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret  *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDingtalkOpenPlatformConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetAppId(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.AppId = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetAppSecret(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.AppSecret = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetCreateTime(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetUpdateTime(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.UpdateTime = &v
	return s
}

type ListDingtalkOpenPlatformConfigsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDingtalkOpenPlatformConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDingtalkOpenPlatformConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsResponse) SetHeaders(v map[string]*string) *ListDingtalkOpenPlatformConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponse) SetBody(v *ListDingtalkOpenPlatformConfigsResponseBody) *ListDingtalkOpenPlatformConfigsResponse {
	s.Body = v
	return s
}

type ListExtensionsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExtensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsRequest) GoString() string {
	return s.String()
}

func (s *ListExtensionsRequest) SetPageNumber(v int32) *ListExtensionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExtensionsRequest) SetPageSize(v int32) *ListExtensionsRequest {
	s.PageSize = &v
	return s
}

type ListExtensionsResponseBody struct {
	Extensions []*ListExtensionsResponseBodyExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExtensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBody) SetExtensions(v []*ListExtensionsResponseBodyExtensions) *ListExtensionsResponseBody {
	s.Extensions = v
	return s
}

func (s *ListExtensionsResponseBody) SetPageNumber(v int32) *ListExtensionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExtensionsResponseBody) SetPageSize(v int32) *ListExtensionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExtensionsResponseBody) SetRequestId(v string) *ListExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExtensionsResponseBody) SetTotalCount(v int32) *ListExtensionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExtensionsResponseBodyExtensions struct {
	Enabled                    *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExtensionDesc              *string `json:"ExtensionDesc,omitempty" xml:"ExtensionDesc,omitempty"`
	ExtensionDocumentationLink *string `json:"ExtensionDocumentationLink,omitempty" xml:"ExtensionDocumentationLink,omitempty"`
	ExtensionId                *string `json:"ExtensionId,omitempty" xml:"ExtensionId,omitempty"`
	ExtensionName              *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
}

func (s ListExtensionsResponseBodyExtensions) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsResponseBodyExtensions) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBodyExtensions) SetEnabled(v string) *ListExtensionsResponseBodyExtensions {
	s.Enabled = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionDesc(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionDesc = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionDocumentationLink(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionDocumentationLink = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionId(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionId = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionName(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionName = &v
	return s
}

type ListExtensionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExtensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsResponse) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponse) SetHeaders(v map[string]*string) *ListExtensionsResponse {
	s.Headers = v
	return s
}

func (s *ListExtensionsResponse) SetBody(v *ListExtensionsResponseBody) *ListExtensionsResponse {
	s.Body = v
	return s
}

type ListFileRequest struct {
	FileId   *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileRequest) GoString() string {
	return s.String()
}

func (s *ListFileRequest) SetFileId(v string) *ListFileRequest {
	s.FileId = &v
	return s
}

func (s *ListFileRequest) SetKeyword(v string) *ListFileRequest {
	s.Keyword = &v
	return s
}

func (s *ListFileRequest) SetPageNum(v int32) *ListFileRequest {
	s.PageNum = &v
	return s
}

func (s *ListFileRequest) SetPageSize(v int32) *ListFileRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileRequest) SetSpaceId(v string) *ListFileRequest {
	s.SpaceId = &v
	return s
}

type ListFileResponseBody struct {
	DataList  []*ListFileResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Paginator *ListFileResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileResponseBody) SetDataList(v []*ListFileResponseBodyDataList) *ListFileResponseBody {
	s.DataList = v
	return s
}

func (s *ListFileResponseBody) SetPaginator(v *ListFileResponseBodyPaginator) *ListFileResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFileResponseBody) SetRequestId(v string) *ListFileResponseBody {
	s.RequestId = &v
	return s
}

type ListFileResponseBodyDataList struct {
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size        *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListFileResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyDataList) SetGmtCreate(v string) *ListFileResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetGmtModified(v string) *ListFileResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetId(v string) *ListFileResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetName(v string) *ListFileResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetSize(v int32) *ListFileResponseBodyDataList {
	s.Size = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetType(v string) *ListFileResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetUrl(v string) *ListFileResponseBodyDataList {
	s.Url = &v
	return s
}

type ListFileResponseBodyPaginator struct {
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFileResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyPaginator) SetPageCount(v int32) *ListFileResponseBodyPaginator {
	s.PageCount = &v
	return s
}

func (s *ListFileResponseBodyPaginator) SetPageNum(v int32) *ListFileResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFileResponseBodyPaginator) SetPageSize(v int32) *ListFileResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFileResponseBodyPaginator) SetTotal(v int32) *ListFileResponseBodyPaginator {
	s.Total = &v
	return s
}

type ListFileResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponse) GoString() string {
	return s.String()
}

func (s *ListFileResponse) SetHeaders(v map[string]*string) *ListFileResponse {
	s.Headers = v
	return s
}

func (s *ListFileResponse) SetBody(v *ListFileResponseBody) *ListFileResponse {
	s.Body = v
	return s
}

type ListFunctionRequest struct {
	FilterBy *string `json:"FilterBy,omitempty" xml:"FilterBy,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionRequest) SetFilterBy(v string) *ListFunctionRequest {
	s.FilterBy = &v
	return s
}

func (s *ListFunctionRequest) SetPageNum(v int32) *ListFunctionRequest {
	s.PageNum = &v
	return s
}

func (s *ListFunctionRequest) SetPageSize(v int32) *ListFunctionRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionRequest) SetSpaceId(v string) *ListFunctionRequest {
	s.SpaceId = &v
	return s
}

type ListFunctionResponseBody struct {
	DataList  []*ListFunctionResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Paginator *ListFunctionResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBody) SetDataList(v []*ListFunctionResponseBodyDataList) *ListFunctionResponseBody {
	s.DataList = v
	return s
}

func (s *ListFunctionResponseBody) SetPaginator(v *ListFunctionResponseBodyPaginator) *ListFunctionResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFunctionResponseBody) SetRequestId(v string) *ListFunctionResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionResponseBodyDataList struct {
	CreatedAt           *string                               `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Desc                *string                               `json:"Desc,omitempty" xml:"Desc,omitempty"`
	HttpTriggerPath     *string                               `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	ModifiedAt          *string                               `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Name                *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Spec                *ListFunctionResponseBodyDataListSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	TimingTriggerConfig *string                               `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
}

func (s ListFunctionResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBodyDataList) SetCreatedAt(v string) *ListFunctionResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetDesc(v string) *ListFunctionResponseBodyDataList {
	s.Desc = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetHttpTriggerPath(v string) *ListFunctionResponseBodyDataList {
	s.HttpTriggerPath = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetModifiedAt(v string) *ListFunctionResponseBodyDataList {
	s.ModifiedAt = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetName(v string) *ListFunctionResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetSpec(v *ListFunctionResponseBodyDataListSpec) *ListFunctionResponseBodyDataList {
	s.Spec = v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetTimingTriggerConfig(v string) *ListFunctionResponseBodyDataList {
	s.TimingTriggerConfig = &v
	return s
}

type ListFunctionResponseBodyDataListSpec struct {
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListFunctionResponseBodyDataListSpec) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBodyDataListSpec) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBodyDataListSpec) SetInstanceConcurrency(v int32) *ListFunctionResponseBodyDataListSpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *ListFunctionResponseBodyDataListSpec) SetMemory(v string) *ListFunctionResponseBodyDataListSpec {
	s.Memory = &v
	return s
}

func (s *ListFunctionResponseBodyDataListSpec) SetRuntime(v string) *ListFunctionResponseBodyDataListSpec {
	s.Runtime = &v
	return s
}

func (s *ListFunctionResponseBodyDataListSpec) SetTimeout(v string) *ListFunctionResponseBodyDataListSpec {
	s.Timeout = &v
	return s
}

type ListFunctionResponseBodyPaginator struct {
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFunctionResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBodyPaginator) SetPageCount(v int32) *ListFunctionResponseBodyPaginator {
	s.PageCount = &v
	return s
}

func (s *ListFunctionResponseBodyPaginator) SetPageNum(v int32) *ListFunctionResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFunctionResponseBodyPaginator) SetPageSize(v int32) *ListFunctionResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFunctionResponseBodyPaginator) SetTotal(v int32) *ListFunctionResponseBodyPaginator {
	s.Total = &v
	return s
}

type ListFunctionResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionResponse) SetHeaders(v map[string]*string) *ListFunctionResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionResponse) SetBody(v *ListFunctionResponseBody) *ListFunctionResponse {
	s.Body = v
	return s
}

type ListFunctionDeploymentRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentRequest) SetName(v string) *ListFunctionDeploymentRequest {
	s.Name = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetPageNum(v int32) *ListFunctionDeploymentRequest {
	s.PageNum = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetPageSize(v int32) *ListFunctionDeploymentRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetSpaceId(v string) *ListFunctionDeploymentRequest {
	s.SpaceId = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetStatus(v string) *ListFunctionDeploymentRequest {
	s.Status = &v
	return s
}

type ListFunctionDeploymentResponseBody struct {
	DataList  []*ListFunctionDeploymentResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Paginator *ListFunctionDeploymentResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBody) SetDataList(v []*ListFunctionDeploymentResponseBodyDataList) *ListFunctionDeploymentResponseBody {
	s.DataList = v
	return s
}

func (s *ListFunctionDeploymentResponseBody) SetPaginator(v *ListFunctionDeploymentResponseBodyPaginator) *ListFunctionDeploymentResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFunctionDeploymentResponseBody) SetRequestId(v string) *ListFunctionDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionDeploymentResponseBodyDataList struct {
	CreatedAt         *string                                           `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DeploymentId      *string                                           `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	DownloadSignedUrl *string                                           `json:"DownloadSignedUrl,omitempty" xml:"DownloadSignedUrl,omitempty"`
	ModifiedAt        *string                                           `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Status            *ListFunctionDeploymentResponseBodyDataListStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	VersionNo         *string                                           `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s ListFunctionDeploymentResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetCreatedAt(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetDeploymentId(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.DeploymentId = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetDownloadSignedUrl(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.DownloadSignedUrl = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetModifiedAt(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.ModifiedAt = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetStatus(v *ListFunctionDeploymentResponseBodyDataListStatus) *ListFunctionDeploymentResponseBodyDataList {
	s.Status = v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetVersionNo(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.VersionNo = &v
	return s
}

type ListFunctionDeploymentResponseBodyDataListStatus struct {
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionDeploymentResponseBodyDataListStatus) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBodyDataListStatus) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBodyDataListStatus) SetLabel(v string) *ListFunctionDeploymentResponseBodyDataListStatus {
	s.Label = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataListStatus) SetStatus(v string) *ListFunctionDeploymentResponseBodyDataListStatus {
	s.Status = &v
	return s
}

type ListFunctionDeploymentResponseBodyPaginator struct {
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFunctionDeploymentResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetPageCount(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.PageCount = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetPageNum(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetPageSize(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetTotal(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.Total = &v
	return s
}

type ListFunctionDeploymentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponse) SetHeaders(v map[string]*string) *ListFunctionDeploymentResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionDeploymentResponse) SetBody(v *ListFunctionDeploymentResponseBody) *ListFunctionDeploymentResponse {
	s.Body = v
	return s
}

type ListFunctionLogRequest struct {
	FromDate     *int64  `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ToDate       *int64  `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
}

func (s ListFunctionLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionLogRequest) SetFromDate(v int64) *ListFunctionLogRequest {
	s.FromDate = &v
	return s
}

func (s *ListFunctionLogRequest) SetLogRequestId(v string) *ListFunctionLogRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListFunctionLogRequest) SetName(v string) *ListFunctionLogRequest {
	s.Name = &v
	return s
}

func (s *ListFunctionLogRequest) SetPageNum(v int32) *ListFunctionLogRequest {
	s.PageNum = &v
	return s
}

func (s *ListFunctionLogRequest) SetPageSize(v int32) *ListFunctionLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionLogRequest) SetSpaceId(v string) *ListFunctionLogRequest {
	s.SpaceId = &v
	return s
}

func (s *ListFunctionLogRequest) SetStatus(v string) *ListFunctionLogRequest {
	s.Status = &v
	return s
}

func (s *ListFunctionLogRequest) SetToDate(v int64) *ListFunctionLogRequest {
	s.ToDate = &v
	return s
}

type ListFunctionLogResponseBody struct {
	DataList  []*ListFunctionLogResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Paginator *ListFunctionLogResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponseBody) SetDataList(v []*ListFunctionLogResponseBodyDataList) *ListFunctionLogResponseBody {
	s.DataList = v
	return s
}

func (s *ListFunctionLogResponseBody) SetPaginator(v *ListFunctionLogResponseBodyPaginator) *ListFunctionLogResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFunctionLogResponseBody) SetRequestId(v string) *ListFunctionLogResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionLogResponseBodyDataList struct {
	Contents     []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	FunctionName *string   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Levels       []*string `json:"Levels,omitempty" xml:"Levels,omitempty" type:"Repeated"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpaceId      *string   `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Timestamps   []*string `json:"Timestamps,omitempty" xml:"Timestamps,omitempty" type:"Repeated"`
}

func (s ListFunctionLogResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponseBodyDataList) SetContents(v []*string) *ListFunctionLogResponseBodyDataList {
	s.Contents = v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetFunctionName(v string) *ListFunctionLogResponseBodyDataList {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetLevels(v []*string) *ListFunctionLogResponseBodyDataList {
	s.Levels = v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetRequestId(v string) *ListFunctionLogResponseBodyDataList {
	s.RequestId = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetSpaceId(v string) *ListFunctionLogResponseBodyDataList {
	s.SpaceId = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetStatus(v string) *ListFunctionLogResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetTimestamps(v []*string) *ListFunctionLogResponseBodyDataList {
	s.Timestamps = v
	return s
}

type ListFunctionLogResponseBodyPaginator struct {
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFunctionLogResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponseBodyPaginator) SetPageCount(v int32) *ListFunctionLogResponseBodyPaginator {
	s.PageCount = &v
	return s
}

func (s *ListFunctionLogResponseBodyPaginator) SetPageNum(v int32) *ListFunctionLogResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFunctionLogResponseBodyPaginator) SetPageSize(v int32) *ListFunctionLogResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFunctionLogResponseBodyPaginator) SetTotal(v int32) *ListFunctionLogResponseBodyPaginator {
	s.Total = &v
	return s
}

type ListFunctionLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponse) SetHeaders(v map[string]*string) *ListFunctionLogResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionLogResponse) SetBody(v *ListFunctionLogResponseBody) *ListFunctionLogResponse {
	s.Body = v
	return s
}

type ListOpenPlatformConfigRequest struct {
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigRequest) SetPlatform(v string) *ListOpenPlatformConfigRequest {
	s.Platform = &v
	return s
}

func (s *ListOpenPlatformConfigRequest) SetSpaceId(v string) *ListOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type ListOpenPlatformConfigResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretList []*ListOpenPlatformConfigResponseBodySecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Repeated"`
}

func (s ListOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigResponseBody) SetRequestId(v string) *ListOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBody) SetSecretList(v []*ListOpenPlatformConfigResponseBodySecretList) *ListOpenPlatformConfigResponseBody {
	s.SecretList = v
	return s
}

type ListOpenPlatformConfigResponseBodySecretList struct {
	AppCert    *string `json:"AppCert,omitempty" xml:"AppCert,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret  *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PublicCert *string `json:"PublicCert,omitempty" xml:"PublicCert,omitempty"`
	PublicKey  *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	RootCert   *string `json:"RootCert,omitempty" xml:"RootCert,omitempty"`
	SignMode   *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListOpenPlatformConfigResponseBodySecretList) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigResponseBodySecretList) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetAppCert(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.AppCert = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetAppId(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.AppId = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetAppSecret(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.AppSecret = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPlatform(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.Platform = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPrivateKey(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.PrivateKey = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPublicCert(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.PublicCert = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPublicKey(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.PublicKey = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetRootCert(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.RootCert = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetSignMode(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.SignMode = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetSpaceId(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.SpaceId = &v
	return s
}

type ListOpenPlatformConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *ListOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *ListOpenPlatformConfigResponse) SetBody(v *ListOpenPlatformConfigResponseBody) *ListOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type ListSpaceRequest struct {
	PageNum  *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceIds []*string `json:"SpaceIds,omitempty" xml:"SpaceIds,omitempty" type:"Repeated"`
}

func (s ListSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceRequest) GoString() string {
	return s.String()
}

func (s *ListSpaceRequest) SetPageNum(v int32) *ListSpaceRequest {
	s.PageNum = &v
	return s
}

func (s *ListSpaceRequest) SetPageSize(v int32) *ListSpaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListSpaceRequest) SetSpaceIds(v []*string) *ListSpaceRequest {
	s.SpaceIds = v
	return s
}

type ListSpaceShrinkRequest struct {
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceIdsShrink *string `json:"SpaceIds,omitempty" xml:"SpaceIds,omitempty"`
}

func (s ListSpaceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSpaceShrinkRequest) SetPageNum(v int32) *ListSpaceShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListSpaceShrinkRequest) SetPageSize(v int32) *ListSpaceShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSpaceShrinkRequest) SetSpaceIdsShrink(v string) *ListSpaceShrinkRequest {
	s.SpaceIdsShrink = &v
	return s
}

type ListSpaceResponseBody struct {
	Count     *int32                         `json:"Count,omitempty" xml:"Count,omitempty"`
	GmtCreate *string                        `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spaces    []*ListSpaceResponseBodySpaces `json:"Spaces,omitempty" xml:"Spaces,omitempty" type:"Repeated"`
}

func (s ListSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpaceResponseBody) SetCount(v int32) *ListSpaceResponseBody {
	s.Count = &v
	return s
}

func (s *ListSpaceResponseBody) SetGmtCreate(v string) *ListSpaceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *ListSpaceResponseBody) SetRequestId(v string) *ListSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpaceResponseBody) SetSpaces(v []*ListSpaceResponseBodySpaces) *ListSpaceResponseBody {
	s.Spaces = v
	return s
}

type ListSpaceResponseBodySpaces struct {
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtLastAccess *int64  `json:"GmtLastAccess,omitempty" xml:"GmtLastAccess,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId       *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSpaceResponseBodySpaces) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *ListSpaceResponseBodySpaces) SetDesc(v string) *ListSpaceResponseBodySpaces {
	s.Desc = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetGmtCreate(v int64) *ListSpaceResponseBodySpaces {
	s.GmtCreate = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetGmtLastAccess(v int64) *ListSpaceResponseBodySpaces {
	s.GmtLastAccess = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetName(v string) *ListSpaceResponseBodySpaces {
	s.Name = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetSpaceId(v string) *ListSpaceResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetStatus(v string) *ListSpaceResponseBodySpaces {
	s.Status = &v
	return s
}

type ListSpaceResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceResponse) GoString() string {
	return s.String()
}

func (s *ListSpaceResponse) SetHeaders(v map[string]*string) *ListSpaceResponse {
	s.Headers = v
	return s
}

func (s *ListSpaceResponse) SetBody(v *ListSpaceResponseBody) *ListSpaceResponse {
	s.Body = v
	return s
}

type ListWebHostingCustomDomainsRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListWebHostingCustomDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsRequest) SetSpaceId(v string) *ListWebHostingCustomDomainsRequest {
	s.SpaceId = &v
	return s
}

type ListWebHostingCustomDomainsResponseBody struct {
	Data      []*ListWebHostingCustomDomainsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWebHostingCustomDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsResponseBody) SetData(v []*ListWebHostingCustomDomainsResponseBodyData) *ListWebHostingCustomDomainsResponseBody {
	s.Data = v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBody) SetRequestId(v string) *ListWebHostingCustomDomainsResponseBody {
	s.RequestId = &v
	return s
}

type ListWebHostingCustomDomainsResponseBodyData struct {
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	Cname                    *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	CreateTime               *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Domain                   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EnableCors               *bool   `json:"EnableCors,omitempty" xml:"EnableCors,omitempty"`
	ForceRedirectType        *string `json:"ForceRedirectType,omitempty" xml:"ForceRedirectType,omitempty"`
	SslProtocol              *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime               *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWebHostingCustomDomainsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetAccessControlAllowOrigin(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.AccessControlAllowOrigin = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetCname(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Cname = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetCreateTime(v int64) *ListWebHostingCustomDomainsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetDescription(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetDomain(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetEnableCors(v bool) *ListWebHostingCustomDomainsResponseBodyData {
	s.EnableCors = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetForceRedirectType(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.ForceRedirectType = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetSslProtocol(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.SslProtocol = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetStatus(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetUpdateTime(v int64) *ListWebHostingCustomDomainsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListWebHostingCustomDomainsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWebHostingCustomDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWebHostingCustomDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsResponse) SetHeaders(v map[string]*string) *ListWebHostingCustomDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListWebHostingCustomDomainsResponse) SetBody(v *ListWebHostingCustomDomainsResponseBody) *ListWebHostingCustomDomainsResponse {
	s.Body = v
	return s
}

type ListWebHostingFilesRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Prefix   *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListWebHostingFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesRequest) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesRequest) SetMarker(v string) *ListWebHostingFilesRequest {
	s.Marker = &v
	return s
}

func (s *ListWebHostingFilesRequest) SetPageSize(v int32) *ListWebHostingFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWebHostingFilesRequest) SetPrefix(v string) *ListWebHostingFilesRequest {
	s.Prefix = &v
	return s
}

func (s *ListWebHostingFilesRequest) SetSpaceId(v string) *ListWebHostingFilesRequest {
	s.SpaceId = &v
	return s
}

type ListWebHostingFilesResponseBody struct {
	Data      *ListWebHostingFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWebHostingFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponseBody) SetData(v *ListWebHostingFilesResponseBodyData) *ListWebHostingFilesResponseBody {
	s.Data = v
	return s
}

func (s *ListWebHostingFilesResponseBody) SetRequestId(v string) *ListWebHostingFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListWebHostingFilesResponseBodyData struct {
	Count           *int32                                                `json:"Count,omitempty" xml:"Count,omitempty"`
	NextMarker      *string                                               `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	WebHostingFiles []*ListWebHostingFilesResponseBodyDataWebHostingFiles `json:"WebHostingFiles,omitempty" xml:"WebHostingFiles,omitempty" type:"Repeated"`
}

func (s ListWebHostingFilesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponseBodyData) SetCount(v int32) *ListWebHostingFilesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyData) SetNextMarker(v string) *ListWebHostingFilesResponseBodyData {
	s.NextMarker = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyData) SetWebHostingFiles(v []*ListWebHostingFilesResponseBodyDataWebHostingFiles) *ListWebHostingFilesResponseBodyData {
	s.WebHostingFiles = v
	return s
}

type ListWebHostingFilesResponseBodyDataWebHostingFiles struct {
	ContentType      *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ETag             *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	FilePath         *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	LastModifiedTime *int64  `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	SignedUrl        *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListWebHostingFilesResponseBodyDataWebHostingFiles) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponseBodyDataWebHostingFiles) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetContentType(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.ContentType = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetETag(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.ETag = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetFilePath(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.FilePath = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetLastModifiedTime(v int64) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.LastModifiedTime = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetSignedUrl(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.SignedUrl = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetSize(v int64) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.Size = &v
	return s
}

type ListWebHostingFilesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWebHostingFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWebHostingFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponse) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponse) SetHeaders(v map[string]*string) *ListWebHostingFilesResponse {
	s.Headers = v
	return s
}

func (s *ListWebHostingFilesResponse) SetBody(v *ListWebHostingFilesResponseBody) *ListWebHostingFilesResponse {
	s.Body = v
	return s
}

type ModifyWebHostingConfigRequest struct {
	AllowedIps      *string `json:"AllowedIps,omitempty" xml:"AllowedIps,omitempty"`
	ErrorPath       *string `json:"ErrorPath,omitempty" xml:"ErrorPath,omitempty"`
	HistoryModePath *string `json:"HistoryModePath,omitempty" xml:"HistoryModePath,omitempty"`
	IndexPath       *string `json:"IndexPath,omitempty" xml:"IndexPath,omitempty"`
	SpaceId         *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ModifyWebHostingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebHostingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebHostingConfigRequest) SetAllowedIps(v string) *ModifyWebHostingConfigRequest {
	s.AllowedIps = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetErrorPath(v string) *ModifyWebHostingConfigRequest {
	s.ErrorPath = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetHistoryModePath(v string) *ModifyWebHostingConfigRequest {
	s.HistoryModePath = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetIndexPath(v string) *ModifyWebHostingConfigRequest {
	s.IndexPath = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetSpaceId(v string) *ModifyWebHostingConfigRequest {
	s.SpaceId = &v
	return s
}

type ModifyWebHostingConfigResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebHostingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebHostingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebHostingConfigResponseBody) SetData(v bool) *ModifyWebHostingConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyWebHostingConfigResponseBody) SetRequestId(v string) *ModifyWebHostingConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebHostingConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebHostingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebHostingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebHostingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebHostingConfigResponse) SetHeaders(v map[string]*string) *ModifyWebHostingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebHostingConfigResponse) SetBody(v *ModifyWebHostingConfigResponseBody) *ModifyWebHostingConfigResponse {
	s.Body = v
	return s
}

type OpenServiceRequest struct {
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s OpenServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenServiceRequest) SetServiceName(v string) *OpenServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *OpenServiceRequest) SetSpaceId(v string) *OpenServiceRequest {
	s.SpaceId = &v
	return s
}

type OpenServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBody) SetRequestId(v string) *OpenServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenServiceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenServiceResponse) SetHeaders(v map[string]*string) *OpenServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenServiceResponse) SetBody(v *OpenServiceResponseBody) *OpenServiceResponse {
	s.Body = v
	return s
}

type OpenWebHostingServiceRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s OpenWebHostingServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenWebHostingServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenWebHostingServiceRequest) SetSpaceId(v string) *OpenWebHostingServiceRequest {
	s.SpaceId = &v
	return s
}

type OpenWebHostingServiceResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenWebHostingServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenWebHostingServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenWebHostingServiceResponseBody) SetData(v bool) *OpenWebHostingServiceResponseBody {
	s.Data = &v
	return s
}

func (s *OpenWebHostingServiceResponseBody) SetRequestId(v string) *OpenWebHostingServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenWebHostingServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenWebHostingServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenWebHostingServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenWebHostingServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenWebHostingServiceResponse) SetHeaders(v map[string]*string) *OpenWebHostingServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenWebHostingServiceResponse) SetBody(v *OpenWebHostingServiceResponseBody) *OpenWebHostingServiceResponse {
	s.Body = v
	return s
}

type QueryDBBackupCollectionsRequest struct {
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s QueryDBBackupCollectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupCollectionsRequest) GoString() string {
	return s.String()
}

func (s *QueryDBBackupCollectionsRequest) SetBackupId(v string) *QueryDBBackupCollectionsRequest {
	s.BackupId = &v
	return s
}

func (s *QueryDBBackupCollectionsRequest) SetSpaceId(v string) *QueryDBBackupCollectionsRequest {
	s.SpaceId = &v
	return s
}

type QueryDBBackupCollectionsResponseBody struct {
	Collections []*string `json:"Collections,omitempty" xml:"Collections,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDBBackupCollectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBBackupCollectionsResponseBody) SetCollections(v []*string) *QueryDBBackupCollectionsResponseBody {
	s.Collections = v
	return s
}

func (s *QueryDBBackupCollectionsResponseBody) SetRequestId(v string) *QueryDBBackupCollectionsResponseBody {
	s.RequestId = &v
	return s
}

type QueryDBBackupCollectionsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBBackupCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBBackupCollectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupCollectionsResponse) GoString() string {
	return s.String()
}

func (s *QueryDBBackupCollectionsResponse) SetHeaders(v map[string]*string) *QueryDBBackupCollectionsResponse {
	s.Headers = v
	return s
}

func (s *QueryDBBackupCollectionsResponse) SetBody(v *QueryDBBackupCollectionsResponseBody) *QueryDBBackupCollectionsResponse {
	s.Body = v
	return s
}

type QueryDBBackupDumpTimesRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s QueryDBBackupDumpTimesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesRequest) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesRequest) SetSpaceId(v string) *QueryDBBackupDumpTimesRequest {
	s.SpaceId = &v
	return s
}

type QueryDBBackupDumpTimesResponseBody struct {
	BackupDumpTimes []*QueryDBBackupDumpTimesResponseBodyBackupDumpTimes `json:"BackupDumpTimes,omitempty" xml:"BackupDumpTimes,omitempty" type:"Repeated"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDBBackupDumpTimesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesResponseBody) SetBackupDumpTimes(v []*QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) *QueryDBBackupDumpTimesResponseBody {
	s.BackupDumpTimes = v
	return s
}

func (s *QueryDBBackupDumpTimesResponseBody) SetRequestId(v string) *QueryDBBackupDumpTimesResponseBody {
	s.RequestId = &v
	return s
}

type QueryDBBackupDumpTimesResponseBodyBackupDumpTimes struct {
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DumpTime *string `json:"DumpTime,omitempty" xml:"DumpTime,omitempty"`
}

func (s QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) SetBackupId(v string) *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes {
	s.BackupId = &v
	return s
}

func (s *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) SetDumpTime(v string) *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes {
	s.DumpTime = &v
	return s
}

type QueryDBBackupDumpTimesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBBackupDumpTimesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBBackupDumpTimesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesResponse) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesResponse) SetHeaders(v map[string]*string) *QueryDBBackupDumpTimesResponse {
	s.Headers = v
	return s
}

func (s *QueryDBBackupDumpTimesResponse) SetBody(v *QueryDBBackupDumpTimesResponseBody) *QueryDBBackupDumpTimesResponse {
	s.Body = v
	return s
}

type QueryDBExportTaskStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryDBExportTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBExportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDBExportTaskStatusRequest) SetSpaceId(v string) *QueryDBExportTaskStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBExportTaskStatusRequest) SetTaskId(v string) *QueryDBExportTaskStatusRequest {
	s.TaskId = &v
	return s
}

type QueryDBExportTaskStatusResponseBody struct {
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
	DownloadUrl   *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExportedCount *string `json:"ExportedCount,omitempty" xml:"ExportedCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDBExportTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBExportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBExportTaskStatusResponseBody) SetDetailMessage(v string) *QueryDBExportTaskStatusResponseBody {
	s.DetailMessage = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetDownloadUrl(v string) *QueryDBExportTaskStatusResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetExportedCount(v string) *QueryDBExportTaskStatusResponseBody {
	s.ExportedCount = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetRequestId(v string) *QueryDBExportTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetStatus(v string) *QueryDBExportTaskStatusResponseBody {
	s.Status = &v
	return s
}

type QueryDBExportTaskStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBExportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBExportTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBExportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDBExportTaskStatusResponse) SetHeaders(v map[string]*string) *QueryDBExportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDBExportTaskStatusResponse) SetBody(v *QueryDBExportTaskStatusResponseBody) *QueryDBExportTaskStatusResponse {
	s.Body = v
	return s
}

type QueryDBImportTaskStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryDBImportTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBImportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDBImportTaskStatusRequest) SetSpaceId(v string) *QueryDBImportTaskStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBImportTaskStatusRequest) SetTaskId(v string) *QueryDBImportTaskStatusRequest {
	s.TaskId = &v
	return s
}

type QueryDBImportTaskStatusResponseBody struct {
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
	FailedCount   *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessCount  *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s QueryDBImportTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBImportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBImportTaskStatusResponseBody) SetDetailMessage(v string) *QueryDBImportTaskStatusResponseBody {
	s.DetailMessage = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetFailedCount(v string) *QueryDBImportTaskStatusResponseBody {
	s.FailedCount = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetRequestId(v string) *QueryDBImportTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetStatus(v string) *QueryDBImportTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetSuccessCount(v string) *QueryDBImportTaskStatusResponseBody {
	s.SuccessCount = &v
	return s
}

type QueryDBImportTaskStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBImportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBImportTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBImportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDBImportTaskStatusResponse) SetHeaders(v map[string]*string) *QueryDBImportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDBImportTaskStatusResponse) SetBody(v *QueryDBImportTaskStatusResponseBody) *QueryDBImportTaskStatusResponse {
	s.Body = v
	return s
}

type QueryDBRestoreTaskStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryDBRestoreTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBRestoreTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDBRestoreTaskStatusRequest) SetSpaceId(v string) *QueryDBRestoreTaskStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBRestoreTaskStatusRequest) SetTaskId(v string) *QueryDBRestoreTaskStatusRequest {
	s.TaskId = &v
	return s
}

type QueryDBRestoreTaskStatusResponseBody struct {
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
	FailedCount   *int64  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessCount  *int64  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s QueryDBRestoreTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBRestoreTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetDetailMessage(v string) *QueryDBRestoreTaskStatusResponseBody {
	s.DetailMessage = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetFailedCount(v int64) *QueryDBRestoreTaskStatusResponseBody {
	s.FailedCount = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetRequestId(v string) *QueryDBRestoreTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetStatus(v string) *QueryDBRestoreTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetSuccessCount(v int64) *QueryDBRestoreTaskStatusResponseBody {
	s.SuccessCount = &v
	return s
}

type QueryDBRestoreTaskStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBRestoreTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBRestoreTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBRestoreTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDBRestoreTaskStatusResponse) SetHeaders(v map[string]*string) *QueryDBRestoreTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDBRestoreTaskStatusResponse) SetBody(v *QueryDBRestoreTaskStatusResponseBody) *QueryDBRestoreTaskStatusResponse {
	s.Body = v
	return s
}

type QueryServiceStatusRequest struct {
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s QueryServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceStatusRequest) SetServiceName(v string) *QueryServiceStatusRequest {
	s.ServiceName = &v
	return s
}

func (s *QueryServiceStatusRequest) SetSpaceId(v string) *QueryServiceStatusRequest {
	s.SpaceId = &v
	return s
}

type QueryServiceStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s QueryServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceStatusResponseBody) SetRequestId(v string) *QueryServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServiceStatusResponseBody) SetServiceStatus(v string) *QueryServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type QueryServiceStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceStatusResponse) SetHeaders(v map[string]*string) *QueryServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceStatusResponse) SetBody(v *QueryServiceStatusResponseBody) *QueryServiceStatusResponse {
	s.Body = v
	return s
}

type RegisterFileRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s RegisterFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterFileRequest) GoString() string {
	return s.String()
}

func (s *RegisterFileRequest) SetId(v string) *RegisterFileRequest {
	s.Id = &v
	return s
}

func (s *RegisterFileRequest) SetSpaceId(v string) *RegisterFileRequest {
	s.SpaceId = &v
	return s
}

type RegisterFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterFileResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterFileResponseBody) SetRequestId(v string) *RegisterFileResponseBody {
	s.RequestId = &v
	return s
}

type RegisterFileResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterFileResponse) GoString() string {
	return s.String()
}

func (s *RegisterFileResponse) SetHeaders(v map[string]*string) *RegisterFileResponse {
	s.Headers = v
	return s
}

func (s *RegisterFileResponse) SetBody(v *RegisterFileResponseBody) *RegisterFileResponse {
	s.Body = v
	return s
}

type RenameDBCollectionRequest struct {
	NewCollection    *string `json:"NewCollection,omitempty" xml:"NewCollection,omitempty"`
	OriginCollection *string `json:"OriginCollection,omitempty" xml:"OriginCollection,omitempty"`
	SpaceId          *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s RenameDBCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDBCollectionRequest) GoString() string {
	return s.String()
}

func (s *RenameDBCollectionRequest) SetNewCollection(v string) *RenameDBCollectionRequest {
	s.NewCollection = &v
	return s
}

func (s *RenameDBCollectionRequest) SetOriginCollection(v string) *RenameDBCollectionRequest {
	s.OriginCollection = &v
	return s
}

func (s *RenameDBCollectionRequest) SetSpaceId(v string) *RenameDBCollectionRequest {
	s.SpaceId = &v
	return s
}

type RenameDBCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameDBCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDBCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDBCollectionResponseBody) SetRequestId(v string) *RenameDBCollectionResponseBody {
	s.RequestId = &v
	return s
}

type RenameDBCollectionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameDBCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameDBCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDBCollectionResponse) GoString() string {
	return s.String()
}

func (s *RenameDBCollectionResponse) SetHeaders(v map[string]*string) *RenameDBCollectionResponse {
	s.Headers = v
	return s
}

func (s *RenameDBCollectionResponse) SetBody(v *RenameDBCollectionResponseBody) *RenameDBCollectionResponse {
	s.Body = v
	return s
}

type ResetServerSecretRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ResetServerSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetServerSecretRequest) GoString() string {
	return s.String()
}

func (s *ResetServerSecretRequest) SetSpaceId(v string) *ResetServerSecretRequest {
	s.SpaceId = &v
	return s
}

type ResetServerSecretResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetServerSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetServerSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ResetServerSecretResponseBody) SetRequestId(v string) *ResetServerSecretResponseBody {
	s.RequestId = &v
	return s
}

type ResetServerSecretResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetServerSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetServerSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetServerSecretResponse) GoString() string {
	return s.String()
}

func (s *ResetServerSecretResponse) SetHeaders(v map[string]*string) *ResetServerSecretResponse {
	s.Headers = v
	return s
}

func (s *ResetServerSecretResponse) SetBody(v *ResetServerSecretResponseBody) *ResetServerSecretResponse {
	s.Body = v
	return s
}

type RunDBCommandRequest struct {
	Body    *string `json:"Body,omitempty" xml:"Body,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s RunDBCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDBCommandRequest) GoString() string {
	return s.String()
}

func (s *RunDBCommandRequest) SetBody(v string) *RunDBCommandRequest {
	s.Body = &v
	return s
}

func (s *RunDBCommandRequest) SetSpaceId(v string) *RunDBCommandRequest {
	s.SpaceId = &v
	return s
}

type RunDBCommandResponseBody struct {
	AffectedDocs *int32  `json:"AffectedDocs,omitempty" xml:"AffectedDocs,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RunDBCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunDBCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunDBCommandResponseBody) SetAffectedDocs(v int32) *RunDBCommandResponseBody {
	s.AffectedDocs = &v
	return s
}

func (s *RunDBCommandResponseBody) SetRequestId(v string) *RunDBCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDBCommandResponseBody) SetResult(v string) *RunDBCommandResponseBody {
	s.Result = &v
	return s
}

type RunDBCommandResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunDBCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunDBCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunDBCommandResponse) GoString() string {
	return s.String()
}

func (s *RunDBCommandResponse) SetHeaders(v map[string]*string) *RunDBCommandResponse {
	s.Headers = v
	return s
}

func (s *RunDBCommandResponse) SetBody(v *RunDBCommandResponseBody) *RunDBCommandResponse {
	s.Body = v
	return s
}

type RunFunctionRequest struct {
	Body    *string `json:"Body,omitempty" xml:"Body,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s RunFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionRequest) GoString() string {
	return s.String()
}

func (s *RunFunctionRequest) SetBody(v string) *RunFunctionRequest {
	s.Body = &v
	return s
}

func (s *RunFunctionRequest) SetSpaceId(v string) *RunFunctionRequest {
	s.SpaceId = &v
	return s
}

type RunFunctionResponseBody struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result      *string                             `json:"Result,omitempty" xml:"Result,omitempty"`
	RuntimeMeta *RunFunctionResponseBodyRuntimeMeta `json:"RuntimeMeta,omitempty" xml:"RuntimeMeta,omitempty" type:"Struct"`
}

func (s RunFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RunFunctionResponseBody) SetRequestId(v string) *RunFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunFunctionResponseBody) SetResult(v string) *RunFunctionResponseBody {
	s.Result = &v
	return s
}

func (s *RunFunctionResponseBody) SetRuntimeMeta(v *RunFunctionResponseBodyRuntimeMeta) *RunFunctionResponseBody {
	s.RuntimeMeta = v
	return s
}

type RunFunctionResponseBodyRuntimeMeta struct {
	BillingDuration    *int32  `json:"BillingDuration,omitempty" xml:"BillingDuration,omitempty"`
	InvocationDuration *int32  `json:"InvocationDuration,omitempty" xml:"InvocationDuration,omitempty"`
	MaxMemoryUsage     *int32  `json:"MaxMemoryUsage,omitempty" xml:"MaxMemoryUsage,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunFunctionResponseBodyRuntimeMeta) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionResponseBodyRuntimeMeta) GoString() string {
	return s.String()
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetBillingDuration(v int32) *RunFunctionResponseBodyRuntimeMeta {
	s.BillingDuration = &v
	return s
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetInvocationDuration(v int32) *RunFunctionResponseBodyRuntimeMeta {
	s.InvocationDuration = &v
	return s
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetMaxMemoryUsage(v int32) *RunFunctionResponseBodyRuntimeMeta {
	s.MaxMemoryUsage = &v
	return s
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetRequestId(v string) *RunFunctionResponseBodyRuntimeMeta {
	s.RequestId = &v
	return s
}

type RunFunctionResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionResponse) GoString() string {
	return s.String()
}

func (s *RunFunctionResponse) SetHeaders(v map[string]*string) *RunFunctionResponse {
	s.Headers = v
	return s
}

func (s *RunFunctionResponse) SetBody(v *RunFunctionResponseBody) *RunFunctionResponse {
	s.Body = v
	return s
}

type SaveAntOpenPlatformConfigRequest struct {
	AppCert    *string `json:"AppCert,omitempty" xml:"AppCert,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PublicCert *string `json:"PublicCert,omitempty" xml:"PublicCert,omitempty"`
	PublicKey  *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	RootCert   *string `json:"RootCert,omitempty" xml:"RootCert,omitempty"`
	SignMode   *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s SaveAntOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAntOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveAntOpenPlatformConfigRequest) SetAppCert(v string) *SaveAntOpenPlatformConfigRequest {
	s.AppCert = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetAppId(v string) *SaveAntOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetPrivateKey(v string) *SaveAntOpenPlatformConfigRequest {
	s.PrivateKey = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetPublicCert(v string) *SaveAntOpenPlatformConfigRequest {
	s.PublicCert = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetPublicKey(v string) *SaveAntOpenPlatformConfigRequest {
	s.PublicKey = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetRootCert(v string) *SaveAntOpenPlatformConfigRequest {
	s.RootCert = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetSignMode(v string) *SaveAntOpenPlatformConfigRequest {
	s.SignMode = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetSpaceId(v string) *SaveAntOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type SaveAntOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveAntOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAntOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAntOpenPlatformConfigResponseBody) SetRequestId(v string) *SaveAntOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveAntOpenPlatformConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveAntOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAntOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAntOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveAntOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *SaveAntOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveAntOpenPlatformConfigResponse) SetBody(v *SaveAntOpenPlatformConfigResponseBody) *SaveAntOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type SaveAppAuthTokenRequest struct {
	AppAuthToken *string `json:"AppAuthToken,omitempty" xml:"AppAuthToken,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	IsvAppId     *string `json:"IsvAppId,omitempty" xml:"IsvAppId,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s SaveAppAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAppAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *SaveAppAuthTokenRequest) SetAppAuthToken(v string) *SaveAppAuthTokenRequest {
	s.AppAuthToken = &v
	return s
}

func (s *SaveAppAuthTokenRequest) SetAppId(v string) *SaveAppAuthTokenRequest {
	s.AppId = &v
	return s
}

func (s *SaveAppAuthTokenRequest) SetIsvAppId(v string) *SaveAppAuthTokenRequest {
	s.IsvAppId = &v
	return s
}

func (s *SaveAppAuthTokenRequest) SetSpaceId(v string) *SaveAppAuthTokenRequest {
	s.SpaceId = &v
	return s
}

type SaveAppAuthTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveAppAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAppAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAppAuthTokenResponseBody) SetRequestId(v string) *SaveAppAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

type SaveAppAuthTokenResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveAppAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAppAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAppAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *SaveAppAuthTokenResponse) SetHeaders(v map[string]*string) *SaveAppAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *SaveAppAuthTokenResponse) SetBody(v *SaveAppAuthTokenResponseBody) *SaveAppAuthTokenResponse {
	s.Body = v
	return s
}

type SaveWebHostingCustomDomainConfigRequest struct {
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ForceRedirectType *string `json:"ForceRedirectType,omitempty" xml:"ForceRedirectType,omitempty"`
	SpaceId           *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s SaveWebHostingCustomDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainConfigRequest) SetDomainName(v string) *SaveWebHostingCustomDomainConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SaveWebHostingCustomDomainConfigRequest) SetForceRedirectType(v string) *SaveWebHostingCustomDomainConfigRequest {
	s.ForceRedirectType = &v
	return s
}

func (s *SaveWebHostingCustomDomainConfigRequest) SetSpaceId(v string) *SaveWebHostingCustomDomainConfigRequest {
	s.SpaceId = &v
	return s
}

type SaveWebHostingCustomDomainConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveWebHostingCustomDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainConfigResponseBody) SetRequestId(v string) *SaveWebHostingCustomDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveWebHostingCustomDomainConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWebHostingCustomDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebHostingCustomDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainConfigResponse) SetHeaders(v map[string]*string) *SaveWebHostingCustomDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveWebHostingCustomDomainConfigResponse) SetBody(v *SaveWebHostingCustomDomainConfigResponseBody) *SaveWebHostingCustomDomainConfigResponse {
	s.Body = v
	return s
}

type SaveWebHostingCustomDomainCorsConfigRequest struct {
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	DomainName               *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableCors               *bool   `json:"EnableCors,omitempty" xml:"EnableCors,omitempty"`
	SpaceId                  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s SaveWebHostingCustomDomainCorsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainCorsConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetAccessControlAllowOrigin(v string) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.AccessControlAllowOrigin = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetDomainName(v string) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetEnableCors(v bool) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.EnableCors = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetSpaceId(v string) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.SpaceId = &v
	return s
}

type SaveWebHostingCustomDomainCorsConfigResponseBody struct {
	// Id of the request
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveWebHostingCustomDomainCorsConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainCorsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetCode(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetData(v bool) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetHttpStatusCode(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetMessage(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetRequestId(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetSuccess(v bool) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Success = &v
	return s
}

type SaveWebHostingCustomDomainCorsConfigResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWebHostingCustomDomainCorsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebHostingCustomDomainCorsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainCorsConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainCorsConfigResponse) SetHeaders(v map[string]*string) *SaveWebHostingCustomDomainCorsConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponse) SetBody(v *SaveWebHostingCustomDomainCorsConfigResponseBody) *SaveWebHostingCustomDomainCorsConfigResponse {
	s.Body = v
	return s
}

type SaveWechatOpenPlatformConfigRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s SaveWechatOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWechatOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveWechatOpenPlatformConfigRequest) SetAppId(v string) *SaveWechatOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *SaveWechatOpenPlatformConfigRequest) SetAppSecret(v string) *SaveWechatOpenPlatformConfigRequest {
	s.AppSecret = &v
	return s
}

func (s *SaveWechatOpenPlatformConfigRequest) SetSpaceId(v string) *SaveWechatOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type SaveWechatOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveWechatOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWechatOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWechatOpenPlatformConfigResponseBody) SetRequestId(v string) *SaveWechatOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveWechatOpenPlatformConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWechatOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWechatOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWechatOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveWechatOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *SaveWechatOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveWechatOpenPlatformConfigResponse) SetBody(v *SaveWechatOpenPlatformConfigResponseBody) *SaveWechatOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type UnbindWebHostingCustomDomainRequest struct {
	CustomDomain *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UnbindWebHostingCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindWebHostingCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *UnbindWebHostingCustomDomainRequest) SetCustomDomain(v string) *UnbindWebHostingCustomDomainRequest {
	s.CustomDomain = &v
	return s
}

func (s *UnbindWebHostingCustomDomainRequest) SetSpaceId(v string) *UnbindWebHostingCustomDomainRequest {
	s.SpaceId = &v
	return s
}

type UnbindWebHostingCustomDomainResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindWebHostingCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindWebHostingCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindWebHostingCustomDomainResponseBody) SetData(v bool) *UnbindWebHostingCustomDomainResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindWebHostingCustomDomainResponseBody) SetRequestId(v string) *UnbindWebHostingCustomDomainResponseBody {
	s.RequestId = &v
	return s
}

type UnbindWebHostingCustomDomainResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindWebHostingCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindWebHostingCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindWebHostingCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *UnbindWebHostingCustomDomainResponse) SetHeaders(v map[string]*string) *UnbindWebHostingCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *UnbindWebHostingCustomDomainResponse) SetBody(v *UnbindWebHostingCustomDomainResponseBody) *UnbindWebHostingCustomDomainResponse {
	s.Body = v
	return s
}

type UpdateDingtalkOpenPlatformConfigRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UpdateDingtalkOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingtalkOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDingtalkOpenPlatformConfigRequest) SetAppId(v string) *UpdateDingtalkOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateDingtalkOpenPlatformConfigRequest) SetAppSecret(v string) *UpdateDingtalkOpenPlatformConfigRequest {
	s.AppSecret = &v
	return s
}

func (s *UpdateDingtalkOpenPlatformConfigRequest) SetSpaceId(v string) *UpdateDingtalkOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type UpdateDingtalkOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDingtalkOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingtalkOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDingtalkOpenPlatformConfigResponseBody) SetRequestId(v string) *UpdateDingtalkOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDingtalkOpenPlatformConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDingtalkOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDingtalkOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingtalkOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDingtalkOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *UpdateDingtalkOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDingtalkOpenPlatformConfigResponse) SetBody(v *UpdateDingtalkOpenPlatformConfigResponseBody) *UpdateDingtalkOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	Desc                *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	HttpTriggerPath     *string `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Timeout             *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	TimingTriggerConfig *string `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetDesc(v string) *UpdateFunctionRequest {
	s.Desc = &v
	return s
}

func (s *UpdateFunctionRequest) SetHttpTriggerPath(v string) *UpdateFunctionRequest {
	s.HttpTriggerPath = &v
	return s
}

func (s *UpdateFunctionRequest) SetInstanceConcurrency(v int32) *UpdateFunctionRequest {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionRequest) SetMemory(v int32) *UpdateFunctionRequest {
	s.Memory = &v
	return s
}

func (s *UpdateFunctionRequest) SetName(v string) *UpdateFunctionRequest {
	s.Name = &v
	return s
}

func (s *UpdateFunctionRequest) SetRuntime(v string) *UpdateFunctionRequest {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionRequest) SetSpaceId(v string) *UpdateFunctionRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateFunctionRequest) SetTimeout(v int32) *UpdateFunctionRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateFunctionRequest) SetTimingTriggerConfig(v string) *UpdateFunctionRequest {
	s.TimingTriggerConfig = &v
	return s
}

type UpdateFunctionResponseBody struct {
	CreatedAt           *string                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Desc                *string                         `json:"Desc,omitempty" xml:"Desc,omitempty"`
	HttpTriggerPath     *string                         `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	ModifiedAt          *string                         `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Name                *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId           *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spec                *UpdateFunctionResponseBodySpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	TimingTriggerConfig *string                         `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetCreatedAt(v string) *UpdateFunctionResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetDesc(v string) *UpdateFunctionResponseBody {
	s.Desc = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetHttpTriggerPath(v string) *UpdateFunctionResponseBody {
	s.HttpTriggerPath = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetModifiedAt(v string) *UpdateFunctionResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetName(v string) *UpdateFunctionResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSpec(v *UpdateFunctionResponseBodySpec) *UpdateFunctionResponseBody {
	s.Spec = v
	return s
}

func (s *UpdateFunctionResponseBody) SetTimingTriggerConfig(v string) *UpdateFunctionResponseBody {
	s.TimingTriggerConfig = &v
	return s
}

type UpdateFunctionResponseBodySpec struct {
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateFunctionResponseBodySpec) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBodySpec) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBodySpec) SetInstanceConcurrency(v int32) *UpdateFunctionResponseBodySpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionResponseBodySpec) SetMemory(v string) *UpdateFunctionResponseBodySpec {
	s.Memory = &v
	return s
}

func (s *UpdateFunctionResponseBodySpec) SetRuntime(v string) *UpdateFunctionResponseBodySpec {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionResponseBodySpec) SetTimeout(v string) *UpdateFunctionResponseBodySpec {
	s.Timeout = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateHttpTriggerConfigRequest struct {
	CustomDomain            *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	CustomDomainCertificate *string `json:"CustomDomainCertificate,omitempty" xml:"CustomDomainCertificate,omitempty"`
	CustomDomainPrivateKey  *string `json:"CustomDomainPrivateKey,omitempty" xml:"CustomDomainPrivateKey,omitempty"`
	EnableService           *bool   `json:"EnableService,omitempty" xml:"EnableService,omitempty"`
	SpaceId                 *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UpdateHttpTriggerConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpTriggerConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpTriggerConfigRequest) SetCustomDomain(v string) *UpdateHttpTriggerConfigRequest {
	s.CustomDomain = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetCustomDomainCertificate(v string) *UpdateHttpTriggerConfigRequest {
	s.CustomDomainCertificate = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetCustomDomainPrivateKey(v string) *UpdateHttpTriggerConfigRequest {
	s.CustomDomainPrivateKey = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetEnableService(v bool) *UpdateHttpTriggerConfigRequest {
	s.EnableService = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetSpaceId(v string) *UpdateHttpTriggerConfigRequest {
	s.SpaceId = &v
	return s
}

type UpdateHttpTriggerConfigResponseBody struct {
	CustomDomain                *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	CustomDomainCertificateInfo *string `json:"CustomDomainCertificateInfo,omitempty" xml:"CustomDomainCertificateInfo,omitempty"`
	CustomDomainCname           *string `json:"CustomDomainCname,omitempty" xml:"CustomDomainCname,omitempty"`
	DefaultEndpoint             *string `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty"`
	EnableService               *bool   `json:"EnableService,omitempty" xml:"EnableService,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpTriggerConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpTriggerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpTriggerConfigResponseBody) SetCustomDomain(v string) *UpdateHttpTriggerConfigResponseBody {
	s.CustomDomain = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetCustomDomainCertificateInfo(v string) *UpdateHttpTriggerConfigResponseBody {
	s.CustomDomainCertificateInfo = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetCustomDomainCname(v string) *UpdateHttpTriggerConfigResponseBody {
	s.CustomDomainCname = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetDefaultEndpoint(v string) *UpdateHttpTriggerConfigResponseBody {
	s.DefaultEndpoint = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetEnableService(v bool) *UpdateHttpTriggerConfigResponseBody {
	s.EnableService = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetRequestId(v string) *UpdateHttpTriggerConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHttpTriggerConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateHttpTriggerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHttpTriggerConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpTriggerConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpTriggerConfigResponse) SetHeaders(v map[string]*string) *UpdateHttpTriggerConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpTriggerConfigResponse) SetBody(v *UpdateHttpTriggerConfigResponseBody) *UpdateHttpTriggerConfigResponse {
	s.Body = v
	return s
}

type UpdateServicePolicyRequest struct {
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UpdateServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateServicePolicyRequest) SetCollectionName(v string) *UpdateServicePolicyRequest {
	s.CollectionName = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetPolicy(v string) *UpdateServicePolicyRequest {
	s.Policy = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetPolicyName(v string) *UpdateServicePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetServiceName(v string) *UpdateServicePolicyRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetSpaceId(v string) *UpdateServicePolicyRequest {
	s.SpaceId = &v
	return s
}

type UpdateServicePolicyResponseBody struct {
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UpdateServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServicePolicyResponseBody) SetCollectionName(v string) *UpdateServicePolicyResponseBody {
	s.CollectionName = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetPolicy(v string) *UpdateServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetPolicyName(v string) *UpdateServicePolicyResponseBody {
	s.PolicyName = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetRequestId(v string) *UpdateServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetServiceName(v string) *UpdateServicePolicyResponseBody {
	s.ServiceName = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetSpaceId(v string) *UpdateServicePolicyResponseBody {
	s.SpaceId = &v
	return s
}

type UpdateServicePolicyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateServicePolicyResponse) SetHeaders(v map[string]*string) *UpdateServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateServicePolicyResponse) SetBody(v *UpdateServicePolicyResponseBody) *UpdateServicePolicyResponse {
	s.Body = v
	return s
}

type UpdateSpaceRequest struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpaceRequest) SetDesc(v string) *UpdateSpaceRequest {
	s.Desc = &v
	return s
}

func (s *UpdateSpaceRequest) SetSpaceId(v string) *UpdateSpaceRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSpaceRequest) SetStatus(v string) *UpdateSpaceRequest {
	s.Status = &v
	return s
}

type UpdateSpaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSpaceResponseBody) SetRequestId(v string) *UpdateSpaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSpaceResponse) SetHeaders(v map[string]*string) *UpdateSpaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSpaceResponse) SetBody(v *UpdateSpaceResponseBody) *UpdateSpaceResponse {
	s.Body = v
	return s
}

type VerifyWebHostingDomainOwnerRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyWebHostingDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebHostingDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyWebHostingDomainOwnerRequest) SetDomain(v string) *VerifyWebHostingDomainOwnerRequest {
	s.Domain = &v
	return s
}

func (s *VerifyWebHostingDomainOwnerRequest) SetSpaceId(v string) *VerifyWebHostingDomainOwnerRequest {
	s.SpaceId = &v
	return s
}

func (s *VerifyWebHostingDomainOwnerRequest) SetVerifyType(v string) *VerifyWebHostingDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

type VerifyWebHostingDomainOwnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyWebHostingDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebHostingDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWebHostingDomainOwnerResponseBody) SetRequestId(v string) *VerifyWebHostingDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyWebHostingDomainOwnerResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyWebHostingDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyWebHostingDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebHostingDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyWebHostingDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyWebHostingDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyWebHostingDomainOwnerResponse) SetBody(v *VerifyWebHostingDomainOwnerResponseBody) *VerifyWebHostingDomainOwnerResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mpserverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddCorsDomainWithOptions(request *AddCorsDomainRequest, runtime *util.RuntimeOptions) (_result *AddCorsDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCorsDomain"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCorsDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCorsDomain(request *AddCorsDomainRequest) (_result *AddCorsDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCorsDomainResponse{}
	_body, _err := client.AddCorsDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDingtalkOpenPlatformConfigWithOptions(request *AddDingtalkOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *AddDingtalkOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		body["AppSecret"] = request.AppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDingtalkOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDingtalkOpenPlatformConfig(request *AddDingtalkOpenPlatformConfigRequest) (_result *AddDingtalkOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.AddDingtalkOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachWebHostingCertificateWithOptions(request *AttachWebHostingCertificateRequest, runtime *util.RuntimeOptions) (_result *AttachWebHostingCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		body["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		body["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		body["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificate)) {
		body["ServerCertificate"] = request.ServerCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachWebHostingCertificate"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachWebHostingCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachWebHostingCertificate(request *AttachWebHostingCertificateRequest) (_result *AttachWebHostingCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachWebHostingCertificateResponse{}
	_body, _err := client.AttachWebHostingCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteWebHostingFilesWithOptions(request *BatchDeleteWebHostingFilesRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteWebHostingFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilePaths)) {
		body["FilePaths"] = request.FilePaths
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteWebHostingFiles"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteWebHostingFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteWebHostingFiles(request *BatchDeleteWebHostingFilesRequest) (_result *BatchDeleteWebHostingFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteWebHostingFilesResponse{}
	_body, _err := client.BatchDeleteWebHostingFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindWebHostingCustomDomainWithOptions(request *BindWebHostingCustomDomainRequest, runtime *util.RuntimeOptions) (_result *BindWebHostingCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomDomain)) {
		body["CustomDomain"] = request.CustomDomain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindWebHostingCustomDomain"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindWebHostingCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindWebHostingCustomDomain(request *BindWebHostingCustomDomainRequest) (_result *BindWebHostingCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindWebHostingCustomDomainResponse{}
	_body, _err := client.BindWebHostingCustomDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMpServerlessRoleExistsWithOptions(request *CheckMpServerlessRoleExistsRequest, runtime *util.RuntimeOptions) (_result *CheckMpServerlessRoleExistsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckMpServerlessRoleExists"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMpServerlessRoleExistsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMpServerlessRoleExists(request *CheckMpServerlessRoleExistsRequest) (_result *CheckMpServerlessRoleExistsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMpServerlessRoleExistsResponse{}
	_body, _err := client.CheckMpServerlessRoleExistsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBExportTaskWithOptions(request *CreateDBExportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDBExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		body["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBExportTask"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBExportTask(request *CreateDBExportTaskRequest) (_result *CreateDBExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBExportTaskResponse{}
	_body, _err := client.CreateDBExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBImportTaskWithOptions(request *CreateDBImportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDBImportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collection)) {
		body["Collection"] = request.Collection
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBImportTask"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBImportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBImportTask(request *CreateDBImportTaskRequest) (_result *CreateDBImportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBImportTaskResponse{}
	_body, _err := client.CreateDBImportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBRestoreTaskWithOptions(request *CreateDBRestoreTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDBRestoreTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		body["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewCollections)) {
		body["NewCollections"] = request.NewCollections
	}

	if !tea.BoolValue(util.IsUnset(request.OriginCollections)) {
		body["OriginCollections"] = request.OriginCollections
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBRestoreTask"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBRestoreTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBRestoreTask(request *CreateDBRestoreTaskRequest) (_result *CreateDBRestoreTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBRestoreTaskResponse{}
	_body, _err := client.CreateDBRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Runtime)) {
		body["Runtime"] = request.Runtime
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionDeploymentWithOptions(request *CreateFunctionDeploymentRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunctionDeployment"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunctionDeployment(request *CreateFunctionDeploymentRequest) (_result *CreateFunctionDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionDeploymentResponse{}
	_body, _err := client.CreateFunctionDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSpaceWithOptions(request *CreateSpaceRequest, runtime *util.RuntimeOptions) (_result *CreateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSpace"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSpace(request *CreateSpaceRequest) (_result *CreateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSpaceResponse{}
	_body, _err := client.CreateSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAntOpenPlatformConfigWithOptions(request *DeleteAntOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteAntOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAntOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAntOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAntOpenPlatformConfig(request *DeleteAntOpenPlatformConfigRequest) (_result *DeleteAntOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAntOpenPlatformConfigResponse{}
	_body, _err := client.DeleteAntOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCorsDomainWithOptions(request *DeleteCorsDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteCorsDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		body["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCorsDomain"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCorsDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCorsDomain(request *DeleteCorsDomainRequest) (_result *DeleteCorsDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCorsDomainResponse{}
	_body, _err := client.DeleteCorsDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBCollectionWithOptions(request *DeleteDBCollectionRequest, runtime *util.RuntimeOptions) (_result *DeleteDBCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBCollection"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBCollection(request *DeleteDBCollectionRequest) (_result *DeleteDBCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBCollectionResponse{}
	_body, _err := client.DeleteDBCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDingtalkOpenPlatformConfigWithOptions(request *DeleteDingtalkOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDingtalkOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDingtalkOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDingtalkOpenPlatformConfig(request *DeleteDingtalkOpenPlatformConfigRequest) (_result *DeleteDingtalkOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.DeleteDingtalkOpenPlatformConfigWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSpaceWithOptions(request *DeleteSpaceRequest, runtime *util.RuntimeOptions) (_result *DeleteSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSpace"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpace(request *DeleteSpaceRequest) (_result *DeleteSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DeleteSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebHostingCertificateWithOptions(request *DeleteWebHostingCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteWebHostingCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWebHostingCertificate"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWebHostingCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebHostingCertificate(request *DeleteWebHostingCertificateRequest) (_result *DeleteWebHostingCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebHostingCertificateResponse{}
	_body, _err := client.DeleteWebHostingCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebHostingFileWithOptions(request *DeleteWebHostingFileRequest, runtime *util.RuntimeOptions) (_result *DeleteWebHostingFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		body["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWebHostingFile"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWebHostingFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebHostingFile(request *DeleteWebHostingFileRequest) (_result *DeleteWebHostingFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebHostingFileResponse{}
	_body, _err := client.DeleteWebHostingFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWechatOpenPlatformConfigWithOptions(request *DeleteWechatOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteWechatOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWechatOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWechatOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWechatOpenPlatformConfig(request *DeleteWechatOpenPlatformConfigRequest) (_result *DeleteWechatOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWechatOpenPlatformConfigResponse{}
	_body, _err := client.DeleteWechatOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployFunctionWithOptions(request *DeployFunctionRequest, runtime *util.RuntimeOptions) (_result *DeployFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentId)) {
		body["DeploymentId"] = request.DeploymentId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployFunction(request *DeployFunctionRequest) (_result *DeployFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployFunctionResponse{}
	_body, _err := client.DeployFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFCOpenStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFCOpenStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeFCOpenStatus"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFCOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFCOpenStatus() (_result *DescribeFCOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFCOpenStatusResponse{}
	_body, _err := client.DescribeFCOpenStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileUploadSignedUrlWithOptions(request *DescribeFileUploadSignedUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeFileUploadSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Filename)) {
		body["Filename"] = request.Filename
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileUploadSignedUrl"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileUploadSignedUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileUploadSignedUrl(request *DescribeFileUploadSignedUrlRequest) (_result *DescribeFileUploadSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileUploadSignedUrlResponse{}
	_body, _err := client.DescribeFileUploadSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFunctionWithOptions(request *DescribeFunctionRequest, runtime *util.RuntimeOptions) (_result *DescribeFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFunction(request *DescribeFunctionRequest) (_result *DescribeFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFunctionResponse{}
	_body, _err := client.DescribeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHttpTriggerConfigWithOptions(request *DescribeHttpTriggerConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeHttpTriggerConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHttpTriggerConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHttpTriggerConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHttpTriggerConfig(request *DescribeHttpTriggerConfigRequest) (_result *DescribeHttpTriggerConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHttpTriggerConfigResponse{}
	_body, _err := client.DescribeHttpTriggerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceQuotaWithOptions(request *DescribeResourceQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceQuota"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceQuota(request *DescribeResourceQuotaRequest) (_result *DescribeResourceQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceQuotaResponse{}
	_body, _err := client.DescribeResourceQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceUsageWithOptions(request *DescribeResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Format)) {
		body["Format"] = request.Format
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceUsage"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (_result *DescribeResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.DescribeResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServicePolicyWithOptions(request *DescribeServicePolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CollectionName)) {
		body["CollectionName"] = request.CollectionName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServicePolicy"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServicePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServicePolicy(request *DescribeServicePolicyRequest) (_result *DescribeServicePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServicePolicyResponse{}
	_body, _err := client.DescribeServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSpaceClientConfigWithOptions(request *DescribeSpaceClientConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeSpaceClientConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSpaceClientConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSpaceClientConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSpaceClientConfig(request *DescribeSpaceClientConfigRequest) (_result *DescribeSpaceClientConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpaceClientConfigResponse{}
	_body, _err := client.DescribeSpaceClientConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebHostingFileWithOptions(request *DescribeWebHostingFileRequest, runtime *util.RuntimeOptions) (_result *DescribeWebHostingFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		body["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebHostingFile"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebHostingFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebHostingFile(request *DescribeWebHostingFileRequest) (_result *DescribeWebHostingFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebHostingFileResponse{}
	_body, _err := client.DescribeWebHostingFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableExtensionWithOptions(request *EnableExtensionRequest, runtime *util.RuntimeOptions) (_result *EnableExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtensionId)) {
		body["ExtensionId"] = request.ExtensionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableExtension"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableExtension(request *EnableExtensionRequest) (_result *EnableExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableExtensionResponse{}
	_body, _err := client.EnableExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingCertificateDetailWithOptions(request *GetWebHostingCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomDomain)) {
		body["CustomDomain"] = request.CustomDomain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebHostingCertificateDetail"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebHostingCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingCertificateDetail(request *GetWebHostingCertificateDetailRequest) (_result *GetWebHostingCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingCertificateDetailResponse{}
	_body, _err := client.GetWebHostingCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingConfigWithOptions(request *GetWebHostingConfigRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebHostingConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebHostingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingConfig(request *GetWebHostingConfigRequest) (_result *GetWebHostingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingConfigResponse{}
	_body, _err := client.GetWebHostingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingDomainVerificationContentWithOptions(request *GetWebHostingDomainVerificationContentRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingDomainVerificationContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebHostingDomainVerificationContent"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebHostingDomainVerificationContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingDomainVerificationContent(request *GetWebHostingDomainVerificationContentRequest) (_result *GetWebHostingDomainVerificationContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingDomainVerificationContentResponse{}
	_body, _err := client.GetWebHostingDomainVerificationContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingStatusWithOptions(request *GetWebHostingStatusRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebHostingStatus"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebHostingStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingStatus(request *GetWebHostingStatusRequest) (_result *GetWebHostingStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingStatusResponse{}
	_body, _err := client.GetWebHostingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingUploadCredentialWithOptions(request *GetWebHostingUploadCredentialRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingUploadCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		body["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebHostingUploadCredential"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebHostingUploadCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingUploadCredential(request *GetWebHostingUploadCredentialRequest) (_result *GetWebHostingUploadCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingUploadCredentialResponse{}
	_body, _err := client.GetWebHostingUploadCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableCertificatesWithOptions(request *ListAvailableCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListAvailableCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableCertificates"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableCertificates(request *ListAvailableCertificatesRequest) (_result *ListAvailableCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableCertificatesResponse{}
	_body, _err := client.ListAvailableCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorsDomainsWithOptions(request *ListCorsDomainsRequest, runtime *util.RuntimeOptions) (_result *ListCorsDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCorsDomains"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCorsDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorsDomains(request *ListCorsDomainsRequest) (_result *ListCorsDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorsDomainsResponse{}
	_body, _err := client.ListCorsDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDingtalkOpenPlatformConfigsWithOptions(request *ListDingtalkOpenPlatformConfigsRequest, runtime *util.RuntimeOptions) (_result *ListDingtalkOpenPlatformConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDingtalkOpenPlatformConfigs"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDingtalkOpenPlatformConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDingtalkOpenPlatformConfigs(request *ListDingtalkOpenPlatformConfigsRequest) (_result *ListDingtalkOpenPlatformConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDingtalkOpenPlatformConfigsResponse{}
	_body, _err := client.ListDingtalkOpenPlatformConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExtensionsWithOptions(request *ListExtensionsRequest, runtime *util.RuntimeOptions) (_result *ListExtensionsResponse, _err error) {
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
		Action:      tea.String("ListExtensions"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExtensions(request *ListExtensionsRequest) (_result *ListExtensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExtensionsResponse{}
	_body, _err := client.ListExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFileWithOptions(request *ListFileRequest, runtime *util.RuntimeOptions) (_result *ListFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFile"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFile(request *ListFileRequest) (_result *ListFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFileResponse{}
	_body, _err := client.ListFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionWithOptions(request *ListFunctionRequest, runtime *util.RuntimeOptions) (_result *ListFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterBy)) {
		body["FilterBy"] = request.FilterBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunction(request *ListFunctionRequest) (_result *ListFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionResponse{}
	_body, _err := client.ListFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionDeploymentWithOptions(request *ListFunctionDeploymentRequest, runtime *util.RuntimeOptions) (_result *ListFunctionDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionDeployment"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionDeployment(request *ListFunctionDeploymentRequest) (_result *ListFunctionDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionDeploymentResponse{}
	_body, _err := client.ListFunctionDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionLogWithOptions(request *ListFunctionLogRequest, runtime *util.RuntimeOptions) (_result *ListFunctionLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionLog"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionLog(request *ListFunctionLogRequest) (_result *ListFunctionLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionLogResponse{}
	_body, _err := client.ListFunctionLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenPlatformConfigWithOptions(request *ListOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *ListOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenPlatformConfig(request *ListOpenPlatformConfigRequest) (_result *ListOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenPlatformConfigResponse{}
	_body, _err := client.ListOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpaceWithOptions(tmpReq *ListSpaceRequest, runtime *util.RuntimeOptions) (_result *ListSpaceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSpaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SpaceIds)) {
		request.SpaceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpaceIds, tea.String("SpaceIds"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceIdsShrink)) {
		body["SpaceIds"] = request.SpaceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSpace"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpace(request *ListSpaceRequest) (_result *ListSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSpaceResponse{}
	_body, _err := client.ListSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWebHostingCustomDomainsWithOptions(request *ListWebHostingCustomDomainsRequest, runtime *util.RuntimeOptions) (_result *ListWebHostingCustomDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWebHostingCustomDomains"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWebHostingCustomDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWebHostingCustomDomains(request *ListWebHostingCustomDomainsRequest) (_result *ListWebHostingCustomDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWebHostingCustomDomainsResponse{}
	_body, _err := client.ListWebHostingCustomDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWebHostingFilesWithOptions(request *ListWebHostingFilesRequest, runtime *util.RuntimeOptions) (_result *ListWebHostingFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		body["Prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWebHostingFiles"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWebHostingFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWebHostingFiles(request *ListWebHostingFilesRequest) (_result *ListWebHostingFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWebHostingFilesResponse{}
	_body, _err := client.ListWebHostingFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebHostingConfigWithOptions(request *ModifyWebHostingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyWebHostingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedIps)) {
		body["AllowedIps"] = request.AllowedIps
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorPath)) {
		body["ErrorPath"] = request.ErrorPath
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryModePath)) {
		body["HistoryModePath"] = request.HistoryModePath
	}

	if !tea.BoolValue(util.IsUnset(request.IndexPath)) {
		body["IndexPath"] = request.IndexPath
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyWebHostingConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyWebHostingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebHostingConfig(request *ModifyWebHostingConfigRequest) (_result *ModifyWebHostingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebHostingConfigResponse{}
	_body, _err := client.ModifyWebHostingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenServiceWithOptions(request *OpenServiceRequest, runtime *util.RuntimeOptions) (_result *OpenServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenService"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenService(request *OpenServiceRequest) (_result *OpenServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenServiceResponse{}
	_body, _err := client.OpenServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenWebHostingServiceWithOptions(request *OpenWebHostingServiceRequest, runtime *util.RuntimeOptions) (_result *OpenWebHostingServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenWebHostingService"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenWebHostingServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenWebHostingService(request *OpenWebHostingServiceRequest) (_result *OpenWebHostingServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenWebHostingServiceResponse{}
	_body, _err := client.OpenWebHostingServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBBackupCollectionsWithOptions(request *QueryDBBackupCollectionsRequest, runtime *util.RuntimeOptions) (_result *QueryDBBackupCollectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		body["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDBBackupCollections"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDBBackupCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBBackupCollections(request *QueryDBBackupCollectionsRequest) (_result *QueryDBBackupCollectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBBackupCollectionsResponse{}
	_body, _err := client.QueryDBBackupCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBBackupDumpTimesWithOptions(request *QueryDBBackupDumpTimesRequest, runtime *util.RuntimeOptions) (_result *QueryDBBackupDumpTimesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDBBackupDumpTimes"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDBBackupDumpTimesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBBackupDumpTimes(request *QueryDBBackupDumpTimesRequest) (_result *QueryDBBackupDumpTimesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBBackupDumpTimesResponse{}
	_body, _err := client.QueryDBBackupDumpTimesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBExportTaskStatusWithOptions(request *QueryDBExportTaskStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDBExportTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDBExportTaskStatus"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDBExportTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBExportTaskStatus(request *QueryDBExportTaskStatusRequest) (_result *QueryDBExportTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBExportTaskStatusResponse{}
	_body, _err := client.QueryDBExportTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBImportTaskStatusWithOptions(request *QueryDBImportTaskStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDBImportTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDBImportTaskStatus"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDBImportTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBImportTaskStatus(request *QueryDBImportTaskStatusRequest) (_result *QueryDBImportTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBImportTaskStatusResponse{}
	_body, _err := client.QueryDBImportTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBRestoreTaskStatusWithOptions(request *QueryDBRestoreTaskStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDBRestoreTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDBRestoreTaskStatus"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDBRestoreTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBRestoreTaskStatus(request *QueryDBRestoreTaskStatusRequest) (_result *QueryDBRestoreTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBRestoreTaskStatusResponse{}
	_body, _err := client.QueryDBRestoreTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceStatusWithOptions(request *QueryServiceStatusRequest, runtime *util.RuntimeOptions) (_result *QueryServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryServiceStatus"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceStatus(request *QueryServiceStatusRequest) (_result *QueryServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServiceStatusResponse{}
	_body, _err := client.QueryServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterFileWithOptions(request *RegisterFileRequest, runtime *util.RuntimeOptions) (_result *RegisterFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterFile"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterFile(request *RegisterFileRequest) (_result *RegisterFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterFileResponse{}
	_body, _err := client.RegisterFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDBCollectionWithOptions(request *RenameDBCollectionRequest, runtime *util.RuntimeOptions) (_result *RenameDBCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewCollection)) {
		body["NewCollection"] = request.NewCollection
	}

	if !tea.BoolValue(util.IsUnset(request.OriginCollection)) {
		body["OriginCollection"] = request.OriginCollection
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameDBCollection"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameDBCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDBCollection(request *RenameDBCollectionRequest) (_result *RenameDBCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameDBCollectionResponse{}
	_body, _err := client.RenameDBCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetServerSecretWithOptions(request *ResetServerSecretRequest, runtime *util.RuntimeOptions) (_result *ResetServerSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetServerSecret"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetServerSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetServerSecret(request *ResetServerSecretRequest) (_result *ResetServerSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetServerSecretResponse{}
	_body, _err := client.ResetServerSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunDBCommandWithOptions(request *RunDBCommandRequest, runtime *util.RuntimeOptions) (_result *RunDBCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunDBCommand"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunDBCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunDBCommand(request *RunDBCommandRequest) (_result *RunDBCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunDBCommandResponse{}
	_body, _err := client.RunDBCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunFunctionWithOptions(request *RunFunctionRequest, runtime *util.RuntimeOptions) (_result *RunFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunFunction(request *RunFunctionRequest) (_result *RunFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunFunctionResponse{}
	_body, _err := client.RunFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAntOpenPlatformConfigWithOptions(request *SaveAntOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *SaveAntOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCert)) {
		body["AppCert"] = request.AppCert
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		body["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.PublicCert)) {
		body["PublicCert"] = request.PublicCert
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKey)) {
		body["PublicKey"] = request.PublicKey
	}

	if !tea.BoolValue(util.IsUnset(request.RootCert)) {
		body["RootCert"] = request.RootCert
	}

	if !tea.BoolValue(util.IsUnset(request.SignMode)) {
		body["SignMode"] = request.SignMode
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveAntOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveAntOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveAntOpenPlatformConfig(request *SaveAntOpenPlatformConfigRequest) (_result *SaveAntOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveAntOpenPlatformConfigResponse{}
	_body, _err := client.SaveAntOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAppAuthTokenWithOptions(request *SaveAppAuthTokenRequest, runtime *util.RuntimeOptions) (_result *SaveAppAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAuthToken)) {
		body["AppAuthToken"] = request.AppAuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvAppId)) {
		body["IsvAppId"] = request.IsvAppId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveAppAuthToken"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveAppAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveAppAuthToken(request *SaveAppAuthTokenRequest) (_result *SaveAppAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveAppAuthTokenResponse{}
	_body, _err := client.SaveAppAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainConfigWithOptions(request *SaveWebHostingCustomDomainConfigRequest, runtime *util.RuntimeOptions) (_result *SaveWebHostingCustomDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ForceRedirectType)) {
		body["ForceRedirectType"] = request.ForceRedirectType
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveWebHostingCustomDomainConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveWebHostingCustomDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainConfig(request *SaveWebHostingCustomDomainConfigRequest) (_result *SaveWebHostingCustomDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebHostingCustomDomainConfigResponse{}
	_body, _err := client.SaveWebHostingCustomDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainCorsConfigWithOptions(request *SaveWebHostingCustomDomainCorsConfigRequest, runtime *util.RuntimeOptions) (_result *SaveWebHostingCustomDomainCorsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessControlAllowOrigin)) {
		body["AccessControlAllowOrigin"] = request.AccessControlAllowOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCors)) {
		body["EnableCors"] = request.EnableCors
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveWebHostingCustomDomainCorsConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveWebHostingCustomDomainCorsConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainCorsConfig(request *SaveWebHostingCustomDomainCorsConfigRequest) (_result *SaveWebHostingCustomDomainCorsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebHostingCustomDomainCorsConfigResponse{}
	_body, _err := client.SaveWebHostingCustomDomainCorsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWechatOpenPlatformConfigWithOptions(request *SaveWechatOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *SaveWechatOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		body["AppSecret"] = request.AppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveWechatOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveWechatOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWechatOpenPlatformConfig(request *SaveWechatOpenPlatformConfigRequest) (_result *SaveWechatOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWechatOpenPlatformConfigResponse{}
	_body, _err := client.SaveWechatOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindWebHostingCustomDomainWithOptions(request *UnbindWebHostingCustomDomainRequest, runtime *util.RuntimeOptions) (_result *UnbindWebHostingCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomDomain)) {
		body["CustomDomain"] = request.CustomDomain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindWebHostingCustomDomain"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindWebHostingCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindWebHostingCustomDomain(request *UnbindWebHostingCustomDomainRequest) (_result *UnbindWebHostingCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindWebHostingCustomDomainResponse{}
	_body, _err := client.UnbindWebHostingCustomDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDingtalkOpenPlatformConfigWithOptions(request *UpdateDingtalkOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateDingtalkOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		body["AppSecret"] = request.AppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDingtalkOpenPlatformConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDingtalkOpenPlatformConfig(request *UpdateDingtalkOpenPlatformConfigRequest) (_result *UpdateDingtalkOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.UpdateDingtalkOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.HttpTriggerPath)) {
		body["HttpTriggerPath"] = request.HttpTriggerPath
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceConcurrency)) {
		body["InstanceConcurrency"] = request.InstanceConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Runtime)) {
		body["Runtime"] = request.Runtime
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.TimingTriggerConfig)) {
		body["TimingTriggerConfig"] = request.TimingTriggerConfig
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHttpTriggerConfigWithOptions(request *UpdateHttpTriggerConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateHttpTriggerConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomDomain)) {
		body["CustomDomain"] = request.CustomDomain
	}

	if !tea.BoolValue(util.IsUnset(request.CustomDomainCertificate)) {
		body["CustomDomainCertificate"] = request.CustomDomainCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.CustomDomainPrivateKey)) {
		body["CustomDomainPrivateKey"] = request.CustomDomainPrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.EnableService)) {
		body["EnableService"] = request.EnableService
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHttpTriggerConfig"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHttpTriggerConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHttpTriggerConfig(request *UpdateHttpTriggerConfigRequest) (_result *UpdateHttpTriggerConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHttpTriggerConfigResponse{}
	_body, _err := client.UpdateHttpTriggerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServicePolicyWithOptions(request *UpdateServicePolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CollectionName)) {
		body["CollectionName"] = request.CollectionName
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		body["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServicePolicy"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServicePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServicePolicy(request *UpdateServicePolicyRequest) (_result *UpdateServicePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServicePolicyResponse{}
	_body, _err := client.UpdateServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSpaceWithOptions(request *UpdateSpaceRequest, runtime *util.RuntimeOptions) (_result *UpdateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSpace"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSpace(request *UpdateSpaceRequest) (_result *UpdateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSpaceResponse{}
	_body, _err := client.UpdateSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyWebHostingDomainOwnerWithOptions(request *VerifyWebHostingDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyWebHostingDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["SpaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		body["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyWebHostingDomainOwner"),
		Version:     tea.String("2019-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyWebHostingDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyWebHostingDomainOwner(request *VerifyWebHostingDomainOwnerRequest) (_result *VerifyWebHostingDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyWebHostingDomainOwnerResponse{}
	_body, _err := client.VerifyWebHostingDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

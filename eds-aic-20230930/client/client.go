// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataImageRegionDistributeMapValue struct {
	// The status of the image distribution task.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The task is ready.
	//
	// 	- DELETE: The task is deleted.
	//
	// 	- INIT: The task is being initialized.
	//
	// 	- CREATE_FAILED: The task failed to be created.
	//
	// 	- CREATING: The task is being created.
	//
	// example:
	//
	// AVAILABLE
	DistributeStatus *string `json:"DistributeStatus,omitempty" xml:"DistributeStatus,omitempty"`
	// The distribution progress of the image.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DataImageRegionDistributeMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataImageRegionDistributeMapValue) GoString() string {
	return s.String()
}

func (s *DataImageRegionDistributeMapValue) SetDistributeStatus(v string) *DataImageRegionDistributeMapValue {
	s.DistributeStatus = &v
	return s
}

func (s *DataImageRegionDistributeMapValue) SetProgress(v string) *DataImageRegionDistributeMapValue {
	s.Progress = &v
	return s
}

type AttachKeyPairRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
}

func (s AttachKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *AttachKeyPairRequest) SetInstanceIds(v []*string) *AttachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairId(v string) *AttachKeyPairRequest {
	s.KeyPairId = &v
	return s
}

type AttachKeyPairResponseBody struct {
	Data *AttachKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBody) SetData(v *AttachKeyPairResponseBodyData) *AttachKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *AttachKeyPairResponseBody) SetRequestId(v string) *AttachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type AttachKeyPairResponseBodyData struct {
	AttachedInstanceIds []*string `json:"AttachedInstanceIds,omitempty" xml:"AttachedInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AttachKeyPairResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyData) SetAttachedInstanceIds(v []*string) *AttachKeyPairResponseBodyData {
	s.AttachedInstanceIds = v
	return s
}

func (s *AttachKeyPairResponseBodyData) SetFailCount(v int32) *AttachKeyPairResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *AttachKeyPairResponseBodyData) SetKeyPairId(v string) *AttachKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *AttachKeyPairResponseBodyData) SetTotalCount(v int32) *AttachKeyPairResponseBodyData {
	s.TotalCount = &v
	return s
}

type AttachKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairResponse) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponse) SetHeaders(v map[string]*string) *AttachKeyPairResponse {
	s.Headers = v
	return s
}

func (s *AttachKeyPairResponse) SetStatusCode(v int32) *AttachKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachKeyPairResponse) SetBody(v *AttachKeyPairResponseBody) *AttachKeyPairResponse {
	s.Body = v
	return s
}

type AuthorizeAndroidInstanceRequest struct {
	// List of instance IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// User ID to be assigned.
	//
	// example:
	//
	// test
	AuthorizeUserId *string `json:"AuthorizeUserId,omitempty" xml:"AuthorizeUserId,omitempty"`
	// User ID to be unassigned.
	//
	// example:
	//
	// test
	UnAuthorizeUserId *string `json:"UnAuthorizeUserId,omitempty" xml:"UnAuthorizeUserId,omitempty"`
}

func (s AuthorizeAndroidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeAndroidInstanceRequest) SetAndroidInstanceIds(v []*string) *AuthorizeAndroidInstanceRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *AuthorizeAndroidInstanceRequest) SetAuthorizeUserId(v string) *AuthorizeAndroidInstanceRequest {
	s.AuthorizeUserId = &v
	return s
}

func (s *AuthorizeAndroidInstanceRequest) SetUnAuthorizeUserId(v string) *AuthorizeAndroidInstanceRequest {
	s.UnAuthorizeUserId = &v
	return s
}

type AuthorizeAndroidInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeAndroidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeAndroidInstanceResponseBody) SetRequestId(v string) *AuthorizeAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeAndroidInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeAndroidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeAndroidInstanceResponse) SetHeaders(v map[string]*string) *AuthorizeAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeAndroidInstanceResponse) SetStatusCode(v int32) *AuthorizeAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeAndroidInstanceResponse) SetBody(v *AuthorizeAndroidInstanceResponseBody) *AuthorizeAndroidInstanceResponse {
	s.Body = v
	return s
}

type BackupFileRequest struct {
	// The IDs of the instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Whether all data is to be backed up.
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// Backup file name.
	//
	// example:
	//
	// defaultBackupFile
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The OSS path of the backup file.
	//
	// >  To upload a backup file to an OSS bucket, you must obtain the name of the bucket. When calling the describeBuckets operation to retrieve a bucket name, you must also call the ossObjectList operation to obtain the object key. Combine these to form the full path: oss://${bucketName}/${key}.
	//
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The description of the backup file.
	//
	// example:
	//
	// This is a backup file description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// List of apps to be backed up.
	SourceAppList []*string `json:"SourceAppList,omitempty" xml:"SourceAppList,omitempty" type:"Repeated"`
	// The paths to the source files.
	SourceFilePathList []*string `json:"SourceFilePathList,omitempty" xml:"SourceFilePathList,omitempty" type:"Repeated"`
	// The endpoint of the OSS bucket to which you want to upload the backup file.
	//
	// > : When calling the DescribeBuckets operation to query buckets, retrieve the IntranetEndpoint value if the cloud phone and the OSS bucket are in the same region. If they are in different regions, retrieve the ExtranetEndpoint value instead.
	//
	// example:
	//
	// oss-cn-shanghai-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the backup.
	//
	// Valid values:
	//
	// 	- OSS: uploads the backup file to an OSS bucket.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s BackupFileRequest) String() string {
	return tea.Prettify(s)
}

func (s BackupFileRequest) GoString() string {
	return s.String()
}

func (s *BackupFileRequest) SetAndroidInstanceIdList(v []*string) *BackupFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *BackupFileRequest) SetBackupAll(v bool) *BackupFileRequest {
	s.BackupAll = &v
	return s
}

func (s *BackupFileRequest) SetBackupFileName(v string) *BackupFileRequest {
	s.BackupFileName = &v
	return s
}

func (s *BackupFileRequest) SetBackupFilePath(v string) *BackupFileRequest {
	s.BackupFilePath = &v
	return s
}

func (s *BackupFileRequest) SetDescription(v string) *BackupFileRequest {
	s.Description = &v
	return s
}

func (s *BackupFileRequest) SetSourceAppList(v []*string) *BackupFileRequest {
	s.SourceAppList = v
	return s
}

func (s *BackupFileRequest) SetSourceFilePathList(v []*string) *BackupFileRequest {
	s.SourceFilePathList = v
	return s
}

func (s *BackupFileRequest) SetUploadEndpoint(v string) *BackupFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *BackupFileRequest) SetUploadType(v string) *BackupFileRequest {
	s.UploadType = &v
	return s
}

type BackupFileResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The objects that are returned.
	//
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	Data []*BackupFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The batch task ID.
	//
	// example:
	//
	// t-22ex666a5mco5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *BackupFileResponseBody) SetCount(v int64) *BackupFileResponseBody {
	s.Count = &v
	return s
}

func (s *BackupFileResponseBody) SetData(v []*BackupFileResponseBodyData) *BackupFileResponseBody {
	s.Data = v
	return s
}

func (s *BackupFileResponseBody) SetRequestId(v string) *BackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackupFileResponseBody) SetTaskId(v string) *BackupFileResponseBody {
	s.TaskId = &v
	return s
}

type BackupFileResponseBodyData struct {
	// Instance id.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// Backup file id.
	//
	// example:
	//
	// bf-b0qbg3pbpjkn7****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// Backup file name.
	//
	// example:
	//
	// a-58ftsoo90p0qa****.ab
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-22ex666a5mco5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackupFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BackupFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *BackupFileResponseBodyData) SetAndroidInstanceId(v string) *BackupFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *BackupFileResponseBodyData) SetBackupFileId(v string) *BackupFileResponseBodyData {
	s.BackupFileId = &v
	return s
}

func (s *BackupFileResponseBodyData) SetBackupFileName(v string) *BackupFileResponseBodyData {
	s.BackupFileName = &v
	return s
}

func (s *BackupFileResponseBodyData) SetTaskId(v string) *BackupFileResponseBodyData {
	s.TaskId = &v
	return s
}

type BackupFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackupFileResponse) String() string {
	return tea.Prettify(s)
}

func (s BackupFileResponse) GoString() string {
	return s.String()
}

func (s *BackupFileResponse) SetHeaders(v map[string]*string) *BackupFileResponse {
	s.Headers = v
	return s
}

func (s *BackupFileResponse) SetStatusCode(v int32) *BackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *BackupFileResponse) SetBody(v *BackupFileResponseBody) *BackupFileResponse {
	s.Body = v
	return s
}

type BatchGetAcpConnectionTicketRequest struct {
	// example:
	//
	// user
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// ag-25nt4kk9whjh****
	InstanceGroupId *string                                            `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	InstanceIds     []*string                                          `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	InstanceTasks   []*BatchGetAcpConnectionTicketRequestInstanceTasks `json:"InstanceTasks,omitempty" xml:"InstanceTasks,omitempty" type:"Repeated"`
}

func (s BatchGetAcpConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAcpConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketRequest) SetEndUserId(v string) *BatchGetAcpConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) SetInstanceGroupId(v string) *BatchGetAcpConnectionTicketRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) SetInstanceIds(v []*string) *BatchGetAcpConnectionTicketRequest {
	s.InstanceIds = v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) SetInstanceTasks(v []*BatchGetAcpConnectionTicketRequestInstanceTasks) *BatchGetAcpConnectionTicketRequest {
	s.InstanceTasks = v
	return s
}

type BatchGetAcpConnectionTicketRequestInstanceTasks struct {
	// example:
	//
	// acp-fkuit0cmyfvzz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou@c9f5c2e8-f5c4-4b01-8602-000cae94****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchGetAcpConnectionTicketRequestInstanceTasks) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAcpConnectionTicketRequestInstanceTasks) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) SetInstanceId(v string) *BatchGetAcpConnectionTicketRequestInstanceTasks {
	s.InstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) SetTaskId(v string) *BatchGetAcpConnectionTicketRequestInstanceTasks {
	s.TaskId = &v
	return s
}

type BatchGetAcpConnectionTicketResponseBody struct {
	InstanceConnectionModels []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels `json:"InstanceConnectionModels,omitempty" xml:"InstanceConnectionModels,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetAcpConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAcpConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketResponseBody) SetInstanceConnectionModels(v []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) *BatchGetAcpConnectionTicketResponseBody {
	s.InstanceConnectionModels = v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBody) SetRequestId(v string) *BatchGetAcpConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

type BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels struct {
	// example:
	//
	// aig-1uzb6heg797z3****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// acp-ajxvwo1u0hqvd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou@c9f5c2e8-f5c4-4b01-8602-000cae94****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// FINISHED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// piVE58_AdmVSVW7SEW3*AE5*p8mmO5gvItsNOmv4S_f_cNpoU_BOTwChTBoNM1ZJeedfK9zxYnbN5hossqIZCr6t7SGxRigm2Cb4fGaCdBZWIzmgdHq6sXXZQg4KFWufyvpeV*0*Cm58slMT1tJw3****
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetAppInstanceGroupId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.InstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetTaskId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.TaskId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetTaskStatus(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.TaskStatus = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetTicket(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.Ticket = &v
	return s
}

type BatchGetAcpConnectionTicketResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetAcpConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetAcpConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetAcpConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketResponse) SetHeaders(v map[string]*string) *BatchGetAcpConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *BatchGetAcpConnectionTicketResponse) SetStatusCode(v int32) *BatchGetAcpConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponse) SetBody(v *BatchGetAcpConnectionTicketResponseBody) *BatchGetAcpConnectionTicketResponse {
	s.Body = v
	return s
}

type CheckResourceStockRequest struct {
	// Specification ID.
	//
	// example:
	//
	// acp.basic.small
	AcpSpecId *string `json:"AcpSpecId,omitempty" xml:"AcpSpecId,omitempty"`
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId     *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	GpuAcceleration *bool   `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The availability zone of the resource.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckResourceStockRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceStockRequest) GoString() string {
	return s.String()
}

func (s *CheckResourceStockRequest) SetAcpSpecId(v string) *CheckResourceStockRequest {
	s.AcpSpecId = &v
	return s
}

func (s *CheckResourceStockRequest) SetAmount(v int32) *CheckResourceStockRequest {
	s.Amount = &v
	return s
}

func (s *CheckResourceStockRequest) SetBizRegionId(v string) *CheckResourceStockRequest {
	s.BizRegionId = &v
	return s
}

func (s *CheckResourceStockRequest) SetGpuAcceleration(v bool) *CheckResourceStockRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CheckResourceStockRequest) SetZoneId(v string) *CheckResourceStockRequest {
	s.ZoneId = &v
	return s
}

type CheckResourceStockResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 805D8FB6-512A-531C-9E4D-2A807D3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of resource inventory.
	ResourceStockModels []*CheckResourceStockResponseBodyResourceStockModels `json:"ResourceStockModels,omitempty" xml:"ResourceStockModels,omitempty" type:"Repeated"`
}

func (s CheckResourceStockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceStockResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourceStockResponseBody) SetRequestId(v string) *CheckResourceStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourceStockResponseBody) SetResourceStockModels(v []*CheckResourceStockResponseBodyResourceStockModels) *CheckResourceStockResponseBody {
	s.ResourceStockModels = v
	return s
}

type CheckResourceStockResponseBodyResourceStockModels struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Inventory status of the instance group.
	//
	// example:
	//
	// Available
	StockStatus *string `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
	// Zone ID.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckResourceStockResponseBodyResourceStockModels) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceStockResponseBodyResourceStockModels) GoString() string {
	return s.String()
}

func (s *CheckResourceStockResponseBodyResourceStockModels) SetRegionId(v string) *CheckResourceStockResponseBodyResourceStockModels {
	s.RegionId = &v
	return s
}

func (s *CheckResourceStockResponseBodyResourceStockModels) SetStockStatus(v string) *CheckResourceStockResponseBodyResourceStockModels {
	s.StockStatus = &v
	return s
}

func (s *CheckResourceStockResponseBodyResourceStockModels) SetZoneId(v string) *CheckResourceStockResponseBodyResourceStockModels {
	s.ZoneId = &v
	return s
}

type CheckResourceStockResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResourceStockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResourceStockResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceStockResponse) GoString() string {
	return s.String()
}

func (s *CheckResourceStockResponse) SetHeaders(v map[string]*string) *CheckResourceStockResponse {
	s.Headers = v
	return s
}

func (s *CheckResourceStockResponse) SetStatusCode(v int32) *CheckResourceStockResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResourceStockResponse) SetBody(v *CheckResourceStockResponseBody) *CheckResourceStockResponse {
	s.Body = v
	return s
}

type CreateAndroidInstanceGroupRequest struct {
	// The number of instance groups. Default value: 1. Maximum value: 1.
	//
	// example:
	//
	// 8
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment. Default value: false.
	//
	// Valid values:
	//
	// 	- true: enables automatic payment. Make sure that your Alibaba Cloud account has sufficient balance.
	//
	// 	- false: disables automatic payment. You must manually complete the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Default value: false.
	//
	// Valid values:
	//
	// 	- true: automatically renew resource upon expiration.
	//
	// 	- false: manually renew resources upon expiration.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query the regions where Cloud Phone is supported.
	//
	// Valid values:
	//
	// 	- cn-shenzhen: China (Shenzhen).
	//
	// 	- cn-beijing: China (Beijing).
	//
	// 	- cn-shanghai: China (Shanghai).
	//
	// 	- cn-hongkong: China (Hong Kong).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value cannot exceed 100 characters in length.
	//
	// example:
	//
	// asadbuvwiabdbvchjsbj
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EnableIpv6  *bool   `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Specifies whether to enable GPU acceleration.
	//
	// 	- true
	//
	// 	- false (true)
	//
	// example:
	//
	// false
	GpuAcceleration *bool `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// The ID of the image. You can call the [DescribeImageList](https://help.aliyun.com/document_detail/2807324.html) operation to query images.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the instance group.
	//
	// > The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with http:// or https://.
	//
	// example:
	//
	// defaultInstanceGroup
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The specifications of the instance group. You can call the [DescribeSpec](https://help.aliyun.com/document_detail/2807299.html) operation to query the available specifications.
	//
	// Valid values:
	//
	// 	- acp.perf.large: Performance (8 vCPUs, 16 GiB of memory, and 32 GiB of storage.
	//
	// 	- acp.basic.small: Lightweight (2 vCPUs, 4 GiB of memory, and 32 GiB of storage).
	//
	// 	- acp.std.large: Standard (4 vCPUs, 8 GiB of memory, and 32 GiB of storage).
	//
	// This parameter is required.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	Ipv6Bandwidth     *int32  `json:"Ipv6Bandwidth,omitempty" xml:"Ipv6Bandwidth,omitempty"`
	// The ID of the key pair. When you create an instance group and specify a valid key pair ID, all cloud phone instances within the group will automatically be bound to that key pair upon creation. This eliminates the need to manually call the operation to bind key pairs to individual cloud phone instances.
	//
	// Take note that binding key pairs to cloud phone instances is currently not supported during instance group resizing.
	//
	// example:
	//
	// kp-7o9xywwfutc1l****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The number of cloud phones in the instance group. Maximum value: 100.
	//
	// example:
	//
	// 1
	NumberOfInstances *int32 `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The ID of the network.
	//
	// 	- This parameter is required if you assign a shared network to cloud phones. You can go to the [Network](https://wya.wuying.aliyun.com/network) page of the Cloud Phone console to retrieve the ID of a **shared network**. If no shared network is available in the Cloud Phone console, you can leave this parameter empty. The system automatically creates one when you create an instance group.
	//
	// 	- This parameter is required if you assign a virtual private cloud (VPC) to cloud phones. You can go to the [Network](https://wya.wuying.aliyun.com/network) page of the Cloud Phone console to retrieve the ID of a **VPC**. If no VPC is available in the Cloud Phone console, you must first create one.
	//
	// example:
	//
	// cn-hangzhou+dir-745976****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The subscription duration. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Hour (Note that this unit is supported only by pay-as-you-go.)
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy. You can call the [ListPolicyGroups](https://help.aliyun.com/document_detail/2807352.html) operation to query policies.
	//
	// example:
	//
	// pg-b7bxrrwxkijjh****
	PolicyGroupId *string                                 `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	Tag           []*CreateAndroidInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/448774.html) operation to query vSwitches.
	//
	// 	- This parameter is not required if you assign a shared network to cloud phones.
	//
	// 	- This parameter is required if you assign a VPC to cloud phones. The vSwitch specified by this parameter is used to create cloud phones.
	//
	// example:
	//
	// vsw-uf61uvzhz8ejaw776****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateAndroidInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupRequest) SetAmount(v int32) *CreateAndroidInstanceGroupRequest {
	s.Amount = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetAutoPay(v bool) *CreateAndroidInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetAutoRenew(v bool) *CreateAndroidInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetBizRegionId(v string) *CreateAndroidInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetChargeType(v string) *CreateAndroidInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetClientToken(v string) *CreateAndroidInstanceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetEnableIpv6(v bool) *CreateAndroidInstanceGroupRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetGpuAcceleration(v bool) *CreateAndroidInstanceGroupRequest {
	s.GpuAcceleration = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetImageId(v string) *CreateAndroidInstanceGroupRequest {
	s.ImageId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetInstanceGroupName(v string) *CreateAndroidInstanceGroupRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetInstanceGroupSpec(v string) *CreateAndroidInstanceGroupRequest {
	s.InstanceGroupSpec = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetIpv6Bandwidth(v int32) *CreateAndroidInstanceGroupRequest {
	s.Ipv6Bandwidth = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetKeyPairId(v string) *CreateAndroidInstanceGroupRequest {
	s.KeyPairId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetNumberOfInstances(v int32) *CreateAndroidInstanceGroupRequest {
	s.NumberOfInstances = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetOfficeSiteId(v string) *CreateAndroidInstanceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetPeriod(v int32) *CreateAndroidInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetPeriodUnit(v string) *CreateAndroidInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetPolicyGroupId(v string) *CreateAndroidInstanceGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetTag(v []*CreateAndroidInstanceGroupRequestTag) *CreateAndroidInstanceGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateAndroidInstanceGroupRequest) SetVSwitchId(v string) *CreateAndroidInstanceGroupRequest {
	s.VSwitchId = &v
	return s
}

type CreateAndroidInstanceGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAndroidInstanceGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateAndroidInstanceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupRequestTag) SetKey(v string) *CreateAndroidInstanceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAndroidInstanceGroupRequestTag) SetValue(v string) *CreateAndroidInstanceGroupRequestTag {
	s.Value = &v
	return s
}

type CreateAndroidInstanceGroupResponseBody struct {
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The instance groups.
	InstanceGroupInfos []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos `json:"InstanceGroupInfos,omitempty" xml:"InstanceGroupInfos,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 22365781890****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAndroidInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponseBody) SetInstanceGroupIds(v []*string) *CreateAndroidInstanceGroupResponseBody {
	s.InstanceGroupIds = v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetInstanceGroupInfos(v []*CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) *CreateAndroidInstanceGroupResponseBody {
	s.InstanceGroupInfos = v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetOrderId(v string) *CreateAndroidInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBody) SetRequestId(v string) *CreateAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos struct {
	// The ID of the instance group.
	//
	// example:
	//
	// ag-cuv4scs4obxch****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The IDs of the instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) SetInstanceGroupId(v string) *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos {
	s.InstanceGroupId = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos) SetInstanceIds(v []*string) *CreateAndroidInstanceGroupResponseBodyInstanceGroupInfos {
	s.InstanceIds = v
	return s
}

type CreateAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndroidInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAndroidInstanceGroupResponse) SetStatusCode(v int32) *CreateAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponse) SetBody(v *CreateAndroidInstanceGroupResponseBody) *CreateAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	// The name of the application.
	//
	// example:
	//
	// Application Name 1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId   *string                        `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	CustomAppInfo *CreateAppRequestCustomAppInfo `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty" type:"Struct"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name used by the application file in OSS. This parameter, combined with `FilePath`, uniquely identifies the OSS path of the application file.
	//
	// >
	//
	// 	- Log on to the [Cloud Phone console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the value of this parameter.
	//
	// 	- If you do not specify `OssAppUrl`, you must specify `FileName` and `FilePath`.
	//
	// example:
	//
	// testApp.apk
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS bucket path to the application file. This parameter, combined with `FileName`, uniquely identifies the OSS path of the application file.
	//
	// >
	//
	// 	- Log on to the [Cloud Phone console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the value of this parameter.
	//
	// 	- If you do not specify `OssAppUrl`, you must specify `FileName` and `FilePath`.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://www.example.com/icon.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The parameters used for installing the application. By default, the `-r` parameter is included when you install an application.
	//
	// example:
	//
	// -d
	InstallParam *string `json:"InstallParam,omitempty" xml:"InstallParam,omitempty"`
	// The endpoint of the OSS bucket to which you want to upload the application file.
	//
	// >
	//
	// 	- Log on to the [Cloud Phone console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the value of this parameter.
	//
	// 	- If you do not specify `FileName` or `FilePath`, you must specify this parameter.
	//
	// example:
	//
	// http://testApp.apk
	OssAppUrl *string `json:"OssAppUrl,omitempty" xml:"OssAppUrl,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetBizRegionId(v string) *CreateAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppRequest) SetCustomAppInfo(v *CreateAppRequestCustomAppInfo) *CreateAppRequest {
	s.CustomAppInfo = v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetFileName(v string) *CreateAppRequest {
	s.FileName = &v
	return s
}

func (s *CreateAppRequest) SetFilePath(v string) *CreateAppRequest {
	s.FilePath = &v
	return s
}

func (s *CreateAppRequest) SetIconUrl(v string) *CreateAppRequest {
	s.IconUrl = &v
	return s
}

func (s *CreateAppRequest) SetInstallParam(v string) *CreateAppRequest {
	s.InstallParam = &v
	return s
}

func (s *CreateAppRequest) SetOssAppUrl(v string) *CreateAppRequest {
	s.OssAppUrl = &v
	return s
}

type CreateAppRequestCustomAppInfo struct {
	ApkSize     *string `json:"ApkSize,omitempty" xml:"ApkSize,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateAppRequestCustomAppInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestCustomAppInfo) GoString() string {
	return s.String()
}

func (s *CreateAppRequestCustomAppInfo) SetApkSize(v string) *CreateAppRequestCustomAppInfo {
	s.ApkSize = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetDownloadUrl(v string) *CreateAppRequestCustomAppInfo {
	s.DownloadUrl = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetMd5(v string) *CreateAppRequestCustomAppInfo {
	s.Md5 = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetPackageName(v string) *CreateAppRequestCustomAppInfo {
	s.PackageName = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetVersion(v string) *CreateAppRequestCustomAppInfo {
	s.Version = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetVersionCode(v string) *CreateAppRequestCustomAppInfo {
	s.VersionCode = &v
	return s
}

type CreateAppShrinkRequest struct {
	// The name of the application.
	//
	// example:
	//
	// Application Name 1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId         *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	CustomAppInfoShrink *string `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name used by the application file in OSS. This parameter, combined with `FilePath`, uniquely identifies the OSS path of the application file.
	//
	// >
	//
	// 	- Log on to the [Cloud Phone console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the value of this parameter.
	//
	// 	- If you do not specify `OssAppUrl`, you must specify `FileName` and `FilePath`.
	//
	// example:
	//
	// testApp.apk
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS bucket path to the application file. This parameter, combined with `FileName`, uniquely identifies the OSS path of the application file.
	//
	// >
	//
	// 	- Log on to the [Cloud Phone console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the value of this parameter.
	//
	// 	- If you do not specify `OssAppUrl`, you must specify `FileName` and `FilePath`.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://www.example.com/icon.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The parameters used for installing the application. By default, the `-r` parameter is included when you install an application.
	//
	// example:
	//
	// -d
	InstallParam *string `json:"InstallParam,omitempty" xml:"InstallParam,omitempty"`
	// The endpoint of the OSS bucket to which you want to upload the application file.
	//
	// >
	//
	// 	- Log on to the [Cloud Phone console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the value of this parameter.
	//
	// 	- If you do not specify `FileName` or `FilePath`, you must specify this parameter.
	//
	// example:
	//
	// http://testApp.apk
	OssAppUrl *string `json:"OssAppUrl,omitempty" xml:"OssAppUrl,omitempty"`
}

func (s CreateAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppShrinkRequest) SetAppName(v string) *CreateAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppShrinkRequest) SetBizRegionId(v string) *CreateAppShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppShrinkRequest) SetCustomAppInfoShrink(v string) *CreateAppShrinkRequest {
	s.CustomAppInfoShrink = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDescription(v string) *CreateAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAppShrinkRequest) SetFileName(v string) *CreateAppShrinkRequest {
	s.FileName = &v
	return s
}

func (s *CreateAppShrinkRequest) SetFilePath(v string) *CreateAppShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *CreateAppShrinkRequest) SetIconUrl(v string) *CreateAppShrinkRequest {
	s.IconUrl = &v
	return s
}

func (s *CreateAppShrinkRequest) SetInstallParam(v string) *CreateAppShrinkRequest {
	s.InstallParam = &v
	return s
}

func (s *CreateAppShrinkRequest) SetOssAppUrl(v string) *CreateAppShrinkRequest {
	s.OssAppUrl = &v
	return s
}

type CreateAppResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// 1234
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetAppId(v int32) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateCustomImageRequest struct {
	// Idempotent parameter. Default is empty, with a maximum length of 100 characters.
	//
	// example:
	//
	// 20393E53-8FF1-524C-B494-B478A5369733
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Image description.
	//
	// example:
	//
	// create for cc5g group auth rules test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom image name
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCustomImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequest) SetClientToken(v string) *CreateCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomImageRequest) SetDescription(v string) *CreateCustomImageRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomImageRequest) SetImageName(v string) *CreateCustomImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateCustomImageRequest) SetInstanceId(v string) *CreateCustomImageRequest {
	s.InstanceId = &v
	return s
}

type CreateCustomImageResponseBody struct {
	// Image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 20393E53-8FF1-524C-B494-B478A5369733
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponseBody) SetImageId(v string) *CreateCustomImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateCustomImageResponseBody) SetRequestId(v string) *CreateCustomImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponse) SetHeaders(v map[string]*string) *CreateCustomImageResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomImageResponse) SetStatusCode(v int32) *CreateCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomImageResponse) SetBody(v *CreateCustomImageResponseBody) *CreateCustomImageResponse {
	s.Body = v
	return s
}

type CreateKeyPairRequest struct {
	// The name of the key pair. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s CreateKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

type CreateKeyPairResponseBody struct {
	// The objects that are returned.
	Data *CreateKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) SetData(v *CreateKeyPairResponseBodyData) *CreateKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *CreateKeyPairResponseBody) SetRequestId(v string) *CreateKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateKeyPairResponseBodyData struct {
	// The time when the key pair was created.
	//
	// example:
	//
	// 2024-06-30 08:45:09.0
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3*****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key of the key pair. The PEM-encoded private key that is in PKCS#8 format and adheres to the ADB connection specification.
	//
	// example:
	//
	// MIIEpAIBAAKCAQEAtReyMzLIcBH78EV2zj****
	PrivateKeyBody *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
}

func (s CreateKeyPairResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBodyData) SetGmtCreated(v string) *CreateKeyPairResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) SetKeyPairId(v string) *CreateKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) SetKeyPairName(v string) *CreateKeyPairResponseBodyData {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) SetPrivateKeyBody(v string) *CreateKeyPairResponseBodyData {
	s.PrivateKeyBody = &v
	return s
}

type CreateKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponse) SetHeaders(v map[string]*string) *CreateKeyPairResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyPairResponse) SetStatusCode(v int32) *CreateKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyPairResponse) SetBody(v *CreateKeyPairResponseBody) *CreateKeyPairResponse {
	s.Body = v
	return s
}

type CreatePolicyGroupRequest struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: read and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy of the Alibaba Cloud Workspace web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicy *CreatePolicyGroupRequestNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s CreatePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) SetCameraRedirect(v string) *CreatePolicyGroupRequest {
	s.CameraRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLockResolution(v string) *CreatePolicyGroupRequest {
	s.LockResolution = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetNetRedirectPolicy(v *CreatePolicyGroupRequestNetRedirectPolicy) *CreatePolicyGroupRequest {
	s.NetRedirectPolicy = v
	return s
}

func (s *CreatePolicyGroupRequest) SetPolicyGroupName(v string) *CreatePolicyGroupRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetResolutionHeight(v int32) *CreatePolicyGroupRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetResolutionWidth(v int32) *CreatePolicyGroupRequest {
	s.ResolutionWidth = &v
	return s
}

type CreatePolicyGroupRequestNetRedirectPolicy struct {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	CustomProxy   *string                                           `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	HostAddr      *string                                           `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	NetRedirect   *string                                           `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	Port          *string                                           `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyPassword *string                                           `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	ProxyType     *string                                           `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	ProxyUserName *string                                           `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	Rules         []*CreatePolicyGroupRequestNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
=======
=======
>>>>>>> Stashed changes
	// Specifies whether to manually configure a custom proxy.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CustomProxy *string `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	// The IPv4 address of the custom proxy.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddr *string `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	// Specifies whether to enable the network redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The port of the custom proxy. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1145
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The password of the proxy. The password must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// password
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The type of the proxy protocol.
	//
	// Valid values:
	//
	// 	- socks5.
	//
	// example:
	//
	// socks5
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The username of the proxy. The name must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// username
	ProxyUserName *string `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	// The proxy rules. You can create up to 100 proxy rules.
	Rules []*CreatePolicyGroupRequestNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
}

func (s CreatePolicyGroupRequestNetRedirectPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestNetRedirectPolicy) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetCustomProxy(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.CustomProxy = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetHostAddr(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.HostAddr = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetNetRedirect(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.NetRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetPort(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.Port = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetProxyPassword(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.ProxyPassword = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetProxyType(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.ProxyType = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetProxyUserName(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.ProxyUserName = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetRules(v []*CreatePolicyGroupRequestNetRedirectPolicyRules) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.Rules = v
	return s
}

type CreatePolicyGroupRequestNetRedirectPolicyRules struct {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Target   *string `json:"Target,omitempty" xml:"Target,omitempty"`
=======
=======
>>>>>>> Stashed changes
	// The type of the rule.
	//
	// Valid values:
	//
	// 	- prc: an application package name.
	//
	// 	- domain: a domain name.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The name of the application package or domain name.
	//
	// example:
	//
	// *.example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
}

func (s CreatePolicyGroupRequestNetRedirectPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestNetRedirectPolicyRules) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) SetRuleType(v string) *CreatePolicyGroupRequestNetRedirectPolicyRules {
	s.RuleType = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) SetTarget(v string) *CreatePolicyGroupRequestNetRedirectPolicyRules {
	s.Target = &v
	return s
}

type CreatePolicyGroupShrinkRequest struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: read and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy of the Alibaba Cloud Workspace web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicyShrink *string `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s CreatePolicyGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupShrinkRequest) SetCameraRedirect(v string) *CreatePolicyGroupShrinkRequest {
	s.CameraRedirect = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetClipboard(v string) *CreatePolicyGroupShrinkRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupShrinkRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetLocalDrive(v string) *CreatePolicyGroupShrinkRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetLockResolution(v string) *CreatePolicyGroupShrinkRequest {
	s.LockResolution = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetNetRedirectPolicyShrink(v string) *CreatePolicyGroupShrinkRequest {
	s.NetRedirectPolicyShrink = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetPolicyGroupName(v string) *CreatePolicyGroupShrinkRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetResolutionHeight(v int32) *CreatePolicyGroupShrinkRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetResolutionWidth(v int32) *CreatePolicyGroupShrinkRequest {
	s.ResolutionWidth = &v
	return s
}

type CreatePolicyGroupResponseBody struct {
	// The ID of the policy.
	//
	// example:
	//
	// pg-exbuu6yrpvb******
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponseBody) SetPolicyGroupId(v string) *CreatePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *CreatePolicyGroupResponseBody) SetRequestId(v string) *CreatePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponse) SetHeaders(v map[string]*string) *CreatePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyGroupResponse) SetStatusCode(v int32) *CreatePolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyGroupResponse) SetBody(v *CreatePolicyGroupResponseBody) *CreatePolicyGroupResponse {
	s.Body = v
	return s
}

type CreateScreenshotRequest struct {
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// cloudphone-saved-bucket-cn-shanghai-default
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// false
	SkipCheckPolicyConfig *string `json:"SkipCheckPolicyConfig,omitempty" xml:"SkipCheckPolicyConfig,omitempty"`
}

func (s CreateScreenshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenshotRequest) GoString() string {
	return s.String()
}

func (s *CreateScreenshotRequest) SetAndroidInstanceIdList(v []*string) *CreateScreenshotRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *CreateScreenshotRequest) SetOssBucketName(v string) *CreateScreenshotRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateScreenshotRequest) SetSkipCheckPolicyConfig(v string) *CreateScreenshotRequest {
	s.SkipCheckPolicyConfig = &v
	return s
}

type CreateScreenshotResponseBody struct {
	// example:
	//
	// 3AF82CE1-2801-52CE-BF64-B491DD7C****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*CreateScreenshotResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CreateScreenshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScreenshotResponseBody) SetRequestId(v string) *CreateScreenshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScreenshotResponseBody) SetTasks(v []*CreateScreenshotResponseBodyTasks) *CreateScreenshotResponseBody {
	s.Tasks = v
	return s
}

type CreateScreenshotResponseBodyTasks struct {
	// example:
	//
	// acp-bwhtebzah2fse****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateScreenshotResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenshotResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *CreateScreenshotResponseBodyTasks) SetAndroidInstanceId(v string) *CreateScreenshotResponseBodyTasks {
	s.AndroidInstanceId = &v
	return s
}

func (s *CreateScreenshotResponseBodyTasks) SetTaskId(v string) *CreateScreenshotResponseBodyTasks {
	s.TaskId = &v
	return s
}

type CreateScreenshotResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScreenshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScreenshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenshotResponse) GoString() string {
	return s.String()
}

func (s *CreateScreenshotResponse) SetHeaders(v map[string]*string) *CreateScreenshotResponse {
	s.Headers = v
	return s
}

func (s *CreateScreenshotResponse) SetStatusCode(v int32) *CreateScreenshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScreenshotResponse) SetBody(v *CreateScreenshotResponseBody) *CreateScreenshotResponse {
	s.Body = v
	return s
}

type DeleteAndroidInstanceGroupRequest struct {
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
}

func (s DeleteAndroidInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAndroidInstanceGroupRequest) SetInstanceGroupIds(v []*string) *DeleteAndroidInstanceGroupRequest {
	s.InstanceGroupIds = v
	return s
}

type DeleteAndroidInstanceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB95E410-FD1D-53C5-9F7D-93CC44D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAndroidInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAndroidInstanceGroupResponseBody) SetRequestId(v string) *DeleteAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAndroidInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *DeleteAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAndroidInstanceGroupResponse) SetStatusCode(v int32) *DeleteAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAndroidInstanceGroupResponse) SetBody(v *DeleteAndroidInstanceGroupResponseBody) *DeleteAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

type DeleteAppsRequest struct {
	// The IDs of the applications.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
}

func (s DeleteAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppsRequest) SetAppIdList(v []*string) *DeleteAppsRequest {
	s.AppIdList = v
	return s
}

type DeleteAppsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 310A783E-CC46-5452-A8A3-71AE5DB5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppsResponseBody) SetRequestId(v string) *DeleteAppsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppsResponse) SetHeaders(v map[string]*string) *DeleteAppsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppsResponse) SetStatusCode(v int32) *DeleteAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppsResponse) SetBody(v *DeleteAppsResponseBody) *DeleteAppsResponse {
	s.Body = v
	return s
}

type DeleteImagesRequest struct {
	// The IDs of the images.
	//
	// This parameter is required.
	ImageIds []*string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty" type:"Repeated"`
}

func (s DeleteImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) SetImageIds(v []*string) *DeleteImagesRequest {
	s.ImageIds = v
	return s
}

type DeleteImagesShrinkRequest struct {
	// The IDs of the images.
	//
	// This parameter is required.
	ImageIdsShrink *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
}

func (s DeleteImagesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesShrinkRequest) SetImageIdsShrink(v string) *DeleteImagesShrinkRequest {
	s.ImageIdsShrink = &v
	return s
}

type DeleteImagesResponseBody struct {
	// The images.
	Data *DeleteImagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4610632D-D661-5982-B3D7-5D3FD183F595
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBody) SetData(v *DeleteImagesResponseBodyData) *DeleteImagesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImagesResponseBody) SetRequestId(v string) *DeleteImagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImagesResponseBodyData struct {
	// The IDs of the images that failed to be deleted.
	FailDeleteImageIds []*string `json:"FailDeleteImageIds,omitempty" xml:"FailDeleteImageIds,omitempty" type:"Repeated"`
	// The IDs of the images that are successfully deleted.
	SuccessDeleteImageIds []*string `json:"SuccessDeleteImageIds,omitempty" xml:"SuccessDeleteImageIds,omitempty" type:"Repeated"`
}

func (s DeleteImagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBodyData) SetFailDeleteImageIds(v []*string) *DeleteImagesResponseBodyData {
	s.FailDeleteImageIds = v
	return s
}

func (s *DeleteImagesResponseBodyData) SetSuccessDeleteImageIds(v []*string) *DeleteImagesResponseBodyData {
	s.SuccessDeleteImageIds = v
	return s
}

type DeleteImagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) SetHeaders(v map[string]*string) *DeleteImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesResponse) SetStatusCode(v int32) *DeleteImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImagesResponse) SetBody(v *DeleteImagesResponseBody) *DeleteImagesResponse {
	s.Body = v
	return s
}

type DeleteKeyPairsRequest struct {
	// The IDs of the ADB key pairs.
	KeyPairIds []*string `json:"KeyPairIds,omitempty" xml:"KeyPairIds,omitempty" type:"Repeated"`
}

func (s DeleteKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) SetKeyPairIds(v []*string) *DeleteKeyPairsRequest {
	s.KeyPairIds = v
	return s
}

type DeleteKeyPairsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponseBody) SetRequestId(v string) *DeleteKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeyPairsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponse) SetHeaders(v map[string]*string) *DeleteKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyPairsResponse) SetStatusCode(v int32) *DeleteKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyPairsResponse) SetBody(v *DeleteKeyPairsResponseBody) *DeleteKeyPairsResponse {
	s.Body = v
	return s
}

type DeletePolicyGroupRequest struct {
	// The IDs of the policies.
	//
	// This parameter is required.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
}

func (s DeletePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupRequest) SetPolicyGroupIds(v []*string) *DeletePolicyGroupRequest {
	s.PolicyGroupIds = v
	return s
}

type DeletePolicyGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 17C731AB-AAEE-5844-A352-D8D0352D3F0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupResponseBody) SetRequestId(v string) *DeletePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupResponse) SetHeaders(v map[string]*string) *DeletePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyGroupResponse) SetStatusCode(v int32) *DeletePolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyGroupResponse) SetBody(v *DeletePolicyGroupResponseBody) *DeletePolicyGroupResponse {
	s.Body = v
	return s
}

type DescribeAndroidInstanceGroupsRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The name of the instance group. Instance groups support fuzzy search by name.
	//
	// example:
	//
	// defaultInstanceGroup
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-5htf0ymsrnb7q****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The maximum number of entries per page. Value range: 0 to 100. Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-1b77w6xrqfubi****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- standard
	//
	// example:
	//
	// standard
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The status of the instance group.
	//
	// Valid values:
	//
	// 	- UPDATING_FAILED: The image update for the instance group failed.
	//
	// 	- FAILED: The instance group failed to be created.
	//
	// 	- RUNNING: The instance group is available.
	//
	// 	- EXPIRED: The instance group expired.
	//
	// 	- DELETING: The instance group is being deleted.
	//
	// 	- DELETED: The instance group is deleted.
	//
	// 	- UPDATING: The instance group is undergoing an image update.
	//
	// 	- CREATING: The instance group is being created.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAndroidInstanceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsRequest) SetBizRegionId(v string) *DescribeAndroidInstanceGroupsRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetChargeType(v string) *DescribeAndroidInstanceGroupsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetInstanceGroupIds(v []*string) *DescribeAndroidInstanceGroupsRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetInstanceGroupName(v string) *DescribeAndroidInstanceGroupsRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetKeyPairId(v string) *DescribeAndroidInstanceGroupsRequest {
	s.KeyPairId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetMaxResults(v int32) *DescribeAndroidInstanceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetNextToken(v string) *DescribeAndroidInstanceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetPolicyGroupId(v string) *DescribeAndroidInstanceGroupsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetSaleMode(v string) *DescribeAndroidInstanceGroupsRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsRequest) SetStatus(v string) *DescribeAndroidInstanceGroupsRequest {
	s.Status = &v
	return s
}

type DescribeAndroidInstanceGroupsResponseBody struct {
	// The details of the instance group.
	InstanceGroupModel []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel `json:"InstanceGroupModel,omitempty" xml:"InstanceGroupModel,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetInstanceGroupModel(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) *DescribeAndroidInstanceGroupsResponseBody {
	s.InstanceGroupModel = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetNextToken(v string) *DescribeAndroidInstanceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetRequestId(v string) *DescribeAndroidInstanceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetTotalCount(v int32) *DescribeAndroidInstanceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-48xr63m4dybjk****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The type of the architecture.
	//
	// example:
	//
	// ARM
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// Number of instances.
	//
	// example:
	//
	// 5
	AvailableInstanceAmount *int32 `json:"AvailableInstanceAmount,omitempty" xml:"AvailableInstanceAmount,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 8
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disks.
	Disks []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The time when the instance group was created.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the subscription instance group expires.
	//
	// example:
	//
	// 2027-06-29 07:25:31
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the instance group was updated.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The list of installed applications.
	//
	// example:
	//
	// "TikTok","WeChat"
	InstalledAppList *string `json:"InstalledAppList,omitempty" xml:"InstalledAppList,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-h67a2cs0zprfdh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// defaultInstanceGroup
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The specifications of the instance group.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	// The description of the instance group specifications.
	//
	// example:
	//
	// ARM-2vCPU4GiB 32GiB
	InstanceGroupSpecDescribe *string `json:"InstanceGroupSpecDescribe,omitempty" xml:"InstanceGroupSpecDescribe,omitempty"`
	// The status of the instance group.
	//
	// example:
	//
	// RUNNING
	InstanceGroupStatus *string `json:"InstanceGroupStatus,omitempty" xml:"InstanceGroupStatus,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 8
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of instances in the instance group.
	//
	// example:
	//
	// 10
	NumberOfInstances *string `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// cn-shanghai+dir-030598****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-c6n38xucps8kl****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rendering type.
	//
	// example:
	//
	// CPU
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	// The height of the resolution.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The sales mode.
	//
	// example:
	//
	// standard
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The version of the operating system.
	//
	// example:
	//
	// Android 12
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-t4n0yqs009ho024wt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetAppInstanceGroupId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetArchitectureType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetAvailableInstanceAmount(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.AvailableInstanceAmount = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetChargeType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetCpu(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Cpu = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetDisks(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Disks = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetErrorCode(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetGmtCreate(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetGmtExpired(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.GmtExpired = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetGmtModified(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetImageId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ImageId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstalledAppList(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstalledAppList = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupName(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupSpec(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupSpec = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupSpecDescribe(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupSpecDescribe = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupStatus(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupStatus = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetMemory(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Memory = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetNumberOfInstances(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.NumberOfInstances = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetOfficeSiteId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetPolicyGroupId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetRegionId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.RegionId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetRenderingType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.RenderingType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetResolutionHeight(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetResolutionWidth(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetSaleMode(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.SaleMode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetSystemVersion(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.SystemVersion = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetVSwitchId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.VSwitchId = &v
	return s
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks struct {
	// The size of the disk. Unit: GB.
	//
	// example:
	//
	// 32
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The type of the disk.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) SetDiskSize(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) SetDiskType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks {
	s.DiskType = &v
	return s
}

type DescribeAndroidInstanceGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAndroidInstanceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponse) SetHeaders(v map[string]*string) *DescribeAndroidInstanceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponse) SetStatusCode(v int32) *DescribeAndroidInstanceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponse) SetBody(v *DescribeAndroidInstanceGroupsResponseBody) *DescribeAndroidInstanceGroupsResponse {
	s.Body = v
	return s
}

type DescribeAndroidInstancesRequest struct {
	// The IDs of the instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// name
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query the regions where Cloud Phone is supported.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-25nt4kk9whjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The name of the instance group.
	//
	// example:
	//
	// test
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The ID of the bound key pair.
	//
	// example:
	//
	// kp-5hh431emkpuoi****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	NodeId    *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName  *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
=======
	// The ID of the node.
	//
>>>>>>> Stashed changes
=======
	// The ID of the node.
	//
>>>>>>> Stashed changes
	// example:
	//
	// node_id
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- Instance: the standard mode.
	//
	// 	- Node: the node mode.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The state of the instance.
	//
	// Valid values:
	//
	// 	- BACKUPING: The instance is being backed up.
	//
	// 	- STARTING: The instance is being started.
	//
	// 	- RUNNING: The instance group is available.
	//
	// 	- DELETING: The instance is being deleted.
	//
	// 	- BACKUP_FAILED: The backup operation failed.
	//
	// 	- DELETED: The instance is deleted.
	//
	// 	- FAILED: The instance failed to be created.
	//
	// 	- STOPPED: The instance is stopped.
	//
	// 	- RECOVERING: The instance has an ongoing file recovery task.
	//
	// 	- UNAVAILABLE: The instance has an exception.
	//
	// 	- REBOOTING: The instance is being restarted.
	//
	// 	- RESETTING: The instance is being reset.
	//
	// 	- STOPPING: The instance is being stopped.
	//
	// 	- RECOVER_FAILED: The file recovery task failed.
	//
	// 	- CREATING: The instance is being created.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resources.
	Tag []*DescribeAndroidInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAndroidInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesRequest) SetAndroidInstanceIds(v []*string) *DescribeAndroidInstancesRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetAndroidInstanceName(v string) *DescribeAndroidInstancesRequest {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetBizRegionId(v string) *DescribeAndroidInstancesRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetChargeType(v string) *DescribeAndroidInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetInstanceGroupId(v string) *DescribeAndroidInstancesRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetInstanceGroupIds(v []*string) *DescribeAndroidInstancesRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetInstanceGroupName(v string) *DescribeAndroidInstancesRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetKeyPairId(v string) *DescribeAndroidInstancesRequest {
	s.KeyPairId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetMaxResults(v int32) *DescribeAndroidInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetNextToken(v string) *DescribeAndroidInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetNodeId(v string) *DescribeAndroidInstancesRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetNodeName(v string) *DescribeAndroidInstancesRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetSaleMode(v string) *DescribeAndroidInstancesRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetStatus(v string) *DescribeAndroidInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetTag(v []*DescribeAndroidInstancesRequestTag) *DescribeAndroidInstancesRequest {
	s.Tag = v
	return s
}

type DescribeAndroidInstancesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesRequestTag) SetKey(v string) *DescribeAndroidInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstancesRequestTag) SetValue(v string) *DescribeAndroidInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeAndroidInstancesResponseBody struct {
	// The instances.
	InstanceModel []*DescribeAndroidInstancesResponseBodyInstanceModel `json:"InstanceModel,omitempty" xml:"InstanceModel,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kmma/xxE9WtwL/ADvZ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAndroidInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBody) SetInstanceModel(v []*DescribeAndroidInstancesResponseBodyInstanceModel) *DescribeAndroidInstancesResponseBody {
	s.InstanceModel = v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) SetNextToken(v string) *DescribeAndroidInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) SetRequestId(v string) *DescribeAndroidInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) SetTotalCount(v int32) *DescribeAndroidInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAndroidInstancesResponseBodyInstanceModel struct {
	// The ID of the instance group.
	//
	// example:
	//
	// ag-ayyhomlal7po****
	AndroidInstanceGroupId *string `json:"AndroidInstanceGroupId,omitempty" xml:"AndroidInstanceGroupId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// AndroidInstanceGroupName
	AndroidInstanceGroupName *string `json:"AndroidInstanceGroupName,omitempty" xml:"AndroidInstanceGroupName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// acp-8at8h6ejkadjh****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// name
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// The state of the instance.
	//
	// example:
	//
	// RUNNING
	AndroidInstanceStatus *string `json:"AndroidInstanceStatus,omitempty" xml:"AndroidInstanceStatus,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-i7yv6tkn7kh8dv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the physical instance.
	//
	// example:
	//
	// ai-9ey6io0q58rcd****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the user to whom the instance is assigned.
	//
	// example:
	//
	// test
	AuthorizedUserId *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The ID of the bound user.
	//
	// example:
	//
	// test
	BindUserId *string `json:"BindUserId,omitempty" xml:"BindUserId,omitempty"`
	// The billing method of the instance.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disks.
	Disks []*DescribeAndroidInstancesResponseBodyInstanceModelDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The cause of the instance data backup failure or restoration failure.
	//
	// example:
	//
	// FilePathNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-05-06 10:42:10
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the subscription instance group expires.
	//
	// example:
	//
	// 2024-07-15T02:03:33Z
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the instance was modified.
	//
	// example:
	//
	// 2023-05-06 10:42:10
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// 3.5.3.867
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// acp.basic.small
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-5hh431emkpucs****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The IP address of the ENI.
	//
	// example:
	//
	// 192.168.22.48
	NetworkInterfaceIp          *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	NetworkInterfaceIpv6Address *string `json:"NetworkInterfaceIpv6Address,omitempty" xml:"NetworkInterfaceIpv6Address,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// cn-shenzhen+dir-211620****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the persistent session.
	//
	// example:
	//
	// p-0btrd5zj8epo****
	PersistentAppInstanceId *string `json:"PersistentAppInstanceId,omitempty" xml:"PersistentAppInstanceId,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-0bszojpu0seql****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 10.32.1.41
	PublicIpAddress   *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	PublicIpv6Address *string `json:"PublicIpv6Address,omitempty" xml:"PublicIpv6Address,omitempty"`
	// The progress of instance data backup or restoration.
	//
	// example:
	//
	// 100
	Rate *int32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	RegionId      *string                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenderingType *string                                                  `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	SessionStatus *string                                                  `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	Tags          []*DescribeAndroidInstancesResponseBodyInstanceModelTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
=======
=======
>>>>>>> Stashed changes
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rendering type.
	//
	// example:
	//
	// local
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	// The status of the session connection.
	//
	// 	- connect
	//
	// 	- disConnect
	//
	// example:
	//
	// connect
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The tags.
	Tags []*DescribeAndroidInstancesResponseBodyInstanceModelTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
}

func (s DescribeAndroidInstancesResponseBodyInstanceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModel) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceGroupId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceGroupName(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceName(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceStatus(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceStatus = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAppInstanceGroupId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAppInstanceId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AppInstanceId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAuthorizedUserId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetBindUserId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.BindUserId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetChargeType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetCpu(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Cpu = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetDisks(v []*DescribeAndroidInstancesResponseBodyInstanceModelDisks) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Disks = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetErrorCode(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetGmtCreate(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetGmtExpired(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.GmtExpired = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetGmtModified(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetImageVersion(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ImageVersion = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetInstanceType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.InstanceType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetKeyPairId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.KeyPairId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetMemory(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Memory = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetNetworkInterfaceIp(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetNetworkInterfaceIpv6Address(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.NetworkInterfaceIpv6Address = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetOfficeSiteId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPersistentAppInstanceId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PersistentAppInstanceId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPolicyGroupId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPublicIpAddress(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPublicIpv6Address(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PublicIpv6Address = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetRate(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Rate = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetRegionId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.RegionId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetRenderingType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.RenderingType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetSessionStatus(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.SessionStatus = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetTags(v []*DescribeAndroidInstancesResponseBodyInstanceModelTags) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Tags = v
	return s
}

type DescribeAndroidInstancesResponseBodyInstanceModelDisks struct {
	// The size of the disk. Unit: GB.
	//
	// example:
	//
	// 32
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The type of the disk.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisks) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) SetDiskSize(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) SetDiskType(v string) *DescribeAndroidInstancesResponseBodyInstanceModelDisks {
	s.DiskType = &v
	return s
}

type DescribeAndroidInstancesResponseBodyInstanceModelTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelTags) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) SetKey(v string) *DescribeAndroidInstancesResponseBodyInstanceModelTags {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) SetValue(v string) *DescribeAndroidInstancesResponseBodyInstanceModelTags {
	s.Value = &v
	return s
}

type DescribeAndroidInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAndroidInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAndroidInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponse) SetHeaders(v map[string]*string) *DescribeAndroidInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAndroidInstancesResponse) SetStatusCode(v int32) *DescribeAndroidInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAndroidInstancesResponse) SetBody(v *DescribeAndroidInstancesResponseBody) *DescribeAndroidInstancesResponse {
	s.Body = v
	return s
}

type DescribeAppsRequest struct {
	// The IDs of the applications.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The name of the application.
	//
	// example:
	//
	// defaultAppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Region id.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The installation/uninstallation status of the application.
	//
	// Valid values:
	//
	// 	- INSTALLFAILED: The application failed to be installed.
	//
	// 	- UNINSTALLING: The application is being uninstalled.
	//
	// 	- INSTALLING: The application is being installed.
	//
	// 	- UNINSTALLED: The application is uninstalled.
	//
	// 	- INSTALLED: The application is installed.
	//
	// 	- UNINSTALLFAILED: The application failed to be uninstalled.
	//
	// example:
	//
	// INSTALLING
	InstallationStatus *string `json:"InstallationStatus,omitempty" xml:"InstallationStatus,omitempty"`
	// The value of MD5.
	//
	// example:
	//
	// THCIEH73KEK3334
	MD5 *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the application.
	//
	// Valid values:
	//
	// 	- FAILED: The application failed to be created.
	//
	// 	- NORMAL: The application is available.
	//
	// 	- CREATING: The application is being created.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) SetAppIdList(v []*string) *DescribeAppsRequest {
	s.AppIdList = v
	return s
}

func (s *DescribeAppsRequest) SetAppName(v string) *DescribeAppsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAppsRequest) SetBizRegionId(v string) *DescribeAppsRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAppsRequest) SetInstallationStatus(v string) *DescribeAppsRequest {
	s.InstallationStatus = &v
	return s
}

func (s *DescribeAppsRequest) SetMD5(v string) *DescribeAppsRequest {
	s.MD5 = &v
	return s
}

func (s *DescribeAppsRequest) SetMaxResults(v int32) *DescribeAppsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAppsRequest) SetNextToken(v string) *DescribeAppsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAppsRequest) SetStatus(v string) *DescribeAppsRequest {
	s.Status = &v
	return s
}

type DescribeAppsResponseBody struct {
	// The objects that are returned.
	Data []*DescribeAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB95E410-FD1D-53C5-9F7D-93CC44D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetData(v []*DescribeAppsResponseBodyData) *DescribeAppsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAppsResponseBody) SetNextToken(v string) *DescribeAppsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalCount(v string) *DescribeAppsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppsResponseBodyData struct {
	// The version of the application.
	//
	// example:
	//
	// 1.0.0
	AndroidAppVersion *string `json:"AndroidAppVersion,omitempty" xml:"AndroidAppVersion,omitempty"`
	// Apk size.
	//
	// example:
	//
	// 10244893
	ApkSize *string `json:"ApkSize,omitempty" xml:"ApkSize,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 10404
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Region id.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the application was last modified.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://test.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The installation/uninstallation status of the application.
	//
	// Valid values:
	//
	// 	- INSTALLFAILED: The application failed to be installed.
	//
	// 	- UNINSTALLING: The application is being uninstalled.
	//
	// 	- INSTALLING: The application is being installed.
	//
	// 	- UNINSTALLED: The application is uninstalled.
	//
	// 	- INSTALLED: The application is installed.
	//
	// 	- UNINSTALLFAILED: The application failed to be uninstalled.
	//
	// example:
	//
	// INSTALLING
	InstallationStatus *string `json:"InstallationStatus,omitempty" xml:"InstallationStatus,omitempty"`
	// The list of instance groups where the application is installed.
	InstanceGroupList []*string `json:"InstanceGroupList,omitempty" xml:"InstanceGroupList,omitempty" type:"Repeated"`
	// The value of MD5.
	//
	// example:
	//
	// THCIEH73KEK3334
	MD5 *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	// The name of the application package.
	//
	// example:
	//
	// cn.rdstar.rdstarandroid
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The status of the application.
	//
	// Valid values:
	//
	// 	- FAILED: The application failed to be created.
	//
	// 	- NORMAL: The application is available.
	//
	// 	- CREATING: The application is being created.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyData) SetAndroidAppVersion(v string) *DescribeAppsResponseBodyData {
	s.AndroidAppVersion = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetApkSize(v string) *DescribeAppsResponseBodyData {
	s.ApkSize = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetAppId(v int32) *DescribeAppsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetAppName(v string) *DescribeAppsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetBizRegionId(v string) *DescribeAppsResponseBodyData {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetDescription(v string) *DescribeAppsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetGmtCreate(v string) *DescribeAppsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetGmtModified(v string) *DescribeAppsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetIconUrl(v string) *DescribeAppsResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetInstallationStatus(v string) *DescribeAppsResponseBodyData {
	s.InstallationStatus = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetInstanceGroupList(v []*string) *DescribeAppsResponseBodyData {
	s.InstanceGroupList = v
	return s
}

func (s *DescribeAppsResponseBodyData) SetMD5(v string) *DescribeAppsResponseBodyData {
	s.MD5 = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetPackageName(v string) *DescribeAppsResponseBodyData {
	s.PackageName = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetStatus(v string) *DescribeAppsResponseBodyData {
	s.Status = &v
	return s
}

type DescribeAppsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetStatusCode(v int32) *DescribeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeBackupFilesRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The name of the instance. Instances support fuzzy search by name.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// Is all data to be backed up.
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-dxrh5jrv0zpb8****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The name of the backup file. Backup files support fuzzy search by name.
	//
	// example:
	//
	// defaulBackupFile
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The description of the backup file. Backup files support fuzzy search by description.
	//
	// example:
	//
	// default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end of the period for querying generated backup files.
	//
	// example:
	//
	// 2024-05-20 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The owner of the backup file.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-fxdx91jsfyiy3****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the period for querying generated backup files.
	//
	// example:
	//
	// 2024-05-23 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of backup file status.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeBackupFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesRequest) SetAndroidInstanceId(v string) *DescribeBackupFilesRequest {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetAndroidInstanceName(v string) *DescribeBackupFilesRequest {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetBackupAll(v bool) *DescribeBackupFilesRequest {
	s.BackupAll = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetBackupFileId(v string) *DescribeBackupFilesRequest {
	s.BackupFileId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetBackupFileName(v string) *DescribeBackupFilesRequest {
	s.BackupFileName = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetDescription(v string) *DescribeBackupFilesRequest {
	s.Description = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetEndTime(v string) *DescribeBackupFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetEndUserId(v string) *DescribeBackupFilesRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetInstanceGroupId(v string) *DescribeBackupFilesRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetMaxResults(v int64) *DescribeBackupFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetNextToken(v string) *DescribeBackupFilesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetStartTime(v string) *DescribeBackupFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetStatusList(v []*string) *DescribeBackupFilesRequest {
	s.StatusList = v
	return s
}

type DescribeBackupFilesResponseBody struct {
	// The backup files that are returned.
	Data []*DescribeBackupFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request. If the request fails, provide this ID to technical support to assist in diagnosing the issue.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 91
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBody) SetData(v []*DescribeBackupFilesResponseBodyData) *DescribeBackupFilesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetMaxResults(v string) *DescribeBackupFilesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetNextToken(v string) *DescribeBackupFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetRequestId(v string) *DescribeBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetTotalCount(v int64) *DescribeBackupFilesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupFilesResponseBodyData struct {
	// The ID of the instance.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// defaultInstanceName
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// Is all data to be backed up.
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-b0qbg3pbpjkn7****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The name of the backup file.
	//
	// example:
	//
	// a-58ftsoo90p0qa****.ab
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The directory in which the backup file is stored.
	//
	// example:
	//
	// oss://cloudphone-saved-bucket-cn-shanghai/backup/aic-58ftsoo90p0qa****.ab
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The description of the backup file.
	//
	// example:
	//
	// This is default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The owner of the backup file.
	//
	// example:
	//
	// test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The total size of the source files.
	//
	// example:
	//
	// 10227168
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The time when the backup file was created.
	//
	// example:
	//
	// 2024-05-15 17:33:59
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the backup file was last updated.
	//
	// example:
	//
	// 2024-05-15 17:33:59
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-58ftsoo90p0qi****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of apps to be backed up.
	SourceAppInfoList []*string `json:"SourceAppInfoList,omitempty" xml:"SourceAppInfoList,omitempty" type:"Repeated"`
	// The directories of the source files.
	SourceFilePathList []*string `json:"SourceFilePathList,omitempty" xml:"SourceFilePathList,omitempty" type:"Repeated"`
	// The status of the backup file.
	//
	// Valid values:
	//
	// 	- AVAILABLE
	//
	// 	- RECOVERING
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The endpoint of the OSS bucket that stores the backup file.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the backup.
	//
	// Valid values:
	//
	// 	- OSS: backup files are stored in OSS buckets. .
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s DescribeBackupFilesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBodyData) SetAndroidInstanceId(v string) *DescribeBackupFilesResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetAndroidInstanceName(v string) *DescribeBackupFilesResponseBodyData {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupAll(v bool) *DescribeBackupFilesResponseBodyData {
	s.BackupAll = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupFileId(v string) *DescribeBackupFilesResponseBodyData {
	s.BackupFileId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupFileName(v string) *DescribeBackupFilesResponseBodyData {
	s.BackupFileName = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupFilePath(v string) *DescribeBackupFilesResponseBodyData {
	s.BackupFilePath = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetDescription(v string) *DescribeBackupFilesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetEndUserId(v string) *DescribeBackupFilesResponseBodyData {
	s.EndUserId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetFileSize(v int64) *DescribeBackupFilesResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetGmtCreated(v string) *DescribeBackupFilesResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetGmtModified(v string) *DescribeBackupFilesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetInstanceGroupId(v string) *DescribeBackupFilesResponseBodyData {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetRegionId(v string) *DescribeBackupFilesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetSourceAppInfoList(v []*string) *DescribeBackupFilesResponseBodyData {
	s.SourceAppInfoList = v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetSourceFilePathList(v []*string) *DescribeBackupFilesResponseBodyData {
	s.SourceFilePathList = v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetStatus(v string) *DescribeBackupFilesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetTaskId(v string) *DescribeBackupFilesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetUploadEndpoint(v string) *DescribeBackupFilesResponseBodyData {
	s.UploadEndpoint = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetUploadType(v string) *DescribeBackupFilesResponseBodyData {
	s.UploadType = &v
	return s
}

type DescribeBackupFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponse) SetHeaders(v map[string]*string) *DescribeBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupFilesResponse) SetStatusCode(v int32) *DescribeBackupFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupFilesResponse) SetBody(v *DescribeBackupFilesResponseBody) *DescribeBackupFilesResponse {
	s.Body = v
	return s
}

type DescribeImageListRequest struct {
	// The ID of the image.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	ImageId          *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName        *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImagePackageType *string `json:"ImagePackageType,omitempty" xml:"ImagePackageType,omitempty"`
=======
=======
>>>>>>> Stashed changes
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// Android 12 image
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Image package type.
	//
	// example:
	//
	// VM
	ImagePackageType *string `json:"ImagePackageType,omitempty" xml:"ImagePackageType,omitempty"`
	// The type of the image.
	//
	// Valid values:
	//
	// 	- User: custom images.
	//
	// 	- System: system images.
	//
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	// This parameter is required.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The maximum number of entries per page. Value range: 1 to 100. Default value: 20.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The state of the image.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The image is available.
	//
	// 	- DELETE: The image is deleted.
	//
	// 	- INIT: The image is being initialized.
	//
	// 	- CREATE_FAILED: The image failed to be created.
	//
	// 	- CREATING: The image is being created.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageListRequest) SetImageId(v string) *DescribeImageListRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageListRequest) SetImageName(v string) *DescribeImageListRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImageListRequest) SetImagePackageType(v string) *DescribeImageListRequest {
	s.ImagePackageType = &v
	return s
}

func (s *DescribeImageListRequest) SetImageType(v string) *DescribeImageListRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImageListRequest) SetMaxResults(v int32) *DescribeImageListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImageListRequest) SetNextToken(v string) *DescribeImageListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImageListRequest) SetStatus(v string) *DescribeImageListRequest {
	s.Status = &v
	return s
}

type DescribeImageListResponseBody struct {
	// The objects that are returned.
	Data []*DescribeImageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 620740FF-492F-5956-B1BA-361E966C0269
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponseBody) SetData(v []*DescribeImageListResponseBodyData) *DescribeImageListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageListResponseBody) SetNextToken(v string) *DescribeImageListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImageListResponseBody) SetRequestId(v string) *DescribeImageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageListResponseBody) SetTotalCount(v int32) *DescribeImageListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeImageListResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 117819727354****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the image was last modified.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// imgc-075cllfeuazh****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// IMAGE
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The region where the image is distributed. The key is the region and the value is the distribution information.
	ImageRegionDistributeMap map[string]*DataImageRegionDistributeMapValue `json:"ImageRegionDistributeMap,omitempty" xml:"ImageRegionDistributeMap,omitempty"`
	// The list of regions.
	ImageRegionList []*string `json:"ImageRegionList,omitempty" xml:"ImageRegionList,omitempty" type:"Repeated"`
	// The type of the image.
	//
	// Valid values:
	//
	// 	- User: custom images.
	//
	// 	- System: system images.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The language of the image.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time when the image was published.
	//
	// example:
	//
	// 2024-07-25 10:06:45
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The rendering type.
	//
	// Valid values:
	//
	// 	- GPURemote
	//
	// 	- CPU
	//
	// 	- GPULocal
	//
	// example:
	//
	// CPU
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	// The state of the image.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The image is available.
	//
	// 	- DELETE: The image is deleted.
	//
	// 	- INIT: The image is being initialized.
	//
	// 	- CREATE_FAILED: The image failed to be created.
	//
	// 	- CREATING: The image is being created.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The OS type of the image.
	//
	// example:
	//
	// Android 12
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s DescribeImageListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponseBodyData) SetAliUid(v int64) *DescribeImageListResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetDescription(v string) *DescribeImageListResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetGmtCreate(v string) *DescribeImageListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetGmtModified(v string) *DescribeImageListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageId(v string) *DescribeImageListResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageName(v string) *DescribeImageListResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageRegionDistributeMap(v map[string]*DataImageRegionDistributeMapValue) *DescribeImageListResponseBodyData {
	s.ImageRegionDistributeMap = v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageRegionList(v []*string) *DescribeImageListResponseBodyData {
	s.ImageRegionList = v
	return s
}

func (s *DescribeImageListResponseBodyData) SetImageType(v string) *DescribeImageListResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetLanguage(v string) *DescribeImageListResponseBodyData {
	s.Language = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetReleaseTime(v string) *DescribeImageListResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetRenderingType(v string) *DescribeImageListResponseBodyData {
	s.RenderingType = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetStatus(v string) *DescribeImageListResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeImageListResponseBodyData) SetSystemType(v string) *DescribeImageListResponseBodyData {
	s.SystemType = &v
	return s
}

type DescribeImageListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponse) SetHeaders(v map[string]*string) *DescribeImageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageListResponse) SetStatusCode(v int32) *DescribeImageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageListResponse) SetBody(v *DescribeImageListResponseBody) *DescribeImageListResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t-4e98eeb5****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetInstanceIds(v []*string) *DescribeInvocationsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInvocationsRequest) SetInvocationId(v string) *DescribeInvocationsRequest {
	s.InvocationId = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	Data []*DescribeInvocationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 440D7342-5E7C-B2DB-D0B4EAC2BDF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetData(v []*DescribeInvocationsResponseBodyData) *DescribeInvocationsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetTotalCount(v string) *DescribeInvocationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInvocationsResponseBodyData struct {
	// example:
	//
	// 2022-08-11 17:45:03
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// t-15775dc8****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// example:
	//
	// RUNNING
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// example:
	//
	// success
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 2022-08-11 17:45:03
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyData) SetFinishTime(v string) *DescribeInvocationsResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetInstanceId(v string) *DescribeInvocationsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetInvocationId(v string) *DescribeInvocationsResponseBodyData {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyData {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetOutput(v string) *DescribeInvocationsResponseBodyData {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetStartTime(v string) *DescribeInvocationsResponseBodyData {
	s.StartTime = &v
	return s
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetStatusCode(v int32) *DescribeInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type DescribeKeyPairsRequest struct {
	// The IDs of the ADB key pairs.
	KeyPairIds []*string `json:"KeyPairIds,omitempty" xml:"KeyPairIds,omitempty" type:"Repeated"`
	// The name of the ADB key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAYRHtOLVQzCYj17y+OP7LZQBUVVbi0GTu8g5****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequest) SetKeyPairIds(v []*string) *DescribeKeyPairsRequest {
	s.KeyPairIds = v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairName(v string) *DescribeKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetMaxResults(v int32) *DescribeKeyPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetNextToken(v string) *DescribeKeyPairsRequest {
	s.NextToken = &v
	return s
}

type DescribeKeyPairsResponseBody struct {
	// The objects that are returned.
	Data []*DescribeKeyPairsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// FFbc8N4E1iOlcSxC+8boa0HHH2LKWbggYUinyrZWvtS1oTrMYCg1HuMLGuftj0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 565FB06A-AE04-5AD0-8A32-5BA92CA5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBody) SetData(v []*DescribeKeyPairsResponseBodyData) *DescribeKeyPairsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetNextToken(v string) *DescribeKeyPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetRequestId(v string) *DescribeKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetTotalCount(v int32) *DescribeKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeKeyPairsResponseBodyData struct {
	// The time when the ADB key pair was created.
	//
	// example:
	//
	// 2022-10-11T08:53:32Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the ADB key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DescribeKeyPairsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyData) SetGmtCreated(v string) *DescribeKeyPairsResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyData) SetKeyPairId(v string) *DescribeKeyPairsResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyData) SetKeyPairName(v string) *DescribeKeyPairsResponseBodyData {
	s.KeyPairName = &v
	return s
}

type DescribeKeyPairsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponse) SetHeaders(v map[string]*string) *DescribeKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyPairsResponse) SetStatusCode(v int32) *DescribeKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeyPairsResponse) SetBody(v *DescribeKeyPairsResponseBody) *DescribeKeyPairsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	SaleMode       *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetSaleMode(v string) *DescribeRegionsRequest {
	s.SaleMode = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// Available regions.
	RegionModels []*DescribeRegionsResponseBodyRegionModels `json:"RegionModels,omitempty" xml:"RegionModels,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// A87B3769-0D05-5383-B236-5798B455****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegionModels(v []*DescribeRegionsResponseBodyRegionModels) *DescribeRegionsResponseBody {
	s.RegionModels = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegionModels struct {
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionModels) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionModels) SetRegionId(v string) *DescribeRegionsResponseBodyRegionModels {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModels) SetRegionName(v string) *DescribeRegionsResponseBodyRegionModels {
	s.RegionName = &v
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

type DescribeSpecRequest struct {
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	MatrixSpec  *string `json:"MatrixSpec,omitempty" xml:"MatrixSpec,omitempty"`
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
	// The maximum number of items to return per page in a paginated query. The value range is 1 to 100, with a default value of 100.
	//
>>>>>>> Stashed changes
=======
	// The maximum number of items to return per page in a paginated query. The value range is 1 to 100, with a default value of 100.
	//
>>>>>>> Stashed changes
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates the starting position for reading. If left empty, it starts from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SaleMode  *string   `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	SpecIds   []*string `json:"SpecIds,omitempty" xml:"SpecIds,omitempty" type:"Repeated"`
=======
=======
>>>>>>> Stashed changes
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SaleMode  *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// List of specification IDs.
	SpecIds []*string `json:"SpecIds,omitempty" xml:"SpecIds,omitempty" type:"Repeated"`
	// Specification status.
	//
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	// example:
	//
	// Available
	SpecStatus *string `json:"SpecStatus,omitempty" xml:"SpecStatus,omitempty"`
	// Specification type.
	//
	// example:
	//
	// ARM
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
}

func (s DescribeSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpecRequest) SetBizRegionId(v string) *DescribeSpecRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeSpecRequest) SetMatrixSpec(v string) *DescribeSpecRequest {
	s.MatrixSpec = &v
	return s
}

func (s *DescribeSpecRequest) SetMaxResults(v int32) *DescribeSpecRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSpecRequest) SetNextToken(v string) *DescribeSpecRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSpecRequest) SetSaleMode(v string) *DescribeSpecRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeSpecRequest) SetSpecIds(v []*string) *DescribeSpecRequest {
	s.SpecIds = v
	return s
}

func (s *DescribeSpecRequest) SetSpecStatus(v string) *DescribeSpecRequest {
	s.SpecStatus = &v
	return s
}

func (s *DescribeSpecRequest) SetSpecType(v string) *DescribeSpecRequest {
	s.SpecType = &v
	return s
}

type DescribeSpecResponseBody struct {
	// Indicates the current read position returned by this call. An empty value means that all data has been read.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// D9888DAD-331E-5FBC-B5A0-F2445115****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specification information.
	SpecInfoModel []*DescribeSpecResponseBodySpecInfoModel `json:"SpecInfoModel,omitempty" xml:"SpecInfoModel,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpecResponseBody) SetNextToken(v string) *DescribeSpecResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSpecResponseBody) SetRequestId(v string) *DescribeSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpecResponseBody) SetSpecInfoModel(v []*DescribeSpecResponseBodySpecInfoModel) *DescribeSpecResponseBody {
	s.SpecInfoModel = v
	return s
}

func (s *DescribeSpecResponseBody) SetTotalCount(v int32) *DescribeSpecResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSpecResponseBodySpecInfoModel struct {
	// Number of CPU cores.
	//
	// example:
	//
	// 8
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// Memory size.
	//
	// example:
	//
	// 16
	Memory     *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PhoneCount *string `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
	// Specification ID.
	//
>>>>>>> Stashed changes
=======
	// Specification ID.
	//
>>>>>>> Stashed changes
	// example:
	//
	// acp.basic.small
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// Specification status.
	//
	// example:
	//
	// Available
	SpecStatus *string `json:"SpecStatus,omitempty" xml:"SpecStatus,omitempty"`
	// Specification type.
	//
	// example:
	//
	// ARM
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// System disk size, in GB.
	//
	// example:
	//
	// 32
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeSpecResponseBodySpecInfoModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecResponseBodySpecInfoModel) GoString() string {
	return s.String()
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetCore(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.Core = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetMemory(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.Memory = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetPhoneCount(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.PhoneCount = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetResolution(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.Resolution = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSpecId(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.SpecId = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSpecStatus(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.SpecStatus = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSpecType(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.SpecType = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSystemDiskSize(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.SystemDiskSize = &v
	return s
}

type DescribeSpecResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpecResponse) SetHeaders(v map[string]*string) *DescribeSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpecResponse) SetStatusCode(v int32) *DescribeSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpecResponse) SetBody(v *DescribeSpecResponseBody) *DescribeSpecResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	Level    *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFbc8N4E1iOlcSxC+8boa0HHH2LKWbggYUinyrZWvtS1oTrMYCg1HuMLGuftj0****
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Param        *string   `json:"Param,omitempty" xml:"Param,omitempty"`
	ParentTaskId *string   `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	ResourceIds  []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	TaskIds      []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// Processing
	TaskStatus   *string   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatuses []*string `json:"TaskStatuses,omitempty" xml:"TaskStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// StartInstance
	TaskType  *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetInstanceId(v string) *DescribeTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTasksRequest) SetInstanceName(v string) *DescribeTasksRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeTasksRequest) SetInvokeId(v string) *DescribeTasksRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeTasksRequest) SetLevel(v int32) *DescribeTasksRequest {
	s.Level = &v
	return s
}

func (s *DescribeTasksRequest) SetMaxResults(v int32) *DescribeTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTasksRequest) SetNextToken(v string) *DescribeTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksRequest) SetParam(v string) *DescribeTasksRequest {
	s.Param = &v
	return s
}

func (s *DescribeTasksRequest) SetParentTaskId(v string) *DescribeTasksRequest {
	s.ParentTaskId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceIds(v []*string) *DescribeTasksRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeTasksRequest) SetTaskIds(v []*string) *DescribeTasksRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatus(v string) *DescribeTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatuses(v []*string) *DescribeTasksRequest {
	s.TaskStatuses = v
	return s
}

func (s *DescribeTasksRequest) SetTaskType(v string) *DescribeTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskTypes(v []*string) *DescribeTasksRequest {
	s.TaskTypes = v
	return s
}

type DescribeTasksResponseBody struct {
	Data []*DescribeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetData(v []*DescribeTasksResponseBodyData) *DescribeTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTasksResponseBody) SetNextToken(v string) *DescribeTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalCount(v int32) *DescribeTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTasksResponseBodyData struct {
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg         *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FailedChildCount *int32  `json:"FailedChildCount,omitempty" xml:"FailedChildCount,omitempty"`
	// example:
	//
	// 2022-10-11T08:53:32Z
	FinishTime     *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName   *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId     *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	Level        *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Operator     *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Param        *string `json:"Param,omitempty" xml:"Param,omitempty"`
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// acp-25nt4kk9whhok****
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Result            *string `json:"Result,omitempty" xml:"Result,omitempty"`
	RunningChildCount *int32  `json:"RunningChildCount,omitempty" xml:"RunningChildCount,omitempty"`
	// example:
	//
	// 2022-10-11T08:53:32Z
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SuccessChildCount *int32  `json:"SuccessChildCount,omitempty" xml:"SuccessChildCount,omitempty"`
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// StartInstance
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TotalChildCount *int32  `json:"TotalChildCount,omitempty" xml:"TotalChildCount,omitempty"`
}

func (s DescribeTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyData) SetErrorCode(v string) *DescribeTasksResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetErrorMsg(v string) *DescribeTasksResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetFailedChildCount(v int32) *DescribeTasksResponseBodyData {
	s.FailedChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetFinishTime(v string) *DescribeTasksResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInstanceId(v string) *DescribeTasksResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInstanceName(v string) *DescribeTasksResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInstanceStatus(v string) *DescribeTasksResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInvokeId(v string) *DescribeTasksResponseBodyData {
	s.InvokeId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetLevel(v int32) *DescribeTasksResponseBodyData {
	s.Level = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetOperator(v string) *DescribeTasksResponseBodyData {
	s.Operator = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetParam(v string) *DescribeTasksResponseBodyData {
	s.Param = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetParentTaskId(v string) *DescribeTasksResponseBodyData {
	s.ParentTaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetRegionId(v string) *DescribeTasksResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetResourceId(v string) *DescribeTasksResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetResult(v string) *DescribeTasksResponseBodyData {
	s.Result = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetRunningChildCount(v int32) *DescribeTasksResponseBodyData {
	s.RunningChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetStartTime(v string) *DescribeTasksResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetSuccessChildCount(v int32) *DescribeTasksResponseBodyData {
	s.SuccessChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTaskId(v string) *DescribeTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTaskStatus(v string) *DescribeTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTaskType(v string) *DescribeTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTotalChildCount(v int32) *DescribeTasksResponseBodyData {
	s.TotalChildCount = &v
	return s
}

type DescribeTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DetachKeyPairRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
}

func (s DetachKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DetachKeyPairRequest) SetInstanceIds(v []*string) *DetachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairId(v string) *DetachKeyPairRequest {
	s.KeyPairId = &v
	return s
}

type DetachKeyPairResponseBody struct {
	Data *DetachKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBody) SetData(v *DetachKeyPairResponseBodyData) *DetachKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *DetachKeyPairResponseBody) SetRequestId(v string) *DetachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type DetachKeyPairResponseBodyData struct {
	DetachedInstanceIds []*string `json:"DetachedInstanceIds,omitempty" xml:"DetachedInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DetachKeyPairResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyData) SetDetachedInstanceIds(v []*string) *DetachKeyPairResponseBodyData {
	s.DetachedInstanceIds = v
	return s
}

func (s *DetachKeyPairResponseBodyData) SetFailCount(v int32) *DetachKeyPairResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *DetachKeyPairResponseBodyData) SetKeyPairId(v string) *DetachKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *DetachKeyPairResponseBodyData) SetTotalCount(v int32) *DetachKeyPairResponseBodyData {
	s.TotalCount = &v
	return s
}

type DetachKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponse) SetHeaders(v map[string]*string) *DetachKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DetachKeyPairResponse) SetStatusCode(v int32) *DetachKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachKeyPairResponse) SetBody(v *DetachKeyPairResponseBody) *DetachKeyPairResponse {
	s.Body = v
	return s
}

type DistributeImageRequest struct {
	// This parameter is required.
	DistributeRegionList []*string `json:"DistributeRegionList,omitempty" xml:"DistributeRegionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DistributeImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DistributeImageRequest) GoString() string {
	return s.String()
}

func (s *DistributeImageRequest) SetDistributeRegionList(v []*string) *DistributeImageRequest {
	s.DistributeRegionList = v
	return s
}

func (s *DistributeImageRequest) SetImageId(v string) *DistributeImageRequest {
	s.ImageId = &v
	return s
}

type DistributeImageResponseBody struct {
	// example:
	//
	// 440D7342-5FC2-5E7C-B2DB-D0B4EAC2BDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DistributeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DistributeImageResponseBody) GoString() string {
	return s.String()
}

func (s *DistributeImageResponseBody) SetRequestId(v string) *DistributeImageResponseBody {
	s.RequestId = &v
	return s
}

type DistributeImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DistributeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DistributeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DistributeImageResponse) GoString() string {
	return s.String()
}

func (s *DistributeImageResponse) SetHeaders(v map[string]*string) *DistributeImageResponse {
	s.Headers = v
	return s
}

func (s *DistributeImageResponse) SetStatusCode(v int32) *DistributeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DistributeImageResponse) SetBody(v *DistributeImageResponseBody) *DistributeImageResponse {
	s.Body = v
	return s
}

type DowngradeAndroidInstanceGroupRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ag-cuv4scs4obxhs****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
}

func (s DowngradeAndroidInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DowngradeAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DowngradeAndroidInstanceGroupRequest) SetAndroidInstanceIds(v []*string) *DowngradeAndroidInstanceGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DowngradeAndroidInstanceGroupRequest) SetAutoPay(v bool) *DowngradeAndroidInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupRequest) SetInstanceGroupId(v string) *DowngradeAndroidInstanceGroupRequest {
	s.InstanceGroupId = &v
	return s
}

type DowngradeAndroidInstanceGroupResponseBody struct {
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3AF82CE1-2801-52CE-BF64-B491DD7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DowngradeAndroidInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DowngradeAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradeAndroidInstanceGroupResponseBody) SetOrderId(v string) *DowngradeAndroidInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponseBody) SetRequestId(v string) *DowngradeAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DowngradeAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradeAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradeAndroidInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DowngradeAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DowngradeAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *DowngradeAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponse) SetStatusCode(v int32) *DowngradeAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponse) SetBody(v *DowngradeAndroidInstanceGroupResponseBody) *DowngradeAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

type FetchFileRequest struct {
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// /data/a.txt
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// This parameter is required.
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s FetchFileRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchFileRequest) GoString() string {
	return s.String()
}

func (s *FetchFileRequest) SetAndroidInstanceIdList(v []*string) *FetchFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *FetchFileRequest) SetSourceFilePath(v string) *FetchFileRequest {
	s.SourceFilePath = &v
	return s
}

func (s *FetchFileRequest) SetUploadEndpoint(v string) *FetchFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *FetchFileRequest) SetUploadType(v string) *FetchFileRequest {
	s.UploadType = &v
	return s
}

func (s *FetchFileRequest) SetUploadUrl(v string) *FetchFileRequest {
	s.UploadUrl = &v
	return s
}

type FetchFileResponseBody struct {
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	Data []*FetchFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FetchFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchFileResponseBody) GoString() string {
	return s.String()
}

func (s *FetchFileResponseBody) SetData(v []*FetchFileResponseBodyData) *FetchFileResponseBody {
	s.Data = v
	return s
}

func (s *FetchFileResponseBody) SetRequestId(v string) *FetchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchFileResponseBody) SetTaskId(v string) *FetchFileResponseBody {
	s.TaskId = &v
	return s
}

type FetchFileResponseBodyData struct {
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FetchFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FetchFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchFileResponseBodyData) SetAndroidInstanceId(v string) *FetchFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *FetchFileResponseBodyData) SetTaskId(v string) *FetchFileResponseBodyData {
	s.TaskId = &v
	return s
}

type FetchFileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchFileResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchFileResponse) GoString() string {
	return s.String()
}

func (s *FetchFileResponse) SetHeaders(v map[string]*string) *FetchFileResponse {
	s.Headers = v
	return s
}

func (s *FetchFileResponse) SetStatusCode(v int32) *FetchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchFileResponse) SetBody(v *FetchFileResponseBody) *FetchFileResponse {
	s.Body = v
	return s
}

type ImportKeyPairRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABC1234567*****
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
}

func (s ImportKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

type ImportKeyPairResponseBody struct {
	Data *ImportKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) SetData(v *ImportKeyPairResponseBodyData) *ImportKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type ImportKeyPairResponseBodyData struct {
	// example:
	//
	// 2023-03-05T10:29:22Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// kp-6v2q33ae4tw3*****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s ImportKeyPairResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBodyData) SetGmtCreated(v string) *ImportKeyPairResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *ImportKeyPairResponseBodyData) SetKeyPairId(v string) *ImportKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *ImportKeyPairResponseBodyData) SetKeyPairName(v string) *ImportKeyPairResponseBodyData {
	s.KeyPairName = &v
	return s
}

type ImportKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponse) SetHeaders(v map[string]*string) *ImportKeyPairResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyPairResponse) SetStatusCode(v int32) *ImportKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeyPairResponse) SetBody(v *ImportKeyPairResponseBody) *ImportKeyPairResponse {
	s.Body = v
	return s
}

type InstallAppRequest struct {
	AppIdList           []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	InstanceGroupIdList []*string `json:"InstanceGroupIdList,omitempty" xml:"InstanceGroupIdList,omitempty" type:"Repeated"`
	InstanceIdList      []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s InstallAppRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAppRequest) GoString() string {
	return s.String()
}

func (s *InstallAppRequest) SetAppIdList(v []*string) *InstallAppRequest {
	s.AppIdList = v
	return s
}

func (s *InstallAppRequest) SetInstanceGroupIdList(v []*string) *InstallAppRequest {
	s.InstanceGroupIdList = v
	return s
}

func (s *InstallAppRequest) SetInstanceIdList(v []*string) *InstallAppRequest {
	s.InstanceIdList = v
	return s
}

type InstallAppResponseBody struct {
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InstallAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAppResponseBody) SetRequestId(v string) *InstallAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAppResponseBody) SetTaskId(v string) *InstallAppResponseBody {
	s.TaskId = &v
	return s
}

type InstallAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAppResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponse) GoString() string {
	return s.String()
}

func (s *InstallAppResponse) SetHeaders(v map[string]*string) *InstallAppResponse {
	s.Headers = v
	return s
}

func (s *InstallAppResponse) SetStatusCode(v int32) *InstallAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAppResponse) SetBody(v *InstallAppResponseBody) *InstallAppResponse {
	s.Body = v
	return s
}

type ListPolicyGroupsRequest struct {
	// The maximum number of entries per page. Value range: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the policies.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicyGroup
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
}

func (s ListPolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsRequest) SetMaxResults(v int32) *ListPolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPolicyGroupsRequest) SetNextToken(v string) *ListPolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPolicyGroupsRequest) SetPolicyGroupIds(v []*string) *ListPolicyGroupsRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *ListPolicyGroupsRequest) SetPolicyGroupName(v string) *ListPolicyGroupsRequest {
	s.PolicyGroupName = &v
	return s
}

type ListPolicyGroupsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policies.
	PolicyGroupModel []*ListPolicyGroupsResponseBodyPolicyGroupModel `json:"PolicyGroupModel,omitempty" xml:"PolicyGroupModel,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPolicyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBody) SetNextToken(v string) *ListPolicyGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPolicyGroupsResponseBody) SetPolicyGroupModel(v []*ListPolicyGroupsResponseBodyPolicyGroupModel) *ListPolicyGroupsResponseBody {
	s.PolicyGroupModel = v
	return s
}

func (s *ListPolicyGroupsResponseBody) SetRequestId(v string) *ListPolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyGroupsResponseBody) SetTotalCount(v int32) *ListPolicyGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPolicyGroupsResponseBodyPolicyGroupModel struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: read and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The time when the policy was created.
	//
	// example:
	//
	// 2024-06-04 10:28:54
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The file transfer policy of the HTML5 client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// download
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write denied.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Identifies whether the resolution is locked.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicy *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-9q6o8qpiy8opkj****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// Default Policy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution.
	//
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
	// The width of the resolution.
	//
>>>>>>> Stashed changes
=======
	// The width of the resolution.
	//
>>>>>>> Stashed changes
	// example:
	//
	// 1920
	SessionResolutionWidth *int32 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModel) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModel) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetCameraRedirect(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.CameraRedirect = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetClipboard(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.Clipboard = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetGmtCreate(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.GmtCreate = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetHtml5FileTransfer(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.Html5FileTransfer = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetLocalDrive(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.LocalDrive = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetLockResolution(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.LockResolution = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetNetRedirectPolicy(v *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.NetRedirectPolicy = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetPolicyGroupId(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.PolicyGroupId = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetPolicyGroupName(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.PolicyGroupName = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetSessionResolutionHeight(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.SessionResolutionHeight = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetSessionResolutionWidth(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.SessionResolutionWidth = &v
	return s
}

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy struct {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	CustomProxy   *string                                                               `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	HostAddr      *string                                                               `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	NetRedirect   *string                                                               `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	Port          *string                                                               `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyPassword *string                                                               `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	ProxyType     *string                                                               `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	ProxyUserName *string                                                               `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	Rules         []*ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
=======
=======
>>>>>>> Stashed changes
	// Indicates whether a custom proxy is manually configured.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CustomProxy *string `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	// The IPv4 address of the custom proxy.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddr *string `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	// Indicates whether the network redirection feature is enabled. When this feature is enabled, network traffic is automatically redirected to the on-premises network by default.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The port of the custom proxy. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1145
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The password of the proxy. The password must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// password
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The type of the proxy protocol.
	//
	// Valid values:
	//
	// 	- socks5.
	//
	// example:
	//
	// socks5
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The username of the proxy. The name must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// username
	ProxyUserName *string `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	// The proxy rules.
	Rules []*ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetCustomProxy(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.CustomProxy = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetHostAddr(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.HostAddr = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetNetRedirect(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.NetRedirect = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetPort(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.Port = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetProxyPassword(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.ProxyPassword = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetProxyType(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.ProxyType = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetProxyUserName(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.ProxyUserName = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetRules(v []*ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.Rules = v
	return s
}

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules struct {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Target   *string `json:"Target,omitempty" xml:"Target,omitempty"`
=======
=======
>>>>>>> Stashed changes
	// The type of the rule.
	//
	// Valid values:
	//
	// 	- prc: an application package name.
	//
	// 	- domain: a domain name.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The name of the application package or domain name.
	//
	// example:
	//
	// *.example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) SetRuleType(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules {
	s.RuleType = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) SetTarget(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules {
	s.Target = &v
	return s
}

type ListPolicyGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponse) SetHeaders(v map[string]*string) *ListPolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyGroupsResponse) SetStatusCode(v int32) *ListPolicyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyGroupsResponse) SetBody(v *ListPolicyGroupsResponseBody) *ListPolicyGroupsResponse {
	s.Body = v
	return s
}

type ModifyAndroidInstanceRequest struct {
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-8v5bjld0r7tkl****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The new name of the cloud phone instance.
	//
	// >  The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with http:// or https://.
	//
	// example:
	//
	// new_name
	NewAndroidInstanceName *string `json:"NewAndroidInstanceName,omitempty" xml:"NewAndroidInstanceName,omitempty"`
}

func (s ModifyAndroidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceRequest) SetAndroidInstanceId(v string) *ModifyAndroidInstanceRequest {
	s.AndroidInstanceId = &v
	return s
}

func (s *ModifyAndroidInstanceRequest) SetNewAndroidInstanceName(v string) *ModifyAndroidInstanceRequest {
	s.NewAndroidInstanceName = &v
	return s
}

type ModifyAndroidInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAndroidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceResponseBody) SetRequestId(v string) *ModifyAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAndroidInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAndroidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceResponse) SetHeaders(v map[string]*string) *ModifyAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyAndroidInstanceResponse) SetStatusCode(v int32) *ModifyAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAndroidInstanceResponse) SetBody(v *ModifyAndroidInstanceResponseBody) *ModifyAndroidInstanceResponse {
	s.Body = v
	return s
}

type ModifyAndroidInstanceGroupRequest struct {
	// The ID of the instance group.
	//
	// example:
	//
	// ag-cuv4scs4obxhs****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The new name of the instance group.
	//
	// >
	//
	// 	- The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with http:// or https://.
	//
	// example:
	//
	// newName
	NewInstanceGroupName *string `json:"NewInstanceGroupName,omitempty" xml:"NewInstanceGroupName,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-2w97kp89gnsif****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
}

func (s ModifyAndroidInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceGroupRequest) SetInstanceGroupId(v string) *ModifyAndroidInstanceGroupRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *ModifyAndroidInstanceGroupRequest) SetNewInstanceGroupName(v string) *ModifyAndroidInstanceGroupRequest {
	s.NewInstanceGroupName = &v
	return s
}

func (s *ModifyAndroidInstanceGroupRequest) SetPolicyGroupId(v string) *ModifyAndroidInstanceGroupRequest {
	s.PolicyGroupId = &v
	return s
}

type ModifyAndroidInstanceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6C83EBE3-F267-5F11-ABF8-4E7B90B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAndroidInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceGroupResponseBody) SetRequestId(v string) *ModifyAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAndroidInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *ModifyAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAndroidInstanceGroupResponse) SetStatusCode(v int32) *ModifyAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAndroidInstanceGroupResponse) SetBody(v *ModifyAndroidInstanceGroupResponseBody) *ModifyAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

type ModifyAppRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// 1234
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// defaultAppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the icon.
	//
	// example:
	//
	// https://defaultIcon.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) SetAppId(v int32) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetDescription(v string) *ModifyAppRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppRequest) SetIconUrl(v string) *ModifyAppRequest {
	s.IconUrl = &v
	return s
}

type ModifyAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 83418504-5A82-5896-A24C-B2D468F0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppResponse) SetStatusCode(v int32) *ModifyAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyKeyPairNameRequest struct {
	// The ID of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// newKeyPairName
	NewKeyPairName *string `json:"NewKeyPairName,omitempty" xml:"NewKeyPairName,omitempty"`
}

func (s ModifyKeyPairNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyKeyPairNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyKeyPairNameRequest) SetKeyPairId(v string) *ModifyKeyPairNameRequest {
	s.KeyPairId = &v
	return s
}

func (s *ModifyKeyPairNameRequest) SetNewKeyPairName(v string) *ModifyKeyPairNameRequest {
	s.NewKeyPairName = &v
	return s
}

type ModifyKeyPairNameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyKeyPairNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyKeyPairNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyKeyPairNameResponseBody) SetRequestId(v string) *ModifyKeyPairNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyKeyPairNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyKeyPairNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyKeyPairNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyKeyPairNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyKeyPairNameResponse) SetHeaders(v map[string]*string) *ModifyKeyPairNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyKeyPairNameResponse) SetStatusCode(v int32) *ModifyKeyPairNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyKeyPairNameResponse) SetBody(v *ModifyKeyPairNameResponseBody) *ModifyKeyPairNameResponse {
	s.Body = v
	return s
}

type ModifyPolicyGroupRequest struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy of the Alibaba Cloud Workspace web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicy *ModifyPolicyGroupRequestNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-4bi18ebi9tfjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicyGroup
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) SetCameraRedirect(v string) *ModifyPolicyGroupRequest {
	s.CameraRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLockResolution(v string) *ModifyPolicyGroupRequest {
	s.LockResolution = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetNetRedirectPolicy(v *ModifyPolicyGroupRequestNetRedirectPolicy) *ModifyPolicyGroupRequest {
	s.NetRedirectPolicy = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupName(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetResolutionHeight(v int32) *ModifyPolicyGroupRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetResolutionWidth(v int32) *ModifyPolicyGroupRequest {
	s.ResolutionWidth = &v
	return s
}

type ModifyPolicyGroupRequestNetRedirectPolicy struct {
	// Specifies whether to manually configure a custom proxy.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CustomProxy *string `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	// The IPv4 address of the custom proxy.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddr *string `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	// Specifies whether to enable network redirection.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The port of the custom proxy. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1145
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The password of the proxy. The password must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// password
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The type of the proxy protocol.
	//
	// Valid values:
	//
	// 	- socks5.
	//
	// example:
	//
	// socks5
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The username of the proxy. The name must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// username
	ProxyUserName *string `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
}

func (s ModifyPolicyGroupRequestNetRedirectPolicy) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestNetRedirectPolicy) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetCustomProxy(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.CustomProxy = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetHostAddr(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.HostAddr = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetNetRedirect(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.NetRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetPort(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.Port = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetProxyPassword(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.ProxyPassword = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetProxyType(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.ProxyType = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetProxyUserName(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.ProxyUserName = &v
	return s
}

type ModifyPolicyGroupShrinkRequest struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy of the Alibaba Cloud Workspace web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicyShrink *string `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-4bi18ebi9tfjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicyGroup
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s ModifyPolicyGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupShrinkRequest) SetCameraRedirect(v string) *ModifyPolicyGroupShrinkRequest {
	s.CameraRedirect = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetClipboard(v string) *ModifyPolicyGroupShrinkRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupShrinkRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetLocalDrive(v string) *ModifyPolicyGroupShrinkRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetLockResolution(v string) *ModifyPolicyGroupShrinkRequest {
	s.LockResolution = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetNetRedirectPolicyShrink(v string) *ModifyPolicyGroupShrinkRequest {
	s.NetRedirectPolicyShrink = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupShrinkRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetPolicyGroupName(v string) *ModifyPolicyGroupShrinkRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetResolutionHeight(v int32) *ModifyPolicyGroupShrinkRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetResolutionWidth(v int32) *ModifyPolicyGroupShrinkRequest {
	s.ResolutionWidth = &v
	return s
}

type ModifyPolicyGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 552B7EED-D434-511F-B838-29EA4E906034
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponseBody) SetRequestId(v string) *ModifyPolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyGroupResponse) SetStatusCode(v int32) *ModifyPolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyGroupResponse) SetBody(v *ModifyPolicyGroupResponseBody) *ModifyPolicyGroupResponse {
	s.Body = v
	return s
}

type OperateAppRequest struct {
	// example:
	//
	// 1234
	AppId          *int32    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// start
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s OperateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateAppRequest) GoString() string {
	return s.String()
}

func (s *OperateAppRequest) SetAppId(v int32) *OperateAppRequest {
	s.AppId = &v
	return s
}

func (s *OperateAppRequest) SetInstanceIdList(v []*string) *OperateAppRequest {
	s.InstanceIdList = v
	return s
}

func (s *OperateAppRequest) SetOperateType(v string) *OperateAppRequest {
	s.OperateType = &v
	return s
}

type OperateAppResponseBody struct {
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-imr0fufqgac2z****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateAppResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppResponseBody) SetRequestId(v string) *OperateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppResponseBody) SetTaskId(v string) *OperateAppResponseBody {
	s.TaskId = &v
	return s
}

type OperateAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateAppResponse) GoString() string {
	return s.String()
}

func (s *OperateAppResponse) SetHeaders(v map[string]*string) *OperateAppResponse {
	s.Headers = v
	return s
}

func (s *OperateAppResponse) SetStatusCode(v int32) *OperateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppResponse) SetBody(v *OperateAppResponseBody) *OperateAppResponse {
	s.Body = v
	return s
}

type RebootAndroidInstancesInGroupRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	ForceStop          *bool     `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s RebootAndroidInstancesInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootAndroidInstancesInGroupRequest) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupRequest) SetAndroidInstanceIds(v []*string) *RebootAndroidInstancesInGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *RebootAndroidInstancesInGroupRequest) SetForceStop(v bool) *RebootAndroidInstancesInGroupRequest {
	s.ForceStop = &v
	return s
}

type RebootAndroidInstancesInGroupResponseBody struct {
	// example:
	//
	// 227CBB4C-F5DC-589D-A667-C5CA3D52****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootAndroidInstancesInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponseBody) SetRequestId(v string) *RebootAndroidInstancesInGroupResponseBody {
	s.RequestId = &v
	return s
}

type RebootAndroidInstancesInGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootAndroidInstancesInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootAndroidInstancesInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponse) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponse) SetHeaders(v map[string]*string) *RebootAndroidInstancesInGroupResponse {
	s.Headers = v
	return s
}

func (s *RebootAndroidInstancesInGroupResponse) SetStatusCode(v int32) *RebootAndroidInstancesInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponse) SetBody(v *RebootAndroidInstancesInGroupResponseBody) *RebootAndroidInstancesInGroupResponse {
	s.Body = v
	return s
}

type RecoveryFileRequest struct {
	// The IDs of the instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Whether all data is to be backed up.
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-azhps4rdyi2th****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The OSS path to which the backup file is uploaded.
	//
	// >  When calling the describeBuckets operation to retrieve a bucket name, you must also call the ossObjectList operation to obtain the object key. Combine these to form the full path: oss://${bucketName}/${key}.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The endpoint of the OSS bucket that stores the backup file.
	//
	// > : When calling the DescribeBuckets operation to query buckets, retrieve the IntranetEndpoint value if the cloud phone and the OSS bucket are in the same region. If they are in different regions, retrieve the ExtranetEndpoint value instead.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the backup.
	//
	// Valid values:
	//
	// 	- OSS: backup files are stored in OSS buckets.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s RecoveryFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoveryFileRequest) GoString() string {
	return s.String()
}

func (s *RecoveryFileRequest) SetAndroidInstanceIdList(v []*string) *RecoveryFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *RecoveryFileRequest) SetBackupAll(v bool) *RecoveryFileRequest {
	s.BackupAll = &v
	return s
}

func (s *RecoveryFileRequest) SetBackupFileId(v string) *RecoveryFileRequest {
	s.BackupFileId = &v
	return s
}

func (s *RecoveryFileRequest) SetBackupFilePath(v string) *RecoveryFileRequest {
	s.BackupFilePath = &v
	return s
}

func (s *RecoveryFileRequest) SetUploadEndpoint(v string) *RecoveryFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *RecoveryFileRequest) SetUploadType(v string) *RecoveryFileRequest {
	s.UploadType = &v
	return s
}

type RecoveryFileResponseBody struct {
	// The number of entries.
	//
	// example:
	//
	// 97
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The backup file that is restored.
	//
	// example:
	//
	// 6AD56E39-430B-5401-AB4A-7B086454****
	Data []*RecoveryFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6AD56E39-430B-5401-AB4A-7B086454****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-5prhfo7wv1gjx****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoveryFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoveryFileResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryFileResponseBody) SetCount(v int64) *RecoveryFileResponseBody {
	s.Count = &v
	return s
}

func (s *RecoveryFileResponseBody) SetData(v []*RecoveryFileResponseBodyData) *RecoveryFileResponseBody {
	s.Data = v
	return s
}

func (s *RecoveryFileResponseBody) SetRequestId(v string) *RecoveryFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoveryFileResponseBody) SetTaskId(v string) *RecoveryFileResponseBody {
	s.TaskId = &v
	return s
}

type RecoveryFileResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-5prhfo7wv1gjx****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoveryFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecoveryFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecoveryFileResponseBodyData) SetAndroidInstanceId(v string) *RecoveryFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *RecoveryFileResponseBodyData) SetTaskId(v string) *RecoveryFileResponseBodyData {
	s.TaskId = &v
	return s
}

type RecoveryFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoveryFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoveryFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoveryFileResponse) GoString() string {
	return s.String()
}

func (s *RecoveryFileResponse) SetHeaders(v map[string]*string) *RecoveryFileResponse {
	s.Headers = v
	return s
}

func (s *RecoveryFileResponse) SetStatusCode(v int32) *RecoveryFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryFileResponse) SetBody(v *RecoveryFileResponseBody) *RecoveryFileResponse {
	s.Body = v
	return s
}

type RenewAndroidInstanceGroupsRequest struct {
	// example:
	//
	// true
	AutoPay          *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s RenewAndroidInstanceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewAndroidInstanceGroupsRequest) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsRequest) SetAutoPay(v bool) *RenewAndroidInstanceGroupsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetInstanceGroupIds(v []*string) *RenewAndroidInstanceGroupsRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPeriod(v int32) *RenewAndroidInstanceGroupsRequest {
	s.Period = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPeriodUnit(v string) *RenewAndroidInstanceGroupsRequest {
	s.PeriodUnit = &v
	return s
}

type RenewAndroidInstanceGroupsResponseBody struct {
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 4B886792-2051-5DB4-8AE6-C8E45D3B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAndroidInstanceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewAndroidInstanceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsResponseBody) SetOrderId(v string) *RenewAndroidInstanceGroupsResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewAndroidInstanceGroupsResponseBody) SetRequestId(v string) *RenewAndroidInstanceGroupsResponseBody {
	s.RequestId = &v
	return s
}

type RenewAndroidInstanceGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAndroidInstanceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAndroidInstanceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewAndroidInstanceGroupsResponse) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsResponse) SetHeaders(v map[string]*string) *RenewAndroidInstanceGroupsResponse {
	s.Headers = v
	return s
}

func (s *RenewAndroidInstanceGroupsResponse) SetStatusCode(v int32) *RenewAndroidInstanceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAndroidInstanceGroupsResponse) SetBody(v *RenewAndroidInstanceGroupsResponseBody) *RenewAndroidInstanceGroupsResponse {
	s.Body = v
	return s
}

type ResetAndroidInstancesInGroupRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
}

func (s ResetAndroidInstancesInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAndroidInstancesInGroupRequest) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupRequest) SetAndroidInstanceIds(v []*string) *ResetAndroidInstancesInGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

type ResetAndroidInstancesInGroupResponseBody struct {
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAndroidInstancesInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponseBody) SetRequestId(v string) *ResetAndroidInstancesInGroupResponseBody {
	s.RequestId = &v
	return s
}

type ResetAndroidInstancesInGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAndroidInstancesInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAndroidInstancesInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponse) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponse) SetHeaders(v map[string]*string) *ResetAndroidInstancesInGroupResponse {
	s.Headers = v
	return s
}

func (s *ResetAndroidInstancesInGroupResponse) SetStatusCode(v int32) *ResetAndroidInstancesInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponse) SetBody(v *ResetAndroidInstancesInGroupResponseBody) *ResetAndroidInstancesInGroupResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	// example:
	//
	// ls
	CommandContent  *string   `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	ContentEncoding *string   `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetInstanceIds(v []*string) *RunCommandRequest {
	s.InstanceIds = v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

type RunCommandResponseBody struct {
	// example:
	//
	// t-gov2ujrk32v4****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// example:
	//
	// 440D7342-5E7C-B2DB-D0B4EAC2BDF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type SendFileRequest struct {
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// /data
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// This parameter is required.
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s SendFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) SetAndroidInstanceIdList(v []*string) *SendFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *SendFileRequest) SetSourceFilePath(v string) *SendFileRequest {
	s.SourceFilePath = &v
	return s
}

func (s *SendFileRequest) SetUploadEndpoint(v string) *SendFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *SendFileRequest) SetUploadType(v string) *SendFileRequest {
	s.UploadType = &v
	return s
}

func (s *SendFileRequest) SetUploadUrl(v string) *SendFileRequest {
	s.UploadUrl = &v
	return s
}

type SendFileResponseBody struct {
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	Data []*SendFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) SetData(v []*SendFileResponseBodyData) *SendFileResponseBody {
	s.Data = v
	return s
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendFileResponseBody) SetTaskId(v string) *SendFileResponseBody {
	s.TaskId = &v
	return s
}

type SendFileResponseBodyData struct {
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendFileResponseBodyData) SetAndroidInstanceId(v string) *SendFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *SendFileResponseBodyData) SetTaskId(v string) *SendFileResponseBodyData {
	s.TaskId = &v
	return s
}

type SendFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponse) GoString() string {
	return s.String()
}

func (s *SendFileResponse) SetHeaders(v map[string]*string) *SendFileResponse {
	s.Headers = v
	return s
}

func (s *SendFileResponse) SetStatusCode(v int32) *SendFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFileResponse) SetBody(v *SendFileResponseBody) *SendFileResponse {
	s.Body = v
	return s
}

type SetAdbSecureRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetAdbSecureRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAdbSecureRequest) GoString() string {
	return s.String()
}

func (s *SetAdbSecureRequest) SetInstanceIds(v []*string) *SetAdbSecureRequest {
	s.InstanceIds = v
	return s
}

func (s *SetAdbSecureRequest) SetStatus(v int32) *SetAdbSecureRequest {
	s.Status = &v
	return s
}

type SetAdbSecureResponseBody struct {
	Data *SetAdbSecureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAdbSecureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAdbSecureResponseBody) GoString() string {
	return s.String()
}

func (s *SetAdbSecureResponseBody) SetData(v *SetAdbSecureResponseBodyData) *SetAdbSecureResponseBody {
	s.Data = v
	return s
}

func (s *SetAdbSecureResponseBody) SetRequestId(v string) *SetAdbSecureResponseBody {
	s.RequestId = &v
	return s
}

type SetAdbSecureResponseBodyData struct {
	// example:
	//
	// 0
	FailCount   *int32    `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SetAdbSecureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetAdbSecureResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetAdbSecureResponseBodyData) SetFailCount(v int32) *SetAdbSecureResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *SetAdbSecureResponseBodyData) SetInstanceIds(v []*string) *SetAdbSecureResponseBodyData {
	s.InstanceIds = v
	return s
}

func (s *SetAdbSecureResponseBodyData) SetTotalCount(v int32) *SetAdbSecureResponseBodyData {
	s.TotalCount = &v
	return s
}

type SetAdbSecureResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAdbSecureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAdbSecureResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAdbSecureResponse) GoString() string {
	return s.String()
}

func (s *SetAdbSecureResponse) SetHeaders(v map[string]*string) *SetAdbSecureResponse {
	s.Headers = v
	return s
}

func (s *SetAdbSecureResponse) SetStatusCode(v int32) *SetAdbSecureResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAdbSecureResponse) SetBody(v *SetAdbSecureResponseBody) *SetAdbSecureResponse {
	s.Body = v
	return s
}

type StartAndroidInstanceRequest struct {
	// List of instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
}

func (s StartAndroidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartAndroidInstanceRequest) SetAndroidInstanceIds(v []*string) *StartAndroidInstanceRequest {
	s.AndroidInstanceIds = v
	return s
}

type StartAndroidInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 17C731AB-AAEE-5844-A352-D8D0352D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAndroidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartAndroidInstanceResponseBody) SetRequestId(v string) *StartAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartAndroidInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAndroidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartAndroidInstanceResponse) SetHeaders(v map[string]*string) *StartAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartAndroidInstanceResponse) SetStatusCode(v int32) *StartAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAndroidInstanceResponse) SetBody(v *StartAndroidInstanceResponseBody) *StartAndroidInstanceResponse {
	s.Body = v
	return s
}

type StopAndroidInstanceRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	ForceStop          *bool     `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s StopAndroidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopAndroidInstanceRequest) SetAndroidInstanceIds(v []*string) *StopAndroidInstanceRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *StopAndroidInstanceRequest) SetForceStop(v bool) *StopAndroidInstanceRequest {
	s.ForceStop = &v
	return s
}

type StopAndroidInstanceResponseBody struct {
	// example:
	//
	// E38B41A8-8E00-5AE4-A957-6636ACB8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAndroidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopAndroidInstanceResponseBody) SetRequestId(v string) *StopAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopAndroidInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAndroidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopAndroidInstanceResponse) SetHeaders(v map[string]*string) *StopAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopAndroidInstanceResponse) SetStatusCode(v int32) *StopAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAndroidInstanceResponse) SetBody(v *StopAndroidInstanceResponseBody) *StopAndroidInstanceResponse {
	s.Body = v
	return s
}

type UninstallAppRequest struct {
	AppIdList           []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	InstanceGroupIdList []*string `json:"InstanceGroupIdList,omitempty" xml:"InstanceGroupIdList,omitempty" type:"Repeated"`
	InstanceIdList      []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s UninstallAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallAppRequest) GoString() string {
	return s.String()
}

func (s *UninstallAppRequest) SetAppIdList(v []*string) *UninstallAppRequest {
	s.AppIdList = v
	return s
}

func (s *UninstallAppRequest) SetInstanceGroupIdList(v []*string) *UninstallAppRequest {
	s.InstanceGroupIdList = v
	return s
}

func (s *UninstallAppRequest) SetInstanceIdList(v []*string) *UninstallAppRequest {
	s.InstanceIdList = v
	return s
}

type UninstallAppResponseBody struct {
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAppResponseBody) SetRequestId(v string) *UninstallAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAppResponseBody) SetTaskId(v string) *UninstallAppResponseBody {
	s.TaskId = &v
	return s
}

type UninstallAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallAppResponse) GoString() string {
	return s.String()
}

func (s *UninstallAppResponse) SetHeaders(v map[string]*string) *UninstallAppResponse {
	s.Headers = v
	return s
}

func (s *UninstallAppResponse) SetStatusCode(v int32) *UninstallAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAppResponse) SetBody(v *UninstallAppResponseBody) *UninstallAppResponse {
	s.Body = v
	return s
}

type UpdateCustomImageNameRequest struct {
	// The ID of the image.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// imagename
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s UpdateCustomImageNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomImageNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomImageNameRequest) SetImageId(v string) *UpdateCustomImageNameRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateCustomImageNameRequest) SetImageName(v string) *UpdateCustomImageNameRequest {
	s.ImageName = &v
	return s
}

type UpdateCustomImageNameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 552B7EED-D434-511F-B838-29EA4E906034
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomImageNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomImageNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomImageNameResponseBody) SetRequestId(v string) *UpdateCustomImageNameResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomImageNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomImageNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomImageNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomImageNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomImageNameResponse) SetHeaders(v map[string]*string) *UpdateCustomImageNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomImageNameResponse) SetStatusCode(v int32) *UpdateCustomImageNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomImageNameResponse) SetBody(v *UpdateCustomImageNameResponseBody) *UpdateCustomImageNameResponse {
	s.Body = v
	return s
}

type UpdateInstanceGroupImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// imgc-075cllfeuazh****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateInstanceGroupImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceGroupImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceGroupImageRequest) SetImageId(v string) *UpdateInstanceGroupImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceGroupImageRequest) SetInstanceGroupIds(v []*string) *UpdateInstanceGroupImageRequest {
	s.InstanceGroupIds = v
	return s
}

type UpdateInstanceGroupImageResponseBody struct {
	// example:
	//
	// 55726272-E40B-530D-914F-5126B19C79B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceGroupImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceGroupImageResponseBody) SetRequestId(v string) *UpdateInstanceGroupImageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceGroupImageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceGroupImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceGroupImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceGroupImageResponse) SetHeaders(v map[string]*string) *UpdateInstanceGroupImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceGroupImageResponse) SetStatusCode(v int32) *UpdateInstanceGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceGroupImageResponse) SetBody(v *UpdateInstanceGroupImageResponseBody) *UpdateInstanceGroupImageResponse {
	s.Body = v
	return s
}

type UpgradeAndroidInstanceGroupRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Make sure that your Alibaba Cloud account has sufficient balance.
	//
	// 	- false: disables the auto-payment feature. You need to manually complete the payment process.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The number of instances that you want to increase.
	//
	// example:
	//
	// 10
	IncreaseNumberOfInstance *int32 `json:"IncreaseNumberOfInstance,omitempty" xml:"IncreaseNumberOfInstance,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-asguicdjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
}

func (s UpgradeAndroidInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAndroidInstanceGroupRequest) SetAutoPay(v bool) *UpgradeAndroidInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupRequest) SetIncreaseNumberOfInstance(v int32) *UpgradeAndroidInstanceGroupRequest {
	s.IncreaseNumberOfInstance = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupRequest) SetInstanceGroupId(v string) *UpgradeAndroidInstanceGroupRequest {
	s.InstanceGroupId = &v
	return s
}

type UpgradeAndroidInstanceGroupResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// [\\"acp-3vzqq4y3f31f3z3df\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 223684716098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 310A783E-CC46-5452-A8A3-71AE5DB59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeAndroidInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAndroidInstanceGroupResponseBody) SetInstanceIds(v string) *UpgradeAndroidInstanceGroupResponseBody {
	s.InstanceIds = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponseBody) SetOrderId(v string) *UpgradeAndroidInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponseBody) SetRequestId(v string) *UpgradeAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAndroidInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *UpgradeAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponse) SetStatusCode(v int32) *UpgradeAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupResponse) SetBody(v *UpgradeAndroidInstanceGroupResponseBody) *UpgradeAndroidInstanceGroupResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eds-aic"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 绑定密钥对
//
// @param request - AttachKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachKeyPairResponse
func (client *Client) AttachKeyPairWithOptions(request *AttachKeyPairRequest, runtime *util.RuntimeOptions) (_result *AttachKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairId)) {
		query["KeyPairId"] = request.KeyPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachKeyPair"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AttachKeyPairResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AttachKeyPairResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 绑定密钥对
//
// @param request - AttachKeyPairRequest
//
// @return AttachKeyPairResponse
func (client *Client) AttachKeyPair(request *AttachKeyPairRequest) (_result *AttachKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachKeyPairResponse{}
	_body, _err := client.AttachKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorize/unauthorize Android instances for users.
//
// Description:
//
// Instance states that support user assignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed.
//
// Instance states that support unassignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed, Expired, Overdue, Deleted.
//
// @param request - AuthorizeAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeAndroidInstanceResponse
func (client *Client) AuthorizeAndroidInstanceWithOptions(request *AuthorizeAndroidInstanceRequest, runtime *util.RuntimeOptions) (_result *AuthorizeAndroidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeUserId)) {
		query["AuthorizeUserId"] = request.AuthorizeUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UnAuthorizeUserId)) {
		query["UnAuthorizeUserId"] = request.UnAuthorizeUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeAndroidInstance"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AuthorizeAndroidInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AuthorizeAndroidInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Authorize/unauthorize Android instances for users.
//
// Description:
//
// Instance states that support user assignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed.
//
// Instance states that support unassignment: Available, Shutting Down, Stopped, Starting, Backing Up, Restoring, Backup Failed, Restore Failed, Expired, Overdue, Deleted.
//
// @param request - AuthorizeAndroidInstanceRequest
//
// @return AuthorizeAndroidInstanceResponse
func (client *Client) AuthorizeAndroidInstance(request *AuthorizeAndroidInstanceRequest) (_result *AuthorizeAndroidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeAndroidInstanceResponse{}
	_body, _err := client.AuthorizeAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates and uploads backup files.
//
// Description:
//
// Currently, this operation allows you to upload only backup files generated by cloud phones to Object Storage Service (OSS) buckets.
//
// @param request - BackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackupFileResponse
func (client *Client) BackupFileWithOptions(request *BackupFileRequest, runtime *util.RuntimeOptions) (_result *BackupFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIdList)) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.BackupAll)) {
		query["BackupAll"] = request.BackupAll
	}

	if !tea.BoolValue(util.IsUnset(request.BackupFileName)) {
		query["BackupFileName"] = request.BackupFileName
	}

	if !tea.BoolValue(util.IsUnset(request.BackupFilePath)) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SourceAppList)) {
		query["SourceAppList"] = request.SourceAppList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFilePathList)) {
		query["SourceFilePathList"] = request.SourceFilePathList
	}

	if !tea.BoolValue(util.IsUnset(request.UploadEndpoint)) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["UploadType"] = request.UploadType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BackupFile"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BackupFileResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BackupFileResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Generates and uploads backup files.
//
// Description:
//
// Currently, this operation allows you to upload only backup files generated by cloud phones to Object Storage Service (OSS) buckets.
//
// @param request - BackupFileRequest
//
// @return BackupFileResponse
func (client *Client) BackupFile(request *BackupFileRequest) (_result *BackupFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BackupFileResponse{}
	_body, _err := client.BackupFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取ticket
//
// @param request - BatchGetAcpConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetAcpConnectionTicketResponse
func (client *Client) BatchGetAcpConnectionTicketWithOptions(request *BatchGetAcpConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *BatchGetAcpConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupId)) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTasks)) {
		query["InstanceTasks"] = request.InstanceTasks
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetAcpConnectionTicket"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BatchGetAcpConnectionTicketResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BatchGetAcpConnectionTicketResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 批量获取ticket
//
// @param request - BatchGetAcpConnectionTicketRequest
//
// @return BatchGetAcpConnectionTicketResponse
func (client *Client) BatchGetAcpConnectionTicket(request *BatchGetAcpConnectionTicketRequest) (_result *BatchGetAcpConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchGetAcpConnectionTicketResponse{}
	_body, _err := client.BatchGetAcpConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check the resource inventory.
//
// @param request - CheckResourceStockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourceStockResponse
func (client *Client) CheckResourceStockWithOptions(request *CheckResourceStockRequest, runtime *util.RuntimeOptions) (_result *CheckResourceStockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcpSpecId)) {
		query["AcpSpecId"] = request.AcpSpecId
	}

	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.GpuAcceleration)) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckResourceStock"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckResourceStockResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckResourceStockResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Check the resource inventory.
//
// @param request - CheckResourceStockRequest
//
// @return CheckResourceStockResponse
func (client *Client) CheckResourceStock(request *CheckResourceStockRequest) (_result *CheckResourceStockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResourceStockResponse{}
	_body, _err := client.CheckResourceStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates pay-as-you-go or subscription instance groups.
//
// Description:
//
// Before creating an instance group, ensure you understand the [billing methods](https://help.aliyun.com/document_detail/2807121.html) supported by Cloud Phone.
//
// 	- If the billing method of an instance group is PrePaid, AutoPay is set to false by default. In this case, you need to go to [Expenses and Costs](https://usercenter2-intl.aliyun.com/order/list) to manually complete the payment.
//
// 	- You can also set AutoPay to true based on your business requirements.
//
// @param request - CreateAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndroidInstanceGroupResponse
func (client *Client) CreateAndroidInstanceGroupWithOptions(request *CreateAndroidInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAndroidInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIpv6)) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !tea.BoolValue(util.IsUnset(request.GpuAcceleration)) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupName)) {
		query["InstanceGroupName"] = request.InstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupSpec)) {
		query["InstanceGroupSpec"] = request.InstanceGroupSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6Bandwidth)) {
		query["Ipv6Bandwidth"] = request.Ipv6Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairId)) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !tea.BoolValue(util.IsUnset(request.NumberOfInstances)) {
		query["NumberOfInstances"] = request.NumberOfInstances
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAndroidInstanceGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAndroidInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAndroidInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates pay-as-you-go or subscription instance groups.
//
// Description:
//
// Before creating an instance group, ensure you understand the [billing methods](https://help.aliyun.com/document_detail/2807121.html) supported by Cloud Phone.
//
// 	- If the billing method of an instance group is PrePaid, AutoPay is set to false by default. In this case, you need to go to [Expenses and Costs](https://usercenter2-intl.aliyun.com/order/list) to manually complete the payment.
//
// 	- You can also set AutoPay to true based on your business requirements.
//
// @param request - CreateAndroidInstanceGroupRequest
//
// @return CreateAndroidInstanceGroupResponse
func (client *Client) CreateAndroidInstanceGroup(request *CreateAndroidInstanceGroupRequest) (_result *CreateAndroidInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAndroidInstanceGroupResponse{}
	_body, _err := client.CreateAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Android application.
//
// Description:
//
// ### [](#)Preparations
//
// Before you proceed, log on to the [Elastic Desktop Service (EDS) Enterprise console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the values of request parameters `FileName`, `FilePath`, and `OssAppUrl`.
//
// @param tmpReq - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(tmpReq *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomAppInfo)) {
		request.CustomAppInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAppInfo, tea.String("CustomAppInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomAppInfoShrink)) {
		query["CustomAppInfo"] = request.CustomAppInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.IconUrl)) {
		query["IconUrl"] = request.IconUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstallParam)) {
		query["InstallParam"] = request.InstallParam
	}

	if !tea.BoolValue(util.IsUnset(request.OssAppUrl)) {
		query["OssAppUrl"] = request.OssAppUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an Android application.
//
// Description:
//
// ### [](#)Preparations
//
// Before you proceed, log on to the [Elastic Desktop Service (EDS) Enterprise console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the application file to Application Center to obtain the values of request parameters `FileName`, `FilePath`, and `OssAppUrl`.
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create Custom Image
//
// @param request - CreateCustomImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomImageResponse
func (client *Client) CreateCustomImageWithOptions(request *CreateCustomImageRequest, runtime *util.RuntimeOptions) (_result *CreateCustomImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		body["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomImage"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateCustomImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateCustomImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Create Custom Image
//
// @param request - CreateCustomImageRequest
//
// @return CreateCustomImageResponse
func (client *Client) CreateCustomImage(request *CreateCustomImageRequest) (_result *CreateCustomImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CreateCustomImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Android Debug Bridge (ADB) key pair. The system retains the public key and provides a PEM-encoded private key in PKCS#8 format, adhering to the ADB connection specification. You must securely store the private key.
//
// Description:
//
// In addition to using the CreateKeyPair operation to generate a key pair, you can also create one by using the ADB tool and upload it to the Cloud Phone console. The usage of this key pair is identical to that of a system-generated key pair.
//
// Each tenant can create up to 500 key pairs.
//
// @param request - CreateKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPairWithOptions(request *CreateKeyPairRequest, runtime *util.RuntimeOptions) (_result *CreateKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKeyPair"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateKeyPairResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateKeyPairResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an Android Debug Bridge (ADB) key pair. The system retains the public key and provides a PEM-encoded private key in PKCS#8 format, adhering to the ADB connection specification. You must securely store the private key.
//
// Description:
//
// In addition to using the CreateKeyPair operation to generate a key pair, you can also create one by using the ADB tool and upload it to the Cloud Phone console. The usage of this key pair is identical to that of a system-generated key pair.
//
// Each tenant can create up to 500 key pairs.
//
// @param request - CreateKeyPairRequest
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPair(request *CreateKeyPairRequest) (_result *CreateKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CreateKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a policy.
//
// @param tmpReq - CreatePolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyGroupResponse
func (client *Client) CreatePolicyGroupWithOptions(tmpReq *CreatePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePolicyGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NetRedirectPolicy)) {
		request.NetRedirectPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetRedirectPolicy, tea.String("NetRedirectPolicy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CameraRedirect)) {
		body["CameraRedirect"] = request.CameraRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		body["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.Html5FileTransfer)) {
		body["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.LocalDrive)) {
		body["LocalDrive"] = request.LocalDrive
	}

	if !tea.BoolValue(util.IsUnset(request.LockResolution)) {
		body["LockResolution"] = request.LockResolution
	}

	if !tea.BoolValue(util.IsUnset(request.NetRedirectPolicyShrink)) {
		body["NetRedirectPolicy"] = request.NetRedirectPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupName)) {
		body["PolicyGroupName"] = request.PolicyGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionHeight)) {
		body["ResolutionHeight"] = request.ResolutionHeight
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionWidth)) {
		body["ResolutionWidth"] = request.ResolutionWidth
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicyGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePolicyGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePolicyGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a policy.
//
// @param request - CreatePolicyGroupRequest
//
// @return CreatePolicyGroupResponse
func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建云手机截图接口
//
// @param request - CreateScreenshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScreenshotResponse
func (client *Client) CreateScreenshotWithOptions(request *CreateScreenshotRequest, runtime *util.RuntimeOptions) (_result *CreateScreenshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIdList)) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.SkipCheckPolicyConfig)) {
		query["SkipCheckPolicyConfig"] = request.SkipCheckPolicyConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScreenshot"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateScreenshotResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateScreenshotResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建云手机截图接口
//
// @param request - CreateScreenshotRequest
//
// @return CreateScreenshotResponse
func (client *Client) CreateScreenshot(request *CreateScreenshotRequest) (_result *CreateScreenshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScreenshotResponse{}
	_body, _err := client.CreateScreenshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
<<<<<<< Updated upstream
<<<<<<< Updated upstream
// 删除安卓实例组
=======
=======
>>>>>>> Stashed changes
// Delete an instance group.
//
// Description:
//
// You can delete only pay-as-you-go instance groups.
//
// You can delete subscription instance groups only after they expire.
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
//
// @param request - DeleteAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAndroidInstanceGroupResponse
func (client *Client) DeleteAndroidInstanceGroupWithOptions(request *DeleteAndroidInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAndroidInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIds)) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAndroidInstanceGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAndroidInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAndroidInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Delete an instance group.
//
// Description:
//
// You can delete only pay-as-you-go instance groups.
//
// You can delete subscription instance groups only after they expire.
//
// @param request - DeleteAndroidInstanceGroupRequest
//
// @return DeleteAndroidInstanceGroupResponse
func (client *Client) DeleteAndroidInstanceGroup(request *DeleteAndroidInstanceGroupRequest) (_result *DeleteAndroidInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAndroidInstanceGroupResponse{}
	_body, _err := client.DeleteAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application. Before you delete an application, make sure that the application is not installed on any instances.
//
// @param request - DeleteAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppsResponse
func (client *Client) DeleteAppsWithOptions(request *DeleteAppsRequest, runtime *util.RuntimeOptions) (_result *DeleteAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIdList)) {
		query["AppIdList"] = request.AppIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApps"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAppsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAppsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an application. Before you delete an application, make sure that the application is not installed on any instances.
//
// @param request - DeleteAppsRequest
//
// @return DeleteAppsResponse
func (client *Client) DeleteApps(request *DeleteAppsRequest) (_result *DeleteAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppsResponse{}
	_body, _err := client.DeleteAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom image.
//
// Description:
//
// You cannot delete an image that is currently in use by an instance group.
//
// @param tmpReq - DeleteImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImagesResponse
func (client *Client) DeleteImagesWithOptions(tmpReq *DeleteImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImageIds)) {
		request.ImageIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageIds, tea.String("ImageIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageIdsShrink)) {
		body["ImageIds"] = request.ImageIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImages"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteImagesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteImagesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a custom image.
//
// Description:
//
// You cannot delete an image that is currently in use by an instance group.
//
// @param request - DeleteImagesRequest
//
// @return DeleteImagesResponse
func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes Android Debug Bridge (ADB) key pairs.
//
// Description:
//
//   If a cloud phone instance is currently associated with the ADB key pair you intend to delete, the ADB key pair cannot be deleted.
//
// 	- Once an ADB key pair is deleted, it cannot be retrieved or queried by using the DescribeKeyPairs operation.
//
// @param request - DeleteKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairsWithOptions(request *DeleteKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairIds)) {
		query["KeyPairIds"] = request.KeyPairIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeyPairs"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteKeyPairsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteKeyPairsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes Android Debug Bridge (ADB) key pairs.
//
// Description:
//
//   If a cloud phone instance is currently associated with the ADB key pair you intend to delete, the ADB key pair cannot be deleted.
//
// 	- Once an ADB key pair is deleted, it cannot be retrieved or queried by using the DescribeKeyPairs operation.
//
// @param request - DeleteKeyPairsRequest
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (_result *DeleteKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DeleteKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @param request - DeletePolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyGroupResponse
func (client *Client) DeletePolicyGroupWithOptions(request *DeletePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyGroupIds)) {
		query["PolicyGroupIds"] = request.PolicyGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicyGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeletePolicyGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeletePolicyGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a policy.
//
// @param request - DeletePolicyGroupRequest
//
// @return DeletePolicyGroupResponse
func (client *Client) DeletePolicyGroup(request *DeletePolicyGroupRequest) (_result *DeletePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyGroupResponse{}
	_body, _err := client.DeletePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance group.
//
// @param request - DescribeAndroidInstanceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAndroidInstanceGroupsResponse
func (client *Client) DescribeAndroidInstanceGroupsWithOptions(request *DescribeAndroidInstanceGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeAndroidInstanceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIds)) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupName)) {
		query["InstanceGroupName"] = request.InstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairId)) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAndroidInstanceGroups"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAndroidInstanceGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAndroidInstanceGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of an instance group.
//
// @param request - DescribeAndroidInstanceGroupsRequest
//
// @return DescribeAndroidInstanceGroupsResponse
func (client *Client) DescribeAndroidInstanceGroups(request *DescribeAndroidInstanceGroupsRequest) (_result *DescribeAndroidInstanceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAndroidInstanceGroupsResponse{}
	_body, _err := client.DescribeAndroidInstanceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
<<<<<<< Updated upstream
<<<<<<< Updated upstream
// 查询安卓的实例列表
=======
// Queries cloud phone instances.
>>>>>>> Stashed changes
=======
// Queries cloud phone instances.
>>>>>>> Stashed changes
//
// @param request - DescribeAndroidInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAndroidInstancesResponse
func (client *Client) DescribeAndroidInstancesWithOptions(request *DescribeAndroidInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeAndroidInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceName)) {
		query["AndroidInstanceName"] = request.AndroidInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupId)) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIds)) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupName)) {
		query["InstanceGroupName"] = request.InstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairId)) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAndroidInstances"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAndroidInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAndroidInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
<<<<<<< Updated upstream
<<<<<<< Updated upstream
// 查询安卓的实例列表
=======
// Queries cloud phone instances.
>>>>>>> Stashed changes
=======
// Queries cloud phone instances.
>>>>>>> Stashed changes
//
// @param request - DescribeAndroidInstancesRequest
//
// @return DescribeAndroidInstancesResponse
func (client *Client) DescribeAndroidInstances(request *DescribeAndroidInstancesRequest) (_result *DescribeAndroidInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAndroidInstancesResponse{}
	_body, _err := client.DescribeAndroidInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries applications.
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIdList)) {
		query["AppIdList"] = request.AppIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstallationStatus)) {
		query["InstallationStatus"] = request.InstallationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MD5)) {
		query["MD5"] = request.MD5
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApps"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAppsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAppsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries applications.
//
// @param request - DescribeAppsRequest
//
// @return DescribeAppsResponse
func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backup files.
//
// Description:
//
// Currently, this operation allows you to query only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - DescribeBackupFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupFilesResponse
func (client *Client) DescribeBackupFilesWithOptions(request *DescribeBackupFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceId)) {
		query["AndroidInstanceId"] = request.AndroidInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceName)) {
		query["AndroidInstanceName"] = request.AndroidInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.BackupAll)) {
		query["BackupAll"] = request.BackupAll
	}

	if !tea.BoolValue(util.IsUnset(request.BackupFileId)) {
		query["BackupFileId"] = request.BackupFileId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupFileName)) {
		query["BackupFileName"] = request.BackupFileName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupId)) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		query["StatusList"] = request.StatusList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupFiles"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBackupFilesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBackupFilesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries backup files.
//
// Description:
//
// Currently, this operation allows you to query only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - DescribeBackupFilesRequest
//
// @return DescribeBackupFilesResponse
func (client *Client) DescribeBackupFiles(request *DescribeBackupFilesRequest) (_result *DescribeBackupFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupFilesResponse{}
	_body, _err := client.DescribeBackupFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
<<<<<<< Updated upstream
<<<<<<< Updated upstream
// 查询镜像列表
=======
// Queries images.
>>>>>>> Stashed changes
=======
// Queries images.
>>>>>>> Stashed changes
//
// @param request - DescribeImageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageListResponse
func (client *Client) DescribeImageListWithOptions(request *DescribeImageListRequest, runtime *util.RuntimeOptions) (_result *DescribeImageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImagePackageType)) {
		query["ImagePackageType"] = request.ImagePackageType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		body["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		body["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageList"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeImageListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeImageListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
<<<<<<< Updated upstream
<<<<<<< Updated upstream
// 查询镜像列表
=======
// Queries images.
>>>>>>> Stashed changes
=======
// Queries images.
>>>>>>> Stashed changes
//
// @param request - DescribeImageListRequest
//
// @return DescribeImageListResponse
func (client *Client) DescribeImageList(request *DescribeImageListRequest) (_result *DescribeImageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageListResponse{}
	_body, _err := client.DescribeImageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询命令结果
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InvocationId)) {
		query["InvocationId"] = request.InvocationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocations"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInvocationsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInvocationsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询命令结果
//
// @param request - DescribeInvocationsRequest
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more key pairs.
//
// @param request - DescribeKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyPairsResponse
func (client *Client) DescribeKeyPairsWithOptions(request *DescribeKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairIds)) {
		query["KeyPairIds"] = request.KeyPairIds
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKeyPairs"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeKeyPairsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeKeyPairsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries one or more key pairs.
//
// @param request - DescribeKeyPairsRequest
//
// @return DescribeKeyPairsResponse
func (client *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (_result *DescribeKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.DescribeKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query available regions.
//
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
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Query available regions.
//
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
// Query available specifications.
//
// @param request - DescribeSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSpecResponse
func (client *Client) DescribeSpecWithOptions(request *DescribeSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.MatrixSpec)) {
		query["MatrixSpec"] = request.MatrixSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
	}

	if !tea.BoolValue(util.IsUnset(request.SpecIds)) {
		query["SpecIds"] = request.SpecIds
	}

	if !tea.BoolValue(util.IsUnset(request.SpecStatus)) {
		query["SpecStatus"] = request.SpecStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["SpecType"] = request.SpecType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSpec"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSpecResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSpecResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Query available specifications.
//
// @param request - DescribeSpecRequest
//
// @return DescribeSpecResponse
func (client *Client) DescribeSpec(request *DescribeSpecRequest) (_result *DescribeSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpecResponse{}
	_body, _err := client.DescribeSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询异步任务
//
// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		query["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		query["Param"] = request.Param
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		query["ParentTaskId"] = request.ParentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		query["TaskIds"] = request.TaskIds
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatuses)) {
		query["TaskStatuses"] = request.TaskStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTypes)) {
		query["TaskTypes"] = request.TaskTypes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTasks"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询异步任务
//
// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
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

// Summary:
//
// 解绑密钥对
//
// @param request - DetachKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachKeyPairResponse
func (client *Client) DetachKeyPairWithOptions(request *DetachKeyPairRequest, runtime *util.RuntimeOptions) (_result *DetachKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairId)) {
		query["KeyPairId"] = request.KeyPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachKeyPair"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DetachKeyPairResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DetachKeyPairResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 解绑密钥对
//
// @param request - DetachKeyPairRequest
//
// @return DetachKeyPairResponse
func (client *Client) DetachKeyPair(request *DetachKeyPairRequest) (_result *DetachKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachKeyPairResponse{}
	_body, _err := client.DetachKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义镜像分发
//
// @param request - DistributeImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DistributeImageResponse
func (client *Client) DistributeImageWithOptions(request *DistributeImageRequest, runtime *util.RuntimeOptions) (_result *DistributeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributeRegionList)) {
		body["DistributeRegionList"] = request.DistributeRegionList
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DistributeImage"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DistributeImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DistributeImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 自定义镜像分发
//
// @param request - DistributeImageRequest
//
// @return DistributeImageResponse
func (client *Client) DistributeImage(request *DistributeImageRequest) (_result *DistributeImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DistributeImageResponse{}
	_body, _err := client.DistributeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例组缩容
//
// @param request - DowngradeAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DowngradeAndroidInstanceGroupResponse
func (client *Client) DowngradeAndroidInstanceGroupWithOptions(request *DowngradeAndroidInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *DowngradeAndroidInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupId)) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DowngradeAndroidInstanceGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DowngradeAndroidInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DowngradeAndroidInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 实例组缩容
//
// @param request - DowngradeAndroidInstanceGroupRequest
//
// @return DowngradeAndroidInstanceGroupResponse
func (client *Client) DowngradeAndroidInstanceGroup(request *DowngradeAndroidInstanceGroupRequest) (_result *DowngradeAndroidInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DowngradeAndroidInstanceGroupResponse{}
	_body, _err := client.DowngradeAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云手机拉取文件到OSS
//
// @param request - FetchFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchFileResponse
func (client *Client) FetchFileWithOptions(request *FetchFileRequest, runtime *util.RuntimeOptions) (_result *FetchFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIdList)) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFilePath)) {
		query["SourceFilePath"] = request.SourceFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.UploadEndpoint)) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["UploadType"] = request.UploadType
	}

	if !tea.BoolValue(util.IsUnset(request.UploadUrl)) {
		query["UploadUrl"] = request.UploadUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchFile"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &FetchFileResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &FetchFileResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 云手机拉取文件到OSS
//
// @param request - FetchFileRequest
//
// @return FetchFileResponse
func (client *Client) FetchFile(request *FetchFileRequest) (_result *FetchFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchFileResponse{}
	_body, _err := client.FetchFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入秘钥
//
// @param request - ImportKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPairWithOptions(request *ImportKeyPairRequest, runtime *util.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyBody)) {
		query["PublicKeyBody"] = request.PublicKeyBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportKeyPair"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ImportKeyPairResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ImportKeyPairResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 导入秘钥
//
// @param request - ImportKeyPairRequest
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (_result *ImportKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.ImportKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装app到实例组
//
// @param request - InstallAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAppResponse
func (client *Client) InstallAppWithOptions(request *InstallAppRequest, runtime *util.RuntimeOptions) (_result *InstallAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIdList)) {
		query["AppIdList"] = request.AppIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIdList)) {
		query["InstanceGroupIdList"] = request.InstanceGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallApp"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InstallAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InstallAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 安装app到实例组
//
// @param request - InstallAppRequest
//
// @return InstallAppResponse
func (client *Client) InstallApp(request *InstallAppRequest) (_result *InstallAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallAppResponse{}
	_body, _err := client.InstallAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries policies.
//
// @param request - ListPolicyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyGroupsResponse
func (client *Client) ListPolicyGroupsWithOptions(request *ListPolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *ListPolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupIds)) {
		body["PolicyGroupIds"] = request.PolicyGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupName)) {
		body["PolicyGroupName"] = request.PolicyGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicyGroups"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListPolicyGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListPolicyGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries policies.
//
// @param request - ListPolicyGroupsRequest
//
// @return ListPolicyGroupsResponse
func (client *Client) ListPolicyGroups(request *ListPolicyGroupsRequest) (_result *ListPolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicyGroupsResponse{}
	_body, _err := client.ListPolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies attributes of a cloud phone instance. Currently, this operation allows you to modify only the name of a cloud phone instance.
//
// @param request - ModifyAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAndroidInstanceResponse
func (client *Client) ModifyAndroidInstanceWithOptions(request *ModifyAndroidInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyAndroidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceId)) {
		query["AndroidInstanceId"] = request.AndroidInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NewAndroidInstanceName)) {
		query["NewAndroidInstanceName"] = request.NewAndroidInstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAndroidInstance"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAndroidInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAndroidInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies attributes of a cloud phone instance. Currently, this operation allows you to modify only the name of a cloud phone instance.
//
// @param request - ModifyAndroidInstanceRequest
//
// @return ModifyAndroidInstanceResponse
func (client *Client) ModifyAndroidInstance(request *ModifyAndroidInstanceRequest) (_result *ModifyAndroidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAndroidInstanceResponse{}
	_body, _err := client.ModifyAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies attributes of an instance group.
//
// @param request - ModifyAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAndroidInstanceGroupResponse
func (client *Client) ModifyAndroidInstanceGroupWithOptions(request *ModifyAndroidInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyAndroidInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceGroupId)) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewInstanceGroupName)) {
		query["NewInstanceGroupName"] = request.NewInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAndroidInstanceGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAndroidInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAndroidInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies attributes of an instance group.
//
// @param request - ModifyAndroidInstanceGroupRequest
//
// @return ModifyAndroidInstanceGroupResponse
func (client *Client) ModifyAndroidInstanceGroup(request *ModifyAndroidInstanceGroupRequest) (_result *ModifyAndroidInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAndroidInstanceGroupResponse{}
	_body, _err := client.ModifyAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify attributes of an application.
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IconUrl)) {
		query["IconUrl"] = request.IconUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApp"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modify attributes of an application.
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies Android Debug Bridge (ADB) key pairs.
//
// @param request - ModifyKeyPairNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyKeyPairNameResponse
func (client *Client) ModifyKeyPairNameWithOptions(request *ModifyKeyPairNameRequest, runtime *util.RuntimeOptions) (_result *ModifyKeyPairNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairId)) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !tea.BoolValue(util.IsUnset(request.NewKeyPairName)) {
		query["NewKeyPairName"] = request.NewKeyPairName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyKeyPairName"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyKeyPairNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyKeyPairNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies Android Debug Bridge (ADB) key pairs.
//
// @param request - ModifyKeyPairNameRequest
//
// @return ModifyKeyPairNameResponse
func (client *Client) ModifyKeyPairName(request *ModifyKeyPairNameRequest) (_result *ModifyKeyPairNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyKeyPairNameResponse{}
	_body, _err := client.ModifyKeyPairNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a policy.
//
// @param tmpReq - ModifyPolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyGroupResponse
func (client *Client) ModifyPolicyGroupWithOptions(tmpReq *ModifyPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyPolicyGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NetRedirectPolicy)) {
		request.NetRedirectPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetRedirectPolicy, tea.String("NetRedirectPolicy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CameraRedirect)) {
		body["CameraRedirect"] = request.CameraRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		body["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.Html5FileTransfer)) {
		body["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.LocalDrive)) {
		body["LocalDrive"] = request.LocalDrive
	}

	if !tea.BoolValue(util.IsUnset(request.LockResolution)) {
		body["LockResolution"] = request.LockResolution
	}

	if !tea.BoolValue(util.IsUnset(request.NetRedirectPolicyShrink)) {
		body["NetRedirectPolicy"] = request.NetRedirectPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupId)) {
		body["PolicyGroupId"] = request.PolicyGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroupName)) {
		body["PolicyGroupName"] = request.PolicyGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionHeight)) {
		body["ResolutionHeight"] = request.ResolutionHeight
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionWidth)) {
		body["ResolutionWidth"] = request.ResolutionWidth
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicyGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyPolicyGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyPolicyGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies a policy.
//
// @param request - ModifyPolicyGroupRequest
//
// @return ModifyPolicyGroupResponse
func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作App
//
// @param request - OperateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppResponse
func (client *Client) OperateAppWithOptions(request *OperateAppRequest, runtime *util.RuntimeOptions) (_result *OperateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateApp"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OperateAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OperateAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 操作App
//
// @param request - OperateAppRequest
//
// @return OperateAppResponse
func (client *Client) OperateApp(request *OperateAppRequest) (_result *OperateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateAppResponse{}
	_body, _err := client.OperateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启安卓实例
//
// @param request - RebootAndroidInstancesInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootAndroidInstancesInGroupResponse
func (client *Client) RebootAndroidInstancesInGroupWithOptions(request *RebootAndroidInstancesInGroupRequest, runtime *util.RuntimeOptions) (_result *RebootAndroidInstancesInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ForceStop)) {
		query["ForceStop"] = request.ForceStop
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootAndroidInstancesInGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RebootAndroidInstancesInGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RebootAndroidInstancesInGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 重启安卓实例
//
// @param request - RebootAndroidInstancesInGroupRequest
//
// @return RebootAndroidInstancesInGroupResponse
func (client *Client) RebootAndroidInstancesInGroup(request *RebootAndroidInstancesInGroupRequest) (_result *RebootAndroidInstancesInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootAndroidInstancesInGroupResponse{}
	_body, _err := client.RebootAndroidInstancesInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores backup files.
//
// Description:
//
// Currently, this operation allows you to restore only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - RecoveryFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoveryFileResponse
func (client *Client) RecoveryFileWithOptions(request *RecoveryFileRequest, runtime *util.RuntimeOptions) (_result *RecoveryFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIdList)) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.BackupAll)) {
		query["BackupAll"] = request.BackupAll
	}

	if !tea.BoolValue(util.IsUnset(request.BackupFileId)) {
		query["BackupFileId"] = request.BackupFileId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupFilePath)) {
		query["BackupFilePath"] = request.BackupFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.UploadEndpoint)) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["UploadType"] = request.UploadType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoveryFile"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RecoveryFileResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RecoveryFileResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Restores backup files.
//
// Description:
//
// Currently, this operation allows you to restore only backup files generated by cloud phones that are stored in Object Storage Service (OSS) buckets.
//
// @param request - RecoveryFileRequest
//
// @return RecoveryFileResponse
func (client *Client) RecoveryFile(request *RecoveryFileRequest) (_result *RecoveryFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoveryFileResponse{}
	_body, _err := client.RecoveryFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 续费安卓实例组
//
// @param request - RenewAndroidInstanceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAndroidInstanceGroupsResponse
func (client *Client) RenewAndroidInstanceGroupsWithOptions(request *RenewAndroidInstanceGroupsRequest, runtime *util.RuntimeOptions) (_result *RenewAndroidInstanceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIds)) {
		query["InstanceGroupIds"] = request.InstanceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewAndroidInstanceGroups"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RenewAndroidInstanceGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RenewAndroidInstanceGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 续费安卓实例组
//
// @param request - RenewAndroidInstanceGroupsRequest
//
// @return RenewAndroidInstanceGroupsResponse
func (client *Client) RenewAndroidInstanceGroups(request *RenewAndroidInstanceGroupsRequest) (_result *RenewAndroidInstanceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewAndroidInstanceGroupsResponse{}
	_body, _err := client.RenewAndroidInstanceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置安卓实例
//
// @param request - ResetAndroidInstancesInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAndroidInstancesInGroupResponse
func (client *Client) ResetAndroidInstancesInGroupWithOptions(request *ResetAndroidInstancesInGroupRequest, runtime *util.RuntimeOptions) (_result *ResetAndroidInstancesInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAndroidInstancesInGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResetAndroidInstancesInGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResetAndroidInstancesInGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 重置安卓实例
//
// @param request - ResetAndroidInstancesInGroupRequest
//
// @return ResetAndroidInstancesInGroupResponse
func (client *Client) ResetAndroidInstancesInGroup(request *ResetAndroidInstancesInGroupRequest) (_result *ResetAndroidInstancesInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAndroidInstancesInGroupResponse{}
	_body, _err := client.ResetAndroidInstancesInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过eds agent通道下发命令
//
// @param request - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunCommandResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunCommandResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 通过eds agent通道下发命令
//
// @param request - RunCommandRequest
//
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送文件到云手机
//
// @param request - SendFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendFileResponse
func (client *Client) SendFileWithOptions(request *SendFileRequest, runtime *util.RuntimeOptions) (_result *SendFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIdList)) {
		query["AndroidInstanceIdList"] = request.AndroidInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFilePath)) {
		query["SourceFilePath"] = request.SourceFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.UploadEndpoint)) {
		query["UploadEndpoint"] = request.UploadEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["UploadType"] = request.UploadType
	}

	if !tea.BoolValue(util.IsUnset(request.UploadUrl)) {
		query["UploadUrl"] = request.UploadUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendFile"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SendFileResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SendFileResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 推送文件到云手机
//
// @param request - SendFileRequest
//
// @return SendFileResponse
func (client *Client) SendFile(request *SendFileRequest) (_result *SendFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendFileResponse{}
	_body, _err := client.SendFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetAdbSecureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAdbSecureResponse
func (client *Client) SetAdbSecureWithOptions(request *SetAdbSecureRequest, runtime *util.RuntimeOptions) (_result *SetAdbSecureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAdbSecure"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SetAdbSecureResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SetAdbSecureResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - SetAdbSecureRequest
//
// @return SetAdbSecureResponse
func (client *Client) SetAdbSecure(request *SetAdbSecureRequest) (_result *SetAdbSecureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAdbSecureResponse{}
	_body, _err := client.SetAdbSecureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start instances.
//
// Description:
//
// Only supports starting when the instance is in the **Stopped, Backup Failed, or Recovery Failed*	- state.
//
// @param request - StartAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAndroidInstanceResponse
func (client *Client) StartAndroidInstanceWithOptions(request *StartAndroidInstanceRequest, runtime *util.RuntimeOptions) (_result *StartAndroidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAndroidInstance"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartAndroidInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartAndroidInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Start instances.
//
// Description:
//
// Only supports starting when the instance is in the **Stopped, Backup Failed, or Recovery Failed*	- state.
//
// @param request - StartAndroidInstanceRequest
//
// @return StartAndroidInstanceResponse
func (client *Client) StartAndroidInstance(request *StartAndroidInstanceRequest) (_result *StartAndroidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAndroidInstanceResponse{}
	_body, _err := client.StartAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例关机
//
// @param request - StopAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAndroidInstanceResponse
func (client *Client) StopAndroidInstanceWithOptions(request *StopAndroidInstanceRequest, runtime *util.RuntimeOptions) (_result *StopAndroidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		query["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ForceStop)) {
		query["ForceStop"] = request.ForceStop
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAndroidInstance"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopAndroidInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopAndroidInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 实例关机
//
// @param request - StopAndroidInstanceRequest
//
// @return StopAndroidInstanceResponse
func (client *Client) StopAndroidInstance(request *StopAndroidInstanceRequest) (_result *StopAndroidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAndroidInstanceResponse{}
	_body, _err := client.StopAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载app
//
// @param request - UninstallAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAppResponse
func (client *Client) UninstallAppWithOptions(request *UninstallAppRequest, runtime *util.RuntimeOptions) (_result *UninstallAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIdList)) {
		query["AppIdList"] = request.AppIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIdList)) {
		query["InstanceGroupIdList"] = request.InstanceGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallApp"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UninstallAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UninstallAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 卸载app
//
// @param request - UninstallAppRequest
//
// @return UninstallAppResponse
func (client *Client) UninstallApp(request *UninstallAppRequest) (_result *UninstallAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallAppResponse{}
	_body, _err := client.UninstallAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of a custom image.
//
// @param request - UpdateCustomImageNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomImageNameResponse
func (client *Client) UpdateCustomImageNameWithOptions(request *UpdateCustomImageNameRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomImageNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		body["ImageName"] = request.ImageName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomImageName"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateCustomImageNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateCustomImageNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the name of a custom image.
//
// @param request - UpdateCustomImageNameRequest
//
// @return UpdateCustomImageNameResponse
func (client *Client) UpdateCustomImageName(request *UpdateCustomImageNameRequest) (_result *UpdateCustomImageNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomImageNameResponse{}
	_body, _err := client.UpdateCustomImageNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例组变更镜像
//
// @param request - UpdateInstanceGroupImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceGroupImageResponse
func (client *Client) UpdateInstanceGroupImageWithOptions(request *UpdateInstanceGroupImageRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceGroupImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupIds)) {
		body["InstanceGroupIds"] = request.InstanceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceGroupImage"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceGroupImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceGroupImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 实例组变更镜像
//
// @param request - UpdateInstanceGroupImageRequest
//
// @return UpdateInstanceGroupImageResponse
func (client *Client) UpdateInstanceGroupImage(request *UpdateInstanceGroupImageRequest) (_result *UpdateInstanceGroupImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceGroupImageResponse{}
	_body, _err := client.UpdateInstanceGroupImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades an instance group. Currently, this operation allows you to only increase the number of instances in an instance group.
//
// Description:
//
// Currently, this operation allows you to only increase the size of an instance group.
//
// @param request - UpgradeAndroidInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAndroidInstanceGroupResponse
func (client *Client) UpgradeAndroidInstanceGroupWithOptions(request *UpgradeAndroidInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *UpgradeAndroidInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.IncreaseNumberOfInstance)) {
		query["IncreaseNumberOfInstance"] = request.IncreaseNumberOfInstance
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGroupId)) {
		query["InstanceGroupId"] = request.InstanceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeAndroidInstanceGroup"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeAndroidInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeAndroidInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Upgrades an instance group. Currently, this operation allows you to only increase the number of instances in an instance group.
//
// Description:
//
// Currently, this operation allows you to only increase the size of an instance group.
//
// @param request - UpgradeAndroidInstanceGroupRequest
//
// @return UpgradeAndroidInstanceGroupResponse
func (client *Client) UpgradeAndroidInstanceGroup(request *UpgradeAndroidInstanceGroupRequest) (_result *UpgradeAndroidInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeAndroidInstanceGroupResponse{}
	_body, _err := client.UpgradeAndroidInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

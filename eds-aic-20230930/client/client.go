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
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the ADB key pair.
	//
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
	// The object that is returned.
	Data *AttachKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
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
	// The IDs of the cloud phone instances to which the ADB key pair is successfully attached.
	AttachedInstanceIds []*string `json:"AttachedInstanceIds,omitempty" xml:"AttachedInstanceIds,omitempty" type:"Repeated"`
	// The number of the cloud phone instances to which the ADB key pair failed to be attached.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The total number of the cloud phone instances.
	//
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
	// Specifies whether to back up the whole instance.
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The name of the backup file.
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
	// The names of the application packages that you want to back up.
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
	// The number of instances that are backed up.
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The object that is returned.
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
	// The ID of the batch task.
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
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
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
	// The ID of the user to whom the cloud phone instance is assigned.
	//
	// example:
	//
	// user
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-25nt4kk9whjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The IDs of the cloud phone instances. You can specify 1 to 100 IDs of cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The instance connection tasks.
	InstanceTasks []*BatchGetAcpConnectionTicketRequestInstanceTasks `json:"InstanceTasks,omitempty" xml:"InstanceTasks,omitempty" type:"Repeated"`
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
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-fkuit0cmyfvzz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the task.
	//
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
	// The results of the instance connection tasks.
	InstanceConnectionModels []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels `json:"InstanceConnectionModels,omitempty" xml:"InstanceConnectionModels,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-1uzb6heg797z3****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ErrorCode          *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-ajxvwo1u0hqvd****
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PersistentAppInstanceId *string `json:"PersistentAppInstanceId,omitempty" xml:"PersistentAppInstanceId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// cn-hangzhou@c9f5c2e8-f5c4-4b01-8602-000cae94****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The state of the task.
	//
	// example:
	//
	// FINISHED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The ticket used to connect to the cloud phone instance.
	//
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

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetAppInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.AppInstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetErrorCode(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.ErrorCode = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.InstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetPersistentAppInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.PersistentAppInstanceId = &v
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

type ChangeCloudPhoneNodeRequest struct {
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cpn-0ugbptfu473fy****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 20
	PhoneCount *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
}

func (s ChangeCloudPhoneNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeRequest) SetInstanceType(v string) *ChangeCloudPhoneNodeRequest {
	s.InstanceType = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetNodeId(v string) *ChangeCloudPhoneNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetPhoneCount(v int32) *ChangeCloudPhoneNodeRequest {
	s.PhoneCount = &v
	return s
}

type ChangeCloudPhoneNodeResponseBody struct {
	NodeInfos []*ChangeCloudPhoneNodeResponseBodyNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 4610632D-D661-5982-B3D7-5D3FD183F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeCloudPhoneNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponseBody) SetNodeInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfos) *ChangeCloudPhoneNodeResponseBody {
	s.NodeInfos = v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBody) SetRequestId(v string) *ChangeCloudPhoneNodeResponseBody {
	s.RequestId = &v
	return s
}

type ChangeCloudPhoneNodeResponseBodyNodeInfos struct {
	InstanceInfos []*ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// cpn-e5kxgjyt8s1mb****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfos) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetInstanceInfos(v []*ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.InstanceInfos = v
	return s
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfos) SetNodeId(v string) *ChangeCloudPhoneNodeResponseBodyNodeInfos {
	s.NodeId = &v
	return s
}

type ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos struct {
	// example:
	//
	// cpn-jewjt8xryuitu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) String() string {
	return tea.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos) SetInstanceId(v string) *ChangeCloudPhoneNodeResponseBodyNodeInfosInstanceInfos {
	s.InstanceId = &v
	return s
}

type ChangeCloudPhoneNodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCloudPhoneNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCloudPhoneNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponse) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponse) SetHeaders(v map[string]*string) *ChangeCloudPhoneNodeResponse {
	s.Headers = v
	return s
}

func (s *ChangeCloudPhoneNodeResponse) SetStatusCode(v int32) *ChangeCloudPhoneNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponse) SetBody(v *ChangeCloudPhoneNodeResponseBody) *ChangeCloudPhoneNodeResponse {
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
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Specifies whether to enable GPU acceleration.
	//
	// Valid values:
	//
	// 	- true: enables GPU acceleration.
	//
	// 	- false (default): disables GPU acceleration.
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
	// >  The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with `http://` or `https://`.
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
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	Ipv6Bandwidth *int32 `json:"Ipv6Bandwidth,omitempty" xml:"Ipv6Bandwidth,omitempty"`
	// The ID of the key pair. When you create an instance group and specify a valid key pair ID, all cloud phone instances within the group will automatically be bound to that key pair upon creation. This eliminates the need to manually bind key pairs to individual cloud phone instances.
	//
	// >  Binding key pairs to cloud phone instances is currently not supported during instance group resizing.
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
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The tags
	Tag []*CreateAndroidInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
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
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The information about the custom app.
	//
	// >
	//
	// 	- If you want to pass in a custom app, configure the `CustomAppInfo` parameter. Take note that the six fields within it are mandatory.
	//
	// 	- A custom app has a higher priority than an app from the Alibaba Cloud Workspace Application Center. If you configure the `CustomAppInfo` parameter, the `FileName` and `FilePath` pair or the `OssAppUrl` will not take effect.
	CustomAppInfo *CreateAppRequestCustomAppInfo `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty" type:"Struct"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name used by the app file in Object Storage Service (OSS). This parameter, combined with `FilePath`, uniquely identifies the OSS path of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [Elastic Desktop Service (EDS) Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	//
	// example:
	//
	// testApp.apk
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS bucket path to the app file. This parameter, combined with `FileName`, uniquely identifies the OSS path of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [EDS Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
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
	// The OSS bucket endpoint of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [EDS Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	//
	// example:
	//
	// http://testApp.apk
	OssAppUrl *string `json:"OssAppUrl,omitempty" xml:"OssAppUrl,omitempty"`
	SignApk   *string `json:"SignApk,omitempty" xml:"SignApk,omitempty"`
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

func (s *CreateAppRequest) SetSignApk(v string) *CreateAppRequest {
	s.SignApk = &v
	return s
}

type CreateAppRequestCustomAppInfo struct {
	// The size of the .apk file. Unit: MB.
	//
	// example:
	//
	// 10
	ApkSize *string `json:"ApkSize,omitempty" xml:"ApkSize,omitempty"`
	// The download URL of the app.
	//
	// example:
	//
	// http://testApp.apk
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The MD5 value of the .apk file.
	//
	// example:
	//
	// df3f46ce5844ddb278f14c5a9cd2****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the app package.
	//
	// example:
	//
	// com.example.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The version of the app.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The code of the app version.
	//
	// example:
	//
	// 10000
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
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The information about the custom app.
	//
	// >
	//
	// 	- If you want to pass in a custom app, configure the `CustomAppInfo` parameter. Take note that the six fields within it are mandatory.
	//
	// 	- A custom app has a higher priority than an app from the Alibaba Cloud Workspace Application Center. If you configure the `CustomAppInfo` parameter, the `FileName` and `FilePath` pair or the `OssAppUrl` will not take effect.
	CustomAppInfoShrink *string `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name used by the app file in Object Storage Service (OSS). This parameter, combined with `FilePath`, uniquely identifies the OSS path of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [Elastic Desktop Service (EDS) Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	//
	// example:
	//
	// testApp.apk
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS bucket path to the app file. This parameter, combined with `FileName`, uniquely identifies the OSS path of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [EDS Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
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
	// The OSS bucket endpoint of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [EDS Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	//
	// example:
	//
	// http://testApp.apk
	OssAppUrl *string `json:"OssAppUrl,omitempty" xml:"OssAppUrl,omitempty"`
	SignApk   *string `json:"SignApk,omitempty" xml:"SignApk,omitempty"`
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

func (s *CreateAppShrinkRequest) SetSignApk(v string) *CreateAppShrinkRequest {
	s.SignApk = &v
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

type CreateCloudPhoneNodeRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- False (default): You must manually complete the payment in the Alibaba Cloud Expenses and Costs console.
	//
	// 	- true: enables the auto-payment feature.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-renewal feature. In this case, the system automatically renews instances upon expiration.
	//
	// 	- false (default): disables the auto-renewal feature. In this case, you need to manually renew instances upon expiration.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method. Only the subscription billing method is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of cloud phone matrixes you want to purchase.
	//
	// example:
	//
	// 1
	Count         *string                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	DisplayConfig *CreateCloudPhoneNodeRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	// The image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance specification.
	//
	// Valid values:
	//
	// 	- ac.max: By default, this specification allows up to 25 instances. You can adjust this number by using PhoneCount (Value range: 4 to 40).
	//
	// 	- ac.plus: By default, this specification allows up to 40 instances. You can adjust this number by using PhoneCount (Value range: 4 to 40).
	//
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-5mwr9azebliva****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The name of the cloud phone matrix.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The subscription duration. The unit is specified by `PeriodUnit`. Valid values:
	//
	// 	- When `PeriodUnit` is set to **year**: 1.
	//
	// 	- When `PeriodUnit` is set to **month**: 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The number of instances per cloud phone matrix.
	//
	// example:
	//
	// 25
	PhoneCount      *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	// The resolution height. Unit: pixel.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution width. Unit: pixel.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The shared storage size Unit: GiB.
	//
	// example:
	//
	// 200
	ServerShareDataVolume *int32 `json:"ServerShareDataVolume,omitempty" xml:"ServerShareDataVolume,omitempty"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// This parameter is required.
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	StreamMode *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	// The resource tags.
	Tag []*CreateCloudPhoneNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeekryyc1q3sm72l****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateCloudPhoneNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequest) SetAutoPay(v bool) *CreateCloudPhoneNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetAutoRenew(v bool) *CreateCloudPhoneNodeRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetBizRegionId(v string) *CreateCloudPhoneNodeRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetChargeType(v string) *CreateCloudPhoneNodeRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetCount(v string) *CreateCloudPhoneNodeRequest {
	s.Count = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetDisplayConfig(v *CreateCloudPhoneNodeRequestDisplayConfig) *CreateCloudPhoneNodeRequest {
	s.DisplayConfig = v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetImageId(v string) *CreateCloudPhoneNodeRequest {
	s.ImageId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetInstanceType(v string) *CreateCloudPhoneNodeRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetNetworkId(v string) *CreateCloudPhoneNodeRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetNodeName(v string) *CreateCloudPhoneNodeRequest {
	s.NodeName = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPeriod(v int32) *CreateCloudPhoneNodeRequest {
	s.Period = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPeriodUnit(v string) *CreateCloudPhoneNodeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPhoneCount(v int32) *CreateCloudPhoneNodeRequest {
	s.PhoneCount = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeRequest {
	s.PhoneDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetResolutionHeight(v int32) *CreateCloudPhoneNodeRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetResolutionWidth(v int32) *CreateCloudPhoneNodeRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetServerShareDataVolume(v int32) *CreateCloudPhoneNodeRequest {
	s.ServerShareDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetServerType(v string) *CreateCloudPhoneNodeRequest {
	s.ServerType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetStreamMode(v int32) *CreateCloudPhoneNodeRequest {
	s.StreamMode = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetTag(v []*CreateCloudPhoneNodeRequestTag) *CreateCloudPhoneNodeRequest {
	s.Tag = v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetVSwitchId(v string) *CreateCloudPhoneNodeRequest {
	s.VSwitchId = &v
	return s
}

type CreateCloudPhoneNodeRequestDisplayConfig struct {
	Dpi            *int32  `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	Fps            *int32  `json:"Fps,omitempty" xml:"Fps,omitempty"`
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
}

func (s CreateCloudPhoneNodeRequestDisplayConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeRequestDisplayConfig) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) SetDpi(v int32) *CreateCloudPhoneNodeRequestDisplayConfig {
	s.Dpi = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) SetFps(v int32) *CreateCloudPhoneNodeRequestDisplayConfig {
	s.Fps = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) SetLockResolution(v string) *CreateCloudPhoneNodeRequestDisplayConfig {
	s.LockResolution = &v
	return s
}

type CreateCloudPhoneNodeRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// keyname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// valuename
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCloudPhoneNodeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequestTag) SetKey(v string) *CreateCloudPhoneNodeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestTag) SetValue(v string) *CreateCloudPhoneNodeRequestTag {
	s.Value = &v
	return s
}

type CreateCloudPhoneNodeShrinkRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- False (default): You must manually complete the payment in the Alibaba Cloud Expenses and Costs console.
	//
	// 	- true: enables the auto-payment feature.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-renewal feature. In this case, the system automatically renews instances upon expiration.
	//
	// 	- false (default): disables the auto-renewal feature. In this case, you need to manually renew instances upon expiration.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method. Only the subscription billing method is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of cloud phone matrixes you want to purchase.
	//
	// example:
	//
	// 1
	Count               *string `json:"Count,omitempty" xml:"Count,omitempty"`
	DisplayConfigShrink *string `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance specification.
	//
	// Valid values:
	//
	// 	- ac.max: By default, this specification allows up to 25 instances. You can adjust this number by using PhoneCount (Value range: 4 to 40).
	//
	// 	- ac.plus: By default, this specification allows up to 40 instances. You can adjust this number by using PhoneCount (Value range: 4 to 40).
	//
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-5mwr9azebliva****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The name of the cloud phone matrix.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The subscription duration. The unit is specified by `PeriodUnit`. Valid values:
	//
	// 	- When `PeriodUnit` is set to **year**: 1.
	//
	// 	- When `PeriodUnit` is set to **month**: 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The number of instances per cloud phone matrix.
	//
	// example:
	//
	// 25
	PhoneCount      *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	// The resolution height. Unit: pixel.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution width. Unit: pixel.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The shared storage size Unit: GiB.
	//
	// example:
	//
	// 200
	ServerShareDataVolume *int32 `json:"ServerShareDataVolume,omitempty" xml:"ServerShareDataVolume,omitempty"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// This parameter is required.
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	StreamMode *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	// The resource tags.
	Tag []*CreateCloudPhoneNodeShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeekryyc1q3sm72l****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateCloudPhoneNodeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetAutoPay(v bool) *CreateCloudPhoneNodeShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetAutoRenew(v bool) *CreateCloudPhoneNodeShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetBizRegionId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetChargeType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetCount(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.Count = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetDisplayConfigShrink(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.DisplayConfigShrink = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetImageId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetInstanceType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetNetworkId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetNodeName(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.NodeName = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPeriod(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPeriodUnit(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPhoneCount(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.PhoneCount = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.PhoneDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetResolutionHeight(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetResolutionWidth(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetServerShareDataVolume(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.ServerShareDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetServerType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ServerType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetStreamMode(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.StreamMode = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetTag(v []*CreateCloudPhoneNodeShrinkRequestTag) *CreateCloudPhoneNodeShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetVSwitchId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.VSwitchId = &v
	return s
}

type CreateCloudPhoneNodeShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// keyname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// valuename
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCloudPhoneNodeShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) SetKey(v string) *CreateCloudPhoneNodeShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) SetValue(v string) *CreateCloudPhoneNodeShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateCloudPhoneNodeResponseBody struct {
	// The cloud phone matrixes.
	NodeInfos []*CreateCloudPhoneNodeResponseBodyNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	// The order ID.
	//
	// example:
	//
	// 223684716098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudPhoneNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponseBody) SetNodeInfos(v []*CreateCloudPhoneNodeResponseBodyNodeInfos) *CreateCloudPhoneNodeResponseBody {
	s.NodeInfos = v
	return s
}

func (s *CreateCloudPhoneNodeResponseBody) SetOrderId(v string) *CreateCloudPhoneNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCloudPhoneNodeResponseBody) SetRequestId(v string) *CreateCloudPhoneNodeResponseBody {
	s.RequestId = &v
	return s
}

type CreateCloudPhoneNodeResponseBodyNodeInfos struct {
	// The IDs of the cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the cloud phone matrix.
	//
	// example:
	//
	// cpn-e5kxgjyt8s1mb****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateCloudPhoneNodeResponseBodyNodeInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeResponseBodyNodeInfos) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) SetInstanceIds(v []*string) *CreateCloudPhoneNodeResponseBodyNodeInfos {
	s.InstanceIds = v
	return s
}

func (s *CreateCloudPhoneNodeResponseBodyNodeInfos) SetNodeId(v string) *CreateCloudPhoneNodeResponseBodyNodeInfos {
	s.NodeId = &v
	return s
}

type CreateCloudPhoneNodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudPhoneNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudPhoneNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudPhoneNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponse) SetHeaders(v map[string]*string) *CreateCloudPhoneNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudPhoneNodeResponse) SetStatusCode(v int32) *CreateCloudPhoneNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudPhoneNodeResponse) SetBody(v *CreateCloudPhoneNodeResponseBody) *CreateCloudPhoneNodeResponse {
	s.Body = v
	return s
}

type CreateCustomImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. By default, this parameter is left empty. The token cannot exceed 64 characters in length.
	//
	// example:
	//
	// 20393E53-8FF1-524C-B494-B478A5369733
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the custom image.
	//
	// example:
	//
	// create for cc5g group auth rules test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom image name
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the cloud phone instance.
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
	// The ID of the custom image.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the request.
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
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	ResolutionWidth *int32                             `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	Watermark       *CreatePolicyGroupRequestWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
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

func (s *CreatePolicyGroupRequest) SetPolicyType(v string) *CreatePolicyGroupRequest {
	s.PolicyType = &v
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

func (s *CreatePolicyGroupRequest) SetWatermark(v *CreatePolicyGroupRequestWatermark) *CreatePolicyGroupRequest {
	s.Watermark = v
	return s
}

type CreatePolicyGroupRequestNetRedirectPolicy struct {
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

type CreatePolicyGroupRequestWatermark struct {
	WatermarkColor             *int32    `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	WatermarkCustomText        *string   `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkFontSize          *int32    `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	WatermarkSwitch            *string   `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	WatermarkTransparencyValue *int32    `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	WatermarkTypes             []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s CreatePolicyGroupRequestWatermark) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestWatermark) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkColor(v int32) *CreatePolicyGroupRequestWatermark {
	s.WatermarkColor = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkCustomText(v string) *CreatePolicyGroupRequestWatermark {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkFontSize(v int32) *CreatePolicyGroupRequestWatermark {
	s.WatermarkFontSize = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkSwitch(v string) *CreatePolicyGroupRequestWatermark {
	s.WatermarkSwitch = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkTransparencyValue(v int32) *CreatePolicyGroupRequestWatermark {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkTypes(v []*string) *CreatePolicyGroupRequestWatermark {
	s.WatermarkTypes = v
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
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	ResolutionWidth *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
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

func (s *CreatePolicyGroupShrinkRequest) SetPolicyType(v string) *CreatePolicyGroupShrinkRequest {
	s.PolicyType = &v
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

func (s *CreatePolicyGroupShrinkRequest) SetWatermarkShrink(v string) *CreatePolicyGroupShrinkRequest {
	s.WatermarkShrink = &v
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
	// The IDs of the cloud phone instances. You can create multiple snapshots simultaneously.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The name of the OSS bucket. The name must start with "cloudphone-saved-bucket-". The OSS bucket and the cloud phone instance must be in the same region. If you leave this parameter empty, the system will create a default OSS bucket named “cloudphone-saved-bucket-{Region of the cloud phone instance}-{AliUid}.”
	//
	// example:
	//
	// cloudphone-saved-bucket-cn-shanghai-default
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Specifies whether to bypass the snapshot policy control. Default value: false.
	//
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
	// The ID of the request. If the request fails, share this ID with technical support to help diagnose the issue.
	//
	// example:
	//
	// 3AF82CE1-2801-52CE-BF64-B491DD7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*CreateScreenshotResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-bwhtebzah2fse****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The ID of the task. You can use the task ID with the DescribeTasks operation to get the download link for the screenshot.
	//
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

type DeleteBackupFileRequest struct {
	// This parameter is required.
	BackupFileIdList []*string `json:"BackupFileIdList,omitempty" xml:"BackupFileIdList,omitempty" type:"Repeated"`
}

func (s DeleteBackupFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileRequest) SetBackupFileIdList(v []*string) *DeleteBackupFileRequest {
	s.BackupFileIdList = v
	return s
}

type DeleteBackupFileResponseBody struct {
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileResponseBody) SetRequestId(v string) *DeleteBackupFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBackupFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileResponse) SetHeaders(v map[string]*string) *DeleteBackupFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupFileResponse) SetStatusCode(v int32) *DeleteBackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupFileResponse) SetBody(v *DeleteBackupFileResponseBody) *DeleteBackupFileResponse {
	s.Body = v
	return s
}

type DeleteCloudPhoneNodesRequest struct {
	// The cloud phone matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s DeleteCloudPhoneNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudPhoneNodesRequest) SetNodeIds(v []*string) *DeleteCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

type DeleteCloudPhoneNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudPhoneNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudPhoneNodesResponseBody) SetRequestId(v string) *DeleteCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCloudPhoneNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudPhoneNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *DeleteCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudPhoneNodesResponse) SetStatusCode(v int32) *DeleteCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudPhoneNodesResponse) SetBody(v *DeleteCloudPhoneNodesResponseBody) *DeleteCloudPhoneNodesResponse {
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
	// The purchase mode of cloud phone instances.
	//
	// Valid values:
	//
	// 	- Instance (default): the instance group mode.
	//
	// 	- Node: the matrix mode [whitelisted].
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
	// The instance group.
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
	// The number of available instances.
	//
	// >  Available instances are those not in the Deleting or Deleted state.
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
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The cause of the creation failure.
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
	// example:
	//
	// 50
	Ipv6Bandwidth *int32 `json:"Ipv6Bandwidth,omitempty" xml:"Ipv6Bandwidth,omitempty"`
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
	// The rendering mode of the instance group.
	//
	// Valid values:
	//
	// 	- GPURemote: GPU remote rendering.
	//
	// 	- CPU: CPU rendering.
	//
	// 	- GPUocal: GPU local rendering.
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

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetEnableIpv6(v bool) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.EnableIpv6 = &v
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

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetIpv6Bandwidth(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Ipv6Bandwidth = &v
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
	AppManagePolicyId   *string `json:"AppManagePolicyId,omitempty" xml:"AppManagePolicyId,omitempty"`
	AuthorizedUserId    *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2807298.html) operation to query the regions where Cloud Phone is supported.
	//
	// example:
	//
	// cn-shanghai
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
	// The ID of the node.
	//
	// example:
	//
	// node_id
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// node_name
	NodeName      *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OfficeSiteIds []*string `json:"OfficeSiteIds,omitempty" xml:"OfficeSiteIds,omitempty" type:"Repeated"`
	QosRuleIds    []*string `json:"QosRuleIds,omitempty" xml:"QosRuleIds,omitempty" type:"Repeated"`
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

func (s *DescribeAndroidInstancesRequest) SetAppManagePolicyId(v string) *DescribeAndroidInstancesRequest {
	s.AppManagePolicyId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetAuthorizedUserId(v string) *DescribeAndroidInstancesRequest {
	s.AuthorizedUserId = &v
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

func (s *DescribeAndroidInstancesRequest) SetOfficeSiteIds(v []*string) *DescribeAndroidInstancesRequest {
	s.OfficeSiteIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetQosRuleIds(v []*string) *DescribeAndroidInstancesRequest {
	s.QosRuleIds = v
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
	// The cloud phone instances.
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
	AppInstanceId   *string                                                           `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppManagePolicy *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy `json:"AppManagePolicy,omitempty" xml:"AppManagePolicy,omitempty" type:"Struct"`
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
	Disks         []*DescribeAndroidInstancesResponseBodyInstanceModelDisks       `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	DisplayConfig *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
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
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	NetworkInterfaceIpv6Address *string `json:"NetworkInterfaceIpv6Address,omitempty" xml:"NetworkInterfaceIpv6Address,omitempty"`
	// The office network ID.
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
	PersistentAppInstanceId *string                                                         `json:"PersistentAppInstanceId,omitempty" xml:"PersistentAppInstanceId,omitempty"`
	PhoneDataInfo           *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo `json:"PhoneDataInfo,omitempty" xml:"PhoneDataInfo,omitempty" type:"Struct"`
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
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	PublicIpv6Address *string `json:"PublicIpv6Address,omitempty" xml:"PublicIpv6Address,omitempty"`
	QosRuleId         *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rendering type.
	//
	// example:
	//
	// local
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	// The session status.
	//
	// Valid values:
	//
	// 	- disConnect: The session is disconnected.
	//
	// 	- connect: The session is connected.
	//
	// example:
	//
	// connect
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	StreamMode    *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	// The tags.
	Tags      []*DescribeAndroidInstancesResponseBodyInstanceModelTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VSwitchId *string                                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAppManagePolicy(v *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AppManagePolicy = v
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

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetDisplayConfig(v *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.DisplayConfig = v
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

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetImageId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ImageId = &v
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

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPhoneDataInfo(v *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PhoneDataInfo = v
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

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetQosRuleId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.QosRuleId = &v
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

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetStreamMode(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.StreamMode = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetTags(v []*DescribeAndroidInstancesResponseBodyInstanceModelTags) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Tags = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetVSwitchId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetZoneId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ZoneId = &v
	return s
}

type DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy struct {
	AppManagePolicyId   *string `json:"AppManagePolicyId,omitempty" xml:"AppManagePolicyId,omitempty"`
	AppManagePolicyName *string `json:"AppManagePolicyName,omitempty" xml:"AppManagePolicyName,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) SetAppManagePolicyId(v string) *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy {
	s.AppManagePolicyId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) SetAppManagePolicyName(v string) *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy {
	s.AppManagePolicyName = &v
	return s
}

type DescribeAndroidInstancesResponseBodyInstanceModelDisks struct {
	// The disk size. Unit: GB.
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

type DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig struct {
	Dpi              *int32  `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	Fps              *int32  `json:"Fps,omitempty" xml:"Fps,omitempty"`
	LockResolution   *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	ResolutionHeight *int32  `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	ResolutionWidth  *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetDpi(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.Dpi = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetFps(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.Fps = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetLockResolution(v string) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.LockResolution = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetResolutionHeight(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetResolutionWidth(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.ResolutionWidth = &v
	return s
}

type DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo struct {
	PhoneDataId     *string `json:"PhoneDataId,omitempty" xml:"PhoneDataId,omitempty"`
	PhoneDataVolume *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) SetPhoneDataId(v string) *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo {
	s.PhoneDataId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) SetPhoneDataVolume(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo {
	s.PhoneDataVolume = &v
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
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
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

func (s *DescribeAppsRequest) SetAppType(v string) *DescribeAppsRequest {
	s.AppType = &v
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
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
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

func (s *DescribeAppsResponseBodyData) SetAppType(v string) *DescribeAppsResponseBodyData {
	s.AppType = &v
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
	// The name of the instance. Fuzzy match is supported.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// Specifies whether the whole instance is backed up.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The name of the backup file. Fuzzy match is supported.
	//
	// example:
	//
	// defaulBackupFile
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The description of the backup file. Fuzzy match is supported.
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
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
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
	// The status of the backup files.
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
	// Indicates whether the whole instance is backed up.
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
	// The names of the application packages that are backed up.
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

type DescribeCloudPhoneNodesRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method. Only the subscription billing method is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If a query doesn\\"t return all results, the response includes a NextToken value for pagination. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The matrix name.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The matrix status.
	//
	// Valid values:
	//
	// 	- FAILED: The matrix failed to be created.
	//
	// 	- RUNNING: The matrix is available.
	//
	// 	- DELETING: The matrix is being deleted.
	//
	// 	- NODE_READY: The matrix is ready, and cloud phone instances are being created.
	//
	// 	- DELETED: The matrix is deleted.
	//
	// 	- CREATING: The matrix is being created.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudPhoneNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesRequest) SetBizRegionId(v string) *DescribeCloudPhoneNodesRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetChargeType(v string) *DescribeCloudPhoneNodesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetMaxResults(v string) *DescribeCloudPhoneNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNextToken(v string) *DescribeCloudPhoneNodesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNodeIds(v []*string) *DescribeCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetNodeName(v string) *DescribeCloudPhoneNodesRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetServerType(v string) *DescribeCloudPhoneNodesRequest {
	s.ServerType = &v
	return s
}

func (s *DescribeCloudPhoneNodesRequest) SetStatus(v string) *DescribeCloudPhoneNodesRequest {
	s.Status = &v
	return s
}

type DescribeCloudPhoneNodesResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- ****
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The matrixes.
	NodeModel []*DescribeCloudPhoneNodesResponseBodyNodeModel `json:"NodeModel,omitempty" xml:"NodeModel,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud phone instances.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBody) SetMaxResults(v int32) *DescribeCloudPhoneNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetNextToken(v string) *DescribeCloudPhoneNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetNodeModel(v []*DescribeCloudPhoneNodesResponseBodyNodeModel) *DescribeCloudPhoneNodesResponseBody {
	s.NodeModel = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetRequestId(v string) *DescribeCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetTotalCount(v int32) *DescribeCloudPhoneNodesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCloudPhoneNodesResponseBodyNodeModel struct {
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-11-13 02:03:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the subscription matrix.
	//
	// example:
	//
	// 2025-03-09 02:00:34
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2025-02-13 02:03:14
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 32
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-5mwr9azebliva****
	NetworkId    *string                                                     `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkInfos []*DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos `json:"NetworkInfos,omitempty" xml:"NetworkInfos,omitempty" type:"Repeated"`
	// The matrix ID.
	//
	// example:
	//
	// cpn-ehs0yoedq8ntm****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The matrix name.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of cloud phone instances per matrix.
	//
	// example:
	//
	// 25
	PhoneCount    *int32                                                     `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	PhoneDataInfo *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo `json:"PhoneDataInfo,omitempty" xml:"PhoneDataInfo,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The height of the resolution. Unit: pixel.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixel.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The matrix specification.
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The size of the shared storage. Unit: GiB.
	//
	// example:
	//
	// 100
	ShareDataVolume *int32 `json:"ShareDataVolume,omitempty" xml:"ShareDataVolume,omitempty"`
	// The matrix status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeekryyc1q3sm72l****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModel) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetChargeType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ChargeType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetCpu(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.Cpu = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetGmtCreate(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetGmtExpired(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.GmtExpired = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetGmtModified(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetInstanceType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.InstanceType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetMemory(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.Memory = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNetworkId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NetworkId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNetworkInfos(v []*DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NetworkInfos = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNodeId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NodeId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNodeName(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NodeName = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetPhoneCount(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.PhoneCount = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetPhoneDataInfo(v *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.PhoneDataInfo = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetRegionId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetResolutionHeight(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetResolutionWidth(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetServerType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ServerType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetShareDataVolume(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ShareDataVolume = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetStatus(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.Status = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetVSwitchId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.VSwitchId = &v
	return s
}

type DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos struct {
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetNetworkId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.NetworkId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetVSwitchId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.VSwitchId = &v
	return s
}

type DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo struct {
	PhoneDataId     *string `json:"PhoneDataId,omitempty" xml:"PhoneDataId,omitempty"`
	PhoneDataVolume *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) SetPhoneDataId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo {
	s.PhoneDataId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) SetPhoneDataVolume(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo {
	s.PhoneDataVolume = &v
	return s
}

type DescribeCloudPhoneNodesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudPhoneNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *DescribeCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudPhoneNodesResponse) SetStatusCode(v int32) *DescribeCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponse) SetBody(v *DescribeCloudPhoneNodesResponseBody) *DescribeCloudPhoneNodesResponse {
	s.Body = v
	return s
}

type DescribeDisplayConfigRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeDisplayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisplayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigRequest) SetAndroidInstanceIds(v []*string) *DescribeDisplayConfigRequest {
	s.AndroidInstanceIds = v
	return s
}

type DescribeDisplayConfigResponseBody struct {
	DisplayConfigModel []*DescribeDisplayConfigResponseBodyDisplayConfigModel `json:"DisplayConfigModel,omitempty" xml:"DisplayConfigModel,omitempty" type:"Repeated"`
	// example:
	//
	// FFEF7EFE-1E36-56D1-B5BF-5BACE43B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDisplayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisplayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigResponseBody) SetDisplayConfigModel(v []*DescribeDisplayConfigResponseBodyDisplayConfigModel) *DescribeDisplayConfigResponseBody {
	s.DisplayConfigModel = v
	return s
}

func (s *DescribeDisplayConfigResponseBody) SetRequestId(v string) *DescribeDisplayConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDisplayConfigResponseBodyDisplayConfigModel struct {
	// example:
	//
	// cpn-jewjt8xryuituz4qn-****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// example:
	//
	// 240
	Dpi *int32 `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	// example:
	//
	// null
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s DescribeDisplayConfigResponseBodyDisplayConfigModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisplayConfigResponseBodyDisplayConfigModel) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetAndroidInstanceId(v string) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetDpi(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.Dpi = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetFps(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.Fps = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetLockResolution(v string) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.LockResolution = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetResolutionHeight(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetResolutionWidth(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.ResolutionWidth = &v
	return s
}

type DescribeDisplayConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisplayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisplayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisplayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigResponse) SetHeaders(v map[string]*string) *DescribeDisplayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisplayConfigResponse) SetStatusCode(v int32) *DescribeDisplayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisplayConfigResponse) SetBody(v *DescribeDisplayConfigResponseBody) *DescribeDisplayConfigResponse {
	s.Body = v
	return s
}

type DescribeImageListRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// System
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
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
	// The images.
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
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
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

func (s *DescribeImageListResponseBodyData) SetImageVersion(v string) *DescribeImageListResponseBodyData {
	s.ImageVersion = &v
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
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the execution. You can retrieve the output of a command once by using either the execution ID or the cloud phone instance ID.
	//
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
	// The objects that are returned.
	Data []*DescribeInvocationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 440D7342-5E7C-B2DB-D0B4EAC2BDF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	// The end time of the command execution.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the cloud phone instance on which the command is executed.
	//
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// t-15775dc8****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// The execution state of the command.
	//
	// Valid values:
	//
	// 	- Failed: The execution of the command failed.
	//
	// 	- Timeout: The execution of the command timed out.
	//
	// 	- Running: The command is being executed.
	//
	// 	- Success: The execution of the command is successful.
	//
	// 	- Pending: The command is waiting to be executed.
	//
	// example:
	//
	// RUNNING
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The output of the command execution.
	//
	// example:
	//
	// success
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The start time of the command execution.
	//
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

type DescribeMetricLastRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2019-01-31 11:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	MetricNames []*string `json:"MetricNames,omitempty" xml:"MetricNames,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 2019-01-31 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricLastRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastRequest) SetAndroidInstanceIds(v []*string) *DescribeMetricLastRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeMetricLastRequest) SetEndTime(v string) *DescribeMetricLastRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricLastRequest) SetLength(v string) *DescribeMetricLastRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricLastRequest) SetMetricNames(v []*string) *DescribeMetricLastRequest {
	s.MetricNames = v
	return s
}

func (s *DescribeMetricLastRequest) SetNextToken(v string) *DescribeMetricLastRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastRequest) SetPeriod(v int32) *DescribeMetricLastRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricLastRequest) SetStartTime(v string) *DescribeMetricLastRequest {
	s.StartTime = &v
	return s
}

type DescribeMetricLastResponseBody struct {
	// example:
	//
	// 100
	Count            *int32                                            `json:"Count,omitempty" xml:"Count,omitempty"`
	MetricTotalModel []*DescribeMetricLastResponseBodyMetricTotalModel `json:"MetricTotalModel,omitempty" xml:"MetricTotalModel,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2B9E6946-0E2A-5D2B-B275-361DF81F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetricLastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBody) SetCount(v int32) *DescribeMetricLastResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetMetricTotalModel(v []*DescribeMetricLastResponseBodyMetricTotalModel) *DescribeMetricLastResponseBody {
	s.MetricTotalModel = v
	return s
}

func (s *DescribeMetricLastResponseBody) SetNextToken(v string) *DescribeMetricLastResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetRequestId(v string) *DescribeMetricLastResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMetricLastResponseBodyMetricTotalModel struct {
	// example:
	//
	// acp-fkuit0cmyru4p****
	AndroidInstanceId *string                                                          `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	MetricModelList   []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelList `json:"MetricModelList,omitempty" xml:"MetricModelList,omitempty" type:"Repeated"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModel) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) SetAndroidInstanceId(v string) *DescribeMetricLastResponseBodyMetricTotalModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) SetMetricModelList(v []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) *DescribeMetricLastResponseBodyMetricTotalModel {
	s.MetricModelList = v
	return s
}

type DescribeMetricLastResponseBodyMetricTotalModelMetricModelList struct {
	DataPoints []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// example:
	//
	// cpu_utilization
	MetricName       *string                                                                          `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	ProcessLastInfos []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos `json:"ProcessLastInfos,omitempty" xml:"ProcessLastInfos,omitempty" type:"Repeated"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) SetDataPoints(v []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	s.DataPoints = v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) SetMetricName(v string) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) SetProcessLastInfos(v []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	s.ProcessLastInfos = v
	return s
}

type DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints struct {
	// example:
	//
	// 99.52
	Average *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	// example:
	//
	// 100
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// example:
	//
	// 93.1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// example:
	//
	// 1548777660000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetAverage(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Average = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetMaximum(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Maximum = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetMinimum(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Minimum = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetTimestamp(v int64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Timestamp = &v
	return s
}

type DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos struct {
	// example:
	//
	// 50
	CpuUsage *float64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// example:
	//
	// 50
	MemoryUsage *float64 `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// example:
	//
	// com.offerup
	Name       *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessIds []*int32 `json:"ProcessIds,omitempty" xml:"ProcessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1548777660000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetCpuUsage(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.CpuUsage = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetMemoryUsage(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.MemoryUsage = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetName(v string) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.Name = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetProcessIds(v []*int32) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.ProcessIds = v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetTimestamp(v int64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.Timestamp = &v
	return s
}

type DescribeMetricLastResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricLastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricLastResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponse) SetHeaders(v map[string]*string) *DescribeMetricLastResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricLastResponse) SetStatusCode(v int32) *DescribeMetricLastResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricLastResponse) SetBody(v *DescribeMetricLastResponseBody) *DescribeMetricLastResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The display language of the console. Valid values:
	//
	// 	- cn: Simplified Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// en
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- Instance: the instance group mode. [Default]
	//
	// 	- Node: the matrix mode. [Whitelist required]
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
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
	// The request ID.
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
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
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// example:
	//
	// cpm.gn6.gx1
	MatrixSpec *string `json:"MatrixSpec,omitempty" xml:"MatrixSpec,omitempty"`
	// The maximum number of items to return per page in a paginated query. The value range is 1 to 100, with a default value of 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates the starting position for reading. If left empty, it starts from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The purchase mode of cloud mobile phones.
	//
	// Valid values:
	//
	// 	- Instance (default): the instance group mode.
	//
	// 	- Node: the matrix mode [whitelisted].
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// List of specification IDs.
	SpecIds []*string `json:"SpecIds,omitempty" xml:"SpecIds,omitempty" type:"Repeated"`
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
	// The specifications.
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
	// The maximum number of cloud phone instances.
	//
	// example:
	//
	// 40
	MaxPhoneCount *string `json:"MaxPhoneCount,omitempty" xml:"MaxPhoneCount,omitempty"`
	// Memory size.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum number of cloud phone instances.
	//
	// example:
	//
	// 4
	MinPhoneCount *string `json:"MinPhoneCount,omitempty" xml:"MinPhoneCount,omitempty"`
	// example:
	//
	// 2
	PhoneCount *string `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	// example:
	//
	// 1920*1080
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// Specification ID.
	//
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

func (s *DescribeSpecResponseBodySpecInfoModel) SetMaxPhoneCount(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.MaxPhoneCount = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetMemory(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.Memory = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetMinPhoneCount(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.MinPhoneCount = &v
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
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the cloud phone instance.
	//
	// example:
	//
	// defaultInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the command execution. You can set the value to the last returned request ID.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The level of the task. A value of 1 specifies a batch task. A value of 2 specifies an instance-level task.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
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
	// FFbc8N4E1iOlcSxC+8boa0HHH2LKWbggYUinyrZWvtS1oTrMYCg1HuMLGuftj0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The extension field.
	//
	// example:
	//
	// param
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the parent task.
	//
	// example:
	//
	// t-iaej5dkbnmivx****
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// The IDs of the resources.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The IDs of the tasks.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The state of the task.
	//
	// Valid values:
	//
	// 	- PartFinished: The task is partially successful.
	//
	// 	- Finished: The task is completed.
	//
	// 	- Failed: The task failed.
	//
	// 	- Skipped: The task is skipped.
	//
	// 	- Processing: The task is running.
	//
	// 	- Waiting: The task is in queue.
	//
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The status of the tasks.
	TaskStatuses []*string `json:"TaskStatuses,omitempty" xml:"TaskStatuses,omitempty" type:"Repeated"`
	// The type of the task.
	//
	// Valid values:
	//
	// 	- BackupFile: backs up files.
	//
	// 	- StopInstance: stops cloud phone instances.
	//
	// 	- RebootInstance: restarts cloud phone instances.
	//
	// 	- StartApp: starts apps.
	//
	// 	- SendFile: uploads files.
	//
	// 	- RunCommand: sends remote command.
	//
	// 	- RestartApp: restarts apps.
	//
	// 	- ResetInstance: resets cloud phone instances.
	//
	// 	- RecoverFile: recovers files.
	//
	// 	- UninstallApp: uninstalls apps.
	//
	// 	- StopApp: stops apps.
	//
	// 	- Screenshot: takes screenshots.
	//
	// 	- InstallApp: installs apps.
	//
	// 	- FetchFile: downloads files.
	//
	// 	- UpdateGroupImage: replaces images.
	//
	// 	- StartInstance: starts instances.
	//
	// example:
	//
	// StartInstance
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The types of the tasks.
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
	// The objects that are returned.
	Data []*DescribeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	// The error code.
	//
	// example:
	//
	// SendFileFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// connect error.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The total number of failed subtasks.
	//
	// example:
	//
	// 2
	FailedChildCount *int32 `json:"FailedChildCount,omitempty" xml:"FailedChildCount,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2022-10-11T08:53:32Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the cloud phone instance.
	//
	// example:
	//
	// defaultInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The state of the cloud phone instance.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The ID of the command execution.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The level of the task.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The operator.
	//
	// example:
	//
	// test
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameters of the task.
	//
	// example:
	//
	// param
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the parent task.
	//
	// example:
	//
	// t-41oan3tza16vs****
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// acp-25nt4kk9whhok****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The execution result of the task.
	//
	// example:
	//
	// {\\"Success\\": True}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The total number of the subtasks that are running.
	//
	// example:
	//
	// 0
	RunningChildCount *int32 `json:"RunningChildCount,omitempty" xml:"RunningChildCount,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2022-10-11T08:53:32Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of successfully executed subtasks.
	//
	// example:
	//
	// 98
	SuccessChildCount *int32 `json:"SuccessChildCount,omitempty" xml:"SuccessChildCount,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The state of the task.
	//
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// StartInstance
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of subtasks of the current batch task.
	//
	// example:
	//
	// 100
	TotalChildCount *int32 `json:"TotalChildCount,omitempty" xml:"TotalChildCount,omitempty"`
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
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the ADB key pair.
	//
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
	// The object that is returned.
	Data *DetachKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
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
	// The IDs of the cloud phone instances from which the ADB key pair is successfully detached.
	DetachedInstanceIds []*string `json:"DetachedInstanceIds,omitempty" xml:"DetachedInstanceIds,omitempty" type:"Repeated"`
	// The number of the cloud phone instances from which the ADB key pair failed to be detached.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The total number of the cloud phone instances.
	//
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

type DisconnectAndroidInstanceRequest struct {
	EndUserId   *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DisconnectAndroidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DisconnectAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *DisconnectAndroidInstanceRequest) SetEndUserId(v string) *DisconnectAndroidInstanceRequest {
	s.EndUserId = &v
	return s
}

func (s *DisconnectAndroidInstanceRequest) SetInstanceIds(v []*string) *DisconnectAndroidInstanceRequest {
	s.InstanceIds = v
	return s
}

type DisconnectAndroidInstanceResponseBody struct {
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisconnectAndroidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisconnectAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DisconnectAndroidInstanceResponseBody) SetRequestId(v string) *DisconnectAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DisconnectAndroidInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisconnectAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisconnectAndroidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisconnectAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *DisconnectAndroidInstanceResponse) SetHeaders(v map[string]*string) *DisconnectAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *DisconnectAndroidInstanceResponse) SetStatusCode(v int32) *DisconnectAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisconnectAndroidInstanceResponse) SetBody(v *DisconnectAndroidInstanceResponseBody) *DisconnectAndroidInstanceResponse {
	s.Body = v
	return s
}

type DistributeImageRequest struct {
	// The regions to which you want to distribute an image.
	//
	// This parameter is required.
	DistributeRegionList []*string `json:"DistributeRegionList,omitempty" xml:"DistributeRegionList,omitempty" type:"Repeated"`
	// The ID of the image that you want to distribute.
	//
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
	// The ID of the request.
	//
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
	// The IDs of the cloud phone instances that you want to delete.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to enable the auto-payment feature. Default value: false.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Ensure your account has sufficient balance to use this feature.
	//
	// 	- false: disables the auto-payment feature. This requires manual payment each time you place an order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the instance group.
	//
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
	// The ID of the order.
	//
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
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

type EndCoordinationRequest struct {
	// example:
	//
	// lina
	CoordinatorUserId *string `json:"CoordinatorUserId,omitempty" xml:"CoordinatorUserId,omitempty"`
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// xiaoming
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
}

func (s EndCoordinationRequest) String() string {
	return tea.Prettify(s)
}

func (s EndCoordinationRequest) GoString() string {
	return s.String()
}

func (s *EndCoordinationRequest) SetCoordinatorUserId(v string) *EndCoordinationRequest {
	s.CoordinatorUserId = &v
	return s
}

func (s *EndCoordinationRequest) SetInstanceId(v string) *EndCoordinationRequest {
	s.InstanceId = &v
	return s
}

func (s *EndCoordinationRequest) SetOwnerUserId(v string) *EndCoordinationRequest {
	s.OwnerUserId = &v
	return s
}

type EndCoordinationResponseBody struct {
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndCoordinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndCoordinationResponseBody) GoString() string {
	return s.String()
}

func (s *EndCoordinationResponseBody) SetRequestId(v string) *EndCoordinationResponseBody {
	s.RequestId = &v
	return s
}

type EndCoordinationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EndCoordinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndCoordinationResponse) String() string {
	return tea.Prettify(s)
}

func (s EndCoordinationResponse) GoString() string {
	return s.String()
}

func (s *EndCoordinationResponse) SetHeaders(v map[string]*string) *EndCoordinationResponse {
	s.Headers = v
	return s
}

func (s *EndCoordinationResponse) SetStatusCode(v int32) *EndCoordinationResponse {
	s.StatusCode = &v
	return s
}

func (s *EndCoordinationResponse) SetBody(v *EndCoordinationResponseBody) *EndCoordinationResponse {
	s.Body = v
	return s
}

type ExpandDataVolumeRequest struct {
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string   `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NodeIds     []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	ShareDataVolume *int32 `json:"ShareDataVolume,omitempty" xml:"ShareDataVolume,omitempty"`
}

func (s ExpandDataVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpandDataVolumeRequest) GoString() string {
	return s.String()
}

func (s *ExpandDataVolumeRequest) SetAutoPay(v bool) *ExpandDataVolumeRequest {
	s.AutoPay = &v
	return s
}

func (s *ExpandDataVolumeRequest) SetBizRegionId(v string) *ExpandDataVolumeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ExpandDataVolumeRequest) SetNodeIds(v []*string) *ExpandDataVolumeRequest {
	s.NodeIds = v
	return s
}

func (s *ExpandDataVolumeRequest) SetShareDataVolume(v int32) *ExpandDataVolumeRequest {
	s.ShareDataVolume = &v
	return s
}

type ExpandDataVolumeResponseBody struct {
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExpandDataVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExpandDataVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ExpandDataVolumeResponseBody) SetOrderId(v string) *ExpandDataVolumeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ExpandDataVolumeResponseBody) SetRequestId(v string) *ExpandDataVolumeResponseBody {
	s.RequestId = &v
	return s
}

type ExpandDataVolumeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExpandDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandDataVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpandDataVolumeResponse) GoString() string {
	return s.String()
}

func (s *ExpandDataVolumeResponse) SetHeaders(v map[string]*string) *ExpandDataVolumeResponse {
	s.Headers = v
	return s
}

func (s *ExpandDataVolumeResponse) SetStatusCode(v int32) *ExpandDataVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpandDataVolumeResponse) SetBody(v *ExpandDataVolumeResponseBody) *ExpandDataVolumeResponse {
	s.Body = v
	return s
}

type FetchFileRequest struct {
	// The IDs of the cloud phone instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The path to the file that you want to pull from the cloud phone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data/a.txt
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The endpoint of the OSS bucket in which you want to store the pulled file.
	//
	// >  Set the value to an internal endpoint when the cloud phone instance and the OSS bucket are in the same region to improve upload speed without incurring public traffic fees. Sample endpoint: `oss-cn-hangzhou-internal.aliyuncs.com`. For more information, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the storage service.
	//
	// >  Currently, only OSS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// The OSS URL of the pulled file.
	//
	// >  The OSS bucket name must start with "cloudphone-saved-bucket-", for example, "cloudphone-saved-bucket-example". You must also create an OSS directory to store the backup data. Set the value for UploadUrl in this format: oss://\\<BucketName>/\\<OSSDirectoryName>.
	//
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
	// The objects that are returned.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	Data []*FetchFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request. If the request fails, share this ID with technical support to help diagnose the issue.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// t-ehs0yoedj0xe9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

type GenerateCoordinationCodeRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// acp-2zecay9ponatdc4m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the user to whom the current instance is assigned.
	//
	// example:
	//
	// xiaoming
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
}

func (s GenerateCoordinationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCoordinationCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeRequest) SetInstanceId(v string) *GenerateCoordinationCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateCoordinationCodeRequest) SetOwnerUserId(v string) *GenerateCoordinationCodeRequest {
	s.OwnerUserId = &v
	return s
}

type GenerateCoordinationCodeResponseBody struct {
	// The collaboration code.
	//
	// example:
	//
	// CSHGDK
	CoordinatorCode *string `json:"CoordinatorCode,omitempty" xml:"CoordinatorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCoordinationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCoordinationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponseBody) SetCoordinatorCode(v string) *GenerateCoordinationCodeResponseBody {
	s.CoordinatorCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponseBody) SetRequestId(v string) *GenerateCoordinationCodeResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCoordinationCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCoordinationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCoordinationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCoordinationCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateCoordinationCodeResponse) SetHeaders(v map[string]*string) *GenerateCoordinationCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateCoordinationCodeResponse) SetStatusCode(v int32) *GenerateCoordinationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCoordinationCodeResponse) SetBody(v *GenerateCoordinationCodeResponseBody) *GenerateCoordinationCodeResponse {
	s.Body = v
	return s
}

type ImportKeyPairRequest struct {
	// The name of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key of the ADB key pair.
	//
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
	// The object that is returned.
	Data *ImportKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
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
	// The time when the ADB key pair was created.
	//
	// example:
	//
	// 2023-03-05T10:29:22Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3*****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the ADB key pair.
	//
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
	// The IDs of the apps that you want to install.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The IDs of the instance groups.
	InstanceGroupIdList []*string `json:"InstanceGroupIdList,omitempty" xml:"InstanceGroupIdList,omitempty" type:"Repeated"`
	// The IDs of the cloud phone instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
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
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-14xwibw7yp73q****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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

func (s *ListPolicyGroupsRequest) SetPolicyType(v string) *ListPolicyGroupsRequest {
	s.PolicyType = &v
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
	PolicyGroupName        *string                                                             `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	PolicyRelatedResources *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources `json:"PolicyRelatedResources,omitempty" xml:"PolicyRelatedResources,omitempty" type:"Struct"`
	// The height of the resolution.
	//
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	// The width of the resolution.
	//
	// example:
	//
	// 1920
	SessionResolutionWidth *int32                                                 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	Watermark              *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
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

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetPolicyRelatedResources(v *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.PolicyRelatedResources = v
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

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetWatermark(v *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.Watermark = v
	return s
}

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy struct {
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

type ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources struct {
	AndroidInstanceGroupIds []*string `json:"AndroidInstanceGroupIds,omitempty" xml:"AndroidInstanceGroupIds,omitempty" type:"Repeated"`
	CloudPhoneMatrixIds     []*string `json:"CloudPhoneMatrixIds,omitempty" xml:"CloudPhoneMatrixIds,omitempty" type:"Repeated"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) SetAndroidInstanceGroupIds(v []*string) *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources {
	s.AndroidInstanceGroupIds = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) SetCloudPhoneMatrixIds(v []*string) *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources {
	s.CloudPhoneMatrixIds = v
	return s
}

type ListPolicyGroupsResponseBodyPolicyGroupModelWatermark struct {
	WatermarkColor             *int32    `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	WatermarkCustomText        *string   `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkFontSize          *int32    `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	WatermarkSwitch            *string   `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	WatermarkTransparencyValue *int32    `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	WatermarkTypes             []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkColor(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkColor = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkCustomText(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkCustomText = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkFontSize(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkFontSize = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkSwitch(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkSwitch = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkTransparencyValue(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkTypes(v []*string) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkTypes = v
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

type ModifyCloudPhoneNodeRequest struct {
	// The name that you want to assign to the cloud phone matrix.
	//
	// example:
	//
	// node_name_new
	NewNodeName *string `json:"NewNodeName,omitempty" xml:"NewNodeName,omitempty"`
	// The ID of the cloud phone matrix.
	//
	// example:
	//
	// cpn-0ugbptfu473fy****
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StreamMode *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
}

func (s ModifyCloudPhoneNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudPhoneNodeRequest) SetNewNodeName(v string) *ModifyCloudPhoneNodeRequest {
	s.NewNodeName = &v
	return s
}

func (s *ModifyCloudPhoneNodeRequest) SetNodeId(v string) *ModifyCloudPhoneNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyCloudPhoneNodeRequest) SetStreamMode(v int32) *ModifyCloudPhoneNodeRequest {
	s.StreamMode = &v
	return s
}

type ModifyCloudPhoneNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudPhoneNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCloudPhoneNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudPhoneNodeResponseBody) SetRequestId(v string) *ModifyCloudPhoneNodeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCloudPhoneNodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudPhoneNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudPhoneNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCloudPhoneNodeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudPhoneNodeResponse) SetHeaders(v map[string]*string) *ModifyCloudPhoneNodeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudPhoneNodeResponse) SetStatusCode(v int32) *ModifyCloudPhoneNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudPhoneNodeResponse) SetBody(v *ModifyCloudPhoneNodeResponseBody) *ModifyCloudPhoneNodeResponse {
	s.Body = v
	return s
}

type ModifyDisplayConfigRequest struct {
	AndroidInstanceIds []*string                                `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	DisplayConfig      *ModifyDisplayConfigRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
}

func (s ModifyDisplayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDisplayConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigRequest) SetAndroidInstanceIds(v []*string) *ModifyDisplayConfigRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ModifyDisplayConfigRequest) SetDisplayConfig(v *ModifyDisplayConfigRequestDisplayConfig) *ModifyDisplayConfigRequest {
	s.DisplayConfig = v
	return s
}

type ModifyDisplayConfigRequestDisplayConfig struct {
	// example:
	//
	// 240
	Dpi *int32 `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	// example:
	//
	// null
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// example:
	//
	// 1920
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s ModifyDisplayConfigRequestDisplayConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyDisplayConfigRequestDisplayConfig) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetDpi(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.Dpi = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetFps(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.Fps = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetLockResolution(v string) *ModifyDisplayConfigRequestDisplayConfig {
	s.LockResolution = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetResolutionHeight(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyDisplayConfigRequestDisplayConfig) SetResolutionWidth(v int32) *ModifyDisplayConfigRequestDisplayConfig {
	s.ResolutionWidth = &v
	return s
}

type ModifyDisplayConfigShrinkRequest struct {
	AndroidInstanceIds  []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	DisplayConfigShrink *string   `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
}

func (s ModifyDisplayConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDisplayConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigShrinkRequest) SetAndroidInstanceIds(v []*string) *ModifyDisplayConfigShrinkRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ModifyDisplayConfigShrinkRequest) SetDisplayConfigShrink(v string) *ModifyDisplayConfigShrinkRequest {
	s.DisplayConfigShrink = &v
	return s
}

type ModifyDisplayConfigResponseBody struct {
	// example:
	//
	// A578AD3A-8E7C-54FE-A09F-B060941*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDisplayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDisplayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigResponseBody) SetRequestId(v string) *ModifyDisplayConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDisplayConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDisplayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDisplayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDisplayConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigResponse) SetHeaders(v map[string]*string) *ModifyDisplayConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDisplayConfigResponse) SetStatusCode(v int32) *ModifyDisplayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDisplayConfigResponse) SetBody(v *ModifyDisplayConfigResponseBody) *ModifyDisplayConfigResponse {
	s.Body = v
	return s
}

type ModifyInstanceChargeTypeRequest struct {
	// Specifies whether to enable the auto-payment feature. Default value: false.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature. Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing method. Valid values:
	//
	// >  Currently, this operation only allows you to change the billing method from **pay-as-you-go to subscription**.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the instance groups.
	//
	// This parameter is required.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The subscription duration. The unit is specified by PeriodUnit. Valid values: 1 Month, 2 Months, 3 Months, 6 Months, and 1 Year.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Month**
	//
	// 	- **Year**
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s ModifyInstanceChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeRequest) SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetAutoRenew(v bool) *ModifyInstanceChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetChargeType(v string) *ModifyInstanceChargeTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceGroupIds(v []*string) *ModifyInstanceChargeTypeRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriod(v int32) *ModifyInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

type ModifyInstanceChargeTypeResponseBody struct {
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponseBody) SetInstanceGroupIds(v []*string) *ModifyInstanceChargeTypeResponseBody {
	s.InstanceGroupIds = v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetOrderId(v string) *ModifyInstanceChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetRequestId(v string) *ModifyInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceChargeTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceChargeTypeResponse) SetStatusCode(v int32) *ModifyInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponse) SetBody(v *ModifyInstanceChargeTypeResponseBody) *ModifyInstanceChargeTypeResponse {
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
	ResolutionWidth *int32                             `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	Watermark       *ModifyPolicyGroupRequestWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
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

func (s *ModifyPolicyGroupRequest) SetWatermark(v *ModifyPolicyGroupRequestWatermark) *ModifyPolicyGroupRequest {
	s.Watermark = v
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
	ProxyUserName *string                                           `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	Rules         []*ModifyPolicyGroupRequestNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
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

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetRules(v []*ModifyPolicyGroupRequestNetRedirectPolicyRules) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.Rules = v
	return s
}

type ModifyPolicyGroupRequestNetRedirectPolicyRules struct {
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Target   *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyPolicyGroupRequestNetRedirectPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestNetRedirectPolicyRules) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) SetRuleType(v string) *ModifyPolicyGroupRequestNetRedirectPolicyRules {
	s.RuleType = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) SetTarget(v string) *ModifyPolicyGroupRequestNetRedirectPolicyRules {
	s.Target = &v
	return s
}

type ModifyPolicyGroupRequestWatermark struct {
	WatermarkColor             *int32    `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	WatermarkCustomText        *string   `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkFontSize          *int32    `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	WatermarkSwitch            *string   `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	WatermarkTransparencyValue *int32    `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	WatermarkTypes             []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s ModifyPolicyGroupRequestWatermark) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestWatermark) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkColor(v int32) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkColor = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkCustomText(v string) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkFontSize(v int32) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkFontSize = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkSwitch(v string) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkSwitch = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkTransparencyValue(v int32) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkTypes(v []*string) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkTypes = v
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
	ResolutionWidth *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
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

func (s *ModifyPolicyGroupShrinkRequest) SetWatermarkShrink(v string) *ModifyPolicyGroupShrinkRequest {
	s.WatermarkShrink = &v
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
	// The ID of the app.
	//
	// example:
	//
	// 1234
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IDs of the cloud phone instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The type of the operation.
	//
	// Valid values:
	//
	// 	- stop: closes the app.
	//
	// 	- restart: reopens the app.
	//
	// 	- start: open the app.
	//
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
	// The ID of the request.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
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
	// The IDs of the cloud phone instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to enforce a restart operation. If a cloud phone instance fails to stop due to system or network issues, a forced restart can be triggered, though it may result in data loss.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ForceStop *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	SaleMode  *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
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

func (s *RebootAndroidInstancesInGroupRequest) SetSaleMode(v string) *RebootAndroidInstancesInGroupRequest {
	s.SaleMode = &v
	return s
}

type RebootAndroidInstancesInGroupResponseBody struct {
	// The ID of the request.
	//
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
	// Specifies whether to back up the whole instance.
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
	// The number of restored instances.
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
	// The ID of the batch task.
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
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Ensure your account has sufficient balance to use this feature.
	//
	// 	- false: disables the auto-payment feature. You need to manually complete the payment process.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The duration of the renewal, measured in units defined by PeriodUnit.
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration. Default value: Month.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
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
	// The ID of the order.
	//
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
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

type RenewCloudPhoneNodesRequest struct {
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-renewal feature. In this case, the system automatically renews the instance upon expiration.
	//
	// 	- false (default): disables the auto-renewal feature. In this case, you need to manually renew the instance upon expiration.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cloud phone matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The subscription duration. The unit is specified by `PeriodUnit`. Valid values:
	//
	// 	- When `PeriodUnit` is set to **year**: 1.
	//
	// 	- When `PeriodUnit` is set to **month**: 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s RenewCloudPhoneNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesRequest) SetAutoPay(v bool) *RenewCloudPhoneNodesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetAutoRenew(v bool) *RenewCloudPhoneNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetNodeIds(v []*string) *RenewCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPeriod(v int32) *RenewCloudPhoneNodesRequest {
	s.Period = &v
	return s
}

func (s *RenewCloudPhoneNodesRequest) SetPeriodUnit(v string) *RenewCloudPhoneNodesRequest {
	s.PeriodUnit = &v
	return s
}

type RenewCloudPhoneNodesResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 22365781890****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewCloudPhoneNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesResponseBody) SetOrderId(v string) *RenewCloudPhoneNodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewCloudPhoneNodesResponseBody) SetRequestId(v string) *RenewCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

type RenewCloudPhoneNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewCloudPhoneNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *RenewCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *RenewCloudPhoneNodesResponse) SetStatusCode(v int32) *RenewCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewCloudPhoneNodesResponse) SetBody(v *RenewCloudPhoneNodesResponseBody) *RenewCloudPhoneNodesResponse {
	s.Body = v
	return s
}

type ResetAndroidInstancesInGroupRequest struct {
	// The IDs of the cloud phone instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	SaleMode           *string   `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	SettingResetType   *int32    `json:"SettingResetType,omitempty" xml:"SettingResetType,omitempty"`
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

func (s *ResetAndroidInstancesInGroupRequest) SetSaleMode(v string) *ResetAndroidInstancesInGroupRequest {
	s.SaleMode = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetSettingResetType(v int32) *ResetAndroidInstancesInGroupRequest {
	s.SettingResetType = &v
	return s
}

type ResetAndroidInstancesInGroupResponseBody struct {
	// The ID of the request.
	//
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
	// The content of the command.
	//
	// example:
	//
	// ls
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding method of the command content (`CommandContent`). The value is not case-sensitive.
	//
	// >  If you set the value to an invalid encoding method, the system will process the command content as `PlainText`.
	//
	// Valid values:
	//
	// 	- Base64: encodes the command content in Base64.
	//
	// 	- PlainText (default): does not encode the command content. The command content is input as plain text.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The timeout period of the command execution. If the command execution exceeds the timeout period, it will be considered timed out. If you leave this parameter empty, it defaults to 60.
	//
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
	// The ID of the command execution. You can use the command execution ID to query the output of a command.
	//
	// example:
	//
	// t-gov2ujrk32v4****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The ID of the request.
	//
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
	// The IDs of the cloud phone instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The path to which you want to upload the pushed file in the cloud phone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The name of the file uploaded from the Object Storage Service (OSS) to the cloud phone instance.
	//
	// >  If UploadType is set to OSS, you must specify TargetFileName. If TargetFileName is empty, the file uploaded from the OSS bucket to the cloud phone instance retains its original name. If TargetFileName is provided with a value, the uploaded file in the SourceFilePath directory uses the specified name (TargetFileName). If UploadType is set to DOWNLOAD_URL, TargetFileName does not take effect.
	//
	// example:
	//
	// test.txt
	TargetFileName *string `json:"TargetFileName,omitempty" xml:"TargetFileName,omitempty"`
	// The endpoint of the OSS bucket in which the file is stored.
	//
	// >  Set the value to an internal endpoint when the cloud phone instance and the OSS bucket are in the same region to improve transfer speed without incurring public traffic fees. Sample endpoint: `oss-cn-hangzhou-internal.aliyuncs.com`. For more information, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The storage type of the file that you want to upload.
	//
	// 	- Set the value to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// The OSS URL of the file.
	//
	// >  The OSS bucket name must start with "cloudphone-saved-bucket-", for example, "cloudphone-saved-bucket-example". You must also create an OSS directory to store the backup data. Set the value for UploadUrl in this format: oss://\\<BucketName>/\\<OSSDirectoryName>\\<FileName>.
	//
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

func (s *SendFileRequest) SetTargetFileName(v string) *SendFileRequest {
	s.TargetFileName = &v
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
	// The objects that are returned.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	Data []*SendFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request. If the request fails, share this ID with technical support to help diagnose the issue.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// t-ehs0yoedj0xe9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The status of the ADB authentication feature.
	//
	// Valid values:
	//
	// 	- 0: The ADB authentication feature is disabled.
	//
	// 	- 1: The ADB authentication feature is enabled.
	//
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
	// The returned object.
	Data *SetAdbSecureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
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
	// The number of the cloud phone instances for which the ADB authentication feature failed to be enabled.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The IDs of the cloud phone instances for which the ADB authentication feature is enabled.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The total number of the cloud phone instances.
	//
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
	SaleMode           *string   `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
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

func (s *StartAndroidInstanceRequest) SetSaleMode(v string) *StartAndroidInstanceRequest {
	s.SaleMode = &v
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
	// The IDs of the cloud phone instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to enforce a stop operation. If a cloud phone instance fails to stop due to system or network issues, a forced stop can be triggered, though it may result in data loss.
	//
	// example:
	//
	// false
	ForceStop *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	SaleMode  *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
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

func (s *StopAndroidInstanceRequest) SetSaleMode(v string) *StopAndroidInstanceRequest {
	s.SaleMode = &v
	return s
}

type StopAndroidInstanceResponseBody struct {
	// The ID of the request.
	//
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
	// The IDs of the apps.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The ID of the instance groups. If you specify this parameter, you cannot specify InstanceIdList.
	InstanceGroupIdList []*string `json:"InstanceGroupIdList,omitempty" xml:"InstanceGroupIdList,omitempty" type:"Repeated"`
	// The IDs of the cloud phone instances. If you specify this parameter, you cannot specify InstanceGroupIdList.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
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
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-075cllfeuazh****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The IDs of the instance groups.
	//
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
	// The ID of the request.
	//
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

type UpdateInstanceImageRequest struct {
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId        *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s UpdateInstanceImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageRequest) SetImageId(v string) *UpdateInstanceImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceImageRequest) SetInstanceIdList(v []*string) *UpdateInstanceImageRequest {
	s.InstanceIdList = v
	return s
}

type UpdateInstanceImageResponseBody struct {
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateInstanceImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponseBody) SetRequestId(v string) *UpdateInstanceImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceImageResponseBody) SetTaskId(v string) *UpdateInstanceImageResponseBody {
	s.TaskId = &v
	return s
}

type UpdateInstanceImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponse) SetHeaders(v map[string]*string) *UpdateInstanceImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceImageResponse) SetStatusCode(v int32) *UpdateInstanceImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceImageResponse) SetBody(v *UpdateInstanceImageResponseBody) *UpdateInstanceImageResponse {
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
// Attaches an Android Debug Bridge (ADB) key pair to one or more cloud phone instances.
//
// Description:
//
//	  You can attach to an ADB key pair only to cloud phone instances in the Running state.
//
//		- After you attach an ADB key pair, make sure the private key of the ADB key pair is copied to the ~/.android directory (macOS or Linux operating systems) or the C:\\Users\\Username.android directory (Windows operating systems). In addition, you must run the adb kill-server command to restart the ADB process to ensure correct ADB connection. Otherwise, ADB connection may fail due to authentication exceptions.
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
	_result = &AttachKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an Android Debug Bridge (ADB) key pair to one or more cloud phone instances.
//
// Description:
//
//	  You can attach to an ADB key pair only to cloud phone instances in the Running state.
//
//		- After you attach an ADB key pair, make sure the private key of the ADB key pair is copied to the ~/.android directory (macOS or Linux operating systems) or the C:\\Users\\Username.android directory (Windows operating systems). In addition, you must run the adb kill-server command to restart the ADB process to ensure correct ADB connection. Otherwise, ADB connection may fail due to authentication exceptions.
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
	_result = &AuthorizeAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	_result = &BackupFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
// Retrieves connection tickets in batch.
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
	_result = &BatchGetAcpConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves connection tickets in batch.
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
// 修改云手机矩阵的配置
//
// @param request - ChangeCloudPhoneNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCloudPhoneNodeResponse
func (client *Client) ChangeCloudPhoneNodeWithOptions(request *ChangeCloudPhoneNodeRequest, runtime *util.RuntimeOptions) (_result *ChangeCloudPhoneNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneCount)) {
		query["PhoneCount"] = request.PhoneCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeCloudPhoneNode"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeCloudPhoneNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改云手机矩阵的配置
//
// @param request - ChangeCloudPhoneNodeRequest
//
// @return ChangeCloudPhoneNodeResponse
func (client *Client) ChangeCloudPhoneNode(request *ChangeCloudPhoneNodeRequest) (_result *ChangeCloudPhoneNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeCloudPhoneNodeResponse{}
	_body, _err := client.ChangeCloudPhoneNodeWithOptions(request, runtime)
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
	_result = &CheckResourceStockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
//   - If the billing method of an instance group is PrePaid, AutoPay is set to false by default. In this case, you need to go to [Expenses and Costs](https://usercenter2-intl.aliyun.com/order/list) to manually complete the payment.
//
//   - You can also set AutoPay to true based on your business requirements.
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
	_result = &CreateAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
//   - If the billing method of an instance group is PrePaid, AutoPay is set to false by default. In this case, you need to go to [Expenses and Costs](https://usercenter2-intl.aliyun.com/order/list) to manually complete the payment.
//
//   - You can also set AutoPay to true based on your business requirements.
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
// When creating an app, you can provide app information to the system in one of the following ways:
//
//   - Way 1: Apps from the Application Center
//
//   - You can use one of the following methods:
//
//   - Method 1: Pass in the `FileName` and `FilePath` parameters at the same time.
//
//   - Method 2: Pass in the `OssAppUrl` parameter
//
//   - Rule: If your app is from the Alibaba Cloud Workspace Application Center, you must use either Method 1 or Method 2. If both are used, Method 1 takes priority.
//
//   - Condition: Before you proceed, log on to the [Elastic Desktop Service (EDS) Enterprise console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the app file to the Application Center to obtain the values of the `FileName`, `FilePath`, and `OssAppUrl` parameters.
//
//   - Way 2: Custom apps
//
//   - Pass in the `CustomAppInfo` parameter.
//
//   - Rule: If you pass in the `CustomAppInfo` parameter, all six fields within it are required.
//
// >  If Way 1 and Way 2 are adopted simultaneously, the information from Way 2 takes priority.
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

	if !tea.BoolValue(util.IsUnset(request.SignApk)) {
		query["SignApk"] = request.SignApk
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
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Android application.
//
// Description:
//
// When creating an app, you can provide app information to the system in one of the following ways:
//
//   - Way 1: Apps from the Application Center
//
//   - You can use one of the following methods:
//
//   - Method 1: Pass in the `FileName` and `FilePath` parameters at the same time.
//
//   - Method 2: Pass in the `OssAppUrl` parameter
//
//   - Rule: If your app is from the Alibaba Cloud Workspace Application Center, you must use either Method 1 or Method 2. If both are used, Method 1 takes priority.
//
//   - Condition: Before you proceed, log on to the [Elastic Desktop Service (EDS) Enterprise console](https://eds.console.aliyun.com/osshelp) and follow the on-screen instructions to upload the app file to the Application Center to obtain the values of the `FileName`, `FilePath`, and `OssAppUrl` parameters.
//
//   - Way 2: Custom apps
//
//   - Pass in the `CustomAppInfo` parameter.
//
//   - Rule: If you pass in the `CustomAppInfo` parameter, all six fields within it are required.
//
// >  If Way 1 and Way 2 are adopted simultaneously, the information from Way 2 takes priority.
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
// Creates a cloud phone matrix.
//
// @param tmpReq - CreateCloudPhoneNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudPhoneNodeResponse
func (client *Client) CreateCloudPhoneNodeWithOptions(tmpReq *CreateCloudPhoneNodeRequest, runtime *util.RuntimeOptions) (_result *CreateCloudPhoneNodeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCloudPhoneNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DisplayConfig)) {
		request.DisplayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayConfig, tea.String("DisplayConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkId)) {
		query["NetworkId"] = request.NetworkId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneCount)) {
		query["PhoneCount"] = request.PhoneCount
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneDataVolume)) {
		query["PhoneDataVolume"] = request.PhoneDataVolume
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionHeight)) {
		query["ResolutionHeight"] = request.ResolutionHeight
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionWidth)) {
		query["ResolutionWidth"] = request.ResolutionWidth
	}

	if !tea.BoolValue(util.IsUnset(request.ServerShareDataVolume)) {
		query["ServerShareDataVolume"] = request.ServerShareDataVolume
	}

	if !tea.BoolValue(util.IsUnset(request.ServerType)) {
		query["ServerType"] = request.ServerType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamMode)) {
		query["StreamMode"] = request.StreamMode
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayConfigShrink)) {
		body["DisplayConfig"] = request.DisplayConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCloudPhoneNode"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCloudPhoneNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud phone matrix.
//
// @param request - CreateCloudPhoneNodeRequest
//
// @return CreateCloudPhoneNodeResponse
func (client *Client) CreateCloudPhoneNode(request *CreateCloudPhoneNodeRequest) (_result *CreateCloudPhoneNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCloudPhoneNodeResponse{}
	_body, _err := client.CreateCloudPhoneNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom image from a cloud phone instance.
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
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image from a cloud phone instance.
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
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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

	if !tea.BoolValue(util.IsUnset(tmpReq.Watermark)) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Watermark, tea.String("Watermark"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		body["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionHeight)) {
		body["ResolutionHeight"] = request.ResolutionHeight
	}

	if !tea.BoolValue(util.IsUnset(request.ResolutionWidth)) {
		body["ResolutionWidth"] = request.ResolutionWidth
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkShrink)) {
		body["Watermark"] = request.WatermarkShrink
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
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// Creates a screenshot of a cloud phone instance.
//
// Description:
//
// You can call this operation to create a screenshot of a cloud phone instance and upload it to the default Object Storage Service (OSS) bucket. The operation returns a task ID, which you can use with the DescribeTasks operation to get the download link for the screenshot.
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
	_result = &CreateScreenshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a screenshot of a cloud phone instance.
//
// Description:
//
// You can call this operation to create a screenshot of a cloud phone instance and upload it to the default Object Storage Service (OSS) bucket. The operation returns a task ID, which you can use with the DescribeTasks operation to get the download link for the screenshot.
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
	_result = &DeleteAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &DeleteAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// 删除备份文件
//
// @param request - DeleteBackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupFileResponse
func (client *Client) DeleteBackupFileWithOptions(request *DeleteBackupFileRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupFileIdList)) {
		query["BackupFileIdList"] = request.BackupFileIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBackupFile"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBackupFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除备份文件
//
// @param request - DeleteBackupFileRequest
//
// @return DeleteBackupFileResponse
func (client *Client) DeleteBackupFile(request *DeleteBackupFileRequest) (_result *DeleteBackupFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupFileResponse{}
	_body, _err := client.DeleteBackupFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a cloud phone matrix.
//
// Description:
//
// Before you proceed, make sure that the cloud phone matrix that you want to delete expired.
//
// @param request - DeleteCloudPhoneNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudPhoneNodesResponse
func (client *Client) DeleteCloudPhoneNodesWithOptions(request *DeleteCloudPhoneNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteCloudPhoneNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		body["NodeIds"] = request.NodeIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCloudPhoneNodes"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCloudPhoneNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cloud phone matrix.
//
// Description:
//
// Before you proceed, make sure that the cloud phone matrix that you want to delete expired.
//
// @param request - DeleteCloudPhoneNodesRequest
//
// @return DeleteCloudPhoneNodesResponse
func (client *Client) DeleteCloudPhoneNodes(request *DeleteCloudPhoneNodesRequest) (_result *DeleteCloudPhoneNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCloudPhoneNodesResponse{}
	_body, _err := client.DeleteCloudPhoneNodesWithOptions(request, runtime)
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
	_result = &DeleteImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
//	  If a cloud phone instance is currently associated with the ADB key pair you intend to delete, the ADB key pair cannot be deleted.
//
//		- Once an ADB key pair is deleted, it cannot be retrieved or queried by using the DescribeKeyPairs operation.
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
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Android Debug Bridge (ADB) key pairs.
//
// Description:
//
//	  If a cloud phone instance is currently associated with the ADB key pair you intend to delete, the ADB key pair cannot be deleted.
//
//		- Once an ADB key pair is deleted, it cannot be retrieved or queried by using the DescribeKeyPairs operation.
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
	_result = &DeletePolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &DescribeAndroidInstanceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// Queries cloud phone instances.
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

	if !tea.BoolValue(util.IsUnset(request.AppManagePolicyId)) {
		query["AppManagePolicyId"] = request.AppManagePolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserId)) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
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

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteIds)) {
		query["OfficeSiteIds"] = request.OfficeSiteIds
	}

	if !tea.BoolValue(util.IsUnset(request.QosRuleIds)) {
		query["QosRuleIds"] = request.QosRuleIds
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
	_result = &DescribeAndroidInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud phone instances.
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

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
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
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &DescribeBackupFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
// Queries the details of a cloud phone matrix.
//
// @param request - DescribeCloudPhoneNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudPhoneNodesResponse
func (client *Client) DescribeCloudPhoneNodesWithOptions(request *DescribeCloudPhoneNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudPhoneNodesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		query["NodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.ServerType)) {
		query["ServerType"] = request.ServerType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudPhoneNodes"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudPhoneNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud phone matrix.
//
// @param request - DescribeCloudPhoneNodesRequest
//
// @return DescribeCloudPhoneNodesResponse
func (client *Client) DescribeCloudPhoneNodes(request *DescribeCloudPhoneNodesRequest) (_result *DescribeCloudPhoneNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudPhoneNodesResponse{}
	_body, _err := client.DescribeCloudPhoneNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询显示设置
//
// @param request - DescribeDisplayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisplayConfigResponse
func (client *Client) DescribeDisplayConfigWithOptions(request *DescribeDisplayConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDisplayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisplayConfig"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisplayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询显示设置
//
// @param request - DescribeDisplayConfigRequest
//
// @return DescribeDisplayConfigResponse
func (client *Client) DescribeDisplayConfig(request *DescribeDisplayConfigRequest) (_result *DescribeDisplayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisplayConfigResponse{}
	_body, _err := client.DescribeDisplayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries images.
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
	_result = &DescribeImageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries images.
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
// Queries the execution results of commands.
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
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution results of commands.
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
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// 查询指定监控项的最新监控数据
//
// @param request - DescribeMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLastWithOptions(request *DescribeMetricLastRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricLastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		body["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.MetricNames)) {
		body["MetricNames"] = request.MetricNames
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMetricLast"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定监控项的最新监控数据
//
// @param request - DescribeMetricLastRequest
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLast(request *DescribeMetricLastRequest) (_result *DescribeMetricLastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.DescribeMetricLastWithOptions(request, runtime)
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
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &DescribeSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// Queries tasks created for a cloud phone instance.
//
// Description:
//
//	  You can call the DescribeTasks operation to query the tasks created for one or more cloud phone instances.
//
//		- The system currently supports various tasks, including starting, stopping, restarting, and resetting cloud phone instances; backing up and restoring data; installing apps; and executing remote commands.
//
//		- You can use the Level field to specify the type of task. If Level is set to 1, it represents a batch task. If Level is set to 2, it represents an instance-level task.
//
// **Example**
//
// Assume you restart two cloud phone instances with the instance IDs acp-25nt4kk9whhok\\*\\*\\*\\	- and acp-j2taq887orj8l\\*\\*\\*\\*, and the returned request ID is B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you want to check the operation outcomes of the two cloud phone instances, you can call the DescribeTasks operation. You need to set the InvokeId request parameter to B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you only want to check the cloud phone instance with the ID acp-25nt4kk9whhok\\*\\*\\*\\*, you must set the ParentTaskId request parameter to the ID of the batch task and the AndroidInstanceId request parameter to acp-25nt4kk9whhok\\*\\*\\*\\	- when calling the DescribeTasks operation.
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
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tasks created for a cloud phone instance.
//
// Description:
//
//	  You can call the DescribeTasks operation to query the tasks created for one or more cloud phone instances.
//
//		- The system currently supports various tasks, including starting, stopping, restarting, and resetting cloud phone instances; backing up and restoring data; installing apps; and executing remote commands.
//
//		- You can use the Level field to specify the type of task. If Level is set to 1, it represents a batch task. If Level is set to 2, it represents an instance-level task.
//
// **Example**
//
// Assume you restart two cloud phone instances with the instance IDs acp-25nt4kk9whhok\\*\\*\\*\\	- and acp-j2taq887orj8l\\*\\*\\*\\*, and the returned request ID is B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you want to check the operation outcomes of the two cloud phone instances, you can call the DescribeTasks operation. You need to set the InvokeId request parameter to B8ED2BA9-0C6A-5643-818F-B5D60A64\\*\\*\\*\\*. If you only want to check the cloud phone instance with the ID acp-25nt4kk9whhok\\*\\*\\*\\*, you must set the ParentTaskId request parameter to the ID of the batch task and the AndroidInstanceId request parameter to acp-25nt4kk9whhok\\*\\*\\*\\	- when calling the DescribeTasks operation.
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
// Detaches an Android Debug Bridge (ADB) key pair from one or more cloud phone instances.
//
// Description:
//
//	After you detach an ADB key pair from a cloud phone instance, the ADB connection will fail. This occurs because the system can no longer authenticate using a valid ADB public key, leading to authentication errors.
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
	_result = &DetachKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches an Android Debug Bridge (ADB) key pair from one or more cloud phone instances.
//
// Description:
//
//	After you detach an ADB key pair from a cloud phone instance, the ADB connection will fail. This occurs because the system can no longer authenticate using a valid ADB public key, leading to authentication errors.
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
// 实例断开连接
//
// @param request - DisconnectAndroidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisconnectAndroidInstanceResponse
func (client *Client) DisconnectAndroidInstanceWithOptions(request *DisconnectAndroidInstanceRequest, runtime *util.RuntimeOptions) (_result *DisconnectAndroidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisconnectAndroidInstance"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisconnectAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例断开连接
//
// @param request - DisconnectAndroidInstanceRequest
//
// @return DisconnectAndroidInstanceResponse
func (client *Client) DisconnectAndroidInstance(request *DisconnectAndroidInstanceRequest) (_result *DisconnectAndroidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisconnectAndroidInstanceResponse{}
	_body, _err := client.DisconnectAndroidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Distributes an image.
//
// Description:
//
// After you distribute an image in supported regions, the distribution cannot be canceled.
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
	_result = &DistributeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Distributes an image.
//
// Description:
//
// After you distribute an image in supported regions, the distribution cannot be canceled.
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
// Downgrades an instance group. Currently, this operation allows you to only delete specific cloud phone instances from an instance group.
//
// Description:
//
// This operation only allows you to scale down an instance group.
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
	_result = &DowngradeAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downgrades an instance group. Currently, this operation allows you to only delete specific cloud phone instances from an instance group.
//
// Description:
//
// This operation only allows you to scale down an instance group.
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
// 结束协同
//
// @param request - EndCoordinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndCoordinationResponse
func (client *Client) EndCoordinationWithOptions(request *EndCoordinationRequest, runtime *util.RuntimeOptions) (_result *EndCoordinationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoordinatorUserId)) {
		query["CoordinatorUserId"] = request.CoordinatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EndCoordination"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EndCoordinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 结束协同
//
// @param request - EndCoordinationRequest
//
// @return EndCoordinationResponse
func (client *Client) EndCoordination(request *EndCoordinationRequest) (_result *EndCoordinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EndCoordinationResponse{}
	_body, _err := client.EndCoordinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 存储扩容
//
// @param request - ExpandDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpandDataVolumeResponse
func (client *Client) ExpandDataVolumeWithOptions(request *ExpandDataVolumeRequest, runtime *util.RuntimeOptions) (_result *ExpandDataVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		query["NodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShareDataVolume)) {
		query["ShareDataVolume"] = request.ShareDataVolume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpandDataVolume"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpandDataVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 存储扩容
//
// @param request - ExpandDataVolumeRequest
//
// @return ExpandDataVolumeResponse
func (client *Client) ExpandDataVolume(request *ExpandDataVolumeRequest) (_result *ExpandDataVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExpandDataVolumeResponse{}
	_body, _err := client.ExpandDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pulls a file from a cloud phone instance and stores it in Object Storage Service (OSS).
//
// Description:
//
// Currently, this operation allows you to retrieve files or folders from cloud phone instances and save them directly to OSS.
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
	_result = &FetchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pulls a file from a cloud phone instance and stores it in Object Storage Service (OSS).
//
// Description:
//
// Currently, this operation allows you to retrieve files or folders from cloud phone instances and save them directly to OSS.
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
// Generates a collaboration code for the cloud phone being accessed by using the current convenience account, and shares this code with other convenience accounts to allow them to access the same cloud phone.
//
// Description:
//
// You can call this operation to generate a collaboration code for a cloud phone accessed by your current account and share this code with other convenience users to allow them to access the same cloud phone over the desktop, mobile, or web client.
//
// @param request - GenerateCoordinationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCoordinationCodeResponse
func (client *Client) GenerateCoordinationCodeWithOptions(request *GenerateCoordinationCodeRequest, runtime *util.RuntimeOptions) (_result *GenerateCoordinationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCoordinationCode"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCoordinationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a collaboration code for the cloud phone being accessed by using the current convenience account, and shares this code with other convenience accounts to allow them to access the same cloud phone.
//
// Description:
//
// You can call this operation to generate a collaboration code for a cloud phone accessed by your current account and share this code with other convenience users to allow them to access the same cloud phone over the desktop, mobile, or web client.
//
// @param request - GenerateCoordinationCodeRequest
//
// @return GenerateCoordinationCodeResponse
func (client *Client) GenerateCoordinationCode(request *GenerateCoordinationCodeRequest) (_result *GenerateCoordinationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateCoordinationCodeResponse{}
	_body, _err := client.GenerateCoordinationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports the public key of an Android Debug Bridge (ADB) key pair.
//
// Description:
//
// To avoid authorization errors that could cause ADB connection failures, you must import the public key of an ADB key pair.
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
	_result = &ImportKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports the public key of an Android Debug Bridge (ADB) key pair.
//
// Description:
//
// To avoid authorization errors that could cause ADB connection failures, you must import the public key of an ADB key pair.
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
// Installs an app on multiple cloud phone instances at the same time.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
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
	_result = &InstallAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an app on multiple cloud phone instances at the same time.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
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

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		body["PolicyType"] = request.PolicyType
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
	_result = &ListPolicyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &ModifyAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &ModifyAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// Modifies a cloud phone matrix. Currently, you can only modify the name of a cloud phone matrix.
//
// @param request - ModifyCloudPhoneNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudPhoneNodeResponse
func (client *Client) ModifyCloudPhoneNodeWithOptions(request *ModifyCloudPhoneNodeRequest, runtime *util.RuntimeOptions) (_result *ModifyCloudPhoneNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewNodeName)) {
		query["NewNodeName"] = request.NewNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamMode)) {
		query["StreamMode"] = request.StreamMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCloudPhoneNode"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCloudPhoneNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a cloud phone matrix. Currently, you can only modify the name of a cloud phone matrix.
//
// @param request - ModifyCloudPhoneNodeRequest
//
// @return ModifyCloudPhoneNodeResponse
func (client *Client) ModifyCloudPhoneNode(request *ModifyCloudPhoneNodeRequest) (_result *ModifyCloudPhoneNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCloudPhoneNodeResponse{}
	_body, _err := client.ModifyCloudPhoneNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改显示设置
//
// @param tmpReq - ModifyDisplayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDisplayConfigResponse
func (client *Client) ModifyDisplayConfigWithOptions(tmpReq *ModifyDisplayConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDisplayConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDisplayConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DisplayConfig)) {
		request.DisplayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayConfig, tea.String("DisplayConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidInstanceIds)) {
		body["AndroidInstanceIds"] = request.AndroidInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayConfigShrink)) {
		body["DisplayConfig"] = request.DisplayConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDisplayConfig"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDisplayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改显示设置
//
// @param request - ModifyDisplayConfigRequest
//
// @return ModifyDisplayConfigResponse
func (client *Client) ModifyDisplayConfig(request *ModifyDisplayConfigRequest) (_result *ModifyDisplayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDisplayConfigResponse{}
	_body, _err := client.ModifyDisplayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the billing method. Currently, this operation only allows you to change the billing method from pay-as-you-go to subscription.
//
// @param request - ModifyInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceChargeTypeResponse
func (client *Client) ModifyInstanceChargeTypeWithOptions(request *ModifyInstanceChargeTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
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
		Action:      tea.String("ModifyInstanceChargeType"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the billing method. Currently, this operation only allows you to change the billing method from pay-as-you-go to subscription.
//
// @param request - ModifyInstanceChargeTypeRequest
//
// @return ModifyInstanceChargeTypeResponse
func (client *Client) ModifyInstanceChargeType(request *ModifyInstanceChargeTypeRequest) (_result *ModifyInstanceChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceChargeTypeResponse{}
	_body, _err := client.ModifyInstanceChargeTypeWithOptions(request, runtime)
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
	_result = &ModifyKeyPairNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

	if !tea.BoolValue(util.IsUnset(tmpReq.Watermark)) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Watermark, tea.String("Watermark"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.WatermarkShrink)) {
		body["Watermark"] = request.WatermarkShrink
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
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// Operates apps in a cloud phone, such as opening, closing, and reopening apps.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
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
	_result = &OperateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Operates apps in a cloud phone, such as opening, closing, and reopening apps.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
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
// Restarts one or more cloud phone instances.
//
// Description:
//
// Before you restart a cloud phone instance, make sure it is in one of the following states: **Available, Abnormal, Backup failure, and Restoration failure**.
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

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
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
	_result = &RebootAndroidInstancesInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts one or more cloud phone instances.
//
// Description:
//
// Before you restart a cloud phone instance, make sure it is in one of the following states: **Available, Abnormal, Backup failure, and Restoration failure**.
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
	_result = &RecoveryFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
// Renews instance groups.
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
	_result = &RenewAndroidInstanceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews instance groups.
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
// Renews a cloud mobile matrix.
//
// @param request - RenewCloudPhoneNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewCloudPhoneNodesResponse
func (client *Client) RenewCloudPhoneNodesWithOptions(request *RenewCloudPhoneNodesRequest, runtime *util.RuntimeOptions) (_result *RenewCloudPhoneNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		body["NodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewCloudPhoneNodes"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewCloudPhoneNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a cloud mobile matrix.
//
// @param request - RenewCloudPhoneNodesRequest
//
// @return RenewCloudPhoneNodesResponse
func (client *Client) RenewCloudPhoneNodes(request *RenewCloudPhoneNodesRequest) (_result *RenewCloudPhoneNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewCloudPhoneNodesResponse{}
	_body, _err := client.RenewCloudPhoneNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets one or more cloud phone instances.
//
// Description:
//
// Before you reset a cloud phone instance, make sure it is in one of the following states: **Available, Stopped, Abnormal, Backup failure, and Restoration failure**.
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

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
	}

	if !tea.BoolValue(util.IsUnset(request.SettingResetType)) {
		query["SettingResetType"] = request.SettingResetType
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
	_result = &ResetAndroidInstancesInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets one or more cloud phone instances.
//
// Description:
//
// Before you reset a cloud phone instance, make sure it is in one of the following states: **Available, Stopped, Abnormal, Backup failure, and Restoration failure**.
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
// Executes a command on a cloud phone instance.
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
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a command on a cloud phone instance.
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
// Pushes files from Object Storage Service (OSS) buckets to cloud phone instances.
//
// Description:
//
// Currently, this operation allows you to only push files or folders from OSS buckets to cloud phone instances.
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

	if !tea.BoolValue(util.IsUnset(request.TargetFileName)) {
		query["TargetFileName"] = request.TargetFileName
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
	_result = &SendFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes files from Object Storage Service (OSS) buckets to cloud phone instances.
//
// Description:
//
// Currently, this operation allows you to only push files or folders from OSS buckets to cloud phone instances.
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

// Summary:
//
// Sets the authentication status for cloud phone instances. If you enable Android Debug Bridge (ADB) authentication for cloud phone instances, the system will verify the validity of the ADB key pairs provided by end users when they connect to the instances over ADB. To ensure successful authentication and a proper connection, we recommend that you attach ADB key pairs to cloud phone instances. If you disable ADB authentication for cloud phone instances, the system will no longer verify the validity of any ADB key pairs. As a result, end users can connect to the cloud phone instances over ADB without authentication, provided the network connection is functioning properly.
//
// Description:
//
// Before you call this operation, make sure that the desired cloud phone instance is in the Running state.
//
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
	_result = &SetAdbSecureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the authentication status for cloud phone instances. If you enable Android Debug Bridge (ADB) authentication for cloud phone instances, the system will verify the validity of the ADB key pairs provided by end users when they connect to the instances over ADB. To ensure successful authentication and a proper connection, we recommend that you attach ADB key pairs to cloud phone instances. If you disable ADB authentication for cloud phone instances, the system will no longer verify the validity of any ADB key pairs. As a result, end users can connect to the cloud phone instances over ADB without authentication, provided the network connection is functioning properly.
//
// Description:
//
// Before you call this operation, make sure that the desired cloud phone instance is in the Running state.
//
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

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
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
	_result = &StartAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
// Stops a cloud phone instance.
//
// Description:
//
// Before you stop a cloud phone instance, make sure it is in one of the following states: **Available, Backup failure, and Restoration failure**.
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

	if !tea.BoolValue(util.IsUnset(request.SaleMode)) {
		query["SaleMode"] = request.SaleMode
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
	_result = &StopAndroidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a cloud phone instance.
//
// Description:
//
// Before you stop a cloud phone instance, make sure it is in one of the following states: **Available, Backup failure, and Restoration failure**.
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
// Uninstalls an app from multiple cloud phone instances.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, you can visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
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
	_result = &UninstallAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an app from multiple cloud phone instances.
//
// Description:
//
// This operation runs asynchronously. To check the operation result, you can visit the Task Center. To retrieve task details, call the [DescribeTasks](~~DescribeTasks~~) operation.
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
	_result = &UpdateCustomImageNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
// Changes the image of an instance group.
//
// Description:
//
// Before you call this operation, make sure the image is in the Available state and the region of the image is included in the region list of the desired instance group. In addition, the instance group itself is available.
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
	_result = &UpdateInstanceGroupImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the image of an instance group.
//
// Description:
//
// Before you call this operation, make sure the image is in the Available state and the region of the image is included in the region list of the desired instance group. In addition, the instance group itself is available.
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
// 更新实例镜像
//
// @param request - UpdateInstanceImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceImageResponse
func (client *Client) UpdateInstanceImageWithOptions(request *UpdateInstanceImageRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceImage"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例镜像
//
// @param request - UpdateInstanceImageRequest
//
// @return UpdateInstanceImageResponse
func (client *Client) UpdateInstanceImage(request *UpdateInstanceImageRequest) (_result *UpdateInstanceImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceImageResponse{}
	_body, _err := client.UpdateInstanceImageWithOptions(request, runtime)
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
	_result = &UpgradeAndroidInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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

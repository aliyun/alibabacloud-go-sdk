// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDeploymentResponseBodyData) *GetDeploymentResponseBody
	GetData() *GetDeploymentResponseBodyData
	SetErrorCode(v string) *GetDeploymentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetDeploymentResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetDeploymentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentResponseBody
	GetSuccess() *bool
}

type GetDeploymentResponseBody struct {
	// The details of the deployment package.
	Data *GetDeploymentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can troubleshoot errors based on the ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBody) GetData() *GetDeploymentResponseBodyData {
	return s.Data
}

func (s *GetDeploymentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentResponseBody) SetData(v *GetDeploymentResponseBodyData) *GetDeploymentResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentResponseBody) SetErrorCode(v string) *GetDeploymentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentResponseBody) SetErrorMessage(v string) *GetDeploymentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentResponseBody) SetHttpStatusCode(v int32) *GetDeploymentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDeploymentResponseBody) SetRequestId(v string) *GetDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentResponseBody) SetSuccess(v bool) *GetDeploymentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeploymentResponseBodyData struct {
	// The deployed items.
	DeployedItems []*GetDeploymentResponseBodyDataDeployedItems `json:"DeployedItems,omitempty" xml:"DeployedItems,omitempty" type:"Repeated"`
	// The details of the deployment package.
	Deployment *GetDeploymentResponseBodyDataDeployment `json:"Deployment,omitempty" xml:"Deployment,omitempty" type:"Struct"`
}

func (s GetDeploymentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyData) GetDeployedItems() []*GetDeploymentResponseBodyDataDeployedItems {
	return s.DeployedItems
}

func (s *GetDeploymentResponseBodyData) GetDeployment() *GetDeploymentResponseBodyDataDeployment {
	return s.Deployment
}

func (s *GetDeploymentResponseBodyData) SetDeployedItems(v []*GetDeploymentResponseBodyDataDeployedItems) *GetDeploymentResponseBodyData {
	s.DeployedItems = v
	return s
}

func (s *GetDeploymentResponseBodyData) SetDeployment(v *GetDeploymentResponseBodyDataDeployment) *GetDeploymentResponseBodyData {
	s.Deployment = v
	return s
}

func (s *GetDeploymentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDeploymentResponseBodyDataDeployedItems struct {
	// The file ID.
	//
	// example:
	//
	// 507642378
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The version of the file.
	//
	// example:
	//
	// 7
	FileVersion *int64 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// - UNPUBLISHED(0): not published
	//
	// - SUCCESS(1): Published
	//
	// - ERROR(2): Publishing failed
	//
	// - CLONED(3): successfully CLONED
	//
	// - DEPLOY_ERROR(4): Publishing failed
	//
	// - CLONING(5): CLONING
	//
	// - REJECT(6): release rejected
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeploymentResponseBodyDataDeployedItems) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentResponseBodyDataDeployedItems) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyDataDeployedItems) GetFileId() *int64 {
	return s.FileId
}

func (s *GetDeploymentResponseBodyDataDeployedItems) GetFileVersion() *int64 {
	return s.FileVersion
}

func (s *GetDeploymentResponseBodyDataDeployedItems) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeploymentResponseBodyDataDeployedItems) SetFileId(v int64) *GetDeploymentResponseBodyDataDeployedItems {
	s.FileId = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployedItems) SetFileVersion(v int64) *GetDeploymentResponseBodyDataDeployedItems {
	s.FileVersion = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployedItems) SetStatus(v int32) *GetDeploymentResponseBodyDataDeployedItems {
	s.Status = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployedItems) Validate() error {
	return dara.Validate(s)
}

type GetDeploymentResponseBodyDataDeployment struct {
	// The check status of one or more files in the deployment task. If the value of the ToEnvironment parameter is 1, the files can be deployed to the production environment only when the value of the Status parameter is 1 and the CheckingStatus parameter is empty. Valid values:
	//
	// 	- 7: The file failed the check.
	//
	// 	- 8: The file is being checked.
	//
	// example:
	//
	// 7
	CheckingStatus *int32 `json:"CheckingStatus,omitempty" xml:"CheckingStatus,omitempty"`
	// The time when the deployment task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593877765000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who created the deployment task.
	//
	// example:
	//
	// 20030****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The error message that was returned when the deployment package failed. In this case, the value of the Status parameter is 2.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the deployment task was run. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593877765000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The environment in which the deployment task was run. Valid values: 0 and 1. A value of 0 indicates the on-premises environment. A value of 1 indicates the development environment.
	//
	// example:
	//
	// 0
	FromEnvironment *int32 `json:"FromEnvironment,omitempty" xml:"FromEnvironment,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who ran the deployment task.
	//
	// example:
	//
	// 2003****
	HandlerId *string `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The name of the deployment task. The value is the same as the name of the specific deployment task that is displayed on the Release Package page in the Deploy module.
	//
	// example:
	//
	// ods_user_info_d-2020-07-04_20030****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the deployment task. Valid values: 0, 1, and 2. A value of 0 indicates that the task is ready. A value of 1 indicates that the task was successful. A value of 2 indicates that the task failed.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The environment to which the file was deployed. Valid values: 1 and 2. A value of 1 indicates the development environment. A value of 2 indicates the production environment.
	//
	// example:
	//
	// 1
	ToEnvironment *int32 `json:"ToEnvironment,omitempty" xml:"ToEnvironment,omitempty"`
}

func (s GetDeploymentResponseBodyDataDeployment) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentResponseBodyDataDeployment) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyDataDeployment) GetCheckingStatus() *int32 {
	return s.CheckingStatus
}

func (s *GetDeploymentResponseBodyDataDeployment) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDeploymentResponseBodyDataDeployment) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDeploymentResponseBodyDataDeployment) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentResponseBodyDataDeployment) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *GetDeploymentResponseBodyDataDeployment) GetFromEnvironment() *int32 {
	return s.FromEnvironment
}

func (s *GetDeploymentResponseBodyDataDeployment) GetHandlerId() *string {
	return s.HandlerId
}

func (s *GetDeploymentResponseBodyDataDeployment) GetName() *string {
	return s.Name
}

func (s *GetDeploymentResponseBodyDataDeployment) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeploymentResponseBodyDataDeployment) GetToEnvironment() *int32 {
	return s.ToEnvironment
}

func (s *GetDeploymentResponseBodyDataDeployment) SetCheckingStatus(v int32) *GetDeploymentResponseBodyDataDeployment {
	s.CheckingStatus = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetCreateTime(v int64) *GetDeploymentResponseBodyDataDeployment {
	s.CreateTime = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetCreatorId(v string) *GetDeploymentResponseBodyDataDeployment {
	s.CreatorId = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetErrorMessage(v string) *GetDeploymentResponseBodyDataDeployment {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetExecuteTime(v int64) *GetDeploymentResponseBodyDataDeployment {
	s.ExecuteTime = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetFromEnvironment(v int32) *GetDeploymentResponseBodyDataDeployment {
	s.FromEnvironment = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetHandlerId(v string) *GetDeploymentResponseBodyDataDeployment {
	s.HandlerId = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetName(v string) *GetDeploymentResponseBodyDataDeployment {
	s.Name = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetStatus(v int32) *GetDeploymentResponseBodyDataDeployment {
	s.Status = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) SetToEnvironment(v int32) *GetDeploymentResponseBodyDataDeployment {
	s.ToEnvironment = &v
	return s
}

func (s *GetDeploymentResponseBodyDataDeployment) Validate() error {
	return dara.Validate(s)
}

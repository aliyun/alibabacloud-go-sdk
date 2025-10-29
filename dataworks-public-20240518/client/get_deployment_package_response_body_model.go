// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDeploymentPackageResponseBodyData) *GetDeploymentPackageResponseBody
	GetData() *GetDeploymentPackageResponseBodyData
	SetErrorCode(v string) *GetDeploymentPackageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentPackageResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetDeploymentPackageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetDeploymentPackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentPackageResponseBody
	GetSuccess() *bool
}

type GetDeploymentPackageResponseBody struct {
	// The deployment package details.
	Data *GetDeploymentPackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. Use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeploymentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentPackageResponseBody) GetData() *GetDeploymentPackageResponseBodyData {
	return s.Data
}

func (s *GetDeploymentPackageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentPackageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentPackageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDeploymentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentPackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentPackageResponseBody) SetData(v *GetDeploymentPackageResponseBodyData) *GetDeploymentPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentPackageResponseBody) SetErrorCode(v string) *GetDeploymentPackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentPackageResponseBody) SetErrorMessage(v string) *GetDeploymentPackageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentPackageResponseBody) SetHttpStatusCode(v int32) *GetDeploymentPackageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDeploymentPackageResponseBody) SetRequestId(v string) *GetDeploymentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentPackageResponseBody) SetSuccess(v bool) *GetDeploymentPackageResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentPackageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeploymentPackageResponseBodyData struct {
	// The deployment item details.
	DeployedItems []*GetDeploymentPackageResponseBodyDataDeployedItems `json:"DeployedItems,omitempty" xml:"DeployedItems,omitempty" type:"Repeated"`
	// The deployment package details.
	Deployment *GetDeploymentPackageResponseBodyDataDeployment `json:"Deployment,omitempty" xml:"Deployment,omitempty" type:"Struct"`
}

func (s GetDeploymentPackageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeploymentPackageResponseBodyData) GetDeployedItems() []*GetDeploymentPackageResponseBodyDataDeployedItems {
	return s.DeployedItems
}

func (s *GetDeploymentPackageResponseBodyData) GetDeployment() *GetDeploymentPackageResponseBodyDataDeployment {
	return s.Deployment
}

func (s *GetDeploymentPackageResponseBodyData) SetDeployedItems(v []*GetDeploymentPackageResponseBodyDataDeployedItems) *GetDeploymentPackageResponseBodyData {
	s.DeployedItems = v
	return s
}

func (s *GetDeploymentPackageResponseBodyData) SetDeployment(v *GetDeploymentPackageResponseBodyDataDeployment) *GetDeploymentPackageResponseBodyData {
	s.Deployment = v
	return s
}

func (s *GetDeploymentPackageResponseBodyData) Validate() error {
	if s.DeployedItems != nil {
		for _, item := range s.DeployedItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Deployment != nil {
		if err := s.Deployment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeploymentPackageResponseBodyDataDeployedItems struct {
	// The file ID.
	//
	// example:
	//
	// 5076****
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file version.
	//
	// example:
	//
	// 7
	FileVersion *int64 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// 	- UNPUBLISHED(0)
	//
	// 	- SUCCESS(1)
	//
	// 	- ERROR(2)
	//
	// 	- CLONED(3)
	//
	// 	- DEPLOY_ERROR(4)
	//
	// 	- CLONING(5)
	//
	// 	- REJECT(6)
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeploymentPackageResponseBodyDataDeployedItems) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentPackageResponseBodyDataDeployedItems) GoString() string {
	return s.String()
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) GetFileId() *int64 {
	return s.FileId
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) GetFileVersion() *int64 {
	return s.FileVersion
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) SetFileId(v int64) *GetDeploymentPackageResponseBodyDataDeployedItems {
	s.FileId = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) SetFileVersion(v int64) *GetDeploymentPackageResponseBodyDataDeployedItems {
	s.FileVersion = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) SetStatus(v int32) *GetDeploymentPackageResponseBodyDataDeployedItems {
	s.Status = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployedItems) Validate() error {
	return dara.Validate(s)
}

type GetDeploymentPackageResponseBodyDataDeployment struct {
	// The validation status of nodes in the deployment package. For packages deployed to the development environment (toEnviroment=1), you can only proceed to deploy to production if the package Status is 1 (succeeded) and CheckingStatus is empty (validation complete).
	//
	// 	- 7: Validation failed
	//
	// 	- 8: Validation in progress
	//
	// example:
	//
	// 7
	CheckingStatus *int32 `json:"CheckingStatus,omitempty" xml:"CheckingStatus,omitempty"`
	// The timestamp (in milliseconds) when the deployment package was created.
	//
	// example:
	//
	// 1593877765000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the user who created the deployment package.
	//
	// example:
	//
	// 20030****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The detailed error message when the deployment package fails (status is 2).
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The timestamp (in milliseconds) when the deployment started.
	//
	// example:
	//
	// 1593877765000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The environment where the deployment is executed. Valid values: 0 (local) and 1 (development).
	//
	// example:
	//
	// 0
	FromEnvironment *int32 `json:"FromEnvironment,omitempty" xml:"FromEnvironment,omitempty"`
	// The Alibaba Cloud account ID of the user who executed the deployment.
	//
	// example:
	//
	// 2003****
	HandlerId *string `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The deployment package name, displayed on the Deploy Center > Deployment Packages page.
	//
	// example:
	//
	// ods_user_info_d-2020-07-04_20030****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The current status of the deployment package. Valid values: 0 (ready), 1 (succeeded), and 2 (failed).
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The target environment for the deployment. Valid values: 1 (development) and 2 (production).
	//
	// example:
	//
	// 1
	ToEnvironment *int32 `json:"ToEnvironment,omitempty" xml:"ToEnvironment,omitempty"`
}

func (s GetDeploymentPackageResponseBodyDataDeployment) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentPackageResponseBodyDataDeployment) GoString() string {
	return s.String()
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetCheckingStatus() *int32 {
	return s.CheckingStatus
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetFromEnvironment() *int32 {
	return s.FromEnvironment
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetHandlerId() *string {
	return s.HandlerId
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetName() *string {
	return s.Name
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) GetToEnvironment() *int32 {
	return s.ToEnvironment
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetCheckingStatus(v int32) *GetDeploymentPackageResponseBodyDataDeployment {
	s.CheckingStatus = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetCreateTime(v int64) *GetDeploymentPackageResponseBodyDataDeployment {
	s.CreateTime = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetCreatorId(v string) *GetDeploymentPackageResponseBodyDataDeployment {
	s.CreatorId = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetErrorMessage(v string) *GetDeploymentPackageResponseBodyDataDeployment {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetExecuteTime(v int64) *GetDeploymentPackageResponseBodyDataDeployment {
	s.ExecuteTime = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetFromEnvironment(v int32) *GetDeploymentPackageResponseBodyDataDeployment {
	s.FromEnvironment = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetHandlerId(v string) *GetDeploymentPackageResponseBodyDataDeployment {
	s.HandlerId = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetName(v string) *GetDeploymentPackageResponseBodyDataDeployment {
	s.Name = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetStatus(v int32) *GetDeploymentPackageResponseBodyDataDeployment {
	s.Status = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) SetToEnvironment(v int32) *GetDeploymentPackageResponseBodyDataDeployment {
	s.ToEnvironment = &v
	return s
}

func (s *GetDeploymentPackageResponseBodyDataDeployment) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *UpdateImageTaskRequest
	GetBizRegionId() *string
	SetCode(v string) *UpdateImageTaskRequest
	GetCode() *string
	SetStatus(v string) *UpdateImageTaskRequest
	GetStatus() *string
	SetTaskId(v string) *UpdateImageTaskRequest
	GetTaskId() *string
	SetTaskStep(v string) *UpdateImageTaskRequest
	GetTaskStep() *string
	SetTaskType(v string) *UpdateImageTaskRequest
	GetTaskType() *string
	SetVersionId(v string) *UpdateImageTaskRequest
	GetVersionId() *string
}

type UpdateImageTaskRequest struct {
	// **Not enabled. Do not pass this parameter.*	- The task region is determined by the associated creation resource. This parameter is not used to switch regions.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status and is not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// **Not enabled. Do not pass this parameter.*	- This parameter is not involved in task status updates.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status and is not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status identifier reported by the creation resource. It must match the task type and the current creation stage. You cannot directly use the overall task status from GetTask as a substitute.
	//
	// Valid values for image creation tasks (IMAGE_MANUAL and IMAGE_MANUAL_RETRY):
	//
	// - **BUILDER_INIT_SUCCESS**: The creation environment is initialized.
	//
	// - **DOWNLOAD_SUCCESS**: The creation resources are downloaded.
	//
	// - **DOWNLOAD_FAILED**: The creation resources fail to be downloaded.
	//
	// - **BUILD_START**: The image creation starts.
	//
	// - **BUILD_SUCCESS**: The application installation and image configuration are successful.
	//
	// - **BUILD_FAILED**: The application installation or image configuration fails.
	//
	// - **BUILD_COMPLETE**: The creation is completed and the subsequent verification process starts.
	//
	// - **BUILD_COMPLETE_SKIP_VALIDATE**: The creation is completed in compatibility mode. Whether subsequent verification is performed is subject to the task query result.
	//
	// Valid values for application adaptation tasks:
	//
	// - **BUILDER_INIT_SUCCESS**: The creation environment is initialized.
	//
	// - **DOWNLOAD_SUCCESS**: The application resources are downloaded.
	//
	// - **DOWNLOAD_FAILED**: The application resources fail to be downloaded.
	//
	// - **BUILD_START**: The application recording starts.
	//
	// - **BUILD_SUCCESS**: The application recording is successful.
	//
	// - **BUILD_FAILED**: The application recording fails.
	//
	// - **NAS_DISPATCH_SUCCESS**: The application files are distributed to NAS.
	//
	// - **NAS_DISPATCH_FAILED**: The application files fail to be distributed to NAS.
	//
	// - **OSS_UPLOAD_SUCCESS**: The application files are uploaded to OSS.
	//
	// - **OSS_UPLOAD_FAILED**: The application files fail to be uploaded to OSS.
	//
	// - **UPLOAD_SUCCESS**: The application files are uploaded.
	//
	// - **UPLOAD_FAILED**: The application files fail to be uploaded.
	//
	// - **BUILD_COMPLETE**: The application creation is completed.
	//
	// example:
	//
	// BUILD_START
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task to be updated. The task and its creation resource must already exist.
	//
	// This parameter is required.
	//
	// example:
	//
	// tid-bp1example****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// **Not enabled. Do not pass this parameter.*	- This parameter is not supported for updating the task step.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status and is not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	TaskStep *string `json:"TaskStep,omitempty" xml:"TaskStep,omitempty"`
	// The actual type of the task to be updated. It must match the task corresponding to TaskId. Status updates support the following application adaptation and manual image creation tasks:
	//
	// - **ADAPT_AUTO**: Creates an application adaptation automatically.
	//
	// - **ADAPT_AUTO_RETRY**: Retries the automatic application adaptation.
	//
	// - **ADAPT_MANUAL**: Creates an application adaptation manually.
	//
	// - **ADAPT_MANUAL_RETRY**: Retries the manual application adaptation.
	//
	// - **IMAGE_MANUAL**: Creates an image manually.
	//
	// - **IMAGE_MANUAL_RETRY**: Retries the manual image creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE_MANUAL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// **Not enabled. Do not pass this parameter.*	- This parameter is not supported for modifying the image version associated with the task.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status and is not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateImageTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageTaskRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *UpdateImageTaskRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateImageTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateImageTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateImageTaskRequest) GetTaskStep() *string {
	return s.TaskStep
}

func (s *UpdateImageTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *UpdateImageTaskRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateImageTaskRequest) SetBizRegionId(v string) *UpdateImageTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *UpdateImageTaskRequest) SetCode(v string) *UpdateImageTaskRequest {
	s.Code = &v
	return s
}

func (s *UpdateImageTaskRequest) SetStatus(v string) *UpdateImageTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateImageTaskRequest) SetTaskId(v string) *UpdateImageTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateImageTaskRequest) SetTaskStep(v string) *UpdateImageTaskRequest {
	s.TaskStep = &v
	return s
}

func (s *UpdateImageTaskRequest) SetTaskType(v string) *UpdateImageTaskRequest {
	s.TaskType = &v
	return s
}

func (s *UpdateImageTaskRequest) SetVersionId(v string) *UpdateImageTaskRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateImageTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskForCreateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageCreateMode(v string) *StartTaskForCreateImageRequest
	GetImageCreateMode() *string
	SetImageId(v string) *StartTaskForCreateImageRequest
	GetImageId() *string
	SetInstanceId(v string) *StartTaskForCreateImageRequest
	GetInstanceId() *string
	SetJumpStep(v string) *StartTaskForCreateImageRequest
	GetJumpStep() *string
	SetPerformanceType(v string) *StartTaskForCreateImageRequest
	GetPerformanceType() *string
	SetProductType(v string) *StartTaskForCreateImageRequest
	GetProductType() *string
	SetResourceList(v []*StartTaskForCreateImageRequestResourceList) *StartTaskForCreateImageRequest
	GetResourceList() []*StartTaskForCreateImageRequestResourceList
	SetRetryType(v string) *StartTaskForCreateImageRequest
	GetRetryType() *string
	SetRunMode(v string) *StartTaskForCreateImageRequest
	GetRunMode() *string
	SetSourceRegion(v string) *StartTaskForCreateImageRequest
	GetSourceRegion() *string
	SetSubInstanceId(v string) *StartTaskForCreateImageRequest
	GetSubInstanceId() *string
	SetSystemDiskSize(v int32) *StartTaskForCreateImageRequest
	GetSystemDiskSize() *int32
	SetVersionId(v string) *StartTaskForCreateImageRequest
	GetVersionId() *string
}

type StartTaskForCreateImageRequest struct {
	// The image creation mode. Currently, use the following mode:
	//
	// - **BUILDER_MANUAL**: Manually enter the creation environment to install applications and configure the image.
	//
	// example:
	//
	// BUILDER_MANUAL
	ImageCreateMode *string `json:"ImageCreateMode,omitempty" xml:"ImageCreateMode,omitempty"`
	// The ID of the image. This parameter is not exposed by default.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The WUYING instance ID. The instance can be a cloud desktop instance or a workstation instance. To ensure data consistency in the image, stop the instance before initiating image creation.
	//
	// example:
	//
	// ecd-5u87dainamz54yk10
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The step to jump to.
	//
	// example:
	//
	// CREATE_VALIDATE_ENV
	JumpStep *string `json:"JumpStep,omitempty" xml:"JumpStep,omitempty"`
	// The resource type.
	//
	// example:
	//
	// CPU
	PerformanceType *string `json:"PerformanceType,omitempty" xml:"PerformanceType,omitempty"`
	// The product type. Set this parameter to the following fixed value when you call the operation:
	//
	// - CloudApp: WUYING Cloud App.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of resources required for image creation, such as the download URLs of applications to be installed. You can specify up to 100 items. The ResourceType and Resource of each item must correspond to each other.
	//
	// example:
	//
	// List of Object
	ResourceList []*StartTaskForCreateImageRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The retry type.
	//
	// example:
	//
	// RETRY
	RetryType *string `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	// The running mode of the image creation environment. Valid values:
	//
	// - **NORMAL**: Normal running mode.
	//
	// - **COMPATIBLE**: Compatible running mode.
	//
	// example:
	//
	// NORMAL
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The source region where the image is created. If you do not specify this parameter, a region is randomly selected.
	//
	// example:
	//
	// cn-shanghai
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The sub-instance ID. This parameter is not applicable to cloud desktop scenarios. In workstation scenarios, pass the persistent session ID to ensure that a specific instance is located.
	//
	// example:
	//
	// p-0cc7s3n1l*****
	SubInstanceId *string `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
	// The size of the system cloud disk. Unit: GB.
	//
	// example:
	//
	// 100
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The image version. This parameter is currently not exposed.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s StartTaskForCreateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForCreateImageRequest) GoString() string {
	return s.String()
}

func (s *StartTaskForCreateImageRequest) GetImageCreateMode() *string {
	return s.ImageCreateMode
}

func (s *StartTaskForCreateImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *StartTaskForCreateImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartTaskForCreateImageRequest) GetJumpStep() *string {
	return s.JumpStep
}

func (s *StartTaskForCreateImageRequest) GetPerformanceType() *string {
	return s.PerformanceType
}

func (s *StartTaskForCreateImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *StartTaskForCreateImageRequest) GetResourceList() []*StartTaskForCreateImageRequestResourceList {
	return s.ResourceList
}

func (s *StartTaskForCreateImageRequest) GetRetryType() *string {
	return s.RetryType
}

func (s *StartTaskForCreateImageRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *StartTaskForCreateImageRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *StartTaskForCreateImageRequest) GetSubInstanceId() *string {
	return s.SubInstanceId
}

func (s *StartTaskForCreateImageRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *StartTaskForCreateImageRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *StartTaskForCreateImageRequest) SetImageCreateMode(v string) *StartTaskForCreateImageRequest {
	s.ImageCreateMode = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetImageId(v string) *StartTaskForCreateImageRequest {
	s.ImageId = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetInstanceId(v string) *StartTaskForCreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetJumpStep(v string) *StartTaskForCreateImageRequest {
	s.JumpStep = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetPerformanceType(v string) *StartTaskForCreateImageRequest {
	s.PerformanceType = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetProductType(v string) *StartTaskForCreateImageRequest {
	s.ProductType = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetResourceList(v []*StartTaskForCreateImageRequestResourceList) *StartTaskForCreateImageRequest {
	s.ResourceList = v
	return s
}

func (s *StartTaskForCreateImageRequest) SetRetryType(v string) *StartTaskForCreateImageRequest {
	s.RetryType = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetRunMode(v string) *StartTaskForCreateImageRequest {
	s.RunMode = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetSourceRegion(v string) *StartTaskForCreateImageRequest {
	s.SourceRegion = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetSubInstanceId(v string) *StartTaskForCreateImageRequest {
	s.SubInstanceId = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetSystemDiskSize(v int32) *StartTaskForCreateImageRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *StartTaskForCreateImageRequest) SetVersionId(v string) *StartTaskForCreateImageRequest {
	s.VersionId = &v
	return s
}

func (s *StartTaskForCreateImageRequest) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartTaskForCreateImageRequestResourceList struct {
	// The resource information.
	//
	// example:
	//
	// https://example.com/apps/business/setup.exe
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The resource type.
	//
	// example:
	//
	// URL
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s StartTaskForCreateImageRequestResourceList) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForCreateImageRequestResourceList) GoString() string {
	return s.String()
}

func (s *StartTaskForCreateImageRequestResourceList) GetResource() *string {
	return s.Resource
}

func (s *StartTaskForCreateImageRequestResourceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *StartTaskForCreateImageRequestResourceList) SetResource(v string) *StartTaskForCreateImageRequestResourceList {
	s.Resource = &v
	return s
}

func (s *StartTaskForCreateImageRequestResourceList) SetResourceType(v string) *StartTaskForCreateImageRequestResourceList {
	s.ResourceType = &v
	return s
}

func (s *StartTaskForCreateImageRequestResourceList) Validate() error {
	return dara.Validate(s)
}

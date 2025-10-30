// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageByInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCleanUserdata(v bool) *CreateImageByInstanceRequest
	GetAutoCleanUserdata() *bool
	SetBizType(v int32) *CreateImageByInstanceRequest
	GetBizType() *int32
	SetDescription(v string) *CreateImageByInstanceRequest
	GetDescription() *string
	SetDiskType(v string) *CreateImageByInstanceRequest
	GetDiskType() *string
	SetImageName(v string) *CreateImageByInstanceRequest
	GetImageName() *string
	SetInstanceId(v string) *CreateImageByInstanceRequest
	GetInstanceId() *string
	SetInstanceType(v string) *CreateImageByInstanceRequest
	GetInstanceType() *string
	SetProductType(v string) *CreateImageByInstanceRequest
	GetProductType() *string
	SetSubInstanceId(v string) *CreateImageByInstanceRequest
	GetSubInstanceId() *string
}

type CreateImageByInstanceRequest struct {
	// This parameter is applicable only to scenarios in which the instance type is Cloud Desktop. Specifies whether to clear private data of users. If this parameter is set to true, the created image clears data in directories other than Administrator and Public in the C:\\Users directory.
	//
	// Valid values:
	//
	// 	- true: cleanup.
	//
	// 	- false: does not clear.
	//
	// example:
	//
	// false
	AutoCleanUserdata *bool `json:"AutoCleanUserdata,omitempty" xml:"AutoCleanUserdata,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// my test image v1.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of disk data contained in the image. By default, the system disk and data disk of the instance are included.
	//
	// Valid values:
	//
	// 	- SYSTEM: only system disk.
	//
	// 	- ALL: system disk + data disk
	//
	// example:
	//
	// ALL
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the RDS instance. The instance can be a CloudDesktop instance, a workstation instance. To ensure data consistency in the image, we recommend that you shut down the instance before you create an image.
	//
	// example:
	//
	// ws-0buj1s9gm******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type.
	//
	// Valid values:
	//
	// 	- CloudDesktop: Cloud Desktop.
	//
	// 	- WuyingServer: Workstation
	//
	// example:
	//
	// WuyingServer
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the child instance. This parameter is not used in cloud computing scenarios. Workstation scenarios, you need to specify a persistent session ID to ensure that a specific instance is located.
	//
	// example:
	//
	// p-0cc7s3n1l*****
	SubInstanceId *string `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
}

func (s CreateImageByInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageByInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateImageByInstanceRequest) GetAutoCleanUserdata() *bool {
	return s.AutoCleanUserdata
}

func (s *CreateImageByInstanceRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *CreateImageByInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageByInstanceRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateImageByInstanceRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImageByInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageByInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateImageByInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateImageByInstanceRequest) GetSubInstanceId() *string {
	return s.SubInstanceId
}

func (s *CreateImageByInstanceRequest) SetAutoCleanUserdata(v bool) *CreateImageByInstanceRequest {
	s.AutoCleanUserdata = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetBizType(v int32) *CreateImageByInstanceRequest {
	s.BizType = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetDescription(v string) *CreateImageByInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetDiskType(v string) *CreateImageByInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetImageName(v string) *CreateImageByInstanceRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetInstanceId(v string) *CreateImageByInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetInstanceType(v string) *CreateImageByInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetProductType(v string) *CreateImageByInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetSubInstanceId(v string) *CreateImageByInstanceRequest {
	s.SubInstanceId = &v
	return s
}

func (s *CreateImageByInstanceRequest) Validate() error {
	return dara.Validate(s)
}

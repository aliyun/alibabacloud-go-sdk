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
	SetCopyProfile(v bool) *CreateImageByInstanceRequest
	GetCopyProfile() *bool
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
	SetTagList(v []*CreateImageByInstanceRequestTagList) *CreateImageByInstanceRequest
	GetTagList() []*CreateImageByInstanceRequestTagList
}

type CreateImageByInstanceRequest struct {
	// This parameter applies only to Cloud Desktop scenarios. Specifies whether to clear user personal data. If set to true, the created image clears data in all directories under C:\\Users except Administrator and Public.
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
	BizType     *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CopyProfile *bool  `json:"CopyProfile,omitempty" xml:"CopyProfile,omitempty"`
	// The image description.
	//
	// example:
	//
	// my test image v1.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of disk data included in the image. By default, both the system cloud disk and data cloud disk of the instance are included.
	//
	// example:
	//
	// ALL
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The image name.
	//
	// example:
	//
	// test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The WUYING instance ID. The instance can be a Cloud Desktop instance or a workstation instance. To ensure data consistency in the image, stop the instance before creating the image.
	//
	// example:
	//
	// ws-0buj1s9gm******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type.
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
	// The sub-instance ID. This parameter does not apply to Cloud Desktop scenarios. In workstation scenarios, specify the persistent session ID to identify a specific instance.
	//
	// example:
	//
	// p-0cc7s3n1l*****
	SubInstanceId *string                                `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
	TagList       []*CreateImageByInstanceRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
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

func (s *CreateImageByInstanceRequest) GetCopyProfile() *bool {
	return s.CopyProfile
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

func (s *CreateImageByInstanceRequest) GetTagList() []*CreateImageByInstanceRequestTagList {
	return s.TagList
}

func (s *CreateImageByInstanceRequest) SetAutoCleanUserdata(v bool) *CreateImageByInstanceRequest {
	s.AutoCleanUserdata = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetBizType(v int32) *CreateImageByInstanceRequest {
	s.BizType = &v
	return s
}

func (s *CreateImageByInstanceRequest) SetCopyProfile(v bool) *CreateImageByInstanceRequest {
	s.CopyProfile = &v
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

func (s *CreateImageByInstanceRequest) SetTagList(v []*CreateImageByInstanceRequestTagList) *CreateImageByInstanceRequest {
	s.TagList = v
	return s
}

func (s *CreateImageByInstanceRequest) Validate() error {
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImageByInstanceRequestTagList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageByInstanceRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s CreateImageByInstanceRequestTagList) GoString() string {
	return s.String()
}

func (s *CreateImageByInstanceRequestTagList) GetKey() *string {
	return s.Key
}

func (s *CreateImageByInstanceRequestTagList) GetValue() *string {
	return s.Value
}

func (s *CreateImageByInstanceRequestTagList) SetKey(v string) *CreateImageByInstanceRequestTagList {
	s.Key = &v
	return s
}

func (s *CreateImageByInstanceRequestTagList) SetValue(v string) *CreateImageByInstanceRequestTagList {
	s.Value = &v
	return s
}

func (s *CreateImageByInstanceRequestTagList) Validate() error {
	return dara.Validate(s)
}

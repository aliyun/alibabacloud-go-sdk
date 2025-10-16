// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*DescribeImagesResponseBodyImages) *DescribeImagesResponseBody
	GetImages() []*DescribeImagesResponseBodyImages
	SetNextToken(v string) *DescribeImagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeImagesResponseBody
	GetRequestId() *string
}

type DescribeImagesResponseBody struct {
	// The details of the images.
	Images []*DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The token that determines the start point of the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4636DBE0-BBB4-4076-8B8E-94D21A9A3CFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) GetImages() []*DescribeImagesResponseBodyImages {
	return s.Images
}

func (s *DescribeImagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImagesResponseBody) SetImages(v []*DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImagesResponseBody) SetNextToken(v string) *DescribeImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagesResponseBodyImages struct {
	// The version of the image.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2018-01-10T01:01:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 150
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the image is a GPU-accelerated image.
	//
	// example:
	//
	// false
	GpuCategory *bool `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The version number of the GPU driver.
	//
	// example:
	//
	// 417.22
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image.
	//
	// Valid values:
	//
	// 	- SYSTEM
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CUSTOM
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// testImageName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// WINDOWS
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The operating system type of the image.
	//
	// Valid values:
	//
	// 	- Ubuntu
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows Server 2022
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- UOS
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CentOS
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows Server 2019
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- SQL Server 2016
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows 10
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows Server 2019
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The creation progress of the image. Unit: %.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX: High-definition Experience (HDX) protocol
	//
	// 	- ASP: in-house Adaptive Streaming Protocol (ASP) (recommended)
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The type of the image session.
	//
	// Valid values:
	//
	// 	- SINGLE_SESSION: single-session image.
	//
	// 	- MULTIPLE_SESSION: multi-session image.
	//
	// example:
	//
	// MULTIPLE_SESSION
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The number of shared images.
	//
	// example:
	//
	// 1
	SharedCount *int32 `json:"SharedCount,omitempty" xml:"SharedCount,omitempty"`
	// The size of the image. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the image.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Available
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The languages of the operating system.
	SupportedLanguages []*string `json:"SupportedLanguages,omitempty" xml:"SupportedLanguages,omitempty" type:"Repeated"`
	// The time when the image was last modified.
	//
	// example:
	//
	// 2021-12-22T02:48:43Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used when disk encryption is enabled. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to query the list of KMS keys.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeImagesResponseBodyImages) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImagesResponseBodyImages) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribeImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *DescribeImagesResponseBodyImages) GetGpuCategory() *bool {
	return s.GpuCategory
}

func (s *DescribeImagesResponseBodyImages) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagesResponseBodyImages) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeImagesResponseBodyImages) GetName() *string {
	return s.Name
}

func (s *DescribeImagesResponseBodyImages) GetOsType() *string {
	return s.OsType
}

func (s *DescribeImagesResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeImagesResponseBodyImages) GetProgress() *string {
	return s.Progress
}

func (s *DescribeImagesResponseBodyImages) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeImagesResponseBodyImages) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeImagesResponseBodyImages) GetSharedCount() *int32 {
	return s.SharedCount
}

func (s *DescribeImagesResponseBodyImages) GetSize() *int32 {
	return s.Size
}

func (s *DescribeImagesResponseBodyImages) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagesResponseBodyImages) GetSupportedLanguages() []*string {
	return s.SupportedLanguages
}

func (s *DescribeImagesResponseBodyImages) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeImagesResponseBodyImages) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *DescribeImagesResponseBodyImages) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *DescribeImagesResponseBodyImages) SetAppVersion(v string) *DescribeImagesResponseBodyImages {
	s.AppVersion = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetCreationTime(v string) *DescribeImagesResponseBodyImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetDataDiskSize(v int32) *DescribeImagesResponseBodyImages {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetDescription(v string) *DescribeImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetGpuCategory(v bool) *DescribeImagesResponseBodyImages {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetGpuDriverVersion(v string) *DescribeImagesResponseBodyImages {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageId(v string) *DescribeImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageType(v string) *DescribeImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetName(v string) *DescribeImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetOsType(v string) *DescribeImagesResponseBodyImages {
	s.OsType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetPlatform(v string) *DescribeImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProgress(v string) *DescribeImagesResponseBodyImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProtocolType(v string) *DescribeImagesResponseBodyImages {
	s.ProtocolType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSessionType(v string) *DescribeImagesResponseBodyImages {
	s.SessionType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSharedCount(v int32) *DescribeImagesResponseBodyImages {
	s.SharedCount = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSize(v int32) *DescribeImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetStatus(v string) *DescribeImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSupportedLanguages(v []*string) *DescribeImagesResponseBodyImages {
	s.SupportedLanguages = v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetUpdateTime(v string) *DescribeImagesResponseBodyImages {
	s.UpdateTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetVolumeEncryptionEnabled(v bool) *DescribeImagesResponseBodyImages {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetVolumeEncryptionKey(v string) *DescribeImagesResponseBodyImages {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}

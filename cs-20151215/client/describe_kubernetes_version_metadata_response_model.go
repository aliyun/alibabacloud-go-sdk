// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKubernetesVersionMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKubernetesVersionMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKubernetesVersionMetadataResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeKubernetesVersionMetadataResponseBody) *DescribeKubernetesVersionMetadataResponse
	GetBody() []*DescribeKubernetesVersionMetadataResponseBody
}

type DescribeKubernetesVersionMetadataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeKubernetesVersionMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeKubernetesVersionMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKubernetesVersionMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKubernetesVersionMetadataResponse) GetBody() []*DescribeKubernetesVersionMetadataResponseBody {
	return s.Body
}

func (s *DescribeKubernetesVersionMetadataResponse) SetHeaders(v map[string]*string) *DescribeKubernetesVersionMetadataResponse {
	s.Headers = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponse) SetStatusCode(v int32) *DescribeKubernetesVersionMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponse) SetBody(v []*DescribeKubernetesVersionMetadataResponseBody) *DescribeKubernetesVersionMetadataResponse {
	s.Body = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKubernetesVersionMetadataResponseBody struct {
	// Features of the queried Kubernetes version.
	//
	// example:
	//
	// {
	//
	//       "AnyAZ": true,
	//
	//       "ChargeType": "PostPaid"
	//
	// }
	Capabilities map[string]interface{} `json:"capabilities,omitempty" xml:"capabilities,omitempty"`
	// The OS images that are returned.
	Images []*DescribeKubernetesVersionMetadataResponseBodyImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The metadata of the Kubernetes version.
	//
	// example:
	//
	// {
	//
	//       "KubernetesVersion": "1.31.1-aliyun.1",
	//
	//       "SubClass": "default",
	//
	//       "ServiceCIDR": ""
	//
	// }
	MetaData map[string]interface{} `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// The container runtime configurations.
	Runtimes []*Runtime `json:"runtimes,omitempty" xml:"runtimes,omitempty" type:"Repeated"`
	// The Kubernetes version supported by ACK. For more information, see [Release notes for Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.16.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The release date of the Kubernetes version.
	//
	// example:
	//
	// 2023-04-30T00:00:00Z
	ReleaseDate *string `json:"release_date,omitempty" xml:"release_date,omitempty"`
	// The expiration date of the Kubernetes version.
	//
	// example:
	//
	// 2025-04-30T00:00:00Z
	ExpirationDate *string `json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
	// Indicates whether you can create clusters that run the Kubernetes version.
	//
	// example:
	//
	// true
	Creatable *bool `json:"creatable,omitempty" xml:"creatable,omitempty"`
	// The list of available Kubernetes versions for updates.
	UpgradableVersions []*string `json:"upgradable_versions,omitempty" xml:"upgradable_versions,omitempty" type:"Repeated"`
}

func (s DescribeKubernetesVersionMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetCapabilities() map[string]interface{} {
	return s.Capabilities
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetImages() []*DescribeKubernetesVersionMetadataResponseBodyImages {
	return s.Images
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetRuntimes() []*Runtime {
	return s.Runtimes
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetCreatable() *bool {
	return s.Creatable
}

func (s *DescribeKubernetesVersionMetadataResponseBody) GetUpgradableVersions() []*string {
	return s.UpgradableVersions
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetCapabilities(v map[string]interface{}) *DescribeKubernetesVersionMetadataResponseBody {
	s.Capabilities = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetImages(v []*DescribeKubernetesVersionMetadataResponseBodyImages) *DescribeKubernetesVersionMetadataResponseBody {
	s.Images = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetMetaData(v map[string]interface{}) *DescribeKubernetesVersionMetadataResponseBody {
	s.MetaData = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetRuntimes(v []*Runtime) *DescribeKubernetesVersionMetadataResponseBody {
	s.Runtimes = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetVersion(v string) *DescribeKubernetesVersionMetadataResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetReleaseDate(v string) *DescribeKubernetesVersionMetadataResponseBody {
	s.ReleaseDate = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetExpirationDate(v string) *DescribeKubernetesVersionMetadataResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetCreatable(v bool) *DescribeKubernetesVersionMetadataResponseBody {
	s.Creatable = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) SetUpgradableVersions(v []*string) *DescribeKubernetesVersionMetadataResponseBody {
	s.UpgradableVersions = v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Runtimes != nil {
		for _, item := range s.Runtimes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKubernetesVersionMetadataResponseBodyImages struct {
	// The ID of the image.
	//
	// example:
	//
	// centos_7_7_x64_20G_alibase_20200426.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The image name.
	//
	// example:
	//
	// CentOS 7.7
	ImageName *string `json:"image_name,omitempty" xml:"image_name,omitempty"`
	// The OS platform. You can obtain the terminal ID by calling one of the following operations:
	//
	// 	- `AliyunLinux`
	//
	// 	- `CentOS`
	//
	// 	- `Windows`
	//
	// 	- `WindowsCore`
	//
	// example:
	//
	// CentOS
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// 7.7
	OsVersion *string `json:"os_version,omitempty" xml:"os_version,omitempty"`
	// The type of operating system distribution that you want to use. We recommend that you use this parameter to specify the node operating system. You can obtain the terminal ID by calling one of the following operations:
	//
	// 	- `CentOS`
	//
	// 	- `AliyunLinux`
	//
	// 	- `AliyunLinux Qboot`
	//
	// 	- `AliyunLinuxUEFI`
	//
	// 	- `AliyunLinux3`
	//
	// 	- `Windows`
	//
	// 	- `WindowsCore`
	//
	// 	- `AliyunLinux3Arm64`
	//
	// 	- `ContainerOS`
	//
	// example:
	//
	// centos_7_7_20
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The type of OS. Examples:
	//
	// 	- `Windows`
	//
	// 	- `Linux`
	//
	// example:
	//
	// Linux
	OsType *string `json:"os_type,omitempty" xml:"os_type,omitempty"`
	// The type of image. Valid values:
	//
	// 	- `system`: public image
	//
	// 	- `self`: custom image
	//
	// 	- `others`: shared image from other Alibaba Cloud accounts
	//
	// 	- `marketplace`: image from the marketplace
	//
	// example:
	//
	// system
	ImageCategory *string `json:"image_category,omitempty" xml:"image_category,omitempty"`
	// The architecture of the image.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
}

func (s DescribeKubernetesVersionMetadataResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetOsVersion() *string {
	return s.OsVersion
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetOsType() *string {
	return s.OsType
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetImageId(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetImageName(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetPlatform(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetOsVersion(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.OsVersion = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetImageType(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetOsType(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.OsType = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetImageCategory(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.ImageCategory = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) SetArchitecture(v string) *DescribeKubernetesVersionMetadataResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataResponseBodyImages) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopInstanceType(v string) *DescribeImagesRequest
	GetDesktopInstanceType() *string
	SetFotaVersion(v string) *DescribeImagesRequest
	GetFotaVersion() *string
	SetGpuCategory(v bool) *DescribeImagesRequest
	GetGpuCategory() *bool
	SetGpuDriverVersion(v string) *DescribeImagesRequest
	GetGpuDriverVersion() *string
	SetImageId(v []*string) *DescribeImagesRequest
	GetImageId() []*string
	SetImageName(v string) *DescribeImagesRequest
	GetImageName() *string
	SetImageStatus(v string) *DescribeImagesRequest
	GetImageStatus() *string
	SetImageType(v string) *DescribeImagesRequest
	GetImageType() *string
	SetLanguageType(v string) *DescribeImagesRequest
	GetLanguageType() *string
	SetMaxResults(v int32) *DescribeImagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImagesRequest
	GetNextToken() *string
	SetOsType(v string) *DescribeImagesRequest
	GetOsType() *string
	SetProtocolType(v string) *DescribeImagesRequest
	GetProtocolType() *string
	SetRegionId(v string) *DescribeImagesRequest
	GetRegionId() *string
	SetSessionType(v string) *DescribeImagesRequest
	GetSessionType() *string
}

type DescribeImagesRequest struct {
	// The cloud desktop specification. You can call [DescribeDesktopTypes](https://help.aliyun.com/document_detail/436816.html) to obtain this parameter.
	//
	// example:
	//
	// ecd.graphics.xlarge
	DesktopInstanceType *string `json:"DesktopInstanceType,omitempty" xml:"DesktopInstanceType,omitempty"`
	// The image version information.
	//
	// example:
	//
	// 0.0.3-R-20220616.133609
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// Specifies whether the image is a GPU image.
	//
	// example:
	//
	// false
	GpuCategory *bool `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The GPU driver version.
	//
	// example:
	//
	// 417.22
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The image ID. You can specify one or more image IDs.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The image name.
	//
	// example:
	//
	// Win_01
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image status.
	//
	// example:
	//
	// Available
	ImageStatus *string `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	// The image type.
	//
	// example:
	//
	// SYSTEM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The operating system language.
	//
	// example:
	//
	// en-US
	LanguageType *string `json:"LanguageType,omitempty" xml:"LanguageType,omitempty"`
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 100
	//
	// - Default value: 10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. An empty value indicates that there is no next page.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operating system type. Default value: `null`.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session type.
	//
	// example:
	//
	// SINGLE_SESSION
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) GetDesktopInstanceType() *string {
	return s.DesktopInstanceType
}

func (s *DescribeImagesRequest) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *DescribeImagesRequest) GetGpuCategory() *bool {
	return s.GpuCategory
}

func (s *DescribeImagesRequest) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeImagesRequest) GetImageId() []*string {
	return s.ImageId
}

func (s *DescribeImagesRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagesRequest) GetImageStatus() *string {
	return s.ImageStatus
}

func (s *DescribeImagesRequest) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeImagesRequest) GetLanguageType() *string {
	return s.LanguageType
}

func (s *DescribeImagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImagesRequest) GetOsType() *string {
	return s.OsType
}

func (s *DescribeImagesRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagesRequest) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeImagesRequest) SetDesktopInstanceType(v string) *DescribeImagesRequest {
	s.DesktopInstanceType = &v
	return s
}

func (s *DescribeImagesRequest) SetFotaVersion(v string) *DescribeImagesRequest {
	s.FotaVersion = &v
	return s
}

func (s *DescribeImagesRequest) SetGpuCategory(v bool) *DescribeImagesRequest {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesRequest) SetGpuDriverVersion(v string) *DescribeImagesRequest {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v []*string) *DescribeImagesRequest {
	s.ImageId = v
	return s
}

func (s *DescribeImagesRequest) SetImageName(v string) *DescribeImagesRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesRequest) SetImageStatus(v string) *DescribeImagesRequest {
	s.ImageStatus = &v
	return s
}

func (s *DescribeImagesRequest) SetImageType(v string) *DescribeImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesRequest) SetLanguageType(v string) *DescribeImagesRequest {
	s.LanguageType = &v
	return s
}

func (s *DescribeImagesRequest) SetMaxResults(v int32) *DescribeImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagesRequest) SetNextToken(v string) *DescribeImagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesRequest) SetOsType(v string) *DescribeImagesRequest {
	s.OsType = &v
	return s
}

func (s *DescribeImagesRequest) SetProtocolType(v string) *DescribeImagesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeImagesRequest) SetRegionId(v string) *DescribeImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetSessionType(v string) *DescribeImagesRequest {
	s.SessionType = &v
	return s
}

func (s *DescribeImagesRequest) Validate() error {
	return dara.Validate(s)
}

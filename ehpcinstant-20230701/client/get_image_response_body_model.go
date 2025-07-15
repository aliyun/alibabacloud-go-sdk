// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody
	GetImage() *GetImageResponseBodyImage
	SetRequestId(v string) *GetImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetImageResponseBody
	GetTotalCount() *int32
}

type GetImageResponseBody struct {
	Image *GetImageResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) GetImage() *GetImageResponseBodyImage {
	return s.Image
}

func (s *GetImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetImageResponseBody) SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody {
	s.Image = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSuccess(v bool) *GetImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageResponseBody) SetTotalCount(v int32) *GetImageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImage struct {
	AdditionalRegionsInfo []*GetImageResponseBodyImageAdditionalRegionsInfo `json:"AdditionalRegionsInfo,omitempty" xml:"AdditionalRegionsInfo,omitempty" type:"Repeated"`
	AppId                 *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerImageSpec    *GetImageResponseBodyImageContainerImageSpec      `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	// example:
	//
	// 2022-12-23T09:51:39Z
	CreateTime   *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentInfo *GetImageResponseBodyImageDocumentInfo `json:"DocumentInfo,omitempty" xml:"DocumentInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// app-image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 40 GiB
	Size        *string                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Status      *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	VMImageSpec *GetImageResponseBodyImageVMImageSpec `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty" type:"Struct"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetImageResponseBodyImage) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImage) GetAdditionalRegionsInfo() []*GetImageResponseBodyImageAdditionalRegionsInfo {
	return s.AdditionalRegionsInfo
}

func (s *GetImageResponseBodyImage) GetAppId() *string {
	return s.AppId
}

func (s *GetImageResponseBodyImage) GetContainerImageSpec() *GetImageResponseBodyImageContainerImageSpec {
	return s.ContainerImageSpec
}

func (s *GetImageResponseBodyImage) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetImageResponseBodyImage) GetDescription() *string {
	return s.Description
}

func (s *GetImageResponseBodyImage) GetDocumentInfo() *GetImageResponseBodyImageDocumentInfo {
	return s.DocumentInfo
}

func (s *GetImageResponseBodyImage) GetImageType() *string {
	return s.ImageType
}

func (s *GetImageResponseBodyImage) GetName() *string {
	return s.Name
}

func (s *GetImageResponseBodyImage) GetSize() *string {
	return s.Size
}

func (s *GetImageResponseBodyImage) GetStatus() *string {
	return s.Status
}

func (s *GetImageResponseBodyImage) GetVMImageSpec() *GetImageResponseBodyImageVMImageSpec {
	return s.VMImageSpec
}

func (s *GetImageResponseBodyImage) GetVersion() *string {
	return s.Version
}

func (s *GetImageResponseBodyImage) SetAdditionalRegionsInfo(v []*GetImageResponseBodyImageAdditionalRegionsInfo) *GetImageResponseBodyImage {
	s.AdditionalRegionsInfo = v
	return s
}

func (s *GetImageResponseBodyImage) SetAppId(v string) *GetImageResponseBodyImage {
	s.AppId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetContainerImageSpec(v *GetImageResponseBodyImageContainerImageSpec) *GetImageResponseBodyImage {
	s.ContainerImageSpec = v
	return s
}

func (s *GetImageResponseBodyImage) SetCreateTime(v string) *GetImageResponseBodyImage {
	s.CreateTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetDescription(v string) *GetImageResponseBodyImage {
	s.Description = &v
	return s
}

func (s *GetImageResponseBodyImage) SetDocumentInfo(v *GetImageResponseBodyImageDocumentInfo) *GetImageResponseBodyImage {
	s.DocumentInfo = v
	return s
}

func (s *GetImageResponseBodyImage) SetImageType(v string) *GetImageResponseBodyImage {
	s.ImageType = &v
	return s
}

func (s *GetImageResponseBodyImage) SetName(v string) *GetImageResponseBodyImage {
	s.Name = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSize(v string) *GetImageResponseBodyImage {
	s.Size = &v
	return s
}

func (s *GetImageResponseBodyImage) SetStatus(v string) *GetImageResponseBodyImage {
	s.Status = &v
	return s
}

func (s *GetImageResponseBodyImage) SetVMImageSpec(v *GetImageResponseBodyImageVMImageSpec) *GetImageResponseBodyImage {
	s.VMImageSpec = v
	return s
}

func (s *GetImageResponseBodyImage) SetVersion(v string) *GetImageResponseBodyImage {
	s.Version = &v
	return s
}

func (s *GetImageResponseBodyImage) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImageAdditionalRegionsInfo struct {
	ImageId  *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetImageResponseBodyImageAdditionalRegionsInfo) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageAdditionalRegionsInfo) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) GetStatus() *string {
	return s.Status
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) SetImageId(v string) *GetImageResponseBodyImageAdditionalRegionsInfo {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) SetRegionId(v string) *GetImageResponseBodyImageAdditionalRegionsInfo {
	s.RegionId = &v
	return s
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) SetStatus(v string) *GetImageResponseBodyImageAdditionalRegionsInfo {
	s.Status = &v
	return s
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImageContainerImageSpec struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// True
	IsACREnterprise *bool `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	// example:
	//
	// True
	IsACRRegistry      *bool                                                          `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	OsTag              *string                                                        `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform           *string                                                        `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegistryCredential *GetImageResponseBodyImageContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	// example:
	//
	// cri-xyz795ygf8k9****
	RegistryCriId *string `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc_open/nginx:latest
	RegistryUrl *string `json:"RegistryUrl,omitempty" xml:"RegistryUrl,omitempty"`
}

func (s GetImageResponseBodyImageContainerImageSpec) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageContainerImageSpec) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetArchitecture() *string {
	return s.Architecture
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetIsACREnterprise() *bool {
	return s.IsACREnterprise
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetIsACRRegistry() *bool {
	return s.IsACRRegistry
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetOsTag() *string {
	return s.OsTag
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetPlatform() *string {
	return s.Platform
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetRegistryCredential() *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	return s.RegistryCredential
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetRegistryCriId() *string {
	return s.RegistryCriId
}

func (s *GetImageResponseBodyImageContainerImageSpec) GetRegistryUrl() *string {
	return s.RegistryUrl
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetArchitecture(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.Architecture = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetIsACREnterprise(v bool) *GetImageResponseBodyImageContainerImageSpec {
	s.IsACREnterprise = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetIsACRRegistry(v bool) *GetImageResponseBodyImageContainerImageSpec {
	s.IsACRRegistry = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetOsTag(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.OsTag = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetPlatform(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.Platform = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetRegistryCredential(v *GetImageResponseBodyImageContainerImageSpecRegistryCredential) *GetImageResponseBodyImageContainerImageSpec {
	s.RegistryCredential = v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetRegistryCriId(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.RegistryCriId = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetRegistryUrl(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.RegistryUrl = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImageContainerImageSpecRegistryCredential struct {
	// example:
	//
	// userpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetImageResponseBodyImageContainerImageSpecRegistryCredential) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageContainerImageSpecRegistryCredential) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) GetPassword() *string {
	return s.Password
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) GetServer() *string {
	return s.Server
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) GetUserName() *string {
	return s.UserName
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) SetPassword(v string) *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	s.Password = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) SetServer(v string) *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	s.Server = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) SetUserName(v string) *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	s.UserName = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImageDocumentInfo struct {
	Document     *string `json:"Document,omitempty" xml:"Document,omitempty"`
	DocumentId   *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	EncodingMode *string `json:"EncodingMode,omitempty" xml:"EncodingMode,omitempty"`
}

func (s GetImageResponseBodyImageDocumentInfo) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageDocumentInfo) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageDocumentInfo) GetDocument() *string {
	return s.Document
}

func (s *GetImageResponseBodyImageDocumentInfo) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetImageResponseBodyImageDocumentInfo) GetEncodingMode() *string {
	return s.EncodingMode
}

func (s *GetImageResponseBodyImageDocumentInfo) SetDocument(v string) *GetImageResponseBodyImageDocumentInfo {
	s.Document = &v
	return s
}

func (s *GetImageResponseBodyImageDocumentInfo) SetDocumentId(v string) *GetImageResponseBodyImageDocumentInfo {
	s.DocumentId = &v
	return s
}

func (s *GetImageResponseBodyImageDocumentInfo) SetEncodingMode(v string) *GetImageResponseBodyImageDocumentInfo {
	s.EncodingMode = &v
	return s
}

func (s *GetImageResponseBodyImageDocumentInfo) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImageVMImageSpec struct {
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// m-uf60twafjtaart******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// CentOS  7.6 64 bit
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// example:
	//
	// CentOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s GetImageResponseBodyImageVMImageSpec) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageVMImageSpec) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageVMImageSpec) GetArchitecture() *string {
	return s.Architecture
}

func (s *GetImageResponseBodyImageVMImageSpec) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageResponseBodyImageVMImageSpec) GetOsTag() *string {
	return s.OsTag
}

func (s *GetImageResponseBodyImageVMImageSpec) GetPlatform() *string {
	return s.Platform
}

func (s *GetImageResponseBodyImageVMImageSpec) SetArchitecture(v string) *GetImageResponseBodyImageVMImageSpec {
	s.Architecture = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) SetImageId(v string) *GetImageResponseBodyImageVMImageSpec {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) SetOsTag(v string) *GetImageResponseBodyImageVMImageSpec {
	s.OsTag = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) SetPlatform(v string) *GetImageResponseBodyImageVMImageSpec {
	s.Platform = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) Validate() error {
	return dara.Validate(s)
}

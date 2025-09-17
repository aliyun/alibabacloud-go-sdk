// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrRegistryInfos(v []*CreateImageCacheRequestAcrRegistryInfos) *CreateImageCacheRequest
	GetAcrRegistryInfos() []*CreateImageCacheRequestAcrRegistryInfos
	SetClientToken(v string) *CreateImageCacheRequest
	GetClientToken() *string
	SetImageCacheName(v string) *CreateImageCacheRequest
	GetImageCacheName() *string
	SetImageRegistryCredentials(v []*CreateImageCacheRequestImageRegistryCredentials) *CreateImageCacheRequest
	GetImageRegistryCredentials() []*CreateImageCacheRequestImageRegistryCredentials
	SetImages(v []*string) *CreateImageCacheRequest
	GetImages() []*string
	SetNetworkConfig(v *CreateImageCacheRequestNetworkConfig) *CreateImageCacheRequest
	GetNetworkConfig() *CreateImageCacheRequestNetworkConfig
	SetRegionId(v string) *CreateImageCacheRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateImageCacheRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateImageCacheRequestTags) *CreateImageCacheRequest
	GetTags() []*CreateImageCacheRequestTags
}

type CreateImageCacheRequest struct {
	AcrRegistryInfos []*CreateImageCacheRequestAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	// example:
	//
	// *****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-imc
	ImageCacheName           *string                                            `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	ImageRegistryCredentials []*CreateImageCacheRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	// This parameter is required.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// This parameter is required.
	NetworkConfig *CreateImageCacheRequestNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*CreateImageCacheRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequest) GetAcrRegistryInfos() []*CreateImageCacheRequestAcrRegistryInfos {
	return s.AcrRegistryInfos
}

func (s *CreateImageCacheRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageCacheRequest) GetImageCacheName() *string {
	return s.ImageCacheName
}

func (s *CreateImageCacheRequest) GetImageRegistryCredentials() []*CreateImageCacheRequestImageRegistryCredentials {
	return s.ImageRegistryCredentials
}

func (s *CreateImageCacheRequest) GetImages() []*string {
	return s.Images
}

func (s *CreateImageCacheRequest) GetNetworkConfig() *CreateImageCacheRequestNetworkConfig {
	return s.NetworkConfig
}

func (s *CreateImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageCacheRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateImageCacheRequest) GetTags() []*CreateImageCacheRequestTags {
	return s.Tags
}

func (s *CreateImageCacheRequest) SetAcrRegistryInfos(v []*CreateImageCacheRequestAcrRegistryInfos) *CreateImageCacheRequest {
	s.AcrRegistryInfos = v
	return s
}

func (s *CreateImageCacheRequest) SetClientToken(v string) *CreateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheName(v string) *CreateImageCacheRequest {
	s.ImageCacheName = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageRegistryCredentials(v []*CreateImageCacheRequestImageRegistryCredentials) *CreateImageCacheRequest {
	s.ImageRegistryCredentials = v
	return s
}

func (s *CreateImageCacheRequest) SetImages(v []*string) *CreateImageCacheRequest {
	s.Images = v
	return s
}

func (s *CreateImageCacheRequest) SetNetworkConfig(v *CreateImageCacheRequestNetworkConfig) *CreateImageCacheRequest {
	s.NetworkConfig = v
	return s
}

func (s *CreateImageCacheRequest) SetRegionId(v string) *CreateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceGroupId(v string) *CreateImageCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetTags(v []*CreateImageCacheRequestTags) *CreateImageCacheRequest {
	s.Tags = v
	return s
}

func (s *CreateImageCacheRequest) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheRequestAcrRegistryInfos struct {
	// example:
	//
	// cri-nwj395hgf6f*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageCacheRequestAcrRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequestAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestAcrRegistryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageCacheRequestAcrRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageCacheRequestAcrRegistryInfos) SetInstanceId(v string) *CreateImageCacheRequestAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfos) SetRegionId(v string) *CreateImageCacheRequestAcrRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfos) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheRequestImageRegistryCredentials struct {
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// example:
	//
	// false
	SkipCertVerification *bool `json:"SkipCertVerification,omitempty" xml:"SkipCertVerification,omitempty"`
	// example:
	//
	// false
	UsePlainHttp *bool `json:"UsePlainHttp,omitempty" xml:"UsePlainHttp,omitempty"`
	// example:
	//
	// myusername
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateImageCacheRequestImageRegistryCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequestImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestImageRegistryCredentials) GetPassword() *string {
	return s.Password
}

func (s *CreateImageCacheRequestImageRegistryCredentials) GetServer() *string {
	return s.Server
}

func (s *CreateImageCacheRequestImageRegistryCredentials) GetSkipCertVerification() *bool {
	return s.SkipCertVerification
}

func (s *CreateImageCacheRequestImageRegistryCredentials) GetUsePlainHttp() *bool {
	return s.UsePlainHttp
}

func (s *CreateImageCacheRequestImageRegistryCredentials) GetUsername() *string {
	return s.Username
}

func (s *CreateImageCacheRequestImageRegistryCredentials) SetPassword(v string) *CreateImageCacheRequestImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredentials) SetServer(v string) *CreateImageCacheRequestImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredentials) SetSkipCertVerification(v bool) *CreateImageCacheRequestImageRegistryCredentials {
	s.SkipCertVerification = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredentials) SetUsePlainHttp(v bool) *CreateImageCacheRequestImageRegistryCredentials {
	s.UsePlainHttp = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredentials) SetUsername(v string) *CreateImageCacheRequestImageRegistryCredentials {
	s.Username = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredentials) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheRequestNetworkConfig struct {
	EipInstance *CreateImageCacheRequestNetworkConfigEipInstance `json:"EipInstance,omitempty" xml:"EipInstance,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// sg-0jlgektkddwa42n*****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateImageCacheRequestNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequestNetworkConfig) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestNetworkConfig) GetEipInstance() *CreateImageCacheRequestNetworkConfigEipInstance {
	return s.EipInstance
}

func (s *CreateImageCacheRequestNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateImageCacheRequestNetworkConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateImageCacheRequestNetworkConfig) SetEipInstance(v *CreateImageCacheRequestNetworkConfigEipInstance) *CreateImageCacheRequestNetworkConfig {
	s.EipInstance = v
	return s
}

func (s *CreateImageCacheRequestNetworkConfig) SetSecurityGroupId(v string) *CreateImageCacheRequestNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateImageCacheRequestNetworkConfig) SetVSwitchIds(v []*string) *CreateImageCacheRequestNetworkConfig {
	s.VSwitchIds = v
	return s
}

func (s *CreateImageCacheRequestNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheRequestNetworkConfigEipInstance struct {
	// example:
	//
	// true
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// example:
	//
	// 100
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// eip-0jl0bx3fnpnjc9i4*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateImageCacheRequestNetworkConfigEipInstance) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequestNetworkConfigEipInstance) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) GetAutoCreate() *bool {
	return s.AutoCreate
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) SetAutoCreate(v bool) *CreateImageCacheRequestNetworkConfigEipInstance {
	s.AutoCreate = &v
	return s
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) SetBandwidth(v int32) *CreateImageCacheRequestNetworkConfigEipInstance {
	s.Bandwidth = &v
	return s
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) SetInstanceId(v string) *CreateImageCacheRequestNetworkConfigEipInstance {
	s.InstanceId = &v
	return s
}

func (s *CreateImageCacheRequestNetworkConfigEipInstance) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheRequestTags struct {
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageCacheRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequestTags) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateImageCacheRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateImageCacheRequestTags) SetKey(v string) *CreateImageCacheRequestTags {
	s.Key = &v
	return s
}

func (s *CreateImageCacheRequestTags) SetValue(v string) *CreateImageCacheRequestTags {
	s.Value = &v
	return s
}

func (s *CreateImageCacheRequestTags) Validate() error {
	return dara.Validate(s)
}

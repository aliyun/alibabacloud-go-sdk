// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrRegistryInfos(v []*CreateImageCacheShrinkRequestAcrRegistryInfos) *CreateImageCacheShrinkRequest
	GetAcrRegistryInfos() []*CreateImageCacheShrinkRequestAcrRegistryInfos
	SetClientToken(v string) *CreateImageCacheShrinkRequest
	GetClientToken() *string
	SetImageCacheName(v string) *CreateImageCacheShrinkRequest
	GetImageCacheName() *string
	SetImageRegistryCredentials(v []*CreateImageCacheShrinkRequestImageRegistryCredentials) *CreateImageCacheShrinkRequest
	GetImageRegistryCredentials() []*CreateImageCacheShrinkRequestImageRegistryCredentials
	SetImages(v []*string) *CreateImageCacheShrinkRequest
	GetImages() []*string
	SetNetworkConfigShrink(v string) *CreateImageCacheShrinkRequest
	GetNetworkConfigShrink() *string
	SetRegionId(v string) *CreateImageCacheShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateImageCacheShrinkRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateImageCacheShrinkRequestTags) *CreateImageCacheShrinkRequest
	GetTags() []*CreateImageCacheShrinkRequestTags
}

type CreateImageCacheShrinkRequest struct {
	AcrRegistryInfos []*CreateImageCacheShrinkRequestAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	// example:
	//
	// *****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-imc
	ImageCacheName           *string                                                  `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	ImageRegistryCredentials []*CreateImageCacheShrinkRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	// This parameter is required.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// This parameter is required.
	NetworkConfigShrink *string `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*CreateImageCacheShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateImageCacheShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheShrinkRequest) GetAcrRegistryInfos() []*CreateImageCacheShrinkRequestAcrRegistryInfos {
	return s.AcrRegistryInfos
}

func (s *CreateImageCacheShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageCacheShrinkRequest) GetImageCacheName() *string {
	return s.ImageCacheName
}

func (s *CreateImageCacheShrinkRequest) GetImageRegistryCredentials() []*CreateImageCacheShrinkRequestImageRegistryCredentials {
	return s.ImageRegistryCredentials
}

func (s *CreateImageCacheShrinkRequest) GetImages() []*string {
	return s.Images
}

func (s *CreateImageCacheShrinkRequest) GetNetworkConfigShrink() *string {
	return s.NetworkConfigShrink
}

func (s *CreateImageCacheShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageCacheShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateImageCacheShrinkRequest) GetTags() []*CreateImageCacheShrinkRequestTags {
	return s.Tags
}

func (s *CreateImageCacheShrinkRequest) SetAcrRegistryInfos(v []*CreateImageCacheShrinkRequestAcrRegistryInfos) *CreateImageCacheShrinkRequest {
	s.AcrRegistryInfos = v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetClientToken(v string) *CreateImageCacheShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetImageCacheName(v string) *CreateImageCacheShrinkRequest {
	s.ImageCacheName = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetImageRegistryCredentials(v []*CreateImageCacheShrinkRequestImageRegistryCredentials) *CreateImageCacheShrinkRequest {
	s.ImageRegistryCredentials = v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetImages(v []*string) *CreateImageCacheShrinkRequest {
	s.Images = v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetNetworkConfigShrink(v string) *CreateImageCacheShrinkRequest {
	s.NetworkConfigShrink = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetRegionId(v string) *CreateImageCacheShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetResourceGroupId(v string) *CreateImageCacheShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetTags(v []*CreateImageCacheShrinkRequestTags) *CreateImageCacheShrinkRequest {
	s.Tags = v
	return s
}

func (s *CreateImageCacheShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheShrinkRequestAcrRegistryInfos struct {
	// example:
	//
	// cri-nwj395hgf6f*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageCacheShrinkRequestAcrRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheShrinkRequestAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *CreateImageCacheShrinkRequestAcrRegistryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageCacheShrinkRequestAcrRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageCacheShrinkRequestAcrRegistryInfos) SetInstanceId(v string) *CreateImageCacheShrinkRequestAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *CreateImageCacheShrinkRequestAcrRegistryInfos) SetRegionId(v string) *CreateImageCacheShrinkRequestAcrRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheShrinkRequestAcrRegistryInfos) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheShrinkRequestImageRegistryCredentials struct {
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
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateImageCacheShrinkRequestImageRegistryCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheShrinkRequestImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) GetPassword() *string {
	return s.Password
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) GetServer() *string {
	return s.Server
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) GetSkipCertVerification() *bool {
	return s.SkipCertVerification
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) GetUsePlainHttp() *bool {
	return s.UsePlainHttp
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) GetUserName() *string {
	return s.UserName
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) SetPassword(v string) *CreateImageCacheShrinkRequestImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) SetServer(v string) *CreateImageCacheShrinkRequestImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) SetSkipCertVerification(v bool) *CreateImageCacheShrinkRequestImageRegistryCredentials {
	s.SkipCertVerification = &v
	return s
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) SetUsePlainHttp(v bool) *CreateImageCacheShrinkRequestImageRegistryCredentials {
	s.UsePlainHttp = &v
	return s
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) SetUserName(v string) *CreateImageCacheShrinkRequestImageRegistryCredentials {
	s.UserName = &v
	return s
}

func (s *CreateImageCacheShrinkRequestImageRegistryCredentials) Validate() error {
	return dara.Validate(s)
}

type CreateImageCacheShrinkRequestTags struct {
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageCacheShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateImageCacheShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateImageCacheShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateImageCacheShrinkRequestTags) SetKey(v string) *CreateImageCacheShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateImageCacheShrinkRequestTags) SetValue(v string) *CreateImageCacheShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *CreateImageCacheShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}

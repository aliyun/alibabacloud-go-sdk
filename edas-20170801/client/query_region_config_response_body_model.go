// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryRegionConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryRegionConfigResponseBody
	GetMessage() *string
	SetRegionConfig(v *QueryRegionConfigResponseBodyRegionConfig) *QueryRegionConfigResponseBody
	GetRegionConfig() *QueryRegionConfigResponseBodyRegionConfig
	SetRequestId(v string) *QueryRegionConfigResponseBody
	GetRequestId() *string
}

type QueryRegionConfigResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about region configurations.
	RegionConfig *QueryRegionConfigResponseBodyRegionConfig `json:"RegionConfig,omitempty" xml:"RegionConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRegionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRegionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryRegionConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRegionConfigResponseBody) GetRegionConfig() *QueryRegionConfigResponseBodyRegionConfig {
	return s.RegionConfig
}

func (s *QueryRegionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRegionConfigResponseBody) SetCode(v int32) *QueryRegionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRegionConfigResponseBody) SetMessage(v string) *QueryRegionConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRegionConfigResponseBody) SetRegionConfig(v *QueryRegionConfigResponseBodyRegionConfig) *QueryRegionConfigResponseBody {
	s.RegionConfig = v
	return s
}

func (s *QueryRegionConfigResponseBody) SetRequestId(v string) *QueryRegionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegionConfigResponseBody) Validate() error {
	if s.RegionConfig != nil {
		if err := s.RegionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRegionConfigResponseBodyRegionConfig struct {
	// The domain name of Address Server.
	//
	// example:
	//
	// ****.edas.aliyun.com
	AddressServerHost *string `json:"AddressServerHost,omitempty" xml:"AddressServerHost,omitempty"`
	// The installation path of the script for EDAS Agent.
	//
	// example:
	//
	// http://edas-qd.oss-cn-****-internal.aliyuncs.com/****sh
	AgentInstallScript *string `json:"AgentInstallScript,omitempty" xml:"AgentInstallScript,omitempty"`
	// The information about the file server.
	FileServerConfig *QueryRegionConfigResponseBodyRegionConfigFileServerConfig `json:"FileServerConfig,omitempty" xml:"FileServerConfig,omitempty" type:"Struct"`
	// The type of the file server.
	//
	// example:
	//
	// oss
	FileServerType *string `json:"FileServerType,omitempty" xml:"FileServerType,omitempty"`
	// The configured ID of the region.
	//
	// example:
	//
	// cn-beijing
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the official image.
	//
	// example:
	//
	// m-2zea4hx8f9zxqah2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The configured name of the region.
	//
	// example:
	//
	// China (Beijing)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The serial number of the region. This parameter is deprecated.
	//
	// example:
	//
	// 1
	No *int32 `json:"No,omitempty" xml:"No,omitempty"`
	// The tag of the region. The value is fixed to `ALIYUN_SHARE`.
	//
	// example:
	//
	// ALIYUN_SHARE
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s QueryRegionConfigResponseBodyRegionConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryRegionConfigResponseBodyRegionConfig) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetAddressServerHost() *string {
	return s.AddressServerHost
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetAgentInstallScript() *string {
	return s.AgentInstallScript
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetFileServerConfig() *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	return s.FileServerConfig
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetFileServerType() *string {
	return s.FileServerType
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetId() *string {
	return s.Id
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetImageId() *string {
	return s.ImageId
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetName() *string {
	return s.Name
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetNo() *int32 {
	return s.No
}

func (s *QueryRegionConfigResponseBodyRegionConfig) GetTag() *string {
	return s.Tag
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetAddressServerHost(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.AddressServerHost = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetAgentInstallScript(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.AgentInstallScript = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetFileServerConfig(v *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) *QueryRegionConfigResponseBodyRegionConfig {
	s.FileServerConfig = v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetFileServerType(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.FileServerType = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetId(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.Id = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetImageId(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.ImageId = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetName(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.Name = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetNo(v int32) *QueryRegionConfigResponseBodyRegionConfig {
	s.No = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetTag(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.Tag = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) Validate() error {
	if s.FileServerConfig != nil {
		if err := s.FileServerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRegionConfigResponseBodyRegionConfigFileServerConfig struct {
	// The Object Storage Service (OSS) bucket of the file server.
	//
	// example:
	//
	// edas-bj
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The internal endpoint of the file server.
	//
	// example:
	//
	// oss-cn-beijing-****.aliyuncs.com
	InternalUrl *string `json:"InternalUrl,omitempty" xml:"InternalUrl,omitempty"`
	// The public endpoint of the file server.
	//
	// example:
	//
	// oss-cn-beijing.aliyuncs.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// The virtual private cloud (VPC) endpoint of the file server.
	//
	// example:
	//
	// v*****-oss-cn-beijing.aliyuncs.com
	VpcUrl *string `json:"VpcUrl,omitempty" xml:"VpcUrl,omitempty"`
}

func (s QueryRegionConfigResponseBodyRegionConfigFileServerConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryRegionConfigResponseBodyRegionConfigFileServerConfig) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) GetBucket() *string {
	return s.Bucket
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) GetInternalUrl() *string {
	return s.InternalUrl
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) GetVpcUrl() *string {
	return s.VpcUrl
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetBucket(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.Bucket = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetInternalUrl(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.InternalUrl = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetPublicUrl(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.PublicUrl = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetVpcUrl(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.VpcUrl = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) Validate() error {
	return dara.Validate(s)
}

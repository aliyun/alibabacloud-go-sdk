// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJenkinsImageRegistryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CreateJenkinsImageRegistryRequest
	GetDomainName() *string
	SetExtraParam(v string) *CreateJenkinsImageRegistryRequest
	GetExtraParam() *string
	SetNetType(v int32) *CreateJenkinsImageRegistryRequest
	GetNetType() *int32
	SetPassword(v string) *CreateJenkinsImageRegistryRequest
	GetPassword() *string
	SetPersistenceDay(v int32) *CreateJenkinsImageRegistryRequest
	GetPersistenceDay() *int32
	SetProtocolType(v int32) *CreateJenkinsImageRegistryRequest
	GetProtocolType() *int32
	SetRegionId(v string) *CreateJenkinsImageRegistryRequest
	GetRegionId() *string
	SetRegistryHostIp(v string) *CreateJenkinsImageRegistryRequest
	GetRegistryHostIp() *string
	SetRegistryName(v string) *CreateJenkinsImageRegistryRequest
	GetRegistryName() *string
	SetRegistryType(v string) *CreateJenkinsImageRegistryRequest
	GetRegistryType() *string
	SetRegistryVersion(v string) *CreateJenkinsImageRegistryRequest
	GetRegistryVersion() *string
	SetSourceIp(v string) *CreateJenkinsImageRegistryRequest
	GetSourceIp() *string
	SetTransPerHour(v int32) *CreateJenkinsImageRegistryRequest
	GetTransPerHour() *int32
	SetUserName(v string) *CreateJenkinsImageRegistryRequest
	GetUserName() *string
	SetVpcId(v string) *CreateJenkinsImageRegistryRequest
	GetVpcId() *string
	SetWhiteList(v string) *CreateJenkinsImageRegistryRequest
	GetWhiteList() *string
}

type CreateJenkinsImageRegistryRequest struct {
	// The domain name of the image repository.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The additional parameters of the image repository. The value of this parameter contains the following fields:
	//
	// 	- **namespace**: the namespace
	//
	// 	- **authToken**: the authorization token
	//
	// example:
	//
	// [{\\"namespace\\":\\"aa\\",\\"authToken\\":\\"aa\\"}]
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The network type. Valid values:
	//
	// 	- **1**: Internet
	//
	// 	- **2**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// 1
	NetType *int32 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The password.
	//
	// example:
	//
	// Harbor********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The number of days during which assets can be retained.
	//
	// example:
	//
	// 30
	PersistenceDay *int32 `json:"PersistenceDay,omitempty" xml:"PersistenceDay,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **1**: HTTP
	//
	// 	- **2**: HTTPS
	//
	// example:
	//
	// 1
	ProtocolType *int32 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address of the image repository.
	//
	// example:
	//
	// 114.55.XXX.XXX
	RegistryHostIp *string `json:"RegistryHostIp,omitempty" xml:"RegistryHostIp,omitempty"`
	// The alias of the image repository.
	//
	// example:
	//
	// testRepo
	RegistryName *string `json:"RegistryName,omitempty" xml:"RegistryName,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- **CI/CD**: Jenkins
	//
	// example:
	//
	// CI/CD
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The version of the image repository. Default value: -. Valid values:
	//
	// 	- **-**: the default version
	//
	// 	- **V1**: V1.0
	//
	// 	- **V2**: V2.0
	//
	// example:
	//
	// V1
	RegistryVersion *string `json:"RegistryVersion,omitempty" xml:"RegistryVersion,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 41.121.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The number of images that can be scanned per hour.
	//
	// example:
	//
	// 30
	TransPerHour *int32 `json:"TransPerHour,omitempty" xml:"TransPerHour,omitempty"`
	// The username.
	//
	// example:
	//
	// RegistryUser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2ze4aoqgeu51ydfb8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The whitelist of IP addresses.
	//
	// example:
	//
	// 192.168.XXX.XXX
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s CreateJenkinsImageRegistryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageRegistryRequest) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageRegistryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateJenkinsImageRegistryRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *CreateJenkinsImageRegistryRequest) GetNetType() *int32 {
	return s.NetType
}

func (s *CreateJenkinsImageRegistryRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateJenkinsImageRegistryRequest) GetPersistenceDay() *int32 {
	return s.PersistenceDay
}

func (s *CreateJenkinsImageRegistryRequest) GetProtocolType() *int32 {
	return s.ProtocolType
}

func (s *CreateJenkinsImageRegistryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateJenkinsImageRegistryRequest) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *CreateJenkinsImageRegistryRequest) GetRegistryName() *string {
	return s.RegistryName
}

func (s *CreateJenkinsImageRegistryRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *CreateJenkinsImageRegistryRequest) GetRegistryVersion() *string {
	return s.RegistryVersion
}

func (s *CreateJenkinsImageRegistryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateJenkinsImageRegistryRequest) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *CreateJenkinsImageRegistryRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateJenkinsImageRegistryRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateJenkinsImageRegistryRequest) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateJenkinsImageRegistryRequest) SetDomainName(v string) *CreateJenkinsImageRegistryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetExtraParam(v string) *CreateJenkinsImageRegistryRequest {
	s.ExtraParam = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetNetType(v int32) *CreateJenkinsImageRegistryRequest {
	s.NetType = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetPassword(v string) *CreateJenkinsImageRegistryRequest {
	s.Password = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetPersistenceDay(v int32) *CreateJenkinsImageRegistryRequest {
	s.PersistenceDay = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetProtocolType(v int32) *CreateJenkinsImageRegistryRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetRegionId(v string) *CreateJenkinsImageRegistryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetRegistryHostIp(v string) *CreateJenkinsImageRegistryRequest {
	s.RegistryHostIp = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetRegistryName(v string) *CreateJenkinsImageRegistryRequest {
	s.RegistryName = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetRegistryType(v string) *CreateJenkinsImageRegistryRequest {
	s.RegistryType = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetRegistryVersion(v string) *CreateJenkinsImageRegistryRequest {
	s.RegistryVersion = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetSourceIp(v string) *CreateJenkinsImageRegistryRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetTransPerHour(v int32) *CreateJenkinsImageRegistryRequest {
	s.TransPerHour = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetUserName(v string) *CreateJenkinsImageRegistryRequest {
	s.UserName = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetVpcId(v string) *CreateJenkinsImageRegistryRequest {
	s.VpcId = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) SetWhiteList(v string) *CreateJenkinsImageRegistryRequest {
	s.WhiteList = &v
	return s
}

func (s *CreateJenkinsImageRegistryRequest) Validate() error {
	return dara.Validate(s)
}

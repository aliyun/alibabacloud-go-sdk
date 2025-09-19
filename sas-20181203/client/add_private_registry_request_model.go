// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivateRegistryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddPrivateRegistryRequest
	GetDomainName() *string
	SetExtraParam(v string) *AddPrivateRegistryRequest
	GetExtraParam() *string
	SetNetType(v int64) *AddPrivateRegistryRequest
	GetNetType() *int64
	SetPassword(v string) *AddPrivateRegistryRequest
	GetPassword() *string
	SetPort(v int32) *AddPrivateRegistryRequest
	GetPort() *int32
	SetProtocolType(v int64) *AddPrivateRegistryRequest
	GetProtocolType() *int64
	SetRegistryHostIp(v string) *AddPrivateRegistryRequest
	GetRegistryHostIp() *string
	SetRegistryRegionId(v string) *AddPrivateRegistryRequest
	GetRegistryRegionId() *string
	SetRegistryType(v string) *AddPrivateRegistryRequest
	GetRegistryType() *string
	SetRegistryVersion(v string) *AddPrivateRegistryRequest
	GetRegistryVersion() *string
	SetTransPerHour(v int32) *AddPrivateRegistryRequest
	GetTransPerHour() *int32
	SetUserName(v string) *AddPrivateRegistryRequest
	GetUserName() *string
	SetVpcId(v string) *AddPrivateRegistryRequest
	GetVpcId() *string
}

type AddPrivateRegistryRequest struct {
	// The domain name of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The additional parameter of the image repository. This parameter is required when you set the RegistryType parameter to **quay**. Valid values:
	//
	// 	- **namespace**
	//
	// 	- **authToken**
	//
	// example:
	//
	// [{"namespace":"aa","authToken":"aa"}]
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The network type. Valid values:
	//
	// 	- **1**: Internet
	//
	// 	- **2**: virtual private cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NetType *int64 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The password that is used to log on to the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **1**: HTTP
	//
	// 	- **2**: HTTPS
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProtocolType *int64 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The IP address of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// ``114.55.**.**``
	RegistryHostIp *string `json:"RegistryHostIp,omitempty" xml:"RegistryHostIp,omitempty"`
	// The region ID.
	//
	// >  You can call the [ListImageRegistryRegion](~~ListImageRegistryRegion~~) operation to query the IDs of supported regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegistryRegionId *string `json:"RegistryRegionId,omitempty" xml:"RegistryRegionId,omitempty"`
	// The type of the private image repository. Valid values:
	//
	// 	- **harbor**
	//
	// 	- **quay**
	//
	// This parameter is required.
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The version of the image repository. Valid values:
	//
	// 	- **V1**
	//
	// 	- **V2**
	//
	// This parameter is required.
	//
	// example:
	//
	// V2
	RegistryVersion *string `json:"RegistryVersion,omitempty" xml:"RegistryVersion,omitempty"`
	// The number of images that are scanned per hour.
	//
	// example:
	//
	// 10
	TransPerHour *int32 `json:"TransPerHour,omitempty" xml:"TransPerHour,omitempty"`
	// The username that is used to log on to the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-wz9hs3e5*******908kd
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddPrivateRegistryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateRegistryRequest) GoString() string {
	return s.String()
}

func (s *AddPrivateRegistryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddPrivateRegistryRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *AddPrivateRegistryRequest) GetNetType() *int64 {
	return s.NetType
}

func (s *AddPrivateRegistryRequest) GetPassword() *string {
	return s.Password
}

func (s *AddPrivateRegistryRequest) GetPort() *int32 {
	return s.Port
}

func (s *AddPrivateRegistryRequest) GetProtocolType() *int64 {
	return s.ProtocolType
}

func (s *AddPrivateRegistryRequest) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *AddPrivateRegistryRequest) GetRegistryRegionId() *string {
	return s.RegistryRegionId
}

func (s *AddPrivateRegistryRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *AddPrivateRegistryRequest) GetRegistryVersion() *string {
	return s.RegistryVersion
}

func (s *AddPrivateRegistryRequest) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *AddPrivateRegistryRequest) GetUserName() *string {
	return s.UserName
}

func (s *AddPrivateRegistryRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddPrivateRegistryRequest) SetDomainName(v string) *AddPrivateRegistryRequest {
	s.DomainName = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetExtraParam(v string) *AddPrivateRegistryRequest {
	s.ExtraParam = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetNetType(v int64) *AddPrivateRegistryRequest {
	s.NetType = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetPassword(v string) *AddPrivateRegistryRequest {
	s.Password = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetPort(v int32) *AddPrivateRegistryRequest {
	s.Port = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetProtocolType(v int64) *AddPrivateRegistryRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetRegistryHostIp(v string) *AddPrivateRegistryRequest {
	s.RegistryHostIp = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetRegistryRegionId(v string) *AddPrivateRegistryRequest {
	s.RegistryRegionId = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetRegistryType(v string) *AddPrivateRegistryRequest {
	s.RegistryType = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetRegistryVersion(v string) *AddPrivateRegistryRequest {
	s.RegistryVersion = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetTransPerHour(v int32) *AddPrivateRegistryRequest {
	s.TransPerHour = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetUserName(v string) *AddPrivateRegistryRequest {
	s.UserName = &v
	return s
}

func (s *AddPrivateRegistryRequest) SetVpcId(v string) *AddPrivateRegistryRequest {
	s.VpcId = &v
	return s
}

func (s *AddPrivateRegistryRequest) Validate() error {
	return dara.Validate(s)
}

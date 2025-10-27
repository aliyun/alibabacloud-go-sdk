// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivateRegistryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddPrivateRegistryResponseBodyData) *AddPrivateRegistryResponseBody
	GetData() *AddPrivateRegistryResponseBodyData
	SetRequestId(v string) *AddPrivateRegistryResponseBody
	GetRequestId() *string
}

type AddPrivateRegistryResponseBody struct {
	// The handling result.
	Data *AddPrivateRegistryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrivateRegistryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrivateRegistryResponseBody) GetData() *AddPrivateRegistryResponseBodyData {
	return s.Data
}

func (s *AddPrivateRegistryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrivateRegistryResponseBody) SetData(v *AddPrivateRegistryResponseBodyData) *AddPrivateRegistryResponseBody {
	s.Data = v
	return s
}

func (s *AddPrivateRegistryResponseBody) SetRequestId(v string) *AddPrivateRegistryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrivateRegistryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddPrivateRegistryResponseBodyData struct {
	// The ID of the user.
	//
	// example:
	//
	// 1766185894104***
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The domain name of the image repository.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// 273698***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The network type. Valid values:
	//
	// 	- **1**: Internet
	//
	// 	- **2**: VPC
	//
	// example:
	//
	// 2
	NetType *int64 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The password.
	//
	// example:
	//
	// ***********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **1**: HTTP
	//
	// 	- **2**: HTTPS
	//
	// example:
	//
	// 2
	ProtocolType *int64 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address of the image repository.
	//
	// example:
	//
	// ``114.55.**.**``
	RegistryHostIp *string `json:"RegistryHostIp,omitempty" xml:"RegistryHostIp,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- **harbor**
	//
	// 	- **quay**
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The version of the image repository. Default value: -. Valid values:
	//
	// 	- **-**: the default version
	//
	// 	- **V1**
	//
	// 	- **V2**
	//
	// example:
	//
	// V2
	RegistryVersion *string `json:"RegistryVersion,omitempty" xml:"RegistryVersion,omitempty"`
	// The value of the token.
	//
	// example:
	//
	// 3c3c602c-fa1f-4bc0-992f-b4b2cac7****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The number of scan tasks that are performed per hour.
	//
	// example:
	//
	// 10
	TransPerHour *int32 `json:"TransPerHour,omitempty" xml:"TransPerHour,omitempty"`
	// The username.
	//
	// example:
	//
	// ******
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2vchkxmf2j9yjt3x2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddPrivateRegistryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateRegistryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddPrivateRegistryResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *AddPrivateRegistryResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *AddPrivateRegistryResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *AddPrivateRegistryResponseBodyData) GetNetType() *int64 {
	return s.NetType
}

func (s *AddPrivateRegistryResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *AddPrivateRegistryResponseBodyData) GetProtocolType() *int64 {
	return s.ProtocolType
}

func (s *AddPrivateRegistryResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrivateRegistryResponseBodyData) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *AddPrivateRegistryResponseBodyData) GetRegistryType() *string {
	return s.RegistryType
}

func (s *AddPrivateRegistryResponseBodyData) GetRegistryVersion() *string {
	return s.RegistryVersion
}

func (s *AddPrivateRegistryResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *AddPrivateRegistryResponseBodyData) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *AddPrivateRegistryResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *AddPrivateRegistryResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *AddPrivateRegistryResponseBodyData) SetAliUid(v int64) *AddPrivateRegistryResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetDomainName(v string) *AddPrivateRegistryResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetId(v int64) *AddPrivateRegistryResponseBodyData {
	s.Id = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetNetType(v int64) *AddPrivateRegistryResponseBodyData {
	s.NetType = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetPassword(v string) *AddPrivateRegistryResponseBodyData {
	s.Password = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetProtocolType(v int64) *AddPrivateRegistryResponseBodyData {
	s.ProtocolType = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetRegionId(v string) *AddPrivateRegistryResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetRegistryHostIp(v string) *AddPrivateRegistryResponseBodyData {
	s.RegistryHostIp = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetRegistryType(v string) *AddPrivateRegistryResponseBodyData {
	s.RegistryType = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetRegistryVersion(v string) *AddPrivateRegistryResponseBodyData {
	s.RegistryVersion = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetToken(v string) *AddPrivateRegistryResponseBodyData {
	s.Token = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetTransPerHour(v int32) *AddPrivateRegistryResponseBodyData {
	s.TransPerHour = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetUserName(v string) *AddPrivateRegistryResponseBodyData {
	s.UserName = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) SetVpcId(v string) *AddPrivateRegistryResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *AddPrivateRegistryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

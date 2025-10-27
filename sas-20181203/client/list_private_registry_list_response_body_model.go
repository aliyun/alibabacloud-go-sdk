// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateRegistryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageRegistryInfos(v []*ListPrivateRegistryListResponseBodyImageRegistryInfos) *ListPrivateRegistryListResponseBody
	GetImageRegistryInfos() []*ListPrivateRegistryListResponseBodyImageRegistryInfos
	SetRequestId(v string) *ListPrivateRegistryListResponseBody
	GetRequestId() *string
}

type ListPrivateRegistryListResponseBody struct {
	// An array that consists of the image repositories.
	ImageRegistryInfos []*ListPrivateRegistryListResponseBodyImageRegistryInfos `json:"ImageRegistryInfos,omitempty" xml:"ImageRegistryInfos,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateRegistryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryListResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryListResponseBody) GetImageRegistryInfos() []*ListPrivateRegistryListResponseBodyImageRegistryInfos {
	return s.ImageRegistryInfos
}

func (s *ListPrivateRegistryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateRegistryListResponseBody) SetImageRegistryInfos(v []*ListPrivateRegistryListResponseBodyImageRegistryInfos) *ListPrivateRegistryListResponseBody {
	s.ImageRegistryInfos = v
	return s
}

func (s *ListPrivateRegistryListResponseBody) SetRequestId(v string) *ListPrivateRegistryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateRegistryListResponseBody) Validate() error {
	if s.ImageRegistryInfos != nil {
		for _, item := range s.ImageRegistryInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateRegistryListResponseBodyImageRegistryInfos struct {
	// The ID of the user.
	//
	// example:
	//
	// 1766185894******
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
	// 66485
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the Jenkins environment.
	//
	// example:
	//
	// JenkinsInfo
	JenkinsEnv *string `json:"JenkinsEnv,omitempty" xml:"JenkinsEnv,omitempty"`
	// The network type. Valid values:
	//
	// 	- **1**: Internet
	//
	// 	- **2**: VPC
	//
	// example:
	//
	// 1
	NetType *int64 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The password used to log on to the image repository.
	//
	// example:
	//
	// Harbor******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The number of days during which assets can be retained.
	//
	// example:
	//
	// 90
	PersistenceDay *int64 `json:"PersistenceDay,omitempty" xml:"PersistenceDay,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **1**: HTTP
	//
	// 	- **2**: HTTPS
	//
	// example:
	//
	// 1
	ProtocolType *int64 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID of the server.
	//
	// example:
	//
	// cn-hangzhou
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
	// test1
	RegistryName *string `json:"RegistryName,omitempty" xml:"RegistryName,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- **acr**: Container Registry
	//
	// 	- **harbor**: Harbor
	//
	// 	- **quay**: Quay
	//
	// 	- **CI/CD**: Jenkins
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The version of the image repository. Valid values:
	//
	// 	- **V1**: V1.0
	//
	// 	- **V2**: V2.0
	//
	// example:
	//
	// V1
	RegistryVersion *string `json:"RegistryVersion,omitempty" xml:"RegistryVersion,omitempty"`
	// The authentication token of the user.
	//
	// example:
	//
	// 0da12bce-cc36-4c48-b3e6-2215fc3a****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The number of images that can be scanned per hour.
	//
	// example:
	//
	// 30
	TransPerHour *int32 `json:"TransPerHour,omitempty" xml:"TransPerHour,omitempty"`
	// The username used to log on to the image repository.
	//
	// example:
	//
	// RegistryUser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp12897gqrex01zn0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The whitelist of IP addresses.
	//
	// example:
	//
	// 100.104.XXX.XXX
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s ListPrivateRegistryListResponseBodyImageRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryListResponseBodyImageRegistryInfos) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetDomainName() *string {
	return s.DomainName
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetId() *int64 {
	return s.Id
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetJenkinsEnv() *string {
	return s.JenkinsEnv
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetNetType() *int64 {
	return s.NetType
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetPassword() *string {
	return s.Password
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetPersistenceDay() *int64 {
	return s.PersistenceDay
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetProtocolType() *int64 {
	return s.ProtocolType
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetRegistryName() *string {
	return s.RegistryName
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetRegistryVersion() *string {
	return s.RegistryVersion
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetToken() *string {
	return s.Token
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetUserName() *string {
	return s.UserName
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) GetWhiteList() *string {
	return s.WhiteList
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetAliUid(v int64) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.AliUid = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetDomainName(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.DomainName = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetId(v int64) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.Id = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetJenkinsEnv(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.JenkinsEnv = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetNetType(v int64) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.NetType = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetPassword(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.Password = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetPersistenceDay(v int64) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.PersistenceDay = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetProtocolType(v int64) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.ProtocolType = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetRegionId(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetRegistryHostIp(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.RegistryHostIp = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetRegistryName(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.RegistryName = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetRegistryType(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.RegistryType = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetRegistryVersion(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.RegistryVersion = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetToken(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.Token = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetTransPerHour(v int32) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.TransPerHour = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetUserName(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.UserName = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetVpcId(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.VpcId = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) SetWhiteList(v string) *ListPrivateRegistryListResponseBodyImageRegistryInfos {
	s.WhiteList = &v
	return s
}

func (s *ListPrivateRegistryListResponseBodyImageRegistryInfos) Validate() error {
	return dara.Validate(s)
}

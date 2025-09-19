// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJenkinsImageRegistryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateJenkinsImageRegistryResponseBodyData) *CreateJenkinsImageRegistryResponseBody
	GetData() *CreateJenkinsImageRegistryResponseBodyData
	SetHttpStatusCode(v int32) *CreateJenkinsImageRegistryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateJenkinsImageRegistryResponseBody
	GetRequestId() *string
	SetTimeCost(v int64) *CreateJenkinsImageRegistryResponseBody
	GetTimeCost() *int64
}

type CreateJenkinsImageRegistryResponseBody struct {
	// The result of creating the image repository.
	Data *CreateJenkinsImageRegistryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AF1E723-53F1-55BF-A4B2-15CB7A32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time consumed. Unit: seconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s CreateJenkinsImageRegistryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageRegistryResponseBody) GetData() *CreateJenkinsImageRegistryResponseBodyData {
	return s.Data
}

func (s *CreateJenkinsImageRegistryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateJenkinsImageRegistryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJenkinsImageRegistryResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *CreateJenkinsImageRegistryResponseBody) SetData(v *CreateJenkinsImageRegistryResponseBodyData) *CreateJenkinsImageRegistryResponseBody {
	s.Data = v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBody) SetHttpStatusCode(v int32) *CreateJenkinsImageRegistryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBody) SetRequestId(v string) *CreateJenkinsImageRegistryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBody) SetTimeCost(v int64) *CreateJenkinsImageRegistryResponseBody {
	s.TimeCost = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateJenkinsImageRegistryResponseBodyData struct {
	// The blacklist.
	//
	// example:
	//
	// 61.9.XXX.XXX
	BlackList *string `json:"BlackList,omitempty" xml:"BlackList,omitempty"`
	// The domain name of the image repository.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The creation time. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2022-10-16 18:17:16
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2022-11-21 10:40:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// 443496
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The network type. Valid values:
	//
	// 	- **1**: Internet
	//
	// 	- **2**: VPC
	//
	// example:
	//
	// 1
	NetType *int32 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The password.
	//
	// example:
	//
	// Harbor******
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
	// 1.13.XXX.XXX
	RegistryHostIp *string `json:"RegistryHostIp,omitempty" xml:"RegistryHostIp,omitempty"`
	// The alias of the image repository.
	//
	// example:
	//
	// fanyi
	RegistryName *string `json:"RegistryName,omitempty" xml:"RegistryName,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- **CI/CD**: Jenkins
	//
	// example:
	//
	// CI/CD
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The authentication token of the user.
	//
	// example:
	//
	// 3c3c602c-fa1f-4bc0-992f-b4b2cac7****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	// vpc-2vchkxmf2j9yjt3x2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The whitelist.
	//
	// example:
	//
	// 192.168.XXX.XXX
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s CreateJenkinsImageRegistryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageRegistryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetBlackList() *string {
	return s.BlackList
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetNetType() *int32 {
	return s.NetType
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetPersistenceDay() *int32 {
	return s.PersistenceDay
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetProtocolType() *int32 {
	return s.ProtocolType
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetRegistryName() *string {
	return s.RegistryName
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetRegistryType() *string {
	return s.RegistryType
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateJenkinsImageRegistryResponseBodyData) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetBlackList(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.BlackList = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetDomainName(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetGmtCreate(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetGmtModified(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetId(v int64) *CreateJenkinsImageRegistryResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetNetType(v int32) *CreateJenkinsImageRegistryResponseBodyData {
	s.NetType = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetPassword(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.Password = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetPersistenceDay(v int32) *CreateJenkinsImageRegistryResponseBodyData {
	s.PersistenceDay = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetProtocolType(v int32) *CreateJenkinsImageRegistryResponseBodyData {
	s.ProtocolType = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetRegionId(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetRegistryHostIp(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.RegistryHostIp = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetRegistryName(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.RegistryName = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetRegistryType(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.RegistryType = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetToken(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.Token = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetTransPerHour(v int32) *CreateJenkinsImageRegistryResponseBodyData {
	s.TransPerHour = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetUserName(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.UserName = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetVpcId(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) SetWhiteList(v string) *CreateJenkinsImageRegistryResponseBodyData {
	s.WhiteList = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

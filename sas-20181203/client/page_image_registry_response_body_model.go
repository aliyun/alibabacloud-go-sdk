// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageImageRegistryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*PageImageRegistryResponseBodyList) *PageImageRegistryResponseBody
	GetList() []*PageImageRegistryResponseBodyList
	SetPageInfo(v *PageImageRegistryResponseBodyPageInfo) *PageImageRegistryResponseBody
	GetPageInfo() *PageImageRegistryResponseBodyPageInfo
	SetRequestId(v string) *PageImageRegistryResponseBody
	GetRequestId() *string
}

type PageImageRegistryResponseBody struct {
	// The image repository list data.
	List []*PageImageRegistryResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information for the paged query.
	PageInfo *PageImageRegistryResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// FDA9E37C-6114-5945-8FF1-E3D4D397****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PageImageRegistryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageImageRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *PageImageRegistryResponseBody) GetList() []*PageImageRegistryResponseBodyList {
	return s.List
}

func (s *PageImageRegistryResponseBody) GetPageInfo() *PageImageRegistryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *PageImageRegistryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageImageRegistryResponseBody) SetList(v []*PageImageRegistryResponseBodyList) *PageImageRegistryResponseBody {
	s.List = v
	return s
}

func (s *PageImageRegistryResponseBody) SetPageInfo(v *PageImageRegistryResponseBodyPageInfo) *PageImageRegistryResponseBody {
	s.PageInfo = v
	return s
}

func (s *PageImageRegistryResponseBody) SetRequestId(v string) *PageImageRegistryResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageImageRegistryResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PageImageRegistryResponseBodyList struct {
	// The blacklist.
	//
	// example:
	//
	// 129.211.XXX.XXX
	BlackList *string `json:"BlackList,omitempty" xml:"BlackList,omitempty"`
	// The domain name of the repository.
	//
	// example:
	//
	// sinochem.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The creation time, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2022-08-30 10:23:30
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2022-09-30 10:23:30
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique ID of the image repository.
	//
	// example:
	//
	// 1078312
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of images in the repository.
	//
	// example:
	//
	// 1
	ImageCount *int32 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// The Jenkins environment context.
	//
	// example:
	//
	// projectInfo
	JenkinsEnv *string `json:"JenkinsEnv,omitempty" xml:"JenkinsEnv,omitempty"`
	// The network type. Valid values:
	//
	// - **1**: public network
	//
	// - **2**: VPC.
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
	// The number of days that assets are retained.
	//
	// example:
	//
	// 30
	PersistenceDay *int32 `json:"PersistenceDay,omitempty" xml:"PersistenceDay,omitempty"`
	// The protocol type. Valid values:
	//
	// - **1**: HTTP
	//
	// - **2**: HTTPS.
	//
	// example:
	//
	// 1
	ProtocolType *int32 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID of the repository.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address of the repository.
	//
	// example:
	//
	// 39.104.XXX.XXX
	RegistryHostIp *string `json:"RegistryHostIp,omitempty" xml:"RegistryHostIp,omitempty"`
	// The alias of the image repository.
	//
	// example:
	//
	// test1
	RegistryName *string `json:"RegistryName,omitempty" xml:"RegistryName,omitempty"`
	// The image repository type. Valid values:
	//
	// - **acr**: ACR
	//
	// - **harbor**: Harbor
	//
	// - **quay**: Quay
	//
	// - **CI/CD**: Jenkins.
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The authentication token of the user.
	//
	// example:
	//
	// c7b90d29-632f-4e58-88b8-00ad77f6****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The number of scan tasks per hour.
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
	// The instance ID of the VPC.
	//
	// example:
	//
	// vpc-5gu8iu68w9b472jbb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The whitelist.
	//
	// example:
	//
	// 192.168.XXX.XXX
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s PageImageRegistryResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s PageImageRegistryResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageImageRegistryResponseBodyList) GetBlackList() *string {
	return s.BlackList
}

func (s *PageImageRegistryResponseBodyList) GetDomainName() *string {
	return s.DomainName
}

func (s *PageImageRegistryResponseBodyList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PageImageRegistryResponseBodyList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PageImageRegistryResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *PageImageRegistryResponseBodyList) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *PageImageRegistryResponseBodyList) GetJenkinsEnv() *string {
	return s.JenkinsEnv
}

func (s *PageImageRegistryResponseBodyList) GetNetType() *int32 {
	return s.NetType
}

func (s *PageImageRegistryResponseBodyList) GetPassword() *string {
	return s.Password
}

func (s *PageImageRegistryResponseBodyList) GetPersistenceDay() *int32 {
	return s.PersistenceDay
}

func (s *PageImageRegistryResponseBodyList) GetProtocolType() *int32 {
	return s.ProtocolType
}

func (s *PageImageRegistryResponseBodyList) GetRegionId() *string {
	return s.RegionId
}

func (s *PageImageRegistryResponseBodyList) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *PageImageRegistryResponseBodyList) GetRegistryName() *string {
	return s.RegistryName
}

func (s *PageImageRegistryResponseBodyList) GetRegistryType() *string {
	return s.RegistryType
}

func (s *PageImageRegistryResponseBodyList) GetToken() *string {
	return s.Token
}

func (s *PageImageRegistryResponseBodyList) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *PageImageRegistryResponseBodyList) GetUserName() *string {
	return s.UserName
}

func (s *PageImageRegistryResponseBodyList) GetVpcId() *string {
	return s.VpcId
}

func (s *PageImageRegistryResponseBodyList) GetWhiteList() *string {
	return s.WhiteList
}

func (s *PageImageRegistryResponseBodyList) SetBlackList(v string) *PageImageRegistryResponseBodyList {
	s.BlackList = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetDomainName(v string) *PageImageRegistryResponseBodyList {
	s.DomainName = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetGmtCreate(v string) *PageImageRegistryResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetGmtModified(v string) *PageImageRegistryResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetId(v int64) *PageImageRegistryResponseBodyList {
	s.Id = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetImageCount(v int32) *PageImageRegistryResponseBodyList {
	s.ImageCount = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetJenkinsEnv(v string) *PageImageRegistryResponseBodyList {
	s.JenkinsEnv = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetNetType(v int32) *PageImageRegistryResponseBodyList {
	s.NetType = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetPassword(v string) *PageImageRegistryResponseBodyList {
	s.Password = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetPersistenceDay(v int32) *PageImageRegistryResponseBodyList {
	s.PersistenceDay = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetProtocolType(v int32) *PageImageRegistryResponseBodyList {
	s.ProtocolType = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetRegionId(v string) *PageImageRegistryResponseBodyList {
	s.RegionId = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetRegistryHostIp(v string) *PageImageRegistryResponseBodyList {
	s.RegistryHostIp = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetRegistryName(v string) *PageImageRegistryResponseBodyList {
	s.RegistryName = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetRegistryType(v string) *PageImageRegistryResponseBodyList {
	s.RegistryType = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetToken(v string) *PageImageRegistryResponseBodyList {
	s.Token = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetTransPerHour(v int32) *PageImageRegistryResponseBodyList {
	s.TransPerHour = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetUserName(v string) *PageImageRegistryResponseBodyList {
	s.UserName = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetVpcId(v string) *PageImageRegistryResponseBodyList {
	s.VpcId = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) SetWhiteList(v string) *PageImageRegistryResponseBodyList {
	s.WhiteList = &v
	return s
}

func (s *PageImageRegistryResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type PageImageRegistryResponseBodyPageInfo struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageImageRegistryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s PageImageRegistryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *PageImageRegistryResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *PageImageRegistryResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PageImageRegistryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageImageRegistryResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *PageImageRegistryResponseBodyPageInfo) SetCount(v int32) *PageImageRegistryResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *PageImageRegistryResponseBodyPageInfo) SetCurrentPage(v int32) *PageImageRegistryResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *PageImageRegistryResponseBodyPageInfo) SetPageSize(v int32) *PageImageRegistryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *PageImageRegistryResponseBodyPageInfo) SetTotalCount(v int32) *PageImageRegistryResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *PageImageRegistryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

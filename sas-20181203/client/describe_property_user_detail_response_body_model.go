// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUserDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyUserDetailResponseBodyPageInfo) *DescribePropertyUserDetailResponseBody
	GetPageInfo() *DescribePropertyUserDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertyUserDetailResponseBodyPropertys) *DescribePropertyUserDetailResponseBody
	GetPropertys() []*DescribePropertyUserDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertyUserDetailResponseBody
	GetRequestId() *string
}

type DescribePropertyUserDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyUserDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The details of asset fingerprints for the account.
	Propertys []*DescribePropertyUserDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 33A71BE3-2CC2-14CB-B460-33A1DD82953A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyUserDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponseBody) GetPageInfo() *DescribePropertyUserDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyUserDetailResponseBody) GetPropertys() []*DescribePropertyUserDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertyUserDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyUserDetailResponseBody) SetPageInfo(v *DescribePropertyUserDetailResponseBodyPageInfo) *DescribePropertyUserDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyUserDetailResponseBody) SetPropertys(v []*DescribePropertyUserDetailResponseBodyPropertys) *DescribePropertyUserDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyUserDetailResponseBody) SetRequestId(v string) *DescribePropertyUserDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Propertys != nil {
		for _, item := range s.Propertys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyUserDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyUserDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetNextToken(v string) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyUserDetailResponseBodyPropertys struct {
	// The date on which the account expires.
	//
	// example:
	//
	// never
	AccountsExpirationDate *string `json:"AccountsExpirationDate,omitempty" xml:"AccountsExpirationDate,omitempty"`
	// The timestamp at which the last asset fingerprint collection is performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649149566000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The details of the user groups to which the account belongs.
	GroupNames []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
	// The ID of the server.
	//
	// example:
	//
	// i-hp35tftuh52wbp1g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// hc-host-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 100.104.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP addresses of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Indicates whether the account is an interactive logon account. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	IsCouldLogin *int32 `json:"IsCouldLogin,omitempty" xml:"IsCouldLogin,omitempty"`
	// Indicates whether the password expires. Valid values:
	//
	// 	- **0**: yes
	//
	// 	- **1**: no
	//
	// example:
	//
	// 1
	IsPasswdExpired *int32 `json:"IsPasswdExpired,omitempty" xml:"IsPasswdExpired,omitempty"`
	// Indicates whether the password is locked. Valid values:
	//
	// 	- **0**: yes
	//
	// 	- **1**: no
	//
	// example:
	//
	// 1
	IsPasswdLocked *int32 `json:"IsPasswdLocked,omitempty" xml:"IsPasswdLocked,omitempty"`
	// Indicates whether the account has root permissions. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	IsRoot *string `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	// Indicates whether the account is a sudo account. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	IsSudoer *int32 `json:"IsSudoer,omitempty" xml:"IsSudoer,omitempty"`
	// Indicates whether the account expires. Valid values:
	//
	// 	- **0**: yes
	//
	// 	- **1**: no
	//
	// example:
	//
	// 1
	IsUserExpired *int32 `json:"IsUserExpired,omitempty" xml:"IsUserExpired,omitempty"`
	// The source IP address of the last logon to the account.
	//
	// example:
	//
	// 192.168.XX.XX
	LastLoginIp *string `json:"LastLoginIp,omitempty" xml:"LastLoginIp,omitempty"`
	// The last logon time of the account.
	//
	// example:
	//
	// 2022-04-04 18:07:06
	LastLoginTime *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	// The timestamp of the last logon to the account. Unit: milliseconds.
	//
	// example:
	//
	// 1649066826000
	LastLoginTimeDt *int64 `json:"LastLoginTimeDt,omitempty" xml:"LastLoginTimeDt,omitempty"`
	// The timestamp of the last logon to the account. Unit: milliseconds.
	//
	// example:
	//
	// 1649066826000
	LastLoginTimestamp *int64 `json:"LastLoginTimestamp,omitempty" xml:"LastLoginTimestamp,omitempty"`
	// The date on which the password of the account expires.
	//
	// example:
	//
	// never
	PasswordExpirationDate *string `json:"PasswordExpirationDate,omitempty" xml:"PasswordExpirationDate,omitempty"`
	// This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// **
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// bin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyUserDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetAccountsExpirationDate() *string {
	return s.AccountsExpirationDate
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetGroupNames() []*string {
	return s.GroupNames
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIp() *string {
	return s.Ip
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIsCouldLogin() *int32 {
	return s.IsCouldLogin
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIsPasswdExpired() *int32 {
	return s.IsPasswdExpired
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIsPasswdLocked() *int32 {
	return s.IsPasswdLocked
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIsRoot() *string {
	return s.IsRoot
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIsSudoer() *int32 {
	return s.IsSudoer
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetIsUserExpired() *int32 {
	return s.IsUserExpired
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetLastLoginIp() *string {
	return s.LastLoginIp
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetLastLoginTimeDt() *int64 {
	return s.LastLoginTimeDt
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetLastLoginTimestamp() *int64 {
	return s.LastLoginTimestamp
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetPasswordExpirationDate() *string {
	return s.PasswordExpirationDate
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetStatus() *string {
	return s.Status
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetUser() *string {
	return s.User
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetAccountsExpirationDate(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.AccountsExpirationDate = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyUserDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetGroupNames(v []*string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.GroupNames = v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsCouldLogin(v int32) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsCouldLogin = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsPasswdExpired(v int32) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsPasswdExpired = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsPasswdLocked(v int32) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsPasswdLocked = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsRoot(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsRoot = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsSudoer(v int32) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsSudoer = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsUserExpired(v int32) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsUserExpired = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginIp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginTime(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginTime = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginTimeDt(v int64) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginTimeDt = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginTimestamp(v int64) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginTimestamp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetPasswordExpirationDate(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.PasswordExpirationDate = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetStatus(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Status = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetUser(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}

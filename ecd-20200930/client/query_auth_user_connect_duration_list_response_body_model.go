// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthUserConnectDurationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthUserConnectDurationList(v []*QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) *QueryAuthUserConnectDurationListResponseBody
	GetAuthUserConnectDurationList() []*QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList
	SetNextToken(v string) *QueryAuthUserConnectDurationListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryAuthUserConnectDurationListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryAuthUserConnectDurationListResponseBody
	GetTotalCount() *int64
}

type QueryAuthUserConnectDurationListResponseBody struct {
	// The connection duration list of authorized users.
	AuthUserConnectDurationList []*QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList `json:"AuthUserConnectDurationList,omitempty" xml:"AuthUserConnectDurationList,omitempty" type:"Repeated"`
	// The pagination token for the next page. This parameter is returned when the results span multiple pages. Pass this value as the NextToken in the next request to retrieve the next page. This parameter is returned only when statistics are collected by individual session details.
	//
	// example:
	//
	// d129c6c0e8c04c8a9f0e2b7c1a3f5e6d
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F0F0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the specified conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAuthUserConnectDurationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthUserConnectDurationListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuthUserConnectDurationListResponseBody) GetAuthUserConnectDurationList() []*QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	return s.AuthUserConnectDurationList
}

func (s *QueryAuthUserConnectDurationListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryAuthUserConnectDurationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAuthUserConnectDurationListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryAuthUserConnectDurationListResponseBody) SetAuthUserConnectDurationList(v []*QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) *QueryAuthUserConnectDurationListResponseBody {
	s.AuthUserConnectDurationList = v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBody) SetNextToken(v string) *QueryAuthUserConnectDurationListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBody) SetRequestId(v string) *QueryAuthUserConnectDurationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBody) SetTotalCount(v int64) *QueryAuthUserConnectDurationListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBody) Validate() error {
	if s.AuthUserConnectDurationList != nil {
		for _, item := range s.AuthUserConnectDurationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList struct {
	// The connection duration of the user, in seconds.
	//
	// example:
	//
	// 3600
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// The end time of the connection, as a UNIX timestamp in milliseconds. This parameter is returned only when statistics are collected by individual session details (StatisticType=SingleSession).
	//
	// example:
	//
	// 1719208800000
	ConnectEndTime *string `json:"ConnectEndTime,omitempty" xml:"ConnectEndTime,omitempty"`
	// The start time of the connection, as a UNIX timestamp in milliseconds. This parameter is returned only when statistics are collected by individual session details (StatisticType=SingleSession).
	//
	// example:
	//
	// 1719205200000
	ConnectStartTime *string `json:"ConnectStartTime,omitempty" xml:"ConnectStartTime,omitempty"`
	// The remarks of the user. This parameter is returned only for convenience users when WithDetail is set to true.
	//
	// example:
	//
	// R&D department employee
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cloud desktop ID.
	//
	// example:
	//
	// ecd-gx2x1dhsmusr2****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud desktop name.
	//
	// example:
	//
	// test-desktop
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The type of the directory to which the user belongs. Valid values:
	//
	// - 1: convenience account.
	//
	// - 2: RAM account.
	//
	// - 3: AD account.
	//
	// - 4: personal edition.
	//
	// example:
	//
	// 3
	DirectoryType *int32 `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The display name of the user. This parameter is returned only for AD users when WithDetail is set to true.
	//
	// example:
	//
	// Zhang San
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The new display name of the user. This parameter is returned only for AD users when WithDetail is set to true.
	//
	// example:
	//
	// Zhang San
	DisplayNameNew *string `json:"DisplayNameNew,omitempty" xml:"DisplayNameNew,omitempty"`
	// The AD domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The nickname of the user. This parameter is returned only for convenience users when WithDetail is set to true.
	//
	// example:
	//
	// Xiao Zhang
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user principal name (UPN). This parameter is returned only for AD users when WithDetail is set to true.
	//
	// example:
	//
	// alice@example.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GoString() string {
	return s.String()
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetConnectDuration() *int64 {
	return s.ConnectDuration
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetConnectEndTime() *string {
	return s.ConnectEndTime
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetConnectStartTime() *string {
	return s.ConnectStartTime
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDescription() *string {
	return s.Description
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDesktopId() *string {
	return s.DesktopId
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDesktopName() *string {
	return s.DesktopName
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDirectoryType() *int32 {
	return s.DirectoryType
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDisplayNameNew() *string {
	return s.DisplayNameNew
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetEndUserId() *string {
	return s.EndUserId
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetNickName() *string {
	return s.NickName
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetConnectDuration(v int64) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.ConnectDuration = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetConnectEndTime(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.ConnectEndTime = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetConnectStartTime(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.ConnectStartTime = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDescription(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.Description = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDesktopId(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.DesktopId = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDesktopName(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.DesktopName = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDirectoryType(v int32) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.DirectoryType = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDisplayName(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.DisplayName = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDisplayNameNew(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.DisplayNameNew = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetDomainName(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.DomainName = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetEndUserId(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.EndUserId = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetNickName(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.NickName = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetRegionId(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.RegionId = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) SetUserPrincipalName(v string) *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList {
	s.UserPrincipalName = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponseBodyAuthUserConnectDurationList) Validate() error {
	return dara.Validate(s)
}

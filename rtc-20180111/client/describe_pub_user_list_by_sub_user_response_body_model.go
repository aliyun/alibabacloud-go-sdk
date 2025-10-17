// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePubUserListBySubUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallStatus(v string) *DescribePubUserListBySubUserResponseBody
	GetCallStatus() *string
	SetPubUserDetailList(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailList) *DescribePubUserListBySubUserResponseBody
	GetPubUserDetailList() []*DescribePubUserListBySubUserResponseBodyPubUserDetailList
	SetRequestId(v string) *DescribePubUserListBySubUserResponseBody
	GetRequestId() *string
	SetSubUserDetail(v *DescribePubUserListBySubUserResponseBodySubUserDetail) *DescribePubUserListBySubUserResponseBody
	GetSubUserDetail() *DescribePubUserListBySubUserResponseBodySubUserDetail
}

type DescribePubUserListBySubUserResponseBody struct {
	// example:
	//
	// IN
	CallStatus        *string                                                      `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	PubUserDetailList []*DescribePubUserListBySubUserResponseBodyPubUserDetailList `json:"PubUserDetailList,omitempty" xml:"PubUserDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubUserDetail *DescribePubUserListBySubUserResponseBodySubUserDetail `json:"SubUserDetail,omitempty" xml:"SubUserDetail,omitempty" type:"Struct"`
}

func (s DescribePubUserListBySubUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBody) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribePubUserListBySubUserResponseBody) GetPubUserDetailList() []*DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	return s.PubUserDetailList
}

func (s *DescribePubUserListBySubUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePubUserListBySubUserResponseBody) GetSubUserDetail() *DescribePubUserListBySubUserResponseBodySubUserDetail {
	return s.SubUserDetail
}

func (s *DescribePubUserListBySubUserResponseBody) SetCallStatus(v string) *DescribePubUserListBySubUserResponseBody {
	s.CallStatus = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetPubUserDetailList(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailList) *DescribePubUserListBySubUserResponseBody {
	s.PubUserDetailList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetRequestId(v string) *DescribePubUserListBySubUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetSubUserDetail(v *DescribePubUserListBySubUserResponseBodySubUserDetail) *DescribePubUserListBySubUserResponseBody {
	s.SubUserDetail = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) Validate() error {
	if s.PubUserDetailList != nil {
		for _, item := range s.PubUserDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubUserDetail != nil {
		if err := s.SubUserDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePubUserListBySubUserResponseBodyPubUserDetailList struct {
	CallIdList []*string `json:"CallIdList,omitempty" xml:"CallIdList,omitempty" type:"Repeated"`
	// example:
	//
	// NATIVE
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList []*string `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                                    `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os     *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList []*string `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0.0
	SdkVersion     *string   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 旁路转推
	UserIdAlias *string `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetCallIdList() []*string {
	return s.CallIdList
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetClientType() *string {
	return s.ClientType
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetLocation() *string {
	return s.Location
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetNetwork() *string {
	return s.Network
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetNetworkList() []*string {
	return s.NetworkList
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetOnlineDuration() *int64 {
	return s.OnlineDuration
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetOnlinePeriods() []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	return s.OnlinePeriods
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetOs() *string {
	return s.Os
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetOsList() []*string {
	return s.OsList
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetRoles() []*string {
	return s.Roles
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetSdkVersionList() []*string {
	return s.SdkVersionList
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetUserId() *string {
	return s.UserId
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) GetUserIdAlias() *string {
	return s.UserIdAlias
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetCallIdList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.CallIdList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetClientType(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.ClientType = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetCreatedTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetDestroyedTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetDuration(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetLocation(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetNetwork(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Network = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetNetworkList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.NetworkList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOnlineDuration(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOnlinePeriods(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOs(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Os = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOsList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OsList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetRoles(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Roles = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetSdkVersion(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.SdkVersion = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetSdkVersionList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.SdkVersionList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetUserId(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.UserId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetUserIdAlias(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.UserIdAlias = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) Validate() error {
	if s.OnlinePeriods != nil {
		for _, item := range s.OnlinePeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods struct {
	// example:
	//
	// 1614936817
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1614936817
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) GetJoinTs() *int64 {
	return s.JoinTs
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) GetLeaveTs() *int64 {
	return s.LeaveTs
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) Validate() error {
	return dara.Validate(s)
}

type DescribePubUserListBySubUserResponseBodySubUserDetail struct {
	// example:
	//
	// NATIVE
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList []*string `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                                `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os     *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList []*string `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0.0
	SdkVersion     *string   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 旁路转推
	UserIdAlias *string `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetail) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetClientType() *string {
	return s.ClientType
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetLocation() *string {
	return s.Location
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetNetwork() *string {
	return s.Network
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetNetworkList() []*string {
	return s.NetworkList
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetOnlineDuration() *int64 {
	return s.OnlineDuration
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetOnlinePeriods() []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	return s.OnlinePeriods
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetOs() *string {
	return s.Os
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetOsList() []*string {
	return s.OsList
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetRoles() []*string {
	return s.Roles
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetSdkVersionList() []*string {
	return s.SdkVersionList
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetUserId() *string {
	return s.UserId
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) GetUserIdAlias() *string {
	return s.UserIdAlias
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetClientType(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.ClientType = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetCreatedTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetDestroyedTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetDuration(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Duration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetLocation(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Location = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetNetwork(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Network = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetNetworkList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.NetworkList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOnlineDuration(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OnlineDuration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOnlinePeriods(v []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OnlinePeriods = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOs(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Os = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOsList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OsList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetRoles(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Roles = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetSdkVersion(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.SdkVersion = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetSdkVersionList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.SdkVersionList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetUserId(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.UserId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetUserIdAlias(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.UserIdAlias = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) Validate() error {
	if s.OnlinePeriods != nil {
		for _, item := range s.OnlinePeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods struct {
	// example:
	//
	// 1614936817
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1614936817
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) GetJoinTs() *int64 {
	return s.JoinTs
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) GetLeaveTs() *int64 {
	return s.LeaveTs
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) SetJoinTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) SetLeaveTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	s.LeaveTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMeetingRoomResponseBody
	GetRequestId() *string
	SetResult(v *QueryMeetingRoomResponseBodyResult) *QueryMeetingRoomResponseBody
	GetResult() *QueryMeetingRoomResponseBodyResult
	SetVendorRequestId(v string) *QueryMeetingRoomResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryMeetingRoomResponseBody
	GetVendorType() *string
}

type QueryMeetingRoomResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *QueryMeetingRoomResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryMeetingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMeetingRoomResponseBody) GetResult() *QueryMeetingRoomResponseBodyResult {
	return s.Result
}

func (s *QueryMeetingRoomResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryMeetingRoomResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryMeetingRoomResponseBody) SetRequestId(v string) *QueryMeetingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMeetingRoomResponseBody) SetResult(v *QueryMeetingRoomResponseBodyResult) *QueryMeetingRoomResponseBody {
	s.Result = v
	return s
}

func (s *QueryMeetingRoomResponseBody) SetVendorRequestId(v string) *QueryMeetingRoomResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryMeetingRoomResponseBody) SetVendorType(v string) *QueryMeetingRoomResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryMeetingRoomResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMeetingRoomResponseBodyResult struct {
	// example:
	//
	// ding994axxxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// [ "2iPOLbpxxxxuwggiiqiPwiEiF" ]
	DeviceUnionIds []*string `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableCycleReservation *bool `json:"EnableCycleReservation,omitempty" xml:"EnableCycleReservation,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId            *string                                                 `json:"IsvRoomId,omitempty" xml:"IsvRoomId,omitempty"`
	ReservationAuthority *QueryMeetingRoomResponseBodyResultReservationAuthority `json:"ReservationAuthority,omitempty" xml:"ReservationAuthority,omitempty" type:"Struct"`
	// example:
	//
	// 10
	RoomCapacity *int32                                       `json:"RoomCapacity,omitempty" xml:"RoomCapacity,omitempty"`
	RoomGroup    *QueryMeetingRoomResponseBodyResultRoomGroup `json:"RoomGroup,omitempty" xml:"RoomGroup,omitempty" type:"Struct"`
	// example:
	//
	// 0ffb7184xxxxx
	RoomId       *string                                         `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	RoomLabels   []*QueryMeetingRoomResponseBodyResultRoomLabels `json:"RoomLabels,omitempty" xml:"RoomLabels,omitempty" type:"Repeated"`
	RoomLocation *QueryMeetingRoomResponseBodyResultRoomLocation `json:"RoomLocation,omitempty" xml:"RoomLocation,omitempty" type:"Struct"`
	// example:
	//
	// 测试会议室
	RoomName *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADxxxxx.jpg
	RoomPicture *string `json:"RoomPicture,omitempty" xml:"RoomPicture,omitempty"`
	// example:
	//
	// 0122414
	RoomStaffId *string `json:"RoomStaffId,omitempty" xml:"RoomStaffId,omitempty"`
	// example:
	//
	// 0
	RoomStatus  *int32  `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	RoomUnionId *string `json:"RoomUnionId,omitempty" xml:"RoomUnionId,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResult) GetCorpId() *string {
	return s.CorpId
}

func (s *QueryMeetingRoomResponseBodyResult) GetDeviceUnionIds() []*string {
	return s.DeviceUnionIds
}

func (s *QueryMeetingRoomResponseBodyResult) GetEnableCycleReservation() *bool {
	return s.EnableCycleReservation
}

func (s *QueryMeetingRoomResponseBodyResult) GetIsvRoomId() *string {
	return s.IsvRoomId
}

func (s *QueryMeetingRoomResponseBodyResult) GetReservationAuthority() *QueryMeetingRoomResponseBodyResultReservationAuthority {
	return s.ReservationAuthority
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomCapacity() *int32 {
	return s.RoomCapacity
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomGroup() *QueryMeetingRoomResponseBodyResultRoomGroup {
	return s.RoomGroup
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomId() *string {
	return s.RoomId
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomLabels() []*QueryMeetingRoomResponseBodyResultRoomLabels {
	return s.RoomLabels
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomLocation() *QueryMeetingRoomResponseBodyResultRoomLocation {
	return s.RoomLocation
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomName() *string {
	return s.RoomName
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomPicture() *string {
	return s.RoomPicture
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomStaffId() *string {
	return s.RoomStaffId
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomStatus() *int32 {
	return s.RoomStatus
}

func (s *QueryMeetingRoomResponseBodyResult) GetRoomUnionId() *string {
	return s.RoomUnionId
}

func (s *QueryMeetingRoomResponseBodyResult) SetCorpId(v string) *QueryMeetingRoomResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetDeviceUnionIds(v []*string) *QueryMeetingRoomResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetEnableCycleReservation(v bool) *QueryMeetingRoomResponseBodyResult {
	s.EnableCycleReservation = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetIsvRoomId(v string) *QueryMeetingRoomResponseBodyResult {
	s.IsvRoomId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetReservationAuthority(v *QueryMeetingRoomResponseBodyResultReservationAuthority) *QueryMeetingRoomResponseBodyResult {
	s.ReservationAuthority = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomCapacity(v int32) *QueryMeetingRoomResponseBodyResult {
	s.RoomCapacity = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomGroup(v *QueryMeetingRoomResponseBodyResultRoomGroup) *QueryMeetingRoomResponseBodyResult {
	s.RoomGroup = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomId(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomLabels(v []*QueryMeetingRoomResponseBodyResultRoomLabels) *QueryMeetingRoomResponseBodyResult {
	s.RoomLabels = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomLocation(v *QueryMeetingRoomResponseBodyResultRoomLocation) *QueryMeetingRoomResponseBodyResult {
	s.RoomLocation = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomName(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomPicture(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomPicture = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomStaffId(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomStaffId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomStatus(v int32) *QueryMeetingRoomResponseBodyResult {
	s.RoomStatus = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomUnionId(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomUnionId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) Validate() error {
	if s.ReservationAuthority != nil {
		if err := s.ReservationAuthority.Validate(); err != nil {
			return err
		}
	}
	if s.RoomGroup != nil {
		if err := s.RoomGroup.Validate(); err != nil {
			return err
		}
	}
	if s.RoomLabels != nil {
		for _, item := range s.RoomLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RoomLocation != nil {
		if err := s.RoomLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMeetingRoomResponseBodyResultReservationAuthority struct {
	AuthorizedMembers []*QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers `json:"AuthorizedMembers,omitempty" xml:"AuthorizedMembers,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthority) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthority) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthority) GetAuthorizedMembers() []*QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	return s.AuthorizedMembers
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthority) SetAuthorizedMembers(v []*QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) *QueryMeetingRoomResponseBodyResultReservationAuthority {
	s.AuthorizedMembers = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthority) Validate() error {
	if s.AuthorizedMembers != nil {
		for _, item := range s.AuthorizedMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) GetMemberName() *string {
	return s.MemberName
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) SetMemberId(v string) *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	s.MemberId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) SetMemberName(v string) *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	s.MemberName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) SetMemberType(v string) *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	s.MemberType = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomResponseBodyResultRoomGroup struct {
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultRoomGroup) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultRoomGroup) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) SetGroupId(v int64) *QueryMeetingRoomResponseBodyResultRoomGroup {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) SetGroupName(v string) *QueryMeetingRoomResponseBodyResultRoomGroup {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) SetParentId(v int64) *QueryMeetingRoomResponseBodyResultRoomGroup {
	s.ParentId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomResponseBodyResultRoomLabels struct {
	// example:
	//
	// 1
	LabelId *int64 `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// example:
	//
	// 电视
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultRoomLabels) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultRoomLabels) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) GetLabelId() *int64 {
	return s.LabelId
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) GetLabelName() *string {
	return s.LabelName
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) SetLabelId(v int64) *QueryMeetingRoomResponseBodyResultRoomLabels {
	s.LabelId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) SetLabelName(v string) *QueryMeetingRoomResponseBodyResultRoomLabels {
	s.LabelName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomResponseBodyResultRoomLocation struct {
	// example:
	//
	// xx市xx区xx街道xx号
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// xxx公司
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultRoomLocation) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultRoomLocation) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) GetDesc() *string {
	return s.Desc
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) GetTitle() *string {
	return s.Title
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) SetDesc(v string) *QueryMeetingRoomResponseBodyResultRoomLocation {
	s.Desc = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) SetTitle(v string) *QueryMeetingRoomResponseBodyResultRoomLocation {
	s.Title = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) Validate() error {
	return dara.Validate(s)
}

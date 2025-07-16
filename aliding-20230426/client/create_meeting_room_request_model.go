// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCycleReservation(v bool) *CreateMeetingRoomRequest
	GetEnableCycleReservation() *bool
	SetGroupId(v int64) *CreateMeetingRoomRequest
	GetGroupId() *int64
	SetIsvRoomId(v string) *CreateMeetingRoomRequest
	GetIsvRoomId() *string
	SetReservationAuthority(v *CreateMeetingRoomRequestReservationAuthority) *CreateMeetingRoomRequest
	GetReservationAuthority() *CreateMeetingRoomRequestReservationAuthority
	SetRoomCapacity(v int32) *CreateMeetingRoomRequest
	GetRoomCapacity() *int32
	SetRoomLabelIds(v []*int64) *CreateMeetingRoomRequest
	GetRoomLabelIds() []*int64
	SetRoomLocation(v *CreateMeetingRoomRequestRoomLocation) *CreateMeetingRoomRequest
	GetRoomLocation() *CreateMeetingRoomRequestRoomLocation
	SetRoomName(v string) *CreateMeetingRoomRequest
	GetRoomName() *string
	SetRoomPicture(v string) *CreateMeetingRoomRequest
	GetRoomPicture() *string
	SetRoomStatus(v int32) *CreateMeetingRoomRequest
	GetRoomStatus() *int32
	SetTenantContext(v *CreateMeetingRoomRequestTenantContext) *CreateMeetingRoomRequest
	GetTenantContext() *CreateMeetingRoomRequestTenantContext
}

type CreateMeetingRoomRequest struct {
	EnableCycleReservation *bool `json:"EnableCycleReservation,omitempty" xml:"EnableCycleReservation,omitempty"`
	// example:
	//
	// 4644
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId            *string                                       `json:"IsvRoomId,omitempty" xml:"IsvRoomId,omitempty"`
	ReservationAuthority *CreateMeetingRoomRequestReservationAuthority `json:"ReservationAuthority,omitempty" xml:"ReservationAuthority,omitempty" type:"Struct"`
	// example:
	//
	// 100
	RoomCapacity *int32                                `json:"RoomCapacity,omitempty" xml:"RoomCapacity,omitempty"`
	RoomLabelIds []*int64                              `json:"RoomLabelIds,omitempty" xml:"RoomLabelIds,omitempty" type:"Repeated"`
	RoomLocation *CreateMeetingRoomRequestRoomLocation `json:"RoomLocation,omitempty" xml:"RoomLocation,omitempty" type:"Struct"`
	RoomName     *string                               `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPxxxxx.jpg
	RoomPicture *string `json:"RoomPicture,omitempty" xml:"RoomPicture,omitempty"`
	// example:
	//
	// 1
	RoomStatus    *int32                                 `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	TenantContext *CreateMeetingRoomRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CreateMeetingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequest) GetEnableCycleReservation() *bool {
	return s.EnableCycleReservation
}

func (s *CreateMeetingRoomRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateMeetingRoomRequest) GetIsvRoomId() *string {
	return s.IsvRoomId
}

func (s *CreateMeetingRoomRequest) GetReservationAuthority() *CreateMeetingRoomRequestReservationAuthority {
	return s.ReservationAuthority
}

func (s *CreateMeetingRoomRequest) GetRoomCapacity() *int32 {
	return s.RoomCapacity
}

func (s *CreateMeetingRoomRequest) GetRoomLabelIds() []*int64 {
	return s.RoomLabelIds
}

func (s *CreateMeetingRoomRequest) GetRoomLocation() *CreateMeetingRoomRequestRoomLocation {
	return s.RoomLocation
}

func (s *CreateMeetingRoomRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *CreateMeetingRoomRequest) GetRoomPicture() *string {
	return s.RoomPicture
}

func (s *CreateMeetingRoomRequest) GetRoomStatus() *int32 {
	return s.RoomStatus
}

func (s *CreateMeetingRoomRequest) GetTenantContext() *CreateMeetingRoomRequestTenantContext {
	return s.TenantContext
}

func (s *CreateMeetingRoomRequest) SetEnableCycleReservation(v bool) *CreateMeetingRoomRequest {
	s.EnableCycleReservation = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetGroupId(v int64) *CreateMeetingRoomRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetIsvRoomId(v string) *CreateMeetingRoomRequest {
	s.IsvRoomId = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetReservationAuthority(v *CreateMeetingRoomRequestReservationAuthority) *CreateMeetingRoomRequest {
	s.ReservationAuthority = v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomCapacity(v int32) *CreateMeetingRoomRequest {
	s.RoomCapacity = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomLabelIds(v []*int64) *CreateMeetingRoomRequest {
	s.RoomLabelIds = v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomLocation(v *CreateMeetingRoomRequestRoomLocation) *CreateMeetingRoomRequest {
	s.RoomLocation = v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomName(v string) *CreateMeetingRoomRequest {
	s.RoomName = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomPicture(v string) *CreateMeetingRoomRequest {
	s.RoomPicture = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomStatus(v int32) *CreateMeetingRoomRequest {
	s.RoomStatus = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetTenantContext(v *CreateMeetingRoomRequestTenantContext) *CreateMeetingRoomRequest {
	s.TenantContext = v
	return s
}

func (s *CreateMeetingRoomRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMeetingRoomRequestReservationAuthority struct {
	AuthorizedMembers []*CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers `json:"AuthorizedMembers,omitempty" xml:"AuthorizedMembers,omitempty" type:"Repeated"`
}

func (s CreateMeetingRoomRequestReservationAuthority) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomRequestReservationAuthority) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestReservationAuthority) GetAuthorizedMembers() []*CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	return s.AuthorizedMembers
}

func (s *CreateMeetingRoomRequestReservationAuthority) SetAuthorizedMembers(v []*CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) *CreateMeetingRoomRequestReservationAuthority {
	s.AuthorizedMembers = v
	return s
}

func (s *CreateMeetingRoomRequestReservationAuthority) Validate() error {
	return dara.Validate(s)
}

type CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberId(v string) *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberId = &v
	return s
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberName(v string) *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberName = &v
	return s
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberType(v string) *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberType = &v
	return s
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) Validate() error {
	return dara.Validate(s)
}

type CreateMeetingRoomRequestRoomLocation struct {
	Desc  *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateMeetingRoomRequestRoomLocation) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomRequestRoomLocation) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestRoomLocation) GetDesc() *string {
	return s.Desc
}

func (s *CreateMeetingRoomRequestRoomLocation) GetTitle() *string {
	return s.Title
}

func (s *CreateMeetingRoomRequestRoomLocation) SetDesc(v string) *CreateMeetingRoomRequestRoomLocation {
	s.Desc = &v
	return s
}

func (s *CreateMeetingRoomRequestRoomLocation) SetTitle(v string) *CreateMeetingRoomRequestRoomLocation {
	s.Title = &v
	return s
}

func (s *CreateMeetingRoomRequestRoomLocation) Validate() error {
	return dara.Validate(s)
}

type CreateMeetingRoomRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateMeetingRoomRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMeetingRoomRequestTenantContext) SetTenantId(v string) *CreateMeetingRoomRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateMeetingRoomRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

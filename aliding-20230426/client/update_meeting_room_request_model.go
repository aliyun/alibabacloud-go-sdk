// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCycleReservation(v bool) *UpdateMeetingRoomRequest
	GetEnableCycleReservation() *bool
	SetGroupId(v int64) *UpdateMeetingRoomRequest
	GetGroupId() *int64
	SetIsvRoomId(v string) *UpdateMeetingRoomRequest
	GetIsvRoomId() *string
	SetReservationAuthority(v *UpdateMeetingRoomRequestReservationAuthority) *UpdateMeetingRoomRequest
	GetReservationAuthority() *UpdateMeetingRoomRequestReservationAuthority
	SetRoomCapacity(v int32) *UpdateMeetingRoomRequest
	GetRoomCapacity() *int32
	SetRoomId(v string) *UpdateMeetingRoomRequest
	GetRoomId() *string
	SetRoomLabelIds(v []*int64) *UpdateMeetingRoomRequest
	GetRoomLabelIds() []*int64
	SetRoomLocation(v *UpdateMeetingRoomRequestRoomLocation) *UpdateMeetingRoomRequest
	GetRoomLocation() *UpdateMeetingRoomRequestRoomLocation
	SetRoomName(v string) *UpdateMeetingRoomRequest
	GetRoomName() *string
	SetRoomPicture(v string) *UpdateMeetingRoomRequest
	GetRoomPicture() *string
	SetRoomStatus(v int32) *UpdateMeetingRoomRequest
	GetRoomStatus() *int32
	SetTenantContext(v *UpdateMeetingRoomRequestTenantContext) *UpdateMeetingRoomRequest
	GetTenantContext() *UpdateMeetingRoomRequestTenantContext
}

type UpdateMeetingRoomRequest struct {
	EnableCycleReservation *bool `json:"EnableCycleReservation,omitempty" xml:"EnableCycleReservation,omitempty"`
	// example:
	//
	// 0
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId            *string                                       `json:"IsvRoomId,omitempty" xml:"IsvRoomId,omitempty"`
	ReservationAuthority *UpdateMeetingRoomRequestReservationAuthority `json:"ReservationAuthority,omitempty" xml:"ReservationAuthority,omitempty" type:"Struct"`
	// example:
	//
	// 100
	RoomCapacity *int32 `json:"RoomCapacity,omitempty" xml:"RoomCapacity,omitempty"`
	// example:
	//
	// 0ffbxxxxx
	RoomId       *string                               `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	RoomLabelIds []*int64                              `json:"RoomLabelIds,omitempty" xml:"RoomLabelIds,omitempty" type:"Repeated"`
	RoomLocation *UpdateMeetingRoomRequestRoomLocation `json:"RoomLocation,omitempty" xml:"RoomLocation,omitempty" type:"Struct"`
	RoomName     *string                               `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPxxxxx.jpg
	RoomPicture *string `json:"RoomPicture,omitempty" xml:"RoomPicture,omitempty"`
	// example:
	//
	// 1
	RoomStatus    *int32                                 `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	TenantContext *UpdateMeetingRoomRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateMeetingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequest) GetEnableCycleReservation() *bool {
	return s.EnableCycleReservation
}

func (s *UpdateMeetingRoomRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateMeetingRoomRequest) GetIsvRoomId() *string {
	return s.IsvRoomId
}

func (s *UpdateMeetingRoomRequest) GetReservationAuthority() *UpdateMeetingRoomRequestReservationAuthority {
	return s.ReservationAuthority
}

func (s *UpdateMeetingRoomRequest) GetRoomCapacity() *int32 {
	return s.RoomCapacity
}

func (s *UpdateMeetingRoomRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *UpdateMeetingRoomRequest) GetRoomLabelIds() []*int64 {
	return s.RoomLabelIds
}

func (s *UpdateMeetingRoomRequest) GetRoomLocation() *UpdateMeetingRoomRequestRoomLocation {
	return s.RoomLocation
}

func (s *UpdateMeetingRoomRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *UpdateMeetingRoomRequest) GetRoomPicture() *string {
	return s.RoomPicture
}

func (s *UpdateMeetingRoomRequest) GetRoomStatus() *int32 {
	return s.RoomStatus
}

func (s *UpdateMeetingRoomRequest) GetTenantContext() *UpdateMeetingRoomRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateMeetingRoomRequest) SetEnableCycleReservation(v bool) *UpdateMeetingRoomRequest {
	s.EnableCycleReservation = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetGroupId(v int64) *UpdateMeetingRoomRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetIsvRoomId(v string) *UpdateMeetingRoomRequest {
	s.IsvRoomId = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetReservationAuthority(v *UpdateMeetingRoomRequestReservationAuthority) *UpdateMeetingRoomRequest {
	s.ReservationAuthority = v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomCapacity(v int32) *UpdateMeetingRoomRequest {
	s.RoomCapacity = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomId(v string) *UpdateMeetingRoomRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomLabelIds(v []*int64) *UpdateMeetingRoomRequest {
	s.RoomLabelIds = v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomLocation(v *UpdateMeetingRoomRequestRoomLocation) *UpdateMeetingRoomRequest {
	s.RoomLocation = v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomName(v string) *UpdateMeetingRoomRequest {
	s.RoomName = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomPicture(v string) *UpdateMeetingRoomRequest {
	s.RoomPicture = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomStatus(v int32) *UpdateMeetingRoomRequest {
	s.RoomStatus = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetTenantContext(v *UpdateMeetingRoomRequestTenantContext) *UpdateMeetingRoomRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateMeetingRoomRequest) Validate() error {
	if s.ReservationAuthority != nil {
		if err := s.ReservationAuthority.Validate(); err != nil {
			return err
		}
	}
	if s.RoomLocation != nil {
		if err := s.RoomLocation.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMeetingRoomRequestReservationAuthority struct {
	AuthorizedMembers []*UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers `json:"AuthorizedMembers,omitempty" xml:"AuthorizedMembers,omitempty" type:"Repeated"`
}

func (s UpdateMeetingRoomRequestReservationAuthority) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomRequestReservationAuthority) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestReservationAuthority) GetAuthorizedMembers() []*UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	return s.AuthorizedMembers
}

func (s *UpdateMeetingRoomRequestReservationAuthority) SetAuthorizedMembers(v []*UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) *UpdateMeetingRoomRequestReservationAuthority {
	s.AuthorizedMembers = v
	return s
}

func (s *UpdateMeetingRoomRequestReservationAuthority) Validate() error {
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

type UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GetMemberName() *string {
	return s.MemberName
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberId(v string) *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberName(v string) *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberName = &v
	return s
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberType(v string) *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) Validate() error {
	return dara.Validate(s)
}

type UpdateMeetingRoomRequestRoomLocation struct {
	Desc  *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateMeetingRoomRequestRoomLocation) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomRequestRoomLocation) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestRoomLocation) GetDesc() *string {
	return s.Desc
}

func (s *UpdateMeetingRoomRequestRoomLocation) GetTitle() *string {
	return s.Title
}

func (s *UpdateMeetingRoomRequestRoomLocation) SetDesc(v string) *UpdateMeetingRoomRequestRoomLocation {
	s.Desc = &v
	return s
}

func (s *UpdateMeetingRoomRequestRoomLocation) SetTitle(v string) *UpdateMeetingRoomRequestRoomLocation {
	s.Title = &v
	return s
}

func (s *UpdateMeetingRoomRequestRoomLocation) Validate() error {
	return dara.Validate(s)
}

type UpdateMeetingRoomRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateMeetingRoomRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMeetingRoomRequestTenantContext) SetTenantId(v string) *UpdateMeetingRoomRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateMeetingRoomRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

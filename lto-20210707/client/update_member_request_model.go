// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedCount(v int64) *UpdateMemberRequest
	GetAuthorizedCount() *int64
	SetAuthorizedDeviceCount(v int32) *UpdateMemberRequest
	GetAuthorizedDeviceCount() *int32
	SetContactor(v string) *UpdateMemberRequest
	GetContactor() *string
	SetMemberId(v string) *UpdateMemberRequest
	GetMemberId() *string
	SetName(v string) *UpdateMemberRequest
	GetName() *string
	SetRegionId(v string) *UpdateMemberRequest
	GetRegionId() *string
	SetRemark(v string) *UpdateMemberRequest
	GetRemark() *string
	SetTelephony(v string) *UpdateMemberRequest
	GetTelephony() *string
	SetUid(v string) *UpdateMemberRequest
	GetUid() *string
}

type UpdateMemberRequest struct {
	// This parameter is required.
	AuthorizedCount       *int64 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	AuthorizedDeviceCount *int32 `json:"AuthorizedDeviceCount,omitempty" xml:"AuthorizedDeviceCount,omitempty"`
	// This parameter is required.
	Contactor *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	// This parameter is required.
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	Telephony *string `json:"Telephony,omitempty" xml:"Telephony,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s UpdateMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberRequest) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *UpdateMemberRequest) GetAuthorizedDeviceCount() *int32 {
	return s.AuthorizedDeviceCount
}

func (s *UpdateMemberRequest) GetContactor() *string {
	return s.Contactor
}

func (s *UpdateMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateMemberRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateMemberRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateMemberRequest) GetTelephony() *string {
	return s.Telephony
}

func (s *UpdateMemberRequest) GetUid() *string {
	return s.Uid
}

func (s *UpdateMemberRequest) SetAuthorizedCount(v int64) *UpdateMemberRequest {
	s.AuthorizedCount = &v
	return s
}

func (s *UpdateMemberRequest) SetAuthorizedDeviceCount(v int32) *UpdateMemberRequest {
	s.AuthorizedDeviceCount = &v
	return s
}

func (s *UpdateMemberRequest) SetContactor(v string) *UpdateMemberRequest {
	s.Contactor = &v
	return s
}

func (s *UpdateMemberRequest) SetMemberId(v string) *UpdateMemberRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateMemberRequest) SetName(v string) *UpdateMemberRequest {
	s.Name = &v
	return s
}

func (s *UpdateMemberRequest) SetRegionId(v string) *UpdateMemberRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMemberRequest) SetRemark(v string) *UpdateMemberRequest {
	s.Remark = &v
	return s
}

func (s *UpdateMemberRequest) SetTelephony(v string) *UpdateMemberRequest {
	s.Telephony = &v
	return s
}

func (s *UpdateMemberRequest) SetUid(v string) *UpdateMemberRequest {
	s.Uid = &v
	return s
}

func (s *UpdateMemberRequest) Validate() error {
	return dara.Validate(s)
}

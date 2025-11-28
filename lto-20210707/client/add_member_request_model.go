// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedCount(v int64) *AddMemberRequest
	GetAuthorizedCount() *int64
	SetAuthorizedDeviceCount(v int32) *AddMemberRequest
	GetAuthorizedDeviceCount() *int32
	SetContactor(v string) *AddMemberRequest
	GetContactor() *string
	SetName(v string) *AddMemberRequest
	GetName() *string
	SetRegionId(v string) *AddMemberRequest
	GetRegionId() *string
	SetRemark(v string) *AddMemberRequest
	GetRemark() *string
	SetTelephony(v string) *AddMemberRequest
	GetTelephony() *string
	SetUid(v string) *AddMemberRequest
	GetUid() *string
}

type AddMemberRequest struct {
	// This parameter is required.
	AuthorizedCount       *int64 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	AuthorizedDeviceCount *int32 `json:"AuthorizedDeviceCount,omitempty" xml:"AuthorizedDeviceCount,omitempty"`
	// This parameter is required.
	Contactor *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	Telephony *string `json:"Telephony,omitempty" xml:"Telephony,omitempty"`
	// This parameter is required.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s AddMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMemberRequest) GoString() string {
	return s.String()
}

func (s *AddMemberRequest) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *AddMemberRequest) GetAuthorizedDeviceCount() *int32 {
	return s.AuthorizedDeviceCount
}

func (s *AddMemberRequest) GetContactor() *string {
	return s.Contactor
}

func (s *AddMemberRequest) GetName() *string {
	return s.Name
}

func (s *AddMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddMemberRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddMemberRequest) GetTelephony() *string {
	return s.Telephony
}

func (s *AddMemberRequest) GetUid() *string {
	return s.Uid
}

func (s *AddMemberRequest) SetAuthorizedCount(v int64) *AddMemberRequest {
	s.AuthorizedCount = &v
	return s
}

func (s *AddMemberRequest) SetAuthorizedDeviceCount(v int32) *AddMemberRequest {
	s.AuthorizedDeviceCount = &v
	return s
}

func (s *AddMemberRequest) SetContactor(v string) *AddMemberRequest {
	s.Contactor = &v
	return s
}

func (s *AddMemberRequest) SetName(v string) *AddMemberRequest {
	s.Name = &v
	return s
}

func (s *AddMemberRequest) SetRegionId(v string) *AddMemberRequest {
	s.RegionId = &v
	return s
}

func (s *AddMemberRequest) SetRemark(v string) *AddMemberRequest {
	s.Remark = &v
	return s
}

func (s *AddMemberRequest) SetTelephony(v string) *AddMemberRequest {
	s.Telephony = &v
	return s
}

func (s *AddMemberRequest) SetUid(v string) *AddMemberRequest {
	s.Uid = &v
	return s
}

func (s *AddMemberRequest) Validate() error {
	return dara.Validate(s)
}

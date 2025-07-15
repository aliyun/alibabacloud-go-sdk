// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *CreateDriveRequest
	GetAliUid() *int64
	SetDescription(v string) *CreateDriveRequest
	GetDescription() *string
	SetDomainId(v string) *CreateDriveRequest
	GetDomainId() *string
	SetDriveName(v string) *CreateDriveRequest
	GetDriveName() *string
	SetExternalDomainId(v string) *CreateDriveRequest
	GetExternalDomainId() *string
	SetProfileRoaming(v bool) *CreateDriveRequest
	GetProfileRoaming() *bool
	SetRegionId(v string) *CreateDriveRequest
	GetRegionId() *string
	SetResourceType(v string) *CreateDriveRequest
	GetResourceType() *string
	SetType(v string) *CreateDriveRequest
	GetType() *string
	SetUserId(v string) *CreateDriveRequest
	GetUserId() *string
}

type CreateDriveRequest struct {
	// example:
	//
	// 1202****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// test01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dom-aaaa****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// test01
	DriveName *string `json:"DriveName,omitempty" xml:"DriveName,omitempty"`
	// example:
	//
	// 1234****
	ExternalDomainId *string `json:"ExternalDomainId,omitempty" xml:"ExternalDomainId,omitempty"`
	// example:
	//
	// -
	ProfileRoaming *bool `json:"ProfileRoaming,omitempty" xml:"ProfileRoaming,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// NAS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// USER_PROFILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDriveRequest) GoString() string {
	return s.String()
}

func (s *CreateDriveRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *CreateDriveRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDriveRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDriveRequest) GetDriveName() *string {
	return s.DriveName
}

func (s *CreateDriveRequest) GetExternalDomainId() *string {
	return s.ExternalDomainId
}

func (s *CreateDriveRequest) GetProfileRoaming() *bool {
	return s.ProfileRoaming
}

func (s *CreateDriveRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDriveRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateDriveRequest) GetType() *string {
	return s.Type
}

func (s *CreateDriveRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateDriveRequest) SetAliUid(v int64) *CreateDriveRequest {
	s.AliUid = &v
	return s
}

func (s *CreateDriveRequest) SetDescription(v string) *CreateDriveRequest {
	s.Description = &v
	return s
}

func (s *CreateDriveRequest) SetDomainId(v string) *CreateDriveRequest {
	s.DomainId = &v
	return s
}

func (s *CreateDriveRequest) SetDriveName(v string) *CreateDriveRequest {
	s.DriveName = &v
	return s
}

func (s *CreateDriveRequest) SetExternalDomainId(v string) *CreateDriveRequest {
	s.ExternalDomainId = &v
	return s
}

func (s *CreateDriveRequest) SetProfileRoaming(v bool) *CreateDriveRequest {
	s.ProfileRoaming = &v
	return s
}

func (s *CreateDriveRequest) SetRegionId(v string) *CreateDriveRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDriveRequest) SetResourceType(v string) *CreateDriveRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateDriveRequest) SetType(v string) *CreateDriveRequest {
	s.Type = &v
	return s
}

func (s *CreateDriveRequest) SetUserId(v string) *CreateDriveRequest {
	s.UserId = &v
	return s
}

func (s *CreateDriveRequest) Validate() error {
	return dara.Validate(s)
}

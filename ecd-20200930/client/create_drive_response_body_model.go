// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDriveResponseBody
	GetCode() *string
	SetDrive(v *CreateDriveResponseBodyDrive) *CreateDriveResponseBody
	GetDrive() *CreateDriveResponseBodyDrive
	SetMessage(v string) *CreateDriveResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDriveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDriveResponseBody
	GetSuccess() *bool
}

type CreateDriveResponseBody struct {
	// The response code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The user-level storage resource.
	Drive *CreateDriveResponseBodyDrive `json:"Drive,omitempty" xml:"Drive,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B7AA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDriveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDriveResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDriveResponseBody) GetDrive() *CreateDriveResponseBodyDrive {
	return s.Drive
}

func (s *CreateDriveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDriveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDriveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDriveResponseBody) SetCode(v string) *CreateDriveResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDriveResponseBody) SetDrive(v *CreateDriveResponseBodyDrive) *CreateDriveResponseBody {
	s.Drive = v
	return s
}

func (s *CreateDriveResponseBody) SetMessage(v string) *CreateDriveResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDriveResponseBody) SetRequestId(v string) *CreateDriveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDriveResponseBody) SetSuccess(v bool) *CreateDriveResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDriveResponseBody) Validate() error {
	if s.Drive != nil {
		if err := s.Drive.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDriveResponseBodyDrive struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1202****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The description of the storage resource.
	//
	// example:
	//
	// test****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the storage resource.
	//
	// example:
	//
	// dom-aaaa****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the user-level storage resource.
	//
	// example:
	//
	// dri-aaaa****
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// The ID of the external user-level storage resource.
	//
	// >  A value is returned for this parameter only if ResourceType is set to PDS.
	//
	// example:
	//
	// 1100****
	ExternalDriveId *string `json:"ExternalDriveId,omitempty" xml:"ExternalDriveId,omitempty"`
	// The ID of the external user.
	//
	// example:
	//
	// user01@cn-hangzhou.120****
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// The time when the storage resource was created.
	//
	// example:
	//
	// 2025-07-02T08:42:26.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the storage resource was last modified.
	//
	// example:
	//
	// 2025-07-07T02:46:04.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the storage resource.
	//
	// example:
	//
	// test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// -
	ProfileRoaming *bool `json:"ProfileRoaming,omitempty" xml:"ProfileRoaming,omitempty"`
	// The status of the user-level storage resource.
	//
	// Valid values:
	//
	// 	- INIT: initializing
	//
	// 	- MAINTAIN: maintaining
	//
	// 	- DELETING
	//
	// 	- INVALID
	//
	// 	- NORMAL
	//
	// 	- FAIL: failed
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total amount of the user-level storage resource capacity.
	//
	// >  Unit: bytes.
	//
	// example:
	//
	// 536870912000
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The usage of the storage resource.
	//
	// Valid values:
	//
	// 	- DESKTOP: data disk space
	//
	// 	- USER_PROFILE: space for personal data of the user
	//
	// example:
	//
	// USER_PROFILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The amount of user-level storage resource capacity used.
	//
	// >  Unit: bytes.
	//
	// example:
	//
	// 243175936
	UsedSize *int64 `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateDriveResponseBodyDrive) String() string {
	return dara.Prettify(s)
}

func (s CreateDriveResponseBodyDrive) GoString() string {
	return s.String()
}

func (s *CreateDriveResponseBodyDrive) GetAliUid() *string {
	return s.AliUid
}

func (s *CreateDriveResponseBodyDrive) GetDescription() *string {
	return s.Description
}

func (s *CreateDriveResponseBodyDrive) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDriveResponseBodyDrive) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateDriveResponseBodyDrive) GetExternalDriveId() *string {
	return s.ExternalDriveId
}

func (s *CreateDriveResponseBodyDrive) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *CreateDriveResponseBodyDrive) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateDriveResponseBodyDrive) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateDriveResponseBodyDrive) GetId() *string {
	return s.Id
}

func (s *CreateDriveResponseBodyDrive) GetName() *string {
	return s.Name
}

func (s *CreateDriveResponseBodyDrive) GetProfileRoaming() *bool {
	return s.ProfileRoaming
}

func (s *CreateDriveResponseBodyDrive) GetStatus() *string {
	return s.Status
}

func (s *CreateDriveResponseBodyDrive) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CreateDriveResponseBodyDrive) GetType() *string {
	return s.Type
}

func (s *CreateDriveResponseBodyDrive) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *CreateDriveResponseBodyDrive) GetUserId() *string {
	return s.UserId
}

func (s *CreateDriveResponseBodyDrive) SetAliUid(v string) *CreateDriveResponseBodyDrive {
	s.AliUid = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetDescription(v string) *CreateDriveResponseBodyDrive {
	s.Description = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetDomainId(v string) *CreateDriveResponseBodyDrive {
	s.DomainId = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetDriveId(v string) *CreateDriveResponseBodyDrive {
	s.DriveId = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetExternalDriveId(v string) *CreateDriveResponseBodyDrive {
	s.ExternalDriveId = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetExternalUserId(v string) *CreateDriveResponseBodyDrive {
	s.ExternalUserId = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetGmtCreate(v string) *CreateDriveResponseBodyDrive {
	s.GmtCreate = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetGmtModified(v string) *CreateDriveResponseBodyDrive {
	s.GmtModified = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetId(v string) *CreateDriveResponseBodyDrive {
	s.Id = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetName(v string) *CreateDriveResponseBodyDrive {
	s.Name = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetProfileRoaming(v bool) *CreateDriveResponseBodyDrive {
	s.ProfileRoaming = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetStatus(v string) *CreateDriveResponseBodyDrive {
	s.Status = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetTotalSize(v int64) *CreateDriveResponseBodyDrive {
	s.TotalSize = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetType(v string) *CreateDriveResponseBodyDrive {
	s.Type = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetUsedSize(v int64) *CreateDriveResponseBodyDrive {
	s.UsedSize = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) SetUserId(v string) *CreateDriveResponseBodyDrive {
	s.UserId = &v
	return s
}

func (s *CreateDriveResponseBodyDrive) Validate() error {
	return dara.Validate(s)
}

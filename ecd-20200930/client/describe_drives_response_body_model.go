// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrivesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDrivesResponseBody
	GetCode() *string
	SetCount(v int32) *DescribeDrivesResponseBody
	GetCount() *int32
	SetDrives(v []*DescribeDrivesResponseBodyDrives) *DescribeDrivesResponseBody
	GetDrives() []*DescribeDrivesResponseBodyDrives
	SetMessage(v string) *DescribeDrivesResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeDrivesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDrivesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrivesResponseBody
	GetSuccess() *bool
}

type DescribeDrivesResponseBody struct {
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of user-level storage resources.
	Drives []*DescribeDrivesResponseBodyDrives `json:"Drives,omitempty" xml:"Drives,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token for the next query. An empty value indicates that there are no more results.
	//
	// example:
	//
	// AAAA****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
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

func (s DescribeDrivesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrivesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrivesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDrivesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDrivesResponseBody) GetDrives() []*DescribeDrivesResponseBodyDrives {
	return s.Drives
}

func (s *DescribeDrivesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDrivesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDrivesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrivesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrivesResponseBody) SetCode(v string) *DescribeDrivesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDrivesResponseBody) SetCount(v int32) *DescribeDrivesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDrivesResponseBody) SetDrives(v []*DescribeDrivesResponseBodyDrives) *DescribeDrivesResponseBody {
	s.Drives = v
	return s
}

func (s *DescribeDrivesResponseBody) SetMessage(v string) *DescribeDrivesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDrivesResponseBody) SetNextToken(v string) *DescribeDrivesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDrivesResponseBody) SetRequestId(v string) *DescribeDrivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrivesResponseBody) SetSuccess(v bool) *DescribeDrivesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrivesResponseBody) Validate() error {
	if s.Drives != nil {
		for _, item := range s.Drives {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDrivesResponseBodyDrives struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1202****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The storage resource description.
	//
	// example:
	//
	// test****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of associated cloud computer pools.
	//
	// > This parameter is returned only when the storage resource is NAS and the purpose is USER_PROFILE.
	//
	// example:
	//
	// 1
	DesktopGroupCount *int32 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The list of associated cloud computer pool details.
	//
	// > This parameter is returned only when the storage resource is NAS and the purpose is USER_PROFILE.
	DesktopGroups []*DescribeDrivesResponseBodyDrivesDesktopGroups `json:"DesktopGroups,omitempty" xml:"DesktopGroups,omitempty" type:"Repeated"`
	// The storage resource ID.
	//
	// example:
	//
	// dom-aaaa****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The user-level storage resource ID.
	//
	// example:
	//
	// dri-aaaa****
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// Indicates whether the User Profile Management (UPM) feature is enabled.
	//
	// example:
	//
	// true
	EnableProfileManagement *bool `json:"EnableProfileManagement,omitempty" xml:"EnableProfileManagement,omitempty"`
	// The external storage resource ID.
	//
	// - If the storage resource is NAS, this parameter returns the NAS ID.
	//
	// - If the storage resource is PDS, this parameter returns the PDS ID.
	//
	// example:
	//
	// 0976****
	ExternalDomainId *string `json:"ExternalDomainId,omitempty" xml:"ExternalDomainId,omitempty"`
	// The external user-level storage resource ID.
	//
	// > This parameter is returned only when the storage resource is PDS.
	//
	// example:
	//
	// 1100****
	ExternalDriveId *string `json:"ExternalDriveId,omitempty" xml:"ExternalDriveId,omitempty"`
	// The external user ID.
	//
	// > This parameter is returned only when the storage resource is PDS.
	//
	// example:
	//
	// user01@cn-hangzhou.120****
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-07-02T08:42:26.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-07-07T02:46:04.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID.
	//
	// > You can ignore this parameter.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The storage resource name.
	//
	// example:
	//
	// test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// -
	ProfileRoaming *bool `json:"ProfileRoaming,omitempty" xml:"ProfileRoaming,omitempty"`
	// The status of the user-level storage resource.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total capacity of the user-level storage resource.
	//
	// example:
	//
	// 536870912000
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The purpose of the storage resource.
	//
	// example:
	//
	// USER_PROFILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The used capacity of the user-level storage resource.
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

func (s DescribeDrivesResponseBodyDrives) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrivesResponseBodyDrives) GoString() string {
	return s.String()
}

func (s *DescribeDrivesResponseBodyDrives) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDrivesResponseBodyDrives) GetDescription() *string {
	return s.Description
}

func (s *DescribeDrivesResponseBodyDrives) GetDesktopGroupCount() *int32 {
	return s.DesktopGroupCount
}

func (s *DescribeDrivesResponseBodyDrives) GetDesktopGroups() []*DescribeDrivesResponseBodyDrivesDesktopGroups {
	return s.DesktopGroups
}

func (s *DescribeDrivesResponseBodyDrives) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDrivesResponseBodyDrives) GetDriveId() *string {
	return s.DriveId
}

func (s *DescribeDrivesResponseBodyDrives) GetEnableProfileManagement() *bool {
	return s.EnableProfileManagement
}

func (s *DescribeDrivesResponseBodyDrives) GetExternalDomainId() *string {
	return s.ExternalDomainId
}

func (s *DescribeDrivesResponseBodyDrives) GetExternalDriveId() *string {
	return s.ExternalDriveId
}

func (s *DescribeDrivesResponseBodyDrives) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *DescribeDrivesResponseBodyDrives) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDrivesResponseBodyDrives) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDrivesResponseBodyDrives) GetId() *string {
	return s.Id
}

func (s *DescribeDrivesResponseBodyDrives) GetName() *string {
	return s.Name
}

func (s *DescribeDrivesResponseBodyDrives) GetProfileRoaming() *bool {
	return s.ProfileRoaming
}

func (s *DescribeDrivesResponseBodyDrives) GetStatus() *string {
	return s.Status
}

func (s *DescribeDrivesResponseBodyDrives) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeDrivesResponseBodyDrives) GetType() *string {
	return s.Type
}

func (s *DescribeDrivesResponseBodyDrives) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *DescribeDrivesResponseBodyDrives) GetUserId() *string {
	return s.UserId
}

func (s *DescribeDrivesResponseBodyDrives) SetAliUid(v int64) *DescribeDrivesResponseBodyDrives {
	s.AliUid = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetDescription(v string) *DescribeDrivesResponseBodyDrives {
	s.Description = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetDesktopGroupCount(v int32) *DescribeDrivesResponseBodyDrives {
	s.DesktopGroupCount = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetDesktopGroups(v []*DescribeDrivesResponseBodyDrivesDesktopGroups) *DescribeDrivesResponseBodyDrives {
	s.DesktopGroups = v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetDomainId(v string) *DescribeDrivesResponseBodyDrives {
	s.DomainId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetDriveId(v string) *DescribeDrivesResponseBodyDrives {
	s.DriveId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetEnableProfileManagement(v bool) *DescribeDrivesResponseBodyDrives {
	s.EnableProfileManagement = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetExternalDomainId(v string) *DescribeDrivesResponseBodyDrives {
	s.ExternalDomainId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetExternalDriveId(v string) *DescribeDrivesResponseBodyDrives {
	s.ExternalDriveId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetExternalUserId(v string) *DescribeDrivesResponseBodyDrives {
	s.ExternalUserId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetGmtCreate(v string) *DescribeDrivesResponseBodyDrives {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetGmtModified(v string) *DescribeDrivesResponseBodyDrives {
	s.GmtModified = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetId(v string) *DescribeDrivesResponseBodyDrives {
	s.Id = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetName(v string) *DescribeDrivesResponseBodyDrives {
	s.Name = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetProfileRoaming(v bool) *DescribeDrivesResponseBodyDrives {
	s.ProfileRoaming = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetStatus(v string) *DescribeDrivesResponseBodyDrives {
	s.Status = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetTotalSize(v int64) *DescribeDrivesResponseBodyDrives {
	s.TotalSize = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetType(v string) *DescribeDrivesResponseBodyDrives {
	s.Type = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetUsedSize(v int64) *DescribeDrivesResponseBodyDrives {
	s.UsedSize = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) SetUserId(v string) *DescribeDrivesResponseBodyDrives {
	s.UserId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrives) Validate() error {
	if s.DesktopGroups != nil {
		for _, item := range s.DesktopGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDrivesResponseBodyDrivesDesktopGroups struct {
	// The cloud computer pool ID.
	//
	// example:
	//
	// dg-aaaa****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The cloud computer pool name.
	//
	// example:
	//
	// group01
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
}

func (s DescribeDrivesResponseBodyDrivesDesktopGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrivesResponseBodyDrivesDesktopGroups) GoString() string {
	return s.String()
}

func (s *DescribeDrivesResponseBodyDrivesDesktopGroups) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDrivesResponseBodyDrivesDesktopGroups) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeDrivesResponseBodyDrivesDesktopGroups) SetDesktopGroupId(v string) *DescribeDrivesResponseBodyDrivesDesktopGroups {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrivesDesktopGroups) SetDesktopGroupName(v string) *DescribeDrivesResponseBodyDrivesDesktopGroups {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDrivesResponseBodyDrivesDesktopGroups) Validate() error {
	return dara.Validate(s)
}

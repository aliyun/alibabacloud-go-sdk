// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDriveGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudDriveGroups(v []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) *DescribeCloudDriveGroupsResponseBody
	GetCloudDriveGroups() []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroups
	SetCount(v int64) *DescribeCloudDriveGroupsResponseBody
	GetCount() *int64
	SetNextToken(v string) *DescribeCloudDriveGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCloudDriveGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudDriveGroupsResponseBody
	GetSuccess() *bool
}

type DescribeCloudDriveGroupsResponseBody struct {
	// The list of cloud drive team spaces.
	CloudDriveGroups []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroups `json:"CloudDriveGroups,omitempty" xml:"CloudDriveGroups,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The token for the next query. If `NextToken` is empty, no more results exist.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EF015AE5-B30A-5189-B519-735CEE40****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudDriveGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveGroupsResponseBody) GetCloudDriveGroups() []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	return s.CloudDriveGroups
}

func (s *DescribeCloudDriveGroupsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCloudDriveGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDriveGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDriveGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudDriveGroupsResponseBody) SetCloudDriveGroups(v []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) *DescribeCloudDriveGroupsResponseBody {
	s.CloudDriveGroups = v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBody) SetCount(v int64) *DescribeCloudDriveGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBody) SetNextToken(v string) *DescribeCloudDriveGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBody) SetRequestId(v string) *DescribeCloudDriveGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBody) SetSuccess(v bool) *DescribeCloudDriveGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBody) Validate() error {
	if s.CloudDriveGroups != nil {
		for _, item := range s.CloudDriveGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudDriveGroupsResponseBodyCloudDriveGroups struct {
	// The list of team administrators.
	//
	// example:
	//
	// ["user01"]
	AdminUserIds *string `json:"AdminUserIds,omitempty" xml:"AdminUserIds,omitempty"`
	// The team storage management administrator information.
	AdminUserInfos []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos `json:"AdminUserInfos,omitempty" xml:"AdminUserInfos,omitempty" type:"Repeated"`
	// The creation time. The time is in the ISO 8601 standard in the UTC format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-04-11T07:44:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-959593****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The space ID.
	//
	// example:
	//
	// sh1234
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// The team space ID.
	//
	// example:
	//
	// cg-e70ga4ixp30ur****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The team space name.
	//
	// example:
	//
	// TestTeam1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The organization ID of the team.
	//
	// example:
	//
	// org-aliyun-wy-org-id
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The size of the team space recycle bin. Unit: bytes.
	//
	// example:
	//
	// 1024000
	RecycleBinSize *string `json:"RecycleBinSize,omitempty" xml:"RecycleBinSize,omitempty"`
	// The team space status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total capacity of the cloud drive team space.
	//
	// example:
	//
	// 5368709120
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The used space size. Unit: bytes.
	//
	// example:
	//
	// 1024000000
	UsedSize *string `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
}

func (s DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetAdminUserIds() *string {
	return s.AdminUserIds
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetAdminUserInfos() []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	return s.AdminUserInfos
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetDriveId() *string {
	return s.DriveId
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetRecycleBinSize() *string {
	return s.RecycleBinSize
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) GetUsedSize() *string {
	return s.UsedSize
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetAdminUserIds(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.AdminUserIds = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetAdminUserInfos(v []*DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.AdminUserInfos = v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetCreateTime(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetDirectoryId(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.DirectoryId = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetDriveId(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.DriveId = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetGroupId(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetGroupName(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetOrgId(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.OrgId = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetRecycleBinSize(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.RecycleBinSize = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetStatus(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.Status = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetTotalSize(v int64) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.TotalSize = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) SetUsedSize(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups {
	s.UsedSize = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroups) Validate() error {
	if s.AdminUserInfos != nil {
		for _, item := range s.AdminUserInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos struct {
	// The email address.
	//
	// example:
	//
	// 123@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The administrator username. This value may not be readable when imported from a third party.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The employee ID (DingTalk).
	//
	// example:
	//
	// 12345
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// The administrator nickname.
	//
	// example:
	//
	// John
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 12345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The administrator nickname.
	//
	// example:
	//
	// John
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks.
	//
	// example:
	//
	// John
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetEmail() *string {
	return s.Email
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetJobNumber() *string {
	return s.JobNumber
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetNickName() *string {
	return s.NickName
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetPhone() *string {
	return s.Phone
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetRealNickName() *string {
	return s.RealNickName
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetEmail(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.Email = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetEndUserId(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.EndUserId = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetJobNumber(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.JobNumber = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetNickName(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.NickName = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetPhone(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.Phone = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetRealNickName(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.RealNickName = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) SetRemark(v string) *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos {
	s.Remark = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponseBodyCloudDriveGroupsAdminUserInfos) Validate() error {
	return dara.Validate(s)
}

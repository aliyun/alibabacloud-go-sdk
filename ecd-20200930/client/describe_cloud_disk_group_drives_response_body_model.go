// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskGroupDrivesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudDriveGroups(v []*DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) *DescribeCloudDiskGroupDrivesResponseBody
	GetCloudDriveGroups() []*DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups
	SetCount(v int64) *DescribeCloudDiskGroupDrivesResponseBody
	GetCount() *int64
	SetNextToken(v string) *DescribeCloudDiskGroupDrivesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCloudDiskGroupDrivesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudDiskGroupDrivesResponseBody
	GetSuccess() *bool
}

type DescribeCloudDiskGroupDrivesResponseBody struct {
	CloudDriveGroups []*DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups `json:"CloudDriveGroups,omitempty" xml:"CloudDriveGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// MTA0MjA=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// B9F9CBBE-8A9F-5FE5-8A72-0E81C2401A91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudDiskGroupDrivesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupDrivesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) GetCloudDriveGroups() []*DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	return s.CloudDriveGroups
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) SetCloudDriveGroups(v []*DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) *DescribeCloudDiskGroupDrivesResponseBody {
	s.CloudDriveGroups = v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) SetCount(v int64) *DescribeCloudDiskGroupDrivesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) SetNextToken(v string) *DescribeCloudDiskGroupDrivesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) SetRequestId(v string) *DescribeCloudDiskGroupDrivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) SetSuccess(v bool) *DescribeCloudDiskGroupDrivesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBody) Validate() error {
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

type DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups struct {
	// example:
	//
	// 2022-04-11T07:44:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-959593****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 1234
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// example:
	//
	// cg-e70ga4ixp30ur****
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// org-aliyun-wy-org-id
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 5368709120
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// example:
	//
	// 1024000000
	UsedSize *string `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
}

func (s DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetDriveId() *string {
	return s.DriveId
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) GetUsedSize() *string {
	return s.UsedSize
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetCreateTime(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetDirectoryId(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.DirectoryId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetDriveId(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.DriveId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetGroupId(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetGroupName(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetOrgId(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.OrgId = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetStatus(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.Status = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetTotalSize(v int64) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.TotalSize = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) SetUsedSize(v string) *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups {
	s.UsedSize = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponseBodyCloudDriveGroups) Validate() error {
	return dara.Validate(s)
}
